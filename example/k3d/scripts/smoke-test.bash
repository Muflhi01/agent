#!/usr/bin/env bash
#
# Usage:
#   smoke-test.bash [-i] [-d] [-s] [-t <duration>]
#
# Dependencies:
#   k3d >=3.0
#   Tanka
#   jq
#
# smoke-test.bash performs smoke tests that can be used to validate a release.
# It only validates correctness, and does not attempt to do significant load
# testing or performance benchmarks.
#
# It works by deploying a k3d cluster with two pairs of Agent deployments: one
# using the scraping service, and one using host filtering. Each Agent deployment
# monitors a unique set of correctness tools (Loki Canary, Tempo Vulture,
# Cortex text-exporter). These tools expose correctness metrics, generated by
# by querying the backends that the Grafana Agents are configured to send
# telemetry data to.
#
# Grafana and Prometheus are deployed to the cluster and are responsible for
# internal monitoring. Prometheus is configured with a set of alert rules that serve
# as test cases for the smoke tests; alerts generated within the span of the testing
# period are treated as failures.
#
# After the smoke-test duration period (defaulting to 3h), alerts will
# be checked, and the script will end.
#
# The k3d cluster is kept alive after the test for analysis. To clean up assets created
# by the script, re-run the script with the -d flag.

set -euo pipefail

# Constants
ROOT=$(git rev-parse --show-toplevel)
K3D_CLUSTER_NAME="agent-smoke-test"
MUTATION_FREQUENCY="5m"
CHAOS_FREQUENCY="30m"

# Variables

# Which function will be called
ENTRYPOINT="run"
TEST_DURATION="3h"
IMPORT_IMAGES=""

while getopts "dt:ish" opt; do
  case $opt in
    d) ENTRYPOINT="cleanup" ;;
    t) TEST_DURATION=$OPTARG ;;
    i) IMPORT_IMAGES="yes" ;;
    s) SKIP_CREATE="yes" ;;
    h)
      echo "Usage: $0 [-i] [-d] [-s] [-t <duration>]"
      exit 0
      ;;
    *)
      echo "Usage: $0 [-i] [-d] [-s] [-t <duration>]"
      exit 1
      ;;
  esac
done

# Run runs the smoke test for $TEST_DURATION.
run() {
  if [[ -z "$SKIP_CREATE" ]]; then
    echo "--- Creating k3d cluster $K3D_CLUSTER_NAME"
    k3d cluster create $K3D_CLUSTER_NAME \
      --port 50080:80@loadbalancer \
      --api-port 50443 \
      --kubeconfig-update-default=true \
      --kubeconfig-switch-context=true \
      --wait >/dev/null
  fi

  # Give the cluster a little bit of time to settle before
  # applying the environment
  echo "--- Waiting for cluster to warm up"
  sleep 10

  if [[ ! -z "$IMPORT_IMAGES" ]]; then
    echo "--- Importing local images"

    k3d image import -c $K3D_CLUSTER_NAME \
      grafana/agent:main \
      grafana/agentctl:main \
      grafana/agent-crow:main
  fi

  (cd $ROOT/example/k3d && jb install)
  tk apply $ROOT/example/k3d/smoke --dangerous-auto-approve

  echo "--- Spawning background tasks"
  mutation_loop &
  MUTATION_PID=$!
  chaos_loop &
  CHAOS_PID=$!

  # Immediately create a job to sync configs so our two Agent deployments
  # are synced up as closely as possible.
  kubectl --context=k3d-$K3D_CLUSTER_NAME --namespace=smoke \
    create job --from=cronjob/grafana-agent-syncer \
    grafana-agent-syncer-startup

  echo "Your environment is now running for the next $TEST_DURATION."
  echo "Grafana URL: http://grafana.k3d.localhost:50080"
  echo "Prometheus URL: http://prometheus.k3d.localhost:50080"
  sleep $TEST_DURATION

  echo "--- Stopping background tasks"
  kill $CHAOS_PID $MUTATION_PID

  echo "Smoke tests complete!"
  echo "Grafana URL: http://grafana.k3d.localhost:50080"
  echo "Prometheus URL: http://prometheus.k3d.localhost:50080"
  echo ""
  echo "Getting results..."

  get_results
}

# mutation_loop changes the number of replicas of Avalanche servers
# randomly between 0-10 replicas (inclusive) every $MUTATION_FREQUENCY.
mutation_loop() {
  while true; do
    sleep $MUTATION_FREQUENCY

    NEW_REPLICAS=$(($RANDOM % 11))
    echo "--- Scaling Avalanche to $NEW_REPLICAS replicas"
    kubectl --context=k3d-$K3D_CLUSTER_NAME --namespace=smoke \
      scale --replicas=$NEW_REPLICAS deployment/avalanche
  done
}

# chaos loop deletes pods every $CHAOS_FREQUENCY.
chaos_loop() {
  while true; do
    sleep $CHAOS_FREQUENCY

    echo "--- Force deleting pod/grafana-agent-0"
    kubectl --context=k3d-$K3D_CLUSTER_NAME --namespace=smoke \
      delete --grace-period=0 --force pod/grafana-agent-0

    REPLICA=$(($RANDOM % 3))
    echo "--- Force deleting pod/agent-cluster-$REPLICA"
    kubectl --context=k3d-$K3D_CLUSTER_NAME --namespace=smoke \
      delete --grace-period=0 --force pod/grafana-agent-cluster-$REPLICA
  done
}

get_results() {
  NUM_ALERTS=$(curl -s -G \
    -H "Host: prometheus.k3d.localhost"    \
    -d "query=count_over_time(ALERTS{alertstate=\"firing\"}[$TEST_DURATION])" \
    'http://localhost:50080/api/v1/query'  \
      | jq '.data.result | length'         \
  )

  if test $NUM_ALERTS -ne 0; then
    echo "FAIL: $NUM_ALERTS alerts found over the last $TEST_DURATION."
    echo "More information: http://prometheus.k3d.localhost:50080/graph?g0.expr=count_over_time(ALERTS{alertstate%3D%22firing%22}[$TEST_DURATION])"

    exit 1
  else
    echo "PASS: 0 alerts found over the last $TEST_DURATION. You're good to go!"

    exit 0
  fi
}

cleanup() {
  echo "--- Deleting k3d cluster $K3D_CLUSTER_NAME"
  k3d cluster delete $K3D_CLUSTER_NAME >/dev/null
}

$ENTRYPOINT

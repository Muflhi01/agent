package v2 //nolint:golint

// Generated by tool/codegen to add a new integration add to tools/codegen/registry, then run make generate

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/prometheus/common/model"

	"github.com/go-kit/log"
	"github.com/grafana/agent/pkg/integrations/shared"
	"github.com/grafana/agent/pkg/integrations/v2/common"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/yaml.v2"

	"github.com/grafana/agent/pkg/integrations/v1/cadvisor"
	"github.com/grafana/agent/pkg/integrations/v1/consul_exporter"
	"github.com/grafana/agent/pkg/integrations/v1/dnsmasq_exporter"
	"github.com/grafana/agent/pkg/integrations/v1/elasticsearch_exporter"
	"github.com/grafana/agent/pkg/integrations/v1/github_exporter"
	"github.com/grafana/agent/pkg/integrations/v1/kafka_exporter"
	"github.com/grafana/agent/pkg/integrations/v1/memcached_exporter"
	"github.com/grafana/agent/pkg/integrations/v1/mongodb_exporter"
	"github.com/grafana/agent/pkg/integrations/v1/mysqld_exporter"
	"github.com/grafana/agent/pkg/integrations/v1/node_exporter"
	"github.com/grafana/agent/pkg/integrations/v1/postgres_exporter"
	"github.com/grafana/agent/pkg/integrations/v1/process_exporter"
	"github.com/grafana/agent/pkg/integrations/v1/redis_exporter"
	"github.com/grafana/agent/pkg/integrations/v1/statsd_exporter"
	"github.com/grafana/agent/pkg/integrations/v1/windows_exporter"
	"github.com/grafana/agent/pkg/integrations/v2/agent"
)

// Integrations is a wrapper to hold all the integrations for marshalling
type Integrations struct {
	Agent                        *Agent                   `yaml:"agent,omitempty"`
	Cadvisor                     *Cadvisor                `yaml:"cadvisor,omitempty"`
	ConsulExporterConfigs        []*ConsulExporter        `yaml:"consul_exporter_configs,omitempty"`
	DnsmasqExporterConfigs       []*DnsmasqExporter       `yaml:"dnsmasq_exporter_configs,omitempty"`
	ElasticsearchExporterConfigs []*ElasticsearchExporter `yaml:"elasticsearch_exporter_configs,omitempty"`
	GithubExporterConfigs        []*GithubExporter        `yaml:"github_exporter_configs,omitempty"`
	KafkaExporterConfigs         []*KafkaExporter         `yaml:"kafka_exporter_configs,omitempty"`
	MemcachedExporterConfigs     []*MemcachedExporter     `yaml:"memcached_exporter_configs,omitempty"`
	MongodbExporterConfigs       []*MongodbExporter       `yaml:"mongodb_exporter_configs,omitempty"`
	MysqldExporterConfigs        []*MysqldExporter        `yaml:"mysqld_exporter_configs,omitempty"`
	NodeExporter                 *NodeExporter            `yaml:"node_exporter,omitempty"`
	PostgresExporterConfigs      []*PostgresExporter      `yaml:"postgres_exporter_configs,omitempty"`
	ProcessExporter              *ProcessExporter         `yaml:"process_exporter,omitempty"`
	RedisExporterConfigs         []*RedisExporter         `yaml:"redis_exporter_configs,omitempty"`
	StatsdExporter               *StatsdExporter          `yaml:"statsd_exporter,omitempty"`
	WindowsExporter              *WindowsExporter         `yaml:"windows_exporter,omitempty"`
	TestConfigs                  []Config                 `yaml:"-,omitempty"`
}

// ActiveConfigs returns an array of integrations that have been set
func (v *Integrations) ActiveConfigs() []Config {
	activeConfigs := make([]Config, 0)
	if v.Agent != nil {
		activeConfigs = append(activeConfigs, v.Agent)
	}
	if v.Cadvisor != nil {
		activeConfigs = append(activeConfigs, newConfigWrapper(v.Cadvisor, v.Cadvisor.Cmn, v.Cadvisor.NewIntegration, v.Cadvisor.InstanceKey))
	}
	for _, i := range v.ConsulExporterConfigs {
		activeConfigs = append(activeConfigs, newConfigWrapper(i, i.Cmn, i.NewIntegration, i.InstanceKey))
	}
	for _, i := range v.DnsmasqExporterConfigs {
		activeConfigs = append(activeConfigs, newConfigWrapper(i, i.Cmn, i.NewIntegration, i.InstanceKey))
	}
	for _, i := range v.ElasticsearchExporterConfigs {
		activeConfigs = append(activeConfigs, newConfigWrapper(i, i.Cmn, i.NewIntegration, i.InstanceKey))
	}
	for _, i := range v.GithubExporterConfigs {
		activeConfigs = append(activeConfigs, newConfigWrapper(i, i.Cmn, i.NewIntegration, i.InstanceKey))
	}
	for _, i := range v.KafkaExporterConfigs {
		activeConfigs = append(activeConfigs, newConfigWrapper(i, i.Cmn, i.NewIntegration, i.InstanceKey))
	}
	for _, i := range v.MemcachedExporterConfigs {
		activeConfigs = append(activeConfigs, newConfigWrapper(i, i.Cmn, i.NewIntegration, i.InstanceKey))
	}
	for _, i := range v.MongodbExporterConfigs {
		activeConfigs = append(activeConfigs, newConfigWrapper(i, i.Cmn, i.NewIntegration, i.InstanceKey))
	}
	for _, i := range v.MysqldExporterConfigs {
		activeConfigs = append(activeConfigs, newConfigWrapper(i, i.Cmn, i.NewIntegration, i.InstanceKey))
	}
	if v.NodeExporter != nil {
		activeConfigs = append(activeConfigs, newConfigWrapper(v.NodeExporter, v.NodeExporter.Cmn, v.NodeExporter.NewIntegration, v.NodeExporter.InstanceKey))
	}
	for _, i := range v.PostgresExporterConfigs {
		activeConfigs = append(activeConfigs, newConfigWrapper(i, i.Cmn, i.NewIntegration, i.InstanceKey))
	}
	if v.ProcessExporter != nil {
		activeConfigs = append(activeConfigs, newConfigWrapper(v.ProcessExporter, v.ProcessExporter.Cmn, v.ProcessExporter.NewIntegration, v.ProcessExporter.InstanceKey))
	}
	for _, i := range v.RedisExporterConfigs {
		activeConfigs = append(activeConfigs, newConfigWrapper(i, i.Cmn, i.NewIntegration, i.InstanceKey))
	}
	if v.StatsdExporter != nil {
		activeConfigs = append(activeConfigs, newConfigWrapper(v.StatsdExporter, v.StatsdExporter.Cmn, v.StatsdExporter.NewIntegration, v.StatsdExporter.InstanceKey))
	}
	if v.WindowsExporter != nil {
		activeConfigs = append(activeConfigs, newConfigWrapper(v.WindowsExporter, v.WindowsExporter.Cmn, v.WindowsExporter.NewIntegration, v.WindowsExporter.InstanceKey))
	}
	for _, i := range v.TestConfigs {
		activeConfigs = append(activeConfigs, i)
	}
	return activeConfigs
}

type Agent struct {
	agent.Config `yaml:",omitempty,inline"`
}

// UnmarshalYaml is required for unmarshalling
func (c *Agent) UnmarshalYAML(unmarshal func(interface{}) error) error {
	err := implementsMarshaller(c.Config)
	if err != nil {
		return err
	}
	type plain Agent
	return unmarshal((*plain)(c))
}
func (c *Agent) NewIntegration(l log.Logger, globals shared.Globals) (Integration, error) {
	return NewMetricsHandlerIntegration(l, c, c.Common, globals, promhttp.Handler())
}

type Cadvisor struct {
	cadvisor.Config `yaml:",omitempty,inline"`

	Cmn common.MetricsConfig `yaml:",inline"`
}

// UnmarshalYaml is required for unmarshalling
func (c *Cadvisor) UnmarshalYAML(unmarshal func(interface{}) error) error {
	c.Config = cadvisor.DefaultConfig
	err := implementsMarshaller(c.Config)
	if err != nil {
		return err
	}
	type plain Cadvisor
	return unmarshal((*plain)(c))
}

type ConsulExporter struct {
	consul_exporter.Config `yaml:",omitempty,inline"`

	Cmn common.MetricsConfig `yaml:",inline"`
}

// UnmarshalYaml is required for unmarshalling
func (c *ConsulExporter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	c.Config = consul_exporter.DefaultConfig
	err := implementsMarshaller(c.Config)
	if err != nil {
		return err
	}
	type plain ConsulExporter
	return unmarshal((*plain)(c))
}

type DnsmasqExporter struct {
	dnsmasq_exporter.Config `yaml:",omitempty,inline"`

	Cmn common.MetricsConfig `yaml:",inline"`
}

// UnmarshalYaml is required for unmarshalling
func (c *DnsmasqExporter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	c.Config = dnsmasq_exporter.DefaultConfig
	err := implementsMarshaller(c.Config)
	if err != nil {
		return err
	}
	type plain DnsmasqExporter
	return unmarshal((*plain)(c))
}

type ElasticsearchExporter struct {
	elasticsearch_exporter.Config `yaml:",omitempty,inline"`

	Cmn common.MetricsConfig `yaml:",inline"`
}

// UnmarshalYaml is required for unmarshalling
func (c *ElasticsearchExporter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	c.Config = elasticsearch_exporter.DefaultConfig
	err := implementsMarshaller(c.Config)
	if err != nil {
		return err
	}
	type plain ElasticsearchExporter
	return unmarshal((*plain)(c))
}

type GithubExporter struct {
	github_exporter.Config `yaml:",omitempty,inline"`

	Cmn common.MetricsConfig `yaml:",inline"`
}

// UnmarshalYaml is required for unmarshalling
func (c *GithubExporter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	c.Config = github_exporter.DefaultConfig
	err := implementsMarshaller(c.Config)
	if err != nil {
		return err
	}
	type plain GithubExporter
	return unmarshal((*plain)(c))
}

type KafkaExporter struct {
	kafka_exporter.Config `yaml:",omitempty,inline"`

	Cmn common.MetricsConfig `yaml:",inline"`
}

// UnmarshalYaml is required for unmarshalling
func (c *KafkaExporter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	c.Config = kafka_exporter.DefaultConfig
	err := implementsMarshaller(c.Config)
	if err != nil {
		return err
	}
	type plain KafkaExporter
	return unmarshal((*plain)(c))
}

type MemcachedExporter struct {
	memcached_exporter.Config `yaml:",omitempty,inline"`

	Cmn common.MetricsConfig `yaml:",inline"`
}

// UnmarshalYaml is required for unmarshalling
func (c *MemcachedExporter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	c.Config = memcached_exporter.DefaultConfig
	err := implementsMarshaller(c.Config)
	if err != nil {
		return err
	}
	type plain MemcachedExporter
	return unmarshal((*plain)(c))
}

type MongodbExporter struct {
	mongodb_exporter.Config `yaml:",omitempty,inline"`

	Cmn common.MetricsConfig `yaml:",inline"`
}

// UnmarshalYaml is required for unmarshalling
func (c *MongodbExporter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	err := implementsMarshaller(c.Config)
	if err != nil {
		return err
	}
	type plain MongodbExporter
	return unmarshal((*plain)(c))
}

type MysqldExporter struct {
	mysqld_exporter.Config `yaml:",omitempty,inline"`

	Cmn common.MetricsConfig `yaml:",inline"`
}

// UnmarshalYaml is required for unmarshalling
func (c *MysqldExporter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	c.Config = mysqld_exporter.DefaultConfig
	err := implementsMarshaller(c.Config)
	if err != nil {
		return err
	}
	type plain MysqldExporter
	return unmarshal((*plain)(c))
}

type NodeExporter struct {
	node_exporter.Config `yaml:",omitempty,inline"`

	Cmn common.MetricsConfig `yaml:",inline"`
}

// UnmarshalYaml is required for unmarshalling
func (c *NodeExporter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	c.Config = node_exporter.DefaultConfig
	err := implementsMarshaller(c.Config)
	if err != nil {
		return err
	}
	type plain NodeExporter
	return unmarshal((*plain)(c))
}

type PostgresExporter struct {
	postgres_exporter.Config `yaml:",omitempty,inline"`

	Cmn common.MetricsConfig `yaml:",inline"`
}

// UnmarshalYaml is required for unmarshalling
func (c *PostgresExporter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	err := implementsMarshaller(c.Config)
	if err != nil {
		return err
	}
	type plain PostgresExporter
	return unmarshal((*plain)(c))
}

type ProcessExporter struct {
	process_exporter.Config `yaml:",omitempty,inline"`

	Cmn common.MetricsConfig `yaml:",inline"`
}

// UnmarshalYaml is required for unmarshalling
func (c *ProcessExporter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	c.Config = process_exporter.DefaultConfig
	err := implementsMarshaller(c.Config)
	if err != nil {
		return err
	}
	type plain ProcessExporter
	return unmarshal((*plain)(c))
}

type RedisExporter struct {
	redis_exporter.Config `yaml:",omitempty,inline"`

	Cmn common.MetricsConfig `yaml:",inline"`
}

// UnmarshalYaml is required for unmarshalling
func (c *RedisExporter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	c.Config = redis_exporter.DefaultConfig
	err := implementsMarshaller(c.Config)
	if err != nil {
		return err
	}
	type plain RedisExporter
	return unmarshal((*plain)(c))
}

type StatsdExporter struct {
	statsd_exporter.Config `yaml:",omitempty,inline"`

	Cmn common.MetricsConfig `yaml:",inline"`
}

// UnmarshalYaml is required for unmarshalling
func (c *StatsdExporter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	c.Config = statsd_exporter.DefaultConfig
	err := implementsMarshaller(c.Config)
	if err != nil {
		return err
	}
	type plain StatsdExporter
	return unmarshal((*plain)(c))
}

type WindowsExporter struct {
	windows_exporter.Config `yaml:",omitempty,inline"`

	Cmn common.MetricsConfig `yaml:",inline"`
}

// UnmarshalYaml is required for unmarshalling
func (c *WindowsExporter) UnmarshalYAML(unmarshal func(interface{}) error) error {
	c.Config = windows_exporter.DefaultConfig
	err := implementsMarshaller(c.Config)
	if err != nil {
		return err
	}
	type plain WindowsExporter
	return unmarshal((*plain)(c))
}

type configWrapper struct {
	cfg                shared.Config
	cmn                common.MetricsConfig
	configInstanceFunc configInstance
	newInstanceFunc    newIntegration
}

func (c *configWrapper) ApplyDefaults(globals shared.Globals) error {
	c.cmn.ApplyDefaults(globals.Autoscrape)
	if id, err := c.Identifier(globals); err == nil {
		c.cmn.InstanceKey = &id
	}
	return nil
}

func (c *configWrapper) Identifier(globals shared.Globals) (string, error) {
	if c.cmn.InstanceKey != nil {
		return *c.cmn.InstanceKey, nil
	}
	return c.configInstanceFunc(globals.AgentIdentifier)
}

func (c *configWrapper) NewIntegration(logger log.Logger, globals shared.Globals) (Integration, error) {
	return newIntegrationFromV1(c, logger, globals, c.newInstanceFunc)
}

func (c *configWrapper) Cfg() Config {
	return c
}

func (c *configWrapper) Name() string {
	return c.cfg.Name()
}

func (c *configWrapper) Common() common.MetricsConfig {
	return c.cmn
}

type newIntegration func(l log.Logger) (shared.Integration, error)

type configInstance func(agentKey string) (string, error)

func newConfigWrapper(cfg shared.Config, cmn common.MetricsConfig, ni newIntegration, ci configInstance) *configWrapper {
	return &configWrapper{
		cfg:                cfg,
		cmn:                cmn,
		configInstanceFunc: ci,
		newInstanceFunc:    ni,
	}
}

func newIntegrationFromV1(c IntegrationConfig, logger log.Logger, globals shared.Globals, newInt func(l log.Logger) (shared.Integration, error)) (Integration, error) {

	v1Integration, err := newInt(logger)
	if err != nil {
		return nil, err
	}

	id, err := c.Cfg().Identifier(globals)
	if err != nil {
		return nil, err
	}

	// Generate our handler. Original integrations didn't accept a prefix, and
	// just assumed that they would be wired to /metrics somewhere.
	handler, err := v1Integration.MetricsHandler()
	if err != nil {
		return nil, fmt.Errorf("generating http handler: %w", err)
	} else if handler == nil {
		handler = http.NotFoundHandler()
	}

	// Generate targets. Original integrations used a static set of targets,
	// so this mapping can always be generated just once.
	//
	// Targets are generated from the result of ScrapeConfigs(), which returns a
	// tuple of job name and relative metrics path.
	//
	// Job names were prefixed at the subsystem level with integrations/, so we
	// will retain that behavior here.
	v1ScrapeConfigs := v1Integration.ScrapeConfigs()
	targets := make([]handlerTarget, 0, len(v1ScrapeConfigs))
	for _, sc := range v1ScrapeConfigs {
		targets = append(targets, handlerTarget{
			MetricsPath: sc.MetricsPath,
			Labels: model.LabelSet{
				model.JobLabel: model.LabelValue("integrations/" + sc.JobName),
			},
		})
	}

	// Convert he run function. Original integrations sometimes returned
	// ctx.Err() on exit. This isn't recommended anymore, but we need to hide the
	// error if it happens, since the error was previously ignored.
	runFunc := func(ctx context.Context) error {
		err := v1Integration.Run(ctx)
		switch {
		case err == nil:
			return nil
		case errors.Is(err, context.Canceled) && ctx.Err() != nil:
			// Hide error that no longer happens in newer integrations.
			return nil
		default:
			return err
		}
	}

	// Aggregate our converted settings into a v2 integration.
	return &metricsHandlerIntegration{
		integrationName: c.Cfg().Name(),
		instanceID:      id,

		common:  c.Common(),
		globals: globals,
		handler: handler,
		targets: targets,

		runFunc: runFunc,
	}, nil
}

func implementsMarshaller(i interface{}) error {
	_, ok := i.(yaml.Unmarshaler)
	if ok {
		return errors.New(fmt.Sprintf("%T cannot implement custom unmarshaler", i))
	}
	return nil
}

// nolint
// code generated by sponge.

package config

import (
	"github.com/zhufuyi/sponge/pkg/conf"
)

var config *Config

func Init(configFile string, fs ...func()) error {
	config = &Config{}
	return conf.Parse(configFile, config, fs...)
}

func Show() string {
	return conf.Show(config)
}

func Get() *Config {
	if config == nil {
		panic("config is nil")
	}
	return config
}

func Set(conf *Config) {
	config = conf
}

type Config struct {
	App     App     `yaml:"app" json:"app"`
	Consul  Consul  `yaml:"consul" json:"consul"`
	Etcd    Etcd    `yaml:"etcd" json:"etcd"`
	Grpc    Grpc    `yaml:"grpc" json:"grpc"`
	HTTP    HTTP    `yaml:"http" json:"http"`
	Jaeger  Jaeger  `yaml:"jaeger" json:"jaeger"`
	Logger  Logger  `yaml:"logger" json:"logger"`
	Mysql   Mysql   `yaml:"mysql" json:"mysql"`
	NacosRd NacosRd `yaml:"nacosRd" json:"nacosRd"`
	Redis   Redis   `yaml:"redis" json:"redis"`
}

type Consul struct {
	Addr string `yaml:"addr" json:"addr"`
}

type Etcd struct {
	Addrs []string `yaml:"addrs" json:"addrs"`
}

type Jaeger struct {
	AgentHost string `yaml:"agentHost" json:"agentHost"`
	AgentPort int    `yaml:"agentPort" json:"agentPort"`
}

type Mysql struct {
	ConnMaxLifetime int    `yaml:"connMaxLifetime" json:"connMaxLifetime"`
	Dsn             string `yaml:"dsn" json:"dsn"`
	EnableLog       bool   `yaml:"enableLog" json:"enableLog"`
	MaxIdleConns    int    `yaml:"maxIdleConns" json:"maxIdleConns"`
	MaxOpenConns    int    `yaml:"maxOpenConns" json:"maxOpenConns"`
	SlowThreshold   int    `yaml:"slowThreshold" json:"slowThreshold"`
}

type Redis struct {
	DialTimeout  int    `yaml:"dialTimeout" json:"dialTimeout"`
	Dsn          string `yaml:"dsn" json:"dsn"`
	ReadTimeout  int    `yaml:"readTimeout" json:"readTimeout"`
	WriteTimeout int    `yaml:"writeTimeout" json:"writeTimeout"`
}

type App struct {
	CacheType               string  `yaml:"cacheType" json:"cacheType"`
	EnableCircuitBreaker    bool    `yaml:"enableCircuitBreaker" json:"enableCircuitBreaker"`
	EnableLimit             bool    `yaml:"enableLimit" json:"enableLimit"`
	EnableMetrics           bool    `yaml:"enableMetrics" json:"enableMetrics"`
	EnablePprof             bool    `yaml:"enablePprof" json:"enablePprof"`
	EnableRegistryDiscovery bool    `yaml:"enableRegistryDiscovery" json:"enableRegistryDiscovery"`
	EnableTracing           bool    `yaml:"enableTracing" json:"enableTracing"`
	Env                     string  `yaml:"env" json:"env"`
	Host                    string  `yaml:"host" json:"host"`
	Name                    string  `yaml:"name" json:"name"`
	RegistryDiscoveryType   string  `yaml:"registryDiscoveryType" json:"registryDiscoveryType"`
	TracingSamplingRate     float64 `yaml:"tracingSamplingRate" json:"tracingSamplingRate"`
	Version                 string  `yaml:"version" json:"version"`
}

type Logger struct {
	Format string `yaml:"format" json:"format"`
	IsSave bool   `yaml:"isSave" json:"isSave"`
	Level  string `yaml:"level" json:"level"`
}

type NacosRd struct {
	IPAddr      string `yaml:"ipAddr" json:"ipAddr"`
	NamespaceID string `yaml:"namespaceID" json:"namespaceID"`
	Port        int    `yaml:"port" json:"port"`
}

type Grpc struct {
	MetricsPort  int `yaml:"metricsPort" json:"metricsPort"`
	Port         int `yaml:"port" json:"port"`
	PprofPort    int `yaml:"pprofPort" json:"pprofPort"`
	ReadTimeout  int `yaml:"readTimeout" json:"readTimeout"`
	WriteTimeout int `yaml:"writeTimeout" json:"writeTimeout"`
}

type HTTP struct {
	Port         int `yaml:"port" json:"port"`
	ReadTimeout  int `yaml:"readTimeout" json:"readTimeout"`
	WriteTimeout int `yaml:"writeTimeout" json:"writeTimeout"`
}

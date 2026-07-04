package config

import (
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server       ServerConfig       `yaml:"server"`
	Storage      StorageConfig      `yaml:"storage"`
	RateLimit    RateLimitConfig    `yaml:"ratelimit"`
	Cache        CacheConfig        `yaml:"cache"`
	Auth         AuthConfig         `yaml:"auth"`
	LoadBalancer LoadBalancerConfig `yaml:"loadbalancer"`
	Observe      ObserveConfig      `yaml:"observe"`
	Docs         DocsConfig         `yaml:"docs"`
	Versioning   VersioningConfig   `yaml:"versioning"`
	Batch        BatchConfig        `yaml:"batch"`
}

type StorageConfig struct {
	Enabled bool   `yaml:"enabled"`
	Path    string `yaml:"path"`
	Seed    bool   `yaml:"seed"`
}

type ServerConfig struct {
	Host           string   `yaml:"host"`
	Port           int      `yaml:"port"`
	AllowedOrigins []string `yaml:"allowedOrigins"`
}

type RateLimitConfig struct {
	Enabled bool          `yaml:"enabled"`
	Default string        `yaml:"default"`
	Rules   []RateRule    `yaml:"rules"`
}

type RateRule struct {
	Path  string `yaml:"path"`
	Limit string `yaml:"limit"`
}

type CacheConfig struct {
	Enabled bool        `yaml:"enabled"`
	Driver  string      `yaml:"driver"`
	Redis   RedisConfig `yaml:"redis"`
	TTL     string      `yaml:"ttl"`
}

type RedisConfig struct {
	Addr string `yaml:"addr"`
}

type AuthConfig struct {
	Enabled   bool           `yaml:"enabled"`
	Providers []AuthProvider `yaml:"providers"`
}

type AuthProvider struct {
	Type   string `yaml:"type"`
	Secret string `yaml:"secret"`
	Issuer string `yaml:"issuer"`
	Header string `yaml:"header"`
}

type LoadBalancerConfig struct {
	Enabled  bool       `yaml:"enabled"`
	Strategy string     `yaml:"strategy"`
	Backends []Backend  `yaml:"backends"`
}

type Backend struct {
	URL    string `yaml:"url"`
	Weight int    `yaml:"weight"`
}

type ObserveConfig struct {
	Metrics MetricsConfig `yaml:"metrics"`
	Logs    LogsConfig    `yaml:"logs"`
	Tracing TracingConfig `yaml:"tracing"`
}

type MetricsConfig struct {
	Enabled bool   `yaml:"enabled"`
	Path    string `yaml:"path"`
}

type LogsConfig struct {
	Enabled bool   `yaml:"enabled"`
	Level   string `yaml:"level"`
	Format  string `yaml:"format"`
}

type TracingConfig struct {
	Enabled bool    `yaml:"enabled"`
	Sample  float64 `yaml:"sample"`
}

type DocsConfig struct {
	Enabled bool   `yaml:"enabled"`
	Path    string `yaml:"path"`
	Title   string `yaml:"title"`
	Version string `yaml:"version"`
}

type VersioningConfig struct {
	Strategy string `yaml:"strategy"`
	Header   string `yaml:"header"`
	Default  string `yaml:"default"`
}

type BatchConfig struct {
	Enabled       bool     `yaml:"enabled"`
	Window        string   `yaml:"window"`
	MaxBatchSize  int      `yaml:"maxBatchSize"`
	FlushInterval string   `yaml:"flushInterval"`
	Classifiers   []string `yaml:"classifiers"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Server: ServerConfig{
			Host:           "0.0.0.0",
			Port:           8080,
			AllowedOrigins: []string{"*"},
		},
		Storage: StorageConfig{
			Enabled: true,
			Path:    "./data/velo.db",
			Seed:    true,
		},
		RateLimit: RateLimitConfig{
			Enabled: true,
			Default: "100/min",
		},
		Cache: CacheConfig{
			Enabled: true,
			Driver:  "memory",
			TTL:     "5m",
		},
		Auth: AuthConfig{
			Enabled: false,
		},
		LoadBalancer: LoadBalancerConfig{
			Enabled:  false,
			Strategy: "round-robin",
		},
		Observe: ObserveConfig{
			Metrics: MetricsConfig{
				Enabled: true,
				Path:    "/metrics",
			},
			Logs: LogsConfig{
				Enabled: true,
				Level:   "info",
				Format:  "json",
			},
		},
		Docs: DocsConfig{
			Enabled: true,
			Path:    "/docs",
			Title:   "Velo API",
			Version: "1.0.0",
		},
		Versioning: VersioningConfig{
			Strategy: "header",
			Header:   "X-API-Version",
			Default:  "v1",
		},
		Batch: BatchConfig{
			Enabled:       true,
			Window:        "5ms",
			MaxBatchSize:  100,
			FlushInterval: "10ms",
			Classifiers:   []string{"endpoint"},
		},
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	if port := os.Getenv("PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			cfg.Server.Port = p
		}
	}

	return cfg, nil
}

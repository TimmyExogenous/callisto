package dogfood

import (
	"gopkg.in/yaml.v3"
)

type Config struct {
	ChainIDWithoutRevision string `yaml:"chain_id_without_revision"`
}

// NewConfig allows to build a new Config instance
func NewConfig(chainIDWithoutRevision string) *Config {
	return &Config{
		ChainIDWithoutRevision: chainIDWithoutRevision,
	}
}

func ParseConfig(bz []byte) (*Config, error) {
	type T struct {
		Config *Config `yaml:"dogfood"`
	}
	var cfg T
	err := yaml.Unmarshal(bz, &cfg)
	return cfg.Config, err
}

package bootstrap

import (
	"gopkg.in/yaml.v3"
)

type Config struct {
	ETHHttp       string `yaml:"eth_http"`
	ETHWebsocket  string `yaml:"eth_websocket"`
	BTCRPC        string `yaml:"btc_rpc"`
	XRPRPC        string `yaml:"xrp_rpc"`
	BootstrapAddr string `yaml:"bootstrap_addr"`
	// UpdateInterval specifies the interval in minutes at which the
	// bootstrap states are automatically refreshed.
	UpdateInterval int64  `yaml:"update_interval"`
	ETHLZChainID   uint64 `yaml:"eth_lz_chain_id"`
}

// NewConfig allows to build a new Config instance
func NewConfig(ethHttp, ethWebsocket, btcRPC, xrpRPC string) *Config {
	return &Config{
		ETHHttp:      ethHttp,
		ETHWebsocket: ethWebsocket,
		BTCRPC:       btcRPC,
		XRPRPC:       xrpRPC,
	}
}

func ParseConfig(bz []byte) (*Config, error) {
	type T struct {
		Config *Config `yaml:"bootstrap"`
	}
	var cfg T
	err := yaml.Unmarshal(bz, &cfg)
	return cfg.Config, err
}

package bootstrap

import (
	"fmt"

	callistotypes "github.com/forbole/callisto/v4/types"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ETHHttp       string `yaml:"eth_http"`
	ETHWebsocket  string `yaml:"eth_websocket"`
	BTCRPC        string `yaml:"btc_rpc"`
	XRPRPC        string `yaml:"xrp_rpc"`
	BootstrapAddr string `yaml:"bootstrap_addr"`
	// The following UpdateIntervals specify the interval in minutes at which the
	// bootstrap states are automatically refreshed.
	ETHUpdateInterval   int                                  `yaml:"eth_update_interval"`
	BTCUpdateInterval   int                                  `yaml:"btc_update_interval"`
	XRPUpdateInterval   int                                  `yaml:"xrp_update_interval"`
	PriceUpdateInterval int                                  `yaml:"price_update_interval"`
	ETHLZChainID        uint64                               `yaml:"eth_lz_chain_id"`
	ClientChainInfos    []callistotypes.BootstrapClientChain `yaml:"client_chain_infos"`
	StakingTokenInfos   []callistotypes.BootstrapToken       `yaml:"staking_token_infos"`
	TokenOracleFeeds    []callistotypes.OracleFeed           `yaml:"token_oracle_feeds"`
}

// NewConfig allows to build a new Config instance
func NewConfig(ethHTTP, ethWebsocket, btcRPC, xrpRPC string) *Config {
	return &Config{
		ETHHttp:      ethHTTP,
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
	if err != nil {
		return nil, fmt.Errorf("parse bootstrap config: %w", err)
	}
	if cfg.Config == nil {
		return nil, fmt.Errorf("bootstrap config missing")
	}
	return cfg.Config, nil
}

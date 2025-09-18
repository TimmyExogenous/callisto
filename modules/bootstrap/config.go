package bootstrap

import (
	"gopkg.in/yaml.v3"
)

const (
	// VirtualAddress is the unified virtual address used for all chain asset ID construction
	VirtualAddress = "0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
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
	// BTC vault address for bootstrap deposits
	BTCVaultAddr string `yaml:"btc_vault_addr"`
	// XRP vault address for bootstrap deposits
	XRPVaultAddr string `yaml:"xrp_vault_addr"`
	// Minimum confirmations required for BTC transactions
	BTCMinConfirmations int `yaml:"btc_min_confirmations"`
	// Minimum confirmations required for XRP transactions
	XRPMinConfirmations int `yaml:"xrp_min_confirmations"`
	// Minimum BTC amount in satoshis (0.001 BTC = 100000 satoshis)
	BTCMinAmount int64 `yaml:"btc_min_amount"`
	// Minimum XRP amount in drops (50 XRP = 50000000 drops)
	XRPMinAmount int64 `yaml:"xrp_min_amount"`
	// XRP destination tag for bootstrap deposits
	XRPDestinationTag int64 `yaml:"xrp_destination_tag"`
	// Maximum iterations for transaction fetching to prevent infinite loops
	MaxFetchIterations int `yaml:"max_fetch_iterations"`
	// Incremental scanning configuration
	BTCStartHeight int64 `yaml:"btc_start_height"` // BTC scanning start height
	XRPStartLedger int64 `yaml:"xrp_start_ledger"` // XRP scanning start ledger
	ScanBatchSize  int   `yaml:"scan_batch_size"`  // Batch size for scanning
	ReorgDepth     int   `yaml:"reorg_depth"`      // Reorganization detection depth
}

// NewConfig allows to build a new Config instance
func NewConfig(ethHttp, ethWebsocket, btcRPC, xrpRPC string) *Config {
	return &Config{
		ETHHttp:             ethHttp,
		ETHWebsocket:        ethWebsocket,
		BTCRPC:              btcRPC,
		XRPRPC:              xrpRPC,
		BTCMinConfirmations: 6,        // Default 6 confirmations for BTC
		XRPMinConfirmations: 6,        // Default 6 confirmations for XRP
		BTCMinAmount:        100000,   // Default 0.001 BTC in satoshis
		XRPMinAmount:        50000000, // Default 50 XRP in drops
		XRPDestinationTag:   9999,     // Default XRP destination tag
		MaxFetchIterations:  1000,     // Default max iterations to prevent infinite loops
		BTCStartHeight:      0,        // Default: scan from genesis
		XRPStartLedger:      0,        // Default: scan from genesis
		ScanBatchSize:       100,      // Default batch size
		ReorgDepth:          6,        // Default reorg detection depth
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

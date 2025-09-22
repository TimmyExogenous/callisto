package types

import "time"

type BootstrapValidator struct {
	ValidatorEthAddress string
	ValidatorIMAddress  string
	ValidatorName       string
	ConsensusPubKey     string
	Rate                string
	MaxRate             string
	MaxChangeRate       string
	UpdatedAt           time.Time
}

type BootstrapClientChain struct {
	Name      string `yaml:"name"`
	MetaInfo  string `yaml:"meta_info"`
	LZChainID uint64 `yaml:"lz_chain_id"`
}

type BootstrapToken struct {
	AssetID   string `yaml:"asset_id"`
	Name      string `yaml:"name"`
	Symbol    string `yaml:"symbol"`
	Address   string `yaml:"address"`
	Decimals  uint8  `yaml:"decimals"`
	LZChainID uint64 `yaml:"lz_chain_id"`
}
type BootstrapTokenState struct {
	BootstrapToken
	StakingTotalAmount string
	UpdatedAt          time.Time
}

type OracleFeed struct {
	AssetID    string `yaml:"asset_id"`
	OracleAddr string `yaml:"oracle_addr"` // contract address of Chainlink aggregator
}

type BootstrapStakerAsset struct {
	StakerID     string
	AssetID      string
	Deposited    string
	Withdrawable string
	Delegated    string
	UpdatedAt    time.Time
}

type BootstrapDelegationState struct {
	StakerID     string
	AssetID      string
	OperatorAddr string
	Delegated    string
	UpdatedAt    time.Time
}

type BootstrapOperatorAsset struct {
	OperatorAddr string
	AssetID      string
	TotalAmount  string
	SelfAmount   string
	OtherAmount  string
	UpdatedAt    time.Time
}

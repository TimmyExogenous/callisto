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
	Name             string
	MetaInfo         string
	LayerZeroChainID int64
	UpdatedAt        time.Time
}

type BootstrapToken struct {
	AssetID            string
	Name               string
	Symbol             string
	Address            string
	Decimals           uint8
	LayerZeroChainID   uint64
	StakingTotalAmount string
	UpdatedAt          time.Time
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

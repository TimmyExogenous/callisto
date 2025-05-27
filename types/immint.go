package types

import (
	imminttypes "github.com/imua-xyz/imuachain/x/immint/types"
)

// ImmintParams represents the x/immint parameters
type ImmintParams struct {
	imminttypes.Params
	Height int64
}

// NewImmintParams allows to build a new ImmintParams instance
func NewImmintParams(params imminttypes.Params, height int64) *ImmintParams {
	return &ImmintParams{
		Params: params,
		Height: height,
	}
}

// MintHistory represents the mint history
type MintHistory struct {
	// this must be a string and not sdkmath.Int because a string can be directly
	// inserted into the database, while an sdkmath.Int cannot.
	Amount      string
	Height      int64
	EpochID     string
	EpochNumber int64
	Denom       string
}

// NewMintHistory allows to build a new MintHistory instance
func NewMintHistory(
	height int64, amount string, epochID string, epochNumber int64, denom string,
) *MintHistory {
	return &MintHistory{
		Height:      height,
		Amount:      amount,
		EpochID:     epochID,
		EpochNumber: epochNumber,
		Denom:       denom,
	}
}

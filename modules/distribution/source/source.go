package source

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/imua-xyz/imuachain/x/feedistribution/types"
)

type Source interface {
	AVSCommunityPool(height int64, avsAddr string) (sdk.DecCoins, error)
	Params(height int64) (distrtypes.Params, error)
	StakerAllClaimedRewards(height int64, stakerID string) ([]distrtypes.StakerClaimedRewardsPerAVS, error)
	StakerUnclaimedRewards(height int64, stakerID string) (distrtypes.CommonAVSRewards, error)
}

package types

import "time"

type TokenomicsParams struct {
	GenesisSupply                     *string
	GenesisPoolRatio                  *string
	GenesisPoolAirdropDuration        *int64
	GenesisPoolAirdropInterval        *int64
	LiquidityIncentiveRatios          []string
	LiquidityIncentiveAirdropDuration *int64
	LiquidityIncentiveAirdropInterval *int64
	GenesisValidatorRewardRatio       *string
	CreatedAt                         *time.Time
}

// GenesisPoolAirdropRound represents a single round of genesis pool airdrop distribution.
// Each round includes the snapshot block height, total value, reward allocation, and completion status.
type GenesisPoolAirdropRound struct {
	AirdropRound       int       // Index of the airdrop round (primary key)
	BlockHeight        int64     // Snapshot block height for this round
	TotalStakers       int       // Total number of stakers eligible in this round
	TotalUSDValue      string    // Total USD value of all eligible stakers (NUMERIC in DB)
	TotalRewardAmount  string    // Total reward amount allocated for this round (NUMERIC in DB)
	RoundDuration      int64     // Round duration for this round
	CreatedAt          time.Time // Timestamp when this record was created (default: now())
	DistributedStakers int       // Number of stakers who have received their rewards
	IsCompleted        bool      // Indicates whether the reward distribution is finished
}

// GenesisStakerAirdrop represents the reward state for a single staker in a specific airdrop round.
type GenesisStakerAirdrop struct {
	StakerID      string     // Unique identifier of the staker
	AirdropRound  int        // Airdrop round index
	USDValue      string     // Total USD value of the staker's assets at snapshot time (NUMERIC)
	RewardAmount  string     // Allocated reward amount (NUMERIC)
	IsDistributed bool       // Whether the reward has been distributed
	DistributedAt *time.Time // When the reward was distributed (nullable)
}

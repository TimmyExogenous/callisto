package distribution

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	distrtypes "github.com/imua-xyz/imuachain/x/feedistribution/types"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	juno "github.com/forbole/juno/v5/types"
	"github.com/rs/zerolog/log"
)

// HandleBlock implements BlockModule
func (m *Module) HandleBlock(
	block *tmctypes.ResultBlock, res *tmctypes.ResultBlockResults, _ []*juno.Tx, _ *tmctypes.ResultValidators,
) error {
	log.Debug().Str("module", m.Name()).Int64("height", block.Block.Height).
		Msg(fmt.Sprintf("updating %s", m.Name()))
	// the events about allocating rewards to operator are emitted during EndBlock
	if err := m.handleAllocateRewardsToOperator(res.EndBlockEvents); err != nil {
		return fmt.Errorf("error while handling events for allocating rewards to operator: %s", err)
	}
	// handleUpdateStakerRewards
	if err := m.handleUpdateStakerRewards(block); err != nil {
		return fmt.Errorf("error while updating staker rewards: %s", err)
	}
	return nil
}

// handleAllocateRewardsToOperator filters, parses, and stores events emitted
// when rewards are allocated to an operator.
func (m *Module) handleAllocateRewardsToOperator(events []abci.Event) error {
	events = juno.FindEventsByType(events, distrtypes.EventTypeAllocateRewardsToOperator)

	for _, event := range events {
		// Extract attributes
		avsAddrAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyAvsAddress)
		if err != nil {
			return fmt.Errorf("failed to get AVS address: %w", err)
		}

		operatorAddrAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyOperator)
		if err != nil {
			return fmt.Errorf("failed to get operator address: %w", err)
		}

		totalRewardAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyOperatorTotalReward)
		if err != nil {
			return fmt.Errorf("failed to get operator total reward: %w", err)
		}

		commissionAttr, err := juno.FindAttributeByKey(event, distrtypes.AttributeKeyOperatorCommission)
		if err != nil {
			return fmt.Errorf("failed to get operator commission: %w", err)
		}

		// Parse rewards and commission
		totalReward, err := sdk.ParseDecCoins(totalRewardAttr.Value)
		if err != nil {
			return fmt.Errorf("failed to parse operator total reward: %w", err)
		}

		commission, err := sdk.ParseDecCoins(commissionAttr.Value)
		if err != nil {
			return fmt.Errorf("failed to parse operator commission: %w", err)
		}

		// Upsert into operator_rewards table
		err = m.db.UpsertOperatorRewards(
			operatorAddrAttr.Value,
			avsAddrAttr.Value,
			totalReward,
			commission,
		)
		if err != nil {
			return fmt.Errorf("failed to upsert operator rewards for %s/%s: %w",
				operatorAddrAttr.Value, avsAddrAttr.Value, err)
		}
	}

	return nil
}

// handleUpdateStakerRewards periodically updates rewards for all stakers by fetching
// reward states from the RPC. We intentionally avoid emitting too many events during
// reward distribution to delegations, so we cannot rely on events for indexing.
// Moreover, due to the F1 distribution model, unclaimed rewards can only be obtained via RPC.
//
// TODO: Consider the potential performance concern—periodically processing all stakers
// may put pressure on both the indexer and the source full node. This only
// affects the indexer itself and the node it fetches data from.
// Additionally, since the rewards state is updated periodically, it may not be real-time.
// The accuracy depends on the update frequency, but it should be sufficient for calculating
// long-term liquidity incentives over a 10-year period.
func (m *Module) handleUpdateStakerRewards(block *tmctypes.ResultBlock) error {
	// TODO: Skip update if the current block is too far behind the latest block on the IMUA chain.
	// Updating from RPC would be meaningless in this case.

	// get distribution indexer parameters
	params, err := m.db.GetDistributionIndexerParams()
	if err != nil {
		return err
	}

	if params.StakerRewardsUpdateInterval == nil {
		return fmt.Errorf("the staker rewards update interval is null")
	}

	oneDay := int64(24 * time.Hour)
	updateInterval := *params.StakerRewardsUpdateInterval
	if params.LastStakerRewardUpdateTime != nil {
		lastUpdateTime := *params.LastStakerRewardUpdateTime
		nextUpdateTime := lastUpdateTime.Add(time.Duration(oneDay * updateInterval))
		if nextUpdateTime.After(block.Block.Time) {
			// do nothing: too early.
			return nil
		}
	}

	// update the rewards for all stakers
	// get all stakers
	stakers, err := m.db.GetAllStakersFromDelegationStates()
	if err != nil {
		return err
	}
	for _, stakerID := range stakers {
		// get all claimed rewards
		allClaimedRewards, err := m.source.StakerAllClaimedRewards(block.Block.Height, stakerID)
		if err != nil {
			return fmt.Errorf("failed to get all outstanding rewards for staker:%s,err:%w", stakerID, err)
		}
		// get all unclaimed rewards
		allUnclaimedRewards, err := m.source.StakerUnclaimedRewards(block.Block.Height, stakerID)
		if err != nil {
			return fmt.Errorf("failed to get all unclaimed rewards for staker:%s,err:%w", stakerID, err)
		}

		// Collect all unique AVS addresses from both maps
		avsSet := make(map[string]struct{})
		// Build a map of claimed rewards by AVS address
		claimedRewardsMap := make(map[string]distrtypes.StakerClaimedRewards)
		for _, r := range allClaimedRewards {
			claimedRewardsMap[r.AVSAddress] = r.ClaimedRewards
			avsSet[r.AVSAddress] = struct{}{}
		}

		// Build a map of unclaimed rewards by AVS address
		unclaimedMap := make(map[string]sdk.DecCoins)
		for _, r := range allUnclaimedRewards {
			unclaimedMap[r.AVSAddress] = r.Rewards
			avsSet[r.AVSAddress] = struct{}{}
		}

		// Upsert reward data for each AVS address
		for avsAddr := range avsSet {
			claimedRewards, ok := claimedRewardsMap[avsAddr]
			if !ok {
				claimedRewards = distrtypes.StakerClaimedRewards{
					OutstandingRewards: sdk.NewDecCoins(),
					WithdrawnRewards:   sdk.NewDecCoins(),
				}
			}
			unclaimed, ok := unclaimedMap[avsAddr]
			if !ok {
				unclaimed = sdk.NewDecCoins()
			}

			if err := m.db.UpsertStakerRewards(
				stakerID, avsAddr,
				claimedRewards.OutstandingRewards,
				claimedRewards.WithdrawnRewards,
				unclaimed); err != nil {
				return fmt.Errorf("failed to upsert staker rewards for staker: %s, avs: %s, err: %w", stakerID, avsAddr, err)
			}
		}
	}
	return nil
}

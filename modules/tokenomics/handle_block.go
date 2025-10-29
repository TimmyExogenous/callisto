package tokenomics

import (
	sdkmath "cosmossdk.io/math"
	"fmt"
	"github.com/forbole/callisto/v4/types"
	operatorkeeper "github.com/imua-xyz/imuachain/x/operator/keeper"
	oracletypes "github.com/imua-xyz/imuachain/x/oracle/types"
	"time"

	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	juno "github.com/forbole/juno/v5/types"
	"github.com/rs/zerolog/log"
)

// HandleBlock implements BlockModule
func (m *Module) HandleBlock(
	block *tmctypes.ResultBlock, _ *tmctypes.ResultBlockResults, _ []*juno.Tx, _ *tmctypes.ResultValidators,
) error {
	log.Debug().Str("module", m.Name()).Int64("height", block.Block.Height).
		Msg(fmt.Sprintf("updating %s", m.Name()))

	return m.HandleGenesisPoolAirdropRound(block)
}

func (m *Module) HandleGenesisPoolAirdropForStakers(roundReward sdkmath.Int, round *types.CommonAirdropRound) error {
	var stakerAirdrop types.GenesisStakerAirdrop
	totalUSDValue := sdkmath.LegacyZeroDec()
	stakerUSDValue := sdkmath.LegacyZeroDec()
	assetPrices := make(map[string]*oracletypes.Price)
	assetDecimals := make(map[string]int)
	totalStakers := 0
	// iterate over all genesis stakers and calculate the USD value
	// the iteration is ordered by staker ID and asset ID, allowing the total USD value of
	// each staker to be calculated in a single pass.
	opFunc := func(sa types.ParsedStakerAsset) error {
		var err error
		if stakerAirdrop.StakerID == "" {
			stakerAirdrop.StakerID = sa.StakerID
		} else if stakerAirdrop.StakerID != sa.StakerID {
			// the airdrop for the previous staker has been fully processed; save it to the database.
			stakerAirdrop.AirdropRound = round.AirdropRound
			stakerAirdrop.USDValue = stakerUSDValue.String()
			err = m.db.SaveGenesisStakerAirdrop(&stakerAirdrop)
			if err != nil {
				return fmt.Errorf("error saving genesis staker airdrop: %w", err)
			}
			// update the total USD value and total number of stakers
			totalUSDValue.AddMut(stakerUSDValue)
			totalStakers++
			// clear the stakerUSDValue and airdrop info for next staker
			stakerUSDValue = sdkmath.LegacyZeroDec()
			stakerAirdrop = types.GenesisStakerAirdrop{
				StakerID: sa.StakerID,
			}
		}

		// get price info for the asset
		var price *oracletypes.Price
		var assetDecimal int
		var ok bool
		price, ok = assetPrices[sa.AssetID]
		if !ok {
			price, err = m.db.GetLatestPriceByAssetID(sa.AssetID)
			if err != nil {
				return fmt.Errorf("error getting latest price by assetID: %w,assetID:%s", err, sa.AssetID)
			}
			assetPrices[sa.AssetID] = price
		}
		// get asset decimal
		assetDecimal, ok = assetDecimals[sa.AssetID]
		if !ok {
			assetDecimal, err = m.db.GetTokenDecimalsByID(sa.AssetID)
			if err != nil {
				return fmt.Errorf("error getting asset decimal: %w,assetID:%s", err, sa.AssetID)
			}
			assetDecimals[sa.AssetID] = assetDecimal
		}
		// calculate the USD value for the stakerID and assetID
		validAssetAmount := sdkmath.MinInt(sa.GenesisDeposited, sa.Deposited)
		assetUSDValue := operatorkeeper.CalculateUSDValue(validAssetAmount, price.Value, uint32(assetDecimal), price.Decimal)
		stakerUSDValue.AddMut(assetUSDValue)

		return nil
	}
	err := m.db.IterateGenesisStakerAssets(opFunc)
	if err != nil {
		return fmt.Errorf("HandleGenesisPoolAirdropForStakers: error iterating over genesis stakers: %w", err)
	}
	// update the total USD value and staker number in the input round info.
	round.TotalUSDValue = totalUSDValue.String()
	round.TotalStakers = totalStakers

	// iterate over all stakers' airdrop info to calculate and update the rewards.
	err = m.db.UpdateGenesisAirdropRewardsByRound(round.AirdropRound, func(usdValue sdkmath.LegacyDec) (sdkmath.Int, error) {
		return usdValue.MulInt(roundReward).Quo(totalUSDValue).TruncateInt(), nil
	})
	if err != nil {
		return fmt.Errorf("HandleGenesisPoolAirdropForStakers: error updating airdrop reward for all genesis stakers: %w", err)
	}
	return nil
}
func (m *Module) calculateAirdropRound(block *tmctypes.ResultBlock, airdropType types.AirdropType) (*types.GeneralRoundInfo, error) {
	oneDay := int64(24 * time.Hour)
	oneYear := 365 * oneDay
	var interval, duration int64
	var latestAirdropRound func() (*types.CommonAirdropRound, error)
	switch airdropType {
	case types.GenesisPoolAirdrop:
		interval = m.cfg.GenesisPoolAirdropInterval * oneDay
		duration = m.cfg.GenesisPoolAirdropDuration * oneDay
		latestAirdropRound = m.db.GetLatestGenesisPoolAirdropRound
	case types.LiquidityIncentivesAirdrop:
		interval = m.cfg.LiquidityIncentiveAirdropInterval.MulInt64(oneYear).TruncateInt64()
		duration = m.cfg.LiquidityIncentiveAirdropDuration * oneYear
		latestAirdropRound = m.db.GetLatestLiquidityIncentivesAirdropRound
	default:
		return nil, fmt.Errorf("invalid airdrop type:%d", airdropType)
	}

	if interval == 0 || duration == 0 {
		// do nothing when the interval or duration is zero.
		log.Info().Msg("the interval or duration is zero")
		return nil, nil
	}

	latestAirdropRoundInfo, err := latestAirdropRound()
	if err != nil {
		return nil, err
	}

	var roundStart time.Time
	var latestRoundID int
	// get the genesis time
	genesis, err := m.db.GetGenesis()
	if err != nil {
		return nil, fmt.Errorf("error getting genesis: %w", err)
	}
	blockTime := block.Block.Time
	if latestAirdropRoundInfo == nil {
		roundStart = genesis.Time
	} else {
		roundStart = latestAirdropRoundInfo.CreatedAt
		latestRoundID = latestAirdropRoundInfo.AirdropRound
	}

	// calculate the round id by the genesis and block time.
	durFromStart := int64(blockTime.Sub(roundStart))
	airdropEndTime := genesis.Time.Add(time.Duration(duration))
	if (durFromStart < interval && blockTime.Before(airdropEndTime)) ||
		!roundStart.Before(airdropEndTime) {
		// Do nothing because the airdrop round is not due yet or the airdrop has already ended.
		return nil, nil
	}
	// the roundID will start from 1.
	roundID := latestRoundID + 1
	roundDur := durFromStart

	return &types.GeneralRoundInfo{
		RoundID:         roundID,
		RoundDur:        roundDur,
		RoundStart:      roundStart,
		GenesisTime:     genesis.Time,
		PreRound:        latestAirdropRoundInfo,
		AirdropDur:      duration,
		AirdropInterval: interval,
	}, nil
}

func (m *Module) HandleGenesisPoolAirdropRound(block *tmctypes.ResultBlock) error {
	if m.cfg.GenesisSupply == 0 || m.cfg.GenesisPoolRatio.IsZero() {
		// do nothing when the genesis supply or the genesis pool ratio is zero.
		log.Info().Msg("the genesis supply or the genesis pool ratio is zero")
		return nil
	}

	generalRoundInfo, err := m.calculateAirdropRound(block, types.GenesisPoolAirdrop)
	if err != nil {
		return err
	}

	if generalRoundInfo == nil {
		return nil
	}

	round := &types.CommonAirdropRound{
		AirdropRound:  generalRoundInfo.RoundID,
		BlockHeight:   block.Block.Height,
		RoundDuration: generalRoundInfo.RoundDur,
		CreatedAt:     block.Block.Time,
	}
	// calculate the total reward amount for this round
	genesisSupplyInt := sdkmath.NewInt(m.cfg.GenesisSupply)

	roundReward := m.cfg.GenesisPoolRatio.MulInt(genesisSupplyInt).MulInt64(generalRoundInfo.RoundDur).QuoInt64(generalRoundInfo.AirdropDur).TruncateInt()
	round.TotalRewardAmount = roundReward.String()

	// calculate and address the airdrop for all genesis stakers.
	err = m.HandleGenesisPoolAirdropForStakers(roundReward, round)
	if err != nil {
		return fmt.Errorf("error addressing genesis pool airdrop for all stakers: %w", err)
	}
	err = m.db.SaveGenesisPoolAirdropRound(round)
	if err != nil {
		return fmt.Errorf("error saving genesis pool airdrop round: %w", err)
	}

	return nil
}

func (m *Module) HandleLiquidityIncentivesAirdropRound(block *tmctypes.ResultBlock) error {
	if m.cfg.GenesisSupply == 0 {
		// do nothing when the genesis supply is zero.
		log.Info().Msg("the genesis supply is zero")
		return nil
	}
	generalRoundInfo, err := m.calculateAirdropRound(block, types.LiquidityIncentivesAirdrop)
	if err != nil {
		return err
	}
	if generalRoundInfo == nil {
		return nil
	}

	round := &types.CommonAirdropRound{
		AirdropRound:  generalRoundInfo.RoundID,
		BlockHeight:   block.Block.Height,
		RoundDuration: generalRoundInfo.RoundDur,
		CreatedAt:     block.Block.Time,
	}
	// calculate the total reward amount for this round
	genesisSupplyInt := sdkmath.NewInt(m.cfg.GenesisSupply)
	oneYear := 365 * 24 * time.Hour
	var rewardRatioIndex int
	if generalRoundInfo.PreRound != nil {
		// The start time of this round should be the end of the previous round, and each round is always created at the end.
		rewardRatioIndex = int(generalRoundInfo.PreRound.CreatedAt.Sub(generalRoundInfo.GenesisTime) / oneYear)
	}
	rewardRatiosLength := len(m.cfg.LiquidityIncentiveRatios)
	rewardRatio := sdkmath.LegacyZeroDec()
}

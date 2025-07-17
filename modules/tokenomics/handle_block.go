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

func (m *Module) HandleGenesisPoolAirdropForStakers(roundReward sdkmath.Int, round *types.GenesisPoolAirdropRound) error {
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

func (m *Module) HandleGenesisPoolAirdropRound(block *tmctypes.ResultBlock) error {
	// get the genesis time
	genesis, err := m.db.GetGenesis()
	if err != nil {
		return fmt.Errorf("error getting genesis: %w", err)
	}
	// calculate the round id by the genesis and block time.
	blockTime := block.Block.Time
	durFromGenesis := int64(blockTime.Sub(genesis.Time))
	tokenomicParams, err := m.db.GetTokenomicsParams()
	if err != nil {
		return fmt.Errorf("error getting tokenomic paramters: %w", err)
	}
	if tokenomicParams.GenesisPoolAirdropInterval == nil {
		return fmt.Errorf("error: GenesisPoolAirdropInterval is nil")
	} else if *tokenomicParams.GenesisPoolAirdropInterval <= 0 {
		return fmt.Errorf("error invalid interval for genesis pool airdrop: %d", *tokenomicParams.GenesisPoolAirdropInterval)
	}

	oneDay := int64(24 * time.Hour)
	interval := *tokenomicParams.LiquidityIncentiveAirdropInterval
	duration := *tokenomicParams.GenesisPoolAirdropDuration
	roundNumber := int((duration + interval - 1) / interval)
	airdropEndTime := genesis.Time.Add(time.Duration(duration * oneDay))

	roundID := int(durFromGenesis / (interval * oneDay))
	roundDur := interval
	if roundID == 0 || roundID > roundNumber {
		// Do nothing because the first airdrop round is not due yet or the airdrop has already ended.
		return nil
	} else if roundID == roundNumber-1 && blockTime.Compare(airdropEndTime) >= 0 {
		// Since the interval may not divide the duration evenly, the final round needs to be processed
		// right after the airdrop end time, which means its duration might be shorter than the interval.
		roundID = roundNumber
		roundDur = duration - interval*(int64(roundNumber)-1)
	}
	// check if this round has already addressed
	exist, err := m.db.HasGenesisPoolAirdropRound(roundID)
	if err != nil {
		return fmt.Errorf("error when checking genesis pool airdrop round: %w", err)
	}
	if exist {
		// Do nothing because this round has been addressed.
		return nil
	}

	round := &types.GenesisPoolAirdropRound{
		AirdropRound:  roundID,
		BlockHeight:   block.Block.Height,
		RoundDuration: roundDur,
		CreatedAt:     time.Now(),
	}
	// calculate the total reward amount for this round
	if tokenomicParams.GenesisSupply == nil {
		return fmt.Errorf("error: GenesisSupply is nil")
	}
	genesisSupplyInt, ok := sdkmath.NewIntFromString(*tokenomicParams.GenesisSupply)
	if !ok {
		return fmt.Errorf("invalid GenesisSupply: %s", *tokenomicParams.GenesisSupply)
	}
	if tokenomicParams.GenesisPoolRatio == nil {
		return fmt.Errorf("error: GenesisPoolRatio is nil")
	}
	genesisPoolRatio, err := sdkmath.LegacyNewDecFromStr(*tokenomicParams.GenesisPoolRatio)
	if !ok {
		return fmt.Errorf("invalid GenesisPoolRatio: %s,err:%w", *tokenomicParams.GenesisPoolRatio, err)
	}
	roundReward := genesisPoolRatio.MulInt(genesisSupplyInt).MulInt64(roundDur).QuoInt64(duration).TruncateInt()
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

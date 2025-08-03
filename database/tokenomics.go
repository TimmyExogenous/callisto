package database

import (
	sdkmath "cosmossdk.io/math"
	"database/sql"
	"errors"
	"fmt"
	"github.com/forbole/callisto/v4/types"
	"github.com/lib/pq"
)

func (db *Db) SaveTokenomicsParams(params *types.TokenomicsParams) error {
	stmt := `
INSERT INTO tokenomics_params (
    one_row_id,
    genesis_supply,
    genesis_pool_ratio,
    genesis_pool_airdrop_duration,
    genesis_pool_airdrop_interval,
    liquidity_incentive_ratios,
    liquidity_incentive_airdrop_duration,
    liquidity_incentive_airdrop_interval,
    genesis_validator_reward_ratio,
    created_at
) 
VALUES (
    TRUE,  -- one_row_id, always true
    COALESCE($1, DEFAULT),
    COALESCE($2, DEFAULT),
    COALESCE($3, DEFAULT),
    COALESCE($4, DEFAULT),
    COALESCE($5, DEFAULT),
    COALESCE($6, DEFAULT),
    COALESCE($7, DEFAULT),
    COALESCE($8, DEFAULT),
 	NOW(),
)
ON CONFLICT (one_row_id) DO UPDATE
	SET
		genesis_pool_ratio = COALESCE(EXCLUDED.genesis_pool_ratio, tokenomics_params.genesis_pool_ratio),
    	genesis_pool_airdrop_duration = COALESCE(EXCLUDED.genesis_pool_airdrop_duration, tokenomics_params.genesis_pool_airdrop_duration),
		genesis_pool_airdrop_interval = COALESCE(EXCLUDED.genesis_pool_airdrop_interval, tokenomics_params.genesis_pool_airdrop_interval),
		liquidity_incentive_ratios = COALESCE(EXCLUDED.liquidity_incentive_ratios, tokenomics_params.liquidity_incentive_ratios),
        liquidity_incentive_airdrop_duration = COALESCE(EXCLUDED.liquidity_incentive_airdrop_duration, tokenomics_params.liquidity_incentive_airdrop_duration),
		liquidity_incentive_airdrop_interval = COALESCE(EXCLUDED.liquidity_incentive_airdrop_interval, tokenomics_params.liquidity_incentive_airdrop_interval),
		genesis_validator_reward_ratio = COALESCE(EXCLUDED.genesis_validator_reward_ratio, tokenomics_params.genesis_validator_reward_ratio),
		created_at = NOW()`

	_, err := db.SQL.Exec(stmt,
		params.GenesisSupply,
		params.GenesisPoolRatio,
		params.GenesisPoolAirdropDuration,
		params.GenesisPoolAirdropInterval,
		pq.Array(params.LiquidityIncentiveRatios),
		params.LiquidityIncentiveAirdropDuration,
		params.LiquidityIncentiveAirdropInterval,
		params.GenesisValidatorRewardRatio,
	)

	if err != nil {
		return fmt.Errorf("error while saving tokenomics params: %w", err)
	}

	return nil
}

func (db *Db) GetTokenomicsParams() (*types.TokenomicsParams, error) {
	stmt := `
SELECT
    genesis_supply,
    genesis_pool_ratio,
    genesis_pool_airdrop_duration,
    genesis_pool_airdrop_interval,
    liquidity_incentive_ratios,
    liquidity_incentive_airdrop_duration,
    liquidity_incentive_airdrop_interval,
    genesis_validator_reward_ratio,
    created_at
FROM tokenomics_params
WHERE one_row_id = TRUE
LIMIT 1;`

	var params types.TokenomicsParams

	err := db.SQL.QueryRow(stmt).Scan(
		&params.GenesisSupply,
		&params.GenesisPoolRatio,
		&params.GenesisPoolAirdropDuration,
		&params.GenesisPoolAirdropInterval,
		pq.Array(&params.LiquidityIncentiveRatios),
		&params.LiquidityIncentiveAirdropDuration,
		&params.LiquidityIncentiveAirdropInterval,
		&params.GenesisValidatorRewardRatio,
		&params.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("tokenomics parameters not found")
		}
		return nil, fmt.Errorf("failed to query tokenomics parameters: %w", err)
	}

	return &params, nil
}

func (db *Db) SaveGenesisPoolAirdropRound(round *types.GenesisPoolAirdropRound) error {
	stmt := `
INSERT INTO genesis_pool_airdrop_rounds (
    airdrop_round,
    block_height,
    total_stakers,
    total_usd_value,
    total_reward_amount,
    distributed_stakers,
    round_duration,
    created_at,
    is_completed
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
ON CONFLICT (airdrop_round) DO NOTHING;`

	_, err := db.SQL.Exec(stmt,
		round.AirdropRound,
		round.BlockHeight,
		round.TotalStakers,
		round.TotalUSDValue,
		round.TotalRewardAmount,
		round.DistributedStakers,
		round.RoundDuration,
		round.CreatedAt,
		round.IsCompleted,
	)
	if err != nil {
		return fmt.Errorf("failed to save genesis pool airdrop round %d: %w", round.AirdropRound, err)
	}
	return nil
}

func (db *Db) HasGenesisPoolAirdropRound(roundIndex int) (bool, error) {
	stmt := `
SELECT 1
FROM genesis_pool_airdrop_rounds
WHERE airdrop_round = $1
LIMIT 1;`

	var exists int
	err := db.SQL.QueryRow(stmt, roundIndex).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil // round does not exist
		}
		return false, fmt.Errorf("failed to check if airdrop round %d exists: %w", roundIndex, err)
	}

	return true, nil
}

func (db *Db) GetGenesisPoolAirdropRound(roundIndex int) (*types.GenesisPoolAirdropRound, error) {
	stmt := `
SELECT
    airdrop_round,
    block_height,
    total_stakers,
    total_usd_value,
    total_reward_amount,
    distributed_stakers,
    round_duration,
    created_at,
    is_completed
FROM genesis_pool_airdrop_rounds
WHERE airdrop_round = $1;`

	var round types.GenesisPoolAirdropRound
	err := db.SQL.QueryRow(stmt, roundIndex).Scan(
		&round.AirdropRound,
		&round.BlockHeight,
		&round.TotalStakers,
		&round.TotalUSDValue,
		&round.TotalRewardAmount,
		&round.DistributedStakers,
		&round.RoundDuration,
		&round.CreatedAt,
		&round.IsCompleted,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("airdrop round %d not found", roundIndex)
		}
		return nil, fmt.Errorf("failed to query airdrop round %d: %w", roundIndex, err)
	}

	return &round, nil
}

func (db *Db) MarkGenesisPoolAirdropRoundCompleted(roundIndex int) error {
	stmt := `
UPDATE genesis_pool_airdrop_rounds
SET is_completed = TRUE
WHERE airdrop_round = $1;`

	result, err := db.SQL.Exec(stmt, roundIndex)
	if err != nil {
		return fmt.Errorf("failed to mark airdrop round %d as completed: %w", roundIndex, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected for round %d: %w", roundIndex, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no airdrop round found with index %d", roundIndex)
	}

	return nil
}

func (db *Db) SaveGenesisStakerAirdrop(airdrop *types.GenesisStakerAirdrop) error {
	stmt := `
INSERT INTO genesis_staker_airdrops (
    staker_id,
    airdrop_round,
    usd_value,
    reward_amount,
    is_distributed,
    distributed_at
) VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (staker_id, airdrop_round) DO NOTHING;`

	_, err := db.SQL.Exec(stmt,
		airdrop.StakerID,
		airdrop.AirdropRound,
		airdrop.USDValue,
		airdrop.RewardAmount,
		airdrop.IsDistributed,
		airdrop.DistributedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save genesis staker airdrop (%s, round %d): %w", airdrop.StakerID, airdrop.AirdropRound, err)
	}
	return nil
}

func (db *Db) MarkGenesisStakerAirdropDistributed(stakerID string, roundIndex int) error {
	stmt := `
UPDATE genesis_staker_airdrops
SET is_distributed = TRUE,
    distributed_at = NOW()
WHERE staker_id = $1 AND airdrop_round = $2;`

	result, err := db.SQL.Exec(stmt, stakerID, roundIndex)
	if err != nil {
		return fmt.Errorf("failed to mark airdrop as distributed for staker %s, round %d: %w", stakerID, roundIndex, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected for staker %s, round %d: %w", stakerID, roundIndex, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no record found for staker %s in round %d", stakerID, roundIndex)
	}

	return nil
}

func (db *Db) UpdateGenesisAirdropRewardsByRound(
	round int,
	calcReward func(usdValue sdkmath.LegacyDec) (sdkmath.Int, error),
) error {
	stmtSelect := `
    SELECT staker_id, usd_value
    FROM genesis_staker_airdrops
    WHERE airdrop_round = $1;
    `
	rows, err := db.SQL.Query(stmtSelect, round)
	if err != nil {
		return fmt.Errorf("failed to query airdrops for round %d: %w", round, err)
	}
	defer rows.Close()

	for rows.Next() {
		var stakerID string
		var usdValue string

		err := rows.Scan(&stakerID, &usdValue)
		if err != nil {
			return fmt.Errorf("failed to scan airdrop row: %w", err)
		}

		usdValueDec, err := sdkmath.LegacyNewDecFromStr(usdValue)
		if err != nil {
			return fmt.Errorf("invalid staker USD value: %w, usdValue:%s", err, usdValue)
		}
		newReward, err := calcReward(usdValueDec)
		if err != nil {
			return fmt.Errorf("failed to calculate reward for staker %s: %w", stakerID, err)
		}

		stmtUpdate := `
        UPDATE genesis_staker_airdrops
        SET reward_amount = $1
        WHERE staker_id = $2 AND airdrop_round = $3;
        `
		_, err = db.SQL.Exec(stmtUpdate, newReward.String(), stakerID, round)
		if err != nil {
			return fmt.Errorf("failed to update reward for staker %s: %w", stakerID, err)
		}
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("row iteration error: %w", err)
	}

	return nil
}

CREATE TABLE tokenomics_params
(
    one_row_id                           BOOLEAN   NOT NULL DEFAULT TRUE PRIMARY KEY,
    -- default genesis supply value is from the tokenomics documentation
    genesis_supply                       NUMERIC   NOT NULL DEFAULT 314159265,

    -- default ratio is from the tokenomics documentation
    genesis_pool_ratio                   NUMERIC   NOT NULL DEFAULT 0.0300,
    -- The genesis pool airdrop will last for 90 days after TGE.
    genesis_pool_airdrop_duration        NUMERIC   NOT NULL DEFAULT 90,
    -- We calculate and execute an airdrop per week (every 7 days).
    -- We can adjust the default value if needed.
    genesis_pool_airdrop_interval        NUMERIC   NOT NULL DEFAULT 7,

    -- default ratios is from the tokenomics documentation
    -- it's similar to the `annual_inflation` in the parameter of immint module.
    -- The rate is selected based on the time since genesis time.
    -- This list stores all annual inflation ratios starting from genesis time.
    -- If the current time exceeds genesis_time + len(list) * 1 year, the last ratio in
    -- the list will be used.
    liquidity_incentive_ratios           NUMERIC[] NOT NULL DEFAULT ARRAY[0.0200, 0.0100, 0.0050, 0.0025, 0.0013],
    -- The liquidity incentive airdrop will last for ten years(ignore the leap year) after TGE.
    liquidity_incentive_airdrop_duration NUMERIC   NOT NULL DEFAULT 3650,
    -- We calculate and execute an airdrop each quarter (every 90 days).
    -- We can adjust the default value if needed.
    liquidity_incentive_airdrop_interval NUMERIC   NOT NULL DEFAULT 90,

    -- default ratio is from the tokenomics documentation
    genesis_validator_reward_ratio       NUMERIC   NOT NULL DEFAULT 0.0200,
    -- Record creation timestamp
    created_at                           TIMESTAMP NOT NULL DEFAULT now(),

    CHECK (one_row_id = TRUE),
    CHECK (genesis_pool_ratio < 1),
    CHECK (genesis_validator_reward_ratio < 1),
    CHECK (
        (SELECT bool_and(ratio < 1)
         FROM unnest(liquidity_incentive_ratios) AS ratio)
        )
);

CREATE TABLE genesis_pool_airdrop_rounds
(
    airdrop_round       INTEGER PRIMARY KEY,              -- Airdrop round index
    block_height        BIGINT    NOT NULL,               -- Snapshot block height
    total_stakers       INTEGER   NOT NULL,               -- Total number of stakers in this round
    total_usd_value     NUMERIC   NOT NULL,               -- Total USD value aggregated across stakers
    total_reward_amount NUMERIC   NOT NULL,               -- Total reward allocated for this round
    distributed_stakers INTEGER   NOT NULL DEFAULT 0,     -- Number of stakers who have received rewards
    round_duration      NUMERIC   NOT NULL,               -- the duration for this round
    created_at          TIMESTAMP NOT NULL DEFAULT now(), -- Record creation timestamp
    is_completed        BOOLEAN   NOT NULL DEFAULT FALSE  -- Whether this round's distribution is finished
);

-- These are the airdrop states for genesis stakers, used to distribute the genesis pool reward.
CREATE TABLE genesis_staker_airdrops
(
    staker_id      TEXT    NOT NULL,               -- The unique identifier of the staker
    airdrop_round  INTEGER NOT NULL,               -- The airdrop round (e.g., weekly interval index)
    usd_value      NUMERIC NOT NULL,               -- The staker's total asset value in USD at snapshot time
    reward_amount  NUMERIC NOT NULL,               -- The reward amount allocated to the staker
    is_distributed BOOLEAN NOT NULL DEFAULT FALSE, -- Whether the reward has been distributed
    distributed_at TIMESTAMP,                      -- Timestamp when the reward was distributed (nullable)

    PRIMARY KEY (staker_id, airdrop_round),
    CONSTRAINT fk_airdrop_round FOREIGN KEY (airdrop_round) REFERENCES genesis_pool_airdrop_rounds (airdrop_round),
);
CREATE INDEX idx_genesis_airdrops_round ON genesis_staker_airdrops (airdrop_round);
CREATE INDEX idx_genesis_airdrops_staker ON genesis_staker_airdrops (staker_id);


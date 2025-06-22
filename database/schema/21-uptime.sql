-- for each height, create a row for each validator that is supposed to sign the block.
-- record whether the validator signed the block.
CREATE TABLE block_uptime (
    height BIGINT NOT NULL,
    consensus_address TEXT NOT NULL,
    signed BOOLEAN NOT NULL,
    PRIMARY KEY (height, consensus_address),
    -- each cons addr will be in the validator set at some height
    FOREIGN KEY (consensus_address) REFERENCES validator(consensus_address)
    -- each cons addr will be part of the consensus keys history
    -- which contains consensus keys + addrs across all chains and heights
    -- we are examining a subset of that, however, since it is not the primary
    -- key, sql will complain
    -- FOREIGN KEY (consensus_address) REFERENCES consensus_keys_history(cons_addr),
    -- each height will have a block, identified uniquely by height.
    -- not true for the genesis + 1 block, since there will not be a genesis block.
    -- FOREIGN KEY (height) REFERENCES block(height)
);

-- View is required by Hasura to track the function
CREATE OR REPLACE VIEW validator_uptime_structure AS
SELECT
    NULL::TEXT AS operator_addr,
    NULL::BIGINT AS expected_blocks,
    NULL::BIGINT AS signed_blocks,
    NULL::NUMERIC AS uptime;
CREATE OR REPLACE FUNCTION get_validator_uptime(
    in_operator_addr TEXT,
    in_chain_id TEXT
)
RETURNS SETOF validator_uptime_structure
LANGUAGE plpgsql
STABLE
AS $$
BEGIN
    RETURN QUERY
    WITH relevant_consensus_keys AS (
        SELECT
            h.cons_addr,
            h.first_activation_height,
            h.last_active_height
        FROM consensus_keys_history h
        WHERE h.operator_addr = in_operator_addr
          AND h.chain_id = in_chain_id
    ),
    signing_events AS (
        SELECT bu.signed
        FROM block_uptime bu
        JOIN relevant_consensus_keys rck
          ON bu.consensus_address = rck.cons_addr
         AND (
             (rck.first_activation_height IS NULL OR bu.height >= rck.first_activation_height)
             AND
             (rck.last_active_height IS NULL OR bu.height <= rck.last_active_height)
         )
    )

    SELECT
        in_operator_addr::TEXT,
        COUNT(*)::BIGINT,
        COUNT(*) FILTER (WHERE signed)::BIGINT,
        ROUND(
            (COUNT(*) FILTER (WHERE signed))::NUMERIC / NULLIF(COUNT(*), 0),
            4
        )::NUMERIC
    FROM signing_events
    GROUP BY in_operator_addr;
END;
$$;
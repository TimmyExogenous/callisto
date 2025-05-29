CREATE TABLE oracle_params
(
    one_row_id BOOLEAN NOT NULL DEFAULT TRUE PRIMARY KEY,
    params     JSONB   NOT NULL,
    height     BIGINT  NOT NULL,
    CHECK (one_row_id)
);

-- For any token, we will store the next round id
-- TODO add more configuration fields?
CREATE TABLE oracle_token_config (
    token_id BIGINT PRIMARY KEY,
    next_round_id BIGINT NOT NULL DEFAULT 1
);

-- we should have a mapping from asset id to token id
-- each asset id will map to one token id
-- however, a token id can map to multiple asset ids
CREATE TABLE oracle_asset_id_to_token_id (
    asset_id TEXT PRIMARY KEY,
    token_id BIGINT NOT NULL,
    FOREIGN KEY (asset_id)
        REFERENCES assets_tokens (asset_id),
    FOREIGN KEY (token_id)
        REFERENCES oracle_token_config (token_id)
);

-- For any token, we will store the price history for each round.
CREATE TABLE oracle_price_history (
    token_id BIGINT NOT NULL,
    round_id BIGINT NOT NULL,
    price NUMERIC NOT NULL,
    price_decimals INTEGER NOT NULL,
    price_timestamp TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    PRIMARY KEY (token_id, round_id),
    CONSTRAINT fk_oracle_price_history_token
        FOREIGN KEY (token_id)
        REFERENCES oracle_token_config (token_id)
);

CREATE INDEX IF NOT EXISTS idx_oracle_price_history_token_id ON oracle_price_history(token_id);
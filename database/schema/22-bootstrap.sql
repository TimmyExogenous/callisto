CREATE TABLE bootstrap_validator
(
    validator_eth_addr  TEXT    NOT NULL PRIMARY KEY,
    validator_im_addr   TEXT    NOT NULL UNIQUE,
    validator_name      TEXT    NOT NULL UNIQUE,
    consensus_pub_key   TEXT    NOT NULL UNIQUE,
    commission_rate     NUMERIC NOT NULL,
    max_commission_rate NUMERIC NOT NULL,
    max_change_rate     NUMERIC NOT NULL,
    updated_at          TIMESTAMP WITHOUT TIME ZONE,
);

CREATE INDEX idx_validator_im_addr ON bootstrap_validator (validator_im_addr);

CREATE TABLE bootstrap_client_chains
(
    name                TEXT NOT NULL,
    meta_info           TEXT NOT NULL,
    layer_zero_chain_id BIGINT PRIMARY KEY,
    updated_at          TIMESTAMP WITHOUT TIME ZONE,
);

CREATE TABLE bootstrap_tokens
(
    -- generated for ease; not required to be part of the schema
    asset_id             TEXT PRIMARY KEY,
    name                 TEXT    NOT NULL,
    symbol               TEXT    NOT NULL,
    address              TEXT    NOT NULL CHECK (address = lower(address)),
    decimals             INT     NOT NULL,
    layer_zero_chain_id  BIGINT  NOT NULL,
    staking_total_amount NUMERIC NOT NULL DEFAULT 0,
    updated_at           TIMESTAMP WITHOUT TIME ZONE,
    -- relational constraint
    CONSTRAINT fk_layer_zero_chain_id FOREIGN KEY (layer_zero_chain_id) REFERENCES bootstrap_client_chains (layer_zero_chain_id)
);

CREATE TABLE bootstrap_token_prices
(
    asset_id   TEXT PRIMARY KEY,
    price      NUMERIC NOT NULL DEFAULT 0,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    CONSTRAINT fk_asset_id FOREIGN KEY (asset_id) REFERENCES bootstrap_tokens (asset_id)
);


CREATE TABLE bootstrap_staker_assets
(
    staker_id    TEXT    NOT NULL,
    asset_id     TEXT    NOT NULL,
    deposited    NUMERIC NOT NULL DEFAULT 0,
    withdrawable NUMERIC NOT NULL DEFAULT 0,
    delegated    NUMERIC NOT NULL DEFAULT 0,
    updated_at   TIMESTAMP WITHOUT TIME ZONE,
    PRIMARY KEY (staker_id, asset_id),
    CONSTRAINT chk_total CHECK (deposited = withdrawable + delegated),
    CONSTRAINT fk_asset_id FOREIGN KEY (asset_id) REFERENCES bootstrap_tokens (asset_id)
);

CREATE INDEX idx_deposits_staker_id ON bootstrap_staker_assets (staker_id);
CREATE INDEX idx_deposits_asset_id ON bootstrap_staker_assets (asset_id);

CREATE TABLE bootstrap_delegation_states
(
    staker_id     TEXT    NOT NULL,
    asset_id      TEXT    NOT NULL,
    operator_addr TEXT    NOT NULL,
    delegated     NUMERIC NOT NULL DEFAULT 0,
    updated_at    TIMESTAMP WITHOUT TIME ZONE,
    PRIMARY KEY (staker_id, asset_id, operator_addr),
    CONSTRAINT fk_operator FOREIGN KEY (operator_addr) REFERENCES bootstrap_validator (validator_im_addr),
    CONSTRAINT fk_asset_id FOREIGN KEY (asset_id) REFERENCES bootstrap_tokens (asset_id),
    CONSTRAINT fk_staker_asset FOREIGN KEY (staker_id, asset_id) REFERENCES bootstrap_staker_assets (staker_id, asset_id)
);

CREATE INDEX idx_delegations_staker_id ON bootstrap_delegation_states (staker_id);
CREATE INDEX idx_delegations_asset_id ON bootstrap_delegation_states (asset_id);

CREATE TABLE bootstrap_operator_assets
(
    operator_addr TEXT    NOT NULL,
    asset_id      TEXT    NOT NULL,
    total_amount  NUMERIC NOT NULL,
    self_amount   NUMERIC NOT NULL DEFAULT 0,
    other_amount  NUMERIC NOT NULL DEFAULT 0,
    updated_at    TIMESTAMP WITHOUT TIME ZONE,
    PRIMARY KEY (operator_addr, asset_id),
    CONSTRAINT fk_asset_id FOREIGN KEY (asset_id) REFERENCES bootstrap_tokens (asset_id),
    CONSTRAINT fk_operator FOREIGN KEY (operator_addr) REFERENCES bootstrap_validator (validator_im_addr),
    CONSTRAINT chk_total_amount CHECK (total_amount = self_amount + other_amount)
);
CREATE INDEX idx_operator_assets_operator ON bootstrap_operator_assets (operator_addr);
CREATE INDEX idx_operator_assets_asset_id ON bootstrap_operator_assets (asset_id);

CREATE TABLE bootstrap_statistics
(
    one_row_id BOOLEAN NOT NULL DEFAULT TRUE PRIMARY KEY,
    tvl        NUMERIC NOT NULL,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    CHECK (one_row_id)
);

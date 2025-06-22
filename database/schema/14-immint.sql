-- single row because the params history is not relevant and not changing often
CREATE TABLE immint_params
(
    one_row_id BOOLEAN NOT NULL DEFAULT TRUE PRIMARY KEY,
    params     JSONB   NOT NULL,
    height     BIGINT  NOT NULL,
    CHECK (one_row_id)
);

-- track mint history for showing it as events
CREATE TABLE immint_history (
    block_height BIGINT PRIMARY KEY,
    quantity_minted NUMERIC NOT NULL,
    epoch_id TEXT NOT NULL,
    epoch_number BIGINT NOT NULL,
    denom TEXT NOT NULL,
    CONSTRAINT fk_epoch_id FOREIGN KEY (epoch_id) REFERENCES epoch_states (identifier),
    CONSTRAINT unique_epoch_id_epoch_number UNIQUE (epoch_id, epoch_number)
);
CREATE INDEX idx_immint_epoch_number ON immint_history (epoch_number);
CREATE INDEX idx_immint_block_height ON immint_history (block_height);

CREATE OR REPLACE VIEW total_minted_structure AS
SELECT NULL::NUMERIC AS total;
CREATE OR REPLACE FUNCTION total_minted()
RETURNS SETOF total_minted_structure
LANGUAGE plpgsql
STABLE
AS $$
BEGIN
    RETURN QUERY
    SELECT
        COALESCE(SUM(quantity_minted), 0) AS total -- Alias matches the view's column
    FROM
        immint_history;
END;
$$;

CREATE OR REPLACE VIEW minting_per_block_structure AS
SELECT
    NULL::BIGINT AS block_height,
    NULL::NUMERIC AS quantity_minted;
CREATE OR REPLACE FUNCTION minting_per_block()
RETURNS SETOF minting_per_block_structure
LANGUAGE plpgsql
STABLE
AS $$
BEGIN
    RETURN QUERY
    SELECT
        ih.block_height,
        ih.quantity_minted
    FROM
        immint_history ih
    ORDER BY
        ih.block_height DESC;
END;
$$;

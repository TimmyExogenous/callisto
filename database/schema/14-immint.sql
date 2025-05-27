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

CREATE OR REPLACE FUNCTION total_minted()
RETURNS NUMERIC AS $$
BEGIN
    RETURN (SELECT COALESCE(SUM(quantity_minted), 0) FROM immint_history);
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION minting_per_block()
RETURNS TABLE(block_height BIGINT, quantity_minted NUMERIC) AS $$
BEGIN
    RETURN QUERY
    SELECT block_height, quantity_minted
    FROM immint_history
    ORDER BY block_height;
END;
$$ LANGUAGE plpgsql;


-- static, unchanging data
CREATE TABLE epoch_definitions (
    identifier TEXT PRIMARY KEY,
    start_time TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    duration INTERVAL NOT NULL
);

-- dynamic data
CREATE TABLE epoch_states (
    identifier TEXT PRIMARY KEY REFERENCES epoch_definitions(identifier),
    current_epoch BIGINT NOT NULL,
    current_epoch_start_time TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    epoch_counting_started BOOLEAN NOT NULL DEFAULT FALSE,
    current_epoch_start_height BIGINT NOT NULL
);

-- These indexes help the functions below
CREATE INDEX idx_epoch_identifier_time ON epoch_states (identifier, current_epoch_start_time);
CREATE INDEX idx_epoch_identifier_height ON epoch_states (identifier, current_epoch_start_height);

-- The view is required because of Hasura.
CREATE OR REPLACE VIEW epoch_number_structure AS
SELECT NULL::BIGINT AS epoch_number;

-- epoch by height
CREATE OR REPLACE FUNCTION get_epoch_number_by_height(
    in_identifier TEXT,
    in_block_height BIGINT
)
RETURNS SETOF epoch_number_structure
LANGUAGE plpgsql
STABLE
AS $$
BEGIN
    RETURN QUERY
    SELECT es.current_epoch
    FROM epoch_states es
    WHERE es.identifier = in_identifier
      AND es.current_epoch_start_height <= in_block_height
    ORDER BY es.current_epoch_start_height DESC
    LIMIT 1;
END;
$$;

-- epoch by time, though not exactly used.
CREATE OR REPLACE FUNCTION get_epoch_number_by_time(
    in_identifier TEXT,
    in_block_time TIMESTAMP
)
RETURNS SETOF epoch_number_structure
LANGUAGE plpgsql
STABLE
AS $$
BEGIN
    RETURN QUERY
    SELECT es.current_epoch AS epoch_number
    FROM epoch_states es
    WHERE es.identifier = in_identifier
      AND es.current_epoch_start_time <= in_block_time
    ORDER BY es.current_epoch_start_time DESC
    LIMIT 1;
END;
$$;
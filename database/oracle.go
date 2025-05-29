package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/forbole/callisto/v4/types"
	oracletypes "github.com/imua-xyz/imuachain/x/oracle/types"
)

// getExistingOracleParamsJSON retrieves the current oracle parameters as a
// JSON string and the corresponding height.
func (db *Db) getExistingOracleParamsJSON() (paramsJSON string, height int64, err error) {
	stmt := `SELECT params, height FROM oracle_params WHERE one_row_id = TRUE LIMIT 1`
	err = db.SQL.QueryRow(stmt).Scan(&paramsJSON, &height)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", 0, sql.ErrNoRows
		}
		return "", 0, fmt.Errorf("error querying existing oracle params: %w", err)
	}
	return paramsJSON, height, nil
}

// UpdateOracleParamsIfNeeded compares new OracleParams with existing ones in the DB
// and saves if they are different or if the new height is greater.
func (db *Db) UpdateOracleParamsIfNeeded(newParams *types.OracleParams) error {
	existingParamsJSON, _, err := db.getExistingOracleParamsJSON()
	newParamsBz, marshalErr := json.Marshal(&newParams.Params)
	if marshalErr != nil {
		return fmt.Errorf("error marshaling new oracle params for comparison: %w", marshalErr)
	}

	if err != nil {
		if err == sql.ErrNoRows {
			// no existing params, save new ones
			return db.SaveOracleParams(newParams)
		}
		return fmt.Errorf("error getting existing oracle params: %w", err)
	}

	if string(newParamsBz) != existingParamsJSON {
		return db.SaveOracleParams(newParams)
	}

	return nil
}

// SaveOracleParams allows to store the given params inside the database
func (db *Db) SaveOracleParams(params *types.OracleParams) error {
	// step 1. raw value of params
	paramsBz, err := json.Marshal(&params.Params)
	if err != nil {
		return fmt.Errorf("error while marshaling oracle params: %s", err)
	}

	stmt := `
INSERT INTO oracle_params (params, height) 
VALUES ($1, $2)
ON CONFLICT (one_row_id) DO UPDATE 
    SET params = excluded.params,
        height = excluded.height
WHERE oracle_params.height <= excluded.height`

	_, err = db.SQL.Exec(stmt, string(paramsBz), params.Height)
	if err != nil {
		return fmt.Errorf("error while storing oracle params: %s", err)
	}

	for i, token := range params.Params.Tokens {
		if i == 0 {
			// skip the first token because it is reserved
			continue
		}
		tokenID := i

		// step 2. token config
		tokenConfig := types.NewOracleTokenConfigFromStr(
			fmt.Sprintf("%d", tokenID),
			// default value of 1
			fmt.Sprintf("%d", 1),
		)
		// if it exists, we will not save it again
		if err := db.SaveOracleTokenConfigIfNeeded(tokenConfig); err != nil {
			return fmt.Errorf("error while storing oracle token config: %s", err)
		}

		// step 3. mapping from token id to asset id
		assetIDs := strings.Split(token.AssetID, ",")
		for _, assetID := range assetIDs {
			if err := db.SaveOracleAssetIDToTokenID(assetID, tokenID); err != nil {
				return fmt.Errorf("error while storing oracle asset id to token id: %s", err)
			}
		}
	}

	return nil
}

// SaveOracleAssetIDToTokenID stores the given asset id to token id mapping inside the database
func (db *Db) SaveOracleAssetIDToTokenID(assetID string, tokenID int) error {
	stmt := `
INSERT INTO oracle_asset_id_to_token_id (asset_id, token_id)
VALUES ($1, $2)
ON CONFLICT (asset_id) DO UPDATE
	SET token_id = excluded.token_id
WHERE oracle_asset_id_to_token_id.asset_id = excluded.asset_id;`

	_, err := db.SQL.Exec(stmt, assetID, tokenID)
	if err != nil {
		return fmt.Errorf("error while storing oracle asset id to token id: %s", err)
	}
	return nil
}

// SaveOracleTokenConfigIfNeeded stores the given token config inside the database
func (db *Db) SaveOracleTokenConfigIfNeeded(tokenConfig *types.OracleTokenConfig) error {
	stmt := `
INSERT INTO oracle_token_config (token_id, next_round_id)
VALUES ($1, $2)
ON CONFLICT (token_id) DO NOTHING;`

	_, err := db.SQL.Exec(stmt, tokenConfig.TokenID, tokenConfig.NextRoundID)
	if err != nil {
		return fmt.Errorf("error while storing oracle token config: %s", err)
	}
	return nil
}

// IncreaseNextRoundID increases the next round id for the given token id
func (db *Db) IncreaseNextRoundID(tokenID string) error {
	stmt := `
UPDATE oracle_token_config
SET next_round_id = next_round_id + 1
WHERE oracle_token_config.token_id = $1;`

	_, err := db.SQL.Exec(stmt, tokenID)
	if err != nil {
		return fmt.Errorf("error while increasing next round id: %s", err)
	}
	return nil
}

// SaveOraclePriceHistory stores the given price history inside the database
func (db *Db) SaveOraclePriceHistory(priceHistory *types.OraclePriceHistory) error {
	parsed, err := time.Parse(oracletypes.TimeLayout, priceHistory.PriceTimestamp)
	if err != nil {
		return fmt.Errorf(
			"error while parsing oracle price history timestamp '%s': %w",
			priceHistory.PriceTimestamp,
			err,
		)
	}

	stmt := `
INSERT INTO oracle_price_history (token_id, round_id, price, price_decimals, price_timestamp)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (token_id, round_id) DO NOTHING;`

	if _, err := db.SQL.Exec(
		stmt,
		priceHistory.TokenID,       // $1
		priceHistory.RoundID,       // $2
		priceHistory.Price,         // $3 (string, assuming PostgreSQL converts to NUMERIC)
		priceHistory.PriceDecimals, // $4 (int, assuming PostgreSQL converts to INTEGER)
		parsed,                     // $5 (time.Time, pq driver converts to TIMESTAMP)
	); err != nil {
		// Use %w for error wrapping
		return fmt.Errorf(
			"error while storing oracle price history (token: %s, round: %s): %w",
			priceHistory.TokenID,
			priceHistory.RoundID,
			err,
		)
	}
	return nil
}

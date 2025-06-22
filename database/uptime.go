package database

import "fmt"

// SaveBlockUptime saves the signing status of a validator for a specific block height.
// If an entry already exists for the (height, consensus_address) pair, it will be updated.
// TODO: aggregate writing across multiple validators at once.
func (db *Db) SaveBlockUptime(
	height int64, consensusAddr string, signed bool,
) error {
	stmt := `
	INSERT INTO block_uptime (height, consensus_address, signed)
	VALUES ($1, $2, $3)
	ON CONFLICT (height, consensus_address) DO UPDATE
	SET signed = $3;`
	_, err := db.SQL.Exec(stmt, height, consensusAddr, signed)
	if err != nil {
		return fmt.Errorf("failed to save block uptime: %w", err)
	}
	return nil
}

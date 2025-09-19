package database

import (
	"database/sql"
	"fmt"
	"github.com/forbole/callisto/v4/types"
	"time"
)

func (db *Db) SaveBootstrapValidator(v *types.BootstrapValidator) error {
	stmt := `
INSERT INTO bootstrap_validator (
    validator_eth_addr,
    validator_im_addr,
    validator_name,
    consensus_pub_key,
    commission_rate,
    max_commission_rate,
    max_change_rate,
    updated_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
ON CONFLICT (validator_eth_addr) DO UPDATE
SET validator_im_addr   = EXCLUDED.validator_im_addr,
    validator_name      = EXCLUDED.validator_name,
    consensus_pub_key   = EXCLUDED.consensus_pub_key,
    commission_rate     = EXCLUDED.commission_rate,
    max_commission_rate = EXCLUDED.max_commission_rate,
    max_change_rate     = EXCLUDED.max_change_rate,
    updated_at          = EXCLUDED.updated_at;`

	_, err := db.SQL.Exec(
		stmt,
		v.ValidatorEthAddress,
		v.ValidatorIMAddress,
		v.ValidatorName,
		v.ConsensusPubKey,
		v.Rate,
		v.MaxRate,
		v.MaxChangeRate,
		v.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save bootstrap validator: %w", err)
	}

	return nil
}

// UpdateCommissionRate updates the commission rate for a bootstrap validator identified
// by validatorEthAddr. It also updates the updated_at timestamp to the current time.
func (db *Db) UpdateCommissionRate(validatorEthAddr string, newRate string) error {
	stmt := `
UPDATE bootstrap_validator
SET commission_rate     = $1,
    updated_at          = $2
WHERE validator_eth_addr = $3;
`

	_, err := db.SQL.Exec(
		stmt,
		newRate,
		time.Now(),
		validatorEthAddr,
	)
	if err != nil {
		return fmt.Errorf("failed to update commission rate for %s: %w", validatorEthAddr, err)
	}
	return nil
}

// UpdateConsensusPubKey updates the consensus public key for a bootstrap validator
// identified by validatorIMAddr. The updated_at timestamp is also set to the current time.
func (db *Db) UpdateConsensusPubKey(validatorIMAddr string, newPubKey string) error {
	stmt := `
UPDATE bootstrap_validator
SET consensus_pub_key = $1,
    updated_at        = $2
WHERE validator_im_addr = $3;
`

	_, err := db.SQL.Exec(
		stmt,
		newPubKey,
		time.Now(),
		validatorIMAddr,
	)
	if err != nil {
		return fmt.Errorf("failed to update consensus pub key for %s: %w", validatorIMAddr, err)
	}
	return nil
}

func (db *Db) SaveBootstrapClientChain(c *types.BootstrapClientChain) error {
	stmt := `
INSERT INTO bootstrap_client_chains (
    name, meta_info, layer_zero_chain_id, updated_at
) VALUES ($1, $2, $3, $4)
ON CONFLICT (layer_zero_chain_id) DO UPDATE
SET name       = EXCLUDED.name,
    meta_info  = EXCLUDED.meta_info,
    updated_at = EXCLUDED.updated_at;`

	_, err := db.SQL.Exec(stmt,
		c.Name,
		c.MetaInfo,
		c.LayerZeroChainID,
		c.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save bootstrap client chain: %w", err)
	}
	return nil
}

func (db *Db) SaveBootstrapToken(t *types.BootstrapToken) error {
	stmt := `
INSERT INTO bootstrap_tokens (
    asset_id, name, symbol, address, decimals,
    layer_zero_chain_id, staking_total_amount, updated_at
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
ON CONFLICT (asset_id) DO UPDATE
SET name                 = EXCLUDED.name,
    symbol               = EXCLUDED.symbol,
    address              = EXCLUDED.address,
    decimals             = EXCLUDED.decimals,
    layer_zero_chain_id  = EXCLUDED.layer_zero_chain_id,
    staking_total_amount = EXCLUDED.staking_total_amount,
    updated_at           = EXCLUDED.updated_at;`

	_, err := db.SQL.Exec(stmt,
		t.AssetID,
		t.Name,
		t.Symbol,
		t.Address,
		t.Decimals,
		t.LayerZeroChainID,
		t.StakingTotalAmount,
		t.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save bootstrap token: %w", err)
	}
	return nil
}

func (db *Db) SaveBootstrapStakerAsset(a *types.BootstrapStakerAsset) error {
	stmt := `
INSERT INTO bootstrap_staker_assets (
    staker_id, asset_id, deposited, withdrawable, delegated, updated_at
) VALUES ($1,$2,$3,$4,$5,$6)
ON CONFLICT (staker_id, asset_id) DO UPDATE
SET deposited    = EXCLUDED.deposited,
    withdrawable = EXCLUDED.withdrawable,
    delegated    = EXCLUDED.delegated,
    updated_at   = EXCLUDED.updated_at;`

	_, err := db.SQL.Exec(stmt,
		a.StakerID,
		a.AssetID,
		a.Deposited,
		a.Withdrawable,
		a.Delegated,
		a.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save bootstrap staker asset: %w", err)
	}
	return nil
}

// DepositBootstrapStakerAsset increases the deposited and withdrawable amount
// for the given staker and asset. If no record exists, a new one will be created
// with deposited = depositAmount, withdrawable = depositAmount, delegated = 0.
// The updated_at timestamp is always set to the current time.
// TODO: The following two functions might not be used because we fetch the states
// from bootstrap directly instead of calculating the new states based on the delta value
// in events. They can be removed if they are not used for handling the states
// of BTC and XRP either.

func (db *Db) DepositBootstrapStakerAsset(stakerID, assetID string, depositAmount int64) error {
	stmt := `
INSERT INTO bootstrap_staker_assets (
    staker_id, asset_id, deposited, withdrawable, delegated, updated_at
) VALUES ($1, $2, $3, $3, 0, $4)
ON CONFLICT (staker_id, asset_id) DO UPDATE
SET deposited    = bootstrap_staker_assets.deposited + EXCLUDED.deposited,
    withdrawable = bootstrap_staker_assets.withdrawable + EXCLUDED.withdrawable,
    updated_at   = EXCLUDED.updated_at;`

	_, err := db.SQL.Exec(
		stmt,
		stakerID,
		assetID,
		depositAmount,
		time.Now(),
	)
	if err != nil {
		return fmt.Errorf("failed to deposit bootstrap staker asset: %w", err)
	}
	return nil
}

func (db *Db) ClaimBootstrapStakerAsset(stakerID, assetID string, claimAmount int64, updatedAt time.Time) error {
	stmt := `
UPDATE bootstrap_staker_assets
SET 
    deposited    = deposited - $3,
    withdrawable = withdrawable - $3,
    updated_at   = $4
WHERE staker_id = $1
  AND asset_id  = $2
  AND withdrawable >= $3
  AND deposited >= $3;`

	res, err := db.SQL.Exec(stmt,
		stakerID,
		assetID,
		claimAmount,
		updatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to claim bootstrap staker asset: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check claim rows affected: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("claim failed: either record does not exist or insufficient balance")
	}

	return nil
}

func (db *Db) SaveBootstrapDelegationState(d *types.BootstrapDelegationState) error {
	stmt := `
INSERT INTO bootstrap_delegation_states (
    staker_id, asset_id, operator_addr, delegated, updated_at
) VALUES ($1,$2,$3,$4,$5)
ON CONFLICT (staker_id, asset_id, operator_addr) DO UPDATE
SET delegated  = EXCLUDED.delegated,
    updated_at = EXCLUDED.updated_at;`

	_, err := db.SQL.Exec(stmt,
		d.StakerID,
		d.AssetID,
		d.OperatorAddr,
		d.Delegated,
		d.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save bootstrap delegation state: %w", err)
	}
	return nil
}

func (db *Db) SaveBootstrapOperatorAsset(o *types.BootstrapOperatorAsset) error {
	stmt := `
INSERT INTO bootstrap_operator_assets (
    operator_addr, asset_id, total_amount, self_amount, other_amount, updated_at
) VALUES ($1,$2,$3,$4,$5,$6)
ON CONFLICT (operator_addr, asset_id) DO UPDATE
SET total_amount = EXCLUDED.total_amount,
    self_amount  = EXCLUDED.self_amount,
    other_amount = EXCLUDED.other_amount,
    updated_at   = EXCLUDED.updated_at;`

	_, err := db.SQL.Exec(stmt,
		o.OperatorAddr,
		o.AssetID,
		o.TotalAmount,
		o.SelfAmount,
		o.OtherAmount,
		o.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save bootstrap operator asset: %w", err)
	}
	return nil
}

// Incremental scanning state management functions

// GetScanState retrieves the scanning state for a specific chain type
func (db *Db) GetScanState(chainType string) (*types.ScanState, error) {
	stmt := `SELECT chain_type, last_height, last_hash, safe_height, updated_at, created_at
             FROM bootstrap_scan_state WHERE chain_type = $1`

	var state types.ScanState
	err := db.SQL.QueryRow(stmt, chainType).Scan(
		&state.ChainType,
		&state.LastHeight,
		&state.LastHash,
		&state.SafeHeight,
		&state.UpdatedAt,
		&state.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get scan state for %s: %w", chainType, err)
	}

	return &state, nil
}

// UpdateScanState updates or inserts the scanning state for a specific chain type
func (db *Db) UpdateScanState(state *types.ScanState) error {
	stmt := `
INSERT INTO bootstrap_scan_state (chain_type, last_height, last_hash, safe_height, updated_at)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (chain_type) DO UPDATE
SET last_height = EXCLUDED.last_height,
    last_hash   = EXCLUDED.last_hash,
    safe_height = EXCLUDED.safe_height,
    updated_at  = EXCLUDED.updated_at`

	_, err := db.SQL.Exec(stmt,
		state.ChainType,
		state.LastHeight,
		state.LastHash,
		state.SafeHeight,
		state.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to update scan state for %s: %w", state.ChainType, err)
	}

	return nil
}

// IsTransactionProcessed checks if a transaction has already been processed
func (db *Db) IsTransactionProcessed(chainType, txHash string) (bool, error) {
	stmt := `SELECT EXISTS(SELECT 1 FROM bootstrap_processed_transactions
                          WHERE chain_type = $1 AND tx_hash = $2)`

	var exists bool
	err := db.SQL.QueryRow(stmt, chainType, txHash).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check if transaction is processed: %w", err)
	}

	return exists, nil
}

// MarkTransactionProcessed marks a transaction as processed
func (db *Db) MarkTransactionProcessed(chainType, txHash string, blockHeight int64) error {
	stmt := `INSERT INTO bootstrap_processed_transactions (chain_type, tx_hash, block_height)
             VALUES ($1, $2, $3)
             ON CONFLICT (chain_type, tx_hash) DO NOTHING`

	_, err := db.SQL.Exec(stmt, chainType, txHash, blockHeight)
	if err != nil {
		return fmt.Errorf("failed to mark transaction as processed: %w", err)
	}

	return nil
}

// CleanupOldProcessedTransactions removes old processed transaction records
func (db *Db) CleanupOldProcessedTransactions(chainType string, olderThanDays int) error {
	stmt := `DELETE FROM bootstrap_processed_transactions
             WHERE chain_type = $1 AND processed_at < NOW() - INTERVAL '%d days'`

	_, err := db.SQL.Exec(fmt.Sprintf(stmt, olderThanDays), chainType)
	if err != nil {
		return fmt.Errorf("failed to cleanup old processed transactions: %w", err)
	}

	return nil
}

// GetProcessedTransactionsByHeight retrieves processed transactions for a specific height range
func (db *Db) GetProcessedTransactionsByHeight(chainType string, fromHeight, toHeight int64) ([]types.ProcessedTransaction, error) {
	stmt := `SELECT chain_type, tx_hash, block_height, processed_at
             FROM bootstrap_processed_transactions
             WHERE chain_type = $1 AND block_height BETWEEN $2 AND $3
             ORDER BY block_height, processed_at`

	rows, err := db.SQL.Query(stmt, chainType, fromHeight, toHeight)
	if err != nil {
		return nil, fmt.Errorf("failed to get processed transactions: %w", err)
	}
	defer rows.Close()

	var transactions []types.ProcessedTransaction
	for rows.Next() {
		var tx types.ProcessedTransaction
		err := rows.Scan(
			&tx.ChainType,
			&tx.TxHash,
			&tx.BlockHeight,
			&tx.ProcessedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan processed transaction: %w", err)
		}
		transactions = append(transactions, tx)
	}

	return transactions, nil
}

// Address binding management functions

// GetAddressBindings retrieves all address bindings for a specific chain type
func (db *Db) GetAddressBindings(chainType string) ([]types.AddressBinding, error) {
	stmt := `SELECT chain_type, source_addr, target_addr, created_at, updated_at
             FROM bootstrap_address_bindings WHERE chain_type = $1
             ORDER BY created_at ASC`

	rows, err := db.SQL.Query(stmt, chainType)
	if err != nil {
		return nil, fmt.Errorf("failed to get address bindings for %s: %w", chainType, err)
	}
	defer rows.Close()

	var bindings []types.AddressBinding
	for rows.Next() {
		var binding types.AddressBinding
		err := rows.Scan(
			&binding.ChainType,
			&binding.SourceAddr,
			&binding.TargetAddr,
			&binding.CreatedAt,
			&binding.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan address binding: %w", err)
		}
		bindings = append(bindings, binding)
	}

	return bindings, nil
}

// SaveAddressBinding saves or updates an address binding
func (db *Db) SaveAddressBinding(binding *types.AddressBinding) error {
	stmt := `
INSERT INTO bootstrap_address_bindings (chain_type, source_addr, target_addr, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (chain_type, source_addr) DO UPDATE
SET target_addr = EXCLUDED.target_addr,
    updated_at  = EXCLUDED.updated_at`

	_, err := db.SQL.Exec(
		stmt,
		binding.ChainType,
		binding.SourceAddr,
		binding.TargetAddr,
		binding.CreatedAt,
		binding.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to save address binding: %w", err)
	}

	return nil
}

// GetAddressBinding retrieves a specific address binding
func (db *Db) GetAddressBinding(chainType, sourceAddr string) (*types.AddressBinding, error) {
	stmt := `SELECT chain_type, source_addr, target_addr, created_at, updated_at
             FROM bootstrap_address_bindings
             WHERE chain_type = $1 AND source_addr = $2`

	var binding types.AddressBinding
	err := db.SQL.QueryRow(stmt, chainType, sourceAddr).Scan(
		&binding.ChainType,
		&binding.SourceAddr,
		&binding.TargetAddr,
		&binding.CreatedAt,
		&binding.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found, not an error
		}
		return nil, fmt.Errorf("failed to get address binding for %s/%s: %w", chainType, sourceAddr, err)
	}

	return &binding, nil
}

// CheckTargetAddressBinding checks if a target address is already bound to a different source address
func (db *Db) CheckTargetAddressBinding(chainType, targetAddr, excludeSourceAddr string) (*types.AddressBinding, error) {
	stmt := `SELECT chain_type, source_addr, target_addr, created_at, updated_at
             FROM bootstrap_address_bindings
             WHERE chain_type = $1 AND target_addr = $2 AND source_addr != $3`

	var binding types.AddressBinding
	err := db.SQL.QueryRow(stmt, chainType, targetAddr, excludeSourceAddr).Scan(
		&binding.ChainType,
		&binding.SourceAddr,
		&binding.TargetAddr,
		&binding.CreatedAt,
		&binding.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found, not an error
		}
		return nil, fmt.Errorf("failed to check target address binding for %s/%s: %w", chainType, targetAddr, err)
	}

	return &binding, nil
}

// DeleteAddressBinding removes an address binding
func (db *Db) DeleteAddressBinding(chainType, sourceAddr string) error {
	stmt := `DELETE FROM bootstrap_address_bindings WHERE chain_type = $1 AND source_addr = $2`

	result, err := db.SQL.Exec(stmt, chainType, sourceAddr)
	if err != nil {
		return fmt.Errorf("failed to delete address binding: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check deletion result: %w", err)
	}

	if rows == 0 {
		return fmt.Errorf("address binding not found for %s/%s", chainType, sourceAddr)
	}

	return nil
}

// Transaction support methods

// WithTransaction executes a function within a database transaction
func (db *Db) WithTransaction(fn func(tx *sql.Tx) error) error {
	tx, err := db.SQL.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("failed to rollback transaction after error %v: %w", err, rollbackErr)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// MarkTransactionProcessedInTx marks a transaction as processed within an existing transaction
func (db *Db) MarkTransactionProcessedInTx(tx *sql.Tx, chainType, txHash string, blockHeight int64) error {
	stmt := `
INSERT INTO bootstrap_processed_transactions (chain_type, tx_hash, block_height, processed_at)
VALUES ($1, $2, $3, NOW())
ON CONFLICT (chain_type, tx_hash) DO NOTHING`

	_, err := tx.Exec(stmt, chainType, txHash, blockHeight)
	if err != nil {
		return fmt.Errorf("failed to mark transaction as processed: %w", err)
	}

	return nil
}

package database

import (
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

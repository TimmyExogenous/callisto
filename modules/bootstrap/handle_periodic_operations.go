package bootstrap

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/forbole/callisto/v4/types"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "bootstrap").Msg("setting up periodic tasks")

	// Schedule a cron job to run.
	if _, err := scheduler.Every(m.Config.UpdateInterval).Minutes().Do(func() {
		m.refetchBootstrapStates()
	}); err != nil {
		return fmt.Errorf("error while setting up daily refetch periodic operation: %s", err)
	}

	return nil
}

func (m *Module) refetchETHStates() error {
	log.Debug().Str("module", "bootstrap").Str("refetching", "ETH states").
		Msg("refetching ETH states")

	// refetch the validators
	validatorCount, err := m.bootstrapSession.GetValidatorsCount()
	if err != nil {
		return err
	}
	for i := int64(0); i < validatorCount.Int64(); i++ {
		validatorEthAddr, err := m.bootstrapSession.RegisteredValidators(big.NewInt(i))
		if err != nil {
			log.Err(err).Msg("call RegisteredValidators")
			continue
		}
		validatorIMAddr, err := m.bootstrapSession.EthToImAddress(validatorEthAddr)
		if err != nil {
			log.Err(err).Msg("call EthToImAddress")
			continue
		}
		validatorInfo, err := m.bootstrapSession.Validators(validatorIMAddr)
		if err != nil {
			log.Err(err).Msg("call Validators")
			continue
		}
		m.database.SaveBootstrapValidator(&types.BootstrapValidator{
			ValidatorEthAddress: validatorEthAddr.String(),
			ValidatorIMAddress:  validatorIMAddr,
			ValidatorName:       validatorInfo.Name,
			ConsensusPubKey:     hexutil.Encode(validatorInfo.ConsensusPublicKey[:]),
			Rate:                validatorInfo.Commission.Rate.String(),
			MaxRate:             validatorInfo.Commission.MaxRate.String(),
			MaxChangeRate:       validatorInfo.Commission.MaxChangeRate.String(),
			UpdatedAt:           time.Now(),
		})
	}
	//
	return nil
}

// Helper functions for safe data parsing and validation

// safeStringExtract safely extracts string value from map
func safeStringExtract(data map[string]interface{}, key string) (string, error) {
	value, exists := data[key]
	if !exists {
		return "", fmt.Errorf("missing key: %s", key)
	}
	str, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("invalid type for key %s: expected string, got %T", key, value)
	}
	return str, nil
}

// safeInt64Extract safely extracts int64 value from map
func safeInt64Extract(data map[string]interface{}, key string) (int64, error) {
	value, exists := data[key]
	if !exists {
		return 0, fmt.Errorf("missing key: %s", key)
	}
	switch v := value.(type) {
	case float64:
		return int64(v), nil
	case int64:
		return v, nil
	case int:
		return int64(v), nil
	default:
		return 0, fmt.Errorf("invalid type for key %s: expected number, got %T", key, value)
	}
}

// validateBTCAddressBinding validates and stores BTC address binding with database persistence
func (m *Module) validateBTCAddressBinding(senderAddr, imuachainAddr, txid string) bool {
	if senderAddr == "" || imuachainAddr == "" {
		log.Warn().Str("txid", txid).Msg("empty address in binding validation")
		return false
	}

	// Normalize addresses for consistent comparison
	senderAddr = strings.ToLower(strings.TrimSpace(senderAddr))
	imuachainAddr = strings.ToLower(strings.TrimSpace(imuachainAddr))

	m.btcMappingMutex.Lock()
	defer m.btcMappingMutex.Unlock()

	// Check existing binding in memory first
	if existing, exists := m.btcAddressMappings[senderAddr]; exists {
		if existing != imuachainAddr {
			log.Warn().Str("txid", txid).
				Str("bitcoin_addr", senderAddr).
				Str("existing_binding", existing).
				Str("new_binding", imuachainAddr).
				Msg("rejecting BTC transaction: address already bound to different imuachain address")
			return false
		}
		// Address already bound correctly, allow
		return true
	}

	// Check reverse binding in memory
	for btcAddr, imuaAddr := range m.btcAddressMappings {
		if imuaAddr == imuachainAddr && btcAddr != senderAddr {
			log.Warn().Str("txid", txid).
				Str("imuachain_addr", imuachainAddr).
				Str("existing_btc_binding", btcAddr).
				Str("new_btc_addr", senderAddr).
				Msg("rejecting BTC transaction: imuachain address already bound to different bitcoin address")
			return false
		}
	}

	// Double-check with database for consistency (in case memory was cleared)
	existingBinding, err := m.database.CheckTargetAddressBinding("BTC", imuachainAddr, senderAddr)
	if err != nil {
		log.Err(err).Str("txid", txid).Msg("error checking target address binding in database")
		return false
	}
	if existingBinding != nil {
		log.Warn().Str("txid", txid).
			Str("imuachain_addr", imuachainAddr).
			Str("existing_btc_binding", existingBinding.SourceAddr).
			Str("new_btc_addr", senderAddr).
			Msg("rejecting BTC transaction: imuachain address already bound to different bitcoin address in database")
		return false
	}

	// Establish new binding in both memory and database
	m.btcAddressMappings[senderAddr] = imuachainAddr

	// Save to database
	binding := &types.AddressBinding{
		ChainType:  "BTC",
		SourceAddr: senderAddr,
		TargetAddr: imuachainAddr,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := m.database.SaveAddressBinding(binding); err != nil {
		log.Err(err).Str("txid", txid).Msg("failed to save BTC address binding to database")
		// Remove from memory if database save failed
		delete(m.btcAddressMappings, senderAddr)
		return false
	}

	log.Info().Str("txid", txid).
		Str("bitcoin_addr", senderAddr).
		Str("imuachain_addr", imuachainAddr).
		Msg("established new BTC address binding")
	return true
}

// validateXRPAddressBinding validates and stores XRP address binding with database persistence
func (m *Module) validateXRPAddressBinding(senderAddr, imuachainAddr, txHash string) bool {
	if senderAddr == "" || imuachainAddr == "" {
		log.Warn().Str("hash", txHash).Msg("empty address in binding validation")
		return false
	}

	// Normalize addresses for consistent comparison
	senderAddr = strings.ToLower(strings.TrimSpace(senderAddr))
	imuachainAddr = strings.ToLower(strings.TrimSpace(imuachainAddr))

	m.xrpMappingMutex.Lock()
	defer m.xrpMappingMutex.Unlock()

	// Check existing binding in memory first
	if existing, exists := m.xrpAddressMappings[senderAddr]; exists {
		if existing != imuachainAddr {
			log.Warn().Str("hash", txHash).
				Str("xrp_addr", senderAddr).
				Str("existing_binding", existing).
				Str("new_binding", imuachainAddr).
				Msg("rejecting XRP transaction: address already bound to different imuachain address")
			return false
		}
		// Address already bound correctly, allow
		return true
	}

	// Check reverse binding in memory
	for xrpAddr, imuaAddr := range m.xrpAddressMappings {
		if imuaAddr == imuachainAddr && xrpAddr != senderAddr {
			log.Warn().Str("hash", txHash).
				Str("imuachain_addr", imuachainAddr).
				Str("existing_xrp_binding", xrpAddr).
				Str("new_xrp_addr", senderAddr).
				Msg("rejecting XRP transaction: imuachain address already bound to different XRP address")
			return false
		}
	}

	// Double-check with database for consistency (in case memory was cleared)
	existingBinding, err := m.database.CheckTargetAddressBinding("XRP", imuachainAddr, senderAddr)
	if err != nil {
		log.Err(err).Str("hash", txHash).Msg("error checking target address binding in database")
		return false
	}
	if existingBinding != nil {
		log.Warn().Str("hash", txHash).
			Str("imuachain_addr", imuachainAddr).
			Str("existing_xrp_binding", existingBinding.SourceAddr).
			Str("new_xrp_addr", senderAddr).
			Msg("rejecting XRP transaction: imuachain address already bound to different XRP address in database")
		return false
	}

	// Establish new binding in both memory and database
	m.xrpAddressMappings[senderAddr] = imuachainAddr

	// Save to database
	binding := &types.AddressBinding{
		ChainType:  "XRP",
		SourceAddr: senderAddr,
		TargetAddr: imuachainAddr,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := m.database.SaveAddressBinding(binding); err != nil {
		log.Err(err).Str("hash", txHash).Msg("failed to save XRP address binding to database")
		// Remove from memory if database save failed
		delete(m.xrpAddressMappings, senderAddr)
		return false
	}

	log.Info().Str("hash", txHash).
		Str("xrp_addr", senderAddr).
		Str("imuachain_addr", imuachainAddr).
		Msg("established new XRP address binding")
	return true
}

// processTransactionWithRetry executes a transaction function with retry logic for database errors
func (m *Module) processTransactionWithRetry(chainType string, txFunc func() error) error {
	const maxRetries = 3

	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		err := txFunc()
		if err == nil {
			return nil // Success
		}

		lastErr = err

		// Check if error is retryable
		if !isRetryableDatabaseError(err) {
			// Non-retryable error, fail immediately
			return err
		}

		if attempt < maxRetries-1 {
			// Exponential backoff: 1s, 2s, 4s...
			backoff := time.Duration(1<<uint(attempt)) * time.Second
			log.Warn().
				Err(err).
				Str("chain", chainType).
				Int("attempt", attempt+1).
				Dur("backoff", backoff).
				Msg("database transaction failed, retrying")

			time.Sleep(backoff)
		}
	}

	return fmt.Errorf("transaction failed after %d retries: %w", maxRetries, lastErr)
}

// isRetryableDatabaseError checks if a database error should be retried
func isRetryableDatabaseError(err error) bool {
	if err == nil {
		return false
	}

	errStr := strings.ToLower(err.Error())

	// Deadlock errors
	if strings.Contains(errStr, "deadlock") ||
		strings.Contains(errStr, "lock wait timeout") {
		return true
	}

	// Connection errors
	if strings.Contains(errStr, "connection refused") ||
		strings.Contains(errStr, "connection reset") ||
		strings.Contains(errStr, "connection lost") ||
		strings.Contains(errStr, "broken pipe") {
		return true
	}

	// Serialization failures (PostgreSQL)
	if strings.Contains(errStr, "serialization failure") ||
		strings.Contains(errStr, "could not serialize") {
		return true
	}

	// Lock acquisition timeouts
	if strings.Contains(errStr, "lock acquisition") ||
		strings.Contains(errStr, "timeout") {
		return true
	}

	return false
}

// sortBTCTransactions sorts BTC transactions by block height and transaction index
func sortBTCTransactions(txs []types.BTCTx) {
	sort.Slice(txs, func(i, j int) bool {
		if txs[i].BlockHeight != txs[j].BlockHeight {
			return txs[i].BlockHeight < txs[j].BlockHeight
		}
		// Use TxID for deterministic order if same block
		return txs[i].TxID < txs[j].TxID
	})
}

// sortXRPTransactions sorts XRP transactions by ledger index and transaction index
func sortXRPTransactions(txs []types.XRPTransaction) {
	sort.Slice(txs, func(i, j int) bool {
		if txs[i].LedgerIndex != txs[j].LedgerIndex {
			return txs[i].LedgerIndex < txs[j].LedgerIndex
		}
		return txs[i].Meta.TransactionIndex < txs[j].Meta.TransactionIndex
	})
}

// BTC vault transaction fetching and processing

func (m *Module) refetchBTCStates() error {
	log.Debug().Str("module", "bootstrap").Str("refetching", "BTC states").
		Msg("refetching BTC states")

	if m.Config.BTCVaultAddr == "" {
		log.Debug().Str("module", "bootstrap").Msg("BTC vault address not configured, skipping")
		return nil
	}

	// Get current block height
	currentHeight, err := m.getBTCCurrentBlockHeight()
	if err != nil {
		return fmt.Errorf("error getting current BTC block height: %s", err)
	}

	// Get or initialize scan state
	scanState, err := m.database.GetScanState("BTC")
	if err != nil {
		// First time scanning, initialize with start height
		startHeight := m.Config.BTCStartHeight
		if startHeight == 0 {
			// Start from current height minus confirmation depth for safety
			startHeight = currentHeight - int64(m.Config.BTCMinConfirmations)
			if startHeight < 0 {
				startHeight = 0
			}
		}

		scanState = &types.ScanState{
			ChainType:  "BTC",
			LastHeight: startHeight,
			SafeHeight: startHeight,
			UpdatedAt:  time.Now(),
		}

		log.Info().Int64("start_height", startHeight).Msg("initializing BTC scan state")
		if err := m.database.UpdateScanState(scanState); err != nil {
			return fmt.Errorf("failed to initialize BTC scan state: %s", err)
		}
	}

	// Calculate safe current height (with confirmation buffer)
	safeCurrentHeight := currentHeight - int64(m.Config.BTCMinConfirmations)
	if safeCurrentHeight <= scanState.SafeHeight {
		log.Debug().Int64("safe_height", safeCurrentHeight).Int64("last_safe", scanState.SafeHeight).
			Msg("no new confirmed BTC blocks to scan")
		return nil
	}

	log.Info().Int64("from_height", scanState.SafeHeight).Int64("to_height", safeCurrentHeight).
		Msg("starting incremental BTC scan")

	// Get transactions for height range (incremental scan)
	transactions, err := m.getBTCVaultTransactionsFromHeight(scanState.SafeHeight+1, safeCurrentHeight)
	if err != nil {
		return fmt.Errorf("error getting BTC vault transactions from height %d to %d: %s",
			scanState.SafeHeight+1, safeCurrentHeight, err)
	}

	// Sort transactions by block height and order for consistent processing
	sortBTCTransactions(transactions)

	processedCount := 0
	duplicateCount := 0

	// Process transactions with address binding validation and deduplication
	for _, tx := range transactions {
		// Check if transaction already processed
		if processed, err := m.database.IsTransactionProcessed("BTC", tx.TxID); err != nil {
			log.Err(err).Str("txid", tx.TxID).Msg("error checking if BTC transaction is processed")
			continue
		} else if processed {
			duplicateCount++
			log.Debug().Str("txid", tx.TxID).Msg("BTC transaction already processed, skipping")
			continue
		}

		if err := m.processBTCTxWithTransaction(tx, currentHeight); err != nil {
			log.Err(err).Str("txid", tx.TxID).Msg("error processing BTC transaction")
			continue
		}

		processedCount++
	}

	// Log processing statistics
	totalTxs := len(transactions)
	duplicateRate := float64(duplicateCount) / float64(totalTxs) * 100
	log.Info().
		Int("total_txs", totalTxs).
		Int("processed", processedCount).
		Int("duplicates", duplicateCount).
		Float64("duplicate_rate_percent", duplicateRate).
		Int64("from_height", scanState.SafeHeight).
		Int64("to_height", safeCurrentHeight).
		Msg("BTC transaction processing statistics")

	// Update scan state
	newScanState := &types.ScanState{
		ChainType:  "BTC",
		LastHeight: currentHeight,
		SafeHeight: safeCurrentHeight,
		UpdatedAt:  time.Now(),
	}

	if err := m.database.UpdateScanState(newScanState); err != nil {
		return fmt.Errorf("failed to update BTC scan state: %s", err)
	}

	log.Info().Int("processed_count", processedCount).
		Int64("from_height", scanState.SafeHeight).
		Int64("to_height", safeCurrentHeight).
		Msg("completed incremental BTC scan")

	return nil
}

// getBTCCurrentBlockHeight gets the current BTC block height from Esplora API
func (m *Module) getBTCCurrentBlockHeight() (int64, error) {
	url := fmt.Sprintf("%s/api/blocks/tip/height", m.Config.BTCRPC)

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Create request with context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch block height: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var height int64
	if err := json.NewDecoder(resp.Body).Decode(&height); err != nil {
		return 0, fmt.Errorf("failed to decode response: %s", err)
	}

	if height <= 0 {
		return 0, fmt.Errorf("invalid block height: %d", height)
	}

	return height, nil
}

// getBTCVaultTransactionsFromHeight gets BTC vault transactions from a specific height range
func (m *Module) getBTCVaultTransactionsFromHeight(fromHeight, toHeight int64) ([]types.BTCTx, error) {
	var allTxs []types.BTCTx

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Process blocks in batches to avoid overwhelming the API
	batchSize := int64(m.Config.ScanBatchSize)
	if batchSize <= 0 {
		batchSize = 10 // Default batch size
	}

	for currentHeight := fromHeight; currentHeight <= toHeight; currentHeight += batchSize {
		endHeight := currentHeight + batchSize - 1
		if endHeight > toHeight {
			endHeight = toHeight
		}

		log.Debug().Int64("from", currentHeight).Int64("to", endHeight).
			Msg("scanning BTC blocks for vault transactions")

		// Get transactions for each block in the range
		for blockHeight := currentHeight; blockHeight <= endHeight; blockHeight++ {
			blockTxs, err := m.getBTCBlockTransactions(client, blockHeight)
			if err != nil {
				log.Err(err).Int64("height", blockHeight).Msg("error getting BTC block transactions")
				continue
			}

			// Filter transactions involving vault address
			for _, tx := range blockTxs {
				if m.isBTCVaultTransaction(tx) {
					allTxs = append(allTxs, tx)
				}
			}
		}

		// Add delay between batches to be API-friendly
		if endHeight < toHeight {
			time.Sleep(100 * time.Millisecond)
		}
	}

	log.Info().Int("total_txs", len(allTxs)).
		Int64("from_height", fromHeight).
		Int64("to_height", toHeight).
		Msg("completed BTC vault transaction scanning")

	return allTxs, nil
}

// getBTCBlockTransactions gets all transactions for a specific block height
func (m *Module) getBTCBlockTransactions(client *http.Client, blockHeight int64) ([]types.BTCTx, error) {
	// First get block hash
	url := fmt.Sprintf("%s/api/block-height/%d", m.Config.BTCRPC, blockHeight)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create block hash request: %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get block hash: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d for block height %d", resp.StatusCode, blockHeight)
	}

	var blockHash string
	if err := json.NewDecoder(resp.Body).Decode(&blockHash); err != nil {
		return nil, fmt.Errorf("failed to decode block hash: %s", err)
	}

	// Then get block transactions
	url = fmt.Sprintf("%s/api/block/%s/txs", m.Config.BTCRPC, blockHash)

	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req, err = http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create block txs request: %s", err)
	}

	resp, err = client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get block transactions: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d for block %s", resp.StatusCode, blockHash)
	}

	var txs []types.BTCTx
	if err := json.NewDecoder(resp.Body).Decode(&txs); err != nil {
		return nil, fmt.Errorf("failed to decode block transactions: %s", err)
	}

	return txs, nil
}

// isBTCVaultTransaction checks if a transaction involves the vault address
func (m *Module) isBTCVaultTransaction(tx types.BTCTx) bool {
	vaultAddr := strings.ToLower(strings.TrimSpace(m.Config.BTCVaultAddr))

	// Check if any output goes to vault address
	for _, vout := range tx.Vout {
		if strings.ToLower(strings.TrimSpace(vout.ScriptPubKeyAddr)) == vaultAddr {
			return true
		}
	}

	// Check if any input comes from vault address
	for _, vin := range tx.Vin {
		if strings.ToLower(strings.TrimSpace(vin.Prevout.ScriptPubKeyAddr)) == vaultAddr {
			return true
		}
	}

	return false
}

// processBTCTxWithTransaction processes a single BTC transaction within a database transaction
// Note: Duplicate processing check is already done at the caller level, but we keep
// the database-level atomicity guarantee with ON CONFLICT DO NOTHING
func (m *Module) processBTCTxWithTransaction(tx types.BTCTx, currentHeight int64) error {
	return m.processTransactionWithRetry("BTC", func() error {
		return m.database.WithTransaction(func(dbTx *sql.Tx) error {
			// 1. Validate transaction first (no database writes)
			// Check confirmations
			if !tx.Confirmed || tx.BlockHeight <= 0 {
				return nil // Skip unconfirmed transactions
			}

			confirmations := currentHeight - tx.BlockHeight + 1
			if confirmations < int64(m.Config.BTCMinConfirmations) {
				return nil // Not enough confirmations
			}

			// Validate transaction
			isValid, opReturnData, err := m.validateBTCTx(tx)
			if err != nil {
				return fmt.Errorf("error validating BTC transaction: %w", err)
			}

			if !isValid {
				return nil // Invalid transaction, skip silently
			}

			// 2. Save business data
			if err := m.saveBTCTransaction(tx, opReturnData); err != nil {
				return fmt.Errorf("failed to save BTC transaction data: %w", err)
			}

			// 3. Mark as processed last (atomicity guarantee)
			// Note: ON CONFLICT DO NOTHING in the database handles concurrent processing
			if err := m.database.MarkTransactionProcessedInTx(dbTx, "BTC", tx.TxID, tx.BlockHeight); err != nil {
				return fmt.Errorf("failed to mark BTC transaction as processed: %w", err)
			}

			return nil
		})
	})
}

// processBTCTx processes a single BTC transaction
func (m *Module) processBTCTx(tx types.BTCTx, currentHeight int64) error {
	// Check confirmations
	if !tx.Confirmed || tx.BlockHeight <= 0 {
		return nil // Skip unconfirmed transactions
	}

	confirmations := currentHeight - tx.BlockHeight + 1
	if confirmations < int64(m.Config.BTCMinConfirmations) {
		return nil // Not enough confirmations
	}

	// Validate transaction
	isValid, opReturnData, err := m.validateBTCTx(tx)
	if err != nil {
		return fmt.Errorf("error validating BTC transaction: %s", err)
	}

	if !isValid {
		return nil // Invalid transaction
	}

	// Save transaction data to database
	return m.saveBTCTransaction(tx, opReturnData)
}

// validateBTCTx validates a BTC transaction for bootstrap deposits
func (m *Module) validateBTCTx(tx types.BTCTx) (bool, *BTCOPReturnData, error) {
	// Check if it's from vault (should not be)
	for _, vin := range tx.Vin {
		if strings.ToLower(strings.TrimSpace(vin.Prevout.ScriptPubKeyAddr)) ==
			strings.ToLower(strings.TrimSpace(m.Config.BTCVaultAddr)) {
			return false, nil, nil // From vault, invalid
		}
	}

	// Find vault output
	var vaultOutput *types.BTCVout
	for _, vout := range tx.Vout {
		if strings.ToLower(strings.TrimSpace(vout.ScriptPubKeyAddr)) ==
			strings.ToLower(strings.TrimSpace(m.Config.BTCVaultAddr)) &&
			vout.Value >= m.Config.BTCMinAmount {
			vaultOutput = &vout
			break
		}
	}

	if vaultOutput == nil {
		return false, nil, nil // No valid vault output
	}

	// Find OP_RETURN output
	var opReturnOutput *types.BTCVout
	for _, vout := range tx.Vout {
		if vout.ScriptPubKeyType == "op_return" {
			opReturnOutput = &vout
			break
		}
	}

	if opReturnOutput == nil {
		return false, nil, nil // No OP_RETURN output
	}

	// Parse OP_RETURN data
	opReturnData, err := parseBTCOPReturn(opReturnOutput.ScriptPubKey)
	if err != nil {
		log.Err(err).Str("txid", tx.TxID).Msg("failed to parse OP_RETURN")
		return false, nil, nil
	}

	// Validate validator address
	if !m.isValidValidatorAddress(opReturnData.ValidatorAddress) {
		log.Warn().Str("txid", tx.TxID).Str("validator", opReturnData.ValidatorAddress).
			Msg("invalid validator address")
		return false, nil, nil
	}

	// Check if validator is registered
	isRegistered, err := m.isValidatorRegistered(opReturnData.ValidatorAddress)
	if err != nil {
		return false, nil, fmt.Errorf("error checking validator registration: %s", err)
	}

	if !isRegistered {
		log.Warn().Str("txid", tx.TxID).Str("validator", opReturnData.ValidatorAddress).
			Msg("validator not registered")
		return false, nil, nil
	}

	// Validate address binding (1-1 mapping between Bitcoin and Imuachain addresses)
	senderAddr := ""
	for _, vin := range tx.Vin {
		if vin.Prevout.ScriptPubKeyAddr != "" {
			senderAddr = vin.Prevout.ScriptPubKeyAddr
			break
		}
	}
	if senderAddr == "" {
		log.Warn().Str("txid", tx.TxID).Msg("no sender address found")
		return false, nil, nil
	}

	if !m.validateBTCAddressBinding(senderAddr, opReturnData.ImuachainAddress, tx.TxID) {
		return false, nil, nil
	}

	return true, opReturnData, nil
}

// BTCOPReturnData represents parsed OP_RETURN data
type BTCOPReturnData struct {
	ImuachainAddress string
	ValidatorAddress string
}

// parseBTCOPReturn parses OP_RETURN data from BTC transaction
func parseBTCOPReturn(scriptPubKey string) (*BTCOPReturnData, error) {
	// Validate OP_RETURN format
	if !strings.HasPrefix(scriptPubKey, "6a3d") {
		return nil, fmt.Errorf("invalid OP_RETURN prefix")
	}

	hexData := scriptPubKey[4:]
	if len(hexData) != 122 {
		return nil, fmt.Errorf("invalid OP_RETURN data length: expected 122, got %d", len(hexData))
	}

	// Extract imuachain address (first 40 hex chars = 20 bytes)
	imuachainHex := "0x" + hexData[:40]
	if !common.IsHexAddress(imuachainHex) {
		return nil, fmt.Errorf("invalid imuachain address format: %s", imuachainHex)
	}

	// Extract validator address (remaining 41 bytes as UTF-8 string)
	validatorHex := hexData[40:]
	validatorBytes, err := hex.DecodeString(validatorHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode validator address: %s", err)
	}

	validatorAddress := string(validatorBytes)
	// Remove null bytes and validate length
	validatorAddress = strings.TrimRight(validatorAddress, "\x00")
	if len(validatorAddress) != 41 {
		return nil, fmt.Errorf("invalid validator address length: expected 41, got %d", len(validatorAddress))
	}

	// Additional validation: check for valid bech32 characters and prefix
	if !strings.HasPrefix(validatorAddress, "im") {
		return nil, fmt.Errorf("invalid validator address prefix: expected 'im', got %s", validatorAddress[:2])
	}

	// Validate bech32 character set
	for _, char := range validatorAddress {
		if !((char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')) {
			return nil, fmt.Errorf("invalid character in validator address: %c", char)
		}
	}

	return &BTCOPReturnData{
		ImuachainAddress: strings.ToLower(strings.TrimSpace(imuachainHex)),
		ValidatorAddress: strings.TrimSpace(validatorAddress),
	}, nil
}

// isValidValidatorAddress validates validator address format (bech32 with 'im' prefix)
func (m *Module) isValidValidatorAddress(address string) bool {
	// Simple validation - should start with 'im' and be 41 characters
	return len(address) == 41 && strings.HasPrefix(address, "im")
}

// isValidatorRegistered checks if validator is registered in bootstrap contract
func (m *Module) isValidatorRegistered(validatorAddr string) (bool, error) {
	validatorInfo, err := m.bootstrapSession.Validators(validatorAddr)
	if err != nil {
		return false, err
	}
	return validatorInfo.Name != "", nil
}

// saveBTCTransaction saves BTC transaction data to database
func (m *Module) saveBTCTransaction(tx types.BTCTx, opReturnData *BTCOPReturnData) error {
	// Find vault output
	var vaultOutput *types.BTCVout
	for _, vout := range tx.Vout {
		if strings.ToLower(strings.TrimSpace(vout.ScriptPubKeyAddr)) ==
			strings.ToLower(strings.TrimSpace(m.Config.BTCVaultAddr)) {
			vaultOutput = &vout
			break
		}
	}

	if vaultOutput == nil {
		return fmt.Errorf("vault output not found")
	}

	// Create staker ID
	stakerID := opReturnData.ImuachainAddress + "_0x1" // BTC chain ID = 1

	// Save staker asset
	stakerAsset := &types.BootstrapStakerAsset{
		StakerID:     stakerID,
		AssetID:      VirtualAddress + "_0x1", // BTC asset ID: virtualAddress + chainID
		Deposited:    strconv.FormatInt(vaultOutput.Value, 10),
		Withdrawable: "0", // All stakes must be delegated
		Delegated:    strconv.FormatInt(vaultOutput.Value, 10),
		UpdatedAt:    time.Now(),
	}

	if err := m.database.SaveBootstrapStakerAsset(stakerAsset); err != nil {
		return fmt.Errorf("failed to save staker asset: %s", err)
	}

	// Save delegation state
	delegationState := &types.BootstrapDelegationState{
		StakerID:     stakerID,
		AssetID:      VirtualAddress + "_0x1",
		OperatorAddr: opReturnData.ValidatorAddress,
		Delegated:    strconv.FormatInt(vaultOutput.Value, 10),
		UpdatedAt:    time.Now(),
	}

	if err := m.database.SaveBootstrapDelegationState(delegationState); err != nil {
		return fmt.Errorf("failed to save delegation state: %s", err)
	}

	log.Info().
		Str("txid", tx.TxID).
		Str("staker", stakerID).
		Str("validator", opReturnData.ValidatorAddress).
		Int64("amount", vaultOutput.Value).
		Msg("processed BTC bootstrap transaction")

	return nil
}

// XRP vault transaction fetching and processing
func (m *Module) refetchXRPStates() error {
	log.Debug().Str("module", "bootstrap").Str("refetching", "XRP states").
		Msg("refetching XRP states")

	if m.Config.XRPVaultAddr == "" {
		log.Debug().Str("module", "bootstrap").Msg("XRP vault address not configured, skipping")
		return nil
	}

	// Get current ledger index
	currentLedger, err := m.getXRPCurrentLedger()
	if err != nil {
		return fmt.Errorf("error getting current XRP ledger: %s", err)
	}

	// Get or initialize scan state
	scanState, err := m.database.GetScanState("XRP")
	if err != nil {
		// First time scanning, initialize with start ledger
		startLedger := m.Config.XRPStartLedger
		if startLedger == 0 {
			// Start from current ledger minus confirmation depth for safety
			startLedger = currentLedger - int64(m.Config.XRPMinConfirmations)
			if startLedger < 0 {
				startLedger = 0
			}
		}

		scanState = &types.ScanState{
			ChainType:  "XRP",
			LastHeight: startLedger,
			SafeHeight: startLedger,
			UpdatedAt:  time.Now(),
		}

		log.Info().Int64("start_ledger", startLedger).Msg("initializing XRP scan state")
		if err := m.database.UpdateScanState(scanState); err != nil {
			return fmt.Errorf("failed to initialize XRP scan state: %s", err)
		}
	}

	// Calculate safe current ledger (with confirmation buffer)
	safeCurrentLedger := currentLedger - int64(m.Config.XRPMinConfirmations)
	if safeCurrentLedger <= scanState.SafeHeight {
		log.Debug().Int64("safe_ledger", safeCurrentLedger).Int64("last_safe", scanState.SafeHeight).
			Msg("no new confirmed XRP ledgers to scan")
		return nil
	}

	log.Info().Int64("from_ledger", scanState.SafeHeight).Int64("to_ledger", safeCurrentLedger).
		Msg("starting incremental XRP scan")

	// Get transactions for ledger range (incremental scan)
	transactions, err := m.getXRPVaultTransactionsFromLedger(scanState.SafeHeight+1, safeCurrentLedger)
	if err != nil {
		return fmt.Errorf("error getting XRP vault transactions from ledger %d to %d: %s",
			scanState.SafeHeight+1, safeCurrentLedger, err)
	}

	// Sort transactions by ledger index and order for consistent processing
	sortXRPTransactions(transactions)

	processedCount := 0
	duplicateCount := 0

	// Process transactions with address binding validation and deduplication
	for _, tx := range transactions {
		// Check if transaction already processed
		if processed, err := m.database.IsTransactionProcessed("XRP", tx.Hash); err != nil {
			log.Err(err).Str("hash", tx.Hash).Msg("error checking if XRP transaction is processed")
			continue
		} else if processed {
			duplicateCount++
			log.Debug().Str("hash", tx.Hash).Msg("XRP transaction already processed, skipping")
			continue
		}

		if err := m.processXRPTxWithTransaction(tx, currentLedger); err != nil {
			log.Err(err).Str("hash", tx.Hash).Msg("error processing XRP transaction")
			continue
		}

		processedCount++
	}

	// Log processing statistics
	totalTxs := len(transactions)
	duplicateRate := float64(duplicateCount) / float64(totalTxs) * 100
	log.Info().
		Int("total_txs", totalTxs).
		Int("processed", processedCount).
		Int("duplicates", duplicateCount).
		Float64("duplicate_rate_percent", duplicateRate).
		Int64("from_ledger", scanState.SafeHeight).
		Int64("to_ledger", safeCurrentLedger).
		Msg("XRP transaction processing statistics")

	// Update scan state
	newScanState := &types.ScanState{
		ChainType:  "XRP",
		LastHeight: currentLedger,
		SafeHeight: safeCurrentLedger,
		UpdatedAt:  time.Now(),
	}

	if err := m.database.UpdateScanState(newScanState); err != nil {
		return fmt.Errorf("failed to update XRP scan state: %s", err)
	}

	log.Info().Int("processed_count", processedCount).
		Int64("from_ledger", scanState.SafeHeight).
		Int64("to_ledger", safeCurrentLedger).
		Msg("completed incremental XRP scan")

	return nil
}

// getXRPCurrentLedger gets the current XRP ledger index
func (m *Module) getXRPCurrentLedger() (int64, error) {
	requestBody := map[string]interface{}{
		"method": "ledger",
		"params": []map[string]interface{}{
			{
				"ledger_index": "validated",
			},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return 0, err
	}

	resp, err := http.Post(m.Config.XRPRPC, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, err
	}

	result, ok := response["result"].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("invalid response format")
	}

	ledger, ok := result["ledger"].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("invalid ledger format")
	}

	ledgerIndex, ok := ledger["ledger_index"]
	if !ok {
		return 0, fmt.Errorf("ledger_index not found")
	}

	switch v := ledgerIndex.(type) {
	case float64:
		return int64(v), nil
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	default:
		return 0, fmt.Errorf("invalid ledger_index type")
	}
}

// parseXRPTransaction safely parses XRP transaction data from JSON map
func parseXRPTransaction(txData map[string]interface{}) (*types.XRPTransaction, error) {
	// Extract required fields with safe type assertions
	hash, err := safeStringExtract(txData, "hash")
	if err != nil {
		return nil, fmt.Errorf("failed to extract hash: %s", err)
	}

	ledgerIndexFloat, ok := txData["ledger_index"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid ledger_index type")
	}
	ledgerIndex := int64(ledgerIndexFloat)

	dateFloat, ok := txData["date"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid date type")
	}
	date := int64(dateFloat)

	validated, _ := txData["validated"].(bool)
	if !validated {
		return nil, fmt.Errorf("transaction not validated")
	}

	// Parse tx_json
	txJSON, ok := txData["tx_json"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid tx_json type")
	}

	// Extract transaction type and account
	transactionType, err := safeStringExtract(txJSON, "TransactionType")
	if err != nil {
		return nil, fmt.Errorf("failed to extract TransactionType: %s", err)
	}

	account, err := safeStringExtract(txJSON, "Account")
	if err != nil {
		return nil, fmt.Errorf("failed to extract Account: %s", err)
	}

	// Validate account format (XRP addresses should be r... format)
	if !strings.HasPrefix(account, "r") || len(account) < 25 || len(account) > 35 {
		return nil, fmt.Errorf("invalid XRP account format: %s", account)
	}

	// Parse meta data
	metaData, ok := txData["meta"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid meta type")
	}

	transactionResult, err := safeStringExtract(metaData, "TransactionResult")
	if err != nil {
		return nil, fmt.Errorf("failed to extract TransactionResult: %s", err)
	}

	transactionIndexFloat, ok := metaData["TransactionIndex"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid TransactionIndex type")
	}
	transactionIndex := int64(transactionIndexFloat)

	// Create transaction struct
	tx := &types.XRPTransaction{
		Hash:        hash,
		LedgerIndex: ledgerIndex,
		Date:        date,
		Validated:   validated,
		Tx: types.XRPTx{
			TransactionType: transactionType,
			Account:         account,
		},
		Meta: types.XRPMeta{
			TransactionResult: transactionResult,
			TransactionIndex:  transactionIndex,
		},
	}

	// Extract optional fields safely
	if destination, ok := txJSON["Destination"].(string); ok {
		tx.Tx.Destination = destination
	}

	if amount, ok := txJSON["Amount"]; ok {
		tx.Tx.Amount = amount
	}

	if fee, ok := txJSON["Fee"].(string); ok {
		tx.Tx.Fee = fee
	}

	if destTag, ok := txJSON["DestinationTag"].(float64); ok {
		tx.Tx.DestinationTag = int64(destTag)
	}

	// Parse memos safely
	if memosInterface, ok := txJSON["Memos"].([]interface{}); ok {
		memos := make([]types.XRPMemo, 0, len(memosInterface))
		for _, memoInterface := range memosInterface {
			memoData, ok := memoInterface.(map[string]interface{})
			if !ok {
				continue
			}
			memo, ok := memoData["Memo"].(map[string]interface{})
			if !ok {
				continue
			}

			xrpMemo := types.XRPMemo{}
			if memoType, ok := memo["MemoType"].(string); ok {
				xrpMemo.Memo.MemoType = memoType
			}
			if memoDataStr, ok := memo["MemoData"].(string); ok {
				xrpMemo.Memo.MemoData = memoDataStr
			}
			if memoFormat, ok := memo["MemoFormat"].(string); ok {
				xrpMemo.Memo.MemoFormat = memoFormat
			}

			memos = append(memos, xrpMemo)
		}
		tx.Tx.Memos = memos
	}

	return tx, nil
}

// processXRPTxWithTransaction processes a single XRP transaction within a database transaction
// Note: Duplicate processing check is already done at the caller level, but we keep
// the database-level atomicity guarantee with ON CONFLICT DO NOTHING
func (m *Module) processXRPTxWithTransaction(tx types.XRPTransaction, currentLedger int64) error {
	return m.processTransactionWithRetry("XRP", func() error {
		return m.database.WithTransaction(func(dbTx *sql.Tx) error {
			// 1. Validate transaction first (no database writes)
			// Check confirmations
			if !tx.Validated || tx.LedgerIndex <= 0 {
				return nil // Skip unvalidated transactions
			}

			confirmations := currentLedger - tx.LedgerIndex + 1
			if confirmations < int64(m.Config.XRPMinConfirmations) {
				return nil // Not enough confirmations
			}

			// Validate transaction
			isValid, memoData, err := m.validateXRPTx(tx)
			if err != nil {
				return fmt.Errorf("error validating XRP transaction: %w", err)
			}

			if !isValid {
				return nil // Invalid transaction, skip silently
			}

			// 2. Save business data
			if err := m.saveXRPTransaction(tx, memoData); err != nil {
				return fmt.Errorf("failed to save XRP transaction data: %w", err)
			}

			// 3. Mark as processed last (atomicity guarantee)
			// Note: ON CONFLICT DO NOTHING in the database handles concurrent processing
			if err := m.database.MarkTransactionProcessedInTx(dbTx, "XRP", tx.Hash, tx.LedgerIndex); err != nil {
				return fmt.Errorf("failed to mark XRP transaction as processed: %w", err)
			}

			return nil
		})
	})
}

// processXRPTx processes a single XRP transaction
func (m *Module) processXRPTx(tx types.XRPTransaction, currentLedger int64) error {
	// Check confirmations
	if !tx.Validated || tx.LedgerIndex <= 0 {
		return nil // Skip unvalidated transactions
	}

	confirmations := currentLedger - tx.LedgerIndex + 1
	if confirmations < int64(m.Config.XRPMinConfirmations) {
		return nil // Not enough confirmations
	}

	// Validate transaction
	isValid, memoData, err := m.validateXRPTx(tx)
	if err != nil {
		return fmt.Errorf("error validating XRP transaction: %s", err)
	}

	if !isValid {
		return nil // Invalid transaction
	}

	// Save transaction data to database
	return m.saveXRPTransaction(tx, memoData)
}

// validateXRPTx validates a XRP transaction for bootstrap deposits
func (m *Module) validateXRPTx(tx types.XRPTransaction) (bool, *XRPMemoData, error) {
	// Must be Payment transaction
	if tx.Tx.TransactionType != "Payment" {
		return false, nil, nil
	}

	// Must be successful
	if tx.Meta.TransactionResult != "tesSUCCESS" {
		return false, nil, nil
	}

	// Must be sent to our vault address
	if tx.Tx.Destination != m.Config.XRPVaultAddr {
		return false, nil, nil
	}

	// Must not be from vault address (no self-transfers)
	if tx.Tx.Account == m.Config.XRPVaultAddr {
		return false, nil, nil
	}

	// Must be XRP payment (not token)
	amountStr, ok := tx.Tx.Amount.(string)
	if !ok {
		return false, nil, nil // Token payment
	}

	// Check minimum amount
	amount, err := strconv.ParseInt(amountStr, 10, 64)
	if err != nil {
		return false, nil, fmt.Errorf("invalid amount format: %s", err)
	}

	if amount < m.Config.XRPMinAmount {
		return false, nil, nil
	}

	// Check DestinationTag (configurable)
	if tx.Tx.DestinationTag != m.Config.XRPDestinationTag {
		log.Debug().Int64("expected", m.Config.XRPDestinationTag).Int64("actual", tx.Tx.DestinationTag).
			Str("hash", tx.Hash).Msg("invalid destination tag")
		return false, nil, nil
	}

	// Must have memo with validator info
	if tx.Tx.Memos == nil || len(tx.Tx.Memos) == 0 {
		return false, nil, nil
	}

	// Parse memo data
	memoData, err := parseXRPMemo(tx.Tx.Memos)
	if err != nil {
		log.Err(err).Str("hash", tx.Hash).Msg("failed to parse XRP memo")
		return false, nil, nil
	}

	// Validate validator address
	if !m.isValidValidatorAddress(memoData.ValidatorAddress) {
		log.Warn().Str("hash", tx.Hash).Str("validator", memoData.ValidatorAddress).
			Msg("invalid validator address")
		return false, nil, nil
	}

	// Check if validator is registered
	isRegistered, err := m.isValidatorRegistered(memoData.ValidatorAddress)
	if err != nil {
		return false, nil, fmt.Errorf("error checking validator registration: %s", err)
	}

	if !isRegistered {
		log.Warn().Str("hash", tx.Hash).Str("validator", memoData.ValidatorAddress).
			Msg("validator not registered")
		return false, nil, nil
	}

	// Validate 1-1 address binding for XRP
	if !m.validateXRPAddressBinding(tx.Tx.Account, memoData.ImuachainAddress, tx.Hash) {
		log.Warn().Str("hash", tx.Hash).Str("xrp_addr", tx.Tx.Account).
			Str("imuachain_addr", memoData.ImuachainAddress).Msg("XRP address binding validation failed")
		return false, nil, nil
	}

	return true, memoData, nil
}

// XRPMemoData represents parsed memo data
type XRPMemoData struct {
	ImuachainAddress string
	ValidatorAddress string
}

// parseXRPMemo parses memo data from XRP transaction
func parseXRPMemo(memos []types.XRPMemo) (*XRPMemoData, error) {
	for _, memo := range memos {
		// Validate MemoType is "Description" (hex: 4465736372697074696F6E)
		if memo.Memo.MemoType != "4465736372697074696F6E" {
			continue
		}

		// Decode memo data
		buffer, err := hex.DecodeString(memo.Memo.MemoData)
		if err != nil {
			continue
		}

		// Validate minimum length (41 bytes validator + 20 bytes ethereum address)
		if len(buffer) < 61 {
			continue
		}

		// Extract ethereum address (last 20 bytes)
		ethBytes := buffer[len(buffer)-20:]
		imuachainAddress := "0x" + hex.EncodeToString(ethBytes)

		if !common.IsHexAddress(imuachainAddress) {
			continue
		}

		// Extract validator address (remaining bytes before ethereum address)
		validatorBytes := buffer[:len(buffer)-20]
		validatorAddress := string(validatorBytes)

		if len(validatorAddress) != 41 {
			continue
		}

		return &XRPMemoData{
			ImuachainAddress: strings.ToLower(imuachainAddress),
			ValidatorAddress: validatorAddress,
		}, nil
	}

	return nil, fmt.Errorf("no valid memo data found")
}

// saveXRPTransaction saves XRP transaction data to database
func (m *Module) saveXRPTransaction(tx types.XRPTransaction, memoData *XRPMemoData) error {
	// Get amount
	amountStr, ok := tx.Tx.Amount.(string)
	if !ok {
		return fmt.Errorf("invalid amount type")
	}

	amount, err := strconv.ParseInt(amountStr, 10, 64)
	if err != nil {
		return fmt.Errorf("failed to parse amount: %s", err)
	}

	// Create staker ID
	stakerID := memoData.ImuachainAddress + "_0x2" // XRP chain ID = 2

	// Save staker asset
	stakerAsset := &types.BootstrapStakerAsset{
		StakerID:     stakerID,
		AssetID:      VirtualAddress + "_0x2", // XRP asset ID: virtualAddress + chainID
		Deposited:    amountStr,
		Withdrawable: "0", // All stakes must be delegated
		Delegated:    amountStr,
		UpdatedAt:    time.Now(),
	}

	if err := m.database.SaveBootstrapStakerAsset(stakerAsset); err != nil {
		return fmt.Errorf("failed to save staker asset: %s", err)
	}

	// Save delegation state
	delegationState := &types.BootstrapDelegationState{
		StakerID:     stakerID,
		AssetID:      VirtualAddress + "_0x2",
		OperatorAddr: memoData.ValidatorAddress,
		Delegated:    amountStr,
		UpdatedAt:    time.Now(),
	}

	if err := m.database.SaveBootstrapDelegationState(delegationState); err != nil {
		return fmt.Errorf("failed to save delegation state: %s", err)
	}

	log.Info().
		Str("hash", tx.Hash).
		Str("staker", stakerID).
		Str("validator", memoData.ValidatorAddress).
		Int64("amount", amount).
		Msg("processed XRP bootstrap transaction")

	return nil
}

// getXRPVaultTransactionsFromLedger gets XRP vault transactions from a specific ledger range
func (m *Module) getXRPVaultTransactionsFromLedger(fromLedger, toLedger int64) ([]types.XRPTransaction, error) {
	var allTxs []types.XRPTransaction

	// Process ledgers in batches to avoid overwhelming the API
	batchSize := int64(m.Config.ScanBatchSize)
	if batchSize <= 0 {
		batchSize = 10 // Default batch size
	}

	for currentLedger := fromLedger; currentLedger <= toLedger; currentLedger += batchSize {
		endLedger := currentLedger + batchSize - 1
		if endLedger > toLedger {
			endLedger = toLedger
		}

		log.Debug().Int64("from", currentLedger).Int64("to", endLedger).
			Msg("scanning XRP ledgers for vault transactions")

		// Get transactions for each ledger in the range
		for ledgerIndex := currentLedger; ledgerIndex <= endLedger; ledgerIndex++ {
			ledgerTxs, err := m.getXRPLedgerTransactions(ledgerIndex)
			if err != nil {
				log.Err(err).Int64("ledger", ledgerIndex).Msg("error getting XRP ledger transactions")
				continue
			}

			// Filter transactions involving vault address
			for _, tx := range ledgerTxs {
				if m.isXRPVaultTransaction(tx) {
					allTxs = append(allTxs, tx)
				}
			}
		}

		// Add delay between batches to be API-friendly
		if endLedger < toLedger {
			time.Sleep(100 * time.Millisecond)
		}
	}

	log.Info().Int("total_txs", len(allTxs)).
		Int64("from_ledger", fromLedger).
		Int64("to_ledger", toLedger).
		Msg("completed XRP vault transaction scanning")

	return allTxs, nil
}

// getXRPLedgerTransactions gets all transactions for a specific ledger index
func (m *Module) getXRPLedgerTransactions(ledgerIndex int64) ([]types.XRPTransaction, error) {
	requestBody := map[string]interface{}{
		"method": "ledger",
		"params": []map[string]interface{}{
			{
				"ledger_index": ledgerIndex,
				"transactions": true,
				"expand":       true,
			},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %s", err)
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", m.Config.XRPRPC, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get ledger transactions: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d for ledger %d", resp.StatusCode, ledgerIndex)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %s", err)
	}

	result, ok := response["result"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format")
	}

	ledger, ok := result["ledger"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid ledger format")
	}

	transactions, ok := ledger["transactions"].([]interface{})
	if !ok {
		// No transactions in this ledger
		return []types.XRPTransaction{}, nil
	}

	var txs []types.XRPTransaction
	for _, txData := range transactions {
		txMap, ok := txData.(map[string]interface{})
		if !ok {
			continue
		}

		// Parse transaction using safe parsing function
		tx, err := parseXRPTransaction(txMap)
		if err != nil {
			log.Debug().Str("error", err.Error()).Int64("ledger", ledgerIndex).
				Msg("skipping invalid XRP transaction in ledger")
			continue
		}

		txs = append(txs, *tx)
	}

	return txs, nil
}

// isXRPVaultTransaction checks if a transaction involves the vault address
func (m *Module) isXRPVaultTransaction(tx types.XRPTransaction) bool {
	vaultAddr := strings.ToLower(strings.TrimSpace(m.Config.XRPVaultAddr))

	// Check if transaction is sent to vault address
	if strings.ToLower(strings.TrimSpace(tx.Tx.Destination)) == vaultAddr {
		return true
	}

	// Check if transaction is sent from vault address
	if strings.ToLower(strings.TrimSpace(tx.Tx.Account)) == vaultAddr {
		return true
	}

	return false
}

// refetchBootstrapStates refetches all bootstrap states and update them using parallel processing
// Note: BTC and XRP processing now uses separate mutexes for true parallelism in address mapping
// Database operations are handled per-chain and should not significantly compete with each other
func (m *Module) refetchBootstrapStates() error {
	log.Debug().Str("module", "bootstrap").Msg("starting parallel bootstrap states refetch")

	// Create a context with timeout for overall operation (10 minutes)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	var wg sync.WaitGroup

	// Define a struct to hold chain processing results
	type chainResult struct {
		chainName string
		err       error
	}

	resultChan := make(chan chainResult, 3)

	// Channel to signal when all goroutines are done
	doneChan := make(chan struct{})

	// Parallel processing for ETH states
	wg.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Error().Interface("panic", r).Str("chain", "ETH").
					Msg("ETH states refetch goroutine panicked")
				resultChan <- chainResult{chainName: "ETH", err: fmt.Errorf("goroutine panic: %v", r)}
			}
			wg.Done()
		}()

		// Check for cancellation before starting
		select {
		case <-ctx.Done():
			log.Warn().Str("chain", "ETH").Msg("ETH states refetch cancelled before starting")
			resultChan <- chainResult{chainName: "ETH", err: ctx.Err()}
			return
		default:
		}

		log.Debug().Str("chain", "ETH").Msg("starting ETH states refetch")
		start := time.Now()

		err := m.refetchETHStates()
		duration := time.Since(start)

		if err != nil {
			if err == context.DeadlineExceeded || err == context.Canceled {
				log.Warn().Err(err).Str("chain", "ETH").Dur("duration", duration).
					Msg("ETH states refetch cancelled or timed out")
			} else {
				log.Error().Err(err).Str("chain", "ETH").Dur("duration", duration).
					Msg("ETH states refetch failed")
			}
			resultChan <- chainResult{chainName: "ETH", err: err}
		} else {
			log.Info().Str("chain", "ETH").Dur("duration", duration).
				Msg("ETH states refetch completed successfully")
			resultChan <- chainResult{chainName: "ETH", err: nil}
		}
	}()

	// Parallel processing for BTC states
	wg.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Error().Interface("panic", r).Str("chain", "BTC").
					Msg("BTC states refetch goroutine panicked")
				resultChan <- chainResult{chainName: "BTC", err: fmt.Errorf("goroutine panic: %v", r)}
			}
			wg.Done()
		}()

		// Check for cancellation before starting
		select {
		case <-ctx.Done():
			log.Warn().Str("chain", "BTC").Msg("BTC states refetch cancelled before starting")
			resultChan <- chainResult{chainName: "BTC", err: ctx.Err()}
			return
		default:
		}

		log.Debug().Str("chain", "BTC").Msg("starting BTC states refetch")
		start := time.Now()

		err := m.refetchBTCStates()
		duration := time.Since(start)

		if err != nil {
			if err == context.DeadlineExceeded || err == context.Canceled {
				log.Warn().Err(err).Str("chain", "BTC").Dur("duration", duration).
					Msg("BTC states refetch cancelled or timed out")
			} else {
				log.Error().Err(err).Str("chain", "BTC").Dur("duration", duration).
					Msg("BTC states refetch failed")
			}
			resultChan <- chainResult{chainName: "BTC", err: err}
		} else {
			log.Info().Str("chain", "BTC").Dur("duration", duration).
				Msg("BTC states refetch completed successfully")
			resultChan <- chainResult{chainName: "BTC", err: nil}
		}
	}()

	// Parallel processing for XRP states
	wg.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Error().Interface("panic", r).Str("chain", "XRP").
					Msg("XRP states refetch goroutine panicked")
				resultChan <- chainResult{chainName: "XRP", err: fmt.Errorf("goroutine panic: %v", r)}
			}
			wg.Done()
		}()

		// Check for cancellation before starting
		select {
		case <-ctx.Done():
			log.Warn().Str("chain", "XRP").Msg("XRP states refetch cancelled before starting")
			resultChan <- chainResult{chainName: "XRP", err: ctx.Err()}
			return
		default:
		}

		log.Debug().Str("chain", "XRP").Msg("starting XRP states refetch")
		start := time.Now()

		err := m.refetchXRPStates()
		duration := time.Since(start)

		if err != nil {
			if err == context.DeadlineExceeded || err == context.Canceled {
				log.Warn().Err(err).Str("chain", "XRP").Dur("duration", duration).
					Msg("XRP states refetch cancelled or timed out")
			} else {
				log.Error().Err(err).Str("chain", "XRP").Dur("duration", duration).
					Msg("XRP states refetch failed")
			}
			resultChan <- chainResult{chainName: "XRP", err: err}
		} else {
			log.Info().Str("chain", "XRP").Dur("duration", duration).
				Msg("XRP states refetch completed successfully")
			resultChan <- chainResult{chainName: "XRP", err: nil}
		}
	}()

	// Start a goroutine to signal when all work is done
	go func() {
		wg.Wait()
		close(doneChan)
	}()

	// Wait for either all goroutines to complete or timeout
	select {
	case <-doneChan:
		// All goroutines completed normally
		log.Debug().Msg("all bootstrap refetch goroutines completed")
	case <-ctx.Done():
		// Timeout occurred
		log.Warn().Err(ctx.Err()).Msg("bootstrap refetch operation timed out, waiting for goroutines to finish")
		// Cancel context to signal goroutines to stop
		cancel()
		// Wait a bit more for graceful shutdown
		gracefulCtx, gracefulCancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer gracefulCancel()
		select {
		case <-doneChan:
			log.Info().Msg("all bootstrap refetch goroutines finished gracefully after timeout")
		case <-gracefulCtx.Done():
			log.Error().Msg("some bootstrap refetch goroutines did not finish gracefully - continuing anyway")
		}
	}

	close(resultChan)

	// Collect and process results
	var errors []string
	successCount := 0

	for result := range resultChan {
		if result.err != nil {
			if result.err == context.DeadlineExceeded || result.err == context.Canceled {
				errors = append(errors, fmt.Sprintf("%s: timed out", result.chainName))
			} else {
				errors = append(errors, fmt.Sprintf("%s: %s", result.chainName, result.err.Error()))
			}
		} else {
			successCount++
		}
	}

	// Log summary of results
	totalChains := 3
	log.Info().Int("successful", successCount).Int("failed", len(errors)).Int("total", totalChains).
		Msg("parallel bootstrap states refetch completed")

	// Return error if any chain failed
	if len(errors) > 0 {
		if len(errors) == totalChains {
			// All chains failed
			return fmt.Errorf("all chains failed to refetch states: %s", strings.Join(errors, "; "))
		} else {
			// Some chains failed - log warning but don't fail the entire operation
			log.Warn().Strs("failed_chains", errors).
				Msg("some chains failed during parallel refetch, but continuing")

			// You can choose to return error here if you want strict failure handling
			// return fmt.Errorf("some chains failed: %s", strings.Join(errors, "; "))
		}
	}

	return nil
}

// refetchBootstrapStatesSequential refetches all bootstrap states sequentially (backup implementation)
// This function can be used if parallel processing causes issues or if strict error handling is needed
func (m *Module) refetchBootstrapStatesSequential() error {
	log.Debug().Str("module", "bootstrap").Msg("starting sequential bootstrap states refetch")

	// Process ETH states
	log.Debug().Str("chain", "ETH").Msg("starting ETH states refetch")
	start := time.Now()
	err := m.refetchETHStates()
	if err != nil {
		log.Error().Err(err).Str("chain", "ETH").Dur("duration", time.Since(start)).
			Msg("ETH states refetch failed")
		return fmt.Errorf("error while refetching states from ETH: %s", err)
	}
	log.Info().Str("chain", "ETH").Dur("duration", time.Since(start)).
		Msg("ETH states refetch completed successfully")

	// Process BTC states
	log.Debug().Str("chain", "BTC").Msg("starting BTC states refetch")
	start = time.Now()
	err = m.refetchBTCStates()
	if err != nil {
		log.Error().Err(err).Str("chain", "BTC").Dur("duration", time.Since(start)).
			Msg("BTC states refetch failed")
		return fmt.Errorf("error while refetching states from BTC: %s", err)
	}
	log.Info().Str("chain", "BTC").Dur("duration", time.Since(start)).
		Msg("BTC states refetch completed successfully")

	// Process XRP states
	log.Debug().Str("chain", "XRP").Msg("starting XRP states refetch")
	start = time.Now()
	err = m.refetchXRPStates()
	if err != nil {
		log.Error().Err(err).Str("chain", "XRP").Dur("duration", time.Since(start)).
			Msg("XRP states refetch failed")
		return fmt.Errorf("error while refetching states from XRP: %s", err)
	}
	log.Info().Str("chain", "XRP").Dur("duration", time.Since(start)).
		Msg("XRP states refetch completed successfully")

	log.Info().Msg("sequential bootstrap states refetch completed successfully")
	return nil
}

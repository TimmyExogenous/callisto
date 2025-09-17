package bootstrap

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethcoretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/forbole/callisto/v4/modules/bootstrap/bootstrap_binding"
	"github.com/forbole/callisto/v4/types"
	assetstypes "github.com/imua-xyz/imuachain/x/assets/types"
	"github.com/rs/zerolog/log"
	"math/big"
	"time"
)

func (m *Module) getSenderByTransactionRawLog(ctx context.Context, rawLog ethcoretypes.Log) (common.Address, error) {
	tx, _, err := m.EthHttpClient.TransactionByHash(ctx, rawLog.TxHash)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get the commission update transaction,err:%s", err)
	}
	txSender, err := m.EthHttpClient.TransactionSender(ctx, tx, rawLog.BlockHash, rawLog.TxIndex)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get the sender of commission update transaction,err:%s", err)
	}
	return txSender, nil
}

func (m *Module) updateStatesAfterDepositOrClaim(stakerAddr, assetAddr common.Address) error {
	totalDepositAmount, err := m.bootstrapSession.TotalDepositAmounts(stakerAddr, assetAddr)
	if err != nil {
		return err
	}
	withdrawableAmount, err := m.bootstrapSession.WithdrawableAmounts(stakerAddr, assetAddr)
	if err != nil {
		return err
	}
	if totalDepositAmount.Cmp(withdrawableAmount) < 0 {
		return fmt.Errorf("total deposit amount:%s is less than withdrawable amount:%s", totalDepositAmount, withdrawableAmount)
	}
	delegationAmount := big.NewInt(0).Sub(totalDepositAmount, withdrawableAmount)
	stakerID, assetID := assetstypes.GetStakerIDAndAssetID(m.Config.ETHLZChainID, stakerAddr[:], assetAddr[:])
	err = m.database.SaveBootstrapStakerAsset(&types.BootstrapStakerAsset{
		StakerID:     stakerID,
		AssetID:      assetID,
		Deposited:    totalDepositAmount.String(),
		Withdrawable: withdrawableAmount.String(),
		Delegated:    delegationAmount.String(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		return err
	}
	return nil
}

// RunAsyncOperations implements modules.AsyncOperationsModule
func (m *Module) RunAsyncOperations() {
	commonWatchCtx := &bind.WatchOpts{
		Context: m.ctx,
	}
	// create event channels for validator
	newValidatorCh := make(chan *bootstrap_binding.BootstrapValidatorRegistered)
	commissionUpdatedCh := make(chan *bootstrap_binding.BootstrapValidatorCommissionUpdated)
	keyReplaceCh := make(chan *bootstrap_binding.BootstrapValidatorKeyReplaced)

	newValidatorSub, err := m.bootstrapFilterer.WatchValidatorRegistered(commonWatchCtx, newValidatorCh)
	if err != nil {
		panic(fmt.Errorf("failed to watch validator registeration,err:%s", err))
	}
	defer newValidatorSub.Unsubscribe()

	commissionSub, err := m.bootstrapFilterer.WatchValidatorCommissionUpdated(commonWatchCtx, commissionUpdatedCh)
	if err != nil {
		panic(fmt.Errorf("failed to watch commission update,err:%s", err))
	}
	defer commissionSub.Unsubscribe()

	keyReplaceSub, err := m.bootstrapFilterer.WatchValidatorKeyReplaced(commonWatchCtx, keyReplaceCh)
	if err != nil {
		panic(fmt.Errorf("failed to watch key replace,err:%s", err))
	}
	defer keyReplaceSub.Unsubscribe()

	// create event channels for whitelist assets
	newAssetCh := make(chan *bootstrap_binding.BootstrapWhitelistTokenAdded)
	newAssetSub, err := m.bootstrapFilterer.WatchWhitelistTokenAdded(commonWatchCtx, newAssetCh)
	if err != nil {
		panic(fmt.Errorf("failed to watch whitelist token addition,err:%s", err))
	}
	defer keyReplaceSub.Unsubscribe()

	// create event channels for deposit, claim, delegation and undelegation
	depositCh := make(chan *bootstrap_binding.BootstrapDepositResult)
	depositSub, err := m.bootstrapFilterer.WatchDepositResult(commonWatchCtx, depositCh, nil, nil, nil)
	if err != nil {
		panic(fmt.Errorf("failed to watch token deposit,err:%s", err))
	}
	defer depositSub.Unsubscribe()

	claimCh := make(chan *bootstrap_binding.BootstrapClaimPrincipalResult)
	claimSub, err := m.bootstrapFilterer.WatchClaimPrincipalResult(commonWatchCtx, claimCh, nil, nil, nil)
	if err != nil {
		panic(fmt.Errorf("failed to watch token claim,err:%s", err))
	}
	defer claimSub.Unsubscribe()

	delegationCh := make(chan *bootstrap_binding.BootstrapDelegateResult)
	delegationSub, err := m.bootstrapFilterer.WatchDelegateResult(commonWatchCtx, delegationCh, nil, nil, nil)
	if err != nil {
		panic(fmt.Errorf("failed to watch token delegation,err:%s", err))
	}
	defer delegationSub.Unsubscribe()

	undelegationCh := make(chan *bootstrap_binding.BootstrapUndelegateResult)
	undelegationSub, err := m.bootstrapFilterer.WatchUndelegateResult(commonWatchCtx, undelegationCh, nil, nil, nil)
	if err != nil {
		panic(fmt.Errorf("failed to watch token undelegation,err:%s", err))
	}
	defer undelegationSub.Unsubscribe()

	for {
		select {
		case err := <-newValidatorSub.Err():
			log.Err(err).Msg("new validator subscription error")
			return
		case err := <-commissionSub.Err():
			log.Err(err).Msg("commission update subscription error")
			return
		case err := <-keyReplaceSub.Err():
			log.Err(err).Msg("key replace subscription error")
			return
		case err := <-newAssetSub.Err():
			log.Err(err).Msg("whitelist token addition subscription error")
			return
		case err := <-depositSub.Err():
			log.Err(err).Msg("token deposit subscription error")
			return
		case err := <-claimSub.Err():
			log.Err(err).Msg("token claim subscription error")
			return
		case err := <-delegationSub.Err():
			log.Err(err).Msg("token delegation subscription error")
			return
		case err := <-undelegationSub.Err():
			log.Err(err).Msg("token undelegation subscription error")
			return
		case e := <-newValidatorCh:
			// save the new validator
			err := m.database.SaveBootstrapValidator(&types.BootstrapValidator{
				ValidatorEthAddress: e.EthAddress.String(),
				ValidatorIMAddress:  e.ValidatorAddress,
				ValidatorName:       e.Name,
				ConsensusPubKey:     hexutil.Encode(e.ConsensusPublicKey[:]),
				Rate:                e.Commission.Rate.String(),
				MaxRate:             e.Commission.MaxRate.String(),
				MaxChangeRate:       e.Commission.MaxChangeRate.String(),
				UpdatedAt:           time.Now(),
			})
			if err != nil {
				log.Err(err).Msg("failed to saving the new validator")
			}
		case e := <-commissionUpdatedCh:
			// get the tx sender as the validator ETH address
			txSender, err := m.getSenderByTransactionRawLog(m.ctx, e.Raw)
			if err != nil {
				log.Err(err)
				continue
			}
			err = m.database.UpdateCommissionRate(txSender.String(), e.NewRate.String())
			if err != nil {
				log.Err(err).Str("validatorEthAddr", txSender.String()).Msg("failed to update the commission rate")
			}
		case e := <-keyReplaceCh:
			err = m.database.UpdateConsensusPubKey(e.ValidatorAddress, hexutil.Encode(e.NewConsensusPublicKey[:]))
			if err != nil {
				log.Err(err).Str("validatorAddr", e.ValidatorAddress).Msg("failed to replace the consensus key")
			}
		case e := <-newAssetCh:
			tokenCount, err := m.bootstrapSession.GetWhitelistedTokensCount()
			if err != nil {
				log.Err(err).Msg("failed to get the count of whitelisted tokens")
				continue
			}
			// get the token info by iterating all indexes
			for i := int64(0); i < tokenCount.Int64(); i++ {
				tokenInfo, err := m.bootstrapSession.GetWhitelistedTokenAtIndex(big.NewInt(i))
				if err != nil {
					log.Err(err).Msg("failed to get the count of whitelisted tokens")
					continue
				}
				if tokenInfo.TokenAddress == e.Token {
					_, assetID := assetstypes.GetStakerIDAndAssetID(m.Config.ETHLZChainID, nil, e.Token[:])
					err = m.database.SaveBootstrapToken(&types.BootstrapToken{
						AssetID:            assetID,
						Address:            e.Token.String(),
						Name:               tokenInfo.Name,
						Symbol:             tokenInfo.Symbol,
						Decimals:           tokenInfo.Decimals,
						UpdatedAt:          time.Now(),
						LayerZeroChainID:   m.Config.ETHLZChainID,
						StakingTotalAmount: big.NewInt(0).String(),
					})
					if err != nil {
						log.Err(err).Str("token", e.Token.String()).Msg("failed to save the whitelist asset")
					}
					break
				}
			}
		case e := <-depositCh:
			if e.Success {
				err := m.updateStatesAfterDepositOrClaim(e.Depositor, e.Token)
				if err != nil {
					log.Err(err).Str("depositor", e.Depositor.String()).Str("token", e.Token.String()).Str("amount", e.Amount.String()).Msg("failed to handle the deposit event")
				}
			}
		case e := <-claimCh:
			if e.Success {
				err := m.updateStatesAfterDepositOrClaim(e.Withdrawer, e.Token)
				if err != nil {
					log.Err(err).Str("withdrawer", e.Withdrawer.String()).Str("token", e.Token.String()).Str("amount", e.Amount.String()).Msg("failed to handle the claim event")
				}
			}
		}
	}

}

// subscribeAndHandleEvents subscribes and handles all events from the bootstrap contract.
func (m *Module) subscribeAndHandleEvents() error {

	return nil
}

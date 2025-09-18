package bootstrap

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/forbole/callisto/v4/types"
	"github.com/go-co-op/gocron"
	assetstypes "github.com/imua-xyz/imuachain/x/assets/types"
	"github.com/rs/zerolog/log"
	"math/big"
	"time"
)

func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "bootstrap").Msg("setting up periodic tasks")

	// Schedule a cron job to run.
	if _, err := scheduler.Every(m.Config.ETHUpdateInterval).Minutes().Do(func() {
		m.refetchETHStates()
	}); err != nil {
		return fmt.Errorf("failed to set up the daily ETH states refetch operation: %s", err)
	}

	if _, err := scheduler.Every(m.Config.BTCUpdateInterval).Minutes().Do(func() {
		m.refetchBTCStates()
	}); err != nil {
		return fmt.Errorf("failed to set up the daily BTC states refetch operation: %s", err)
	}

	if _, err := scheduler.Every(m.Config.XRPUpdateInterval).Minutes().Do(func() {
		m.refetchXRPStates()
	}); err != nil {
		return fmt.Errorf("failed to set up the daily XRP states refetch operation: %s", err)
	}

	return nil
}

func (m *Module) updateStakerAsset(stakerAddr, assetAddr common.Address) (string, string, error) {
	totalDepositAmount, err := m.bootstrapSession.TotalDepositAmounts(stakerAddr, assetAddr)
	if err != nil {
		return "", "", err
	}

	withdrawableAmount, err := m.bootstrapSession.WithdrawableAmounts(stakerAddr, assetAddr)
	if err != nil {
		return "", "", err
	}
	if totalDepositAmount.Cmp(withdrawableAmount) < 0 {
		return "", "", fmt.Errorf("total deposit amount:%s is less than withdrawable amount:%s", totalDepositAmount, withdrawableAmount)
	}
	delegationAmount := big.NewInt(0).Sub(totalDepositAmount, withdrawableAmount)
	stakerID, assetID := assetstypes.GetStakerIDAndAssetID(m.Config.ETHLZChainID, stakerAddr[:], assetAddr[:])

	stakerAssetExist, err := m.database.BootstrapStakerAssetExists(stakerID, assetID)
	if err != nil {
		return "", "", err
	}
	if !stakerAssetExist && totalDepositAmount.Cmp(big.NewInt(0)) == 0 {
		// In the refetch case, do nothing since the staker has no deposit for this asset.
		return "", "", nil
	}
	err = m.database.SaveBootstrapStakerAsset(&types.BootstrapStakerAsset{
		StakerID:     stakerID,
		AssetID:      assetID,
		Deposited:    totalDepositAmount.String(),
		Withdrawable: withdrawableAmount.String(),
		Delegated:    delegationAmount.String(),
		UpdatedAt:    time.Now(),
	})
	if err != nil {
		return "", "", err
	}
	return stakerID, assetID, nil
}

func (m *Module) refetchETHStates() error {
	log.Debug().Str("module", "bootstrap").Str("refetching", "ETH states").
		Msg("refetching ETH states")

	// refetch all validators
	validatorCount, err := m.bootstrapSession.GetValidatorsCount()
	if err != nil {
		return err
	}
	validatorIMAddresses := make([]string, validatorCount.Int64())
	validatorETHAddresses := make([]common.Address, validatorCount.Int64())
	for i := int64(0); i < validatorCount.Int64(); i++ {
		validatorEthAddr, err := m.bootstrapSession.RegisteredValidators(big.NewInt(i))
		if err != nil {
			return fmt.Errorf("failed to call RegisteredValidators,index:%d,err:%s", i, err)
		}
		validatorIMAddr, err := m.bootstrapSession.EthToImAddress(validatorEthAddr)
		if err != nil {
			return fmt.Errorf("failed to call EthToImAddress,validatorEthAddr:%s,err:%s", validatorEthAddr, err)
		}
		validatorInfo, err := m.bootstrapSession.Validators(validatorIMAddr)
		if err != nil {
			return fmt.Errorf("failed to call Validators,validatorIMAddr:%s,err:%s", validatorIMAddr, err)
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
		validatorIMAddresses[i] = validatorIMAddr
		validatorETHAddresses[i] = validatorEthAddr
	}

	// refetch all staking assets
	tokenCount, err := m.bootstrapSession.GetWhitelistedTokensCount()
	if err != nil {
		return err
	}
	stakingAssets := make([]common.Address, tokenCount.Int64())
	stakingAssetIDs := make([]string, tokenCount.Int64())
	// get the token info by iterating all indexes
	for i := int64(0); i < tokenCount.Int64(); i++ {
		tokenInfo, err := m.bootstrapSession.GetWhitelistedTokenAtIndex(big.NewInt(i))
		if err != nil {
			return fmt.Errorf("failed to call GetWhitelistedTokenAtIndex,index:%d,err:%s", i, err)
		}

		_, assetID := assetstypes.GetStakerIDAndAssetID(m.Config.ETHLZChainID, nil, tokenInfo.TokenAddress[:])
		assetDepositAmount, err := m.bootstrapSession.DepositsByToken(tokenInfo.TokenAddress)
		if err != nil {
			return fmt.Errorf("failed to call DepositsByToken,tokenAddr:%s,err:%s", tokenInfo.TokenAddress, err)
		}
		err = m.database.SaveBootstrapToken(&types.BootstrapToken{
			AssetID:            assetID,
			Address:            tokenInfo.TokenAddress.String(),
			Name:               tokenInfo.Name,
			Symbol:             tokenInfo.Symbol,
			Decimals:           tokenInfo.Decimals,
			UpdatedAt:          time.Now(),
			LayerZeroChainID:   m.Config.ETHLZChainID,
			StakingTotalAmount: assetDepositAmount.String(),
		})
		if err != nil {
			return err
		}
		stakingAssets[i] = tokenInfo.TokenAddress
		stakingAssetIDs[i] = assetID
	}

	// refetch all staker assets
	depositorCount, err := m.bootstrapSession.GetDepositorsCount()
	if err != nil {
		return fmt.Errorf("failed to call GetDepositorsCount,err:%s", err)
	}
	for i := int64(0); i < depositorCount.Int64(); i++ {
		depositer, err := m.bootstrapSession.Depositors(big.NewInt(i))
		if err != nil {
			return fmt.Errorf("failed to call Depositors,index:%d,err:%s", i, err)
		}
		for _, assetAddr := range stakingAssets {
			stakerID, assetID, err := m.updateStakerAsset(depositer, assetAddr)
			if err != nil {
				return fmt.Errorf("failed to call Depositors,index:%d,err:%s", i, err)
			}
			if stakerID != "" && assetID != "" {
				for _, validator := range validatorIMAddresses {
					delegationAmount, err := m.bootstrapSession.Delegations(depositer, validator, assetAddr)
					if err != nil {
						return err
					}
					delegationExist, err := m.database.BootstrapDelegationExists(stakerID, assetID, validator)
					if err != nil {
						return err
					}
					if !delegationExist && delegationAmount.Cmp(big.NewInt(0)) == 0 {
						// Skip the validator if the delegation hasn't been saved in the database and the fetched amount
						// is zero. This avoids saving delegations that don't exist. Since there is no flag indicating
						// whether a delegation exists, the bootstrap map will always return zero and nil error when the
						// delegation does not exist.
						// With this approach, it is possible that a delegation which was delegated in the bootstrap
						// but later fully undelegated—and whose actions were not captured by events—will not be
						// saved to the database. This is acceptable, because such a staker effectively has no delegated assets.
						// The database only records the current latest state and does not store operation history.
						continue
					}
					// update the delegation states
					err = m.database.SaveBootstrapDelegationState(&types.BootstrapDelegationState{
						StakerID:     stakerID,
						AssetID:      assetID,
						OperatorAddr: validator,
						Delegated:    delegationAmount.String(),
						UpdatedAt:    time.Now(),
					})
					if err != nil {
						return err
					}
				}
			}
		}
	}

	// fetch all validator assets
	for i, validator := range validatorIMAddresses {
		for j, assetAddr := range stakingAssets {
			validatorAssetAmount, err := m.bootstrapSession.DelegationsByValidator(validator, assetAddr)
			if err != nil {
				return err
			}
			operatorAssetExist, err := m.database.OperatorAssetExists(validator, stakingAssetIDs[j])
			if err != nil {
				return err
			}
			if !operatorAssetExist && validatorAssetAmount.Cmp(big.NewInt(0)) == 0 {
				continue
			}
			// get the self delegation amount
			selfDelegation, err := m.bootstrapSession.Delegations(validatorETHAddresses[i], validator, assetAddr)
			if err != nil {
				return err
			}
			err = m.database.SaveBootstrapOperatorAsset(&types.BootstrapOperatorAsset{
				OperatorAddr: validator,
				AssetID:      stakingAssetIDs[j],
				TotalAmount:  validatorAssetAmount.String(),
				SelfAmount:   selfDelegation.String(),
				OtherAmount:  big.NewInt(0).Sub(validatorAssetAmount, selfDelegation).String(),
				UpdatedAt:    time.Now(),
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *Module) refetchBTCStates() error {
	log.Debug().Str("module", "bootstrap").Str("refetching", "BTC states").
		Msg("refetching BTC states")
	//
	return nil
}
func (m *Module) refetchXRPStates() error {
	log.Debug().Str("module", "bootstrap").Str("refetching", "XRP states").
		Msg("refetching XRP states")
	//
	return nil
}

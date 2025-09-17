package bootstrap

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/forbole/callisto/v4/types"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
	"math/big"
	"time"
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

// refetchBootstrapStates refetches all bootstrap states and update them
func (m *Module) refetchBootstrapStates() error {
	err := m.refetchETHStates()
	if err != nil {
		return fmt.Errorf("error while refetching states from ETH: %s", err)
	}
	err = m.refetchBTCStates()
	if err != nil {
		return fmt.Errorf("error while refetching states from BTC: %s", err)
	}
	err = m.refetchXRPStates()
	if err != nil {
		return fmt.Errorf("error while refetching states from XRP: %s", err)
	}
	return nil
}

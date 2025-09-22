package bootstrap

import (
	"github.com/forbole/callisto/v4/types"
	"math/big"
	"time"
)

// RunAdditionalOperations implements modules.AdditionalOperationsModule
func (m *Module) RunAdditionalOperations() error {
	// Save default client chains from the config
	for _, clientChain := range m.Config.ClientChainInfos {
		err := m.database.SaveBootstrapClientChain(&types.BootstrapClientChain{
			Name:      clientChain.Name,
			MetaInfo:  clientChain.MetaInfo,
			LZChainID: clientChain.LZChainID,
		})
		if err != nil {
			return err
		}
	}
	// save default staking tokens from the config
	for _, stakingToken := range m.Config.StakingTokenInfos {
		err := m.database.SaveBootstrapToken(&types.BootstrapTokenState{
			BootstrapToken:     stakingToken,
			StakingTotalAmount: big.NewInt(0).String(),
			UpdatedAt:          time.Now(),
		})
		if err != nil {
			return err
		}
	}
	// Fetch all states at indexer startup.
	return m.refetchETHStates()
}

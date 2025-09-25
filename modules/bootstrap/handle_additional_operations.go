package bootstrap

import (
	"math/big"
	"time"

	"github.com/forbole/callisto/v4/types"
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
	return nil
}

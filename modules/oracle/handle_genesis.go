package oracle

import (
	"encoding/json"
	"fmt"

	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/forbole/callisto/v4/types"
	oracletypes "github.com/imua-xyz/imuachain/x/oracle/types"
	"github.com/rs/zerolog/log"
)

// HandleGenesis implements modules.GenesisModule
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", m.Name()).Msg("parsing genesis")

	// Read the genesis state
	var genState oracletypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[oracletypes.ModuleName], &genState)
	if err != nil {
		return fmt.Errorf("error while reading oracle genesis data: %s", err)
	}

	// Save the
	// 1. genesis params
	// 2. token config
	// 3. mapping from asset id to token id
	err = m.db.SaveOracleParams(types.NewOracleParams(genState.Params, doc.InitialHeight))
	if err != nil {
		return fmt.Errorf("error while storing genesis oracle params: %s", err)
	}

	for _, token := range genState.PricesList {
		tokenConfig := types.NewOracleTokenConfig(token)
		for _, price := range token.PriceList {
			priceHistory := types.NewOraclePriceHistory(tokenConfig, price)
			err = m.db.SaveOraclePriceHistory(priceHistory)
			if err != nil {
				return fmt.Errorf("error while storing genesis oracle price history: %s", err)
			}
		}
	}

	return nil
}

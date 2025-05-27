package immint

import (
	"encoding/json"
	"fmt"

	tmtypes "github.com/cometbft/cometbft/types"

	"github.com/forbole/callisto/v4/types"

	imminttypes "github.com/imua-xyz/imuachain/x/immint/types"
	"github.com/rs/zerolog/log"
)

// HandleGenesis implements modules.Module
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", m.Name()).Msg("parsing genesis")

	// Read the genesis state
	var genState imminttypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[imminttypes.ModuleName], &genState)
	if err != nil {
		return fmt.Errorf("error while reading immint genesis data: %s", err)
	}

	// Save the params
	err = m.db.SaveImmintParams(types.NewImmintParams(genState.Params, doc.InitialHeight))
	if err != nil {
		return fmt.Errorf("error while storing genesis immint params: %s", err)
	}

	return nil
}

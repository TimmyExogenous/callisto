package tokenomics

import (
	"encoding/json"
	"fmt"
	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/forbole/callisto/v4/types"

	"github.com/rs/zerolog/log"
)

// HandleGenesis implements modules.GenesisModule
func (m *Module) HandleGenesis(_ *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "tokenomics").Msg("parsing genesis")
	err := m.db.SaveTokenomicsParams(&types.TokenomicsParams{
		Height: 0,
	})
	if err != nil {
		return fmt.Errorf("error while initializing tokenomics paramters: %s", err)
	}
	return nil
}

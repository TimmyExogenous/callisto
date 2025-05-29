package oracle

import (
	"fmt"

	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/forbole/callisto/v4/types"
	juno "github.com/forbole/juno/v5/types"
	"github.com/rs/zerolog/log"
)

// HandleBlock implements BlockModule
func (m *Module) HandleBlock(
	block *tmctypes.ResultBlock, res *tmctypes.ResultBlockResults,
	_ []*juno.Tx, _ *tmctypes.ResultValidators,
) error {
	log.Debug().Str("module", m.Name()).Int64("height", block.Block.Height).
		Msg(fmt.Sprintf("updating %s", m.Name()))
	params, err := m.source.GetParams(block.Block.Height)
	if err != nil {
		return fmt.Errorf("error getting oracle params: %w", err)
	}

	// update params every block, if needed
	// TODO: replace with event-based update post 1.1.2 onwards
	oracleParams := types.NewOracleParams(params, block.Block.Height)
	if err := m.db.UpdateOracleParamsIfNeeded(oracleParams); err != nil {
		return fmt.Errorf("error updating oracle params: %w", err)
	}

	return nil
}

package oracle

import (
	juno "github.com/forbole/juno/v5/types"
)

// HandleTx implements modules.TransactionModule
func (m *Module) HandleTx(tx *juno.Tx) error {
	// we return nil here because we instead handle the message.
	return nil
}

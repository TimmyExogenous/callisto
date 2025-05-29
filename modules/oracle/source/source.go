package source

import (
	oracletypes "github.com/imua-xyz/imuachain/x/oracle/types"
)

type Source interface {
	GetParams(height int64) (oracletypes.Params, error)
}

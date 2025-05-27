package source

import (
	assetstypes "github.com/imua-xyz/imuachain/x/assets/types"
)

type Source interface {
	GetParams(height int64) (assetstypes.Params, error)
}

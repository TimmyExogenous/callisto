package source

import (
	imminttypes "github.com/imua-xyz/imuachain/x/immint/types"
)

type Source interface {
	GetParams(height int64) (imminttypes.Params, error)
}

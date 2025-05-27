package source

import (
	dogfoodtypes "github.com/imua-xyz/imuachain/x/dogfood/types"
)

type Source interface {
	GetParams(height int64) (dogfoodtypes.Params, error)
	GetValidators(height int64) ([]dogfoodtypes.ImuachainValidator, error)
}

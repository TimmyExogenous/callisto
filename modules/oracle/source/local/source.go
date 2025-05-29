package local

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/juno/v5/node/local"
	oracletypes "github.com/imua-xyz/imuachain/x/oracle/types"

	oraclesource "github.com/forbole/callisto/v4/modules/oracle/source"
)

// interface guard
var (
	_ oraclesource.Source = &Source{}
)

// Source implements oraclesource.Source using a local node
type Source struct {
	*local.Source
	querier oracletypes.QueryServer
}

// NewSource implements a new Source instance
func NewSource(source *local.Source, querier oracletypes.QueryServer) *Source {
	return &Source{
		Source:  source,
		querier: querier,
	}
}

// GetParams implements oraclesource.Source
func (s Source) GetParams(height int64) (oracletypes.Params, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return oracletypes.Params{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.querier.Params(
		sdk.WrapSDKContext(ctx),
		&oracletypes.QueryParamsRequest{},
	)
	if err != nil {
		return oracletypes.Params{}, err
	}

	return res.Params, nil
}

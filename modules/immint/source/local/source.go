package local

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/juno/v5/node/local"
	imminttypes "github.com/imua-xyz/imuachain/x/immint/types"

	immintsource "github.com/forbole/callisto/v4/modules/immint/source"
)

// interface guard
var (
	_ immintsource.Source = &Source{}
)

// Source implements immintsource.Source using a local node
type Source struct {
	*local.Source
	querier imminttypes.QueryServer
}

// NewSource implements a new Source instance
func NewSource(source *local.Source, querier imminttypes.QueryServer) *Source {
	return &Source{
		Source:  source,
		querier: querier,
	}
}

// GetParams implements immintsource.Source
func (s Source) GetParams(height int64) (imminttypes.Params, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return imminttypes.Params{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.querier.Params(
		sdk.WrapSDKContext(ctx),
		&imminttypes.QueryParamsRequest{},
	)
	if err != nil {
		return imminttypes.Params{}, err
	}

	return res.Params, nil
}

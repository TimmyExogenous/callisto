package remote

import (
	"github.com/forbole/juno/v5/node/remote"
	oracletypes "github.com/imua-xyz/imuachain/x/oracle/types"

	oraclesource "github.com/forbole/callisto/v4/modules/oracle/source"
)

// interface guard
var (
	_ oraclesource.Source = &Source{}
)

// Source implements oraclesource.Source using a remote node
type Source struct {
	*remote.Source
	querier oracletypes.QueryClient
}

// NewSource implements a new Source instance
func NewSource(source *remote.Source, querier oracletypes.QueryClient) *Source {
	return &Source{
		Source:  source,
		querier: querier,
	}
}

// GetParams implements oraclesource.Source
func (s Source) GetParams(height int64) (oracletypes.Params, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)

	res, err := s.querier.Params(
		ctx, &oracletypes.QueryParamsRequest{},
	)
	if err != nil {
		return oracletypes.Params{}, err
	}

	return res.Params, nil
}

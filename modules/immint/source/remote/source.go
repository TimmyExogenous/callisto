package remote

import (
	"github.com/forbole/juno/v5/node/remote"
	imminttypes "github.com/imua-xyz/imuachain/x/immint/types"

	immintsource "github.com/forbole/callisto/v4/modules/immint/source"
)

// interface guard
var (
	_ immintsource.Source = &Source{}
)

// Source implements immintsource.Source using a remote node
type Source struct {
	*remote.Source
	querier imminttypes.QueryClient
}

// NewSource implements a new Source instance
func NewSource(source *remote.Source, querier imminttypes.QueryClient) *Source {
	return &Source{
		Source:  source,
		querier: querier,
	}
}

// GetParams implements immintsource.Source
func (s Source) GetParams(height int64) (imminttypes.Params, error) {
	ctx := remote.GetHeightRequestContext(s.Ctx, height)

	res, err := s.querier.Params(
		ctx, &imminttypes.QueryParamsRequest{},
	)
	if err != nil {
		return imminttypes.Params{}, err
	}

	return res.Params, nil
}

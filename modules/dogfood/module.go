package dogfood

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/callisto/v4/database"
	dogfoodsource "github.com/forbole/callisto/v4/modules/dogfood/source"
	"github.com/forbole/juno/v5/modules"
	"github.com/forbole/juno/v5/types/config"
)

var (
	_ modules.Module            = &Module{}
	_ modules.GenesisModule     = &Module{}
	_ modules.TransactionModule = &Module{}
	_ modules.BlockModule       = &Module{}
	_ modules.MessageModule     = &Module{}
)

// Module implements x/assets module indexer
type Module struct {
	cdc    codec.Codec
	db     *database.Db
	source dogfoodsource.Source
	cfg    *Config
}

// NeawModule builds a new Module instance
func NewModule(cfg config.Config, source dogfoodsource.Source, cdc codec.Codec, db *database.Db) *Module {
	bz, err := cfg.GetBytes()
	if err != nil {
		panic(err)
	}
	dogfoodCfg, err := ParseConfig(bz)
	if err != nil {
		panic(err)
	}
	return &Module{
		cdc:    cdc,
		db:     db,
		cfg:    dogfoodCfg,
		source: source,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "dogfood"
}

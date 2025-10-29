package tokenomics

import (
	"github.com/forbole/callisto/v4/database"
	"github.com/forbole/juno/v5/modules"
	"github.com/forbole/juno/v5/types/config"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
)

// Module represents the virtual tokenomics module
type Module struct {
	db  *database.Db
	cfg *Config
}

// NewModule builds a new Module instance
func NewModule(cfg config.Config, db *database.Db) *Module {
	bz, err := cfg.GetBytes()
	if err != nil {
		panic(err)
	}
	tokenomicsCfg, err := ParseConfig(bz)
	if err != nil {
		panic(err)
	}
	return &Module{
		db:  db,
		cfg: tokenomicsCfg,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "tokenomics"
}

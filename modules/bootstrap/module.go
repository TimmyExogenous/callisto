package bootstrap

import (
	"context"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	callistodb "github.com/forbole/callisto/v4/database"
	"github.com/forbole/callisto/v4/modules/bootstrap/bootstrap_binding"
	"github.com/forbole/callisto/v4/modules/bootstrap/storage_binding"
	"github.com/forbole/juno/v5/types/config"

	"github.com/forbole/juno/v5/modules"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
)

type Module struct {
	database      *callistodb.Db
	EthHttpClient *ethclient.Client
	EthWSClient   *ethclient.Client
	Config        Config
	BootstrapAddr common.Address
	// No dedicated clients for BTC and XRP, since we may call their RPCs
	// directly using the http package.
	ctx               context.Context
	bootstrapSession  *bootstrap_binding.BootstrapCallerSession
	storageSession    *storage_binding.BootstrapStorageCallerSession
	bootstrapFilterer *bootstrap_binding.BootstrapFilterer
	storageFilterer   *storage_binding.BootstrapStorageFilterer
	// Address mappings for 1-1 binding validation (separate mutexes for true parallelism)
	btcAddressMappings map[string]string // bitcoin -> imuachain
	xrpAddressMappings map[string]string // xrp -> imuachain
	btcMappingMutex    sync.RWMutex      // Separate mutex for BTC mappings
	xrpMappingMutex    sync.RWMutex      // Separate mutex for XRP mappings
}

// NewModule builds a new Module instance
func NewModule(
	cfg config.Config,
	database *callistodb.Db,
) *Module {
	bz, err := cfg.GetBytes()
	if err != nil {
		panic(err)
	}

	bootstrapCfg, err := ParseConfig(bz)
	if err != nil {
		panic(err)
	}
	if !common.IsHexAddress(bootstrapCfg.BootstrapAddr) {
		panic(fmt.Sprintf("invalid bootstrap address:%s", bootstrapCfg.BootstrapAddr))
	}
	bootstrapAddr := common.HexToAddress(bootstrapCfg.BootstrapAddr)

	httpRC, err := rpc.DialContext(context.Background(), bootstrapCfg.ETHHttp)
	if err != nil {
		panic(err)
	}
	ethHttpClient := ethclient.NewClient(httpRC)

	websocketRC, err := rpc.DialContext(context.Background(), bootstrapCfg.ETHWebsocket)
	if err != nil {
		panic(err)
	}
	ethWSClient := ethclient.NewClient(websocketRC)

	// create the sessions for bootstrap and storage contracts.
	ctx := context.Background()
	bootstrapCaller, err := bootstrap_binding.NewBootstrapCaller(bootstrapAddr, ethHttpClient)
	if err != nil {
		panic(fmt.Errorf("failed to new bootstrap caller,err:%s", err))
	}
	bootstrapSession := &bootstrap_binding.BootstrapCallerSession{
		Contract: bootstrapCaller,
		CallOpts: bind.CallOpts{Context: ctx},
	}
	storageCaller, err := storage_binding.NewBootstrapStorageCaller(bootstrapAddr, ethHttpClient)
	if err != nil {
		panic(fmt.Errorf("failed to new bootstrap storage caller,err:%s", err))
	}
	storageSession := &storage_binding.BootstrapStorageCallerSession{
		Contract: storageCaller,
		CallOpts: bind.CallOpts{Context: ctx},
	}

	// create the filterer to subscribe all related events
	bootstrapFilterer, err := bootstrap_binding.NewBootstrapFilterer(bootstrapAddr, ethWSClient)
	if err != nil {
		panic(fmt.Errorf("failed to new bootstrap filterer,err:%s", err))
	}
	storageFilterer, err := storage_binding.NewBootstrapStorageFilterer(bootstrapAddr, ethWSClient)
	if err != nil {
		panic(fmt.Errorf("failed to new bootstrap storage filterer,err:%s", err))
	}

	module := &Module{
		database:           database,
		EthHttpClient:      ethHttpClient,
		EthWSClient:        ethWSClient,
		Config:             *bootstrapCfg,
		BootstrapAddr:      bootstrapAddr,
		ctx:                ctx,
		bootstrapSession:   bootstrapSession,
		storageSession:     storageSession,
		bootstrapFilterer:  bootstrapFilterer,
		storageFilterer:    storageFilterer,
		btcAddressMappings: make(map[string]string),
		xrpAddressMappings: make(map[string]string),
		// Note: mutexes are zero-valued, no need to initialize explicitly
	}

	// Initialize address bindings from database
	if err := module.loadExistingBindings(); err != nil {
		panic(fmt.Errorf("failed to load existing address bindings: %w", err))
	}

	return module
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "bootstrap"
}

// loadExistingBindings loads existing address bindings from database into memory
func (m *Module) loadExistingBindings() error {
	// Load BTC bindings
	btcBindings, err := m.database.GetAddressBindings("BTC")
	if err != nil {
		return fmt.Errorf("failed to load BTC address bindings: %w", err)
	}

	m.btcMappingMutex.Lock()
	for _, binding := range btcBindings {
		m.btcAddressMappings[binding.SourceAddr] = binding.TargetAddr
	}
	m.btcMappingMutex.Unlock()

	// Load XRP bindings
	xrpBindings, err := m.database.GetAddressBindings("XRP")
	if err != nil {
		return fmt.Errorf("failed to load XRP address bindings: %w", err)
	}

	m.xrpMappingMutex.Lock()
	for _, binding := range xrpBindings {
		m.xrpAddressMappings[binding.SourceAddr] = binding.TargetAddr
	}
	m.xrpMappingMutex.Unlock()

	return nil
}

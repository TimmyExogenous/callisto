package bootstrap

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	callistodb "github.com/forbole/callisto/v4/database"
	"github.com/forbole/callisto/v4/modules/bootstrap/bootstrap_binding"
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
	bootstrapFilterer *bootstrap_binding.BootstrapFilterer
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

	// create the filterer to subscribe all related events
	bootstrapFilterer, err := bootstrap_binding.NewBootstrapFilterer(bootstrapAddr, ethWSClient)
	if err != nil {
		panic(fmt.Errorf("failed to new bootstrap filterer,err:%s", err))
	}

	return &Module{
		database:          database,
		EthHttpClient:     ethHttpClient,
		EthWSClient:       ethWSClient,
		Config:            *bootstrapCfg,
		BootstrapAddr:     bootstrapAddr,
		ctx:               ctx,
		bootstrapSession:  bootstrapSession,
		bootstrapFilterer: bootstrapFilterer,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "bootstrap"
}

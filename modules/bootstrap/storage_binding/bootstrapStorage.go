// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package storage_binding

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IValidatorRegistryCommission is an auto generated low-level Go binding around an user-defined struct.
type IValidatorRegistryCommission struct {
	Rate          *big.Int
	MaxRate       *big.Int
	MaxChangeRate *big.Int
}

// BootstrapStorageABI is the input ABI used to generate the binding from.
const BootstrapStorageABI = "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structBootstrapStorage.ImmutableConfig\",\"components\":[{\"name\":\"imuachainChainId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"beaconOracleAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultBeacon\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"imuaCapsuleBeacon\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"beaconProxyBytecode\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"networkConfig\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"AFTER_PECTRA_MAX_EFFECTIVE_BALANCE_ETH_PER_VALIDATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"AFTER_PECTRA_MIN_ACTIVATION_BALANCE_ETH_PER_VALIDATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BEACON_ORACLE_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"BEACON_PROXY_BYTECODE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractBeaconProxyBytecode\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IMUACHAIN_CHAIN_ID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IMUA_ADDRESS_PREFIX\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IMUA_CAPSULE_BEACON\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VAULT_BEACON\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBeacon\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bootstrapped\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clientChainGatewayLogic\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"clientChainInitializationData\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"commissionEdited\",\"inputs\":[{\"name\":\"imAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"hasEdited\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"consensusPublicKeyInUse\",\"inputs\":[{\"name\":\"consensusKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"used\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"customProxyAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegations\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"imAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegationsByValidator\",\"inputs\":[{\"name\":\"imAddress\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositors\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositsByToken\",\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ethToImAddress\",\"inputs\":[{\"name\":\"ethAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"imAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inboundNonce\",\"inputs\":[{\"name\":\"eid\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"sender\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isDepositor\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"hasDeposited\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isValidImAddress\",\"inputs\":[{\"name\":\"addressToValidate\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isWhitelistedToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"whitelisted\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"offsetDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerToCapsule\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"capsule\",\"type\":\"address\",\"internalType\":\"contractIImuaCapsule\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registeredValidators\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"spawnTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerToPubkeyIDs\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakerToTokenToValidators\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenToVault\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"vault\",\"type\":\"address\",\"internalType\":\"contractIVault\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalDepositAmounts\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorNameInUse\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"used\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validators\",\"inputs\":[{\"name\":\"imAddress\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"commission\",\"type\":\"tuple\",\"internalType\":\"structIValidatorRegistry.Commission\",\"components\":[{\"name\":\"rate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxChangeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"consensusPublicKey\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelistTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawableAmounts\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"BootstrapNotTimeYet\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BootstrapUpgradeFailed\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Bootstrapped\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BootstrappedAlready\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CapsuleCreated\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"capsule\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClaimPrincipalResult\",\"inputs\":[{\"name\":\"success\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClientChainGatewayLogicUpdated\",\"inputs\":[{\"name\":\"newLogic\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"initializationData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegateResult\",\"inputs\":[{\"name\":\"success\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegatee\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositResult\",\"inputs\":[{\"name\":\"success\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DepositThenDelegateResult\",\"inputs\":[{\"name\":\"delegateSuccess\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegatee\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"delegatedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageExecuted\",\"inputs\":[{\"name\":\"act\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumAction\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MessageSent\",\"inputs\":[{\"name\":\"act\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumAction\"},{\"name\":\"packetId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"nativeFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OffsetDurationUpdated\",\"inputs\":[{\"name\":\"newOffsetDuration\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SpawnTimeUpdated\",\"inputs\":[{\"name\":\"newSpawnTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakedWithCapsule\",\"inputs\":[{\"name\":\"staker\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"capsule\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UndelegateResult\",\"inputs\":[{\"name\":\"success\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"undelegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"undelegatee\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VaultCreated\",\"inputs\":[{\"name\":\"underlyingToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WhitelistTokenAdded\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]"

// BootstrapStorage is an auto generated Go binding around an Ethereum contract.
type BootstrapStorage struct {
	BootstrapStorageCaller     // Read-only binding to the contract
	BootstrapStorageTransactor // Write-only binding to the contract
	BootstrapStorageFilterer   // Log filterer for contract events
}

// BootstrapStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type BootstrapStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BootstrapStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BootstrapStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BootstrapStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BootstrapStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BootstrapStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BootstrapStorageSession struct {
	Contract     *BootstrapStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BootstrapStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BootstrapStorageCallerSession struct {
	Contract *BootstrapStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// BootstrapStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BootstrapStorageTransactorSession struct {
	Contract     *BootstrapStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// BootstrapStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type BootstrapStorageRaw struct {
	Contract *BootstrapStorage // Generic contract binding to access the raw methods on
}

// BootstrapStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BootstrapStorageCallerRaw struct {
	Contract *BootstrapStorageCaller // Generic read-only contract binding to access the raw methods on
}

// BootstrapStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BootstrapStorageTransactorRaw struct {
	Contract *BootstrapStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBootstrapStorage creates a new instance of BootstrapStorage, bound to a specific deployed contract.
func NewBootstrapStorage(address common.Address, backend bind.ContractBackend) (*BootstrapStorage, error) {
	contract, err := bindBootstrapStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BootstrapStorage{BootstrapStorageCaller: BootstrapStorageCaller{contract: contract}, BootstrapStorageTransactor: BootstrapStorageTransactor{contract: contract}, BootstrapStorageFilterer: BootstrapStorageFilterer{contract: contract}}, nil
}

// NewBootstrapStorageCaller creates a new read-only instance of BootstrapStorage, bound to a specific deployed contract.
func NewBootstrapStorageCaller(address common.Address, caller bind.ContractCaller) (*BootstrapStorageCaller, error) {
	contract, err := bindBootstrapStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageCaller{contract: contract}, nil
}

// NewBootstrapStorageTransactor creates a new write-only instance of BootstrapStorage, bound to a specific deployed contract.
func NewBootstrapStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*BootstrapStorageTransactor, error) {
	contract, err := bindBootstrapStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageTransactor{contract: contract}, nil
}

// NewBootstrapStorageFilterer creates a new log filterer instance of BootstrapStorage, bound to a specific deployed contract.
func NewBootstrapStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*BootstrapStorageFilterer, error) {
	contract, err := bindBootstrapStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageFilterer{contract: contract}, nil
}

// bindBootstrapStorage binds a generic wrapper to an already deployed contract.
func bindBootstrapStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BootstrapStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BootstrapStorage *BootstrapStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BootstrapStorage.Contract.BootstrapStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BootstrapStorage *BootstrapStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BootstrapStorage.Contract.BootstrapStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BootstrapStorage *BootstrapStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BootstrapStorage.Contract.BootstrapStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BootstrapStorage *BootstrapStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BootstrapStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BootstrapStorage *BootstrapStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BootstrapStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BootstrapStorage *BootstrapStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BootstrapStorage.Contract.contract.Transact(opts, method, params...)
}

// AFTERPECTRAMAXEFFECTIVEBALANCEETHPERVALIDATOR is a free data retrieval call binding the contract method 0x5fe60b0d.
//
// Solidity: function AFTER_PECTRA_MAX_EFFECTIVE_BALANCE_ETH_PER_VALIDATOR() view returns(uint256)
func (_BootstrapStorage *BootstrapStorageCaller) AFTERPECTRAMAXEFFECTIVEBALANCEETHPERVALIDATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "AFTER_PECTRA_MAX_EFFECTIVE_BALANCE_ETH_PER_VALIDATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AFTERPECTRAMAXEFFECTIVEBALANCEETHPERVALIDATOR is a free data retrieval call binding the contract method 0x5fe60b0d.
//
// Solidity: function AFTER_PECTRA_MAX_EFFECTIVE_BALANCE_ETH_PER_VALIDATOR() view returns(uint256)
func (_BootstrapStorage *BootstrapStorageSession) AFTERPECTRAMAXEFFECTIVEBALANCEETHPERVALIDATOR() (*big.Int, error) {
	return _BootstrapStorage.Contract.AFTERPECTRAMAXEFFECTIVEBALANCEETHPERVALIDATOR(&_BootstrapStorage.CallOpts)
}

// AFTERPECTRAMAXEFFECTIVEBALANCEETHPERVALIDATOR is a free data retrieval call binding the contract method 0x5fe60b0d.
//
// Solidity: function AFTER_PECTRA_MAX_EFFECTIVE_BALANCE_ETH_PER_VALIDATOR() view returns(uint256)
func (_BootstrapStorage *BootstrapStorageCallerSession) AFTERPECTRAMAXEFFECTIVEBALANCEETHPERVALIDATOR() (*big.Int, error) {
	return _BootstrapStorage.Contract.AFTERPECTRAMAXEFFECTIVEBALANCEETHPERVALIDATOR(&_BootstrapStorage.CallOpts)
}

// AFTERPECTRAMINACTIVATIONBALANCEETHPERVALIDATOR is a free data retrieval call binding the contract method 0xf442193e.
//
// Solidity: function AFTER_PECTRA_MIN_ACTIVATION_BALANCE_ETH_PER_VALIDATOR() view returns(uint256)
func (_BootstrapStorage *BootstrapStorageCaller) AFTERPECTRAMINACTIVATIONBALANCEETHPERVALIDATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "AFTER_PECTRA_MIN_ACTIVATION_BALANCE_ETH_PER_VALIDATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AFTERPECTRAMINACTIVATIONBALANCEETHPERVALIDATOR is a free data retrieval call binding the contract method 0xf442193e.
//
// Solidity: function AFTER_PECTRA_MIN_ACTIVATION_BALANCE_ETH_PER_VALIDATOR() view returns(uint256)
func (_BootstrapStorage *BootstrapStorageSession) AFTERPECTRAMINACTIVATIONBALANCEETHPERVALIDATOR() (*big.Int, error) {
	return _BootstrapStorage.Contract.AFTERPECTRAMINACTIVATIONBALANCEETHPERVALIDATOR(&_BootstrapStorage.CallOpts)
}

// AFTERPECTRAMINACTIVATIONBALANCEETHPERVALIDATOR is a free data retrieval call binding the contract method 0xf442193e.
//
// Solidity: function AFTER_PECTRA_MIN_ACTIVATION_BALANCE_ETH_PER_VALIDATOR() view returns(uint256)
func (_BootstrapStorage *BootstrapStorageCallerSession) AFTERPECTRAMINACTIVATIONBALANCEETHPERVALIDATOR() (*big.Int, error) {
	return _BootstrapStorage.Contract.AFTERPECTRAMINACTIVATIONBALANCEETHPERVALIDATOR(&_BootstrapStorage.CallOpts)
}

// BEACONORACLEADDRESS is a free data retrieval call binding the contract method 0x90a9b42e.
//
// Solidity: function BEACON_ORACLE_ADDRESS() view returns(address)
func (_BootstrapStorage *BootstrapStorageCaller) BEACONORACLEADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "BEACON_ORACLE_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BEACONORACLEADDRESS is a free data retrieval call binding the contract method 0x90a9b42e.
//
// Solidity: function BEACON_ORACLE_ADDRESS() view returns(address)
func (_BootstrapStorage *BootstrapStorageSession) BEACONORACLEADDRESS() (common.Address, error) {
	return _BootstrapStorage.Contract.BEACONORACLEADDRESS(&_BootstrapStorage.CallOpts)
}

// BEACONORACLEADDRESS is a free data retrieval call binding the contract method 0x90a9b42e.
//
// Solidity: function BEACON_ORACLE_ADDRESS() view returns(address)
func (_BootstrapStorage *BootstrapStorageCallerSession) BEACONORACLEADDRESS() (common.Address, error) {
	return _BootstrapStorage.Contract.BEACONORACLEADDRESS(&_BootstrapStorage.CallOpts)
}

// BEACONPROXYBYTECODE is a free data retrieval call binding the contract method 0xcda3ef36.
//
// Solidity: function BEACON_PROXY_BYTECODE() view returns(address)
func (_BootstrapStorage *BootstrapStorageCaller) BEACONPROXYBYTECODE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "BEACON_PROXY_BYTECODE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BEACONPROXYBYTECODE is a free data retrieval call binding the contract method 0xcda3ef36.
//
// Solidity: function BEACON_PROXY_BYTECODE() view returns(address)
func (_BootstrapStorage *BootstrapStorageSession) BEACONPROXYBYTECODE() (common.Address, error) {
	return _BootstrapStorage.Contract.BEACONPROXYBYTECODE(&_BootstrapStorage.CallOpts)
}

// BEACONPROXYBYTECODE is a free data retrieval call binding the contract method 0xcda3ef36.
//
// Solidity: function BEACON_PROXY_BYTECODE() view returns(address)
func (_BootstrapStorage *BootstrapStorageCallerSession) BEACONPROXYBYTECODE() (common.Address, error) {
	return _BootstrapStorage.Contract.BEACONPROXYBYTECODE(&_BootstrapStorage.CallOpts)
}

// IMUACHAINCHAINID is a free data retrieval call binding the contract method 0x1ef3a2df.
//
// Solidity: function IMUACHAIN_CHAIN_ID() view returns(uint32)
func (_BootstrapStorage *BootstrapStorageCaller) IMUACHAINCHAINID(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "IMUACHAIN_CHAIN_ID")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// IMUACHAINCHAINID is a free data retrieval call binding the contract method 0x1ef3a2df.
//
// Solidity: function IMUACHAIN_CHAIN_ID() view returns(uint32)
func (_BootstrapStorage *BootstrapStorageSession) IMUACHAINCHAINID() (uint32, error) {
	return _BootstrapStorage.Contract.IMUACHAINCHAINID(&_BootstrapStorage.CallOpts)
}

// IMUACHAINCHAINID is a free data retrieval call binding the contract method 0x1ef3a2df.
//
// Solidity: function IMUACHAIN_CHAIN_ID() view returns(uint32)
func (_BootstrapStorage *BootstrapStorageCallerSession) IMUACHAINCHAINID() (uint32, error) {
	return _BootstrapStorage.Contract.IMUACHAINCHAINID(&_BootstrapStorage.CallOpts)
}

// IMUAADDRESSPREFIX is a free data retrieval call binding the contract method 0xe09b8274.
//
// Solidity: function IMUA_ADDRESS_PREFIX() view returns(bytes)
func (_BootstrapStorage *BootstrapStorageCaller) IMUAADDRESSPREFIX(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "IMUA_ADDRESS_PREFIX")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// IMUAADDRESSPREFIX is a free data retrieval call binding the contract method 0xe09b8274.
//
// Solidity: function IMUA_ADDRESS_PREFIX() view returns(bytes)
func (_BootstrapStorage *BootstrapStorageSession) IMUAADDRESSPREFIX() ([]byte, error) {
	return _BootstrapStorage.Contract.IMUAADDRESSPREFIX(&_BootstrapStorage.CallOpts)
}

// IMUAADDRESSPREFIX is a free data retrieval call binding the contract method 0xe09b8274.
//
// Solidity: function IMUA_ADDRESS_PREFIX() view returns(bytes)
func (_BootstrapStorage *BootstrapStorageCallerSession) IMUAADDRESSPREFIX() ([]byte, error) {
	return _BootstrapStorage.Contract.IMUAADDRESSPREFIX(&_BootstrapStorage.CallOpts)
}

// IMUACAPSULEBEACON is a free data retrieval call binding the contract method 0xf67fadda.
//
// Solidity: function IMUA_CAPSULE_BEACON() view returns(address)
func (_BootstrapStorage *BootstrapStorageCaller) IMUACAPSULEBEACON(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "IMUA_CAPSULE_BEACON")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IMUACAPSULEBEACON is a free data retrieval call binding the contract method 0xf67fadda.
//
// Solidity: function IMUA_CAPSULE_BEACON() view returns(address)
func (_BootstrapStorage *BootstrapStorageSession) IMUACAPSULEBEACON() (common.Address, error) {
	return _BootstrapStorage.Contract.IMUACAPSULEBEACON(&_BootstrapStorage.CallOpts)
}

// IMUACAPSULEBEACON is a free data retrieval call binding the contract method 0xf67fadda.
//
// Solidity: function IMUA_CAPSULE_BEACON() view returns(address)
func (_BootstrapStorage *BootstrapStorageCallerSession) IMUACAPSULEBEACON() (common.Address, error) {
	return _BootstrapStorage.Contract.IMUACAPSULEBEACON(&_BootstrapStorage.CallOpts)
}

// VAULTBEACON is a free data retrieval call binding the contract method 0x979c75d9.
//
// Solidity: function VAULT_BEACON() view returns(address)
func (_BootstrapStorage *BootstrapStorageCaller) VAULTBEACON(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "VAULT_BEACON")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VAULTBEACON is a free data retrieval call binding the contract method 0x979c75d9.
//
// Solidity: function VAULT_BEACON() view returns(address)
func (_BootstrapStorage *BootstrapStorageSession) VAULTBEACON() (common.Address, error) {
	return _BootstrapStorage.Contract.VAULTBEACON(&_BootstrapStorage.CallOpts)
}

// VAULTBEACON is a free data retrieval call binding the contract method 0x979c75d9.
//
// Solidity: function VAULT_BEACON() view returns(address)
func (_BootstrapStorage *BootstrapStorageCallerSession) VAULTBEACON() (common.Address, error) {
	return _BootstrapStorage.Contract.VAULTBEACON(&_BootstrapStorage.CallOpts)
}

// Bootstrapped is a free data retrieval call binding the contract method 0x35142c8c.
//
// Solidity: function bootstrapped() view returns(bool)
func (_BootstrapStorage *BootstrapStorageCaller) Bootstrapped(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "bootstrapped")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Bootstrapped is a free data retrieval call binding the contract method 0x35142c8c.
//
// Solidity: function bootstrapped() view returns(bool)
func (_BootstrapStorage *BootstrapStorageSession) Bootstrapped() (bool, error) {
	return _BootstrapStorage.Contract.Bootstrapped(&_BootstrapStorage.CallOpts)
}

// Bootstrapped is a free data retrieval call binding the contract method 0x35142c8c.
//
// Solidity: function bootstrapped() view returns(bool)
func (_BootstrapStorage *BootstrapStorageCallerSession) Bootstrapped() (bool, error) {
	return _BootstrapStorage.Contract.Bootstrapped(&_BootstrapStorage.CallOpts)
}

// ClientChainGatewayLogic is a free data retrieval call binding the contract method 0x669aac44.
//
// Solidity: function clientChainGatewayLogic() view returns(address)
func (_BootstrapStorage *BootstrapStorageCaller) ClientChainGatewayLogic(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "clientChainGatewayLogic")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClientChainGatewayLogic is a free data retrieval call binding the contract method 0x669aac44.
//
// Solidity: function clientChainGatewayLogic() view returns(address)
func (_BootstrapStorage *BootstrapStorageSession) ClientChainGatewayLogic() (common.Address, error) {
	return _BootstrapStorage.Contract.ClientChainGatewayLogic(&_BootstrapStorage.CallOpts)
}

// ClientChainGatewayLogic is a free data retrieval call binding the contract method 0x669aac44.
//
// Solidity: function clientChainGatewayLogic() view returns(address)
func (_BootstrapStorage *BootstrapStorageCallerSession) ClientChainGatewayLogic() (common.Address, error) {
	return _BootstrapStorage.Contract.ClientChainGatewayLogic(&_BootstrapStorage.CallOpts)
}

// ClientChainInitializationData is a free data retrieval call binding the contract method 0xf613fbec.
//
// Solidity: function clientChainInitializationData() view returns(bytes)
func (_BootstrapStorage *BootstrapStorageCaller) ClientChainInitializationData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "clientChainInitializationData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// ClientChainInitializationData is a free data retrieval call binding the contract method 0xf613fbec.
//
// Solidity: function clientChainInitializationData() view returns(bytes)
func (_BootstrapStorage *BootstrapStorageSession) ClientChainInitializationData() ([]byte, error) {
	return _BootstrapStorage.Contract.ClientChainInitializationData(&_BootstrapStorage.CallOpts)
}

// ClientChainInitializationData is a free data retrieval call binding the contract method 0xf613fbec.
//
// Solidity: function clientChainInitializationData() view returns(bytes)
func (_BootstrapStorage *BootstrapStorageCallerSession) ClientChainInitializationData() ([]byte, error) {
	return _BootstrapStorage.Contract.ClientChainInitializationData(&_BootstrapStorage.CallOpts)
}

// CommissionEdited is a free data retrieval call binding the contract method 0x31f13d30.
//
// Solidity: function commissionEdited(string imAddress) view returns(bool hasEdited)
func (_BootstrapStorage *BootstrapStorageCaller) CommissionEdited(opts *bind.CallOpts, imAddress string) (bool, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "commissionEdited", imAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CommissionEdited is a free data retrieval call binding the contract method 0x31f13d30.
//
// Solidity: function commissionEdited(string imAddress) view returns(bool hasEdited)
func (_BootstrapStorage *BootstrapStorageSession) CommissionEdited(imAddress string) (bool, error) {
	return _BootstrapStorage.Contract.CommissionEdited(&_BootstrapStorage.CallOpts, imAddress)
}

// CommissionEdited is a free data retrieval call binding the contract method 0x31f13d30.
//
// Solidity: function commissionEdited(string imAddress) view returns(bool hasEdited)
func (_BootstrapStorage *BootstrapStorageCallerSession) CommissionEdited(imAddress string) (bool, error) {
	return _BootstrapStorage.Contract.CommissionEdited(&_BootstrapStorage.CallOpts, imAddress)
}

// ConsensusPublicKeyInUse is a free data retrieval call binding the contract method 0xc489f42f.
//
// Solidity: function consensusPublicKeyInUse(bytes32 consensusKey) view returns(bool used)
func (_BootstrapStorage *BootstrapStorageCaller) ConsensusPublicKeyInUse(opts *bind.CallOpts, consensusKey [32]byte) (bool, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "consensusPublicKeyInUse", consensusKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ConsensusPublicKeyInUse is a free data retrieval call binding the contract method 0xc489f42f.
//
// Solidity: function consensusPublicKeyInUse(bytes32 consensusKey) view returns(bool used)
func (_BootstrapStorage *BootstrapStorageSession) ConsensusPublicKeyInUse(consensusKey [32]byte) (bool, error) {
	return _BootstrapStorage.Contract.ConsensusPublicKeyInUse(&_BootstrapStorage.CallOpts, consensusKey)
}

// ConsensusPublicKeyInUse is a free data retrieval call binding the contract method 0xc489f42f.
//
// Solidity: function consensusPublicKeyInUse(bytes32 consensusKey) view returns(bool used)
func (_BootstrapStorage *BootstrapStorageCallerSession) ConsensusPublicKeyInUse(consensusKey [32]byte) (bool, error) {
	return _BootstrapStorage.Contract.ConsensusPublicKeyInUse(&_BootstrapStorage.CallOpts, consensusKey)
}

// CustomProxyAdmin is a free data retrieval call binding the contract method 0x4caa408d.
//
// Solidity: function customProxyAdmin() view returns(address)
func (_BootstrapStorage *BootstrapStorageCaller) CustomProxyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "customProxyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CustomProxyAdmin is a free data retrieval call binding the contract method 0x4caa408d.
//
// Solidity: function customProxyAdmin() view returns(address)
func (_BootstrapStorage *BootstrapStorageSession) CustomProxyAdmin() (common.Address, error) {
	return _BootstrapStorage.Contract.CustomProxyAdmin(&_BootstrapStorage.CallOpts)
}

// CustomProxyAdmin is a free data retrieval call binding the contract method 0x4caa408d.
//
// Solidity: function customProxyAdmin() view returns(address)
func (_BootstrapStorage *BootstrapStorageCallerSession) CustomProxyAdmin() (common.Address, error) {
	return _BootstrapStorage.Contract.CustomProxyAdmin(&_BootstrapStorage.CallOpts)
}

// Delegations is a free data retrieval call binding the contract method 0xd428dc46.
//
// Solidity: function delegations(address delegator, string imAddress, address tokenAddress) view returns(uint256 amount)
func (_BootstrapStorage *BootstrapStorageCaller) Delegations(opts *bind.CallOpts, delegator common.Address, imAddress string, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "delegations", delegator, imAddress, tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Delegations is a free data retrieval call binding the contract method 0xd428dc46.
//
// Solidity: function delegations(address delegator, string imAddress, address tokenAddress) view returns(uint256 amount)
func (_BootstrapStorage *BootstrapStorageSession) Delegations(delegator common.Address, imAddress string, tokenAddress common.Address) (*big.Int, error) {
	return _BootstrapStorage.Contract.Delegations(&_BootstrapStorage.CallOpts, delegator, imAddress, tokenAddress)
}

// Delegations is a free data retrieval call binding the contract method 0xd428dc46.
//
// Solidity: function delegations(address delegator, string imAddress, address tokenAddress) view returns(uint256 amount)
func (_BootstrapStorage *BootstrapStorageCallerSession) Delegations(delegator common.Address, imAddress string, tokenAddress common.Address) (*big.Int, error) {
	return _BootstrapStorage.Contract.Delegations(&_BootstrapStorage.CallOpts, delegator, imAddress, tokenAddress)
}

// DelegationsByValidator is a free data retrieval call binding the contract method 0x2795ae23.
//
// Solidity: function delegationsByValidator(string imAddress, address tokenAddress) view returns(uint256 amount)
func (_BootstrapStorage *BootstrapStorageCaller) DelegationsByValidator(opts *bind.CallOpts, imAddress string, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "delegationsByValidator", imAddress, tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegationsByValidator is a free data retrieval call binding the contract method 0x2795ae23.
//
// Solidity: function delegationsByValidator(string imAddress, address tokenAddress) view returns(uint256 amount)
func (_BootstrapStorage *BootstrapStorageSession) DelegationsByValidator(imAddress string, tokenAddress common.Address) (*big.Int, error) {
	return _BootstrapStorage.Contract.DelegationsByValidator(&_BootstrapStorage.CallOpts, imAddress, tokenAddress)
}

// DelegationsByValidator is a free data retrieval call binding the contract method 0x2795ae23.
//
// Solidity: function delegationsByValidator(string imAddress, address tokenAddress) view returns(uint256 amount)
func (_BootstrapStorage *BootstrapStorageCallerSession) DelegationsByValidator(imAddress string, tokenAddress common.Address) (*big.Int, error) {
	return _BootstrapStorage.Contract.DelegationsByValidator(&_BootstrapStorage.CallOpts, imAddress, tokenAddress)
}

// Depositors is a free data retrieval call binding the contract method 0xe4b2fb79.
//
// Solidity: function depositors(uint256 ) view returns(address)
func (_BootstrapStorage *BootstrapStorageCaller) Depositors(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "depositors", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Depositors is a free data retrieval call binding the contract method 0xe4b2fb79.
//
// Solidity: function depositors(uint256 ) view returns(address)
func (_BootstrapStorage *BootstrapStorageSession) Depositors(arg0 *big.Int) (common.Address, error) {
	return _BootstrapStorage.Contract.Depositors(&_BootstrapStorage.CallOpts, arg0)
}

// Depositors is a free data retrieval call binding the contract method 0xe4b2fb79.
//
// Solidity: function depositors(uint256 ) view returns(address)
func (_BootstrapStorage *BootstrapStorageCallerSession) Depositors(arg0 *big.Int) (common.Address, error) {
	return _BootstrapStorage.Contract.Depositors(&_BootstrapStorage.CallOpts, arg0)
}

// DepositsByToken is a free data retrieval call binding the contract method 0x6edad9a6.
//
// Solidity: function depositsByToken(address tokenAddress) view returns(uint256 amount)
func (_BootstrapStorage *BootstrapStorageCaller) DepositsByToken(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "depositsByToken", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositsByToken is a free data retrieval call binding the contract method 0x6edad9a6.
//
// Solidity: function depositsByToken(address tokenAddress) view returns(uint256 amount)
func (_BootstrapStorage *BootstrapStorageSession) DepositsByToken(tokenAddress common.Address) (*big.Int, error) {
	return _BootstrapStorage.Contract.DepositsByToken(&_BootstrapStorage.CallOpts, tokenAddress)
}

// DepositsByToken is a free data retrieval call binding the contract method 0x6edad9a6.
//
// Solidity: function depositsByToken(address tokenAddress) view returns(uint256 amount)
func (_BootstrapStorage *BootstrapStorageCallerSession) DepositsByToken(tokenAddress common.Address) (*big.Int, error) {
	return _BootstrapStorage.Contract.DepositsByToken(&_BootstrapStorage.CallOpts, tokenAddress)
}

// EthToImAddress is a free data retrieval call binding the contract method 0x92f026f9.
//
// Solidity: function ethToImAddress(address ethAddress) view returns(string imAddress)
func (_BootstrapStorage *BootstrapStorageCaller) EthToImAddress(opts *bind.CallOpts, ethAddress common.Address) (string, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "ethToImAddress", ethAddress)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// EthToImAddress is a free data retrieval call binding the contract method 0x92f026f9.
//
// Solidity: function ethToImAddress(address ethAddress) view returns(string imAddress)
func (_BootstrapStorage *BootstrapStorageSession) EthToImAddress(ethAddress common.Address) (string, error) {
	return _BootstrapStorage.Contract.EthToImAddress(&_BootstrapStorage.CallOpts, ethAddress)
}

// EthToImAddress is a free data retrieval call binding the contract method 0x92f026f9.
//
// Solidity: function ethToImAddress(address ethAddress) view returns(string imAddress)
func (_BootstrapStorage *BootstrapStorageCallerSession) EthToImAddress(ethAddress common.Address) (string, error) {
	return _BootstrapStorage.Contract.EthToImAddress(&_BootstrapStorage.CallOpts, ethAddress)
}

// InboundNonce is a free data retrieval call binding the contract method 0x632284fd.
//
// Solidity: function inboundNonce(uint32 eid, bytes32 sender) view returns(uint64 nonce)
func (_BootstrapStorage *BootstrapStorageCaller) InboundNonce(opts *bind.CallOpts, eid uint32, sender [32]byte) (uint64, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "inboundNonce", eid, sender)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InboundNonce is a free data retrieval call binding the contract method 0x632284fd.
//
// Solidity: function inboundNonce(uint32 eid, bytes32 sender) view returns(uint64 nonce)
func (_BootstrapStorage *BootstrapStorageSession) InboundNonce(eid uint32, sender [32]byte) (uint64, error) {
	return _BootstrapStorage.Contract.InboundNonce(&_BootstrapStorage.CallOpts, eid, sender)
}

// InboundNonce is a free data retrieval call binding the contract method 0x632284fd.
//
// Solidity: function inboundNonce(uint32 eid, bytes32 sender) view returns(uint64 nonce)
func (_BootstrapStorage *BootstrapStorageCallerSession) InboundNonce(eid uint32, sender [32]byte) (uint64, error) {
	return _BootstrapStorage.Contract.InboundNonce(&_BootstrapStorage.CallOpts, eid, sender)
}

// IsDepositor is a free data retrieval call binding the contract method 0x2f70d1ba.
//
// Solidity: function isDepositor(address depositor) view returns(bool hasDeposited)
func (_BootstrapStorage *BootstrapStorageCaller) IsDepositor(opts *bind.CallOpts, depositor common.Address) (bool, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "isDepositor", depositor)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDepositor is a free data retrieval call binding the contract method 0x2f70d1ba.
//
// Solidity: function isDepositor(address depositor) view returns(bool hasDeposited)
func (_BootstrapStorage *BootstrapStorageSession) IsDepositor(depositor common.Address) (bool, error) {
	return _BootstrapStorage.Contract.IsDepositor(&_BootstrapStorage.CallOpts, depositor)
}

// IsDepositor is a free data retrieval call binding the contract method 0x2f70d1ba.
//
// Solidity: function isDepositor(address depositor) view returns(bool hasDeposited)
func (_BootstrapStorage *BootstrapStorageCallerSession) IsDepositor(depositor common.Address) (bool, error) {
	return _BootstrapStorage.Contract.IsDepositor(&_BootstrapStorage.CallOpts, depositor)
}

// IsValidImAddress is a free data retrieval call binding the contract method 0x72ac3ab6.
//
// Solidity: function isValidImAddress(string addressToValidate) pure returns(bool)
func (_BootstrapStorage *BootstrapStorageCaller) IsValidImAddress(opts *bind.CallOpts, addressToValidate string) (bool, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "isValidImAddress", addressToValidate)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidImAddress is a free data retrieval call binding the contract method 0x72ac3ab6.
//
// Solidity: function isValidImAddress(string addressToValidate) pure returns(bool)
func (_BootstrapStorage *BootstrapStorageSession) IsValidImAddress(addressToValidate string) (bool, error) {
	return _BootstrapStorage.Contract.IsValidImAddress(&_BootstrapStorage.CallOpts, addressToValidate)
}

// IsValidImAddress is a free data retrieval call binding the contract method 0x72ac3ab6.
//
// Solidity: function isValidImAddress(string addressToValidate) pure returns(bool)
func (_BootstrapStorage *BootstrapStorageCallerSession) IsValidImAddress(addressToValidate string) (bool, error) {
	return _BootstrapStorage.Contract.IsValidImAddress(&_BootstrapStorage.CallOpts, addressToValidate)
}

// IsWhitelistedToken is a free data retrieval call binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address token) view returns(bool whitelisted)
func (_BootstrapStorage *BootstrapStorageCaller) IsWhitelistedToken(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "isWhitelistedToken", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWhitelistedToken is a free data retrieval call binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address token) view returns(bool whitelisted)
func (_BootstrapStorage *BootstrapStorageSession) IsWhitelistedToken(token common.Address) (bool, error) {
	return _BootstrapStorage.Contract.IsWhitelistedToken(&_BootstrapStorage.CallOpts, token)
}

// IsWhitelistedToken is a free data retrieval call binding the contract method 0xab37f486.
//
// Solidity: function isWhitelistedToken(address token) view returns(bool whitelisted)
func (_BootstrapStorage *BootstrapStorageCallerSession) IsWhitelistedToken(token common.Address) (bool, error) {
	return _BootstrapStorage.Contract.IsWhitelistedToken(&_BootstrapStorage.CallOpts, token)
}

// OffsetDuration is a free data retrieval call binding the contract method 0xec246031.
//
// Solidity: function offsetDuration() view returns(uint256)
func (_BootstrapStorage *BootstrapStorageCaller) OffsetDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "offsetDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetDuration is a free data retrieval call binding the contract method 0xec246031.
//
// Solidity: function offsetDuration() view returns(uint256)
func (_BootstrapStorage *BootstrapStorageSession) OffsetDuration() (*big.Int, error) {
	return _BootstrapStorage.Contract.OffsetDuration(&_BootstrapStorage.CallOpts)
}

// OffsetDuration is a free data retrieval call binding the contract method 0xec246031.
//
// Solidity: function offsetDuration() view returns(uint256)
func (_BootstrapStorage *BootstrapStorageCallerSession) OffsetDuration() (*big.Int, error) {
	return _BootstrapStorage.Contract.OffsetDuration(&_BootstrapStorage.CallOpts)
}

// OwnerToCapsule is a free data retrieval call binding the contract method 0x098ad41a.
//
// Solidity: function ownerToCapsule(address owner) view returns(address capsule)
func (_BootstrapStorage *BootstrapStorageCaller) OwnerToCapsule(opts *bind.CallOpts, owner common.Address) (common.Address, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "ownerToCapsule", owner)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerToCapsule is a free data retrieval call binding the contract method 0x098ad41a.
//
// Solidity: function ownerToCapsule(address owner) view returns(address capsule)
func (_BootstrapStorage *BootstrapStorageSession) OwnerToCapsule(owner common.Address) (common.Address, error) {
	return _BootstrapStorage.Contract.OwnerToCapsule(&_BootstrapStorage.CallOpts, owner)
}

// OwnerToCapsule is a free data retrieval call binding the contract method 0x098ad41a.
//
// Solidity: function ownerToCapsule(address owner) view returns(address capsule)
func (_BootstrapStorage *BootstrapStorageCallerSession) OwnerToCapsule(owner common.Address) (common.Address, error) {
	return _BootstrapStorage.Contract.OwnerToCapsule(&_BootstrapStorage.CallOpts, owner)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xc306d47d.
//
// Solidity: function registeredValidators(uint256 ) view returns(address)
func (_BootstrapStorage *BootstrapStorageCaller) RegisteredValidators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "registeredValidators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegisteredValidators is a free data retrieval call binding the contract method 0xc306d47d.
//
// Solidity: function registeredValidators(uint256 ) view returns(address)
func (_BootstrapStorage *BootstrapStorageSession) RegisteredValidators(arg0 *big.Int) (common.Address, error) {
	return _BootstrapStorage.Contract.RegisteredValidators(&_BootstrapStorage.CallOpts, arg0)
}

// RegisteredValidators is a free data retrieval call binding the contract method 0xc306d47d.
//
// Solidity: function registeredValidators(uint256 ) view returns(address)
func (_BootstrapStorage *BootstrapStorageCallerSession) RegisteredValidators(arg0 *big.Int) (common.Address, error) {
	return _BootstrapStorage.Contract.RegisteredValidators(&_BootstrapStorage.CallOpts, arg0)
}

// SpawnTime is a free data retrieval call binding the contract method 0x76fd464f.
//
// Solidity: function spawnTime() view returns(uint256)
func (_BootstrapStorage *BootstrapStorageCaller) SpawnTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "spawnTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpawnTime is a free data retrieval call binding the contract method 0x76fd464f.
//
// Solidity: function spawnTime() view returns(uint256)
func (_BootstrapStorage *BootstrapStorageSession) SpawnTime() (*big.Int, error) {
	return _BootstrapStorage.Contract.SpawnTime(&_BootstrapStorage.CallOpts)
}

// SpawnTime is a free data retrieval call binding the contract method 0x76fd464f.
//
// Solidity: function spawnTime() view returns(uint256)
func (_BootstrapStorage *BootstrapStorageCallerSession) SpawnTime() (*big.Int, error) {
	return _BootstrapStorage.Contract.SpawnTime(&_BootstrapStorage.CallOpts)
}

// StakerToPubkeyIDs is a free data retrieval call binding the contract method 0x7e4af2d2.
//
// Solidity: function stakerToPubkeyIDs(address staker, uint256 ) view returns(bytes32)
func (_BootstrapStorage *BootstrapStorageCaller) StakerToPubkeyIDs(opts *bind.CallOpts, staker common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "stakerToPubkeyIDs", staker, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StakerToPubkeyIDs is a free data retrieval call binding the contract method 0x7e4af2d2.
//
// Solidity: function stakerToPubkeyIDs(address staker, uint256 ) view returns(bytes32)
func (_BootstrapStorage *BootstrapStorageSession) StakerToPubkeyIDs(staker common.Address, arg1 *big.Int) ([32]byte, error) {
	return _BootstrapStorage.Contract.StakerToPubkeyIDs(&_BootstrapStorage.CallOpts, staker, arg1)
}

// StakerToPubkeyIDs is a free data retrieval call binding the contract method 0x7e4af2d2.
//
// Solidity: function stakerToPubkeyIDs(address staker, uint256 ) view returns(bytes32)
func (_BootstrapStorage *BootstrapStorageCallerSession) StakerToPubkeyIDs(staker common.Address, arg1 *big.Int) ([32]byte, error) {
	return _BootstrapStorage.Contract.StakerToPubkeyIDs(&_BootstrapStorage.CallOpts, staker, arg1)
}

// StakerToTokenToValidators is a free data retrieval call binding the contract method 0x304884d7.
//
// Solidity: function stakerToTokenToValidators(address staker, address token, uint256 ) view returns(string)
func (_BootstrapStorage *BootstrapStorageCaller) StakerToTokenToValidators(opts *bind.CallOpts, staker common.Address, token common.Address, arg2 *big.Int) (string, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "stakerToTokenToValidators", staker, token, arg2)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StakerToTokenToValidators is a free data retrieval call binding the contract method 0x304884d7.
//
// Solidity: function stakerToTokenToValidators(address staker, address token, uint256 ) view returns(string)
func (_BootstrapStorage *BootstrapStorageSession) StakerToTokenToValidators(staker common.Address, token common.Address, arg2 *big.Int) (string, error) {
	return _BootstrapStorage.Contract.StakerToTokenToValidators(&_BootstrapStorage.CallOpts, staker, token, arg2)
}

// StakerToTokenToValidators is a free data retrieval call binding the contract method 0x304884d7.
//
// Solidity: function stakerToTokenToValidators(address staker, address token, uint256 ) view returns(string)
func (_BootstrapStorage *BootstrapStorageCallerSession) StakerToTokenToValidators(staker common.Address, token common.Address, arg2 *big.Int) (string, error) {
	return _BootstrapStorage.Contract.StakerToTokenToValidators(&_BootstrapStorage.CallOpts, staker, token, arg2)
}

// TokenToVault is a free data retrieval call binding the contract method 0x0c7e1725.
//
// Solidity: function tokenToVault(address token) view returns(address vault)
func (_BootstrapStorage *BootstrapStorageCaller) TokenToVault(opts *bind.CallOpts, token common.Address) (common.Address, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "tokenToVault", token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenToVault is a free data retrieval call binding the contract method 0x0c7e1725.
//
// Solidity: function tokenToVault(address token) view returns(address vault)
func (_BootstrapStorage *BootstrapStorageSession) TokenToVault(token common.Address) (common.Address, error) {
	return _BootstrapStorage.Contract.TokenToVault(&_BootstrapStorage.CallOpts, token)
}

// TokenToVault is a free data retrieval call binding the contract method 0x0c7e1725.
//
// Solidity: function tokenToVault(address token) view returns(address vault)
func (_BootstrapStorage *BootstrapStorageCallerSession) TokenToVault(token common.Address) (common.Address, error) {
	return _BootstrapStorage.Contract.TokenToVault(&_BootstrapStorage.CallOpts, token)
}

// TotalDepositAmounts is a free data retrieval call binding the contract method 0x1af20d25.
//
// Solidity: function totalDepositAmounts(address depositor, address tokenAddress) view returns(uint256 amount)
func (_BootstrapStorage *BootstrapStorageCaller) TotalDepositAmounts(opts *bind.CallOpts, depositor common.Address, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "totalDepositAmounts", depositor, tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDepositAmounts is a free data retrieval call binding the contract method 0x1af20d25.
//
// Solidity: function totalDepositAmounts(address depositor, address tokenAddress) view returns(uint256 amount)
func (_BootstrapStorage *BootstrapStorageSession) TotalDepositAmounts(depositor common.Address, tokenAddress common.Address) (*big.Int, error) {
	return _BootstrapStorage.Contract.TotalDepositAmounts(&_BootstrapStorage.CallOpts, depositor, tokenAddress)
}

// TotalDepositAmounts is a free data retrieval call binding the contract method 0x1af20d25.
//
// Solidity: function totalDepositAmounts(address depositor, address tokenAddress) view returns(uint256 amount)
func (_BootstrapStorage *BootstrapStorageCallerSession) TotalDepositAmounts(depositor common.Address, tokenAddress common.Address) (*big.Int, error) {
	return _BootstrapStorage.Contract.TotalDepositAmounts(&_BootstrapStorage.CallOpts, depositor, tokenAddress)
}

// ValidatorNameInUse is a free data retrieval call binding the contract method 0x598e7fbd.
//
// Solidity: function validatorNameInUse(string name) view returns(bool used)
func (_BootstrapStorage *BootstrapStorageCaller) ValidatorNameInUse(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "validatorNameInUse", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidatorNameInUse is a free data retrieval call binding the contract method 0x598e7fbd.
//
// Solidity: function validatorNameInUse(string name) view returns(bool used)
func (_BootstrapStorage *BootstrapStorageSession) ValidatorNameInUse(name string) (bool, error) {
	return _BootstrapStorage.Contract.ValidatorNameInUse(&_BootstrapStorage.CallOpts, name)
}

// ValidatorNameInUse is a free data retrieval call binding the contract method 0x598e7fbd.
//
// Solidity: function validatorNameInUse(string name) view returns(bool used)
func (_BootstrapStorage *BootstrapStorageCallerSession) ValidatorNameInUse(name string) (bool, error) {
	return _BootstrapStorage.Contract.ValidatorNameInUse(&_BootstrapStorage.CallOpts, name)
}

// Validators is a free data retrieval call binding the contract method 0xfacbd0e3.
//
// Solidity: function validators(string imAddress) view returns(string name, (uint256,uint256,uint256) commission, bytes32 consensusPublicKey)
func (_BootstrapStorage *BootstrapStorageCaller) Validators(opts *bind.CallOpts, imAddress string) (struct {
	Name               string
	Commission         IValidatorRegistryCommission
	ConsensusPublicKey [32]byte
}, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "validators", imAddress)

	outstruct := new(struct {
		Name               string
		Commission         IValidatorRegistryCommission
		ConsensusPublicKey [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Commission = *abi.ConvertType(out[1], new(IValidatorRegistryCommission)).(*IValidatorRegistryCommission)
	outstruct.ConsensusPublicKey = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0xfacbd0e3.
//
// Solidity: function validators(string imAddress) view returns(string name, (uint256,uint256,uint256) commission, bytes32 consensusPublicKey)
func (_BootstrapStorage *BootstrapStorageSession) Validators(imAddress string) (struct {
	Name               string
	Commission         IValidatorRegistryCommission
	ConsensusPublicKey [32]byte
}, error) {
	return _BootstrapStorage.Contract.Validators(&_BootstrapStorage.CallOpts, imAddress)
}

// Validators is a free data retrieval call binding the contract method 0xfacbd0e3.
//
// Solidity: function validators(string imAddress) view returns(string name, (uint256,uint256,uint256) commission, bytes32 consensusPublicKey)
func (_BootstrapStorage *BootstrapStorageCallerSession) Validators(imAddress string) (struct {
	Name               string
	Commission         IValidatorRegistryCommission
	ConsensusPublicKey [32]byte
}, error) {
	return _BootstrapStorage.Contract.Validators(&_BootstrapStorage.CallOpts, imAddress)
}

// WhitelistTokens is a free data retrieval call binding the contract method 0x602a70f1.
//
// Solidity: function whitelistTokens(uint256 ) view returns(address)
func (_BootstrapStorage *BootstrapStorageCaller) WhitelistTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "whitelistTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistTokens is a free data retrieval call binding the contract method 0x602a70f1.
//
// Solidity: function whitelistTokens(uint256 ) view returns(address)
func (_BootstrapStorage *BootstrapStorageSession) WhitelistTokens(arg0 *big.Int) (common.Address, error) {
	return _BootstrapStorage.Contract.WhitelistTokens(&_BootstrapStorage.CallOpts, arg0)
}

// WhitelistTokens is a free data retrieval call binding the contract method 0x602a70f1.
//
// Solidity: function whitelistTokens(uint256 ) view returns(address)
func (_BootstrapStorage *BootstrapStorageCallerSession) WhitelistTokens(arg0 *big.Int) (common.Address, error) {
	return _BootstrapStorage.Contract.WhitelistTokens(&_BootstrapStorage.CallOpts, arg0)
}

// WithdrawableAmounts is a free data retrieval call binding the contract method 0x3a5a6389.
//
// Solidity: function withdrawableAmounts(address depositor, address tokenAddress) view returns(uint256 amount)
func (_BootstrapStorage *BootstrapStorageCaller) WithdrawableAmounts(opts *bind.CallOpts, depositor common.Address, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BootstrapStorage.contract.Call(opts, &out, "withdrawableAmounts", depositor, tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableAmounts is a free data retrieval call binding the contract method 0x3a5a6389.
//
// Solidity: function withdrawableAmounts(address depositor, address tokenAddress) view returns(uint256 amount)
func (_BootstrapStorage *BootstrapStorageSession) WithdrawableAmounts(depositor common.Address, tokenAddress common.Address) (*big.Int, error) {
	return _BootstrapStorage.Contract.WithdrawableAmounts(&_BootstrapStorage.CallOpts, depositor, tokenAddress)
}

// WithdrawableAmounts is a free data retrieval call binding the contract method 0x3a5a6389.
//
// Solidity: function withdrawableAmounts(address depositor, address tokenAddress) view returns(uint256 amount)
func (_BootstrapStorage *BootstrapStorageCallerSession) WithdrawableAmounts(depositor common.Address, tokenAddress common.Address) (*big.Int, error) {
	return _BootstrapStorage.Contract.WithdrawableAmounts(&_BootstrapStorage.CallOpts, depositor, tokenAddress)
}

// BootstrapStorageBootstrapNotTimeYetIterator is returned from FilterBootstrapNotTimeYet and is used to iterate over the raw logs and unpacked data for BootstrapNotTimeYet events raised by the BootstrapStorage contract.
type BootstrapStorageBootstrapNotTimeYetIterator struct {
	Event *BootstrapStorageBootstrapNotTimeYet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapStorageBootstrapNotTimeYetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapStorageBootstrapNotTimeYet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapStorageBootstrapNotTimeYet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapStorageBootstrapNotTimeYetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapStorageBootstrapNotTimeYetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapStorageBootstrapNotTimeYet represents a BootstrapNotTimeYet event raised by the BootstrapStorage contract.
type BootstrapStorageBootstrapNotTimeYet struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBootstrapNotTimeYet is a free log retrieval operation binding the contract event 0xd8bb15fbd9470bdd3978999d7cefc45ea8566f88468b9ca55852fc9cdee48845.
//
// Solidity: event BootstrapNotTimeYet()
func (_BootstrapStorage *BootstrapStorageFilterer) FilterBootstrapNotTimeYet(opts *bind.FilterOpts) (*BootstrapStorageBootstrapNotTimeYetIterator, error) {

	logs, sub, err := _BootstrapStorage.contract.FilterLogs(opts, "BootstrapNotTimeYet")
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageBootstrapNotTimeYetIterator{contract: _BootstrapStorage.contract, event: "BootstrapNotTimeYet", logs: logs, sub: sub}, nil
}

// WatchBootstrapNotTimeYet is a free log subscription operation binding the contract event 0xd8bb15fbd9470bdd3978999d7cefc45ea8566f88468b9ca55852fc9cdee48845.
//
// Solidity: event BootstrapNotTimeYet()
func (_BootstrapStorage *BootstrapStorageFilterer) WatchBootstrapNotTimeYet(opts *bind.WatchOpts, sink chan<- *BootstrapStorageBootstrapNotTimeYet) (event.Subscription, error) {

	logs, sub, err := _BootstrapStorage.contract.WatchLogs(opts, "BootstrapNotTimeYet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapStorageBootstrapNotTimeYet)
				if err := _BootstrapStorage.contract.UnpackLog(event, "BootstrapNotTimeYet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBootstrapNotTimeYet is a log parse operation binding the contract event 0xd8bb15fbd9470bdd3978999d7cefc45ea8566f88468b9ca55852fc9cdee48845.
//
// Solidity: event BootstrapNotTimeYet()
func (_BootstrapStorage *BootstrapStorageFilterer) ParseBootstrapNotTimeYet(log types.Log) (*BootstrapStorageBootstrapNotTimeYet, error) {
	event := new(BootstrapStorageBootstrapNotTimeYet)
	if err := _BootstrapStorage.contract.UnpackLog(event, "BootstrapNotTimeYet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapStorageBootstrapUpgradeFailedIterator is returned from FilterBootstrapUpgradeFailed and is used to iterate over the raw logs and unpacked data for BootstrapUpgradeFailed events raised by the BootstrapStorage contract.
type BootstrapStorageBootstrapUpgradeFailedIterator struct {
	Event *BootstrapStorageBootstrapUpgradeFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapStorageBootstrapUpgradeFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapStorageBootstrapUpgradeFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapStorageBootstrapUpgradeFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapStorageBootstrapUpgradeFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapStorageBootstrapUpgradeFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapStorageBootstrapUpgradeFailed represents a BootstrapUpgradeFailed event raised by the BootstrapStorage contract.
type BootstrapStorageBootstrapUpgradeFailed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBootstrapUpgradeFailed is a free log retrieval operation binding the contract event 0xdcb7e6fbb96e73875566de62d28669c83aff5eb7dfae2271b7f78ab7727d4b84.
//
// Solidity: event BootstrapUpgradeFailed()
func (_BootstrapStorage *BootstrapStorageFilterer) FilterBootstrapUpgradeFailed(opts *bind.FilterOpts) (*BootstrapStorageBootstrapUpgradeFailedIterator, error) {

	logs, sub, err := _BootstrapStorage.contract.FilterLogs(opts, "BootstrapUpgradeFailed")
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageBootstrapUpgradeFailedIterator{contract: _BootstrapStorage.contract, event: "BootstrapUpgradeFailed", logs: logs, sub: sub}, nil
}

// WatchBootstrapUpgradeFailed is a free log subscription operation binding the contract event 0xdcb7e6fbb96e73875566de62d28669c83aff5eb7dfae2271b7f78ab7727d4b84.
//
// Solidity: event BootstrapUpgradeFailed()
func (_BootstrapStorage *BootstrapStorageFilterer) WatchBootstrapUpgradeFailed(opts *bind.WatchOpts, sink chan<- *BootstrapStorageBootstrapUpgradeFailed) (event.Subscription, error) {

	logs, sub, err := _BootstrapStorage.contract.WatchLogs(opts, "BootstrapUpgradeFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapStorageBootstrapUpgradeFailed)
				if err := _BootstrapStorage.contract.UnpackLog(event, "BootstrapUpgradeFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBootstrapUpgradeFailed is a log parse operation binding the contract event 0xdcb7e6fbb96e73875566de62d28669c83aff5eb7dfae2271b7f78ab7727d4b84.
//
// Solidity: event BootstrapUpgradeFailed()
func (_BootstrapStorage *BootstrapStorageFilterer) ParseBootstrapUpgradeFailed(log types.Log) (*BootstrapStorageBootstrapUpgradeFailed, error) {
	event := new(BootstrapStorageBootstrapUpgradeFailed)
	if err := _BootstrapStorage.contract.UnpackLog(event, "BootstrapUpgradeFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapStorageBootstrappedIterator is returned from FilterBootstrapped and is used to iterate over the raw logs and unpacked data for Bootstrapped events raised by the BootstrapStorage contract.
type BootstrapStorageBootstrappedIterator struct {
	Event *BootstrapStorageBootstrapped // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapStorageBootstrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapStorageBootstrapped)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapStorageBootstrapped)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapStorageBootstrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapStorageBootstrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapStorageBootstrapped represents a Bootstrapped event raised by the BootstrapStorage contract.
type BootstrapStorageBootstrapped struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBootstrapped is a free log retrieval operation binding the contract event 0x90af5282d84f2aaaa645bb396ffbb1a1929fda2af47defb09244b31d458a719d.
//
// Solidity: event Bootstrapped()
func (_BootstrapStorage *BootstrapStorageFilterer) FilterBootstrapped(opts *bind.FilterOpts) (*BootstrapStorageBootstrappedIterator, error) {

	logs, sub, err := _BootstrapStorage.contract.FilterLogs(opts, "Bootstrapped")
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageBootstrappedIterator{contract: _BootstrapStorage.contract, event: "Bootstrapped", logs: logs, sub: sub}, nil
}

// WatchBootstrapped is a free log subscription operation binding the contract event 0x90af5282d84f2aaaa645bb396ffbb1a1929fda2af47defb09244b31d458a719d.
//
// Solidity: event Bootstrapped()
func (_BootstrapStorage *BootstrapStorageFilterer) WatchBootstrapped(opts *bind.WatchOpts, sink chan<- *BootstrapStorageBootstrapped) (event.Subscription, error) {

	logs, sub, err := _BootstrapStorage.contract.WatchLogs(opts, "Bootstrapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapStorageBootstrapped)
				if err := _BootstrapStorage.contract.UnpackLog(event, "Bootstrapped", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBootstrapped is a log parse operation binding the contract event 0x90af5282d84f2aaaa645bb396ffbb1a1929fda2af47defb09244b31d458a719d.
//
// Solidity: event Bootstrapped()
func (_BootstrapStorage *BootstrapStorageFilterer) ParseBootstrapped(log types.Log) (*BootstrapStorageBootstrapped, error) {
	event := new(BootstrapStorageBootstrapped)
	if err := _BootstrapStorage.contract.UnpackLog(event, "Bootstrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapStorageBootstrappedAlreadyIterator is returned from FilterBootstrappedAlready and is used to iterate over the raw logs and unpacked data for BootstrappedAlready events raised by the BootstrapStorage contract.
type BootstrapStorageBootstrappedAlreadyIterator struct {
	Event *BootstrapStorageBootstrappedAlready // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapStorageBootstrappedAlreadyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapStorageBootstrappedAlready)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapStorageBootstrappedAlready)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapStorageBootstrappedAlreadyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapStorageBootstrappedAlreadyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapStorageBootstrappedAlready represents a BootstrappedAlready event raised by the BootstrapStorage contract.
type BootstrapStorageBootstrappedAlready struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBootstrappedAlready is a free log retrieval operation binding the contract event 0x6e5e65216df72fd325095b47d04105d3ca27805ca44babb89b131cb8640892ff.
//
// Solidity: event BootstrappedAlready()
func (_BootstrapStorage *BootstrapStorageFilterer) FilterBootstrappedAlready(opts *bind.FilterOpts) (*BootstrapStorageBootstrappedAlreadyIterator, error) {

	logs, sub, err := _BootstrapStorage.contract.FilterLogs(opts, "BootstrappedAlready")
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageBootstrappedAlreadyIterator{contract: _BootstrapStorage.contract, event: "BootstrappedAlready", logs: logs, sub: sub}, nil
}

// WatchBootstrappedAlready is a free log subscription operation binding the contract event 0x6e5e65216df72fd325095b47d04105d3ca27805ca44babb89b131cb8640892ff.
//
// Solidity: event BootstrappedAlready()
func (_BootstrapStorage *BootstrapStorageFilterer) WatchBootstrappedAlready(opts *bind.WatchOpts, sink chan<- *BootstrapStorageBootstrappedAlready) (event.Subscription, error) {

	logs, sub, err := _BootstrapStorage.contract.WatchLogs(opts, "BootstrappedAlready")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapStorageBootstrappedAlready)
				if err := _BootstrapStorage.contract.UnpackLog(event, "BootstrappedAlready", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBootstrappedAlready is a log parse operation binding the contract event 0x6e5e65216df72fd325095b47d04105d3ca27805ca44babb89b131cb8640892ff.
//
// Solidity: event BootstrappedAlready()
func (_BootstrapStorage *BootstrapStorageFilterer) ParseBootstrappedAlready(log types.Log) (*BootstrapStorageBootstrappedAlready, error) {
	event := new(BootstrapStorageBootstrappedAlready)
	if err := _BootstrapStorage.contract.UnpackLog(event, "BootstrappedAlready", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapStorageCapsuleCreatedIterator is returned from FilterCapsuleCreated and is used to iterate over the raw logs and unpacked data for CapsuleCreated events raised by the BootstrapStorage contract.
type BootstrapStorageCapsuleCreatedIterator struct {
	Event *BootstrapStorageCapsuleCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapStorageCapsuleCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapStorageCapsuleCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapStorageCapsuleCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapStorageCapsuleCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapStorageCapsuleCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapStorageCapsuleCreated represents a CapsuleCreated event raised by the BootstrapStorage contract.
type BootstrapStorageCapsuleCreated struct {
	Owner   common.Address
	Capsule common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCapsuleCreated is a free log retrieval operation binding the contract event 0x9b7b96ddfa41a5d76e277b661fe9d90e452022c554009c095d97269a615a11a3.
//
// Solidity: event CapsuleCreated(address indexed owner, address indexed capsule)
func (_BootstrapStorage *BootstrapStorageFilterer) FilterCapsuleCreated(opts *bind.FilterOpts, owner []common.Address, capsule []common.Address) (*BootstrapStorageCapsuleCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var capsuleRule []interface{}
	for _, capsuleItem := range capsule {
		capsuleRule = append(capsuleRule, capsuleItem)
	}

	logs, sub, err := _BootstrapStorage.contract.FilterLogs(opts, "CapsuleCreated", ownerRule, capsuleRule)
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageCapsuleCreatedIterator{contract: _BootstrapStorage.contract, event: "CapsuleCreated", logs: logs, sub: sub}, nil
}

// WatchCapsuleCreated is a free log subscription operation binding the contract event 0x9b7b96ddfa41a5d76e277b661fe9d90e452022c554009c095d97269a615a11a3.
//
// Solidity: event CapsuleCreated(address indexed owner, address indexed capsule)
func (_BootstrapStorage *BootstrapStorageFilterer) WatchCapsuleCreated(opts *bind.WatchOpts, sink chan<- *BootstrapStorageCapsuleCreated, owner []common.Address, capsule []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var capsuleRule []interface{}
	for _, capsuleItem := range capsule {
		capsuleRule = append(capsuleRule, capsuleItem)
	}

	logs, sub, err := _BootstrapStorage.contract.WatchLogs(opts, "CapsuleCreated", ownerRule, capsuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapStorageCapsuleCreated)
				if err := _BootstrapStorage.contract.UnpackLog(event, "CapsuleCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCapsuleCreated is a log parse operation binding the contract event 0x9b7b96ddfa41a5d76e277b661fe9d90e452022c554009c095d97269a615a11a3.
//
// Solidity: event CapsuleCreated(address indexed owner, address indexed capsule)
func (_BootstrapStorage *BootstrapStorageFilterer) ParseCapsuleCreated(log types.Log) (*BootstrapStorageCapsuleCreated, error) {
	event := new(BootstrapStorageCapsuleCreated)
	if err := _BootstrapStorage.contract.UnpackLog(event, "CapsuleCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapStorageClaimPrincipalResultIterator is returned from FilterClaimPrincipalResult and is used to iterate over the raw logs and unpacked data for ClaimPrincipalResult events raised by the BootstrapStorage contract.
type BootstrapStorageClaimPrincipalResultIterator struct {
	Event *BootstrapStorageClaimPrincipalResult // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapStorageClaimPrincipalResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapStorageClaimPrincipalResult)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapStorageClaimPrincipalResult)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapStorageClaimPrincipalResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapStorageClaimPrincipalResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapStorageClaimPrincipalResult represents a ClaimPrincipalResult event raised by the BootstrapStorage contract.
type BootstrapStorageClaimPrincipalResult struct {
	Success    bool
	Token      common.Address
	Withdrawer common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimPrincipalResult is a free log retrieval operation binding the contract event 0xeca0f64c3bab2d9808a4cb927796dd3c69892b2a9ae47cb5a078d6a4405f27f5.
//
// Solidity: event ClaimPrincipalResult(bool indexed success, address indexed token, address indexed withdrawer, uint256 amount)
func (_BootstrapStorage *BootstrapStorageFilterer) FilterClaimPrincipalResult(opts *bind.FilterOpts, success []bool, token []common.Address, withdrawer []common.Address) (*BootstrapStorageClaimPrincipalResultIterator, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _BootstrapStorage.contract.FilterLogs(opts, "ClaimPrincipalResult", successRule, tokenRule, withdrawerRule)
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageClaimPrincipalResultIterator{contract: _BootstrapStorage.contract, event: "ClaimPrincipalResult", logs: logs, sub: sub}, nil
}

// WatchClaimPrincipalResult is a free log subscription operation binding the contract event 0xeca0f64c3bab2d9808a4cb927796dd3c69892b2a9ae47cb5a078d6a4405f27f5.
//
// Solidity: event ClaimPrincipalResult(bool indexed success, address indexed token, address indexed withdrawer, uint256 amount)
func (_BootstrapStorage *BootstrapStorageFilterer) WatchClaimPrincipalResult(opts *bind.WatchOpts, sink chan<- *BootstrapStorageClaimPrincipalResult, success []bool, token []common.Address, withdrawer []common.Address) (event.Subscription, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}

	logs, sub, err := _BootstrapStorage.contract.WatchLogs(opts, "ClaimPrincipalResult", successRule, tokenRule, withdrawerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapStorageClaimPrincipalResult)
				if err := _BootstrapStorage.contract.UnpackLog(event, "ClaimPrincipalResult", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimPrincipalResult is a log parse operation binding the contract event 0xeca0f64c3bab2d9808a4cb927796dd3c69892b2a9ae47cb5a078d6a4405f27f5.
//
// Solidity: event ClaimPrincipalResult(bool indexed success, address indexed token, address indexed withdrawer, uint256 amount)
func (_BootstrapStorage *BootstrapStorageFilterer) ParseClaimPrincipalResult(log types.Log) (*BootstrapStorageClaimPrincipalResult, error) {
	event := new(BootstrapStorageClaimPrincipalResult)
	if err := _BootstrapStorage.contract.UnpackLog(event, "ClaimPrincipalResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapStorageClientChainGatewayLogicUpdatedIterator is returned from FilterClientChainGatewayLogicUpdated and is used to iterate over the raw logs and unpacked data for ClientChainGatewayLogicUpdated events raised by the BootstrapStorage contract.
type BootstrapStorageClientChainGatewayLogicUpdatedIterator struct {
	Event *BootstrapStorageClientChainGatewayLogicUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapStorageClientChainGatewayLogicUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapStorageClientChainGatewayLogicUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapStorageClientChainGatewayLogicUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapStorageClientChainGatewayLogicUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapStorageClientChainGatewayLogicUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapStorageClientChainGatewayLogicUpdated represents a ClientChainGatewayLogicUpdated event raised by the BootstrapStorage contract.
type BootstrapStorageClientChainGatewayLogicUpdated struct {
	NewLogic           common.Address
	InitializationData []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterClientChainGatewayLogicUpdated is a free log retrieval operation binding the contract event 0xfef38cb604939fc77b2392a748b93e1d47d9c7220e77266894c4fe56c585564d.
//
// Solidity: event ClientChainGatewayLogicUpdated(address newLogic, bytes initializationData)
func (_BootstrapStorage *BootstrapStorageFilterer) FilterClientChainGatewayLogicUpdated(opts *bind.FilterOpts) (*BootstrapStorageClientChainGatewayLogicUpdatedIterator, error) {

	logs, sub, err := _BootstrapStorage.contract.FilterLogs(opts, "ClientChainGatewayLogicUpdated")
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageClientChainGatewayLogicUpdatedIterator{contract: _BootstrapStorage.contract, event: "ClientChainGatewayLogicUpdated", logs: logs, sub: sub}, nil
}

// WatchClientChainGatewayLogicUpdated is a free log subscription operation binding the contract event 0xfef38cb604939fc77b2392a748b93e1d47d9c7220e77266894c4fe56c585564d.
//
// Solidity: event ClientChainGatewayLogicUpdated(address newLogic, bytes initializationData)
func (_BootstrapStorage *BootstrapStorageFilterer) WatchClientChainGatewayLogicUpdated(opts *bind.WatchOpts, sink chan<- *BootstrapStorageClientChainGatewayLogicUpdated) (event.Subscription, error) {

	logs, sub, err := _BootstrapStorage.contract.WatchLogs(opts, "ClientChainGatewayLogicUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapStorageClientChainGatewayLogicUpdated)
				if err := _BootstrapStorage.contract.UnpackLog(event, "ClientChainGatewayLogicUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClientChainGatewayLogicUpdated is a log parse operation binding the contract event 0xfef38cb604939fc77b2392a748b93e1d47d9c7220e77266894c4fe56c585564d.
//
// Solidity: event ClientChainGatewayLogicUpdated(address newLogic, bytes initializationData)
func (_BootstrapStorage *BootstrapStorageFilterer) ParseClientChainGatewayLogicUpdated(log types.Log) (*BootstrapStorageClientChainGatewayLogicUpdated, error) {
	event := new(BootstrapStorageClientChainGatewayLogicUpdated)
	if err := _BootstrapStorage.contract.UnpackLog(event, "ClientChainGatewayLogicUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapStorageDelegateResultIterator is returned from FilterDelegateResult and is used to iterate over the raw logs and unpacked data for DelegateResult events raised by the BootstrapStorage contract.
type BootstrapStorageDelegateResultIterator struct {
	Event *BootstrapStorageDelegateResult // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapStorageDelegateResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapStorageDelegateResult)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapStorageDelegateResult)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapStorageDelegateResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapStorageDelegateResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapStorageDelegateResult represents a DelegateResult event raised by the BootstrapStorage contract.
type BootstrapStorageDelegateResult struct {
	Success   bool
	Delegator common.Address
	Delegatee common.Hash
	Token     common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegateResult is a free log retrieval operation binding the contract event 0xdc4fbf33793e5efc89a1c209786465733043b48c000e0d0611a72abe9a871bbc.
//
// Solidity: event DelegateResult(bool indexed success, address indexed delegator, string indexed delegatee, address token, uint256 amount)
func (_BootstrapStorage *BootstrapStorageFilterer) FilterDelegateResult(opts *bind.FilterOpts, success []bool, delegator []common.Address, delegatee []string) (*BootstrapStorageDelegateResultIterator, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _BootstrapStorage.contract.FilterLogs(opts, "DelegateResult", successRule, delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageDelegateResultIterator{contract: _BootstrapStorage.contract, event: "DelegateResult", logs: logs, sub: sub}, nil
}

// WatchDelegateResult is a free log subscription operation binding the contract event 0xdc4fbf33793e5efc89a1c209786465733043b48c000e0d0611a72abe9a871bbc.
//
// Solidity: event DelegateResult(bool indexed success, address indexed delegator, string indexed delegatee, address token, uint256 amount)
func (_BootstrapStorage *BootstrapStorageFilterer) WatchDelegateResult(opts *bind.WatchOpts, sink chan<- *BootstrapStorageDelegateResult, success []bool, delegator []common.Address, delegatee []string) (event.Subscription, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _BootstrapStorage.contract.WatchLogs(opts, "DelegateResult", successRule, delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapStorageDelegateResult)
				if err := _BootstrapStorage.contract.UnpackLog(event, "DelegateResult", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegateResult is a log parse operation binding the contract event 0xdc4fbf33793e5efc89a1c209786465733043b48c000e0d0611a72abe9a871bbc.
//
// Solidity: event DelegateResult(bool indexed success, address indexed delegator, string indexed delegatee, address token, uint256 amount)
func (_BootstrapStorage *BootstrapStorageFilterer) ParseDelegateResult(log types.Log) (*BootstrapStorageDelegateResult, error) {
	event := new(BootstrapStorageDelegateResult)
	if err := _BootstrapStorage.contract.UnpackLog(event, "DelegateResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapStorageDepositResultIterator is returned from FilterDepositResult and is used to iterate over the raw logs and unpacked data for DepositResult events raised by the BootstrapStorage contract.
type BootstrapStorageDepositResultIterator struct {
	Event *BootstrapStorageDepositResult // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapStorageDepositResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapStorageDepositResult)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapStorageDepositResult)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapStorageDepositResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapStorageDepositResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapStorageDepositResult represents a DepositResult event raised by the BootstrapStorage contract.
type BootstrapStorageDepositResult struct {
	Success   bool
	Token     common.Address
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDepositResult is a free log retrieval operation binding the contract event 0x96a7936e872abc3ec0271aa99dce327d8625f783cb4c455a36a7cd848d9ff1a9.
//
// Solidity: event DepositResult(bool indexed success, address indexed token, address indexed depositor, uint256 amount)
func (_BootstrapStorage *BootstrapStorageFilterer) FilterDepositResult(opts *bind.FilterOpts, success []bool, token []common.Address, depositor []common.Address) (*BootstrapStorageDepositResultIterator, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _BootstrapStorage.contract.FilterLogs(opts, "DepositResult", successRule, tokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageDepositResultIterator{contract: _BootstrapStorage.contract, event: "DepositResult", logs: logs, sub: sub}, nil
}

// WatchDepositResult is a free log subscription operation binding the contract event 0x96a7936e872abc3ec0271aa99dce327d8625f783cb4c455a36a7cd848d9ff1a9.
//
// Solidity: event DepositResult(bool indexed success, address indexed token, address indexed depositor, uint256 amount)
func (_BootstrapStorage *BootstrapStorageFilterer) WatchDepositResult(opts *bind.WatchOpts, sink chan<- *BootstrapStorageDepositResult, success []bool, token []common.Address, depositor []common.Address) (event.Subscription, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _BootstrapStorage.contract.WatchLogs(opts, "DepositResult", successRule, tokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapStorageDepositResult)
				if err := _BootstrapStorage.contract.UnpackLog(event, "DepositResult", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositResult is a log parse operation binding the contract event 0x96a7936e872abc3ec0271aa99dce327d8625f783cb4c455a36a7cd848d9ff1a9.
//
// Solidity: event DepositResult(bool indexed success, address indexed token, address indexed depositor, uint256 amount)
func (_BootstrapStorage *BootstrapStorageFilterer) ParseDepositResult(log types.Log) (*BootstrapStorageDepositResult, error) {
	event := new(BootstrapStorageDepositResult)
	if err := _BootstrapStorage.contract.UnpackLog(event, "DepositResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapStorageDepositThenDelegateResultIterator is returned from FilterDepositThenDelegateResult and is used to iterate over the raw logs and unpacked data for DepositThenDelegateResult events raised by the BootstrapStorage contract.
type BootstrapStorageDepositThenDelegateResultIterator struct {
	Event *BootstrapStorageDepositThenDelegateResult // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapStorageDepositThenDelegateResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapStorageDepositThenDelegateResult)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapStorageDepositThenDelegateResult)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapStorageDepositThenDelegateResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapStorageDepositThenDelegateResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapStorageDepositThenDelegateResult represents a DepositThenDelegateResult event raised by the BootstrapStorage contract.
type BootstrapStorageDepositThenDelegateResult struct {
	DelegateSuccess bool
	Delegator       common.Address
	Delegatee       common.Hash
	Token           common.Address
	DelegatedAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDepositThenDelegateResult is a free log retrieval operation binding the contract event 0x7a8415b264ee200190fa69106c4629c425ee8398b8bb4bfec724b8de85cff5f8.
//
// Solidity: event DepositThenDelegateResult(bool indexed delegateSuccess, address indexed delegator, string indexed delegatee, address token, uint256 delegatedAmount)
func (_BootstrapStorage *BootstrapStorageFilterer) FilterDepositThenDelegateResult(opts *bind.FilterOpts, delegateSuccess []bool, delegator []common.Address, delegatee []string) (*BootstrapStorageDepositThenDelegateResultIterator, error) {

	var delegateSuccessRule []interface{}
	for _, delegateSuccessItem := range delegateSuccess {
		delegateSuccessRule = append(delegateSuccessRule, delegateSuccessItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _BootstrapStorage.contract.FilterLogs(opts, "DepositThenDelegateResult", delegateSuccessRule, delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageDepositThenDelegateResultIterator{contract: _BootstrapStorage.contract, event: "DepositThenDelegateResult", logs: logs, sub: sub}, nil
}

// WatchDepositThenDelegateResult is a free log subscription operation binding the contract event 0x7a8415b264ee200190fa69106c4629c425ee8398b8bb4bfec724b8de85cff5f8.
//
// Solidity: event DepositThenDelegateResult(bool indexed delegateSuccess, address indexed delegator, string indexed delegatee, address token, uint256 delegatedAmount)
func (_BootstrapStorage *BootstrapStorageFilterer) WatchDepositThenDelegateResult(opts *bind.WatchOpts, sink chan<- *BootstrapStorageDepositThenDelegateResult, delegateSuccess []bool, delegator []common.Address, delegatee []string) (event.Subscription, error) {

	var delegateSuccessRule []interface{}
	for _, delegateSuccessItem := range delegateSuccess {
		delegateSuccessRule = append(delegateSuccessRule, delegateSuccessItem)
	}
	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var delegateeRule []interface{}
	for _, delegateeItem := range delegatee {
		delegateeRule = append(delegateeRule, delegateeItem)
	}

	logs, sub, err := _BootstrapStorage.contract.WatchLogs(opts, "DepositThenDelegateResult", delegateSuccessRule, delegatorRule, delegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapStorageDepositThenDelegateResult)
				if err := _BootstrapStorage.contract.UnpackLog(event, "DepositThenDelegateResult", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositThenDelegateResult is a log parse operation binding the contract event 0x7a8415b264ee200190fa69106c4629c425ee8398b8bb4bfec724b8de85cff5f8.
//
// Solidity: event DepositThenDelegateResult(bool indexed delegateSuccess, address indexed delegator, string indexed delegatee, address token, uint256 delegatedAmount)
func (_BootstrapStorage *BootstrapStorageFilterer) ParseDepositThenDelegateResult(log types.Log) (*BootstrapStorageDepositThenDelegateResult, error) {
	event := new(BootstrapStorageDepositThenDelegateResult)
	if err := _BootstrapStorage.contract.UnpackLog(event, "DepositThenDelegateResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapStorageMessageExecutedIterator is returned from FilterMessageExecuted and is used to iterate over the raw logs and unpacked data for MessageExecuted events raised by the BootstrapStorage contract.
type BootstrapStorageMessageExecutedIterator struct {
	Event *BootstrapStorageMessageExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapStorageMessageExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapStorageMessageExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapStorageMessageExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapStorageMessageExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapStorageMessageExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapStorageMessageExecuted represents a MessageExecuted event raised by the BootstrapStorage contract.
type BootstrapStorageMessageExecuted struct {
	Act   uint8
	Nonce uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMessageExecuted is a free log retrieval operation binding the contract event 0xb82f0bf216aec4c87941891caba2e306587945a146feb0cbcd67b55d22b88501.
//
// Solidity: event MessageExecuted(uint8 indexed act, uint64 nonce)
func (_BootstrapStorage *BootstrapStorageFilterer) FilterMessageExecuted(opts *bind.FilterOpts, act []uint8) (*BootstrapStorageMessageExecutedIterator, error) {

	var actRule []interface{}
	for _, actItem := range act {
		actRule = append(actRule, actItem)
	}

	logs, sub, err := _BootstrapStorage.contract.FilterLogs(opts, "MessageExecuted", actRule)
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageMessageExecutedIterator{contract: _BootstrapStorage.contract, event: "MessageExecuted", logs: logs, sub: sub}, nil
}

// WatchMessageExecuted is a free log subscription operation binding the contract event 0xb82f0bf216aec4c87941891caba2e306587945a146feb0cbcd67b55d22b88501.
//
// Solidity: event MessageExecuted(uint8 indexed act, uint64 nonce)
func (_BootstrapStorage *BootstrapStorageFilterer) WatchMessageExecuted(opts *bind.WatchOpts, sink chan<- *BootstrapStorageMessageExecuted, act []uint8) (event.Subscription, error) {

	var actRule []interface{}
	for _, actItem := range act {
		actRule = append(actRule, actItem)
	}

	logs, sub, err := _BootstrapStorage.contract.WatchLogs(opts, "MessageExecuted", actRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapStorageMessageExecuted)
				if err := _BootstrapStorage.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageExecuted is a log parse operation binding the contract event 0xb82f0bf216aec4c87941891caba2e306587945a146feb0cbcd67b55d22b88501.
//
// Solidity: event MessageExecuted(uint8 indexed act, uint64 nonce)
func (_BootstrapStorage *BootstrapStorageFilterer) ParseMessageExecuted(log types.Log) (*BootstrapStorageMessageExecuted, error) {
	event := new(BootstrapStorageMessageExecuted)
	if err := _BootstrapStorage.contract.UnpackLog(event, "MessageExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapStorageMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the BootstrapStorage contract.
type BootstrapStorageMessageSentIterator struct {
	Event *BootstrapStorageMessageSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapStorageMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapStorageMessageSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapStorageMessageSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapStorageMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapStorageMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapStorageMessageSent represents a MessageSent event raised by the BootstrapStorage contract.
type BootstrapStorageMessageSent struct {
	Act       uint8
	PacketId  [32]byte
	Nonce     uint64
	NativeFee *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0xfba7df935971916019b62d3b0eee5d6989cdb9f8081346465209d2887e4c0968.
//
// Solidity: event MessageSent(uint8 indexed act, bytes32 packetId, uint64 nonce, uint256 nativeFee)
func (_BootstrapStorage *BootstrapStorageFilterer) FilterMessageSent(opts *bind.FilterOpts, act []uint8) (*BootstrapStorageMessageSentIterator, error) {

	var actRule []interface{}
	for _, actItem := range act {
		actRule = append(actRule, actItem)
	}

	logs, sub, err := _BootstrapStorage.contract.FilterLogs(opts, "MessageSent", actRule)
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageMessageSentIterator{contract: _BootstrapStorage.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0xfba7df935971916019b62d3b0eee5d6989cdb9f8081346465209d2887e4c0968.
//
// Solidity: event MessageSent(uint8 indexed act, bytes32 packetId, uint64 nonce, uint256 nativeFee)
func (_BootstrapStorage *BootstrapStorageFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *BootstrapStorageMessageSent, act []uint8) (event.Subscription, error) {

	var actRule []interface{}
	for _, actItem := range act {
		actRule = append(actRule, actItem)
	}

	logs, sub, err := _BootstrapStorage.contract.WatchLogs(opts, "MessageSent", actRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapStorageMessageSent)
				if err := _BootstrapStorage.contract.UnpackLog(event, "MessageSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMessageSent is a log parse operation binding the contract event 0xfba7df935971916019b62d3b0eee5d6989cdb9f8081346465209d2887e4c0968.
//
// Solidity: event MessageSent(uint8 indexed act, bytes32 packetId, uint64 nonce, uint256 nativeFee)
func (_BootstrapStorage *BootstrapStorageFilterer) ParseMessageSent(log types.Log) (*BootstrapStorageMessageSent, error) {
	event := new(BootstrapStorageMessageSent)
	if err := _BootstrapStorage.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapStorageOffsetDurationUpdatedIterator is returned from FilterOffsetDurationUpdated and is used to iterate over the raw logs and unpacked data for OffsetDurationUpdated events raised by the BootstrapStorage contract.
type BootstrapStorageOffsetDurationUpdatedIterator struct {
	Event *BootstrapStorageOffsetDurationUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapStorageOffsetDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapStorageOffsetDurationUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapStorageOffsetDurationUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapStorageOffsetDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapStorageOffsetDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapStorageOffsetDurationUpdated represents a OffsetDurationUpdated event raised by the BootstrapStorage contract.
type BootstrapStorageOffsetDurationUpdated struct {
	NewOffsetDuration *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOffsetDurationUpdated is a free log retrieval operation binding the contract event 0x33ea0c40bbb70e43d4a933b05cb895a4a5f876debd3b8ed39c3c4d1e016f4ad3.
//
// Solidity: event OffsetDurationUpdated(uint256 newOffsetDuration)
func (_BootstrapStorage *BootstrapStorageFilterer) FilterOffsetDurationUpdated(opts *bind.FilterOpts) (*BootstrapStorageOffsetDurationUpdatedIterator, error) {

	logs, sub, err := _BootstrapStorage.contract.FilterLogs(opts, "OffsetDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageOffsetDurationUpdatedIterator{contract: _BootstrapStorage.contract, event: "OffsetDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchOffsetDurationUpdated is a free log subscription operation binding the contract event 0x33ea0c40bbb70e43d4a933b05cb895a4a5f876debd3b8ed39c3c4d1e016f4ad3.
//
// Solidity: event OffsetDurationUpdated(uint256 newOffsetDuration)
func (_BootstrapStorage *BootstrapStorageFilterer) WatchOffsetDurationUpdated(opts *bind.WatchOpts, sink chan<- *BootstrapStorageOffsetDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _BootstrapStorage.contract.WatchLogs(opts, "OffsetDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapStorageOffsetDurationUpdated)
				if err := _BootstrapStorage.contract.UnpackLog(event, "OffsetDurationUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOffsetDurationUpdated is a log parse operation binding the contract event 0x33ea0c40bbb70e43d4a933b05cb895a4a5f876debd3b8ed39c3c4d1e016f4ad3.
//
// Solidity: event OffsetDurationUpdated(uint256 newOffsetDuration)
func (_BootstrapStorage *BootstrapStorageFilterer) ParseOffsetDurationUpdated(log types.Log) (*BootstrapStorageOffsetDurationUpdated, error) {
	event := new(BootstrapStorageOffsetDurationUpdated)
	if err := _BootstrapStorage.contract.UnpackLog(event, "OffsetDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapStorageSpawnTimeUpdatedIterator is returned from FilterSpawnTimeUpdated and is used to iterate over the raw logs and unpacked data for SpawnTimeUpdated events raised by the BootstrapStorage contract.
type BootstrapStorageSpawnTimeUpdatedIterator struct {
	Event *BootstrapStorageSpawnTimeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapStorageSpawnTimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapStorageSpawnTimeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapStorageSpawnTimeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapStorageSpawnTimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapStorageSpawnTimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapStorageSpawnTimeUpdated represents a SpawnTimeUpdated event raised by the BootstrapStorage contract.
type BootstrapStorageSpawnTimeUpdated struct {
	NewSpawnTime *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSpawnTimeUpdated is a free log retrieval operation binding the contract event 0xa2d6431b356061bf0720994ce72f8bb72d17ed14a040f92c3600c55b857095db.
//
// Solidity: event SpawnTimeUpdated(uint256 newSpawnTime)
func (_BootstrapStorage *BootstrapStorageFilterer) FilterSpawnTimeUpdated(opts *bind.FilterOpts) (*BootstrapStorageSpawnTimeUpdatedIterator, error) {

	logs, sub, err := _BootstrapStorage.contract.FilterLogs(opts, "SpawnTimeUpdated")
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageSpawnTimeUpdatedIterator{contract: _BootstrapStorage.contract, event: "SpawnTimeUpdated", logs: logs, sub: sub}, nil
}

// WatchSpawnTimeUpdated is a free log subscription operation binding the contract event 0xa2d6431b356061bf0720994ce72f8bb72d17ed14a040f92c3600c55b857095db.
//
// Solidity: event SpawnTimeUpdated(uint256 newSpawnTime)
func (_BootstrapStorage *BootstrapStorageFilterer) WatchSpawnTimeUpdated(opts *bind.WatchOpts, sink chan<- *BootstrapStorageSpawnTimeUpdated) (event.Subscription, error) {

	logs, sub, err := _BootstrapStorage.contract.WatchLogs(opts, "SpawnTimeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapStorageSpawnTimeUpdated)
				if err := _BootstrapStorage.contract.UnpackLog(event, "SpawnTimeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSpawnTimeUpdated is a log parse operation binding the contract event 0xa2d6431b356061bf0720994ce72f8bb72d17ed14a040f92c3600c55b857095db.
//
// Solidity: event SpawnTimeUpdated(uint256 newSpawnTime)
func (_BootstrapStorage *BootstrapStorageFilterer) ParseSpawnTimeUpdated(log types.Log) (*BootstrapStorageSpawnTimeUpdated, error) {
	event := new(BootstrapStorageSpawnTimeUpdated)
	if err := _BootstrapStorage.contract.UnpackLog(event, "SpawnTimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapStorageStakedWithCapsuleIterator is returned from FilterStakedWithCapsule and is used to iterate over the raw logs and unpacked data for StakedWithCapsule events raised by the BootstrapStorage contract.
type BootstrapStorageStakedWithCapsuleIterator struct {
	Event *BootstrapStorageStakedWithCapsule // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapStorageStakedWithCapsuleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapStorageStakedWithCapsule)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapStorageStakedWithCapsule)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapStorageStakedWithCapsuleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapStorageStakedWithCapsuleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapStorageStakedWithCapsule represents a StakedWithCapsule event raised by the BootstrapStorage contract.
type BootstrapStorageStakedWithCapsule struct {
	Staker  common.Address
	Capsule common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakedWithCapsule is a free log retrieval operation binding the contract event 0x9baefaaa262c2c2f0c3acf8ee9293d5043ef654d71c2aba83024f1322c72e5ee.
//
// Solidity: event StakedWithCapsule(address indexed staker, address indexed capsule)
func (_BootstrapStorage *BootstrapStorageFilterer) FilterStakedWithCapsule(opts *bind.FilterOpts, staker []common.Address, capsule []common.Address) (*BootstrapStorageStakedWithCapsuleIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var capsuleRule []interface{}
	for _, capsuleItem := range capsule {
		capsuleRule = append(capsuleRule, capsuleItem)
	}

	logs, sub, err := _BootstrapStorage.contract.FilterLogs(opts, "StakedWithCapsule", stakerRule, capsuleRule)
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageStakedWithCapsuleIterator{contract: _BootstrapStorage.contract, event: "StakedWithCapsule", logs: logs, sub: sub}, nil
}

// WatchStakedWithCapsule is a free log subscription operation binding the contract event 0x9baefaaa262c2c2f0c3acf8ee9293d5043ef654d71c2aba83024f1322c72e5ee.
//
// Solidity: event StakedWithCapsule(address indexed staker, address indexed capsule)
func (_BootstrapStorage *BootstrapStorageFilterer) WatchStakedWithCapsule(opts *bind.WatchOpts, sink chan<- *BootstrapStorageStakedWithCapsule, staker []common.Address, capsule []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var capsuleRule []interface{}
	for _, capsuleItem := range capsule {
		capsuleRule = append(capsuleRule, capsuleItem)
	}

	logs, sub, err := _BootstrapStorage.contract.WatchLogs(opts, "StakedWithCapsule", stakerRule, capsuleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapStorageStakedWithCapsule)
				if err := _BootstrapStorage.contract.UnpackLog(event, "StakedWithCapsule", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakedWithCapsule is a log parse operation binding the contract event 0x9baefaaa262c2c2f0c3acf8ee9293d5043ef654d71c2aba83024f1322c72e5ee.
//
// Solidity: event StakedWithCapsule(address indexed staker, address indexed capsule)
func (_BootstrapStorage *BootstrapStorageFilterer) ParseStakedWithCapsule(log types.Log) (*BootstrapStorageStakedWithCapsule, error) {
	event := new(BootstrapStorageStakedWithCapsule)
	if err := _BootstrapStorage.contract.UnpackLog(event, "StakedWithCapsule", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapStorageUndelegateResultIterator is returned from FilterUndelegateResult and is used to iterate over the raw logs and unpacked data for UndelegateResult events raised by the BootstrapStorage contract.
type BootstrapStorageUndelegateResultIterator struct {
	Event *BootstrapStorageUndelegateResult // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapStorageUndelegateResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapStorageUndelegateResult)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapStorageUndelegateResult)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapStorageUndelegateResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapStorageUndelegateResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapStorageUndelegateResult represents a UndelegateResult event raised by the BootstrapStorage contract.
type BootstrapStorageUndelegateResult struct {
	Success     bool
	Undelegator common.Address
	Undelegatee common.Hash
	Token       common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUndelegateResult is a free log retrieval operation binding the contract event 0x8bc7dddfddc42a489a24fbcbb5e45de61c15f3228d343959e1b9c176f835531a.
//
// Solidity: event UndelegateResult(bool indexed success, address indexed undelegator, string indexed undelegatee, address token, uint256 amount)
func (_BootstrapStorage *BootstrapStorageFilterer) FilterUndelegateResult(opts *bind.FilterOpts, success []bool, undelegator []common.Address, undelegatee []string) (*BootstrapStorageUndelegateResultIterator, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var undelegatorRule []interface{}
	for _, undelegatorItem := range undelegator {
		undelegatorRule = append(undelegatorRule, undelegatorItem)
	}
	var undelegateeRule []interface{}
	for _, undelegateeItem := range undelegatee {
		undelegateeRule = append(undelegateeRule, undelegateeItem)
	}

	logs, sub, err := _BootstrapStorage.contract.FilterLogs(opts, "UndelegateResult", successRule, undelegatorRule, undelegateeRule)
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageUndelegateResultIterator{contract: _BootstrapStorage.contract, event: "UndelegateResult", logs: logs, sub: sub}, nil
}

// WatchUndelegateResult is a free log subscription operation binding the contract event 0x8bc7dddfddc42a489a24fbcbb5e45de61c15f3228d343959e1b9c176f835531a.
//
// Solidity: event UndelegateResult(bool indexed success, address indexed undelegator, string indexed undelegatee, address token, uint256 amount)
func (_BootstrapStorage *BootstrapStorageFilterer) WatchUndelegateResult(opts *bind.WatchOpts, sink chan<- *BootstrapStorageUndelegateResult, success []bool, undelegator []common.Address, undelegatee []string) (event.Subscription, error) {

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}
	var undelegatorRule []interface{}
	for _, undelegatorItem := range undelegator {
		undelegatorRule = append(undelegatorRule, undelegatorItem)
	}
	var undelegateeRule []interface{}
	for _, undelegateeItem := range undelegatee {
		undelegateeRule = append(undelegateeRule, undelegateeItem)
	}

	logs, sub, err := _BootstrapStorage.contract.WatchLogs(opts, "UndelegateResult", successRule, undelegatorRule, undelegateeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapStorageUndelegateResult)
				if err := _BootstrapStorage.contract.UnpackLog(event, "UndelegateResult", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUndelegateResult is a log parse operation binding the contract event 0x8bc7dddfddc42a489a24fbcbb5e45de61c15f3228d343959e1b9c176f835531a.
//
// Solidity: event UndelegateResult(bool indexed success, address indexed undelegator, string indexed undelegatee, address token, uint256 amount)
func (_BootstrapStorage *BootstrapStorageFilterer) ParseUndelegateResult(log types.Log) (*BootstrapStorageUndelegateResult, error) {
	event := new(BootstrapStorageUndelegateResult)
	if err := _BootstrapStorage.contract.UnpackLog(event, "UndelegateResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapStorageVaultCreatedIterator is returned from FilterVaultCreated and is used to iterate over the raw logs and unpacked data for VaultCreated events raised by the BootstrapStorage contract.
type BootstrapStorageVaultCreatedIterator struct {
	Event *BootstrapStorageVaultCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapStorageVaultCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapStorageVaultCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapStorageVaultCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapStorageVaultCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapStorageVaultCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapStorageVaultCreated represents a VaultCreated event raised by the BootstrapStorage contract.
type BootstrapStorageVaultCreated struct {
	UnderlyingToken common.Address
	Vault           common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVaultCreated is a free log retrieval operation binding the contract event 0x5d9c31ffa0fecffd7cf379989a3c7af252f0335e0d2a1320b55245912c781f53.
//
// Solidity: event VaultCreated(address underlyingToken, address vault)
func (_BootstrapStorage *BootstrapStorageFilterer) FilterVaultCreated(opts *bind.FilterOpts) (*BootstrapStorageVaultCreatedIterator, error) {

	logs, sub, err := _BootstrapStorage.contract.FilterLogs(opts, "VaultCreated")
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageVaultCreatedIterator{contract: _BootstrapStorage.contract, event: "VaultCreated", logs: logs, sub: sub}, nil
}

// WatchVaultCreated is a free log subscription operation binding the contract event 0x5d9c31ffa0fecffd7cf379989a3c7af252f0335e0d2a1320b55245912c781f53.
//
// Solidity: event VaultCreated(address underlyingToken, address vault)
func (_BootstrapStorage *BootstrapStorageFilterer) WatchVaultCreated(opts *bind.WatchOpts, sink chan<- *BootstrapStorageVaultCreated) (event.Subscription, error) {

	logs, sub, err := _BootstrapStorage.contract.WatchLogs(opts, "VaultCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapStorageVaultCreated)
				if err := _BootstrapStorage.contract.UnpackLog(event, "VaultCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVaultCreated is a log parse operation binding the contract event 0x5d9c31ffa0fecffd7cf379989a3c7af252f0335e0d2a1320b55245912c781f53.
//
// Solidity: event VaultCreated(address underlyingToken, address vault)
func (_BootstrapStorage *BootstrapStorageFilterer) ParseVaultCreated(log types.Log) (*BootstrapStorageVaultCreated, error) {
	event := new(BootstrapStorageVaultCreated)
	if err := _BootstrapStorage.contract.UnpackLog(event, "VaultCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BootstrapStorageWhitelistTokenAddedIterator is returned from FilterWhitelistTokenAdded and is used to iterate over the raw logs and unpacked data for WhitelistTokenAdded events raised by the BootstrapStorage contract.
type BootstrapStorageWhitelistTokenAddedIterator struct {
	Event *BootstrapStorageWhitelistTokenAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BootstrapStorageWhitelistTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BootstrapStorageWhitelistTokenAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BootstrapStorageWhitelistTokenAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BootstrapStorageWhitelistTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BootstrapStorageWhitelistTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BootstrapStorageWhitelistTokenAdded represents a WhitelistTokenAdded event raised by the BootstrapStorage contract.
type BootstrapStorageWhitelistTokenAdded struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWhitelistTokenAdded is a free log retrieval operation binding the contract event 0x3c96ec6d36b4038b363b8d7e8ad3b2fc7204f60317b8d0a4082094b6295fff2c.
//
// Solidity: event WhitelistTokenAdded(address _token)
func (_BootstrapStorage *BootstrapStorageFilterer) FilterWhitelistTokenAdded(opts *bind.FilterOpts) (*BootstrapStorageWhitelistTokenAddedIterator, error) {

	logs, sub, err := _BootstrapStorage.contract.FilterLogs(opts, "WhitelistTokenAdded")
	if err != nil {
		return nil, err
	}
	return &BootstrapStorageWhitelistTokenAddedIterator{contract: _BootstrapStorage.contract, event: "WhitelistTokenAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistTokenAdded is a free log subscription operation binding the contract event 0x3c96ec6d36b4038b363b8d7e8ad3b2fc7204f60317b8d0a4082094b6295fff2c.
//
// Solidity: event WhitelistTokenAdded(address _token)
func (_BootstrapStorage *BootstrapStorageFilterer) WatchWhitelistTokenAdded(opts *bind.WatchOpts, sink chan<- *BootstrapStorageWhitelistTokenAdded) (event.Subscription, error) {

	logs, sub, err := _BootstrapStorage.contract.WatchLogs(opts, "WhitelistTokenAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BootstrapStorageWhitelistTokenAdded)
				if err := _BootstrapStorage.contract.UnpackLog(event, "WhitelistTokenAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWhitelistTokenAdded is a log parse operation binding the contract event 0x3c96ec6d36b4038b363b8d7e8ad3b2fc7204f60317b8d0a4082094b6295fff2c.
//
// Solidity: event WhitelistTokenAdded(address _token)
func (_BootstrapStorage *BootstrapStorageFilterer) ParseWhitelistTokenAdded(log types.Log) (*BootstrapStorageWhitelistTokenAdded, error) {
	event := new(BootstrapStorageWhitelistTokenAdded)
	if err := _BootstrapStorage.contract.UnpackLog(event, "WhitelistTokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

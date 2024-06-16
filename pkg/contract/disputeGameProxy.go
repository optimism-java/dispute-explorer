// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IDisputeGameFactoryGameSearchResult is an auto generated low-level Go binding around an user-defined struct.
type IDisputeGameFactoryGameSearchResult struct {
	Index     *big.Int
	Metadata  [32]byte
	Timestamp uint64
	RootClaim [32]byte
	ExtraData []byte
}

// DisputeGameProxyMetaData contains all meta data concerning the DisputeGameProxy contract.
var DisputeGameProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"Hash\",\"name\":\"uuid\",\"type\":\"bytes32\"}],\"name\":\"GameAlreadyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectBondAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"GameType\",\"name\":\"gameType\",\"type\":\"uint32\"}],\"name\":\"NoImplementation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"disputeProxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"GameType\",\"name\":\"gameType\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"Claim\",\"name\":\"rootClaim\",\"type\":\"bytes32\"}],\"name\":\"DisputeGameCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"GameType\",\"name\":\"gameType\",\"type\":\"uint32\"}],\"name\":\"ImplementationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"GameType\",\"name\":\"gameType\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newBond\",\"type\":\"uint256\"}],\"name\":\"InitBondUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"GameType\",\"name\":\"_gameType\",\"type\":\"uint32\"},{\"internalType\":\"Claim\",\"name\":\"_rootClaim\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"contractIDisputeGame\",\"name\":\"proxy_\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"GameType\",\"name\":\"_gameType\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_n\",\"type\":\"uint256\"}],\"name\":\"findLatestGames\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"GameId\",\"name\":\"metadata\",\"type\":\"bytes32\"},{\"internalType\":\"Timestamp\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"Claim\",\"name\":\"rootClaim\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraData\",\"type\":\"bytes\"}],\"internalType\":\"structIDisputeGameFactory.GameSearchResult[]\",\"name\":\"games_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"gameAtIndex\",\"outputs\":[{\"internalType\":\"GameType\",\"name\":\"gameType_\",\"type\":\"uint32\"},{\"internalType\":\"Timestamp\",\"name\":\"timestamp_\",\"type\":\"uint64\"},{\"internalType\":\"contractIDisputeGame\",\"name\":\"proxy_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gameCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gameCount_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"GameType\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"gameImpls\",\"outputs\":[{\"internalType\":\"contractIDisputeGame\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"GameType\",\"name\":\"_gameType\",\"type\":\"uint32\"},{\"internalType\":\"Claim\",\"name\":\"_rootClaim\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"games\",\"outputs\":[{\"internalType\":\"contractIDisputeGame\",\"name\":\"proxy_\",\"type\":\"address\"},{\"internalType\":\"Timestamp\",\"name\":\"timestamp_\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"GameType\",\"name\":\"_gameType\",\"type\":\"uint32\"},{\"internalType\":\"Claim\",\"name\":\"_rootClaim\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"getGameUUID\",\"outputs\":[{\"internalType\":\"Hash\",\"name\":\"uuid_\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"GameType\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"initBonds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"GameType\",\"name\":\"_gameType\",\"type\":\"uint32\"},{\"internalType\":\"contractIDisputeGame\",\"name\":\"_impl\",\"type\":\"address\"}],\"name\":\"setImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"GameType\",\"name\":\"_gameType\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_initBond\",\"type\":\"uint256\"}],\"name\":\"setInitBond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DisputeGameProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use DisputeGameProxyMetaData.ABI instead.
var DisputeGameProxyABI = DisputeGameProxyMetaData.ABI

// DisputeGameProxy is an auto generated Go binding around an Ethereum contract.
type DisputeGameProxy struct {
	DisputeGameProxyCaller     // Read-only binding to the contract
	DisputeGameProxyTransactor // Write-only binding to the contract
	DisputeGameProxyFilterer   // Log filterer for contract events
}

// DisputeGameProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type DisputeGameProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisputeGameProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DisputeGameProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisputeGameProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DisputeGameProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DisputeGameProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DisputeGameProxySession struct {
	Contract     *DisputeGameProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DisputeGameProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DisputeGameProxyCallerSession struct {
	Contract *DisputeGameProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DisputeGameProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DisputeGameProxyTransactorSession struct {
	Contract     *DisputeGameProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DisputeGameProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type DisputeGameProxyRaw struct {
	Contract *DisputeGameProxy // Generic contract binding to access the raw methods on
}

// DisputeGameProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DisputeGameProxyCallerRaw struct {
	Contract *DisputeGameProxyCaller // Generic read-only contract binding to access the raw methods on
}

// DisputeGameProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DisputeGameProxyTransactorRaw struct {
	Contract *DisputeGameProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDisputeGameProxy creates a new instance of DisputeGameProxy, bound to a specific deployed contract.
func NewDisputeGameProxy(address common.Address, backend bind.ContractBackend) (*DisputeGameProxy, error) {
	contract, err := bindDisputeGameProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DisputeGameProxy{DisputeGameProxyCaller: DisputeGameProxyCaller{contract: contract}, DisputeGameProxyTransactor: DisputeGameProxyTransactor{contract: contract}, DisputeGameProxyFilterer: DisputeGameProxyFilterer{contract: contract}}, nil
}

// NewDisputeGameProxyCaller creates a new read-only instance of DisputeGameProxy, bound to a specific deployed contract.
func NewDisputeGameProxyCaller(address common.Address, caller bind.ContractCaller) (*DisputeGameProxyCaller, error) {
	contract, err := bindDisputeGameProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DisputeGameProxyCaller{contract: contract}, nil
}

// NewDisputeGameProxyTransactor creates a new write-only instance of DisputeGameProxy, bound to a specific deployed contract.
func NewDisputeGameProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*DisputeGameProxyTransactor, error) {
	contract, err := bindDisputeGameProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DisputeGameProxyTransactor{contract: contract}, nil
}

// NewDisputeGameProxyFilterer creates a new log filterer instance of DisputeGameProxy, bound to a specific deployed contract.
func NewDisputeGameProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*DisputeGameProxyFilterer, error) {
	contract, err := bindDisputeGameProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DisputeGameProxyFilterer{contract: contract}, nil
}

// bindDisputeGameProxy binds a generic wrapper to an already deployed contract.
func bindDisputeGameProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DisputeGameProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DisputeGameProxy *DisputeGameProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DisputeGameProxy.Contract.DisputeGameProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DisputeGameProxy *DisputeGameProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DisputeGameProxy.Contract.DisputeGameProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DisputeGameProxy *DisputeGameProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DisputeGameProxy.Contract.DisputeGameProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DisputeGameProxy *DisputeGameProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DisputeGameProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DisputeGameProxy *DisputeGameProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DisputeGameProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DisputeGameProxy *DisputeGameProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DisputeGameProxy.Contract.contract.Transact(opts, method, params...)
}

// FindLatestGames is a free data retrieval call binding the contract method 0x254bd683.
//
// Solidity: function findLatestGames(uint32 _gameType, uint256 _start, uint256 _n) view returns((uint256,bytes32,uint64,bytes32,bytes)[] games_)
func (_DisputeGameProxy *DisputeGameProxyCaller) FindLatestGames(opts *bind.CallOpts, _gameType uint32, _start *big.Int, _n *big.Int) ([]IDisputeGameFactoryGameSearchResult, error) {
	var out []interface{}
	err := _DisputeGameProxy.contract.Call(opts, &out, "findLatestGames", _gameType, _start, _n)

	if err != nil {
		return *new([]IDisputeGameFactoryGameSearchResult), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDisputeGameFactoryGameSearchResult)).(*[]IDisputeGameFactoryGameSearchResult)

	return out0, err

}

// FindLatestGames is a free data retrieval call binding the contract method 0x254bd683.
//
// Solidity: function findLatestGames(uint32 _gameType, uint256 _start, uint256 _n) view returns((uint256,bytes32,uint64,bytes32,bytes)[] games_)
func (_DisputeGameProxy *DisputeGameProxySession) FindLatestGames(_gameType uint32, _start *big.Int, _n *big.Int) ([]IDisputeGameFactoryGameSearchResult, error) {
	return _DisputeGameProxy.Contract.FindLatestGames(&_DisputeGameProxy.CallOpts, _gameType, _start, _n)
}

// FindLatestGames is a free data retrieval call binding the contract method 0x254bd683.
//
// Solidity: function findLatestGames(uint32 _gameType, uint256 _start, uint256 _n) view returns((uint256,bytes32,uint64,bytes32,bytes)[] games_)
func (_DisputeGameProxy *DisputeGameProxyCallerSession) FindLatestGames(_gameType uint32, _start *big.Int, _n *big.Int) ([]IDisputeGameFactoryGameSearchResult, error) {
	return _DisputeGameProxy.Contract.FindLatestGames(&_DisputeGameProxy.CallOpts, _gameType, _start, _n)
}

// GameAtIndex is a free data retrieval call binding the contract method 0xbb8aa1fc.
//
// Solidity: function gameAtIndex(uint256 _index) view returns(uint32 gameType_, uint64 timestamp_, address proxy_)
func (_DisputeGameProxy *DisputeGameProxyCaller) GameAtIndex(opts *bind.CallOpts, _index *big.Int) (struct {
	GameType  uint32
	Timestamp uint64
	Proxy     common.Address
}, error) {
	var out []interface{}
	err := _DisputeGameProxy.contract.Call(opts, &out, "gameAtIndex", _index)

	outstruct := new(struct {
		GameType  uint32
		Timestamp uint64
		Proxy     common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GameType = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Proxy = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GameAtIndex is a free data retrieval call binding the contract method 0xbb8aa1fc.
//
// Solidity: function gameAtIndex(uint256 _index) view returns(uint32 gameType_, uint64 timestamp_, address proxy_)
func (_DisputeGameProxy *DisputeGameProxySession) GameAtIndex(_index *big.Int) (struct {
	GameType  uint32
	Timestamp uint64
	Proxy     common.Address
}, error) {
	return _DisputeGameProxy.Contract.GameAtIndex(&_DisputeGameProxy.CallOpts, _index)
}

// GameAtIndex is a free data retrieval call binding the contract method 0xbb8aa1fc.
//
// Solidity: function gameAtIndex(uint256 _index) view returns(uint32 gameType_, uint64 timestamp_, address proxy_)
func (_DisputeGameProxy *DisputeGameProxyCallerSession) GameAtIndex(_index *big.Int) (struct {
	GameType  uint32
	Timestamp uint64
	Proxy     common.Address
}, error) {
	return _DisputeGameProxy.Contract.GameAtIndex(&_DisputeGameProxy.CallOpts, _index)
}

// GameCount is a free data retrieval call binding the contract method 0x4d1975b4.
//
// Solidity: function gameCount() view returns(uint256 gameCount_)
func (_DisputeGameProxy *DisputeGameProxyCaller) GameCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DisputeGameProxy.contract.Call(opts, &out, "gameCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GameCount is a free data retrieval call binding the contract method 0x4d1975b4.
//
// Solidity: function gameCount() view returns(uint256 gameCount_)
func (_DisputeGameProxy *DisputeGameProxySession) GameCount() (*big.Int, error) {
	return _DisputeGameProxy.Contract.GameCount(&_DisputeGameProxy.CallOpts)
}

// GameCount is a free data retrieval call binding the contract method 0x4d1975b4.
//
// Solidity: function gameCount() view returns(uint256 gameCount_)
func (_DisputeGameProxy *DisputeGameProxyCallerSession) GameCount() (*big.Int, error) {
	return _DisputeGameProxy.Contract.GameCount(&_DisputeGameProxy.CallOpts)
}

// GameImpls is a free data retrieval call binding the contract method 0x1b685b9e.
//
// Solidity: function gameImpls(uint32 ) view returns(address)
func (_DisputeGameProxy *DisputeGameProxyCaller) GameImpls(opts *bind.CallOpts, arg0 uint32) (common.Address, error) {
	var out []interface{}
	err := _DisputeGameProxy.contract.Call(opts, &out, "gameImpls", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GameImpls is a free data retrieval call binding the contract method 0x1b685b9e.
//
// Solidity: function gameImpls(uint32 ) view returns(address)
func (_DisputeGameProxy *DisputeGameProxySession) GameImpls(arg0 uint32) (common.Address, error) {
	return _DisputeGameProxy.Contract.GameImpls(&_DisputeGameProxy.CallOpts, arg0)
}

// GameImpls is a free data retrieval call binding the contract method 0x1b685b9e.
//
// Solidity: function gameImpls(uint32 ) view returns(address)
func (_DisputeGameProxy *DisputeGameProxyCallerSession) GameImpls(arg0 uint32) (common.Address, error) {
	return _DisputeGameProxy.Contract.GameImpls(&_DisputeGameProxy.CallOpts, arg0)
}

// Games is a free data retrieval call binding the contract method 0x5f0150cb.
//
// Solidity: function games(uint32 _gameType, bytes32 _rootClaim, bytes _extraData) view returns(address proxy_, uint64 timestamp_)
func (_DisputeGameProxy *DisputeGameProxyCaller) Games(opts *bind.CallOpts, _gameType uint32, _rootClaim [32]byte, _extraData []byte) (struct {
	Proxy     common.Address
	Timestamp uint64
}, error) {
	var out []interface{}
	err := _DisputeGameProxy.contract.Call(opts, &out, "games", _gameType, _rootClaim, _extraData)

	outstruct := new(struct {
		Proxy     common.Address
		Timestamp uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Proxy = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// Games is a free data retrieval call binding the contract method 0x5f0150cb.
//
// Solidity: function games(uint32 _gameType, bytes32 _rootClaim, bytes _extraData) view returns(address proxy_, uint64 timestamp_)
func (_DisputeGameProxy *DisputeGameProxySession) Games(_gameType uint32, _rootClaim [32]byte, _extraData []byte) (struct {
	Proxy     common.Address
	Timestamp uint64
}, error) {
	return _DisputeGameProxy.Contract.Games(&_DisputeGameProxy.CallOpts, _gameType, _rootClaim, _extraData)
}

// Games is a free data retrieval call binding the contract method 0x5f0150cb.
//
// Solidity: function games(uint32 _gameType, bytes32 _rootClaim, bytes _extraData) view returns(address proxy_, uint64 timestamp_)
func (_DisputeGameProxy *DisputeGameProxyCallerSession) Games(_gameType uint32, _rootClaim [32]byte, _extraData []byte) (struct {
	Proxy     common.Address
	Timestamp uint64
}, error) {
	return _DisputeGameProxy.Contract.Games(&_DisputeGameProxy.CallOpts, _gameType, _rootClaim, _extraData)
}

// GetGameUUID is a free data retrieval call binding the contract method 0x96cd9720.
//
// Solidity: function getGameUUID(uint32 _gameType, bytes32 _rootClaim, bytes _extraData) pure returns(bytes32 uuid_)
func (_DisputeGameProxy *DisputeGameProxyCaller) GetGameUUID(opts *bind.CallOpts, _gameType uint32, _rootClaim [32]byte, _extraData []byte) ([32]byte, error) {
	var out []interface{}
	err := _DisputeGameProxy.contract.Call(opts, &out, "getGameUUID", _gameType, _rootClaim, _extraData)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetGameUUID is a free data retrieval call binding the contract method 0x96cd9720.
//
// Solidity: function getGameUUID(uint32 _gameType, bytes32 _rootClaim, bytes _extraData) pure returns(bytes32 uuid_)
func (_DisputeGameProxy *DisputeGameProxySession) GetGameUUID(_gameType uint32, _rootClaim [32]byte, _extraData []byte) ([32]byte, error) {
	return _DisputeGameProxy.Contract.GetGameUUID(&_DisputeGameProxy.CallOpts, _gameType, _rootClaim, _extraData)
}

// GetGameUUID is a free data retrieval call binding the contract method 0x96cd9720.
//
// Solidity: function getGameUUID(uint32 _gameType, bytes32 _rootClaim, bytes _extraData) pure returns(bytes32 uuid_)
func (_DisputeGameProxy *DisputeGameProxyCallerSession) GetGameUUID(_gameType uint32, _rootClaim [32]byte, _extraData []byte) ([32]byte, error) {
	return _DisputeGameProxy.Contract.GetGameUUID(&_DisputeGameProxy.CallOpts, _gameType, _rootClaim, _extraData)
}

// InitBonds is a free data retrieval call binding the contract method 0x6593dc6e.
//
// Solidity: function initBonds(uint32 ) view returns(uint256)
func (_DisputeGameProxy *DisputeGameProxyCaller) InitBonds(opts *bind.CallOpts, arg0 uint32) (*big.Int, error) {
	var out []interface{}
	err := _DisputeGameProxy.contract.Call(opts, &out, "initBonds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitBonds is a free data retrieval call binding the contract method 0x6593dc6e.
//
// Solidity: function initBonds(uint32 ) view returns(uint256)
func (_DisputeGameProxy *DisputeGameProxySession) InitBonds(arg0 uint32) (*big.Int, error) {
	return _DisputeGameProxy.Contract.InitBonds(&_DisputeGameProxy.CallOpts, arg0)
}

// InitBonds is a free data retrieval call binding the contract method 0x6593dc6e.
//
// Solidity: function initBonds(uint32 ) view returns(uint256)
func (_DisputeGameProxy *DisputeGameProxyCallerSession) InitBonds(arg0 uint32) (*big.Int, error) {
	return _DisputeGameProxy.Contract.InitBonds(&_DisputeGameProxy.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DisputeGameProxy *DisputeGameProxyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DisputeGameProxy.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DisputeGameProxy *DisputeGameProxySession) Owner() (common.Address, error) {
	return _DisputeGameProxy.Contract.Owner(&_DisputeGameProxy.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DisputeGameProxy *DisputeGameProxyCallerSession) Owner() (common.Address, error) {
	return _DisputeGameProxy.Contract.Owner(&_DisputeGameProxy.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DisputeGameProxy *DisputeGameProxyCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DisputeGameProxy.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DisputeGameProxy *DisputeGameProxySession) Version() (string, error) {
	return _DisputeGameProxy.Contract.Version(&_DisputeGameProxy.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_DisputeGameProxy *DisputeGameProxyCallerSession) Version() (string, error) {
	return _DisputeGameProxy.Contract.Version(&_DisputeGameProxy.CallOpts)
}

// Create is a paid mutator transaction binding the contract method 0x82ecf2f6.
//
// Solidity: function create(uint32 _gameType, bytes32 _rootClaim, bytes _extraData) payable returns(address proxy_)
func (_DisputeGameProxy *DisputeGameProxyTransactor) Create(opts *bind.TransactOpts, _gameType uint32, _rootClaim [32]byte, _extraData []byte) (*types.Transaction, error) {
	return _DisputeGameProxy.contract.Transact(opts, "create", _gameType, _rootClaim, _extraData)
}

// Create is a paid mutator transaction binding the contract method 0x82ecf2f6.
//
// Solidity: function create(uint32 _gameType, bytes32 _rootClaim, bytes _extraData) payable returns(address proxy_)
func (_DisputeGameProxy *DisputeGameProxySession) Create(_gameType uint32, _rootClaim [32]byte, _extraData []byte) (*types.Transaction, error) {
	return _DisputeGameProxy.Contract.Create(&_DisputeGameProxy.TransactOpts, _gameType, _rootClaim, _extraData)
}

// Create is a paid mutator transaction binding the contract method 0x82ecf2f6.
//
// Solidity: function create(uint32 _gameType, bytes32 _rootClaim, bytes _extraData) payable returns(address proxy_)
func (_DisputeGameProxy *DisputeGameProxyTransactorSession) Create(_gameType uint32, _rootClaim [32]byte, _extraData []byte) (*types.Transaction, error) {
	return _DisputeGameProxy.Contract.Create(&_DisputeGameProxy.TransactOpts, _gameType, _rootClaim, _extraData)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_DisputeGameProxy *DisputeGameProxyTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _DisputeGameProxy.contract.Transact(opts, "initialize", _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_DisputeGameProxy *DisputeGameProxySession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _DisputeGameProxy.Contract.Initialize(&_DisputeGameProxy.TransactOpts, _owner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _owner) returns()
func (_DisputeGameProxy *DisputeGameProxyTransactorSession) Initialize(_owner common.Address) (*types.Transaction, error) {
	return _DisputeGameProxy.Contract.Initialize(&_DisputeGameProxy.TransactOpts, _owner)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DisputeGameProxy *DisputeGameProxyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DisputeGameProxy.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DisputeGameProxy *DisputeGameProxySession) RenounceOwnership() (*types.Transaction, error) {
	return _DisputeGameProxy.Contract.RenounceOwnership(&_DisputeGameProxy.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DisputeGameProxy *DisputeGameProxyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DisputeGameProxy.Contract.RenounceOwnership(&_DisputeGameProxy.TransactOpts)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x14f6b1a3.
//
// Solidity: function setImplementation(uint32 _gameType, address _impl) returns()
func (_DisputeGameProxy *DisputeGameProxyTransactor) SetImplementation(opts *bind.TransactOpts, _gameType uint32, _impl common.Address) (*types.Transaction, error) {
	return _DisputeGameProxy.contract.Transact(opts, "setImplementation", _gameType, _impl)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x14f6b1a3.
//
// Solidity: function setImplementation(uint32 _gameType, address _impl) returns()
func (_DisputeGameProxy *DisputeGameProxySession) SetImplementation(_gameType uint32, _impl common.Address) (*types.Transaction, error) {
	return _DisputeGameProxy.Contract.SetImplementation(&_DisputeGameProxy.TransactOpts, _gameType, _impl)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x14f6b1a3.
//
// Solidity: function setImplementation(uint32 _gameType, address _impl) returns()
func (_DisputeGameProxy *DisputeGameProxyTransactorSession) SetImplementation(_gameType uint32, _impl common.Address) (*types.Transaction, error) {
	return _DisputeGameProxy.Contract.SetImplementation(&_DisputeGameProxy.TransactOpts, _gameType, _impl)
}

// SetInitBond is a paid mutator transaction binding the contract method 0x1e334240.
//
// Solidity: function setInitBond(uint32 _gameType, uint256 _initBond) returns()
func (_DisputeGameProxy *DisputeGameProxyTransactor) SetInitBond(opts *bind.TransactOpts, _gameType uint32, _initBond *big.Int) (*types.Transaction, error) {
	return _DisputeGameProxy.contract.Transact(opts, "setInitBond", _gameType, _initBond)
}

// SetInitBond is a paid mutator transaction binding the contract method 0x1e334240.
//
// Solidity: function setInitBond(uint32 _gameType, uint256 _initBond) returns()
func (_DisputeGameProxy *DisputeGameProxySession) SetInitBond(_gameType uint32, _initBond *big.Int) (*types.Transaction, error) {
	return _DisputeGameProxy.Contract.SetInitBond(&_DisputeGameProxy.TransactOpts, _gameType, _initBond)
}

// SetInitBond is a paid mutator transaction binding the contract method 0x1e334240.
//
// Solidity: function setInitBond(uint32 _gameType, uint256 _initBond) returns()
func (_DisputeGameProxy *DisputeGameProxyTransactorSession) SetInitBond(_gameType uint32, _initBond *big.Int) (*types.Transaction, error) {
	return _DisputeGameProxy.Contract.SetInitBond(&_DisputeGameProxy.TransactOpts, _gameType, _initBond)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DisputeGameProxy *DisputeGameProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DisputeGameProxy.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DisputeGameProxy *DisputeGameProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DisputeGameProxy.Contract.TransferOwnership(&_DisputeGameProxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DisputeGameProxy *DisputeGameProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DisputeGameProxy.Contract.TransferOwnership(&_DisputeGameProxy.TransactOpts, newOwner)
}

// DisputeGameProxyDisputeGameCreatedIterator is returned from FilterDisputeGameCreated and is used to iterate over the raw logs and unpacked data for DisputeGameCreated events raised by the DisputeGameProxy contract.
type DisputeGameProxyDisputeGameCreatedIterator struct {
	Event *DisputeGameProxyDisputeGameCreated // Event containing the contract specifics and raw log

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
func (it *DisputeGameProxyDisputeGameCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputeGameProxyDisputeGameCreated)
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
		it.Event = new(DisputeGameProxyDisputeGameCreated)
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
func (it *DisputeGameProxyDisputeGameCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputeGameProxyDisputeGameCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputeGameProxyDisputeGameCreated represents a DisputeGameCreated event raised by the DisputeGameProxy contract.
type DisputeGameProxyDisputeGameCreated struct {
	DisputeProxy common.Address
	GameType     uint32
	RootClaim    [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDisputeGameCreated is a free log retrieval operation binding the contract event 0x5b565efe82411da98814f356d0e7bcb8f0219b8d970307c5afb4a6903a8b2e35.
//
// Solidity: event DisputeGameCreated(address indexed disputeProxy, uint32 indexed gameType, bytes32 indexed rootClaim)
func (_DisputeGameProxy *DisputeGameProxyFilterer) FilterDisputeGameCreated(opts *bind.FilterOpts, disputeProxy []common.Address, gameType []uint32, rootClaim [][32]byte) (*DisputeGameProxyDisputeGameCreatedIterator, error) {

	var disputeProxyRule []interface{}
	for _, disputeProxyItem := range disputeProxy {
		disputeProxyRule = append(disputeProxyRule, disputeProxyItem)
	}
	var gameTypeRule []interface{}
	for _, gameTypeItem := range gameType {
		gameTypeRule = append(gameTypeRule, gameTypeItem)
	}
	var rootClaimRule []interface{}
	for _, rootClaimItem := range rootClaim {
		rootClaimRule = append(rootClaimRule, rootClaimItem)
	}

	logs, sub, err := _DisputeGameProxy.contract.FilterLogs(opts, "DisputeGameCreated", disputeProxyRule, gameTypeRule, rootClaimRule)
	if err != nil {
		return nil, err
	}
	return &DisputeGameProxyDisputeGameCreatedIterator{contract: _DisputeGameProxy.contract, event: "DisputeGameCreated", logs: logs, sub: sub}, nil
}

// WatchDisputeGameCreated is a free log subscription operation binding the contract event 0x5b565efe82411da98814f356d0e7bcb8f0219b8d970307c5afb4a6903a8b2e35.
//
// Solidity: event DisputeGameCreated(address indexed disputeProxy, uint32 indexed gameType, bytes32 indexed rootClaim)
func (_DisputeGameProxy *DisputeGameProxyFilterer) WatchDisputeGameCreated(opts *bind.WatchOpts, sink chan<- *DisputeGameProxyDisputeGameCreated, disputeProxy []common.Address, gameType []uint32, rootClaim [][32]byte) (event.Subscription, error) {

	var disputeProxyRule []interface{}
	for _, disputeProxyItem := range disputeProxy {
		disputeProxyRule = append(disputeProxyRule, disputeProxyItem)
	}
	var gameTypeRule []interface{}
	for _, gameTypeItem := range gameType {
		gameTypeRule = append(gameTypeRule, gameTypeItem)
	}
	var rootClaimRule []interface{}
	for _, rootClaimItem := range rootClaim {
		rootClaimRule = append(rootClaimRule, rootClaimItem)
	}

	logs, sub, err := _DisputeGameProxy.contract.WatchLogs(opts, "DisputeGameCreated", disputeProxyRule, gameTypeRule, rootClaimRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputeGameProxyDisputeGameCreated)
				if err := _DisputeGameProxy.contract.UnpackLog(event, "DisputeGameCreated", log); err != nil {
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

// ParseDisputeGameCreated is a log parse operation binding the contract event 0x5b565efe82411da98814f356d0e7bcb8f0219b8d970307c5afb4a6903a8b2e35.
//
// Solidity: event DisputeGameCreated(address indexed disputeProxy, uint32 indexed gameType, bytes32 indexed rootClaim)
func (_DisputeGameProxy *DisputeGameProxyFilterer) ParseDisputeGameCreated(log types.Log) (*DisputeGameProxyDisputeGameCreated, error) {
	event := new(DisputeGameProxyDisputeGameCreated)
	if err := _DisputeGameProxy.contract.UnpackLog(event, "DisputeGameCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DisputeGameProxyImplementationSetIterator is returned from FilterImplementationSet and is used to iterate over the raw logs and unpacked data for ImplementationSet events raised by the DisputeGameProxy contract.
type DisputeGameProxyImplementationSetIterator struct {
	Event *DisputeGameProxyImplementationSet // Event containing the contract specifics and raw log

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
func (it *DisputeGameProxyImplementationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputeGameProxyImplementationSet)
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
		it.Event = new(DisputeGameProxyImplementationSet)
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
func (it *DisputeGameProxyImplementationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputeGameProxyImplementationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputeGameProxyImplementationSet represents a ImplementationSet event raised by the DisputeGameProxy contract.
type DisputeGameProxyImplementationSet struct {
	Impl     common.Address
	GameType uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterImplementationSet is a free log retrieval operation binding the contract event 0xff513d80e2c7fa487608f70a618dfbc0cf415699dc69588c747e8c71566c88de.
//
// Solidity: event ImplementationSet(address indexed impl, uint32 indexed gameType)
func (_DisputeGameProxy *DisputeGameProxyFilterer) FilterImplementationSet(opts *bind.FilterOpts, impl []common.Address, gameType []uint32) (*DisputeGameProxyImplementationSetIterator, error) {

	var implRule []interface{}
	for _, implItem := range impl {
		implRule = append(implRule, implItem)
	}
	var gameTypeRule []interface{}
	for _, gameTypeItem := range gameType {
		gameTypeRule = append(gameTypeRule, gameTypeItem)
	}

	logs, sub, err := _DisputeGameProxy.contract.FilterLogs(opts, "ImplementationSet", implRule, gameTypeRule)
	if err != nil {
		return nil, err
	}
	return &DisputeGameProxyImplementationSetIterator{contract: _DisputeGameProxy.contract, event: "ImplementationSet", logs: logs, sub: sub}, nil
}

// WatchImplementationSet is a free log subscription operation binding the contract event 0xff513d80e2c7fa487608f70a618dfbc0cf415699dc69588c747e8c71566c88de.
//
// Solidity: event ImplementationSet(address indexed impl, uint32 indexed gameType)
func (_DisputeGameProxy *DisputeGameProxyFilterer) WatchImplementationSet(opts *bind.WatchOpts, sink chan<- *DisputeGameProxyImplementationSet, impl []common.Address, gameType []uint32) (event.Subscription, error) {

	var implRule []interface{}
	for _, implItem := range impl {
		implRule = append(implRule, implItem)
	}
	var gameTypeRule []interface{}
	for _, gameTypeItem := range gameType {
		gameTypeRule = append(gameTypeRule, gameTypeItem)
	}

	logs, sub, err := _DisputeGameProxy.contract.WatchLogs(opts, "ImplementationSet", implRule, gameTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputeGameProxyImplementationSet)
				if err := _DisputeGameProxy.contract.UnpackLog(event, "ImplementationSet", log); err != nil {
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

// ParseImplementationSet is a log parse operation binding the contract event 0xff513d80e2c7fa487608f70a618dfbc0cf415699dc69588c747e8c71566c88de.
//
// Solidity: event ImplementationSet(address indexed impl, uint32 indexed gameType)
func (_DisputeGameProxy *DisputeGameProxyFilterer) ParseImplementationSet(log types.Log) (*DisputeGameProxyImplementationSet, error) {
	event := new(DisputeGameProxyImplementationSet)
	if err := _DisputeGameProxy.contract.UnpackLog(event, "ImplementationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DisputeGameProxyInitBondUpdatedIterator is returned from FilterInitBondUpdated and is used to iterate over the raw logs and unpacked data for InitBondUpdated events raised by the DisputeGameProxy contract.
type DisputeGameProxyInitBondUpdatedIterator struct {
	Event *DisputeGameProxyInitBondUpdated // Event containing the contract specifics and raw log

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
func (it *DisputeGameProxyInitBondUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputeGameProxyInitBondUpdated)
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
		it.Event = new(DisputeGameProxyInitBondUpdated)
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
func (it *DisputeGameProxyInitBondUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputeGameProxyInitBondUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputeGameProxyInitBondUpdated represents a InitBondUpdated event raised by the DisputeGameProxy contract.
type DisputeGameProxyInitBondUpdated struct {
	GameType uint32
	NewBond  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInitBondUpdated is a free log retrieval operation binding the contract event 0x74d6665c4b26d5596a5aa13d3014e0c06af4d322075a797f87b03cd4c5bc91ca.
//
// Solidity: event InitBondUpdated(uint32 indexed gameType, uint256 indexed newBond)
func (_DisputeGameProxy *DisputeGameProxyFilterer) FilterInitBondUpdated(opts *bind.FilterOpts, gameType []uint32, newBond []*big.Int) (*DisputeGameProxyInitBondUpdatedIterator, error) {

	var gameTypeRule []interface{}
	for _, gameTypeItem := range gameType {
		gameTypeRule = append(gameTypeRule, gameTypeItem)
	}
	var newBondRule []interface{}
	for _, newBondItem := range newBond {
		newBondRule = append(newBondRule, newBondItem)
	}

	logs, sub, err := _DisputeGameProxy.contract.FilterLogs(opts, "InitBondUpdated", gameTypeRule, newBondRule)
	if err != nil {
		return nil, err
	}
	return &DisputeGameProxyInitBondUpdatedIterator{contract: _DisputeGameProxy.contract, event: "InitBondUpdated", logs: logs, sub: sub}, nil
}

// WatchInitBondUpdated is a free log subscription operation binding the contract event 0x74d6665c4b26d5596a5aa13d3014e0c06af4d322075a797f87b03cd4c5bc91ca.
//
// Solidity: event InitBondUpdated(uint32 indexed gameType, uint256 indexed newBond)
func (_DisputeGameProxy *DisputeGameProxyFilterer) WatchInitBondUpdated(opts *bind.WatchOpts, sink chan<- *DisputeGameProxyInitBondUpdated, gameType []uint32, newBond []*big.Int) (event.Subscription, error) {

	var gameTypeRule []interface{}
	for _, gameTypeItem := range gameType {
		gameTypeRule = append(gameTypeRule, gameTypeItem)
	}
	var newBondRule []interface{}
	for _, newBondItem := range newBond {
		newBondRule = append(newBondRule, newBondItem)
	}

	logs, sub, err := _DisputeGameProxy.contract.WatchLogs(opts, "InitBondUpdated", gameTypeRule, newBondRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputeGameProxyInitBondUpdated)
				if err := _DisputeGameProxy.contract.UnpackLog(event, "InitBondUpdated", log); err != nil {
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

// ParseInitBondUpdated is a log parse operation binding the contract event 0x74d6665c4b26d5596a5aa13d3014e0c06af4d322075a797f87b03cd4c5bc91ca.
//
// Solidity: event InitBondUpdated(uint32 indexed gameType, uint256 indexed newBond)
func (_DisputeGameProxy *DisputeGameProxyFilterer) ParseInitBondUpdated(log types.Log) (*DisputeGameProxyInitBondUpdated, error) {
	event := new(DisputeGameProxyInitBondUpdated)
	if err := _DisputeGameProxy.contract.UnpackLog(event, "InitBondUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DisputeGameProxyInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DisputeGameProxy contract.
type DisputeGameProxyInitializedIterator struct {
	Event *DisputeGameProxyInitialized // Event containing the contract specifics and raw log

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
func (it *DisputeGameProxyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputeGameProxyInitialized)
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
		it.Event = new(DisputeGameProxyInitialized)
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
func (it *DisputeGameProxyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputeGameProxyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputeGameProxyInitialized represents a Initialized event raised by the DisputeGameProxy contract.
type DisputeGameProxyInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DisputeGameProxy *DisputeGameProxyFilterer) FilterInitialized(opts *bind.FilterOpts) (*DisputeGameProxyInitializedIterator, error) {

	logs, sub, err := _DisputeGameProxy.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DisputeGameProxyInitializedIterator{contract: _DisputeGameProxy.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DisputeGameProxy *DisputeGameProxyFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DisputeGameProxyInitialized) (event.Subscription, error) {

	logs, sub, err := _DisputeGameProxy.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputeGameProxyInitialized)
				if err := _DisputeGameProxy.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DisputeGameProxy *DisputeGameProxyFilterer) ParseInitialized(log types.Log) (*DisputeGameProxyInitialized, error) {
	event := new(DisputeGameProxyInitialized)
	if err := _DisputeGameProxy.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DisputeGameProxyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DisputeGameProxy contract.
type DisputeGameProxyOwnershipTransferredIterator struct {
	Event *DisputeGameProxyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DisputeGameProxyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DisputeGameProxyOwnershipTransferred)
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
		it.Event = new(DisputeGameProxyOwnershipTransferred)
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
func (it *DisputeGameProxyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DisputeGameProxyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DisputeGameProxyOwnershipTransferred represents a OwnershipTransferred event raised by the DisputeGameProxy contract.
type DisputeGameProxyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DisputeGameProxy *DisputeGameProxyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DisputeGameProxyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DisputeGameProxy.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DisputeGameProxyOwnershipTransferredIterator{contract: _DisputeGameProxy.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DisputeGameProxy *DisputeGameProxyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DisputeGameProxyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DisputeGameProxy.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DisputeGameProxyOwnershipTransferred)
				if err := _DisputeGameProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DisputeGameProxy *DisputeGameProxyFilterer) ParseOwnershipTransferred(log types.Log) (*DisputeGameProxyOwnershipTransferred, error) {
	event := new(DisputeGameProxyOwnershipTransferred)
	if err := _DisputeGameProxy.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

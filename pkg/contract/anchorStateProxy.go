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

// AnchorStateRegistryStartingAnchorRoot is an auto generated low-level Go binding around an user-defined struct.
type AnchorStateRegistryStartingAnchorRoot struct {
	GameType   uint32
	OutputRoot OutputRoot
}

// OutputRoot is an auto generated low-level Go binding around an user-defined struct.
type OutputRoot struct {
	Root          [32]byte
	L2BlockNumber *big.Int
}

// AnchorStateProxyMetaData contains all meta data concerning the AnchorStateProxy contract.
var AnchorStateProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIDisputeGameFactory\",\"name\":\"_disputeGameFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"GameType\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"anchors\",\"outputs\":[{\"internalType\":\"Hash\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"l2BlockNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disputeGameFactory\",\"outputs\":[{\"internalType\":\"contractIDisputeGameFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"GameType\",\"name\":\"gameType\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"Hash\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"l2BlockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structOutputRoot\",\"name\":\"outputRoot\",\"type\":\"tuple\"}],\"internalType\":\"structAnchorStateRegistry.StartingAnchorRoot[]\",\"name\":\"_startingAnchorRoots\",\"type\":\"tuple[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tryUpdateAnchorState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AnchorStateProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use AnchorStateProxyMetaData.ABI instead.
var AnchorStateProxyABI = AnchorStateProxyMetaData.ABI

// AnchorStateProxy is an auto generated Go binding around an Ethereum contract.
type AnchorStateProxy struct {
	AnchorStateProxyCaller     // Read-only binding to the contract
	AnchorStateProxyTransactor // Write-only binding to the contract
	AnchorStateProxyFilterer   // Log filterer for contract events
}

// AnchorStateProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type AnchorStateProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnchorStateProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AnchorStateProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnchorStateProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AnchorStateProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AnchorStateProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AnchorStateProxySession struct {
	Contract     *AnchorStateProxy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AnchorStateProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AnchorStateProxyCallerSession struct {
	Contract *AnchorStateProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// AnchorStateProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AnchorStateProxyTransactorSession struct {
	Contract     *AnchorStateProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AnchorStateProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type AnchorStateProxyRaw struct {
	Contract *AnchorStateProxy // Generic contract binding to access the raw methods on
}

// AnchorStateProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AnchorStateProxyCallerRaw struct {
	Contract *AnchorStateProxyCaller // Generic read-only contract binding to access the raw methods on
}

// AnchorStateProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AnchorStateProxyTransactorRaw struct {
	Contract *AnchorStateProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAnchorStateProxy creates a new instance of AnchorStateProxy, bound to a specific deployed contract.
func NewAnchorStateProxy(address common.Address, backend bind.ContractBackend) (*AnchorStateProxy, error) {
	contract, err := bindAnchorStateProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AnchorStateProxy{AnchorStateProxyCaller: AnchorStateProxyCaller{contract: contract}, AnchorStateProxyTransactor: AnchorStateProxyTransactor{contract: contract}, AnchorStateProxyFilterer: AnchorStateProxyFilterer{contract: contract}}, nil
}

// NewAnchorStateProxyCaller creates a new read-only instance of AnchorStateProxy, bound to a specific deployed contract.
func NewAnchorStateProxyCaller(address common.Address, caller bind.ContractCaller) (*AnchorStateProxyCaller, error) {
	contract, err := bindAnchorStateProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AnchorStateProxyCaller{contract: contract}, nil
}

// NewAnchorStateProxyTransactor creates a new write-only instance of AnchorStateProxy, bound to a specific deployed contract.
func NewAnchorStateProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*AnchorStateProxyTransactor, error) {
	contract, err := bindAnchorStateProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AnchorStateProxyTransactor{contract: contract}, nil
}

// NewAnchorStateProxyFilterer creates a new log filterer instance of AnchorStateProxy, bound to a specific deployed contract.
func NewAnchorStateProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*AnchorStateProxyFilterer, error) {
	contract, err := bindAnchorStateProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AnchorStateProxyFilterer{contract: contract}, nil
}

// bindAnchorStateProxy binds a generic wrapper to an already deployed contract.
func bindAnchorStateProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AnchorStateProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AnchorStateProxy *AnchorStateProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AnchorStateProxy.Contract.AnchorStateProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AnchorStateProxy *AnchorStateProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnchorStateProxy.Contract.AnchorStateProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AnchorStateProxy *AnchorStateProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AnchorStateProxy.Contract.AnchorStateProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AnchorStateProxy *AnchorStateProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AnchorStateProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AnchorStateProxy *AnchorStateProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnchorStateProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AnchorStateProxy *AnchorStateProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AnchorStateProxy.Contract.contract.Transact(opts, method, params...)
}

// Anchors is a free data retrieval call binding the contract method 0x7258a807.
//
// Solidity: function anchors(uint32 ) view returns(bytes32 root, uint256 l2BlockNumber)
func (_AnchorStateProxy *AnchorStateProxyCaller) Anchors(opts *bind.CallOpts, arg0 uint32) (struct {
	Root          [32]byte
	L2BlockNumber *big.Int
}, error) {
	var out []interface{}
	err := _AnchorStateProxy.contract.Call(opts, &out, "anchors", arg0)

	outstruct := new(struct {
		Root          [32]byte
		L2BlockNumber *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Root = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.L2BlockNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Anchors is a free data retrieval call binding the contract method 0x7258a807.
//
// Solidity: function anchors(uint32 ) view returns(bytes32 root, uint256 l2BlockNumber)
func (_AnchorStateProxy *AnchorStateProxySession) Anchors(arg0 uint32) (struct {
	Root          [32]byte
	L2BlockNumber *big.Int
}, error) {
	return _AnchorStateProxy.Contract.Anchors(&_AnchorStateProxy.CallOpts, arg0)
}

// Anchors is a free data retrieval call binding the contract method 0x7258a807.
//
// Solidity: function anchors(uint32 ) view returns(bytes32 root, uint256 l2BlockNumber)
func (_AnchorStateProxy *AnchorStateProxyCallerSession) Anchors(arg0 uint32) (struct {
	Root          [32]byte
	L2BlockNumber *big.Int
}, error) {
	return _AnchorStateProxy.Contract.Anchors(&_AnchorStateProxy.CallOpts, arg0)
}

// DisputeGameFactory is a free data retrieval call binding the contract method 0xf2b4e617.
//
// Solidity: function disputeGameFactory() view returns(address)
func (_AnchorStateProxy *AnchorStateProxyCaller) DisputeGameFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AnchorStateProxy.contract.Call(opts, &out, "disputeGameFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DisputeGameFactory is a free data retrieval call binding the contract method 0xf2b4e617.
//
// Solidity: function disputeGameFactory() view returns(address)
func (_AnchorStateProxy *AnchorStateProxySession) DisputeGameFactory() (common.Address, error) {
	return _AnchorStateProxy.Contract.DisputeGameFactory(&_AnchorStateProxy.CallOpts)
}

// DisputeGameFactory is a free data retrieval call binding the contract method 0xf2b4e617.
//
// Solidity: function disputeGameFactory() view returns(address)
func (_AnchorStateProxy *AnchorStateProxyCallerSession) DisputeGameFactory() (common.Address, error) {
	return _AnchorStateProxy.Contract.DisputeGameFactory(&_AnchorStateProxy.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_AnchorStateProxy *AnchorStateProxyCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AnchorStateProxy.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_AnchorStateProxy *AnchorStateProxySession) Version() (string, error) {
	return _AnchorStateProxy.Contract.Version(&_AnchorStateProxy.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_AnchorStateProxy *AnchorStateProxyCallerSession) Version() (string, error) {
	return _AnchorStateProxy.Contract.Version(&_AnchorStateProxy.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc303f0df.
//
// Solidity: function initialize((uint32,(bytes32,uint256))[] _startingAnchorRoots) returns()
func (_AnchorStateProxy *AnchorStateProxyTransactor) Initialize(opts *bind.TransactOpts, _startingAnchorRoots []AnchorStateRegistryStartingAnchorRoot) (*types.Transaction, error) {
	return _AnchorStateProxy.contract.Transact(opts, "initialize", _startingAnchorRoots)
}

// Initialize is a paid mutator transaction binding the contract method 0xc303f0df.
//
// Solidity: function initialize((uint32,(bytes32,uint256))[] _startingAnchorRoots) returns()
func (_AnchorStateProxy *AnchorStateProxySession) Initialize(_startingAnchorRoots []AnchorStateRegistryStartingAnchorRoot) (*types.Transaction, error) {
	return _AnchorStateProxy.Contract.Initialize(&_AnchorStateProxy.TransactOpts, _startingAnchorRoots)
}

// Initialize is a paid mutator transaction binding the contract method 0xc303f0df.
//
// Solidity: function initialize((uint32,(bytes32,uint256))[] _startingAnchorRoots) returns()
func (_AnchorStateProxy *AnchorStateProxyTransactorSession) Initialize(_startingAnchorRoots []AnchorStateRegistryStartingAnchorRoot) (*types.Transaction, error) {
	return _AnchorStateProxy.Contract.Initialize(&_AnchorStateProxy.TransactOpts, _startingAnchorRoots)
}

// TryUpdateAnchorState is a paid mutator transaction binding the contract method 0x838c2d1e.
//
// Solidity: function tryUpdateAnchorState() returns()
func (_AnchorStateProxy *AnchorStateProxyTransactor) TryUpdateAnchorState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AnchorStateProxy.contract.Transact(opts, "tryUpdateAnchorState")
}

// TryUpdateAnchorState is a paid mutator transaction binding the contract method 0x838c2d1e.
//
// Solidity: function tryUpdateAnchorState() returns()
func (_AnchorStateProxy *AnchorStateProxySession) TryUpdateAnchorState() (*types.Transaction, error) {
	return _AnchorStateProxy.Contract.TryUpdateAnchorState(&_AnchorStateProxy.TransactOpts)
}

// TryUpdateAnchorState is a paid mutator transaction binding the contract method 0x838c2d1e.
//
// Solidity: function tryUpdateAnchorState() returns()
func (_AnchorStateProxy *AnchorStateProxyTransactorSession) TryUpdateAnchorState() (*types.Transaction, error) {
	return _AnchorStateProxy.Contract.TryUpdateAnchorState(&_AnchorStateProxy.TransactOpts)
}

// AnchorStateProxyInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AnchorStateProxy contract.
type AnchorStateProxyInitializedIterator struct {
	Event *AnchorStateProxyInitialized // Event containing the contract specifics and raw log

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
func (it *AnchorStateProxyInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AnchorStateProxyInitialized)
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
		it.Event = new(AnchorStateProxyInitialized)
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
func (it *AnchorStateProxyInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AnchorStateProxyInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AnchorStateProxyInitialized represents a Initialized event raised by the AnchorStateProxy contract.
type AnchorStateProxyInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AnchorStateProxy *AnchorStateProxyFilterer) FilterInitialized(opts *bind.FilterOpts) (*AnchorStateProxyInitializedIterator, error) {

	logs, sub, err := _AnchorStateProxy.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AnchorStateProxyInitializedIterator{contract: _AnchorStateProxy.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AnchorStateProxy *AnchorStateProxyFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AnchorStateProxyInitialized) (event.Subscription, error) {

	logs, sub, err := _AnchorStateProxy.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AnchorStateProxyInitialized)
				if err := _AnchorStateProxy.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AnchorStateProxy *AnchorStateProxyFilterer) ParseInitialized(log types.Log) (*AnchorStateProxyInitialized, error) {
	event := new(AnchorStateProxyInitialized)
	if err := _AnchorStateProxy.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

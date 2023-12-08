// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Router

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

// RouterMetaData contains all meta data concerning the Router contract.
var RouterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdg\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"addPlugin\",\"inputs\":[{\"name\":\"_plugin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approvePlugin\",\"inputs\":[{\"name\":\"_plugin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approvedPlugins\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreasePosition\",\"inputs\":[{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decreasePositionAndSwap\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decreasePositionAndSwapETH\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decreasePositionETH\",\"inputs\":[{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"denyPlugin\",\"inputs\":[{\"name\":\"_plugin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"directPoolDeposit\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increasePosition\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increasePositionETH\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"pluginDecreasePosition\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pluginIncreasePosition\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pluginTransfer\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"plugins\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removePlugin\",\"inputs\":[{\"name\":\"_plugin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swapETHToTokens\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"swapTokensToETH\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"usdg\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vault\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"weth\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Swap\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenIn\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenOut\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterMetaData.ABI instead.
var RouterABI = RouterMetaData.ABI

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// ApprovedPlugins is a free data retrieval call binding the contract method 0x956f285e.
//
// Solidity: function approvedPlugins(address , address ) view returns(bool)
func (_Router *RouterCaller) ApprovedPlugins(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "approvedPlugins", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedPlugins is a free data retrieval call binding the contract method 0x956f285e.
//
// Solidity: function approvedPlugins(address , address ) view returns(bool)
func (_Router *RouterSession) ApprovedPlugins(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Router.Contract.ApprovedPlugins(&_Router.CallOpts, arg0, arg1)
}

// ApprovedPlugins is a free data retrieval call binding the contract method 0x956f285e.
//
// Solidity: function approvedPlugins(address , address ) view returns(bool)
func (_Router *RouterCallerSession) ApprovedPlugins(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Router.Contract.ApprovedPlugins(&_Router.CallOpts, arg0, arg1)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Router *RouterCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Router *RouterSession) Gov() (common.Address, error) {
	return _Router.Contract.Gov(&_Router.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Router *RouterCallerSession) Gov() (common.Address, error) {
	return _Router.Contract.Gov(&_Router.CallOpts)
}

// Plugins is a free data retrieval call binding the contract method 0x4b12e643.
//
// Solidity: function plugins(address ) view returns(bool)
func (_Router *RouterCaller) Plugins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "plugins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Plugins is a free data retrieval call binding the contract method 0x4b12e643.
//
// Solidity: function plugins(address ) view returns(bool)
func (_Router *RouterSession) Plugins(arg0 common.Address) (bool, error) {
	return _Router.Contract.Plugins(&_Router.CallOpts, arg0)
}

// Plugins is a free data retrieval call binding the contract method 0x4b12e643.
//
// Solidity: function plugins(address ) view returns(bool)
func (_Router *RouterCallerSession) Plugins(arg0 common.Address) (bool, error) {
	return _Router.Contract.Plugins(&_Router.CallOpts, arg0)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_Router *RouterCaller) Usdg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "usdg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_Router *RouterSession) Usdg() (common.Address, error) {
	return _Router.Contract.Usdg(&_Router.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_Router *RouterCallerSession) Usdg() (common.Address, error) {
	return _Router.Contract.Usdg(&_Router.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Router *RouterCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Router *RouterSession) Vault() (common.Address, error) {
	return _Router.Contract.Vault(&_Router.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Router *RouterCallerSession) Vault() (common.Address, error) {
	return _Router.Contract.Vault(&_Router.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Router *RouterCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Router *RouterSession) Weth() (common.Address, error) {
	return _Router.Contract.Weth(&_Router.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Router *RouterCallerSession) Weth() (common.Address, error) {
	return _Router.Contract.Weth(&_Router.CallOpts)
}

// AddPlugin is a paid mutator transaction binding the contract method 0xd8867fc8.
//
// Solidity: function addPlugin(address _plugin) returns()
func (_Router *RouterTransactor) AddPlugin(opts *bind.TransactOpts, _plugin common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "addPlugin", _plugin)
}

// AddPlugin is a paid mutator transaction binding the contract method 0xd8867fc8.
//
// Solidity: function addPlugin(address _plugin) returns()
func (_Router *RouterSession) AddPlugin(_plugin common.Address) (*types.Transaction, error) {
	return _Router.Contract.AddPlugin(&_Router.TransactOpts, _plugin)
}

// AddPlugin is a paid mutator transaction binding the contract method 0xd8867fc8.
//
// Solidity: function addPlugin(address _plugin) returns()
func (_Router *RouterTransactorSession) AddPlugin(_plugin common.Address) (*types.Transaction, error) {
	return _Router.Contract.AddPlugin(&_Router.TransactOpts, _plugin)
}

// ApprovePlugin is a paid mutator transaction binding the contract method 0x38c74dd9.
//
// Solidity: function approvePlugin(address _plugin) returns()
func (_Router *RouterTransactor) ApprovePlugin(opts *bind.TransactOpts, _plugin common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "approvePlugin", _plugin)
}

// ApprovePlugin is a paid mutator transaction binding the contract method 0x38c74dd9.
//
// Solidity: function approvePlugin(address _plugin) returns()
func (_Router *RouterSession) ApprovePlugin(_plugin common.Address) (*types.Transaction, error) {
	return _Router.Contract.ApprovePlugin(&_Router.TransactOpts, _plugin)
}

// ApprovePlugin is a paid mutator transaction binding the contract method 0x38c74dd9.
//
// Solidity: function approvePlugin(address _plugin) returns()
func (_Router *RouterTransactorSession) ApprovePlugin(_plugin common.Address) (*types.Transaction, error) {
	return _Router.Contract.ApprovePlugin(&_Router.TransactOpts, _plugin)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x90205d8c.
//
// Solidity: function decreasePosition(address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price) returns()
func (_Router *RouterTransactor) DecreasePosition(opts *bind.TransactOpts, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "decreasePosition", _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x90205d8c.
//
// Solidity: function decreasePosition(address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price) returns()
func (_Router *RouterSession) DecreasePosition(_collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Router.Contract.DecreasePosition(&_Router.TransactOpts, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x90205d8c.
//
// Solidity: function decreasePosition(address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price) returns()
func (_Router *RouterTransactorSession) DecreasePosition(_collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Router.Contract.DecreasePosition(&_Router.TransactOpts, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price)
}

// DecreasePositionAndSwap is a paid mutator transaction binding the contract method 0x5fc8500e.
//
// Solidity: function decreasePositionAndSwap(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price, uint256 _minOut) returns()
func (_Router *RouterTransactor) DecreasePositionAndSwap(opts *bind.TransactOpts, _path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int, _minOut *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "decreasePositionAndSwap", _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price, _minOut)
}

// DecreasePositionAndSwap is a paid mutator transaction binding the contract method 0x5fc8500e.
//
// Solidity: function decreasePositionAndSwap(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price, uint256 _minOut) returns()
func (_Router *RouterSession) DecreasePositionAndSwap(_path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int, _minOut *big.Int) (*types.Transaction, error) {
	return _Router.Contract.DecreasePositionAndSwap(&_Router.TransactOpts, _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price, _minOut)
}

// DecreasePositionAndSwap is a paid mutator transaction binding the contract method 0x5fc8500e.
//
// Solidity: function decreasePositionAndSwap(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price, uint256 _minOut) returns()
func (_Router *RouterTransactorSession) DecreasePositionAndSwap(_path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int, _minOut *big.Int) (*types.Transaction, error) {
	return _Router.Contract.DecreasePositionAndSwap(&_Router.TransactOpts, _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price, _minOut)
}

// DecreasePositionAndSwapETH is a paid mutator transaction binding the contract method 0x3039e37f.
//
// Solidity: function decreasePositionAndSwapETH(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price, uint256 _minOut) returns()
func (_Router *RouterTransactor) DecreasePositionAndSwapETH(opts *bind.TransactOpts, _path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int, _minOut *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "decreasePositionAndSwapETH", _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price, _minOut)
}

// DecreasePositionAndSwapETH is a paid mutator transaction binding the contract method 0x3039e37f.
//
// Solidity: function decreasePositionAndSwapETH(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price, uint256 _minOut) returns()
func (_Router *RouterSession) DecreasePositionAndSwapETH(_path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int, _minOut *big.Int) (*types.Transaction, error) {
	return _Router.Contract.DecreasePositionAndSwapETH(&_Router.TransactOpts, _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price, _minOut)
}

// DecreasePositionAndSwapETH is a paid mutator transaction binding the contract method 0x3039e37f.
//
// Solidity: function decreasePositionAndSwapETH(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price, uint256 _minOut) returns()
func (_Router *RouterTransactorSession) DecreasePositionAndSwapETH(_path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int, _minOut *big.Int) (*types.Transaction, error) {
	return _Router.Contract.DecreasePositionAndSwapETH(&_Router.TransactOpts, _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price, _minOut)
}

// DecreasePositionETH is a paid mutator transaction binding the contract method 0x430ed37c.
//
// Solidity: function decreasePositionETH(address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price) returns()
func (_Router *RouterTransactor) DecreasePositionETH(opts *bind.TransactOpts, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "decreasePositionETH", _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price)
}

// DecreasePositionETH is a paid mutator transaction binding the contract method 0x430ed37c.
//
// Solidity: function decreasePositionETH(address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price) returns()
func (_Router *RouterSession) DecreasePositionETH(_collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Router.Contract.DecreasePositionETH(&_Router.TransactOpts, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price)
}

// DecreasePositionETH is a paid mutator transaction binding the contract method 0x430ed37c.
//
// Solidity: function decreasePositionETH(address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price) returns()
func (_Router *RouterTransactorSession) DecreasePositionETH(_collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Router.Contract.DecreasePositionETH(&_Router.TransactOpts, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price)
}

// DenyPlugin is a paid mutator transaction binding the contract method 0xcedd4375.
//
// Solidity: function denyPlugin(address _plugin) returns()
func (_Router *RouterTransactor) DenyPlugin(opts *bind.TransactOpts, _plugin common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "denyPlugin", _plugin)
}

// DenyPlugin is a paid mutator transaction binding the contract method 0xcedd4375.
//
// Solidity: function denyPlugin(address _plugin) returns()
func (_Router *RouterSession) DenyPlugin(_plugin common.Address) (*types.Transaction, error) {
	return _Router.Contract.DenyPlugin(&_Router.TransactOpts, _plugin)
}

// DenyPlugin is a paid mutator transaction binding the contract method 0xcedd4375.
//
// Solidity: function denyPlugin(address _plugin) returns()
func (_Router *RouterTransactorSession) DenyPlugin(_plugin common.Address) (*types.Transaction, error) {
	return _Router.Contract.DenyPlugin(&_Router.TransactOpts, _plugin)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x90b64ad3.
//
// Solidity: function directPoolDeposit(address _token, uint256 _amount) returns()
func (_Router *RouterTransactor) DirectPoolDeposit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "directPoolDeposit", _token, _amount)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x90b64ad3.
//
// Solidity: function directPoolDeposit(address _token, uint256 _amount) returns()
func (_Router *RouterSession) DirectPoolDeposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.DirectPoolDeposit(&_Router.TransactOpts, _token, _amount)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x90b64ad3.
//
// Solidity: function directPoolDeposit(address _token, uint256 _amount) returns()
func (_Router *RouterTransactorSession) DirectPoolDeposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.DirectPoolDeposit(&_Router.TransactOpts, _token, _amount)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0xb7ddc992.
//
// Solidity: function increasePosition(address[] _path, address _indexToken, uint256 _amountIn, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _price) returns()
func (_Router *RouterTransactor) IncreasePosition(opts *bind.TransactOpts, _path []common.Address, _indexToken common.Address, _amountIn *big.Int, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _price *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "increasePosition", _path, _indexToken, _amountIn, _minOut, _sizeDelta, _isLong, _price)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0xb7ddc992.
//
// Solidity: function increasePosition(address[] _path, address _indexToken, uint256 _amountIn, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _price) returns()
func (_Router *RouterSession) IncreasePosition(_path []common.Address, _indexToken common.Address, _amountIn *big.Int, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _price *big.Int) (*types.Transaction, error) {
	return _Router.Contract.IncreasePosition(&_Router.TransactOpts, _path, _indexToken, _amountIn, _minOut, _sizeDelta, _isLong, _price)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0xb7ddc992.
//
// Solidity: function increasePosition(address[] _path, address _indexToken, uint256 _amountIn, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _price) returns()
func (_Router *RouterTransactorSession) IncreasePosition(_path []common.Address, _indexToken common.Address, _amountIn *big.Int, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _price *big.Int) (*types.Transaction, error) {
	return _Router.Contract.IncreasePosition(&_Router.TransactOpts, _path, _indexToken, _amountIn, _minOut, _sizeDelta, _isLong, _price)
}

// IncreasePositionETH is a paid mutator transaction binding the contract method 0xb32755de.
//
// Solidity: function increasePositionETH(address[] _path, address _indexToken, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _price) payable returns()
func (_Router *RouterTransactor) IncreasePositionETH(opts *bind.TransactOpts, _path []common.Address, _indexToken common.Address, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _price *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "increasePositionETH", _path, _indexToken, _minOut, _sizeDelta, _isLong, _price)
}

// IncreasePositionETH is a paid mutator transaction binding the contract method 0xb32755de.
//
// Solidity: function increasePositionETH(address[] _path, address _indexToken, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _price) payable returns()
func (_Router *RouterSession) IncreasePositionETH(_path []common.Address, _indexToken common.Address, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _price *big.Int) (*types.Transaction, error) {
	return _Router.Contract.IncreasePositionETH(&_Router.TransactOpts, _path, _indexToken, _minOut, _sizeDelta, _isLong, _price)
}

// IncreasePositionETH is a paid mutator transaction binding the contract method 0xb32755de.
//
// Solidity: function increasePositionETH(address[] _path, address _indexToken, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _price) payable returns()
func (_Router *RouterTransactorSession) IncreasePositionETH(_path []common.Address, _indexToken common.Address, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _price *big.Int) (*types.Transaction, error) {
	return _Router.Contract.IncreasePositionETH(&_Router.TransactOpts, _path, _indexToken, _minOut, _sizeDelta, _isLong, _price)
}

// PluginDecreasePosition is a paid mutator transaction binding the contract method 0x2662166b.
//
// Solidity: function pluginDecreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_Router *RouterTransactor) PluginDecreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "pluginDecreasePosition", _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// PluginDecreasePosition is a paid mutator transaction binding the contract method 0x2662166b.
//
// Solidity: function pluginDecreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_Router *RouterSession) PluginDecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _Router.Contract.PluginDecreasePosition(&_Router.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// PluginDecreasePosition is a paid mutator transaction binding the contract method 0x2662166b.
//
// Solidity: function pluginDecreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_Router *RouterTransactorSession) PluginDecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _Router.Contract.PluginDecreasePosition(&_Router.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// PluginIncreasePosition is a paid mutator transaction binding the contract method 0x1f1dd176.
//
// Solidity: function pluginIncreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_Router *RouterTransactor) PluginIncreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "pluginIncreasePosition", _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// PluginIncreasePosition is a paid mutator transaction binding the contract method 0x1f1dd176.
//
// Solidity: function pluginIncreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_Router *RouterSession) PluginIncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _Router.Contract.PluginIncreasePosition(&_Router.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// PluginIncreasePosition is a paid mutator transaction binding the contract method 0x1f1dd176.
//
// Solidity: function pluginIncreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_Router *RouterTransactorSession) PluginIncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _Router.Contract.PluginIncreasePosition(&_Router.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// PluginTransfer is a paid mutator transaction binding the contract method 0x1b827878.
//
// Solidity: function pluginTransfer(address _token, address _account, address _receiver, uint256 _amount) returns()
func (_Router *RouterTransactor) PluginTransfer(opts *bind.TransactOpts, _token common.Address, _account common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "pluginTransfer", _token, _account, _receiver, _amount)
}

// PluginTransfer is a paid mutator transaction binding the contract method 0x1b827878.
//
// Solidity: function pluginTransfer(address _token, address _account, address _receiver, uint256 _amount) returns()
func (_Router *RouterSession) PluginTransfer(_token common.Address, _account common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.PluginTransfer(&_Router.TransactOpts, _token, _account, _receiver, _amount)
}

// PluginTransfer is a paid mutator transaction binding the contract method 0x1b827878.
//
// Solidity: function pluginTransfer(address _token, address _account, address _receiver, uint256 _amount) returns()
func (_Router *RouterTransactorSession) PluginTransfer(_token common.Address, _account common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.PluginTransfer(&_Router.TransactOpts, _token, _account, _receiver, _amount)
}

// RemovePlugin is a paid mutator transaction binding the contract method 0xa4d95b64.
//
// Solidity: function removePlugin(address _plugin) returns()
func (_Router *RouterTransactor) RemovePlugin(opts *bind.TransactOpts, _plugin common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "removePlugin", _plugin)
}

// RemovePlugin is a paid mutator transaction binding the contract method 0xa4d95b64.
//
// Solidity: function removePlugin(address _plugin) returns()
func (_Router *RouterSession) RemovePlugin(_plugin common.Address) (*types.Transaction, error) {
	return _Router.Contract.RemovePlugin(&_Router.TransactOpts, _plugin)
}

// RemovePlugin is a paid mutator transaction binding the contract method 0xa4d95b64.
//
// Solidity: function removePlugin(address _plugin) returns()
func (_Router *RouterTransactorSession) RemovePlugin(_plugin common.Address) (*types.Transaction, error) {
	return _Router.Contract.RemovePlugin(&_Router.TransactOpts, _plugin)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Router *RouterTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Router *RouterSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetGov(&_Router.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Router *RouterTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetGov(&_Router.TransactOpts, _gov)
}

// Swap is a paid mutator transaction binding the contract method 0x6023e966.
//
// Solidity: function swap(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_Router *RouterTransactor) Swap(opts *bind.TransactOpts, _path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swap", _path, _amountIn, _minOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x6023e966.
//
// Solidity: function swap(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_Router *RouterSession) Swap(_path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Router.Contract.Swap(&_Router.TransactOpts, _path, _amountIn, _minOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x6023e966.
//
// Solidity: function swap(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_Router *RouterTransactorSession) Swap(_path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Router.Contract.Swap(&_Router.TransactOpts, _path, _amountIn, _minOut, _receiver)
}

// SwapETHToTokens is a paid mutator transaction binding the contract method 0xabe68eaa.
//
// Solidity: function swapETHToTokens(address[] _path, uint256 _minOut, address _receiver) payable returns()
func (_Router *RouterTransactor) SwapETHToTokens(opts *bind.TransactOpts, _path []common.Address, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapETHToTokens", _path, _minOut, _receiver)
}

// SwapETHToTokens is a paid mutator transaction binding the contract method 0xabe68eaa.
//
// Solidity: function swapETHToTokens(address[] _path, uint256 _minOut, address _receiver) payable returns()
func (_Router *RouterSession) SwapETHToTokens(_path []common.Address, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Router.Contract.SwapETHToTokens(&_Router.TransactOpts, _path, _minOut, _receiver)
}

// SwapETHToTokens is a paid mutator transaction binding the contract method 0xabe68eaa.
//
// Solidity: function swapETHToTokens(address[] _path, uint256 _minOut, address _receiver) payable returns()
func (_Router *RouterTransactorSession) SwapETHToTokens(_path []common.Address, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Router.Contract.SwapETHToTokens(&_Router.TransactOpts, _path, _minOut, _receiver)
}

// SwapTokensToETH is a paid mutator transaction binding the contract method 0x2d4ba6a7.
//
// Solidity: function swapTokensToETH(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_Router *RouterTransactor) SwapTokensToETH(opts *bind.TransactOpts, _path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapTokensToETH", _path, _amountIn, _minOut, _receiver)
}

// SwapTokensToETH is a paid mutator transaction binding the contract method 0x2d4ba6a7.
//
// Solidity: function swapTokensToETH(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_Router *RouterSession) SwapTokensToETH(_path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Router.Contract.SwapTokensToETH(&_Router.TransactOpts, _path, _amountIn, _minOut, _receiver)
}

// SwapTokensToETH is a paid mutator transaction binding the contract method 0x2d4ba6a7.
//
// Solidity: function swapTokensToETH(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_Router *RouterTransactorSession) SwapTokensToETH(_path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Router.Contract.SwapTokensToETH(&_Router.TransactOpts, _path, _amountIn, _minOut, _receiver)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactorSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}

// RouterSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Router contract.
type RouterSwapIterator struct {
	Event *RouterSwap // Event containing the contract specifics and raw log

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
func (it *RouterSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterSwap)
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
		it.Event = new(RouterSwap)
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
func (it *RouterSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterSwap represents a Swap event raised by the Router contract.
type RouterSwap struct {
	Account   common.Address
	TokenIn   common.Address
	TokenOut  common.Address
	AmountIn  *big.Int
	AmountOut *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xcd3829a3813dc3cdd188fd3d01dcf3268c16be2fdd2dd21d0665418816e46062.
//
// Solidity: event Swap(address account, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut)
func (_Router *RouterFilterer) FilterSwap(opts *bind.FilterOpts) (*RouterSwapIterator, error) {

	logs, sub, err := _Router.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &RouterSwapIterator{contract: _Router.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xcd3829a3813dc3cdd188fd3d01dcf3268c16be2fdd2dd21d0665418816e46062.
//
// Solidity: event Swap(address account, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut)
func (_Router *RouterFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *RouterSwap) (event.Subscription, error) {

	logs, sub, err := _Router.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterSwap)
				if err := _Router.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xcd3829a3813dc3cdd188fd3d01dcf3268c16be2fdd2dd21d0665418816e46062.
//
// Solidity: event Swap(address account, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut)
func (_Router *RouterFilterer) ParseSwap(log types.Log) (*RouterSwap, error) {
	event := new(RouterSwap)
	if err := _Router.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

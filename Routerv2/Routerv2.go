// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Routerv2

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

// Routerv2MetaData contains all meta data concerning the Routerv2 contract.
var Routerv2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdg\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"addPlugin\",\"inputs\":[{\"name\":\"_plugin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approvePlugin\",\"inputs\":[{\"name\":\"_plugin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approvedPlugins\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreasePosition\",\"inputs\":[{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decreasePositionAndSwap\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decreasePositionAndSwapETH\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decreasePositionETH\",\"inputs\":[{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"denyPlugin\",\"inputs\":[{\"name\":\"_plugin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"directPoolDeposit\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increasePosition\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increasePositionETH\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"pluginDecreasePosition\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pluginIncreasePosition\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pluginTransfer\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"plugins\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removePlugin\",\"inputs\":[{\"name\":\"_plugin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swapETHToTokens\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"swapTokensToETH\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"usdg\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vault\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"weth\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Swap\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenIn\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenOut\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// Routerv2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Routerv2MetaData.ABI instead.
var Routerv2ABI = Routerv2MetaData.ABI

// Routerv2 is an auto generated Go binding around an Ethereum contract.
type Routerv2 struct {
	Routerv2Caller     // Read-only binding to the contract
	Routerv2Transactor // Write-only binding to the contract
	Routerv2Filterer   // Log filterer for contract events
}

// Routerv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Routerv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Routerv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Routerv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Routerv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Routerv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Routerv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Routerv2Session struct {
	Contract     *Routerv2         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Routerv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Routerv2CallerSession struct {
	Contract *Routerv2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// Routerv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Routerv2TransactorSession struct {
	Contract     *Routerv2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Routerv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Routerv2Raw struct {
	Contract *Routerv2 // Generic contract binding to access the raw methods on
}

// Routerv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Routerv2CallerRaw struct {
	Contract *Routerv2Caller // Generic read-only contract binding to access the raw methods on
}

// Routerv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Routerv2TransactorRaw struct {
	Contract *Routerv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewRouterv2 creates a new instance of Routerv2, bound to a specific deployed contract.
func NewRouterv2(address common.Address, backend bind.ContractBackend) (*Routerv2, error) {
	contract, err := bindRouterv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Routerv2{Routerv2Caller: Routerv2Caller{contract: contract}, Routerv2Transactor: Routerv2Transactor{contract: contract}, Routerv2Filterer: Routerv2Filterer{contract: contract}}, nil
}

// NewRouterv2Caller creates a new read-only instance of Routerv2, bound to a specific deployed contract.
func NewRouterv2Caller(address common.Address, caller bind.ContractCaller) (*Routerv2Caller, error) {
	contract, err := bindRouterv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Routerv2Caller{contract: contract}, nil
}

// NewRouterv2Transactor creates a new write-only instance of Routerv2, bound to a specific deployed contract.
func NewRouterv2Transactor(address common.Address, transactor bind.ContractTransactor) (*Routerv2Transactor, error) {
	contract, err := bindRouterv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Routerv2Transactor{contract: contract}, nil
}

// NewRouterv2Filterer creates a new log filterer instance of Routerv2, bound to a specific deployed contract.
func NewRouterv2Filterer(address common.Address, filterer bind.ContractFilterer) (*Routerv2Filterer, error) {
	contract, err := bindRouterv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Routerv2Filterer{contract: contract}, nil
}

// bindRouterv2 binds a generic wrapper to an already deployed contract.
func bindRouterv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Routerv2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Routerv2 *Routerv2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Routerv2.Contract.Routerv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Routerv2 *Routerv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Routerv2.Contract.Routerv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Routerv2 *Routerv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Routerv2.Contract.Routerv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Routerv2 *Routerv2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Routerv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Routerv2 *Routerv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Routerv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Routerv2 *Routerv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Routerv2.Contract.contract.Transact(opts, method, params...)
}

// ApprovedPlugins is a free data retrieval call binding the contract method 0x956f285e.
//
// Solidity: function approvedPlugins(address , address ) view returns(bool)
func (_Routerv2 *Routerv2Caller) ApprovedPlugins(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Routerv2.contract.Call(opts, &out, "approvedPlugins", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedPlugins is a free data retrieval call binding the contract method 0x956f285e.
//
// Solidity: function approvedPlugins(address , address ) view returns(bool)
func (_Routerv2 *Routerv2Session) ApprovedPlugins(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Routerv2.Contract.ApprovedPlugins(&_Routerv2.CallOpts, arg0, arg1)
}

// ApprovedPlugins is a free data retrieval call binding the contract method 0x956f285e.
//
// Solidity: function approvedPlugins(address , address ) view returns(bool)
func (_Routerv2 *Routerv2CallerSession) ApprovedPlugins(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Routerv2.Contract.ApprovedPlugins(&_Routerv2.CallOpts, arg0, arg1)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Routerv2 *Routerv2Caller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Routerv2.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Routerv2 *Routerv2Session) Gov() (common.Address, error) {
	return _Routerv2.Contract.Gov(&_Routerv2.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Routerv2 *Routerv2CallerSession) Gov() (common.Address, error) {
	return _Routerv2.Contract.Gov(&_Routerv2.CallOpts)
}

// Plugins is a free data retrieval call binding the contract method 0x4b12e643.
//
// Solidity: function plugins(address ) view returns(bool)
func (_Routerv2 *Routerv2Caller) Plugins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Routerv2.contract.Call(opts, &out, "plugins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Plugins is a free data retrieval call binding the contract method 0x4b12e643.
//
// Solidity: function plugins(address ) view returns(bool)
func (_Routerv2 *Routerv2Session) Plugins(arg0 common.Address) (bool, error) {
	return _Routerv2.Contract.Plugins(&_Routerv2.CallOpts, arg0)
}

// Plugins is a free data retrieval call binding the contract method 0x4b12e643.
//
// Solidity: function plugins(address ) view returns(bool)
func (_Routerv2 *Routerv2CallerSession) Plugins(arg0 common.Address) (bool, error) {
	return _Routerv2.Contract.Plugins(&_Routerv2.CallOpts, arg0)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_Routerv2 *Routerv2Caller) Usdg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Routerv2.contract.Call(opts, &out, "usdg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_Routerv2 *Routerv2Session) Usdg() (common.Address, error) {
	return _Routerv2.Contract.Usdg(&_Routerv2.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_Routerv2 *Routerv2CallerSession) Usdg() (common.Address, error) {
	return _Routerv2.Contract.Usdg(&_Routerv2.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Routerv2 *Routerv2Caller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Routerv2.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Routerv2 *Routerv2Session) Vault() (common.Address, error) {
	return _Routerv2.Contract.Vault(&_Routerv2.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Routerv2 *Routerv2CallerSession) Vault() (common.Address, error) {
	return _Routerv2.Contract.Vault(&_Routerv2.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Routerv2 *Routerv2Caller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Routerv2.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Routerv2 *Routerv2Session) Weth() (common.Address, error) {
	return _Routerv2.Contract.Weth(&_Routerv2.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Routerv2 *Routerv2CallerSession) Weth() (common.Address, error) {
	return _Routerv2.Contract.Weth(&_Routerv2.CallOpts)
}

// AddPlugin is a paid mutator transaction binding the contract method 0xd8867fc8.
//
// Solidity: function addPlugin(address _plugin) returns()
func (_Routerv2 *Routerv2Transactor) AddPlugin(opts *bind.TransactOpts, _plugin common.Address) (*types.Transaction, error) {
	return _Routerv2.contract.Transact(opts, "addPlugin", _plugin)
}

// AddPlugin is a paid mutator transaction binding the contract method 0xd8867fc8.
//
// Solidity: function addPlugin(address _plugin) returns()
func (_Routerv2 *Routerv2Session) AddPlugin(_plugin common.Address) (*types.Transaction, error) {
	return _Routerv2.Contract.AddPlugin(&_Routerv2.TransactOpts, _plugin)
}

// AddPlugin is a paid mutator transaction binding the contract method 0xd8867fc8.
//
// Solidity: function addPlugin(address _plugin) returns()
func (_Routerv2 *Routerv2TransactorSession) AddPlugin(_plugin common.Address) (*types.Transaction, error) {
	return _Routerv2.Contract.AddPlugin(&_Routerv2.TransactOpts, _plugin)
}

// ApprovePlugin is a paid mutator transaction binding the contract method 0x38c74dd9.
//
// Solidity: function approvePlugin(address _plugin) returns()
func (_Routerv2 *Routerv2Transactor) ApprovePlugin(opts *bind.TransactOpts, _plugin common.Address) (*types.Transaction, error) {
	return _Routerv2.contract.Transact(opts, "approvePlugin", _plugin)
}

// ApprovePlugin is a paid mutator transaction binding the contract method 0x38c74dd9.
//
// Solidity: function approvePlugin(address _plugin) returns()
func (_Routerv2 *Routerv2Session) ApprovePlugin(_plugin common.Address) (*types.Transaction, error) {
	return _Routerv2.Contract.ApprovePlugin(&_Routerv2.TransactOpts, _plugin)
}

// ApprovePlugin is a paid mutator transaction binding the contract method 0x38c74dd9.
//
// Solidity: function approvePlugin(address _plugin) returns()
func (_Routerv2 *Routerv2TransactorSession) ApprovePlugin(_plugin common.Address) (*types.Transaction, error) {
	return _Routerv2.Contract.ApprovePlugin(&_Routerv2.TransactOpts, _plugin)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x90205d8c.
//
// Solidity: function decreasePosition(address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price) returns()
func (_Routerv2 *Routerv2Transactor) DecreasePosition(opts *bind.TransactOpts, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Routerv2.contract.Transact(opts, "decreasePosition", _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x90205d8c.
//
// Solidity: function decreasePosition(address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price) returns()
func (_Routerv2 *Routerv2Session) DecreasePosition(_collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Routerv2.Contract.DecreasePosition(&_Routerv2.TransactOpts, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x90205d8c.
//
// Solidity: function decreasePosition(address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price) returns()
func (_Routerv2 *Routerv2TransactorSession) DecreasePosition(_collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Routerv2.Contract.DecreasePosition(&_Routerv2.TransactOpts, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price)
}

// DecreasePositionAndSwap is a paid mutator transaction binding the contract method 0x5fc8500e.
//
// Solidity: function decreasePositionAndSwap(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price, uint256 _minOut) returns()
func (_Routerv2 *Routerv2Transactor) DecreasePositionAndSwap(opts *bind.TransactOpts, _path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int, _minOut *big.Int) (*types.Transaction, error) {
	return _Routerv2.contract.Transact(opts, "decreasePositionAndSwap", _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price, _minOut)
}

// DecreasePositionAndSwap is a paid mutator transaction binding the contract method 0x5fc8500e.
//
// Solidity: function decreasePositionAndSwap(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price, uint256 _minOut) returns()
func (_Routerv2 *Routerv2Session) DecreasePositionAndSwap(_path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int, _minOut *big.Int) (*types.Transaction, error) {
	return _Routerv2.Contract.DecreasePositionAndSwap(&_Routerv2.TransactOpts, _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price, _minOut)
}

// DecreasePositionAndSwap is a paid mutator transaction binding the contract method 0x5fc8500e.
//
// Solidity: function decreasePositionAndSwap(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price, uint256 _minOut) returns()
func (_Routerv2 *Routerv2TransactorSession) DecreasePositionAndSwap(_path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int, _minOut *big.Int) (*types.Transaction, error) {
	return _Routerv2.Contract.DecreasePositionAndSwap(&_Routerv2.TransactOpts, _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price, _minOut)
}

// DecreasePositionAndSwapETH is a paid mutator transaction binding the contract method 0x3039e37f.
//
// Solidity: function decreasePositionAndSwapETH(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price, uint256 _minOut) returns()
func (_Routerv2 *Routerv2Transactor) DecreasePositionAndSwapETH(opts *bind.TransactOpts, _path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int, _minOut *big.Int) (*types.Transaction, error) {
	return _Routerv2.contract.Transact(opts, "decreasePositionAndSwapETH", _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price, _minOut)
}

// DecreasePositionAndSwapETH is a paid mutator transaction binding the contract method 0x3039e37f.
//
// Solidity: function decreasePositionAndSwapETH(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price, uint256 _minOut) returns()
func (_Routerv2 *Routerv2Session) DecreasePositionAndSwapETH(_path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int, _minOut *big.Int) (*types.Transaction, error) {
	return _Routerv2.Contract.DecreasePositionAndSwapETH(&_Routerv2.TransactOpts, _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price, _minOut)
}

// DecreasePositionAndSwapETH is a paid mutator transaction binding the contract method 0x3039e37f.
//
// Solidity: function decreasePositionAndSwapETH(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price, uint256 _minOut) returns()
func (_Routerv2 *Routerv2TransactorSession) DecreasePositionAndSwapETH(_path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int, _minOut *big.Int) (*types.Transaction, error) {
	return _Routerv2.Contract.DecreasePositionAndSwapETH(&_Routerv2.TransactOpts, _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price, _minOut)
}

// DecreasePositionETH is a paid mutator transaction binding the contract method 0x430ed37c.
//
// Solidity: function decreasePositionETH(address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price) returns()
func (_Routerv2 *Routerv2Transactor) DecreasePositionETH(opts *bind.TransactOpts, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Routerv2.contract.Transact(opts, "decreasePositionETH", _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price)
}

// DecreasePositionETH is a paid mutator transaction binding the contract method 0x430ed37c.
//
// Solidity: function decreasePositionETH(address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price) returns()
func (_Routerv2 *Routerv2Session) DecreasePositionETH(_collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Routerv2.Contract.DecreasePositionETH(&_Routerv2.TransactOpts, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price)
}

// DecreasePositionETH is a paid mutator transaction binding the contract method 0x430ed37c.
//
// Solidity: function decreasePositionETH(address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price) returns()
func (_Routerv2 *Routerv2TransactorSession) DecreasePositionETH(_collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _Routerv2.Contract.DecreasePositionETH(&_Routerv2.TransactOpts, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price)
}

// DenyPlugin is a paid mutator transaction binding the contract method 0xcedd4375.
//
// Solidity: function denyPlugin(address _plugin) returns()
func (_Routerv2 *Routerv2Transactor) DenyPlugin(opts *bind.TransactOpts, _plugin common.Address) (*types.Transaction, error) {
	return _Routerv2.contract.Transact(opts, "denyPlugin", _plugin)
}

// DenyPlugin is a paid mutator transaction binding the contract method 0xcedd4375.
//
// Solidity: function denyPlugin(address _plugin) returns()
func (_Routerv2 *Routerv2Session) DenyPlugin(_plugin common.Address) (*types.Transaction, error) {
	return _Routerv2.Contract.DenyPlugin(&_Routerv2.TransactOpts, _plugin)
}

// DenyPlugin is a paid mutator transaction binding the contract method 0xcedd4375.
//
// Solidity: function denyPlugin(address _plugin) returns()
func (_Routerv2 *Routerv2TransactorSession) DenyPlugin(_plugin common.Address) (*types.Transaction, error) {
	return _Routerv2.Contract.DenyPlugin(&_Routerv2.TransactOpts, _plugin)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x90b64ad3.
//
// Solidity: function directPoolDeposit(address _token, uint256 _amount) returns()
func (_Routerv2 *Routerv2Transactor) DirectPoolDeposit(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Routerv2.contract.Transact(opts, "directPoolDeposit", _token, _amount)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x90b64ad3.
//
// Solidity: function directPoolDeposit(address _token, uint256 _amount) returns()
func (_Routerv2 *Routerv2Session) DirectPoolDeposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Routerv2.Contract.DirectPoolDeposit(&_Routerv2.TransactOpts, _token, _amount)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x90b64ad3.
//
// Solidity: function directPoolDeposit(address _token, uint256 _amount) returns()
func (_Routerv2 *Routerv2TransactorSession) DirectPoolDeposit(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Routerv2.Contract.DirectPoolDeposit(&_Routerv2.TransactOpts, _token, _amount)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0xb7ddc992.
//
// Solidity: function increasePosition(address[] _path, address _indexToken, uint256 _amountIn, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _price) returns()
func (_Routerv2 *Routerv2Transactor) IncreasePosition(opts *bind.TransactOpts, _path []common.Address, _indexToken common.Address, _amountIn *big.Int, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _price *big.Int) (*types.Transaction, error) {
	return _Routerv2.contract.Transact(opts, "increasePosition", _path, _indexToken, _amountIn, _minOut, _sizeDelta, _isLong, _price)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0xb7ddc992.
//
// Solidity: function increasePosition(address[] _path, address _indexToken, uint256 _amountIn, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _price) returns()
func (_Routerv2 *Routerv2Session) IncreasePosition(_path []common.Address, _indexToken common.Address, _amountIn *big.Int, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _price *big.Int) (*types.Transaction, error) {
	return _Routerv2.Contract.IncreasePosition(&_Routerv2.TransactOpts, _path, _indexToken, _amountIn, _minOut, _sizeDelta, _isLong, _price)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0xb7ddc992.
//
// Solidity: function increasePosition(address[] _path, address _indexToken, uint256 _amountIn, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _price) returns()
func (_Routerv2 *Routerv2TransactorSession) IncreasePosition(_path []common.Address, _indexToken common.Address, _amountIn *big.Int, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _price *big.Int) (*types.Transaction, error) {
	return _Routerv2.Contract.IncreasePosition(&_Routerv2.TransactOpts, _path, _indexToken, _amountIn, _minOut, _sizeDelta, _isLong, _price)
}

// IncreasePositionETH is a paid mutator transaction binding the contract method 0xb32755de.
//
// Solidity: function increasePositionETH(address[] _path, address _indexToken, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _price) payable returns()
func (_Routerv2 *Routerv2Transactor) IncreasePositionETH(opts *bind.TransactOpts, _path []common.Address, _indexToken common.Address, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _price *big.Int) (*types.Transaction, error) {
	return _Routerv2.contract.Transact(opts, "increasePositionETH", _path, _indexToken, _minOut, _sizeDelta, _isLong, _price)
}

// IncreasePositionETH is a paid mutator transaction binding the contract method 0xb32755de.
//
// Solidity: function increasePositionETH(address[] _path, address _indexToken, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _price) payable returns()
func (_Routerv2 *Routerv2Session) IncreasePositionETH(_path []common.Address, _indexToken common.Address, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _price *big.Int) (*types.Transaction, error) {
	return _Routerv2.Contract.IncreasePositionETH(&_Routerv2.TransactOpts, _path, _indexToken, _minOut, _sizeDelta, _isLong, _price)
}

// IncreasePositionETH is a paid mutator transaction binding the contract method 0xb32755de.
//
// Solidity: function increasePositionETH(address[] _path, address _indexToken, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _price) payable returns()
func (_Routerv2 *Routerv2TransactorSession) IncreasePositionETH(_path []common.Address, _indexToken common.Address, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _price *big.Int) (*types.Transaction, error) {
	return _Routerv2.Contract.IncreasePositionETH(&_Routerv2.TransactOpts, _path, _indexToken, _minOut, _sizeDelta, _isLong, _price)
}

// PluginDecreasePosition is a paid mutator transaction binding the contract method 0x2662166b.
//
// Solidity: function pluginDecreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_Routerv2 *Routerv2Transactor) PluginDecreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _Routerv2.contract.Transact(opts, "pluginDecreasePosition", _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// PluginDecreasePosition is a paid mutator transaction binding the contract method 0x2662166b.
//
// Solidity: function pluginDecreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_Routerv2 *Routerv2Session) PluginDecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _Routerv2.Contract.PluginDecreasePosition(&_Routerv2.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// PluginDecreasePosition is a paid mutator transaction binding the contract method 0x2662166b.
//
// Solidity: function pluginDecreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_Routerv2 *Routerv2TransactorSession) PluginDecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _Routerv2.Contract.PluginDecreasePosition(&_Routerv2.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// PluginIncreasePosition is a paid mutator transaction binding the contract method 0x1f1dd176.
//
// Solidity: function pluginIncreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_Routerv2 *Routerv2Transactor) PluginIncreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _Routerv2.contract.Transact(opts, "pluginIncreasePosition", _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// PluginIncreasePosition is a paid mutator transaction binding the contract method 0x1f1dd176.
//
// Solidity: function pluginIncreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_Routerv2 *Routerv2Session) PluginIncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _Routerv2.Contract.PluginIncreasePosition(&_Routerv2.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// PluginIncreasePosition is a paid mutator transaction binding the contract method 0x1f1dd176.
//
// Solidity: function pluginIncreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_Routerv2 *Routerv2TransactorSession) PluginIncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _Routerv2.Contract.PluginIncreasePosition(&_Routerv2.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// PluginTransfer is a paid mutator transaction binding the contract method 0x1b827878.
//
// Solidity: function pluginTransfer(address _token, address _account, address _receiver, uint256 _amount) returns()
func (_Routerv2 *Routerv2Transactor) PluginTransfer(opts *bind.TransactOpts, _token common.Address, _account common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Routerv2.contract.Transact(opts, "pluginTransfer", _token, _account, _receiver, _amount)
}

// PluginTransfer is a paid mutator transaction binding the contract method 0x1b827878.
//
// Solidity: function pluginTransfer(address _token, address _account, address _receiver, uint256 _amount) returns()
func (_Routerv2 *Routerv2Session) PluginTransfer(_token common.Address, _account common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Routerv2.Contract.PluginTransfer(&_Routerv2.TransactOpts, _token, _account, _receiver, _amount)
}

// PluginTransfer is a paid mutator transaction binding the contract method 0x1b827878.
//
// Solidity: function pluginTransfer(address _token, address _account, address _receiver, uint256 _amount) returns()
func (_Routerv2 *Routerv2TransactorSession) PluginTransfer(_token common.Address, _account common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Routerv2.Contract.PluginTransfer(&_Routerv2.TransactOpts, _token, _account, _receiver, _amount)
}

// RemovePlugin is a paid mutator transaction binding the contract method 0xa4d95b64.
//
// Solidity: function removePlugin(address _plugin) returns()
func (_Routerv2 *Routerv2Transactor) RemovePlugin(opts *bind.TransactOpts, _plugin common.Address) (*types.Transaction, error) {
	return _Routerv2.contract.Transact(opts, "removePlugin", _plugin)
}

// RemovePlugin is a paid mutator transaction binding the contract method 0xa4d95b64.
//
// Solidity: function removePlugin(address _plugin) returns()
func (_Routerv2 *Routerv2Session) RemovePlugin(_plugin common.Address) (*types.Transaction, error) {
	return _Routerv2.Contract.RemovePlugin(&_Routerv2.TransactOpts, _plugin)
}

// RemovePlugin is a paid mutator transaction binding the contract method 0xa4d95b64.
//
// Solidity: function removePlugin(address _plugin) returns()
func (_Routerv2 *Routerv2TransactorSession) RemovePlugin(_plugin common.Address) (*types.Transaction, error) {
	return _Routerv2.Contract.RemovePlugin(&_Routerv2.TransactOpts, _plugin)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Routerv2 *Routerv2Transactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _Routerv2.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Routerv2 *Routerv2Session) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Routerv2.Contract.SetGov(&_Routerv2.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Routerv2 *Routerv2TransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Routerv2.Contract.SetGov(&_Routerv2.TransactOpts, _gov)
}

// Swap is a paid mutator transaction binding the contract method 0x6023e966.
//
// Solidity: function swap(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_Routerv2 *Routerv2Transactor) Swap(opts *bind.TransactOpts, _path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Routerv2.contract.Transact(opts, "swap", _path, _amountIn, _minOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x6023e966.
//
// Solidity: function swap(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_Routerv2 *Routerv2Session) Swap(_path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Routerv2.Contract.Swap(&_Routerv2.TransactOpts, _path, _amountIn, _minOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x6023e966.
//
// Solidity: function swap(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_Routerv2 *Routerv2TransactorSession) Swap(_path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Routerv2.Contract.Swap(&_Routerv2.TransactOpts, _path, _amountIn, _minOut, _receiver)
}

// SwapETHToTokens is a paid mutator transaction binding the contract method 0xabe68eaa.
//
// Solidity: function swapETHToTokens(address[] _path, uint256 _minOut, address _receiver) payable returns()
func (_Routerv2 *Routerv2Transactor) SwapETHToTokens(opts *bind.TransactOpts, _path []common.Address, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Routerv2.contract.Transact(opts, "swapETHToTokens", _path, _minOut, _receiver)
}

// SwapETHToTokens is a paid mutator transaction binding the contract method 0xabe68eaa.
//
// Solidity: function swapETHToTokens(address[] _path, uint256 _minOut, address _receiver) payable returns()
func (_Routerv2 *Routerv2Session) SwapETHToTokens(_path []common.Address, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Routerv2.Contract.SwapETHToTokens(&_Routerv2.TransactOpts, _path, _minOut, _receiver)
}

// SwapETHToTokens is a paid mutator transaction binding the contract method 0xabe68eaa.
//
// Solidity: function swapETHToTokens(address[] _path, uint256 _minOut, address _receiver) payable returns()
func (_Routerv2 *Routerv2TransactorSession) SwapETHToTokens(_path []common.Address, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Routerv2.Contract.SwapETHToTokens(&_Routerv2.TransactOpts, _path, _minOut, _receiver)
}

// SwapTokensToETH is a paid mutator transaction binding the contract method 0x2d4ba6a7.
//
// Solidity: function swapTokensToETH(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_Routerv2 *Routerv2Transactor) SwapTokensToETH(opts *bind.TransactOpts, _path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Routerv2.contract.Transact(opts, "swapTokensToETH", _path, _amountIn, _minOut, _receiver)
}

// SwapTokensToETH is a paid mutator transaction binding the contract method 0x2d4ba6a7.
//
// Solidity: function swapTokensToETH(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_Routerv2 *Routerv2Session) SwapTokensToETH(_path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Routerv2.Contract.SwapTokensToETH(&_Routerv2.TransactOpts, _path, _amountIn, _minOut, _receiver)
}

// SwapTokensToETH is a paid mutator transaction binding the contract method 0x2d4ba6a7.
//
// Solidity: function swapTokensToETH(address[] _path, uint256 _amountIn, uint256 _minOut, address _receiver) returns()
func (_Routerv2 *Routerv2TransactorSession) SwapTokensToETH(_path []common.Address, _amountIn *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _Routerv2.Contract.SwapTokensToETH(&_Routerv2.TransactOpts, _path, _amountIn, _minOut, _receiver)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Routerv2 *Routerv2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Routerv2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Routerv2 *Routerv2Session) Receive() (*types.Transaction, error) {
	return _Routerv2.Contract.Receive(&_Routerv2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Routerv2 *Routerv2TransactorSession) Receive() (*types.Transaction, error) {
	return _Routerv2.Contract.Receive(&_Routerv2.TransactOpts)
}

// Routerv2SwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Routerv2 contract.
type Routerv2SwapIterator struct {
	Event *Routerv2Swap // Event containing the contract specifics and raw log

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
func (it *Routerv2SwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Routerv2Swap)
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
		it.Event = new(Routerv2Swap)
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
func (it *Routerv2SwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Routerv2SwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Routerv2Swap represents a Swap event raised by the Routerv2 contract.
type Routerv2Swap struct {
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
func (_Routerv2 *Routerv2Filterer) FilterSwap(opts *bind.FilterOpts) (*Routerv2SwapIterator, error) {

	logs, sub, err := _Routerv2.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &Routerv2SwapIterator{contract: _Routerv2.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xcd3829a3813dc3cdd188fd3d01dcf3268c16be2fdd2dd21d0665418816e46062.
//
// Solidity: event Swap(address account, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut)
func (_Routerv2 *Routerv2Filterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *Routerv2Swap) (event.Subscription, error) {

	logs, sub, err := _Routerv2.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Routerv2Swap)
				if err := _Routerv2.contract.UnpackLog(event, "Swap", log); err != nil {
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
func (_Routerv2 *Routerv2Filterer) ParseSwap(log types.Log) (*Routerv2Swap, error) {
	event := new(Routerv2Swap)
	if err := _Routerv2.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package YieldToken

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

// YieldTokenMetaData contains all meta data concerning the YieldToken contract.
var YieldTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_initialSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addAdmin\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addNonStakingAccount\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"admins\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inWhitelistMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonStakingAccounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonStakingSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverClaim\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeAdmin\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeNonStakingAccount\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInWhitelistMode\",\"inputs\":[{\"name\":\"_inWhitelistMode\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInfo\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWhitelistedHandler\",\"inputs\":[{\"name\":\"_handler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isWhitelisted\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setYieldTrackers\",\"inputs\":[{\"name\":\"_yieldTrackers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakedBalance\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalStaked\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"whitelistedHandlers\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"yieldTrackers\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// YieldTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use YieldTokenMetaData.ABI instead.
var YieldTokenABI = YieldTokenMetaData.ABI

// YieldToken is an auto generated Go binding around an Ethereum contract.
type YieldToken struct {
	YieldTokenCaller     // Read-only binding to the contract
	YieldTokenTransactor // Write-only binding to the contract
	YieldTokenFilterer   // Log filterer for contract events
}

// YieldTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type YieldTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YieldTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YieldTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YieldTokenSession struct {
	Contract     *YieldToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YieldTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YieldTokenCallerSession struct {
	Contract *YieldTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// YieldTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YieldTokenTransactorSession struct {
	Contract     *YieldTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// YieldTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type YieldTokenRaw struct {
	Contract *YieldToken // Generic contract binding to access the raw methods on
}

// YieldTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YieldTokenCallerRaw struct {
	Contract *YieldTokenCaller // Generic read-only contract binding to access the raw methods on
}

// YieldTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YieldTokenTransactorRaw struct {
	Contract *YieldTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYieldToken creates a new instance of YieldToken, bound to a specific deployed contract.
func NewYieldToken(address common.Address, backend bind.ContractBackend) (*YieldToken, error) {
	contract, err := bindYieldToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YieldToken{YieldTokenCaller: YieldTokenCaller{contract: contract}, YieldTokenTransactor: YieldTokenTransactor{contract: contract}, YieldTokenFilterer: YieldTokenFilterer{contract: contract}}, nil
}

// NewYieldTokenCaller creates a new read-only instance of YieldToken, bound to a specific deployed contract.
func NewYieldTokenCaller(address common.Address, caller bind.ContractCaller) (*YieldTokenCaller, error) {
	contract, err := bindYieldToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YieldTokenCaller{contract: contract}, nil
}

// NewYieldTokenTransactor creates a new write-only instance of YieldToken, bound to a specific deployed contract.
func NewYieldTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*YieldTokenTransactor, error) {
	contract, err := bindYieldToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YieldTokenTransactor{contract: contract}, nil
}

// NewYieldTokenFilterer creates a new log filterer instance of YieldToken, bound to a specific deployed contract.
func NewYieldTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*YieldTokenFilterer, error) {
	contract, err := bindYieldToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YieldTokenFilterer{contract: contract}, nil
}

// bindYieldToken binds a generic wrapper to an already deployed contract.
func bindYieldToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YieldTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YieldToken *YieldTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YieldToken.Contract.YieldTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YieldToken *YieldTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YieldToken.Contract.YieldTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YieldToken *YieldTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YieldToken.Contract.YieldTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YieldToken *YieldTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YieldToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YieldToken *YieldTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YieldToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YieldToken *YieldTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YieldToken.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_YieldToken *YieldTokenCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_YieldToken *YieldTokenSession) Admins(arg0 common.Address) (bool, error) {
	return _YieldToken.Contract.Admins(&_YieldToken.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_YieldToken *YieldTokenCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _YieldToken.Contract.Admins(&_YieldToken.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_YieldToken *YieldTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_YieldToken *YieldTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _YieldToken.Contract.Allowance(&_YieldToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_YieldToken *YieldTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _YieldToken.Contract.Allowance(&_YieldToken.CallOpts, _owner, _spender)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_YieldToken *YieldTokenCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_YieldToken *YieldTokenSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YieldToken.Contract.Allowances(&_YieldToken.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_YieldToken *YieldTokenCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YieldToken.Contract.Allowances(&_YieldToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_YieldToken *YieldTokenCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_YieldToken *YieldTokenSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _YieldToken.Contract.BalanceOf(&_YieldToken.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_YieldToken *YieldTokenCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _YieldToken.Contract.BalanceOf(&_YieldToken.CallOpts, _account)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_YieldToken *YieldTokenCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_YieldToken *YieldTokenSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _YieldToken.Contract.Balances(&_YieldToken.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_YieldToken *YieldTokenCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _YieldToken.Contract.Balances(&_YieldToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YieldToken *YieldTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YieldToken *YieldTokenSession) Decimals() (uint8, error) {
	return _YieldToken.Contract.Decimals(&_YieldToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YieldToken *YieldTokenCallerSession) Decimals() (uint8, error) {
	return _YieldToken.Contract.Decimals(&_YieldToken.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_YieldToken *YieldTokenCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_YieldToken *YieldTokenSession) Gov() (common.Address, error) {
	return _YieldToken.Contract.Gov(&_YieldToken.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_YieldToken *YieldTokenCallerSession) Gov() (common.Address, error) {
	return _YieldToken.Contract.Gov(&_YieldToken.CallOpts)
}

// InWhitelistMode is a free data retrieval call binding the contract method 0x293d6987.
//
// Solidity: function inWhitelistMode() view returns(bool)
func (_YieldToken *YieldTokenCaller) InWhitelistMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "inWhitelistMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InWhitelistMode is a free data retrieval call binding the contract method 0x293d6987.
//
// Solidity: function inWhitelistMode() view returns(bool)
func (_YieldToken *YieldTokenSession) InWhitelistMode() (bool, error) {
	return _YieldToken.Contract.InWhitelistMode(&_YieldToken.CallOpts)
}

// InWhitelistMode is a free data retrieval call binding the contract method 0x293d6987.
//
// Solidity: function inWhitelistMode() view returns(bool)
func (_YieldToken *YieldTokenCallerSession) InWhitelistMode() (bool, error) {
	return _YieldToken.Contract.InWhitelistMode(&_YieldToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YieldToken *YieldTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YieldToken *YieldTokenSession) Name() (string, error) {
	return _YieldToken.Contract.Name(&_YieldToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YieldToken *YieldTokenCallerSession) Name() (string, error) {
	return _YieldToken.Contract.Name(&_YieldToken.CallOpts)
}

// NonStakingAccounts is a free data retrieval call binding the contract method 0x9554381a.
//
// Solidity: function nonStakingAccounts(address ) view returns(bool)
func (_YieldToken *YieldTokenCaller) NonStakingAccounts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "nonStakingAccounts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NonStakingAccounts is a free data retrieval call binding the contract method 0x9554381a.
//
// Solidity: function nonStakingAccounts(address ) view returns(bool)
func (_YieldToken *YieldTokenSession) NonStakingAccounts(arg0 common.Address) (bool, error) {
	return _YieldToken.Contract.NonStakingAccounts(&_YieldToken.CallOpts, arg0)
}

// NonStakingAccounts is a free data retrieval call binding the contract method 0x9554381a.
//
// Solidity: function nonStakingAccounts(address ) view returns(bool)
func (_YieldToken *YieldTokenCallerSession) NonStakingAccounts(arg0 common.Address) (bool, error) {
	return _YieldToken.Contract.NonStakingAccounts(&_YieldToken.CallOpts, arg0)
}

// NonStakingSupply is a free data retrieval call binding the contract method 0xc93be636.
//
// Solidity: function nonStakingSupply() view returns(uint256)
func (_YieldToken *YieldTokenCaller) NonStakingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "nonStakingSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonStakingSupply is a free data retrieval call binding the contract method 0xc93be636.
//
// Solidity: function nonStakingSupply() view returns(uint256)
func (_YieldToken *YieldTokenSession) NonStakingSupply() (*big.Int, error) {
	return _YieldToken.Contract.NonStakingSupply(&_YieldToken.CallOpts)
}

// NonStakingSupply is a free data retrieval call binding the contract method 0xc93be636.
//
// Solidity: function nonStakingSupply() view returns(uint256)
func (_YieldToken *YieldTokenCallerSession) NonStakingSupply() (*big.Int, error) {
	return _YieldToken.Contract.NonStakingSupply(&_YieldToken.CallOpts)
}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_YieldToken *YieldTokenCaller) StakedBalance(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "stakedBalance", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_YieldToken *YieldTokenSession) StakedBalance(_account common.Address) (*big.Int, error) {
	return _YieldToken.Contract.StakedBalance(&_YieldToken.CallOpts, _account)
}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_YieldToken *YieldTokenCallerSession) StakedBalance(_account common.Address) (*big.Int, error) {
	return _YieldToken.Contract.StakedBalance(&_YieldToken.CallOpts, _account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YieldToken *YieldTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YieldToken *YieldTokenSession) Symbol() (string, error) {
	return _YieldToken.Contract.Symbol(&_YieldToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YieldToken *YieldTokenCallerSession) Symbol() (string, error) {
	return _YieldToken.Contract.Symbol(&_YieldToken.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_YieldToken *YieldTokenCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_YieldToken *YieldTokenSession) TotalStaked() (*big.Int, error) {
	return _YieldToken.Contract.TotalStaked(&_YieldToken.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_YieldToken *YieldTokenCallerSession) TotalStaked() (*big.Int, error) {
	return _YieldToken.Contract.TotalStaked(&_YieldToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YieldToken *YieldTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YieldToken *YieldTokenSession) TotalSupply() (*big.Int, error) {
	return _YieldToken.Contract.TotalSupply(&_YieldToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YieldToken *YieldTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _YieldToken.Contract.TotalSupply(&_YieldToken.CallOpts)
}

// WhitelistedHandlers is a free data retrieval call binding the contract method 0x36300051.
//
// Solidity: function whitelistedHandlers(address ) view returns(bool)
func (_YieldToken *YieldTokenCaller) WhitelistedHandlers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "whitelistedHandlers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedHandlers is a free data retrieval call binding the contract method 0x36300051.
//
// Solidity: function whitelistedHandlers(address ) view returns(bool)
func (_YieldToken *YieldTokenSession) WhitelistedHandlers(arg0 common.Address) (bool, error) {
	return _YieldToken.Contract.WhitelistedHandlers(&_YieldToken.CallOpts, arg0)
}

// WhitelistedHandlers is a free data retrieval call binding the contract method 0x36300051.
//
// Solidity: function whitelistedHandlers(address ) view returns(bool)
func (_YieldToken *YieldTokenCallerSession) WhitelistedHandlers(arg0 common.Address) (bool, error) {
	return _YieldToken.Contract.WhitelistedHandlers(&_YieldToken.CallOpts, arg0)
}

// YieldTrackers is a free data retrieval call binding the contract method 0x52cd38d9.
//
// Solidity: function yieldTrackers(uint256 ) view returns(address)
func (_YieldToken *YieldTokenCaller) YieldTrackers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YieldToken.contract.Call(opts, &out, "yieldTrackers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YieldTrackers is a free data retrieval call binding the contract method 0x52cd38d9.
//
// Solidity: function yieldTrackers(uint256 ) view returns(address)
func (_YieldToken *YieldTokenSession) YieldTrackers(arg0 *big.Int) (common.Address, error) {
	return _YieldToken.Contract.YieldTrackers(&_YieldToken.CallOpts, arg0)
}

// YieldTrackers is a free data retrieval call binding the contract method 0x52cd38d9.
//
// Solidity: function yieldTrackers(uint256 ) view returns(address)
func (_YieldToken *YieldTokenCallerSession) YieldTrackers(arg0 *big.Int) (common.Address, error) {
	return _YieldToken.Contract.YieldTrackers(&_YieldToken.CallOpts, arg0)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_YieldToken *YieldTokenTransactor) AddAdmin(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "addAdmin", _account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_YieldToken *YieldTokenSession) AddAdmin(_account common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.AddAdmin(&_YieldToken.TransactOpts, _account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_YieldToken *YieldTokenTransactorSession) AddAdmin(_account common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.AddAdmin(&_YieldToken.TransactOpts, _account)
}

// AddNonStakingAccount is a paid mutator transaction binding the contract method 0x62289077.
//
// Solidity: function addNonStakingAccount(address _account) returns()
func (_YieldToken *YieldTokenTransactor) AddNonStakingAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "addNonStakingAccount", _account)
}

// AddNonStakingAccount is a paid mutator transaction binding the contract method 0x62289077.
//
// Solidity: function addNonStakingAccount(address _account) returns()
func (_YieldToken *YieldTokenSession) AddNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.AddNonStakingAccount(&_YieldToken.TransactOpts, _account)
}

// AddNonStakingAccount is a paid mutator transaction binding the contract method 0x62289077.
//
// Solidity: function addNonStakingAccount(address _account) returns()
func (_YieldToken *YieldTokenTransactorSession) AddNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.AddNonStakingAccount(&_YieldToken.TransactOpts, _account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_YieldToken *YieldTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_YieldToken *YieldTokenSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.Approve(&_YieldToken.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_YieldToken *YieldTokenTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.Approve(&_YieldToken.TransactOpts, _spender, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns()
func (_YieldToken *YieldTokenTransactor) Claim(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "claim", _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns()
func (_YieldToken *YieldTokenSession) Claim(_receiver common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.Claim(&_YieldToken.TransactOpts, _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns()
func (_YieldToken *YieldTokenTransactorSession) Claim(_receiver common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.Claim(&_YieldToken.TransactOpts, _receiver)
}

// RecoverClaim is a paid mutator transaction binding the contract method 0x996f11ee.
//
// Solidity: function recoverClaim(address _account, address _receiver) returns()
func (_YieldToken *YieldTokenTransactor) RecoverClaim(opts *bind.TransactOpts, _account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "recoverClaim", _account, _receiver)
}

// RecoverClaim is a paid mutator transaction binding the contract method 0x996f11ee.
//
// Solidity: function recoverClaim(address _account, address _receiver) returns()
func (_YieldToken *YieldTokenSession) RecoverClaim(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.RecoverClaim(&_YieldToken.TransactOpts, _account, _receiver)
}

// RecoverClaim is a paid mutator transaction binding the contract method 0x996f11ee.
//
// Solidity: function recoverClaim(address _account, address _receiver) returns()
func (_YieldToken *YieldTokenTransactorSession) RecoverClaim(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.RecoverClaim(&_YieldToken.TransactOpts, _account, _receiver)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_YieldToken *YieldTokenTransactor) RemoveAdmin(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "removeAdmin", _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_YieldToken *YieldTokenSession) RemoveAdmin(_account common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.RemoveAdmin(&_YieldToken.TransactOpts, _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_YieldToken *YieldTokenTransactorSession) RemoveAdmin(_account common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.RemoveAdmin(&_YieldToken.TransactOpts, _account)
}

// RemoveNonStakingAccount is a paid mutator transaction binding the contract method 0xfb30d916.
//
// Solidity: function removeNonStakingAccount(address _account) returns()
func (_YieldToken *YieldTokenTransactor) RemoveNonStakingAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "removeNonStakingAccount", _account)
}

// RemoveNonStakingAccount is a paid mutator transaction binding the contract method 0xfb30d916.
//
// Solidity: function removeNonStakingAccount(address _account) returns()
func (_YieldToken *YieldTokenSession) RemoveNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.RemoveNonStakingAccount(&_YieldToken.TransactOpts, _account)
}

// RemoveNonStakingAccount is a paid mutator transaction binding the contract method 0xfb30d916.
//
// Solidity: function removeNonStakingAccount(address _account) returns()
func (_YieldToken *YieldTokenTransactorSession) RemoveNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.RemoveNonStakingAccount(&_YieldToken.TransactOpts, _account)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_YieldToken *YieldTokenTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_YieldToken *YieldTokenSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.SetGov(&_YieldToken.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_YieldToken *YieldTokenTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.SetGov(&_YieldToken.TransactOpts, _gov)
}

// SetInWhitelistMode is a paid mutator transaction binding the contract method 0x4cb5bbe3.
//
// Solidity: function setInWhitelistMode(bool _inWhitelistMode) returns()
func (_YieldToken *YieldTokenTransactor) SetInWhitelistMode(opts *bind.TransactOpts, _inWhitelistMode bool) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "setInWhitelistMode", _inWhitelistMode)
}

// SetInWhitelistMode is a paid mutator transaction binding the contract method 0x4cb5bbe3.
//
// Solidity: function setInWhitelistMode(bool _inWhitelistMode) returns()
func (_YieldToken *YieldTokenSession) SetInWhitelistMode(_inWhitelistMode bool) (*types.Transaction, error) {
	return _YieldToken.Contract.SetInWhitelistMode(&_YieldToken.TransactOpts, _inWhitelistMode)
}

// SetInWhitelistMode is a paid mutator transaction binding the contract method 0x4cb5bbe3.
//
// Solidity: function setInWhitelistMode(bool _inWhitelistMode) returns()
func (_YieldToken *YieldTokenTransactorSession) SetInWhitelistMode(_inWhitelistMode bool) (*types.Transaction, error) {
	return _YieldToken.Contract.SetInWhitelistMode(&_YieldToken.TransactOpts, _inWhitelistMode)
}

// SetInfo is a paid mutator transaction binding the contract method 0xa923fc40.
//
// Solidity: function setInfo(string _name, string _symbol) returns()
func (_YieldToken *YieldTokenTransactor) SetInfo(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "setInfo", _name, _symbol)
}

// SetInfo is a paid mutator transaction binding the contract method 0xa923fc40.
//
// Solidity: function setInfo(string _name, string _symbol) returns()
func (_YieldToken *YieldTokenSession) SetInfo(_name string, _symbol string) (*types.Transaction, error) {
	return _YieldToken.Contract.SetInfo(&_YieldToken.TransactOpts, _name, _symbol)
}

// SetInfo is a paid mutator transaction binding the contract method 0xa923fc40.
//
// Solidity: function setInfo(string _name, string _symbol) returns()
func (_YieldToken *YieldTokenTransactorSession) SetInfo(_name string, _symbol string) (*types.Transaction, error) {
	return _YieldToken.Contract.SetInfo(&_YieldToken.TransactOpts, _name, _symbol)
}

// SetWhitelistedHandler is a paid mutator transaction binding the contract method 0xd92fc87e.
//
// Solidity: function setWhitelistedHandler(address _handler, bool _isWhitelisted) returns()
func (_YieldToken *YieldTokenTransactor) SetWhitelistedHandler(opts *bind.TransactOpts, _handler common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "setWhitelistedHandler", _handler, _isWhitelisted)
}

// SetWhitelistedHandler is a paid mutator transaction binding the contract method 0xd92fc87e.
//
// Solidity: function setWhitelistedHandler(address _handler, bool _isWhitelisted) returns()
func (_YieldToken *YieldTokenSession) SetWhitelistedHandler(_handler common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _YieldToken.Contract.SetWhitelistedHandler(&_YieldToken.TransactOpts, _handler, _isWhitelisted)
}

// SetWhitelistedHandler is a paid mutator transaction binding the contract method 0xd92fc87e.
//
// Solidity: function setWhitelistedHandler(address _handler, bool _isWhitelisted) returns()
func (_YieldToken *YieldTokenTransactorSession) SetWhitelistedHandler(_handler common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _YieldToken.Contract.SetWhitelistedHandler(&_YieldToken.TransactOpts, _handler, _isWhitelisted)
}

// SetYieldTrackers is a paid mutator transaction binding the contract method 0x276eab4e.
//
// Solidity: function setYieldTrackers(address[] _yieldTrackers) returns()
func (_YieldToken *YieldTokenTransactor) SetYieldTrackers(opts *bind.TransactOpts, _yieldTrackers []common.Address) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "setYieldTrackers", _yieldTrackers)
}

// SetYieldTrackers is a paid mutator transaction binding the contract method 0x276eab4e.
//
// Solidity: function setYieldTrackers(address[] _yieldTrackers) returns()
func (_YieldToken *YieldTokenSession) SetYieldTrackers(_yieldTrackers []common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.SetYieldTrackers(&_YieldToken.TransactOpts, _yieldTrackers)
}

// SetYieldTrackers is a paid mutator transaction binding the contract method 0x276eab4e.
//
// Solidity: function setYieldTrackers(address[] _yieldTrackers) returns()
func (_YieldToken *YieldTokenTransactorSession) SetYieldTrackers(_yieldTrackers []common.Address) (*types.Transaction, error) {
	return _YieldToken.Contract.SetYieldTrackers(&_YieldToken.TransactOpts, _yieldTrackers)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_YieldToken *YieldTokenTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_YieldToken *YieldTokenSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.Transfer(&_YieldToken.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_YieldToken *YieldTokenTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.Transfer(&_YieldToken.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_YieldToken *YieldTokenTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_YieldToken *YieldTokenSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.TransferFrom(&_YieldToken.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_YieldToken *YieldTokenTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.TransferFrom(&_YieldToken.TransactOpts, _sender, _recipient, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_YieldToken *YieldTokenTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldToken.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_YieldToken *YieldTokenSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.WithdrawToken(&_YieldToken.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_YieldToken *YieldTokenTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldToken.Contract.WithdrawToken(&_YieldToken.TransactOpts, _token, _account, _amount)
}

// YieldTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the YieldToken contract.
type YieldTokenApprovalIterator struct {
	Event *YieldTokenApproval // Event containing the contract specifics and raw log

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
func (it *YieldTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YieldTokenApproval)
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
		it.Event = new(YieldTokenApproval)
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
func (it *YieldTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YieldTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YieldTokenApproval represents a Approval event raised by the YieldToken contract.
type YieldTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YieldToken *YieldTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*YieldTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YieldToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &YieldTokenApprovalIterator{contract: _YieldToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YieldToken *YieldTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *YieldTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YieldToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YieldTokenApproval)
				if err := _YieldToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YieldToken *YieldTokenFilterer) ParseApproval(log types.Log) (*YieldTokenApproval, error) {
	event := new(YieldTokenApproval)
	if err := _YieldToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YieldTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the YieldToken contract.
type YieldTokenTransferIterator struct {
	Event *YieldTokenTransfer // Event containing the contract specifics and raw log

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
func (it *YieldTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YieldTokenTransfer)
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
		it.Event = new(YieldTokenTransfer)
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
func (it *YieldTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YieldTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YieldTokenTransfer represents a Transfer event raised by the YieldToken contract.
type YieldTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YieldToken *YieldTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*YieldTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YieldToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &YieldTokenTransferIterator{contract: _YieldToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YieldToken *YieldTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *YieldTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YieldToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YieldTokenTransfer)
				if err := _YieldToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YieldToken *YieldTokenFilterer) ParseTransfer(log types.Log) (*YieldTokenTransfer, error) {
	event := new(YieldTokenTransfer)
	if err := _YieldToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

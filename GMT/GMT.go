// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GMT

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

// GMTMetaData contains all meta data concerning the GMT contract.
var GMTMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_initialSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addAdmin\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addBlockedRecipient\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addMsgSender\",\"inputs\":[{\"name\":\"_msgSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"admins\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowedMsgSenders\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"beginMigration\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blockedRecipients\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"endMigration\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasActiveMigration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrationTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeAdmin\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeBlockedRecipient\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeMsgSender\",\"inputs\":[{\"name\":\"_msgSender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNextMigrationTime\",\"inputs\":[{\"name\":\"_migrationTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// GMTABI is the input ABI used to generate the binding from.
// Deprecated: Use GMTMetaData.ABI instead.
var GMTABI = GMTMetaData.ABI

// GMT is an auto generated Go binding around an Ethereum contract.
type GMT struct {
	GMTCaller     // Read-only binding to the contract
	GMTTransactor // Write-only binding to the contract
	GMTFilterer   // Log filterer for contract events
}

// GMTCaller is an auto generated read-only Go binding around an Ethereum contract.
type GMTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GMTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GMTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GMTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GMTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GMTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GMTSession struct {
	Contract     *GMT              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GMTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GMTCallerSession struct {
	Contract *GMTCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// GMTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GMTTransactorSession struct {
	Contract     *GMTTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GMTRaw is an auto generated low-level Go binding around an Ethereum contract.
type GMTRaw struct {
	Contract *GMT // Generic contract binding to access the raw methods on
}

// GMTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GMTCallerRaw struct {
	Contract *GMTCaller // Generic read-only contract binding to access the raw methods on
}

// GMTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GMTTransactorRaw struct {
	Contract *GMTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGMT creates a new instance of GMT, bound to a specific deployed contract.
func NewGMT(address common.Address, backend bind.ContractBackend) (*GMT, error) {
	contract, err := bindGMT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GMT{GMTCaller: GMTCaller{contract: contract}, GMTTransactor: GMTTransactor{contract: contract}, GMTFilterer: GMTFilterer{contract: contract}}, nil
}

// NewGMTCaller creates a new read-only instance of GMT, bound to a specific deployed contract.
func NewGMTCaller(address common.Address, caller bind.ContractCaller) (*GMTCaller, error) {
	contract, err := bindGMT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GMTCaller{contract: contract}, nil
}

// NewGMTTransactor creates a new write-only instance of GMT, bound to a specific deployed contract.
func NewGMTTransactor(address common.Address, transactor bind.ContractTransactor) (*GMTTransactor, error) {
	contract, err := bindGMT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GMTTransactor{contract: contract}, nil
}

// NewGMTFilterer creates a new log filterer instance of GMT, bound to a specific deployed contract.
func NewGMTFilterer(address common.Address, filterer bind.ContractFilterer) (*GMTFilterer, error) {
	contract, err := bindGMT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GMTFilterer{contract: contract}, nil
}

// bindGMT binds a generic wrapper to an already deployed contract.
func bindGMT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GMTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GMT *GMTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GMT.Contract.GMTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GMT *GMTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GMT.Contract.GMTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GMT *GMTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GMT.Contract.GMTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GMT *GMTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GMT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GMT *GMTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GMT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GMT *GMTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GMT.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_GMT *GMTCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GMT.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_GMT *GMTSession) Admins(arg0 common.Address) (bool, error) {
	return _GMT.Contract.Admins(&_GMT.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_GMT *GMTCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _GMT.Contract.Admins(&_GMT.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_GMT *GMTCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GMT.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_GMT *GMTSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _GMT.Contract.Allowance(&_GMT.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_GMT *GMTCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _GMT.Contract.Allowance(&_GMT.CallOpts, _owner, _spender)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_GMT *GMTCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GMT.contract.Call(opts, &out, "allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_GMT *GMTSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GMT.Contract.Allowances(&_GMT.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_GMT *GMTCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GMT.Contract.Allowances(&_GMT.CallOpts, arg0, arg1)
}

// AllowedMsgSenders is a free data retrieval call binding the contract method 0x01d060b0.
//
// Solidity: function allowedMsgSenders(address ) view returns(bool)
func (_GMT *GMTCaller) AllowedMsgSenders(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GMT.contract.Call(opts, &out, "allowedMsgSenders", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedMsgSenders is a free data retrieval call binding the contract method 0x01d060b0.
//
// Solidity: function allowedMsgSenders(address ) view returns(bool)
func (_GMT *GMTSession) AllowedMsgSenders(arg0 common.Address) (bool, error) {
	return _GMT.Contract.AllowedMsgSenders(&_GMT.CallOpts, arg0)
}

// AllowedMsgSenders is a free data retrieval call binding the contract method 0x01d060b0.
//
// Solidity: function allowedMsgSenders(address ) view returns(bool)
func (_GMT *GMTCallerSession) AllowedMsgSenders(arg0 common.Address) (bool, error) {
	return _GMT.Contract.AllowedMsgSenders(&_GMT.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_GMT *GMTCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GMT.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_GMT *GMTSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _GMT.Contract.BalanceOf(&_GMT.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_GMT *GMTCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _GMT.Contract.BalanceOf(&_GMT.CallOpts, _account)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_GMT *GMTCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GMT.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_GMT *GMTSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _GMT.Contract.Balances(&_GMT.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_GMT *GMTCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _GMT.Contract.Balances(&_GMT.CallOpts, arg0)
}

// BlockedRecipients is a free data retrieval call binding the contract method 0x205b10a0.
//
// Solidity: function blockedRecipients(address ) view returns(bool)
func (_GMT *GMTCaller) BlockedRecipients(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GMT.contract.Call(opts, &out, "blockedRecipients", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlockedRecipients is a free data retrieval call binding the contract method 0x205b10a0.
//
// Solidity: function blockedRecipients(address ) view returns(bool)
func (_GMT *GMTSession) BlockedRecipients(arg0 common.Address) (bool, error) {
	return _GMT.Contract.BlockedRecipients(&_GMT.CallOpts, arg0)
}

// BlockedRecipients is a free data retrieval call binding the contract method 0x205b10a0.
//
// Solidity: function blockedRecipients(address ) view returns(bool)
func (_GMT *GMTCallerSession) BlockedRecipients(arg0 common.Address) (bool, error) {
	return _GMT.Contract.BlockedRecipients(&_GMT.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GMT *GMTCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _GMT.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GMT *GMTSession) Decimals() (uint8, error) {
	return _GMT.Contract.Decimals(&_GMT.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GMT *GMTCallerSession) Decimals() (uint8, error) {
	return _GMT.Contract.Decimals(&_GMT.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_GMT *GMTCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GMT.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_GMT *GMTSession) Gov() (common.Address, error) {
	return _GMT.Contract.Gov(&_GMT.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_GMT *GMTCallerSession) Gov() (common.Address, error) {
	return _GMT.Contract.Gov(&_GMT.CallOpts)
}

// HasActiveMigration is a free data retrieval call binding the contract method 0x8fb998e1.
//
// Solidity: function hasActiveMigration() view returns(bool)
func (_GMT *GMTCaller) HasActiveMigration(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GMT.contract.Call(opts, &out, "hasActiveMigration")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasActiveMigration is a free data retrieval call binding the contract method 0x8fb998e1.
//
// Solidity: function hasActiveMigration() view returns(bool)
func (_GMT *GMTSession) HasActiveMigration() (bool, error) {
	return _GMT.Contract.HasActiveMigration(&_GMT.CallOpts)
}

// HasActiveMigration is a free data retrieval call binding the contract method 0x8fb998e1.
//
// Solidity: function hasActiveMigration() view returns(bool)
func (_GMT *GMTCallerSession) HasActiveMigration() (bool, error) {
	return _GMT.Contract.HasActiveMigration(&_GMT.CallOpts)
}

// MigrationTime is a free data retrieval call binding the contract method 0xff61a51c.
//
// Solidity: function migrationTime() view returns(uint256)
func (_GMT *GMTCaller) MigrationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GMT.contract.Call(opts, &out, "migrationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MigrationTime is a free data retrieval call binding the contract method 0xff61a51c.
//
// Solidity: function migrationTime() view returns(uint256)
func (_GMT *GMTSession) MigrationTime() (*big.Int, error) {
	return _GMT.Contract.MigrationTime(&_GMT.CallOpts)
}

// MigrationTime is a free data retrieval call binding the contract method 0xff61a51c.
//
// Solidity: function migrationTime() view returns(uint256)
func (_GMT *GMTCallerSession) MigrationTime() (*big.Int, error) {
	return _GMT.Contract.MigrationTime(&_GMT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GMT *GMTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GMT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GMT *GMTSession) Name() (string, error) {
	return _GMT.Contract.Name(&_GMT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GMT *GMTCallerSession) Name() (string, error) {
	return _GMT.Contract.Name(&_GMT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GMT *GMTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GMT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GMT *GMTSession) Symbol() (string, error) {
	return _GMT.Contract.Symbol(&_GMT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GMT *GMTCallerSession) Symbol() (string, error) {
	return _GMT.Contract.Symbol(&_GMT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GMT *GMTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GMT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GMT *GMTSession) TotalSupply() (*big.Int, error) {
	return _GMT.Contract.TotalSupply(&_GMT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GMT *GMTCallerSession) TotalSupply() (*big.Int, error) {
	return _GMT.Contract.TotalSupply(&_GMT.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_GMT *GMTTransactor) AddAdmin(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _GMT.contract.Transact(opts, "addAdmin", _account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_GMT *GMTSession) AddAdmin(_account common.Address) (*types.Transaction, error) {
	return _GMT.Contract.AddAdmin(&_GMT.TransactOpts, _account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_GMT *GMTTransactorSession) AddAdmin(_account common.Address) (*types.Transaction, error) {
	return _GMT.Contract.AddAdmin(&_GMT.TransactOpts, _account)
}

// AddBlockedRecipient is a paid mutator transaction binding the contract method 0xb9952455.
//
// Solidity: function addBlockedRecipient(address _recipient) returns()
func (_GMT *GMTTransactor) AddBlockedRecipient(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _GMT.contract.Transact(opts, "addBlockedRecipient", _recipient)
}

// AddBlockedRecipient is a paid mutator transaction binding the contract method 0xb9952455.
//
// Solidity: function addBlockedRecipient(address _recipient) returns()
func (_GMT *GMTSession) AddBlockedRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _GMT.Contract.AddBlockedRecipient(&_GMT.TransactOpts, _recipient)
}

// AddBlockedRecipient is a paid mutator transaction binding the contract method 0xb9952455.
//
// Solidity: function addBlockedRecipient(address _recipient) returns()
func (_GMT *GMTTransactorSession) AddBlockedRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _GMT.Contract.AddBlockedRecipient(&_GMT.TransactOpts, _recipient)
}

// AddMsgSender is a paid mutator transaction binding the contract method 0x08ebd96a.
//
// Solidity: function addMsgSender(address _msgSender) returns()
func (_GMT *GMTTransactor) AddMsgSender(opts *bind.TransactOpts, _msgSender common.Address) (*types.Transaction, error) {
	return _GMT.contract.Transact(opts, "addMsgSender", _msgSender)
}

// AddMsgSender is a paid mutator transaction binding the contract method 0x08ebd96a.
//
// Solidity: function addMsgSender(address _msgSender) returns()
func (_GMT *GMTSession) AddMsgSender(_msgSender common.Address) (*types.Transaction, error) {
	return _GMT.Contract.AddMsgSender(&_GMT.TransactOpts, _msgSender)
}

// AddMsgSender is a paid mutator transaction binding the contract method 0x08ebd96a.
//
// Solidity: function addMsgSender(address _msgSender) returns()
func (_GMT *GMTTransactorSession) AddMsgSender(_msgSender common.Address) (*types.Transaction, error) {
	return _GMT.Contract.AddMsgSender(&_GMT.TransactOpts, _msgSender)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_GMT *GMTTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMT.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_GMT *GMTSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMT.Contract.Approve(&_GMT.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_GMT *GMTTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMT.Contract.Approve(&_GMT.TransactOpts, _spender, _amount)
}

// BeginMigration is a paid mutator transaction binding the contract method 0xfe0194f2.
//
// Solidity: function beginMigration() returns()
func (_GMT *GMTTransactor) BeginMigration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GMT.contract.Transact(opts, "beginMigration")
}

// BeginMigration is a paid mutator transaction binding the contract method 0xfe0194f2.
//
// Solidity: function beginMigration() returns()
func (_GMT *GMTSession) BeginMigration() (*types.Transaction, error) {
	return _GMT.Contract.BeginMigration(&_GMT.TransactOpts)
}

// BeginMigration is a paid mutator transaction binding the contract method 0xfe0194f2.
//
// Solidity: function beginMigration() returns()
func (_GMT *GMTTransactorSession) BeginMigration() (*types.Transaction, error) {
	return _GMT.Contract.BeginMigration(&_GMT.TransactOpts)
}

// EndMigration is a paid mutator transaction binding the contract method 0x6c525d04.
//
// Solidity: function endMigration() returns()
func (_GMT *GMTTransactor) EndMigration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GMT.contract.Transact(opts, "endMigration")
}

// EndMigration is a paid mutator transaction binding the contract method 0x6c525d04.
//
// Solidity: function endMigration() returns()
func (_GMT *GMTSession) EndMigration() (*types.Transaction, error) {
	return _GMT.Contract.EndMigration(&_GMT.TransactOpts)
}

// EndMigration is a paid mutator transaction binding the contract method 0x6c525d04.
//
// Solidity: function endMigration() returns()
func (_GMT *GMTTransactorSession) EndMigration() (*types.Transaction, error) {
	return _GMT.Contract.EndMigration(&_GMT.TransactOpts)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_GMT *GMTTransactor) RemoveAdmin(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _GMT.contract.Transact(opts, "removeAdmin", _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_GMT *GMTSession) RemoveAdmin(_account common.Address) (*types.Transaction, error) {
	return _GMT.Contract.RemoveAdmin(&_GMT.TransactOpts, _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_GMT *GMTTransactorSession) RemoveAdmin(_account common.Address) (*types.Transaction, error) {
	return _GMT.Contract.RemoveAdmin(&_GMT.TransactOpts, _account)
}

// RemoveBlockedRecipient is a paid mutator transaction binding the contract method 0x5ec080fc.
//
// Solidity: function removeBlockedRecipient(address _recipient) returns()
func (_GMT *GMTTransactor) RemoveBlockedRecipient(opts *bind.TransactOpts, _recipient common.Address) (*types.Transaction, error) {
	return _GMT.contract.Transact(opts, "removeBlockedRecipient", _recipient)
}

// RemoveBlockedRecipient is a paid mutator transaction binding the contract method 0x5ec080fc.
//
// Solidity: function removeBlockedRecipient(address _recipient) returns()
func (_GMT *GMTSession) RemoveBlockedRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _GMT.Contract.RemoveBlockedRecipient(&_GMT.TransactOpts, _recipient)
}

// RemoveBlockedRecipient is a paid mutator transaction binding the contract method 0x5ec080fc.
//
// Solidity: function removeBlockedRecipient(address _recipient) returns()
func (_GMT *GMTTransactorSession) RemoveBlockedRecipient(_recipient common.Address) (*types.Transaction, error) {
	return _GMT.Contract.RemoveBlockedRecipient(&_GMT.TransactOpts, _recipient)
}

// RemoveMsgSender is a paid mutator transaction binding the contract method 0x4d9567ee.
//
// Solidity: function removeMsgSender(address _msgSender) returns()
func (_GMT *GMTTransactor) RemoveMsgSender(opts *bind.TransactOpts, _msgSender common.Address) (*types.Transaction, error) {
	return _GMT.contract.Transact(opts, "removeMsgSender", _msgSender)
}

// RemoveMsgSender is a paid mutator transaction binding the contract method 0x4d9567ee.
//
// Solidity: function removeMsgSender(address _msgSender) returns()
func (_GMT *GMTSession) RemoveMsgSender(_msgSender common.Address) (*types.Transaction, error) {
	return _GMT.Contract.RemoveMsgSender(&_GMT.TransactOpts, _msgSender)
}

// RemoveMsgSender is a paid mutator transaction binding the contract method 0x4d9567ee.
//
// Solidity: function removeMsgSender(address _msgSender) returns()
func (_GMT *GMTTransactorSession) RemoveMsgSender(_msgSender common.Address) (*types.Transaction, error) {
	return _GMT.Contract.RemoveMsgSender(&_GMT.TransactOpts, _msgSender)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_GMT *GMTTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _GMT.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_GMT *GMTSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _GMT.Contract.SetGov(&_GMT.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_GMT *GMTTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _GMT.Contract.SetGov(&_GMT.TransactOpts, _gov)
}

// SetNextMigrationTime is a paid mutator transaction binding the contract method 0xa92c6fae.
//
// Solidity: function setNextMigrationTime(uint256 _migrationTime) returns()
func (_GMT *GMTTransactor) SetNextMigrationTime(opts *bind.TransactOpts, _migrationTime *big.Int) (*types.Transaction, error) {
	return _GMT.contract.Transact(opts, "setNextMigrationTime", _migrationTime)
}

// SetNextMigrationTime is a paid mutator transaction binding the contract method 0xa92c6fae.
//
// Solidity: function setNextMigrationTime(uint256 _migrationTime) returns()
func (_GMT *GMTSession) SetNextMigrationTime(_migrationTime *big.Int) (*types.Transaction, error) {
	return _GMT.Contract.SetNextMigrationTime(&_GMT.TransactOpts, _migrationTime)
}

// SetNextMigrationTime is a paid mutator transaction binding the contract method 0xa92c6fae.
//
// Solidity: function setNextMigrationTime(uint256 _migrationTime) returns()
func (_GMT *GMTTransactorSession) SetNextMigrationTime(_migrationTime *big.Int) (*types.Transaction, error) {
	return _GMT.Contract.SetNextMigrationTime(&_GMT.TransactOpts, _migrationTime)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_GMT *GMTTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMT.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_GMT *GMTSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMT.Contract.Transfer(&_GMT.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_GMT *GMTTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMT.Contract.Transfer(&_GMT.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_GMT *GMTTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMT.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_GMT *GMTSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMT.Contract.TransferFrom(&_GMT.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_GMT *GMTTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMT.Contract.TransferFrom(&_GMT.TransactOpts, _sender, _recipient, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_GMT *GMTTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMT.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_GMT *GMTSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMT.Contract.WithdrawToken(&_GMT.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_GMT *GMTTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GMT.Contract.WithdrawToken(&_GMT.TransactOpts, _token, _account, _amount)
}

// GMTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GMT contract.
type GMTApprovalIterator struct {
	Event *GMTApproval // Event containing the contract specifics and raw log

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
func (it *GMTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GMTApproval)
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
		it.Event = new(GMTApproval)
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
func (it *GMTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GMTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GMTApproval represents a Approval event raised by the GMT contract.
type GMTApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GMT *GMTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GMTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GMT.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GMTApprovalIterator{contract: _GMT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GMT *GMTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GMTApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GMT.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GMTApproval)
				if err := _GMT.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_GMT *GMTFilterer) ParseApproval(log types.Log) (*GMTApproval, error) {
	event := new(GMTApproval)
	if err := _GMT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GMTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the GMT contract.
type GMTTransferIterator struct {
	Event *GMTTransfer // Event containing the contract specifics and raw log

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
func (it *GMTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GMTTransfer)
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
		it.Event = new(GMTTransfer)
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
func (it *GMTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GMTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GMTTransfer represents a Transfer event raised by the GMT contract.
type GMTTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GMT *GMTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GMTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GMT.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GMTTransferIterator{contract: _GMT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GMT *GMTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GMTTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GMT.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GMTTransfer)
				if err := _GMT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_GMT *GMTFilterer) ParseTransfer(log types.Log) (*GMTTransfer, error) {
	event := new(GMTTransfer)
	if err := _GMT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

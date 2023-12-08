// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MintableBaseToken

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

// MintableBaseTokenMetaData contains all meta data concerning the MintableBaseToken contract.
var MintableBaseTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_initialSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addAdmin\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addNonStakingAccount\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"admins\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inPrivateTransferMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isHandler\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isMinter\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonStakingAccounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonStakingSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverClaim\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeAdmin\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeNonStakingAccount\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setHandler\",\"inputs\":[{\"name\":\"_handler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInPrivateTransferMode\",\"inputs\":[{\"name\":\"_inPrivateTransferMode\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInfo\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinter\",\"inputs\":[{\"name\":\"_minter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setYieldTrackers\",\"inputs\":[{\"name\":\"_yieldTrackers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakedBalance\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalStaked\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"yieldTrackers\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// MintableBaseTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use MintableBaseTokenMetaData.ABI instead.
var MintableBaseTokenABI = MintableBaseTokenMetaData.ABI

// MintableBaseToken is an auto generated Go binding around an Ethereum contract.
type MintableBaseToken struct {
	MintableBaseTokenCaller     // Read-only binding to the contract
	MintableBaseTokenTransactor // Write-only binding to the contract
	MintableBaseTokenFilterer   // Log filterer for contract events
}

// MintableBaseTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type MintableBaseTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableBaseTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MintableBaseTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableBaseTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MintableBaseTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintableBaseTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MintableBaseTokenSession struct {
	Contract     *MintableBaseToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MintableBaseTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MintableBaseTokenCallerSession struct {
	Contract *MintableBaseTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// MintableBaseTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MintableBaseTokenTransactorSession struct {
	Contract     *MintableBaseTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// MintableBaseTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type MintableBaseTokenRaw struct {
	Contract *MintableBaseToken // Generic contract binding to access the raw methods on
}

// MintableBaseTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MintableBaseTokenCallerRaw struct {
	Contract *MintableBaseTokenCaller // Generic read-only contract binding to access the raw methods on
}

// MintableBaseTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MintableBaseTokenTransactorRaw struct {
	Contract *MintableBaseTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMintableBaseToken creates a new instance of MintableBaseToken, bound to a specific deployed contract.
func NewMintableBaseToken(address common.Address, backend bind.ContractBackend) (*MintableBaseToken, error) {
	contract, err := bindMintableBaseToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MintableBaseToken{MintableBaseTokenCaller: MintableBaseTokenCaller{contract: contract}, MintableBaseTokenTransactor: MintableBaseTokenTransactor{contract: contract}, MintableBaseTokenFilterer: MintableBaseTokenFilterer{contract: contract}}, nil
}

// NewMintableBaseTokenCaller creates a new read-only instance of MintableBaseToken, bound to a specific deployed contract.
func NewMintableBaseTokenCaller(address common.Address, caller bind.ContractCaller) (*MintableBaseTokenCaller, error) {
	contract, err := bindMintableBaseToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MintableBaseTokenCaller{contract: contract}, nil
}

// NewMintableBaseTokenTransactor creates a new write-only instance of MintableBaseToken, bound to a specific deployed contract.
func NewMintableBaseTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*MintableBaseTokenTransactor, error) {
	contract, err := bindMintableBaseToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MintableBaseTokenTransactor{contract: contract}, nil
}

// NewMintableBaseTokenFilterer creates a new log filterer instance of MintableBaseToken, bound to a specific deployed contract.
func NewMintableBaseTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*MintableBaseTokenFilterer, error) {
	contract, err := bindMintableBaseToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MintableBaseTokenFilterer{contract: contract}, nil
}

// bindMintableBaseToken binds a generic wrapper to an already deployed contract.
func bindMintableBaseToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MintableBaseTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintableBaseToken *MintableBaseTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintableBaseToken.Contract.MintableBaseTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintableBaseToken *MintableBaseTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.MintableBaseTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintableBaseToken *MintableBaseTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.MintableBaseTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MintableBaseToken *MintableBaseTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MintableBaseToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MintableBaseToken *MintableBaseTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MintableBaseToken *MintableBaseTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_MintableBaseToken *MintableBaseTokenCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MintableBaseToken.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_MintableBaseToken *MintableBaseTokenSession) Admins(arg0 common.Address) (bool, error) {
	return _MintableBaseToken.Contract.Admins(&_MintableBaseToken.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_MintableBaseToken *MintableBaseTokenCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _MintableBaseToken.Contract.Admins(&_MintableBaseToken.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MintableBaseToken.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MintableBaseToken.Contract.Allowance(&_MintableBaseToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _MintableBaseToken.Contract.Allowance(&_MintableBaseToken.CallOpts, _owner, _spender)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MintableBaseToken.contract.Call(opts, &out, "allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MintableBaseToken.Contract.Allowances(&_MintableBaseToken.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _MintableBaseToken.Contract.Allowances(&_MintableBaseToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MintableBaseToken.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _MintableBaseToken.Contract.BalanceOf(&_MintableBaseToken.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _MintableBaseToken.Contract.BalanceOf(&_MintableBaseToken.CallOpts, _account)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MintableBaseToken.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _MintableBaseToken.Contract.Balances(&_MintableBaseToken.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _MintableBaseToken.Contract.Balances(&_MintableBaseToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MintableBaseToken *MintableBaseTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MintableBaseToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MintableBaseToken *MintableBaseTokenSession) Decimals() (uint8, error) {
	return _MintableBaseToken.Contract.Decimals(&_MintableBaseToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MintableBaseToken *MintableBaseTokenCallerSession) Decimals() (uint8, error) {
	return _MintableBaseToken.Contract.Decimals(&_MintableBaseToken.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_MintableBaseToken *MintableBaseTokenCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MintableBaseToken.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_MintableBaseToken *MintableBaseTokenSession) Gov() (common.Address, error) {
	return _MintableBaseToken.Contract.Gov(&_MintableBaseToken.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_MintableBaseToken *MintableBaseTokenCallerSession) Gov() (common.Address, error) {
	return _MintableBaseToken.Contract.Gov(&_MintableBaseToken.CallOpts)
}

// InPrivateTransferMode is a free data retrieval call binding the contract method 0xdfbaefb1.
//
// Solidity: function inPrivateTransferMode() view returns(bool)
func (_MintableBaseToken *MintableBaseTokenCaller) InPrivateTransferMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MintableBaseToken.contract.Call(opts, &out, "inPrivateTransferMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InPrivateTransferMode is a free data retrieval call binding the contract method 0xdfbaefb1.
//
// Solidity: function inPrivateTransferMode() view returns(bool)
func (_MintableBaseToken *MintableBaseTokenSession) InPrivateTransferMode() (bool, error) {
	return _MintableBaseToken.Contract.InPrivateTransferMode(&_MintableBaseToken.CallOpts)
}

// InPrivateTransferMode is a free data retrieval call binding the contract method 0xdfbaefb1.
//
// Solidity: function inPrivateTransferMode() view returns(bool)
func (_MintableBaseToken *MintableBaseTokenCallerSession) InPrivateTransferMode() (bool, error) {
	return _MintableBaseToken.Contract.InPrivateTransferMode(&_MintableBaseToken.CallOpts)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_MintableBaseToken *MintableBaseTokenCaller) IsHandler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MintableBaseToken.contract.Call(opts, &out, "isHandler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_MintableBaseToken *MintableBaseTokenSession) IsHandler(arg0 common.Address) (bool, error) {
	return _MintableBaseToken.Contract.IsHandler(&_MintableBaseToken.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_MintableBaseToken *MintableBaseTokenCallerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _MintableBaseToken.Contract.IsHandler(&_MintableBaseToken.CallOpts, arg0)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_MintableBaseToken *MintableBaseTokenCaller) IsMinter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MintableBaseToken.contract.Call(opts, &out, "isMinter", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_MintableBaseToken *MintableBaseTokenSession) IsMinter(arg0 common.Address) (bool, error) {
	return _MintableBaseToken.Contract.IsMinter(&_MintableBaseToken.CallOpts, arg0)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address ) view returns(bool)
func (_MintableBaseToken *MintableBaseTokenCallerSession) IsMinter(arg0 common.Address) (bool, error) {
	return _MintableBaseToken.Contract.IsMinter(&_MintableBaseToken.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MintableBaseToken *MintableBaseTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MintableBaseToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MintableBaseToken *MintableBaseTokenSession) Name() (string, error) {
	return _MintableBaseToken.Contract.Name(&_MintableBaseToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MintableBaseToken *MintableBaseTokenCallerSession) Name() (string, error) {
	return _MintableBaseToken.Contract.Name(&_MintableBaseToken.CallOpts)
}

// NonStakingAccounts is a free data retrieval call binding the contract method 0x9554381a.
//
// Solidity: function nonStakingAccounts(address ) view returns(bool)
func (_MintableBaseToken *MintableBaseTokenCaller) NonStakingAccounts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MintableBaseToken.contract.Call(opts, &out, "nonStakingAccounts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NonStakingAccounts is a free data retrieval call binding the contract method 0x9554381a.
//
// Solidity: function nonStakingAccounts(address ) view returns(bool)
func (_MintableBaseToken *MintableBaseTokenSession) NonStakingAccounts(arg0 common.Address) (bool, error) {
	return _MintableBaseToken.Contract.NonStakingAccounts(&_MintableBaseToken.CallOpts, arg0)
}

// NonStakingAccounts is a free data retrieval call binding the contract method 0x9554381a.
//
// Solidity: function nonStakingAccounts(address ) view returns(bool)
func (_MintableBaseToken *MintableBaseTokenCallerSession) NonStakingAccounts(arg0 common.Address) (bool, error) {
	return _MintableBaseToken.Contract.NonStakingAccounts(&_MintableBaseToken.CallOpts, arg0)
}

// NonStakingSupply is a free data retrieval call binding the contract method 0xc93be636.
//
// Solidity: function nonStakingSupply() view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenCaller) NonStakingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintableBaseToken.contract.Call(opts, &out, "nonStakingSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonStakingSupply is a free data retrieval call binding the contract method 0xc93be636.
//
// Solidity: function nonStakingSupply() view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenSession) NonStakingSupply() (*big.Int, error) {
	return _MintableBaseToken.Contract.NonStakingSupply(&_MintableBaseToken.CallOpts)
}

// NonStakingSupply is a free data retrieval call binding the contract method 0xc93be636.
//
// Solidity: function nonStakingSupply() view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenCallerSession) NonStakingSupply() (*big.Int, error) {
	return _MintableBaseToken.Contract.NonStakingSupply(&_MintableBaseToken.CallOpts)
}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenCaller) StakedBalance(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MintableBaseToken.contract.Call(opts, &out, "stakedBalance", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenSession) StakedBalance(_account common.Address) (*big.Int, error) {
	return _MintableBaseToken.Contract.StakedBalance(&_MintableBaseToken.CallOpts, _account)
}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenCallerSession) StakedBalance(_account common.Address) (*big.Int, error) {
	return _MintableBaseToken.Contract.StakedBalance(&_MintableBaseToken.CallOpts, _account)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MintableBaseToken *MintableBaseTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MintableBaseToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MintableBaseToken *MintableBaseTokenSession) Symbol() (string, error) {
	return _MintableBaseToken.Contract.Symbol(&_MintableBaseToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MintableBaseToken *MintableBaseTokenCallerSession) Symbol() (string, error) {
	return _MintableBaseToken.Contract.Symbol(&_MintableBaseToken.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintableBaseToken.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenSession) TotalStaked() (*big.Int, error) {
	return _MintableBaseToken.Contract.TotalStaked(&_MintableBaseToken.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenCallerSession) TotalStaked() (*big.Int, error) {
	return _MintableBaseToken.Contract.TotalStaked(&_MintableBaseToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MintableBaseToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenSession) TotalSupply() (*big.Int, error) {
	return _MintableBaseToken.Contract.TotalSupply(&_MintableBaseToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MintableBaseToken *MintableBaseTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _MintableBaseToken.Contract.TotalSupply(&_MintableBaseToken.CallOpts)
}

// YieldTrackers is a free data retrieval call binding the contract method 0x52cd38d9.
//
// Solidity: function yieldTrackers(uint256 ) view returns(address)
func (_MintableBaseToken *MintableBaseTokenCaller) YieldTrackers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MintableBaseToken.contract.Call(opts, &out, "yieldTrackers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YieldTrackers is a free data retrieval call binding the contract method 0x52cd38d9.
//
// Solidity: function yieldTrackers(uint256 ) view returns(address)
func (_MintableBaseToken *MintableBaseTokenSession) YieldTrackers(arg0 *big.Int) (common.Address, error) {
	return _MintableBaseToken.Contract.YieldTrackers(&_MintableBaseToken.CallOpts, arg0)
}

// YieldTrackers is a free data retrieval call binding the contract method 0x52cd38d9.
//
// Solidity: function yieldTrackers(uint256 ) view returns(address)
func (_MintableBaseToken *MintableBaseTokenCallerSession) YieldTrackers(arg0 *big.Int) (common.Address, error) {
	return _MintableBaseToken.Contract.YieldTrackers(&_MintableBaseToken.CallOpts, arg0)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_MintableBaseToken *MintableBaseTokenTransactor) AddAdmin(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.contract.Transact(opts, "addAdmin", _account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_MintableBaseToken *MintableBaseTokenSession) AddAdmin(_account common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.AddAdmin(&_MintableBaseToken.TransactOpts, _account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_MintableBaseToken *MintableBaseTokenTransactorSession) AddAdmin(_account common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.AddAdmin(&_MintableBaseToken.TransactOpts, _account)
}

// AddNonStakingAccount is a paid mutator transaction binding the contract method 0x62289077.
//
// Solidity: function addNonStakingAccount(address _account) returns()
func (_MintableBaseToken *MintableBaseTokenTransactor) AddNonStakingAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.contract.Transact(opts, "addNonStakingAccount", _account)
}

// AddNonStakingAccount is a paid mutator transaction binding the contract method 0x62289077.
//
// Solidity: function addNonStakingAccount(address _account) returns()
func (_MintableBaseToken *MintableBaseTokenSession) AddNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.AddNonStakingAccount(&_MintableBaseToken.TransactOpts, _account)
}

// AddNonStakingAccount is a paid mutator transaction binding the contract method 0x62289077.
//
// Solidity: function addNonStakingAccount(address _account) returns()
func (_MintableBaseToken *MintableBaseTokenTransactorSession) AddNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.AddNonStakingAccount(&_MintableBaseToken.TransactOpts, _account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_MintableBaseToken *MintableBaseTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableBaseToken.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_MintableBaseToken *MintableBaseTokenSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.Approve(&_MintableBaseToken.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_MintableBaseToken *MintableBaseTokenTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.Approve(&_MintableBaseToken.TransactOpts, _spender, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_MintableBaseToken *MintableBaseTokenTransactor) Burn(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableBaseToken.contract.Transact(opts, "burn", _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_MintableBaseToken *MintableBaseTokenSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.Burn(&_MintableBaseToken.TransactOpts, _account, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _account, uint256 _amount) returns()
func (_MintableBaseToken *MintableBaseTokenTransactorSession) Burn(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.Burn(&_MintableBaseToken.TransactOpts, _account, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns()
func (_MintableBaseToken *MintableBaseTokenTransactor) Claim(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.contract.Transact(opts, "claim", _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns()
func (_MintableBaseToken *MintableBaseTokenSession) Claim(_receiver common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.Claim(&_MintableBaseToken.TransactOpts, _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns()
func (_MintableBaseToken *MintableBaseTokenTransactorSession) Claim(_receiver common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.Claim(&_MintableBaseToken.TransactOpts, _receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_MintableBaseToken *MintableBaseTokenTransactor) Mint(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableBaseToken.contract.Transact(opts, "mint", _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_MintableBaseToken *MintableBaseTokenSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.Mint(&_MintableBaseToken.TransactOpts, _account, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _account, uint256 _amount) returns()
func (_MintableBaseToken *MintableBaseTokenTransactorSession) Mint(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.Mint(&_MintableBaseToken.TransactOpts, _account, _amount)
}

// RecoverClaim is a paid mutator transaction binding the contract method 0x996f11ee.
//
// Solidity: function recoverClaim(address _account, address _receiver) returns()
func (_MintableBaseToken *MintableBaseTokenTransactor) RecoverClaim(opts *bind.TransactOpts, _account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.contract.Transact(opts, "recoverClaim", _account, _receiver)
}

// RecoverClaim is a paid mutator transaction binding the contract method 0x996f11ee.
//
// Solidity: function recoverClaim(address _account, address _receiver) returns()
func (_MintableBaseToken *MintableBaseTokenSession) RecoverClaim(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.RecoverClaim(&_MintableBaseToken.TransactOpts, _account, _receiver)
}

// RecoverClaim is a paid mutator transaction binding the contract method 0x996f11ee.
//
// Solidity: function recoverClaim(address _account, address _receiver) returns()
func (_MintableBaseToken *MintableBaseTokenTransactorSession) RecoverClaim(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.RecoverClaim(&_MintableBaseToken.TransactOpts, _account, _receiver)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_MintableBaseToken *MintableBaseTokenTransactor) RemoveAdmin(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.contract.Transact(opts, "removeAdmin", _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_MintableBaseToken *MintableBaseTokenSession) RemoveAdmin(_account common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.RemoveAdmin(&_MintableBaseToken.TransactOpts, _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_MintableBaseToken *MintableBaseTokenTransactorSession) RemoveAdmin(_account common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.RemoveAdmin(&_MintableBaseToken.TransactOpts, _account)
}

// RemoveNonStakingAccount is a paid mutator transaction binding the contract method 0xfb30d916.
//
// Solidity: function removeNonStakingAccount(address _account) returns()
func (_MintableBaseToken *MintableBaseTokenTransactor) RemoveNonStakingAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.contract.Transact(opts, "removeNonStakingAccount", _account)
}

// RemoveNonStakingAccount is a paid mutator transaction binding the contract method 0xfb30d916.
//
// Solidity: function removeNonStakingAccount(address _account) returns()
func (_MintableBaseToken *MintableBaseTokenSession) RemoveNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.RemoveNonStakingAccount(&_MintableBaseToken.TransactOpts, _account)
}

// RemoveNonStakingAccount is a paid mutator transaction binding the contract method 0xfb30d916.
//
// Solidity: function removeNonStakingAccount(address _account) returns()
func (_MintableBaseToken *MintableBaseTokenTransactorSession) RemoveNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.RemoveNonStakingAccount(&_MintableBaseToken.TransactOpts, _account)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_MintableBaseToken *MintableBaseTokenTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_MintableBaseToken *MintableBaseTokenSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.SetGov(&_MintableBaseToken.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_MintableBaseToken *MintableBaseTokenTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.SetGov(&_MintableBaseToken.TransactOpts, _gov)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_MintableBaseToken *MintableBaseTokenTransactor) SetHandler(opts *bind.TransactOpts, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _MintableBaseToken.contract.Transact(opts, "setHandler", _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_MintableBaseToken *MintableBaseTokenSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.SetHandler(&_MintableBaseToken.TransactOpts, _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_MintableBaseToken *MintableBaseTokenTransactorSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.SetHandler(&_MintableBaseToken.TransactOpts, _handler, _isActive)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x5a47a1a7.
//
// Solidity: function setInPrivateTransferMode(bool _inPrivateTransferMode) returns()
func (_MintableBaseToken *MintableBaseTokenTransactor) SetInPrivateTransferMode(opts *bind.TransactOpts, _inPrivateTransferMode bool) (*types.Transaction, error) {
	return _MintableBaseToken.contract.Transact(opts, "setInPrivateTransferMode", _inPrivateTransferMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x5a47a1a7.
//
// Solidity: function setInPrivateTransferMode(bool _inPrivateTransferMode) returns()
func (_MintableBaseToken *MintableBaseTokenSession) SetInPrivateTransferMode(_inPrivateTransferMode bool) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.SetInPrivateTransferMode(&_MintableBaseToken.TransactOpts, _inPrivateTransferMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x5a47a1a7.
//
// Solidity: function setInPrivateTransferMode(bool _inPrivateTransferMode) returns()
func (_MintableBaseToken *MintableBaseTokenTransactorSession) SetInPrivateTransferMode(_inPrivateTransferMode bool) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.SetInPrivateTransferMode(&_MintableBaseToken.TransactOpts, _inPrivateTransferMode)
}

// SetInfo is a paid mutator transaction binding the contract method 0xa923fc40.
//
// Solidity: function setInfo(string _name, string _symbol) returns()
func (_MintableBaseToken *MintableBaseTokenTransactor) SetInfo(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _MintableBaseToken.contract.Transact(opts, "setInfo", _name, _symbol)
}

// SetInfo is a paid mutator transaction binding the contract method 0xa923fc40.
//
// Solidity: function setInfo(string _name, string _symbol) returns()
func (_MintableBaseToken *MintableBaseTokenSession) SetInfo(_name string, _symbol string) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.SetInfo(&_MintableBaseToken.TransactOpts, _name, _symbol)
}

// SetInfo is a paid mutator transaction binding the contract method 0xa923fc40.
//
// Solidity: function setInfo(string _name, string _symbol) returns()
func (_MintableBaseToken *MintableBaseTokenTransactorSession) SetInfo(_name string, _symbol string) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.SetInfo(&_MintableBaseToken.TransactOpts, _name, _symbol)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address _minter, bool _isActive) returns()
func (_MintableBaseToken *MintableBaseTokenTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address, _isActive bool) (*types.Transaction, error) {
	return _MintableBaseToken.contract.Transact(opts, "setMinter", _minter, _isActive)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address _minter, bool _isActive) returns()
func (_MintableBaseToken *MintableBaseTokenSession) SetMinter(_minter common.Address, _isActive bool) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.SetMinter(&_MintableBaseToken.TransactOpts, _minter, _isActive)
}

// SetMinter is a paid mutator transaction binding the contract method 0xcf456ae7.
//
// Solidity: function setMinter(address _minter, bool _isActive) returns()
func (_MintableBaseToken *MintableBaseTokenTransactorSession) SetMinter(_minter common.Address, _isActive bool) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.SetMinter(&_MintableBaseToken.TransactOpts, _minter, _isActive)
}

// SetYieldTrackers is a paid mutator transaction binding the contract method 0x276eab4e.
//
// Solidity: function setYieldTrackers(address[] _yieldTrackers) returns()
func (_MintableBaseToken *MintableBaseTokenTransactor) SetYieldTrackers(opts *bind.TransactOpts, _yieldTrackers []common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.contract.Transact(opts, "setYieldTrackers", _yieldTrackers)
}

// SetYieldTrackers is a paid mutator transaction binding the contract method 0x276eab4e.
//
// Solidity: function setYieldTrackers(address[] _yieldTrackers) returns()
func (_MintableBaseToken *MintableBaseTokenSession) SetYieldTrackers(_yieldTrackers []common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.SetYieldTrackers(&_MintableBaseToken.TransactOpts, _yieldTrackers)
}

// SetYieldTrackers is a paid mutator transaction binding the contract method 0x276eab4e.
//
// Solidity: function setYieldTrackers(address[] _yieldTrackers) returns()
func (_MintableBaseToken *MintableBaseTokenTransactorSession) SetYieldTrackers(_yieldTrackers []common.Address) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.SetYieldTrackers(&_MintableBaseToken.TransactOpts, _yieldTrackers)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_MintableBaseToken *MintableBaseTokenTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableBaseToken.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_MintableBaseToken *MintableBaseTokenSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.Transfer(&_MintableBaseToken.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_MintableBaseToken *MintableBaseTokenTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.Transfer(&_MintableBaseToken.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_MintableBaseToken *MintableBaseTokenTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableBaseToken.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_MintableBaseToken *MintableBaseTokenSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.TransferFrom(&_MintableBaseToken.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_MintableBaseToken *MintableBaseTokenTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.TransferFrom(&_MintableBaseToken.TransactOpts, _sender, _recipient, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_MintableBaseToken *MintableBaseTokenTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableBaseToken.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_MintableBaseToken *MintableBaseTokenSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.WithdrawToken(&_MintableBaseToken.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_MintableBaseToken *MintableBaseTokenTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MintableBaseToken.Contract.WithdrawToken(&_MintableBaseToken.TransactOpts, _token, _account, _amount)
}

// MintableBaseTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MintableBaseToken contract.
type MintableBaseTokenApprovalIterator struct {
	Event *MintableBaseTokenApproval // Event containing the contract specifics and raw log

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
func (it *MintableBaseTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintableBaseTokenApproval)
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
		it.Event = new(MintableBaseTokenApproval)
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
func (it *MintableBaseTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintableBaseTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintableBaseTokenApproval represents a Approval event raised by the MintableBaseToken contract.
type MintableBaseTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MintableBaseToken *MintableBaseTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MintableBaseTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MintableBaseToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MintableBaseTokenApprovalIterator{contract: _MintableBaseToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MintableBaseToken *MintableBaseTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MintableBaseTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MintableBaseToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintableBaseTokenApproval)
				if err := _MintableBaseToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MintableBaseToken *MintableBaseTokenFilterer) ParseApproval(log types.Log) (*MintableBaseTokenApproval, error) {
	event := new(MintableBaseTokenApproval)
	if err := _MintableBaseToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintableBaseTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MintableBaseToken contract.
type MintableBaseTokenTransferIterator struct {
	Event *MintableBaseTokenTransfer // Event containing the contract specifics and raw log

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
func (it *MintableBaseTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintableBaseTokenTransfer)
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
		it.Event = new(MintableBaseTokenTransfer)
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
func (it *MintableBaseTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintableBaseTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintableBaseTokenTransfer represents a Transfer event raised by the MintableBaseToken contract.
type MintableBaseTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MintableBaseToken *MintableBaseTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MintableBaseTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MintableBaseToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MintableBaseTokenTransferIterator{contract: _MintableBaseToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MintableBaseToken *MintableBaseTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MintableBaseTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MintableBaseToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintableBaseTokenTransfer)
				if err := _MintableBaseToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MintableBaseToken *MintableBaseTokenFilterer) ParseTransfer(log types.Log) (*MintableBaseTokenTransfer, error) {
	event := new(MintableBaseTokenTransfer)
	if err := _MintableBaseToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

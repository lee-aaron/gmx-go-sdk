// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package YieldFarm

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

// YieldFarmMetaData contains all meta data concerning the YieldFarm contract.
var YieldFarmMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_stakingToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addAdmin\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addNonStakingAccount\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"admins\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inWhitelistMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonStakingAccounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonStakingSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"recoverClaim\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeAdmin\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeNonStakingAccount\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInWhitelistMode\",\"inputs\":[{\"name\":\"_inWhitelistMode\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInfo\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWhitelistedHandler\",\"inputs\":[{\"name\":\"_handler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isWhitelisted\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setYieldTrackers\",\"inputs\":[{\"name\":\"_yieldTrackers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakedBalance\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalStaked\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstake\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"whitelistedHandlers\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"yieldTrackers\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// YieldFarmABI is the input ABI used to generate the binding from.
// Deprecated: Use YieldFarmMetaData.ABI instead.
var YieldFarmABI = YieldFarmMetaData.ABI

// YieldFarm is an auto generated Go binding around an Ethereum contract.
type YieldFarm struct {
	YieldFarmCaller     // Read-only binding to the contract
	YieldFarmTransactor // Write-only binding to the contract
	YieldFarmFilterer   // Log filterer for contract events
}

// YieldFarmCaller is an auto generated read-only Go binding around an Ethereum contract.
type YieldFarmCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldFarmTransactor is an auto generated write-only Go binding around an Ethereum contract.
type YieldFarmTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldFarmFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type YieldFarmFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// YieldFarmSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type YieldFarmSession struct {
	Contract     *YieldFarm        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// YieldFarmCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type YieldFarmCallerSession struct {
	Contract *YieldFarmCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// YieldFarmTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type YieldFarmTransactorSession struct {
	Contract     *YieldFarmTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// YieldFarmRaw is an auto generated low-level Go binding around an Ethereum contract.
type YieldFarmRaw struct {
	Contract *YieldFarm // Generic contract binding to access the raw methods on
}

// YieldFarmCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type YieldFarmCallerRaw struct {
	Contract *YieldFarmCaller // Generic read-only contract binding to access the raw methods on
}

// YieldFarmTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type YieldFarmTransactorRaw struct {
	Contract *YieldFarmTransactor // Generic write-only contract binding to access the raw methods on
}

// NewYieldFarm creates a new instance of YieldFarm, bound to a specific deployed contract.
func NewYieldFarm(address common.Address, backend bind.ContractBackend) (*YieldFarm, error) {
	contract, err := bindYieldFarm(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &YieldFarm{YieldFarmCaller: YieldFarmCaller{contract: contract}, YieldFarmTransactor: YieldFarmTransactor{contract: contract}, YieldFarmFilterer: YieldFarmFilterer{contract: contract}}, nil
}

// NewYieldFarmCaller creates a new read-only instance of YieldFarm, bound to a specific deployed contract.
func NewYieldFarmCaller(address common.Address, caller bind.ContractCaller) (*YieldFarmCaller, error) {
	contract, err := bindYieldFarm(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &YieldFarmCaller{contract: contract}, nil
}

// NewYieldFarmTransactor creates a new write-only instance of YieldFarm, bound to a specific deployed contract.
func NewYieldFarmTransactor(address common.Address, transactor bind.ContractTransactor) (*YieldFarmTransactor, error) {
	contract, err := bindYieldFarm(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &YieldFarmTransactor{contract: contract}, nil
}

// NewYieldFarmFilterer creates a new log filterer instance of YieldFarm, bound to a specific deployed contract.
func NewYieldFarmFilterer(address common.Address, filterer bind.ContractFilterer) (*YieldFarmFilterer, error) {
	contract, err := bindYieldFarm(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &YieldFarmFilterer{contract: contract}, nil
}

// bindYieldFarm binds a generic wrapper to an already deployed contract.
func bindYieldFarm(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := YieldFarmMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YieldFarm *YieldFarmRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YieldFarm.Contract.YieldFarmCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YieldFarm *YieldFarmRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YieldFarm.Contract.YieldFarmTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YieldFarm *YieldFarmRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YieldFarm.Contract.YieldFarmTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_YieldFarm *YieldFarmCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _YieldFarm.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_YieldFarm *YieldFarmTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _YieldFarm.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_YieldFarm *YieldFarmTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _YieldFarm.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_YieldFarm *YieldFarmCaller) Admins(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YieldFarm.contract.Call(opts, &out, "admins", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_YieldFarm *YieldFarmSession) Admins(arg0 common.Address) (bool, error) {
	return _YieldFarm.Contract.Admins(&_YieldFarm.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x429b62e5.
//
// Solidity: function admins(address ) view returns(bool)
func (_YieldFarm *YieldFarmCallerSession) Admins(arg0 common.Address) (bool, error) {
	return _YieldFarm.Contract.Admins(&_YieldFarm.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_YieldFarm *YieldFarmCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YieldFarm.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_YieldFarm *YieldFarmSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _YieldFarm.Contract.Allowance(&_YieldFarm.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_YieldFarm *YieldFarmCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _YieldFarm.Contract.Allowance(&_YieldFarm.CallOpts, _owner, _spender)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_YieldFarm *YieldFarmCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YieldFarm.contract.Call(opts, &out, "allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_YieldFarm *YieldFarmSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YieldFarm.Contract.Allowances(&_YieldFarm.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_YieldFarm *YieldFarmCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _YieldFarm.Contract.Allowances(&_YieldFarm.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_YieldFarm *YieldFarmCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YieldFarm.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_YieldFarm *YieldFarmSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _YieldFarm.Contract.BalanceOf(&_YieldFarm.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_YieldFarm *YieldFarmCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _YieldFarm.Contract.BalanceOf(&_YieldFarm.CallOpts, _account)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_YieldFarm *YieldFarmCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YieldFarm.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_YieldFarm *YieldFarmSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _YieldFarm.Contract.Balances(&_YieldFarm.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_YieldFarm *YieldFarmCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _YieldFarm.Contract.Balances(&_YieldFarm.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YieldFarm *YieldFarmCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _YieldFarm.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YieldFarm *YieldFarmSession) Decimals() (uint8, error) {
	return _YieldFarm.Contract.Decimals(&_YieldFarm.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_YieldFarm *YieldFarmCallerSession) Decimals() (uint8, error) {
	return _YieldFarm.Contract.Decimals(&_YieldFarm.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_YieldFarm *YieldFarmCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YieldFarm.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_YieldFarm *YieldFarmSession) Gov() (common.Address, error) {
	return _YieldFarm.Contract.Gov(&_YieldFarm.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_YieldFarm *YieldFarmCallerSession) Gov() (common.Address, error) {
	return _YieldFarm.Contract.Gov(&_YieldFarm.CallOpts)
}

// InWhitelistMode is a free data retrieval call binding the contract method 0x293d6987.
//
// Solidity: function inWhitelistMode() view returns(bool)
func (_YieldFarm *YieldFarmCaller) InWhitelistMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _YieldFarm.contract.Call(opts, &out, "inWhitelistMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InWhitelistMode is a free data retrieval call binding the contract method 0x293d6987.
//
// Solidity: function inWhitelistMode() view returns(bool)
func (_YieldFarm *YieldFarmSession) InWhitelistMode() (bool, error) {
	return _YieldFarm.Contract.InWhitelistMode(&_YieldFarm.CallOpts)
}

// InWhitelistMode is a free data retrieval call binding the contract method 0x293d6987.
//
// Solidity: function inWhitelistMode() view returns(bool)
func (_YieldFarm *YieldFarmCallerSession) InWhitelistMode() (bool, error) {
	return _YieldFarm.Contract.InWhitelistMode(&_YieldFarm.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YieldFarm *YieldFarmCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YieldFarm.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YieldFarm *YieldFarmSession) Name() (string, error) {
	return _YieldFarm.Contract.Name(&_YieldFarm.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_YieldFarm *YieldFarmCallerSession) Name() (string, error) {
	return _YieldFarm.Contract.Name(&_YieldFarm.CallOpts)
}

// NonStakingAccounts is a free data retrieval call binding the contract method 0x9554381a.
//
// Solidity: function nonStakingAccounts(address ) view returns(bool)
func (_YieldFarm *YieldFarmCaller) NonStakingAccounts(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YieldFarm.contract.Call(opts, &out, "nonStakingAccounts", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NonStakingAccounts is a free data retrieval call binding the contract method 0x9554381a.
//
// Solidity: function nonStakingAccounts(address ) view returns(bool)
func (_YieldFarm *YieldFarmSession) NonStakingAccounts(arg0 common.Address) (bool, error) {
	return _YieldFarm.Contract.NonStakingAccounts(&_YieldFarm.CallOpts, arg0)
}

// NonStakingAccounts is a free data retrieval call binding the contract method 0x9554381a.
//
// Solidity: function nonStakingAccounts(address ) view returns(bool)
func (_YieldFarm *YieldFarmCallerSession) NonStakingAccounts(arg0 common.Address) (bool, error) {
	return _YieldFarm.Contract.NonStakingAccounts(&_YieldFarm.CallOpts, arg0)
}

// NonStakingSupply is a free data retrieval call binding the contract method 0xc93be636.
//
// Solidity: function nonStakingSupply() view returns(uint256)
func (_YieldFarm *YieldFarmCaller) NonStakingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YieldFarm.contract.Call(opts, &out, "nonStakingSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NonStakingSupply is a free data retrieval call binding the contract method 0xc93be636.
//
// Solidity: function nonStakingSupply() view returns(uint256)
func (_YieldFarm *YieldFarmSession) NonStakingSupply() (*big.Int, error) {
	return _YieldFarm.Contract.NonStakingSupply(&_YieldFarm.CallOpts)
}

// NonStakingSupply is a free data retrieval call binding the contract method 0xc93be636.
//
// Solidity: function nonStakingSupply() view returns(uint256)
func (_YieldFarm *YieldFarmCallerSession) NonStakingSupply() (*big.Int, error) {
	return _YieldFarm.Contract.NonStakingSupply(&_YieldFarm.CallOpts)
}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_YieldFarm *YieldFarmCaller) StakedBalance(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _YieldFarm.contract.Call(opts, &out, "stakedBalance", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_YieldFarm *YieldFarmSession) StakedBalance(_account common.Address) (*big.Int, error) {
	return _YieldFarm.Contract.StakedBalance(&_YieldFarm.CallOpts, _account)
}

// StakedBalance is a free data retrieval call binding the contract method 0x60217267.
//
// Solidity: function stakedBalance(address _account) view returns(uint256)
func (_YieldFarm *YieldFarmCallerSession) StakedBalance(_account common.Address) (*big.Int, error) {
	return _YieldFarm.Contract.StakedBalance(&_YieldFarm.CallOpts, _account)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_YieldFarm *YieldFarmCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _YieldFarm.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_YieldFarm *YieldFarmSession) StakingToken() (common.Address, error) {
	return _YieldFarm.Contract.StakingToken(&_YieldFarm.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_YieldFarm *YieldFarmCallerSession) StakingToken() (common.Address, error) {
	return _YieldFarm.Contract.StakingToken(&_YieldFarm.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YieldFarm *YieldFarmCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _YieldFarm.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YieldFarm *YieldFarmSession) Symbol() (string, error) {
	return _YieldFarm.Contract.Symbol(&_YieldFarm.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_YieldFarm *YieldFarmCallerSession) Symbol() (string, error) {
	return _YieldFarm.Contract.Symbol(&_YieldFarm.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_YieldFarm *YieldFarmCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YieldFarm.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_YieldFarm *YieldFarmSession) TotalStaked() (*big.Int, error) {
	return _YieldFarm.Contract.TotalStaked(&_YieldFarm.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_YieldFarm *YieldFarmCallerSession) TotalStaked() (*big.Int, error) {
	return _YieldFarm.Contract.TotalStaked(&_YieldFarm.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YieldFarm *YieldFarmCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _YieldFarm.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YieldFarm *YieldFarmSession) TotalSupply() (*big.Int, error) {
	return _YieldFarm.Contract.TotalSupply(&_YieldFarm.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_YieldFarm *YieldFarmCallerSession) TotalSupply() (*big.Int, error) {
	return _YieldFarm.Contract.TotalSupply(&_YieldFarm.CallOpts)
}

// WhitelistedHandlers is a free data retrieval call binding the contract method 0x36300051.
//
// Solidity: function whitelistedHandlers(address ) view returns(bool)
func (_YieldFarm *YieldFarmCaller) WhitelistedHandlers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _YieldFarm.contract.Call(opts, &out, "whitelistedHandlers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedHandlers is a free data retrieval call binding the contract method 0x36300051.
//
// Solidity: function whitelistedHandlers(address ) view returns(bool)
func (_YieldFarm *YieldFarmSession) WhitelistedHandlers(arg0 common.Address) (bool, error) {
	return _YieldFarm.Contract.WhitelistedHandlers(&_YieldFarm.CallOpts, arg0)
}

// WhitelistedHandlers is a free data retrieval call binding the contract method 0x36300051.
//
// Solidity: function whitelistedHandlers(address ) view returns(bool)
func (_YieldFarm *YieldFarmCallerSession) WhitelistedHandlers(arg0 common.Address) (bool, error) {
	return _YieldFarm.Contract.WhitelistedHandlers(&_YieldFarm.CallOpts, arg0)
}

// YieldTrackers is a free data retrieval call binding the contract method 0x52cd38d9.
//
// Solidity: function yieldTrackers(uint256 ) view returns(address)
func (_YieldFarm *YieldFarmCaller) YieldTrackers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _YieldFarm.contract.Call(opts, &out, "yieldTrackers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// YieldTrackers is a free data retrieval call binding the contract method 0x52cd38d9.
//
// Solidity: function yieldTrackers(uint256 ) view returns(address)
func (_YieldFarm *YieldFarmSession) YieldTrackers(arg0 *big.Int) (common.Address, error) {
	return _YieldFarm.Contract.YieldTrackers(&_YieldFarm.CallOpts, arg0)
}

// YieldTrackers is a free data retrieval call binding the contract method 0x52cd38d9.
//
// Solidity: function yieldTrackers(uint256 ) view returns(address)
func (_YieldFarm *YieldFarmCallerSession) YieldTrackers(arg0 *big.Int) (common.Address, error) {
	return _YieldFarm.Contract.YieldTrackers(&_YieldFarm.CallOpts, arg0)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_YieldFarm *YieldFarmTransactor) AddAdmin(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _YieldFarm.contract.Transact(opts, "addAdmin", _account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_YieldFarm *YieldFarmSession) AddAdmin(_account common.Address) (*types.Transaction, error) {
	return _YieldFarm.Contract.AddAdmin(&_YieldFarm.TransactOpts, _account)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address _account) returns()
func (_YieldFarm *YieldFarmTransactorSession) AddAdmin(_account common.Address) (*types.Transaction, error) {
	return _YieldFarm.Contract.AddAdmin(&_YieldFarm.TransactOpts, _account)
}

// AddNonStakingAccount is a paid mutator transaction binding the contract method 0x62289077.
//
// Solidity: function addNonStakingAccount(address _account) returns()
func (_YieldFarm *YieldFarmTransactor) AddNonStakingAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _YieldFarm.contract.Transact(opts, "addNonStakingAccount", _account)
}

// AddNonStakingAccount is a paid mutator transaction binding the contract method 0x62289077.
//
// Solidity: function addNonStakingAccount(address _account) returns()
func (_YieldFarm *YieldFarmSession) AddNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _YieldFarm.Contract.AddNonStakingAccount(&_YieldFarm.TransactOpts, _account)
}

// AddNonStakingAccount is a paid mutator transaction binding the contract method 0x62289077.
//
// Solidity: function addNonStakingAccount(address _account) returns()
func (_YieldFarm *YieldFarmTransactorSession) AddNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _YieldFarm.Contract.AddNonStakingAccount(&_YieldFarm.TransactOpts, _account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_YieldFarm *YieldFarmTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldFarm.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_YieldFarm *YieldFarmSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldFarm.Contract.Approve(&_YieldFarm.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_YieldFarm *YieldFarmTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldFarm.Contract.Approve(&_YieldFarm.TransactOpts, _spender, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns()
func (_YieldFarm *YieldFarmTransactor) Claim(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _YieldFarm.contract.Transact(opts, "claim", _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns()
func (_YieldFarm *YieldFarmSession) Claim(_receiver common.Address) (*types.Transaction, error) {
	return _YieldFarm.Contract.Claim(&_YieldFarm.TransactOpts, _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns()
func (_YieldFarm *YieldFarmTransactorSession) Claim(_receiver common.Address) (*types.Transaction, error) {
	return _YieldFarm.Contract.Claim(&_YieldFarm.TransactOpts, _receiver)
}

// RecoverClaim is a paid mutator transaction binding the contract method 0x996f11ee.
//
// Solidity: function recoverClaim(address _account, address _receiver) returns()
func (_YieldFarm *YieldFarmTransactor) RecoverClaim(opts *bind.TransactOpts, _account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _YieldFarm.contract.Transact(opts, "recoverClaim", _account, _receiver)
}

// RecoverClaim is a paid mutator transaction binding the contract method 0x996f11ee.
//
// Solidity: function recoverClaim(address _account, address _receiver) returns()
func (_YieldFarm *YieldFarmSession) RecoverClaim(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _YieldFarm.Contract.RecoverClaim(&_YieldFarm.TransactOpts, _account, _receiver)
}

// RecoverClaim is a paid mutator transaction binding the contract method 0x996f11ee.
//
// Solidity: function recoverClaim(address _account, address _receiver) returns()
func (_YieldFarm *YieldFarmTransactorSession) RecoverClaim(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _YieldFarm.Contract.RecoverClaim(&_YieldFarm.TransactOpts, _account, _receiver)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_YieldFarm *YieldFarmTransactor) RemoveAdmin(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _YieldFarm.contract.Transact(opts, "removeAdmin", _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_YieldFarm *YieldFarmSession) RemoveAdmin(_account common.Address) (*types.Transaction, error) {
	return _YieldFarm.Contract.RemoveAdmin(&_YieldFarm.TransactOpts, _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address _account) returns()
func (_YieldFarm *YieldFarmTransactorSession) RemoveAdmin(_account common.Address) (*types.Transaction, error) {
	return _YieldFarm.Contract.RemoveAdmin(&_YieldFarm.TransactOpts, _account)
}

// RemoveNonStakingAccount is a paid mutator transaction binding the contract method 0xfb30d916.
//
// Solidity: function removeNonStakingAccount(address _account) returns()
func (_YieldFarm *YieldFarmTransactor) RemoveNonStakingAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _YieldFarm.contract.Transact(opts, "removeNonStakingAccount", _account)
}

// RemoveNonStakingAccount is a paid mutator transaction binding the contract method 0xfb30d916.
//
// Solidity: function removeNonStakingAccount(address _account) returns()
func (_YieldFarm *YieldFarmSession) RemoveNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _YieldFarm.Contract.RemoveNonStakingAccount(&_YieldFarm.TransactOpts, _account)
}

// RemoveNonStakingAccount is a paid mutator transaction binding the contract method 0xfb30d916.
//
// Solidity: function removeNonStakingAccount(address _account) returns()
func (_YieldFarm *YieldFarmTransactorSession) RemoveNonStakingAccount(_account common.Address) (*types.Transaction, error) {
	return _YieldFarm.Contract.RemoveNonStakingAccount(&_YieldFarm.TransactOpts, _account)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_YieldFarm *YieldFarmTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _YieldFarm.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_YieldFarm *YieldFarmSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _YieldFarm.Contract.SetGov(&_YieldFarm.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_YieldFarm *YieldFarmTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _YieldFarm.Contract.SetGov(&_YieldFarm.TransactOpts, _gov)
}

// SetInWhitelistMode is a paid mutator transaction binding the contract method 0x4cb5bbe3.
//
// Solidity: function setInWhitelistMode(bool _inWhitelistMode) returns()
func (_YieldFarm *YieldFarmTransactor) SetInWhitelistMode(opts *bind.TransactOpts, _inWhitelistMode bool) (*types.Transaction, error) {
	return _YieldFarm.contract.Transact(opts, "setInWhitelistMode", _inWhitelistMode)
}

// SetInWhitelistMode is a paid mutator transaction binding the contract method 0x4cb5bbe3.
//
// Solidity: function setInWhitelistMode(bool _inWhitelistMode) returns()
func (_YieldFarm *YieldFarmSession) SetInWhitelistMode(_inWhitelistMode bool) (*types.Transaction, error) {
	return _YieldFarm.Contract.SetInWhitelistMode(&_YieldFarm.TransactOpts, _inWhitelistMode)
}

// SetInWhitelistMode is a paid mutator transaction binding the contract method 0x4cb5bbe3.
//
// Solidity: function setInWhitelistMode(bool _inWhitelistMode) returns()
func (_YieldFarm *YieldFarmTransactorSession) SetInWhitelistMode(_inWhitelistMode bool) (*types.Transaction, error) {
	return _YieldFarm.Contract.SetInWhitelistMode(&_YieldFarm.TransactOpts, _inWhitelistMode)
}

// SetInfo is a paid mutator transaction binding the contract method 0xa923fc40.
//
// Solidity: function setInfo(string _name, string _symbol) returns()
func (_YieldFarm *YieldFarmTransactor) SetInfo(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _YieldFarm.contract.Transact(opts, "setInfo", _name, _symbol)
}

// SetInfo is a paid mutator transaction binding the contract method 0xa923fc40.
//
// Solidity: function setInfo(string _name, string _symbol) returns()
func (_YieldFarm *YieldFarmSession) SetInfo(_name string, _symbol string) (*types.Transaction, error) {
	return _YieldFarm.Contract.SetInfo(&_YieldFarm.TransactOpts, _name, _symbol)
}

// SetInfo is a paid mutator transaction binding the contract method 0xa923fc40.
//
// Solidity: function setInfo(string _name, string _symbol) returns()
func (_YieldFarm *YieldFarmTransactorSession) SetInfo(_name string, _symbol string) (*types.Transaction, error) {
	return _YieldFarm.Contract.SetInfo(&_YieldFarm.TransactOpts, _name, _symbol)
}

// SetWhitelistedHandler is a paid mutator transaction binding the contract method 0xd92fc87e.
//
// Solidity: function setWhitelistedHandler(address _handler, bool _isWhitelisted) returns()
func (_YieldFarm *YieldFarmTransactor) SetWhitelistedHandler(opts *bind.TransactOpts, _handler common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _YieldFarm.contract.Transact(opts, "setWhitelistedHandler", _handler, _isWhitelisted)
}

// SetWhitelistedHandler is a paid mutator transaction binding the contract method 0xd92fc87e.
//
// Solidity: function setWhitelistedHandler(address _handler, bool _isWhitelisted) returns()
func (_YieldFarm *YieldFarmSession) SetWhitelistedHandler(_handler common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _YieldFarm.Contract.SetWhitelistedHandler(&_YieldFarm.TransactOpts, _handler, _isWhitelisted)
}

// SetWhitelistedHandler is a paid mutator transaction binding the contract method 0xd92fc87e.
//
// Solidity: function setWhitelistedHandler(address _handler, bool _isWhitelisted) returns()
func (_YieldFarm *YieldFarmTransactorSession) SetWhitelistedHandler(_handler common.Address, _isWhitelisted bool) (*types.Transaction, error) {
	return _YieldFarm.Contract.SetWhitelistedHandler(&_YieldFarm.TransactOpts, _handler, _isWhitelisted)
}

// SetYieldTrackers is a paid mutator transaction binding the contract method 0x276eab4e.
//
// Solidity: function setYieldTrackers(address[] _yieldTrackers) returns()
func (_YieldFarm *YieldFarmTransactor) SetYieldTrackers(opts *bind.TransactOpts, _yieldTrackers []common.Address) (*types.Transaction, error) {
	return _YieldFarm.contract.Transact(opts, "setYieldTrackers", _yieldTrackers)
}

// SetYieldTrackers is a paid mutator transaction binding the contract method 0x276eab4e.
//
// Solidity: function setYieldTrackers(address[] _yieldTrackers) returns()
func (_YieldFarm *YieldFarmSession) SetYieldTrackers(_yieldTrackers []common.Address) (*types.Transaction, error) {
	return _YieldFarm.Contract.SetYieldTrackers(&_YieldFarm.TransactOpts, _yieldTrackers)
}

// SetYieldTrackers is a paid mutator transaction binding the contract method 0x276eab4e.
//
// Solidity: function setYieldTrackers(address[] _yieldTrackers) returns()
func (_YieldFarm *YieldFarmTransactorSession) SetYieldTrackers(_yieldTrackers []common.Address) (*types.Transaction, error) {
	return _YieldFarm.Contract.SetYieldTrackers(&_YieldFarm.TransactOpts, _yieldTrackers)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_YieldFarm *YieldFarmTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _YieldFarm.contract.Transact(opts, "stake", _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_YieldFarm *YieldFarmSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _YieldFarm.Contract.Stake(&_YieldFarm.TransactOpts, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_YieldFarm *YieldFarmTransactorSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _YieldFarm.Contract.Stake(&_YieldFarm.TransactOpts, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_YieldFarm *YieldFarmTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldFarm.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_YieldFarm *YieldFarmSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldFarm.Contract.Transfer(&_YieldFarm.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_YieldFarm *YieldFarmTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldFarm.Contract.Transfer(&_YieldFarm.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_YieldFarm *YieldFarmTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldFarm.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_YieldFarm *YieldFarmSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldFarm.Contract.TransferFrom(&_YieldFarm.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_YieldFarm *YieldFarmTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldFarm.Contract.TransferFrom(&_YieldFarm.TransactOpts, _sender, _recipient, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns()
func (_YieldFarm *YieldFarmTransactor) Unstake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _YieldFarm.contract.Transact(opts, "unstake", _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns()
func (_YieldFarm *YieldFarmSession) Unstake(_amount *big.Int) (*types.Transaction, error) {
	return _YieldFarm.Contract.Unstake(&_YieldFarm.TransactOpts, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x2e17de78.
//
// Solidity: function unstake(uint256 _amount) returns()
func (_YieldFarm *YieldFarmTransactorSession) Unstake(_amount *big.Int) (*types.Transaction, error) {
	return _YieldFarm.Contract.Unstake(&_YieldFarm.TransactOpts, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_YieldFarm *YieldFarmTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldFarm.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_YieldFarm *YieldFarmSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldFarm.Contract.WithdrawToken(&_YieldFarm.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_YieldFarm *YieldFarmTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _YieldFarm.Contract.WithdrawToken(&_YieldFarm.TransactOpts, _token, _account, _amount)
}

// YieldFarmApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the YieldFarm contract.
type YieldFarmApprovalIterator struct {
	Event *YieldFarmApproval // Event containing the contract specifics and raw log

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
func (it *YieldFarmApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YieldFarmApproval)
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
		it.Event = new(YieldFarmApproval)
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
func (it *YieldFarmApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YieldFarmApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YieldFarmApproval represents a Approval event raised by the YieldFarm contract.
type YieldFarmApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YieldFarm *YieldFarmFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*YieldFarmApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YieldFarm.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &YieldFarmApprovalIterator{contract: _YieldFarm.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_YieldFarm *YieldFarmFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *YieldFarmApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _YieldFarm.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YieldFarmApproval)
				if err := _YieldFarm.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_YieldFarm *YieldFarmFilterer) ParseApproval(log types.Log) (*YieldFarmApproval, error) {
	event := new(YieldFarmApproval)
	if err := _YieldFarm.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// YieldFarmTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the YieldFarm contract.
type YieldFarmTransferIterator struct {
	Event *YieldFarmTransfer // Event containing the contract specifics and raw log

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
func (it *YieldFarmTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(YieldFarmTransfer)
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
		it.Event = new(YieldFarmTransfer)
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
func (it *YieldFarmTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *YieldFarmTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// YieldFarmTransfer represents a Transfer event raised by the YieldFarm contract.
type YieldFarmTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YieldFarm *YieldFarmFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*YieldFarmTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YieldFarm.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &YieldFarmTransferIterator{contract: _YieldFarm.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_YieldFarm *YieldFarmFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *YieldFarmTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _YieldFarm.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(YieldFarmTransfer)
				if err := _YieldFarm.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_YieldFarm *YieldFarmFilterer) ParseTransfer(log types.Log) (*YieldFarmTransfer, error) {
	event := new(YieldFarmTransfer)
	if err := _YieldFarm.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

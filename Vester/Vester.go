// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Vester

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

// VesterMetaData contains all meta data concerning the Vester contract.
var VesterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_vestingDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_esToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_pairToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_claimableToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_rewardTracker\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bonusRewards\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimForAccount\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimable\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimableToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimedAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cumulativeClaimAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cumulativeRewardDeductions\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositForAccount\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"esToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCombinedAverageStakedAmount\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxVestableAmount\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPairAmount\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_esAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalVested\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVestedAmount\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasMaxVestableAmount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasPairToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasRewardTracker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isHandler\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastVestingTimes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pairToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardTracker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBonusRewards\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCumulativeRewardDeductions\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setHandler\",\"inputs\":[{\"name\":\"_handler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setHasMaxVestableAmount\",\"inputs\":[{\"name\":\"_hasMaxVestableAmount\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTransferredAverageStakedAmounts\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTransferredCumulativeRewards\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferStakeValues\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferredAverageStakedAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferredCumulativeRewards\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vestingDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claim\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PairTransfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"claimedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// VesterABI is the input ABI used to generate the binding from.
// Deprecated: Use VesterMetaData.ABI instead.
var VesterABI = VesterMetaData.ABI

// Vester is an auto generated Go binding around an Ethereum contract.
type Vester struct {
	VesterCaller     // Read-only binding to the contract
	VesterTransactor // Write-only binding to the contract
	VesterFilterer   // Log filterer for contract events
}

// VesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type VesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VesterSession struct {
	Contract     *Vester           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VesterCallerSession struct {
	Contract *VesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VesterTransactorSession struct {
	Contract     *VesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type VesterRaw struct {
	Contract *Vester // Generic contract binding to access the raw methods on
}

// VesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VesterCallerRaw struct {
	Contract *VesterCaller // Generic read-only contract binding to access the raw methods on
}

// VesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VesterTransactorRaw struct {
	Contract *VesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVester creates a new instance of Vester, bound to a specific deployed contract.
func NewVester(address common.Address, backend bind.ContractBackend) (*Vester, error) {
	contract, err := bindVester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vester{VesterCaller: VesterCaller{contract: contract}, VesterTransactor: VesterTransactor{contract: contract}, VesterFilterer: VesterFilterer{contract: contract}}, nil
}

// NewVesterCaller creates a new read-only instance of Vester, bound to a specific deployed contract.
func NewVesterCaller(address common.Address, caller bind.ContractCaller) (*VesterCaller, error) {
	contract, err := bindVester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VesterCaller{contract: contract}, nil
}

// NewVesterTransactor creates a new write-only instance of Vester, bound to a specific deployed contract.
func NewVesterTransactor(address common.Address, transactor bind.ContractTransactor) (*VesterTransactor, error) {
	contract, err := bindVester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VesterTransactor{contract: contract}, nil
}

// NewVesterFilterer creates a new log filterer instance of Vester, bound to a specific deployed contract.
func NewVesterFilterer(address common.Address, filterer bind.ContractFilterer) (*VesterFilterer, error) {
	contract, err := bindVester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VesterFilterer{contract: contract}, nil
}

// bindVester binds a generic wrapper to an already deployed contract.
func bindVester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VesterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vester *VesterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vester.Contract.VesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vester *VesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vester.Contract.VesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vester *VesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vester.Contract.VesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vester *VesterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vester *VesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vester *VesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vester.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Vester *VesterCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Vester *VesterSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vester.Contract.Allowance(&_Vester.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Vester *VesterCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vester.Contract.Allowance(&_Vester.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Vester *VesterCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Vester *VesterSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Vester.Contract.BalanceOf(&_Vester.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_Vester *VesterCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _Vester.Contract.BalanceOf(&_Vester.CallOpts, _account)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Vester *VesterCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Vester *VesterSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Vester.Contract.Balances(&_Vester.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Vester *VesterCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Vester.Contract.Balances(&_Vester.CallOpts, arg0)
}

// BonusRewards is a free data retrieval call binding the contract method 0xa2545fa5.
//
// Solidity: function bonusRewards(address ) view returns(uint256)
func (_Vester *VesterCaller) BonusRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "bonusRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BonusRewards is a free data retrieval call binding the contract method 0xa2545fa5.
//
// Solidity: function bonusRewards(address ) view returns(uint256)
func (_Vester *VesterSession) BonusRewards(arg0 common.Address) (*big.Int, error) {
	return _Vester.Contract.BonusRewards(&_Vester.CallOpts, arg0)
}

// BonusRewards is a free data retrieval call binding the contract method 0xa2545fa5.
//
// Solidity: function bonusRewards(address ) view returns(uint256)
func (_Vester *VesterCallerSession) BonusRewards(arg0 common.Address) (*big.Int, error) {
	return _Vester.Contract.BonusRewards(&_Vester.CallOpts, arg0)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_Vester *VesterCaller) Claimable(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "claimable", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_Vester *VesterSession) Claimable(_account common.Address) (*big.Int, error) {
	return _Vester.Contract.Claimable(&_Vester.CallOpts, _account)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_Vester *VesterCallerSession) Claimable(_account common.Address) (*big.Int, error) {
	return _Vester.Contract.Claimable(&_Vester.CallOpts, _account)
}

// ClaimableToken is a free data retrieval call binding the contract method 0xf6d6d5aa.
//
// Solidity: function claimableToken() view returns(address)
func (_Vester *VesterCaller) ClaimableToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "claimableToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ClaimableToken is a free data retrieval call binding the contract method 0xf6d6d5aa.
//
// Solidity: function claimableToken() view returns(address)
func (_Vester *VesterSession) ClaimableToken() (common.Address, error) {
	return _Vester.Contract.ClaimableToken(&_Vester.CallOpts)
}

// ClaimableToken is a free data retrieval call binding the contract method 0xf6d6d5aa.
//
// Solidity: function claimableToken() view returns(address)
func (_Vester *VesterCallerSession) ClaimableToken() (common.Address, error) {
	return _Vester.Contract.ClaimableToken(&_Vester.CallOpts)
}

// ClaimedAmounts is a free data retrieval call binding the contract method 0x71417b32.
//
// Solidity: function claimedAmounts(address ) view returns(uint256)
func (_Vester *VesterCaller) ClaimedAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "claimedAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedAmounts is a free data retrieval call binding the contract method 0x71417b32.
//
// Solidity: function claimedAmounts(address ) view returns(uint256)
func (_Vester *VesterSession) ClaimedAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vester.Contract.ClaimedAmounts(&_Vester.CallOpts, arg0)
}

// ClaimedAmounts is a free data retrieval call binding the contract method 0x71417b32.
//
// Solidity: function claimedAmounts(address ) view returns(uint256)
func (_Vester *VesterCallerSession) ClaimedAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vester.Contract.ClaimedAmounts(&_Vester.CallOpts, arg0)
}

// CumulativeClaimAmounts is a free data retrieval call binding the contract method 0xb5ff136d.
//
// Solidity: function cumulativeClaimAmounts(address ) view returns(uint256)
func (_Vester *VesterCaller) CumulativeClaimAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "cumulativeClaimAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeClaimAmounts is a free data retrieval call binding the contract method 0xb5ff136d.
//
// Solidity: function cumulativeClaimAmounts(address ) view returns(uint256)
func (_Vester *VesterSession) CumulativeClaimAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vester.Contract.CumulativeClaimAmounts(&_Vester.CallOpts, arg0)
}

// CumulativeClaimAmounts is a free data retrieval call binding the contract method 0xb5ff136d.
//
// Solidity: function cumulativeClaimAmounts(address ) view returns(uint256)
func (_Vester *VesterCallerSession) CumulativeClaimAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vester.Contract.CumulativeClaimAmounts(&_Vester.CallOpts, arg0)
}

// CumulativeRewardDeductions is a free data retrieval call binding the contract method 0x387a785d.
//
// Solidity: function cumulativeRewardDeductions(address ) view returns(uint256)
func (_Vester *VesterCaller) CumulativeRewardDeductions(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "cumulativeRewardDeductions", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeRewardDeductions is a free data retrieval call binding the contract method 0x387a785d.
//
// Solidity: function cumulativeRewardDeductions(address ) view returns(uint256)
func (_Vester *VesterSession) CumulativeRewardDeductions(arg0 common.Address) (*big.Int, error) {
	return _Vester.Contract.CumulativeRewardDeductions(&_Vester.CallOpts, arg0)
}

// CumulativeRewardDeductions is a free data retrieval call binding the contract method 0x387a785d.
//
// Solidity: function cumulativeRewardDeductions(address ) view returns(uint256)
func (_Vester *VesterCallerSession) CumulativeRewardDeductions(arg0 common.Address) (*big.Int, error) {
	return _Vester.Contract.CumulativeRewardDeductions(&_Vester.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vester *VesterCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vester *VesterSession) Decimals() (uint8, error) {
	return _Vester.Contract.Decimals(&_Vester.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Vester *VesterCallerSession) Decimals() (uint8, error) {
	return _Vester.Contract.Decimals(&_Vester.CallOpts)
}

// EsToken is a free data retrieval call binding the contract method 0x16ca05c5.
//
// Solidity: function esToken() view returns(address)
func (_Vester *VesterCaller) EsToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "esToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EsToken is a free data retrieval call binding the contract method 0x16ca05c5.
//
// Solidity: function esToken() view returns(address)
func (_Vester *VesterSession) EsToken() (common.Address, error) {
	return _Vester.Contract.EsToken(&_Vester.CallOpts)
}

// EsToken is a free data retrieval call binding the contract method 0x16ca05c5.
//
// Solidity: function esToken() view returns(address)
func (_Vester *VesterCallerSession) EsToken() (common.Address, error) {
	return _Vester.Contract.EsToken(&_Vester.CallOpts)
}

// GetCombinedAverageStakedAmount is a free data retrieval call binding the contract method 0x45f01ee6.
//
// Solidity: function getCombinedAverageStakedAmount(address _account) view returns(uint256)
func (_Vester *VesterCaller) GetCombinedAverageStakedAmount(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "getCombinedAverageStakedAmount", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCombinedAverageStakedAmount is a free data retrieval call binding the contract method 0x45f01ee6.
//
// Solidity: function getCombinedAverageStakedAmount(address _account) view returns(uint256)
func (_Vester *VesterSession) GetCombinedAverageStakedAmount(_account common.Address) (*big.Int, error) {
	return _Vester.Contract.GetCombinedAverageStakedAmount(&_Vester.CallOpts, _account)
}

// GetCombinedAverageStakedAmount is a free data retrieval call binding the contract method 0x45f01ee6.
//
// Solidity: function getCombinedAverageStakedAmount(address _account) view returns(uint256)
func (_Vester *VesterCallerSession) GetCombinedAverageStakedAmount(_account common.Address) (*big.Int, error) {
	return _Vester.Contract.GetCombinedAverageStakedAmount(&_Vester.CallOpts, _account)
}

// GetMaxVestableAmount is a free data retrieval call binding the contract method 0x08f26c76.
//
// Solidity: function getMaxVestableAmount(address _account) view returns(uint256)
func (_Vester *VesterCaller) GetMaxVestableAmount(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "getMaxVestableAmount", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxVestableAmount is a free data retrieval call binding the contract method 0x08f26c76.
//
// Solidity: function getMaxVestableAmount(address _account) view returns(uint256)
func (_Vester *VesterSession) GetMaxVestableAmount(_account common.Address) (*big.Int, error) {
	return _Vester.Contract.GetMaxVestableAmount(&_Vester.CallOpts, _account)
}

// GetMaxVestableAmount is a free data retrieval call binding the contract method 0x08f26c76.
//
// Solidity: function getMaxVestableAmount(address _account) view returns(uint256)
func (_Vester *VesterCallerSession) GetMaxVestableAmount(_account common.Address) (*big.Int, error) {
	return _Vester.Contract.GetMaxVestableAmount(&_Vester.CallOpts, _account)
}

// GetPairAmount is a free data retrieval call binding the contract method 0x7cf8f3b2.
//
// Solidity: function getPairAmount(address _account, uint256 _esAmount) view returns(uint256)
func (_Vester *VesterCaller) GetPairAmount(opts *bind.CallOpts, _account common.Address, _esAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "getPairAmount", _account, _esAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPairAmount is a free data retrieval call binding the contract method 0x7cf8f3b2.
//
// Solidity: function getPairAmount(address _account, uint256 _esAmount) view returns(uint256)
func (_Vester *VesterSession) GetPairAmount(_account common.Address, _esAmount *big.Int) (*big.Int, error) {
	return _Vester.Contract.GetPairAmount(&_Vester.CallOpts, _account, _esAmount)
}

// GetPairAmount is a free data retrieval call binding the contract method 0x7cf8f3b2.
//
// Solidity: function getPairAmount(address _account, uint256 _esAmount) view returns(uint256)
func (_Vester *VesterCallerSession) GetPairAmount(_account common.Address, _esAmount *big.Int) (*big.Int, error) {
	return _Vester.Contract.GetPairAmount(&_Vester.CallOpts, _account, _esAmount)
}

// GetTotalVested is a free data retrieval call binding the contract method 0x93035473.
//
// Solidity: function getTotalVested(address _account) view returns(uint256)
func (_Vester *VesterCaller) GetTotalVested(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "getTotalVested", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalVested is a free data retrieval call binding the contract method 0x93035473.
//
// Solidity: function getTotalVested(address _account) view returns(uint256)
func (_Vester *VesterSession) GetTotalVested(_account common.Address) (*big.Int, error) {
	return _Vester.Contract.GetTotalVested(&_Vester.CallOpts, _account)
}

// GetTotalVested is a free data retrieval call binding the contract method 0x93035473.
//
// Solidity: function getTotalVested(address _account) view returns(uint256)
func (_Vester *VesterCallerSession) GetTotalVested(_account common.Address) (*big.Int, error) {
	return _Vester.Contract.GetTotalVested(&_Vester.CallOpts, _account)
}

// GetVestedAmount is a free data retrieval call binding the contract method 0xd5a73fdd.
//
// Solidity: function getVestedAmount(address _account) view returns(uint256)
func (_Vester *VesterCaller) GetVestedAmount(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "getVestedAmount", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVestedAmount is a free data retrieval call binding the contract method 0xd5a73fdd.
//
// Solidity: function getVestedAmount(address _account) view returns(uint256)
func (_Vester *VesterSession) GetVestedAmount(_account common.Address) (*big.Int, error) {
	return _Vester.Contract.GetVestedAmount(&_Vester.CallOpts, _account)
}

// GetVestedAmount is a free data retrieval call binding the contract method 0xd5a73fdd.
//
// Solidity: function getVestedAmount(address _account) view returns(uint256)
func (_Vester *VesterCallerSession) GetVestedAmount(_account common.Address) (*big.Int, error) {
	return _Vester.Contract.GetVestedAmount(&_Vester.CallOpts, _account)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Vester *VesterCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Vester *VesterSession) Gov() (common.Address, error) {
	return _Vester.Contract.Gov(&_Vester.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Vester *VesterCallerSession) Gov() (common.Address, error) {
	return _Vester.Contract.Gov(&_Vester.CallOpts)
}

// HasMaxVestableAmount is a free data retrieval call binding the contract method 0xacf077a5.
//
// Solidity: function hasMaxVestableAmount() view returns(bool)
func (_Vester *VesterCaller) HasMaxVestableAmount(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "hasMaxVestableAmount")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMaxVestableAmount is a free data retrieval call binding the contract method 0xacf077a5.
//
// Solidity: function hasMaxVestableAmount() view returns(bool)
func (_Vester *VesterSession) HasMaxVestableAmount() (bool, error) {
	return _Vester.Contract.HasMaxVestableAmount(&_Vester.CallOpts)
}

// HasMaxVestableAmount is a free data retrieval call binding the contract method 0xacf077a5.
//
// Solidity: function hasMaxVestableAmount() view returns(bool)
func (_Vester *VesterCallerSession) HasMaxVestableAmount() (bool, error) {
	return _Vester.Contract.HasMaxVestableAmount(&_Vester.CallOpts)
}

// HasPairToken is a free data retrieval call binding the contract method 0xd75abb57.
//
// Solidity: function hasPairToken() view returns(bool)
func (_Vester *VesterCaller) HasPairToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "hasPairToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasPairToken is a free data retrieval call binding the contract method 0xd75abb57.
//
// Solidity: function hasPairToken() view returns(bool)
func (_Vester *VesterSession) HasPairToken() (bool, error) {
	return _Vester.Contract.HasPairToken(&_Vester.CallOpts)
}

// HasPairToken is a free data retrieval call binding the contract method 0xd75abb57.
//
// Solidity: function hasPairToken() view returns(bool)
func (_Vester *VesterCallerSession) HasPairToken() (bool, error) {
	return _Vester.Contract.HasPairToken(&_Vester.CallOpts)
}

// HasRewardTracker is a free data retrieval call binding the contract method 0xf421f62a.
//
// Solidity: function hasRewardTracker() view returns(bool)
func (_Vester *VesterCaller) HasRewardTracker(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "hasRewardTracker")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRewardTracker is a free data retrieval call binding the contract method 0xf421f62a.
//
// Solidity: function hasRewardTracker() view returns(bool)
func (_Vester *VesterSession) HasRewardTracker() (bool, error) {
	return _Vester.Contract.HasRewardTracker(&_Vester.CallOpts)
}

// HasRewardTracker is a free data retrieval call binding the contract method 0xf421f62a.
//
// Solidity: function hasRewardTracker() view returns(bool)
func (_Vester *VesterCallerSession) HasRewardTracker() (bool, error) {
	return _Vester.Contract.HasRewardTracker(&_Vester.CallOpts)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_Vester *VesterCaller) IsHandler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "isHandler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_Vester *VesterSession) IsHandler(arg0 common.Address) (bool, error) {
	return _Vester.Contract.IsHandler(&_Vester.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_Vester *VesterCallerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _Vester.Contract.IsHandler(&_Vester.CallOpts, arg0)
}

// LastVestingTimes is a free data retrieval call binding the contract method 0x0db9ea4a.
//
// Solidity: function lastVestingTimes(address ) view returns(uint256)
func (_Vester *VesterCaller) LastVestingTimes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "lastVestingTimes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastVestingTimes is a free data retrieval call binding the contract method 0x0db9ea4a.
//
// Solidity: function lastVestingTimes(address ) view returns(uint256)
func (_Vester *VesterSession) LastVestingTimes(arg0 common.Address) (*big.Int, error) {
	return _Vester.Contract.LastVestingTimes(&_Vester.CallOpts, arg0)
}

// LastVestingTimes is a free data retrieval call binding the contract method 0x0db9ea4a.
//
// Solidity: function lastVestingTimes(address ) view returns(uint256)
func (_Vester *VesterCallerSession) LastVestingTimes(arg0 common.Address) (*big.Int, error) {
	return _Vester.Contract.LastVestingTimes(&_Vester.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vester *VesterCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vester *VesterSession) Name() (string, error) {
	return _Vester.Contract.Name(&_Vester.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Vester *VesterCallerSession) Name() (string, error) {
	return _Vester.Contract.Name(&_Vester.CallOpts)
}

// PairAmounts is a free data retrieval call binding the contract method 0x5d50e729.
//
// Solidity: function pairAmounts(address ) view returns(uint256)
func (_Vester *VesterCaller) PairAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "pairAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairAmounts is a free data retrieval call binding the contract method 0x5d50e729.
//
// Solidity: function pairAmounts(address ) view returns(uint256)
func (_Vester *VesterSession) PairAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vester.Contract.PairAmounts(&_Vester.CallOpts, arg0)
}

// PairAmounts is a free data retrieval call binding the contract method 0x5d50e729.
//
// Solidity: function pairAmounts(address ) view returns(uint256)
func (_Vester *VesterCallerSession) PairAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vester.Contract.PairAmounts(&_Vester.CallOpts, arg0)
}

// PairSupply is a free data retrieval call binding the contract method 0x15e90a41.
//
// Solidity: function pairSupply() view returns(uint256)
func (_Vester *VesterCaller) PairSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "pairSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PairSupply is a free data retrieval call binding the contract method 0x15e90a41.
//
// Solidity: function pairSupply() view returns(uint256)
func (_Vester *VesterSession) PairSupply() (*big.Int, error) {
	return _Vester.Contract.PairSupply(&_Vester.CallOpts)
}

// PairSupply is a free data retrieval call binding the contract method 0x15e90a41.
//
// Solidity: function pairSupply() view returns(uint256)
func (_Vester *VesterCallerSession) PairSupply() (*big.Int, error) {
	return _Vester.Contract.PairSupply(&_Vester.CallOpts)
}

// PairToken is a free data retrieval call binding the contract method 0x3de35b79.
//
// Solidity: function pairToken() view returns(address)
func (_Vester *VesterCaller) PairToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "pairToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairToken is a free data retrieval call binding the contract method 0x3de35b79.
//
// Solidity: function pairToken() view returns(address)
func (_Vester *VesterSession) PairToken() (common.Address, error) {
	return _Vester.Contract.PairToken(&_Vester.CallOpts)
}

// PairToken is a free data retrieval call binding the contract method 0x3de35b79.
//
// Solidity: function pairToken() view returns(address)
func (_Vester *VesterCallerSession) PairToken() (common.Address, error) {
	return _Vester.Contract.PairToken(&_Vester.CallOpts)
}

// RewardTracker is a free data retrieval call binding the contract method 0x6bcb411a.
//
// Solidity: function rewardTracker() view returns(address)
func (_Vester *VesterCaller) RewardTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "rewardTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTracker is a free data retrieval call binding the contract method 0x6bcb411a.
//
// Solidity: function rewardTracker() view returns(address)
func (_Vester *VesterSession) RewardTracker() (common.Address, error) {
	return _Vester.Contract.RewardTracker(&_Vester.CallOpts)
}

// RewardTracker is a free data retrieval call binding the contract method 0x6bcb411a.
//
// Solidity: function rewardTracker() view returns(address)
func (_Vester *VesterCallerSession) RewardTracker() (common.Address, error) {
	return _Vester.Contract.RewardTracker(&_Vester.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vester *VesterCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vester *VesterSession) Symbol() (string, error) {
	return _Vester.Contract.Symbol(&_Vester.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Vester *VesterCallerSession) Symbol() (string, error) {
	return _Vester.Contract.Symbol(&_Vester.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vester *VesterCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vester *VesterSession) TotalSupply() (*big.Int, error) {
	return _Vester.Contract.TotalSupply(&_Vester.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Vester *VesterCallerSession) TotalSupply() (*big.Int, error) {
	return _Vester.Contract.TotalSupply(&_Vester.CallOpts)
}

// TransferredAverageStakedAmounts is a free data retrieval call binding the contract method 0x7337035c.
//
// Solidity: function transferredAverageStakedAmounts(address ) view returns(uint256)
func (_Vester *VesterCaller) TransferredAverageStakedAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "transferredAverageStakedAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferredAverageStakedAmounts is a free data retrieval call binding the contract method 0x7337035c.
//
// Solidity: function transferredAverageStakedAmounts(address ) view returns(uint256)
func (_Vester *VesterSession) TransferredAverageStakedAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vester.Contract.TransferredAverageStakedAmounts(&_Vester.CallOpts, arg0)
}

// TransferredAverageStakedAmounts is a free data retrieval call binding the contract method 0x7337035c.
//
// Solidity: function transferredAverageStakedAmounts(address ) view returns(uint256)
func (_Vester *VesterCallerSession) TransferredAverageStakedAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vester.Contract.TransferredAverageStakedAmounts(&_Vester.CallOpts, arg0)
}

// TransferredCumulativeRewards is a free data retrieval call binding the contract method 0xb71bce2a.
//
// Solidity: function transferredCumulativeRewards(address ) view returns(uint256)
func (_Vester *VesterCaller) TransferredCumulativeRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "transferredCumulativeRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferredCumulativeRewards is a free data retrieval call binding the contract method 0xb71bce2a.
//
// Solidity: function transferredCumulativeRewards(address ) view returns(uint256)
func (_Vester *VesterSession) TransferredCumulativeRewards(arg0 common.Address) (*big.Int, error) {
	return _Vester.Contract.TransferredCumulativeRewards(&_Vester.CallOpts, arg0)
}

// TransferredCumulativeRewards is a free data retrieval call binding the contract method 0xb71bce2a.
//
// Solidity: function transferredCumulativeRewards(address ) view returns(uint256)
func (_Vester *VesterCallerSession) TransferredCumulativeRewards(arg0 common.Address) (*big.Int, error) {
	return _Vester.Contract.TransferredCumulativeRewards(&_Vester.CallOpts, arg0)
}

// VestingDuration is a free data retrieval call binding the contract method 0x1514617e.
//
// Solidity: function vestingDuration() view returns(uint256)
func (_Vester *VesterCaller) VestingDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vester.contract.Call(opts, &out, "vestingDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VestingDuration is a free data retrieval call binding the contract method 0x1514617e.
//
// Solidity: function vestingDuration() view returns(uint256)
func (_Vester *VesterSession) VestingDuration() (*big.Int, error) {
	return _Vester.Contract.VestingDuration(&_Vester.CallOpts)
}

// VestingDuration is a free data retrieval call binding the contract method 0x1514617e.
//
// Solidity: function vestingDuration() view returns(uint256)
func (_Vester *VesterCallerSession) VestingDuration() (*big.Int, error) {
	return _Vester.Contract.VestingDuration(&_Vester.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_Vester *VesterTransactor) Approve(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Vester.contract.Transact(opts, "approve", arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_Vester *VesterSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.Approve(&_Vester.TransactOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_Vester *VesterTransactorSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.Approve(&_Vester.TransactOpts, arg0, arg1)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_Vester *VesterTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vester.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_Vester *VesterSession) Claim() (*types.Transaction, error) {
	return _Vester.Contract.Claim(&_Vester.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns(uint256)
func (_Vester *VesterTransactorSession) Claim() (*types.Transaction, error) {
	return _Vester.Contract.Claim(&_Vester.TransactOpts)
}

// ClaimForAccount is a paid mutator transaction binding the contract method 0x13e82e7a.
//
// Solidity: function claimForAccount(address _account, address _receiver) returns(uint256)
func (_Vester *VesterTransactor) ClaimForAccount(opts *bind.TransactOpts, _account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vester.contract.Transact(opts, "claimForAccount", _account, _receiver)
}

// ClaimForAccount is a paid mutator transaction binding the contract method 0x13e82e7a.
//
// Solidity: function claimForAccount(address _account, address _receiver) returns(uint256)
func (_Vester *VesterSession) ClaimForAccount(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vester.Contract.ClaimForAccount(&_Vester.TransactOpts, _account, _receiver)
}

// ClaimForAccount is a paid mutator transaction binding the contract method 0x13e82e7a.
//
// Solidity: function claimForAccount(address _account, address _receiver) returns(uint256)
func (_Vester *VesterTransactorSession) ClaimForAccount(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vester.Contract.ClaimForAccount(&_Vester.TransactOpts, _account, _receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Vester *VesterTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Vester.contract.Transact(opts, "deposit", _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Vester *VesterSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.Deposit(&_Vester.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 _amount) returns()
func (_Vester *VesterTransactorSession) Deposit(_amount *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.Deposit(&_Vester.TransactOpts, _amount)
}

// DepositForAccount is a paid mutator transaction binding the contract method 0x342fcda9.
//
// Solidity: function depositForAccount(address _account, uint256 _amount) returns()
func (_Vester *VesterTransactor) DepositForAccount(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vester.contract.Transact(opts, "depositForAccount", _account, _amount)
}

// DepositForAccount is a paid mutator transaction binding the contract method 0x342fcda9.
//
// Solidity: function depositForAccount(address _account, uint256 _amount) returns()
func (_Vester *VesterSession) DepositForAccount(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.DepositForAccount(&_Vester.TransactOpts, _account, _amount)
}

// DepositForAccount is a paid mutator transaction binding the contract method 0x342fcda9.
//
// Solidity: function depositForAccount(address _account, uint256 _amount) returns()
func (_Vester *VesterTransactorSession) DepositForAccount(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.DepositForAccount(&_Vester.TransactOpts, _account, _amount)
}

// SetBonusRewards is a paid mutator transaction binding the contract method 0x41f22724.
//
// Solidity: function setBonusRewards(address _account, uint256 _amount) returns()
func (_Vester *VesterTransactor) SetBonusRewards(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vester.contract.Transact(opts, "setBonusRewards", _account, _amount)
}

// SetBonusRewards is a paid mutator transaction binding the contract method 0x41f22724.
//
// Solidity: function setBonusRewards(address _account, uint256 _amount) returns()
func (_Vester *VesterSession) SetBonusRewards(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.SetBonusRewards(&_Vester.TransactOpts, _account, _amount)
}

// SetBonusRewards is a paid mutator transaction binding the contract method 0x41f22724.
//
// Solidity: function setBonusRewards(address _account, uint256 _amount) returns()
func (_Vester *VesterTransactorSession) SetBonusRewards(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.SetBonusRewards(&_Vester.TransactOpts, _account, _amount)
}

// SetCumulativeRewardDeductions is a paid mutator transaction binding the contract method 0xd89b7007.
//
// Solidity: function setCumulativeRewardDeductions(address _account, uint256 _amount) returns()
func (_Vester *VesterTransactor) SetCumulativeRewardDeductions(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vester.contract.Transact(opts, "setCumulativeRewardDeductions", _account, _amount)
}

// SetCumulativeRewardDeductions is a paid mutator transaction binding the contract method 0xd89b7007.
//
// Solidity: function setCumulativeRewardDeductions(address _account, uint256 _amount) returns()
func (_Vester *VesterSession) SetCumulativeRewardDeductions(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.SetCumulativeRewardDeductions(&_Vester.TransactOpts, _account, _amount)
}

// SetCumulativeRewardDeductions is a paid mutator transaction binding the contract method 0xd89b7007.
//
// Solidity: function setCumulativeRewardDeductions(address _account, uint256 _amount) returns()
func (_Vester *VesterTransactorSession) SetCumulativeRewardDeductions(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.SetCumulativeRewardDeductions(&_Vester.TransactOpts, _account, _amount)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Vester *VesterTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _Vester.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Vester *VesterSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Vester.Contract.SetGov(&_Vester.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Vester *VesterTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Vester.Contract.SetGov(&_Vester.TransactOpts, _gov)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_Vester *VesterTransactor) SetHandler(opts *bind.TransactOpts, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _Vester.contract.Transact(opts, "setHandler", _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_Vester *VesterSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _Vester.Contract.SetHandler(&_Vester.TransactOpts, _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_Vester *VesterTransactorSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _Vester.Contract.SetHandler(&_Vester.TransactOpts, _handler, _isActive)
}

// SetHasMaxVestableAmount is a paid mutator transaction binding the contract method 0x69de9b93.
//
// Solidity: function setHasMaxVestableAmount(bool _hasMaxVestableAmount) returns()
func (_Vester *VesterTransactor) SetHasMaxVestableAmount(opts *bind.TransactOpts, _hasMaxVestableAmount bool) (*types.Transaction, error) {
	return _Vester.contract.Transact(opts, "setHasMaxVestableAmount", _hasMaxVestableAmount)
}

// SetHasMaxVestableAmount is a paid mutator transaction binding the contract method 0x69de9b93.
//
// Solidity: function setHasMaxVestableAmount(bool _hasMaxVestableAmount) returns()
func (_Vester *VesterSession) SetHasMaxVestableAmount(_hasMaxVestableAmount bool) (*types.Transaction, error) {
	return _Vester.Contract.SetHasMaxVestableAmount(&_Vester.TransactOpts, _hasMaxVestableAmount)
}

// SetHasMaxVestableAmount is a paid mutator transaction binding the contract method 0x69de9b93.
//
// Solidity: function setHasMaxVestableAmount(bool _hasMaxVestableAmount) returns()
func (_Vester *VesterTransactorSession) SetHasMaxVestableAmount(_hasMaxVestableAmount bool) (*types.Transaction, error) {
	return _Vester.Contract.SetHasMaxVestableAmount(&_Vester.TransactOpts, _hasMaxVestableAmount)
}

// SetTransferredAverageStakedAmounts is a paid mutator transaction binding the contract method 0xe3ecc4b2.
//
// Solidity: function setTransferredAverageStakedAmounts(address _account, uint256 _amount) returns()
func (_Vester *VesterTransactor) SetTransferredAverageStakedAmounts(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vester.contract.Transact(opts, "setTransferredAverageStakedAmounts", _account, _amount)
}

// SetTransferredAverageStakedAmounts is a paid mutator transaction binding the contract method 0xe3ecc4b2.
//
// Solidity: function setTransferredAverageStakedAmounts(address _account, uint256 _amount) returns()
func (_Vester *VesterSession) SetTransferredAverageStakedAmounts(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.SetTransferredAverageStakedAmounts(&_Vester.TransactOpts, _account, _amount)
}

// SetTransferredAverageStakedAmounts is a paid mutator transaction binding the contract method 0xe3ecc4b2.
//
// Solidity: function setTransferredAverageStakedAmounts(address _account, uint256 _amount) returns()
func (_Vester *VesterTransactorSession) SetTransferredAverageStakedAmounts(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.SetTransferredAverageStakedAmounts(&_Vester.TransactOpts, _account, _amount)
}

// SetTransferredCumulativeRewards is a paid mutator transaction binding the contract method 0xd0b038b7.
//
// Solidity: function setTransferredCumulativeRewards(address _account, uint256 _amount) returns()
func (_Vester *VesterTransactor) SetTransferredCumulativeRewards(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vester.contract.Transact(opts, "setTransferredCumulativeRewards", _account, _amount)
}

// SetTransferredCumulativeRewards is a paid mutator transaction binding the contract method 0xd0b038b7.
//
// Solidity: function setTransferredCumulativeRewards(address _account, uint256 _amount) returns()
func (_Vester *VesterSession) SetTransferredCumulativeRewards(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.SetTransferredCumulativeRewards(&_Vester.TransactOpts, _account, _amount)
}

// SetTransferredCumulativeRewards is a paid mutator transaction binding the contract method 0xd0b038b7.
//
// Solidity: function setTransferredCumulativeRewards(address _account, uint256 _amount) returns()
func (_Vester *VesterTransactorSession) SetTransferredCumulativeRewards(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.SetTransferredCumulativeRewards(&_Vester.TransactOpts, _account, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_Vester *VesterTransactor) Transfer(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Vester.contract.Transact(opts, "transfer", arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_Vester *VesterSession) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.Transfer(&_Vester.TransactOpts, arg0, arg1)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address , uint256 ) returns(bool)
func (_Vester *VesterTransactorSession) Transfer(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.Transfer(&_Vester.TransactOpts, arg0, arg1)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_Vester *VesterTransactor) TransferFrom(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Vester.contract.Transact(opts, "transferFrom", arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_Vester *VesterSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.TransferFrom(&_Vester.TransactOpts, arg0, arg1, arg2)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address , address , uint256 ) returns(bool)
func (_Vester *VesterTransactorSession) TransferFrom(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.TransferFrom(&_Vester.TransactOpts, arg0, arg1, arg2)
}

// TransferStakeValues is a paid mutator transaction binding the contract method 0xf713c230.
//
// Solidity: function transferStakeValues(address _sender, address _receiver) returns()
func (_Vester *VesterTransactor) TransferStakeValues(opts *bind.TransactOpts, _sender common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vester.contract.Transact(opts, "transferStakeValues", _sender, _receiver)
}

// TransferStakeValues is a paid mutator transaction binding the contract method 0xf713c230.
//
// Solidity: function transferStakeValues(address _sender, address _receiver) returns()
func (_Vester *VesterSession) TransferStakeValues(_sender common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vester.Contract.TransferStakeValues(&_Vester.TransactOpts, _sender, _receiver)
}

// TransferStakeValues is a paid mutator transaction binding the contract method 0xf713c230.
//
// Solidity: function transferStakeValues(address _sender, address _receiver) returns()
func (_Vester *VesterTransactorSession) TransferStakeValues(_sender common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vester.Contract.TransferStakeValues(&_Vester.TransactOpts, _sender, _receiver)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Vester *VesterTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vester.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Vester *VesterSession) Withdraw() (*types.Transaction, error) {
	return _Vester.Contract.Withdraw(&_Vester.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Vester *VesterTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Vester.Contract.Withdraw(&_Vester.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_Vester *VesterTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vester.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_Vester *VesterSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.WithdrawToken(&_Vester.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_Vester *VesterTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vester.Contract.WithdrawToken(&_Vester.TransactOpts, _token, _account, _amount)
}

// VesterApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Vester contract.
type VesterApprovalIterator struct {
	Event *VesterApproval // Event containing the contract specifics and raw log

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
func (it *VesterApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VesterApproval)
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
		it.Event = new(VesterApproval)
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
func (it *VesterApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VesterApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VesterApproval represents a Approval event raised by the Vester contract.
type VesterApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Vester *VesterFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*VesterApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vester.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VesterApprovalIterator{contract: _Vester.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Vester *VesterFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VesterApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vester.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VesterApproval)
				if err := _Vester.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Vester *VesterFilterer) ParseApproval(log types.Log) (*VesterApproval, error) {
	event := new(VesterApproval)
	if err := _Vester.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VesterClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Vester contract.
type VesterClaimIterator struct {
	Event *VesterClaim // Event containing the contract specifics and raw log

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
func (it *VesterClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VesterClaim)
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
		it.Event = new(VesterClaim)
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
func (it *VesterClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VesterClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VesterClaim represents a Claim event raised by the Vester contract.
type VesterClaim struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address receiver, uint256 amount)
func (_Vester *VesterFilterer) FilterClaim(opts *bind.FilterOpts) (*VesterClaimIterator, error) {

	logs, sub, err := _Vester.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &VesterClaimIterator{contract: _Vester.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address receiver, uint256 amount)
func (_Vester *VesterFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *VesterClaim) (event.Subscription, error) {

	logs, sub, err := _Vester.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VesterClaim)
				if err := _Vester.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address receiver, uint256 amount)
func (_Vester *VesterFilterer) ParseClaim(log types.Log) (*VesterClaim, error) {
	event := new(VesterClaim)
	if err := _Vester.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VesterDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Vester contract.
type VesterDepositIterator struct {
	Event *VesterDeposit // Event containing the contract specifics and raw log

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
func (it *VesterDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VesterDeposit)
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
		it.Event = new(VesterDeposit)
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
func (it *VesterDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VesterDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VesterDeposit represents a Deposit event raised by the Vester contract.
type VesterDeposit struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address account, uint256 amount)
func (_Vester *VesterFilterer) FilterDeposit(opts *bind.FilterOpts) (*VesterDepositIterator, error) {

	logs, sub, err := _Vester.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &VesterDepositIterator{contract: _Vester.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address account, uint256 amount)
func (_Vester *VesterFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *VesterDeposit) (event.Subscription, error) {

	logs, sub, err := _Vester.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VesterDeposit)
				if err := _Vester.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address account, uint256 amount)
func (_Vester *VesterFilterer) ParseDeposit(log types.Log) (*VesterDeposit, error) {
	event := new(VesterDeposit)
	if err := _Vester.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VesterPairTransferIterator is returned from FilterPairTransfer and is used to iterate over the raw logs and unpacked data for PairTransfer events raised by the Vester contract.
type VesterPairTransferIterator struct {
	Event *VesterPairTransfer // Event containing the contract specifics and raw log

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
func (it *VesterPairTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VesterPairTransfer)
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
		it.Event = new(VesterPairTransfer)
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
func (it *VesterPairTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VesterPairTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VesterPairTransfer represents a PairTransfer event raised by the Vester contract.
type VesterPairTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPairTransfer is a free log retrieval operation binding the contract event 0x659523c479d006050ebc0d0e48fea36d1b2c5d45b2f31402ac6f8671fc84cc04.
//
// Solidity: event PairTransfer(address indexed from, address indexed to, uint256 value)
func (_Vester *VesterFilterer) FilterPairTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VesterPairTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vester.contract.FilterLogs(opts, "PairTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VesterPairTransferIterator{contract: _Vester.contract, event: "PairTransfer", logs: logs, sub: sub}, nil
}

// WatchPairTransfer is a free log subscription operation binding the contract event 0x659523c479d006050ebc0d0e48fea36d1b2c5d45b2f31402ac6f8671fc84cc04.
//
// Solidity: event PairTransfer(address indexed from, address indexed to, uint256 value)
func (_Vester *VesterFilterer) WatchPairTransfer(opts *bind.WatchOpts, sink chan<- *VesterPairTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vester.contract.WatchLogs(opts, "PairTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VesterPairTransfer)
				if err := _Vester.contract.UnpackLog(event, "PairTransfer", log); err != nil {
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

// ParsePairTransfer is a log parse operation binding the contract event 0x659523c479d006050ebc0d0e48fea36d1b2c5d45b2f31402ac6f8671fc84cc04.
//
// Solidity: event PairTransfer(address indexed from, address indexed to, uint256 value)
func (_Vester *VesterFilterer) ParsePairTransfer(log types.Log) (*VesterPairTransfer, error) {
	event := new(VesterPairTransfer)
	if err := _Vester.contract.UnpackLog(event, "PairTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VesterTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Vester contract.
type VesterTransferIterator struct {
	Event *VesterTransfer // Event containing the contract specifics and raw log

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
func (it *VesterTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VesterTransfer)
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
		it.Event = new(VesterTransfer)
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
func (it *VesterTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VesterTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VesterTransfer represents a Transfer event raised by the Vester contract.
type VesterTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Vester *VesterFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VesterTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vester.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VesterTransferIterator{contract: _Vester.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Vester *VesterFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VesterTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vester.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VesterTransfer)
				if err := _Vester.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Vester *VesterFilterer) ParseTransfer(log types.Log) (*VesterTransfer, error) {
	event := new(VesterTransfer)
	if err := _Vester.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VesterWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Vester contract.
type VesterWithdrawIterator struct {
	Event *VesterWithdraw // Event containing the contract specifics and raw log

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
func (it *VesterWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VesterWithdraw)
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
		it.Event = new(VesterWithdraw)
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
func (it *VesterWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VesterWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VesterWithdraw represents a Withdraw event raised by the Vester contract.
type VesterWithdraw struct {
	Account       common.Address
	ClaimedAmount *big.Int
	Balance       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address account, uint256 claimedAmount, uint256 balance)
func (_Vester *VesterFilterer) FilterWithdraw(opts *bind.FilterOpts) (*VesterWithdrawIterator, error) {

	logs, sub, err := _Vester.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &VesterWithdrawIterator{contract: _Vester.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address account, uint256 claimedAmount, uint256 balance)
func (_Vester *VesterFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *VesterWithdraw) (event.Subscription, error) {

	logs, sub, err := _Vester.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VesterWithdraw)
				if err := _Vester.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xf279e6a1f5e320cca91135676d9cb6e44ca8a08c0b88342bcdb1144f6511b568.
//
// Solidity: event Withdraw(address account, uint256 claimedAmount, uint256 balance)
func (_Vester *VesterFilterer) ParseWithdraw(log types.Log) (*VesterWithdraw, error) {
	event := new(VesterWithdraw)
	if err := _Vester.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

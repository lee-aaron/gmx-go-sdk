// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RewardTracker

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

// RewardTrackerMetaData contains all meta data concerning the RewardTracker contract.
var RewardTrackerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BASIS_POINTS_DIVISOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"averageStakedAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimForAccount\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimable\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimableReward\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cumulativeRewardPerToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cumulativeRewards\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositBalances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"distributor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inPrivateClaimingMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inPrivateStakingMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inPrivateTransferMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_depositTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_distributor\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDepositToken\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isHandler\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isInitialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previousCumulatedRewardPerToken\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rewardToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setDepositToken\",\"inputs\":[{\"name\":\"_depositToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isDepositToken\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setHandler\",\"inputs\":[{\"name\":\"_handler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInPrivateClaimingMode\",\"inputs\":[{\"name\":\"_inPrivateClaimingMode\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInPrivateStakingMode\",\"inputs\":[{\"name\":\"_inPrivateStakingMode\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInPrivateTransferMode\",\"inputs\":[{\"name\":\"_inPrivateTransferMode\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stake\",\"inputs\":[{\"name\":\"_depositToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeForAccount\",\"inputs\":[{\"name\":\"_fundingAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_depositToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakedAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokensPerInterval\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalDepositSupply\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstake\",\"inputs\":[{\"name\":\"_depositToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstakeForAccount\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_depositToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateRewards\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claim\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// RewardTrackerABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardTrackerMetaData.ABI instead.
var RewardTrackerABI = RewardTrackerMetaData.ABI

// RewardTracker is an auto generated Go binding around an Ethereum contract.
type RewardTracker struct {
	RewardTrackerCaller     // Read-only binding to the contract
	RewardTrackerTransactor // Write-only binding to the contract
	RewardTrackerFilterer   // Log filterer for contract events
}

// RewardTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardTrackerSession struct {
	Contract     *RewardTracker    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardTrackerCallerSession struct {
	Contract *RewardTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RewardTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardTrackerTransactorSession struct {
	Contract     *RewardTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RewardTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardTrackerRaw struct {
	Contract *RewardTracker // Generic contract binding to access the raw methods on
}

// RewardTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardTrackerCallerRaw struct {
	Contract *RewardTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// RewardTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardTrackerTransactorRaw struct {
	Contract *RewardTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardTracker creates a new instance of RewardTracker, bound to a specific deployed contract.
func NewRewardTracker(address common.Address, backend bind.ContractBackend) (*RewardTracker, error) {
	contract, err := bindRewardTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardTracker{RewardTrackerCaller: RewardTrackerCaller{contract: contract}, RewardTrackerTransactor: RewardTrackerTransactor{contract: contract}, RewardTrackerFilterer: RewardTrackerFilterer{contract: contract}}, nil
}

// NewRewardTrackerCaller creates a new read-only instance of RewardTracker, bound to a specific deployed contract.
func NewRewardTrackerCaller(address common.Address, caller bind.ContractCaller) (*RewardTrackerCaller, error) {
	contract, err := bindRewardTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardTrackerCaller{contract: contract}, nil
}

// NewRewardTrackerTransactor creates a new write-only instance of RewardTracker, bound to a specific deployed contract.
func NewRewardTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardTrackerTransactor, error) {
	contract, err := bindRewardTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardTrackerTransactor{contract: contract}, nil
}

// NewRewardTrackerFilterer creates a new log filterer instance of RewardTracker, bound to a specific deployed contract.
func NewRewardTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardTrackerFilterer, error) {
	contract, err := bindRewardTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardTrackerFilterer{contract: contract}, nil
}

// bindRewardTracker binds a generic wrapper to an already deployed contract.
func bindRewardTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RewardTrackerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardTracker *RewardTrackerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardTracker.Contract.RewardTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardTracker *RewardTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardTracker.Contract.RewardTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardTracker *RewardTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardTracker.Contract.RewardTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardTracker *RewardTrackerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardTracker *RewardTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardTracker *RewardTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardTracker.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_RewardTracker *RewardTrackerCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_RewardTracker *RewardTrackerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _RewardTracker.Contract.BASISPOINTSDIVISOR(&_RewardTracker.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_RewardTracker *RewardTrackerCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _RewardTracker.Contract.BASISPOINTSDIVISOR(&_RewardTracker.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_RewardTracker *RewardTrackerCaller) PRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_RewardTracker *RewardTrackerSession) PRECISION() (*big.Int, error) {
	return _RewardTracker.Contract.PRECISION(&_RewardTracker.CallOpts)
}

// PRECISION is a free data retrieval call binding the contract method 0xaaf5eb68.
//
// Solidity: function PRECISION() view returns(uint256)
func (_RewardTracker *RewardTrackerCallerSession) PRECISION() (*big.Int, error) {
	return _RewardTracker.Contract.PRECISION(&_RewardTracker.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_RewardTracker *RewardTrackerCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_RewardTracker *RewardTrackerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.Allowance(&_RewardTracker.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_RewardTracker *RewardTrackerCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.Allowance(&_RewardTracker.CallOpts, _owner, _spender)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_RewardTracker *RewardTrackerCaller) Allowances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "allowances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_RewardTracker *RewardTrackerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.Allowances(&_RewardTracker.CallOpts, arg0, arg1)
}

// Allowances is a free data retrieval call binding the contract method 0x55b6ed5c.
//
// Solidity: function allowances(address , address ) view returns(uint256)
func (_RewardTracker *RewardTrackerCallerSession) Allowances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.Allowances(&_RewardTracker.CallOpts, arg0, arg1)
}

// AverageStakedAmounts is a free data retrieval call binding the contract method 0xa3180217.
//
// Solidity: function averageStakedAmounts(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerCaller) AverageStakedAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "averageStakedAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AverageStakedAmounts is a free data retrieval call binding the contract method 0xa3180217.
//
// Solidity: function averageStakedAmounts(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerSession) AverageStakedAmounts(arg0 common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.AverageStakedAmounts(&_RewardTracker.CallOpts, arg0)
}

// AverageStakedAmounts is a free data retrieval call binding the contract method 0xa3180217.
//
// Solidity: function averageStakedAmounts(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerCallerSession) AverageStakedAmounts(arg0 common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.AverageStakedAmounts(&_RewardTracker.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_RewardTracker *RewardTrackerCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "balanceOf", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_RewardTracker *RewardTrackerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.BalanceOf(&_RewardTracker.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) view returns(uint256)
func (_RewardTracker *RewardTrackerCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.BalanceOf(&_RewardTracker.CallOpts, _account)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.Balances(&_RewardTracker.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.Balances(&_RewardTracker.CallOpts, arg0)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_RewardTracker *RewardTrackerCaller) Claimable(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "claimable", _account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_RewardTracker *RewardTrackerSession) Claimable(_account common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.Claimable(&_RewardTracker.CallOpts, _account)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address _account) view returns(uint256)
func (_RewardTracker *RewardTrackerCallerSession) Claimable(_account common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.Claimable(&_RewardTracker.CallOpts, _account)
}

// ClaimableReward is a free data retrieval call binding the contract method 0xe9503425.
//
// Solidity: function claimableReward(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerCaller) ClaimableReward(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "claimableReward", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableReward is a free data retrieval call binding the contract method 0xe9503425.
//
// Solidity: function claimableReward(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerSession) ClaimableReward(arg0 common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.ClaimableReward(&_RewardTracker.CallOpts, arg0)
}

// ClaimableReward is a free data retrieval call binding the contract method 0xe9503425.
//
// Solidity: function claimableReward(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerCallerSession) ClaimableReward(arg0 common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.ClaimableReward(&_RewardTracker.CallOpts, arg0)
}

// CumulativeRewardPerToken is a free data retrieval call binding the contract method 0xf5fc5076.
//
// Solidity: function cumulativeRewardPerToken() view returns(uint256)
func (_RewardTracker *RewardTrackerCaller) CumulativeRewardPerToken(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "cumulativeRewardPerToken")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeRewardPerToken is a free data retrieval call binding the contract method 0xf5fc5076.
//
// Solidity: function cumulativeRewardPerToken() view returns(uint256)
func (_RewardTracker *RewardTrackerSession) CumulativeRewardPerToken() (*big.Int, error) {
	return _RewardTracker.Contract.CumulativeRewardPerToken(&_RewardTracker.CallOpts)
}

// CumulativeRewardPerToken is a free data retrieval call binding the contract method 0xf5fc5076.
//
// Solidity: function cumulativeRewardPerToken() view returns(uint256)
func (_RewardTracker *RewardTrackerCallerSession) CumulativeRewardPerToken() (*big.Int, error) {
	return _RewardTracker.Contract.CumulativeRewardPerToken(&_RewardTracker.CallOpts)
}

// CumulativeRewards is a free data retrieval call binding the contract method 0x3792def3.
//
// Solidity: function cumulativeRewards(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerCaller) CumulativeRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "cumulativeRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeRewards is a free data retrieval call binding the contract method 0x3792def3.
//
// Solidity: function cumulativeRewards(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerSession) CumulativeRewards(arg0 common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.CumulativeRewards(&_RewardTracker.CallOpts, arg0)
}

// CumulativeRewards is a free data retrieval call binding the contract method 0x3792def3.
//
// Solidity: function cumulativeRewards(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerCallerSession) CumulativeRewards(arg0 common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.CumulativeRewards(&_RewardTracker.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RewardTracker *RewardTrackerCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RewardTracker *RewardTrackerSession) Decimals() (uint8, error) {
	return _RewardTracker.Contract.Decimals(&_RewardTracker.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_RewardTracker *RewardTrackerCallerSession) Decimals() (uint8, error) {
	return _RewardTracker.Contract.Decimals(&_RewardTracker.CallOpts)
}

// DepositBalances is a free data retrieval call binding the contract method 0xf5d9d63e.
//
// Solidity: function depositBalances(address , address ) view returns(uint256)
func (_RewardTracker *RewardTrackerCaller) DepositBalances(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "depositBalances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositBalances is a free data retrieval call binding the contract method 0xf5d9d63e.
//
// Solidity: function depositBalances(address , address ) view returns(uint256)
func (_RewardTracker *RewardTrackerSession) DepositBalances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.DepositBalances(&_RewardTracker.CallOpts, arg0, arg1)
}

// DepositBalances is a free data retrieval call binding the contract method 0xf5d9d63e.
//
// Solidity: function depositBalances(address , address ) view returns(uint256)
func (_RewardTracker *RewardTrackerCallerSession) DepositBalances(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.DepositBalances(&_RewardTracker.CallOpts, arg0, arg1)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_RewardTracker *RewardTrackerCaller) Distributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "distributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_RewardTracker *RewardTrackerSession) Distributor() (common.Address, error) {
	return _RewardTracker.Contract.Distributor(&_RewardTracker.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() view returns(address)
func (_RewardTracker *RewardTrackerCallerSession) Distributor() (common.Address, error) {
	return _RewardTracker.Contract.Distributor(&_RewardTracker.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_RewardTracker *RewardTrackerCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_RewardTracker *RewardTrackerSession) Gov() (common.Address, error) {
	return _RewardTracker.Contract.Gov(&_RewardTracker.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_RewardTracker *RewardTrackerCallerSession) Gov() (common.Address, error) {
	return _RewardTracker.Contract.Gov(&_RewardTracker.CallOpts)
}

// InPrivateClaimingMode is a free data retrieval call binding the contract method 0xf76033d3.
//
// Solidity: function inPrivateClaimingMode() view returns(bool)
func (_RewardTracker *RewardTrackerCaller) InPrivateClaimingMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "inPrivateClaimingMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InPrivateClaimingMode is a free data retrieval call binding the contract method 0xf76033d3.
//
// Solidity: function inPrivateClaimingMode() view returns(bool)
func (_RewardTracker *RewardTrackerSession) InPrivateClaimingMode() (bool, error) {
	return _RewardTracker.Contract.InPrivateClaimingMode(&_RewardTracker.CallOpts)
}

// InPrivateClaimingMode is a free data retrieval call binding the contract method 0xf76033d3.
//
// Solidity: function inPrivateClaimingMode() view returns(bool)
func (_RewardTracker *RewardTrackerCallerSession) InPrivateClaimingMode() (bool, error) {
	return _RewardTracker.Contract.InPrivateClaimingMode(&_RewardTracker.CallOpts)
}

// InPrivateStakingMode is a free data retrieval call binding the contract method 0xc5fa2730.
//
// Solidity: function inPrivateStakingMode() view returns(bool)
func (_RewardTracker *RewardTrackerCaller) InPrivateStakingMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "inPrivateStakingMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InPrivateStakingMode is a free data retrieval call binding the contract method 0xc5fa2730.
//
// Solidity: function inPrivateStakingMode() view returns(bool)
func (_RewardTracker *RewardTrackerSession) InPrivateStakingMode() (bool, error) {
	return _RewardTracker.Contract.InPrivateStakingMode(&_RewardTracker.CallOpts)
}

// InPrivateStakingMode is a free data retrieval call binding the contract method 0xc5fa2730.
//
// Solidity: function inPrivateStakingMode() view returns(bool)
func (_RewardTracker *RewardTrackerCallerSession) InPrivateStakingMode() (bool, error) {
	return _RewardTracker.Contract.InPrivateStakingMode(&_RewardTracker.CallOpts)
}

// InPrivateTransferMode is a free data retrieval call binding the contract method 0xdfbaefb1.
//
// Solidity: function inPrivateTransferMode() view returns(bool)
func (_RewardTracker *RewardTrackerCaller) InPrivateTransferMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "inPrivateTransferMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InPrivateTransferMode is a free data retrieval call binding the contract method 0xdfbaefb1.
//
// Solidity: function inPrivateTransferMode() view returns(bool)
func (_RewardTracker *RewardTrackerSession) InPrivateTransferMode() (bool, error) {
	return _RewardTracker.Contract.InPrivateTransferMode(&_RewardTracker.CallOpts)
}

// InPrivateTransferMode is a free data retrieval call binding the contract method 0xdfbaefb1.
//
// Solidity: function inPrivateTransferMode() view returns(bool)
func (_RewardTracker *RewardTrackerCallerSession) InPrivateTransferMode() (bool, error) {
	return _RewardTracker.Contract.InPrivateTransferMode(&_RewardTracker.CallOpts)
}

// IsDepositToken is a free data retrieval call binding the contract method 0xb89e45b3.
//
// Solidity: function isDepositToken(address ) view returns(bool)
func (_RewardTracker *RewardTrackerCaller) IsDepositToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "isDepositToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDepositToken is a free data retrieval call binding the contract method 0xb89e45b3.
//
// Solidity: function isDepositToken(address ) view returns(bool)
func (_RewardTracker *RewardTrackerSession) IsDepositToken(arg0 common.Address) (bool, error) {
	return _RewardTracker.Contract.IsDepositToken(&_RewardTracker.CallOpts, arg0)
}

// IsDepositToken is a free data retrieval call binding the contract method 0xb89e45b3.
//
// Solidity: function isDepositToken(address ) view returns(bool)
func (_RewardTracker *RewardTrackerCallerSession) IsDepositToken(arg0 common.Address) (bool, error) {
	return _RewardTracker.Contract.IsDepositToken(&_RewardTracker.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_RewardTracker *RewardTrackerCaller) IsHandler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "isHandler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_RewardTracker *RewardTrackerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _RewardTracker.Contract.IsHandler(&_RewardTracker.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_RewardTracker *RewardTrackerCallerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _RewardTracker.Contract.IsHandler(&_RewardTracker.CallOpts, arg0)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_RewardTracker *RewardTrackerCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_RewardTracker *RewardTrackerSession) IsInitialized() (bool, error) {
	return _RewardTracker.Contract.IsInitialized(&_RewardTracker.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_RewardTracker *RewardTrackerCallerSession) IsInitialized() (bool, error) {
	return _RewardTracker.Contract.IsInitialized(&_RewardTracker.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RewardTracker *RewardTrackerCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RewardTracker *RewardTrackerSession) Name() (string, error) {
	return _RewardTracker.Contract.Name(&_RewardTracker.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RewardTracker *RewardTrackerCallerSession) Name() (string, error) {
	return _RewardTracker.Contract.Name(&_RewardTracker.CallOpts)
}

// PreviousCumulatedRewardPerToken is a free data retrieval call binding the contract method 0x44a08411.
//
// Solidity: function previousCumulatedRewardPerToken(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerCaller) PreviousCumulatedRewardPerToken(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "previousCumulatedRewardPerToken", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviousCumulatedRewardPerToken is a free data retrieval call binding the contract method 0x44a08411.
//
// Solidity: function previousCumulatedRewardPerToken(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerSession) PreviousCumulatedRewardPerToken(arg0 common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.PreviousCumulatedRewardPerToken(&_RewardTracker.CallOpts, arg0)
}

// PreviousCumulatedRewardPerToken is a free data retrieval call binding the contract method 0x44a08411.
//
// Solidity: function previousCumulatedRewardPerToken(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerCallerSession) PreviousCumulatedRewardPerToken(arg0 common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.PreviousCumulatedRewardPerToken(&_RewardTracker.CallOpts, arg0)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_RewardTracker *RewardTrackerCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_RewardTracker *RewardTrackerSession) RewardToken() (common.Address, error) {
	return _RewardTracker.Contract.RewardToken(&_RewardTracker.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_RewardTracker *RewardTrackerCallerSession) RewardToken() (common.Address, error) {
	return _RewardTracker.Contract.RewardToken(&_RewardTracker.CallOpts)
}

// StakedAmounts is a free data retrieval call binding the contract method 0x10c1c103.
//
// Solidity: function stakedAmounts(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerCaller) StakedAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "stakedAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakedAmounts is a free data retrieval call binding the contract method 0x10c1c103.
//
// Solidity: function stakedAmounts(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerSession) StakedAmounts(arg0 common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.StakedAmounts(&_RewardTracker.CallOpts, arg0)
}

// StakedAmounts is a free data retrieval call binding the contract method 0x10c1c103.
//
// Solidity: function stakedAmounts(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerCallerSession) StakedAmounts(arg0 common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.StakedAmounts(&_RewardTracker.CallOpts, arg0)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RewardTracker *RewardTrackerCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RewardTracker *RewardTrackerSession) Symbol() (string, error) {
	return _RewardTracker.Contract.Symbol(&_RewardTracker.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RewardTracker *RewardTrackerCallerSession) Symbol() (string, error) {
	return _RewardTracker.Contract.Symbol(&_RewardTracker.CallOpts)
}

// TokensPerInterval is a free data retrieval call binding the contract method 0xa8d93627.
//
// Solidity: function tokensPerInterval() view returns(uint256)
func (_RewardTracker *RewardTrackerCaller) TokensPerInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "tokensPerInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensPerInterval is a free data retrieval call binding the contract method 0xa8d93627.
//
// Solidity: function tokensPerInterval() view returns(uint256)
func (_RewardTracker *RewardTrackerSession) TokensPerInterval() (*big.Int, error) {
	return _RewardTracker.Contract.TokensPerInterval(&_RewardTracker.CallOpts)
}

// TokensPerInterval is a free data retrieval call binding the contract method 0xa8d93627.
//
// Solidity: function tokensPerInterval() view returns(uint256)
func (_RewardTracker *RewardTrackerCallerSession) TokensPerInterval() (*big.Int, error) {
	return _RewardTracker.Contract.TokensPerInterval(&_RewardTracker.CallOpts)
}

// TotalDepositSupply is a free data retrieval call binding the contract method 0x552ce1dc.
//
// Solidity: function totalDepositSupply(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerCaller) TotalDepositSupply(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "totalDepositSupply", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDepositSupply is a free data retrieval call binding the contract method 0x552ce1dc.
//
// Solidity: function totalDepositSupply(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerSession) TotalDepositSupply(arg0 common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.TotalDepositSupply(&_RewardTracker.CallOpts, arg0)
}

// TotalDepositSupply is a free data retrieval call binding the contract method 0x552ce1dc.
//
// Solidity: function totalDepositSupply(address ) view returns(uint256)
func (_RewardTracker *RewardTrackerCallerSession) TotalDepositSupply(arg0 common.Address) (*big.Int, error) {
	return _RewardTracker.Contract.TotalDepositSupply(&_RewardTracker.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RewardTracker *RewardTrackerCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RewardTracker.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RewardTracker *RewardTrackerSession) TotalSupply() (*big.Int, error) {
	return _RewardTracker.Contract.TotalSupply(&_RewardTracker.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RewardTracker *RewardTrackerCallerSession) TotalSupply() (*big.Int, error) {
	return _RewardTracker.Contract.TotalSupply(&_RewardTracker.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_RewardTracker *RewardTrackerTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_RewardTracker *RewardTrackerSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.Contract.Approve(&_RewardTracker.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_RewardTracker *RewardTrackerTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.Contract.Approve(&_RewardTracker.TransactOpts, _spender, _amount)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns(uint256)
func (_RewardTracker *RewardTrackerTransactor) Claim(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _RewardTracker.contract.Transact(opts, "claim", _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns(uint256)
func (_RewardTracker *RewardTrackerSession) Claim(_receiver common.Address) (*types.Transaction, error) {
	return _RewardTracker.Contract.Claim(&_RewardTracker.TransactOpts, _receiver)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address _receiver) returns(uint256)
func (_RewardTracker *RewardTrackerTransactorSession) Claim(_receiver common.Address) (*types.Transaction, error) {
	return _RewardTracker.Contract.Claim(&_RewardTracker.TransactOpts, _receiver)
}

// ClaimForAccount is a paid mutator transaction binding the contract method 0x13e82e7a.
//
// Solidity: function claimForAccount(address _account, address _receiver) returns(uint256)
func (_RewardTracker *RewardTrackerTransactor) ClaimForAccount(opts *bind.TransactOpts, _account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _RewardTracker.contract.Transact(opts, "claimForAccount", _account, _receiver)
}

// ClaimForAccount is a paid mutator transaction binding the contract method 0x13e82e7a.
//
// Solidity: function claimForAccount(address _account, address _receiver) returns(uint256)
func (_RewardTracker *RewardTrackerSession) ClaimForAccount(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _RewardTracker.Contract.ClaimForAccount(&_RewardTracker.TransactOpts, _account, _receiver)
}

// ClaimForAccount is a paid mutator transaction binding the contract method 0x13e82e7a.
//
// Solidity: function claimForAccount(address _account, address _receiver) returns(uint256)
func (_RewardTracker *RewardTrackerTransactorSession) ClaimForAccount(_account common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _RewardTracker.Contract.ClaimForAccount(&_RewardTracker.TransactOpts, _account, _receiver)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _depositTokens, address _distributor) returns()
func (_RewardTracker *RewardTrackerTransactor) Initialize(opts *bind.TransactOpts, _depositTokens []common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _RewardTracker.contract.Transact(opts, "initialize", _depositTokens, _distributor)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _depositTokens, address _distributor) returns()
func (_RewardTracker *RewardTrackerSession) Initialize(_depositTokens []common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _RewardTracker.Contract.Initialize(&_RewardTracker.TransactOpts, _depositTokens, _distributor)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] _depositTokens, address _distributor) returns()
func (_RewardTracker *RewardTrackerTransactorSession) Initialize(_depositTokens []common.Address, _distributor common.Address) (*types.Transaction, error) {
	return _RewardTracker.Contract.Initialize(&_RewardTracker.TransactOpts, _depositTokens, _distributor)
}

// SetDepositToken is a paid mutator transaction binding the contract method 0xe44b7558.
//
// Solidity: function setDepositToken(address _depositToken, bool _isDepositToken) returns()
func (_RewardTracker *RewardTrackerTransactor) SetDepositToken(opts *bind.TransactOpts, _depositToken common.Address, _isDepositToken bool) (*types.Transaction, error) {
	return _RewardTracker.contract.Transact(opts, "setDepositToken", _depositToken, _isDepositToken)
}

// SetDepositToken is a paid mutator transaction binding the contract method 0xe44b7558.
//
// Solidity: function setDepositToken(address _depositToken, bool _isDepositToken) returns()
func (_RewardTracker *RewardTrackerSession) SetDepositToken(_depositToken common.Address, _isDepositToken bool) (*types.Transaction, error) {
	return _RewardTracker.Contract.SetDepositToken(&_RewardTracker.TransactOpts, _depositToken, _isDepositToken)
}

// SetDepositToken is a paid mutator transaction binding the contract method 0xe44b7558.
//
// Solidity: function setDepositToken(address _depositToken, bool _isDepositToken) returns()
func (_RewardTracker *RewardTrackerTransactorSession) SetDepositToken(_depositToken common.Address, _isDepositToken bool) (*types.Transaction, error) {
	return _RewardTracker.Contract.SetDepositToken(&_RewardTracker.TransactOpts, _depositToken, _isDepositToken)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_RewardTracker *RewardTrackerTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _RewardTracker.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_RewardTracker *RewardTrackerSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _RewardTracker.Contract.SetGov(&_RewardTracker.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_RewardTracker *RewardTrackerTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _RewardTracker.Contract.SetGov(&_RewardTracker.TransactOpts, _gov)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_RewardTracker *RewardTrackerTransactor) SetHandler(opts *bind.TransactOpts, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _RewardTracker.contract.Transact(opts, "setHandler", _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_RewardTracker *RewardTrackerSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _RewardTracker.Contract.SetHandler(&_RewardTracker.TransactOpts, _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_RewardTracker *RewardTrackerTransactorSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _RewardTracker.Contract.SetHandler(&_RewardTracker.TransactOpts, _handler, _isActive)
}

// SetInPrivateClaimingMode is a paid mutator transaction binding the contract method 0x3cd7f700.
//
// Solidity: function setInPrivateClaimingMode(bool _inPrivateClaimingMode) returns()
func (_RewardTracker *RewardTrackerTransactor) SetInPrivateClaimingMode(opts *bind.TransactOpts, _inPrivateClaimingMode bool) (*types.Transaction, error) {
	return _RewardTracker.contract.Transact(opts, "setInPrivateClaimingMode", _inPrivateClaimingMode)
}

// SetInPrivateClaimingMode is a paid mutator transaction binding the contract method 0x3cd7f700.
//
// Solidity: function setInPrivateClaimingMode(bool _inPrivateClaimingMode) returns()
func (_RewardTracker *RewardTrackerSession) SetInPrivateClaimingMode(_inPrivateClaimingMode bool) (*types.Transaction, error) {
	return _RewardTracker.Contract.SetInPrivateClaimingMode(&_RewardTracker.TransactOpts, _inPrivateClaimingMode)
}

// SetInPrivateClaimingMode is a paid mutator transaction binding the contract method 0x3cd7f700.
//
// Solidity: function setInPrivateClaimingMode(bool _inPrivateClaimingMode) returns()
func (_RewardTracker *RewardTrackerTransactorSession) SetInPrivateClaimingMode(_inPrivateClaimingMode bool) (*types.Transaction, error) {
	return _RewardTracker.Contract.SetInPrivateClaimingMode(&_RewardTracker.TransactOpts, _inPrivateClaimingMode)
}

// SetInPrivateStakingMode is a paid mutator transaction binding the contract method 0x1d30d5bc.
//
// Solidity: function setInPrivateStakingMode(bool _inPrivateStakingMode) returns()
func (_RewardTracker *RewardTrackerTransactor) SetInPrivateStakingMode(opts *bind.TransactOpts, _inPrivateStakingMode bool) (*types.Transaction, error) {
	return _RewardTracker.contract.Transact(opts, "setInPrivateStakingMode", _inPrivateStakingMode)
}

// SetInPrivateStakingMode is a paid mutator transaction binding the contract method 0x1d30d5bc.
//
// Solidity: function setInPrivateStakingMode(bool _inPrivateStakingMode) returns()
func (_RewardTracker *RewardTrackerSession) SetInPrivateStakingMode(_inPrivateStakingMode bool) (*types.Transaction, error) {
	return _RewardTracker.Contract.SetInPrivateStakingMode(&_RewardTracker.TransactOpts, _inPrivateStakingMode)
}

// SetInPrivateStakingMode is a paid mutator transaction binding the contract method 0x1d30d5bc.
//
// Solidity: function setInPrivateStakingMode(bool _inPrivateStakingMode) returns()
func (_RewardTracker *RewardTrackerTransactorSession) SetInPrivateStakingMode(_inPrivateStakingMode bool) (*types.Transaction, error) {
	return _RewardTracker.Contract.SetInPrivateStakingMode(&_RewardTracker.TransactOpts, _inPrivateStakingMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x5a47a1a7.
//
// Solidity: function setInPrivateTransferMode(bool _inPrivateTransferMode) returns()
func (_RewardTracker *RewardTrackerTransactor) SetInPrivateTransferMode(opts *bind.TransactOpts, _inPrivateTransferMode bool) (*types.Transaction, error) {
	return _RewardTracker.contract.Transact(opts, "setInPrivateTransferMode", _inPrivateTransferMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x5a47a1a7.
//
// Solidity: function setInPrivateTransferMode(bool _inPrivateTransferMode) returns()
func (_RewardTracker *RewardTrackerSession) SetInPrivateTransferMode(_inPrivateTransferMode bool) (*types.Transaction, error) {
	return _RewardTracker.Contract.SetInPrivateTransferMode(&_RewardTracker.TransactOpts, _inPrivateTransferMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x5a47a1a7.
//
// Solidity: function setInPrivateTransferMode(bool _inPrivateTransferMode) returns()
func (_RewardTracker *RewardTrackerTransactorSession) SetInPrivateTransferMode(_inPrivateTransferMode bool) (*types.Transaction, error) {
	return _RewardTracker.Contract.SetInPrivateTransferMode(&_RewardTracker.TransactOpts, _inPrivateTransferMode)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _depositToken, uint256 _amount) returns()
func (_RewardTracker *RewardTrackerTransactor) Stake(opts *bind.TransactOpts, _depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.contract.Transact(opts, "stake", _depositToken, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _depositToken, uint256 _amount) returns()
func (_RewardTracker *RewardTrackerSession) Stake(_depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.Contract.Stake(&_RewardTracker.TransactOpts, _depositToken, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xadc9772e.
//
// Solidity: function stake(address _depositToken, uint256 _amount) returns()
func (_RewardTracker *RewardTrackerTransactorSession) Stake(_depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.Contract.Stake(&_RewardTracker.TransactOpts, _depositToken, _amount)
}

// StakeForAccount is a paid mutator transaction binding the contract method 0x790b5a6c.
//
// Solidity: function stakeForAccount(address _fundingAccount, address _account, address _depositToken, uint256 _amount) returns()
func (_RewardTracker *RewardTrackerTransactor) StakeForAccount(opts *bind.TransactOpts, _fundingAccount common.Address, _account common.Address, _depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.contract.Transact(opts, "stakeForAccount", _fundingAccount, _account, _depositToken, _amount)
}

// StakeForAccount is a paid mutator transaction binding the contract method 0x790b5a6c.
//
// Solidity: function stakeForAccount(address _fundingAccount, address _account, address _depositToken, uint256 _amount) returns()
func (_RewardTracker *RewardTrackerSession) StakeForAccount(_fundingAccount common.Address, _account common.Address, _depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.Contract.StakeForAccount(&_RewardTracker.TransactOpts, _fundingAccount, _account, _depositToken, _amount)
}

// StakeForAccount is a paid mutator transaction binding the contract method 0x790b5a6c.
//
// Solidity: function stakeForAccount(address _fundingAccount, address _account, address _depositToken, uint256 _amount) returns()
func (_RewardTracker *RewardTrackerTransactorSession) StakeForAccount(_fundingAccount common.Address, _account common.Address, _depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.Contract.StakeForAccount(&_RewardTracker.TransactOpts, _fundingAccount, _account, _depositToken, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_RewardTracker *RewardTrackerTransactor) Transfer(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.contract.Transact(opts, "transfer", _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_RewardTracker *RewardTrackerSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.Contract.Transfer(&_RewardTracker.TransactOpts, _recipient, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _recipient, uint256 _amount) returns(bool)
func (_RewardTracker *RewardTrackerTransactorSession) Transfer(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.Contract.Transfer(&_RewardTracker.TransactOpts, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_RewardTracker *RewardTrackerTransactor) TransferFrom(opts *bind.TransactOpts, _sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.contract.Transact(opts, "transferFrom", _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_RewardTracker *RewardTrackerSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.Contract.TransferFrom(&_RewardTracker.TransactOpts, _sender, _recipient, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _sender, address _recipient, uint256 _amount) returns(bool)
func (_RewardTracker *RewardTrackerTransactorSession) TransferFrom(_sender common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.Contract.TransferFrom(&_RewardTracker.TransactOpts, _sender, _recipient, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _depositToken, uint256 _amount) returns()
func (_RewardTracker *RewardTrackerTransactor) Unstake(opts *bind.TransactOpts, _depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.contract.Transact(opts, "unstake", _depositToken, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _depositToken, uint256 _amount) returns()
func (_RewardTracker *RewardTrackerSession) Unstake(_depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.Contract.Unstake(&_RewardTracker.TransactOpts, _depositToken, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _depositToken, uint256 _amount) returns()
func (_RewardTracker *RewardTrackerTransactorSession) Unstake(_depositToken common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.Contract.Unstake(&_RewardTracker.TransactOpts, _depositToken, _amount)
}

// UnstakeForAccount is a paid mutator transaction binding the contract method 0x098bf59d.
//
// Solidity: function unstakeForAccount(address _account, address _depositToken, uint256 _amount, address _receiver) returns()
func (_RewardTracker *RewardTrackerTransactor) UnstakeForAccount(opts *bind.TransactOpts, _account common.Address, _depositToken common.Address, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _RewardTracker.contract.Transact(opts, "unstakeForAccount", _account, _depositToken, _amount, _receiver)
}

// UnstakeForAccount is a paid mutator transaction binding the contract method 0x098bf59d.
//
// Solidity: function unstakeForAccount(address _account, address _depositToken, uint256 _amount, address _receiver) returns()
func (_RewardTracker *RewardTrackerSession) UnstakeForAccount(_account common.Address, _depositToken common.Address, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _RewardTracker.Contract.UnstakeForAccount(&_RewardTracker.TransactOpts, _account, _depositToken, _amount, _receiver)
}

// UnstakeForAccount is a paid mutator transaction binding the contract method 0x098bf59d.
//
// Solidity: function unstakeForAccount(address _account, address _depositToken, uint256 _amount, address _receiver) returns()
func (_RewardTracker *RewardTrackerTransactorSession) UnstakeForAccount(_account common.Address, _depositToken common.Address, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _RewardTracker.Contract.UnstakeForAccount(&_RewardTracker.TransactOpts, _account, _depositToken, _amount, _receiver)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x3e158b0c.
//
// Solidity: function updateRewards() returns()
func (_RewardTracker *RewardTrackerTransactor) UpdateRewards(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardTracker.contract.Transact(opts, "updateRewards")
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x3e158b0c.
//
// Solidity: function updateRewards() returns()
func (_RewardTracker *RewardTrackerSession) UpdateRewards() (*types.Transaction, error) {
	return _RewardTracker.Contract.UpdateRewards(&_RewardTracker.TransactOpts)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x3e158b0c.
//
// Solidity: function updateRewards() returns()
func (_RewardTracker *RewardTrackerTransactorSession) UpdateRewards() (*types.Transaction, error) {
	return _RewardTracker.Contract.UpdateRewards(&_RewardTracker.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_RewardTracker *RewardTrackerTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_RewardTracker *RewardTrackerSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.Contract.WithdrawToken(&_RewardTracker.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_RewardTracker *RewardTrackerTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardTracker.Contract.WithdrawToken(&_RewardTracker.TransactOpts, _token, _account, _amount)
}

// RewardTrackerApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RewardTracker contract.
type RewardTrackerApprovalIterator struct {
	Event *RewardTrackerApproval // Event containing the contract specifics and raw log

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
func (it *RewardTrackerApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardTrackerApproval)
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
		it.Event = new(RewardTrackerApproval)
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
func (it *RewardTrackerApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardTrackerApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardTrackerApproval represents a Approval event raised by the RewardTracker contract.
type RewardTrackerApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RewardTracker *RewardTrackerFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*RewardTrackerApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RewardTracker.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &RewardTrackerApprovalIterator{contract: _RewardTracker.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_RewardTracker *RewardTrackerFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RewardTrackerApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RewardTracker.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardTrackerApproval)
				if err := _RewardTracker.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_RewardTracker *RewardTrackerFilterer) ParseApproval(log types.Log) (*RewardTrackerApproval, error) {
	event := new(RewardTrackerApproval)
	if err := _RewardTracker.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardTrackerClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the RewardTracker contract.
type RewardTrackerClaimIterator struct {
	Event *RewardTrackerClaim // Event containing the contract specifics and raw log

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
func (it *RewardTrackerClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardTrackerClaim)
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
		it.Event = new(RewardTrackerClaim)
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
func (it *RewardTrackerClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardTrackerClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardTrackerClaim represents a Claim event raised by the RewardTracker contract.
type RewardTrackerClaim struct {
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address receiver, uint256 amount)
func (_RewardTracker *RewardTrackerFilterer) FilterClaim(opts *bind.FilterOpts) (*RewardTrackerClaimIterator, error) {

	logs, sub, err := _RewardTracker.contract.FilterLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return &RewardTrackerClaimIterator{contract: _RewardTracker.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address receiver, uint256 amount)
func (_RewardTracker *RewardTrackerFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *RewardTrackerClaim) (event.Subscription, error) {

	logs, sub, err := _RewardTracker.contract.WatchLogs(opts, "Claim")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardTrackerClaim)
				if err := _RewardTracker.contract.UnpackLog(event, "Claim", log); err != nil {
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
func (_RewardTracker *RewardTrackerFilterer) ParseClaim(log types.Log) (*RewardTrackerClaim, error) {
	event := new(RewardTrackerClaim)
	if err := _RewardTracker.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardTrackerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RewardTracker contract.
type RewardTrackerTransferIterator struct {
	Event *RewardTrackerTransfer // Event containing the contract specifics and raw log

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
func (it *RewardTrackerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardTrackerTransfer)
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
		it.Event = new(RewardTrackerTransfer)
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
func (it *RewardTrackerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardTrackerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardTrackerTransfer represents a Transfer event raised by the RewardTracker contract.
type RewardTrackerTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RewardTracker *RewardTrackerFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RewardTrackerTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RewardTracker.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RewardTrackerTransferIterator{contract: _RewardTracker.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_RewardTracker *RewardTrackerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RewardTrackerTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RewardTracker.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardTrackerTransfer)
				if err := _RewardTracker.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_RewardTracker *RewardTrackerFilterer) ParseTransfer(log types.Log) (*RewardTrackerTransfer, error) {
	event := new(RewardTrackerTransfer)
	if err := _RewardTracker.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

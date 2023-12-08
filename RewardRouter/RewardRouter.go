// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RewardRouter

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

// RewardRouterMetaData contains all meta data concerning the RewardRouter contract.
var RewardRouterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"acceptTransfer\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchCompoundForAccounts\",\"inputs\":[{\"name\":\"_accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchStakeGmxForAccount\",\"inputs\":[{\"name\":\"_accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bnGmx\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bonusGmxTracker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimEsGmx\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimFees\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"compound\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"compoundForAccount\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"esGmx\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeGlpTracker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeGmxTracker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"glp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"glpManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"glpVester\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gmx\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gmxVester\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"handleRewards\",\"inputs\":[{\"name\":\"_shouldClaimGmx\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_shouldStakeGmx\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_shouldClaimEsGmx\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_shouldStakeEsGmx\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_shouldStakeMultiplierPoints\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_shouldClaimWeth\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_shouldConvertWethToEth\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_weth\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_gmx\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_esGmx\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_bnGmx\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_glp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stakedGmxTracker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_bonusGmxTracker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feeGmxTracker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_feeGlpTracker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_stakedGlpTracker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_glpManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_gmxVester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_glpVester\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isInitialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mintAndStakeGlp\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minUsdg\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minGlp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintAndStakeGlpETH\",\"inputs\":[{\"name\":\"_minUsdg\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minGlp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"pendingReceivers\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signalTransfer\",\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeEsGmx\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeGmx\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakeGmxForAccount\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"stakedGlpTracker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stakedGmxTracker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unstakeAndRedeemGlp\",\"inputs\":[{\"name\":\"_tokenOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_glpAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstakeAndRedeemGlpETH\",\"inputs\":[{\"name\":\"_glpAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstakeEsGmx\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unstakeGmx\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"weth\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"StakeGlp\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakeGmx\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnstakeGlp\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UnstakeGmx\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// RewardRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardRouterMetaData.ABI instead.
var RewardRouterABI = RewardRouterMetaData.ABI

// RewardRouter is an auto generated Go binding around an Ethereum contract.
type RewardRouter struct {
	RewardRouterCaller     // Read-only binding to the contract
	RewardRouterTransactor // Write-only binding to the contract
	RewardRouterFilterer   // Log filterer for contract events
}

// RewardRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardRouterSession struct {
	Contract     *RewardRouter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardRouterCallerSession struct {
	Contract *RewardRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RewardRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardRouterTransactorSession struct {
	Contract     *RewardRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RewardRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardRouterRaw struct {
	Contract *RewardRouter // Generic contract binding to access the raw methods on
}

// RewardRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardRouterCallerRaw struct {
	Contract *RewardRouterCaller // Generic read-only contract binding to access the raw methods on
}

// RewardRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardRouterTransactorRaw struct {
	Contract *RewardRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardRouter creates a new instance of RewardRouter, bound to a specific deployed contract.
func NewRewardRouter(address common.Address, backend bind.ContractBackend) (*RewardRouter, error) {
	contract, err := bindRewardRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardRouter{RewardRouterCaller: RewardRouterCaller{contract: contract}, RewardRouterTransactor: RewardRouterTransactor{contract: contract}, RewardRouterFilterer: RewardRouterFilterer{contract: contract}}, nil
}

// NewRewardRouterCaller creates a new read-only instance of RewardRouter, bound to a specific deployed contract.
func NewRewardRouterCaller(address common.Address, caller bind.ContractCaller) (*RewardRouterCaller, error) {
	contract, err := bindRewardRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardRouterCaller{contract: contract}, nil
}

// NewRewardRouterTransactor creates a new write-only instance of RewardRouter, bound to a specific deployed contract.
func NewRewardRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardRouterTransactor, error) {
	contract, err := bindRewardRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardRouterTransactor{contract: contract}, nil
}

// NewRewardRouterFilterer creates a new log filterer instance of RewardRouter, bound to a specific deployed contract.
func NewRewardRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardRouterFilterer, error) {
	contract, err := bindRewardRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardRouterFilterer{contract: contract}, nil
}

// bindRewardRouter binds a generic wrapper to an already deployed contract.
func bindRewardRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RewardRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardRouter *RewardRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardRouter.Contract.RewardRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardRouter *RewardRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardRouter.Contract.RewardRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardRouter *RewardRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardRouter.Contract.RewardRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardRouter *RewardRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardRouter *RewardRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardRouter *RewardRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardRouter.Contract.contract.Transact(opts, method, params...)
}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_RewardRouter *RewardRouterCaller) BnGmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouter.contract.Call(opts, &out, "bnGmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_RewardRouter *RewardRouterSession) BnGmx() (common.Address, error) {
	return _RewardRouter.Contract.BnGmx(&_RewardRouter.CallOpts)
}

// BnGmx is a free data retrieval call binding the contract method 0xe102f564.
//
// Solidity: function bnGmx() view returns(address)
func (_RewardRouter *RewardRouterCallerSession) BnGmx() (common.Address, error) {
	return _RewardRouter.Contract.BnGmx(&_RewardRouter.CallOpts)
}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_RewardRouter *RewardRouterCaller) BonusGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouter.contract.Call(opts, &out, "bonusGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_RewardRouter *RewardRouterSession) BonusGmxTracker() (common.Address, error) {
	return _RewardRouter.Contract.BonusGmxTracker(&_RewardRouter.CallOpts)
}

// BonusGmxTracker is a free data retrieval call binding the contract method 0x1fcd60e5.
//
// Solidity: function bonusGmxTracker() view returns(address)
func (_RewardRouter *RewardRouterCallerSession) BonusGmxTracker() (common.Address, error) {
	return _RewardRouter.Contract.BonusGmxTracker(&_RewardRouter.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_RewardRouter *RewardRouterCaller) EsGmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouter.contract.Call(opts, &out, "esGmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_RewardRouter *RewardRouterSession) EsGmx() (common.Address, error) {
	return _RewardRouter.Contract.EsGmx(&_RewardRouter.CallOpts)
}

// EsGmx is a free data retrieval call binding the contract method 0x6a192a78.
//
// Solidity: function esGmx() view returns(address)
func (_RewardRouter *RewardRouterCallerSession) EsGmx() (common.Address, error) {
	return _RewardRouter.Contract.EsGmx(&_RewardRouter.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_RewardRouter *RewardRouterCaller) FeeGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouter.contract.Call(opts, &out, "feeGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_RewardRouter *RewardRouterSession) FeeGlpTracker() (common.Address, error) {
	return _RewardRouter.Contract.FeeGlpTracker(&_RewardRouter.CallOpts)
}

// FeeGlpTracker is a free data retrieval call binding the contract method 0xe1c363b7.
//
// Solidity: function feeGlpTracker() view returns(address)
func (_RewardRouter *RewardRouterCallerSession) FeeGlpTracker() (common.Address, error) {
	return _RewardRouter.Contract.FeeGlpTracker(&_RewardRouter.CallOpts)
}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_RewardRouter *RewardRouterCaller) FeeGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouter.contract.Call(opts, &out, "feeGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_RewardRouter *RewardRouterSession) FeeGmxTracker() (common.Address, error) {
	return _RewardRouter.Contract.FeeGmxTracker(&_RewardRouter.CallOpts)
}

// FeeGmxTracker is a free data retrieval call binding the contract method 0x51c3e3b4.
//
// Solidity: function feeGmxTracker() view returns(address)
func (_RewardRouter *RewardRouterCallerSession) FeeGmxTracker() (common.Address, error) {
	return _RewardRouter.Contract.FeeGmxTracker(&_RewardRouter.CallOpts)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_RewardRouter *RewardRouterCaller) Glp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouter.contract.Call(opts, &out, "glp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_RewardRouter *RewardRouterSession) Glp() (common.Address, error) {
	return _RewardRouter.Contract.Glp(&_RewardRouter.CallOpts)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_RewardRouter *RewardRouterCallerSession) Glp() (common.Address, error) {
	return _RewardRouter.Contract.Glp(&_RewardRouter.CallOpts)
}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_RewardRouter *RewardRouterCaller) GlpManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouter.contract.Call(opts, &out, "glpManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_RewardRouter *RewardRouterSession) GlpManager() (common.Address, error) {
	return _RewardRouter.Contract.GlpManager(&_RewardRouter.CallOpts)
}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_RewardRouter *RewardRouterCallerSession) GlpManager() (common.Address, error) {
	return _RewardRouter.Contract.GlpManager(&_RewardRouter.CallOpts)
}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_RewardRouter *RewardRouterCaller) GlpVester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouter.contract.Call(opts, &out, "glpVester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_RewardRouter *RewardRouterSession) GlpVester() (common.Address, error) {
	return _RewardRouter.Contract.GlpVester(&_RewardRouter.CallOpts)
}

// GlpVester is a free data retrieval call binding the contract method 0x3671df25.
//
// Solidity: function glpVester() view returns(address)
func (_RewardRouter *RewardRouterCallerSession) GlpVester() (common.Address, error) {
	return _RewardRouter.Contract.GlpVester(&_RewardRouter.CallOpts)
}

// Gmx is a free data retrieval call binding the contract method 0x31e67c71.
//
// Solidity: function gmx() view returns(address)
func (_RewardRouter *RewardRouterCaller) Gmx(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouter.contract.Call(opts, &out, "gmx")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gmx is a free data retrieval call binding the contract method 0x31e67c71.
//
// Solidity: function gmx() view returns(address)
func (_RewardRouter *RewardRouterSession) Gmx() (common.Address, error) {
	return _RewardRouter.Contract.Gmx(&_RewardRouter.CallOpts)
}

// Gmx is a free data retrieval call binding the contract method 0x31e67c71.
//
// Solidity: function gmx() view returns(address)
func (_RewardRouter *RewardRouterCallerSession) Gmx() (common.Address, error) {
	return _RewardRouter.Contract.Gmx(&_RewardRouter.CallOpts)
}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_RewardRouter *RewardRouterCaller) GmxVester(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouter.contract.Call(opts, &out, "gmxVester")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_RewardRouter *RewardRouterSession) GmxVester() (common.Address, error) {
	return _RewardRouter.Contract.GmxVester(&_RewardRouter.CallOpts)
}

// GmxVester is a free data retrieval call binding the contract method 0x41d315b4.
//
// Solidity: function gmxVester() view returns(address)
func (_RewardRouter *RewardRouterCallerSession) GmxVester() (common.Address, error) {
	return _RewardRouter.Contract.GmxVester(&_RewardRouter.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_RewardRouter *RewardRouterCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouter.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_RewardRouter *RewardRouterSession) Gov() (common.Address, error) {
	return _RewardRouter.Contract.Gov(&_RewardRouter.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_RewardRouter *RewardRouterCallerSession) Gov() (common.Address, error) {
	return _RewardRouter.Contract.Gov(&_RewardRouter.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_RewardRouter *RewardRouterCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RewardRouter.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_RewardRouter *RewardRouterSession) IsInitialized() (bool, error) {
	return _RewardRouter.Contract.IsInitialized(&_RewardRouter.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_RewardRouter *RewardRouterCallerSession) IsInitialized() (bool, error) {
	return _RewardRouter.Contract.IsInitialized(&_RewardRouter.CallOpts)
}

// PendingReceivers is a free data retrieval call binding the contract method 0xe1b9db89.
//
// Solidity: function pendingReceivers(address ) view returns(address)
func (_RewardRouter *RewardRouterCaller) PendingReceivers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _RewardRouter.contract.Call(opts, &out, "pendingReceivers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingReceivers is a free data retrieval call binding the contract method 0xe1b9db89.
//
// Solidity: function pendingReceivers(address ) view returns(address)
func (_RewardRouter *RewardRouterSession) PendingReceivers(arg0 common.Address) (common.Address, error) {
	return _RewardRouter.Contract.PendingReceivers(&_RewardRouter.CallOpts, arg0)
}

// PendingReceivers is a free data retrieval call binding the contract method 0xe1b9db89.
//
// Solidity: function pendingReceivers(address ) view returns(address)
func (_RewardRouter *RewardRouterCallerSession) PendingReceivers(arg0 common.Address) (common.Address, error) {
	return _RewardRouter.Contract.PendingReceivers(&_RewardRouter.CallOpts, arg0)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_RewardRouter *RewardRouterCaller) StakedGlpTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouter.contract.Call(opts, &out, "stakedGlpTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_RewardRouter *RewardRouterSession) StakedGlpTracker() (common.Address, error) {
	return _RewardRouter.Contract.StakedGlpTracker(&_RewardRouter.CallOpts)
}

// StakedGlpTracker is a free data retrieval call binding the contract method 0xaf394d00.
//
// Solidity: function stakedGlpTracker() view returns(address)
func (_RewardRouter *RewardRouterCallerSession) StakedGlpTracker() (common.Address, error) {
	return _RewardRouter.Contract.StakedGlpTracker(&_RewardRouter.CallOpts)
}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_RewardRouter *RewardRouterCaller) StakedGmxTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouter.contract.Call(opts, &out, "stakedGmxTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_RewardRouter *RewardRouterSession) StakedGmxTracker() (common.Address, error) {
	return _RewardRouter.Contract.StakedGmxTracker(&_RewardRouter.CallOpts)
}

// StakedGmxTracker is a free data retrieval call binding the contract method 0x0ce4018a.
//
// Solidity: function stakedGmxTracker() view returns(address)
func (_RewardRouter *RewardRouterCallerSession) StakedGmxTracker() (common.Address, error) {
	return _RewardRouter.Contract.StakedGmxTracker(&_RewardRouter.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_RewardRouter *RewardRouterCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardRouter.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_RewardRouter *RewardRouterSession) Weth() (common.Address, error) {
	return _RewardRouter.Contract.Weth(&_RewardRouter.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_RewardRouter *RewardRouterCallerSession) Weth() (common.Address, error) {
	return _RewardRouter.Contract.Weth(&_RewardRouter.CallOpts)
}

// AcceptTransfer is a paid mutator transaction binding the contract method 0x655603a4.
//
// Solidity: function acceptTransfer(address _sender) returns()
func (_RewardRouter *RewardRouterTransactor) AcceptTransfer(opts *bind.TransactOpts, _sender common.Address) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "acceptTransfer", _sender)
}

// AcceptTransfer is a paid mutator transaction binding the contract method 0x655603a4.
//
// Solidity: function acceptTransfer(address _sender) returns()
func (_RewardRouter *RewardRouterSession) AcceptTransfer(_sender common.Address) (*types.Transaction, error) {
	return _RewardRouter.Contract.AcceptTransfer(&_RewardRouter.TransactOpts, _sender)
}

// AcceptTransfer is a paid mutator transaction binding the contract method 0x655603a4.
//
// Solidity: function acceptTransfer(address _sender) returns()
func (_RewardRouter *RewardRouterTransactorSession) AcceptTransfer(_sender common.Address) (*types.Transaction, error) {
	return _RewardRouter.Contract.AcceptTransfer(&_RewardRouter.TransactOpts, _sender)
}

// BatchCompoundForAccounts is a paid mutator transaction binding the contract method 0x1af276a6.
//
// Solidity: function batchCompoundForAccounts(address[] _accounts) returns()
func (_RewardRouter *RewardRouterTransactor) BatchCompoundForAccounts(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "batchCompoundForAccounts", _accounts)
}

// BatchCompoundForAccounts is a paid mutator transaction binding the contract method 0x1af276a6.
//
// Solidity: function batchCompoundForAccounts(address[] _accounts) returns()
func (_RewardRouter *RewardRouterSession) BatchCompoundForAccounts(_accounts []common.Address) (*types.Transaction, error) {
	return _RewardRouter.Contract.BatchCompoundForAccounts(&_RewardRouter.TransactOpts, _accounts)
}

// BatchCompoundForAccounts is a paid mutator transaction binding the contract method 0x1af276a6.
//
// Solidity: function batchCompoundForAccounts(address[] _accounts) returns()
func (_RewardRouter *RewardRouterTransactorSession) BatchCompoundForAccounts(_accounts []common.Address) (*types.Transaction, error) {
	return _RewardRouter.Contract.BatchCompoundForAccounts(&_RewardRouter.TransactOpts, _accounts)
}

// BatchStakeGmxForAccount is a paid mutator transaction binding the contract method 0x0db79e52.
//
// Solidity: function batchStakeGmxForAccount(address[] _accounts, uint256[] _amounts) returns()
func (_RewardRouter *RewardRouterTransactor) BatchStakeGmxForAccount(opts *bind.TransactOpts, _accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "batchStakeGmxForAccount", _accounts, _amounts)
}

// BatchStakeGmxForAccount is a paid mutator transaction binding the contract method 0x0db79e52.
//
// Solidity: function batchStakeGmxForAccount(address[] _accounts, uint256[] _amounts) returns()
func (_RewardRouter *RewardRouterSession) BatchStakeGmxForAccount(_accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _RewardRouter.Contract.BatchStakeGmxForAccount(&_RewardRouter.TransactOpts, _accounts, _amounts)
}

// BatchStakeGmxForAccount is a paid mutator transaction binding the contract method 0x0db79e52.
//
// Solidity: function batchStakeGmxForAccount(address[] _accounts, uint256[] _amounts) returns()
func (_RewardRouter *RewardRouterTransactorSession) BatchStakeGmxForAccount(_accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _RewardRouter.Contract.BatchStakeGmxForAccount(&_RewardRouter.TransactOpts, _accounts, _amounts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_RewardRouter *RewardRouterTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_RewardRouter *RewardRouterSession) Claim() (*types.Transaction, error) {
	return _RewardRouter.Contract.Claim(&_RewardRouter.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_RewardRouter *RewardRouterTransactorSession) Claim() (*types.Transaction, error) {
	return _RewardRouter.Contract.Claim(&_RewardRouter.TransactOpts)
}

// ClaimEsGmx is a paid mutator transaction binding the contract method 0x5fe3945f.
//
// Solidity: function claimEsGmx() returns()
func (_RewardRouter *RewardRouterTransactor) ClaimEsGmx(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "claimEsGmx")
}

// ClaimEsGmx is a paid mutator transaction binding the contract method 0x5fe3945f.
//
// Solidity: function claimEsGmx() returns()
func (_RewardRouter *RewardRouterSession) ClaimEsGmx() (*types.Transaction, error) {
	return _RewardRouter.Contract.ClaimEsGmx(&_RewardRouter.TransactOpts)
}

// ClaimEsGmx is a paid mutator transaction binding the contract method 0x5fe3945f.
//
// Solidity: function claimEsGmx() returns()
func (_RewardRouter *RewardRouterTransactorSession) ClaimEsGmx() (*types.Transaction, error) {
	return _RewardRouter.Contract.ClaimEsGmx(&_RewardRouter.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_RewardRouter *RewardRouterTransactor) ClaimFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "claimFees")
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_RewardRouter *RewardRouterSession) ClaimFees() (*types.Transaction, error) {
	return _RewardRouter.Contract.ClaimFees(&_RewardRouter.TransactOpts)
}

// ClaimFees is a paid mutator transaction binding the contract method 0xd294f093.
//
// Solidity: function claimFees() returns()
func (_RewardRouter *RewardRouterTransactorSession) ClaimFees() (*types.Transaction, error) {
	return _RewardRouter.Contract.ClaimFees(&_RewardRouter.TransactOpts)
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns()
func (_RewardRouter *RewardRouterTransactor) Compound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "compound")
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns()
func (_RewardRouter *RewardRouterSession) Compound() (*types.Transaction, error) {
	return _RewardRouter.Contract.Compound(&_RewardRouter.TransactOpts)
}

// Compound is a paid mutator transaction binding the contract method 0xf69e2046.
//
// Solidity: function compound() returns()
func (_RewardRouter *RewardRouterTransactorSession) Compound() (*types.Transaction, error) {
	return _RewardRouter.Contract.Compound(&_RewardRouter.TransactOpts)
}

// CompoundForAccount is a paid mutator transaction binding the contract method 0x2a9f4083.
//
// Solidity: function compoundForAccount(address _account) returns()
func (_RewardRouter *RewardRouterTransactor) CompoundForAccount(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "compoundForAccount", _account)
}

// CompoundForAccount is a paid mutator transaction binding the contract method 0x2a9f4083.
//
// Solidity: function compoundForAccount(address _account) returns()
func (_RewardRouter *RewardRouterSession) CompoundForAccount(_account common.Address) (*types.Transaction, error) {
	return _RewardRouter.Contract.CompoundForAccount(&_RewardRouter.TransactOpts, _account)
}

// CompoundForAccount is a paid mutator transaction binding the contract method 0x2a9f4083.
//
// Solidity: function compoundForAccount(address _account) returns()
func (_RewardRouter *RewardRouterTransactorSession) CompoundForAccount(_account common.Address) (*types.Transaction, error) {
	return _RewardRouter.Contract.CompoundForAccount(&_RewardRouter.TransactOpts, _account)
}

// HandleRewards is a paid mutator transaction binding the contract method 0x30b70002.
//
// Solidity: function handleRewards(bool _shouldClaimGmx, bool _shouldStakeGmx, bool _shouldClaimEsGmx, bool _shouldStakeEsGmx, bool _shouldStakeMultiplierPoints, bool _shouldClaimWeth, bool _shouldConvertWethToEth) returns()
func (_RewardRouter *RewardRouterTransactor) HandleRewards(opts *bind.TransactOpts, _shouldClaimGmx bool, _shouldStakeGmx bool, _shouldClaimEsGmx bool, _shouldStakeEsGmx bool, _shouldStakeMultiplierPoints bool, _shouldClaimWeth bool, _shouldConvertWethToEth bool) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "handleRewards", _shouldClaimGmx, _shouldStakeGmx, _shouldClaimEsGmx, _shouldStakeEsGmx, _shouldStakeMultiplierPoints, _shouldClaimWeth, _shouldConvertWethToEth)
}

// HandleRewards is a paid mutator transaction binding the contract method 0x30b70002.
//
// Solidity: function handleRewards(bool _shouldClaimGmx, bool _shouldStakeGmx, bool _shouldClaimEsGmx, bool _shouldStakeEsGmx, bool _shouldStakeMultiplierPoints, bool _shouldClaimWeth, bool _shouldConvertWethToEth) returns()
func (_RewardRouter *RewardRouterSession) HandleRewards(_shouldClaimGmx bool, _shouldStakeGmx bool, _shouldClaimEsGmx bool, _shouldStakeEsGmx bool, _shouldStakeMultiplierPoints bool, _shouldClaimWeth bool, _shouldConvertWethToEth bool) (*types.Transaction, error) {
	return _RewardRouter.Contract.HandleRewards(&_RewardRouter.TransactOpts, _shouldClaimGmx, _shouldStakeGmx, _shouldClaimEsGmx, _shouldStakeEsGmx, _shouldStakeMultiplierPoints, _shouldClaimWeth, _shouldConvertWethToEth)
}

// HandleRewards is a paid mutator transaction binding the contract method 0x30b70002.
//
// Solidity: function handleRewards(bool _shouldClaimGmx, bool _shouldStakeGmx, bool _shouldClaimEsGmx, bool _shouldStakeEsGmx, bool _shouldStakeMultiplierPoints, bool _shouldClaimWeth, bool _shouldConvertWethToEth) returns()
func (_RewardRouter *RewardRouterTransactorSession) HandleRewards(_shouldClaimGmx bool, _shouldStakeGmx bool, _shouldClaimEsGmx bool, _shouldStakeEsGmx bool, _shouldStakeMultiplierPoints bool, _shouldClaimWeth bool, _shouldConvertWethToEth bool) (*types.Transaction, error) {
	return _RewardRouter.Contract.HandleRewards(&_RewardRouter.TransactOpts, _shouldClaimGmx, _shouldStakeGmx, _shouldClaimEsGmx, _shouldStakeEsGmx, _shouldStakeMultiplierPoints, _shouldClaimWeth, _shouldConvertWethToEth)
}

// Initialize is a paid mutator transaction binding the contract method 0x2fdd983d.
//
// Solidity: function initialize(address _weth, address _gmx, address _esGmx, address _bnGmx, address _glp, address _stakedGmxTracker, address _bonusGmxTracker, address _feeGmxTracker, address _feeGlpTracker, address _stakedGlpTracker, address _glpManager, address _gmxVester, address _glpVester) returns()
func (_RewardRouter *RewardRouterTransactor) Initialize(opts *bind.TransactOpts, _weth common.Address, _gmx common.Address, _esGmx common.Address, _bnGmx common.Address, _glp common.Address, _stakedGmxTracker common.Address, _bonusGmxTracker common.Address, _feeGmxTracker common.Address, _feeGlpTracker common.Address, _stakedGlpTracker common.Address, _glpManager common.Address, _gmxVester common.Address, _glpVester common.Address) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "initialize", _weth, _gmx, _esGmx, _bnGmx, _glp, _stakedGmxTracker, _bonusGmxTracker, _feeGmxTracker, _feeGlpTracker, _stakedGlpTracker, _glpManager, _gmxVester, _glpVester)
}

// Initialize is a paid mutator transaction binding the contract method 0x2fdd983d.
//
// Solidity: function initialize(address _weth, address _gmx, address _esGmx, address _bnGmx, address _glp, address _stakedGmxTracker, address _bonusGmxTracker, address _feeGmxTracker, address _feeGlpTracker, address _stakedGlpTracker, address _glpManager, address _gmxVester, address _glpVester) returns()
func (_RewardRouter *RewardRouterSession) Initialize(_weth common.Address, _gmx common.Address, _esGmx common.Address, _bnGmx common.Address, _glp common.Address, _stakedGmxTracker common.Address, _bonusGmxTracker common.Address, _feeGmxTracker common.Address, _feeGlpTracker common.Address, _stakedGlpTracker common.Address, _glpManager common.Address, _gmxVester common.Address, _glpVester common.Address) (*types.Transaction, error) {
	return _RewardRouter.Contract.Initialize(&_RewardRouter.TransactOpts, _weth, _gmx, _esGmx, _bnGmx, _glp, _stakedGmxTracker, _bonusGmxTracker, _feeGmxTracker, _feeGlpTracker, _stakedGlpTracker, _glpManager, _gmxVester, _glpVester)
}

// Initialize is a paid mutator transaction binding the contract method 0x2fdd983d.
//
// Solidity: function initialize(address _weth, address _gmx, address _esGmx, address _bnGmx, address _glp, address _stakedGmxTracker, address _bonusGmxTracker, address _feeGmxTracker, address _feeGlpTracker, address _stakedGlpTracker, address _glpManager, address _gmxVester, address _glpVester) returns()
func (_RewardRouter *RewardRouterTransactorSession) Initialize(_weth common.Address, _gmx common.Address, _esGmx common.Address, _bnGmx common.Address, _glp common.Address, _stakedGmxTracker common.Address, _bonusGmxTracker common.Address, _feeGmxTracker common.Address, _feeGlpTracker common.Address, _stakedGlpTracker common.Address, _glpManager common.Address, _gmxVester common.Address, _glpVester common.Address) (*types.Transaction, error) {
	return _RewardRouter.Contract.Initialize(&_RewardRouter.TransactOpts, _weth, _gmx, _esGmx, _bnGmx, _glp, _stakedGmxTracker, _bonusGmxTracker, _feeGmxTracker, _feeGlpTracker, _stakedGlpTracker, _glpManager, _gmxVester, _glpVester)
}

// MintAndStakeGlp is a paid mutator transaction binding the contract method 0x364e2311.
//
// Solidity: function mintAndStakeGlp(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_RewardRouter *RewardRouterTransactor) MintAndStakeGlp(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "mintAndStakeGlp", _token, _amount, _minUsdg, _minGlp)
}

// MintAndStakeGlp is a paid mutator transaction binding the contract method 0x364e2311.
//
// Solidity: function mintAndStakeGlp(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_RewardRouter *RewardRouterSession) MintAndStakeGlp(_token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _RewardRouter.Contract.MintAndStakeGlp(&_RewardRouter.TransactOpts, _token, _amount, _minUsdg, _minGlp)
}

// MintAndStakeGlp is a paid mutator transaction binding the contract method 0x364e2311.
//
// Solidity: function mintAndStakeGlp(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_RewardRouter *RewardRouterTransactorSession) MintAndStakeGlp(_token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _RewardRouter.Contract.MintAndStakeGlp(&_RewardRouter.TransactOpts, _token, _amount, _minUsdg, _minGlp)
}

// MintAndStakeGlpETH is a paid mutator transaction binding the contract method 0x53a8aa03.
//
// Solidity: function mintAndStakeGlpETH(uint256 _minUsdg, uint256 _minGlp) payable returns(uint256)
func (_RewardRouter *RewardRouterTransactor) MintAndStakeGlpETH(opts *bind.TransactOpts, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "mintAndStakeGlpETH", _minUsdg, _minGlp)
}

// MintAndStakeGlpETH is a paid mutator transaction binding the contract method 0x53a8aa03.
//
// Solidity: function mintAndStakeGlpETH(uint256 _minUsdg, uint256 _minGlp) payable returns(uint256)
func (_RewardRouter *RewardRouterSession) MintAndStakeGlpETH(_minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _RewardRouter.Contract.MintAndStakeGlpETH(&_RewardRouter.TransactOpts, _minUsdg, _minGlp)
}

// MintAndStakeGlpETH is a paid mutator transaction binding the contract method 0x53a8aa03.
//
// Solidity: function mintAndStakeGlpETH(uint256 _minUsdg, uint256 _minGlp) payable returns(uint256)
func (_RewardRouter *RewardRouterTransactorSession) MintAndStakeGlpETH(_minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _RewardRouter.Contract.MintAndStakeGlpETH(&_RewardRouter.TransactOpts, _minUsdg, _minGlp)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_RewardRouter *RewardRouterTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_RewardRouter *RewardRouterSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _RewardRouter.Contract.SetGov(&_RewardRouter.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_RewardRouter *RewardRouterTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _RewardRouter.Contract.SetGov(&_RewardRouter.TransactOpts, _gov)
}

// SignalTransfer is a paid mutator transaction binding the contract method 0xef9aacfd.
//
// Solidity: function signalTransfer(address _receiver) returns()
func (_RewardRouter *RewardRouterTransactor) SignalTransfer(opts *bind.TransactOpts, _receiver common.Address) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "signalTransfer", _receiver)
}

// SignalTransfer is a paid mutator transaction binding the contract method 0xef9aacfd.
//
// Solidity: function signalTransfer(address _receiver) returns()
func (_RewardRouter *RewardRouterSession) SignalTransfer(_receiver common.Address) (*types.Transaction, error) {
	return _RewardRouter.Contract.SignalTransfer(&_RewardRouter.TransactOpts, _receiver)
}

// SignalTransfer is a paid mutator transaction binding the contract method 0xef9aacfd.
//
// Solidity: function signalTransfer(address _receiver) returns()
func (_RewardRouter *RewardRouterTransactorSession) SignalTransfer(_receiver common.Address) (*types.Transaction, error) {
	return _RewardRouter.Contract.SignalTransfer(&_RewardRouter.TransactOpts, _receiver)
}

// StakeEsGmx is a paid mutator transaction binding the contract method 0xef8c5994.
//
// Solidity: function stakeEsGmx(uint256 _amount) returns()
func (_RewardRouter *RewardRouterTransactor) StakeEsGmx(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "stakeEsGmx", _amount)
}

// StakeEsGmx is a paid mutator transaction binding the contract method 0xef8c5994.
//
// Solidity: function stakeEsGmx(uint256 _amount) returns()
func (_RewardRouter *RewardRouterSession) StakeEsGmx(_amount *big.Int) (*types.Transaction, error) {
	return _RewardRouter.Contract.StakeEsGmx(&_RewardRouter.TransactOpts, _amount)
}

// StakeEsGmx is a paid mutator transaction binding the contract method 0xef8c5994.
//
// Solidity: function stakeEsGmx(uint256 _amount) returns()
func (_RewardRouter *RewardRouterTransactorSession) StakeEsGmx(_amount *big.Int) (*types.Transaction, error) {
	return _RewardRouter.Contract.StakeEsGmx(&_RewardRouter.TransactOpts, _amount)
}

// StakeGmx is a paid mutator transaction binding the contract method 0xf3daeacc.
//
// Solidity: function stakeGmx(uint256 _amount) returns()
func (_RewardRouter *RewardRouterTransactor) StakeGmx(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "stakeGmx", _amount)
}

// StakeGmx is a paid mutator transaction binding the contract method 0xf3daeacc.
//
// Solidity: function stakeGmx(uint256 _amount) returns()
func (_RewardRouter *RewardRouterSession) StakeGmx(_amount *big.Int) (*types.Transaction, error) {
	return _RewardRouter.Contract.StakeGmx(&_RewardRouter.TransactOpts, _amount)
}

// StakeGmx is a paid mutator transaction binding the contract method 0xf3daeacc.
//
// Solidity: function stakeGmx(uint256 _amount) returns()
func (_RewardRouter *RewardRouterTransactorSession) StakeGmx(_amount *big.Int) (*types.Transaction, error) {
	return _RewardRouter.Contract.StakeGmx(&_RewardRouter.TransactOpts, _amount)
}

// StakeGmxForAccount is a paid mutator transaction binding the contract method 0x5da4b8dd.
//
// Solidity: function stakeGmxForAccount(address _account, uint256 _amount) returns()
func (_RewardRouter *RewardRouterTransactor) StakeGmxForAccount(opts *bind.TransactOpts, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "stakeGmxForAccount", _account, _amount)
}

// StakeGmxForAccount is a paid mutator transaction binding the contract method 0x5da4b8dd.
//
// Solidity: function stakeGmxForAccount(address _account, uint256 _amount) returns()
func (_RewardRouter *RewardRouterSession) StakeGmxForAccount(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardRouter.Contract.StakeGmxForAccount(&_RewardRouter.TransactOpts, _account, _amount)
}

// StakeGmxForAccount is a paid mutator transaction binding the contract method 0x5da4b8dd.
//
// Solidity: function stakeGmxForAccount(address _account, uint256 _amount) returns()
func (_RewardRouter *RewardRouterTransactorSession) StakeGmxForAccount(_account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardRouter.Contract.StakeGmxForAccount(&_RewardRouter.TransactOpts, _account, _amount)
}

// UnstakeAndRedeemGlp is a paid mutator transaction binding the contract method 0x0f3aa554.
//
// Solidity: function unstakeAndRedeemGlp(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_RewardRouter *RewardRouterTransactor) UnstakeAndRedeemGlp(opts *bind.TransactOpts, _tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "unstakeAndRedeemGlp", _tokenOut, _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlp is a paid mutator transaction binding the contract method 0x0f3aa554.
//
// Solidity: function unstakeAndRedeemGlp(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_RewardRouter *RewardRouterSession) UnstakeAndRedeemGlp(_tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _RewardRouter.Contract.UnstakeAndRedeemGlp(&_RewardRouter.TransactOpts, _tokenOut, _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlp is a paid mutator transaction binding the contract method 0x0f3aa554.
//
// Solidity: function unstakeAndRedeemGlp(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_RewardRouter *RewardRouterTransactorSession) UnstakeAndRedeemGlp(_tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _RewardRouter.Contract.UnstakeAndRedeemGlp(&_RewardRouter.TransactOpts, _tokenOut, _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlpETH is a paid mutator transaction binding the contract method 0xabb5e5e2.
//
// Solidity: function unstakeAndRedeemGlpETH(uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_RewardRouter *RewardRouterTransactor) UnstakeAndRedeemGlpETH(opts *bind.TransactOpts, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "unstakeAndRedeemGlpETH", _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlpETH is a paid mutator transaction binding the contract method 0xabb5e5e2.
//
// Solidity: function unstakeAndRedeemGlpETH(uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_RewardRouter *RewardRouterSession) UnstakeAndRedeemGlpETH(_glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _RewardRouter.Contract.UnstakeAndRedeemGlpETH(&_RewardRouter.TransactOpts, _glpAmount, _minOut, _receiver)
}

// UnstakeAndRedeemGlpETH is a paid mutator transaction binding the contract method 0xabb5e5e2.
//
// Solidity: function unstakeAndRedeemGlpETH(uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_RewardRouter *RewardRouterTransactorSession) UnstakeAndRedeemGlpETH(_glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _RewardRouter.Contract.UnstakeAndRedeemGlpETH(&_RewardRouter.TransactOpts, _glpAmount, _minOut, _receiver)
}

// UnstakeEsGmx is a paid mutator transaction binding the contract method 0x64f64467.
//
// Solidity: function unstakeEsGmx(uint256 _amount) returns()
func (_RewardRouter *RewardRouterTransactor) UnstakeEsGmx(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "unstakeEsGmx", _amount)
}

// UnstakeEsGmx is a paid mutator transaction binding the contract method 0x64f64467.
//
// Solidity: function unstakeEsGmx(uint256 _amount) returns()
func (_RewardRouter *RewardRouterSession) UnstakeEsGmx(_amount *big.Int) (*types.Transaction, error) {
	return _RewardRouter.Contract.UnstakeEsGmx(&_RewardRouter.TransactOpts, _amount)
}

// UnstakeEsGmx is a paid mutator transaction binding the contract method 0x64f64467.
//
// Solidity: function unstakeEsGmx(uint256 _amount) returns()
func (_RewardRouter *RewardRouterTransactorSession) UnstakeEsGmx(_amount *big.Int) (*types.Transaction, error) {
	return _RewardRouter.Contract.UnstakeEsGmx(&_RewardRouter.TransactOpts, _amount)
}

// UnstakeGmx is a paid mutator transaction binding the contract method 0x078580d2.
//
// Solidity: function unstakeGmx(uint256 _amount) returns()
func (_RewardRouter *RewardRouterTransactor) UnstakeGmx(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "unstakeGmx", _amount)
}

// UnstakeGmx is a paid mutator transaction binding the contract method 0x078580d2.
//
// Solidity: function unstakeGmx(uint256 _amount) returns()
func (_RewardRouter *RewardRouterSession) UnstakeGmx(_amount *big.Int) (*types.Transaction, error) {
	return _RewardRouter.Contract.UnstakeGmx(&_RewardRouter.TransactOpts, _amount)
}

// UnstakeGmx is a paid mutator transaction binding the contract method 0x078580d2.
//
// Solidity: function unstakeGmx(uint256 _amount) returns()
func (_RewardRouter *RewardRouterTransactorSession) UnstakeGmx(_amount *big.Int) (*types.Transaction, error) {
	return _RewardRouter.Contract.UnstakeGmx(&_RewardRouter.TransactOpts, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_RewardRouter *RewardRouterTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardRouter.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_RewardRouter *RewardRouterSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardRouter.Contract.WithdrawToken(&_RewardRouter.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_RewardRouter *RewardRouterTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardRouter.Contract.WithdrawToken(&_RewardRouter.TransactOpts, _token, _account, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewardRouter *RewardRouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardRouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewardRouter *RewardRouterSession) Receive() (*types.Transaction, error) {
	return _RewardRouter.Contract.Receive(&_RewardRouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RewardRouter *RewardRouterTransactorSession) Receive() (*types.Transaction, error) {
	return _RewardRouter.Contract.Receive(&_RewardRouter.TransactOpts)
}

// RewardRouterStakeGlpIterator is returned from FilterStakeGlp and is used to iterate over the raw logs and unpacked data for StakeGlp events raised by the RewardRouter contract.
type RewardRouterStakeGlpIterator struct {
	Event *RewardRouterStakeGlp // Event containing the contract specifics and raw log

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
func (it *RewardRouterStakeGlpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardRouterStakeGlp)
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
		it.Event = new(RewardRouterStakeGlp)
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
func (it *RewardRouterStakeGlpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardRouterStakeGlpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardRouterStakeGlp represents a StakeGlp event raised by the RewardRouter contract.
type RewardRouterStakeGlp struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeGlp is a free log retrieval operation binding the contract event 0xa4725d47fa458d9222498e4d63f34527cf7318c1506f89d9092b35fdbcb64f3a.
//
// Solidity: event StakeGlp(address account, uint256 amount)
func (_RewardRouter *RewardRouterFilterer) FilterStakeGlp(opts *bind.FilterOpts) (*RewardRouterStakeGlpIterator, error) {

	logs, sub, err := _RewardRouter.contract.FilterLogs(opts, "StakeGlp")
	if err != nil {
		return nil, err
	}
	return &RewardRouterStakeGlpIterator{contract: _RewardRouter.contract, event: "StakeGlp", logs: logs, sub: sub}, nil
}

// WatchStakeGlp is a free log subscription operation binding the contract event 0xa4725d47fa458d9222498e4d63f34527cf7318c1506f89d9092b35fdbcb64f3a.
//
// Solidity: event StakeGlp(address account, uint256 amount)
func (_RewardRouter *RewardRouterFilterer) WatchStakeGlp(opts *bind.WatchOpts, sink chan<- *RewardRouterStakeGlp) (event.Subscription, error) {

	logs, sub, err := _RewardRouter.contract.WatchLogs(opts, "StakeGlp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardRouterStakeGlp)
				if err := _RewardRouter.contract.UnpackLog(event, "StakeGlp", log); err != nil {
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

// ParseStakeGlp is a log parse operation binding the contract event 0xa4725d47fa458d9222498e4d63f34527cf7318c1506f89d9092b35fdbcb64f3a.
//
// Solidity: event StakeGlp(address account, uint256 amount)
func (_RewardRouter *RewardRouterFilterer) ParseStakeGlp(log types.Log) (*RewardRouterStakeGlp, error) {
	event := new(RewardRouterStakeGlp)
	if err := _RewardRouter.contract.UnpackLog(event, "StakeGlp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardRouterStakeGmxIterator is returned from FilterStakeGmx and is used to iterate over the raw logs and unpacked data for StakeGmx events raised by the RewardRouter contract.
type RewardRouterStakeGmxIterator struct {
	Event *RewardRouterStakeGmx // Event containing the contract specifics and raw log

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
func (it *RewardRouterStakeGmxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardRouterStakeGmx)
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
		it.Event = new(RewardRouterStakeGmx)
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
func (it *RewardRouterStakeGmxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardRouterStakeGmxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardRouterStakeGmx represents a StakeGmx event raised by the RewardRouter contract.
type RewardRouterStakeGmx struct {
	Account common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeGmx is a free log retrieval operation binding the contract event 0xad0723806aa1e5a8fb826fc9f0c5b589e585a6b60dc768a1b20691c95062d2d6.
//
// Solidity: event StakeGmx(address account, address token, uint256 amount)
func (_RewardRouter *RewardRouterFilterer) FilterStakeGmx(opts *bind.FilterOpts) (*RewardRouterStakeGmxIterator, error) {

	logs, sub, err := _RewardRouter.contract.FilterLogs(opts, "StakeGmx")
	if err != nil {
		return nil, err
	}
	return &RewardRouterStakeGmxIterator{contract: _RewardRouter.contract, event: "StakeGmx", logs: logs, sub: sub}, nil
}

// WatchStakeGmx is a free log subscription operation binding the contract event 0xad0723806aa1e5a8fb826fc9f0c5b589e585a6b60dc768a1b20691c95062d2d6.
//
// Solidity: event StakeGmx(address account, address token, uint256 amount)
func (_RewardRouter *RewardRouterFilterer) WatchStakeGmx(opts *bind.WatchOpts, sink chan<- *RewardRouterStakeGmx) (event.Subscription, error) {

	logs, sub, err := _RewardRouter.contract.WatchLogs(opts, "StakeGmx")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardRouterStakeGmx)
				if err := _RewardRouter.contract.UnpackLog(event, "StakeGmx", log); err != nil {
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

// ParseStakeGmx is a log parse operation binding the contract event 0xad0723806aa1e5a8fb826fc9f0c5b589e585a6b60dc768a1b20691c95062d2d6.
//
// Solidity: event StakeGmx(address account, address token, uint256 amount)
func (_RewardRouter *RewardRouterFilterer) ParseStakeGmx(log types.Log) (*RewardRouterStakeGmx, error) {
	event := new(RewardRouterStakeGmx)
	if err := _RewardRouter.contract.UnpackLog(event, "StakeGmx", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardRouterUnstakeGlpIterator is returned from FilterUnstakeGlp and is used to iterate over the raw logs and unpacked data for UnstakeGlp events raised by the RewardRouter contract.
type RewardRouterUnstakeGlpIterator struct {
	Event *RewardRouterUnstakeGlp // Event containing the contract specifics and raw log

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
func (it *RewardRouterUnstakeGlpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardRouterUnstakeGlp)
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
		it.Event = new(RewardRouterUnstakeGlp)
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
func (it *RewardRouterUnstakeGlpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardRouterUnstakeGlpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardRouterUnstakeGlp represents a UnstakeGlp event raised by the RewardRouter contract.
type RewardRouterUnstakeGlp struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstakeGlp is a free log retrieval operation binding the contract event 0x1cb6202519b6b6c72ba5ed11e2c3f53af3cea010f96bfc705584e53e75cf034c.
//
// Solidity: event UnstakeGlp(address account, uint256 amount)
func (_RewardRouter *RewardRouterFilterer) FilterUnstakeGlp(opts *bind.FilterOpts) (*RewardRouterUnstakeGlpIterator, error) {

	logs, sub, err := _RewardRouter.contract.FilterLogs(opts, "UnstakeGlp")
	if err != nil {
		return nil, err
	}
	return &RewardRouterUnstakeGlpIterator{contract: _RewardRouter.contract, event: "UnstakeGlp", logs: logs, sub: sub}, nil
}

// WatchUnstakeGlp is a free log subscription operation binding the contract event 0x1cb6202519b6b6c72ba5ed11e2c3f53af3cea010f96bfc705584e53e75cf034c.
//
// Solidity: event UnstakeGlp(address account, uint256 amount)
func (_RewardRouter *RewardRouterFilterer) WatchUnstakeGlp(opts *bind.WatchOpts, sink chan<- *RewardRouterUnstakeGlp) (event.Subscription, error) {

	logs, sub, err := _RewardRouter.contract.WatchLogs(opts, "UnstakeGlp")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardRouterUnstakeGlp)
				if err := _RewardRouter.contract.UnpackLog(event, "UnstakeGlp", log); err != nil {
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

// ParseUnstakeGlp is a log parse operation binding the contract event 0x1cb6202519b6b6c72ba5ed11e2c3f53af3cea010f96bfc705584e53e75cf034c.
//
// Solidity: event UnstakeGlp(address account, uint256 amount)
func (_RewardRouter *RewardRouterFilterer) ParseUnstakeGlp(log types.Log) (*RewardRouterUnstakeGlp, error) {
	event := new(RewardRouterUnstakeGlp)
	if err := _RewardRouter.contract.UnpackLog(event, "UnstakeGlp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardRouterUnstakeGmxIterator is returned from FilterUnstakeGmx and is used to iterate over the raw logs and unpacked data for UnstakeGmx events raised by the RewardRouter contract.
type RewardRouterUnstakeGmxIterator struct {
	Event *RewardRouterUnstakeGmx // Event containing the contract specifics and raw log

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
func (it *RewardRouterUnstakeGmxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardRouterUnstakeGmx)
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
		it.Event = new(RewardRouterUnstakeGmx)
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
func (it *RewardRouterUnstakeGmxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardRouterUnstakeGmxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardRouterUnstakeGmx represents a UnstakeGmx event raised by the RewardRouter contract.
type RewardRouterUnstakeGmx struct {
	Account common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstakeGmx is a free log retrieval operation binding the contract event 0xce8eb393006add0768cc6cefb3ca0fc4787015ce1ac86bd800e72a7999310345.
//
// Solidity: event UnstakeGmx(address account, address token, uint256 amount)
func (_RewardRouter *RewardRouterFilterer) FilterUnstakeGmx(opts *bind.FilterOpts) (*RewardRouterUnstakeGmxIterator, error) {

	logs, sub, err := _RewardRouter.contract.FilterLogs(opts, "UnstakeGmx")
	if err != nil {
		return nil, err
	}
	return &RewardRouterUnstakeGmxIterator{contract: _RewardRouter.contract, event: "UnstakeGmx", logs: logs, sub: sub}, nil
}

// WatchUnstakeGmx is a free log subscription operation binding the contract event 0xce8eb393006add0768cc6cefb3ca0fc4787015ce1ac86bd800e72a7999310345.
//
// Solidity: event UnstakeGmx(address account, address token, uint256 amount)
func (_RewardRouter *RewardRouterFilterer) WatchUnstakeGmx(opts *bind.WatchOpts, sink chan<- *RewardRouterUnstakeGmx) (event.Subscription, error) {

	logs, sub, err := _RewardRouter.contract.WatchLogs(opts, "UnstakeGmx")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardRouterUnstakeGmx)
				if err := _RewardRouter.contract.UnpackLog(event, "UnstakeGmx", log); err != nil {
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

// ParseUnstakeGmx is a log parse operation binding the contract event 0xce8eb393006add0768cc6cefb3ca0fc4787015ce1ac86bd800e72a7999310345.
//
// Solidity: event UnstakeGmx(address account, address token, uint256 amount)
func (_RewardRouter *RewardRouterFilterer) ParseUnstakeGmx(log types.Log) (*RewardRouterUnstakeGmx, error) {
	event := new(RewardRouterUnstakeGmx)
	if err := _RewardRouter.contract.UnpackLog(event, "UnstakeGmx", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

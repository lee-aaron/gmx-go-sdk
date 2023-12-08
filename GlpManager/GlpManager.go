// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GlpManager

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

// GlpManagerMetaData contains all meta data concerning the GlpManager contract.
var GlpManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdg\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_glp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_shortsTracker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_cooldownDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BASIS_POINTS_DIVISOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"GLP_PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_COOLDOWN_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PRICE_PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"USDG_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addLiquidity\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minUsdg\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minGlp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addLiquidityForAccount\",\"inputs\":[{\"name\":\"_fundingAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minUsdg\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minGlp\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"aumAddition\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"aumDeduction\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cooldownDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAum\",\"inputs\":[{\"name\":\"maximise\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAumInUsdg\",\"inputs\":[{\"name\":\"maximise\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAums\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalShortAveragePrice\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalShortDelta\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_size\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPrice\",\"inputs\":[{\"name\":\"_maximise\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"glp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inPrivateMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isHandler\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastAddedAt\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeLiquidity\",\"inputs\":[{\"name\":\"_tokenOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_glpAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeLiquidityForAccount\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_glpAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAumAdjustment\",\"inputs\":[{\"name\":\"_aumAddition\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_aumDeduction\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCooldownDuration\",\"inputs\":[{\"name\":\"_cooldownDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setHandler\",\"inputs\":[{\"name\":\"_handler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInPrivateMode\",\"inputs\":[{\"name\":\"_inPrivateMode\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setShortsTracker\",\"inputs\":[{\"name\":\"_shortsTracker\",\"type\":\"address\",\"internalType\":\"contractIShortsTracker\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setShortsTrackerAveragePriceWeight\",\"inputs\":[{\"name\":\"_shortsTrackerAveragePriceWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shortsTracker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIShortsTracker\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"shortsTrackerAveragePriceWeight\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdg\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vault\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIVault\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AddLiquidity\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"aumInUsdg\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"glpSupply\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"usdgAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"mintAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemoveLiquidity\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"glpAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"aumInUsdg\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"glpSupply\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"usdgAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// GlpManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use GlpManagerMetaData.ABI instead.
var GlpManagerABI = GlpManagerMetaData.ABI

// GlpManager is an auto generated Go binding around an Ethereum contract.
type GlpManager struct {
	GlpManagerCaller     // Read-only binding to the contract
	GlpManagerTransactor // Write-only binding to the contract
	GlpManagerFilterer   // Log filterer for contract events
}

// GlpManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlpManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlpManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlpManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlpManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlpManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlpManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlpManagerSession struct {
	Contract     *GlpManager       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GlpManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlpManagerCallerSession struct {
	Contract *GlpManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GlpManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlpManagerTransactorSession struct {
	Contract     *GlpManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GlpManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlpManagerRaw struct {
	Contract *GlpManager // Generic contract binding to access the raw methods on
}

// GlpManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlpManagerCallerRaw struct {
	Contract *GlpManagerCaller // Generic read-only contract binding to access the raw methods on
}

// GlpManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlpManagerTransactorRaw struct {
	Contract *GlpManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlpManager creates a new instance of GlpManager, bound to a specific deployed contract.
func NewGlpManager(address common.Address, backend bind.ContractBackend) (*GlpManager, error) {
	contract, err := bindGlpManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlpManager{GlpManagerCaller: GlpManagerCaller{contract: contract}, GlpManagerTransactor: GlpManagerTransactor{contract: contract}, GlpManagerFilterer: GlpManagerFilterer{contract: contract}}, nil
}

// NewGlpManagerCaller creates a new read-only instance of GlpManager, bound to a specific deployed contract.
func NewGlpManagerCaller(address common.Address, caller bind.ContractCaller) (*GlpManagerCaller, error) {
	contract, err := bindGlpManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlpManagerCaller{contract: contract}, nil
}

// NewGlpManagerTransactor creates a new write-only instance of GlpManager, bound to a specific deployed contract.
func NewGlpManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*GlpManagerTransactor, error) {
	contract, err := bindGlpManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlpManagerTransactor{contract: contract}, nil
}

// NewGlpManagerFilterer creates a new log filterer instance of GlpManager, bound to a specific deployed contract.
func NewGlpManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*GlpManagerFilterer, error) {
	contract, err := bindGlpManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlpManagerFilterer{contract: contract}, nil
}

// bindGlpManager binds a generic wrapper to an already deployed contract.
func bindGlpManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GlpManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlpManager *GlpManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GlpManager.Contract.GlpManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlpManager *GlpManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlpManager.Contract.GlpManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlpManager *GlpManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlpManager.Contract.GlpManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlpManager *GlpManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GlpManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlpManager *GlpManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlpManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlpManager *GlpManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlpManager.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_GlpManager *GlpManagerCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_GlpManager *GlpManagerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _GlpManager.Contract.BASISPOINTSDIVISOR(&_GlpManager.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_GlpManager *GlpManagerCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _GlpManager.Contract.BASISPOINTSDIVISOR(&_GlpManager.CallOpts)
}

// GLPPRECISION is a free data retrieval call binding the contract method 0x662f1c68.
//
// Solidity: function GLP_PRECISION() view returns(uint256)
func (_GlpManager *GlpManagerCaller) GLPPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "GLP_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GLPPRECISION is a free data retrieval call binding the contract method 0x662f1c68.
//
// Solidity: function GLP_PRECISION() view returns(uint256)
func (_GlpManager *GlpManagerSession) GLPPRECISION() (*big.Int, error) {
	return _GlpManager.Contract.GLPPRECISION(&_GlpManager.CallOpts)
}

// GLPPRECISION is a free data retrieval call binding the contract method 0x662f1c68.
//
// Solidity: function GLP_PRECISION() view returns(uint256)
func (_GlpManager *GlpManagerCallerSession) GLPPRECISION() (*big.Int, error) {
	return _GlpManager.Contract.GLPPRECISION(&_GlpManager.CallOpts)
}

// MAXCOOLDOWNDURATION is a free data retrieval call binding the contract method 0x1e9049cf.
//
// Solidity: function MAX_COOLDOWN_DURATION() view returns(uint256)
func (_GlpManager *GlpManagerCaller) MAXCOOLDOWNDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "MAX_COOLDOWN_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCOOLDOWNDURATION is a free data retrieval call binding the contract method 0x1e9049cf.
//
// Solidity: function MAX_COOLDOWN_DURATION() view returns(uint256)
func (_GlpManager *GlpManagerSession) MAXCOOLDOWNDURATION() (*big.Int, error) {
	return _GlpManager.Contract.MAXCOOLDOWNDURATION(&_GlpManager.CallOpts)
}

// MAXCOOLDOWNDURATION is a free data retrieval call binding the contract method 0x1e9049cf.
//
// Solidity: function MAX_COOLDOWN_DURATION() view returns(uint256)
func (_GlpManager *GlpManagerCallerSession) MAXCOOLDOWNDURATION() (*big.Int, error) {
	return _GlpManager.Contract.MAXCOOLDOWNDURATION(&_GlpManager.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_GlpManager *GlpManagerCaller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_GlpManager *GlpManagerSession) PRICEPRECISION() (*big.Int, error) {
	return _GlpManager.Contract.PRICEPRECISION(&_GlpManager.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_GlpManager *GlpManagerCallerSession) PRICEPRECISION() (*big.Int, error) {
	return _GlpManager.Contract.PRICEPRECISION(&_GlpManager.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_GlpManager *GlpManagerCaller) USDGDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "USDG_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_GlpManager *GlpManagerSession) USDGDECIMALS() (*big.Int, error) {
	return _GlpManager.Contract.USDGDECIMALS(&_GlpManager.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_GlpManager *GlpManagerCallerSession) USDGDECIMALS() (*big.Int, error) {
	return _GlpManager.Contract.USDGDECIMALS(&_GlpManager.CallOpts)
}

// AumAddition is a free data retrieval call binding the contract method 0x196b68cb.
//
// Solidity: function aumAddition() view returns(uint256)
func (_GlpManager *GlpManagerCaller) AumAddition(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "aumAddition")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AumAddition is a free data retrieval call binding the contract method 0x196b68cb.
//
// Solidity: function aumAddition() view returns(uint256)
func (_GlpManager *GlpManagerSession) AumAddition() (*big.Int, error) {
	return _GlpManager.Contract.AumAddition(&_GlpManager.CallOpts)
}

// AumAddition is a free data retrieval call binding the contract method 0x196b68cb.
//
// Solidity: function aumAddition() view returns(uint256)
func (_GlpManager *GlpManagerCallerSession) AumAddition() (*big.Int, error) {
	return _GlpManager.Contract.AumAddition(&_GlpManager.CallOpts)
}

// AumDeduction is a free data retrieval call binding the contract method 0xb172bb0c.
//
// Solidity: function aumDeduction() view returns(uint256)
func (_GlpManager *GlpManagerCaller) AumDeduction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "aumDeduction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AumDeduction is a free data retrieval call binding the contract method 0xb172bb0c.
//
// Solidity: function aumDeduction() view returns(uint256)
func (_GlpManager *GlpManagerSession) AumDeduction() (*big.Int, error) {
	return _GlpManager.Contract.AumDeduction(&_GlpManager.CallOpts)
}

// AumDeduction is a free data retrieval call binding the contract method 0xb172bb0c.
//
// Solidity: function aumDeduction() view returns(uint256)
func (_GlpManager *GlpManagerCallerSession) AumDeduction() (*big.Int, error) {
	return _GlpManager.Contract.AumDeduction(&_GlpManager.CallOpts)
}

// CooldownDuration is a free data retrieval call binding the contract method 0x35269315.
//
// Solidity: function cooldownDuration() view returns(uint256)
func (_GlpManager *GlpManagerCaller) CooldownDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "cooldownDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CooldownDuration is a free data retrieval call binding the contract method 0x35269315.
//
// Solidity: function cooldownDuration() view returns(uint256)
func (_GlpManager *GlpManagerSession) CooldownDuration() (*big.Int, error) {
	return _GlpManager.Contract.CooldownDuration(&_GlpManager.CallOpts)
}

// CooldownDuration is a free data retrieval call binding the contract method 0x35269315.
//
// Solidity: function cooldownDuration() view returns(uint256)
func (_GlpManager *GlpManagerCallerSession) CooldownDuration() (*big.Int, error) {
	return _GlpManager.Contract.CooldownDuration(&_GlpManager.CallOpts)
}

// GetAum is a free data retrieval call binding the contract method 0x03391476.
//
// Solidity: function getAum(bool maximise) view returns(uint256)
func (_GlpManager *GlpManagerCaller) GetAum(opts *bind.CallOpts, maximise bool) (*big.Int, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "getAum", maximise)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAum is a free data retrieval call binding the contract method 0x03391476.
//
// Solidity: function getAum(bool maximise) view returns(uint256)
func (_GlpManager *GlpManagerSession) GetAum(maximise bool) (*big.Int, error) {
	return _GlpManager.Contract.GetAum(&_GlpManager.CallOpts, maximise)
}

// GetAum is a free data retrieval call binding the contract method 0x03391476.
//
// Solidity: function getAum(bool maximise) view returns(uint256)
func (_GlpManager *GlpManagerCallerSession) GetAum(maximise bool) (*big.Int, error) {
	return _GlpManager.Contract.GetAum(&_GlpManager.CallOpts, maximise)
}

// GetAumInUsdg is a free data retrieval call binding the contract method 0x68a0a3e0.
//
// Solidity: function getAumInUsdg(bool maximise) view returns(uint256)
func (_GlpManager *GlpManagerCaller) GetAumInUsdg(opts *bind.CallOpts, maximise bool) (*big.Int, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "getAumInUsdg", maximise)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAumInUsdg is a free data retrieval call binding the contract method 0x68a0a3e0.
//
// Solidity: function getAumInUsdg(bool maximise) view returns(uint256)
func (_GlpManager *GlpManagerSession) GetAumInUsdg(maximise bool) (*big.Int, error) {
	return _GlpManager.Contract.GetAumInUsdg(&_GlpManager.CallOpts, maximise)
}

// GetAumInUsdg is a free data retrieval call binding the contract method 0x68a0a3e0.
//
// Solidity: function getAumInUsdg(bool maximise) view returns(uint256)
func (_GlpManager *GlpManagerCallerSession) GetAumInUsdg(maximise bool) (*big.Int, error) {
	return _GlpManager.Contract.GetAumInUsdg(&_GlpManager.CallOpts, maximise)
}

// GetAums is a free data retrieval call binding the contract method 0xed0d1c04.
//
// Solidity: function getAums() view returns(uint256[])
func (_GlpManager *GlpManagerCaller) GetAums(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "getAums")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAums is a free data retrieval call binding the contract method 0xed0d1c04.
//
// Solidity: function getAums() view returns(uint256[])
func (_GlpManager *GlpManagerSession) GetAums() ([]*big.Int, error) {
	return _GlpManager.Contract.GetAums(&_GlpManager.CallOpts)
}

// GetAums is a free data retrieval call binding the contract method 0xed0d1c04.
//
// Solidity: function getAums() view returns(uint256[])
func (_GlpManager *GlpManagerCallerSession) GetAums() ([]*big.Int, error) {
	return _GlpManager.Contract.GetAums(&_GlpManager.CallOpts)
}

// GetGlobalShortAveragePrice is a free data retrieval call binding the contract method 0x440d828a.
//
// Solidity: function getGlobalShortAveragePrice(address _token) view returns(uint256)
func (_GlpManager *GlpManagerCaller) GetGlobalShortAveragePrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "getGlobalShortAveragePrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGlobalShortAveragePrice is a free data retrieval call binding the contract method 0x440d828a.
//
// Solidity: function getGlobalShortAveragePrice(address _token) view returns(uint256)
func (_GlpManager *GlpManagerSession) GetGlobalShortAveragePrice(_token common.Address) (*big.Int, error) {
	return _GlpManager.Contract.GetGlobalShortAveragePrice(&_GlpManager.CallOpts, _token)
}

// GetGlobalShortAveragePrice is a free data retrieval call binding the contract method 0x440d828a.
//
// Solidity: function getGlobalShortAveragePrice(address _token) view returns(uint256)
func (_GlpManager *GlpManagerCallerSession) GetGlobalShortAveragePrice(_token common.Address) (*big.Int, error) {
	return _GlpManager.Contract.GetGlobalShortAveragePrice(&_GlpManager.CallOpts, _token)
}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xa1acd3d5.
//
// Solidity: function getGlobalShortDelta(address _token, uint256 _price, uint256 _size) view returns(uint256, bool)
func (_GlpManager *GlpManagerCaller) GetGlobalShortDelta(opts *bind.CallOpts, _token common.Address, _price *big.Int, _size *big.Int) (*big.Int, bool, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "getGlobalShortDelta", _token, _price, _size)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xa1acd3d5.
//
// Solidity: function getGlobalShortDelta(address _token, uint256 _price, uint256 _size) view returns(uint256, bool)
func (_GlpManager *GlpManagerSession) GetGlobalShortDelta(_token common.Address, _price *big.Int, _size *big.Int) (*big.Int, bool, error) {
	return _GlpManager.Contract.GetGlobalShortDelta(&_GlpManager.CallOpts, _token, _price, _size)
}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xa1acd3d5.
//
// Solidity: function getGlobalShortDelta(address _token, uint256 _price, uint256 _size) view returns(uint256, bool)
func (_GlpManager *GlpManagerCallerSession) GetGlobalShortDelta(_token common.Address, _price *big.Int, _size *big.Int) (*big.Int, bool, error) {
	return _GlpManager.Contract.GetGlobalShortDelta(&_GlpManager.CallOpts, _token, _price, _size)
}

// GetPrice is a free data retrieval call binding the contract method 0xe245b5af.
//
// Solidity: function getPrice(bool _maximise) view returns(uint256)
func (_GlpManager *GlpManagerCaller) GetPrice(opts *bind.CallOpts, _maximise bool) (*big.Int, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "getPrice", _maximise)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0xe245b5af.
//
// Solidity: function getPrice(bool _maximise) view returns(uint256)
func (_GlpManager *GlpManagerSession) GetPrice(_maximise bool) (*big.Int, error) {
	return _GlpManager.Contract.GetPrice(&_GlpManager.CallOpts, _maximise)
}

// GetPrice is a free data retrieval call binding the contract method 0xe245b5af.
//
// Solidity: function getPrice(bool _maximise) view returns(uint256)
func (_GlpManager *GlpManagerCallerSession) GetPrice(_maximise bool) (*big.Int, error) {
	return _GlpManager.Contract.GetPrice(&_GlpManager.CallOpts, _maximise)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_GlpManager *GlpManagerCaller) Glp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "glp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_GlpManager *GlpManagerSession) Glp() (common.Address, error) {
	return _GlpManager.Contract.Glp(&_GlpManager.CallOpts)
}

// Glp is a free data retrieval call binding the contract method 0x78a207ee.
//
// Solidity: function glp() view returns(address)
func (_GlpManager *GlpManagerCallerSession) Glp() (common.Address, error) {
	return _GlpManager.Contract.Glp(&_GlpManager.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_GlpManager *GlpManagerCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_GlpManager *GlpManagerSession) Gov() (common.Address, error) {
	return _GlpManager.Contract.Gov(&_GlpManager.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_GlpManager *GlpManagerCallerSession) Gov() (common.Address, error) {
	return _GlpManager.Contract.Gov(&_GlpManager.CallOpts)
}

// InPrivateMode is a free data retrieval call binding the contract method 0x070eacee.
//
// Solidity: function inPrivateMode() view returns(bool)
func (_GlpManager *GlpManagerCaller) InPrivateMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "inPrivateMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InPrivateMode is a free data retrieval call binding the contract method 0x070eacee.
//
// Solidity: function inPrivateMode() view returns(bool)
func (_GlpManager *GlpManagerSession) InPrivateMode() (bool, error) {
	return _GlpManager.Contract.InPrivateMode(&_GlpManager.CallOpts)
}

// InPrivateMode is a free data retrieval call binding the contract method 0x070eacee.
//
// Solidity: function inPrivateMode() view returns(bool)
func (_GlpManager *GlpManagerCallerSession) InPrivateMode() (bool, error) {
	return _GlpManager.Contract.InPrivateMode(&_GlpManager.CallOpts)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_GlpManager *GlpManagerCaller) IsHandler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "isHandler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_GlpManager *GlpManagerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _GlpManager.Contract.IsHandler(&_GlpManager.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_GlpManager *GlpManagerCallerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _GlpManager.Contract.IsHandler(&_GlpManager.CallOpts, arg0)
}

// LastAddedAt is a free data retrieval call binding the contract method 0x8b770e11.
//
// Solidity: function lastAddedAt(address ) view returns(uint256)
func (_GlpManager *GlpManagerCaller) LastAddedAt(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "lastAddedAt", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastAddedAt is a free data retrieval call binding the contract method 0x8b770e11.
//
// Solidity: function lastAddedAt(address ) view returns(uint256)
func (_GlpManager *GlpManagerSession) LastAddedAt(arg0 common.Address) (*big.Int, error) {
	return _GlpManager.Contract.LastAddedAt(&_GlpManager.CallOpts, arg0)
}

// LastAddedAt is a free data retrieval call binding the contract method 0x8b770e11.
//
// Solidity: function lastAddedAt(address ) view returns(uint256)
func (_GlpManager *GlpManagerCallerSession) LastAddedAt(arg0 common.Address) (*big.Int, error) {
	return _GlpManager.Contract.LastAddedAt(&_GlpManager.CallOpts, arg0)
}

// ShortsTracker is a free data retrieval call binding the contract method 0x657bc5d0.
//
// Solidity: function shortsTracker() view returns(address)
func (_GlpManager *GlpManagerCaller) ShortsTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "shortsTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ShortsTracker is a free data retrieval call binding the contract method 0x657bc5d0.
//
// Solidity: function shortsTracker() view returns(address)
func (_GlpManager *GlpManagerSession) ShortsTracker() (common.Address, error) {
	return _GlpManager.Contract.ShortsTracker(&_GlpManager.CallOpts)
}

// ShortsTracker is a free data retrieval call binding the contract method 0x657bc5d0.
//
// Solidity: function shortsTracker() view returns(address)
func (_GlpManager *GlpManagerCallerSession) ShortsTracker() (common.Address, error) {
	return _GlpManager.Contract.ShortsTracker(&_GlpManager.CallOpts)
}

// ShortsTrackerAveragePriceWeight is a free data retrieval call binding the contract method 0x64e6617f.
//
// Solidity: function shortsTrackerAveragePriceWeight() view returns(uint256)
func (_GlpManager *GlpManagerCaller) ShortsTrackerAveragePriceWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "shortsTrackerAveragePriceWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ShortsTrackerAveragePriceWeight is a free data retrieval call binding the contract method 0x64e6617f.
//
// Solidity: function shortsTrackerAveragePriceWeight() view returns(uint256)
func (_GlpManager *GlpManagerSession) ShortsTrackerAveragePriceWeight() (*big.Int, error) {
	return _GlpManager.Contract.ShortsTrackerAveragePriceWeight(&_GlpManager.CallOpts)
}

// ShortsTrackerAveragePriceWeight is a free data retrieval call binding the contract method 0x64e6617f.
//
// Solidity: function shortsTrackerAveragePriceWeight() view returns(uint256)
func (_GlpManager *GlpManagerCallerSession) ShortsTrackerAveragePriceWeight() (*big.Int, error) {
	return _GlpManager.Contract.ShortsTrackerAveragePriceWeight(&_GlpManager.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_GlpManager *GlpManagerCaller) Usdg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "usdg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_GlpManager *GlpManagerSession) Usdg() (common.Address, error) {
	return _GlpManager.Contract.Usdg(&_GlpManager.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_GlpManager *GlpManagerCallerSession) Usdg() (common.Address, error) {
	return _GlpManager.Contract.Usdg(&_GlpManager.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_GlpManager *GlpManagerCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GlpManager.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_GlpManager *GlpManagerSession) Vault() (common.Address, error) {
	return _GlpManager.Contract.Vault(&_GlpManager.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_GlpManager *GlpManagerCallerSession) Vault() (common.Address, error) {
	return _GlpManager.Contract.Vault(&_GlpManager.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ece366a.
//
// Solidity: function addLiquidity(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_GlpManager *GlpManagerTransactor) AddLiquidity(opts *bind.TransactOpts, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _GlpManager.contract.Transact(opts, "addLiquidity", _token, _amount, _minUsdg, _minGlp)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ece366a.
//
// Solidity: function addLiquidity(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_GlpManager *GlpManagerSession) AddLiquidity(_token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _GlpManager.Contract.AddLiquidity(&_GlpManager.TransactOpts, _token, _amount, _minUsdg, _minGlp)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ece366a.
//
// Solidity: function addLiquidity(address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_GlpManager *GlpManagerTransactorSession) AddLiquidity(_token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _GlpManager.Contract.AddLiquidity(&_GlpManager.TransactOpts, _token, _amount, _minUsdg, _minGlp)
}

// AddLiquidityForAccount is a paid mutator transaction binding the contract method 0x17eb2a15.
//
// Solidity: function addLiquidityForAccount(address _fundingAccount, address _account, address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_GlpManager *GlpManagerTransactor) AddLiquidityForAccount(opts *bind.TransactOpts, _fundingAccount common.Address, _account common.Address, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _GlpManager.contract.Transact(opts, "addLiquidityForAccount", _fundingAccount, _account, _token, _amount, _minUsdg, _minGlp)
}

// AddLiquidityForAccount is a paid mutator transaction binding the contract method 0x17eb2a15.
//
// Solidity: function addLiquidityForAccount(address _fundingAccount, address _account, address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_GlpManager *GlpManagerSession) AddLiquidityForAccount(_fundingAccount common.Address, _account common.Address, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _GlpManager.Contract.AddLiquidityForAccount(&_GlpManager.TransactOpts, _fundingAccount, _account, _token, _amount, _minUsdg, _minGlp)
}

// AddLiquidityForAccount is a paid mutator transaction binding the contract method 0x17eb2a15.
//
// Solidity: function addLiquidityForAccount(address _fundingAccount, address _account, address _token, uint256 _amount, uint256 _minUsdg, uint256 _minGlp) returns(uint256)
func (_GlpManager *GlpManagerTransactorSession) AddLiquidityForAccount(_fundingAccount common.Address, _account common.Address, _token common.Address, _amount *big.Int, _minUsdg *big.Int, _minGlp *big.Int) (*types.Transaction, error) {
	return _GlpManager.Contract.AddLiquidityForAccount(&_GlpManager.TransactOpts, _fundingAccount, _account, _token, _amount, _minUsdg, _minGlp)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x8fed0b2c.
//
// Solidity: function removeLiquidity(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_GlpManager *GlpManagerTransactor) RemoveLiquidity(opts *bind.TransactOpts, _tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _GlpManager.contract.Transact(opts, "removeLiquidity", _tokenOut, _glpAmount, _minOut, _receiver)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x8fed0b2c.
//
// Solidity: function removeLiquidity(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_GlpManager *GlpManagerSession) RemoveLiquidity(_tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _GlpManager.Contract.RemoveLiquidity(&_GlpManager.TransactOpts, _tokenOut, _glpAmount, _minOut, _receiver)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x8fed0b2c.
//
// Solidity: function removeLiquidity(address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_GlpManager *GlpManagerTransactorSession) RemoveLiquidity(_tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _GlpManager.Contract.RemoveLiquidity(&_GlpManager.TransactOpts, _tokenOut, _glpAmount, _minOut, _receiver)
}

// RemoveLiquidityForAccount is a paid mutator transaction binding the contract method 0x71d597ad.
//
// Solidity: function removeLiquidityForAccount(address _account, address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_GlpManager *GlpManagerTransactor) RemoveLiquidityForAccount(opts *bind.TransactOpts, _account common.Address, _tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _GlpManager.contract.Transact(opts, "removeLiquidityForAccount", _account, _tokenOut, _glpAmount, _minOut, _receiver)
}

// RemoveLiquidityForAccount is a paid mutator transaction binding the contract method 0x71d597ad.
//
// Solidity: function removeLiquidityForAccount(address _account, address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_GlpManager *GlpManagerSession) RemoveLiquidityForAccount(_account common.Address, _tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _GlpManager.Contract.RemoveLiquidityForAccount(&_GlpManager.TransactOpts, _account, _tokenOut, _glpAmount, _minOut, _receiver)
}

// RemoveLiquidityForAccount is a paid mutator transaction binding the contract method 0x71d597ad.
//
// Solidity: function removeLiquidityForAccount(address _account, address _tokenOut, uint256 _glpAmount, uint256 _minOut, address _receiver) returns(uint256)
func (_GlpManager *GlpManagerTransactorSession) RemoveLiquidityForAccount(_account common.Address, _tokenOut common.Address, _glpAmount *big.Int, _minOut *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _GlpManager.Contract.RemoveLiquidityForAccount(&_GlpManager.TransactOpts, _account, _tokenOut, _glpAmount, _minOut, _receiver)
}

// SetAumAdjustment is a paid mutator transaction binding the contract method 0x9116c4ae.
//
// Solidity: function setAumAdjustment(uint256 _aumAddition, uint256 _aumDeduction) returns()
func (_GlpManager *GlpManagerTransactor) SetAumAdjustment(opts *bind.TransactOpts, _aumAddition *big.Int, _aumDeduction *big.Int) (*types.Transaction, error) {
	return _GlpManager.contract.Transact(opts, "setAumAdjustment", _aumAddition, _aumDeduction)
}

// SetAumAdjustment is a paid mutator transaction binding the contract method 0x9116c4ae.
//
// Solidity: function setAumAdjustment(uint256 _aumAddition, uint256 _aumDeduction) returns()
func (_GlpManager *GlpManagerSession) SetAumAdjustment(_aumAddition *big.Int, _aumDeduction *big.Int) (*types.Transaction, error) {
	return _GlpManager.Contract.SetAumAdjustment(&_GlpManager.TransactOpts, _aumAddition, _aumDeduction)
}

// SetAumAdjustment is a paid mutator transaction binding the contract method 0x9116c4ae.
//
// Solidity: function setAumAdjustment(uint256 _aumAddition, uint256 _aumDeduction) returns()
func (_GlpManager *GlpManagerTransactorSession) SetAumAdjustment(_aumAddition *big.Int, _aumDeduction *big.Int) (*types.Transaction, error) {
	return _GlpManager.Contract.SetAumAdjustment(&_GlpManager.TransactOpts, _aumAddition, _aumDeduction)
}

// SetCooldownDuration is a paid mutator transaction binding the contract method 0x966be075.
//
// Solidity: function setCooldownDuration(uint256 _cooldownDuration) returns()
func (_GlpManager *GlpManagerTransactor) SetCooldownDuration(opts *bind.TransactOpts, _cooldownDuration *big.Int) (*types.Transaction, error) {
	return _GlpManager.contract.Transact(opts, "setCooldownDuration", _cooldownDuration)
}

// SetCooldownDuration is a paid mutator transaction binding the contract method 0x966be075.
//
// Solidity: function setCooldownDuration(uint256 _cooldownDuration) returns()
func (_GlpManager *GlpManagerSession) SetCooldownDuration(_cooldownDuration *big.Int) (*types.Transaction, error) {
	return _GlpManager.Contract.SetCooldownDuration(&_GlpManager.TransactOpts, _cooldownDuration)
}

// SetCooldownDuration is a paid mutator transaction binding the contract method 0x966be075.
//
// Solidity: function setCooldownDuration(uint256 _cooldownDuration) returns()
func (_GlpManager *GlpManagerTransactorSession) SetCooldownDuration(_cooldownDuration *big.Int) (*types.Transaction, error) {
	return _GlpManager.Contract.SetCooldownDuration(&_GlpManager.TransactOpts, _cooldownDuration)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_GlpManager *GlpManagerTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _GlpManager.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_GlpManager *GlpManagerSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _GlpManager.Contract.SetGov(&_GlpManager.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_GlpManager *GlpManagerTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _GlpManager.Contract.SetGov(&_GlpManager.TransactOpts, _gov)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_GlpManager *GlpManagerTransactor) SetHandler(opts *bind.TransactOpts, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _GlpManager.contract.Transact(opts, "setHandler", _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_GlpManager *GlpManagerSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _GlpManager.Contract.SetHandler(&_GlpManager.TransactOpts, _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_GlpManager *GlpManagerTransactorSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _GlpManager.Contract.SetHandler(&_GlpManager.TransactOpts, _handler, _isActive)
}

// SetInPrivateMode is a paid mutator transaction binding the contract method 0x6a86da19.
//
// Solidity: function setInPrivateMode(bool _inPrivateMode) returns()
func (_GlpManager *GlpManagerTransactor) SetInPrivateMode(opts *bind.TransactOpts, _inPrivateMode bool) (*types.Transaction, error) {
	return _GlpManager.contract.Transact(opts, "setInPrivateMode", _inPrivateMode)
}

// SetInPrivateMode is a paid mutator transaction binding the contract method 0x6a86da19.
//
// Solidity: function setInPrivateMode(bool _inPrivateMode) returns()
func (_GlpManager *GlpManagerSession) SetInPrivateMode(_inPrivateMode bool) (*types.Transaction, error) {
	return _GlpManager.Contract.SetInPrivateMode(&_GlpManager.TransactOpts, _inPrivateMode)
}

// SetInPrivateMode is a paid mutator transaction binding the contract method 0x6a86da19.
//
// Solidity: function setInPrivateMode(bool _inPrivateMode) returns()
func (_GlpManager *GlpManagerTransactorSession) SetInPrivateMode(_inPrivateMode bool) (*types.Transaction, error) {
	return _GlpManager.Contract.SetInPrivateMode(&_GlpManager.TransactOpts, _inPrivateMode)
}

// SetShortsTracker is a paid mutator transaction binding the contract method 0xd34ee093.
//
// Solidity: function setShortsTracker(address _shortsTracker) returns()
func (_GlpManager *GlpManagerTransactor) SetShortsTracker(opts *bind.TransactOpts, _shortsTracker common.Address) (*types.Transaction, error) {
	return _GlpManager.contract.Transact(opts, "setShortsTracker", _shortsTracker)
}

// SetShortsTracker is a paid mutator transaction binding the contract method 0xd34ee093.
//
// Solidity: function setShortsTracker(address _shortsTracker) returns()
func (_GlpManager *GlpManagerSession) SetShortsTracker(_shortsTracker common.Address) (*types.Transaction, error) {
	return _GlpManager.Contract.SetShortsTracker(&_GlpManager.TransactOpts, _shortsTracker)
}

// SetShortsTracker is a paid mutator transaction binding the contract method 0xd34ee093.
//
// Solidity: function setShortsTracker(address _shortsTracker) returns()
func (_GlpManager *GlpManagerTransactorSession) SetShortsTracker(_shortsTracker common.Address) (*types.Transaction, error) {
	return _GlpManager.Contract.SetShortsTracker(&_GlpManager.TransactOpts, _shortsTracker)
}

// SetShortsTrackerAveragePriceWeight is a paid mutator transaction binding the contract method 0x4f5f6b5e.
//
// Solidity: function setShortsTrackerAveragePriceWeight(uint256 _shortsTrackerAveragePriceWeight) returns()
func (_GlpManager *GlpManagerTransactor) SetShortsTrackerAveragePriceWeight(opts *bind.TransactOpts, _shortsTrackerAveragePriceWeight *big.Int) (*types.Transaction, error) {
	return _GlpManager.contract.Transact(opts, "setShortsTrackerAveragePriceWeight", _shortsTrackerAveragePriceWeight)
}

// SetShortsTrackerAveragePriceWeight is a paid mutator transaction binding the contract method 0x4f5f6b5e.
//
// Solidity: function setShortsTrackerAveragePriceWeight(uint256 _shortsTrackerAveragePriceWeight) returns()
func (_GlpManager *GlpManagerSession) SetShortsTrackerAveragePriceWeight(_shortsTrackerAveragePriceWeight *big.Int) (*types.Transaction, error) {
	return _GlpManager.Contract.SetShortsTrackerAveragePriceWeight(&_GlpManager.TransactOpts, _shortsTrackerAveragePriceWeight)
}

// SetShortsTrackerAveragePriceWeight is a paid mutator transaction binding the contract method 0x4f5f6b5e.
//
// Solidity: function setShortsTrackerAveragePriceWeight(uint256 _shortsTrackerAveragePriceWeight) returns()
func (_GlpManager *GlpManagerTransactorSession) SetShortsTrackerAveragePriceWeight(_shortsTrackerAveragePriceWeight *big.Int) (*types.Transaction, error) {
	return _GlpManager.Contract.SetShortsTrackerAveragePriceWeight(&_GlpManager.TransactOpts, _shortsTrackerAveragePriceWeight)
}

// GlpManagerAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the GlpManager contract.
type GlpManagerAddLiquidityIterator struct {
	Event *GlpManagerAddLiquidity // Event containing the contract specifics and raw log

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
func (it *GlpManagerAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlpManagerAddLiquidity)
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
		it.Event = new(GlpManagerAddLiquidity)
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
func (it *GlpManagerAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlpManagerAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlpManagerAddLiquidity represents a AddLiquidity event raised by the GlpManager contract.
type GlpManagerAddLiquidity struct {
	Account    common.Address
	Token      common.Address
	Amount     *big.Int
	AumInUsdg  *big.Int
	GlpSupply  *big.Int
	UsdgAmount *big.Int
	MintAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x38dc38b96482be64113daffd8d464ebda93e856b70ccfc605e69ccf892ab981e.
//
// Solidity: event AddLiquidity(address account, address token, uint256 amount, uint256 aumInUsdg, uint256 glpSupply, uint256 usdgAmount, uint256 mintAmount)
func (_GlpManager *GlpManagerFilterer) FilterAddLiquidity(opts *bind.FilterOpts) (*GlpManagerAddLiquidityIterator, error) {

	logs, sub, err := _GlpManager.contract.FilterLogs(opts, "AddLiquidity")
	if err != nil {
		return nil, err
	}
	return &GlpManagerAddLiquidityIterator{contract: _GlpManager.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x38dc38b96482be64113daffd8d464ebda93e856b70ccfc605e69ccf892ab981e.
//
// Solidity: event AddLiquidity(address account, address token, uint256 amount, uint256 aumInUsdg, uint256 glpSupply, uint256 usdgAmount, uint256 mintAmount)
func (_GlpManager *GlpManagerFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *GlpManagerAddLiquidity) (event.Subscription, error) {

	logs, sub, err := _GlpManager.contract.WatchLogs(opts, "AddLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlpManagerAddLiquidity)
				if err := _GlpManager.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x38dc38b96482be64113daffd8d464ebda93e856b70ccfc605e69ccf892ab981e.
//
// Solidity: event AddLiquidity(address account, address token, uint256 amount, uint256 aumInUsdg, uint256 glpSupply, uint256 usdgAmount, uint256 mintAmount)
func (_GlpManager *GlpManagerFilterer) ParseAddLiquidity(log types.Log) (*GlpManagerAddLiquidity, error) {
	event := new(GlpManagerAddLiquidity)
	if err := _GlpManager.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlpManagerRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the GlpManager contract.
type GlpManagerRemoveLiquidityIterator struct {
	Event *GlpManagerRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *GlpManagerRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlpManagerRemoveLiquidity)
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
		it.Event = new(GlpManagerRemoveLiquidity)
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
func (it *GlpManagerRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlpManagerRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlpManagerRemoveLiquidity represents a RemoveLiquidity event raised by the GlpManager contract.
type GlpManagerRemoveLiquidity struct {
	Account    common.Address
	Token      common.Address
	GlpAmount  *big.Int
	AumInUsdg  *big.Int
	GlpSupply  *big.Int
	UsdgAmount *big.Int
	AmountOut  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x87b9679bb9a4944bafa98c267e7cd4a00ab29fed48afdefae25f0fca5da27940.
//
// Solidity: event RemoveLiquidity(address account, address token, uint256 glpAmount, uint256 aumInUsdg, uint256 glpSupply, uint256 usdgAmount, uint256 amountOut)
func (_GlpManager *GlpManagerFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts) (*GlpManagerRemoveLiquidityIterator, error) {

	logs, sub, err := _GlpManager.contract.FilterLogs(opts, "RemoveLiquidity")
	if err != nil {
		return nil, err
	}
	return &GlpManagerRemoveLiquidityIterator{contract: _GlpManager.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x87b9679bb9a4944bafa98c267e7cd4a00ab29fed48afdefae25f0fca5da27940.
//
// Solidity: event RemoveLiquidity(address account, address token, uint256 glpAmount, uint256 aumInUsdg, uint256 glpSupply, uint256 usdgAmount, uint256 amountOut)
func (_GlpManager *GlpManagerFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *GlpManagerRemoveLiquidity) (event.Subscription, error) {

	logs, sub, err := _GlpManager.contract.WatchLogs(opts, "RemoveLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlpManagerRemoveLiquidity)
				if err := _GlpManager.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x87b9679bb9a4944bafa98c267e7cd4a00ab29fed48afdefae25f0fca5da27940.
//
// Solidity: event RemoveLiquidity(address account, address token, uint256 glpAmount, uint256 aumInUsdg, uint256 glpSupply, uint256 usdgAmount, uint256 amountOut)
func (_GlpManager *GlpManagerFilterer) ParseRemoveLiquidity(log types.Log) (*GlpManagerRemoveLiquidity, error) {
	event := new(GlpManagerRemoveLiquidity)
	if err := _GlpManager.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

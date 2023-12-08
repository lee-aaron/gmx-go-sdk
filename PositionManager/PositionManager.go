// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PositionManager

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

// PositionManagerMetaData contains all meta data concerning the PositionManager contract.
var PositionManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_shortsTracker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_depositFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_orderBook\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"BASIS_POINTS_DIVISOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decreasePosition\",\"inputs\":[{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decreasePositionAndSwap\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decreasePositionAndSwapETH\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decreasePositionETH\",\"inputs\":[{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ethTransferGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeDecreaseOrder\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_orderIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_feeReceiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeIncreaseOrder\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_orderIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_feeReceiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSwapOrder\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_orderIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_feeReceiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeReserves\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inLegacyMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increasePosition\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increasePositionBufferBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increasePositionETH\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"isLiquidator\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isOrderKeeper\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPartner\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"liquidatePosition\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxGlobalLongSizes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxGlobalShortSizes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"orderBook\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"referralStorage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sendValue\",\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAdmin\",\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDepositFee\",\"inputs\":[{\"name\":\"_depositFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEthTransferGasLimit\",\"inputs\":[{\"name\":\"_ethTransferGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInLegacyMode\",\"inputs\":[{\"name\":\"_inLegacyMode\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIncreasePositionBufferBps\",\"inputs\":[{\"name\":\"_increasePositionBufferBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLiquidator\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxGlobalSizes\",\"inputs\":[{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_longSizes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_shortSizes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOrderKeeper\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPartner\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setReferralStorage\",\"inputs\":[{\"name\":\"_referralStorage\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setShouldValidateIncreaseOrder\",\"inputs\":[{\"name\":\"_shouldValidateIncreaseOrder\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shortsTracker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"shouldValidateIncreaseOrder\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vault\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"weth\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DecreasePositionReferral\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"marginFeeBasisPoints\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"referralCode\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"referrer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncreasePositionReferral\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"marginFeeBasisPoints\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"referralCode\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"referrer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetDepositFee\",\"inputs\":[{\"name\":\"depositFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetEthTransferGasLimit\",\"inputs\":[{\"name\":\"ethTransferGasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetInLegacyMode\",\"inputs\":[{\"name\":\"inLegacyMode\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetIncreasePositionBufferBps\",\"inputs\":[{\"name\":\"increasePositionBufferBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetLiquidator\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"isActive\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetMaxGlobalSizes\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"longSizes\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"shortSizes\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetOrderKeeper\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"isActive\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetPartner\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"isActive\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetReferralStorage\",\"inputs\":[{\"name\":\"referralStorage\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetShouldValidateIncreaseOrder\",\"inputs\":[{\"name\":\"shouldValidateIncreaseOrder\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawFees\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// PositionManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use PositionManagerMetaData.ABI instead.
var PositionManagerABI = PositionManagerMetaData.ABI

// PositionManager is an auto generated Go binding around an Ethereum contract.
type PositionManager struct {
	PositionManagerCaller     // Read-only binding to the contract
	PositionManagerTransactor // Write-only binding to the contract
	PositionManagerFilterer   // Log filterer for contract events
}

// PositionManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PositionManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PositionManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PositionManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PositionManagerSession struct {
	Contract     *PositionManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PositionManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PositionManagerCallerSession struct {
	Contract *PositionManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// PositionManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PositionManagerTransactorSession struct {
	Contract     *PositionManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// PositionManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PositionManagerRaw struct {
	Contract *PositionManager // Generic contract binding to access the raw methods on
}

// PositionManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PositionManagerCallerRaw struct {
	Contract *PositionManagerCaller // Generic read-only contract binding to access the raw methods on
}

// PositionManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PositionManagerTransactorRaw struct {
	Contract *PositionManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPositionManager creates a new instance of PositionManager, bound to a specific deployed contract.
func NewPositionManager(address common.Address, backend bind.ContractBackend) (*PositionManager, error) {
	contract, err := bindPositionManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PositionManager{PositionManagerCaller: PositionManagerCaller{contract: contract}, PositionManagerTransactor: PositionManagerTransactor{contract: contract}, PositionManagerFilterer: PositionManagerFilterer{contract: contract}}, nil
}

// NewPositionManagerCaller creates a new read-only instance of PositionManager, bound to a specific deployed contract.
func NewPositionManagerCaller(address common.Address, caller bind.ContractCaller) (*PositionManagerCaller, error) {
	contract, err := bindPositionManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PositionManagerCaller{contract: contract}, nil
}

// NewPositionManagerTransactor creates a new write-only instance of PositionManager, bound to a specific deployed contract.
func NewPositionManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PositionManagerTransactor, error) {
	contract, err := bindPositionManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PositionManagerTransactor{contract: contract}, nil
}

// NewPositionManagerFilterer creates a new log filterer instance of PositionManager, bound to a specific deployed contract.
func NewPositionManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PositionManagerFilterer, error) {
	contract, err := bindPositionManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PositionManagerFilterer{contract: contract}, nil
}

// bindPositionManager binds a generic wrapper to an already deployed contract.
func bindPositionManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PositionManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PositionManager *PositionManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PositionManager.Contract.PositionManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PositionManager *PositionManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionManager.Contract.PositionManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PositionManager *PositionManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PositionManager.Contract.PositionManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PositionManager *PositionManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PositionManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PositionManager *PositionManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PositionManager *PositionManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PositionManager.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_PositionManager *PositionManagerCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_PositionManager *PositionManagerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _PositionManager.Contract.BASISPOINTSDIVISOR(&_PositionManager.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_PositionManager *PositionManagerCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _PositionManager.Contract.BASISPOINTSDIVISOR(&_PositionManager.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PositionManager *PositionManagerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PositionManager *PositionManagerSession) Admin() (common.Address, error) {
	return _PositionManager.Contract.Admin(&_PositionManager.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PositionManager *PositionManagerCallerSession) Admin() (common.Address, error) {
	return _PositionManager.Contract.Admin(&_PositionManager.CallOpts)
}

// DepositFee is a free data retrieval call binding the contract method 0x67a52793.
//
// Solidity: function depositFee() view returns(uint256)
func (_PositionManager *PositionManagerCaller) DepositFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "depositFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositFee is a free data retrieval call binding the contract method 0x67a52793.
//
// Solidity: function depositFee() view returns(uint256)
func (_PositionManager *PositionManagerSession) DepositFee() (*big.Int, error) {
	return _PositionManager.Contract.DepositFee(&_PositionManager.CallOpts)
}

// DepositFee is a free data retrieval call binding the contract method 0x67a52793.
//
// Solidity: function depositFee() view returns(uint256)
func (_PositionManager *PositionManagerCallerSession) DepositFee() (*big.Int, error) {
	return _PositionManager.Contract.DepositFee(&_PositionManager.CallOpts)
}

// EthTransferGasLimit is a free data retrieval call binding the contract method 0x2d79cf42.
//
// Solidity: function ethTransferGasLimit() view returns(uint256)
func (_PositionManager *PositionManagerCaller) EthTransferGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "ethTransferGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthTransferGasLimit is a free data retrieval call binding the contract method 0x2d79cf42.
//
// Solidity: function ethTransferGasLimit() view returns(uint256)
func (_PositionManager *PositionManagerSession) EthTransferGasLimit() (*big.Int, error) {
	return _PositionManager.Contract.EthTransferGasLimit(&_PositionManager.CallOpts)
}

// EthTransferGasLimit is a free data retrieval call binding the contract method 0x2d79cf42.
//
// Solidity: function ethTransferGasLimit() view returns(uint256)
func (_PositionManager *PositionManagerCallerSession) EthTransferGasLimit() (*big.Int, error) {
	return _PositionManager.Contract.EthTransferGasLimit(&_PositionManager.CallOpts)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_PositionManager *PositionManagerCaller) FeeReserves(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "feeReserves", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_PositionManager *PositionManagerSession) FeeReserves(arg0 common.Address) (*big.Int, error) {
	return _PositionManager.Contract.FeeReserves(&_PositionManager.CallOpts, arg0)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_PositionManager *PositionManagerCallerSession) FeeReserves(arg0 common.Address) (*big.Int, error) {
	return _PositionManager.Contract.FeeReserves(&_PositionManager.CallOpts, arg0)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_PositionManager *PositionManagerCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_PositionManager *PositionManagerSession) Gov() (common.Address, error) {
	return _PositionManager.Contract.Gov(&_PositionManager.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_PositionManager *PositionManagerCallerSession) Gov() (common.Address, error) {
	return _PositionManager.Contract.Gov(&_PositionManager.CallOpts)
}

// InLegacyMode is a free data retrieval call binding the contract method 0xd4ca83f9.
//
// Solidity: function inLegacyMode() view returns(bool)
func (_PositionManager *PositionManagerCaller) InLegacyMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "inLegacyMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InLegacyMode is a free data retrieval call binding the contract method 0xd4ca83f9.
//
// Solidity: function inLegacyMode() view returns(bool)
func (_PositionManager *PositionManagerSession) InLegacyMode() (bool, error) {
	return _PositionManager.Contract.InLegacyMode(&_PositionManager.CallOpts)
}

// InLegacyMode is a free data retrieval call binding the contract method 0xd4ca83f9.
//
// Solidity: function inLegacyMode() view returns(bool)
func (_PositionManager *PositionManagerCallerSession) InLegacyMode() (bool, error) {
	return _PositionManager.Contract.InLegacyMode(&_PositionManager.CallOpts)
}

// IncreasePositionBufferBps is a free data retrieval call binding the contract method 0x98d1e03a.
//
// Solidity: function increasePositionBufferBps() view returns(uint256)
func (_PositionManager *PositionManagerCaller) IncreasePositionBufferBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "increasePositionBufferBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncreasePositionBufferBps is a free data retrieval call binding the contract method 0x98d1e03a.
//
// Solidity: function increasePositionBufferBps() view returns(uint256)
func (_PositionManager *PositionManagerSession) IncreasePositionBufferBps() (*big.Int, error) {
	return _PositionManager.Contract.IncreasePositionBufferBps(&_PositionManager.CallOpts)
}

// IncreasePositionBufferBps is a free data retrieval call binding the contract method 0x98d1e03a.
//
// Solidity: function increasePositionBufferBps() view returns(uint256)
func (_PositionManager *PositionManagerCallerSession) IncreasePositionBufferBps() (*big.Int, error) {
	return _PositionManager.Contract.IncreasePositionBufferBps(&_PositionManager.CallOpts)
}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address ) view returns(bool)
func (_PositionManager *PositionManagerCaller) IsLiquidator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "isLiquidator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address ) view returns(bool)
func (_PositionManager *PositionManagerSession) IsLiquidator(arg0 common.Address) (bool, error) {
	return _PositionManager.Contract.IsLiquidator(&_PositionManager.CallOpts, arg0)
}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address ) view returns(bool)
func (_PositionManager *PositionManagerCallerSession) IsLiquidator(arg0 common.Address) (bool, error) {
	return _PositionManager.Contract.IsLiquidator(&_PositionManager.CallOpts, arg0)
}

// IsOrderKeeper is a free data retrieval call binding the contract method 0x3833f5f5.
//
// Solidity: function isOrderKeeper(address ) view returns(bool)
func (_PositionManager *PositionManagerCaller) IsOrderKeeper(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "isOrderKeeper", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOrderKeeper is a free data retrieval call binding the contract method 0x3833f5f5.
//
// Solidity: function isOrderKeeper(address ) view returns(bool)
func (_PositionManager *PositionManagerSession) IsOrderKeeper(arg0 common.Address) (bool, error) {
	return _PositionManager.Contract.IsOrderKeeper(&_PositionManager.CallOpts, arg0)
}

// IsOrderKeeper is a free data retrieval call binding the contract method 0x3833f5f5.
//
// Solidity: function isOrderKeeper(address ) view returns(bool)
func (_PositionManager *PositionManagerCallerSession) IsOrderKeeper(arg0 common.Address) (bool, error) {
	return _PositionManager.Contract.IsOrderKeeper(&_PositionManager.CallOpts, arg0)
}

// IsPartner is a free data retrieval call binding the contract method 0x8c0f9aac.
//
// Solidity: function isPartner(address ) view returns(bool)
func (_PositionManager *PositionManagerCaller) IsPartner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "isPartner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPartner is a free data retrieval call binding the contract method 0x8c0f9aac.
//
// Solidity: function isPartner(address ) view returns(bool)
func (_PositionManager *PositionManagerSession) IsPartner(arg0 common.Address) (bool, error) {
	return _PositionManager.Contract.IsPartner(&_PositionManager.CallOpts, arg0)
}

// IsPartner is a free data retrieval call binding the contract method 0x8c0f9aac.
//
// Solidity: function isPartner(address ) view returns(bool)
func (_PositionManager *PositionManagerCallerSession) IsPartner(arg0 common.Address) (bool, error) {
	return _PositionManager.Contract.IsPartner(&_PositionManager.CallOpts, arg0)
}

// MaxGlobalLongSizes is a free data retrieval call binding the contract method 0x1045c74e.
//
// Solidity: function maxGlobalLongSizes(address ) view returns(uint256)
func (_PositionManager *PositionManagerCaller) MaxGlobalLongSizes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "maxGlobalLongSizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGlobalLongSizes is a free data retrieval call binding the contract method 0x1045c74e.
//
// Solidity: function maxGlobalLongSizes(address ) view returns(uint256)
func (_PositionManager *PositionManagerSession) MaxGlobalLongSizes(arg0 common.Address) (*big.Int, error) {
	return _PositionManager.Contract.MaxGlobalLongSizes(&_PositionManager.CallOpts, arg0)
}

// MaxGlobalLongSizes is a free data retrieval call binding the contract method 0x1045c74e.
//
// Solidity: function maxGlobalLongSizes(address ) view returns(uint256)
func (_PositionManager *PositionManagerCallerSession) MaxGlobalLongSizes(arg0 common.Address) (*big.Int, error) {
	return _PositionManager.Contract.MaxGlobalLongSizes(&_PositionManager.CallOpts, arg0)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_PositionManager *PositionManagerCaller) MaxGlobalShortSizes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "maxGlobalShortSizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_PositionManager *PositionManagerSession) MaxGlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _PositionManager.Contract.MaxGlobalShortSizes(&_PositionManager.CallOpts, arg0)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_PositionManager *PositionManagerCallerSession) MaxGlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _PositionManager.Contract.MaxGlobalShortSizes(&_PositionManager.CallOpts, arg0)
}

// OrderBook is a free data retrieval call binding the contract method 0x776af5ba.
//
// Solidity: function orderBook() view returns(address)
func (_PositionManager *PositionManagerCaller) OrderBook(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "orderBook")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OrderBook is a free data retrieval call binding the contract method 0x776af5ba.
//
// Solidity: function orderBook() view returns(address)
func (_PositionManager *PositionManagerSession) OrderBook() (common.Address, error) {
	return _PositionManager.Contract.OrderBook(&_PositionManager.CallOpts)
}

// OrderBook is a free data retrieval call binding the contract method 0x776af5ba.
//
// Solidity: function orderBook() view returns(address)
func (_PositionManager *PositionManagerCallerSession) OrderBook() (common.Address, error) {
	return _PositionManager.Contract.OrderBook(&_PositionManager.CallOpts)
}

// ReferralStorage is a free data retrieval call binding the contract method 0x006cc35e.
//
// Solidity: function referralStorage() view returns(address)
func (_PositionManager *PositionManagerCaller) ReferralStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "referralStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReferralStorage is a free data retrieval call binding the contract method 0x006cc35e.
//
// Solidity: function referralStorage() view returns(address)
func (_PositionManager *PositionManagerSession) ReferralStorage() (common.Address, error) {
	return _PositionManager.Contract.ReferralStorage(&_PositionManager.CallOpts)
}

// ReferralStorage is a free data retrieval call binding the contract method 0x006cc35e.
//
// Solidity: function referralStorage() view returns(address)
func (_PositionManager *PositionManagerCallerSession) ReferralStorage() (common.Address, error) {
	return _PositionManager.Contract.ReferralStorage(&_PositionManager.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_PositionManager *PositionManagerCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_PositionManager *PositionManagerSession) Router() (common.Address, error) {
	return _PositionManager.Contract.Router(&_PositionManager.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_PositionManager *PositionManagerCallerSession) Router() (common.Address, error) {
	return _PositionManager.Contract.Router(&_PositionManager.CallOpts)
}

// ShortsTracker is a free data retrieval call binding the contract method 0x657bc5d0.
//
// Solidity: function shortsTracker() view returns(address)
func (_PositionManager *PositionManagerCaller) ShortsTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "shortsTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ShortsTracker is a free data retrieval call binding the contract method 0x657bc5d0.
//
// Solidity: function shortsTracker() view returns(address)
func (_PositionManager *PositionManagerSession) ShortsTracker() (common.Address, error) {
	return _PositionManager.Contract.ShortsTracker(&_PositionManager.CallOpts)
}

// ShortsTracker is a free data retrieval call binding the contract method 0x657bc5d0.
//
// Solidity: function shortsTracker() view returns(address)
func (_PositionManager *PositionManagerCallerSession) ShortsTracker() (common.Address, error) {
	return _PositionManager.Contract.ShortsTracker(&_PositionManager.CallOpts)
}

// ShouldValidateIncreaseOrder is a free data retrieval call binding the contract method 0x4584bd4b.
//
// Solidity: function shouldValidateIncreaseOrder() view returns(bool)
func (_PositionManager *PositionManagerCaller) ShouldValidateIncreaseOrder(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "shouldValidateIncreaseOrder")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ShouldValidateIncreaseOrder is a free data retrieval call binding the contract method 0x4584bd4b.
//
// Solidity: function shouldValidateIncreaseOrder() view returns(bool)
func (_PositionManager *PositionManagerSession) ShouldValidateIncreaseOrder() (bool, error) {
	return _PositionManager.Contract.ShouldValidateIncreaseOrder(&_PositionManager.CallOpts)
}

// ShouldValidateIncreaseOrder is a free data retrieval call binding the contract method 0x4584bd4b.
//
// Solidity: function shouldValidateIncreaseOrder() view returns(bool)
func (_PositionManager *PositionManagerCallerSession) ShouldValidateIncreaseOrder() (bool, error) {
	return _PositionManager.Contract.ShouldValidateIncreaseOrder(&_PositionManager.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_PositionManager *PositionManagerCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_PositionManager *PositionManagerSession) Vault() (common.Address, error) {
	return _PositionManager.Contract.Vault(&_PositionManager.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_PositionManager *PositionManagerCallerSession) Vault() (common.Address, error) {
	return _PositionManager.Contract.Vault(&_PositionManager.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PositionManager *PositionManagerCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionManager.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PositionManager *PositionManagerSession) Weth() (common.Address, error) {
	return _PositionManager.Contract.Weth(&_PositionManager.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PositionManager *PositionManagerCallerSession) Weth() (common.Address, error) {
	return _PositionManager.Contract.Weth(&_PositionManager.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount) returns()
func (_PositionManager *PositionManagerTransactor) Approve(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "approve", _token, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount) returns()
func (_PositionManager *PositionManagerSession) Approve(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.Approve(&_PositionManager.TransactOpts, _token, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount) returns()
func (_PositionManager *PositionManagerTransactorSession) Approve(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.Approve(&_PositionManager.TransactOpts, _token, _spender, _amount)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x90205d8c.
//
// Solidity: function decreasePosition(address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price) returns()
func (_PositionManager *PositionManagerTransactor) DecreasePosition(opts *bind.TransactOpts, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "decreasePosition", _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x90205d8c.
//
// Solidity: function decreasePosition(address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price) returns()
func (_PositionManager *PositionManagerSession) DecreasePosition(_collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.DecreasePosition(&_PositionManager.TransactOpts, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x90205d8c.
//
// Solidity: function decreasePosition(address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price) returns()
func (_PositionManager *PositionManagerTransactorSession) DecreasePosition(_collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.DecreasePosition(&_PositionManager.TransactOpts, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price)
}

// DecreasePositionAndSwap is a paid mutator transaction binding the contract method 0x5fc8500e.
//
// Solidity: function decreasePositionAndSwap(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price, uint256 _minOut) returns()
func (_PositionManager *PositionManagerTransactor) DecreasePositionAndSwap(opts *bind.TransactOpts, _path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int, _minOut *big.Int) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "decreasePositionAndSwap", _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price, _minOut)
}

// DecreasePositionAndSwap is a paid mutator transaction binding the contract method 0x5fc8500e.
//
// Solidity: function decreasePositionAndSwap(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price, uint256 _minOut) returns()
func (_PositionManager *PositionManagerSession) DecreasePositionAndSwap(_path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int, _minOut *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.DecreasePositionAndSwap(&_PositionManager.TransactOpts, _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price, _minOut)
}

// DecreasePositionAndSwap is a paid mutator transaction binding the contract method 0x5fc8500e.
//
// Solidity: function decreasePositionAndSwap(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price, uint256 _minOut) returns()
func (_PositionManager *PositionManagerTransactorSession) DecreasePositionAndSwap(_path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int, _minOut *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.DecreasePositionAndSwap(&_PositionManager.TransactOpts, _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price, _minOut)
}

// DecreasePositionAndSwapETH is a paid mutator transaction binding the contract method 0x3039e37f.
//
// Solidity: function decreasePositionAndSwapETH(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price, uint256 _minOut) returns()
func (_PositionManager *PositionManagerTransactor) DecreasePositionAndSwapETH(opts *bind.TransactOpts, _path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int, _minOut *big.Int) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "decreasePositionAndSwapETH", _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price, _minOut)
}

// DecreasePositionAndSwapETH is a paid mutator transaction binding the contract method 0x3039e37f.
//
// Solidity: function decreasePositionAndSwapETH(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price, uint256 _minOut) returns()
func (_PositionManager *PositionManagerSession) DecreasePositionAndSwapETH(_path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int, _minOut *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.DecreasePositionAndSwapETH(&_PositionManager.TransactOpts, _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price, _minOut)
}

// DecreasePositionAndSwapETH is a paid mutator transaction binding the contract method 0x3039e37f.
//
// Solidity: function decreasePositionAndSwapETH(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price, uint256 _minOut) returns()
func (_PositionManager *PositionManagerTransactorSession) DecreasePositionAndSwapETH(_path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int, _minOut *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.DecreasePositionAndSwapETH(&_PositionManager.TransactOpts, _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price, _minOut)
}

// DecreasePositionETH is a paid mutator transaction binding the contract method 0x430ed37c.
//
// Solidity: function decreasePositionETH(address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price) returns()
func (_PositionManager *PositionManagerTransactor) DecreasePositionETH(opts *bind.TransactOpts, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "decreasePositionETH", _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price)
}

// DecreasePositionETH is a paid mutator transaction binding the contract method 0x430ed37c.
//
// Solidity: function decreasePositionETH(address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price) returns()
func (_PositionManager *PositionManagerSession) DecreasePositionETH(_collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.DecreasePositionETH(&_PositionManager.TransactOpts, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price)
}

// DecreasePositionETH is a paid mutator transaction binding the contract method 0x430ed37c.
//
// Solidity: function decreasePositionETH(address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _price) returns()
func (_PositionManager *PositionManagerTransactorSession) DecreasePositionETH(_collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _price *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.DecreasePositionETH(&_PositionManager.TransactOpts, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _price)
}

// ExecuteDecreaseOrder is a paid mutator transaction binding the contract method 0x11d9444a.
//
// Solidity: function executeDecreaseOrder(address _account, uint256 _orderIndex, address _feeReceiver) returns()
func (_PositionManager *PositionManagerTransactor) ExecuteDecreaseOrder(opts *bind.TransactOpts, _account common.Address, _orderIndex *big.Int, _feeReceiver common.Address) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "executeDecreaseOrder", _account, _orderIndex, _feeReceiver)
}

// ExecuteDecreaseOrder is a paid mutator transaction binding the contract method 0x11d9444a.
//
// Solidity: function executeDecreaseOrder(address _account, uint256 _orderIndex, address _feeReceiver) returns()
func (_PositionManager *PositionManagerSession) ExecuteDecreaseOrder(_account common.Address, _orderIndex *big.Int, _feeReceiver common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.ExecuteDecreaseOrder(&_PositionManager.TransactOpts, _account, _orderIndex, _feeReceiver)
}

// ExecuteDecreaseOrder is a paid mutator transaction binding the contract method 0x11d9444a.
//
// Solidity: function executeDecreaseOrder(address _account, uint256 _orderIndex, address _feeReceiver) returns()
func (_PositionManager *PositionManagerTransactorSession) ExecuteDecreaseOrder(_account common.Address, _orderIndex *big.Int, _feeReceiver common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.ExecuteDecreaseOrder(&_PositionManager.TransactOpts, _account, _orderIndex, _feeReceiver)
}

// ExecuteIncreaseOrder is a paid mutator transaction binding the contract method 0xd38ab519.
//
// Solidity: function executeIncreaseOrder(address _account, uint256 _orderIndex, address _feeReceiver) returns()
func (_PositionManager *PositionManagerTransactor) ExecuteIncreaseOrder(opts *bind.TransactOpts, _account common.Address, _orderIndex *big.Int, _feeReceiver common.Address) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "executeIncreaseOrder", _account, _orderIndex, _feeReceiver)
}

// ExecuteIncreaseOrder is a paid mutator transaction binding the contract method 0xd38ab519.
//
// Solidity: function executeIncreaseOrder(address _account, uint256 _orderIndex, address _feeReceiver) returns()
func (_PositionManager *PositionManagerSession) ExecuteIncreaseOrder(_account common.Address, _orderIndex *big.Int, _feeReceiver common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.ExecuteIncreaseOrder(&_PositionManager.TransactOpts, _account, _orderIndex, _feeReceiver)
}

// ExecuteIncreaseOrder is a paid mutator transaction binding the contract method 0xd38ab519.
//
// Solidity: function executeIncreaseOrder(address _account, uint256 _orderIndex, address _feeReceiver) returns()
func (_PositionManager *PositionManagerTransactorSession) ExecuteIncreaseOrder(_account common.Address, _orderIndex *big.Int, _feeReceiver common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.ExecuteIncreaseOrder(&_PositionManager.TransactOpts, _account, _orderIndex, _feeReceiver)
}

// ExecuteSwapOrder is a paid mutator transaction binding the contract method 0x07c7edc3.
//
// Solidity: function executeSwapOrder(address _account, uint256 _orderIndex, address _feeReceiver) returns()
func (_PositionManager *PositionManagerTransactor) ExecuteSwapOrder(opts *bind.TransactOpts, _account common.Address, _orderIndex *big.Int, _feeReceiver common.Address) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "executeSwapOrder", _account, _orderIndex, _feeReceiver)
}

// ExecuteSwapOrder is a paid mutator transaction binding the contract method 0x07c7edc3.
//
// Solidity: function executeSwapOrder(address _account, uint256 _orderIndex, address _feeReceiver) returns()
func (_PositionManager *PositionManagerSession) ExecuteSwapOrder(_account common.Address, _orderIndex *big.Int, _feeReceiver common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.ExecuteSwapOrder(&_PositionManager.TransactOpts, _account, _orderIndex, _feeReceiver)
}

// ExecuteSwapOrder is a paid mutator transaction binding the contract method 0x07c7edc3.
//
// Solidity: function executeSwapOrder(address _account, uint256 _orderIndex, address _feeReceiver) returns()
func (_PositionManager *PositionManagerTransactorSession) ExecuteSwapOrder(_account common.Address, _orderIndex *big.Int, _feeReceiver common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.ExecuteSwapOrder(&_PositionManager.TransactOpts, _account, _orderIndex, _feeReceiver)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0xb7ddc992.
//
// Solidity: function increasePosition(address[] _path, address _indexToken, uint256 _amountIn, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _price) returns()
func (_PositionManager *PositionManagerTransactor) IncreasePosition(opts *bind.TransactOpts, _path []common.Address, _indexToken common.Address, _amountIn *big.Int, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _price *big.Int) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "increasePosition", _path, _indexToken, _amountIn, _minOut, _sizeDelta, _isLong, _price)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0xb7ddc992.
//
// Solidity: function increasePosition(address[] _path, address _indexToken, uint256 _amountIn, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _price) returns()
func (_PositionManager *PositionManagerSession) IncreasePosition(_path []common.Address, _indexToken common.Address, _amountIn *big.Int, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _price *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.IncreasePosition(&_PositionManager.TransactOpts, _path, _indexToken, _amountIn, _minOut, _sizeDelta, _isLong, _price)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0xb7ddc992.
//
// Solidity: function increasePosition(address[] _path, address _indexToken, uint256 _amountIn, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _price) returns()
func (_PositionManager *PositionManagerTransactorSession) IncreasePosition(_path []common.Address, _indexToken common.Address, _amountIn *big.Int, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _price *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.IncreasePosition(&_PositionManager.TransactOpts, _path, _indexToken, _amountIn, _minOut, _sizeDelta, _isLong, _price)
}

// IncreasePositionETH is a paid mutator transaction binding the contract method 0xb32755de.
//
// Solidity: function increasePositionETH(address[] _path, address _indexToken, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _price) payable returns()
func (_PositionManager *PositionManagerTransactor) IncreasePositionETH(opts *bind.TransactOpts, _path []common.Address, _indexToken common.Address, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _price *big.Int) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "increasePositionETH", _path, _indexToken, _minOut, _sizeDelta, _isLong, _price)
}

// IncreasePositionETH is a paid mutator transaction binding the contract method 0xb32755de.
//
// Solidity: function increasePositionETH(address[] _path, address _indexToken, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _price) payable returns()
func (_PositionManager *PositionManagerSession) IncreasePositionETH(_path []common.Address, _indexToken common.Address, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _price *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.IncreasePositionETH(&_PositionManager.TransactOpts, _path, _indexToken, _minOut, _sizeDelta, _isLong, _price)
}

// IncreasePositionETH is a paid mutator transaction binding the contract method 0xb32755de.
//
// Solidity: function increasePositionETH(address[] _path, address _indexToken, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _price) payable returns()
func (_PositionManager *PositionManagerTransactorSession) IncreasePositionETH(_path []common.Address, _indexToken common.Address, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _price *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.IncreasePositionETH(&_PositionManager.TransactOpts, _path, _indexToken, _minOut, _sizeDelta, _isLong, _price)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_PositionManager *PositionManagerTransactor) LiquidatePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "liquidatePosition", _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_PositionManager *PositionManagerSession) LiquidatePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.LiquidatePosition(&_PositionManager.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_PositionManager *PositionManagerTransactorSession) LiquidatePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.LiquidatePosition(&_PositionManager.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// SendValue is a paid mutator transaction binding the contract method 0x24a084df.
//
// Solidity: function sendValue(address _receiver, uint256 _amount) returns()
func (_PositionManager *PositionManagerTransactor) SendValue(opts *bind.TransactOpts, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "sendValue", _receiver, _amount)
}

// SendValue is a paid mutator transaction binding the contract method 0x24a084df.
//
// Solidity: function sendValue(address _receiver, uint256 _amount) returns()
func (_PositionManager *PositionManagerSession) SendValue(_receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.SendValue(&_PositionManager.TransactOpts, _receiver, _amount)
}

// SendValue is a paid mutator transaction binding the contract method 0x24a084df.
//
// Solidity: function sendValue(address _receiver, uint256 _amount) returns()
func (_PositionManager *PositionManagerTransactorSession) SendValue(_receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.SendValue(&_PositionManager.TransactOpts, _receiver, _amount)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_PositionManager *PositionManagerTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_PositionManager *PositionManagerSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.SetAdmin(&_PositionManager.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_PositionManager *PositionManagerTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.SetAdmin(&_PositionManager.TransactOpts, _admin)
}

// SetDepositFee is a paid mutator transaction binding the contract method 0x490ae210.
//
// Solidity: function setDepositFee(uint256 _depositFee) returns()
func (_PositionManager *PositionManagerTransactor) SetDepositFee(opts *bind.TransactOpts, _depositFee *big.Int) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setDepositFee", _depositFee)
}

// SetDepositFee is a paid mutator transaction binding the contract method 0x490ae210.
//
// Solidity: function setDepositFee(uint256 _depositFee) returns()
func (_PositionManager *PositionManagerSession) SetDepositFee(_depositFee *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.SetDepositFee(&_PositionManager.TransactOpts, _depositFee)
}

// SetDepositFee is a paid mutator transaction binding the contract method 0x490ae210.
//
// Solidity: function setDepositFee(uint256 _depositFee) returns()
func (_PositionManager *PositionManagerTransactorSession) SetDepositFee(_depositFee *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.SetDepositFee(&_PositionManager.TransactOpts, _depositFee)
}

// SetEthTransferGasLimit is a paid mutator transaction binding the contract method 0x3a9b52ad.
//
// Solidity: function setEthTransferGasLimit(uint256 _ethTransferGasLimit) returns()
func (_PositionManager *PositionManagerTransactor) SetEthTransferGasLimit(opts *bind.TransactOpts, _ethTransferGasLimit *big.Int) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setEthTransferGasLimit", _ethTransferGasLimit)
}

// SetEthTransferGasLimit is a paid mutator transaction binding the contract method 0x3a9b52ad.
//
// Solidity: function setEthTransferGasLimit(uint256 _ethTransferGasLimit) returns()
func (_PositionManager *PositionManagerSession) SetEthTransferGasLimit(_ethTransferGasLimit *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.SetEthTransferGasLimit(&_PositionManager.TransactOpts, _ethTransferGasLimit)
}

// SetEthTransferGasLimit is a paid mutator transaction binding the contract method 0x3a9b52ad.
//
// Solidity: function setEthTransferGasLimit(uint256 _ethTransferGasLimit) returns()
func (_PositionManager *PositionManagerTransactorSession) SetEthTransferGasLimit(_ethTransferGasLimit *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.SetEthTransferGasLimit(&_PositionManager.TransactOpts, _ethTransferGasLimit)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_PositionManager *PositionManagerTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_PositionManager *PositionManagerSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.SetGov(&_PositionManager.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_PositionManager *PositionManagerTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.SetGov(&_PositionManager.TransactOpts, _gov)
}

// SetInLegacyMode is a paid mutator transaction binding the contract method 0x9c95332f.
//
// Solidity: function setInLegacyMode(bool _inLegacyMode) returns()
func (_PositionManager *PositionManagerTransactor) SetInLegacyMode(opts *bind.TransactOpts, _inLegacyMode bool) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setInLegacyMode", _inLegacyMode)
}

// SetInLegacyMode is a paid mutator transaction binding the contract method 0x9c95332f.
//
// Solidity: function setInLegacyMode(bool _inLegacyMode) returns()
func (_PositionManager *PositionManagerSession) SetInLegacyMode(_inLegacyMode bool) (*types.Transaction, error) {
	return _PositionManager.Contract.SetInLegacyMode(&_PositionManager.TransactOpts, _inLegacyMode)
}

// SetInLegacyMode is a paid mutator transaction binding the contract method 0x9c95332f.
//
// Solidity: function setInLegacyMode(bool _inLegacyMode) returns()
func (_PositionManager *PositionManagerTransactorSession) SetInLegacyMode(_inLegacyMode bool) (*types.Transaction, error) {
	return _PositionManager.Contract.SetInLegacyMode(&_PositionManager.TransactOpts, _inLegacyMode)
}

// SetIncreasePositionBufferBps is a paid mutator transaction binding the contract method 0x233bfe3b.
//
// Solidity: function setIncreasePositionBufferBps(uint256 _increasePositionBufferBps) returns()
func (_PositionManager *PositionManagerTransactor) SetIncreasePositionBufferBps(opts *bind.TransactOpts, _increasePositionBufferBps *big.Int) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setIncreasePositionBufferBps", _increasePositionBufferBps)
}

// SetIncreasePositionBufferBps is a paid mutator transaction binding the contract method 0x233bfe3b.
//
// Solidity: function setIncreasePositionBufferBps(uint256 _increasePositionBufferBps) returns()
func (_PositionManager *PositionManagerSession) SetIncreasePositionBufferBps(_increasePositionBufferBps *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.SetIncreasePositionBufferBps(&_PositionManager.TransactOpts, _increasePositionBufferBps)
}

// SetIncreasePositionBufferBps is a paid mutator transaction binding the contract method 0x233bfe3b.
//
// Solidity: function setIncreasePositionBufferBps(uint256 _increasePositionBufferBps) returns()
func (_PositionManager *PositionManagerTransactorSession) SetIncreasePositionBufferBps(_increasePositionBufferBps *big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.SetIncreasePositionBufferBps(&_PositionManager.TransactOpts, _increasePositionBufferBps)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _account, bool _isActive) returns()
func (_PositionManager *PositionManagerTransactor) SetLiquidator(opts *bind.TransactOpts, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setLiquidator", _account, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _account, bool _isActive) returns()
func (_PositionManager *PositionManagerSession) SetLiquidator(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PositionManager.Contract.SetLiquidator(&_PositionManager.TransactOpts, _account, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _account, bool _isActive) returns()
func (_PositionManager *PositionManagerTransactorSession) SetLiquidator(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PositionManager.Contract.SetLiquidator(&_PositionManager.TransactOpts, _account, _isActive)
}

// SetMaxGlobalSizes is a paid mutator transaction binding the contract method 0xef12c67e.
//
// Solidity: function setMaxGlobalSizes(address[] _tokens, uint256[] _longSizes, uint256[] _shortSizes) returns()
func (_PositionManager *PositionManagerTransactor) SetMaxGlobalSizes(opts *bind.TransactOpts, _tokens []common.Address, _longSizes []*big.Int, _shortSizes []*big.Int) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setMaxGlobalSizes", _tokens, _longSizes, _shortSizes)
}

// SetMaxGlobalSizes is a paid mutator transaction binding the contract method 0xef12c67e.
//
// Solidity: function setMaxGlobalSizes(address[] _tokens, uint256[] _longSizes, uint256[] _shortSizes) returns()
func (_PositionManager *PositionManagerSession) SetMaxGlobalSizes(_tokens []common.Address, _longSizes []*big.Int, _shortSizes []*big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.SetMaxGlobalSizes(&_PositionManager.TransactOpts, _tokens, _longSizes, _shortSizes)
}

// SetMaxGlobalSizes is a paid mutator transaction binding the contract method 0xef12c67e.
//
// Solidity: function setMaxGlobalSizes(address[] _tokens, uint256[] _longSizes, uint256[] _shortSizes) returns()
func (_PositionManager *PositionManagerTransactorSession) SetMaxGlobalSizes(_tokens []common.Address, _longSizes []*big.Int, _shortSizes []*big.Int) (*types.Transaction, error) {
	return _PositionManager.Contract.SetMaxGlobalSizes(&_PositionManager.TransactOpts, _tokens, _longSizes, _shortSizes)
}

// SetOrderKeeper is a paid mutator transaction binding the contract method 0x1e261538.
//
// Solidity: function setOrderKeeper(address _account, bool _isActive) returns()
func (_PositionManager *PositionManagerTransactor) SetOrderKeeper(opts *bind.TransactOpts, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setOrderKeeper", _account, _isActive)
}

// SetOrderKeeper is a paid mutator transaction binding the contract method 0x1e261538.
//
// Solidity: function setOrderKeeper(address _account, bool _isActive) returns()
func (_PositionManager *PositionManagerSession) SetOrderKeeper(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PositionManager.Contract.SetOrderKeeper(&_PositionManager.TransactOpts, _account, _isActive)
}

// SetOrderKeeper is a paid mutator transaction binding the contract method 0x1e261538.
//
// Solidity: function setOrderKeeper(address _account, bool _isActive) returns()
func (_PositionManager *PositionManagerTransactorSession) SetOrderKeeper(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PositionManager.Contract.SetOrderKeeper(&_PositionManager.TransactOpts, _account, _isActive)
}

// SetPartner is a paid mutator transaction binding the contract method 0x21acf659.
//
// Solidity: function setPartner(address _account, bool _isActive) returns()
func (_PositionManager *PositionManagerTransactor) SetPartner(opts *bind.TransactOpts, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setPartner", _account, _isActive)
}

// SetPartner is a paid mutator transaction binding the contract method 0x21acf659.
//
// Solidity: function setPartner(address _account, bool _isActive) returns()
func (_PositionManager *PositionManagerSession) SetPartner(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PositionManager.Contract.SetPartner(&_PositionManager.TransactOpts, _account, _isActive)
}

// SetPartner is a paid mutator transaction binding the contract method 0x21acf659.
//
// Solidity: function setPartner(address _account, bool _isActive) returns()
func (_PositionManager *PositionManagerTransactorSession) SetPartner(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PositionManager.Contract.SetPartner(&_PositionManager.TransactOpts, _account, _isActive)
}

// SetReferralStorage is a paid mutator transaction binding the contract method 0xae4d7f9a.
//
// Solidity: function setReferralStorage(address _referralStorage) returns()
func (_PositionManager *PositionManagerTransactor) SetReferralStorage(opts *bind.TransactOpts, _referralStorage common.Address) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setReferralStorage", _referralStorage)
}

// SetReferralStorage is a paid mutator transaction binding the contract method 0xae4d7f9a.
//
// Solidity: function setReferralStorage(address _referralStorage) returns()
func (_PositionManager *PositionManagerSession) SetReferralStorage(_referralStorage common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.SetReferralStorage(&_PositionManager.TransactOpts, _referralStorage)
}

// SetReferralStorage is a paid mutator transaction binding the contract method 0xae4d7f9a.
//
// Solidity: function setReferralStorage(address _referralStorage) returns()
func (_PositionManager *PositionManagerTransactorSession) SetReferralStorage(_referralStorage common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.SetReferralStorage(&_PositionManager.TransactOpts, _referralStorage)
}

// SetShouldValidateIncreaseOrder is a paid mutator transaction binding the contract method 0x1b904359.
//
// Solidity: function setShouldValidateIncreaseOrder(bool _shouldValidateIncreaseOrder) returns()
func (_PositionManager *PositionManagerTransactor) SetShouldValidateIncreaseOrder(opts *bind.TransactOpts, _shouldValidateIncreaseOrder bool) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "setShouldValidateIncreaseOrder", _shouldValidateIncreaseOrder)
}

// SetShouldValidateIncreaseOrder is a paid mutator transaction binding the contract method 0x1b904359.
//
// Solidity: function setShouldValidateIncreaseOrder(bool _shouldValidateIncreaseOrder) returns()
func (_PositionManager *PositionManagerSession) SetShouldValidateIncreaseOrder(_shouldValidateIncreaseOrder bool) (*types.Transaction, error) {
	return _PositionManager.Contract.SetShouldValidateIncreaseOrder(&_PositionManager.TransactOpts, _shouldValidateIncreaseOrder)
}

// SetShouldValidateIncreaseOrder is a paid mutator transaction binding the contract method 0x1b904359.
//
// Solidity: function setShouldValidateIncreaseOrder(bool _shouldValidateIncreaseOrder) returns()
func (_PositionManager *PositionManagerTransactorSession) SetShouldValidateIncreaseOrder(_shouldValidateIncreaseOrder bool) (*types.Transaction, error) {
	return _PositionManager.Contract.SetShouldValidateIncreaseOrder(&_PositionManager.TransactOpts, _shouldValidateIncreaseOrder)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns()
func (_PositionManager *PositionManagerTransactor) WithdrawFees(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _PositionManager.contract.Transact(opts, "withdrawFees", _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns()
func (_PositionManager *PositionManagerSession) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.WithdrawFees(&_PositionManager.TransactOpts, _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns()
func (_PositionManager *PositionManagerTransactorSession) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _PositionManager.Contract.WithdrawFees(&_PositionManager.TransactOpts, _token, _receiver)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PositionManager *PositionManagerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionManager.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PositionManager *PositionManagerSession) Receive() (*types.Transaction, error) {
	return _PositionManager.Contract.Receive(&_PositionManager.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PositionManager *PositionManagerTransactorSession) Receive() (*types.Transaction, error) {
	return _PositionManager.Contract.Receive(&_PositionManager.TransactOpts)
}

// PositionManagerDecreasePositionReferralIterator is returned from FilterDecreasePositionReferral and is used to iterate over the raw logs and unpacked data for DecreasePositionReferral events raised by the PositionManager contract.
type PositionManagerDecreasePositionReferralIterator struct {
	Event *PositionManagerDecreasePositionReferral // Event containing the contract specifics and raw log

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
func (it *PositionManagerDecreasePositionReferralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerDecreasePositionReferral)
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
		it.Event = new(PositionManagerDecreasePositionReferral)
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
func (it *PositionManagerDecreasePositionReferralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerDecreasePositionReferralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerDecreasePositionReferral represents a DecreasePositionReferral event raised by the PositionManager contract.
type PositionManagerDecreasePositionReferral struct {
	Account              common.Address
	SizeDelta            *big.Int
	MarginFeeBasisPoints *big.Int
	ReferralCode         [32]byte
	Referrer             common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDecreasePositionReferral is a free log retrieval operation binding the contract event 0x474c763ff84bf2c2039a6d9fea955ecd0f724030e3c365b91169c6a16fe751b7.
//
// Solidity: event DecreasePositionReferral(address account, uint256 sizeDelta, uint256 marginFeeBasisPoints, bytes32 referralCode, address referrer)
func (_PositionManager *PositionManagerFilterer) FilterDecreasePositionReferral(opts *bind.FilterOpts) (*PositionManagerDecreasePositionReferralIterator, error) {

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "DecreasePositionReferral")
	if err != nil {
		return nil, err
	}
	return &PositionManagerDecreasePositionReferralIterator{contract: _PositionManager.contract, event: "DecreasePositionReferral", logs: logs, sub: sub}, nil
}

// WatchDecreasePositionReferral is a free log subscription operation binding the contract event 0x474c763ff84bf2c2039a6d9fea955ecd0f724030e3c365b91169c6a16fe751b7.
//
// Solidity: event DecreasePositionReferral(address account, uint256 sizeDelta, uint256 marginFeeBasisPoints, bytes32 referralCode, address referrer)
func (_PositionManager *PositionManagerFilterer) WatchDecreasePositionReferral(opts *bind.WatchOpts, sink chan<- *PositionManagerDecreasePositionReferral) (event.Subscription, error) {

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "DecreasePositionReferral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerDecreasePositionReferral)
				if err := _PositionManager.contract.UnpackLog(event, "DecreasePositionReferral", log); err != nil {
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

// ParseDecreasePositionReferral is a log parse operation binding the contract event 0x474c763ff84bf2c2039a6d9fea955ecd0f724030e3c365b91169c6a16fe751b7.
//
// Solidity: event DecreasePositionReferral(address account, uint256 sizeDelta, uint256 marginFeeBasisPoints, bytes32 referralCode, address referrer)
func (_PositionManager *PositionManagerFilterer) ParseDecreasePositionReferral(log types.Log) (*PositionManagerDecreasePositionReferral, error) {
	event := new(PositionManagerDecreasePositionReferral)
	if err := _PositionManager.contract.UnpackLog(event, "DecreasePositionReferral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerIncreasePositionReferralIterator is returned from FilterIncreasePositionReferral and is used to iterate over the raw logs and unpacked data for IncreasePositionReferral events raised by the PositionManager contract.
type PositionManagerIncreasePositionReferralIterator struct {
	Event *PositionManagerIncreasePositionReferral // Event containing the contract specifics and raw log

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
func (it *PositionManagerIncreasePositionReferralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerIncreasePositionReferral)
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
		it.Event = new(PositionManagerIncreasePositionReferral)
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
func (it *PositionManagerIncreasePositionReferralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerIncreasePositionReferralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerIncreasePositionReferral represents a IncreasePositionReferral event raised by the PositionManager contract.
type PositionManagerIncreasePositionReferral struct {
	Account              common.Address
	SizeDelta            *big.Int
	MarginFeeBasisPoints *big.Int
	ReferralCode         [32]byte
	Referrer             common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterIncreasePositionReferral is a free log retrieval operation binding the contract event 0xc2414023ce7002ee98557d1e7be21e5559073336f2217ee5f9b2e50fd85f71ee.
//
// Solidity: event IncreasePositionReferral(address account, uint256 sizeDelta, uint256 marginFeeBasisPoints, bytes32 referralCode, address referrer)
func (_PositionManager *PositionManagerFilterer) FilterIncreasePositionReferral(opts *bind.FilterOpts) (*PositionManagerIncreasePositionReferralIterator, error) {

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "IncreasePositionReferral")
	if err != nil {
		return nil, err
	}
	return &PositionManagerIncreasePositionReferralIterator{contract: _PositionManager.contract, event: "IncreasePositionReferral", logs: logs, sub: sub}, nil
}

// WatchIncreasePositionReferral is a free log subscription operation binding the contract event 0xc2414023ce7002ee98557d1e7be21e5559073336f2217ee5f9b2e50fd85f71ee.
//
// Solidity: event IncreasePositionReferral(address account, uint256 sizeDelta, uint256 marginFeeBasisPoints, bytes32 referralCode, address referrer)
func (_PositionManager *PositionManagerFilterer) WatchIncreasePositionReferral(opts *bind.WatchOpts, sink chan<- *PositionManagerIncreasePositionReferral) (event.Subscription, error) {

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "IncreasePositionReferral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerIncreasePositionReferral)
				if err := _PositionManager.contract.UnpackLog(event, "IncreasePositionReferral", log); err != nil {
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

// ParseIncreasePositionReferral is a log parse operation binding the contract event 0xc2414023ce7002ee98557d1e7be21e5559073336f2217ee5f9b2e50fd85f71ee.
//
// Solidity: event IncreasePositionReferral(address account, uint256 sizeDelta, uint256 marginFeeBasisPoints, bytes32 referralCode, address referrer)
func (_PositionManager *PositionManagerFilterer) ParseIncreasePositionReferral(log types.Log) (*PositionManagerIncreasePositionReferral, error) {
	event := new(PositionManagerIncreasePositionReferral)
	if err := _PositionManager.contract.UnpackLog(event, "IncreasePositionReferral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerSetAdminIterator is returned from FilterSetAdmin and is used to iterate over the raw logs and unpacked data for SetAdmin events raised by the PositionManager contract.
type PositionManagerSetAdminIterator struct {
	Event *PositionManagerSetAdmin // Event containing the contract specifics and raw log

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
func (it *PositionManagerSetAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerSetAdmin)
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
		it.Event = new(PositionManagerSetAdmin)
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
func (it *PositionManagerSetAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerSetAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerSetAdmin represents a SetAdmin event raised by the PositionManager contract.
type PositionManagerSetAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetAdmin is a free log retrieval operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_PositionManager *PositionManagerFilterer) FilterSetAdmin(opts *bind.FilterOpts) (*PositionManagerSetAdminIterator, error) {

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return &PositionManagerSetAdminIterator{contract: _PositionManager.contract, event: "SetAdmin", logs: logs, sub: sub}, nil
}

// WatchSetAdmin is a free log subscription operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_PositionManager *PositionManagerFilterer) WatchSetAdmin(opts *bind.WatchOpts, sink chan<- *PositionManagerSetAdmin) (event.Subscription, error) {

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerSetAdmin)
				if err := _PositionManager.contract.UnpackLog(event, "SetAdmin", log); err != nil {
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

// ParseSetAdmin is a log parse operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_PositionManager *PositionManagerFilterer) ParseSetAdmin(log types.Log) (*PositionManagerSetAdmin, error) {
	event := new(PositionManagerSetAdmin)
	if err := _PositionManager.contract.UnpackLog(event, "SetAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerSetDepositFeeIterator is returned from FilterSetDepositFee and is used to iterate over the raw logs and unpacked data for SetDepositFee events raised by the PositionManager contract.
type PositionManagerSetDepositFeeIterator struct {
	Event *PositionManagerSetDepositFee // Event containing the contract specifics and raw log

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
func (it *PositionManagerSetDepositFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerSetDepositFee)
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
		it.Event = new(PositionManagerSetDepositFee)
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
func (it *PositionManagerSetDepositFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerSetDepositFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerSetDepositFee represents a SetDepositFee event raised by the PositionManager contract.
type PositionManagerSetDepositFee struct {
	DepositFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetDepositFee is a free log retrieval operation binding the contract event 0x974fd3c1fcb4653dfc4fb740c4c692cd212d55c28f163f310128cb64d8300675.
//
// Solidity: event SetDepositFee(uint256 depositFee)
func (_PositionManager *PositionManagerFilterer) FilterSetDepositFee(opts *bind.FilterOpts) (*PositionManagerSetDepositFeeIterator, error) {

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "SetDepositFee")
	if err != nil {
		return nil, err
	}
	return &PositionManagerSetDepositFeeIterator{contract: _PositionManager.contract, event: "SetDepositFee", logs: logs, sub: sub}, nil
}

// WatchSetDepositFee is a free log subscription operation binding the contract event 0x974fd3c1fcb4653dfc4fb740c4c692cd212d55c28f163f310128cb64d8300675.
//
// Solidity: event SetDepositFee(uint256 depositFee)
func (_PositionManager *PositionManagerFilterer) WatchSetDepositFee(opts *bind.WatchOpts, sink chan<- *PositionManagerSetDepositFee) (event.Subscription, error) {

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "SetDepositFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerSetDepositFee)
				if err := _PositionManager.contract.UnpackLog(event, "SetDepositFee", log); err != nil {
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

// ParseSetDepositFee is a log parse operation binding the contract event 0x974fd3c1fcb4653dfc4fb740c4c692cd212d55c28f163f310128cb64d8300675.
//
// Solidity: event SetDepositFee(uint256 depositFee)
func (_PositionManager *PositionManagerFilterer) ParseSetDepositFee(log types.Log) (*PositionManagerSetDepositFee, error) {
	event := new(PositionManagerSetDepositFee)
	if err := _PositionManager.contract.UnpackLog(event, "SetDepositFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerSetEthTransferGasLimitIterator is returned from FilterSetEthTransferGasLimit and is used to iterate over the raw logs and unpacked data for SetEthTransferGasLimit events raised by the PositionManager contract.
type PositionManagerSetEthTransferGasLimitIterator struct {
	Event *PositionManagerSetEthTransferGasLimit // Event containing the contract specifics and raw log

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
func (it *PositionManagerSetEthTransferGasLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerSetEthTransferGasLimit)
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
		it.Event = new(PositionManagerSetEthTransferGasLimit)
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
func (it *PositionManagerSetEthTransferGasLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerSetEthTransferGasLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerSetEthTransferGasLimit represents a SetEthTransferGasLimit event raised by the PositionManager contract.
type PositionManagerSetEthTransferGasLimit struct {
	EthTransferGasLimit *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetEthTransferGasLimit is a free log retrieval operation binding the contract event 0x4d371d598d3a13f99ce992a17975bbaf1e1c256e072ec7d2f93ce88e40d9ba1c.
//
// Solidity: event SetEthTransferGasLimit(uint256 ethTransferGasLimit)
func (_PositionManager *PositionManagerFilterer) FilterSetEthTransferGasLimit(opts *bind.FilterOpts) (*PositionManagerSetEthTransferGasLimitIterator, error) {

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "SetEthTransferGasLimit")
	if err != nil {
		return nil, err
	}
	return &PositionManagerSetEthTransferGasLimitIterator{contract: _PositionManager.contract, event: "SetEthTransferGasLimit", logs: logs, sub: sub}, nil
}

// WatchSetEthTransferGasLimit is a free log subscription operation binding the contract event 0x4d371d598d3a13f99ce992a17975bbaf1e1c256e072ec7d2f93ce88e40d9ba1c.
//
// Solidity: event SetEthTransferGasLimit(uint256 ethTransferGasLimit)
func (_PositionManager *PositionManagerFilterer) WatchSetEthTransferGasLimit(opts *bind.WatchOpts, sink chan<- *PositionManagerSetEthTransferGasLimit) (event.Subscription, error) {

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "SetEthTransferGasLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerSetEthTransferGasLimit)
				if err := _PositionManager.contract.UnpackLog(event, "SetEthTransferGasLimit", log); err != nil {
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

// ParseSetEthTransferGasLimit is a log parse operation binding the contract event 0x4d371d598d3a13f99ce992a17975bbaf1e1c256e072ec7d2f93ce88e40d9ba1c.
//
// Solidity: event SetEthTransferGasLimit(uint256 ethTransferGasLimit)
func (_PositionManager *PositionManagerFilterer) ParseSetEthTransferGasLimit(log types.Log) (*PositionManagerSetEthTransferGasLimit, error) {
	event := new(PositionManagerSetEthTransferGasLimit)
	if err := _PositionManager.contract.UnpackLog(event, "SetEthTransferGasLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerSetInLegacyModeIterator is returned from FilterSetInLegacyMode and is used to iterate over the raw logs and unpacked data for SetInLegacyMode events raised by the PositionManager contract.
type PositionManagerSetInLegacyModeIterator struct {
	Event *PositionManagerSetInLegacyMode // Event containing the contract specifics and raw log

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
func (it *PositionManagerSetInLegacyModeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerSetInLegacyMode)
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
		it.Event = new(PositionManagerSetInLegacyMode)
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
func (it *PositionManagerSetInLegacyModeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerSetInLegacyModeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerSetInLegacyMode represents a SetInLegacyMode event raised by the PositionManager contract.
type PositionManagerSetInLegacyMode struct {
	InLegacyMode bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSetInLegacyMode is a free log retrieval operation binding the contract event 0xeac6b3611e79ff0d8ea5daa8439f6b1ab7eea4ebf95f1dd360417f712c3fc304.
//
// Solidity: event SetInLegacyMode(bool inLegacyMode)
func (_PositionManager *PositionManagerFilterer) FilterSetInLegacyMode(opts *bind.FilterOpts) (*PositionManagerSetInLegacyModeIterator, error) {

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "SetInLegacyMode")
	if err != nil {
		return nil, err
	}
	return &PositionManagerSetInLegacyModeIterator{contract: _PositionManager.contract, event: "SetInLegacyMode", logs: logs, sub: sub}, nil
}

// WatchSetInLegacyMode is a free log subscription operation binding the contract event 0xeac6b3611e79ff0d8ea5daa8439f6b1ab7eea4ebf95f1dd360417f712c3fc304.
//
// Solidity: event SetInLegacyMode(bool inLegacyMode)
func (_PositionManager *PositionManagerFilterer) WatchSetInLegacyMode(opts *bind.WatchOpts, sink chan<- *PositionManagerSetInLegacyMode) (event.Subscription, error) {

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "SetInLegacyMode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerSetInLegacyMode)
				if err := _PositionManager.contract.UnpackLog(event, "SetInLegacyMode", log); err != nil {
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

// ParseSetInLegacyMode is a log parse operation binding the contract event 0xeac6b3611e79ff0d8ea5daa8439f6b1ab7eea4ebf95f1dd360417f712c3fc304.
//
// Solidity: event SetInLegacyMode(bool inLegacyMode)
func (_PositionManager *PositionManagerFilterer) ParseSetInLegacyMode(log types.Log) (*PositionManagerSetInLegacyMode, error) {
	event := new(PositionManagerSetInLegacyMode)
	if err := _PositionManager.contract.UnpackLog(event, "SetInLegacyMode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerSetIncreasePositionBufferBpsIterator is returned from FilterSetIncreasePositionBufferBps and is used to iterate over the raw logs and unpacked data for SetIncreasePositionBufferBps events raised by the PositionManager contract.
type PositionManagerSetIncreasePositionBufferBpsIterator struct {
	Event *PositionManagerSetIncreasePositionBufferBps // Event containing the contract specifics and raw log

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
func (it *PositionManagerSetIncreasePositionBufferBpsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerSetIncreasePositionBufferBps)
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
		it.Event = new(PositionManagerSetIncreasePositionBufferBps)
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
func (it *PositionManagerSetIncreasePositionBufferBpsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerSetIncreasePositionBufferBpsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerSetIncreasePositionBufferBps represents a SetIncreasePositionBufferBps event raised by the PositionManager contract.
type PositionManagerSetIncreasePositionBufferBps struct {
	IncreasePositionBufferBps *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetIncreasePositionBufferBps is a free log retrieval operation binding the contract event 0x21167d0d4661af93817ebce920f18986eed3d75d5e1c03f2aed05efcbafbc452.
//
// Solidity: event SetIncreasePositionBufferBps(uint256 increasePositionBufferBps)
func (_PositionManager *PositionManagerFilterer) FilterSetIncreasePositionBufferBps(opts *bind.FilterOpts) (*PositionManagerSetIncreasePositionBufferBpsIterator, error) {

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "SetIncreasePositionBufferBps")
	if err != nil {
		return nil, err
	}
	return &PositionManagerSetIncreasePositionBufferBpsIterator{contract: _PositionManager.contract, event: "SetIncreasePositionBufferBps", logs: logs, sub: sub}, nil
}

// WatchSetIncreasePositionBufferBps is a free log subscription operation binding the contract event 0x21167d0d4661af93817ebce920f18986eed3d75d5e1c03f2aed05efcbafbc452.
//
// Solidity: event SetIncreasePositionBufferBps(uint256 increasePositionBufferBps)
func (_PositionManager *PositionManagerFilterer) WatchSetIncreasePositionBufferBps(opts *bind.WatchOpts, sink chan<- *PositionManagerSetIncreasePositionBufferBps) (event.Subscription, error) {

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "SetIncreasePositionBufferBps")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerSetIncreasePositionBufferBps)
				if err := _PositionManager.contract.UnpackLog(event, "SetIncreasePositionBufferBps", log); err != nil {
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

// ParseSetIncreasePositionBufferBps is a log parse operation binding the contract event 0x21167d0d4661af93817ebce920f18986eed3d75d5e1c03f2aed05efcbafbc452.
//
// Solidity: event SetIncreasePositionBufferBps(uint256 increasePositionBufferBps)
func (_PositionManager *PositionManagerFilterer) ParseSetIncreasePositionBufferBps(log types.Log) (*PositionManagerSetIncreasePositionBufferBps, error) {
	event := new(PositionManagerSetIncreasePositionBufferBps)
	if err := _PositionManager.contract.UnpackLog(event, "SetIncreasePositionBufferBps", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerSetLiquidatorIterator is returned from FilterSetLiquidator and is used to iterate over the raw logs and unpacked data for SetLiquidator events raised by the PositionManager contract.
type PositionManagerSetLiquidatorIterator struct {
	Event *PositionManagerSetLiquidator // Event containing the contract specifics and raw log

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
func (it *PositionManagerSetLiquidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerSetLiquidator)
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
		it.Event = new(PositionManagerSetLiquidator)
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
func (it *PositionManagerSetLiquidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerSetLiquidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerSetLiquidator represents a SetLiquidator event raised by the PositionManager contract.
type PositionManagerSetLiquidator struct {
	Account  common.Address
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetLiquidator is a free log retrieval operation binding the contract event 0x8c0d56805c3b43d441481229dc64bee168253ffe4305f37ab7cfe63b1c4268c6.
//
// Solidity: event SetLiquidator(address indexed account, bool isActive)
func (_PositionManager *PositionManagerFilterer) FilterSetLiquidator(opts *bind.FilterOpts, account []common.Address) (*PositionManagerSetLiquidatorIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "SetLiquidator", accountRule)
	if err != nil {
		return nil, err
	}
	return &PositionManagerSetLiquidatorIterator{contract: _PositionManager.contract, event: "SetLiquidator", logs: logs, sub: sub}, nil
}

// WatchSetLiquidator is a free log subscription operation binding the contract event 0x8c0d56805c3b43d441481229dc64bee168253ffe4305f37ab7cfe63b1c4268c6.
//
// Solidity: event SetLiquidator(address indexed account, bool isActive)
func (_PositionManager *PositionManagerFilterer) WatchSetLiquidator(opts *bind.WatchOpts, sink chan<- *PositionManagerSetLiquidator, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "SetLiquidator", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerSetLiquidator)
				if err := _PositionManager.contract.UnpackLog(event, "SetLiquidator", log); err != nil {
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

// ParseSetLiquidator is a log parse operation binding the contract event 0x8c0d56805c3b43d441481229dc64bee168253ffe4305f37ab7cfe63b1c4268c6.
//
// Solidity: event SetLiquidator(address indexed account, bool isActive)
func (_PositionManager *PositionManagerFilterer) ParseSetLiquidator(log types.Log) (*PositionManagerSetLiquidator, error) {
	event := new(PositionManagerSetLiquidator)
	if err := _PositionManager.contract.UnpackLog(event, "SetLiquidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerSetMaxGlobalSizesIterator is returned from FilterSetMaxGlobalSizes and is used to iterate over the raw logs and unpacked data for SetMaxGlobalSizes events raised by the PositionManager contract.
type PositionManagerSetMaxGlobalSizesIterator struct {
	Event *PositionManagerSetMaxGlobalSizes // Event containing the contract specifics and raw log

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
func (it *PositionManagerSetMaxGlobalSizesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerSetMaxGlobalSizes)
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
		it.Event = new(PositionManagerSetMaxGlobalSizes)
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
func (it *PositionManagerSetMaxGlobalSizesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerSetMaxGlobalSizesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerSetMaxGlobalSizes represents a SetMaxGlobalSizes event raised by the PositionManager contract.
type PositionManagerSetMaxGlobalSizes struct {
	Tokens     []common.Address
	LongSizes  []*big.Int
	ShortSizes []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetMaxGlobalSizes is a free log retrieval operation binding the contract event 0xae32d569b058895b9620d6552b09aaffedc9a6f396be4d595a224ad09f8b2139.
//
// Solidity: event SetMaxGlobalSizes(address[] tokens, uint256[] longSizes, uint256[] shortSizes)
func (_PositionManager *PositionManagerFilterer) FilterSetMaxGlobalSizes(opts *bind.FilterOpts) (*PositionManagerSetMaxGlobalSizesIterator, error) {

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "SetMaxGlobalSizes")
	if err != nil {
		return nil, err
	}
	return &PositionManagerSetMaxGlobalSizesIterator{contract: _PositionManager.contract, event: "SetMaxGlobalSizes", logs: logs, sub: sub}, nil
}

// WatchSetMaxGlobalSizes is a free log subscription operation binding the contract event 0xae32d569b058895b9620d6552b09aaffedc9a6f396be4d595a224ad09f8b2139.
//
// Solidity: event SetMaxGlobalSizes(address[] tokens, uint256[] longSizes, uint256[] shortSizes)
func (_PositionManager *PositionManagerFilterer) WatchSetMaxGlobalSizes(opts *bind.WatchOpts, sink chan<- *PositionManagerSetMaxGlobalSizes) (event.Subscription, error) {

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "SetMaxGlobalSizes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerSetMaxGlobalSizes)
				if err := _PositionManager.contract.UnpackLog(event, "SetMaxGlobalSizes", log); err != nil {
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

// ParseSetMaxGlobalSizes is a log parse operation binding the contract event 0xae32d569b058895b9620d6552b09aaffedc9a6f396be4d595a224ad09f8b2139.
//
// Solidity: event SetMaxGlobalSizes(address[] tokens, uint256[] longSizes, uint256[] shortSizes)
func (_PositionManager *PositionManagerFilterer) ParseSetMaxGlobalSizes(log types.Log) (*PositionManagerSetMaxGlobalSizes, error) {
	event := new(PositionManagerSetMaxGlobalSizes)
	if err := _PositionManager.contract.UnpackLog(event, "SetMaxGlobalSizes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerSetOrderKeeperIterator is returned from FilterSetOrderKeeper and is used to iterate over the raw logs and unpacked data for SetOrderKeeper events raised by the PositionManager contract.
type PositionManagerSetOrderKeeperIterator struct {
	Event *PositionManagerSetOrderKeeper // Event containing the contract specifics and raw log

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
func (it *PositionManagerSetOrderKeeperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerSetOrderKeeper)
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
		it.Event = new(PositionManagerSetOrderKeeper)
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
func (it *PositionManagerSetOrderKeeperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerSetOrderKeeperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerSetOrderKeeper represents a SetOrderKeeper event raised by the PositionManager contract.
type PositionManagerSetOrderKeeper struct {
	Account  common.Address
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetOrderKeeper is a free log retrieval operation binding the contract event 0x1d5bc0255b943d6a5b5279e8a55d74d620baccbceecb25e87a3558f14c4c118e.
//
// Solidity: event SetOrderKeeper(address indexed account, bool isActive)
func (_PositionManager *PositionManagerFilterer) FilterSetOrderKeeper(opts *bind.FilterOpts, account []common.Address) (*PositionManagerSetOrderKeeperIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "SetOrderKeeper", accountRule)
	if err != nil {
		return nil, err
	}
	return &PositionManagerSetOrderKeeperIterator{contract: _PositionManager.contract, event: "SetOrderKeeper", logs: logs, sub: sub}, nil
}

// WatchSetOrderKeeper is a free log subscription operation binding the contract event 0x1d5bc0255b943d6a5b5279e8a55d74d620baccbceecb25e87a3558f14c4c118e.
//
// Solidity: event SetOrderKeeper(address indexed account, bool isActive)
func (_PositionManager *PositionManagerFilterer) WatchSetOrderKeeper(opts *bind.WatchOpts, sink chan<- *PositionManagerSetOrderKeeper, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "SetOrderKeeper", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerSetOrderKeeper)
				if err := _PositionManager.contract.UnpackLog(event, "SetOrderKeeper", log); err != nil {
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

// ParseSetOrderKeeper is a log parse operation binding the contract event 0x1d5bc0255b943d6a5b5279e8a55d74d620baccbceecb25e87a3558f14c4c118e.
//
// Solidity: event SetOrderKeeper(address indexed account, bool isActive)
func (_PositionManager *PositionManagerFilterer) ParseSetOrderKeeper(log types.Log) (*PositionManagerSetOrderKeeper, error) {
	event := new(PositionManagerSetOrderKeeper)
	if err := _PositionManager.contract.UnpackLog(event, "SetOrderKeeper", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerSetPartnerIterator is returned from FilterSetPartner and is used to iterate over the raw logs and unpacked data for SetPartner events raised by the PositionManager contract.
type PositionManagerSetPartnerIterator struct {
	Event *PositionManagerSetPartner // Event containing the contract specifics and raw log

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
func (it *PositionManagerSetPartnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerSetPartner)
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
		it.Event = new(PositionManagerSetPartner)
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
func (it *PositionManagerSetPartnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerSetPartnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerSetPartner represents a SetPartner event raised by the PositionManager contract.
type PositionManagerSetPartner struct {
	Account  common.Address
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPartner is a free log retrieval operation binding the contract event 0xa4e46c70ff429a91de7d1716d736e877c7cca1c22ac850b23d242530dd95e474.
//
// Solidity: event SetPartner(address account, bool isActive)
func (_PositionManager *PositionManagerFilterer) FilterSetPartner(opts *bind.FilterOpts) (*PositionManagerSetPartnerIterator, error) {

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "SetPartner")
	if err != nil {
		return nil, err
	}
	return &PositionManagerSetPartnerIterator{contract: _PositionManager.contract, event: "SetPartner", logs: logs, sub: sub}, nil
}

// WatchSetPartner is a free log subscription operation binding the contract event 0xa4e46c70ff429a91de7d1716d736e877c7cca1c22ac850b23d242530dd95e474.
//
// Solidity: event SetPartner(address account, bool isActive)
func (_PositionManager *PositionManagerFilterer) WatchSetPartner(opts *bind.WatchOpts, sink chan<- *PositionManagerSetPartner) (event.Subscription, error) {

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "SetPartner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerSetPartner)
				if err := _PositionManager.contract.UnpackLog(event, "SetPartner", log); err != nil {
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

// ParseSetPartner is a log parse operation binding the contract event 0xa4e46c70ff429a91de7d1716d736e877c7cca1c22ac850b23d242530dd95e474.
//
// Solidity: event SetPartner(address account, bool isActive)
func (_PositionManager *PositionManagerFilterer) ParseSetPartner(log types.Log) (*PositionManagerSetPartner, error) {
	event := new(PositionManagerSetPartner)
	if err := _PositionManager.contract.UnpackLog(event, "SetPartner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerSetReferralStorageIterator is returned from FilterSetReferralStorage and is used to iterate over the raw logs and unpacked data for SetReferralStorage events raised by the PositionManager contract.
type PositionManagerSetReferralStorageIterator struct {
	Event *PositionManagerSetReferralStorage // Event containing the contract specifics and raw log

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
func (it *PositionManagerSetReferralStorageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerSetReferralStorage)
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
		it.Event = new(PositionManagerSetReferralStorage)
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
func (it *PositionManagerSetReferralStorageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerSetReferralStorageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerSetReferralStorage represents a SetReferralStorage event raised by the PositionManager contract.
type PositionManagerSetReferralStorage struct {
	ReferralStorage common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetReferralStorage is a free log retrieval operation binding the contract event 0x828abcccea18192c21d645e575652c49e20b986dab777906fc473d056b01b6a8.
//
// Solidity: event SetReferralStorage(address referralStorage)
func (_PositionManager *PositionManagerFilterer) FilterSetReferralStorage(opts *bind.FilterOpts) (*PositionManagerSetReferralStorageIterator, error) {

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "SetReferralStorage")
	if err != nil {
		return nil, err
	}
	return &PositionManagerSetReferralStorageIterator{contract: _PositionManager.contract, event: "SetReferralStorage", logs: logs, sub: sub}, nil
}

// WatchSetReferralStorage is a free log subscription operation binding the contract event 0x828abcccea18192c21d645e575652c49e20b986dab777906fc473d056b01b6a8.
//
// Solidity: event SetReferralStorage(address referralStorage)
func (_PositionManager *PositionManagerFilterer) WatchSetReferralStorage(opts *bind.WatchOpts, sink chan<- *PositionManagerSetReferralStorage) (event.Subscription, error) {

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "SetReferralStorage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerSetReferralStorage)
				if err := _PositionManager.contract.UnpackLog(event, "SetReferralStorage", log); err != nil {
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

// ParseSetReferralStorage is a log parse operation binding the contract event 0x828abcccea18192c21d645e575652c49e20b986dab777906fc473d056b01b6a8.
//
// Solidity: event SetReferralStorage(address referralStorage)
func (_PositionManager *PositionManagerFilterer) ParseSetReferralStorage(log types.Log) (*PositionManagerSetReferralStorage, error) {
	event := new(PositionManagerSetReferralStorage)
	if err := _PositionManager.contract.UnpackLog(event, "SetReferralStorage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerSetShouldValidateIncreaseOrderIterator is returned from FilterSetShouldValidateIncreaseOrder and is used to iterate over the raw logs and unpacked data for SetShouldValidateIncreaseOrder events raised by the PositionManager contract.
type PositionManagerSetShouldValidateIncreaseOrderIterator struct {
	Event *PositionManagerSetShouldValidateIncreaseOrder // Event containing the contract specifics and raw log

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
func (it *PositionManagerSetShouldValidateIncreaseOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerSetShouldValidateIncreaseOrder)
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
		it.Event = new(PositionManagerSetShouldValidateIncreaseOrder)
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
func (it *PositionManagerSetShouldValidateIncreaseOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerSetShouldValidateIncreaseOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerSetShouldValidateIncreaseOrder represents a SetShouldValidateIncreaseOrder event raised by the PositionManager contract.
type PositionManagerSetShouldValidateIncreaseOrder struct {
	ShouldValidateIncreaseOrder bool
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterSetShouldValidateIncreaseOrder is a free log retrieval operation binding the contract event 0xa956222e37fe025ff51e5440ac729a9bd417ff91e485e14dcffa2c0ba8894f40.
//
// Solidity: event SetShouldValidateIncreaseOrder(bool shouldValidateIncreaseOrder)
func (_PositionManager *PositionManagerFilterer) FilterSetShouldValidateIncreaseOrder(opts *bind.FilterOpts) (*PositionManagerSetShouldValidateIncreaseOrderIterator, error) {

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "SetShouldValidateIncreaseOrder")
	if err != nil {
		return nil, err
	}
	return &PositionManagerSetShouldValidateIncreaseOrderIterator{contract: _PositionManager.contract, event: "SetShouldValidateIncreaseOrder", logs: logs, sub: sub}, nil
}

// WatchSetShouldValidateIncreaseOrder is a free log subscription operation binding the contract event 0xa956222e37fe025ff51e5440ac729a9bd417ff91e485e14dcffa2c0ba8894f40.
//
// Solidity: event SetShouldValidateIncreaseOrder(bool shouldValidateIncreaseOrder)
func (_PositionManager *PositionManagerFilterer) WatchSetShouldValidateIncreaseOrder(opts *bind.WatchOpts, sink chan<- *PositionManagerSetShouldValidateIncreaseOrder) (event.Subscription, error) {

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "SetShouldValidateIncreaseOrder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerSetShouldValidateIncreaseOrder)
				if err := _PositionManager.contract.UnpackLog(event, "SetShouldValidateIncreaseOrder", log); err != nil {
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

// ParseSetShouldValidateIncreaseOrder is a log parse operation binding the contract event 0xa956222e37fe025ff51e5440ac729a9bd417ff91e485e14dcffa2c0ba8894f40.
//
// Solidity: event SetShouldValidateIncreaseOrder(bool shouldValidateIncreaseOrder)
func (_PositionManager *PositionManagerFilterer) ParseSetShouldValidateIncreaseOrder(log types.Log) (*PositionManagerSetShouldValidateIncreaseOrder, error) {
	event := new(PositionManagerSetShouldValidateIncreaseOrder)
	if err := _PositionManager.contract.UnpackLog(event, "SetShouldValidateIncreaseOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionManagerWithdrawFeesIterator is returned from FilterWithdrawFees and is used to iterate over the raw logs and unpacked data for WithdrawFees events raised by the PositionManager contract.
type PositionManagerWithdrawFeesIterator struct {
	Event *PositionManagerWithdrawFees // Event containing the contract specifics and raw log

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
func (it *PositionManagerWithdrawFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionManagerWithdrawFees)
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
		it.Event = new(PositionManagerWithdrawFees)
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
func (it *PositionManagerWithdrawFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionManagerWithdrawFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionManagerWithdrawFees represents a WithdrawFees event raised by the PositionManager contract.
type PositionManagerWithdrawFees struct {
	Token    common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFees is a free log retrieval operation binding the contract event 0x4f1b51dd7a2fcb861aa2670f668be66835c4ee12b4bbbf037e4d0018f39819e4.
//
// Solidity: event WithdrawFees(address token, address receiver, uint256 amount)
func (_PositionManager *PositionManagerFilterer) FilterWithdrawFees(opts *bind.FilterOpts) (*PositionManagerWithdrawFeesIterator, error) {

	logs, sub, err := _PositionManager.contract.FilterLogs(opts, "WithdrawFees")
	if err != nil {
		return nil, err
	}
	return &PositionManagerWithdrawFeesIterator{contract: _PositionManager.contract, event: "WithdrawFees", logs: logs, sub: sub}, nil
}

// WatchWithdrawFees is a free log subscription operation binding the contract event 0x4f1b51dd7a2fcb861aa2670f668be66835c4ee12b4bbbf037e4d0018f39819e4.
//
// Solidity: event WithdrawFees(address token, address receiver, uint256 amount)
func (_PositionManager *PositionManagerFilterer) WatchWithdrawFees(opts *bind.WatchOpts, sink chan<- *PositionManagerWithdrawFees) (event.Subscription, error) {

	logs, sub, err := _PositionManager.contract.WatchLogs(opts, "WithdrawFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionManagerWithdrawFees)
				if err := _PositionManager.contract.UnpackLog(event, "WithdrawFees", log); err != nil {
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

// ParseWithdrawFees is a log parse operation binding the contract event 0x4f1b51dd7a2fcb861aa2670f668be66835c4ee12b4bbbf037e4d0018f39819e4.
//
// Solidity: event WithdrawFees(address token, address receiver, uint256 amount)
func (_PositionManager *PositionManagerFilterer) ParseWithdrawFees(log types.Log) (*PositionManagerWithdrawFees, error) {
	event := new(PositionManagerWithdrawFees)
	if err := _PositionManager.contract.UnpackLog(event, "WithdrawFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GmxMigrator

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

// GmxMigratorMetaData contains all meta data concerning the GmxMigrator contract.
var GmxMigratorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_minAuthorizations\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"actionsNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ammRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"caps\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"endMigration\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getIouToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenAmounts\",\"inputs\":[{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenPrice\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gmxPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasMaxMigrationLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_ammRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_gmxPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_signers\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_whitelistedTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_iouTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_prices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_caps\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_lpTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_lpTokenAs\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_lpTokenBs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"iouTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isInitialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isMigrationActive\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSigner\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lpTokenAs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lpTokenBs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lpTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxMigrationAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrate\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"migratedAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minAuthorizations\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingActions\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"prices\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setHasMaxMigrationLimit\",\"inputs\":[{\"name\":\"_hasMaxMigrationLimit\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxMigrationAmount\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_maxMigrationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signApprove\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signalApprove\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signedActions\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"signers\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelistedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ClearAction\",\"inputs\":[{\"name\":\"action\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignAction\",\"inputs\":[{\"name\":\"action\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignalApprove\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignalPendingAction\",\"inputs\":[{\"name\":\"action\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// GmxMigratorABI is the input ABI used to generate the binding from.
// Deprecated: Use GmxMigratorMetaData.ABI instead.
var GmxMigratorABI = GmxMigratorMetaData.ABI

// GmxMigrator is an auto generated Go binding around an Ethereum contract.
type GmxMigrator struct {
	GmxMigratorCaller     // Read-only binding to the contract
	GmxMigratorTransactor // Write-only binding to the contract
	GmxMigratorFilterer   // Log filterer for contract events
}

// GmxMigratorCaller is an auto generated read-only Go binding around an Ethereum contract.
type GmxMigratorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmxMigratorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GmxMigratorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmxMigratorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GmxMigratorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GmxMigratorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GmxMigratorSession struct {
	Contract     *GmxMigrator      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GmxMigratorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GmxMigratorCallerSession struct {
	Contract *GmxMigratorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// GmxMigratorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GmxMigratorTransactorSession struct {
	Contract     *GmxMigratorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// GmxMigratorRaw is an auto generated low-level Go binding around an Ethereum contract.
type GmxMigratorRaw struct {
	Contract *GmxMigrator // Generic contract binding to access the raw methods on
}

// GmxMigratorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GmxMigratorCallerRaw struct {
	Contract *GmxMigratorCaller // Generic read-only contract binding to access the raw methods on
}

// GmxMigratorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GmxMigratorTransactorRaw struct {
	Contract *GmxMigratorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGmxMigrator creates a new instance of GmxMigrator, bound to a specific deployed contract.
func NewGmxMigrator(address common.Address, backend bind.ContractBackend) (*GmxMigrator, error) {
	contract, err := bindGmxMigrator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GmxMigrator{GmxMigratorCaller: GmxMigratorCaller{contract: contract}, GmxMigratorTransactor: GmxMigratorTransactor{contract: contract}, GmxMigratorFilterer: GmxMigratorFilterer{contract: contract}}, nil
}

// NewGmxMigratorCaller creates a new read-only instance of GmxMigrator, bound to a specific deployed contract.
func NewGmxMigratorCaller(address common.Address, caller bind.ContractCaller) (*GmxMigratorCaller, error) {
	contract, err := bindGmxMigrator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GmxMigratorCaller{contract: contract}, nil
}

// NewGmxMigratorTransactor creates a new write-only instance of GmxMigrator, bound to a specific deployed contract.
func NewGmxMigratorTransactor(address common.Address, transactor bind.ContractTransactor) (*GmxMigratorTransactor, error) {
	contract, err := bindGmxMigrator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GmxMigratorTransactor{contract: contract}, nil
}

// NewGmxMigratorFilterer creates a new log filterer instance of GmxMigrator, bound to a specific deployed contract.
func NewGmxMigratorFilterer(address common.Address, filterer bind.ContractFilterer) (*GmxMigratorFilterer, error) {
	contract, err := bindGmxMigrator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GmxMigratorFilterer{contract: contract}, nil
}

// bindGmxMigrator binds a generic wrapper to an already deployed contract.
func bindGmxMigrator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GmxMigratorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GmxMigrator *GmxMigratorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GmxMigrator.Contract.GmxMigratorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GmxMigrator *GmxMigratorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GmxMigrator.Contract.GmxMigratorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GmxMigrator *GmxMigratorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GmxMigrator.Contract.GmxMigratorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GmxMigrator *GmxMigratorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GmxMigrator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GmxMigrator *GmxMigratorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GmxMigrator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GmxMigrator *GmxMigratorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GmxMigrator.Contract.contract.Transact(opts, method, params...)
}

// ActionsNonce is a free data retrieval call binding the contract method 0xa4c04bb7.
//
// Solidity: function actionsNonce() view returns(uint256)
func (_GmxMigrator *GmxMigratorCaller) ActionsNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "actionsNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActionsNonce is a free data retrieval call binding the contract method 0xa4c04bb7.
//
// Solidity: function actionsNonce() view returns(uint256)
func (_GmxMigrator *GmxMigratorSession) ActionsNonce() (*big.Int, error) {
	return _GmxMigrator.Contract.ActionsNonce(&_GmxMigrator.CallOpts)
}

// ActionsNonce is a free data retrieval call binding the contract method 0xa4c04bb7.
//
// Solidity: function actionsNonce() view returns(uint256)
func (_GmxMigrator *GmxMigratorCallerSession) ActionsNonce() (*big.Int, error) {
	return _GmxMigrator.Contract.ActionsNonce(&_GmxMigrator.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_GmxMigrator *GmxMigratorCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_GmxMigrator *GmxMigratorSession) Admin() (common.Address, error) {
	return _GmxMigrator.Contract.Admin(&_GmxMigrator.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_GmxMigrator *GmxMigratorCallerSession) Admin() (common.Address, error) {
	return _GmxMigrator.Contract.Admin(&_GmxMigrator.CallOpts)
}

// AmmRouter is a free data retrieval call binding the contract method 0x90b3c4e7.
//
// Solidity: function ammRouter() view returns(address)
func (_GmxMigrator *GmxMigratorCaller) AmmRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "ammRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AmmRouter is a free data retrieval call binding the contract method 0x90b3c4e7.
//
// Solidity: function ammRouter() view returns(address)
func (_GmxMigrator *GmxMigratorSession) AmmRouter() (common.Address, error) {
	return _GmxMigrator.Contract.AmmRouter(&_GmxMigrator.CallOpts)
}

// AmmRouter is a free data retrieval call binding the contract method 0x90b3c4e7.
//
// Solidity: function ammRouter() view returns(address)
func (_GmxMigrator *GmxMigratorCallerSession) AmmRouter() (common.Address, error) {
	return _GmxMigrator.Contract.AmmRouter(&_GmxMigrator.CallOpts)
}

// Caps is a free data retrieval call binding the contract method 0x66d97b21.
//
// Solidity: function caps(address ) view returns(uint256)
func (_GmxMigrator *GmxMigratorCaller) Caps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "caps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Caps is a free data retrieval call binding the contract method 0x66d97b21.
//
// Solidity: function caps(address ) view returns(uint256)
func (_GmxMigrator *GmxMigratorSession) Caps(arg0 common.Address) (*big.Int, error) {
	return _GmxMigrator.Contract.Caps(&_GmxMigrator.CallOpts, arg0)
}

// Caps is a free data retrieval call binding the contract method 0x66d97b21.
//
// Solidity: function caps(address ) view returns(uint256)
func (_GmxMigrator *GmxMigratorCallerSession) Caps(arg0 common.Address) (*big.Int, error) {
	return _GmxMigrator.Contract.Caps(&_GmxMigrator.CallOpts, arg0)
}

// GetIouToken is a free data retrieval call binding the contract method 0x7efdec32.
//
// Solidity: function getIouToken(address _token) view returns(address)
func (_GmxMigrator *GmxMigratorCaller) GetIouToken(opts *bind.CallOpts, _token common.Address) (common.Address, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "getIouToken", _token)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIouToken is a free data retrieval call binding the contract method 0x7efdec32.
//
// Solidity: function getIouToken(address _token) view returns(address)
func (_GmxMigrator *GmxMigratorSession) GetIouToken(_token common.Address) (common.Address, error) {
	return _GmxMigrator.Contract.GetIouToken(&_GmxMigrator.CallOpts, _token)
}

// GetIouToken is a free data retrieval call binding the contract method 0x7efdec32.
//
// Solidity: function getIouToken(address _token) view returns(address)
func (_GmxMigrator *GmxMigratorCallerSession) GetIouToken(_token common.Address) (common.Address, error) {
	return _GmxMigrator.Contract.GetIouToken(&_GmxMigrator.CallOpts, _token)
}

// GetTokenAmounts is a free data retrieval call binding the contract method 0x7b0365a7.
//
// Solidity: function getTokenAmounts(address[] _tokens) view returns(uint256[])
func (_GmxMigrator *GmxMigratorCaller) GetTokenAmounts(opts *bind.CallOpts, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "getTokenAmounts", _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTokenAmounts is a free data retrieval call binding the contract method 0x7b0365a7.
//
// Solidity: function getTokenAmounts(address[] _tokens) view returns(uint256[])
func (_GmxMigrator *GmxMigratorSession) GetTokenAmounts(_tokens []common.Address) ([]*big.Int, error) {
	return _GmxMigrator.Contract.GetTokenAmounts(&_GmxMigrator.CallOpts, _tokens)
}

// GetTokenAmounts is a free data retrieval call binding the contract method 0x7b0365a7.
//
// Solidity: function getTokenAmounts(address[] _tokens) view returns(uint256[])
func (_GmxMigrator *GmxMigratorCallerSession) GetTokenAmounts(_tokens []common.Address) ([]*big.Int, error) {
	return _GmxMigrator.Contract.GetTokenAmounts(&_GmxMigrator.CallOpts, _tokens)
}

// GetTokenPrice is a free data retrieval call binding the contract method 0xd02641a0.
//
// Solidity: function getTokenPrice(address _token) view returns(uint256)
func (_GmxMigrator *GmxMigratorCaller) GetTokenPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "getTokenPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenPrice is a free data retrieval call binding the contract method 0xd02641a0.
//
// Solidity: function getTokenPrice(address _token) view returns(uint256)
func (_GmxMigrator *GmxMigratorSession) GetTokenPrice(_token common.Address) (*big.Int, error) {
	return _GmxMigrator.Contract.GetTokenPrice(&_GmxMigrator.CallOpts, _token)
}

// GetTokenPrice is a free data retrieval call binding the contract method 0xd02641a0.
//
// Solidity: function getTokenPrice(address _token) view returns(uint256)
func (_GmxMigrator *GmxMigratorCallerSession) GetTokenPrice(_token common.Address) (*big.Int, error) {
	return _GmxMigrator.Contract.GetTokenPrice(&_GmxMigrator.CallOpts, _token)
}

// GmxPrice is a free data retrieval call binding the contract method 0xa9d878ee.
//
// Solidity: function gmxPrice() view returns(uint256)
func (_GmxMigrator *GmxMigratorCaller) GmxPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "gmxPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GmxPrice is a free data retrieval call binding the contract method 0xa9d878ee.
//
// Solidity: function gmxPrice() view returns(uint256)
func (_GmxMigrator *GmxMigratorSession) GmxPrice() (*big.Int, error) {
	return _GmxMigrator.Contract.GmxPrice(&_GmxMigrator.CallOpts)
}

// GmxPrice is a free data retrieval call binding the contract method 0xa9d878ee.
//
// Solidity: function gmxPrice() view returns(uint256)
func (_GmxMigrator *GmxMigratorCallerSession) GmxPrice() (*big.Int, error) {
	return _GmxMigrator.Contract.GmxPrice(&_GmxMigrator.CallOpts)
}

// HasMaxMigrationLimit is a free data retrieval call binding the contract method 0x5183a7ca.
//
// Solidity: function hasMaxMigrationLimit() view returns(bool)
func (_GmxMigrator *GmxMigratorCaller) HasMaxMigrationLimit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "hasMaxMigrationLimit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMaxMigrationLimit is a free data retrieval call binding the contract method 0x5183a7ca.
//
// Solidity: function hasMaxMigrationLimit() view returns(bool)
func (_GmxMigrator *GmxMigratorSession) HasMaxMigrationLimit() (bool, error) {
	return _GmxMigrator.Contract.HasMaxMigrationLimit(&_GmxMigrator.CallOpts)
}

// HasMaxMigrationLimit is a free data retrieval call binding the contract method 0x5183a7ca.
//
// Solidity: function hasMaxMigrationLimit() view returns(bool)
func (_GmxMigrator *GmxMigratorCallerSession) HasMaxMigrationLimit() (bool, error) {
	return _GmxMigrator.Contract.HasMaxMigrationLimit(&_GmxMigrator.CallOpts)
}

// IouTokens is a free data retrieval call binding the contract method 0x41410d4a.
//
// Solidity: function iouTokens(address ) view returns(address)
func (_GmxMigrator *GmxMigratorCaller) IouTokens(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "iouTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IouTokens is a free data retrieval call binding the contract method 0x41410d4a.
//
// Solidity: function iouTokens(address ) view returns(address)
func (_GmxMigrator *GmxMigratorSession) IouTokens(arg0 common.Address) (common.Address, error) {
	return _GmxMigrator.Contract.IouTokens(&_GmxMigrator.CallOpts, arg0)
}

// IouTokens is a free data retrieval call binding the contract method 0x41410d4a.
//
// Solidity: function iouTokens(address ) view returns(address)
func (_GmxMigrator *GmxMigratorCallerSession) IouTokens(arg0 common.Address) (common.Address, error) {
	return _GmxMigrator.Contract.IouTokens(&_GmxMigrator.CallOpts, arg0)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_GmxMigrator *GmxMigratorCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_GmxMigrator *GmxMigratorSession) IsInitialized() (bool, error) {
	return _GmxMigrator.Contract.IsInitialized(&_GmxMigrator.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_GmxMigrator *GmxMigratorCallerSession) IsInitialized() (bool, error) {
	return _GmxMigrator.Contract.IsInitialized(&_GmxMigrator.CallOpts)
}

// IsMigrationActive is a free data retrieval call binding the contract method 0x0e913a6c.
//
// Solidity: function isMigrationActive() view returns(bool)
func (_GmxMigrator *GmxMigratorCaller) IsMigrationActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "isMigrationActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMigrationActive is a free data retrieval call binding the contract method 0x0e913a6c.
//
// Solidity: function isMigrationActive() view returns(bool)
func (_GmxMigrator *GmxMigratorSession) IsMigrationActive() (bool, error) {
	return _GmxMigrator.Contract.IsMigrationActive(&_GmxMigrator.CallOpts)
}

// IsMigrationActive is a free data retrieval call binding the contract method 0x0e913a6c.
//
// Solidity: function isMigrationActive() view returns(bool)
func (_GmxMigrator *GmxMigratorCallerSession) IsMigrationActive() (bool, error) {
	return _GmxMigrator.Contract.IsMigrationActive(&_GmxMigrator.CallOpts)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) view returns(bool)
func (_GmxMigrator *GmxMigratorCaller) IsSigner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "isSigner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) view returns(bool)
func (_GmxMigrator *GmxMigratorSession) IsSigner(arg0 common.Address) (bool, error) {
	return _GmxMigrator.Contract.IsSigner(&_GmxMigrator.CallOpts, arg0)
}

// IsSigner is a free data retrieval call binding the contract method 0x7df73e27.
//
// Solidity: function isSigner(address ) view returns(bool)
func (_GmxMigrator *GmxMigratorCallerSession) IsSigner(arg0 common.Address) (bool, error) {
	return _GmxMigrator.Contract.IsSigner(&_GmxMigrator.CallOpts, arg0)
}

// LpTokenAs is a free data retrieval call binding the contract method 0x43b45f35.
//
// Solidity: function lpTokenAs(address ) view returns(address)
func (_GmxMigrator *GmxMigratorCaller) LpTokenAs(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "lpTokenAs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpTokenAs is a free data retrieval call binding the contract method 0x43b45f35.
//
// Solidity: function lpTokenAs(address ) view returns(address)
func (_GmxMigrator *GmxMigratorSession) LpTokenAs(arg0 common.Address) (common.Address, error) {
	return _GmxMigrator.Contract.LpTokenAs(&_GmxMigrator.CallOpts, arg0)
}

// LpTokenAs is a free data retrieval call binding the contract method 0x43b45f35.
//
// Solidity: function lpTokenAs(address ) view returns(address)
func (_GmxMigrator *GmxMigratorCallerSession) LpTokenAs(arg0 common.Address) (common.Address, error) {
	return _GmxMigrator.Contract.LpTokenAs(&_GmxMigrator.CallOpts, arg0)
}

// LpTokenBs is a free data retrieval call binding the contract method 0xe699e048.
//
// Solidity: function lpTokenBs(address ) view returns(address)
func (_GmxMigrator *GmxMigratorCaller) LpTokenBs(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "lpTokenBs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpTokenBs is a free data retrieval call binding the contract method 0xe699e048.
//
// Solidity: function lpTokenBs(address ) view returns(address)
func (_GmxMigrator *GmxMigratorSession) LpTokenBs(arg0 common.Address) (common.Address, error) {
	return _GmxMigrator.Contract.LpTokenBs(&_GmxMigrator.CallOpts, arg0)
}

// LpTokenBs is a free data retrieval call binding the contract method 0xe699e048.
//
// Solidity: function lpTokenBs(address ) view returns(address)
func (_GmxMigrator *GmxMigratorCallerSession) LpTokenBs(arg0 common.Address) (common.Address, error) {
	return _GmxMigrator.Contract.LpTokenBs(&_GmxMigrator.CallOpts, arg0)
}

// LpTokens is a free data retrieval call binding the contract method 0xb17b658d.
//
// Solidity: function lpTokens(address ) view returns(bool)
func (_GmxMigrator *GmxMigratorCaller) LpTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "lpTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// LpTokens is a free data retrieval call binding the contract method 0xb17b658d.
//
// Solidity: function lpTokens(address ) view returns(bool)
func (_GmxMigrator *GmxMigratorSession) LpTokens(arg0 common.Address) (bool, error) {
	return _GmxMigrator.Contract.LpTokens(&_GmxMigrator.CallOpts, arg0)
}

// LpTokens is a free data retrieval call binding the contract method 0xb17b658d.
//
// Solidity: function lpTokens(address ) view returns(bool)
func (_GmxMigrator *GmxMigratorCallerSession) LpTokens(arg0 common.Address) (bool, error) {
	return _GmxMigrator.Contract.LpTokens(&_GmxMigrator.CallOpts, arg0)
}

// MaxMigrationAmounts is a free data retrieval call binding the contract method 0xceb6250e.
//
// Solidity: function maxMigrationAmounts(address , address ) view returns(uint256)
func (_GmxMigrator *GmxMigratorCaller) MaxMigrationAmounts(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "maxMigrationAmounts", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMigrationAmounts is a free data retrieval call binding the contract method 0xceb6250e.
//
// Solidity: function maxMigrationAmounts(address , address ) view returns(uint256)
func (_GmxMigrator *GmxMigratorSession) MaxMigrationAmounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GmxMigrator.Contract.MaxMigrationAmounts(&_GmxMigrator.CallOpts, arg0, arg1)
}

// MaxMigrationAmounts is a free data retrieval call binding the contract method 0xceb6250e.
//
// Solidity: function maxMigrationAmounts(address , address ) view returns(uint256)
func (_GmxMigrator *GmxMigratorCallerSession) MaxMigrationAmounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GmxMigrator.Contract.MaxMigrationAmounts(&_GmxMigrator.CallOpts, arg0, arg1)
}

// MigratedAmounts is a free data retrieval call binding the contract method 0xa758f6dd.
//
// Solidity: function migratedAmounts(address , address ) view returns(uint256)
func (_GmxMigrator *GmxMigratorCaller) MigratedAmounts(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "migratedAmounts", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MigratedAmounts is a free data retrieval call binding the contract method 0xa758f6dd.
//
// Solidity: function migratedAmounts(address , address ) view returns(uint256)
func (_GmxMigrator *GmxMigratorSession) MigratedAmounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GmxMigrator.Contract.MigratedAmounts(&_GmxMigrator.CallOpts, arg0, arg1)
}

// MigratedAmounts is a free data retrieval call binding the contract method 0xa758f6dd.
//
// Solidity: function migratedAmounts(address , address ) view returns(uint256)
func (_GmxMigrator *GmxMigratorCallerSession) MigratedAmounts(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _GmxMigrator.Contract.MigratedAmounts(&_GmxMigrator.CallOpts, arg0, arg1)
}

// MinAuthorizations is a free data retrieval call binding the contract method 0x287800c9.
//
// Solidity: function minAuthorizations() view returns(uint256)
func (_GmxMigrator *GmxMigratorCaller) MinAuthorizations(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "minAuthorizations")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAuthorizations is a free data retrieval call binding the contract method 0x287800c9.
//
// Solidity: function minAuthorizations() view returns(uint256)
func (_GmxMigrator *GmxMigratorSession) MinAuthorizations() (*big.Int, error) {
	return _GmxMigrator.Contract.MinAuthorizations(&_GmxMigrator.CallOpts)
}

// MinAuthorizations is a free data retrieval call binding the contract method 0x287800c9.
//
// Solidity: function minAuthorizations() view returns(uint256)
func (_GmxMigrator *GmxMigratorCallerSession) MinAuthorizations() (*big.Int, error) {
	return _GmxMigrator.Contract.MinAuthorizations(&_GmxMigrator.CallOpts)
}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(bool)
func (_GmxMigrator *GmxMigratorCaller) PendingActions(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "pendingActions", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(bool)
func (_GmxMigrator *GmxMigratorSession) PendingActions(arg0 [32]byte) (bool, error) {
	return _GmxMigrator.Contract.PendingActions(&_GmxMigrator.CallOpts, arg0)
}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(bool)
func (_GmxMigrator *GmxMigratorCallerSession) PendingActions(arg0 [32]byte) (bool, error) {
	return _GmxMigrator.Contract.PendingActions(&_GmxMigrator.CallOpts, arg0)
}

// Prices is a free data retrieval call binding the contract method 0xcfed246b.
//
// Solidity: function prices(address ) view returns(uint256)
func (_GmxMigrator *GmxMigratorCaller) Prices(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "prices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Prices is a free data retrieval call binding the contract method 0xcfed246b.
//
// Solidity: function prices(address ) view returns(uint256)
func (_GmxMigrator *GmxMigratorSession) Prices(arg0 common.Address) (*big.Int, error) {
	return _GmxMigrator.Contract.Prices(&_GmxMigrator.CallOpts, arg0)
}

// Prices is a free data retrieval call binding the contract method 0xcfed246b.
//
// Solidity: function prices(address ) view returns(uint256)
func (_GmxMigrator *GmxMigratorCallerSession) Prices(arg0 common.Address) (*big.Int, error) {
	return _GmxMigrator.Contract.Prices(&_GmxMigrator.CallOpts, arg0)
}

// SignedActions is a free data retrieval call binding the contract method 0x87c6d4f9.
//
// Solidity: function signedActions(address , bytes32 ) view returns(bool)
func (_GmxMigrator *GmxMigratorCaller) SignedActions(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "signedActions", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SignedActions is a free data retrieval call binding the contract method 0x87c6d4f9.
//
// Solidity: function signedActions(address , bytes32 ) view returns(bool)
func (_GmxMigrator *GmxMigratorSession) SignedActions(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _GmxMigrator.Contract.SignedActions(&_GmxMigrator.CallOpts, arg0, arg1)
}

// SignedActions is a free data retrieval call binding the contract method 0x87c6d4f9.
//
// Solidity: function signedActions(address , bytes32 ) view returns(bool)
func (_GmxMigrator *GmxMigratorCallerSession) SignedActions(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _GmxMigrator.Contract.SignedActions(&_GmxMigrator.CallOpts, arg0, arg1)
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_GmxMigrator *GmxMigratorCaller) Signers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "signers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_GmxMigrator *GmxMigratorSession) Signers(arg0 *big.Int) (common.Address, error) {
	return _GmxMigrator.Contract.Signers(&_GmxMigrator.CallOpts, arg0)
}

// Signers is a free data retrieval call binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) view returns(address)
func (_GmxMigrator *GmxMigratorCallerSession) Signers(arg0 *big.Int) (common.Address, error) {
	return _GmxMigrator.Contract.Signers(&_GmxMigrator.CallOpts, arg0)
}

// TokenAmounts is a free data retrieval call binding the contract method 0xa0e2e5f6.
//
// Solidity: function tokenAmounts(address ) view returns(uint256)
func (_GmxMigrator *GmxMigratorCaller) TokenAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "tokenAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenAmounts is a free data retrieval call binding the contract method 0xa0e2e5f6.
//
// Solidity: function tokenAmounts(address ) view returns(uint256)
func (_GmxMigrator *GmxMigratorSession) TokenAmounts(arg0 common.Address) (*big.Int, error) {
	return _GmxMigrator.Contract.TokenAmounts(&_GmxMigrator.CallOpts, arg0)
}

// TokenAmounts is a free data retrieval call binding the contract method 0xa0e2e5f6.
//
// Solidity: function tokenAmounts(address ) view returns(uint256)
func (_GmxMigrator *GmxMigratorCallerSession) TokenAmounts(arg0 common.Address) (*big.Int, error) {
	return _GmxMigrator.Contract.TokenAmounts(&_GmxMigrator.CallOpts, arg0)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_GmxMigrator *GmxMigratorCaller) WhitelistedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GmxMigrator.contract.Call(opts, &out, "whitelistedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_GmxMigrator *GmxMigratorSession) WhitelistedTokens(arg0 common.Address) (bool, error) {
	return _GmxMigrator.Contract.WhitelistedTokens(&_GmxMigrator.CallOpts, arg0)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_GmxMigrator *GmxMigratorCallerSession) WhitelistedTokens(arg0 common.Address) (bool, error) {
	return _GmxMigrator.Contract.WhitelistedTokens(&_GmxMigrator.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x4dc5ecb3.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount, uint256 _nonce) returns()
func (_GmxMigrator *GmxMigratorTransactor) Approve(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _amount *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxMigrator.contract.Transact(opts, "approve", _token, _spender, _amount, _nonce)
}

// Approve is a paid mutator transaction binding the contract method 0x4dc5ecb3.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount, uint256 _nonce) returns()
func (_GmxMigrator *GmxMigratorSession) Approve(_token common.Address, _spender common.Address, _amount *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxMigrator.Contract.Approve(&_GmxMigrator.TransactOpts, _token, _spender, _amount, _nonce)
}

// Approve is a paid mutator transaction binding the contract method 0x4dc5ecb3.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount, uint256 _nonce) returns()
func (_GmxMigrator *GmxMigratorTransactorSession) Approve(_token common.Address, _spender common.Address, _amount *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxMigrator.Contract.Approve(&_GmxMigrator.TransactOpts, _token, _spender, _amount, _nonce)
}

// EndMigration is a paid mutator transaction binding the contract method 0x6c525d04.
//
// Solidity: function endMigration() returns()
func (_GmxMigrator *GmxMigratorTransactor) EndMigration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GmxMigrator.contract.Transact(opts, "endMigration")
}

// EndMigration is a paid mutator transaction binding the contract method 0x6c525d04.
//
// Solidity: function endMigration() returns()
func (_GmxMigrator *GmxMigratorSession) EndMigration() (*types.Transaction, error) {
	return _GmxMigrator.Contract.EndMigration(&_GmxMigrator.TransactOpts)
}

// EndMigration is a paid mutator transaction binding the contract method 0x6c525d04.
//
// Solidity: function endMigration() returns()
func (_GmxMigrator *GmxMigratorTransactorSession) EndMigration() (*types.Transaction, error) {
	return _GmxMigrator.Contract.EndMigration(&_GmxMigrator.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x64d4b173.
//
// Solidity: function initialize(address _ammRouter, uint256 _gmxPrice, address[] _signers, address[] _whitelistedTokens, address[] _iouTokens, uint256[] _prices, uint256[] _caps, address[] _lpTokens, address[] _lpTokenAs, address[] _lpTokenBs) returns()
func (_GmxMigrator *GmxMigratorTransactor) Initialize(opts *bind.TransactOpts, _ammRouter common.Address, _gmxPrice *big.Int, _signers []common.Address, _whitelistedTokens []common.Address, _iouTokens []common.Address, _prices []*big.Int, _caps []*big.Int, _lpTokens []common.Address, _lpTokenAs []common.Address, _lpTokenBs []common.Address) (*types.Transaction, error) {
	return _GmxMigrator.contract.Transact(opts, "initialize", _ammRouter, _gmxPrice, _signers, _whitelistedTokens, _iouTokens, _prices, _caps, _lpTokens, _lpTokenAs, _lpTokenBs)
}

// Initialize is a paid mutator transaction binding the contract method 0x64d4b173.
//
// Solidity: function initialize(address _ammRouter, uint256 _gmxPrice, address[] _signers, address[] _whitelistedTokens, address[] _iouTokens, uint256[] _prices, uint256[] _caps, address[] _lpTokens, address[] _lpTokenAs, address[] _lpTokenBs) returns()
func (_GmxMigrator *GmxMigratorSession) Initialize(_ammRouter common.Address, _gmxPrice *big.Int, _signers []common.Address, _whitelistedTokens []common.Address, _iouTokens []common.Address, _prices []*big.Int, _caps []*big.Int, _lpTokens []common.Address, _lpTokenAs []common.Address, _lpTokenBs []common.Address) (*types.Transaction, error) {
	return _GmxMigrator.Contract.Initialize(&_GmxMigrator.TransactOpts, _ammRouter, _gmxPrice, _signers, _whitelistedTokens, _iouTokens, _prices, _caps, _lpTokens, _lpTokenAs, _lpTokenBs)
}

// Initialize is a paid mutator transaction binding the contract method 0x64d4b173.
//
// Solidity: function initialize(address _ammRouter, uint256 _gmxPrice, address[] _signers, address[] _whitelistedTokens, address[] _iouTokens, uint256[] _prices, uint256[] _caps, address[] _lpTokens, address[] _lpTokenAs, address[] _lpTokenBs) returns()
func (_GmxMigrator *GmxMigratorTransactorSession) Initialize(_ammRouter common.Address, _gmxPrice *big.Int, _signers []common.Address, _whitelistedTokens []common.Address, _iouTokens []common.Address, _prices []*big.Int, _caps []*big.Int, _lpTokens []common.Address, _lpTokenAs []common.Address, _lpTokenBs []common.Address) (*types.Transaction, error) {
	return _GmxMigrator.Contract.Initialize(&_GmxMigrator.TransactOpts, _ammRouter, _gmxPrice, _signers, _whitelistedTokens, _iouTokens, _prices, _caps, _lpTokens, _lpTokenAs, _lpTokenBs)
}

// Migrate is a paid mutator transaction binding the contract method 0xad68ebf7.
//
// Solidity: function migrate(address _token, uint256 _tokenAmount) returns()
func (_GmxMigrator *GmxMigratorTransactor) Migrate(opts *bind.TransactOpts, _token common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _GmxMigrator.contract.Transact(opts, "migrate", _token, _tokenAmount)
}

// Migrate is a paid mutator transaction binding the contract method 0xad68ebf7.
//
// Solidity: function migrate(address _token, uint256 _tokenAmount) returns()
func (_GmxMigrator *GmxMigratorSession) Migrate(_token common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _GmxMigrator.Contract.Migrate(&_GmxMigrator.TransactOpts, _token, _tokenAmount)
}

// Migrate is a paid mutator transaction binding the contract method 0xad68ebf7.
//
// Solidity: function migrate(address _token, uint256 _tokenAmount) returns()
func (_GmxMigrator *GmxMigratorTransactorSession) Migrate(_token common.Address, _tokenAmount *big.Int) (*types.Transaction, error) {
	return _GmxMigrator.Contract.Migrate(&_GmxMigrator.TransactOpts, _token, _tokenAmount)
}

// SetHasMaxMigrationLimit is a paid mutator transaction binding the contract method 0x0122f868.
//
// Solidity: function setHasMaxMigrationLimit(bool _hasMaxMigrationLimit) returns()
func (_GmxMigrator *GmxMigratorTransactor) SetHasMaxMigrationLimit(opts *bind.TransactOpts, _hasMaxMigrationLimit bool) (*types.Transaction, error) {
	return _GmxMigrator.contract.Transact(opts, "setHasMaxMigrationLimit", _hasMaxMigrationLimit)
}

// SetHasMaxMigrationLimit is a paid mutator transaction binding the contract method 0x0122f868.
//
// Solidity: function setHasMaxMigrationLimit(bool _hasMaxMigrationLimit) returns()
func (_GmxMigrator *GmxMigratorSession) SetHasMaxMigrationLimit(_hasMaxMigrationLimit bool) (*types.Transaction, error) {
	return _GmxMigrator.Contract.SetHasMaxMigrationLimit(&_GmxMigrator.TransactOpts, _hasMaxMigrationLimit)
}

// SetHasMaxMigrationLimit is a paid mutator transaction binding the contract method 0x0122f868.
//
// Solidity: function setHasMaxMigrationLimit(bool _hasMaxMigrationLimit) returns()
func (_GmxMigrator *GmxMigratorTransactorSession) SetHasMaxMigrationLimit(_hasMaxMigrationLimit bool) (*types.Transaction, error) {
	return _GmxMigrator.Contract.SetHasMaxMigrationLimit(&_GmxMigrator.TransactOpts, _hasMaxMigrationLimit)
}

// SetMaxMigrationAmount is a paid mutator transaction binding the contract method 0x91835527.
//
// Solidity: function setMaxMigrationAmount(address _account, address _token, uint256 _maxMigrationAmount) returns()
func (_GmxMigrator *GmxMigratorTransactor) SetMaxMigrationAmount(opts *bind.TransactOpts, _account common.Address, _token common.Address, _maxMigrationAmount *big.Int) (*types.Transaction, error) {
	return _GmxMigrator.contract.Transact(opts, "setMaxMigrationAmount", _account, _token, _maxMigrationAmount)
}

// SetMaxMigrationAmount is a paid mutator transaction binding the contract method 0x91835527.
//
// Solidity: function setMaxMigrationAmount(address _account, address _token, uint256 _maxMigrationAmount) returns()
func (_GmxMigrator *GmxMigratorSession) SetMaxMigrationAmount(_account common.Address, _token common.Address, _maxMigrationAmount *big.Int) (*types.Transaction, error) {
	return _GmxMigrator.Contract.SetMaxMigrationAmount(&_GmxMigrator.TransactOpts, _account, _token, _maxMigrationAmount)
}

// SetMaxMigrationAmount is a paid mutator transaction binding the contract method 0x91835527.
//
// Solidity: function setMaxMigrationAmount(address _account, address _token, uint256 _maxMigrationAmount) returns()
func (_GmxMigrator *GmxMigratorTransactorSession) SetMaxMigrationAmount(_account common.Address, _token common.Address, _maxMigrationAmount *big.Int) (*types.Transaction, error) {
	return _GmxMigrator.Contract.SetMaxMigrationAmount(&_GmxMigrator.TransactOpts, _account, _token, _maxMigrationAmount)
}

// SignApprove is a paid mutator transaction binding the contract method 0xf52dc4f7.
//
// Solidity: function signApprove(address _token, address _spender, uint256 _amount, uint256 _nonce) returns()
func (_GmxMigrator *GmxMigratorTransactor) SignApprove(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _amount *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxMigrator.contract.Transact(opts, "signApprove", _token, _spender, _amount, _nonce)
}

// SignApprove is a paid mutator transaction binding the contract method 0xf52dc4f7.
//
// Solidity: function signApprove(address _token, address _spender, uint256 _amount, uint256 _nonce) returns()
func (_GmxMigrator *GmxMigratorSession) SignApprove(_token common.Address, _spender common.Address, _amount *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxMigrator.Contract.SignApprove(&_GmxMigrator.TransactOpts, _token, _spender, _amount, _nonce)
}

// SignApprove is a paid mutator transaction binding the contract method 0xf52dc4f7.
//
// Solidity: function signApprove(address _token, address _spender, uint256 _amount, uint256 _nonce) returns()
func (_GmxMigrator *GmxMigratorTransactorSession) SignApprove(_token common.Address, _spender common.Address, _amount *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _GmxMigrator.Contract.SignApprove(&_GmxMigrator.TransactOpts, _token, _spender, _amount, _nonce)
}

// SignalApprove is a paid mutator transaction binding the contract method 0xdce6e18d.
//
// Solidity: function signalApprove(address _token, address _spender, uint256 _amount) returns()
func (_GmxMigrator *GmxMigratorTransactor) SignalApprove(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxMigrator.contract.Transact(opts, "signalApprove", _token, _spender, _amount)
}

// SignalApprove is a paid mutator transaction binding the contract method 0xdce6e18d.
//
// Solidity: function signalApprove(address _token, address _spender, uint256 _amount) returns()
func (_GmxMigrator *GmxMigratorSession) SignalApprove(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxMigrator.Contract.SignalApprove(&_GmxMigrator.TransactOpts, _token, _spender, _amount)
}

// SignalApprove is a paid mutator transaction binding the contract method 0xdce6e18d.
//
// Solidity: function signalApprove(address _token, address _spender, uint256 _amount) returns()
func (_GmxMigrator *GmxMigratorTransactorSession) SignalApprove(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _GmxMigrator.Contract.SignalApprove(&_GmxMigrator.TransactOpts, _token, _spender, _amount)
}

// GmxMigratorClearActionIterator is returned from FilterClearAction and is used to iterate over the raw logs and unpacked data for ClearAction events raised by the GmxMigrator contract.
type GmxMigratorClearActionIterator struct {
	Event *GmxMigratorClearAction // Event containing the contract specifics and raw log

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
func (it *GmxMigratorClearActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxMigratorClearAction)
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
		it.Event = new(GmxMigratorClearAction)
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
func (it *GmxMigratorClearActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxMigratorClearActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxMigratorClearAction represents a ClearAction event raised by the GmxMigrator contract.
type GmxMigratorClearAction struct {
	Action [32]byte
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClearAction is a free log retrieval operation binding the contract event 0xf4640d39061e643d9b802cb3725953405344555ad6dbb1cbdb0495f3eccb8e68.
//
// Solidity: event ClearAction(bytes32 action, uint256 nonce)
func (_GmxMigrator *GmxMigratorFilterer) FilterClearAction(opts *bind.FilterOpts) (*GmxMigratorClearActionIterator, error) {

	logs, sub, err := _GmxMigrator.contract.FilterLogs(opts, "ClearAction")
	if err != nil {
		return nil, err
	}
	return &GmxMigratorClearActionIterator{contract: _GmxMigrator.contract, event: "ClearAction", logs: logs, sub: sub}, nil
}

// WatchClearAction is a free log subscription operation binding the contract event 0xf4640d39061e643d9b802cb3725953405344555ad6dbb1cbdb0495f3eccb8e68.
//
// Solidity: event ClearAction(bytes32 action, uint256 nonce)
func (_GmxMigrator *GmxMigratorFilterer) WatchClearAction(opts *bind.WatchOpts, sink chan<- *GmxMigratorClearAction) (event.Subscription, error) {

	logs, sub, err := _GmxMigrator.contract.WatchLogs(opts, "ClearAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxMigratorClearAction)
				if err := _GmxMigrator.contract.UnpackLog(event, "ClearAction", log); err != nil {
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

// ParseClearAction is a log parse operation binding the contract event 0xf4640d39061e643d9b802cb3725953405344555ad6dbb1cbdb0495f3eccb8e68.
//
// Solidity: event ClearAction(bytes32 action, uint256 nonce)
func (_GmxMigrator *GmxMigratorFilterer) ParseClearAction(log types.Log) (*GmxMigratorClearAction, error) {
	event := new(GmxMigratorClearAction)
	if err := _GmxMigrator.contract.UnpackLog(event, "ClearAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxMigratorSignActionIterator is returned from FilterSignAction and is used to iterate over the raw logs and unpacked data for SignAction events raised by the GmxMigrator contract.
type GmxMigratorSignActionIterator struct {
	Event *GmxMigratorSignAction // Event containing the contract specifics and raw log

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
func (it *GmxMigratorSignActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxMigratorSignAction)
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
		it.Event = new(GmxMigratorSignAction)
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
func (it *GmxMigratorSignActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxMigratorSignActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxMigratorSignAction represents a SignAction event raised by the GmxMigrator contract.
type GmxMigratorSignAction struct {
	Action [32]byte
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignAction is a free log retrieval operation binding the contract event 0xaae28fe5531fe5dfb8d12409392ec67b50c825dd06233312cb6aeaddd16cbd22.
//
// Solidity: event SignAction(bytes32 action, uint256 nonce)
func (_GmxMigrator *GmxMigratorFilterer) FilterSignAction(opts *bind.FilterOpts) (*GmxMigratorSignActionIterator, error) {

	logs, sub, err := _GmxMigrator.contract.FilterLogs(opts, "SignAction")
	if err != nil {
		return nil, err
	}
	return &GmxMigratorSignActionIterator{contract: _GmxMigrator.contract, event: "SignAction", logs: logs, sub: sub}, nil
}

// WatchSignAction is a free log subscription operation binding the contract event 0xaae28fe5531fe5dfb8d12409392ec67b50c825dd06233312cb6aeaddd16cbd22.
//
// Solidity: event SignAction(bytes32 action, uint256 nonce)
func (_GmxMigrator *GmxMigratorFilterer) WatchSignAction(opts *bind.WatchOpts, sink chan<- *GmxMigratorSignAction) (event.Subscription, error) {

	logs, sub, err := _GmxMigrator.contract.WatchLogs(opts, "SignAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxMigratorSignAction)
				if err := _GmxMigrator.contract.UnpackLog(event, "SignAction", log); err != nil {
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

// ParseSignAction is a log parse operation binding the contract event 0xaae28fe5531fe5dfb8d12409392ec67b50c825dd06233312cb6aeaddd16cbd22.
//
// Solidity: event SignAction(bytes32 action, uint256 nonce)
func (_GmxMigrator *GmxMigratorFilterer) ParseSignAction(log types.Log) (*GmxMigratorSignAction, error) {
	event := new(GmxMigratorSignAction)
	if err := _GmxMigrator.contract.UnpackLog(event, "SignAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxMigratorSignalApproveIterator is returned from FilterSignalApprove and is used to iterate over the raw logs and unpacked data for SignalApprove events raised by the GmxMigrator contract.
type GmxMigratorSignalApproveIterator struct {
	Event *GmxMigratorSignalApprove // Event containing the contract specifics and raw log

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
func (it *GmxMigratorSignalApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxMigratorSignalApprove)
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
		it.Event = new(GmxMigratorSignalApprove)
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
func (it *GmxMigratorSignalApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxMigratorSignalApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxMigratorSignalApprove represents a SignalApprove event raised by the GmxMigrator contract.
type GmxMigratorSignalApprove struct {
	Token   common.Address
	Spender common.Address
	Amount  *big.Int
	Action  [32]byte
	Nonce   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignalApprove is a free log retrieval operation binding the contract event 0xc19251bf5f704ddc3d5babe6f4e5bde0dded20b19f7844716861821ab3163cd7.
//
// Solidity: event SignalApprove(address token, address spender, uint256 amount, bytes32 action, uint256 nonce)
func (_GmxMigrator *GmxMigratorFilterer) FilterSignalApprove(opts *bind.FilterOpts) (*GmxMigratorSignalApproveIterator, error) {

	logs, sub, err := _GmxMigrator.contract.FilterLogs(opts, "SignalApprove")
	if err != nil {
		return nil, err
	}
	return &GmxMigratorSignalApproveIterator{contract: _GmxMigrator.contract, event: "SignalApprove", logs: logs, sub: sub}, nil
}

// WatchSignalApprove is a free log subscription operation binding the contract event 0xc19251bf5f704ddc3d5babe6f4e5bde0dded20b19f7844716861821ab3163cd7.
//
// Solidity: event SignalApprove(address token, address spender, uint256 amount, bytes32 action, uint256 nonce)
func (_GmxMigrator *GmxMigratorFilterer) WatchSignalApprove(opts *bind.WatchOpts, sink chan<- *GmxMigratorSignalApprove) (event.Subscription, error) {

	logs, sub, err := _GmxMigrator.contract.WatchLogs(opts, "SignalApprove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxMigratorSignalApprove)
				if err := _GmxMigrator.contract.UnpackLog(event, "SignalApprove", log); err != nil {
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

// ParseSignalApprove is a log parse operation binding the contract event 0xc19251bf5f704ddc3d5babe6f4e5bde0dded20b19f7844716861821ab3163cd7.
//
// Solidity: event SignalApprove(address token, address spender, uint256 amount, bytes32 action, uint256 nonce)
func (_GmxMigrator *GmxMigratorFilterer) ParseSignalApprove(log types.Log) (*GmxMigratorSignalApprove, error) {
	event := new(GmxMigratorSignalApprove)
	if err := _GmxMigrator.contract.UnpackLog(event, "SignalApprove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GmxMigratorSignalPendingActionIterator is returned from FilterSignalPendingAction and is used to iterate over the raw logs and unpacked data for SignalPendingAction events raised by the GmxMigrator contract.
type GmxMigratorSignalPendingActionIterator struct {
	Event *GmxMigratorSignalPendingAction // Event containing the contract specifics and raw log

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
func (it *GmxMigratorSignalPendingActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GmxMigratorSignalPendingAction)
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
		it.Event = new(GmxMigratorSignalPendingAction)
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
func (it *GmxMigratorSignalPendingActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GmxMigratorSignalPendingActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GmxMigratorSignalPendingAction represents a SignalPendingAction event raised by the GmxMigrator contract.
type GmxMigratorSignalPendingAction struct {
	Action [32]byte
	Nonce  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignalPendingAction is a free log retrieval operation binding the contract event 0x64df01c46eb530dc540770a0b88cc32f0b8c2b371a546ae0b13cc8ca6671fff9.
//
// Solidity: event SignalPendingAction(bytes32 action, uint256 nonce)
func (_GmxMigrator *GmxMigratorFilterer) FilterSignalPendingAction(opts *bind.FilterOpts) (*GmxMigratorSignalPendingActionIterator, error) {

	logs, sub, err := _GmxMigrator.contract.FilterLogs(opts, "SignalPendingAction")
	if err != nil {
		return nil, err
	}
	return &GmxMigratorSignalPendingActionIterator{contract: _GmxMigrator.contract, event: "SignalPendingAction", logs: logs, sub: sub}, nil
}

// WatchSignalPendingAction is a free log subscription operation binding the contract event 0x64df01c46eb530dc540770a0b88cc32f0b8c2b371a546ae0b13cc8ca6671fff9.
//
// Solidity: event SignalPendingAction(bytes32 action, uint256 nonce)
func (_GmxMigrator *GmxMigratorFilterer) WatchSignalPendingAction(opts *bind.WatchOpts, sink chan<- *GmxMigratorSignalPendingAction) (event.Subscription, error) {

	logs, sub, err := _GmxMigrator.contract.WatchLogs(opts, "SignalPendingAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GmxMigratorSignalPendingAction)
				if err := _GmxMigrator.contract.UnpackLog(event, "SignalPendingAction", log); err != nil {
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

// ParseSignalPendingAction is a log parse operation binding the contract event 0x64df01c46eb530dc540770a0b88cc32f0b8c2b371a546ae0b13cc8ca6671fff9.
//
// Solidity: event SignalPendingAction(bytes32 action, uint256 nonce)
func (_GmxMigrator *GmxMigratorFilterer) ParseSignalPendingAction(log types.Log) (*GmxMigratorSignalPendingAction, error) {
	event := new(GmxMigratorSignalPendingAction)
	if err := _GmxMigrator.contract.UnpackLog(event, "SignalPendingAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

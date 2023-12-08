// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ReaderV2

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

// ReaderV2MetaData contains all meta data concerning the ReaderV2 contract.
var ReaderV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"BASIS_POINTS_DIVISOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"POSITION_PROPS_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PRICE_PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"USDG_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAmountOut\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"contractIVault\"},{\"name\":\"_tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeBasisPoints\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"contractIVault\"},{\"name\":\"_tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFees\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFullVaultTokenInfo\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdgAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFundingRates\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxAmountIn\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"contractIVault\"},{\"name\":\"_tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenOut\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPairInfo\",\"inputs\":[{\"name\":\"_factory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPositions\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_indexTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_isLong\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPrices\",\"inputs\":[{\"name\":\"_priceFeed\",\"type\":\"address\",\"internalType\":\"contractIVaultPriceFeed\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakingInfo\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_yieldTrackers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenBalances\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenBalancesWithSupplies\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenSupply\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_excludedAccounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalBalance\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalStaked\",\"inputs\":[{\"name\":\"_yieldTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVaultTokenInfo\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdgAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVaultTokenInfoV2\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdgAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVestingInfo\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_vesters\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasMaxGlobalShortSizes\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setConfig\",\"inputs\":[{\"name\":\"_hasMaxGlobalShortSizes\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// ReaderV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use ReaderV2MetaData.ABI instead.
var ReaderV2ABI = ReaderV2MetaData.ABI

// ReaderV2 is an auto generated Go binding around an Ethereum contract.
type ReaderV2 struct {
	ReaderV2Caller     // Read-only binding to the contract
	ReaderV2Transactor // Write-only binding to the contract
	ReaderV2Filterer   // Log filterer for contract events
}

// ReaderV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type ReaderV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReaderV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ReaderV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReaderV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReaderV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReaderV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReaderV2Session struct {
	Contract     *ReaderV2         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReaderV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReaderV2CallerSession struct {
	Contract *ReaderV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ReaderV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReaderV2TransactorSession struct {
	Contract     *ReaderV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ReaderV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type ReaderV2Raw struct {
	Contract *ReaderV2 // Generic contract binding to access the raw methods on
}

// ReaderV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReaderV2CallerRaw struct {
	Contract *ReaderV2Caller // Generic read-only contract binding to access the raw methods on
}

// ReaderV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReaderV2TransactorRaw struct {
	Contract *ReaderV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewReaderV2 creates a new instance of ReaderV2, bound to a specific deployed contract.
func NewReaderV2(address common.Address, backend bind.ContractBackend) (*ReaderV2, error) {
	contract, err := bindReaderV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReaderV2{ReaderV2Caller: ReaderV2Caller{contract: contract}, ReaderV2Transactor: ReaderV2Transactor{contract: contract}, ReaderV2Filterer: ReaderV2Filterer{contract: contract}}, nil
}

// NewReaderV2Caller creates a new read-only instance of ReaderV2, bound to a specific deployed contract.
func NewReaderV2Caller(address common.Address, caller bind.ContractCaller) (*ReaderV2Caller, error) {
	contract, err := bindReaderV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReaderV2Caller{contract: contract}, nil
}

// NewReaderV2Transactor creates a new write-only instance of ReaderV2, bound to a specific deployed contract.
func NewReaderV2Transactor(address common.Address, transactor bind.ContractTransactor) (*ReaderV2Transactor, error) {
	contract, err := bindReaderV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReaderV2Transactor{contract: contract}, nil
}

// NewReaderV2Filterer creates a new log filterer instance of ReaderV2, bound to a specific deployed contract.
func NewReaderV2Filterer(address common.Address, filterer bind.ContractFilterer) (*ReaderV2Filterer, error) {
	contract, err := bindReaderV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReaderV2Filterer{contract: contract}, nil
}

// bindReaderV2 binds a generic wrapper to an already deployed contract.
func bindReaderV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReaderV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReaderV2 *ReaderV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReaderV2.Contract.ReaderV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReaderV2 *ReaderV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReaderV2.Contract.ReaderV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReaderV2 *ReaderV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReaderV2.Contract.ReaderV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReaderV2 *ReaderV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReaderV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReaderV2 *ReaderV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReaderV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReaderV2 *ReaderV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReaderV2.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_ReaderV2 *ReaderV2Caller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_ReaderV2 *ReaderV2Session) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _ReaderV2.Contract.BASISPOINTSDIVISOR(&_ReaderV2.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_ReaderV2 *ReaderV2CallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _ReaderV2.Contract.BASISPOINTSDIVISOR(&_ReaderV2.CallOpts)
}

// POSITIONPROPSLENGTH is a free data retrieval call binding the contract method 0xad7e5497.
//
// Solidity: function POSITION_PROPS_LENGTH() view returns(uint256)
func (_ReaderV2 *ReaderV2Caller) POSITIONPROPSLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "POSITION_PROPS_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// POSITIONPROPSLENGTH is a free data retrieval call binding the contract method 0xad7e5497.
//
// Solidity: function POSITION_PROPS_LENGTH() view returns(uint256)
func (_ReaderV2 *ReaderV2Session) POSITIONPROPSLENGTH() (*big.Int, error) {
	return _ReaderV2.Contract.POSITIONPROPSLENGTH(&_ReaderV2.CallOpts)
}

// POSITIONPROPSLENGTH is a free data retrieval call binding the contract method 0xad7e5497.
//
// Solidity: function POSITION_PROPS_LENGTH() view returns(uint256)
func (_ReaderV2 *ReaderV2CallerSession) POSITIONPROPSLENGTH() (*big.Int, error) {
	return _ReaderV2.Contract.POSITIONPROPSLENGTH(&_ReaderV2.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_ReaderV2 *ReaderV2Caller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_ReaderV2 *ReaderV2Session) PRICEPRECISION() (*big.Int, error) {
	return _ReaderV2.Contract.PRICEPRECISION(&_ReaderV2.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_ReaderV2 *ReaderV2CallerSession) PRICEPRECISION() (*big.Int, error) {
	return _ReaderV2.Contract.PRICEPRECISION(&_ReaderV2.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_ReaderV2 *ReaderV2Caller) USDGDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "USDG_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_ReaderV2 *ReaderV2Session) USDGDECIMALS() (*big.Int, error) {
	return _ReaderV2.Contract.USDGDECIMALS(&_ReaderV2.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_ReaderV2 *ReaderV2CallerSession) USDGDECIMALS() (*big.Int, error) {
	return _ReaderV2.Contract.USDGDECIMALS(&_ReaderV2.CallOpts)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xd7176ca9.
//
// Solidity: function getAmountOut(address _vault, address _tokenIn, address _tokenOut, uint256 _amountIn) view returns(uint256, uint256)
func (_ReaderV2 *ReaderV2Caller) GetAmountOut(opts *bind.CallOpts, _vault common.Address, _tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "getAmountOut", _vault, _tokenIn, _tokenOut, _amountIn)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0xd7176ca9.
//
// Solidity: function getAmountOut(address _vault, address _tokenIn, address _tokenOut, uint256 _amountIn) view returns(uint256, uint256)
func (_ReaderV2 *ReaderV2Session) GetAmountOut(_vault common.Address, _tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int) (*big.Int, *big.Int, error) {
	return _ReaderV2.Contract.GetAmountOut(&_ReaderV2.CallOpts, _vault, _tokenIn, _tokenOut, _amountIn)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xd7176ca9.
//
// Solidity: function getAmountOut(address _vault, address _tokenIn, address _tokenOut, uint256 _amountIn) view returns(uint256, uint256)
func (_ReaderV2 *ReaderV2CallerSession) GetAmountOut(_vault common.Address, _tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int) (*big.Int, *big.Int, error) {
	return _ReaderV2.Contract.GetAmountOut(&_ReaderV2.CallOpts, _vault, _tokenIn, _tokenOut, _amountIn)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0x440f018c.
//
// Solidity: function getFeeBasisPoints(address _vault, address _tokenIn, address _tokenOut, uint256 _amountIn) view returns(uint256, uint256, uint256)
func (_ReaderV2 *ReaderV2Caller) GetFeeBasisPoints(opts *bind.CallOpts, _vault common.Address, _tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "getFeeBasisPoints", _vault, _tokenIn, _tokenOut, _amountIn)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0x440f018c.
//
// Solidity: function getFeeBasisPoints(address _vault, address _tokenIn, address _tokenOut, uint256 _amountIn) view returns(uint256, uint256, uint256)
func (_ReaderV2 *ReaderV2Session) GetFeeBasisPoints(_vault common.Address, _tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _ReaderV2.Contract.GetFeeBasisPoints(&_ReaderV2.CallOpts, _vault, _tokenIn, _tokenOut, _amountIn)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0x440f018c.
//
// Solidity: function getFeeBasisPoints(address _vault, address _tokenIn, address _tokenOut, uint256 _amountIn) view returns(uint256, uint256, uint256)
func (_ReaderV2 *ReaderV2CallerSession) GetFeeBasisPoints(_vault common.Address, _tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _ReaderV2.Contract.GetFeeBasisPoints(&_ReaderV2.CallOpts, _vault, _tokenIn, _tokenOut, _amountIn)
}

// GetFees is a free data retrieval call binding the contract method 0x86d4d0f5.
//
// Solidity: function getFees(address _vault, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Caller) GetFees(opts *bind.CallOpts, _vault common.Address, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "getFees", _vault, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetFees is a free data retrieval call binding the contract method 0x86d4d0f5.
//
// Solidity: function getFees(address _vault, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Session) GetFees(_vault common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetFees(&_ReaderV2.CallOpts, _vault, _tokens)
}

// GetFees is a free data retrieval call binding the contract method 0x86d4d0f5.
//
// Solidity: function getFees(address _vault, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2CallerSession) GetFees(_vault common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetFees(&_ReaderV2.CallOpts, _vault, _tokens)
}

// GetFullVaultTokenInfo is a free data retrieval call binding the contract method 0x7b906e93.
//
// Solidity: function getFullVaultTokenInfo(address _vault, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Caller) GetFullVaultTokenInfo(opts *bind.CallOpts, _vault common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "getFullVaultTokenInfo", _vault, _weth, _usdgAmount, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetFullVaultTokenInfo is a free data retrieval call binding the contract method 0x7b906e93.
//
// Solidity: function getFullVaultTokenInfo(address _vault, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Session) GetFullVaultTokenInfo(_vault common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetFullVaultTokenInfo(&_ReaderV2.CallOpts, _vault, _weth, _usdgAmount, _tokens)
}

// GetFullVaultTokenInfo is a free data retrieval call binding the contract method 0x7b906e93.
//
// Solidity: function getFullVaultTokenInfo(address _vault, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2CallerSession) GetFullVaultTokenInfo(_vault common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetFullVaultTokenInfo(&_ReaderV2.CallOpts, _vault, _weth, _usdgAmount, _tokens)
}

// GetFundingRates is a free data retrieval call binding the contract method 0x95a7535a.
//
// Solidity: function getFundingRates(address _vault, address _weth, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Caller) GetFundingRates(opts *bind.CallOpts, _vault common.Address, _weth common.Address, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "getFundingRates", _vault, _weth, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetFundingRates is a free data retrieval call binding the contract method 0x95a7535a.
//
// Solidity: function getFundingRates(address _vault, address _weth, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Session) GetFundingRates(_vault common.Address, _weth common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetFundingRates(&_ReaderV2.CallOpts, _vault, _weth, _tokens)
}

// GetFundingRates is a free data retrieval call binding the contract method 0x95a7535a.
//
// Solidity: function getFundingRates(address _vault, address _weth, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2CallerSession) GetFundingRates(_vault common.Address, _weth common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetFundingRates(&_ReaderV2.CallOpts, _vault, _weth, _tokens)
}

// GetMaxAmountIn is a free data retrieval call binding the contract method 0xf3535e6c.
//
// Solidity: function getMaxAmountIn(address _vault, address _tokenIn, address _tokenOut) view returns(uint256)
func (_ReaderV2 *ReaderV2Caller) GetMaxAmountIn(opts *bind.CallOpts, _vault common.Address, _tokenIn common.Address, _tokenOut common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "getMaxAmountIn", _vault, _tokenIn, _tokenOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxAmountIn is a free data retrieval call binding the contract method 0xf3535e6c.
//
// Solidity: function getMaxAmountIn(address _vault, address _tokenIn, address _tokenOut) view returns(uint256)
func (_ReaderV2 *ReaderV2Session) GetMaxAmountIn(_vault common.Address, _tokenIn common.Address, _tokenOut common.Address) (*big.Int, error) {
	return _ReaderV2.Contract.GetMaxAmountIn(&_ReaderV2.CallOpts, _vault, _tokenIn, _tokenOut)
}

// GetMaxAmountIn is a free data retrieval call binding the contract method 0xf3535e6c.
//
// Solidity: function getMaxAmountIn(address _vault, address _tokenIn, address _tokenOut) view returns(uint256)
func (_ReaderV2 *ReaderV2CallerSession) GetMaxAmountIn(_vault common.Address, _tokenIn common.Address, _tokenOut common.Address) (*big.Int, error) {
	return _ReaderV2.Contract.GetMaxAmountIn(&_ReaderV2.CallOpts, _vault, _tokenIn, _tokenOut)
}

// GetPairInfo is a free data retrieval call binding the contract method 0xa4543ead.
//
// Solidity: function getPairInfo(address _factory, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Caller) GetPairInfo(opts *bind.CallOpts, _factory common.Address, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "getPairInfo", _factory, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPairInfo is a free data retrieval call binding the contract method 0xa4543ead.
//
// Solidity: function getPairInfo(address _factory, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Session) GetPairInfo(_factory common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetPairInfo(&_ReaderV2.CallOpts, _factory, _tokens)
}

// GetPairInfo is a free data retrieval call binding the contract method 0xa4543ead.
//
// Solidity: function getPairInfo(address _factory, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2CallerSession) GetPairInfo(_factory common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetPairInfo(&_ReaderV2.CallOpts, _factory, _tokens)
}

// GetPositions is a free data retrieval call binding the contract method 0xdc383cab.
//
// Solidity: function getPositions(address _vault, address _account, address[] _collateralTokens, address[] _indexTokens, bool[] _isLong) view returns(uint256[])
func (_ReaderV2 *ReaderV2Caller) GetPositions(opts *bind.CallOpts, _vault common.Address, _account common.Address, _collateralTokens []common.Address, _indexTokens []common.Address, _isLong []bool) ([]*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "getPositions", _vault, _account, _collateralTokens, _indexTokens, _isLong)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPositions is a free data retrieval call binding the contract method 0xdc383cab.
//
// Solidity: function getPositions(address _vault, address _account, address[] _collateralTokens, address[] _indexTokens, bool[] _isLong) view returns(uint256[])
func (_ReaderV2 *ReaderV2Session) GetPositions(_vault common.Address, _account common.Address, _collateralTokens []common.Address, _indexTokens []common.Address, _isLong []bool) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetPositions(&_ReaderV2.CallOpts, _vault, _account, _collateralTokens, _indexTokens, _isLong)
}

// GetPositions is a free data retrieval call binding the contract method 0xdc383cab.
//
// Solidity: function getPositions(address _vault, address _account, address[] _collateralTokens, address[] _indexTokens, bool[] _isLong) view returns(uint256[])
func (_ReaderV2 *ReaderV2CallerSession) GetPositions(_vault common.Address, _account common.Address, _collateralTokens []common.Address, _indexTokens []common.Address, _isLong []bool) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetPositions(&_ReaderV2.CallOpts, _vault, _account, _collateralTokens, _indexTokens, _isLong)
}

// GetPrices is a free data retrieval call binding the contract method 0x3613d527.
//
// Solidity: function getPrices(address _priceFeed, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Caller) GetPrices(opts *bind.CallOpts, _priceFeed common.Address, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "getPrices", _priceFeed, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPrices is a free data retrieval call binding the contract method 0x3613d527.
//
// Solidity: function getPrices(address _priceFeed, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Session) GetPrices(_priceFeed common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetPrices(&_ReaderV2.CallOpts, _priceFeed, _tokens)
}

// GetPrices is a free data retrieval call binding the contract method 0x3613d527.
//
// Solidity: function getPrices(address _priceFeed, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2CallerSession) GetPrices(_priceFeed common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetPrices(&_ReaderV2.CallOpts, _priceFeed, _tokens)
}

// GetStakingInfo is a free data retrieval call binding the contract method 0x937a0be8.
//
// Solidity: function getStakingInfo(address _account, address[] _yieldTrackers) view returns(uint256[])
func (_ReaderV2 *ReaderV2Caller) GetStakingInfo(opts *bind.CallOpts, _account common.Address, _yieldTrackers []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "getStakingInfo", _account, _yieldTrackers)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetStakingInfo is a free data retrieval call binding the contract method 0x937a0be8.
//
// Solidity: function getStakingInfo(address _account, address[] _yieldTrackers) view returns(uint256[])
func (_ReaderV2 *ReaderV2Session) GetStakingInfo(_account common.Address, _yieldTrackers []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetStakingInfo(&_ReaderV2.CallOpts, _account, _yieldTrackers)
}

// GetStakingInfo is a free data retrieval call binding the contract method 0x937a0be8.
//
// Solidity: function getStakingInfo(address _account, address[] _yieldTrackers) view returns(uint256[])
func (_ReaderV2 *ReaderV2CallerSession) GetStakingInfo(_account common.Address, _yieldTrackers []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetStakingInfo(&_ReaderV2.CallOpts, _account, _yieldTrackers)
}

// GetTokenBalances is a free data retrieval call binding the contract method 0xd802178e.
//
// Solidity: function getTokenBalances(address _account, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Caller) GetTokenBalances(opts *bind.CallOpts, _account common.Address, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "getTokenBalances", _account, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTokenBalances is a free data retrieval call binding the contract method 0xd802178e.
//
// Solidity: function getTokenBalances(address _account, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Session) GetTokenBalances(_account common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetTokenBalances(&_ReaderV2.CallOpts, _account, _tokens)
}

// GetTokenBalances is a free data retrieval call binding the contract method 0xd802178e.
//
// Solidity: function getTokenBalances(address _account, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2CallerSession) GetTokenBalances(_account common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetTokenBalances(&_ReaderV2.CallOpts, _account, _tokens)
}

// GetTokenBalancesWithSupplies is a free data retrieval call binding the contract method 0x2e3e3342.
//
// Solidity: function getTokenBalancesWithSupplies(address _account, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Caller) GetTokenBalancesWithSupplies(opts *bind.CallOpts, _account common.Address, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "getTokenBalancesWithSupplies", _account, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTokenBalancesWithSupplies is a free data retrieval call binding the contract method 0x2e3e3342.
//
// Solidity: function getTokenBalancesWithSupplies(address _account, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Session) GetTokenBalancesWithSupplies(_account common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetTokenBalancesWithSupplies(&_ReaderV2.CallOpts, _account, _tokens)
}

// GetTokenBalancesWithSupplies is a free data retrieval call binding the contract method 0x2e3e3342.
//
// Solidity: function getTokenBalancesWithSupplies(address _account, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2CallerSession) GetTokenBalancesWithSupplies(_account common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetTokenBalancesWithSupplies(&_ReaderV2.CallOpts, _account, _tokens)
}

// GetTokenSupply is a free data retrieval call binding the contract method 0x2ac0184c.
//
// Solidity: function getTokenSupply(address _token, address[] _excludedAccounts) view returns(uint256)
func (_ReaderV2 *ReaderV2Caller) GetTokenSupply(opts *bind.CallOpts, _token common.Address, _excludedAccounts []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "getTokenSupply", _token, _excludedAccounts)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenSupply is a free data retrieval call binding the contract method 0x2ac0184c.
//
// Solidity: function getTokenSupply(address _token, address[] _excludedAccounts) view returns(uint256)
func (_ReaderV2 *ReaderV2Session) GetTokenSupply(_token common.Address, _excludedAccounts []common.Address) (*big.Int, error) {
	return _ReaderV2.Contract.GetTokenSupply(&_ReaderV2.CallOpts, _token, _excludedAccounts)
}

// GetTokenSupply is a free data retrieval call binding the contract method 0x2ac0184c.
//
// Solidity: function getTokenSupply(address _token, address[] _excludedAccounts) view returns(uint256)
func (_ReaderV2 *ReaderV2CallerSession) GetTokenSupply(_token common.Address, _excludedAccounts []common.Address) (*big.Int, error) {
	return _ReaderV2.Contract.GetTokenSupply(&_ReaderV2.CallOpts, _token, _excludedAccounts)
}

// GetTotalBalance is a free data retrieval call binding the contract method 0xfbdb05ca.
//
// Solidity: function getTotalBalance(address _token, address[] _accounts) view returns(uint256)
func (_ReaderV2 *ReaderV2Caller) GetTotalBalance(opts *bind.CallOpts, _token common.Address, _accounts []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "getTotalBalance", _token, _accounts)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBalance is a free data retrieval call binding the contract method 0xfbdb05ca.
//
// Solidity: function getTotalBalance(address _token, address[] _accounts) view returns(uint256)
func (_ReaderV2 *ReaderV2Session) GetTotalBalance(_token common.Address, _accounts []common.Address) (*big.Int, error) {
	return _ReaderV2.Contract.GetTotalBalance(&_ReaderV2.CallOpts, _token, _accounts)
}

// GetTotalBalance is a free data retrieval call binding the contract method 0xfbdb05ca.
//
// Solidity: function getTotalBalance(address _token, address[] _accounts) view returns(uint256)
func (_ReaderV2 *ReaderV2CallerSession) GetTotalBalance(_token common.Address, _accounts []common.Address) (*big.Int, error) {
	return _ReaderV2.Contract.GetTotalBalance(&_ReaderV2.CallOpts, _token, _accounts)
}

// GetTotalStaked is a free data retrieval call binding the contract method 0x2413c8c1.
//
// Solidity: function getTotalStaked(address[] _yieldTokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Caller) GetTotalStaked(opts *bind.CallOpts, _yieldTokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "getTotalStaked", _yieldTokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTotalStaked is a free data retrieval call binding the contract method 0x2413c8c1.
//
// Solidity: function getTotalStaked(address[] _yieldTokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Session) GetTotalStaked(_yieldTokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetTotalStaked(&_ReaderV2.CallOpts, _yieldTokens)
}

// GetTotalStaked is a free data retrieval call binding the contract method 0x2413c8c1.
//
// Solidity: function getTotalStaked(address[] _yieldTokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2CallerSession) GetTotalStaked(_yieldTokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetTotalStaked(&_ReaderV2.CallOpts, _yieldTokens)
}

// GetVaultTokenInfo is a free data retrieval call binding the contract method 0x20542568.
//
// Solidity: function getVaultTokenInfo(address _vault, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Caller) GetVaultTokenInfo(opts *bind.CallOpts, _vault common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "getVaultTokenInfo", _vault, _weth, _usdgAmount, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetVaultTokenInfo is a free data retrieval call binding the contract method 0x20542568.
//
// Solidity: function getVaultTokenInfo(address _vault, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Session) GetVaultTokenInfo(_vault common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetVaultTokenInfo(&_ReaderV2.CallOpts, _vault, _weth, _usdgAmount, _tokens)
}

// GetVaultTokenInfo is a free data retrieval call binding the contract method 0x20542568.
//
// Solidity: function getVaultTokenInfo(address _vault, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2CallerSession) GetVaultTokenInfo(_vault common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetVaultTokenInfo(&_ReaderV2.CallOpts, _vault, _weth, _usdgAmount, _tokens)
}

// GetVaultTokenInfoV2 is a free data retrieval call binding the contract method 0x8e83ca32.
//
// Solidity: function getVaultTokenInfoV2(address _vault, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Caller) GetVaultTokenInfoV2(opts *bind.CallOpts, _vault common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "getVaultTokenInfoV2", _vault, _weth, _usdgAmount, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetVaultTokenInfoV2 is a free data retrieval call binding the contract method 0x8e83ca32.
//
// Solidity: function getVaultTokenInfoV2(address _vault, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2Session) GetVaultTokenInfoV2(_vault common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetVaultTokenInfoV2(&_ReaderV2.CallOpts, _vault, _weth, _usdgAmount, _tokens)
}

// GetVaultTokenInfoV2 is a free data retrieval call binding the contract method 0x8e83ca32.
//
// Solidity: function getVaultTokenInfoV2(address _vault, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_ReaderV2 *ReaderV2CallerSession) GetVaultTokenInfoV2(_vault common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetVaultTokenInfoV2(&_ReaderV2.CallOpts, _vault, _weth, _usdgAmount, _tokens)
}

// GetVestingInfo is a free data retrieval call binding the contract method 0x48211934.
//
// Solidity: function getVestingInfo(address _account, address[] _vesters) view returns(uint256[])
func (_ReaderV2 *ReaderV2Caller) GetVestingInfo(opts *bind.CallOpts, _account common.Address, _vesters []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "getVestingInfo", _account, _vesters)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetVestingInfo is a free data retrieval call binding the contract method 0x48211934.
//
// Solidity: function getVestingInfo(address _account, address[] _vesters) view returns(uint256[])
func (_ReaderV2 *ReaderV2Session) GetVestingInfo(_account common.Address, _vesters []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetVestingInfo(&_ReaderV2.CallOpts, _account, _vesters)
}

// GetVestingInfo is a free data retrieval call binding the contract method 0x48211934.
//
// Solidity: function getVestingInfo(address _account, address[] _vesters) view returns(uint256[])
func (_ReaderV2 *ReaderV2CallerSession) GetVestingInfo(_account common.Address, _vesters []common.Address) ([]*big.Int, error) {
	return _ReaderV2.Contract.GetVestingInfo(&_ReaderV2.CallOpts, _account, _vesters)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_ReaderV2 *ReaderV2Caller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_ReaderV2 *ReaderV2Session) Gov() (common.Address, error) {
	return _ReaderV2.Contract.Gov(&_ReaderV2.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_ReaderV2 *ReaderV2CallerSession) Gov() (common.Address, error) {
	return _ReaderV2.Contract.Gov(&_ReaderV2.CallOpts)
}

// HasMaxGlobalShortSizes is a free data retrieval call binding the contract method 0xc6f1d676.
//
// Solidity: function hasMaxGlobalShortSizes() view returns(bool)
func (_ReaderV2 *ReaderV2Caller) HasMaxGlobalShortSizes(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ReaderV2.contract.Call(opts, &out, "hasMaxGlobalShortSizes")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMaxGlobalShortSizes is a free data retrieval call binding the contract method 0xc6f1d676.
//
// Solidity: function hasMaxGlobalShortSizes() view returns(bool)
func (_ReaderV2 *ReaderV2Session) HasMaxGlobalShortSizes() (bool, error) {
	return _ReaderV2.Contract.HasMaxGlobalShortSizes(&_ReaderV2.CallOpts)
}

// HasMaxGlobalShortSizes is a free data retrieval call binding the contract method 0xc6f1d676.
//
// Solidity: function hasMaxGlobalShortSizes() view returns(bool)
func (_ReaderV2 *ReaderV2CallerSession) HasMaxGlobalShortSizes() (bool, error) {
	return _ReaderV2.Contract.HasMaxGlobalShortSizes(&_ReaderV2.CallOpts)
}

// SetConfig is a paid mutator transaction binding the contract method 0x9b0183c3.
//
// Solidity: function setConfig(bool _hasMaxGlobalShortSizes) returns()
func (_ReaderV2 *ReaderV2Transactor) SetConfig(opts *bind.TransactOpts, _hasMaxGlobalShortSizes bool) (*types.Transaction, error) {
	return _ReaderV2.contract.Transact(opts, "setConfig", _hasMaxGlobalShortSizes)
}

// SetConfig is a paid mutator transaction binding the contract method 0x9b0183c3.
//
// Solidity: function setConfig(bool _hasMaxGlobalShortSizes) returns()
func (_ReaderV2 *ReaderV2Session) SetConfig(_hasMaxGlobalShortSizes bool) (*types.Transaction, error) {
	return _ReaderV2.Contract.SetConfig(&_ReaderV2.TransactOpts, _hasMaxGlobalShortSizes)
}

// SetConfig is a paid mutator transaction binding the contract method 0x9b0183c3.
//
// Solidity: function setConfig(bool _hasMaxGlobalShortSizes) returns()
func (_ReaderV2 *ReaderV2TransactorSession) SetConfig(_hasMaxGlobalShortSizes bool) (*types.Transaction, error) {
	return _ReaderV2.Contract.SetConfig(&_ReaderV2.TransactOpts, _hasMaxGlobalShortSizes)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_ReaderV2 *ReaderV2Transactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _ReaderV2.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_ReaderV2 *ReaderV2Session) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _ReaderV2.Contract.SetGov(&_ReaderV2.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_ReaderV2 *ReaderV2TransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _ReaderV2.Contract.SetGov(&_ReaderV2.TransactOpts, _gov)
}

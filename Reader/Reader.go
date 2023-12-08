// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Reader

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

// ReaderMetaData contains all meta data concerning the Reader contract.
var ReaderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"BASIS_POINTS_DIVISOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"POSITION_PROPS_LENGTH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PRICE_PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"USDG_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAmountOut\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"contractIVault\"},{\"name\":\"_tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeBasisPoints\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"contractIVault\"},{\"name\":\"_tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFees\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFullVaultTokenInfo\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdgAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFundingRates\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxAmountIn\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"contractIVault\"},{\"name\":\"_tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenOut\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPairInfo\",\"inputs\":[{\"name\":\"_factory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPositions\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_indexTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_isLong\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPrices\",\"inputs\":[{\"name\":\"_priceFeed\",\"type\":\"address\",\"internalType\":\"contractIVaultPriceFeed\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakingInfo\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_yieldTrackers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenBalances\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenBalancesWithSupplies\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenSupply\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_excludedAccounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalBalance\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"contractIERC20\"},{\"name\":\"_accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTotalStaked\",\"inputs\":[{\"name\":\"_yieldTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVaultTokenInfo\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdgAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVaultTokenInfoV2\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdgAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVestingInfo\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_vesters\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasMaxGlobalShortSizes\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setConfig\",\"inputs\":[{\"name\":\"_hasMaxGlobalShortSizes\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// ReaderABI is the input ABI used to generate the binding from.
// Deprecated: Use ReaderMetaData.ABI instead.
var ReaderABI = ReaderMetaData.ABI

// Reader is an auto generated Go binding around an Ethereum contract.
type Reader struct {
	ReaderCaller     // Read-only binding to the contract
	ReaderTransactor // Write-only binding to the contract
	ReaderFilterer   // Log filterer for contract events
}

// ReaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReaderSession struct {
	Contract     *Reader           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReaderCallerSession struct {
	Contract *ReaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ReaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReaderTransactorSession struct {
	Contract     *ReaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReaderRaw struct {
	Contract *Reader // Generic contract binding to access the raw methods on
}

// ReaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReaderCallerRaw struct {
	Contract *ReaderCaller // Generic read-only contract binding to access the raw methods on
}

// ReaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReaderTransactorRaw struct {
	Contract *ReaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReader creates a new instance of Reader, bound to a specific deployed contract.
func NewReader(address common.Address, backend bind.ContractBackend) (*Reader, error) {
	contract, err := bindReader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reader{ReaderCaller: ReaderCaller{contract: contract}, ReaderTransactor: ReaderTransactor{contract: contract}, ReaderFilterer: ReaderFilterer{contract: contract}}, nil
}

// NewReaderCaller creates a new read-only instance of Reader, bound to a specific deployed contract.
func NewReaderCaller(address common.Address, caller bind.ContractCaller) (*ReaderCaller, error) {
	contract, err := bindReader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReaderCaller{contract: contract}, nil
}

// NewReaderTransactor creates a new write-only instance of Reader, bound to a specific deployed contract.
func NewReaderTransactor(address common.Address, transactor bind.ContractTransactor) (*ReaderTransactor, error) {
	contract, err := bindReader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReaderTransactor{contract: contract}, nil
}

// NewReaderFilterer creates a new log filterer instance of Reader, bound to a specific deployed contract.
func NewReaderFilterer(address common.Address, filterer bind.ContractFilterer) (*ReaderFilterer, error) {
	contract, err := bindReader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReaderFilterer{contract: contract}, nil
}

// bindReader binds a generic wrapper to an already deployed contract.
func bindReader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReaderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reader *ReaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reader.Contract.ReaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reader *ReaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reader.Contract.ReaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reader *ReaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reader.Contract.ReaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reader *ReaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reader *ReaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reader *ReaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reader.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_Reader *ReaderCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_Reader *ReaderSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _Reader.Contract.BASISPOINTSDIVISOR(&_Reader.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_Reader *ReaderCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _Reader.Contract.BASISPOINTSDIVISOR(&_Reader.CallOpts)
}

// POSITIONPROPSLENGTH is a free data retrieval call binding the contract method 0xad7e5497.
//
// Solidity: function POSITION_PROPS_LENGTH() view returns(uint256)
func (_Reader *ReaderCaller) POSITIONPROPSLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "POSITION_PROPS_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// POSITIONPROPSLENGTH is a free data retrieval call binding the contract method 0xad7e5497.
//
// Solidity: function POSITION_PROPS_LENGTH() view returns(uint256)
func (_Reader *ReaderSession) POSITIONPROPSLENGTH() (*big.Int, error) {
	return _Reader.Contract.POSITIONPROPSLENGTH(&_Reader.CallOpts)
}

// POSITIONPROPSLENGTH is a free data retrieval call binding the contract method 0xad7e5497.
//
// Solidity: function POSITION_PROPS_LENGTH() view returns(uint256)
func (_Reader *ReaderCallerSession) POSITIONPROPSLENGTH() (*big.Int, error) {
	return _Reader.Contract.POSITIONPROPSLENGTH(&_Reader.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_Reader *ReaderCaller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_Reader *ReaderSession) PRICEPRECISION() (*big.Int, error) {
	return _Reader.Contract.PRICEPRECISION(&_Reader.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_Reader *ReaderCallerSession) PRICEPRECISION() (*big.Int, error) {
	return _Reader.Contract.PRICEPRECISION(&_Reader.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_Reader *ReaderCaller) USDGDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "USDG_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_Reader *ReaderSession) USDGDECIMALS() (*big.Int, error) {
	return _Reader.Contract.USDGDECIMALS(&_Reader.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_Reader *ReaderCallerSession) USDGDECIMALS() (*big.Int, error) {
	return _Reader.Contract.USDGDECIMALS(&_Reader.CallOpts)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xd7176ca9.
//
// Solidity: function getAmountOut(address _vault, address _tokenIn, address _tokenOut, uint256 _amountIn) view returns(uint256, uint256)
func (_Reader *ReaderCaller) GetAmountOut(opts *bind.CallOpts, _vault common.Address, _tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getAmountOut", _vault, _tokenIn, _tokenOut, _amountIn)

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
func (_Reader *ReaderSession) GetAmountOut(_vault common.Address, _tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int) (*big.Int, *big.Int, error) {
	return _Reader.Contract.GetAmountOut(&_Reader.CallOpts, _vault, _tokenIn, _tokenOut, _amountIn)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xd7176ca9.
//
// Solidity: function getAmountOut(address _vault, address _tokenIn, address _tokenOut, uint256 _amountIn) view returns(uint256, uint256)
func (_Reader *ReaderCallerSession) GetAmountOut(_vault common.Address, _tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int) (*big.Int, *big.Int, error) {
	return _Reader.Contract.GetAmountOut(&_Reader.CallOpts, _vault, _tokenIn, _tokenOut, _amountIn)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0x440f018c.
//
// Solidity: function getFeeBasisPoints(address _vault, address _tokenIn, address _tokenOut, uint256 _amountIn) view returns(uint256, uint256, uint256)
func (_Reader *ReaderCaller) GetFeeBasisPoints(opts *bind.CallOpts, _vault common.Address, _tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getFeeBasisPoints", _vault, _tokenIn, _tokenOut, _amountIn)

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
func (_Reader *ReaderSession) GetFeeBasisPoints(_vault common.Address, _tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Reader.Contract.GetFeeBasisPoints(&_Reader.CallOpts, _vault, _tokenIn, _tokenOut, _amountIn)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0x440f018c.
//
// Solidity: function getFeeBasisPoints(address _vault, address _tokenIn, address _tokenOut, uint256 _amountIn) view returns(uint256, uint256, uint256)
func (_Reader *ReaderCallerSession) GetFeeBasisPoints(_vault common.Address, _tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Reader.Contract.GetFeeBasisPoints(&_Reader.CallOpts, _vault, _tokenIn, _tokenOut, _amountIn)
}

// GetFees is a free data retrieval call binding the contract method 0x86d4d0f5.
//
// Solidity: function getFees(address _vault, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderCaller) GetFees(opts *bind.CallOpts, _vault common.Address, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getFees", _vault, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetFees is a free data retrieval call binding the contract method 0x86d4d0f5.
//
// Solidity: function getFees(address _vault, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderSession) GetFees(_vault common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetFees(&_Reader.CallOpts, _vault, _tokens)
}

// GetFees is a free data retrieval call binding the contract method 0x86d4d0f5.
//
// Solidity: function getFees(address _vault, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderCallerSession) GetFees(_vault common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetFees(&_Reader.CallOpts, _vault, _tokens)
}

// GetFullVaultTokenInfo is a free data retrieval call binding the contract method 0x7b906e93.
//
// Solidity: function getFullVaultTokenInfo(address _vault, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderCaller) GetFullVaultTokenInfo(opts *bind.CallOpts, _vault common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getFullVaultTokenInfo", _vault, _weth, _usdgAmount, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetFullVaultTokenInfo is a free data retrieval call binding the contract method 0x7b906e93.
//
// Solidity: function getFullVaultTokenInfo(address _vault, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderSession) GetFullVaultTokenInfo(_vault common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetFullVaultTokenInfo(&_Reader.CallOpts, _vault, _weth, _usdgAmount, _tokens)
}

// GetFullVaultTokenInfo is a free data retrieval call binding the contract method 0x7b906e93.
//
// Solidity: function getFullVaultTokenInfo(address _vault, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderCallerSession) GetFullVaultTokenInfo(_vault common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetFullVaultTokenInfo(&_Reader.CallOpts, _vault, _weth, _usdgAmount, _tokens)
}

// GetFundingRates is a free data retrieval call binding the contract method 0x95a7535a.
//
// Solidity: function getFundingRates(address _vault, address _weth, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderCaller) GetFundingRates(opts *bind.CallOpts, _vault common.Address, _weth common.Address, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getFundingRates", _vault, _weth, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetFundingRates is a free data retrieval call binding the contract method 0x95a7535a.
//
// Solidity: function getFundingRates(address _vault, address _weth, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderSession) GetFundingRates(_vault common.Address, _weth common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetFundingRates(&_Reader.CallOpts, _vault, _weth, _tokens)
}

// GetFundingRates is a free data retrieval call binding the contract method 0x95a7535a.
//
// Solidity: function getFundingRates(address _vault, address _weth, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderCallerSession) GetFundingRates(_vault common.Address, _weth common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetFundingRates(&_Reader.CallOpts, _vault, _weth, _tokens)
}

// GetMaxAmountIn is a free data retrieval call binding the contract method 0xf3535e6c.
//
// Solidity: function getMaxAmountIn(address _vault, address _tokenIn, address _tokenOut) view returns(uint256)
func (_Reader *ReaderCaller) GetMaxAmountIn(opts *bind.CallOpts, _vault common.Address, _tokenIn common.Address, _tokenOut common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getMaxAmountIn", _vault, _tokenIn, _tokenOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxAmountIn is a free data retrieval call binding the contract method 0xf3535e6c.
//
// Solidity: function getMaxAmountIn(address _vault, address _tokenIn, address _tokenOut) view returns(uint256)
func (_Reader *ReaderSession) GetMaxAmountIn(_vault common.Address, _tokenIn common.Address, _tokenOut common.Address) (*big.Int, error) {
	return _Reader.Contract.GetMaxAmountIn(&_Reader.CallOpts, _vault, _tokenIn, _tokenOut)
}

// GetMaxAmountIn is a free data retrieval call binding the contract method 0xf3535e6c.
//
// Solidity: function getMaxAmountIn(address _vault, address _tokenIn, address _tokenOut) view returns(uint256)
func (_Reader *ReaderCallerSession) GetMaxAmountIn(_vault common.Address, _tokenIn common.Address, _tokenOut common.Address) (*big.Int, error) {
	return _Reader.Contract.GetMaxAmountIn(&_Reader.CallOpts, _vault, _tokenIn, _tokenOut)
}

// GetPairInfo is a free data retrieval call binding the contract method 0xa4543ead.
//
// Solidity: function getPairInfo(address _factory, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderCaller) GetPairInfo(opts *bind.CallOpts, _factory common.Address, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getPairInfo", _factory, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPairInfo is a free data retrieval call binding the contract method 0xa4543ead.
//
// Solidity: function getPairInfo(address _factory, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderSession) GetPairInfo(_factory common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetPairInfo(&_Reader.CallOpts, _factory, _tokens)
}

// GetPairInfo is a free data retrieval call binding the contract method 0xa4543ead.
//
// Solidity: function getPairInfo(address _factory, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderCallerSession) GetPairInfo(_factory common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetPairInfo(&_Reader.CallOpts, _factory, _tokens)
}

// GetPositions is a free data retrieval call binding the contract method 0xdc383cab.
//
// Solidity: function getPositions(address _vault, address _account, address[] _collateralTokens, address[] _indexTokens, bool[] _isLong) view returns(uint256[])
func (_Reader *ReaderCaller) GetPositions(opts *bind.CallOpts, _vault common.Address, _account common.Address, _collateralTokens []common.Address, _indexTokens []common.Address, _isLong []bool) ([]*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getPositions", _vault, _account, _collateralTokens, _indexTokens, _isLong)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPositions is a free data retrieval call binding the contract method 0xdc383cab.
//
// Solidity: function getPositions(address _vault, address _account, address[] _collateralTokens, address[] _indexTokens, bool[] _isLong) view returns(uint256[])
func (_Reader *ReaderSession) GetPositions(_vault common.Address, _account common.Address, _collateralTokens []common.Address, _indexTokens []common.Address, _isLong []bool) ([]*big.Int, error) {
	return _Reader.Contract.GetPositions(&_Reader.CallOpts, _vault, _account, _collateralTokens, _indexTokens, _isLong)
}

// GetPositions is a free data retrieval call binding the contract method 0xdc383cab.
//
// Solidity: function getPositions(address _vault, address _account, address[] _collateralTokens, address[] _indexTokens, bool[] _isLong) view returns(uint256[])
func (_Reader *ReaderCallerSession) GetPositions(_vault common.Address, _account common.Address, _collateralTokens []common.Address, _indexTokens []common.Address, _isLong []bool) ([]*big.Int, error) {
	return _Reader.Contract.GetPositions(&_Reader.CallOpts, _vault, _account, _collateralTokens, _indexTokens, _isLong)
}

// GetPrices is a free data retrieval call binding the contract method 0x3613d527.
//
// Solidity: function getPrices(address _priceFeed, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderCaller) GetPrices(opts *bind.CallOpts, _priceFeed common.Address, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getPrices", _priceFeed, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetPrices is a free data retrieval call binding the contract method 0x3613d527.
//
// Solidity: function getPrices(address _priceFeed, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderSession) GetPrices(_priceFeed common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetPrices(&_Reader.CallOpts, _priceFeed, _tokens)
}

// GetPrices is a free data retrieval call binding the contract method 0x3613d527.
//
// Solidity: function getPrices(address _priceFeed, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderCallerSession) GetPrices(_priceFeed common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetPrices(&_Reader.CallOpts, _priceFeed, _tokens)
}

// GetStakingInfo is a free data retrieval call binding the contract method 0x937a0be8.
//
// Solidity: function getStakingInfo(address _account, address[] _yieldTrackers) view returns(uint256[])
func (_Reader *ReaderCaller) GetStakingInfo(opts *bind.CallOpts, _account common.Address, _yieldTrackers []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getStakingInfo", _account, _yieldTrackers)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetStakingInfo is a free data retrieval call binding the contract method 0x937a0be8.
//
// Solidity: function getStakingInfo(address _account, address[] _yieldTrackers) view returns(uint256[])
func (_Reader *ReaderSession) GetStakingInfo(_account common.Address, _yieldTrackers []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetStakingInfo(&_Reader.CallOpts, _account, _yieldTrackers)
}

// GetStakingInfo is a free data retrieval call binding the contract method 0x937a0be8.
//
// Solidity: function getStakingInfo(address _account, address[] _yieldTrackers) view returns(uint256[])
func (_Reader *ReaderCallerSession) GetStakingInfo(_account common.Address, _yieldTrackers []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetStakingInfo(&_Reader.CallOpts, _account, _yieldTrackers)
}

// GetTokenBalances is a free data retrieval call binding the contract method 0xd802178e.
//
// Solidity: function getTokenBalances(address _account, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderCaller) GetTokenBalances(opts *bind.CallOpts, _account common.Address, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getTokenBalances", _account, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTokenBalances is a free data retrieval call binding the contract method 0xd802178e.
//
// Solidity: function getTokenBalances(address _account, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderSession) GetTokenBalances(_account common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetTokenBalances(&_Reader.CallOpts, _account, _tokens)
}

// GetTokenBalances is a free data retrieval call binding the contract method 0xd802178e.
//
// Solidity: function getTokenBalances(address _account, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderCallerSession) GetTokenBalances(_account common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetTokenBalances(&_Reader.CallOpts, _account, _tokens)
}

// GetTokenBalancesWithSupplies is a free data retrieval call binding the contract method 0x2e3e3342.
//
// Solidity: function getTokenBalancesWithSupplies(address _account, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderCaller) GetTokenBalancesWithSupplies(opts *bind.CallOpts, _account common.Address, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getTokenBalancesWithSupplies", _account, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTokenBalancesWithSupplies is a free data retrieval call binding the contract method 0x2e3e3342.
//
// Solidity: function getTokenBalancesWithSupplies(address _account, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderSession) GetTokenBalancesWithSupplies(_account common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetTokenBalancesWithSupplies(&_Reader.CallOpts, _account, _tokens)
}

// GetTokenBalancesWithSupplies is a free data retrieval call binding the contract method 0x2e3e3342.
//
// Solidity: function getTokenBalancesWithSupplies(address _account, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderCallerSession) GetTokenBalancesWithSupplies(_account common.Address, _tokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetTokenBalancesWithSupplies(&_Reader.CallOpts, _account, _tokens)
}

// GetTokenSupply is a free data retrieval call binding the contract method 0x2ac0184c.
//
// Solidity: function getTokenSupply(address _token, address[] _excludedAccounts) view returns(uint256)
func (_Reader *ReaderCaller) GetTokenSupply(opts *bind.CallOpts, _token common.Address, _excludedAccounts []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getTokenSupply", _token, _excludedAccounts)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenSupply is a free data retrieval call binding the contract method 0x2ac0184c.
//
// Solidity: function getTokenSupply(address _token, address[] _excludedAccounts) view returns(uint256)
func (_Reader *ReaderSession) GetTokenSupply(_token common.Address, _excludedAccounts []common.Address) (*big.Int, error) {
	return _Reader.Contract.GetTokenSupply(&_Reader.CallOpts, _token, _excludedAccounts)
}

// GetTokenSupply is a free data retrieval call binding the contract method 0x2ac0184c.
//
// Solidity: function getTokenSupply(address _token, address[] _excludedAccounts) view returns(uint256)
func (_Reader *ReaderCallerSession) GetTokenSupply(_token common.Address, _excludedAccounts []common.Address) (*big.Int, error) {
	return _Reader.Contract.GetTokenSupply(&_Reader.CallOpts, _token, _excludedAccounts)
}

// GetTotalBalance is a free data retrieval call binding the contract method 0xfbdb05ca.
//
// Solidity: function getTotalBalance(address _token, address[] _accounts) view returns(uint256)
func (_Reader *ReaderCaller) GetTotalBalance(opts *bind.CallOpts, _token common.Address, _accounts []common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getTotalBalance", _token, _accounts)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBalance is a free data retrieval call binding the contract method 0xfbdb05ca.
//
// Solidity: function getTotalBalance(address _token, address[] _accounts) view returns(uint256)
func (_Reader *ReaderSession) GetTotalBalance(_token common.Address, _accounts []common.Address) (*big.Int, error) {
	return _Reader.Contract.GetTotalBalance(&_Reader.CallOpts, _token, _accounts)
}

// GetTotalBalance is a free data retrieval call binding the contract method 0xfbdb05ca.
//
// Solidity: function getTotalBalance(address _token, address[] _accounts) view returns(uint256)
func (_Reader *ReaderCallerSession) GetTotalBalance(_token common.Address, _accounts []common.Address) (*big.Int, error) {
	return _Reader.Contract.GetTotalBalance(&_Reader.CallOpts, _token, _accounts)
}

// GetTotalStaked is a free data retrieval call binding the contract method 0x2413c8c1.
//
// Solidity: function getTotalStaked(address[] _yieldTokens) view returns(uint256[])
func (_Reader *ReaderCaller) GetTotalStaked(opts *bind.CallOpts, _yieldTokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getTotalStaked", _yieldTokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTotalStaked is a free data retrieval call binding the contract method 0x2413c8c1.
//
// Solidity: function getTotalStaked(address[] _yieldTokens) view returns(uint256[])
func (_Reader *ReaderSession) GetTotalStaked(_yieldTokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetTotalStaked(&_Reader.CallOpts, _yieldTokens)
}

// GetTotalStaked is a free data retrieval call binding the contract method 0x2413c8c1.
//
// Solidity: function getTotalStaked(address[] _yieldTokens) view returns(uint256[])
func (_Reader *ReaderCallerSession) GetTotalStaked(_yieldTokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetTotalStaked(&_Reader.CallOpts, _yieldTokens)
}

// GetVaultTokenInfo is a free data retrieval call binding the contract method 0x20542568.
//
// Solidity: function getVaultTokenInfo(address _vault, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderCaller) GetVaultTokenInfo(opts *bind.CallOpts, _vault common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getVaultTokenInfo", _vault, _weth, _usdgAmount, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetVaultTokenInfo is a free data retrieval call binding the contract method 0x20542568.
//
// Solidity: function getVaultTokenInfo(address _vault, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderSession) GetVaultTokenInfo(_vault common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetVaultTokenInfo(&_Reader.CallOpts, _vault, _weth, _usdgAmount, _tokens)
}

// GetVaultTokenInfo is a free data retrieval call binding the contract method 0x20542568.
//
// Solidity: function getVaultTokenInfo(address _vault, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderCallerSession) GetVaultTokenInfo(_vault common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetVaultTokenInfo(&_Reader.CallOpts, _vault, _weth, _usdgAmount, _tokens)
}

// GetVaultTokenInfoV2 is a free data retrieval call binding the contract method 0x8e83ca32.
//
// Solidity: function getVaultTokenInfoV2(address _vault, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderCaller) GetVaultTokenInfoV2(opts *bind.CallOpts, _vault common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getVaultTokenInfoV2", _vault, _weth, _usdgAmount, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetVaultTokenInfoV2 is a free data retrieval call binding the contract method 0x8e83ca32.
//
// Solidity: function getVaultTokenInfoV2(address _vault, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderSession) GetVaultTokenInfoV2(_vault common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetVaultTokenInfoV2(&_Reader.CallOpts, _vault, _weth, _usdgAmount, _tokens)
}

// GetVaultTokenInfoV2 is a free data retrieval call binding the contract method 0x8e83ca32.
//
// Solidity: function getVaultTokenInfoV2(address _vault, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_Reader *ReaderCallerSession) GetVaultTokenInfoV2(_vault common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetVaultTokenInfoV2(&_Reader.CallOpts, _vault, _weth, _usdgAmount, _tokens)
}

// GetVestingInfo is a free data retrieval call binding the contract method 0x48211934.
//
// Solidity: function getVestingInfo(address _account, address[] _vesters) view returns(uint256[])
func (_Reader *ReaderCaller) GetVestingInfo(opts *bind.CallOpts, _account common.Address, _vesters []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "getVestingInfo", _account, _vesters)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetVestingInfo is a free data retrieval call binding the contract method 0x48211934.
//
// Solidity: function getVestingInfo(address _account, address[] _vesters) view returns(uint256[])
func (_Reader *ReaderSession) GetVestingInfo(_account common.Address, _vesters []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetVestingInfo(&_Reader.CallOpts, _account, _vesters)
}

// GetVestingInfo is a free data retrieval call binding the contract method 0x48211934.
//
// Solidity: function getVestingInfo(address _account, address[] _vesters) view returns(uint256[])
func (_Reader *ReaderCallerSession) GetVestingInfo(_account common.Address, _vesters []common.Address) ([]*big.Int, error) {
	return _Reader.Contract.GetVestingInfo(&_Reader.CallOpts, _account, _vesters)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Reader *ReaderCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Reader *ReaderSession) Gov() (common.Address, error) {
	return _Reader.Contract.Gov(&_Reader.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Reader *ReaderCallerSession) Gov() (common.Address, error) {
	return _Reader.Contract.Gov(&_Reader.CallOpts)
}

// HasMaxGlobalShortSizes is a free data retrieval call binding the contract method 0xc6f1d676.
//
// Solidity: function hasMaxGlobalShortSizes() view returns(bool)
func (_Reader *ReaderCaller) HasMaxGlobalShortSizes(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Reader.contract.Call(opts, &out, "hasMaxGlobalShortSizes")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasMaxGlobalShortSizes is a free data retrieval call binding the contract method 0xc6f1d676.
//
// Solidity: function hasMaxGlobalShortSizes() view returns(bool)
func (_Reader *ReaderSession) HasMaxGlobalShortSizes() (bool, error) {
	return _Reader.Contract.HasMaxGlobalShortSizes(&_Reader.CallOpts)
}

// HasMaxGlobalShortSizes is a free data retrieval call binding the contract method 0xc6f1d676.
//
// Solidity: function hasMaxGlobalShortSizes() view returns(bool)
func (_Reader *ReaderCallerSession) HasMaxGlobalShortSizes() (bool, error) {
	return _Reader.Contract.HasMaxGlobalShortSizes(&_Reader.CallOpts)
}

// SetConfig is a paid mutator transaction binding the contract method 0x9b0183c3.
//
// Solidity: function setConfig(bool _hasMaxGlobalShortSizes) returns()
func (_Reader *ReaderTransactor) SetConfig(opts *bind.TransactOpts, _hasMaxGlobalShortSizes bool) (*types.Transaction, error) {
	return _Reader.contract.Transact(opts, "setConfig", _hasMaxGlobalShortSizes)
}

// SetConfig is a paid mutator transaction binding the contract method 0x9b0183c3.
//
// Solidity: function setConfig(bool _hasMaxGlobalShortSizes) returns()
func (_Reader *ReaderSession) SetConfig(_hasMaxGlobalShortSizes bool) (*types.Transaction, error) {
	return _Reader.Contract.SetConfig(&_Reader.TransactOpts, _hasMaxGlobalShortSizes)
}

// SetConfig is a paid mutator transaction binding the contract method 0x9b0183c3.
//
// Solidity: function setConfig(bool _hasMaxGlobalShortSizes) returns()
func (_Reader *ReaderTransactorSession) SetConfig(_hasMaxGlobalShortSizes bool) (*types.Transaction, error) {
	return _Reader.Contract.SetConfig(&_Reader.TransactOpts, _hasMaxGlobalShortSizes)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Reader *ReaderTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _Reader.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Reader *ReaderSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Reader.Contract.SetGov(&_Reader.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Reader *ReaderTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Reader.Contract.SetGov(&_Reader.TransactOpts, _gov)
}

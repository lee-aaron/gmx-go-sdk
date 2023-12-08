// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Treasury

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

// TreasuryMetaData contains all meta data concerning the Treasury contract.
var TreasuryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addLiquidity\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addWhitelists\",\"inputs\":[{\"name\":\"_accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"busd\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"busdBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"busdHardCap\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"busdReceived\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"busdSlotCap\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"endSwap\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"extendUnlockTime\",\"inputs\":[{\"name\":\"_unlockTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fund\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gmt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gmtListingPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gmtPresalePrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseBusdBasisPoints\",\"inputs\":[{\"name\":\"_busdBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isInitialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isLiquidityAdded\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSwapActive\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeWhitelists\",\"inputs\":[{\"name\":\"_accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setFund\",\"inputs\":[{\"name\":\"_fund\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"_busdAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swapAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"swapWhitelist\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unlockTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateWhitelist\",\"inputs\":[{\"name\":\"prevAccount\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nextAccount\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// TreasuryABI is the input ABI used to generate the binding from.
// Deprecated: Use TreasuryMetaData.ABI instead.
var TreasuryABI = TreasuryMetaData.ABI

// Treasury is an auto generated Go binding around an Ethereum contract.
type Treasury struct {
	TreasuryCaller     // Read-only binding to the contract
	TreasuryTransactor // Write-only binding to the contract
	TreasuryFilterer   // Log filterer for contract events
}

// TreasuryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TreasuryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TreasuryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TreasuryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasurySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TreasurySession struct {
	Contract     *Treasury         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TreasuryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TreasuryCallerSession struct {
	Contract *TreasuryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TreasuryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TreasuryTransactorSession struct {
	Contract     *TreasuryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TreasuryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TreasuryRaw struct {
	Contract *Treasury // Generic contract binding to access the raw methods on
}

// TreasuryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TreasuryCallerRaw struct {
	Contract *TreasuryCaller // Generic read-only contract binding to access the raw methods on
}

// TreasuryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TreasuryTransactorRaw struct {
	Contract *TreasuryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTreasury creates a new instance of Treasury, bound to a specific deployed contract.
func NewTreasury(address common.Address, backend bind.ContractBackend) (*Treasury, error) {
	contract, err := bindTreasury(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Treasury{TreasuryCaller: TreasuryCaller{contract: contract}, TreasuryTransactor: TreasuryTransactor{contract: contract}, TreasuryFilterer: TreasuryFilterer{contract: contract}}, nil
}

// NewTreasuryCaller creates a new read-only instance of Treasury, bound to a specific deployed contract.
func NewTreasuryCaller(address common.Address, caller bind.ContractCaller) (*TreasuryCaller, error) {
	contract, err := bindTreasury(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryCaller{contract: contract}, nil
}

// NewTreasuryTransactor creates a new write-only instance of Treasury, bound to a specific deployed contract.
func NewTreasuryTransactor(address common.Address, transactor bind.ContractTransactor) (*TreasuryTransactor, error) {
	contract, err := bindTreasury(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryTransactor{contract: contract}, nil
}

// NewTreasuryFilterer creates a new log filterer instance of Treasury, bound to a specific deployed contract.
func NewTreasuryFilterer(address common.Address, filterer bind.ContractFilterer) (*TreasuryFilterer, error) {
	contract, err := bindTreasury(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TreasuryFilterer{contract: contract}, nil
}

// bindTreasury binds a generic wrapper to an already deployed contract.
func bindTreasury(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TreasuryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Treasury *TreasuryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Treasury.Contract.TreasuryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Treasury *TreasuryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.Contract.TreasuryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Treasury *TreasuryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Treasury.Contract.TreasuryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Treasury *TreasuryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Treasury.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Treasury *TreasuryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Treasury *TreasuryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Treasury.Contract.contract.Transact(opts, method, params...)
}

// Busd is a free data retrieval call binding the contract method 0x3ca5b234.
//
// Solidity: function busd() view returns(address)
func (_Treasury *TreasuryCaller) Busd(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "busd")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Busd is a free data retrieval call binding the contract method 0x3ca5b234.
//
// Solidity: function busd() view returns(address)
func (_Treasury *TreasurySession) Busd() (common.Address, error) {
	return _Treasury.Contract.Busd(&_Treasury.CallOpts)
}

// Busd is a free data retrieval call binding the contract method 0x3ca5b234.
//
// Solidity: function busd() view returns(address)
func (_Treasury *TreasuryCallerSession) Busd() (common.Address, error) {
	return _Treasury.Contract.Busd(&_Treasury.CallOpts)
}

// BusdBasisPoints is a free data retrieval call binding the contract method 0x77b4da2c.
//
// Solidity: function busdBasisPoints() view returns(uint256)
func (_Treasury *TreasuryCaller) BusdBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "busdBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BusdBasisPoints is a free data retrieval call binding the contract method 0x77b4da2c.
//
// Solidity: function busdBasisPoints() view returns(uint256)
func (_Treasury *TreasurySession) BusdBasisPoints() (*big.Int, error) {
	return _Treasury.Contract.BusdBasisPoints(&_Treasury.CallOpts)
}

// BusdBasisPoints is a free data retrieval call binding the contract method 0x77b4da2c.
//
// Solidity: function busdBasisPoints() view returns(uint256)
func (_Treasury *TreasuryCallerSession) BusdBasisPoints() (*big.Int, error) {
	return _Treasury.Contract.BusdBasisPoints(&_Treasury.CallOpts)
}

// BusdHardCap is a free data retrieval call binding the contract method 0x63bd4bdf.
//
// Solidity: function busdHardCap() view returns(uint256)
func (_Treasury *TreasuryCaller) BusdHardCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "busdHardCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BusdHardCap is a free data retrieval call binding the contract method 0x63bd4bdf.
//
// Solidity: function busdHardCap() view returns(uint256)
func (_Treasury *TreasurySession) BusdHardCap() (*big.Int, error) {
	return _Treasury.Contract.BusdHardCap(&_Treasury.CallOpts)
}

// BusdHardCap is a free data retrieval call binding the contract method 0x63bd4bdf.
//
// Solidity: function busdHardCap() view returns(uint256)
func (_Treasury *TreasuryCallerSession) BusdHardCap() (*big.Int, error) {
	return _Treasury.Contract.BusdHardCap(&_Treasury.CallOpts)
}

// BusdReceived is a free data retrieval call binding the contract method 0xf6f4f55b.
//
// Solidity: function busdReceived() view returns(uint256)
func (_Treasury *TreasuryCaller) BusdReceived(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "busdReceived")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BusdReceived is a free data retrieval call binding the contract method 0xf6f4f55b.
//
// Solidity: function busdReceived() view returns(uint256)
func (_Treasury *TreasurySession) BusdReceived() (*big.Int, error) {
	return _Treasury.Contract.BusdReceived(&_Treasury.CallOpts)
}

// BusdReceived is a free data retrieval call binding the contract method 0xf6f4f55b.
//
// Solidity: function busdReceived() view returns(uint256)
func (_Treasury *TreasuryCallerSession) BusdReceived() (*big.Int, error) {
	return _Treasury.Contract.BusdReceived(&_Treasury.CallOpts)
}

// BusdSlotCap is a free data retrieval call binding the contract method 0xdec9b196.
//
// Solidity: function busdSlotCap() view returns(uint256)
func (_Treasury *TreasuryCaller) BusdSlotCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "busdSlotCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BusdSlotCap is a free data retrieval call binding the contract method 0xdec9b196.
//
// Solidity: function busdSlotCap() view returns(uint256)
func (_Treasury *TreasurySession) BusdSlotCap() (*big.Int, error) {
	return _Treasury.Contract.BusdSlotCap(&_Treasury.CallOpts)
}

// BusdSlotCap is a free data retrieval call binding the contract method 0xdec9b196.
//
// Solidity: function busdSlotCap() view returns(uint256)
func (_Treasury *TreasuryCallerSession) BusdSlotCap() (*big.Int, error) {
	return _Treasury.Contract.BusdSlotCap(&_Treasury.CallOpts)
}

// Fund is a free data retrieval call binding the contract method 0xb60d4288.
//
// Solidity: function fund() view returns(address)
func (_Treasury *TreasuryCaller) Fund(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "fund")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Fund is a free data retrieval call binding the contract method 0xb60d4288.
//
// Solidity: function fund() view returns(address)
func (_Treasury *TreasurySession) Fund() (common.Address, error) {
	return _Treasury.Contract.Fund(&_Treasury.CallOpts)
}

// Fund is a free data retrieval call binding the contract method 0xb60d4288.
//
// Solidity: function fund() view returns(address)
func (_Treasury *TreasuryCallerSession) Fund() (common.Address, error) {
	return _Treasury.Contract.Fund(&_Treasury.CallOpts)
}

// Gmt is a free data retrieval call binding the contract method 0xf8322d24.
//
// Solidity: function gmt() view returns(address)
func (_Treasury *TreasuryCaller) Gmt(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "gmt")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gmt is a free data retrieval call binding the contract method 0xf8322d24.
//
// Solidity: function gmt() view returns(address)
func (_Treasury *TreasurySession) Gmt() (common.Address, error) {
	return _Treasury.Contract.Gmt(&_Treasury.CallOpts)
}

// Gmt is a free data retrieval call binding the contract method 0xf8322d24.
//
// Solidity: function gmt() view returns(address)
func (_Treasury *TreasuryCallerSession) Gmt() (common.Address, error) {
	return _Treasury.Contract.Gmt(&_Treasury.CallOpts)
}

// GmtListingPrice is a free data retrieval call binding the contract method 0x0963e9fc.
//
// Solidity: function gmtListingPrice() view returns(uint256)
func (_Treasury *TreasuryCaller) GmtListingPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "gmtListingPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GmtListingPrice is a free data retrieval call binding the contract method 0x0963e9fc.
//
// Solidity: function gmtListingPrice() view returns(uint256)
func (_Treasury *TreasurySession) GmtListingPrice() (*big.Int, error) {
	return _Treasury.Contract.GmtListingPrice(&_Treasury.CallOpts)
}

// GmtListingPrice is a free data retrieval call binding the contract method 0x0963e9fc.
//
// Solidity: function gmtListingPrice() view returns(uint256)
func (_Treasury *TreasuryCallerSession) GmtListingPrice() (*big.Int, error) {
	return _Treasury.Contract.GmtListingPrice(&_Treasury.CallOpts)
}

// GmtPresalePrice is a free data retrieval call binding the contract method 0x7c098091.
//
// Solidity: function gmtPresalePrice() view returns(uint256)
func (_Treasury *TreasuryCaller) GmtPresalePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "gmtPresalePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GmtPresalePrice is a free data retrieval call binding the contract method 0x7c098091.
//
// Solidity: function gmtPresalePrice() view returns(uint256)
func (_Treasury *TreasurySession) GmtPresalePrice() (*big.Int, error) {
	return _Treasury.Contract.GmtPresalePrice(&_Treasury.CallOpts)
}

// GmtPresalePrice is a free data retrieval call binding the contract method 0x7c098091.
//
// Solidity: function gmtPresalePrice() view returns(uint256)
func (_Treasury *TreasuryCallerSession) GmtPresalePrice() (*big.Int, error) {
	return _Treasury.Contract.GmtPresalePrice(&_Treasury.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Treasury *TreasuryCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Treasury *TreasurySession) Gov() (common.Address, error) {
	return _Treasury.Contract.Gov(&_Treasury.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Treasury *TreasuryCallerSession) Gov() (common.Address, error) {
	return _Treasury.Contract.Gov(&_Treasury.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Treasury *TreasuryCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Treasury *TreasurySession) IsInitialized() (bool, error) {
	return _Treasury.Contract.IsInitialized(&_Treasury.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Treasury *TreasuryCallerSession) IsInitialized() (bool, error) {
	return _Treasury.Contract.IsInitialized(&_Treasury.CallOpts)
}

// IsLiquidityAdded is a free data retrieval call binding the contract method 0x1e17ba39.
//
// Solidity: function isLiquidityAdded() view returns(bool)
func (_Treasury *TreasuryCaller) IsLiquidityAdded(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "isLiquidityAdded")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidityAdded is a free data retrieval call binding the contract method 0x1e17ba39.
//
// Solidity: function isLiquidityAdded() view returns(bool)
func (_Treasury *TreasurySession) IsLiquidityAdded() (bool, error) {
	return _Treasury.Contract.IsLiquidityAdded(&_Treasury.CallOpts)
}

// IsLiquidityAdded is a free data retrieval call binding the contract method 0x1e17ba39.
//
// Solidity: function isLiquidityAdded() view returns(bool)
func (_Treasury *TreasuryCallerSession) IsLiquidityAdded() (bool, error) {
	return _Treasury.Contract.IsLiquidityAdded(&_Treasury.CallOpts)
}

// IsSwapActive is a free data retrieval call binding the contract method 0xe35bff96.
//
// Solidity: function isSwapActive() view returns(bool)
func (_Treasury *TreasuryCaller) IsSwapActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "isSwapActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSwapActive is a free data retrieval call binding the contract method 0xe35bff96.
//
// Solidity: function isSwapActive() view returns(bool)
func (_Treasury *TreasurySession) IsSwapActive() (bool, error) {
	return _Treasury.Contract.IsSwapActive(&_Treasury.CallOpts)
}

// IsSwapActive is a free data retrieval call binding the contract method 0xe35bff96.
//
// Solidity: function isSwapActive() view returns(bool)
func (_Treasury *TreasuryCallerSession) IsSwapActive() (bool, error) {
	return _Treasury.Contract.IsSwapActive(&_Treasury.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Treasury *TreasuryCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Treasury *TreasurySession) Router() (common.Address, error) {
	return _Treasury.Contract.Router(&_Treasury.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Treasury *TreasuryCallerSession) Router() (common.Address, error) {
	return _Treasury.Contract.Router(&_Treasury.CallOpts)
}

// SwapAmounts is a free data retrieval call binding the contract method 0x442bbad6.
//
// Solidity: function swapAmounts(address ) view returns(uint256)
func (_Treasury *TreasuryCaller) SwapAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "swapAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapAmounts is a free data retrieval call binding the contract method 0x442bbad6.
//
// Solidity: function swapAmounts(address ) view returns(uint256)
func (_Treasury *TreasurySession) SwapAmounts(arg0 common.Address) (*big.Int, error) {
	return _Treasury.Contract.SwapAmounts(&_Treasury.CallOpts, arg0)
}

// SwapAmounts is a free data retrieval call binding the contract method 0x442bbad6.
//
// Solidity: function swapAmounts(address ) view returns(uint256)
func (_Treasury *TreasuryCallerSession) SwapAmounts(arg0 common.Address) (*big.Int, error) {
	return _Treasury.Contract.SwapAmounts(&_Treasury.CallOpts, arg0)
}

// SwapWhitelist is a free data retrieval call binding the contract method 0xd0435961.
//
// Solidity: function swapWhitelist(address ) view returns(bool)
func (_Treasury *TreasuryCaller) SwapWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "swapWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SwapWhitelist is a free data retrieval call binding the contract method 0xd0435961.
//
// Solidity: function swapWhitelist(address ) view returns(bool)
func (_Treasury *TreasurySession) SwapWhitelist(arg0 common.Address) (bool, error) {
	return _Treasury.Contract.SwapWhitelist(&_Treasury.CallOpts, arg0)
}

// SwapWhitelist is a free data retrieval call binding the contract method 0xd0435961.
//
// Solidity: function swapWhitelist(address ) view returns(bool)
func (_Treasury *TreasuryCallerSession) SwapWhitelist(arg0 common.Address) (bool, error) {
	return _Treasury.Contract.SwapWhitelist(&_Treasury.CallOpts, arg0)
}

// UnlockTime is a free data retrieval call binding the contract method 0x251c1aa3.
//
// Solidity: function unlockTime() view returns(uint256)
func (_Treasury *TreasuryCaller) UnlockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "unlockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockTime is a free data retrieval call binding the contract method 0x251c1aa3.
//
// Solidity: function unlockTime() view returns(uint256)
func (_Treasury *TreasurySession) UnlockTime() (*big.Int, error) {
	return _Treasury.Contract.UnlockTime(&_Treasury.CallOpts)
}

// UnlockTime is a free data retrieval call binding the contract method 0x251c1aa3.
//
// Solidity: function unlockTime() view returns(uint256)
func (_Treasury *TreasuryCallerSession) UnlockTime() (*big.Int, error) {
	return _Treasury.Contract.UnlockTime(&_Treasury.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8078d94.
//
// Solidity: function addLiquidity() returns()
func (_Treasury *TreasuryTransactor) AddLiquidity(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "addLiquidity")
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8078d94.
//
// Solidity: function addLiquidity() returns()
func (_Treasury *TreasurySession) AddLiquidity() (*types.Transaction, error) {
	return _Treasury.Contract.AddLiquidity(&_Treasury.TransactOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0xe8078d94.
//
// Solidity: function addLiquidity() returns()
func (_Treasury *TreasuryTransactorSession) AddLiquidity() (*types.Transaction, error) {
	return _Treasury.Contract.AddLiquidity(&_Treasury.TransactOpts)
}

// AddWhitelists is a paid mutator transaction binding the contract method 0xc8eaf28f.
//
// Solidity: function addWhitelists(address[] _accounts) returns()
func (_Treasury *TreasuryTransactor) AddWhitelists(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "addWhitelists", _accounts)
}

// AddWhitelists is a paid mutator transaction binding the contract method 0xc8eaf28f.
//
// Solidity: function addWhitelists(address[] _accounts) returns()
func (_Treasury *TreasurySession) AddWhitelists(_accounts []common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.AddWhitelists(&_Treasury.TransactOpts, _accounts)
}

// AddWhitelists is a paid mutator transaction binding the contract method 0xc8eaf28f.
//
// Solidity: function addWhitelists(address[] _accounts) returns()
func (_Treasury *TreasuryTransactorSession) AddWhitelists(_accounts []common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.AddWhitelists(&_Treasury.TransactOpts, _accounts)
}

// EndSwap is a paid mutator transaction binding the contract method 0xe3cb6ea3.
//
// Solidity: function endSwap() returns()
func (_Treasury *TreasuryTransactor) EndSwap(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "endSwap")
}

// EndSwap is a paid mutator transaction binding the contract method 0xe3cb6ea3.
//
// Solidity: function endSwap() returns()
func (_Treasury *TreasurySession) EndSwap() (*types.Transaction, error) {
	return _Treasury.Contract.EndSwap(&_Treasury.TransactOpts)
}

// EndSwap is a paid mutator transaction binding the contract method 0xe3cb6ea3.
//
// Solidity: function endSwap() returns()
func (_Treasury *TreasuryTransactorSession) EndSwap() (*types.Transaction, error) {
	return _Treasury.Contract.EndSwap(&_Treasury.TransactOpts)
}

// ExtendUnlockTime is a paid mutator transaction binding the contract method 0xc2333ee8.
//
// Solidity: function extendUnlockTime(uint256 _unlockTime) returns()
func (_Treasury *TreasuryTransactor) ExtendUnlockTime(opts *bind.TransactOpts, _unlockTime *big.Int) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "extendUnlockTime", _unlockTime)
}

// ExtendUnlockTime is a paid mutator transaction binding the contract method 0xc2333ee8.
//
// Solidity: function extendUnlockTime(uint256 _unlockTime) returns()
func (_Treasury *TreasurySession) ExtendUnlockTime(_unlockTime *big.Int) (*types.Transaction, error) {
	return _Treasury.Contract.ExtendUnlockTime(&_Treasury.TransactOpts, _unlockTime)
}

// ExtendUnlockTime is a paid mutator transaction binding the contract method 0xc2333ee8.
//
// Solidity: function extendUnlockTime(uint256 _unlockTime) returns()
func (_Treasury *TreasuryTransactorSession) ExtendUnlockTime(_unlockTime *big.Int) (*types.Transaction, error) {
	return _Treasury.Contract.ExtendUnlockTime(&_Treasury.TransactOpts, _unlockTime)
}

// IncreaseBusdBasisPoints is a paid mutator transaction binding the contract method 0x21089a50.
//
// Solidity: function increaseBusdBasisPoints(uint256 _busdBasisPoints) returns()
func (_Treasury *TreasuryTransactor) IncreaseBusdBasisPoints(opts *bind.TransactOpts, _busdBasisPoints *big.Int) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "increaseBusdBasisPoints", _busdBasisPoints)
}

// IncreaseBusdBasisPoints is a paid mutator transaction binding the contract method 0x21089a50.
//
// Solidity: function increaseBusdBasisPoints(uint256 _busdBasisPoints) returns()
func (_Treasury *TreasurySession) IncreaseBusdBasisPoints(_busdBasisPoints *big.Int) (*types.Transaction, error) {
	return _Treasury.Contract.IncreaseBusdBasisPoints(&_Treasury.TransactOpts, _busdBasisPoints)
}

// IncreaseBusdBasisPoints is a paid mutator transaction binding the contract method 0x21089a50.
//
// Solidity: function increaseBusdBasisPoints(uint256 _busdBasisPoints) returns()
func (_Treasury *TreasuryTransactorSession) IncreaseBusdBasisPoints(_busdBasisPoints *big.Int) (*types.Transaction, error) {
	return _Treasury.Contract.IncreaseBusdBasisPoints(&_Treasury.TransactOpts, _busdBasisPoints)
}

// Initialize is a paid mutator transaction binding the contract method 0x7fbbe46f.
//
// Solidity: function initialize(address[] _addresses, uint256[] _values) returns()
func (_Treasury *TreasuryTransactor) Initialize(opts *bind.TransactOpts, _addresses []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "initialize", _addresses, _values)
}

// Initialize is a paid mutator transaction binding the contract method 0x7fbbe46f.
//
// Solidity: function initialize(address[] _addresses, uint256[] _values) returns()
func (_Treasury *TreasurySession) Initialize(_addresses []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _Treasury.Contract.Initialize(&_Treasury.TransactOpts, _addresses, _values)
}

// Initialize is a paid mutator transaction binding the contract method 0x7fbbe46f.
//
// Solidity: function initialize(address[] _addresses, uint256[] _values) returns()
func (_Treasury *TreasuryTransactorSession) Initialize(_addresses []common.Address, _values []*big.Int) (*types.Transaction, error) {
	return _Treasury.Contract.Initialize(&_Treasury.TransactOpts, _addresses, _values)
}

// RemoveWhitelists is a paid mutator transaction binding the contract method 0x281bd798.
//
// Solidity: function removeWhitelists(address[] _accounts) returns()
func (_Treasury *TreasuryTransactor) RemoveWhitelists(opts *bind.TransactOpts, _accounts []common.Address) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "removeWhitelists", _accounts)
}

// RemoveWhitelists is a paid mutator transaction binding the contract method 0x281bd798.
//
// Solidity: function removeWhitelists(address[] _accounts) returns()
func (_Treasury *TreasurySession) RemoveWhitelists(_accounts []common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.RemoveWhitelists(&_Treasury.TransactOpts, _accounts)
}

// RemoveWhitelists is a paid mutator transaction binding the contract method 0x281bd798.
//
// Solidity: function removeWhitelists(address[] _accounts) returns()
func (_Treasury *TreasuryTransactorSession) RemoveWhitelists(_accounts []common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.RemoveWhitelists(&_Treasury.TransactOpts, _accounts)
}

// SetFund is a paid mutator transaction binding the contract method 0x0e21750f.
//
// Solidity: function setFund(address _fund) returns()
func (_Treasury *TreasuryTransactor) SetFund(opts *bind.TransactOpts, _fund common.Address) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "setFund", _fund)
}

// SetFund is a paid mutator transaction binding the contract method 0x0e21750f.
//
// Solidity: function setFund(address _fund) returns()
func (_Treasury *TreasurySession) SetFund(_fund common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.SetFund(&_Treasury.TransactOpts, _fund)
}

// SetFund is a paid mutator transaction binding the contract method 0x0e21750f.
//
// Solidity: function setFund(address _fund) returns()
func (_Treasury *TreasuryTransactorSession) SetFund(_fund common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.SetFund(&_Treasury.TransactOpts, _fund)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Treasury *TreasuryTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Treasury *TreasurySession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.SetGov(&_Treasury.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Treasury *TreasuryTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.SetGov(&_Treasury.TransactOpts, _gov)
}

// Swap is a paid mutator transaction binding the contract method 0x94b918de.
//
// Solidity: function swap(uint256 _busdAmount) returns()
func (_Treasury *TreasuryTransactor) Swap(opts *bind.TransactOpts, _busdAmount *big.Int) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "swap", _busdAmount)
}

// Swap is a paid mutator transaction binding the contract method 0x94b918de.
//
// Solidity: function swap(uint256 _busdAmount) returns()
func (_Treasury *TreasurySession) Swap(_busdAmount *big.Int) (*types.Transaction, error) {
	return _Treasury.Contract.Swap(&_Treasury.TransactOpts, _busdAmount)
}

// Swap is a paid mutator transaction binding the contract method 0x94b918de.
//
// Solidity: function swap(uint256 _busdAmount) returns()
func (_Treasury *TreasuryTransactorSession) Swap(_busdAmount *big.Int) (*types.Transaction, error) {
	return _Treasury.Contract.Swap(&_Treasury.TransactOpts, _busdAmount)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0xab2b04cf.
//
// Solidity: function updateWhitelist(address prevAccount, address nextAccount) returns()
func (_Treasury *TreasuryTransactor) UpdateWhitelist(opts *bind.TransactOpts, prevAccount common.Address, nextAccount common.Address) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "updateWhitelist", prevAccount, nextAccount)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0xab2b04cf.
//
// Solidity: function updateWhitelist(address prevAccount, address nextAccount) returns()
func (_Treasury *TreasurySession) UpdateWhitelist(prevAccount common.Address, nextAccount common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.UpdateWhitelist(&_Treasury.TransactOpts, prevAccount, nextAccount)
}

// UpdateWhitelist is a paid mutator transaction binding the contract method 0xab2b04cf.
//
// Solidity: function updateWhitelist(address prevAccount, address nextAccount) returns()
func (_Treasury *TreasuryTransactorSession) UpdateWhitelist(prevAccount common.Address, nextAccount common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.UpdateWhitelist(&_Treasury.TransactOpts, prevAccount, nextAccount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_Treasury *TreasuryTransactor) WithdrawToken(opts *bind.TransactOpts, _token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "withdrawToken", _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_Treasury *TreasurySession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Treasury.Contract.WithdrawToken(&_Treasury.TransactOpts, _token, _account, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x01e33667.
//
// Solidity: function withdrawToken(address _token, address _account, uint256 _amount) returns()
func (_Treasury *TreasuryTransactorSession) WithdrawToken(_token common.Address, _account common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Treasury.Contract.WithdrawToken(&_Treasury.TransactOpts, _token, _account, _amount)
}

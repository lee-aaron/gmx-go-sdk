// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ReferralStorage

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

// ReferralStorageMetaData contains all meta data concerning the ReferralStorage contract.
var ReferralStorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"BASIS_POINTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"codeOwners\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTraderReferralInfo\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"govSetCodeOwner\",\"inputs\":[{\"name\":\"_code\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_newAccount\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isHandler\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"referrerDiscountShares\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"referrerTiers\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerCode\",\"inputs\":[{\"name\":\"_code\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCodeOwner\",\"inputs\":[{\"name\":\"_code\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_newAccount\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setHandler\",\"inputs\":[{\"name\":\"_handler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setReferrerDiscountShare\",\"inputs\":[{\"name\":\"_discountShare\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setReferrerTier\",\"inputs\":[{\"name\":\"_referrer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tierId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTier\",\"inputs\":[{\"name\":\"_tierId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_totalRebate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_discountShare\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTraderReferralCode\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_code\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTraderReferralCodeByUser\",\"inputs\":[{\"name\":\"_code\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tiers\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"totalRebate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"discountShare\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"traderReferralCodes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"GovSetCodeOwner\",\"inputs\":[{\"name\":\"code\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"newAccount\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegisterCode\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"code\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetCodeOwner\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAccount\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"code\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetHandler\",\"inputs\":[{\"name\":\"handler\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"isActive\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetReferrerDiscountShare\",\"inputs\":[{\"name\":\"referrer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"discountShare\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetReferrerTier\",\"inputs\":[{\"name\":\"referrer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tierId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetTier\",\"inputs\":[{\"name\":\"tierId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"totalRebate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"discountShare\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetTraderReferralCode\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"code\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false}]",
}

// ReferralStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use ReferralStorageMetaData.ABI instead.
var ReferralStorageABI = ReferralStorageMetaData.ABI

// ReferralStorage is an auto generated Go binding around an Ethereum contract.
type ReferralStorage struct {
	ReferralStorageCaller     // Read-only binding to the contract
	ReferralStorageTransactor // Write-only binding to the contract
	ReferralStorageFilterer   // Log filterer for contract events
}

// ReferralStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReferralStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReferralStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReferralStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReferralStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReferralStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReferralStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReferralStorageSession struct {
	Contract     *ReferralStorage  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReferralStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReferralStorageCallerSession struct {
	Contract *ReferralStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReferralStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReferralStorageTransactorSession struct {
	Contract     *ReferralStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReferralStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReferralStorageRaw struct {
	Contract *ReferralStorage // Generic contract binding to access the raw methods on
}

// ReferralStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReferralStorageCallerRaw struct {
	Contract *ReferralStorageCaller // Generic read-only contract binding to access the raw methods on
}

// ReferralStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReferralStorageTransactorRaw struct {
	Contract *ReferralStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReferralStorage creates a new instance of ReferralStorage, bound to a specific deployed contract.
func NewReferralStorage(address common.Address, backend bind.ContractBackend) (*ReferralStorage, error) {
	contract, err := bindReferralStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReferralStorage{ReferralStorageCaller: ReferralStorageCaller{contract: contract}, ReferralStorageTransactor: ReferralStorageTransactor{contract: contract}, ReferralStorageFilterer: ReferralStorageFilterer{contract: contract}}, nil
}

// NewReferralStorageCaller creates a new read-only instance of ReferralStorage, bound to a specific deployed contract.
func NewReferralStorageCaller(address common.Address, caller bind.ContractCaller) (*ReferralStorageCaller, error) {
	contract, err := bindReferralStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReferralStorageCaller{contract: contract}, nil
}

// NewReferralStorageTransactor creates a new write-only instance of ReferralStorage, bound to a specific deployed contract.
func NewReferralStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*ReferralStorageTransactor, error) {
	contract, err := bindReferralStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReferralStorageTransactor{contract: contract}, nil
}

// NewReferralStorageFilterer creates a new log filterer instance of ReferralStorage, bound to a specific deployed contract.
func NewReferralStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*ReferralStorageFilterer, error) {
	contract, err := bindReferralStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReferralStorageFilterer{contract: contract}, nil
}

// bindReferralStorage binds a generic wrapper to an already deployed contract.
func bindReferralStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReferralStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReferralStorage *ReferralStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReferralStorage.Contract.ReferralStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReferralStorage *ReferralStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReferralStorage.Contract.ReferralStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReferralStorage *ReferralStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReferralStorage.Contract.ReferralStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReferralStorage *ReferralStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReferralStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReferralStorage *ReferralStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReferralStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReferralStorage *ReferralStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReferralStorage.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_ReferralStorage *ReferralStorageCaller) BASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ReferralStorage.contract.Call(opts, &out, "BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_ReferralStorage *ReferralStorageSession) BASISPOINTS() (*big.Int, error) {
	return _ReferralStorage.Contract.BASISPOINTS(&_ReferralStorage.CallOpts)
}

// BASISPOINTS is a free data retrieval call binding the contract method 0xe1f1c4a7.
//
// Solidity: function BASIS_POINTS() view returns(uint256)
func (_ReferralStorage *ReferralStorageCallerSession) BASISPOINTS() (*big.Int, error) {
	return _ReferralStorage.Contract.BASISPOINTS(&_ReferralStorage.CallOpts)
}

// CodeOwners is a free data retrieval call binding the contract method 0xc8b3c460.
//
// Solidity: function codeOwners(bytes32 ) view returns(address)
func (_ReferralStorage *ReferralStorageCaller) CodeOwners(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ReferralStorage.contract.Call(opts, &out, "codeOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CodeOwners is a free data retrieval call binding the contract method 0xc8b3c460.
//
// Solidity: function codeOwners(bytes32 ) view returns(address)
func (_ReferralStorage *ReferralStorageSession) CodeOwners(arg0 [32]byte) (common.Address, error) {
	return _ReferralStorage.Contract.CodeOwners(&_ReferralStorage.CallOpts, arg0)
}

// CodeOwners is a free data retrieval call binding the contract method 0xc8b3c460.
//
// Solidity: function codeOwners(bytes32 ) view returns(address)
func (_ReferralStorage *ReferralStorageCallerSession) CodeOwners(arg0 [32]byte) (common.Address, error) {
	return _ReferralStorage.Contract.CodeOwners(&_ReferralStorage.CallOpts, arg0)
}

// GetTraderReferralInfo is a free data retrieval call binding the contract method 0x534ef883.
//
// Solidity: function getTraderReferralInfo(address _account) view returns(bytes32, address)
func (_ReferralStorage *ReferralStorageCaller) GetTraderReferralInfo(opts *bind.CallOpts, _account common.Address) ([32]byte, common.Address, error) {
	var out []interface{}
	err := _ReferralStorage.contract.Call(opts, &out, "getTraderReferralInfo", _account)

	if err != nil {
		return *new([32]byte), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// GetTraderReferralInfo is a free data retrieval call binding the contract method 0x534ef883.
//
// Solidity: function getTraderReferralInfo(address _account) view returns(bytes32, address)
func (_ReferralStorage *ReferralStorageSession) GetTraderReferralInfo(_account common.Address) ([32]byte, common.Address, error) {
	return _ReferralStorage.Contract.GetTraderReferralInfo(&_ReferralStorage.CallOpts, _account)
}

// GetTraderReferralInfo is a free data retrieval call binding the contract method 0x534ef883.
//
// Solidity: function getTraderReferralInfo(address _account) view returns(bytes32, address)
func (_ReferralStorage *ReferralStorageCallerSession) GetTraderReferralInfo(_account common.Address) ([32]byte, common.Address, error) {
	return _ReferralStorage.Contract.GetTraderReferralInfo(&_ReferralStorage.CallOpts, _account)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_ReferralStorage *ReferralStorageCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ReferralStorage.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_ReferralStorage *ReferralStorageSession) Gov() (common.Address, error) {
	return _ReferralStorage.Contract.Gov(&_ReferralStorage.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_ReferralStorage *ReferralStorageCallerSession) Gov() (common.Address, error) {
	return _ReferralStorage.Contract.Gov(&_ReferralStorage.CallOpts)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_ReferralStorage *ReferralStorageCaller) IsHandler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ReferralStorage.contract.Call(opts, &out, "isHandler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_ReferralStorage *ReferralStorageSession) IsHandler(arg0 common.Address) (bool, error) {
	return _ReferralStorage.Contract.IsHandler(&_ReferralStorage.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_ReferralStorage *ReferralStorageCallerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _ReferralStorage.Contract.IsHandler(&_ReferralStorage.CallOpts, arg0)
}

// ReferrerDiscountShares is a free data retrieval call binding the contract method 0x71a6a790.
//
// Solidity: function referrerDiscountShares(address ) view returns(uint256)
func (_ReferralStorage *ReferralStorageCaller) ReferrerDiscountShares(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ReferralStorage.contract.Call(opts, &out, "referrerDiscountShares", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReferrerDiscountShares is a free data retrieval call binding the contract method 0x71a6a790.
//
// Solidity: function referrerDiscountShares(address ) view returns(uint256)
func (_ReferralStorage *ReferralStorageSession) ReferrerDiscountShares(arg0 common.Address) (*big.Int, error) {
	return _ReferralStorage.Contract.ReferrerDiscountShares(&_ReferralStorage.CallOpts, arg0)
}

// ReferrerDiscountShares is a free data retrieval call binding the contract method 0x71a6a790.
//
// Solidity: function referrerDiscountShares(address ) view returns(uint256)
func (_ReferralStorage *ReferralStorageCallerSession) ReferrerDiscountShares(arg0 common.Address) (*big.Int, error) {
	return _ReferralStorage.Contract.ReferrerDiscountShares(&_ReferralStorage.CallOpts, arg0)
}

// ReferrerTiers is a free data retrieval call binding the contract method 0x1582a018.
//
// Solidity: function referrerTiers(address ) view returns(uint256)
func (_ReferralStorage *ReferralStorageCaller) ReferrerTiers(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ReferralStorage.contract.Call(opts, &out, "referrerTiers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReferrerTiers is a free data retrieval call binding the contract method 0x1582a018.
//
// Solidity: function referrerTiers(address ) view returns(uint256)
func (_ReferralStorage *ReferralStorageSession) ReferrerTiers(arg0 common.Address) (*big.Int, error) {
	return _ReferralStorage.Contract.ReferrerTiers(&_ReferralStorage.CallOpts, arg0)
}

// ReferrerTiers is a free data retrieval call binding the contract method 0x1582a018.
//
// Solidity: function referrerTiers(address ) view returns(uint256)
func (_ReferralStorage *ReferralStorageCallerSession) ReferrerTiers(arg0 common.Address) (*big.Int, error) {
	return _ReferralStorage.Contract.ReferrerTiers(&_ReferralStorage.CallOpts, arg0)
}

// Tiers is a free data retrieval call binding the contract method 0x039af9eb.
//
// Solidity: function tiers(uint256 ) view returns(uint256 totalRebate, uint256 discountShare)
func (_ReferralStorage *ReferralStorageCaller) Tiers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TotalRebate   *big.Int
	DiscountShare *big.Int
}, error) {
	var out []interface{}
	err := _ReferralStorage.contract.Call(opts, &out, "tiers", arg0)

	outstruct := new(struct {
		TotalRebate   *big.Int
		DiscountShare *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalRebate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.DiscountShare = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Tiers is a free data retrieval call binding the contract method 0x039af9eb.
//
// Solidity: function tiers(uint256 ) view returns(uint256 totalRebate, uint256 discountShare)
func (_ReferralStorage *ReferralStorageSession) Tiers(arg0 *big.Int) (struct {
	TotalRebate   *big.Int
	DiscountShare *big.Int
}, error) {
	return _ReferralStorage.Contract.Tiers(&_ReferralStorage.CallOpts, arg0)
}

// Tiers is a free data retrieval call binding the contract method 0x039af9eb.
//
// Solidity: function tiers(uint256 ) view returns(uint256 totalRebate, uint256 discountShare)
func (_ReferralStorage *ReferralStorageCallerSession) Tiers(arg0 *big.Int) (struct {
	TotalRebate   *big.Int
	DiscountShare *big.Int
}, error) {
	return _ReferralStorage.Contract.Tiers(&_ReferralStorage.CallOpts, arg0)
}

// TraderReferralCodes is a free data retrieval call binding the contract method 0x85725b58.
//
// Solidity: function traderReferralCodes(address ) view returns(bytes32)
func (_ReferralStorage *ReferralStorageCaller) TraderReferralCodes(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _ReferralStorage.contract.Call(opts, &out, "traderReferralCodes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TraderReferralCodes is a free data retrieval call binding the contract method 0x85725b58.
//
// Solidity: function traderReferralCodes(address ) view returns(bytes32)
func (_ReferralStorage *ReferralStorageSession) TraderReferralCodes(arg0 common.Address) ([32]byte, error) {
	return _ReferralStorage.Contract.TraderReferralCodes(&_ReferralStorage.CallOpts, arg0)
}

// TraderReferralCodes is a free data retrieval call binding the contract method 0x85725b58.
//
// Solidity: function traderReferralCodes(address ) view returns(bytes32)
func (_ReferralStorage *ReferralStorageCallerSession) TraderReferralCodes(arg0 common.Address) ([32]byte, error) {
	return _ReferralStorage.Contract.TraderReferralCodes(&_ReferralStorage.CallOpts, arg0)
}

// GovSetCodeOwner is a paid mutator transaction binding the contract method 0xdfcfa250.
//
// Solidity: function govSetCodeOwner(bytes32 _code, address _newAccount) returns()
func (_ReferralStorage *ReferralStorageTransactor) GovSetCodeOwner(opts *bind.TransactOpts, _code [32]byte, _newAccount common.Address) (*types.Transaction, error) {
	return _ReferralStorage.contract.Transact(opts, "govSetCodeOwner", _code, _newAccount)
}

// GovSetCodeOwner is a paid mutator transaction binding the contract method 0xdfcfa250.
//
// Solidity: function govSetCodeOwner(bytes32 _code, address _newAccount) returns()
func (_ReferralStorage *ReferralStorageSession) GovSetCodeOwner(_code [32]byte, _newAccount common.Address) (*types.Transaction, error) {
	return _ReferralStorage.Contract.GovSetCodeOwner(&_ReferralStorage.TransactOpts, _code, _newAccount)
}

// GovSetCodeOwner is a paid mutator transaction binding the contract method 0xdfcfa250.
//
// Solidity: function govSetCodeOwner(bytes32 _code, address _newAccount) returns()
func (_ReferralStorage *ReferralStorageTransactorSession) GovSetCodeOwner(_code [32]byte, _newAccount common.Address) (*types.Transaction, error) {
	return _ReferralStorage.Contract.GovSetCodeOwner(&_ReferralStorage.TransactOpts, _code, _newAccount)
}

// RegisterCode is a paid mutator transaction binding the contract method 0x36def2c8.
//
// Solidity: function registerCode(bytes32 _code) returns()
func (_ReferralStorage *ReferralStorageTransactor) RegisterCode(opts *bind.TransactOpts, _code [32]byte) (*types.Transaction, error) {
	return _ReferralStorage.contract.Transact(opts, "registerCode", _code)
}

// RegisterCode is a paid mutator transaction binding the contract method 0x36def2c8.
//
// Solidity: function registerCode(bytes32 _code) returns()
func (_ReferralStorage *ReferralStorageSession) RegisterCode(_code [32]byte) (*types.Transaction, error) {
	return _ReferralStorage.Contract.RegisterCode(&_ReferralStorage.TransactOpts, _code)
}

// RegisterCode is a paid mutator transaction binding the contract method 0x36def2c8.
//
// Solidity: function registerCode(bytes32 _code) returns()
func (_ReferralStorage *ReferralStorageTransactorSession) RegisterCode(_code [32]byte) (*types.Transaction, error) {
	return _ReferralStorage.Contract.RegisterCode(&_ReferralStorage.TransactOpts, _code)
}

// SetCodeOwner is a paid mutator transaction binding the contract method 0xed843134.
//
// Solidity: function setCodeOwner(bytes32 _code, address _newAccount) returns()
func (_ReferralStorage *ReferralStorageTransactor) SetCodeOwner(opts *bind.TransactOpts, _code [32]byte, _newAccount common.Address) (*types.Transaction, error) {
	return _ReferralStorage.contract.Transact(opts, "setCodeOwner", _code, _newAccount)
}

// SetCodeOwner is a paid mutator transaction binding the contract method 0xed843134.
//
// Solidity: function setCodeOwner(bytes32 _code, address _newAccount) returns()
func (_ReferralStorage *ReferralStorageSession) SetCodeOwner(_code [32]byte, _newAccount common.Address) (*types.Transaction, error) {
	return _ReferralStorage.Contract.SetCodeOwner(&_ReferralStorage.TransactOpts, _code, _newAccount)
}

// SetCodeOwner is a paid mutator transaction binding the contract method 0xed843134.
//
// Solidity: function setCodeOwner(bytes32 _code, address _newAccount) returns()
func (_ReferralStorage *ReferralStorageTransactorSession) SetCodeOwner(_code [32]byte, _newAccount common.Address) (*types.Transaction, error) {
	return _ReferralStorage.Contract.SetCodeOwner(&_ReferralStorage.TransactOpts, _code, _newAccount)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_ReferralStorage *ReferralStorageTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _ReferralStorage.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_ReferralStorage *ReferralStorageSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _ReferralStorage.Contract.SetGov(&_ReferralStorage.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_ReferralStorage *ReferralStorageTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _ReferralStorage.Contract.SetGov(&_ReferralStorage.TransactOpts, _gov)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_ReferralStorage *ReferralStorageTransactor) SetHandler(opts *bind.TransactOpts, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _ReferralStorage.contract.Transact(opts, "setHandler", _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_ReferralStorage *ReferralStorageSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _ReferralStorage.Contract.SetHandler(&_ReferralStorage.TransactOpts, _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x9cb7de4b.
//
// Solidity: function setHandler(address _handler, bool _isActive) returns()
func (_ReferralStorage *ReferralStorageTransactorSession) SetHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _ReferralStorage.Contract.SetHandler(&_ReferralStorage.TransactOpts, _handler, _isActive)
}

// SetReferrerDiscountShare is a paid mutator transaction binding the contract method 0x9c8e2de9.
//
// Solidity: function setReferrerDiscountShare(uint256 _discountShare) returns()
func (_ReferralStorage *ReferralStorageTransactor) SetReferrerDiscountShare(opts *bind.TransactOpts, _discountShare *big.Int) (*types.Transaction, error) {
	return _ReferralStorage.contract.Transact(opts, "setReferrerDiscountShare", _discountShare)
}

// SetReferrerDiscountShare is a paid mutator transaction binding the contract method 0x9c8e2de9.
//
// Solidity: function setReferrerDiscountShare(uint256 _discountShare) returns()
func (_ReferralStorage *ReferralStorageSession) SetReferrerDiscountShare(_discountShare *big.Int) (*types.Transaction, error) {
	return _ReferralStorage.Contract.SetReferrerDiscountShare(&_ReferralStorage.TransactOpts, _discountShare)
}

// SetReferrerDiscountShare is a paid mutator transaction binding the contract method 0x9c8e2de9.
//
// Solidity: function setReferrerDiscountShare(uint256 _discountShare) returns()
func (_ReferralStorage *ReferralStorageTransactorSession) SetReferrerDiscountShare(_discountShare *big.Int) (*types.Transaction, error) {
	return _ReferralStorage.Contract.SetReferrerDiscountShare(&_ReferralStorage.TransactOpts, _discountShare)
}

// SetReferrerTier is a paid mutator transaction binding the contract method 0x3fb8b323.
//
// Solidity: function setReferrerTier(address _referrer, uint256 _tierId) returns()
func (_ReferralStorage *ReferralStorageTransactor) SetReferrerTier(opts *bind.TransactOpts, _referrer common.Address, _tierId *big.Int) (*types.Transaction, error) {
	return _ReferralStorage.contract.Transact(opts, "setReferrerTier", _referrer, _tierId)
}

// SetReferrerTier is a paid mutator transaction binding the contract method 0x3fb8b323.
//
// Solidity: function setReferrerTier(address _referrer, uint256 _tierId) returns()
func (_ReferralStorage *ReferralStorageSession) SetReferrerTier(_referrer common.Address, _tierId *big.Int) (*types.Transaction, error) {
	return _ReferralStorage.Contract.SetReferrerTier(&_ReferralStorage.TransactOpts, _referrer, _tierId)
}

// SetReferrerTier is a paid mutator transaction binding the contract method 0x3fb8b323.
//
// Solidity: function setReferrerTier(address _referrer, uint256 _tierId) returns()
func (_ReferralStorage *ReferralStorageTransactorSession) SetReferrerTier(_referrer common.Address, _tierId *big.Int) (*types.Transaction, error) {
	return _ReferralStorage.Contract.SetReferrerTier(&_ReferralStorage.TransactOpts, _referrer, _tierId)
}

// SetTier is a paid mutator transaction binding the contract method 0x836a0187.
//
// Solidity: function setTier(uint256 _tierId, uint256 _totalRebate, uint256 _discountShare) returns()
func (_ReferralStorage *ReferralStorageTransactor) SetTier(opts *bind.TransactOpts, _tierId *big.Int, _totalRebate *big.Int, _discountShare *big.Int) (*types.Transaction, error) {
	return _ReferralStorage.contract.Transact(opts, "setTier", _tierId, _totalRebate, _discountShare)
}

// SetTier is a paid mutator transaction binding the contract method 0x836a0187.
//
// Solidity: function setTier(uint256 _tierId, uint256 _totalRebate, uint256 _discountShare) returns()
func (_ReferralStorage *ReferralStorageSession) SetTier(_tierId *big.Int, _totalRebate *big.Int, _discountShare *big.Int) (*types.Transaction, error) {
	return _ReferralStorage.Contract.SetTier(&_ReferralStorage.TransactOpts, _tierId, _totalRebate, _discountShare)
}

// SetTier is a paid mutator transaction binding the contract method 0x836a0187.
//
// Solidity: function setTier(uint256 _tierId, uint256 _totalRebate, uint256 _discountShare) returns()
func (_ReferralStorage *ReferralStorageTransactorSession) SetTier(_tierId *big.Int, _totalRebate *big.Int, _discountShare *big.Int) (*types.Transaction, error) {
	return _ReferralStorage.Contract.SetTier(&_ReferralStorage.TransactOpts, _tierId, _totalRebate, _discountShare)
}

// SetTraderReferralCode is a paid mutator transaction binding the contract method 0x56b4b2ad.
//
// Solidity: function setTraderReferralCode(address _account, bytes32 _code) returns()
func (_ReferralStorage *ReferralStorageTransactor) SetTraderReferralCode(opts *bind.TransactOpts, _account common.Address, _code [32]byte) (*types.Transaction, error) {
	return _ReferralStorage.contract.Transact(opts, "setTraderReferralCode", _account, _code)
}

// SetTraderReferralCode is a paid mutator transaction binding the contract method 0x56b4b2ad.
//
// Solidity: function setTraderReferralCode(address _account, bytes32 _code) returns()
func (_ReferralStorage *ReferralStorageSession) SetTraderReferralCode(_account common.Address, _code [32]byte) (*types.Transaction, error) {
	return _ReferralStorage.Contract.SetTraderReferralCode(&_ReferralStorage.TransactOpts, _account, _code)
}

// SetTraderReferralCode is a paid mutator transaction binding the contract method 0x56b4b2ad.
//
// Solidity: function setTraderReferralCode(address _account, bytes32 _code) returns()
func (_ReferralStorage *ReferralStorageTransactorSession) SetTraderReferralCode(_account common.Address, _code [32]byte) (*types.Transaction, error) {
	return _ReferralStorage.Contract.SetTraderReferralCode(&_ReferralStorage.TransactOpts, _account, _code)
}

// SetTraderReferralCodeByUser is a paid mutator transaction binding the contract method 0xe1e01bf3.
//
// Solidity: function setTraderReferralCodeByUser(bytes32 _code) returns()
func (_ReferralStorage *ReferralStorageTransactor) SetTraderReferralCodeByUser(opts *bind.TransactOpts, _code [32]byte) (*types.Transaction, error) {
	return _ReferralStorage.contract.Transact(opts, "setTraderReferralCodeByUser", _code)
}

// SetTraderReferralCodeByUser is a paid mutator transaction binding the contract method 0xe1e01bf3.
//
// Solidity: function setTraderReferralCodeByUser(bytes32 _code) returns()
func (_ReferralStorage *ReferralStorageSession) SetTraderReferralCodeByUser(_code [32]byte) (*types.Transaction, error) {
	return _ReferralStorage.Contract.SetTraderReferralCodeByUser(&_ReferralStorage.TransactOpts, _code)
}

// SetTraderReferralCodeByUser is a paid mutator transaction binding the contract method 0xe1e01bf3.
//
// Solidity: function setTraderReferralCodeByUser(bytes32 _code) returns()
func (_ReferralStorage *ReferralStorageTransactorSession) SetTraderReferralCodeByUser(_code [32]byte) (*types.Transaction, error) {
	return _ReferralStorage.Contract.SetTraderReferralCodeByUser(&_ReferralStorage.TransactOpts, _code)
}

// ReferralStorageGovSetCodeOwnerIterator is returned from FilterGovSetCodeOwner and is used to iterate over the raw logs and unpacked data for GovSetCodeOwner events raised by the ReferralStorage contract.
type ReferralStorageGovSetCodeOwnerIterator struct {
	Event *ReferralStorageGovSetCodeOwner // Event containing the contract specifics and raw log

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
func (it *ReferralStorageGovSetCodeOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralStorageGovSetCodeOwner)
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
		it.Event = new(ReferralStorageGovSetCodeOwner)
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
func (it *ReferralStorageGovSetCodeOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralStorageGovSetCodeOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralStorageGovSetCodeOwner represents a GovSetCodeOwner event raised by the ReferralStorage contract.
type ReferralStorageGovSetCodeOwner struct {
	Code       [32]byte
	NewAccount common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGovSetCodeOwner is a free log retrieval operation binding the contract event 0x6431f88c655dd0e2b8d09b6405c007c515c66d67f2998e69c902873a8c8f3e97.
//
// Solidity: event GovSetCodeOwner(bytes32 code, address newAccount)
func (_ReferralStorage *ReferralStorageFilterer) FilterGovSetCodeOwner(opts *bind.FilterOpts) (*ReferralStorageGovSetCodeOwnerIterator, error) {

	logs, sub, err := _ReferralStorage.contract.FilterLogs(opts, "GovSetCodeOwner")
	if err != nil {
		return nil, err
	}
	return &ReferralStorageGovSetCodeOwnerIterator{contract: _ReferralStorage.contract, event: "GovSetCodeOwner", logs: logs, sub: sub}, nil
}

// WatchGovSetCodeOwner is a free log subscription operation binding the contract event 0x6431f88c655dd0e2b8d09b6405c007c515c66d67f2998e69c902873a8c8f3e97.
//
// Solidity: event GovSetCodeOwner(bytes32 code, address newAccount)
func (_ReferralStorage *ReferralStorageFilterer) WatchGovSetCodeOwner(opts *bind.WatchOpts, sink chan<- *ReferralStorageGovSetCodeOwner) (event.Subscription, error) {

	logs, sub, err := _ReferralStorage.contract.WatchLogs(opts, "GovSetCodeOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralStorageGovSetCodeOwner)
				if err := _ReferralStorage.contract.UnpackLog(event, "GovSetCodeOwner", log); err != nil {
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

// ParseGovSetCodeOwner is a log parse operation binding the contract event 0x6431f88c655dd0e2b8d09b6405c007c515c66d67f2998e69c902873a8c8f3e97.
//
// Solidity: event GovSetCodeOwner(bytes32 code, address newAccount)
func (_ReferralStorage *ReferralStorageFilterer) ParseGovSetCodeOwner(log types.Log) (*ReferralStorageGovSetCodeOwner, error) {
	event := new(ReferralStorageGovSetCodeOwner)
	if err := _ReferralStorage.contract.UnpackLog(event, "GovSetCodeOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralStorageRegisterCodeIterator is returned from FilterRegisterCode and is used to iterate over the raw logs and unpacked data for RegisterCode events raised by the ReferralStorage contract.
type ReferralStorageRegisterCodeIterator struct {
	Event *ReferralStorageRegisterCode // Event containing the contract specifics and raw log

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
func (it *ReferralStorageRegisterCodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralStorageRegisterCode)
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
		it.Event = new(ReferralStorageRegisterCode)
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
func (it *ReferralStorageRegisterCodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralStorageRegisterCodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralStorageRegisterCode represents a RegisterCode event raised by the ReferralStorage contract.
type ReferralStorageRegisterCode struct {
	Account common.Address
	Code    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRegisterCode is a free log retrieval operation binding the contract event 0x04f82286a2a3b2ee5c8555de8304dfe2ea70991613213184b73a9e408d2d8029.
//
// Solidity: event RegisterCode(address account, bytes32 code)
func (_ReferralStorage *ReferralStorageFilterer) FilterRegisterCode(opts *bind.FilterOpts) (*ReferralStorageRegisterCodeIterator, error) {

	logs, sub, err := _ReferralStorage.contract.FilterLogs(opts, "RegisterCode")
	if err != nil {
		return nil, err
	}
	return &ReferralStorageRegisterCodeIterator{contract: _ReferralStorage.contract, event: "RegisterCode", logs: logs, sub: sub}, nil
}

// WatchRegisterCode is a free log subscription operation binding the contract event 0x04f82286a2a3b2ee5c8555de8304dfe2ea70991613213184b73a9e408d2d8029.
//
// Solidity: event RegisterCode(address account, bytes32 code)
func (_ReferralStorage *ReferralStorageFilterer) WatchRegisterCode(opts *bind.WatchOpts, sink chan<- *ReferralStorageRegisterCode) (event.Subscription, error) {

	logs, sub, err := _ReferralStorage.contract.WatchLogs(opts, "RegisterCode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralStorageRegisterCode)
				if err := _ReferralStorage.contract.UnpackLog(event, "RegisterCode", log); err != nil {
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

// ParseRegisterCode is a log parse operation binding the contract event 0x04f82286a2a3b2ee5c8555de8304dfe2ea70991613213184b73a9e408d2d8029.
//
// Solidity: event RegisterCode(address account, bytes32 code)
func (_ReferralStorage *ReferralStorageFilterer) ParseRegisterCode(log types.Log) (*ReferralStorageRegisterCode, error) {
	event := new(ReferralStorageRegisterCode)
	if err := _ReferralStorage.contract.UnpackLog(event, "RegisterCode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralStorageSetCodeOwnerIterator is returned from FilterSetCodeOwner and is used to iterate over the raw logs and unpacked data for SetCodeOwner events raised by the ReferralStorage contract.
type ReferralStorageSetCodeOwnerIterator struct {
	Event *ReferralStorageSetCodeOwner // Event containing the contract specifics and raw log

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
func (it *ReferralStorageSetCodeOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralStorageSetCodeOwner)
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
		it.Event = new(ReferralStorageSetCodeOwner)
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
func (it *ReferralStorageSetCodeOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralStorageSetCodeOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralStorageSetCodeOwner represents a SetCodeOwner event raised by the ReferralStorage contract.
type ReferralStorageSetCodeOwner struct {
	Account    common.Address
	NewAccount common.Address
	Code       [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetCodeOwner is a free log retrieval operation binding the contract event 0x5640856798d41ce9ca0a109b54c20a06eb99ba9c36ab4547115dafb8473cf397.
//
// Solidity: event SetCodeOwner(address account, address newAccount, bytes32 code)
func (_ReferralStorage *ReferralStorageFilterer) FilterSetCodeOwner(opts *bind.FilterOpts) (*ReferralStorageSetCodeOwnerIterator, error) {

	logs, sub, err := _ReferralStorage.contract.FilterLogs(opts, "SetCodeOwner")
	if err != nil {
		return nil, err
	}
	return &ReferralStorageSetCodeOwnerIterator{contract: _ReferralStorage.contract, event: "SetCodeOwner", logs: logs, sub: sub}, nil
}

// WatchSetCodeOwner is a free log subscription operation binding the contract event 0x5640856798d41ce9ca0a109b54c20a06eb99ba9c36ab4547115dafb8473cf397.
//
// Solidity: event SetCodeOwner(address account, address newAccount, bytes32 code)
func (_ReferralStorage *ReferralStorageFilterer) WatchSetCodeOwner(opts *bind.WatchOpts, sink chan<- *ReferralStorageSetCodeOwner) (event.Subscription, error) {

	logs, sub, err := _ReferralStorage.contract.WatchLogs(opts, "SetCodeOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralStorageSetCodeOwner)
				if err := _ReferralStorage.contract.UnpackLog(event, "SetCodeOwner", log); err != nil {
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

// ParseSetCodeOwner is a log parse operation binding the contract event 0x5640856798d41ce9ca0a109b54c20a06eb99ba9c36ab4547115dafb8473cf397.
//
// Solidity: event SetCodeOwner(address account, address newAccount, bytes32 code)
func (_ReferralStorage *ReferralStorageFilterer) ParseSetCodeOwner(log types.Log) (*ReferralStorageSetCodeOwner, error) {
	event := new(ReferralStorageSetCodeOwner)
	if err := _ReferralStorage.contract.UnpackLog(event, "SetCodeOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralStorageSetHandlerIterator is returned from FilterSetHandler and is used to iterate over the raw logs and unpacked data for SetHandler events raised by the ReferralStorage contract.
type ReferralStorageSetHandlerIterator struct {
	Event *ReferralStorageSetHandler // Event containing the contract specifics and raw log

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
func (it *ReferralStorageSetHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralStorageSetHandler)
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
		it.Event = new(ReferralStorageSetHandler)
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
func (it *ReferralStorageSetHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralStorageSetHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralStorageSetHandler represents a SetHandler event raised by the ReferralStorage contract.
type ReferralStorageSetHandler struct {
	Handler  common.Address
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetHandler is a free log retrieval operation binding the contract event 0xd373464a39404e5f98fdb4b152b8ba9a094561e97e5c4b4ea11eb18cd9e6695e.
//
// Solidity: event SetHandler(address handler, bool isActive)
func (_ReferralStorage *ReferralStorageFilterer) FilterSetHandler(opts *bind.FilterOpts) (*ReferralStorageSetHandlerIterator, error) {

	logs, sub, err := _ReferralStorage.contract.FilterLogs(opts, "SetHandler")
	if err != nil {
		return nil, err
	}
	return &ReferralStorageSetHandlerIterator{contract: _ReferralStorage.contract, event: "SetHandler", logs: logs, sub: sub}, nil
}

// WatchSetHandler is a free log subscription operation binding the contract event 0xd373464a39404e5f98fdb4b152b8ba9a094561e97e5c4b4ea11eb18cd9e6695e.
//
// Solidity: event SetHandler(address handler, bool isActive)
func (_ReferralStorage *ReferralStorageFilterer) WatchSetHandler(opts *bind.WatchOpts, sink chan<- *ReferralStorageSetHandler) (event.Subscription, error) {

	logs, sub, err := _ReferralStorage.contract.WatchLogs(opts, "SetHandler")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralStorageSetHandler)
				if err := _ReferralStorage.contract.UnpackLog(event, "SetHandler", log); err != nil {
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

// ParseSetHandler is a log parse operation binding the contract event 0xd373464a39404e5f98fdb4b152b8ba9a094561e97e5c4b4ea11eb18cd9e6695e.
//
// Solidity: event SetHandler(address handler, bool isActive)
func (_ReferralStorage *ReferralStorageFilterer) ParseSetHandler(log types.Log) (*ReferralStorageSetHandler, error) {
	event := new(ReferralStorageSetHandler)
	if err := _ReferralStorage.contract.UnpackLog(event, "SetHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralStorageSetReferrerDiscountShareIterator is returned from FilterSetReferrerDiscountShare and is used to iterate over the raw logs and unpacked data for SetReferrerDiscountShare events raised by the ReferralStorage contract.
type ReferralStorageSetReferrerDiscountShareIterator struct {
	Event *ReferralStorageSetReferrerDiscountShare // Event containing the contract specifics and raw log

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
func (it *ReferralStorageSetReferrerDiscountShareIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralStorageSetReferrerDiscountShare)
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
		it.Event = new(ReferralStorageSetReferrerDiscountShare)
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
func (it *ReferralStorageSetReferrerDiscountShareIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralStorageSetReferrerDiscountShareIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralStorageSetReferrerDiscountShare represents a SetReferrerDiscountShare event raised by the ReferralStorage contract.
type ReferralStorageSetReferrerDiscountShare struct {
	Referrer      common.Address
	DiscountShare *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetReferrerDiscountShare is a free log retrieval operation binding the contract event 0xbd224f3917462b0fa80805fe3ec29be3a37f664955aad6224e5ece036224c429.
//
// Solidity: event SetReferrerDiscountShare(address referrer, uint256 discountShare)
func (_ReferralStorage *ReferralStorageFilterer) FilterSetReferrerDiscountShare(opts *bind.FilterOpts) (*ReferralStorageSetReferrerDiscountShareIterator, error) {

	logs, sub, err := _ReferralStorage.contract.FilterLogs(opts, "SetReferrerDiscountShare")
	if err != nil {
		return nil, err
	}
	return &ReferralStorageSetReferrerDiscountShareIterator{contract: _ReferralStorage.contract, event: "SetReferrerDiscountShare", logs: logs, sub: sub}, nil
}

// WatchSetReferrerDiscountShare is a free log subscription operation binding the contract event 0xbd224f3917462b0fa80805fe3ec29be3a37f664955aad6224e5ece036224c429.
//
// Solidity: event SetReferrerDiscountShare(address referrer, uint256 discountShare)
func (_ReferralStorage *ReferralStorageFilterer) WatchSetReferrerDiscountShare(opts *bind.WatchOpts, sink chan<- *ReferralStorageSetReferrerDiscountShare) (event.Subscription, error) {

	logs, sub, err := _ReferralStorage.contract.WatchLogs(opts, "SetReferrerDiscountShare")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralStorageSetReferrerDiscountShare)
				if err := _ReferralStorage.contract.UnpackLog(event, "SetReferrerDiscountShare", log); err != nil {
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

// ParseSetReferrerDiscountShare is a log parse operation binding the contract event 0xbd224f3917462b0fa80805fe3ec29be3a37f664955aad6224e5ece036224c429.
//
// Solidity: event SetReferrerDiscountShare(address referrer, uint256 discountShare)
func (_ReferralStorage *ReferralStorageFilterer) ParseSetReferrerDiscountShare(log types.Log) (*ReferralStorageSetReferrerDiscountShare, error) {
	event := new(ReferralStorageSetReferrerDiscountShare)
	if err := _ReferralStorage.contract.UnpackLog(event, "SetReferrerDiscountShare", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralStorageSetReferrerTierIterator is returned from FilterSetReferrerTier and is used to iterate over the raw logs and unpacked data for SetReferrerTier events raised by the ReferralStorage contract.
type ReferralStorageSetReferrerTierIterator struct {
	Event *ReferralStorageSetReferrerTier // Event containing the contract specifics and raw log

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
func (it *ReferralStorageSetReferrerTierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralStorageSetReferrerTier)
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
		it.Event = new(ReferralStorageSetReferrerTier)
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
func (it *ReferralStorageSetReferrerTierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralStorageSetReferrerTierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralStorageSetReferrerTier represents a SetReferrerTier event raised by the ReferralStorage contract.
type ReferralStorageSetReferrerTier struct {
	Referrer common.Address
	TierId   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetReferrerTier is a free log retrieval operation binding the contract event 0x7696855cdbb94bb5a44cb0a95caff6f29173aa1cbfba193834b12b90827ce2bc.
//
// Solidity: event SetReferrerTier(address referrer, uint256 tierId)
func (_ReferralStorage *ReferralStorageFilterer) FilterSetReferrerTier(opts *bind.FilterOpts) (*ReferralStorageSetReferrerTierIterator, error) {

	logs, sub, err := _ReferralStorage.contract.FilterLogs(opts, "SetReferrerTier")
	if err != nil {
		return nil, err
	}
	return &ReferralStorageSetReferrerTierIterator{contract: _ReferralStorage.contract, event: "SetReferrerTier", logs: logs, sub: sub}, nil
}

// WatchSetReferrerTier is a free log subscription operation binding the contract event 0x7696855cdbb94bb5a44cb0a95caff6f29173aa1cbfba193834b12b90827ce2bc.
//
// Solidity: event SetReferrerTier(address referrer, uint256 tierId)
func (_ReferralStorage *ReferralStorageFilterer) WatchSetReferrerTier(opts *bind.WatchOpts, sink chan<- *ReferralStorageSetReferrerTier) (event.Subscription, error) {

	logs, sub, err := _ReferralStorage.contract.WatchLogs(opts, "SetReferrerTier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralStorageSetReferrerTier)
				if err := _ReferralStorage.contract.UnpackLog(event, "SetReferrerTier", log); err != nil {
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

// ParseSetReferrerTier is a log parse operation binding the contract event 0x7696855cdbb94bb5a44cb0a95caff6f29173aa1cbfba193834b12b90827ce2bc.
//
// Solidity: event SetReferrerTier(address referrer, uint256 tierId)
func (_ReferralStorage *ReferralStorageFilterer) ParseSetReferrerTier(log types.Log) (*ReferralStorageSetReferrerTier, error) {
	event := new(ReferralStorageSetReferrerTier)
	if err := _ReferralStorage.contract.UnpackLog(event, "SetReferrerTier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralStorageSetTierIterator is returned from FilterSetTier and is used to iterate over the raw logs and unpacked data for SetTier events raised by the ReferralStorage contract.
type ReferralStorageSetTierIterator struct {
	Event *ReferralStorageSetTier // Event containing the contract specifics and raw log

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
func (it *ReferralStorageSetTierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralStorageSetTier)
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
		it.Event = new(ReferralStorageSetTier)
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
func (it *ReferralStorageSetTierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralStorageSetTierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralStorageSetTier represents a SetTier event raised by the ReferralStorage contract.
type ReferralStorageSetTier struct {
	TierId        *big.Int
	TotalRebate   *big.Int
	DiscountShare *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetTier is a free log retrieval operation binding the contract event 0x4ebd23a492b2bd79586cb57cae40d7793408265f20320f68b478b971e696f4c7.
//
// Solidity: event SetTier(uint256 tierId, uint256 totalRebate, uint256 discountShare)
func (_ReferralStorage *ReferralStorageFilterer) FilterSetTier(opts *bind.FilterOpts) (*ReferralStorageSetTierIterator, error) {

	logs, sub, err := _ReferralStorage.contract.FilterLogs(opts, "SetTier")
	if err != nil {
		return nil, err
	}
	return &ReferralStorageSetTierIterator{contract: _ReferralStorage.contract, event: "SetTier", logs: logs, sub: sub}, nil
}

// WatchSetTier is a free log subscription operation binding the contract event 0x4ebd23a492b2bd79586cb57cae40d7793408265f20320f68b478b971e696f4c7.
//
// Solidity: event SetTier(uint256 tierId, uint256 totalRebate, uint256 discountShare)
func (_ReferralStorage *ReferralStorageFilterer) WatchSetTier(opts *bind.WatchOpts, sink chan<- *ReferralStorageSetTier) (event.Subscription, error) {

	logs, sub, err := _ReferralStorage.contract.WatchLogs(opts, "SetTier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralStorageSetTier)
				if err := _ReferralStorage.contract.UnpackLog(event, "SetTier", log); err != nil {
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

// ParseSetTier is a log parse operation binding the contract event 0x4ebd23a492b2bd79586cb57cae40d7793408265f20320f68b478b971e696f4c7.
//
// Solidity: event SetTier(uint256 tierId, uint256 totalRebate, uint256 discountShare)
func (_ReferralStorage *ReferralStorageFilterer) ParseSetTier(log types.Log) (*ReferralStorageSetTier, error) {
	event := new(ReferralStorageSetTier)
	if err := _ReferralStorage.contract.UnpackLog(event, "SetTier", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReferralStorageSetTraderReferralCodeIterator is returned from FilterSetTraderReferralCode and is used to iterate over the raw logs and unpacked data for SetTraderReferralCode events raised by the ReferralStorage contract.
type ReferralStorageSetTraderReferralCodeIterator struct {
	Event *ReferralStorageSetTraderReferralCode // Event containing the contract specifics and raw log

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
func (it *ReferralStorageSetTraderReferralCodeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReferralStorageSetTraderReferralCode)
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
		it.Event = new(ReferralStorageSetTraderReferralCode)
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
func (it *ReferralStorageSetTraderReferralCodeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReferralStorageSetTraderReferralCodeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReferralStorageSetTraderReferralCode represents a SetTraderReferralCode event raised by the ReferralStorage contract.
type ReferralStorageSetTraderReferralCode struct {
	Account common.Address
	Code    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetTraderReferralCode is a free log retrieval operation binding the contract event 0x43825f14567dda057e821be2e51a5aa79aa51f3907a647e3ed2bd486a01050f1.
//
// Solidity: event SetTraderReferralCode(address account, bytes32 code)
func (_ReferralStorage *ReferralStorageFilterer) FilterSetTraderReferralCode(opts *bind.FilterOpts) (*ReferralStorageSetTraderReferralCodeIterator, error) {

	logs, sub, err := _ReferralStorage.contract.FilterLogs(opts, "SetTraderReferralCode")
	if err != nil {
		return nil, err
	}
	return &ReferralStorageSetTraderReferralCodeIterator{contract: _ReferralStorage.contract, event: "SetTraderReferralCode", logs: logs, sub: sub}, nil
}

// WatchSetTraderReferralCode is a free log subscription operation binding the contract event 0x43825f14567dda057e821be2e51a5aa79aa51f3907a647e3ed2bd486a01050f1.
//
// Solidity: event SetTraderReferralCode(address account, bytes32 code)
func (_ReferralStorage *ReferralStorageFilterer) WatchSetTraderReferralCode(opts *bind.WatchOpts, sink chan<- *ReferralStorageSetTraderReferralCode) (event.Subscription, error) {

	logs, sub, err := _ReferralStorage.contract.WatchLogs(opts, "SetTraderReferralCode")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReferralStorageSetTraderReferralCode)
				if err := _ReferralStorage.contract.UnpackLog(event, "SetTraderReferralCode", log); err != nil {
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

// ParseSetTraderReferralCode is a log parse operation binding the contract event 0x43825f14567dda057e821be2e51a5aa79aa51f3907a647e3ed2bd486a01050f1.
//
// Solidity: event SetTraderReferralCode(address account, bytes32 code)
func (_ReferralStorage *ReferralStorageFilterer) ParseSetTraderReferralCode(log types.Log) (*ReferralStorageSetTraderReferralCode, error) {
	event := new(ReferralStorageSetTraderReferralCode)
	if err := _ReferralStorage.contract.UnpackLog(event, "SetTraderReferralCode", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

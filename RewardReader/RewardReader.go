// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RewardReader

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

// RewardReaderMetaData contains all meta data concerning the RewardReader contract.
var RewardReaderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getDepositBalances\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_depositTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_rewardTrackers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakingInfo\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_rewardTrackers\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVestingInfoV2\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_vesters\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"}]",
}

// RewardReaderABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardReaderMetaData.ABI instead.
var RewardReaderABI = RewardReaderMetaData.ABI

// RewardReader is an auto generated Go binding around an Ethereum contract.
type RewardReader struct {
	RewardReaderCaller     // Read-only binding to the contract
	RewardReaderTransactor // Write-only binding to the contract
	RewardReaderFilterer   // Log filterer for contract events
}

// RewardReaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardReaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardReaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardReaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardReaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardReaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardReaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardReaderSession struct {
	Contract     *RewardReader     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardReaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardReaderCallerSession struct {
	Contract *RewardReaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RewardReaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardReaderTransactorSession struct {
	Contract     *RewardReaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RewardReaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardReaderRaw struct {
	Contract *RewardReader // Generic contract binding to access the raw methods on
}

// RewardReaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardReaderCallerRaw struct {
	Contract *RewardReaderCaller // Generic read-only contract binding to access the raw methods on
}

// RewardReaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardReaderTransactorRaw struct {
	Contract *RewardReaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardReader creates a new instance of RewardReader, bound to a specific deployed contract.
func NewRewardReader(address common.Address, backend bind.ContractBackend) (*RewardReader, error) {
	contract, err := bindRewardReader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardReader{RewardReaderCaller: RewardReaderCaller{contract: contract}, RewardReaderTransactor: RewardReaderTransactor{contract: contract}, RewardReaderFilterer: RewardReaderFilterer{contract: contract}}, nil
}

// NewRewardReaderCaller creates a new read-only instance of RewardReader, bound to a specific deployed contract.
func NewRewardReaderCaller(address common.Address, caller bind.ContractCaller) (*RewardReaderCaller, error) {
	contract, err := bindRewardReader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardReaderCaller{contract: contract}, nil
}

// NewRewardReaderTransactor creates a new write-only instance of RewardReader, bound to a specific deployed contract.
func NewRewardReaderTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardReaderTransactor, error) {
	contract, err := bindRewardReader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardReaderTransactor{contract: contract}, nil
}

// NewRewardReaderFilterer creates a new log filterer instance of RewardReader, bound to a specific deployed contract.
func NewRewardReaderFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardReaderFilterer, error) {
	contract, err := bindRewardReader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardReaderFilterer{contract: contract}, nil
}

// bindRewardReader binds a generic wrapper to an already deployed contract.
func bindRewardReader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RewardReaderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardReader *RewardReaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardReader.Contract.RewardReaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardReader *RewardReaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardReader.Contract.RewardReaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardReader *RewardReaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardReader.Contract.RewardReaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardReader *RewardReaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardReader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardReader *RewardReaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardReader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardReader *RewardReaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardReader.Contract.contract.Transact(opts, method, params...)
}

// GetDepositBalances is a free data retrieval call binding the contract method 0x575157e4.
//
// Solidity: function getDepositBalances(address _account, address[] _depositTokens, address[] _rewardTrackers) view returns(uint256[])
func (_RewardReader *RewardReaderCaller) GetDepositBalances(opts *bind.CallOpts, _account common.Address, _depositTokens []common.Address, _rewardTrackers []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _RewardReader.contract.Call(opts, &out, "getDepositBalances", _account, _depositTokens, _rewardTrackers)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetDepositBalances is a free data retrieval call binding the contract method 0x575157e4.
//
// Solidity: function getDepositBalances(address _account, address[] _depositTokens, address[] _rewardTrackers) view returns(uint256[])
func (_RewardReader *RewardReaderSession) GetDepositBalances(_account common.Address, _depositTokens []common.Address, _rewardTrackers []common.Address) ([]*big.Int, error) {
	return _RewardReader.Contract.GetDepositBalances(&_RewardReader.CallOpts, _account, _depositTokens, _rewardTrackers)
}

// GetDepositBalances is a free data retrieval call binding the contract method 0x575157e4.
//
// Solidity: function getDepositBalances(address _account, address[] _depositTokens, address[] _rewardTrackers) view returns(uint256[])
func (_RewardReader *RewardReaderCallerSession) GetDepositBalances(_account common.Address, _depositTokens []common.Address, _rewardTrackers []common.Address) ([]*big.Int, error) {
	return _RewardReader.Contract.GetDepositBalances(&_RewardReader.CallOpts, _account, _depositTokens, _rewardTrackers)
}

// GetStakingInfo is a free data retrieval call binding the contract method 0x937a0be8.
//
// Solidity: function getStakingInfo(address _account, address[] _rewardTrackers) view returns(uint256[])
func (_RewardReader *RewardReaderCaller) GetStakingInfo(opts *bind.CallOpts, _account common.Address, _rewardTrackers []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _RewardReader.contract.Call(opts, &out, "getStakingInfo", _account, _rewardTrackers)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetStakingInfo is a free data retrieval call binding the contract method 0x937a0be8.
//
// Solidity: function getStakingInfo(address _account, address[] _rewardTrackers) view returns(uint256[])
func (_RewardReader *RewardReaderSession) GetStakingInfo(_account common.Address, _rewardTrackers []common.Address) ([]*big.Int, error) {
	return _RewardReader.Contract.GetStakingInfo(&_RewardReader.CallOpts, _account, _rewardTrackers)
}

// GetStakingInfo is a free data retrieval call binding the contract method 0x937a0be8.
//
// Solidity: function getStakingInfo(address _account, address[] _rewardTrackers) view returns(uint256[])
func (_RewardReader *RewardReaderCallerSession) GetStakingInfo(_account common.Address, _rewardTrackers []common.Address) ([]*big.Int, error) {
	return _RewardReader.Contract.GetStakingInfo(&_RewardReader.CallOpts, _account, _rewardTrackers)
}

// GetVestingInfoV2 is a free data retrieval call binding the contract method 0xe1fc2a38.
//
// Solidity: function getVestingInfoV2(address _account, address[] _vesters) view returns(uint256[])
func (_RewardReader *RewardReaderCaller) GetVestingInfoV2(opts *bind.CallOpts, _account common.Address, _vesters []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _RewardReader.contract.Call(opts, &out, "getVestingInfoV2", _account, _vesters)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetVestingInfoV2 is a free data retrieval call binding the contract method 0xe1fc2a38.
//
// Solidity: function getVestingInfoV2(address _account, address[] _vesters) view returns(uint256[])
func (_RewardReader *RewardReaderSession) GetVestingInfoV2(_account common.Address, _vesters []common.Address) ([]*big.Int, error) {
	return _RewardReader.Contract.GetVestingInfoV2(&_RewardReader.CallOpts, _account, _vesters)
}

// GetVestingInfoV2 is a free data retrieval call binding the contract method 0xe1fc2a38.
//
// Solidity: function getVestingInfoV2(address _account, address[] _vesters) view returns(uint256[])
func (_RewardReader *RewardReaderCallerSession) GetVestingInfoV2(_account common.Address, _vesters []common.Address) ([]*big.Int, error) {
	return _RewardReader.Contract.GetVestingInfoV2(&_RewardReader.CallOpts, _account, _vesters)
}

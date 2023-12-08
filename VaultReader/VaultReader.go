// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VaultReader

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

// VaultReaderMetaData contains all meta data concerning the VaultReader contract.
var VaultReaderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getVaultTokenInfoV3\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_positionManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdgAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getVaultTokenInfoV4\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_positionManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdgAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"}]",
}

// VaultReaderABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultReaderMetaData.ABI instead.
var VaultReaderABI = VaultReaderMetaData.ABI

// VaultReader is an auto generated Go binding around an Ethereum contract.
type VaultReader struct {
	VaultReaderCaller     // Read-only binding to the contract
	VaultReaderTransactor // Write-only binding to the contract
	VaultReaderFilterer   // Log filterer for contract events
}

// VaultReaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultReaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultReaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultReaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultReaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultReaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultReaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultReaderSession struct {
	Contract     *VaultReader      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultReaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultReaderCallerSession struct {
	Contract *VaultReaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// VaultReaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultReaderTransactorSession struct {
	Contract     *VaultReaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VaultReaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultReaderRaw struct {
	Contract *VaultReader // Generic contract binding to access the raw methods on
}

// VaultReaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultReaderCallerRaw struct {
	Contract *VaultReaderCaller // Generic read-only contract binding to access the raw methods on
}

// VaultReaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultReaderTransactorRaw struct {
	Contract *VaultReaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultReader creates a new instance of VaultReader, bound to a specific deployed contract.
func NewVaultReader(address common.Address, backend bind.ContractBackend) (*VaultReader, error) {
	contract, err := bindVaultReader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultReader{VaultReaderCaller: VaultReaderCaller{contract: contract}, VaultReaderTransactor: VaultReaderTransactor{contract: contract}, VaultReaderFilterer: VaultReaderFilterer{contract: contract}}, nil
}

// NewVaultReaderCaller creates a new read-only instance of VaultReader, bound to a specific deployed contract.
func NewVaultReaderCaller(address common.Address, caller bind.ContractCaller) (*VaultReaderCaller, error) {
	contract, err := bindVaultReader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultReaderCaller{contract: contract}, nil
}

// NewVaultReaderTransactor creates a new write-only instance of VaultReader, bound to a specific deployed contract.
func NewVaultReaderTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultReaderTransactor, error) {
	contract, err := bindVaultReader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultReaderTransactor{contract: contract}, nil
}

// NewVaultReaderFilterer creates a new log filterer instance of VaultReader, bound to a specific deployed contract.
func NewVaultReaderFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultReaderFilterer, error) {
	contract, err := bindVaultReader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultReaderFilterer{contract: contract}, nil
}

// bindVaultReader binds a generic wrapper to an already deployed contract.
func bindVaultReader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VaultReaderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultReader *VaultReaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultReader.Contract.VaultReaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultReader *VaultReaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultReader.Contract.VaultReaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultReader *VaultReaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultReader.Contract.VaultReaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultReader *VaultReaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultReader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultReader *VaultReaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultReader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultReader *VaultReaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultReader.Contract.contract.Transact(opts, method, params...)
}

// GetVaultTokenInfoV3 is a free data retrieval call binding the contract method 0xf75e8101.
//
// Solidity: function getVaultTokenInfoV3(address _vault, address _positionManager, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_VaultReader *VaultReaderCaller) GetVaultTokenInfoV3(opts *bind.CallOpts, _vault common.Address, _positionManager common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _VaultReader.contract.Call(opts, &out, "getVaultTokenInfoV3", _vault, _positionManager, _weth, _usdgAmount, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetVaultTokenInfoV3 is a free data retrieval call binding the contract method 0xf75e8101.
//
// Solidity: function getVaultTokenInfoV3(address _vault, address _positionManager, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_VaultReader *VaultReaderSession) GetVaultTokenInfoV3(_vault common.Address, _positionManager common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	return _VaultReader.Contract.GetVaultTokenInfoV3(&_VaultReader.CallOpts, _vault, _positionManager, _weth, _usdgAmount, _tokens)
}

// GetVaultTokenInfoV3 is a free data retrieval call binding the contract method 0xf75e8101.
//
// Solidity: function getVaultTokenInfoV3(address _vault, address _positionManager, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_VaultReader *VaultReaderCallerSession) GetVaultTokenInfoV3(_vault common.Address, _positionManager common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	return _VaultReader.Contract.GetVaultTokenInfoV3(&_VaultReader.CallOpts, _vault, _positionManager, _weth, _usdgAmount, _tokens)
}

// GetVaultTokenInfoV4 is a free data retrieval call binding the contract method 0xd3f3265c.
//
// Solidity: function getVaultTokenInfoV4(address _vault, address _positionManager, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_VaultReader *VaultReaderCaller) GetVaultTokenInfoV4(opts *bind.CallOpts, _vault common.Address, _positionManager common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _VaultReader.contract.Call(opts, &out, "getVaultTokenInfoV4", _vault, _positionManager, _weth, _usdgAmount, _tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetVaultTokenInfoV4 is a free data retrieval call binding the contract method 0xd3f3265c.
//
// Solidity: function getVaultTokenInfoV4(address _vault, address _positionManager, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_VaultReader *VaultReaderSession) GetVaultTokenInfoV4(_vault common.Address, _positionManager common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	return _VaultReader.Contract.GetVaultTokenInfoV4(&_VaultReader.CallOpts, _vault, _positionManager, _weth, _usdgAmount, _tokens)
}

// GetVaultTokenInfoV4 is a free data retrieval call binding the contract method 0xd3f3265c.
//
// Solidity: function getVaultTokenInfoV4(address _vault, address _positionManager, address _weth, uint256 _usdgAmount, address[] _tokens) view returns(uint256[])
func (_VaultReader *VaultReaderCallerSession) GetVaultTokenInfoV4(_vault common.Address, _positionManager common.Address, _weth common.Address, _usdgAmount *big.Int, _tokens []common.Address) ([]*big.Int, error) {
	return _VaultReader.Contract.GetVaultTokenInfoV4(&_VaultReader.CallOpts, _vault, _positionManager, _weth, _usdgAmount, _tokens)
}

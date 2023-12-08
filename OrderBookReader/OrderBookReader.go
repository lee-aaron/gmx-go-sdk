// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OrderBookReader

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

// OrderBookReaderMetaData contains all meta data concerning the OrderBookReader contract.
var OrderBookReaderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getDecreaseOrders\",\"inputs\":[{\"name\":\"_orderBookAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIncreaseOrders\",\"inputs\":[{\"name\":\"_orderBookAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSwapOrders\",\"inputs\":[{\"name\":\"_orderBookAddress\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"}]",
}

// OrderBookReaderABI is the input ABI used to generate the binding from.
// Deprecated: Use OrderBookReaderMetaData.ABI instead.
var OrderBookReaderABI = OrderBookReaderMetaData.ABI

// OrderBookReader is an auto generated Go binding around an Ethereum contract.
type OrderBookReader struct {
	OrderBookReaderCaller     // Read-only binding to the contract
	OrderBookReaderTransactor // Write-only binding to the contract
	OrderBookReaderFilterer   // Log filterer for contract events
}

// OrderBookReaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderBookReaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderBookReaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderBookReaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderBookReaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderBookReaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderBookReaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderBookReaderSession struct {
	Contract     *OrderBookReader  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderBookReaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderBookReaderCallerSession struct {
	Contract *OrderBookReaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// OrderBookReaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderBookReaderTransactorSession struct {
	Contract     *OrderBookReaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// OrderBookReaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderBookReaderRaw struct {
	Contract *OrderBookReader // Generic contract binding to access the raw methods on
}

// OrderBookReaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderBookReaderCallerRaw struct {
	Contract *OrderBookReaderCaller // Generic read-only contract binding to access the raw methods on
}

// OrderBookReaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderBookReaderTransactorRaw struct {
	Contract *OrderBookReaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderBookReader creates a new instance of OrderBookReader, bound to a specific deployed contract.
func NewOrderBookReader(address common.Address, backend bind.ContractBackend) (*OrderBookReader, error) {
	contract, err := bindOrderBookReader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrderBookReader{OrderBookReaderCaller: OrderBookReaderCaller{contract: contract}, OrderBookReaderTransactor: OrderBookReaderTransactor{contract: contract}, OrderBookReaderFilterer: OrderBookReaderFilterer{contract: contract}}, nil
}

// NewOrderBookReaderCaller creates a new read-only instance of OrderBookReader, bound to a specific deployed contract.
func NewOrderBookReaderCaller(address common.Address, caller bind.ContractCaller) (*OrderBookReaderCaller, error) {
	contract, err := bindOrderBookReader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderBookReaderCaller{contract: contract}, nil
}

// NewOrderBookReaderTransactor creates a new write-only instance of OrderBookReader, bound to a specific deployed contract.
func NewOrderBookReaderTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderBookReaderTransactor, error) {
	contract, err := bindOrderBookReader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderBookReaderTransactor{contract: contract}, nil
}

// NewOrderBookReaderFilterer creates a new log filterer instance of OrderBookReader, bound to a specific deployed contract.
func NewOrderBookReaderFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderBookReaderFilterer, error) {
	contract, err := bindOrderBookReader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderBookReaderFilterer{contract: contract}, nil
}

// bindOrderBookReader binds a generic wrapper to an already deployed contract.
func bindOrderBookReader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OrderBookReaderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderBookReader *OrderBookReaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderBookReader.Contract.OrderBookReaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderBookReader *OrderBookReaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBookReader.Contract.OrderBookReaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderBookReader *OrderBookReaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderBookReader.Contract.OrderBookReaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderBookReader *OrderBookReaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderBookReader.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderBookReader *OrderBookReaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBookReader.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderBookReader *OrderBookReaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderBookReader.Contract.contract.Transact(opts, method, params...)
}

// GetDecreaseOrders is a free data retrieval call binding the contract method 0x0ce933b9.
//
// Solidity: function getDecreaseOrders(address _orderBookAddress, address _account, uint256[] _indices) view returns(uint256[], address[])
func (_OrderBookReader *OrderBookReaderCaller) GetDecreaseOrders(opts *bind.CallOpts, _orderBookAddress common.Address, _account common.Address, _indices []*big.Int) ([]*big.Int, []common.Address, error) {
	var out []interface{}
	err := _OrderBookReader.contract.Call(opts, &out, "getDecreaseOrders", _orderBookAddress, _account, _indices)

	if err != nil {
		return *new([]*big.Int), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

// GetDecreaseOrders is a free data retrieval call binding the contract method 0x0ce933b9.
//
// Solidity: function getDecreaseOrders(address _orderBookAddress, address _account, uint256[] _indices) view returns(uint256[], address[])
func (_OrderBookReader *OrderBookReaderSession) GetDecreaseOrders(_orderBookAddress common.Address, _account common.Address, _indices []*big.Int) ([]*big.Int, []common.Address, error) {
	return _OrderBookReader.Contract.GetDecreaseOrders(&_OrderBookReader.CallOpts, _orderBookAddress, _account, _indices)
}

// GetDecreaseOrders is a free data retrieval call binding the contract method 0x0ce933b9.
//
// Solidity: function getDecreaseOrders(address _orderBookAddress, address _account, uint256[] _indices) view returns(uint256[], address[])
func (_OrderBookReader *OrderBookReaderCallerSession) GetDecreaseOrders(_orderBookAddress common.Address, _account common.Address, _indices []*big.Int) ([]*big.Int, []common.Address, error) {
	return _OrderBookReader.Contract.GetDecreaseOrders(&_OrderBookReader.CallOpts, _orderBookAddress, _account, _indices)
}

// GetIncreaseOrders is a free data retrieval call binding the contract method 0xc38ccd50.
//
// Solidity: function getIncreaseOrders(address _orderBookAddress, address _account, uint256[] _indices) view returns(uint256[], address[])
func (_OrderBookReader *OrderBookReaderCaller) GetIncreaseOrders(opts *bind.CallOpts, _orderBookAddress common.Address, _account common.Address, _indices []*big.Int) ([]*big.Int, []common.Address, error) {
	var out []interface{}
	err := _OrderBookReader.contract.Call(opts, &out, "getIncreaseOrders", _orderBookAddress, _account, _indices)

	if err != nil {
		return *new([]*big.Int), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

// GetIncreaseOrders is a free data retrieval call binding the contract method 0xc38ccd50.
//
// Solidity: function getIncreaseOrders(address _orderBookAddress, address _account, uint256[] _indices) view returns(uint256[], address[])
func (_OrderBookReader *OrderBookReaderSession) GetIncreaseOrders(_orderBookAddress common.Address, _account common.Address, _indices []*big.Int) ([]*big.Int, []common.Address, error) {
	return _OrderBookReader.Contract.GetIncreaseOrders(&_OrderBookReader.CallOpts, _orderBookAddress, _account, _indices)
}

// GetIncreaseOrders is a free data retrieval call binding the contract method 0xc38ccd50.
//
// Solidity: function getIncreaseOrders(address _orderBookAddress, address _account, uint256[] _indices) view returns(uint256[], address[])
func (_OrderBookReader *OrderBookReaderCallerSession) GetIncreaseOrders(_orderBookAddress common.Address, _account common.Address, _indices []*big.Int) ([]*big.Int, []common.Address, error) {
	return _OrderBookReader.Contract.GetIncreaseOrders(&_OrderBookReader.CallOpts, _orderBookAddress, _account, _indices)
}

// GetSwapOrders is a free data retrieval call binding the contract method 0x2e181469.
//
// Solidity: function getSwapOrders(address _orderBookAddress, address _account, uint256[] _indices) view returns(uint256[], address[])
func (_OrderBookReader *OrderBookReaderCaller) GetSwapOrders(opts *bind.CallOpts, _orderBookAddress common.Address, _account common.Address, _indices []*big.Int) ([]*big.Int, []common.Address, error) {
	var out []interface{}
	err := _OrderBookReader.contract.Call(opts, &out, "getSwapOrders", _orderBookAddress, _account, _indices)

	if err != nil {
		return *new([]*big.Int), *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return out0, out1, err

}

// GetSwapOrders is a free data retrieval call binding the contract method 0x2e181469.
//
// Solidity: function getSwapOrders(address _orderBookAddress, address _account, uint256[] _indices) view returns(uint256[], address[])
func (_OrderBookReader *OrderBookReaderSession) GetSwapOrders(_orderBookAddress common.Address, _account common.Address, _indices []*big.Int) ([]*big.Int, []common.Address, error) {
	return _OrderBookReader.Contract.GetSwapOrders(&_OrderBookReader.CallOpts, _orderBookAddress, _account, _indices)
}

// GetSwapOrders is a free data retrieval call binding the contract method 0x2e181469.
//
// Solidity: function getSwapOrders(address _orderBookAddress, address _account, uint256[] _indices) view returns(uint256[], address[])
func (_OrderBookReader *OrderBookReaderCallerSession) GetSwapOrders(_orderBookAddress common.Address, _account common.Address, _indices []*big.Int) ([]*big.Int, []common.Address, error) {
	return _OrderBookReader.Contract.GetSwapOrders(&_OrderBookReader.CallOpts, _orderBookAddress, _account, _indices)
}

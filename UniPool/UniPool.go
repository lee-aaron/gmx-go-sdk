// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package UniPool

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

// UniPoolMetaData contains all meta data concerning the UniPool contract.
var UniPoolMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"observe\",\"inputs\":[{\"name\":\"\",\"type\":\"uint32[]\",\"internalType\":\"uint32[]\"}],\"outputs\":[{\"name\":\"tickCumulatives\",\"type\":\"int56[]\",\"internalType\":\"int56[]\"},{\"name\":\"secondsPerLiquidityCumulativeX128s\",\"type\":\"uint160[]\",\"internalType\":\"uint160[]\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"slot0\",\"inputs\":[],\"outputs\":[{\"name\":\"sqrtPriceX96\",\"type\":\"uint160\",\"internalType\":\"uint160\"},{\"name\":\"tick\",\"type\":\"int24\",\"internalType\":\"int24\"},{\"name\":\"observationIndex\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"observationCardinality\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"observationCardinalityNext\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"feeProtocol\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"unlocked\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tickSpacing\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"int24\",\"internalType\":\"int24\"}],\"stateMutability\":\"pure\"}]",
}

// UniPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use UniPoolMetaData.ABI instead.
var UniPoolABI = UniPoolMetaData.ABI

// UniPool is an auto generated Go binding around an Ethereum contract.
type UniPool struct {
	UniPoolCaller     // Read-only binding to the contract
	UniPoolTransactor // Write-only binding to the contract
	UniPoolFilterer   // Log filterer for contract events
}

// UniPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniPoolSession struct {
	Contract     *UniPool          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniPoolCallerSession struct {
	Contract *UniPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// UniPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniPoolTransactorSession struct {
	Contract     *UniPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UniPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniPoolRaw struct {
	Contract *UniPool // Generic contract binding to access the raw methods on
}

// UniPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniPoolCallerRaw struct {
	Contract *UniPoolCaller // Generic read-only contract binding to access the raw methods on
}

// UniPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniPoolTransactorRaw struct {
	Contract *UniPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniPool creates a new instance of UniPool, bound to a specific deployed contract.
func NewUniPool(address common.Address, backend bind.ContractBackend) (*UniPool, error) {
	contract, err := bindUniPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniPool{UniPoolCaller: UniPoolCaller{contract: contract}, UniPoolTransactor: UniPoolTransactor{contract: contract}, UniPoolFilterer: UniPoolFilterer{contract: contract}}, nil
}

// NewUniPoolCaller creates a new read-only instance of UniPool, bound to a specific deployed contract.
func NewUniPoolCaller(address common.Address, caller bind.ContractCaller) (*UniPoolCaller, error) {
	contract, err := bindUniPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniPoolCaller{contract: contract}, nil
}

// NewUniPoolTransactor creates a new write-only instance of UniPool, bound to a specific deployed contract.
func NewUniPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*UniPoolTransactor, error) {
	contract, err := bindUniPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniPoolTransactor{contract: contract}, nil
}

// NewUniPoolFilterer creates a new log filterer instance of UniPool, bound to a specific deployed contract.
func NewUniPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*UniPoolFilterer, error) {
	contract, err := bindUniPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniPoolFilterer{contract: contract}, nil
}

// bindUniPool binds a generic wrapper to an already deployed contract.
func bindUniPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniPool *UniPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniPool.Contract.UniPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniPool *UniPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniPool.Contract.UniPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniPool *UniPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniPool.Contract.UniPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniPool *UniPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniPool *UniPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniPool *UniPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniPool.Contract.contract.Transact(opts, method, params...)
}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] ) pure returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulativeX128s)
func (_UniPool *UniPoolCaller) Observe(opts *bind.CallOpts, arg0 []uint32) (struct {
	TickCumulatives                    []*big.Int
	SecondsPerLiquidityCumulativeX128s []*big.Int
}, error) {
	var out []interface{}
	err := _UniPool.contract.Call(opts, &out, "observe", arg0)

	outstruct := new(struct {
		TickCumulatives                    []*big.Int
		SecondsPerLiquidityCumulativeX128s []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TickCumulatives = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.SecondsPerLiquidityCumulativeX128s = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] ) pure returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulativeX128s)
func (_UniPool *UniPoolSession) Observe(arg0 []uint32) (struct {
	TickCumulatives                    []*big.Int
	SecondsPerLiquidityCumulativeX128s []*big.Int
}, error) {
	return _UniPool.Contract.Observe(&_UniPool.CallOpts, arg0)
}

// Observe is a free data retrieval call binding the contract method 0x883bdbfd.
//
// Solidity: function observe(uint32[] ) pure returns(int56[] tickCumulatives, uint160[] secondsPerLiquidityCumulativeX128s)
func (_UniPool *UniPoolCallerSession) Observe(arg0 []uint32) (struct {
	TickCumulatives                    []*big.Int
	SecondsPerLiquidityCumulativeX128s []*big.Int
}, error) {
	return _UniPool.Contract.Observe(&_UniPool.CallOpts, arg0)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_UniPool *UniPoolCaller) Slot0(opts *bind.CallOpts) (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	var out []interface{}
	err := _UniPool.contract.Call(opts, &out, "slot0")

	outstruct := new(struct {
		SqrtPriceX96               *big.Int
		Tick                       *big.Int
		ObservationIndex           uint16
		ObservationCardinality     uint16
		ObservationCardinalityNext uint16
		FeeProtocol                uint8
		Unlocked                   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SqrtPriceX96 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tick = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ObservationIndex = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.ObservationCardinality = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.ObservationCardinalityNext = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.FeeProtocol = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.Unlocked = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_UniPool *UniPoolSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	return _UniPool.Contract.Slot0(&_UniPool.CallOpts)
}

// Slot0 is a free data retrieval call binding the contract method 0x3850c7bd.
//
// Solidity: function slot0() view returns(uint160 sqrtPriceX96, int24 tick, uint16 observationIndex, uint16 observationCardinality, uint16 observationCardinalityNext, uint8 feeProtocol, bool unlocked)
func (_UniPool *UniPoolCallerSession) Slot0() (struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	return _UniPool.Contract.Slot0(&_UniPool.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() pure returns(int24)
func (_UniPool *UniPoolCaller) TickSpacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UniPool.contract.Call(opts, &out, "tickSpacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() pure returns(int24)
func (_UniPool *UniPoolSession) TickSpacing() (*big.Int, error) {
	return _UniPool.Contract.TickSpacing(&_UniPool.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() pure returns(int24)
func (_UniPool *UniPoolCallerSession) TickSpacing() (*big.Int, error) {
	return _UniPool.Contract.TickSpacing(&_UniPool.CallOpts)
}

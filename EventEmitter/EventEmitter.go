// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EventEmitter

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

// EventUtilsAddressArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsAddressArrayKeyValue struct {
	Key   string
	Value []common.Address
}

// EventUtilsAddressItems is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsAddressItems struct {
	Items      []EventUtilsAddressKeyValue
	ArrayItems []EventUtilsAddressArrayKeyValue
}

// EventUtilsAddressKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsAddressKeyValue struct {
	Key   string
	Value common.Address
}

// EventUtilsBoolArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBoolArrayKeyValue struct {
	Key   string
	Value []bool
}

// EventUtilsBoolItems is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBoolItems struct {
	Items      []EventUtilsBoolKeyValue
	ArrayItems []EventUtilsBoolArrayKeyValue
}

// EventUtilsBoolKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBoolKeyValue struct {
	Key   string
	Value bool
}

// EventUtilsBytes32ArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBytes32ArrayKeyValue struct {
	Key   string
	Value [][32]byte
}

// EventUtilsBytes32Items is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBytes32Items struct {
	Items      []EventUtilsBytes32KeyValue
	ArrayItems []EventUtilsBytes32ArrayKeyValue
}

// EventUtilsBytes32KeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBytes32KeyValue struct {
	Key   string
	Value [32]byte
}

// EventUtilsBytesArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBytesArrayKeyValue struct {
	Key   string
	Value [][]byte
}

// EventUtilsBytesItems is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBytesItems struct {
	Items      []EventUtilsBytesKeyValue
	ArrayItems []EventUtilsBytesArrayKeyValue
}

// EventUtilsBytesKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsBytesKeyValue struct {
	Key   string
	Value []byte
}

// EventUtilsEventLogData is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsEventLogData struct {
	AddressItems EventUtilsAddressItems
	UintItems    EventUtilsUintItems
	IntItems     EventUtilsIntItems
	BoolItems    EventUtilsBoolItems
	Bytes32Items EventUtilsBytes32Items
	BytesItems   EventUtilsBytesItems
	StringItems  EventUtilsStringItems
}

// EventUtilsIntArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsIntArrayKeyValue struct {
	Key   string
	Value []*big.Int
}

// EventUtilsIntItems is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsIntItems struct {
	Items      []EventUtilsIntKeyValue
	ArrayItems []EventUtilsIntArrayKeyValue
}

// EventUtilsIntKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsIntKeyValue struct {
	Key   string
	Value *big.Int
}

// EventUtilsStringArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsStringArrayKeyValue struct {
	Key   string
	Value []string
}

// EventUtilsStringItems is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsStringItems struct {
	Items      []EventUtilsStringKeyValue
	ArrayItems []EventUtilsStringArrayKeyValue
}

// EventUtilsStringKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsStringKeyValue struct {
	Key   string
	Value string
}

// EventUtilsUintArrayKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsUintArrayKeyValue struct {
	Key   string
	Value []*big.Int
}

// EventUtilsUintItems is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsUintItems struct {
	Items      []EventUtilsUintKeyValue
	ArrayItems []EventUtilsUintArrayKeyValue
}

// EventUtilsUintKeyValue is an auto generated low-level Go binding around an user-defined struct.
type EventUtilsUintKeyValue struct {
	Key   string
	Value *big.Int
}

// EventEmitterMetaData contains all meta data concerning the EventEmitter contract.
var EventEmitterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_roleStore\",\"type\":\"address\",\"internalType\":\"contractRoleStore\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emitDataLog1\",\"inputs\":[{\"name\":\"topic1\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emitDataLog2\",\"inputs\":[{\"name\":\"topic1\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"topic2\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emitDataLog3\",\"inputs\":[{\"name\":\"topic1\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"topic2\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"topic3\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emitDataLog4\",\"inputs\":[{\"name\":\"topic1\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"topic2\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"topic3\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"topic4\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emitEventLog\",\"inputs\":[{\"name\":\"eventName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"eventData\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.EventLogData\",\"components\":[{\"name\":\"addressItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.AddressItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.AddressKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.AddressArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}]},{\"name\":\"uintItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.UintItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.UintKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.UintArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]},{\"name\":\"intItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.IntItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.IntKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.IntArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}]}]},{\"name\":\"boolItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.BoolItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BoolKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BoolArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}]}]},{\"name\":\"bytes32Items\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.Bytes32Items\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.Bytes32KeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.Bytes32ArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}]},{\"name\":\"bytesItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.BytesItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BytesKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BytesArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]}]},{\"name\":\"stringItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.StringItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.StringKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.StringArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emitEventLog1\",\"inputs\":[{\"name\":\"eventName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"topic1\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"eventData\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.EventLogData\",\"components\":[{\"name\":\"addressItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.AddressItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.AddressKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.AddressArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}]},{\"name\":\"uintItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.UintItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.UintKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.UintArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]},{\"name\":\"intItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.IntItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.IntKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.IntArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}]}]},{\"name\":\"boolItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.BoolItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BoolKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BoolArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}]}]},{\"name\":\"bytes32Items\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.Bytes32Items\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.Bytes32KeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.Bytes32ArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}]},{\"name\":\"bytesItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.BytesItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BytesKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BytesArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]}]},{\"name\":\"stringItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.StringItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.StringKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.StringArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emitEventLog2\",\"inputs\":[{\"name\":\"eventName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"topic1\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"topic2\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"eventData\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.EventLogData\",\"components\":[{\"name\":\"addressItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.AddressItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.AddressKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.AddressArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}]},{\"name\":\"uintItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.UintItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.UintKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.UintArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]},{\"name\":\"intItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.IntItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.IntKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.IntArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}]}]},{\"name\":\"boolItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.BoolItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BoolKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BoolArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}]}]},{\"name\":\"bytes32Items\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.Bytes32Items\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.Bytes32KeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.Bytes32ArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}]},{\"name\":\"bytesItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.BytesItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BytesKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BytesArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]}]},{\"name\":\"stringItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.StringItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.StringKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.StringArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"roleStore\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractRoleStore\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"EventLog\",\"inputs\":[{\"name\":\"msgSender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"eventName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"eventNameHash\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"eventData\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structEventUtils.EventLogData\",\"components\":[{\"name\":\"addressItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.AddressItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.AddressKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.AddressArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}]},{\"name\":\"uintItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.UintItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.UintKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.UintArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]},{\"name\":\"intItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.IntItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.IntKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.IntArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}]}]},{\"name\":\"boolItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.BoolItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BoolKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BoolArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}]}]},{\"name\":\"bytes32Items\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.Bytes32Items\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.Bytes32KeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.Bytes32ArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}]},{\"name\":\"bytesItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.BytesItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BytesKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BytesArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]}]},{\"name\":\"stringItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.StringItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.StringKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.StringArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EventLog1\",\"inputs\":[{\"name\":\"msgSender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"eventName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"eventNameHash\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"topic1\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"eventData\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structEventUtils.EventLogData\",\"components\":[{\"name\":\"addressItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.AddressItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.AddressKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.AddressArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}]},{\"name\":\"uintItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.UintItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.UintKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.UintArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]},{\"name\":\"intItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.IntItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.IntKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.IntArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}]}]},{\"name\":\"boolItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.BoolItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BoolKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BoolArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}]}]},{\"name\":\"bytes32Items\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.Bytes32Items\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.Bytes32KeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.Bytes32ArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}]},{\"name\":\"bytesItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.BytesItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BytesKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BytesArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]}]},{\"name\":\"stringItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.StringItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.StringKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.StringArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EventLog2\",\"inputs\":[{\"name\":\"msgSender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"eventName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"eventNameHash\",\"type\":\"string\",\"indexed\":true,\"internalType\":\"string\"},{\"name\":\"topic1\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"topic2\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"eventData\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structEventUtils.EventLogData\",\"components\":[{\"name\":\"addressItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.AddressItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.AddressKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.AddressArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}]},{\"name\":\"uintItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.UintItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.UintKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.UintArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}]}]},{\"name\":\"intItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.IntItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.IntKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"int256\",\"internalType\":\"int256\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.IntArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}]}]},{\"name\":\"boolItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.BoolItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BoolKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BoolArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}]}]},{\"name\":\"bytes32Items\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.Bytes32Items\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.Bytes32KeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.Bytes32ArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}]},{\"name\":\"bytesItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.BytesItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BytesKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.BytesArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}]}]},{\"name\":\"stringItems\",\"type\":\"tuple\",\"internalType\":\"structEventUtils.StringItems\",\"components\":[{\"name\":\"items\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.StringKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"name\":\"arrayItems\",\"type\":\"tuple[]\",\"internalType\":\"structEventUtils.StringArrayKeyValue[]\",\"components\":[{\"name\":\"key\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"value\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}]}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"msgSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
}

// EventEmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use EventEmitterMetaData.ABI instead.
var EventEmitterABI = EventEmitterMetaData.ABI

// EventEmitter is an auto generated Go binding around an Ethereum contract.
type EventEmitter struct {
	EventEmitterCaller     // Read-only binding to the contract
	EventEmitterTransactor // Write-only binding to the contract
	EventEmitterFilterer   // Log filterer for contract events
}

// EventEmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventEmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventEmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventEmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventEmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventEmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventEmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventEmitterSession struct {
	Contract     *EventEmitter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventEmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventEmitterCallerSession struct {
	Contract *EventEmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// EventEmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventEmitterTransactorSession struct {
	Contract     *EventEmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EventEmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventEmitterRaw struct {
	Contract *EventEmitter // Generic contract binding to access the raw methods on
}

// EventEmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventEmitterCallerRaw struct {
	Contract *EventEmitterCaller // Generic read-only contract binding to access the raw methods on
}

// EventEmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventEmitterTransactorRaw struct {
	Contract *EventEmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventEmitter creates a new instance of EventEmitter, bound to a specific deployed contract.
func NewEventEmitter(address common.Address, backend bind.ContractBackend) (*EventEmitter, error) {
	contract, err := bindEventEmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EventEmitter{EventEmitterCaller: EventEmitterCaller{contract: contract}, EventEmitterTransactor: EventEmitterTransactor{contract: contract}, EventEmitterFilterer: EventEmitterFilterer{contract: contract}}, nil
}

// NewEventEmitterCaller creates a new read-only instance of EventEmitter, bound to a specific deployed contract.
func NewEventEmitterCaller(address common.Address, caller bind.ContractCaller) (*EventEmitterCaller, error) {
	contract, err := bindEventEmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventEmitterCaller{contract: contract}, nil
}

// NewEventEmitterTransactor creates a new write-only instance of EventEmitter, bound to a specific deployed contract.
func NewEventEmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*EventEmitterTransactor, error) {
	contract, err := bindEventEmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventEmitterTransactor{contract: contract}, nil
}

// NewEventEmitterFilterer creates a new log filterer instance of EventEmitter, bound to a specific deployed contract.
func NewEventEmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*EventEmitterFilterer, error) {
	contract, err := bindEventEmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventEmitterFilterer{contract: contract}, nil
}

// bindEventEmitter binds a generic wrapper to an already deployed contract.
func bindEventEmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EventEmitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventEmitter *EventEmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventEmitter.Contract.EventEmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventEmitter *EventEmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventEmitter.Contract.EventEmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventEmitter *EventEmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventEmitter.Contract.EventEmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventEmitter *EventEmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventEmitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventEmitter *EventEmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventEmitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventEmitter *EventEmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventEmitter.Contract.contract.Transact(opts, method, params...)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_EventEmitter *EventEmitterCaller) RoleStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EventEmitter.contract.Call(opts, &out, "roleStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_EventEmitter *EventEmitterSession) RoleStore() (common.Address, error) {
	return _EventEmitter.Contract.RoleStore(&_EventEmitter.CallOpts)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_EventEmitter *EventEmitterCallerSession) RoleStore() (common.Address, error) {
	return _EventEmitter.Contract.RoleStore(&_EventEmitter.CallOpts)
}

// EmitDataLog1 is a paid mutator transaction binding the contract method 0xf9d5c0ea.
//
// Solidity: function emitDataLog1(bytes32 topic1, bytes data) returns()
func (_EventEmitter *EventEmitterTransactor) EmitDataLog1(opts *bind.TransactOpts, topic1 [32]byte, data []byte) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "emitDataLog1", topic1, data)
}

// EmitDataLog1 is a paid mutator transaction binding the contract method 0xf9d5c0ea.
//
// Solidity: function emitDataLog1(bytes32 topic1, bytes data) returns()
func (_EventEmitter *EventEmitterSession) EmitDataLog1(topic1 [32]byte, data []byte) (*types.Transaction, error) {
	return _EventEmitter.Contract.EmitDataLog1(&_EventEmitter.TransactOpts, topic1, data)
}

// EmitDataLog1 is a paid mutator transaction binding the contract method 0xf9d5c0ea.
//
// Solidity: function emitDataLog1(bytes32 topic1, bytes data) returns()
func (_EventEmitter *EventEmitterTransactorSession) EmitDataLog1(topic1 [32]byte, data []byte) (*types.Transaction, error) {
	return _EventEmitter.Contract.EmitDataLog1(&_EventEmitter.TransactOpts, topic1, data)
}

// EmitDataLog2 is a paid mutator transaction binding the contract method 0xdda0db32.
//
// Solidity: function emitDataLog2(bytes32 topic1, bytes32 topic2, bytes data) returns()
func (_EventEmitter *EventEmitterTransactor) EmitDataLog2(opts *bind.TransactOpts, topic1 [32]byte, topic2 [32]byte, data []byte) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "emitDataLog2", topic1, topic2, data)
}

// EmitDataLog2 is a paid mutator transaction binding the contract method 0xdda0db32.
//
// Solidity: function emitDataLog2(bytes32 topic1, bytes32 topic2, bytes data) returns()
func (_EventEmitter *EventEmitterSession) EmitDataLog2(topic1 [32]byte, topic2 [32]byte, data []byte) (*types.Transaction, error) {
	return _EventEmitter.Contract.EmitDataLog2(&_EventEmitter.TransactOpts, topic1, topic2, data)
}

// EmitDataLog2 is a paid mutator transaction binding the contract method 0xdda0db32.
//
// Solidity: function emitDataLog2(bytes32 topic1, bytes32 topic2, bytes data) returns()
func (_EventEmitter *EventEmitterTransactorSession) EmitDataLog2(topic1 [32]byte, topic2 [32]byte, data []byte) (*types.Transaction, error) {
	return _EventEmitter.Contract.EmitDataLog2(&_EventEmitter.TransactOpts, topic1, topic2, data)
}

// EmitDataLog3 is a paid mutator transaction binding the contract method 0xb3ac1c38.
//
// Solidity: function emitDataLog3(bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes data) returns()
func (_EventEmitter *EventEmitterTransactor) EmitDataLog3(opts *bind.TransactOpts, topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, data []byte) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "emitDataLog3", topic1, topic2, topic3, data)
}

// EmitDataLog3 is a paid mutator transaction binding the contract method 0xb3ac1c38.
//
// Solidity: function emitDataLog3(bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes data) returns()
func (_EventEmitter *EventEmitterSession) EmitDataLog3(topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, data []byte) (*types.Transaction, error) {
	return _EventEmitter.Contract.EmitDataLog3(&_EventEmitter.TransactOpts, topic1, topic2, topic3, data)
}

// EmitDataLog3 is a paid mutator transaction binding the contract method 0xb3ac1c38.
//
// Solidity: function emitDataLog3(bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes data) returns()
func (_EventEmitter *EventEmitterTransactorSession) EmitDataLog3(topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, data []byte) (*types.Transaction, error) {
	return _EventEmitter.Contract.EmitDataLog3(&_EventEmitter.TransactOpts, topic1, topic2, topic3, data)
}

// EmitDataLog4 is a paid mutator transaction binding the contract method 0xee288ce8.
//
// Solidity: function emitDataLog4(bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes32 topic4, bytes data) returns()
func (_EventEmitter *EventEmitterTransactor) EmitDataLog4(opts *bind.TransactOpts, topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, topic4 [32]byte, data []byte) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "emitDataLog4", topic1, topic2, topic3, topic4, data)
}

// EmitDataLog4 is a paid mutator transaction binding the contract method 0xee288ce8.
//
// Solidity: function emitDataLog4(bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes32 topic4, bytes data) returns()
func (_EventEmitter *EventEmitterSession) EmitDataLog4(topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, topic4 [32]byte, data []byte) (*types.Transaction, error) {
	return _EventEmitter.Contract.EmitDataLog4(&_EventEmitter.TransactOpts, topic1, topic2, topic3, topic4, data)
}

// EmitDataLog4 is a paid mutator transaction binding the contract method 0xee288ce8.
//
// Solidity: function emitDataLog4(bytes32 topic1, bytes32 topic2, bytes32 topic3, bytes32 topic4, bytes data) returns()
func (_EventEmitter *EventEmitterTransactorSession) EmitDataLog4(topic1 [32]byte, topic2 [32]byte, topic3 [32]byte, topic4 [32]byte, data []byte) (*types.Transaction, error) {
	return _EventEmitter.Contract.EmitDataLog4(&_EventEmitter.TransactOpts, topic1, topic2, topic3, topic4, data)
}

// EmitEventLog is a paid mutator transaction binding the contract method 0x906c49f6.
//
// Solidity: function emitEventLog(string eventName, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_EventEmitter *EventEmitterTransactor) EmitEventLog(opts *bind.TransactOpts, eventName string, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "emitEventLog", eventName, eventData)
}

// EmitEventLog is a paid mutator transaction binding the contract method 0x906c49f6.
//
// Solidity: function emitEventLog(string eventName, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_EventEmitter *EventEmitterSession) EmitEventLog(eventName string, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _EventEmitter.Contract.EmitEventLog(&_EventEmitter.TransactOpts, eventName, eventData)
}

// EmitEventLog is a paid mutator transaction binding the contract method 0x906c49f6.
//
// Solidity: function emitEventLog(string eventName, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_EventEmitter *EventEmitterTransactorSession) EmitEventLog(eventName string, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _EventEmitter.Contract.EmitEventLog(&_EventEmitter.TransactOpts, eventName, eventData)
}

// EmitEventLog1 is a paid mutator transaction binding the contract method 0x24de01e4.
//
// Solidity: function emitEventLog1(string eventName, bytes32 topic1, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_EventEmitter *EventEmitterTransactor) EmitEventLog1(opts *bind.TransactOpts, eventName string, topic1 [32]byte, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "emitEventLog1", eventName, topic1, eventData)
}

// EmitEventLog1 is a paid mutator transaction binding the contract method 0x24de01e4.
//
// Solidity: function emitEventLog1(string eventName, bytes32 topic1, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_EventEmitter *EventEmitterSession) EmitEventLog1(eventName string, topic1 [32]byte, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _EventEmitter.Contract.EmitEventLog1(&_EventEmitter.TransactOpts, eventName, topic1, eventData)
}

// EmitEventLog1 is a paid mutator transaction binding the contract method 0x24de01e4.
//
// Solidity: function emitEventLog1(string eventName, bytes32 topic1, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_EventEmitter *EventEmitterTransactorSession) EmitEventLog1(eventName string, topic1 [32]byte, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _EventEmitter.Contract.EmitEventLog1(&_EventEmitter.TransactOpts, eventName, topic1, eventData)
}

// EmitEventLog2 is a paid mutator transaction binding the contract method 0x63d16363.
//
// Solidity: function emitEventLog2(string eventName, bytes32 topic1, bytes32 topic2, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_EventEmitter *EventEmitterTransactor) EmitEventLog2(opts *bind.TransactOpts, eventName string, topic1 [32]byte, topic2 [32]byte, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _EventEmitter.contract.Transact(opts, "emitEventLog2", eventName, topic1, topic2, eventData)
}

// EmitEventLog2 is a paid mutator transaction binding the contract method 0x63d16363.
//
// Solidity: function emitEventLog2(string eventName, bytes32 topic1, bytes32 topic2, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_EventEmitter *EventEmitterSession) EmitEventLog2(eventName string, topic1 [32]byte, topic2 [32]byte, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _EventEmitter.Contract.EmitEventLog2(&_EventEmitter.TransactOpts, eventName, topic1, topic2, eventData)
}

// EmitEventLog2 is a paid mutator transaction binding the contract method 0x63d16363.
//
// Solidity: function emitEventLog2(string eventName, bytes32 topic1, bytes32 topic2, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData) returns()
func (_EventEmitter *EventEmitterTransactorSession) EmitEventLog2(eventName string, topic1 [32]byte, topic2 [32]byte, eventData EventUtilsEventLogData) (*types.Transaction, error) {
	return _EventEmitter.Contract.EmitEventLog2(&_EventEmitter.TransactOpts, eventName, topic1, topic2, eventData)
}

// EventEmitterEventLogIterator is returned from FilterEventLog and is used to iterate over the raw logs and unpacked data for EventLog events raised by the EventEmitter contract.
type EventEmitterEventLogIterator struct {
	Event *EventEmitterEventLog // Event containing the contract specifics and raw log

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
func (it *EventEmitterEventLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventEmitterEventLog)
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
		it.Event = new(EventEmitterEventLog)
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
func (it *EventEmitterEventLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventEmitterEventLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventEmitterEventLog represents a EventLog event raised by the EventEmitter contract.
type EventEmitterEventLog struct {
	MsgSender     common.Address
	EventName     string
	EventNameHash common.Hash
	EventData     EventUtilsEventLogData
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEventLog is a free log retrieval operation binding the contract event 0x7e3bde2ba7aca4a8499608ca57f3b0c1c1c93ace63ffd3741a9fab204146fc9a.
//
// Solidity: event EventLog(address msgSender, string eventName, string indexed eventNameHash, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_EventEmitter *EventEmitterFilterer) FilterEventLog(opts *bind.FilterOpts, eventNameHash []string) (*EventEmitterEventLogIterator, error) {

	var eventNameHashRule []interface{}
	for _, eventNameHashItem := range eventNameHash {
		eventNameHashRule = append(eventNameHashRule, eventNameHashItem)
	}

	logs, sub, err := _EventEmitter.contract.FilterLogs(opts, "EventLog", eventNameHashRule)
	if err != nil {
		return nil, err
	}
	return &EventEmitterEventLogIterator{contract: _EventEmitter.contract, event: "EventLog", logs: logs, sub: sub}, nil
}

// WatchEventLog is a free log subscription operation binding the contract event 0x7e3bde2ba7aca4a8499608ca57f3b0c1c1c93ace63ffd3741a9fab204146fc9a.
//
// Solidity: event EventLog(address msgSender, string eventName, string indexed eventNameHash, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_EventEmitter *EventEmitterFilterer) WatchEventLog(opts *bind.WatchOpts, sink chan<- *EventEmitterEventLog, eventNameHash []string) (event.Subscription, error) {

	var eventNameHashRule []interface{}
	for _, eventNameHashItem := range eventNameHash {
		eventNameHashRule = append(eventNameHashRule, eventNameHashItem)
	}

	logs, sub, err := _EventEmitter.contract.WatchLogs(opts, "EventLog", eventNameHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventEmitterEventLog)
				if err := _EventEmitter.contract.UnpackLog(event, "EventLog", log); err != nil {
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

// ParseEventLog is a log parse operation binding the contract event 0x7e3bde2ba7aca4a8499608ca57f3b0c1c1c93ace63ffd3741a9fab204146fc9a.
//
// Solidity: event EventLog(address msgSender, string eventName, string indexed eventNameHash, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_EventEmitter *EventEmitterFilterer) ParseEventLog(log types.Log) (*EventEmitterEventLog, error) {
	event := new(EventEmitterEventLog)
	if err := _EventEmitter.contract.UnpackLog(event, "EventLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventEmitterEventLog1Iterator is returned from FilterEventLog1 and is used to iterate over the raw logs and unpacked data for EventLog1 events raised by the EventEmitter contract.
type EventEmitterEventLog1Iterator struct {
	Event *EventEmitterEventLog1 // Event containing the contract specifics and raw log

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
func (it *EventEmitterEventLog1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventEmitterEventLog1)
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
		it.Event = new(EventEmitterEventLog1)
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
func (it *EventEmitterEventLog1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventEmitterEventLog1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventEmitterEventLog1 represents a EventLog1 event raised by the EventEmitter contract.
type EventEmitterEventLog1 struct {
	MsgSender     common.Address
	EventName     string
	EventNameHash common.Hash
	Topic1        [32]byte
	EventData     EventUtilsEventLogData
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEventLog1 is a free log retrieval operation binding the contract event 0x137a44067c8961cd7e1d876f4754a5a3a75989b4552f1843fc69c3b372def160.
//
// Solidity: event EventLog1(address msgSender, string eventName, string indexed eventNameHash, bytes32 indexed topic1, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_EventEmitter *EventEmitterFilterer) FilterEventLog1(opts *bind.FilterOpts, eventNameHash []string, topic1 [][32]byte) (*EventEmitterEventLog1Iterator, error) {

	var eventNameHashRule []interface{}
	for _, eventNameHashItem := range eventNameHash {
		eventNameHashRule = append(eventNameHashRule, eventNameHashItem)
	}
	var topic1Rule []interface{}
	for _, topic1Item := range topic1 {
		topic1Rule = append(topic1Rule, topic1Item)
	}

	logs, sub, err := _EventEmitter.contract.FilterLogs(opts, "EventLog1", eventNameHashRule, topic1Rule)
	if err != nil {
		return nil, err
	}
	return &EventEmitterEventLog1Iterator{contract: _EventEmitter.contract, event: "EventLog1", logs: logs, sub: sub}, nil
}

// WatchEventLog1 is a free log subscription operation binding the contract event 0x137a44067c8961cd7e1d876f4754a5a3a75989b4552f1843fc69c3b372def160.
//
// Solidity: event EventLog1(address msgSender, string eventName, string indexed eventNameHash, bytes32 indexed topic1, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_EventEmitter *EventEmitterFilterer) WatchEventLog1(opts *bind.WatchOpts, sink chan<- *EventEmitterEventLog1, eventNameHash []string, topic1 [][32]byte) (event.Subscription, error) {

	var eventNameHashRule []interface{}
	for _, eventNameHashItem := range eventNameHash {
		eventNameHashRule = append(eventNameHashRule, eventNameHashItem)
	}
	var topic1Rule []interface{}
	for _, topic1Item := range topic1 {
		topic1Rule = append(topic1Rule, topic1Item)
	}

	logs, sub, err := _EventEmitter.contract.WatchLogs(opts, "EventLog1", eventNameHashRule, topic1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventEmitterEventLog1)
				if err := _EventEmitter.contract.UnpackLog(event, "EventLog1", log); err != nil {
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

// ParseEventLog1 is a log parse operation binding the contract event 0x137a44067c8961cd7e1d876f4754a5a3a75989b4552f1843fc69c3b372def160.
//
// Solidity: event EventLog1(address msgSender, string eventName, string indexed eventNameHash, bytes32 indexed topic1, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_EventEmitter *EventEmitterFilterer) ParseEventLog1(log types.Log) (*EventEmitterEventLog1, error) {
	event := new(EventEmitterEventLog1)
	if err := _EventEmitter.contract.UnpackLog(event, "EventLog1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EventEmitterEventLog2Iterator is returned from FilterEventLog2 and is used to iterate over the raw logs and unpacked data for EventLog2 events raised by the EventEmitter contract.
type EventEmitterEventLog2Iterator struct {
	Event *EventEmitterEventLog2 // Event containing the contract specifics and raw log

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
func (it *EventEmitterEventLog2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventEmitterEventLog2)
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
		it.Event = new(EventEmitterEventLog2)
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
func (it *EventEmitterEventLog2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventEmitterEventLog2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventEmitterEventLog2 represents a EventLog2 event raised by the EventEmitter contract.
type EventEmitterEventLog2 struct {
	MsgSender     common.Address
	EventName     string
	EventNameHash common.Hash
	Topic1        [32]byte
	Topic2        [32]byte
	EventData     EventUtilsEventLogData
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEventLog2 is a free log retrieval operation binding the contract event 0x468a25a7ba624ceea6e540ad6f49171b52495b648417ae91bca21676d8a24dc5.
//
// Solidity: event EventLog2(address msgSender, string eventName, string indexed eventNameHash, bytes32 indexed topic1, bytes32 indexed topic2, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_EventEmitter *EventEmitterFilterer) FilterEventLog2(opts *bind.FilterOpts, eventNameHash []string, topic1 [][32]byte, topic2 [][32]byte) (*EventEmitterEventLog2Iterator, error) {

	var eventNameHashRule []interface{}
	for _, eventNameHashItem := range eventNameHash {
		eventNameHashRule = append(eventNameHashRule, eventNameHashItem)
	}
	var topic1Rule []interface{}
	for _, topic1Item := range topic1 {
		topic1Rule = append(topic1Rule, topic1Item)
	}
	var topic2Rule []interface{}
	for _, topic2Item := range topic2 {
		topic2Rule = append(topic2Rule, topic2Item)
	}

	logs, sub, err := _EventEmitter.contract.FilterLogs(opts, "EventLog2", eventNameHashRule, topic1Rule, topic2Rule)
	if err != nil {
		return nil, err
	}
	return &EventEmitterEventLog2Iterator{contract: _EventEmitter.contract, event: "EventLog2", logs: logs, sub: sub}, nil
}

// WatchEventLog2 is a free log subscription operation binding the contract event 0x468a25a7ba624ceea6e540ad6f49171b52495b648417ae91bca21676d8a24dc5.
//
// Solidity: event EventLog2(address msgSender, string eventName, string indexed eventNameHash, bytes32 indexed topic1, bytes32 indexed topic2, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_EventEmitter *EventEmitterFilterer) WatchEventLog2(opts *bind.WatchOpts, sink chan<- *EventEmitterEventLog2, eventNameHash []string, topic1 [][32]byte, topic2 [][32]byte) (event.Subscription, error) {

	var eventNameHashRule []interface{}
	for _, eventNameHashItem := range eventNameHash {
		eventNameHashRule = append(eventNameHashRule, eventNameHashItem)
	}
	var topic1Rule []interface{}
	for _, topic1Item := range topic1 {
		topic1Rule = append(topic1Rule, topic1Item)
	}
	var topic2Rule []interface{}
	for _, topic2Item := range topic2 {
		topic2Rule = append(topic2Rule, topic2Item)
	}

	logs, sub, err := _EventEmitter.contract.WatchLogs(opts, "EventLog2", eventNameHashRule, topic1Rule, topic2Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventEmitterEventLog2)
				if err := _EventEmitter.contract.UnpackLog(event, "EventLog2", log); err != nil {
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

// ParseEventLog2 is a log parse operation binding the contract event 0x468a25a7ba624ceea6e540ad6f49171b52495b648417ae91bca21676d8a24dc5.
//
// Solidity: event EventLog2(address msgSender, string eventName, string indexed eventNameHash, bytes32 indexed topic1, bytes32 indexed topic2, (((string,address)[],(string,address[])[]),((string,uint256)[],(string,uint256[])[]),((string,int256)[],(string,int256[])[]),((string,bool)[],(string,bool[])[]),((string,bytes32)[],(string,bytes32[])[]),((string,bytes)[],(string,bytes[])[]),((string,string)[],(string,string[])[])) eventData)
func (_EventEmitter *EventEmitterFilterer) ParseEventLog2(log types.Log) (*EventEmitterEventLog2, error) {
	event := new(EventEmitterEventLog2)
	if err := _EventEmitter.contract.UnpackLog(event, "EventLog2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

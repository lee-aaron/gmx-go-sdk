// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ExchangeRouter

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

// BaseOrderUtilsCreateOrderParams is an auto generated low-level Go binding around an user-defined struct.
type BaseOrderUtilsCreateOrderParams struct {
	Addresses                BaseOrderUtilsCreateOrderParamsAddresses
	Numbers                  BaseOrderUtilsCreateOrderParamsNumbers
	OrderType                uint8
	DecreasePositionSwapType uint8
	IsLong                   bool
	ShouldUnwrapNativeToken  bool
	ReferralCode             [32]byte
}

// BaseOrderUtilsCreateOrderParamsAddresses is an auto generated low-level Go binding around an user-defined struct.
type BaseOrderUtilsCreateOrderParamsAddresses struct {
	Receiver               common.Address
	CallbackContract       common.Address
	UiFeeReceiver          common.Address
	Market                 common.Address
	InitialCollateralToken common.Address
	SwapPath               []common.Address
}

// BaseOrderUtilsCreateOrderParamsNumbers is an auto generated low-level Go binding around an user-defined struct.
type BaseOrderUtilsCreateOrderParamsNumbers struct {
	SizeDeltaUsd                 *big.Int
	InitialCollateralDeltaAmount *big.Int
	TriggerPrice                 *big.Int
	AcceptablePrice              *big.Int
	ExecutionFee                 *big.Int
	CallbackGasLimit             *big.Int
	MinOutputAmount              *big.Int
}

// DepositUtilsCreateDepositParams is an auto generated low-level Go binding around an user-defined struct.
type DepositUtilsCreateDepositParams struct {
	Receiver                common.Address
	CallbackContract        common.Address
	UiFeeReceiver           common.Address
	Market                  common.Address
	InitialLongToken        common.Address
	InitialShortToken       common.Address
	LongTokenSwapPath       []common.Address
	ShortTokenSwapPath      []common.Address
	MinMarketTokens         *big.Int
	ShouldUnwrapNativeToken bool
	ExecutionFee            *big.Int
	CallbackGasLimit        *big.Int
}

// OracleUtilsSimulatePricesParams is an auto generated low-level Go binding around an user-defined struct.
type OracleUtilsSimulatePricesParams struct {
	PrimaryTokens []common.Address
	PrimaryPrices []PriceProps
}

// PriceProps is an auto generated low-level Go binding around an user-defined struct.
type PriceProps struct {
	Min *big.Int
	Max *big.Int
}

// WithdrawalUtilsCreateWithdrawalParams is an auto generated low-level Go binding around an user-defined struct.
type WithdrawalUtilsCreateWithdrawalParams struct {
	Receiver                common.Address
	CallbackContract        common.Address
	UiFeeReceiver           common.Address
	Market                  common.Address
	LongTokenSwapPath       []common.Address
	ShortTokenSwapPath      []common.Address
	MinLongTokenAmount      *big.Int
	MinShortTokenAmount     *big.Int
	ShouldUnwrapNativeToken bool
	ExecutionFee            *big.Int
	CallbackGasLimit        *big.Int
}

// ExchangeRouterMetaData contains all meta data concerning the ExchangeRouter contract.
var ExchangeRouterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_router\",\"type\":\"address\",\"internalType\":\"contractRouter\"},{\"name\":\"_roleStore\",\"type\":\"address\",\"internalType\":\"contractRoleStore\"},{\"name\":\"_dataStore\",\"type\":\"address\",\"internalType\":\"contractDataStore\"},{\"name\":\"_eventEmitter\",\"type\":\"address\",\"internalType\":\"contractEventEmitter\"},{\"name\":\"_depositHandler\",\"type\":\"address\",\"internalType\":\"contractIDepositHandler\"},{\"name\":\"_withdrawalHandler\",\"type\":\"address\",\"internalType\":\"contractIWithdrawalHandler\"},{\"name\":\"_orderHandler\",\"type\":\"address\",\"internalType\":\"contractIOrderHandler\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelDeposit\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"cancelOrder\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"cancelWithdrawal\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claimAffiliateRewards\",\"inputs\":[{\"name\":\"markets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claimCollateral\",\"inputs\":[{\"name\":\"markets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"timeKeys\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claimFundingFees\",\"inputs\":[{\"name\":\"markets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claimUiFees\",\"inputs\":[{\"name\":\"markets\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"createDeposit\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structDepositUtils.CreateDepositParams\",\"components\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callbackContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"uiFeeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"market\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialLongToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialShortToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"longTokenSwapPath\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"shortTokenSwapPath\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"minMarketTokens\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shouldUnwrapNativeToken\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"callbackGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"createOrder\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structBaseOrderUtils.CreateOrderParams\",\"components\":[{\"name\":\"addresses\",\"type\":\"tuple\",\"internalType\":\"structBaseOrderUtils.CreateOrderParamsAddresses\",\"components\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callbackContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"uiFeeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"market\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialCollateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"swapPath\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"numbers\",\"type\":\"tuple\",\"internalType\":\"structBaseOrderUtils.CreateOrderParamsNumbers\",\"components\":[{\"name\":\"sizeDeltaUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initialCollateralDeltaAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"triggerPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"acceptablePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"callbackGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minOutputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"orderType\",\"type\":\"uint8\",\"internalType\":\"enumOrder.OrderType\"},{\"name\":\"decreasePositionSwapType\",\"type\":\"uint8\",\"internalType\":\"enumOrder.DecreasePositionSwapType\"},{\"name\":\"isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"shouldUnwrapNativeToken\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"referralCode\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"createWithdrawal\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structWithdrawalUtils.CreateWithdrawalParams\",\"components\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callbackContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"uiFeeReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"market\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"longTokenSwapPath\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"shortTokenSwapPath\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"minLongTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minShortTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"shouldUnwrapNativeToken\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"callbackGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"dataStore\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractDataStore\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositHandler\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIDepositHandler\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eventEmitter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractEventEmitter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"multicall\",\"inputs\":[{\"name\":\"data\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[{\"name\":\"results\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"orderHandler\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOrderHandler\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"roleStore\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractRoleStore\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractRouter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sendNativeToken\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"sendTokens\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"sendWnt\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setSavedCallbackContract\",\"inputs\":[{\"name\":\"market\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callbackContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setUiFeeFactor\",\"inputs\":[{\"name\":\"uiFeeFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"simulateExecuteDeposit\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"simulatedOracleParams\",\"type\":\"tuple\",\"internalType\":\"structOracleUtils.SimulatePricesParams\",\"components\":[{\"name\":\"primaryTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"primaryPrices\",\"type\":\"tuple[]\",\"internalType\":\"structPrice.Props[]\",\"components\":[{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"simulateExecuteOrder\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"simulatedOracleParams\",\"type\":\"tuple\",\"internalType\":\"structOracleUtils.SimulatePricesParams\",\"components\":[{\"name\":\"primaryTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"primaryPrices\",\"type\":\"tuple[]\",\"internalType\":\"structPrice.Props[]\",\"components\":[{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"simulateExecuteWithdrawal\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"simulatedOracleParams\",\"type\":\"tuple\",\"internalType\":\"structOracleUtils.SimulatePricesParams\",\"components\":[{\"name\":\"primaryTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"primaryPrices\",\"type\":\"tuple[]\",\"internalType\":\"structPrice.Props[]\",\"components\":[{\"name\":\"min\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"updateOrder\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sizeDeltaUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"acceptablePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"triggerPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minOutputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdrawalHandler\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWithdrawalHandler\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"TokenTransferReverted\",\"inputs\":[{\"name\":\"reason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"returndata\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CollateralAlreadyClaimed\",\"inputs\":[{\"name\":\"adjustedClaimableAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"DisabledFeature\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"DisabledMarket\",\"inputs\":[{\"name\":\"market\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EmptyAddressInMarketTokenBalanceValidation\",\"inputs\":[{\"name\":\"market\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EmptyDeposit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyHoldingAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyMarket\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyReceiver\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyTokenTranferGasLimit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidClaimAffiliateRewardsInput\",\"inputs\":[{\"name\":\"marketsLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidClaimCollateralInput\",\"inputs\":[{\"name\":\"marketsLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"timeKeysLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidClaimFundingFeesInput\",\"inputs\":[{\"name\":\"marketsLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidClaimUiFeesInput\",\"inputs\":[{\"name\":\"marketsLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokensLength\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMarketTokenBalance\",\"inputs\":[{\"name\":\"market\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expectedMinBalance\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMarketTokenBalanceForClaimableFunding\",\"inputs\":[{\"name\":\"market\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"claimableFundingFeeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMarketTokenBalanceForCollateralAmount\",\"inputs\":[{\"name\":\"market\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidUiFeeFactor\",\"inputs\":[{\"name\":\"uiFeeFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxUiFeeFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TokenTransferError\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"msgSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"role\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
}

// ExchangeRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use ExchangeRouterMetaData.ABI instead.
var ExchangeRouterABI = ExchangeRouterMetaData.ABI

// ExchangeRouter is an auto generated Go binding around an Ethereum contract.
type ExchangeRouter struct {
	ExchangeRouterCaller     // Read-only binding to the contract
	ExchangeRouterTransactor // Write-only binding to the contract
	ExchangeRouterFilterer   // Log filterer for contract events
}

// ExchangeRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeRouterSession struct {
	Contract     *ExchangeRouter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeRouterCallerSession struct {
	Contract *ExchangeRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ExchangeRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeRouterTransactorSession struct {
	Contract     *ExchangeRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ExchangeRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRouterRaw struct {
	Contract *ExchangeRouter // Generic contract binding to access the raw methods on
}

// ExchangeRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeRouterCallerRaw struct {
	Contract *ExchangeRouterCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeRouterTransactorRaw struct {
	Contract *ExchangeRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchangeRouter creates a new instance of ExchangeRouter, bound to a specific deployed contract.
func NewExchangeRouter(address common.Address, backend bind.ContractBackend) (*ExchangeRouter, error) {
	contract, err := bindExchangeRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExchangeRouter{ExchangeRouterCaller: ExchangeRouterCaller{contract: contract}, ExchangeRouterTransactor: ExchangeRouterTransactor{contract: contract}, ExchangeRouterFilterer: ExchangeRouterFilterer{contract: contract}}, nil
}

// NewExchangeRouterCaller creates a new read-only instance of ExchangeRouter, bound to a specific deployed contract.
func NewExchangeRouterCaller(address common.Address, caller bind.ContractCaller) (*ExchangeRouterCaller, error) {
	contract, err := bindExchangeRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeRouterCaller{contract: contract}, nil
}

// NewExchangeRouterTransactor creates a new write-only instance of ExchangeRouter, bound to a specific deployed contract.
func NewExchangeRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeRouterTransactor, error) {
	contract, err := bindExchangeRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeRouterTransactor{contract: contract}, nil
}

// NewExchangeRouterFilterer creates a new log filterer instance of ExchangeRouter, bound to a specific deployed contract.
func NewExchangeRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeRouterFilterer, error) {
	contract, err := bindExchangeRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeRouterFilterer{contract: contract}, nil
}

// bindExchangeRouter binds a generic wrapper to an already deployed contract.
func bindExchangeRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExchangeRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeRouter *ExchangeRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExchangeRouter.Contract.ExchangeRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeRouter *ExchangeRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.ExchangeRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeRouter *ExchangeRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.ExchangeRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeRouter *ExchangeRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExchangeRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeRouter *ExchangeRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeRouter *ExchangeRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.contract.Transact(opts, method, params...)
}

// DataStore is a free data retrieval call binding the contract method 0x660d0d67.
//
// Solidity: function dataStore() view returns(address)
func (_ExchangeRouter *ExchangeRouterCaller) DataStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeRouter.contract.Call(opts, &out, "dataStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataStore is a free data retrieval call binding the contract method 0x660d0d67.
//
// Solidity: function dataStore() view returns(address)
func (_ExchangeRouter *ExchangeRouterSession) DataStore() (common.Address, error) {
	return _ExchangeRouter.Contract.DataStore(&_ExchangeRouter.CallOpts)
}

// DataStore is a free data retrieval call binding the contract method 0x660d0d67.
//
// Solidity: function dataStore() view returns(address)
func (_ExchangeRouter *ExchangeRouterCallerSession) DataStore() (common.Address, error) {
	return _ExchangeRouter.Contract.DataStore(&_ExchangeRouter.CallOpts)
}

// DepositHandler is a free data retrieval call binding the contract method 0x9c8b2cfb.
//
// Solidity: function depositHandler() view returns(address)
func (_ExchangeRouter *ExchangeRouterCaller) DepositHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeRouter.contract.Call(opts, &out, "depositHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositHandler is a free data retrieval call binding the contract method 0x9c8b2cfb.
//
// Solidity: function depositHandler() view returns(address)
func (_ExchangeRouter *ExchangeRouterSession) DepositHandler() (common.Address, error) {
	return _ExchangeRouter.Contract.DepositHandler(&_ExchangeRouter.CallOpts)
}

// DepositHandler is a free data retrieval call binding the contract method 0x9c8b2cfb.
//
// Solidity: function depositHandler() view returns(address)
func (_ExchangeRouter *ExchangeRouterCallerSession) DepositHandler() (common.Address, error) {
	return _ExchangeRouter.Contract.DepositHandler(&_ExchangeRouter.CallOpts)
}

// EventEmitter is a free data retrieval call binding the contract method 0x9ff78c30.
//
// Solidity: function eventEmitter() view returns(address)
func (_ExchangeRouter *ExchangeRouterCaller) EventEmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeRouter.contract.Call(opts, &out, "eventEmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EventEmitter is a free data retrieval call binding the contract method 0x9ff78c30.
//
// Solidity: function eventEmitter() view returns(address)
func (_ExchangeRouter *ExchangeRouterSession) EventEmitter() (common.Address, error) {
	return _ExchangeRouter.Contract.EventEmitter(&_ExchangeRouter.CallOpts)
}

// EventEmitter is a free data retrieval call binding the contract method 0x9ff78c30.
//
// Solidity: function eventEmitter() view returns(address)
func (_ExchangeRouter *ExchangeRouterCallerSession) EventEmitter() (common.Address, error) {
	return _ExchangeRouter.Contract.EventEmitter(&_ExchangeRouter.CallOpts)
}

// OrderHandler is a free data retrieval call binding the contract method 0xb5848305.
//
// Solidity: function orderHandler() view returns(address)
func (_ExchangeRouter *ExchangeRouterCaller) OrderHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeRouter.contract.Call(opts, &out, "orderHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OrderHandler is a free data retrieval call binding the contract method 0xb5848305.
//
// Solidity: function orderHandler() view returns(address)
func (_ExchangeRouter *ExchangeRouterSession) OrderHandler() (common.Address, error) {
	return _ExchangeRouter.Contract.OrderHandler(&_ExchangeRouter.CallOpts)
}

// OrderHandler is a free data retrieval call binding the contract method 0xb5848305.
//
// Solidity: function orderHandler() view returns(address)
func (_ExchangeRouter *ExchangeRouterCallerSession) OrderHandler() (common.Address, error) {
	return _ExchangeRouter.Contract.OrderHandler(&_ExchangeRouter.CallOpts)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_ExchangeRouter *ExchangeRouterCaller) RoleStore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeRouter.contract.Call(opts, &out, "roleStore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_ExchangeRouter *ExchangeRouterSession) RoleStore() (common.Address, error) {
	return _ExchangeRouter.Contract.RoleStore(&_ExchangeRouter.CallOpts)
}

// RoleStore is a free data retrieval call binding the contract method 0x4a4a7b04.
//
// Solidity: function roleStore() view returns(address)
func (_ExchangeRouter *ExchangeRouterCallerSession) RoleStore() (common.Address, error) {
	return _ExchangeRouter.Contract.RoleStore(&_ExchangeRouter.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_ExchangeRouter *ExchangeRouterCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeRouter.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_ExchangeRouter *ExchangeRouterSession) Router() (common.Address, error) {
	return _ExchangeRouter.Contract.Router(&_ExchangeRouter.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_ExchangeRouter *ExchangeRouterCallerSession) Router() (common.Address, error) {
	return _ExchangeRouter.Contract.Router(&_ExchangeRouter.CallOpts)
}

// WithdrawalHandler is a free data retrieval call binding the contract method 0x2c2f3c07.
//
// Solidity: function withdrawalHandler() view returns(address)
func (_ExchangeRouter *ExchangeRouterCaller) WithdrawalHandler(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExchangeRouter.contract.Call(opts, &out, "withdrawalHandler")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WithdrawalHandler is a free data retrieval call binding the contract method 0x2c2f3c07.
//
// Solidity: function withdrawalHandler() view returns(address)
func (_ExchangeRouter *ExchangeRouterSession) WithdrawalHandler() (common.Address, error) {
	return _ExchangeRouter.Contract.WithdrawalHandler(&_ExchangeRouter.CallOpts)
}

// WithdrawalHandler is a free data retrieval call binding the contract method 0x2c2f3c07.
//
// Solidity: function withdrawalHandler() view returns(address)
func (_ExchangeRouter *ExchangeRouterCallerSession) WithdrawalHandler() (common.Address, error) {
	return _ExchangeRouter.Contract.WithdrawalHandler(&_ExchangeRouter.CallOpts)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x31404484.
//
// Solidity: function cancelDeposit(bytes32 key) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactor) CancelDeposit(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "cancelDeposit", key)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x31404484.
//
// Solidity: function cancelDeposit(bytes32 key) payable returns()
func (_ExchangeRouter *ExchangeRouterSession) CancelDeposit(key [32]byte) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.CancelDeposit(&_ExchangeRouter.TransactOpts, key)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x31404484.
//
// Solidity: function cancelDeposit(bytes32 key) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactorSession) CancelDeposit(key [32]byte) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.CancelDeposit(&_ExchangeRouter.TransactOpts, key)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(bytes32 key) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactor) CancelOrder(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "cancelOrder", key)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(bytes32 key) payable returns()
func (_ExchangeRouter *ExchangeRouterSession) CancelOrder(key [32]byte) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.CancelOrder(&_ExchangeRouter.TransactOpts, key)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(bytes32 key) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactorSession) CancelOrder(key [32]byte) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.CancelOrder(&_ExchangeRouter.TransactOpts, key)
}

// CancelWithdrawal is a paid mutator transaction binding the contract method 0x7213c5a0.
//
// Solidity: function cancelWithdrawal(bytes32 key) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactor) CancelWithdrawal(opts *bind.TransactOpts, key [32]byte) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "cancelWithdrawal", key)
}

// CancelWithdrawal is a paid mutator transaction binding the contract method 0x7213c5a0.
//
// Solidity: function cancelWithdrawal(bytes32 key) payable returns()
func (_ExchangeRouter *ExchangeRouterSession) CancelWithdrawal(key [32]byte) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.CancelWithdrawal(&_ExchangeRouter.TransactOpts, key)
}

// CancelWithdrawal is a paid mutator transaction binding the contract method 0x7213c5a0.
//
// Solidity: function cancelWithdrawal(bytes32 key) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactorSession) CancelWithdrawal(key [32]byte) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.CancelWithdrawal(&_ExchangeRouter.TransactOpts, key)
}

// ClaimAffiliateRewards is a paid mutator transaction binding the contract method 0x49287a22.
//
// Solidity: function claimAffiliateRewards(address[] markets, address[] tokens, address receiver) payable returns(uint256[])
func (_ExchangeRouter *ExchangeRouterTransactor) ClaimAffiliateRewards(opts *bind.TransactOpts, markets []common.Address, tokens []common.Address, receiver common.Address) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "claimAffiliateRewards", markets, tokens, receiver)
}

// ClaimAffiliateRewards is a paid mutator transaction binding the contract method 0x49287a22.
//
// Solidity: function claimAffiliateRewards(address[] markets, address[] tokens, address receiver) payable returns(uint256[])
func (_ExchangeRouter *ExchangeRouterSession) ClaimAffiliateRewards(markets []common.Address, tokens []common.Address, receiver common.Address) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.ClaimAffiliateRewards(&_ExchangeRouter.TransactOpts, markets, tokens, receiver)
}

// ClaimAffiliateRewards is a paid mutator transaction binding the contract method 0x49287a22.
//
// Solidity: function claimAffiliateRewards(address[] markets, address[] tokens, address receiver) payable returns(uint256[])
func (_ExchangeRouter *ExchangeRouterTransactorSession) ClaimAffiliateRewards(markets []common.Address, tokens []common.Address, receiver common.Address) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.ClaimAffiliateRewards(&_ExchangeRouter.TransactOpts, markets, tokens, receiver)
}

// ClaimCollateral is a paid mutator transaction binding the contract method 0xe9249b57.
//
// Solidity: function claimCollateral(address[] markets, address[] tokens, uint256[] timeKeys, address receiver) payable returns(uint256[])
func (_ExchangeRouter *ExchangeRouterTransactor) ClaimCollateral(opts *bind.TransactOpts, markets []common.Address, tokens []common.Address, timeKeys []*big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "claimCollateral", markets, tokens, timeKeys, receiver)
}

// ClaimCollateral is a paid mutator transaction binding the contract method 0xe9249b57.
//
// Solidity: function claimCollateral(address[] markets, address[] tokens, uint256[] timeKeys, address receiver) payable returns(uint256[])
func (_ExchangeRouter *ExchangeRouterSession) ClaimCollateral(markets []common.Address, tokens []common.Address, timeKeys []*big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.ClaimCollateral(&_ExchangeRouter.TransactOpts, markets, tokens, timeKeys, receiver)
}

// ClaimCollateral is a paid mutator transaction binding the contract method 0xe9249b57.
//
// Solidity: function claimCollateral(address[] markets, address[] tokens, uint256[] timeKeys, address receiver) payable returns(uint256[])
func (_ExchangeRouter *ExchangeRouterTransactorSession) ClaimCollateral(markets []common.Address, tokens []common.Address, timeKeys []*big.Int, receiver common.Address) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.ClaimCollateral(&_ExchangeRouter.TransactOpts, markets, tokens, timeKeys, receiver)
}

// ClaimFundingFees is a paid mutator transaction binding the contract method 0xc41b1ab3.
//
// Solidity: function claimFundingFees(address[] markets, address[] tokens, address receiver) payable returns(uint256[])
func (_ExchangeRouter *ExchangeRouterTransactor) ClaimFundingFees(opts *bind.TransactOpts, markets []common.Address, tokens []common.Address, receiver common.Address) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "claimFundingFees", markets, tokens, receiver)
}

// ClaimFundingFees is a paid mutator transaction binding the contract method 0xc41b1ab3.
//
// Solidity: function claimFundingFees(address[] markets, address[] tokens, address receiver) payable returns(uint256[])
func (_ExchangeRouter *ExchangeRouterSession) ClaimFundingFees(markets []common.Address, tokens []common.Address, receiver common.Address) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.ClaimFundingFees(&_ExchangeRouter.TransactOpts, markets, tokens, receiver)
}

// ClaimFundingFees is a paid mutator transaction binding the contract method 0xc41b1ab3.
//
// Solidity: function claimFundingFees(address[] markets, address[] tokens, address receiver) payable returns(uint256[])
func (_ExchangeRouter *ExchangeRouterTransactorSession) ClaimFundingFees(markets []common.Address, tokens []common.Address, receiver common.Address) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.ClaimFundingFees(&_ExchangeRouter.TransactOpts, markets, tokens, receiver)
}

// ClaimUiFees is a paid mutator transaction binding the contract method 0x01a9cbb2.
//
// Solidity: function claimUiFees(address[] markets, address[] tokens, address receiver) payable returns(uint256[])
func (_ExchangeRouter *ExchangeRouterTransactor) ClaimUiFees(opts *bind.TransactOpts, markets []common.Address, tokens []common.Address, receiver common.Address) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "claimUiFees", markets, tokens, receiver)
}

// ClaimUiFees is a paid mutator transaction binding the contract method 0x01a9cbb2.
//
// Solidity: function claimUiFees(address[] markets, address[] tokens, address receiver) payable returns(uint256[])
func (_ExchangeRouter *ExchangeRouterSession) ClaimUiFees(markets []common.Address, tokens []common.Address, receiver common.Address) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.ClaimUiFees(&_ExchangeRouter.TransactOpts, markets, tokens, receiver)
}

// ClaimUiFees is a paid mutator transaction binding the contract method 0x01a9cbb2.
//
// Solidity: function claimUiFees(address[] markets, address[] tokens, address receiver) payable returns(uint256[])
func (_ExchangeRouter *ExchangeRouterTransactorSession) ClaimUiFees(markets []common.Address, tokens []common.Address, receiver common.Address) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.ClaimUiFees(&_ExchangeRouter.TransactOpts, markets, tokens, receiver)
}

// CreateDeposit is a paid mutator transaction binding the contract method 0x5b4e9561.
//
// Solidity: function createDeposit((address,address,address,address,address,address,address[],address[],uint256,bool,uint256,uint256) params) payable returns(bytes32)
func (_ExchangeRouter *ExchangeRouterTransactor) CreateDeposit(opts *bind.TransactOpts, params DepositUtilsCreateDepositParams) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "createDeposit", params)
}

// CreateDeposit is a paid mutator transaction binding the contract method 0x5b4e9561.
//
// Solidity: function createDeposit((address,address,address,address,address,address,address[],address[],uint256,bool,uint256,uint256) params) payable returns(bytes32)
func (_ExchangeRouter *ExchangeRouterSession) CreateDeposit(params DepositUtilsCreateDepositParams) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.CreateDeposit(&_ExchangeRouter.TransactOpts, params)
}

// CreateDeposit is a paid mutator transaction binding the contract method 0x5b4e9561.
//
// Solidity: function createDeposit((address,address,address,address,address,address,address[],address[],uint256,bool,uint256,uint256) params) payable returns(bytes32)
func (_ExchangeRouter *ExchangeRouterTransactorSession) CreateDeposit(params DepositUtilsCreateDepositParams) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.CreateDeposit(&_ExchangeRouter.TransactOpts, params)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x4a393a41.
//
// Solidity: function createOrder(((address,address,address,address,address,address[]),(uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint8,uint8,bool,bool,bytes32) params) payable returns(bytes32)
func (_ExchangeRouter *ExchangeRouterTransactor) CreateOrder(opts *bind.TransactOpts, params BaseOrderUtilsCreateOrderParams) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "createOrder", params)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x4a393a41.
//
// Solidity: function createOrder(((address,address,address,address,address,address[]),(uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint8,uint8,bool,bool,bytes32) params) payable returns(bytes32)
func (_ExchangeRouter *ExchangeRouterSession) CreateOrder(params BaseOrderUtilsCreateOrderParams) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.CreateOrder(&_ExchangeRouter.TransactOpts, params)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x4a393a41.
//
// Solidity: function createOrder(((address,address,address,address,address,address[]),(uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint8,uint8,bool,bool,bytes32) params) payable returns(bytes32)
func (_ExchangeRouter *ExchangeRouterTransactorSession) CreateOrder(params BaseOrderUtilsCreateOrderParams) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.CreateOrder(&_ExchangeRouter.TransactOpts, params)
}

// CreateWithdrawal is a paid mutator transaction binding the contract method 0xad23c5a1.
//
// Solidity: function createWithdrawal((address,address,address,address,address[],address[],uint256,uint256,bool,uint256,uint256) params) payable returns(bytes32)
func (_ExchangeRouter *ExchangeRouterTransactor) CreateWithdrawal(opts *bind.TransactOpts, params WithdrawalUtilsCreateWithdrawalParams) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "createWithdrawal", params)
}

// CreateWithdrawal is a paid mutator transaction binding the contract method 0xad23c5a1.
//
// Solidity: function createWithdrawal((address,address,address,address,address[],address[],uint256,uint256,bool,uint256,uint256) params) payable returns(bytes32)
func (_ExchangeRouter *ExchangeRouterSession) CreateWithdrawal(params WithdrawalUtilsCreateWithdrawalParams) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.CreateWithdrawal(&_ExchangeRouter.TransactOpts, params)
}

// CreateWithdrawal is a paid mutator transaction binding the contract method 0xad23c5a1.
//
// Solidity: function createWithdrawal((address,address,address,address,address[],address[],uint256,uint256,bool,uint256,uint256) params) payable returns(bytes32)
func (_ExchangeRouter *ExchangeRouterTransactorSession) CreateWithdrawal(params WithdrawalUtilsCreateWithdrawalParams) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.CreateWithdrawal(&_ExchangeRouter.TransactOpts, params)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_ExchangeRouter *ExchangeRouterTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_ExchangeRouter *ExchangeRouterSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.Multicall(&_ExchangeRouter.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_ExchangeRouter *ExchangeRouterTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.Multicall(&_ExchangeRouter.TransactOpts, data)
}

// SendNativeToken is a paid mutator transaction binding the contract method 0x53ead2d3.
//
// Solidity: function sendNativeToken(address receiver, uint256 amount) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactor) SendNativeToken(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "sendNativeToken", receiver, amount)
}

// SendNativeToken is a paid mutator transaction binding the contract method 0x53ead2d3.
//
// Solidity: function sendNativeToken(address receiver, uint256 amount) payable returns()
func (_ExchangeRouter *ExchangeRouterSession) SendNativeToken(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.SendNativeToken(&_ExchangeRouter.TransactOpts, receiver, amount)
}

// SendNativeToken is a paid mutator transaction binding the contract method 0x53ead2d3.
//
// Solidity: function sendNativeToken(address receiver, uint256 amount) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactorSession) SendNativeToken(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.SendNativeToken(&_ExchangeRouter.TransactOpts, receiver, amount)
}

// SendTokens is a paid mutator transaction binding the contract method 0xe6d66ac8.
//
// Solidity: function sendTokens(address token, address receiver, uint256 amount) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactor) SendTokens(opts *bind.TransactOpts, token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "sendTokens", token, receiver, amount)
}

// SendTokens is a paid mutator transaction binding the contract method 0xe6d66ac8.
//
// Solidity: function sendTokens(address token, address receiver, uint256 amount) payable returns()
func (_ExchangeRouter *ExchangeRouterSession) SendTokens(token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.SendTokens(&_ExchangeRouter.TransactOpts, token, receiver, amount)
}

// SendTokens is a paid mutator transaction binding the contract method 0xe6d66ac8.
//
// Solidity: function sendTokens(address token, address receiver, uint256 amount) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactorSession) SendTokens(token common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.SendTokens(&_ExchangeRouter.TransactOpts, token, receiver, amount)
}

// SendWnt is a paid mutator transaction binding the contract method 0x7d39aaf1.
//
// Solidity: function sendWnt(address receiver, uint256 amount) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactor) SendWnt(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "sendWnt", receiver, amount)
}

// SendWnt is a paid mutator transaction binding the contract method 0x7d39aaf1.
//
// Solidity: function sendWnt(address receiver, uint256 amount) payable returns()
func (_ExchangeRouter *ExchangeRouterSession) SendWnt(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.SendWnt(&_ExchangeRouter.TransactOpts, receiver, amount)
}

// SendWnt is a paid mutator transaction binding the contract method 0x7d39aaf1.
//
// Solidity: function sendWnt(address receiver, uint256 amount) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactorSession) SendWnt(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.SendWnt(&_ExchangeRouter.TransactOpts, receiver, amount)
}

// SetSavedCallbackContract is a paid mutator transaction binding the contract method 0x073fb09e.
//
// Solidity: function setSavedCallbackContract(address market, address callbackContract) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactor) SetSavedCallbackContract(opts *bind.TransactOpts, market common.Address, callbackContract common.Address) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "setSavedCallbackContract", market, callbackContract)
}

// SetSavedCallbackContract is a paid mutator transaction binding the contract method 0x073fb09e.
//
// Solidity: function setSavedCallbackContract(address market, address callbackContract) payable returns()
func (_ExchangeRouter *ExchangeRouterSession) SetSavedCallbackContract(market common.Address, callbackContract common.Address) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.SetSavedCallbackContract(&_ExchangeRouter.TransactOpts, market, callbackContract)
}

// SetSavedCallbackContract is a paid mutator transaction binding the contract method 0x073fb09e.
//
// Solidity: function setSavedCallbackContract(address market, address callbackContract) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactorSession) SetSavedCallbackContract(market common.Address, callbackContract common.Address) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.SetSavedCallbackContract(&_ExchangeRouter.TransactOpts, market, callbackContract)
}

// SetUiFeeFactor is a paid mutator transaction binding the contract method 0x5a03cd94.
//
// Solidity: function setUiFeeFactor(uint256 uiFeeFactor) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactor) SetUiFeeFactor(opts *bind.TransactOpts, uiFeeFactor *big.Int) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "setUiFeeFactor", uiFeeFactor)
}

// SetUiFeeFactor is a paid mutator transaction binding the contract method 0x5a03cd94.
//
// Solidity: function setUiFeeFactor(uint256 uiFeeFactor) payable returns()
func (_ExchangeRouter *ExchangeRouterSession) SetUiFeeFactor(uiFeeFactor *big.Int) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.SetUiFeeFactor(&_ExchangeRouter.TransactOpts, uiFeeFactor)
}

// SetUiFeeFactor is a paid mutator transaction binding the contract method 0x5a03cd94.
//
// Solidity: function setUiFeeFactor(uint256 uiFeeFactor) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactorSession) SetUiFeeFactor(uiFeeFactor *big.Int) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.SetUiFeeFactor(&_ExchangeRouter.TransactOpts, uiFeeFactor)
}

// SimulateExecuteDeposit is a paid mutator transaction binding the contract method 0xb9e2f5ee.
//
// Solidity: function simulateExecuteDeposit(bytes32 key, (address[],(uint256,uint256)[]) simulatedOracleParams) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactor) SimulateExecuteDeposit(opts *bind.TransactOpts, key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "simulateExecuteDeposit", key, simulatedOracleParams)
}

// SimulateExecuteDeposit is a paid mutator transaction binding the contract method 0xb9e2f5ee.
//
// Solidity: function simulateExecuteDeposit(bytes32 key, (address[],(uint256,uint256)[]) simulatedOracleParams) payable returns()
func (_ExchangeRouter *ExchangeRouterSession) SimulateExecuteDeposit(key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.SimulateExecuteDeposit(&_ExchangeRouter.TransactOpts, key, simulatedOracleParams)
}

// SimulateExecuteDeposit is a paid mutator transaction binding the contract method 0xb9e2f5ee.
//
// Solidity: function simulateExecuteDeposit(bytes32 key, (address[],(uint256,uint256)[]) simulatedOracleParams) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactorSession) SimulateExecuteDeposit(key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.SimulateExecuteDeposit(&_ExchangeRouter.TransactOpts, key, simulatedOracleParams)
}

// SimulateExecuteOrder is a paid mutator transaction binding the contract method 0x263ea0fa.
//
// Solidity: function simulateExecuteOrder(bytes32 key, (address[],(uint256,uint256)[]) simulatedOracleParams) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactor) SimulateExecuteOrder(opts *bind.TransactOpts, key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "simulateExecuteOrder", key, simulatedOracleParams)
}

// SimulateExecuteOrder is a paid mutator transaction binding the contract method 0x263ea0fa.
//
// Solidity: function simulateExecuteOrder(bytes32 key, (address[],(uint256,uint256)[]) simulatedOracleParams) payable returns()
func (_ExchangeRouter *ExchangeRouterSession) SimulateExecuteOrder(key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.SimulateExecuteOrder(&_ExchangeRouter.TransactOpts, key, simulatedOracleParams)
}

// SimulateExecuteOrder is a paid mutator transaction binding the contract method 0x263ea0fa.
//
// Solidity: function simulateExecuteOrder(bytes32 key, (address[],(uint256,uint256)[]) simulatedOracleParams) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactorSession) SimulateExecuteOrder(key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.SimulateExecuteOrder(&_ExchangeRouter.TransactOpts, key, simulatedOracleParams)
}

// SimulateExecuteWithdrawal is a paid mutator transaction binding the contract method 0x6331d7a7.
//
// Solidity: function simulateExecuteWithdrawal(bytes32 key, (address[],(uint256,uint256)[]) simulatedOracleParams) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactor) SimulateExecuteWithdrawal(opts *bind.TransactOpts, key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "simulateExecuteWithdrawal", key, simulatedOracleParams)
}

// SimulateExecuteWithdrawal is a paid mutator transaction binding the contract method 0x6331d7a7.
//
// Solidity: function simulateExecuteWithdrawal(bytes32 key, (address[],(uint256,uint256)[]) simulatedOracleParams) payable returns()
func (_ExchangeRouter *ExchangeRouterSession) SimulateExecuteWithdrawal(key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.SimulateExecuteWithdrawal(&_ExchangeRouter.TransactOpts, key, simulatedOracleParams)
}

// SimulateExecuteWithdrawal is a paid mutator transaction binding the contract method 0x6331d7a7.
//
// Solidity: function simulateExecuteWithdrawal(bytes32 key, (address[],(uint256,uint256)[]) simulatedOracleParams) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactorSession) SimulateExecuteWithdrawal(key [32]byte, simulatedOracleParams OracleUtilsSimulatePricesParams) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.SimulateExecuteWithdrawal(&_ExchangeRouter.TransactOpts, key, simulatedOracleParams)
}

// UpdateOrder is a paid mutator transaction binding the contract method 0xaab286f8.
//
// Solidity: function updateOrder(bytes32 key, uint256 sizeDeltaUsd, uint256 acceptablePrice, uint256 triggerPrice, uint256 minOutputAmount) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactor) UpdateOrder(opts *bind.TransactOpts, key [32]byte, sizeDeltaUsd *big.Int, acceptablePrice *big.Int, triggerPrice *big.Int, minOutputAmount *big.Int) (*types.Transaction, error) {
	return _ExchangeRouter.contract.Transact(opts, "updateOrder", key, sizeDeltaUsd, acceptablePrice, triggerPrice, minOutputAmount)
}

// UpdateOrder is a paid mutator transaction binding the contract method 0xaab286f8.
//
// Solidity: function updateOrder(bytes32 key, uint256 sizeDeltaUsd, uint256 acceptablePrice, uint256 triggerPrice, uint256 minOutputAmount) payable returns()
func (_ExchangeRouter *ExchangeRouterSession) UpdateOrder(key [32]byte, sizeDeltaUsd *big.Int, acceptablePrice *big.Int, triggerPrice *big.Int, minOutputAmount *big.Int) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.UpdateOrder(&_ExchangeRouter.TransactOpts, key, sizeDeltaUsd, acceptablePrice, triggerPrice, minOutputAmount)
}

// UpdateOrder is a paid mutator transaction binding the contract method 0xaab286f8.
//
// Solidity: function updateOrder(bytes32 key, uint256 sizeDeltaUsd, uint256 acceptablePrice, uint256 triggerPrice, uint256 minOutputAmount) payable returns()
func (_ExchangeRouter *ExchangeRouterTransactorSession) UpdateOrder(key [32]byte, sizeDeltaUsd *big.Int, acceptablePrice *big.Int, triggerPrice *big.Int, minOutputAmount *big.Int) (*types.Transaction, error) {
	return _ExchangeRouter.Contract.UpdateOrder(&_ExchangeRouter.TransactOpts, key, sizeDeltaUsd, acceptablePrice, triggerPrice, minOutputAmount)
}

// ExchangeRouterTokenTransferRevertedIterator is returned from FilterTokenTransferReverted and is used to iterate over the raw logs and unpacked data for TokenTransferReverted events raised by the ExchangeRouter contract.
type ExchangeRouterTokenTransferRevertedIterator struct {
	Event *ExchangeRouterTokenTransferReverted // Event containing the contract specifics and raw log

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
func (it *ExchangeRouterTokenTransferRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeRouterTokenTransferReverted)
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
		it.Event = new(ExchangeRouterTokenTransferReverted)
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
func (it *ExchangeRouterTokenTransferRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeRouterTokenTransferRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeRouterTokenTransferReverted represents a TokenTransferReverted event raised by the ExchangeRouter contract.
type ExchangeRouterTokenTransferReverted struct {
	Reason     string
	Returndata []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenTransferReverted is a free log retrieval operation binding the contract event 0xc9f14d9a0a9b46470c7c0b6c508f8283abaab7f795f153953c58cd4250824dae.
//
// Solidity: event TokenTransferReverted(string reason, bytes returndata)
func (_ExchangeRouter *ExchangeRouterFilterer) FilterTokenTransferReverted(opts *bind.FilterOpts) (*ExchangeRouterTokenTransferRevertedIterator, error) {

	logs, sub, err := _ExchangeRouter.contract.FilterLogs(opts, "TokenTransferReverted")
	if err != nil {
		return nil, err
	}
	return &ExchangeRouterTokenTransferRevertedIterator{contract: _ExchangeRouter.contract, event: "TokenTransferReverted", logs: logs, sub: sub}, nil
}

// WatchTokenTransferReverted is a free log subscription operation binding the contract event 0xc9f14d9a0a9b46470c7c0b6c508f8283abaab7f795f153953c58cd4250824dae.
//
// Solidity: event TokenTransferReverted(string reason, bytes returndata)
func (_ExchangeRouter *ExchangeRouterFilterer) WatchTokenTransferReverted(opts *bind.WatchOpts, sink chan<- *ExchangeRouterTokenTransferReverted) (event.Subscription, error) {

	logs, sub, err := _ExchangeRouter.contract.WatchLogs(opts, "TokenTransferReverted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeRouterTokenTransferReverted)
				if err := _ExchangeRouter.contract.UnpackLog(event, "TokenTransferReverted", log); err != nil {
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

// ParseTokenTransferReverted is a log parse operation binding the contract event 0xc9f14d9a0a9b46470c7c0b6c508f8283abaab7f795f153953c58cd4250824dae.
//
// Solidity: event TokenTransferReverted(string reason, bytes returndata)
func (_ExchangeRouter *ExchangeRouterFilterer) ParseTokenTransferReverted(log types.Log) (*ExchangeRouterTokenTransferReverted, error) {
	event := new(ExchangeRouterTokenTransferReverted)
	if err := _ExchangeRouter.contract.UnpackLog(event, "TokenTransferReverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

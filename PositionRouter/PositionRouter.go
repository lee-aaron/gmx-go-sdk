// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PositionRouter

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

// PositionRouterMetaData contains all meta data concerning the PositionRouter contract.
var PositionRouterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_shortsTracker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_depositFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minExecutionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"BASIS_POINTS_DIVISOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"callbackGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelDecreasePosition\",\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_executionFeeReceiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelIncreasePosition\",\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_executionFeeReceiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createDecreasePosition\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_acceptablePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_withdrawETH\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_callbackTarget\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"createIncreasePosition\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_acceptablePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_referralCode\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_callbackTarget\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"createIncreasePositionETH\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_acceptablePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_referralCode\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_callbackTarget\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"customCallbackGasLimits\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreasePositionRequestKeys\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreasePositionRequestKeysStart\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreasePositionRequests\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"acceptablePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"withdrawETH\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"callbackTarget\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreasePositionsIndex\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ethTransferGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeDecreasePosition\",\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_executionFeeReceiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeDecreasePositions\",\"inputs\":[{\"name\":\"_endIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_executionFeeReceiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeIncreasePosition\",\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_executionFeeReceiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeIncreasePositions\",\"inputs\":[{\"name\":\"_endIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_executionFeeReceiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeReserves\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDecreasePositionRequestPath\",\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIncreasePositionRequestPath\",\"inputs\":[{\"name\":\"_key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequestKey\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getRequestQueueLengths\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increasePositionBufferBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increasePositionRequestKeys\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increasePositionRequestKeysStart\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increasePositionRequests\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"acceptablePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"blockTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hasCollateralInETH\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"callbackTarget\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increasePositionsIndex\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isLeverageEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPositionKeeper\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxGlobalLongSizes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxGlobalShortSizes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxTimeDelay\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minBlockDelayKeeper\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minExecutionFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minTimeDelayPublic\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"referralStorage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sendValue\",\"inputs\":[{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setAdmin\",\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCallbackGasLimit\",\"inputs\":[{\"name\":\"_callbackGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCustomCallbackGasLimit\",\"inputs\":[{\"name\":\"_callbackTarget\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_callbackGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDelayValues\",\"inputs\":[{\"name\":\"_minBlockDelayKeeper\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minTimeDelayPublic\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_maxTimeDelay\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDepositFee\",\"inputs\":[{\"name\":\"_depositFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEthTransferGasLimit\",\"inputs\":[{\"name\":\"_ethTransferGasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIncreasePositionBufferBps\",\"inputs\":[{\"name\":\"_increasePositionBufferBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIsLeverageEnabled\",\"inputs\":[{\"name\":\"_isLeverageEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxGlobalSizes\",\"inputs\":[{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_longSizes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_shortSizes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinExecutionFee\",\"inputs\":[{\"name\":\"_minExecutionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPositionKeeper\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setReferralStorage\",\"inputs\":[{\"name\":\"_referralStorage\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRequestKeysStartValues\",\"inputs\":[{\"name\":\"_increasePositionRequestKeysStart\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_decreasePositionRequestKeysStart\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shortsTracker\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vault\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"weth\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Callback\",\"inputs\":[{\"name\":\"callbackTarget\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"callbackGasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CancelDecreasePosition\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"path\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"acceptablePrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"minOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"blockGap\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timeGap\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CancelIncreasePosition\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"path\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"minOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"acceptablePrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"blockGap\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timeGap\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CreateDecreasePosition\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"path\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"acceptablePrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"minOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"queueIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"blockTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CreateIncreasePosition\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"path\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"minOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"acceptablePrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"queueIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"blockTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecreasePositionReferral\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"marginFeeBasisPoints\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"referralCode\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"referrer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecuteDecreasePosition\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"path\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"acceptablePrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"minOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"blockGap\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timeGap\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecuteIncreasePosition\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"path\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"minOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"acceptablePrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"blockGap\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timeGap\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncreasePositionReferral\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"marginFeeBasisPoints\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"referralCode\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"referrer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetAdmin\",\"inputs\":[{\"name\":\"admin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetCallbackGasLimit\",\"inputs\":[{\"name\":\"callbackGasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetCustomCallbackGasLimit\",\"inputs\":[{\"name\":\"callbackTarget\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"callbackGasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetDelayValues\",\"inputs\":[{\"name\":\"minBlockDelayKeeper\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"minTimeDelayPublic\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"maxTimeDelay\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetDepositFee\",\"inputs\":[{\"name\":\"depositFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetEthTransferGasLimit\",\"inputs\":[{\"name\":\"ethTransferGasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetIncreasePositionBufferBps\",\"inputs\":[{\"name\":\"increasePositionBufferBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetIsLeverageEnabled\",\"inputs\":[{\"name\":\"isLeverageEnabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetMaxGlobalSizes\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"longSizes\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"shortSizes\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetMinExecutionFee\",\"inputs\":[{\"name\":\"minExecutionFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetPositionKeeper\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"isActive\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetReferralStorage\",\"inputs\":[{\"name\":\"referralStorage\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetRequestKeysStartValues\",\"inputs\":[{\"name\":\"increasePositionRequestKeysStart\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decreasePositionRequestKeysStart\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawFees\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// PositionRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use PositionRouterMetaData.ABI instead.
var PositionRouterABI = PositionRouterMetaData.ABI

// PositionRouter is an auto generated Go binding around an Ethereum contract.
type PositionRouter struct {
	PositionRouterCaller     // Read-only binding to the contract
	PositionRouterTransactor // Write-only binding to the contract
	PositionRouterFilterer   // Log filterer for contract events
}

// PositionRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PositionRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PositionRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PositionRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PositionRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PositionRouterSession struct {
	Contract     *PositionRouter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PositionRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PositionRouterCallerSession struct {
	Contract *PositionRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PositionRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PositionRouterTransactorSession struct {
	Contract     *PositionRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PositionRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PositionRouterRaw struct {
	Contract *PositionRouter // Generic contract binding to access the raw methods on
}

// PositionRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PositionRouterCallerRaw struct {
	Contract *PositionRouterCaller // Generic read-only contract binding to access the raw methods on
}

// PositionRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PositionRouterTransactorRaw struct {
	Contract *PositionRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPositionRouter creates a new instance of PositionRouter, bound to a specific deployed contract.
func NewPositionRouter(address common.Address, backend bind.ContractBackend) (*PositionRouter, error) {
	contract, err := bindPositionRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PositionRouter{PositionRouterCaller: PositionRouterCaller{contract: contract}, PositionRouterTransactor: PositionRouterTransactor{contract: contract}, PositionRouterFilterer: PositionRouterFilterer{contract: contract}}, nil
}

// NewPositionRouterCaller creates a new read-only instance of PositionRouter, bound to a specific deployed contract.
func NewPositionRouterCaller(address common.Address, caller bind.ContractCaller) (*PositionRouterCaller, error) {
	contract, err := bindPositionRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PositionRouterCaller{contract: contract}, nil
}

// NewPositionRouterTransactor creates a new write-only instance of PositionRouter, bound to a specific deployed contract.
func NewPositionRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*PositionRouterTransactor, error) {
	contract, err := bindPositionRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PositionRouterTransactor{contract: contract}, nil
}

// NewPositionRouterFilterer creates a new log filterer instance of PositionRouter, bound to a specific deployed contract.
func NewPositionRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*PositionRouterFilterer, error) {
	contract, err := bindPositionRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PositionRouterFilterer{contract: contract}, nil
}

// bindPositionRouter binds a generic wrapper to an already deployed contract.
func bindPositionRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PositionRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PositionRouter *PositionRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PositionRouter.Contract.PositionRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PositionRouter *PositionRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionRouter.Contract.PositionRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PositionRouter *PositionRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PositionRouter.Contract.PositionRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PositionRouter *PositionRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PositionRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PositionRouter *PositionRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PositionRouter *PositionRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PositionRouter.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_PositionRouter *PositionRouterCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_PositionRouter *PositionRouterSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _PositionRouter.Contract.BASISPOINTSDIVISOR(&_PositionRouter.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_PositionRouter *PositionRouterCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _PositionRouter.Contract.BASISPOINTSDIVISOR(&_PositionRouter.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PositionRouter *PositionRouterCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PositionRouter *PositionRouterSession) Admin() (common.Address, error) {
	return _PositionRouter.Contract.Admin(&_PositionRouter.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PositionRouter *PositionRouterCallerSession) Admin() (common.Address, error) {
	return _PositionRouter.Contract.Admin(&_PositionRouter.CallOpts)
}

// CallbackGasLimit is a free data retrieval call binding the contract method 0x24f74697.
//
// Solidity: function callbackGasLimit() view returns(uint256)
func (_PositionRouter *PositionRouterCaller) CallbackGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "callbackGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CallbackGasLimit is a free data retrieval call binding the contract method 0x24f74697.
//
// Solidity: function callbackGasLimit() view returns(uint256)
func (_PositionRouter *PositionRouterSession) CallbackGasLimit() (*big.Int, error) {
	return _PositionRouter.Contract.CallbackGasLimit(&_PositionRouter.CallOpts)
}

// CallbackGasLimit is a free data retrieval call binding the contract method 0x24f74697.
//
// Solidity: function callbackGasLimit() view returns(uint256)
func (_PositionRouter *PositionRouterCallerSession) CallbackGasLimit() (*big.Int, error) {
	return _PositionRouter.Contract.CallbackGasLimit(&_PositionRouter.CallOpts)
}

// CustomCallbackGasLimits is a free data retrieval call binding the contract method 0xc4f38e33.
//
// Solidity: function customCallbackGasLimits(address ) view returns(uint256)
func (_PositionRouter *PositionRouterCaller) CustomCallbackGasLimits(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "customCallbackGasLimits", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CustomCallbackGasLimits is a free data retrieval call binding the contract method 0xc4f38e33.
//
// Solidity: function customCallbackGasLimits(address ) view returns(uint256)
func (_PositionRouter *PositionRouterSession) CustomCallbackGasLimits(arg0 common.Address) (*big.Int, error) {
	return _PositionRouter.Contract.CustomCallbackGasLimits(&_PositionRouter.CallOpts, arg0)
}

// CustomCallbackGasLimits is a free data retrieval call binding the contract method 0xc4f38e33.
//
// Solidity: function customCallbackGasLimits(address ) view returns(uint256)
func (_PositionRouter *PositionRouterCallerSession) CustomCallbackGasLimits(arg0 common.Address) (*big.Int, error) {
	return _PositionRouter.Contract.CustomCallbackGasLimits(&_PositionRouter.CallOpts, arg0)
}

// DecreasePositionRequestKeys is a free data retrieval call binding the contract method 0x4278555f.
//
// Solidity: function decreasePositionRequestKeys(uint256 ) view returns(bytes32)
func (_PositionRouter *PositionRouterCaller) DecreasePositionRequestKeys(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "decreasePositionRequestKeys", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DecreasePositionRequestKeys is a free data retrieval call binding the contract method 0x4278555f.
//
// Solidity: function decreasePositionRequestKeys(uint256 ) view returns(bytes32)
func (_PositionRouter *PositionRouterSession) DecreasePositionRequestKeys(arg0 *big.Int) ([32]byte, error) {
	return _PositionRouter.Contract.DecreasePositionRequestKeys(&_PositionRouter.CallOpts, arg0)
}

// DecreasePositionRequestKeys is a free data retrieval call binding the contract method 0x4278555f.
//
// Solidity: function decreasePositionRequestKeys(uint256 ) view returns(bytes32)
func (_PositionRouter *PositionRouterCallerSession) DecreasePositionRequestKeys(arg0 *big.Int) ([32]byte, error) {
	return _PositionRouter.Contract.DecreasePositionRequestKeys(&_PositionRouter.CallOpts, arg0)
}

// DecreasePositionRequestKeysStart is a free data retrieval call binding the contract method 0x1bca8cf0.
//
// Solidity: function decreasePositionRequestKeysStart() view returns(uint256)
func (_PositionRouter *PositionRouterCaller) DecreasePositionRequestKeysStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "decreasePositionRequestKeysStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecreasePositionRequestKeysStart is a free data retrieval call binding the contract method 0x1bca8cf0.
//
// Solidity: function decreasePositionRequestKeysStart() view returns(uint256)
func (_PositionRouter *PositionRouterSession) DecreasePositionRequestKeysStart() (*big.Int, error) {
	return _PositionRouter.Contract.DecreasePositionRequestKeysStart(&_PositionRouter.CallOpts)
}

// DecreasePositionRequestKeysStart is a free data retrieval call binding the contract method 0x1bca8cf0.
//
// Solidity: function decreasePositionRequestKeysStart() view returns(uint256)
func (_PositionRouter *PositionRouterCallerSession) DecreasePositionRequestKeysStart() (*big.Int, error) {
	return _PositionRouter.Contract.DecreasePositionRequestKeysStart(&_PositionRouter.CallOpts)
}

// DecreasePositionRequests is a free data retrieval call binding the contract method 0x1f285106.
//
// Solidity: function decreasePositionRequests(bytes32 ) view returns(address account, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, address receiver, uint256 acceptablePrice, uint256 minOut, uint256 executionFee, uint256 blockNumber, uint256 blockTime, bool withdrawETH, address callbackTarget)
func (_PositionRouter *PositionRouterCaller) DecreasePositionRequests(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Account         common.Address
	IndexToken      common.Address
	CollateralDelta *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	Receiver        common.Address
	AcceptablePrice *big.Int
	MinOut          *big.Int
	ExecutionFee    *big.Int
	BlockNumber     *big.Int
	BlockTime       *big.Int
	WithdrawETH     bool
	CallbackTarget  common.Address
}, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "decreasePositionRequests", arg0)

	outstruct := new(struct {
		Account         common.Address
		IndexToken      common.Address
		CollateralDelta *big.Int
		SizeDelta       *big.Int
		IsLong          bool
		Receiver        common.Address
		AcceptablePrice *big.Int
		MinOut          *big.Int
		ExecutionFee    *big.Int
		BlockNumber     *big.Int
		BlockTime       *big.Int
		WithdrawETH     bool
		CallbackTarget  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Account = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IndexToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.CollateralDelta = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SizeDelta = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsLong = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Receiver = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.AcceptablePrice = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.MinOut = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.ExecutionFee = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.BlockNumber = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.BlockTime = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.WithdrawETH = *abi.ConvertType(out[11], new(bool)).(*bool)
	outstruct.CallbackTarget = *abi.ConvertType(out[12], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// DecreasePositionRequests is a free data retrieval call binding the contract method 0x1f285106.
//
// Solidity: function decreasePositionRequests(bytes32 ) view returns(address account, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, address receiver, uint256 acceptablePrice, uint256 minOut, uint256 executionFee, uint256 blockNumber, uint256 blockTime, bool withdrawETH, address callbackTarget)
func (_PositionRouter *PositionRouterSession) DecreasePositionRequests(arg0 [32]byte) (struct {
	Account         common.Address
	IndexToken      common.Address
	CollateralDelta *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	Receiver        common.Address
	AcceptablePrice *big.Int
	MinOut          *big.Int
	ExecutionFee    *big.Int
	BlockNumber     *big.Int
	BlockTime       *big.Int
	WithdrawETH     bool
	CallbackTarget  common.Address
}, error) {
	return _PositionRouter.Contract.DecreasePositionRequests(&_PositionRouter.CallOpts, arg0)
}

// DecreasePositionRequests is a free data retrieval call binding the contract method 0x1f285106.
//
// Solidity: function decreasePositionRequests(bytes32 ) view returns(address account, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, address receiver, uint256 acceptablePrice, uint256 minOut, uint256 executionFee, uint256 blockNumber, uint256 blockTime, bool withdrawETH, address callbackTarget)
func (_PositionRouter *PositionRouterCallerSession) DecreasePositionRequests(arg0 [32]byte) (struct {
	Account         common.Address
	IndexToken      common.Address
	CollateralDelta *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	Receiver        common.Address
	AcceptablePrice *big.Int
	MinOut          *big.Int
	ExecutionFee    *big.Int
	BlockNumber     *big.Int
	BlockTime       *big.Int
	WithdrawETH     bool
	CallbackTarget  common.Address
}, error) {
	return _PositionRouter.Contract.DecreasePositionRequests(&_PositionRouter.CallOpts, arg0)
}

// DecreasePositionsIndex is a free data retrieval call binding the contract method 0xfa444577.
//
// Solidity: function decreasePositionsIndex(address ) view returns(uint256)
func (_PositionRouter *PositionRouterCaller) DecreasePositionsIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "decreasePositionsIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecreasePositionsIndex is a free data retrieval call binding the contract method 0xfa444577.
//
// Solidity: function decreasePositionsIndex(address ) view returns(uint256)
func (_PositionRouter *PositionRouterSession) DecreasePositionsIndex(arg0 common.Address) (*big.Int, error) {
	return _PositionRouter.Contract.DecreasePositionsIndex(&_PositionRouter.CallOpts, arg0)
}

// DecreasePositionsIndex is a free data retrieval call binding the contract method 0xfa444577.
//
// Solidity: function decreasePositionsIndex(address ) view returns(uint256)
func (_PositionRouter *PositionRouterCallerSession) DecreasePositionsIndex(arg0 common.Address) (*big.Int, error) {
	return _PositionRouter.Contract.DecreasePositionsIndex(&_PositionRouter.CallOpts, arg0)
}

// DepositFee is a free data retrieval call binding the contract method 0x67a52793.
//
// Solidity: function depositFee() view returns(uint256)
func (_PositionRouter *PositionRouterCaller) DepositFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "depositFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositFee is a free data retrieval call binding the contract method 0x67a52793.
//
// Solidity: function depositFee() view returns(uint256)
func (_PositionRouter *PositionRouterSession) DepositFee() (*big.Int, error) {
	return _PositionRouter.Contract.DepositFee(&_PositionRouter.CallOpts)
}

// DepositFee is a free data retrieval call binding the contract method 0x67a52793.
//
// Solidity: function depositFee() view returns(uint256)
func (_PositionRouter *PositionRouterCallerSession) DepositFee() (*big.Int, error) {
	return _PositionRouter.Contract.DepositFee(&_PositionRouter.CallOpts)
}

// EthTransferGasLimit is a free data retrieval call binding the contract method 0x2d79cf42.
//
// Solidity: function ethTransferGasLimit() view returns(uint256)
func (_PositionRouter *PositionRouterCaller) EthTransferGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "ethTransferGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthTransferGasLimit is a free data retrieval call binding the contract method 0x2d79cf42.
//
// Solidity: function ethTransferGasLimit() view returns(uint256)
func (_PositionRouter *PositionRouterSession) EthTransferGasLimit() (*big.Int, error) {
	return _PositionRouter.Contract.EthTransferGasLimit(&_PositionRouter.CallOpts)
}

// EthTransferGasLimit is a free data retrieval call binding the contract method 0x2d79cf42.
//
// Solidity: function ethTransferGasLimit() view returns(uint256)
func (_PositionRouter *PositionRouterCallerSession) EthTransferGasLimit() (*big.Int, error) {
	return _PositionRouter.Contract.EthTransferGasLimit(&_PositionRouter.CallOpts)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_PositionRouter *PositionRouterCaller) FeeReserves(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "feeReserves", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_PositionRouter *PositionRouterSession) FeeReserves(arg0 common.Address) (*big.Int, error) {
	return _PositionRouter.Contract.FeeReserves(&_PositionRouter.CallOpts, arg0)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_PositionRouter *PositionRouterCallerSession) FeeReserves(arg0 common.Address) (*big.Int, error) {
	return _PositionRouter.Contract.FeeReserves(&_PositionRouter.CallOpts, arg0)
}

// GetDecreasePositionRequestPath is a free data retrieval call binding the contract method 0x5d5c22e8.
//
// Solidity: function getDecreasePositionRequestPath(bytes32 _key) view returns(address[])
func (_PositionRouter *PositionRouterCaller) GetDecreasePositionRequestPath(opts *bind.CallOpts, _key [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "getDecreasePositionRequestPath", _key)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDecreasePositionRequestPath is a free data retrieval call binding the contract method 0x5d5c22e8.
//
// Solidity: function getDecreasePositionRequestPath(bytes32 _key) view returns(address[])
func (_PositionRouter *PositionRouterSession) GetDecreasePositionRequestPath(_key [32]byte) ([]common.Address, error) {
	return _PositionRouter.Contract.GetDecreasePositionRequestPath(&_PositionRouter.CallOpts, _key)
}

// GetDecreasePositionRequestPath is a free data retrieval call binding the contract method 0x5d5c22e8.
//
// Solidity: function getDecreasePositionRequestPath(bytes32 _key) view returns(address[])
func (_PositionRouter *PositionRouterCallerSession) GetDecreasePositionRequestPath(_key [32]byte) ([]common.Address, error) {
	return _PositionRouter.Contract.GetDecreasePositionRequestPath(&_PositionRouter.CallOpts, _key)
}

// GetIncreasePositionRequestPath is a free data retrieval call binding the contract method 0x95e9bbd7.
//
// Solidity: function getIncreasePositionRequestPath(bytes32 _key) view returns(address[])
func (_PositionRouter *PositionRouterCaller) GetIncreasePositionRequestPath(opts *bind.CallOpts, _key [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "getIncreasePositionRequestPath", _key)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetIncreasePositionRequestPath is a free data retrieval call binding the contract method 0x95e9bbd7.
//
// Solidity: function getIncreasePositionRequestPath(bytes32 _key) view returns(address[])
func (_PositionRouter *PositionRouterSession) GetIncreasePositionRequestPath(_key [32]byte) ([]common.Address, error) {
	return _PositionRouter.Contract.GetIncreasePositionRequestPath(&_PositionRouter.CallOpts, _key)
}

// GetIncreasePositionRequestPath is a free data retrieval call binding the contract method 0x95e9bbd7.
//
// Solidity: function getIncreasePositionRequestPath(bytes32 _key) view returns(address[])
func (_PositionRouter *PositionRouterCallerSession) GetIncreasePositionRequestPath(_key [32]byte) ([]common.Address, error) {
	return _PositionRouter.Contract.GetIncreasePositionRequestPath(&_PositionRouter.CallOpts, _key)
}

// GetRequestKey is a free data retrieval call binding the contract method 0x62f8a3fe.
//
// Solidity: function getRequestKey(address _account, uint256 _index) pure returns(bytes32)
func (_PositionRouter *PositionRouterCaller) GetRequestKey(opts *bind.CallOpts, _account common.Address, _index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "getRequestKey", _account, _index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRequestKey is a free data retrieval call binding the contract method 0x62f8a3fe.
//
// Solidity: function getRequestKey(address _account, uint256 _index) pure returns(bytes32)
func (_PositionRouter *PositionRouterSession) GetRequestKey(_account common.Address, _index *big.Int) ([32]byte, error) {
	return _PositionRouter.Contract.GetRequestKey(&_PositionRouter.CallOpts, _account, _index)
}

// GetRequestKey is a free data retrieval call binding the contract method 0x62f8a3fe.
//
// Solidity: function getRequestKey(address _account, uint256 _index) pure returns(bytes32)
func (_PositionRouter *PositionRouterCallerSession) GetRequestKey(_account common.Address, _index *big.Int) ([32]byte, error) {
	return _PositionRouter.Contract.GetRequestKey(&_PositionRouter.CallOpts, _account, _index)
}

// GetRequestQueueLengths is a free data retrieval call binding the contract method 0xf2cea6a5.
//
// Solidity: function getRequestQueueLengths() view returns(uint256, uint256, uint256, uint256)
func (_PositionRouter *PositionRouterCaller) GetRequestQueueLengths(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "getRequestQueueLengths")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetRequestQueueLengths is a free data retrieval call binding the contract method 0xf2cea6a5.
//
// Solidity: function getRequestQueueLengths() view returns(uint256, uint256, uint256, uint256)
func (_PositionRouter *PositionRouterSession) GetRequestQueueLengths() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _PositionRouter.Contract.GetRequestQueueLengths(&_PositionRouter.CallOpts)
}

// GetRequestQueueLengths is a free data retrieval call binding the contract method 0xf2cea6a5.
//
// Solidity: function getRequestQueueLengths() view returns(uint256, uint256, uint256, uint256)
func (_PositionRouter *PositionRouterCallerSession) GetRequestQueueLengths() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _PositionRouter.Contract.GetRequestQueueLengths(&_PositionRouter.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_PositionRouter *PositionRouterCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_PositionRouter *PositionRouterSession) Gov() (common.Address, error) {
	return _PositionRouter.Contract.Gov(&_PositionRouter.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_PositionRouter *PositionRouterCallerSession) Gov() (common.Address, error) {
	return _PositionRouter.Contract.Gov(&_PositionRouter.CallOpts)
}

// IncreasePositionBufferBps is a free data retrieval call binding the contract method 0x98d1e03a.
//
// Solidity: function increasePositionBufferBps() view returns(uint256)
func (_PositionRouter *PositionRouterCaller) IncreasePositionBufferBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "increasePositionBufferBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncreasePositionBufferBps is a free data retrieval call binding the contract method 0x98d1e03a.
//
// Solidity: function increasePositionBufferBps() view returns(uint256)
func (_PositionRouter *PositionRouterSession) IncreasePositionBufferBps() (*big.Int, error) {
	return _PositionRouter.Contract.IncreasePositionBufferBps(&_PositionRouter.CallOpts)
}

// IncreasePositionBufferBps is a free data retrieval call binding the contract method 0x98d1e03a.
//
// Solidity: function increasePositionBufferBps() view returns(uint256)
func (_PositionRouter *PositionRouterCallerSession) IncreasePositionBufferBps() (*big.Int, error) {
	return _PositionRouter.Contract.IncreasePositionBufferBps(&_PositionRouter.CallOpts)
}

// IncreasePositionRequestKeys is a free data retrieval call binding the contract method 0x04225954.
//
// Solidity: function increasePositionRequestKeys(uint256 ) view returns(bytes32)
func (_PositionRouter *PositionRouterCaller) IncreasePositionRequestKeys(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "increasePositionRequestKeys", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IncreasePositionRequestKeys is a free data retrieval call binding the contract method 0x04225954.
//
// Solidity: function increasePositionRequestKeys(uint256 ) view returns(bytes32)
func (_PositionRouter *PositionRouterSession) IncreasePositionRequestKeys(arg0 *big.Int) ([32]byte, error) {
	return _PositionRouter.Contract.IncreasePositionRequestKeys(&_PositionRouter.CallOpts, arg0)
}

// IncreasePositionRequestKeys is a free data retrieval call binding the contract method 0x04225954.
//
// Solidity: function increasePositionRequestKeys(uint256 ) view returns(bytes32)
func (_PositionRouter *PositionRouterCallerSession) IncreasePositionRequestKeys(arg0 *big.Int) ([32]byte, error) {
	return _PositionRouter.Contract.IncreasePositionRequestKeys(&_PositionRouter.CallOpts, arg0)
}

// IncreasePositionRequestKeysStart is a free data retrieval call binding the contract method 0x9b578620.
//
// Solidity: function increasePositionRequestKeysStart() view returns(uint256)
func (_PositionRouter *PositionRouterCaller) IncreasePositionRequestKeysStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "increasePositionRequestKeysStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncreasePositionRequestKeysStart is a free data retrieval call binding the contract method 0x9b578620.
//
// Solidity: function increasePositionRequestKeysStart() view returns(uint256)
func (_PositionRouter *PositionRouterSession) IncreasePositionRequestKeysStart() (*big.Int, error) {
	return _PositionRouter.Contract.IncreasePositionRequestKeysStart(&_PositionRouter.CallOpts)
}

// IncreasePositionRequestKeysStart is a free data retrieval call binding the contract method 0x9b578620.
//
// Solidity: function increasePositionRequestKeysStart() view returns(uint256)
func (_PositionRouter *PositionRouterCallerSession) IncreasePositionRequestKeysStart() (*big.Int, error) {
	return _PositionRouter.Contract.IncreasePositionRequestKeysStart(&_PositionRouter.CallOpts)
}

// IncreasePositionRequests is a free data retrieval call binding the contract method 0xfaf990f3.
//
// Solidity: function increasePositionRequests(bytes32 ) view returns(address account, address indexToken, uint256 amountIn, uint256 minOut, uint256 sizeDelta, bool isLong, uint256 acceptablePrice, uint256 executionFee, uint256 blockNumber, uint256 blockTime, bool hasCollateralInETH, address callbackTarget)
func (_PositionRouter *PositionRouterCaller) IncreasePositionRequests(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Account            common.Address
	IndexToken         common.Address
	AmountIn           *big.Int
	MinOut             *big.Int
	SizeDelta          *big.Int
	IsLong             bool
	AcceptablePrice    *big.Int
	ExecutionFee       *big.Int
	BlockNumber        *big.Int
	BlockTime          *big.Int
	HasCollateralInETH bool
	CallbackTarget     common.Address
}, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "increasePositionRequests", arg0)

	outstruct := new(struct {
		Account            common.Address
		IndexToken         common.Address
		AmountIn           *big.Int
		MinOut             *big.Int
		SizeDelta          *big.Int
		IsLong             bool
		AcceptablePrice    *big.Int
		ExecutionFee       *big.Int
		BlockNumber        *big.Int
		BlockTime          *big.Int
		HasCollateralInETH bool
		CallbackTarget     common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Account = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IndexToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AmountIn = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MinOut = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.SizeDelta = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.IsLong = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.AcceptablePrice = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.ExecutionFee = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.BlockNumber = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.BlockTime = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.HasCollateralInETH = *abi.ConvertType(out[10], new(bool)).(*bool)
	outstruct.CallbackTarget = *abi.ConvertType(out[11], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// IncreasePositionRequests is a free data retrieval call binding the contract method 0xfaf990f3.
//
// Solidity: function increasePositionRequests(bytes32 ) view returns(address account, address indexToken, uint256 amountIn, uint256 minOut, uint256 sizeDelta, bool isLong, uint256 acceptablePrice, uint256 executionFee, uint256 blockNumber, uint256 blockTime, bool hasCollateralInETH, address callbackTarget)
func (_PositionRouter *PositionRouterSession) IncreasePositionRequests(arg0 [32]byte) (struct {
	Account            common.Address
	IndexToken         common.Address
	AmountIn           *big.Int
	MinOut             *big.Int
	SizeDelta          *big.Int
	IsLong             bool
	AcceptablePrice    *big.Int
	ExecutionFee       *big.Int
	BlockNumber        *big.Int
	BlockTime          *big.Int
	HasCollateralInETH bool
	CallbackTarget     common.Address
}, error) {
	return _PositionRouter.Contract.IncreasePositionRequests(&_PositionRouter.CallOpts, arg0)
}

// IncreasePositionRequests is a free data retrieval call binding the contract method 0xfaf990f3.
//
// Solidity: function increasePositionRequests(bytes32 ) view returns(address account, address indexToken, uint256 amountIn, uint256 minOut, uint256 sizeDelta, bool isLong, uint256 acceptablePrice, uint256 executionFee, uint256 blockNumber, uint256 blockTime, bool hasCollateralInETH, address callbackTarget)
func (_PositionRouter *PositionRouterCallerSession) IncreasePositionRequests(arg0 [32]byte) (struct {
	Account            common.Address
	IndexToken         common.Address
	AmountIn           *big.Int
	MinOut             *big.Int
	SizeDelta          *big.Int
	IsLong             bool
	AcceptablePrice    *big.Int
	ExecutionFee       *big.Int
	BlockNumber        *big.Int
	BlockTime          *big.Int
	HasCollateralInETH bool
	CallbackTarget     common.Address
}, error) {
	return _PositionRouter.Contract.IncreasePositionRequests(&_PositionRouter.CallOpts, arg0)
}

// IncreasePositionsIndex is a free data retrieval call binding the contract method 0x633451de.
//
// Solidity: function increasePositionsIndex(address ) view returns(uint256)
func (_PositionRouter *PositionRouterCaller) IncreasePositionsIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "increasePositionsIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncreasePositionsIndex is a free data retrieval call binding the contract method 0x633451de.
//
// Solidity: function increasePositionsIndex(address ) view returns(uint256)
func (_PositionRouter *PositionRouterSession) IncreasePositionsIndex(arg0 common.Address) (*big.Int, error) {
	return _PositionRouter.Contract.IncreasePositionsIndex(&_PositionRouter.CallOpts, arg0)
}

// IncreasePositionsIndex is a free data retrieval call binding the contract method 0x633451de.
//
// Solidity: function increasePositionsIndex(address ) view returns(uint256)
func (_PositionRouter *PositionRouterCallerSession) IncreasePositionsIndex(arg0 common.Address) (*big.Int, error) {
	return _PositionRouter.Contract.IncreasePositionsIndex(&_PositionRouter.CallOpts, arg0)
}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_PositionRouter *PositionRouterCaller) IsLeverageEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "isLeverageEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_PositionRouter *PositionRouterSession) IsLeverageEnabled() (bool, error) {
	return _PositionRouter.Contract.IsLeverageEnabled(&_PositionRouter.CallOpts)
}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_PositionRouter *PositionRouterCallerSession) IsLeverageEnabled() (bool, error) {
	return _PositionRouter.Contract.IsLeverageEnabled(&_PositionRouter.CallOpts)
}

// IsPositionKeeper is a free data retrieval call binding the contract method 0x36eba48a.
//
// Solidity: function isPositionKeeper(address ) view returns(bool)
func (_PositionRouter *PositionRouterCaller) IsPositionKeeper(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "isPositionKeeper", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPositionKeeper is a free data retrieval call binding the contract method 0x36eba48a.
//
// Solidity: function isPositionKeeper(address ) view returns(bool)
func (_PositionRouter *PositionRouterSession) IsPositionKeeper(arg0 common.Address) (bool, error) {
	return _PositionRouter.Contract.IsPositionKeeper(&_PositionRouter.CallOpts, arg0)
}

// IsPositionKeeper is a free data retrieval call binding the contract method 0x36eba48a.
//
// Solidity: function isPositionKeeper(address ) view returns(bool)
func (_PositionRouter *PositionRouterCallerSession) IsPositionKeeper(arg0 common.Address) (bool, error) {
	return _PositionRouter.Contract.IsPositionKeeper(&_PositionRouter.CallOpts, arg0)
}

// MaxGlobalLongSizes is a free data retrieval call binding the contract method 0x1045c74e.
//
// Solidity: function maxGlobalLongSizes(address ) view returns(uint256)
func (_PositionRouter *PositionRouterCaller) MaxGlobalLongSizes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "maxGlobalLongSizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGlobalLongSizes is a free data retrieval call binding the contract method 0x1045c74e.
//
// Solidity: function maxGlobalLongSizes(address ) view returns(uint256)
func (_PositionRouter *PositionRouterSession) MaxGlobalLongSizes(arg0 common.Address) (*big.Int, error) {
	return _PositionRouter.Contract.MaxGlobalLongSizes(&_PositionRouter.CallOpts, arg0)
}

// MaxGlobalLongSizes is a free data retrieval call binding the contract method 0x1045c74e.
//
// Solidity: function maxGlobalLongSizes(address ) view returns(uint256)
func (_PositionRouter *PositionRouterCallerSession) MaxGlobalLongSizes(arg0 common.Address) (*big.Int, error) {
	return _PositionRouter.Contract.MaxGlobalLongSizes(&_PositionRouter.CallOpts, arg0)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_PositionRouter *PositionRouterCaller) MaxGlobalShortSizes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "maxGlobalShortSizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_PositionRouter *PositionRouterSession) MaxGlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _PositionRouter.Contract.MaxGlobalShortSizes(&_PositionRouter.CallOpts, arg0)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_PositionRouter *PositionRouterCallerSession) MaxGlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _PositionRouter.Contract.MaxGlobalShortSizes(&_PositionRouter.CallOpts, arg0)
}

// MaxTimeDelay is a free data retrieval call binding the contract method 0xcb0269c9.
//
// Solidity: function maxTimeDelay() view returns(uint256)
func (_PositionRouter *PositionRouterCaller) MaxTimeDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "maxTimeDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTimeDelay is a free data retrieval call binding the contract method 0xcb0269c9.
//
// Solidity: function maxTimeDelay() view returns(uint256)
func (_PositionRouter *PositionRouterSession) MaxTimeDelay() (*big.Int, error) {
	return _PositionRouter.Contract.MaxTimeDelay(&_PositionRouter.CallOpts)
}

// MaxTimeDelay is a free data retrieval call binding the contract method 0xcb0269c9.
//
// Solidity: function maxTimeDelay() view returns(uint256)
func (_PositionRouter *PositionRouterCallerSession) MaxTimeDelay() (*big.Int, error) {
	return _PositionRouter.Contract.MaxTimeDelay(&_PositionRouter.CallOpts)
}

// MinBlockDelayKeeper is a free data retrieval call binding the contract method 0x5841fcaa.
//
// Solidity: function minBlockDelayKeeper() view returns(uint256)
func (_PositionRouter *PositionRouterCaller) MinBlockDelayKeeper(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "minBlockDelayKeeper")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBlockDelayKeeper is a free data retrieval call binding the contract method 0x5841fcaa.
//
// Solidity: function minBlockDelayKeeper() view returns(uint256)
func (_PositionRouter *PositionRouterSession) MinBlockDelayKeeper() (*big.Int, error) {
	return _PositionRouter.Contract.MinBlockDelayKeeper(&_PositionRouter.CallOpts)
}

// MinBlockDelayKeeper is a free data retrieval call binding the contract method 0x5841fcaa.
//
// Solidity: function minBlockDelayKeeper() view returns(uint256)
func (_PositionRouter *PositionRouterCallerSession) MinBlockDelayKeeper() (*big.Int, error) {
	return _PositionRouter.Contract.MinBlockDelayKeeper(&_PositionRouter.CallOpts)
}

// MinExecutionFee is a free data retrieval call binding the contract method 0x63ae2103.
//
// Solidity: function minExecutionFee() view returns(uint256)
func (_PositionRouter *PositionRouterCaller) MinExecutionFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "minExecutionFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinExecutionFee is a free data retrieval call binding the contract method 0x63ae2103.
//
// Solidity: function minExecutionFee() view returns(uint256)
func (_PositionRouter *PositionRouterSession) MinExecutionFee() (*big.Int, error) {
	return _PositionRouter.Contract.MinExecutionFee(&_PositionRouter.CallOpts)
}

// MinExecutionFee is a free data retrieval call binding the contract method 0x63ae2103.
//
// Solidity: function minExecutionFee() view returns(uint256)
func (_PositionRouter *PositionRouterCallerSession) MinExecutionFee() (*big.Int, error) {
	return _PositionRouter.Contract.MinExecutionFee(&_PositionRouter.CallOpts)
}

// MinTimeDelayPublic is a free data retrieval call binding the contract method 0x3a2a80c7.
//
// Solidity: function minTimeDelayPublic() view returns(uint256)
func (_PositionRouter *PositionRouterCaller) MinTimeDelayPublic(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "minTimeDelayPublic")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinTimeDelayPublic is a free data retrieval call binding the contract method 0x3a2a80c7.
//
// Solidity: function minTimeDelayPublic() view returns(uint256)
func (_PositionRouter *PositionRouterSession) MinTimeDelayPublic() (*big.Int, error) {
	return _PositionRouter.Contract.MinTimeDelayPublic(&_PositionRouter.CallOpts)
}

// MinTimeDelayPublic is a free data retrieval call binding the contract method 0x3a2a80c7.
//
// Solidity: function minTimeDelayPublic() view returns(uint256)
func (_PositionRouter *PositionRouterCallerSession) MinTimeDelayPublic() (*big.Int, error) {
	return _PositionRouter.Contract.MinTimeDelayPublic(&_PositionRouter.CallOpts)
}

// ReferralStorage is a free data retrieval call binding the contract method 0x006cc35e.
//
// Solidity: function referralStorage() view returns(address)
func (_PositionRouter *PositionRouterCaller) ReferralStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "referralStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReferralStorage is a free data retrieval call binding the contract method 0x006cc35e.
//
// Solidity: function referralStorage() view returns(address)
func (_PositionRouter *PositionRouterSession) ReferralStorage() (common.Address, error) {
	return _PositionRouter.Contract.ReferralStorage(&_PositionRouter.CallOpts)
}

// ReferralStorage is a free data retrieval call binding the contract method 0x006cc35e.
//
// Solidity: function referralStorage() view returns(address)
func (_PositionRouter *PositionRouterCallerSession) ReferralStorage() (common.Address, error) {
	return _PositionRouter.Contract.ReferralStorage(&_PositionRouter.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_PositionRouter *PositionRouterCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_PositionRouter *PositionRouterSession) Router() (common.Address, error) {
	return _PositionRouter.Contract.Router(&_PositionRouter.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_PositionRouter *PositionRouterCallerSession) Router() (common.Address, error) {
	return _PositionRouter.Contract.Router(&_PositionRouter.CallOpts)
}

// ShortsTracker is a free data retrieval call binding the contract method 0x657bc5d0.
//
// Solidity: function shortsTracker() view returns(address)
func (_PositionRouter *PositionRouterCaller) ShortsTracker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "shortsTracker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ShortsTracker is a free data retrieval call binding the contract method 0x657bc5d0.
//
// Solidity: function shortsTracker() view returns(address)
func (_PositionRouter *PositionRouterSession) ShortsTracker() (common.Address, error) {
	return _PositionRouter.Contract.ShortsTracker(&_PositionRouter.CallOpts)
}

// ShortsTracker is a free data retrieval call binding the contract method 0x657bc5d0.
//
// Solidity: function shortsTracker() view returns(address)
func (_PositionRouter *PositionRouterCallerSession) ShortsTracker() (common.Address, error) {
	return _PositionRouter.Contract.ShortsTracker(&_PositionRouter.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_PositionRouter *PositionRouterCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_PositionRouter *PositionRouterSession) Vault() (common.Address, error) {
	return _PositionRouter.Contract.Vault(&_PositionRouter.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_PositionRouter *PositionRouterCallerSession) Vault() (common.Address, error) {
	return _PositionRouter.Contract.Vault(&_PositionRouter.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PositionRouter *PositionRouterCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PositionRouter.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PositionRouter *PositionRouterSession) Weth() (common.Address, error) {
	return _PositionRouter.Contract.Weth(&_PositionRouter.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_PositionRouter *PositionRouterCallerSession) Weth() (common.Address, error) {
	return _PositionRouter.Contract.Weth(&_PositionRouter.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount) returns()
func (_PositionRouter *PositionRouterTransactor) Approve(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "approve", _token, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount) returns()
func (_PositionRouter *PositionRouterSession) Approve(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.Approve(&_PositionRouter.TransactOpts, _token, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount) returns()
func (_PositionRouter *PositionRouterTransactorSession) Approve(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.Approve(&_PositionRouter.TransactOpts, _token, _spender, _amount)
}

// CancelDecreasePosition is a paid mutator transaction binding the contract method 0x60a362e2.
//
// Solidity: function cancelDecreasePosition(bytes32 _key, address _executionFeeReceiver) returns(bool)
func (_PositionRouter *PositionRouterTransactor) CancelDecreasePosition(opts *bind.TransactOpts, _key [32]byte, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "cancelDecreasePosition", _key, _executionFeeReceiver)
}

// CancelDecreasePosition is a paid mutator transaction binding the contract method 0x60a362e2.
//
// Solidity: function cancelDecreasePosition(bytes32 _key, address _executionFeeReceiver) returns(bool)
func (_PositionRouter *PositionRouterSession) CancelDecreasePosition(_key [32]byte, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.CancelDecreasePosition(&_PositionRouter.TransactOpts, _key, _executionFeeReceiver)
}

// CancelDecreasePosition is a paid mutator transaction binding the contract method 0x60a362e2.
//
// Solidity: function cancelDecreasePosition(bytes32 _key, address _executionFeeReceiver) returns(bool)
func (_PositionRouter *PositionRouterTransactorSession) CancelDecreasePosition(_key [32]byte, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.CancelDecreasePosition(&_PositionRouter.TransactOpts, _key, _executionFeeReceiver)
}

// CancelIncreasePosition is a paid mutator transaction binding the contract method 0x225fc9fd.
//
// Solidity: function cancelIncreasePosition(bytes32 _key, address _executionFeeReceiver) returns(bool)
func (_PositionRouter *PositionRouterTransactor) CancelIncreasePosition(opts *bind.TransactOpts, _key [32]byte, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "cancelIncreasePosition", _key, _executionFeeReceiver)
}

// CancelIncreasePosition is a paid mutator transaction binding the contract method 0x225fc9fd.
//
// Solidity: function cancelIncreasePosition(bytes32 _key, address _executionFeeReceiver) returns(bool)
func (_PositionRouter *PositionRouterSession) CancelIncreasePosition(_key [32]byte, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.CancelIncreasePosition(&_PositionRouter.TransactOpts, _key, _executionFeeReceiver)
}

// CancelIncreasePosition is a paid mutator transaction binding the contract method 0x225fc9fd.
//
// Solidity: function cancelIncreasePosition(bytes32 _key, address _executionFeeReceiver) returns(bool)
func (_PositionRouter *PositionRouterTransactorSession) CancelIncreasePosition(_key [32]byte, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.CancelIncreasePosition(&_PositionRouter.TransactOpts, _key, _executionFeeReceiver)
}

// CreateDecreasePosition is a paid mutator transaction binding the contract method 0x7be7d141.
//
// Solidity: function createDecreasePosition(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _acceptablePrice, uint256 _minOut, uint256 _executionFee, bool _withdrawETH, address _callbackTarget) payable returns(bytes32)
func (_PositionRouter *PositionRouterTransactor) CreateDecreasePosition(opts *bind.TransactOpts, _path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _acceptablePrice *big.Int, _minOut *big.Int, _executionFee *big.Int, _withdrawETH bool, _callbackTarget common.Address) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "createDecreasePosition", _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _acceptablePrice, _minOut, _executionFee, _withdrawETH, _callbackTarget)
}

// CreateDecreasePosition is a paid mutator transaction binding the contract method 0x7be7d141.
//
// Solidity: function createDecreasePosition(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _acceptablePrice, uint256 _minOut, uint256 _executionFee, bool _withdrawETH, address _callbackTarget) payable returns(bytes32)
func (_PositionRouter *PositionRouterSession) CreateDecreasePosition(_path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _acceptablePrice *big.Int, _minOut *big.Int, _executionFee *big.Int, _withdrawETH bool, _callbackTarget common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.CreateDecreasePosition(&_PositionRouter.TransactOpts, _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _acceptablePrice, _minOut, _executionFee, _withdrawETH, _callbackTarget)
}

// CreateDecreasePosition is a paid mutator transaction binding the contract method 0x7be7d141.
//
// Solidity: function createDecreasePosition(address[] _path, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver, uint256 _acceptablePrice, uint256 _minOut, uint256 _executionFee, bool _withdrawETH, address _callbackTarget) payable returns(bytes32)
func (_PositionRouter *PositionRouterTransactorSession) CreateDecreasePosition(_path []common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address, _acceptablePrice *big.Int, _minOut *big.Int, _executionFee *big.Int, _withdrawETH bool, _callbackTarget common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.CreateDecreasePosition(&_PositionRouter.TransactOpts, _path, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver, _acceptablePrice, _minOut, _executionFee, _withdrawETH, _callbackTarget)
}

// CreateIncreasePosition is a paid mutator transaction binding the contract method 0xf2ae372f.
//
// Solidity: function createIncreasePosition(address[] _path, address _indexToken, uint256 _amountIn, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _acceptablePrice, uint256 _executionFee, bytes32 _referralCode, address _callbackTarget) payable returns(bytes32)
func (_PositionRouter *PositionRouterTransactor) CreateIncreasePosition(opts *bind.TransactOpts, _path []common.Address, _indexToken common.Address, _amountIn *big.Int, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _acceptablePrice *big.Int, _executionFee *big.Int, _referralCode [32]byte, _callbackTarget common.Address) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "createIncreasePosition", _path, _indexToken, _amountIn, _minOut, _sizeDelta, _isLong, _acceptablePrice, _executionFee, _referralCode, _callbackTarget)
}

// CreateIncreasePosition is a paid mutator transaction binding the contract method 0xf2ae372f.
//
// Solidity: function createIncreasePosition(address[] _path, address _indexToken, uint256 _amountIn, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _acceptablePrice, uint256 _executionFee, bytes32 _referralCode, address _callbackTarget) payable returns(bytes32)
func (_PositionRouter *PositionRouterSession) CreateIncreasePosition(_path []common.Address, _indexToken common.Address, _amountIn *big.Int, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _acceptablePrice *big.Int, _executionFee *big.Int, _referralCode [32]byte, _callbackTarget common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.CreateIncreasePosition(&_PositionRouter.TransactOpts, _path, _indexToken, _amountIn, _minOut, _sizeDelta, _isLong, _acceptablePrice, _executionFee, _referralCode, _callbackTarget)
}

// CreateIncreasePosition is a paid mutator transaction binding the contract method 0xf2ae372f.
//
// Solidity: function createIncreasePosition(address[] _path, address _indexToken, uint256 _amountIn, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _acceptablePrice, uint256 _executionFee, bytes32 _referralCode, address _callbackTarget) payable returns(bytes32)
func (_PositionRouter *PositionRouterTransactorSession) CreateIncreasePosition(_path []common.Address, _indexToken common.Address, _amountIn *big.Int, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _acceptablePrice *big.Int, _executionFee *big.Int, _referralCode [32]byte, _callbackTarget common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.CreateIncreasePosition(&_PositionRouter.TransactOpts, _path, _indexToken, _amountIn, _minOut, _sizeDelta, _isLong, _acceptablePrice, _executionFee, _referralCode, _callbackTarget)
}

// CreateIncreasePositionETH is a paid mutator transaction binding the contract method 0x5b88e8c6.
//
// Solidity: function createIncreasePositionETH(address[] _path, address _indexToken, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _acceptablePrice, uint256 _executionFee, bytes32 _referralCode, address _callbackTarget) payable returns(bytes32)
func (_PositionRouter *PositionRouterTransactor) CreateIncreasePositionETH(opts *bind.TransactOpts, _path []common.Address, _indexToken common.Address, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _acceptablePrice *big.Int, _executionFee *big.Int, _referralCode [32]byte, _callbackTarget common.Address) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "createIncreasePositionETH", _path, _indexToken, _minOut, _sizeDelta, _isLong, _acceptablePrice, _executionFee, _referralCode, _callbackTarget)
}

// CreateIncreasePositionETH is a paid mutator transaction binding the contract method 0x5b88e8c6.
//
// Solidity: function createIncreasePositionETH(address[] _path, address _indexToken, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _acceptablePrice, uint256 _executionFee, bytes32 _referralCode, address _callbackTarget) payable returns(bytes32)
func (_PositionRouter *PositionRouterSession) CreateIncreasePositionETH(_path []common.Address, _indexToken common.Address, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _acceptablePrice *big.Int, _executionFee *big.Int, _referralCode [32]byte, _callbackTarget common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.CreateIncreasePositionETH(&_PositionRouter.TransactOpts, _path, _indexToken, _minOut, _sizeDelta, _isLong, _acceptablePrice, _executionFee, _referralCode, _callbackTarget)
}

// CreateIncreasePositionETH is a paid mutator transaction binding the contract method 0x5b88e8c6.
//
// Solidity: function createIncreasePositionETH(address[] _path, address _indexToken, uint256 _minOut, uint256 _sizeDelta, bool _isLong, uint256 _acceptablePrice, uint256 _executionFee, bytes32 _referralCode, address _callbackTarget) payable returns(bytes32)
func (_PositionRouter *PositionRouterTransactorSession) CreateIncreasePositionETH(_path []common.Address, _indexToken common.Address, _minOut *big.Int, _sizeDelta *big.Int, _isLong bool, _acceptablePrice *big.Int, _executionFee *big.Int, _referralCode [32]byte, _callbackTarget common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.CreateIncreasePositionETH(&_PositionRouter.TransactOpts, _path, _indexToken, _minOut, _sizeDelta, _isLong, _acceptablePrice, _executionFee, _referralCode, _callbackTarget)
}

// ExecuteDecreasePosition is a paid mutator transaction binding the contract method 0x0d4d003d.
//
// Solidity: function executeDecreasePosition(bytes32 _key, address _executionFeeReceiver) returns(bool)
func (_PositionRouter *PositionRouterTransactor) ExecuteDecreasePosition(opts *bind.TransactOpts, _key [32]byte, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "executeDecreasePosition", _key, _executionFeeReceiver)
}

// ExecuteDecreasePosition is a paid mutator transaction binding the contract method 0x0d4d003d.
//
// Solidity: function executeDecreasePosition(bytes32 _key, address _executionFeeReceiver) returns(bool)
func (_PositionRouter *PositionRouterSession) ExecuteDecreasePosition(_key [32]byte, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.ExecuteDecreasePosition(&_PositionRouter.TransactOpts, _key, _executionFeeReceiver)
}

// ExecuteDecreasePosition is a paid mutator transaction binding the contract method 0x0d4d003d.
//
// Solidity: function executeDecreasePosition(bytes32 _key, address _executionFeeReceiver) returns(bool)
func (_PositionRouter *PositionRouterTransactorSession) ExecuteDecreasePosition(_key [32]byte, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.ExecuteDecreasePosition(&_PositionRouter.TransactOpts, _key, _executionFeeReceiver)
}

// ExecuteDecreasePositions is a paid mutator transaction binding the contract method 0xf3883d8b.
//
// Solidity: function executeDecreasePositions(uint256 _endIndex, address _executionFeeReceiver) returns()
func (_PositionRouter *PositionRouterTransactor) ExecuteDecreasePositions(opts *bind.TransactOpts, _endIndex *big.Int, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "executeDecreasePositions", _endIndex, _executionFeeReceiver)
}

// ExecuteDecreasePositions is a paid mutator transaction binding the contract method 0xf3883d8b.
//
// Solidity: function executeDecreasePositions(uint256 _endIndex, address _executionFeeReceiver) returns()
func (_PositionRouter *PositionRouterSession) ExecuteDecreasePositions(_endIndex *big.Int, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.ExecuteDecreasePositions(&_PositionRouter.TransactOpts, _endIndex, _executionFeeReceiver)
}

// ExecuteDecreasePositions is a paid mutator transaction binding the contract method 0xf3883d8b.
//
// Solidity: function executeDecreasePositions(uint256 _endIndex, address _executionFeeReceiver) returns()
func (_PositionRouter *PositionRouterTransactorSession) ExecuteDecreasePositions(_endIndex *big.Int, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.ExecuteDecreasePositions(&_PositionRouter.TransactOpts, _endIndex, _executionFeeReceiver)
}

// ExecuteIncreasePosition is a paid mutator transaction binding the contract method 0x27b42c0f.
//
// Solidity: function executeIncreasePosition(bytes32 _key, address _executionFeeReceiver) returns(bool)
func (_PositionRouter *PositionRouterTransactor) ExecuteIncreasePosition(opts *bind.TransactOpts, _key [32]byte, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "executeIncreasePosition", _key, _executionFeeReceiver)
}

// ExecuteIncreasePosition is a paid mutator transaction binding the contract method 0x27b42c0f.
//
// Solidity: function executeIncreasePosition(bytes32 _key, address _executionFeeReceiver) returns(bool)
func (_PositionRouter *PositionRouterSession) ExecuteIncreasePosition(_key [32]byte, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.ExecuteIncreasePosition(&_PositionRouter.TransactOpts, _key, _executionFeeReceiver)
}

// ExecuteIncreasePosition is a paid mutator transaction binding the contract method 0x27b42c0f.
//
// Solidity: function executeIncreasePosition(bytes32 _key, address _executionFeeReceiver) returns(bool)
func (_PositionRouter *PositionRouterTransactorSession) ExecuteIncreasePosition(_key [32]byte, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.ExecuteIncreasePosition(&_PositionRouter.TransactOpts, _key, _executionFeeReceiver)
}

// ExecuteIncreasePositions is a paid mutator transaction binding the contract method 0x9a208100.
//
// Solidity: function executeIncreasePositions(uint256 _endIndex, address _executionFeeReceiver) returns()
func (_PositionRouter *PositionRouterTransactor) ExecuteIncreasePositions(opts *bind.TransactOpts, _endIndex *big.Int, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "executeIncreasePositions", _endIndex, _executionFeeReceiver)
}

// ExecuteIncreasePositions is a paid mutator transaction binding the contract method 0x9a208100.
//
// Solidity: function executeIncreasePositions(uint256 _endIndex, address _executionFeeReceiver) returns()
func (_PositionRouter *PositionRouterSession) ExecuteIncreasePositions(_endIndex *big.Int, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.ExecuteIncreasePositions(&_PositionRouter.TransactOpts, _endIndex, _executionFeeReceiver)
}

// ExecuteIncreasePositions is a paid mutator transaction binding the contract method 0x9a208100.
//
// Solidity: function executeIncreasePositions(uint256 _endIndex, address _executionFeeReceiver) returns()
func (_PositionRouter *PositionRouterTransactorSession) ExecuteIncreasePositions(_endIndex *big.Int, _executionFeeReceiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.ExecuteIncreasePositions(&_PositionRouter.TransactOpts, _endIndex, _executionFeeReceiver)
}

// SendValue is a paid mutator transaction binding the contract method 0x24a084df.
//
// Solidity: function sendValue(address _receiver, uint256 _amount) returns()
func (_PositionRouter *PositionRouterTransactor) SendValue(opts *bind.TransactOpts, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "sendValue", _receiver, _amount)
}

// SendValue is a paid mutator transaction binding the contract method 0x24a084df.
//
// Solidity: function sendValue(address _receiver, uint256 _amount) returns()
func (_PositionRouter *PositionRouterSession) SendValue(_receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SendValue(&_PositionRouter.TransactOpts, _receiver, _amount)
}

// SendValue is a paid mutator transaction binding the contract method 0x24a084df.
//
// Solidity: function sendValue(address _receiver, uint256 _amount) returns()
func (_PositionRouter *PositionRouterTransactorSession) SendValue(_receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SendValue(&_PositionRouter.TransactOpts, _receiver, _amount)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_PositionRouter *PositionRouterTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_PositionRouter *PositionRouterSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetAdmin(&_PositionRouter.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_PositionRouter *PositionRouterTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetAdmin(&_PositionRouter.TransactOpts, _admin)
}

// SetCallbackGasLimit is a paid mutator transaction binding the contract method 0x8a54942f.
//
// Solidity: function setCallbackGasLimit(uint256 _callbackGasLimit) returns()
func (_PositionRouter *PositionRouterTransactor) SetCallbackGasLimit(opts *bind.TransactOpts, _callbackGasLimit *big.Int) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "setCallbackGasLimit", _callbackGasLimit)
}

// SetCallbackGasLimit is a paid mutator transaction binding the contract method 0x8a54942f.
//
// Solidity: function setCallbackGasLimit(uint256 _callbackGasLimit) returns()
func (_PositionRouter *PositionRouterSession) SetCallbackGasLimit(_callbackGasLimit *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetCallbackGasLimit(&_PositionRouter.TransactOpts, _callbackGasLimit)
}

// SetCallbackGasLimit is a paid mutator transaction binding the contract method 0x8a54942f.
//
// Solidity: function setCallbackGasLimit(uint256 _callbackGasLimit) returns()
func (_PositionRouter *PositionRouterTransactorSession) SetCallbackGasLimit(_callbackGasLimit *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetCallbackGasLimit(&_PositionRouter.TransactOpts, _callbackGasLimit)
}

// SetCustomCallbackGasLimit is a paid mutator transaction binding the contract method 0x804fc702.
//
// Solidity: function setCustomCallbackGasLimit(address _callbackTarget, uint256 _callbackGasLimit) returns()
func (_PositionRouter *PositionRouterTransactor) SetCustomCallbackGasLimit(opts *bind.TransactOpts, _callbackTarget common.Address, _callbackGasLimit *big.Int) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "setCustomCallbackGasLimit", _callbackTarget, _callbackGasLimit)
}

// SetCustomCallbackGasLimit is a paid mutator transaction binding the contract method 0x804fc702.
//
// Solidity: function setCustomCallbackGasLimit(address _callbackTarget, uint256 _callbackGasLimit) returns()
func (_PositionRouter *PositionRouterSession) SetCustomCallbackGasLimit(_callbackTarget common.Address, _callbackGasLimit *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetCustomCallbackGasLimit(&_PositionRouter.TransactOpts, _callbackTarget, _callbackGasLimit)
}

// SetCustomCallbackGasLimit is a paid mutator transaction binding the contract method 0x804fc702.
//
// Solidity: function setCustomCallbackGasLimit(address _callbackTarget, uint256 _callbackGasLimit) returns()
func (_PositionRouter *PositionRouterTransactorSession) SetCustomCallbackGasLimit(_callbackTarget common.Address, _callbackGasLimit *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetCustomCallbackGasLimit(&_PositionRouter.TransactOpts, _callbackTarget, _callbackGasLimit)
}

// SetDelayValues is a paid mutator transaction binding the contract method 0x4067b132.
//
// Solidity: function setDelayValues(uint256 _minBlockDelayKeeper, uint256 _minTimeDelayPublic, uint256 _maxTimeDelay) returns()
func (_PositionRouter *PositionRouterTransactor) SetDelayValues(opts *bind.TransactOpts, _minBlockDelayKeeper *big.Int, _minTimeDelayPublic *big.Int, _maxTimeDelay *big.Int) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "setDelayValues", _minBlockDelayKeeper, _minTimeDelayPublic, _maxTimeDelay)
}

// SetDelayValues is a paid mutator transaction binding the contract method 0x4067b132.
//
// Solidity: function setDelayValues(uint256 _minBlockDelayKeeper, uint256 _minTimeDelayPublic, uint256 _maxTimeDelay) returns()
func (_PositionRouter *PositionRouterSession) SetDelayValues(_minBlockDelayKeeper *big.Int, _minTimeDelayPublic *big.Int, _maxTimeDelay *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetDelayValues(&_PositionRouter.TransactOpts, _minBlockDelayKeeper, _minTimeDelayPublic, _maxTimeDelay)
}

// SetDelayValues is a paid mutator transaction binding the contract method 0x4067b132.
//
// Solidity: function setDelayValues(uint256 _minBlockDelayKeeper, uint256 _minTimeDelayPublic, uint256 _maxTimeDelay) returns()
func (_PositionRouter *PositionRouterTransactorSession) SetDelayValues(_minBlockDelayKeeper *big.Int, _minTimeDelayPublic *big.Int, _maxTimeDelay *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetDelayValues(&_PositionRouter.TransactOpts, _minBlockDelayKeeper, _minTimeDelayPublic, _maxTimeDelay)
}

// SetDepositFee is a paid mutator transaction binding the contract method 0x490ae210.
//
// Solidity: function setDepositFee(uint256 _depositFee) returns()
func (_PositionRouter *PositionRouterTransactor) SetDepositFee(opts *bind.TransactOpts, _depositFee *big.Int) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "setDepositFee", _depositFee)
}

// SetDepositFee is a paid mutator transaction binding the contract method 0x490ae210.
//
// Solidity: function setDepositFee(uint256 _depositFee) returns()
func (_PositionRouter *PositionRouterSession) SetDepositFee(_depositFee *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetDepositFee(&_PositionRouter.TransactOpts, _depositFee)
}

// SetDepositFee is a paid mutator transaction binding the contract method 0x490ae210.
//
// Solidity: function setDepositFee(uint256 _depositFee) returns()
func (_PositionRouter *PositionRouterTransactorSession) SetDepositFee(_depositFee *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetDepositFee(&_PositionRouter.TransactOpts, _depositFee)
}

// SetEthTransferGasLimit is a paid mutator transaction binding the contract method 0x3a9b52ad.
//
// Solidity: function setEthTransferGasLimit(uint256 _ethTransferGasLimit) returns()
func (_PositionRouter *PositionRouterTransactor) SetEthTransferGasLimit(opts *bind.TransactOpts, _ethTransferGasLimit *big.Int) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "setEthTransferGasLimit", _ethTransferGasLimit)
}

// SetEthTransferGasLimit is a paid mutator transaction binding the contract method 0x3a9b52ad.
//
// Solidity: function setEthTransferGasLimit(uint256 _ethTransferGasLimit) returns()
func (_PositionRouter *PositionRouterSession) SetEthTransferGasLimit(_ethTransferGasLimit *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetEthTransferGasLimit(&_PositionRouter.TransactOpts, _ethTransferGasLimit)
}

// SetEthTransferGasLimit is a paid mutator transaction binding the contract method 0x3a9b52ad.
//
// Solidity: function setEthTransferGasLimit(uint256 _ethTransferGasLimit) returns()
func (_PositionRouter *PositionRouterTransactorSession) SetEthTransferGasLimit(_ethTransferGasLimit *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetEthTransferGasLimit(&_PositionRouter.TransactOpts, _ethTransferGasLimit)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_PositionRouter *PositionRouterTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_PositionRouter *PositionRouterSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetGov(&_PositionRouter.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_PositionRouter *PositionRouterTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetGov(&_PositionRouter.TransactOpts, _gov)
}

// SetIncreasePositionBufferBps is a paid mutator transaction binding the contract method 0x233bfe3b.
//
// Solidity: function setIncreasePositionBufferBps(uint256 _increasePositionBufferBps) returns()
func (_PositionRouter *PositionRouterTransactor) SetIncreasePositionBufferBps(opts *bind.TransactOpts, _increasePositionBufferBps *big.Int) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "setIncreasePositionBufferBps", _increasePositionBufferBps)
}

// SetIncreasePositionBufferBps is a paid mutator transaction binding the contract method 0x233bfe3b.
//
// Solidity: function setIncreasePositionBufferBps(uint256 _increasePositionBufferBps) returns()
func (_PositionRouter *PositionRouterSession) SetIncreasePositionBufferBps(_increasePositionBufferBps *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetIncreasePositionBufferBps(&_PositionRouter.TransactOpts, _increasePositionBufferBps)
}

// SetIncreasePositionBufferBps is a paid mutator transaction binding the contract method 0x233bfe3b.
//
// Solidity: function setIncreasePositionBufferBps(uint256 _increasePositionBufferBps) returns()
func (_PositionRouter *PositionRouterTransactorSession) SetIncreasePositionBufferBps(_increasePositionBufferBps *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetIncreasePositionBufferBps(&_PositionRouter.TransactOpts, _increasePositionBufferBps)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_PositionRouter *PositionRouterTransactor) SetIsLeverageEnabled(opts *bind.TransactOpts, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "setIsLeverageEnabled", _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_PositionRouter *PositionRouterSession) SetIsLeverageEnabled(_isLeverageEnabled bool) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetIsLeverageEnabled(&_PositionRouter.TransactOpts, _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_PositionRouter *PositionRouterTransactorSession) SetIsLeverageEnabled(_isLeverageEnabled bool) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetIsLeverageEnabled(&_PositionRouter.TransactOpts, _isLeverageEnabled)
}

// SetMaxGlobalSizes is a paid mutator transaction binding the contract method 0xef12c67e.
//
// Solidity: function setMaxGlobalSizes(address[] _tokens, uint256[] _longSizes, uint256[] _shortSizes) returns()
func (_PositionRouter *PositionRouterTransactor) SetMaxGlobalSizes(opts *bind.TransactOpts, _tokens []common.Address, _longSizes []*big.Int, _shortSizes []*big.Int) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "setMaxGlobalSizes", _tokens, _longSizes, _shortSizes)
}

// SetMaxGlobalSizes is a paid mutator transaction binding the contract method 0xef12c67e.
//
// Solidity: function setMaxGlobalSizes(address[] _tokens, uint256[] _longSizes, uint256[] _shortSizes) returns()
func (_PositionRouter *PositionRouterSession) SetMaxGlobalSizes(_tokens []common.Address, _longSizes []*big.Int, _shortSizes []*big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetMaxGlobalSizes(&_PositionRouter.TransactOpts, _tokens, _longSizes, _shortSizes)
}

// SetMaxGlobalSizes is a paid mutator transaction binding the contract method 0xef12c67e.
//
// Solidity: function setMaxGlobalSizes(address[] _tokens, uint256[] _longSizes, uint256[] _shortSizes) returns()
func (_PositionRouter *PositionRouterTransactorSession) SetMaxGlobalSizes(_tokens []common.Address, _longSizes []*big.Int, _shortSizes []*big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetMaxGlobalSizes(&_PositionRouter.TransactOpts, _tokens, _longSizes, _shortSizes)
}

// SetMinExecutionFee is a paid mutator transaction binding the contract method 0xfc2cee62.
//
// Solidity: function setMinExecutionFee(uint256 _minExecutionFee) returns()
func (_PositionRouter *PositionRouterTransactor) SetMinExecutionFee(opts *bind.TransactOpts, _minExecutionFee *big.Int) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "setMinExecutionFee", _minExecutionFee)
}

// SetMinExecutionFee is a paid mutator transaction binding the contract method 0xfc2cee62.
//
// Solidity: function setMinExecutionFee(uint256 _minExecutionFee) returns()
func (_PositionRouter *PositionRouterSession) SetMinExecutionFee(_minExecutionFee *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetMinExecutionFee(&_PositionRouter.TransactOpts, _minExecutionFee)
}

// SetMinExecutionFee is a paid mutator transaction binding the contract method 0xfc2cee62.
//
// Solidity: function setMinExecutionFee(uint256 _minExecutionFee) returns()
func (_PositionRouter *PositionRouterTransactorSession) SetMinExecutionFee(_minExecutionFee *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetMinExecutionFee(&_PositionRouter.TransactOpts, _minExecutionFee)
}

// SetPositionKeeper is a paid mutator transaction binding the contract method 0x3422ead1.
//
// Solidity: function setPositionKeeper(address _account, bool _isActive) returns()
func (_PositionRouter *PositionRouterTransactor) SetPositionKeeper(opts *bind.TransactOpts, _account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "setPositionKeeper", _account, _isActive)
}

// SetPositionKeeper is a paid mutator transaction binding the contract method 0x3422ead1.
//
// Solidity: function setPositionKeeper(address _account, bool _isActive) returns()
func (_PositionRouter *PositionRouterSession) SetPositionKeeper(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetPositionKeeper(&_PositionRouter.TransactOpts, _account, _isActive)
}

// SetPositionKeeper is a paid mutator transaction binding the contract method 0x3422ead1.
//
// Solidity: function setPositionKeeper(address _account, bool _isActive) returns()
func (_PositionRouter *PositionRouterTransactorSession) SetPositionKeeper(_account common.Address, _isActive bool) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetPositionKeeper(&_PositionRouter.TransactOpts, _account, _isActive)
}

// SetReferralStorage is a paid mutator transaction binding the contract method 0xae4d7f9a.
//
// Solidity: function setReferralStorage(address _referralStorage) returns()
func (_PositionRouter *PositionRouterTransactor) SetReferralStorage(opts *bind.TransactOpts, _referralStorage common.Address) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "setReferralStorage", _referralStorage)
}

// SetReferralStorage is a paid mutator transaction binding the contract method 0xae4d7f9a.
//
// Solidity: function setReferralStorage(address _referralStorage) returns()
func (_PositionRouter *PositionRouterSession) SetReferralStorage(_referralStorage common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetReferralStorage(&_PositionRouter.TransactOpts, _referralStorage)
}

// SetReferralStorage is a paid mutator transaction binding the contract method 0xae4d7f9a.
//
// Solidity: function setReferralStorage(address _referralStorage) returns()
func (_PositionRouter *PositionRouterTransactorSession) SetReferralStorage(_referralStorage common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetReferralStorage(&_PositionRouter.TransactOpts, _referralStorage)
}

// SetRequestKeysStartValues is a paid mutator transaction binding the contract method 0x308aa81f.
//
// Solidity: function setRequestKeysStartValues(uint256 _increasePositionRequestKeysStart, uint256 _decreasePositionRequestKeysStart) returns()
func (_PositionRouter *PositionRouterTransactor) SetRequestKeysStartValues(opts *bind.TransactOpts, _increasePositionRequestKeysStart *big.Int, _decreasePositionRequestKeysStart *big.Int) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "setRequestKeysStartValues", _increasePositionRequestKeysStart, _decreasePositionRequestKeysStart)
}

// SetRequestKeysStartValues is a paid mutator transaction binding the contract method 0x308aa81f.
//
// Solidity: function setRequestKeysStartValues(uint256 _increasePositionRequestKeysStart, uint256 _decreasePositionRequestKeysStart) returns()
func (_PositionRouter *PositionRouterSession) SetRequestKeysStartValues(_increasePositionRequestKeysStart *big.Int, _decreasePositionRequestKeysStart *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetRequestKeysStartValues(&_PositionRouter.TransactOpts, _increasePositionRequestKeysStart, _decreasePositionRequestKeysStart)
}

// SetRequestKeysStartValues is a paid mutator transaction binding the contract method 0x308aa81f.
//
// Solidity: function setRequestKeysStartValues(uint256 _increasePositionRequestKeysStart, uint256 _decreasePositionRequestKeysStart) returns()
func (_PositionRouter *PositionRouterTransactorSession) SetRequestKeysStartValues(_increasePositionRequestKeysStart *big.Int, _decreasePositionRequestKeysStart *big.Int) (*types.Transaction, error) {
	return _PositionRouter.Contract.SetRequestKeysStartValues(&_PositionRouter.TransactOpts, _increasePositionRequestKeysStart, _decreasePositionRequestKeysStart)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns()
func (_PositionRouter *PositionRouterTransactor) WithdrawFees(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.contract.Transact(opts, "withdrawFees", _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns()
func (_PositionRouter *PositionRouterSession) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.WithdrawFees(&_PositionRouter.TransactOpts, _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns()
func (_PositionRouter *PositionRouterTransactorSession) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _PositionRouter.Contract.WithdrawFees(&_PositionRouter.TransactOpts, _token, _receiver)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PositionRouter *PositionRouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PositionRouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PositionRouter *PositionRouterSession) Receive() (*types.Transaction, error) {
	return _PositionRouter.Contract.Receive(&_PositionRouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PositionRouter *PositionRouterTransactorSession) Receive() (*types.Transaction, error) {
	return _PositionRouter.Contract.Receive(&_PositionRouter.TransactOpts)
}

// PositionRouterCallbackIterator is returned from FilterCallback and is used to iterate over the raw logs and unpacked data for Callback events raised by the PositionRouter contract.
type PositionRouterCallbackIterator struct {
	Event *PositionRouterCallback // Event containing the contract specifics and raw log

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
func (it *PositionRouterCallbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterCallback)
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
		it.Event = new(PositionRouterCallback)
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
func (it *PositionRouterCallbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterCallbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterCallback represents a Callback event raised by the PositionRouter contract.
type PositionRouterCallback struct {
	CallbackTarget   common.Address
	Success          bool
	CallbackGasLimit *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCallback is a free log retrieval operation binding the contract event 0xc9123a2a8e16684aa24686f2bf8a6d0eb0c601bdc109140c9729916865a58bc4.
//
// Solidity: event Callback(address callbackTarget, bool success, uint256 callbackGasLimit)
func (_PositionRouter *PositionRouterFilterer) FilterCallback(opts *bind.FilterOpts) (*PositionRouterCallbackIterator, error) {

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "Callback")
	if err != nil {
		return nil, err
	}
	return &PositionRouterCallbackIterator{contract: _PositionRouter.contract, event: "Callback", logs: logs, sub: sub}, nil
}

// WatchCallback is a free log subscription operation binding the contract event 0xc9123a2a8e16684aa24686f2bf8a6d0eb0c601bdc109140c9729916865a58bc4.
//
// Solidity: event Callback(address callbackTarget, bool success, uint256 callbackGasLimit)
func (_PositionRouter *PositionRouterFilterer) WatchCallback(opts *bind.WatchOpts, sink chan<- *PositionRouterCallback) (event.Subscription, error) {

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "Callback")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterCallback)
				if err := _PositionRouter.contract.UnpackLog(event, "Callback", log); err != nil {
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

// ParseCallback is a log parse operation binding the contract event 0xc9123a2a8e16684aa24686f2bf8a6d0eb0c601bdc109140c9729916865a58bc4.
//
// Solidity: event Callback(address callbackTarget, bool success, uint256 callbackGasLimit)
func (_PositionRouter *PositionRouterFilterer) ParseCallback(log types.Log) (*PositionRouterCallback, error) {
	event := new(PositionRouterCallback)
	if err := _PositionRouter.contract.UnpackLog(event, "Callback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterCancelDecreasePositionIterator is returned from FilterCancelDecreasePosition and is used to iterate over the raw logs and unpacked data for CancelDecreasePosition events raised by the PositionRouter contract.
type PositionRouterCancelDecreasePositionIterator struct {
	Event *PositionRouterCancelDecreasePosition // Event containing the contract specifics and raw log

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
func (it *PositionRouterCancelDecreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterCancelDecreasePosition)
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
		it.Event = new(PositionRouterCancelDecreasePosition)
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
func (it *PositionRouterCancelDecreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterCancelDecreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterCancelDecreasePosition represents a CancelDecreasePosition event raised by the PositionRouter contract.
type PositionRouterCancelDecreasePosition struct {
	Account         common.Address
	Path            []common.Address
	IndexToken      common.Address
	CollateralDelta *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	Receiver        common.Address
	AcceptablePrice *big.Int
	MinOut          *big.Int
	ExecutionFee    *big.Int
	BlockGap        *big.Int
	TimeGap         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCancelDecreasePosition is a free log retrieval operation binding the contract event 0x87abfd78e844f28318363bdf3da99eab2f4a2da9ff7ae365484507f7b6c3f805.
//
// Solidity: event CancelDecreasePosition(address indexed account, address[] path, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, address receiver, uint256 acceptablePrice, uint256 minOut, uint256 executionFee, uint256 blockGap, uint256 timeGap)
func (_PositionRouter *PositionRouterFilterer) FilterCancelDecreasePosition(opts *bind.FilterOpts, account []common.Address) (*PositionRouterCancelDecreasePositionIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "CancelDecreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return &PositionRouterCancelDecreasePositionIterator{contract: _PositionRouter.contract, event: "CancelDecreasePosition", logs: logs, sub: sub}, nil
}

// WatchCancelDecreasePosition is a free log subscription operation binding the contract event 0x87abfd78e844f28318363bdf3da99eab2f4a2da9ff7ae365484507f7b6c3f805.
//
// Solidity: event CancelDecreasePosition(address indexed account, address[] path, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, address receiver, uint256 acceptablePrice, uint256 minOut, uint256 executionFee, uint256 blockGap, uint256 timeGap)
func (_PositionRouter *PositionRouterFilterer) WatchCancelDecreasePosition(opts *bind.WatchOpts, sink chan<- *PositionRouterCancelDecreasePosition, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "CancelDecreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterCancelDecreasePosition)
				if err := _PositionRouter.contract.UnpackLog(event, "CancelDecreasePosition", log); err != nil {
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

// ParseCancelDecreasePosition is a log parse operation binding the contract event 0x87abfd78e844f28318363bdf3da99eab2f4a2da9ff7ae365484507f7b6c3f805.
//
// Solidity: event CancelDecreasePosition(address indexed account, address[] path, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, address receiver, uint256 acceptablePrice, uint256 minOut, uint256 executionFee, uint256 blockGap, uint256 timeGap)
func (_PositionRouter *PositionRouterFilterer) ParseCancelDecreasePosition(log types.Log) (*PositionRouterCancelDecreasePosition, error) {
	event := new(PositionRouterCancelDecreasePosition)
	if err := _PositionRouter.contract.UnpackLog(event, "CancelDecreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterCancelIncreasePositionIterator is returned from FilterCancelIncreasePosition and is used to iterate over the raw logs and unpacked data for CancelIncreasePosition events raised by the PositionRouter contract.
type PositionRouterCancelIncreasePositionIterator struct {
	Event *PositionRouterCancelIncreasePosition // Event containing the contract specifics and raw log

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
func (it *PositionRouterCancelIncreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterCancelIncreasePosition)
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
		it.Event = new(PositionRouterCancelIncreasePosition)
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
func (it *PositionRouterCancelIncreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterCancelIncreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterCancelIncreasePosition represents a CancelIncreasePosition event raised by the PositionRouter contract.
type PositionRouterCancelIncreasePosition struct {
	Account         common.Address
	Path            []common.Address
	IndexToken      common.Address
	AmountIn        *big.Int
	MinOut          *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	AcceptablePrice *big.Int
	ExecutionFee    *big.Int
	BlockGap        *big.Int
	TimeGap         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCancelIncreasePosition is a free log retrieval operation binding the contract event 0x35b638e650e2328786fb405bd69d2083dbedc018d086662e74b775b4f1dae4bf.
//
// Solidity: event CancelIncreasePosition(address indexed account, address[] path, address indexToken, uint256 amountIn, uint256 minOut, uint256 sizeDelta, bool isLong, uint256 acceptablePrice, uint256 executionFee, uint256 blockGap, uint256 timeGap)
func (_PositionRouter *PositionRouterFilterer) FilterCancelIncreasePosition(opts *bind.FilterOpts, account []common.Address) (*PositionRouterCancelIncreasePositionIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "CancelIncreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return &PositionRouterCancelIncreasePositionIterator{contract: _PositionRouter.contract, event: "CancelIncreasePosition", logs: logs, sub: sub}, nil
}

// WatchCancelIncreasePosition is a free log subscription operation binding the contract event 0x35b638e650e2328786fb405bd69d2083dbedc018d086662e74b775b4f1dae4bf.
//
// Solidity: event CancelIncreasePosition(address indexed account, address[] path, address indexToken, uint256 amountIn, uint256 minOut, uint256 sizeDelta, bool isLong, uint256 acceptablePrice, uint256 executionFee, uint256 blockGap, uint256 timeGap)
func (_PositionRouter *PositionRouterFilterer) WatchCancelIncreasePosition(opts *bind.WatchOpts, sink chan<- *PositionRouterCancelIncreasePosition, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "CancelIncreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterCancelIncreasePosition)
				if err := _PositionRouter.contract.UnpackLog(event, "CancelIncreasePosition", log); err != nil {
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

// ParseCancelIncreasePosition is a log parse operation binding the contract event 0x35b638e650e2328786fb405bd69d2083dbedc018d086662e74b775b4f1dae4bf.
//
// Solidity: event CancelIncreasePosition(address indexed account, address[] path, address indexToken, uint256 amountIn, uint256 minOut, uint256 sizeDelta, bool isLong, uint256 acceptablePrice, uint256 executionFee, uint256 blockGap, uint256 timeGap)
func (_PositionRouter *PositionRouterFilterer) ParseCancelIncreasePosition(log types.Log) (*PositionRouterCancelIncreasePosition, error) {
	event := new(PositionRouterCancelIncreasePosition)
	if err := _PositionRouter.contract.UnpackLog(event, "CancelIncreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterCreateDecreasePositionIterator is returned from FilterCreateDecreasePosition and is used to iterate over the raw logs and unpacked data for CreateDecreasePosition events raised by the PositionRouter contract.
type PositionRouterCreateDecreasePositionIterator struct {
	Event *PositionRouterCreateDecreasePosition // Event containing the contract specifics and raw log

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
func (it *PositionRouterCreateDecreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterCreateDecreasePosition)
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
		it.Event = new(PositionRouterCreateDecreasePosition)
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
func (it *PositionRouterCreateDecreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterCreateDecreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterCreateDecreasePosition represents a CreateDecreasePosition event raised by the PositionRouter contract.
type PositionRouterCreateDecreasePosition struct {
	Account         common.Address
	Path            []common.Address
	IndexToken      common.Address
	CollateralDelta *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	Receiver        common.Address
	AcceptablePrice *big.Int
	MinOut          *big.Int
	ExecutionFee    *big.Int
	Index           *big.Int
	QueueIndex      *big.Int
	BlockNumber     *big.Int
	BlockTime       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreateDecreasePosition is a free log retrieval operation binding the contract event 0x81ed0476a7e785a9e4728fffd679ea97176ca1ac85e1003462558bb5677da57b.
//
// Solidity: event CreateDecreasePosition(address indexed account, address[] path, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, address receiver, uint256 acceptablePrice, uint256 minOut, uint256 executionFee, uint256 index, uint256 queueIndex, uint256 blockNumber, uint256 blockTime)
func (_PositionRouter *PositionRouterFilterer) FilterCreateDecreasePosition(opts *bind.FilterOpts, account []common.Address) (*PositionRouterCreateDecreasePositionIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "CreateDecreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return &PositionRouterCreateDecreasePositionIterator{contract: _PositionRouter.contract, event: "CreateDecreasePosition", logs: logs, sub: sub}, nil
}

// WatchCreateDecreasePosition is a free log subscription operation binding the contract event 0x81ed0476a7e785a9e4728fffd679ea97176ca1ac85e1003462558bb5677da57b.
//
// Solidity: event CreateDecreasePosition(address indexed account, address[] path, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, address receiver, uint256 acceptablePrice, uint256 minOut, uint256 executionFee, uint256 index, uint256 queueIndex, uint256 blockNumber, uint256 blockTime)
func (_PositionRouter *PositionRouterFilterer) WatchCreateDecreasePosition(opts *bind.WatchOpts, sink chan<- *PositionRouterCreateDecreasePosition, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "CreateDecreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterCreateDecreasePosition)
				if err := _PositionRouter.contract.UnpackLog(event, "CreateDecreasePosition", log); err != nil {
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

// ParseCreateDecreasePosition is a log parse operation binding the contract event 0x81ed0476a7e785a9e4728fffd679ea97176ca1ac85e1003462558bb5677da57b.
//
// Solidity: event CreateDecreasePosition(address indexed account, address[] path, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, address receiver, uint256 acceptablePrice, uint256 minOut, uint256 executionFee, uint256 index, uint256 queueIndex, uint256 blockNumber, uint256 blockTime)
func (_PositionRouter *PositionRouterFilterer) ParseCreateDecreasePosition(log types.Log) (*PositionRouterCreateDecreasePosition, error) {
	event := new(PositionRouterCreateDecreasePosition)
	if err := _PositionRouter.contract.UnpackLog(event, "CreateDecreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterCreateIncreasePositionIterator is returned from FilterCreateIncreasePosition and is used to iterate over the raw logs and unpacked data for CreateIncreasePosition events raised by the PositionRouter contract.
type PositionRouterCreateIncreasePositionIterator struct {
	Event *PositionRouterCreateIncreasePosition // Event containing the contract specifics and raw log

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
func (it *PositionRouterCreateIncreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterCreateIncreasePosition)
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
		it.Event = new(PositionRouterCreateIncreasePosition)
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
func (it *PositionRouterCreateIncreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterCreateIncreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterCreateIncreasePosition represents a CreateIncreasePosition event raised by the PositionRouter contract.
type PositionRouterCreateIncreasePosition struct {
	Account         common.Address
	Path            []common.Address
	IndexToken      common.Address
	AmountIn        *big.Int
	MinOut          *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	AcceptablePrice *big.Int
	ExecutionFee    *big.Int
	Index           *big.Int
	QueueIndex      *big.Int
	BlockNumber     *big.Int
	BlockTime       *big.Int
	GasPrice        *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterCreateIncreasePosition is a free log retrieval operation binding the contract event 0x5265bc4952da402633b3fc35f67ab4245493a0ab94dd8ab123667c8d45a4485c.
//
// Solidity: event CreateIncreasePosition(address indexed account, address[] path, address indexToken, uint256 amountIn, uint256 minOut, uint256 sizeDelta, bool isLong, uint256 acceptablePrice, uint256 executionFee, uint256 index, uint256 queueIndex, uint256 blockNumber, uint256 blockTime, uint256 gasPrice)
func (_PositionRouter *PositionRouterFilterer) FilterCreateIncreasePosition(opts *bind.FilterOpts, account []common.Address) (*PositionRouterCreateIncreasePositionIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "CreateIncreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return &PositionRouterCreateIncreasePositionIterator{contract: _PositionRouter.contract, event: "CreateIncreasePosition", logs: logs, sub: sub}, nil
}

// WatchCreateIncreasePosition is a free log subscription operation binding the contract event 0x5265bc4952da402633b3fc35f67ab4245493a0ab94dd8ab123667c8d45a4485c.
//
// Solidity: event CreateIncreasePosition(address indexed account, address[] path, address indexToken, uint256 amountIn, uint256 minOut, uint256 sizeDelta, bool isLong, uint256 acceptablePrice, uint256 executionFee, uint256 index, uint256 queueIndex, uint256 blockNumber, uint256 blockTime, uint256 gasPrice)
func (_PositionRouter *PositionRouterFilterer) WatchCreateIncreasePosition(opts *bind.WatchOpts, sink chan<- *PositionRouterCreateIncreasePosition, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "CreateIncreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterCreateIncreasePosition)
				if err := _PositionRouter.contract.UnpackLog(event, "CreateIncreasePosition", log); err != nil {
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

// ParseCreateIncreasePosition is a log parse operation binding the contract event 0x5265bc4952da402633b3fc35f67ab4245493a0ab94dd8ab123667c8d45a4485c.
//
// Solidity: event CreateIncreasePosition(address indexed account, address[] path, address indexToken, uint256 amountIn, uint256 minOut, uint256 sizeDelta, bool isLong, uint256 acceptablePrice, uint256 executionFee, uint256 index, uint256 queueIndex, uint256 blockNumber, uint256 blockTime, uint256 gasPrice)
func (_PositionRouter *PositionRouterFilterer) ParseCreateIncreasePosition(log types.Log) (*PositionRouterCreateIncreasePosition, error) {
	event := new(PositionRouterCreateIncreasePosition)
	if err := _PositionRouter.contract.UnpackLog(event, "CreateIncreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterDecreasePositionReferralIterator is returned from FilterDecreasePositionReferral and is used to iterate over the raw logs and unpacked data for DecreasePositionReferral events raised by the PositionRouter contract.
type PositionRouterDecreasePositionReferralIterator struct {
	Event *PositionRouterDecreasePositionReferral // Event containing the contract specifics and raw log

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
func (it *PositionRouterDecreasePositionReferralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterDecreasePositionReferral)
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
		it.Event = new(PositionRouterDecreasePositionReferral)
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
func (it *PositionRouterDecreasePositionReferralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterDecreasePositionReferralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterDecreasePositionReferral represents a DecreasePositionReferral event raised by the PositionRouter contract.
type PositionRouterDecreasePositionReferral struct {
	Account              common.Address
	SizeDelta            *big.Int
	MarginFeeBasisPoints *big.Int
	ReferralCode         [32]byte
	Referrer             common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDecreasePositionReferral is a free log retrieval operation binding the contract event 0x474c763ff84bf2c2039a6d9fea955ecd0f724030e3c365b91169c6a16fe751b7.
//
// Solidity: event DecreasePositionReferral(address account, uint256 sizeDelta, uint256 marginFeeBasisPoints, bytes32 referralCode, address referrer)
func (_PositionRouter *PositionRouterFilterer) FilterDecreasePositionReferral(opts *bind.FilterOpts) (*PositionRouterDecreasePositionReferralIterator, error) {

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "DecreasePositionReferral")
	if err != nil {
		return nil, err
	}
	return &PositionRouterDecreasePositionReferralIterator{contract: _PositionRouter.contract, event: "DecreasePositionReferral", logs: logs, sub: sub}, nil
}

// WatchDecreasePositionReferral is a free log subscription operation binding the contract event 0x474c763ff84bf2c2039a6d9fea955ecd0f724030e3c365b91169c6a16fe751b7.
//
// Solidity: event DecreasePositionReferral(address account, uint256 sizeDelta, uint256 marginFeeBasisPoints, bytes32 referralCode, address referrer)
func (_PositionRouter *PositionRouterFilterer) WatchDecreasePositionReferral(opts *bind.WatchOpts, sink chan<- *PositionRouterDecreasePositionReferral) (event.Subscription, error) {

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "DecreasePositionReferral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterDecreasePositionReferral)
				if err := _PositionRouter.contract.UnpackLog(event, "DecreasePositionReferral", log); err != nil {
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

// ParseDecreasePositionReferral is a log parse operation binding the contract event 0x474c763ff84bf2c2039a6d9fea955ecd0f724030e3c365b91169c6a16fe751b7.
//
// Solidity: event DecreasePositionReferral(address account, uint256 sizeDelta, uint256 marginFeeBasisPoints, bytes32 referralCode, address referrer)
func (_PositionRouter *PositionRouterFilterer) ParseDecreasePositionReferral(log types.Log) (*PositionRouterDecreasePositionReferral, error) {
	event := new(PositionRouterDecreasePositionReferral)
	if err := _PositionRouter.contract.UnpackLog(event, "DecreasePositionReferral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterExecuteDecreasePositionIterator is returned from FilterExecuteDecreasePosition and is used to iterate over the raw logs and unpacked data for ExecuteDecreasePosition events raised by the PositionRouter contract.
type PositionRouterExecuteDecreasePositionIterator struct {
	Event *PositionRouterExecuteDecreasePosition // Event containing the contract specifics and raw log

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
func (it *PositionRouterExecuteDecreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterExecuteDecreasePosition)
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
		it.Event = new(PositionRouterExecuteDecreasePosition)
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
func (it *PositionRouterExecuteDecreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterExecuteDecreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterExecuteDecreasePosition represents a ExecuteDecreasePosition event raised by the PositionRouter contract.
type PositionRouterExecuteDecreasePosition struct {
	Account         common.Address
	Path            []common.Address
	IndexToken      common.Address
	CollateralDelta *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	Receiver        common.Address
	AcceptablePrice *big.Int
	MinOut          *big.Int
	ExecutionFee    *big.Int
	BlockGap        *big.Int
	TimeGap         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterExecuteDecreasePosition is a free log retrieval operation binding the contract event 0x21435c5b618d77ff3657140cd3318e2cffaebc5e0e1b7318f56a9ba4044c3ed2.
//
// Solidity: event ExecuteDecreasePosition(address indexed account, address[] path, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, address receiver, uint256 acceptablePrice, uint256 minOut, uint256 executionFee, uint256 blockGap, uint256 timeGap)
func (_PositionRouter *PositionRouterFilterer) FilterExecuteDecreasePosition(opts *bind.FilterOpts, account []common.Address) (*PositionRouterExecuteDecreasePositionIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "ExecuteDecreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return &PositionRouterExecuteDecreasePositionIterator{contract: _PositionRouter.contract, event: "ExecuteDecreasePosition", logs: logs, sub: sub}, nil
}

// WatchExecuteDecreasePosition is a free log subscription operation binding the contract event 0x21435c5b618d77ff3657140cd3318e2cffaebc5e0e1b7318f56a9ba4044c3ed2.
//
// Solidity: event ExecuteDecreasePosition(address indexed account, address[] path, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, address receiver, uint256 acceptablePrice, uint256 minOut, uint256 executionFee, uint256 blockGap, uint256 timeGap)
func (_PositionRouter *PositionRouterFilterer) WatchExecuteDecreasePosition(opts *bind.WatchOpts, sink chan<- *PositionRouterExecuteDecreasePosition, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "ExecuteDecreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterExecuteDecreasePosition)
				if err := _PositionRouter.contract.UnpackLog(event, "ExecuteDecreasePosition", log); err != nil {
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

// ParseExecuteDecreasePosition is a log parse operation binding the contract event 0x21435c5b618d77ff3657140cd3318e2cffaebc5e0e1b7318f56a9ba4044c3ed2.
//
// Solidity: event ExecuteDecreasePosition(address indexed account, address[] path, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, address receiver, uint256 acceptablePrice, uint256 minOut, uint256 executionFee, uint256 blockGap, uint256 timeGap)
func (_PositionRouter *PositionRouterFilterer) ParseExecuteDecreasePosition(log types.Log) (*PositionRouterExecuteDecreasePosition, error) {
	event := new(PositionRouterExecuteDecreasePosition)
	if err := _PositionRouter.contract.UnpackLog(event, "ExecuteDecreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterExecuteIncreasePositionIterator is returned from FilterExecuteIncreasePosition and is used to iterate over the raw logs and unpacked data for ExecuteIncreasePosition events raised by the PositionRouter contract.
type PositionRouterExecuteIncreasePositionIterator struct {
	Event *PositionRouterExecuteIncreasePosition // Event containing the contract specifics and raw log

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
func (it *PositionRouterExecuteIncreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterExecuteIncreasePosition)
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
		it.Event = new(PositionRouterExecuteIncreasePosition)
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
func (it *PositionRouterExecuteIncreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterExecuteIncreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterExecuteIncreasePosition represents a ExecuteIncreasePosition event raised by the PositionRouter contract.
type PositionRouterExecuteIncreasePosition struct {
	Account         common.Address
	Path            []common.Address
	IndexToken      common.Address
	AmountIn        *big.Int
	MinOut          *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	AcceptablePrice *big.Int
	ExecutionFee    *big.Int
	BlockGap        *big.Int
	TimeGap         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterExecuteIncreasePosition is a free log retrieval operation binding the contract event 0x1be316b94d38c07bd41cdb4913772d0a0a82802786a2f8b657b6e85dbcdfc641.
//
// Solidity: event ExecuteIncreasePosition(address indexed account, address[] path, address indexToken, uint256 amountIn, uint256 minOut, uint256 sizeDelta, bool isLong, uint256 acceptablePrice, uint256 executionFee, uint256 blockGap, uint256 timeGap)
func (_PositionRouter *PositionRouterFilterer) FilterExecuteIncreasePosition(opts *bind.FilterOpts, account []common.Address) (*PositionRouterExecuteIncreasePositionIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "ExecuteIncreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return &PositionRouterExecuteIncreasePositionIterator{contract: _PositionRouter.contract, event: "ExecuteIncreasePosition", logs: logs, sub: sub}, nil
}

// WatchExecuteIncreasePosition is a free log subscription operation binding the contract event 0x1be316b94d38c07bd41cdb4913772d0a0a82802786a2f8b657b6e85dbcdfc641.
//
// Solidity: event ExecuteIncreasePosition(address indexed account, address[] path, address indexToken, uint256 amountIn, uint256 minOut, uint256 sizeDelta, bool isLong, uint256 acceptablePrice, uint256 executionFee, uint256 blockGap, uint256 timeGap)
func (_PositionRouter *PositionRouterFilterer) WatchExecuteIncreasePosition(opts *bind.WatchOpts, sink chan<- *PositionRouterExecuteIncreasePosition, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "ExecuteIncreasePosition", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterExecuteIncreasePosition)
				if err := _PositionRouter.contract.UnpackLog(event, "ExecuteIncreasePosition", log); err != nil {
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

// ParseExecuteIncreasePosition is a log parse operation binding the contract event 0x1be316b94d38c07bd41cdb4913772d0a0a82802786a2f8b657b6e85dbcdfc641.
//
// Solidity: event ExecuteIncreasePosition(address indexed account, address[] path, address indexToken, uint256 amountIn, uint256 minOut, uint256 sizeDelta, bool isLong, uint256 acceptablePrice, uint256 executionFee, uint256 blockGap, uint256 timeGap)
func (_PositionRouter *PositionRouterFilterer) ParseExecuteIncreasePosition(log types.Log) (*PositionRouterExecuteIncreasePosition, error) {
	event := new(PositionRouterExecuteIncreasePosition)
	if err := _PositionRouter.contract.UnpackLog(event, "ExecuteIncreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterIncreasePositionReferralIterator is returned from FilterIncreasePositionReferral and is used to iterate over the raw logs and unpacked data for IncreasePositionReferral events raised by the PositionRouter contract.
type PositionRouterIncreasePositionReferralIterator struct {
	Event *PositionRouterIncreasePositionReferral // Event containing the contract specifics and raw log

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
func (it *PositionRouterIncreasePositionReferralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterIncreasePositionReferral)
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
		it.Event = new(PositionRouterIncreasePositionReferral)
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
func (it *PositionRouterIncreasePositionReferralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterIncreasePositionReferralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterIncreasePositionReferral represents a IncreasePositionReferral event raised by the PositionRouter contract.
type PositionRouterIncreasePositionReferral struct {
	Account              common.Address
	SizeDelta            *big.Int
	MarginFeeBasisPoints *big.Int
	ReferralCode         [32]byte
	Referrer             common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterIncreasePositionReferral is a free log retrieval operation binding the contract event 0xc2414023ce7002ee98557d1e7be21e5559073336f2217ee5f9b2e50fd85f71ee.
//
// Solidity: event IncreasePositionReferral(address account, uint256 sizeDelta, uint256 marginFeeBasisPoints, bytes32 referralCode, address referrer)
func (_PositionRouter *PositionRouterFilterer) FilterIncreasePositionReferral(opts *bind.FilterOpts) (*PositionRouterIncreasePositionReferralIterator, error) {

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "IncreasePositionReferral")
	if err != nil {
		return nil, err
	}
	return &PositionRouterIncreasePositionReferralIterator{contract: _PositionRouter.contract, event: "IncreasePositionReferral", logs: logs, sub: sub}, nil
}

// WatchIncreasePositionReferral is a free log subscription operation binding the contract event 0xc2414023ce7002ee98557d1e7be21e5559073336f2217ee5f9b2e50fd85f71ee.
//
// Solidity: event IncreasePositionReferral(address account, uint256 sizeDelta, uint256 marginFeeBasisPoints, bytes32 referralCode, address referrer)
func (_PositionRouter *PositionRouterFilterer) WatchIncreasePositionReferral(opts *bind.WatchOpts, sink chan<- *PositionRouterIncreasePositionReferral) (event.Subscription, error) {

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "IncreasePositionReferral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterIncreasePositionReferral)
				if err := _PositionRouter.contract.UnpackLog(event, "IncreasePositionReferral", log); err != nil {
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

// ParseIncreasePositionReferral is a log parse operation binding the contract event 0xc2414023ce7002ee98557d1e7be21e5559073336f2217ee5f9b2e50fd85f71ee.
//
// Solidity: event IncreasePositionReferral(address account, uint256 sizeDelta, uint256 marginFeeBasisPoints, bytes32 referralCode, address referrer)
func (_PositionRouter *PositionRouterFilterer) ParseIncreasePositionReferral(log types.Log) (*PositionRouterIncreasePositionReferral, error) {
	event := new(PositionRouterIncreasePositionReferral)
	if err := _PositionRouter.contract.UnpackLog(event, "IncreasePositionReferral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterSetAdminIterator is returned from FilterSetAdmin and is used to iterate over the raw logs and unpacked data for SetAdmin events raised by the PositionRouter contract.
type PositionRouterSetAdminIterator struct {
	Event *PositionRouterSetAdmin // Event containing the contract specifics and raw log

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
func (it *PositionRouterSetAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterSetAdmin)
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
		it.Event = new(PositionRouterSetAdmin)
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
func (it *PositionRouterSetAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterSetAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterSetAdmin represents a SetAdmin event raised by the PositionRouter contract.
type PositionRouterSetAdmin struct {
	Admin common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetAdmin is a free log retrieval operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_PositionRouter *PositionRouterFilterer) FilterSetAdmin(opts *bind.FilterOpts) (*PositionRouterSetAdminIterator, error) {

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return &PositionRouterSetAdminIterator{contract: _PositionRouter.contract, event: "SetAdmin", logs: logs, sub: sub}, nil
}

// WatchSetAdmin is a free log subscription operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_PositionRouter *PositionRouterFilterer) WatchSetAdmin(opts *bind.WatchOpts, sink chan<- *PositionRouterSetAdmin) (event.Subscription, error) {

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "SetAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterSetAdmin)
				if err := _PositionRouter.contract.UnpackLog(event, "SetAdmin", log); err != nil {
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

// ParseSetAdmin is a log parse operation binding the contract event 0x5a272403b402d892977df56625f4164ccaf70ca3863991c43ecfe76a6905b0a1.
//
// Solidity: event SetAdmin(address admin)
func (_PositionRouter *PositionRouterFilterer) ParseSetAdmin(log types.Log) (*PositionRouterSetAdmin, error) {
	event := new(PositionRouterSetAdmin)
	if err := _PositionRouter.contract.UnpackLog(event, "SetAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterSetCallbackGasLimitIterator is returned from FilterSetCallbackGasLimit and is used to iterate over the raw logs and unpacked data for SetCallbackGasLimit events raised by the PositionRouter contract.
type PositionRouterSetCallbackGasLimitIterator struct {
	Event *PositionRouterSetCallbackGasLimit // Event containing the contract specifics and raw log

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
func (it *PositionRouterSetCallbackGasLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterSetCallbackGasLimit)
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
		it.Event = new(PositionRouterSetCallbackGasLimit)
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
func (it *PositionRouterSetCallbackGasLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterSetCallbackGasLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterSetCallbackGasLimit represents a SetCallbackGasLimit event raised by the PositionRouter contract.
type PositionRouterSetCallbackGasLimit struct {
	CallbackGasLimit *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetCallbackGasLimit is a free log retrieval operation binding the contract event 0x22bd2c9f980325d046be74aaef5fc76df4a2bc3fbc7c5a1200fcc79fe80dab6c.
//
// Solidity: event SetCallbackGasLimit(uint256 callbackGasLimit)
func (_PositionRouter *PositionRouterFilterer) FilterSetCallbackGasLimit(opts *bind.FilterOpts) (*PositionRouterSetCallbackGasLimitIterator, error) {

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "SetCallbackGasLimit")
	if err != nil {
		return nil, err
	}
	return &PositionRouterSetCallbackGasLimitIterator{contract: _PositionRouter.contract, event: "SetCallbackGasLimit", logs: logs, sub: sub}, nil
}

// WatchSetCallbackGasLimit is a free log subscription operation binding the contract event 0x22bd2c9f980325d046be74aaef5fc76df4a2bc3fbc7c5a1200fcc79fe80dab6c.
//
// Solidity: event SetCallbackGasLimit(uint256 callbackGasLimit)
func (_PositionRouter *PositionRouterFilterer) WatchSetCallbackGasLimit(opts *bind.WatchOpts, sink chan<- *PositionRouterSetCallbackGasLimit) (event.Subscription, error) {

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "SetCallbackGasLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterSetCallbackGasLimit)
				if err := _PositionRouter.contract.UnpackLog(event, "SetCallbackGasLimit", log); err != nil {
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

// ParseSetCallbackGasLimit is a log parse operation binding the contract event 0x22bd2c9f980325d046be74aaef5fc76df4a2bc3fbc7c5a1200fcc79fe80dab6c.
//
// Solidity: event SetCallbackGasLimit(uint256 callbackGasLimit)
func (_PositionRouter *PositionRouterFilterer) ParseSetCallbackGasLimit(log types.Log) (*PositionRouterSetCallbackGasLimit, error) {
	event := new(PositionRouterSetCallbackGasLimit)
	if err := _PositionRouter.contract.UnpackLog(event, "SetCallbackGasLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterSetCustomCallbackGasLimitIterator is returned from FilterSetCustomCallbackGasLimit and is used to iterate over the raw logs and unpacked data for SetCustomCallbackGasLimit events raised by the PositionRouter contract.
type PositionRouterSetCustomCallbackGasLimitIterator struct {
	Event *PositionRouterSetCustomCallbackGasLimit // Event containing the contract specifics and raw log

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
func (it *PositionRouterSetCustomCallbackGasLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterSetCustomCallbackGasLimit)
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
		it.Event = new(PositionRouterSetCustomCallbackGasLimit)
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
func (it *PositionRouterSetCustomCallbackGasLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterSetCustomCallbackGasLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterSetCustomCallbackGasLimit represents a SetCustomCallbackGasLimit event raised by the PositionRouter contract.
type PositionRouterSetCustomCallbackGasLimit struct {
	CallbackTarget   common.Address
	CallbackGasLimit *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSetCustomCallbackGasLimit is a free log retrieval operation binding the contract event 0x0a585bcfd5f265014b902e5350c05f3a465468d433f13009dcf83f17dc1316be.
//
// Solidity: event SetCustomCallbackGasLimit(address callbackTarget, uint256 callbackGasLimit)
func (_PositionRouter *PositionRouterFilterer) FilterSetCustomCallbackGasLimit(opts *bind.FilterOpts) (*PositionRouterSetCustomCallbackGasLimitIterator, error) {

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "SetCustomCallbackGasLimit")
	if err != nil {
		return nil, err
	}
	return &PositionRouterSetCustomCallbackGasLimitIterator{contract: _PositionRouter.contract, event: "SetCustomCallbackGasLimit", logs: logs, sub: sub}, nil
}

// WatchSetCustomCallbackGasLimit is a free log subscription operation binding the contract event 0x0a585bcfd5f265014b902e5350c05f3a465468d433f13009dcf83f17dc1316be.
//
// Solidity: event SetCustomCallbackGasLimit(address callbackTarget, uint256 callbackGasLimit)
func (_PositionRouter *PositionRouterFilterer) WatchSetCustomCallbackGasLimit(opts *bind.WatchOpts, sink chan<- *PositionRouterSetCustomCallbackGasLimit) (event.Subscription, error) {

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "SetCustomCallbackGasLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterSetCustomCallbackGasLimit)
				if err := _PositionRouter.contract.UnpackLog(event, "SetCustomCallbackGasLimit", log); err != nil {
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

// ParseSetCustomCallbackGasLimit is a log parse operation binding the contract event 0x0a585bcfd5f265014b902e5350c05f3a465468d433f13009dcf83f17dc1316be.
//
// Solidity: event SetCustomCallbackGasLimit(address callbackTarget, uint256 callbackGasLimit)
func (_PositionRouter *PositionRouterFilterer) ParseSetCustomCallbackGasLimit(log types.Log) (*PositionRouterSetCustomCallbackGasLimit, error) {
	event := new(PositionRouterSetCustomCallbackGasLimit)
	if err := _PositionRouter.contract.UnpackLog(event, "SetCustomCallbackGasLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterSetDelayValuesIterator is returned from FilterSetDelayValues and is used to iterate over the raw logs and unpacked data for SetDelayValues events raised by the PositionRouter contract.
type PositionRouterSetDelayValuesIterator struct {
	Event *PositionRouterSetDelayValues // Event containing the contract specifics and raw log

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
func (it *PositionRouterSetDelayValuesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterSetDelayValues)
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
		it.Event = new(PositionRouterSetDelayValues)
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
func (it *PositionRouterSetDelayValuesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterSetDelayValuesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterSetDelayValues represents a SetDelayValues event raised by the PositionRouter contract.
type PositionRouterSetDelayValues struct {
	MinBlockDelayKeeper *big.Int
	MinTimeDelayPublic  *big.Int
	MaxTimeDelay        *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetDelayValues is a free log retrieval operation binding the contract event 0xb98e759701eaca2e60c25e91109003c1c7442ef731b5d569037063005da8254d.
//
// Solidity: event SetDelayValues(uint256 minBlockDelayKeeper, uint256 minTimeDelayPublic, uint256 maxTimeDelay)
func (_PositionRouter *PositionRouterFilterer) FilterSetDelayValues(opts *bind.FilterOpts) (*PositionRouterSetDelayValuesIterator, error) {

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "SetDelayValues")
	if err != nil {
		return nil, err
	}
	return &PositionRouterSetDelayValuesIterator{contract: _PositionRouter.contract, event: "SetDelayValues", logs: logs, sub: sub}, nil
}

// WatchSetDelayValues is a free log subscription operation binding the contract event 0xb98e759701eaca2e60c25e91109003c1c7442ef731b5d569037063005da8254d.
//
// Solidity: event SetDelayValues(uint256 minBlockDelayKeeper, uint256 minTimeDelayPublic, uint256 maxTimeDelay)
func (_PositionRouter *PositionRouterFilterer) WatchSetDelayValues(opts *bind.WatchOpts, sink chan<- *PositionRouterSetDelayValues) (event.Subscription, error) {

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "SetDelayValues")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterSetDelayValues)
				if err := _PositionRouter.contract.UnpackLog(event, "SetDelayValues", log); err != nil {
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

// ParseSetDelayValues is a log parse operation binding the contract event 0xb98e759701eaca2e60c25e91109003c1c7442ef731b5d569037063005da8254d.
//
// Solidity: event SetDelayValues(uint256 minBlockDelayKeeper, uint256 minTimeDelayPublic, uint256 maxTimeDelay)
func (_PositionRouter *PositionRouterFilterer) ParseSetDelayValues(log types.Log) (*PositionRouterSetDelayValues, error) {
	event := new(PositionRouterSetDelayValues)
	if err := _PositionRouter.contract.UnpackLog(event, "SetDelayValues", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterSetDepositFeeIterator is returned from FilterSetDepositFee and is used to iterate over the raw logs and unpacked data for SetDepositFee events raised by the PositionRouter contract.
type PositionRouterSetDepositFeeIterator struct {
	Event *PositionRouterSetDepositFee // Event containing the contract specifics and raw log

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
func (it *PositionRouterSetDepositFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterSetDepositFee)
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
		it.Event = new(PositionRouterSetDepositFee)
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
func (it *PositionRouterSetDepositFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterSetDepositFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterSetDepositFee represents a SetDepositFee event raised by the PositionRouter contract.
type PositionRouterSetDepositFee struct {
	DepositFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetDepositFee is a free log retrieval operation binding the contract event 0x974fd3c1fcb4653dfc4fb740c4c692cd212d55c28f163f310128cb64d8300675.
//
// Solidity: event SetDepositFee(uint256 depositFee)
func (_PositionRouter *PositionRouterFilterer) FilterSetDepositFee(opts *bind.FilterOpts) (*PositionRouterSetDepositFeeIterator, error) {

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "SetDepositFee")
	if err != nil {
		return nil, err
	}
	return &PositionRouterSetDepositFeeIterator{contract: _PositionRouter.contract, event: "SetDepositFee", logs: logs, sub: sub}, nil
}

// WatchSetDepositFee is a free log subscription operation binding the contract event 0x974fd3c1fcb4653dfc4fb740c4c692cd212d55c28f163f310128cb64d8300675.
//
// Solidity: event SetDepositFee(uint256 depositFee)
func (_PositionRouter *PositionRouterFilterer) WatchSetDepositFee(opts *bind.WatchOpts, sink chan<- *PositionRouterSetDepositFee) (event.Subscription, error) {

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "SetDepositFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterSetDepositFee)
				if err := _PositionRouter.contract.UnpackLog(event, "SetDepositFee", log); err != nil {
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

// ParseSetDepositFee is a log parse operation binding the contract event 0x974fd3c1fcb4653dfc4fb740c4c692cd212d55c28f163f310128cb64d8300675.
//
// Solidity: event SetDepositFee(uint256 depositFee)
func (_PositionRouter *PositionRouterFilterer) ParseSetDepositFee(log types.Log) (*PositionRouterSetDepositFee, error) {
	event := new(PositionRouterSetDepositFee)
	if err := _PositionRouter.contract.UnpackLog(event, "SetDepositFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterSetEthTransferGasLimitIterator is returned from FilterSetEthTransferGasLimit and is used to iterate over the raw logs and unpacked data for SetEthTransferGasLimit events raised by the PositionRouter contract.
type PositionRouterSetEthTransferGasLimitIterator struct {
	Event *PositionRouterSetEthTransferGasLimit // Event containing the contract specifics and raw log

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
func (it *PositionRouterSetEthTransferGasLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterSetEthTransferGasLimit)
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
		it.Event = new(PositionRouterSetEthTransferGasLimit)
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
func (it *PositionRouterSetEthTransferGasLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterSetEthTransferGasLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterSetEthTransferGasLimit represents a SetEthTransferGasLimit event raised by the PositionRouter contract.
type PositionRouterSetEthTransferGasLimit struct {
	EthTransferGasLimit *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSetEthTransferGasLimit is a free log retrieval operation binding the contract event 0x4d371d598d3a13f99ce992a17975bbaf1e1c256e072ec7d2f93ce88e40d9ba1c.
//
// Solidity: event SetEthTransferGasLimit(uint256 ethTransferGasLimit)
func (_PositionRouter *PositionRouterFilterer) FilterSetEthTransferGasLimit(opts *bind.FilterOpts) (*PositionRouterSetEthTransferGasLimitIterator, error) {

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "SetEthTransferGasLimit")
	if err != nil {
		return nil, err
	}
	return &PositionRouterSetEthTransferGasLimitIterator{contract: _PositionRouter.contract, event: "SetEthTransferGasLimit", logs: logs, sub: sub}, nil
}

// WatchSetEthTransferGasLimit is a free log subscription operation binding the contract event 0x4d371d598d3a13f99ce992a17975bbaf1e1c256e072ec7d2f93ce88e40d9ba1c.
//
// Solidity: event SetEthTransferGasLimit(uint256 ethTransferGasLimit)
func (_PositionRouter *PositionRouterFilterer) WatchSetEthTransferGasLimit(opts *bind.WatchOpts, sink chan<- *PositionRouterSetEthTransferGasLimit) (event.Subscription, error) {

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "SetEthTransferGasLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterSetEthTransferGasLimit)
				if err := _PositionRouter.contract.UnpackLog(event, "SetEthTransferGasLimit", log); err != nil {
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

// ParseSetEthTransferGasLimit is a log parse operation binding the contract event 0x4d371d598d3a13f99ce992a17975bbaf1e1c256e072ec7d2f93ce88e40d9ba1c.
//
// Solidity: event SetEthTransferGasLimit(uint256 ethTransferGasLimit)
func (_PositionRouter *PositionRouterFilterer) ParseSetEthTransferGasLimit(log types.Log) (*PositionRouterSetEthTransferGasLimit, error) {
	event := new(PositionRouterSetEthTransferGasLimit)
	if err := _PositionRouter.contract.UnpackLog(event, "SetEthTransferGasLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterSetIncreasePositionBufferBpsIterator is returned from FilterSetIncreasePositionBufferBps and is used to iterate over the raw logs and unpacked data for SetIncreasePositionBufferBps events raised by the PositionRouter contract.
type PositionRouterSetIncreasePositionBufferBpsIterator struct {
	Event *PositionRouterSetIncreasePositionBufferBps // Event containing the contract specifics and raw log

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
func (it *PositionRouterSetIncreasePositionBufferBpsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterSetIncreasePositionBufferBps)
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
		it.Event = new(PositionRouterSetIncreasePositionBufferBps)
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
func (it *PositionRouterSetIncreasePositionBufferBpsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterSetIncreasePositionBufferBpsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterSetIncreasePositionBufferBps represents a SetIncreasePositionBufferBps event raised by the PositionRouter contract.
type PositionRouterSetIncreasePositionBufferBps struct {
	IncreasePositionBufferBps *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterSetIncreasePositionBufferBps is a free log retrieval operation binding the contract event 0x21167d0d4661af93817ebce920f18986eed3d75d5e1c03f2aed05efcbafbc452.
//
// Solidity: event SetIncreasePositionBufferBps(uint256 increasePositionBufferBps)
func (_PositionRouter *PositionRouterFilterer) FilterSetIncreasePositionBufferBps(opts *bind.FilterOpts) (*PositionRouterSetIncreasePositionBufferBpsIterator, error) {

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "SetIncreasePositionBufferBps")
	if err != nil {
		return nil, err
	}
	return &PositionRouterSetIncreasePositionBufferBpsIterator{contract: _PositionRouter.contract, event: "SetIncreasePositionBufferBps", logs: logs, sub: sub}, nil
}

// WatchSetIncreasePositionBufferBps is a free log subscription operation binding the contract event 0x21167d0d4661af93817ebce920f18986eed3d75d5e1c03f2aed05efcbafbc452.
//
// Solidity: event SetIncreasePositionBufferBps(uint256 increasePositionBufferBps)
func (_PositionRouter *PositionRouterFilterer) WatchSetIncreasePositionBufferBps(opts *bind.WatchOpts, sink chan<- *PositionRouterSetIncreasePositionBufferBps) (event.Subscription, error) {

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "SetIncreasePositionBufferBps")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterSetIncreasePositionBufferBps)
				if err := _PositionRouter.contract.UnpackLog(event, "SetIncreasePositionBufferBps", log); err != nil {
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

// ParseSetIncreasePositionBufferBps is a log parse operation binding the contract event 0x21167d0d4661af93817ebce920f18986eed3d75d5e1c03f2aed05efcbafbc452.
//
// Solidity: event SetIncreasePositionBufferBps(uint256 increasePositionBufferBps)
func (_PositionRouter *PositionRouterFilterer) ParseSetIncreasePositionBufferBps(log types.Log) (*PositionRouterSetIncreasePositionBufferBps, error) {
	event := new(PositionRouterSetIncreasePositionBufferBps)
	if err := _PositionRouter.contract.UnpackLog(event, "SetIncreasePositionBufferBps", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterSetIsLeverageEnabledIterator is returned from FilterSetIsLeverageEnabled and is used to iterate over the raw logs and unpacked data for SetIsLeverageEnabled events raised by the PositionRouter contract.
type PositionRouterSetIsLeverageEnabledIterator struct {
	Event *PositionRouterSetIsLeverageEnabled // Event containing the contract specifics and raw log

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
func (it *PositionRouterSetIsLeverageEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterSetIsLeverageEnabled)
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
		it.Event = new(PositionRouterSetIsLeverageEnabled)
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
func (it *PositionRouterSetIsLeverageEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterSetIsLeverageEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterSetIsLeverageEnabled represents a SetIsLeverageEnabled event raised by the PositionRouter contract.
type PositionRouterSetIsLeverageEnabled struct {
	IsLeverageEnabled bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSetIsLeverageEnabled is a free log retrieval operation binding the contract event 0x4eb87a5935d402aa24c01b45bfb30adefcd2328b480f2d967864de4b64ea929f.
//
// Solidity: event SetIsLeverageEnabled(bool isLeverageEnabled)
func (_PositionRouter *PositionRouterFilterer) FilterSetIsLeverageEnabled(opts *bind.FilterOpts) (*PositionRouterSetIsLeverageEnabledIterator, error) {

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "SetIsLeverageEnabled")
	if err != nil {
		return nil, err
	}
	return &PositionRouterSetIsLeverageEnabledIterator{contract: _PositionRouter.contract, event: "SetIsLeverageEnabled", logs: logs, sub: sub}, nil
}

// WatchSetIsLeverageEnabled is a free log subscription operation binding the contract event 0x4eb87a5935d402aa24c01b45bfb30adefcd2328b480f2d967864de4b64ea929f.
//
// Solidity: event SetIsLeverageEnabled(bool isLeverageEnabled)
func (_PositionRouter *PositionRouterFilterer) WatchSetIsLeverageEnabled(opts *bind.WatchOpts, sink chan<- *PositionRouterSetIsLeverageEnabled) (event.Subscription, error) {

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "SetIsLeverageEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterSetIsLeverageEnabled)
				if err := _PositionRouter.contract.UnpackLog(event, "SetIsLeverageEnabled", log); err != nil {
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

// ParseSetIsLeverageEnabled is a log parse operation binding the contract event 0x4eb87a5935d402aa24c01b45bfb30adefcd2328b480f2d967864de4b64ea929f.
//
// Solidity: event SetIsLeverageEnabled(bool isLeverageEnabled)
func (_PositionRouter *PositionRouterFilterer) ParseSetIsLeverageEnabled(log types.Log) (*PositionRouterSetIsLeverageEnabled, error) {
	event := new(PositionRouterSetIsLeverageEnabled)
	if err := _PositionRouter.contract.UnpackLog(event, "SetIsLeverageEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterSetMaxGlobalSizesIterator is returned from FilterSetMaxGlobalSizes and is used to iterate over the raw logs and unpacked data for SetMaxGlobalSizes events raised by the PositionRouter contract.
type PositionRouterSetMaxGlobalSizesIterator struct {
	Event *PositionRouterSetMaxGlobalSizes // Event containing the contract specifics and raw log

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
func (it *PositionRouterSetMaxGlobalSizesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterSetMaxGlobalSizes)
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
		it.Event = new(PositionRouterSetMaxGlobalSizes)
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
func (it *PositionRouterSetMaxGlobalSizesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterSetMaxGlobalSizesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterSetMaxGlobalSizes represents a SetMaxGlobalSizes event raised by the PositionRouter contract.
type PositionRouterSetMaxGlobalSizes struct {
	Tokens     []common.Address
	LongSizes  []*big.Int
	ShortSizes []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetMaxGlobalSizes is a free log retrieval operation binding the contract event 0xae32d569b058895b9620d6552b09aaffedc9a6f396be4d595a224ad09f8b2139.
//
// Solidity: event SetMaxGlobalSizes(address[] tokens, uint256[] longSizes, uint256[] shortSizes)
func (_PositionRouter *PositionRouterFilterer) FilterSetMaxGlobalSizes(opts *bind.FilterOpts) (*PositionRouterSetMaxGlobalSizesIterator, error) {

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "SetMaxGlobalSizes")
	if err != nil {
		return nil, err
	}
	return &PositionRouterSetMaxGlobalSizesIterator{contract: _PositionRouter.contract, event: "SetMaxGlobalSizes", logs: logs, sub: sub}, nil
}

// WatchSetMaxGlobalSizes is a free log subscription operation binding the contract event 0xae32d569b058895b9620d6552b09aaffedc9a6f396be4d595a224ad09f8b2139.
//
// Solidity: event SetMaxGlobalSizes(address[] tokens, uint256[] longSizes, uint256[] shortSizes)
func (_PositionRouter *PositionRouterFilterer) WatchSetMaxGlobalSizes(opts *bind.WatchOpts, sink chan<- *PositionRouterSetMaxGlobalSizes) (event.Subscription, error) {

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "SetMaxGlobalSizes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterSetMaxGlobalSizes)
				if err := _PositionRouter.contract.UnpackLog(event, "SetMaxGlobalSizes", log); err != nil {
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

// ParseSetMaxGlobalSizes is a log parse operation binding the contract event 0xae32d569b058895b9620d6552b09aaffedc9a6f396be4d595a224ad09f8b2139.
//
// Solidity: event SetMaxGlobalSizes(address[] tokens, uint256[] longSizes, uint256[] shortSizes)
func (_PositionRouter *PositionRouterFilterer) ParseSetMaxGlobalSizes(log types.Log) (*PositionRouterSetMaxGlobalSizes, error) {
	event := new(PositionRouterSetMaxGlobalSizes)
	if err := _PositionRouter.contract.UnpackLog(event, "SetMaxGlobalSizes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterSetMinExecutionFeeIterator is returned from FilterSetMinExecutionFee and is used to iterate over the raw logs and unpacked data for SetMinExecutionFee events raised by the PositionRouter contract.
type PositionRouterSetMinExecutionFeeIterator struct {
	Event *PositionRouterSetMinExecutionFee // Event containing the contract specifics and raw log

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
func (it *PositionRouterSetMinExecutionFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterSetMinExecutionFee)
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
		it.Event = new(PositionRouterSetMinExecutionFee)
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
func (it *PositionRouterSetMinExecutionFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterSetMinExecutionFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterSetMinExecutionFee represents a SetMinExecutionFee event raised by the PositionRouter contract.
type PositionRouterSetMinExecutionFee struct {
	MinExecutionFee *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetMinExecutionFee is a free log retrieval operation binding the contract event 0x52a8358457e20bbb36e4086b83fb0749599f1893fe4c35a876c46dc4886d12db.
//
// Solidity: event SetMinExecutionFee(uint256 minExecutionFee)
func (_PositionRouter *PositionRouterFilterer) FilterSetMinExecutionFee(opts *bind.FilterOpts) (*PositionRouterSetMinExecutionFeeIterator, error) {

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "SetMinExecutionFee")
	if err != nil {
		return nil, err
	}
	return &PositionRouterSetMinExecutionFeeIterator{contract: _PositionRouter.contract, event: "SetMinExecutionFee", logs: logs, sub: sub}, nil
}

// WatchSetMinExecutionFee is a free log subscription operation binding the contract event 0x52a8358457e20bbb36e4086b83fb0749599f1893fe4c35a876c46dc4886d12db.
//
// Solidity: event SetMinExecutionFee(uint256 minExecutionFee)
func (_PositionRouter *PositionRouterFilterer) WatchSetMinExecutionFee(opts *bind.WatchOpts, sink chan<- *PositionRouterSetMinExecutionFee) (event.Subscription, error) {

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "SetMinExecutionFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterSetMinExecutionFee)
				if err := _PositionRouter.contract.UnpackLog(event, "SetMinExecutionFee", log); err != nil {
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

// ParseSetMinExecutionFee is a log parse operation binding the contract event 0x52a8358457e20bbb36e4086b83fb0749599f1893fe4c35a876c46dc4886d12db.
//
// Solidity: event SetMinExecutionFee(uint256 minExecutionFee)
func (_PositionRouter *PositionRouterFilterer) ParseSetMinExecutionFee(log types.Log) (*PositionRouterSetMinExecutionFee, error) {
	event := new(PositionRouterSetMinExecutionFee)
	if err := _PositionRouter.contract.UnpackLog(event, "SetMinExecutionFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterSetPositionKeeperIterator is returned from FilterSetPositionKeeper and is used to iterate over the raw logs and unpacked data for SetPositionKeeper events raised by the PositionRouter contract.
type PositionRouterSetPositionKeeperIterator struct {
	Event *PositionRouterSetPositionKeeper // Event containing the contract specifics and raw log

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
func (it *PositionRouterSetPositionKeeperIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterSetPositionKeeper)
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
		it.Event = new(PositionRouterSetPositionKeeper)
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
func (it *PositionRouterSetPositionKeeperIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterSetPositionKeeperIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterSetPositionKeeper represents a SetPositionKeeper event raised by the PositionRouter contract.
type PositionRouterSetPositionKeeper struct {
	Account  common.Address
	IsActive bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetPositionKeeper is a free log retrieval operation binding the contract event 0xfbabc02389290a451c6e600d05bf9887b99bfad39d8e1237e4e3df042e4941fe.
//
// Solidity: event SetPositionKeeper(address indexed account, bool isActive)
func (_PositionRouter *PositionRouterFilterer) FilterSetPositionKeeper(opts *bind.FilterOpts, account []common.Address) (*PositionRouterSetPositionKeeperIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "SetPositionKeeper", accountRule)
	if err != nil {
		return nil, err
	}
	return &PositionRouterSetPositionKeeperIterator{contract: _PositionRouter.contract, event: "SetPositionKeeper", logs: logs, sub: sub}, nil
}

// WatchSetPositionKeeper is a free log subscription operation binding the contract event 0xfbabc02389290a451c6e600d05bf9887b99bfad39d8e1237e4e3df042e4941fe.
//
// Solidity: event SetPositionKeeper(address indexed account, bool isActive)
func (_PositionRouter *PositionRouterFilterer) WatchSetPositionKeeper(opts *bind.WatchOpts, sink chan<- *PositionRouterSetPositionKeeper, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "SetPositionKeeper", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterSetPositionKeeper)
				if err := _PositionRouter.contract.UnpackLog(event, "SetPositionKeeper", log); err != nil {
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

// ParseSetPositionKeeper is a log parse operation binding the contract event 0xfbabc02389290a451c6e600d05bf9887b99bfad39d8e1237e4e3df042e4941fe.
//
// Solidity: event SetPositionKeeper(address indexed account, bool isActive)
func (_PositionRouter *PositionRouterFilterer) ParseSetPositionKeeper(log types.Log) (*PositionRouterSetPositionKeeper, error) {
	event := new(PositionRouterSetPositionKeeper)
	if err := _PositionRouter.contract.UnpackLog(event, "SetPositionKeeper", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterSetReferralStorageIterator is returned from FilterSetReferralStorage and is used to iterate over the raw logs and unpacked data for SetReferralStorage events raised by the PositionRouter contract.
type PositionRouterSetReferralStorageIterator struct {
	Event *PositionRouterSetReferralStorage // Event containing the contract specifics and raw log

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
func (it *PositionRouterSetReferralStorageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterSetReferralStorage)
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
		it.Event = new(PositionRouterSetReferralStorage)
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
func (it *PositionRouterSetReferralStorageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterSetReferralStorageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterSetReferralStorage represents a SetReferralStorage event raised by the PositionRouter contract.
type PositionRouterSetReferralStorage struct {
	ReferralStorage common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSetReferralStorage is a free log retrieval operation binding the contract event 0x828abcccea18192c21d645e575652c49e20b986dab777906fc473d056b01b6a8.
//
// Solidity: event SetReferralStorage(address referralStorage)
func (_PositionRouter *PositionRouterFilterer) FilterSetReferralStorage(opts *bind.FilterOpts) (*PositionRouterSetReferralStorageIterator, error) {

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "SetReferralStorage")
	if err != nil {
		return nil, err
	}
	return &PositionRouterSetReferralStorageIterator{contract: _PositionRouter.contract, event: "SetReferralStorage", logs: logs, sub: sub}, nil
}

// WatchSetReferralStorage is a free log subscription operation binding the contract event 0x828abcccea18192c21d645e575652c49e20b986dab777906fc473d056b01b6a8.
//
// Solidity: event SetReferralStorage(address referralStorage)
func (_PositionRouter *PositionRouterFilterer) WatchSetReferralStorage(opts *bind.WatchOpts, sink chan<- *PositionRouterSetReferralStorage) (event.Subscription, error) {

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "SetReferralStorage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterSetReferralStorage)
				if err := _PositionRouter.contract.UnpackLog(event, "SetReferralStorage", log); err != nil {
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

// ParseSetReferralStorage is a log parse operation binding the contract event 0x828abcccea18192c21d645e575652c49e20b986dab777906fc473d056b01b6a8.
//
// Solidity: event SetReferralStorage(address referralStorage)
func (_PositionRouter *PositionRouterFilterer) ParseSetReferralStorage(log types.Log) (*PositionRouterSetReferralStorage, error) {
	event := new(PositionRouterSetReferralStorage)
	if err := _PositionRouter.contract.UnpackLog(event, "SetReferralStorage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterSetRequestKeysStartValuesIterator is returned from FilterSetRequestKeysStartValues and is used to iterate over the raw logs and unpacked data for SetRequestKeysStartValues events raised by the PositionRouter contract.
type PositionRouterSetRequestKeysStartValuesIterator struct {
	Event *PositionRouterSetRequestKeysStartValues // Event containing the contract specifics and raw log

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
func (it *PositionRouterSetRequestKeysStartValuesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterSetRequestKeysStartValues)
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
		it.Event = new(PositionRouterSetRequestKeysStartValues)
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
func (it *PositionRouterSetRequestKeysStartValuesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterSetRequestKeysStartValuesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterSetRequestKeysStartValues represents a SetRequestKeysStartValues event raised by the PositionRouter contract.
type PositionRouterSetRequestKeysStartValues struct {
	IncreasePositionRequestKeysStart *big.Int
	DecreasePositionRequestKeysStart *big.Int
	Raw                              types.Log // Blockchain specific contextual infos
}

// FilterSetRequestKeysStartValues is a free log retrieval operation binding the contract event 0xebb0f666150f4be5b60c45df8f3e49992510b0128027fe58eea6110f296493bc.
//
// Solidity: event SetRequestKeysStartValues(uint256 increasePositionRequestKeysStart, uint256 decreasePositionRequestKeysStart)
func (_PositionRouter *PositionRouterFilterer) FilterSetRequestKeysStartValues(opts *bind.FilterOpts) (*PositionRouterSetRequestKeysStartValuesIterator, error) {

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "SetRequestKeysStartValues")
	if err != nil {
		return nil, err
	}
	return &PositionRouterSetRequestKeysStartValuesIterator{contract: _PositionRouter.contract, event: "SetRequestKeysStartValues", logs: logs, sub: sub}, nil
}

// WatchSetRequestKeysStartValues is a free log subscription operation binding the contract event 0xebb0f666150f4be5b60c45df8f3e49992510b0128027fe58eea6110f296493bc.
//
// Solidity: event SetRequestKeysStartValues(uint256 increasePositionRequestKeysStart, uint256 decreasePositionRequestKeysStart)
func (_PositionRouter *PositionRouterFilterer) WatchSetRequestKeysStartValues(opts *bind.WatchOpts, sink chan<- *PositionRouterSetRequestKeysStartValues) (event.Subscription, error) {

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "SetRequestKeysStartValues")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterSetRequestKeysStartValues)
				if err := _PositionRouter.contract.UnpackLog(event, "SetRequestKeysStartValues", log); err != nil {
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

// ParseSetRequestKeysStartValues is a log parse operation binding the contract event 0xebb0f666150f4be5b60c45df8f3e49992510b0128027fe58eea6110f296493bc.
//
// Solidity: event SetRequestKeysStartValues(uint256 increasePositionRequestKeysStart, uint256 decreasePositionRequestKeysStart)
func (_PositionRouter *PositionRouterFilterer) ParseSetRequestKeysStartValues(log types.Log) (*PositionRouterSetRequestKeysStartValues, error) {
	event := new(PositionRouterSetRequestKeysStartValues)
	if err := _PositionRouter.contract.UnpackLog(event, "SetRequestKeysStartValues", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PositionRouterWithdrawFeesIterator is returned from FilterWithdrawFees and is used to iterate over the raw logs and unpacked data for WithdrawFees events raised by the PositionRouter contract.
type PositionRouterWithdrawFeesIterator struct {
	Event *PositionRouterWithdrawFees // Event containing the contract specifics and raw log

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
func (it *PositionRouterWithdrawFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PositionRouterWithdrawFees)
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
		it.Event = new(PositionRouterWithdrawFees)
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
func (it *PositionRouterWithdrawFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PositionRouterWithdrawFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PositionRouterWithdrawFees represents a WithdrawFees event raised by the PositionRouter contract.
type PositionRouterWithdrawFees struct {
	Token    common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFees is a free log retrieval operation binding the contract event 0x4f1b51dd7a2fcb861aa2670f668be66835c4ee12b4bbbf037e4d0018f39819e4.
//
// Solidity: event WithdrawFees(address token, address receiver, uint256 amount)
func (_PositionRouter *PositionRouterFilterer) FilterWithdrawFees(opts *bind.FilterOpts) (*PositionRouterWithdrawFeesIterator, error) {

	logs, sub, err := _PositionRouter.contract.FilterLogs(opts, "WithdrawFees")
	if err != nil {
		return nil, err
	}
	return &PositionRouterWithdrawFeesIterator{contract: _PositionRouter.contract, event: "WithdrawFees", logs: logs, sub: sub}, nil
}

// WatchWithdrawFees is a free log subscription operation binding the contract event 0x4f1b51dd7a2fcb861aa2670f668be66835c4ee12b4bbbf037e4d0018f39819e4.
//
// Solidity: event WithdrawFees(address token, address receiver, uint256 amount)
func (_PositionRouter *PositionRouterFilterer) WatchWithdrawFees(opts *bind.WatchOpts, sink chan<- *PositionRouterWithdrawFees) (event.Subscription, error) {

	logs, sub, err := _PositionRouter.contract.WatchLogs(opts, "WithdrawFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PositionRouterWithdrawFees)
				if err := _PositionRouter.contract.UnpackLog(event, "WithdrawFees", log); err != nil {
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

// ParseWithdrawFees is a log parse operation binding the contract event 0x4f1b51dd7a2fcb861aa2670f668be66835c4ee12b4bbbf037e4d0018f39819e4.
//
// Solidity: event WithdrawFees(address token, address receiver, uint256 amount)
func (_PositionRouter *PositionRouterFilterer) ParseWithdrawFees(log types.Log) (*PositionRouterWithdrawFees, error) {
	event := new(PositionRouterWithdrawFees)
	if err := _PositionRouter.contract.UnpackLog(event, "WithdrawFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

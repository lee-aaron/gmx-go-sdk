// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OrderBook

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

// OrderBookMetaData contains all meta data concerning the OrderBook contract.
var OrderBookMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"PRICE_PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"USDG_PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelDecreaseOrder\",\"inputs\":[{\"name\":\"_orderIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelIncreaseOrder\",\"inputs\":[{\"name\":\"_orderIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelMultiple\",\"inputs\":[{\"name\":\"_swapOrderIndexes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_increaseOrderIndexes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_decreaseOrderIndexes\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelSwapOrder\",\"inputs\":[{\"name\":\"_orderIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createDecreaseOrder\",\"inputs\":[{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_triggerPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_triggerAboveThreshold\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"createIncreaseOrder\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_triggerPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_triggerAboveThreshold\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_shouldWrap\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"createSwapOrder\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_triggerRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_triggerAboveThreshold\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_shouldWrap\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_shouldUnwrap\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"decreaseOrders\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"triggerPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"triggerAboveThreshold\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseOrdersIndex\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeDecreaseOrder\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_orderIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_feeReceiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeIncreaseOrder\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_orderIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_feeReceiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"executeSwapOrder\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_orderIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_feeReceiver\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDecreaseOrder\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_orderIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"triggerPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"triggerAboveThreshold\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIncreaseOrder\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_orderIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"purchaseToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"purchaseTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"triggerPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"triggerAboveThreshold\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSwapOrder\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_orderIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"path0\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"path1\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"path2\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"triggerRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"triggerAboveThreshold\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"shouldUnwrap\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUsdgMinPrice\",\"inputs\":[{\"name\":\"_otherToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseOrders\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"purchaseToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"purchaseTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"triggerPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"triggerAboveThreshold\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseOrdersIndex\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_weth\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdg\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_minExecutionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minPurchaseTokenAmountUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isInitialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minExecutionFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minPurchaseTokenAmountUsd\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinExecutionFee\",\"inputs\":[{\"name\":\"_minExecutionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinPurchaseTokenAmountUsd\",\"inputs\":[{\"name\":\"_minPurchaseTokenAmountUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swapOrders\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"triggerRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"triggerAboveThreshold\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"shouldUnwrap\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"swapOrdersIndex\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateDecreaseOrder\",\"inputs\":[{\"name\":\"_orderIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_triggerPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_triggerAboveThreshold\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateIncreaseOrder\",\"inputs\":[{\"name\":\"_orderIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_triggerPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_triggerAboveThreshold\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateSwapOrder\",\"inputs\":[{\"name\":\"_orderIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minOut\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_triggerRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_triggerAboveThreshold\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"usdg\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatePositionOrderPrice\",\"inputs\":[{\"name\":\"_triggerAboveThreshold\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_triggerPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_maximizePrice\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_raise\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateSwapOrderPriceWithTriggerAboveThreshold\",\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_triggerRatio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vault\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"weth\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CancelDecreaseOrder\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"orderIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"triggerPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"triggerAboveThreshold\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CancelIncreaseOrder\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"orderIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"purchaseToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"purchaseTokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"triggerPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"triggerAboveThreshold\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CancelSwapOrder\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"orderIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"path\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"minOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"triggerRatio\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"triggerAboveThreshold\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"shouldUnwrap\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CreateDecreaseOrder\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"orderIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"triggerPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"triggerAboveThreshold\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CreateIncreaseOrder\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"orderIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"purchaseToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"purchaseTokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"triggerPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"triggerAboveThreshold\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CreateSwapOrder\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"orderIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"path\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"minOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"triggerRatio\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"triggerAboveThreshold\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"shouldUnwrap\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecuteDecreaseOrder\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"orderIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"triggerPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"triggerAboveThreshold\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"executionPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecuteIncreaseOrder\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"orderIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"purchaseToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"purchaseTokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"triggerPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"triggerAboveThreshold\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"executionPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecuteSwapOrder\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"orderIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"path\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"minOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"triggerRatio\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"triggerAboveThreshold\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"shouldUnwrap\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialize\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"weth\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"usdg\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"minExecutionFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"minPurchaseTokenAmountUsd\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateDecreaseOrder\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"orderIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"triggerPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"triggerAboveThreshold\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateGov\",\"inputs\":[{\"name\":\"gov\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateIncreaseOrder\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"orderIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"triggerPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"triggerAboveThreshold\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateMinExecutionFee\",\"inputs\":[{\"name\":\"minExecutionFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateMinPurchaseTokenAmountUsd\",\"inputs\":[{\"name\":\"minPurchaseTokenAmountUsd\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateSwapOrder\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ordexIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"path\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"minOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"triggerRatio\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"triggerAboveThreshold\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"shouldUnwrap\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"executionFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// OrderBookABI is the input ABI used to generate the binding from.
// Deprecated: Use OrderBookMetaData.ABI instead.
var OrderBookABI = OrderBookMetaData.ABI

// OrderBook is an auto generated Go binding around an Ethereum contract.
type OrderBook struct {
	OrderBookCaller     // Read-only binding to the contract
	OrderBookTransactor // Write-only binding to the contract
	OrderBookFilterer   // Log filterer for contract events
}

// OrderBookCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderBookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderBookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderBookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderBookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderBookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderBookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderBookSession struct {
	Contract     *OrderBook        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderBookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderBookCallerSession struct {
	Contract *OrderBookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OrderBookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderBookTransactorSession struct {
	Contract     *OrderBookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OrderBookRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderBookRaw struct {
	Contract *OrderBook // Generic contract binding to access the raw methods on
}

// OrderBookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderBookCallerRaw struct {
	Contract *OrderBookCaller // Generic read-only contract binding to access the raw methods on
}

// OrderBookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderBookTransactorRaw struct {
	Contract *OrderBookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderBook creates a new instance of OrderBook, bound to a specific deployed contract.
func NewOrderBook(address common.Address, backend bind.ContractBackend) (*OrderBook, error) {
	contract, err := bindOrderBook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrderBook{OrderBookCaller: OrderBookCaller{contract: contract}, OrderBookTransactor: OrderBookTransactor{contract: contract}, OrderBookFilterer: OrderBookFilterer{contract: contract}}, nil
}

// NewOrderBookCaller creates a new read-only instance of OrderBook, bound to a specific deployed contract.
func NewOrderBookCaller(address common.Address, caller bind.ContractCaller) (*OrderBookCaller, error) {
	contract, err := bindOrderBook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderBookCaller{contract: contract}, nil
}

// NewOrderBookTransactor creates a new write-only instance of OrderBook, bound to a specific deployed contract.
func NewOrderBookTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderBookTransactor, error) {
	contract, err := bindOrderBook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderBookTransactor{contract: contract}, nil
}

// NewOrderBookFilterer creates a new log filterer instance of OrderBook, bound to a specific deployed contract.
func NewOrderBookFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderBookFilterer, error) {
	contract, err := bindOrderBook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderBookFilterer{contract: contract}, nil
}

// bindOrderBook binds a generic wrapper to an already deployed contract.
func bindOrderBook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OrderBookMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderBook *OrderBookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderBook.Contract.OrderBookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderBook *OrderBookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBook.Contract.OrderBookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderBook *OrderBookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderBook.Contract.OrderBookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderBook *OrderBookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderBook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderBook *OrderBookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderBook *OrderBookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderBook.Contract.contract.Transact(opts, method, params...)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_OrderBook *OrderBookCaller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_OrderBook *OrderBookSession) PRICEPRECISION() (*big.Int, error) {
	return _OrderBook.Contract.PRICEPRECISION(&_OrderBook.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_OrderBook *OrderBookCallerSession) PRICEPRECISION() (*big.Int, error) {
	return _OrderBook.Contract.PRICEPRECISION(&_OrderBook.CallOpts)
}

// USDGPRECISION is a free data retrieval call binding the contract method 0x4a686d67.
//
// Solidity: function USDG_PRECISION() view returns(uint256)
func (_OrderBook *OrderBookCaller) USDGPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "USDG_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// USDGPRECISION is a free data retrieval call binding the contract method 0x4a686d67.
//
// Solidity: function USDG_PRECISION() view returns(uint256)
func (_OrderBook *OrderBookSession) USDGPRECISION() (*big.Int, error) {
	return _OrderBook.Contract.USDGPRECISION(&_OrderBook.CallOpts)
}

// USDGPRECISION is a free data retrieval call binding the contract method 0x4a686d67.
//
// Solidity: function USDG_PRECISION() view returns(uint256)
func (_OrderBook *OrderBookCallerSession) USDGPRECISION() (*big.Int, error) {
	return _OrderBook.Contract.USDGPRECISION(&_OrderBook.CallOpts)
}

// DecreaseOrders is a free data retrieval call binding the contract method 0xf2d2e01b.
//
// Solidity: function decreaseOrders(address , uint256 ) view returns(address account, address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookCaller) DecreaseOrders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Account               common.Address
	CollateralToken       common.Address
	CollateralDelta       *big.Int
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
}, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "decreaseOrders", arg0, arg1)

	outstruct := new(struct {
		Account               common.Address
		CollateralToken       common.Address
		CollateralDelta       *big.Int
		IndexToken            common.Address
		SizeDelta             *big.Int
		IsLong                bool
		TriggerPrice          *big.Int
		TriggerAboveThreshold bool
		ExecutionFee          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Account = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CollateralToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.CollateralDelta = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IndexToken = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.SizeDelta = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.IsLong = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.TriggerPrice = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.TriggerAboveThreshold = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.ExecutionFee = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DecreaseOrders is a free data retrieval call binding the contract method 0xf2d2e01b.
//
// Solidity: function decreaseOrders(address , uint256 ) view returns(address account, address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookSession) DecreaseOrders(arg0 common.Address, arg1 *big.Int) (struct {
	Account               common.Address
	CollateralToken       common.Address
	CollateralDelta       *big.Int
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
}, error) {
	return _OrderBook.Contract.DecreaseOrders(&_OrderBook.CallOpts, arg0, arg1)
}

// DecreaseOrders is a free data retrieval call binding the contract method 0xf2d2e01b.
//
// Solidity: function decreaseOrders(address , uint256 ) view returns(address account, address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookCallerSession) DecreaseOrders(arg0 common.Address, arg1 *big.Int) (struct {
	Account               common.Address
	CollateralToken       common.Address
	CollateralDelta       *big.Int
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
}, error) {
	return _OrderBook.Contract.DecreaseOrders(&_OrderBook.CallOpts, arg0, arg1)
}

// DecreaseOrdersIndex is a free data retrieval call binding the contract method 0xd566d0ca.
//
// Solidity: function decreaseOrdersIndex(address ) view returns(uint256)
func (_OrderBook *OrderBookCaller) DecreaseOrdersIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "decreaseOrdersIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecreaseOrdersIndex is a free data retrieval call binding the contract method 0xd566d0ca.
//
// Solidity: function decreaseOrdersIndex(address ) view returns(uint256)
func (_OrderBook *OrderBookSession) DecreaseOrdersIndex(arg0 common.Address) (*big.Int, error) {
	return _OrderBook.Contract.DecreaseOrdersIndex(&_OrderBook.CallOpts, arg0)
}

// DecreaseOrdersIndex is a free data retrieval call binding the contract method 0xd566d0ca.
//
// Solidity: function decreaseOrdersIndex(address ) view returns(uint256)
func (_OrderBook *OrderBookCallerSession) DecreaseOrdersIndex(arg0 common.Address) (*big.Int, error) {
	return _OrderBook.Contract.DecreaseOrdersIndex(&_OrderBook.CallOpts, arg0)
}

// GetDecreaseOrder is a free data retrieval call binding the contract method 0x026032ee.
//
// Solidity: function getDecreaseOrder(address _account, uint256 _orderIndex) view returns(address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookCaller) GetDecreaseOrder(opts *bind.CallOpts, _account common.Address, _orderIndex *big.Int) (struct {
	CollateralToken       common.Address
	CollateralDelta       *big.Int
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
}, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "getDecreaseOrder", _account, _orderIndex)

	outstruct := new(struct {
		CollateralToken       common.Address
		CollateralDelta       *big.Int
		IndexToken            common.Address
		SizeDelta             *big.Int
		IsLong                bool
		TriggerPrice          *big.Int
		TriggerAboveThreshold bool
		ExecutionFee          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CollateralToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CollateralDelta = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IndexToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.SizeDelta = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsLong = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.TriggerPrice = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TriggerAboveThreshold = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.ExecutionFee = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetDecreaseOrder is a free data retrieval call binding the contract method 0x026032ee.
//
// Solidity: function getDecreaseOrder(address _account, uint256 _orderIndex) view returns(address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookSession) GetDecreaseOrder(_account common.Address, _orderIndex *big.Int) (struct {
	CollateralToken       common.Address
	CollateralDelta       *big.Int
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
}, error) {
	return _OrderBook.Contract.GetDecreaseOrder(&_OrderBook.CallOpts, _account, _orderIndex)
}

// GetDecreaseOrder is a free data retrieval call binding the contract method 0x026032ee.
//
// Solidity: function getDecreaseOrder(address _account, uint256 _orderIndex) view returns(address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookCallerSession) GetDecreaseOrder(_account common.Address, _orderIndex *big.Int) (struct {
	CollateralToken       common.Address
	CollateralDelta       *big.Int
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
}, error) {
	return _OrderBook.Contract.GetDecreaseOrder(&_OrderBook.CallOpts, _account, _orderIndex)
}

// GetIncreaseOrder is a free data retrieval call binding the contract method 0xd3bab1d1.
//
// Solidity: function getIncreaseOrder(address _account, uint256 _orderIndex) view returns(address purchaseToken, uint256 purchaseTokenAmount, address collateralToken, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookCaller) GetIncreaseOrder(opts *bind.CallOpts, _account common.Address, _orderIndex *big.Int) (struct {
	PurchaseToken         common.Address
	PurchaseTokenAmount   *big.Int
	CollateralToken       common.Address
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
}, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "getIncreaseOrder", _account, _orderIndex)

	outstruct := new(struct {
		PurchaseToken         common.Address
		PurchaseTokenAmount   *big.Int
		CollateralToken       common.Address
		IndexToken            common.Address
		SizeDelta             *big.Int
		IsLong                bool
		TriggerPrice          *big.Int
		TriggerAboveThreshold bool
		ExecutionFee          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PurchaseToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PurchaseTokenAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CollateralToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.IndexToken = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.SizeDelta = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.IsLong = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.TriggerPrice = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.TriggerAboveThreshold = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.ExecutionFee = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetIncreaseOrder is a free data retrieval call binding the contract method 0xd3bab1d1.
//
// Solidity: function getIncreaseOrder(address _account, uint256 _orderIndex) view returns(address purchaseToken, uint256 purchaseTokenAmount, address collateralToken, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookSession) GetIncreaseOrder(_account common.Address, _orderIndex *big.Int) (struct {
	PurchaseToken         common.Address
	PurchaseTokenAmount   *big.Int
	CollateralToken       common.Address
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
}, error) {
	return _OrderBook.Contract.GetIncreaseOrder(&_OrderBook.CallOpts, _account, _orderIndex)
}

// GetIncreaseOrder is a free data retrieval call binding the contract method 0xd3bab1d1.
//
// Solidity: function getIncreaseOrder(address _account, uint256 _orderIndex) view returns(address purchaseToken, uint256 purchaseTokenAmount, address collateralToken, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookCallerSession) GetIncreaseOrder(_account common.Address, _orderIndex *big.Int) (struct {
	PurchaseToken         common.Address
	PurchaseTokenAmount   *big.Int
	CollateralToken       common.Address
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
}, error) {
	return _OrderBook.Contract.GetIncreaseOrder(&_OrderBook.CallOpts, _account, _orderIndex)
}

// GetSwapOrder is a free data retrieval call binding the contract method 0xd0d40cd6.
//
// Solidity: function getSwapOrder(address _account, uint256 _orderIndex) view returns(address path0, address path1, address path2, uint256 amountIn, uint256 minOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_OrderBook *OrderBookCaller) GetSwapOrder(opts *bind.CallOpts, _account common.Address, _orderIndex *big.Int) (struct {
	Path0                 common.Address
	Path1                 common.Address
	Path2                 common.Address
	AmountIn              *big.Int
	MinOut                *big.Int
	TriggerRatio          *big.Int
	TriggerAboveThreshold bool
	ShouldUnwrap          bool
	ExecutionFee          *big.Int
}, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "getSwapOrder", _account, _orderIndex)

	outstruct := new(struct {
		Path0                 common.Address
		Path1                 common.Address
		Path2                 common.Address
		AmountIn              *big.Int
		MinOut                *big.Int
		TriggerRatio          *big.Int
		TriggerAboveThreshold bool
		ShouldUnwrap          bool
		ExecutionFee          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Path0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Path1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Path2 = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.AmountIn = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.MinOut = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TriggerRatio = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TriggerAboveThreshold = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.ShouldUnwrap = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.ExecutionFee = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSwapOrder is a free data retrieval call binding the contract method 0xd0d40cd6.
//
// Solidity: function getSwapOrder(address _account, uint256 _orderIndex) view returns(address path0, address path1, address path2, uint256 amountIn, uint256 minOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_OrderBook *OrderBookSession) GetSwapOrder(_account common.Address, _orderIndex *big.Int) (struct {
	Path0                 common.Address
	Path1                 common.Address
	Path2                 common.Address
	AmountIn              *big.Int
	MinOut                *big.Int
	TriggerRatio          *big.Int
	TriggerAboveThreshold bool
	ShouldUnwrap          bool
	ExecutionFee          *big.Int
}, error) {
	return _OrderBook.Contract.GetSwapOrder(&_OrderBook.CallOpts, _account, _orderIndex)
}

// GetSwapOrder is a free data retrieval call binding the contract method 0xd0d40cd6.
//
// Solidity: function getSwapOrder(address _account, uint256 _orderIndex) view returns(address path0, address path1, address path2, uint256 amountIn, uint256 minOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_OrderBook *OrderBookCallerSession) GetSwapOrder(_account common.Address, _orderIndex *big.Int) (struct {
	Path0                 common.Address
	Path1                 common.Address
	Path2                 common.Address
	AmountIn              *big.Int
	MinOut                *big.Int
	TriggerRatio          *big.Int
	TriggerAboveThreshold bool
	ShouldUnwrap          bool
	ExecutionFee          *big.Int
}, error) {
	return _OrderBook.Contract.GetSwapOrder(&_OrderBook.CallOpts, _account, _orderIndex)
}

// GetUsdgMinPrice is a free data retrieval call binding the contract method 0x9e23de5c.
//
// Solidity: function getUsdgMinPrice(address _otherToken) view returns(uint256)
func (_OrderBook *OrderBookCaller) GetUsdgMinPrice(opts *bind.CallOpts, _otherToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "getUsdgMinPrice", _otherToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUsdgMinPrice is a free data retrieval call binding the contract method 0x9e23de5c.
//
// Solidity: function getUsdgMinPrice(address _otherToken) view returns(uint256)
func (_OrderBook *OrderBookSession) GetUsdgMinPrice(_otherToken common.Address) (*big.Int, error) {
	return _OrderBook.Contract.GetUsdgMinPrice(&_OrderBook.CallOpts, _otherToken)
}

// GetUsdgMinPrice is a free data retrieval call binding the contract method 0x9e23de5c.
//
// Solidity: function getUsdgMinPrice(address _otherToken) view returns(uint256)
func (_OrderBook *OrderBookCallerSession) GetUsdgMinPrice(_otherToken common.Address) (*big.Int, error) {
	return _OrderBook.Contract.GetUsdgMinPrice(&_OrderBook.CallOpts, _otherToken)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_OrderBook *OrderBookCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_OrderBook *OrderBookSession) Gov() (common.Address, error) {
	return _OrderBook.Contract.Gov(&_OrderBook.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_OrderBook *OrderBookCallerSession) Gov() (common.Address, error) {
	return _OrderBook.Contract.Gov(&_OrderBook.CallOpts)
}

// IncreaseOrders is a free data retrieval call binding the contract method 0x2b7d6290.
//
// Solidity: function increaseOrders(address , uint256 ) view returns(address account, address purchaseToken, uint256 purchaseTokenAmount, address collateralToken, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookCaller) IncreaseOrders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Account               common.Address
	PurchaseToken         common.Address
	PurchaseTokenAmount   *big.Int
	CollateralToken       common.Address
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
}, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "increaseOrders", arg0, arg1)

	outstruct := new(struct {
		Account               common.Address
		PurchaseToken         common.Address
		PurchaseTokenAmount   *big.Int
		CollateralToken       common.Address
		IndexToken            common.Address
		SizeDelta             *big.Int
		IsLong                bool
		TriggerPrice          *big.Int
		TriggerAboveThreshold bool
		ExecutionFee          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Account = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.PurchaseToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.PurchaseTokenAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CollateralToken = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.IndexToken = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.SizeDelta = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.IsLong = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.TriggerPrice = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.TriggerAboveThreshold = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.ExecutionFee = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// IncreaseOrders is a free data retrieval call binding the contract method 0x2b7d6290.
//
// Solidity: function increaseOrders(address , uint256 ) view returns(address account, address purchaseToken, uint256 purchaseTokenAmount, address collateralToken, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookSession) IncreaseOrders(arg0 common.Address, arg1 *big.Int) (struct {
	Account               common.Address
	PurchaseToken         common.Address
	PurchaseTokenAmount   *big.Int
	CollateralToken       common.Address
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
}, error) {
	return _OrderBook.Contract.IncreaseOrders(&_OrderBook.CallOpts, arg0, arg1)
}

// IncreaseOrders is a free data retrieval call binding the contract method 0x2b7d6290.
//
// Solidity: function increaseOrders(address , uint256 ) view returns(address account, address purchaseToken, uint256 purchaseTokenAmount, address collateralToken, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookCallerSession) IncreaseOrders(arg0 common.Address, arg1 *big.Int) (struct {
	Account               common.Address
	PurchaseToken         common.Address
	PurchaseTokenAmount   *big.Int
	CollateralToken       common.Address
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
}, error) {
	return _OrderBook.Contract.IncreaseOrders(&_OrderBook.CallOpts, arg0, arg1)
}

// IncreaseOrdersIndex is a free data retrieval call binding the contract method 0xaec22455.
//
// Solidity: function increaseOrdersIndex(address ) view returns(uint256)
func (_OrderBook *OrderBookCaller) IncreaseOrdersIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "increaseOrdersIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncreaseOrdersIndex is a free data retrieval call binding the contract method 0xaec22455.
//
// Solidity: function increaseOrdersIndex(address ) view returns(uint256)
func (_OrderBook *OrderBookSession) IncreaseOrdersIndex(arg0 common.Address) (*big.Int, error) {
	return _OrderBook.Contract.IncreaseOrdersIndex(&_OrderBook.CallOpts, arg0)
}

// IncreaseOrdersIndex is a free data retrieval call binding the contract method 0xaec22455.
//
// Solidity: function increaseOrdersIndex(address ) view returns(uint256)
func (_OrderBook *OrderBookCallerSession) IncreaseOrdersIndex(arg0 common.Address) (*big.Int, error) {
	return _OrderBook.Contract.IncreaseOrdersIndex(&_OrderBook.CallOpts, arg0)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_OrderBook *OrderBookCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_OrderBook *OrderBookSession) IsInitialized() (bool, error) {
	return _OrderBook.Contract.IsInitialized(&_OrderBook.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_OrderBook *OrderBookCallerSession) IsInitialized() (bool, error) {
	return _OrderBook.Contract.IsInitialized(&_OrderBook.CallOpts)
}

// MinExecutionFee is a free data retrieval call binding the contract method 0x63ae2103.
//
// Solidity: function minExecutionFee() view returns(uint256)
func (_OrderBook *OrderBookCaller) MinExecutionFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "minExecutionFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinExecutionFee is a free data retrieval call binding the contract method 0x63ae2103.
//
// Solidity: function minExecutionFee() view returns(uint256)
func (_OrderBook *OrderBookSession) MinExecutionFee() (*big.Int, error) {
	return _OrderBook.Contract.MinExecutionFee(&_OrderBook.CallOpts)
}

// MinExecutionFee is a free data retrieval call binding the contract method 0x63ae2103.
//
// Solidity: function minExecutionFee() view returns(uint256)
func (_OrderBook *OrderBookCallerSession) MinExecutionFee() (*big.Int, error) {
	return _OrderBook.Contract.MinExecutionFee(&_OrderBook.CallOpts)
}

// MinPurchaseTokenAmountUsd is a free data retrieval call binding the contract method 0x8de10c2e.
//
// Solidity: function minPurchaseTokenAmountUsd() view returns(uint256)
func (_OrderBook *OrderBookCaller) MinPurchaseTokenAmountUsd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "minPurchaseTokenAmountUsd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinPurchaseTokenAmountUsd is a free data retrieval call binding the contract method 0x8de10c2e.
//
// Solidity: function minPurchaseTokenAmountUsd() view returns(uint256)
func (_OrderBook *OrderBookSession) MinPurchaseTokenAmountUsd() (*big.Int, error) {
	return _OrderBook.Contract.MinPurchaseTokenAmountUsd(&_OrderBook.CallOpts)
}

// MinPurchaseTokenAmountUsd is a free data retrieval call binding the contract method 0x8de10c2e.
//
// Solidity: function minPurchaseTokenAmountUsd() view returns(uint256)
func (_OrderBook *OrderBookCallerSession) MinPurchaseTokenAmountUsd() (*big.Int, error) {
	return _OrderBook.Contract.MinPurchaseTokenAmountUsd(&_OrderBook.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_OrderBook *OrderBookCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_OrderBook *OrderBookSession) Router() (common.Address, error) {
	return _OrderBook.Contract.Router(&_OrderBook.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_OrderBook *OrderBookCallerSession) Router() (common.Address, error) {
	return _OrderBook.Contract.Router(&_OrderBook.CallOpts)
}

// SwapOrders is a free data retrieval call binding the contract method 0x79221fa2.
//
// Solidity: function swapOrders(address , uint256 ) view returns(address account, uint256 amountIn, uint256 minOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_OrderBook *OrderBookCaller) SwapOrders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Account               common.Address
	AmountIn              *big.Int
	MinOut                *big.Int
	TriggerRatio          *big.Int
	TriggerAboveThreshold bool
	ShouldUnwrap          bool
	ExecutionFee          *big.Int
}, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "swapOrders", arg0, arg1)

	outstruct := new(struct {
		Account               common.Address
		AmountIn              *big.Int
		MinOut                *big.Int
		TriggerRatio          *big.Int
		TriggerAboveThreshold bool
		ShouldUnwrap          bool
		ExecutionFee          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Account = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AmountIn = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MinOut = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TriggerRatio = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TriggerAboveThreshold = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.ShouldUnwrap = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.ExecutionFee = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SwapOrders is a free data retrieval call binding the contract method 0x79221fa2.
//
// Solidity: function swapOrders(address , uint256 ) view returns(address account, uint256 amountIn, uint256 minOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_OrderBook *OrderBookSession) SwapOrders(arg0 common.Address, arg1 *big.Int) (struct {
	Account               common.Address
	AmountIn              *big.Int
	MinOut                *big.Int
	TriggerRatio          *big.Int
	TriggerAboveThreshold bool
	ShouldUnwrap          bool
	ExecutionFee          *big.Int
}, error) {
	return _OrderBook.Contract.SwapOrders(&_OrderBook.CallOpts, arg0, arg1)
}

// SwapOrders is a free data retrieval call binding the contract method 0x79221fa2.
//
// Solidity: function swapOrders(address , uint256 ) view returns(address account, uint256 amountIn, uint256 minOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_OrderBook *OrderBookCallerSession) SwapOrders(arg0 common.Address, arg1 *big.Int) (struct {
	Account               common.Address
	AmountIn              *big.Int
	MinOut                *big.Int
	TriggerRatio          *big.Int
	TriggerAboveThreshold bool
	ShouldUnwrap          bool
	ExecutionFee          *big.Int
}, error) {
	return _OrderBook.Contract.SwapOrders(&_OrderBook.CallOpts, arg0, arg1)
}

// SwapOrdersIndex is a free data retrieval call binding the contract method 0x00cf066b.
//
// Solidity: function swapOrdersIndex(address ) view returns(uint256)
func (_OrderBook *OrderBookCaller) SwapOrdersIndex(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "swapOrdersIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapOrdersIndex is a free data retrieval call binding the contract method 0x00cf066b.
//
// Solidity: function swapOrdersIndex(address ) view returns(uint256)
func (_OrderBook *OrderBookSession) SwapOrdersIndex(arg0 common.Address) (*big.Int, error) {
	return _OrderBook.Contract.SwapOrdersIndex(&_OrderBook.CallOpts, arg0)
}

// SwapOrdersIndex is a free data retrieval call binding the contract method 0x00cf066b.
//
// Solidity: function swapOrdersIndex(address ) view returns(uint256)
func (_OrderBook *OrderBookCallerSession) SwapOrdersIndex(arg0 common.Address) (*big.Int, error) {
	return _OrderBook.Contract.SwapOrdersIndex(&_OrderBook.CallOpts, arg0)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_OrderBook *OrderBookCaller) Usdg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "usdg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_OrderBook *OrderBookSession) Usdg() (common.Address, error) {
	return _OrderBook.Contract.Usdg(&_OrderBook.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_OrderBook *OrderBookCallerSession) Usdg() (common.Address, error) {
	return _OrderBook.Contract.Usdg(&_OrderBook.CallOpts)
}

// ValidatePositionOrderPrice is a free data retrieval call binding the contract method 0x4c54f0b0.
//
// Solidity: function validatePositionOrderPrice(bool _triggerAboveThreshold, uint256 _triggerPrice, address _indexToken, bool _maximizePrice, bool _raise) view returns(uint256, bool)
func (_OrderBook *OrderBookCaller) ValidatePositionOrderPrice(opts *bind.CallOpts, _triggerAboveThreshold bool, _triggerPrice *big.Int, _indexToken common.Address, _maximizePrice bool, _raise bool) (*big.Int, bool, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "validatePositionOrderPrice", _triggerAboveThreshold, _triggerPrice, _indexToken, _maximizePrice, _raise)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// ValidatePositionOrderPrice is a free data retrieval call binding the contract method 0x4c54f0b0.
//
// Solidity: function validatePositionOrderPrice(bool _triggerAboveThreshold, uint256 _triggerPrice, address _indexToken, bool _maximizePrice, bool _raise) view returns(uint256, bool)
func (_OrderBook *OrderBookSession) ValidatePositionOrderPrice(_triggerAboveThreshold bool, _triggerPrice *big.Int, _indexToken common.Address, _maximizePrice bool, _raise bool) (*big.Int, bool, error) {
	return _OrderBook.Contract.ValidatePositionOrderPrice(&_OrderBook.CallOpts, _triggerAboveThreshold, _triggerPrice, _indexToken, _maximizePrice, _raise)
}

// ValidatePositionOrderPrice is a free data retrieval call binding the contract method 0x4c54f0b0.
//
// Solidity: function validatePositionOrderPrice(bool _triggerAboveThreshold, uint256 _triggerPrice, address _indexToken, bool _maximizePrice, bool _raise) view returns(uint256, bool)
func (_OrderBook *OrderBookCallerSession) ValidatePositionOrderPrice(_triggerAboveThreshold bool, _triggerPrice *big.Int, _indexToken common.Address, _maximizePrice bool, _raise bool) (*big.Int, bool, error) {
	return _OrderBook.Contract.ValidatePositionOrderPrice(&_OrderBook.CallOpts, _triggerAboveThreshold, _triggerPrice, _indexToken, _maximizePrice, _raise)
}

// ValidateSwapOrderPriceWithTriggerAboveThreshold is a free data retrieval call binding the contract method 0xc4a1821b.
//
// Solidity: function validateSwapOrderPriceWithTriggerAboveThreshold(address[] _path, uint256 _triggerRatio) view returns(bool)
func (_OrderBook *OrderBookCaller) ValidateSwapOrderPriceWithTriggerAboveThreshold(opts *bind.CallOpts, _path []common.Address, _triggerRatio *big.Int) (bool, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "validateSwapOrderPriceWithTriggerAboveThreshold", _path, _triggerRatio)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateSwapOrderPriceWithTriggerAboveThreshold is a free data retrieval call binding the contract method 0xc4a1821b.
//
// Solidity: function validateSwapOrderPriceWithTriggerAboveThreshold(address[] _path, uint256 _triggerRatio) view returns(bool)
func (_OrderBook *OrderBookSession) ValidateSwapOrderPriceWithTriggerAboveThreshold(_path []common.Address, _triggerRatio *big.Int) (bool, error) {
	return _OrderBook.Contract.ValidateSwapOrderPriceWithTriggerAboveThreshold(&_OrderBook.CallOpts, _path, _triggerRatio)
}

// ValidateSwapOrderPriceWithTriggerAboveThreshold is a free data retrieval call binding the contract method 0xc4a1821b.
//
// Solidity: function validateSwapOrderPriceWithTriggerAboveThreshold(address[] _path, uint256 _triggerRatio) view returns(bool)
func (_OrderBook *OrderBookCallerSession) ValidateSwapOrderPriceWithTriggerAboveThreshold(_path []common.Address, _triggerRatio *big.Int) (bool, error) {
	return _OrderBook.Contract.ValidateSwapOrderPriceWithTriggerAboveThreshold(&_OrderBook.CallOpts, _path, _triggerRatio)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_OrderBook *OrderBookCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_OrderBook *OrderBookSession) Vault() (common.Address, error) {
	return _OrderBook.Contract.Vault(&_OrderBook.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_OrderBook *OrderBookCallerSession) Vault() (common.Address, error) {
	return _OrderBook.Contract.Vault(&_OrderBook.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_OrderBook *OrderBookCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_OrderBook *OrderBookSession) Weth() (common.Address, error) {
	return _OrderBook.Contract.Weth(&_OrderBook.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_OrderBook *OrderBookCallerSession) Weth() (common.Address, error) {
	return _OrderBook.Contract.Weth(&_OrderBook.CallOpts)
}

// CancelDecreaseOrder is a paid mutator transaction binding the contract method 0x9e71b0f0.
//
// Solidity: function cancelDecreaseOrder(uint256 _orderIndex) returns()
func (_OrderBook *OrderBookTransactor) CancelDecreaseOrder(opts *bind.TransactOpts, _orderIndex *big.Int) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "cancelDecreaseOrder", _orderIndex)
}

// CancelDecreaseOrder is a paid mutator transaction binding the contract method 0x9e71b0f0.
//
// Solidity: function cancelDecreaseOrder(uint256 _orderIndex) returns()
func (_OrderBook *OrderBookSession) CancelDecreaseOrder(_orderIndex *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.CancelDecreaseOrder(&_OrderBook.TransactOpts, _orderIndex)
}

// CancelDecreaseOrder is a paid mutator transaction binding the contract method 0x9e71b0f0.
//
// Solidity: function cancelDecreaseOrder(uint256 _orderIndex) returns()
func (_OrderBook *OrderBookTransactorSession) CancelDecreaseOrder(_orderIndex *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.CancelDecreaseOrder(&_OrderBook.TransactOpts, _orderIndex)
}

// CancelIncreaseOrder is a paid mutator transaction binding the contract method 0x47e0bbd0.
//
// Solidity: function cancelIncreaseOrder(uint256 _orderIndex) returns()
func (_OrderBook *OrderBookTransactor) CancelIncreaseOrder(opts *bind.TransactOpts, _orderIndex *big.Int) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "cancelIncreaseOrder", _orderIndex)
}

// CancelIncreaseOrder is a paid mutator transaction binding the contract method 0x47e0bbd0.
//
// Solidity: function cancelIncreaseOrder(uint256 _orderIndex) returns()
func (_OrderBook *OrderBookSession) CancelIncreaseOrder(_orderIndex *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.CancelIncreaseOrder(&_OrderBook.TransactOpts, _orderIndex)
}

// CancelIncreaseOrder is a paid mutator transaction binding the contract method 0x47e0bbd0.
//
// Solidity: function cancelIncreaseOrder(uint256 _orderIndex) returns()
func (_OrderBook *OrderBookTransactorSession) CancelIncreaseOrder(_orderIndex *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.CancelIncreaseOrder(&_OrderBook.TransactOpts, _orderIndex)
}

// CancelMultiple is a paid mutator transaction binding the contract method 0x807c5600.
//
// Solidity: function cancelMultiple(uint256[] _swapOrderIndexes, uint256[] _increaseOrderIndexes, uint256[] _decreaseOrderIndexes) returns()
func (_OrderBook *OrderBookTransactor) CancelMultiple(opts *bind.TransactOpts, _swapOrderIndexes []*big.Int, _increaseOrderIndexes []*big.Int, _decreaseOrderIndexes []*big.Int) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "cancelMultiple", _swapOrderIndexes, _increaseOrderIndexes, _decreaseOrderIndexes)
}

// CancelMultiple is a paid mutator transaction binding the contract method 0x807c5600.
//
// Solidity: function cancelMultiple(uint256[] _swapOrderIndexes, uint256[] _increaseOrderIndexes, uint256[] _decreaseOrderIndexes) returns()
func (_OrderBook *OrderBookSession) CancelMultiple(_swapOrderIndexes []*big.Int, _increaseOrderIndexes []*big.Int, _decreaseOrderIndexes []*big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.CancelMultiple(&_OrderBook.TransactOpts, _swapOrderIndexes, _increaseOrderIndexes, _decreaseOrderIndexes)
}

// CancelMultiple is a paid mutator transaction binding the contract method 0x807c5600.
//
// Solidity: function cancelMultiple(uint256[] _swapOrderIndexes, uint256[] _increaseOrderIndexes, uint256[] _decreaseOrderIndexes) returns()
func (_OrderBook *OrderBookTransactorSession) CancelMultiple(_swapOrderIndexes []*big.Int, _increaseOrderIndexes []*big.Int, _decreaseOrderIndexes []*big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.CancelMultiple(&_OrderBook.TransactOpts, _swapOrderIndexes, _increaseOrderIndexes, _decreaseOrderIndexes)
}

// CancelSwapOrder is a paid mutator transaction binding the contract method 0xf882ac07.
//
// Solidity: function cancelSwapOrder(uint256 _orderIndex) returns()
func (_OrderBook *OrderBookTransactor) CancelSwapOrder(opts *bind.TransactOpts, _orderIndex *big.Int) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "cancelSwapOrder", _orderIndex)
}

// CancelSwapOrder is a paid mutator transaction binding the contract method 0xf882ac07.
//
// Solidity: function cancelSwapOrder(uint256 _orderIndex) returns()
func (_OrderBook *OrderBookSession) CancelSwapOrder(_orderIndex *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.CancelSwapOrder(&_OrderBook.TransactOpts, _orderIndex)
}

// CancelSwapOrder is a paid mutator transaction binding the contract method 0xf882ac07.
//
// Solidity: function cancelSwapOrder(uint256 _orderIndex) returns()
func (_OrderBook *OrderBookTransactorSession) CancelSwapOrder(_orderIndex *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.CancelSwapOrder(&_OrderBook.TransactOpts, _orderIndex)
}

// CreateDecreaseOrder is a paid mutator transaction binding the contract method 0xc16cde8a.
//
// Solidity: function createDecreaseOrder(address _indexToken, uint256 _sizeDelta, address _collateralToken, uint256 _collateralDelta, bool _isLong, uint256 _triggerPrice, bool _triggerAboveThreshold) payable returns()
func (_OrderBook *OrderBookTransactor) CreateDecreaseOrder(opts *bind.TransactOpts, _indexToken common.Address, _sizeDelta *big.Int, _collateralToken common.Address, _collateralDelta *big.Int, _isLong bool, _triggerPrice *big.Int, _triggerAboveThreshold bool) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "createDecreaseOrder", _indexToken, _sizeDelta, _collateralToken, _collateralDelta, _isLong, _triggerPrice, _triggerAboveThreshold)
}

// CreateDecreaseOrder is a paid mutator transaction binding the contract method 0xc16cde8a.
//
// Solidity: function createDecreaseOrder(address _indexToken, uint256 _sizeDelta, address _collateralToken, uint256 _collateralDelta, bool _isLong, uint256 _triggerPrice, bool _triggerAboveThreshold) payable returns()
func (_OrderBook *OrderBookSession) CreateDecreaseOrder(_indexToken common.Address, _sizeDelta *big.Int, _collateralToken common.Address, _collateralDelta *big.Int, _isLong bool, _triggerPrice *big.Int, _triggerAboveThreshold bool) (*types.Transaction, error) {
	return _OrderBook.Contract.CreateDecreaseOrder(&_OrderBook.TransactOpts, _indexToken, _sizeDelta, _collateralToken, _collateralDelta, _isLong, _triggerPrice, _triggerAboveThreshold)
}

// CreateDecreaseOrder is a paid mutator transaction binding the contract method 0xc16cde8a.
//
// Solidity: function createDecreaseOrder(address _indexToken, uint256 _sizeDelta, address _collateralToken, uint256 _collateralDelta, bool _isLong, uint256 _triggerPrice, bool _triggerAboveThreshold) payable returns()
func (_OrderBook *OrderBookTransactorSession) CreateDecreaseOrder(_indexToken common.Address, _sizeDelta *big.Int, _collateralToken common.Address, _collateralDelta *big.Int, _isLong bool, _triggerPrice *big.Int, _triggerAboveThreshold bool) (*types.Transaction, error) {
	return _OrderBook.Contract.CreateDecreaseOrder(&_OrderBook.TransactOpts, _indexToken, _sizeDelta, _collateralToken, _collateralDelta, _isLong, _triggerPrice, _triggerAboveThreshold)
}

// CreateIncreaseOrder is a paid mutator transaction binding the contract method 0xb142a4b0.
//
// Solidity: function createIncreaseOrder(address[] _path, uint256 _amountIn, address _indexToken, uint256 _minOut, uint256 _sizeDelta, address _collateralToken, bool _isLong, uint256 _triggerPrice, bool _triggerAboveThreshold, uint256 _executionFee, bool _shouldWrap) payable returns()
func (_OrderBook *OrderBookTransactor) CreateIncreaseOrder(opts *bind.TransactOpts, _path []common.Address, _amountIn *big.Int, _indexToken common.Address, _minOut *big.Int, _sizeDelta *big.Int, _collateralToken common.Address, _isLong bool, _triggerPrice *big.Int, _triggerAboveThreshold bool, _executionFee *big.Int, _shouldWrap bool) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "createIncreaseOrder", _path, _amountIn, _indexToken, _minOut, _sizeDelta, _collateralToken, _isLong, _triggerPrice, _triggerAboveThreshold, _executionFee, _shouldWrap)
}

// CreateIncreaseOrder is a paid mutator transaction binding the contract method 0xb142a4b0.
//
// Solidity: function createIncreaseOrder(address[] _path, uint256 _amountIn, address _indexToken, uint256 _minOut, uint256 _sizeDelta, address _collateralToken, bool _isLong, uint256 _triggerPrice, bool _triggerAboveThreshold, uint256 _executionFee, bool _shouldWrap) payable returns()
func (_OrderBook *OrderBookSession) CreateIncreaseOrder(_path []common.Address, _amountIn *big.Int, _indexToken common.Address, _minOut *big.Int, _sizeDelta *big.Int, _collateralToken common.Address, _isLong bool, _triggerPrice *big.Int, _triggerAboveThreshold bool, _executionFee *big.Int, _shouldWrap bool) (*types.Transaction, error) {
	return _OrderBook.Contract.CreateIncreaseOrder(&_OrderBook.TransactOpts, _path, _amountIn, _indexToken, _minOut, _sizeDelta, _collateralToken, _isLong, _triggerPrice, _triggerAboveThreshold, _executionFee, _shouldWrap)
}

// CreateIncreaseOrder is a paid mutator transaction binding the contract method 0xb142a4b0.
//
// Solidity: function createIncreaseOrder(address[] _path, uint256 _amountIn, address _indexToken, uint256 _minOut, uint256 _sizeDelta, address _collateralToken, bool _isLong, uint256 _triggerPrice, bool _triggerAboveThreshold, uint256 _executionFee, bool _shouldWrap) payable returns()
func (_OrderBook *OrderBookTransactorSession) CreateIncreaseOrder(_path []common.Address, _amountIn *big.Int, _indexToken common.Address, _minOut *big.Int, _sizeDelta *big.Int, _collateralToken common.Address, _isLong bool, _triggerPrice *big.Int, _triggerAboveThreshold bool, _executionFee *big.Int, _shouldWrap bool) (*types.Transaction, error) {
	return _OrderBook.Contract.CreateIncreaseOrder(&_OrderBook.TransactOpts, _path, _amountIn, _indexToken, _minOut, _sizeDelta, _collateralToken, _isLong, _triggerPrice, _triggerAboveThreshold, _executionFee, _shouldWrap)
}

// CreateSwapOrder is a paid mutator transaction binding the contract method 0x269ae6c2.
//
// Solidity: function createSwapOrder(address[] _path, uint256 _amountIn, uint256 _minOut, uint256 _triggerRatio, bool _triggerAboveThreshold, uint256 _executionFee, bool _shouldWrap, bool _shouldUnwrap) payable returns()
func (_OrderBook *OrderBookTransactor) CreateSwapOrder(opts *bind.TransactOpts, _path []common.Address, _amountIn *big.Int, _minOut *big.Int, _triggerRatio *big.Int, _triggerAboveThreshold bool, _executionFee *big.Int, _shouldWrap bool, _shouldUnwrap bool) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "createSwapOrder", _path, _amountIn, _minOut, _triggerRatio, _triggerAboveThreshold, _executionFee, _shouldWrap, _shouldUnwrap)
}

// CreateSwapOrder is a paid mutator transaction binding the contract method 0x269ae6c2.
//
// Solidity: function createSwapOrder(address[] _path, uint256 _amountIn, uint256 _minOut, uint256 _triggerRatio, bool _triggerAboveThreshold, uint256 _executionFee, bool _shouldWrap, bool _shouldUnwrap) payable returns()
func (_OrderBook *OrderBookSession) CreateSwapOrder(_path []common.Address, _amountIn *big.Int, _minOut *big.Int, _triggerRatio *big.Int, _triggerAboveThreshold bool, _executionFee *big.Int, _shouldWrap bool, _shouldUnwrap bool) (*types.Transaction, error) {
	return _OrderBook.Contract.CreateSwapOrder(&_OrderBook.TransactOpts, _path, _amountIn, _minOut, _triggerRatio, _triggerAboveThreshold, _executionFee, _shouldWrap, _shouldUnwrap)
}

// CreateSwapOrder is a paid mutator transaction binding the contract method 0x269ae6c2.
//
// Solidity: function createSwapOrder(address[] _path, uint256 _amountIn, uint256 _minOut, uint256 _triggerRatio, bool _triggerAboveThreshold, uint256 _executionFee, bool _shouldWrap, bool _shouldUnwrap) payable returns()
func (_OrderBook *OrderBookTransactorSession) CreateSwapOrder(_path []common.Address, _amountIn *big.Int, _minOut *big.Int, _triggerRatio *big.Int, _triggerAboveThreshold bool, _executionFee *big.Int, _shouldWrap bool, _shouldUnwrap bool) (*types.Transaction, error) {
	return _OrderBook.Contract.CreateSwapOrder(&_OrderBook.TransactOpts, _path, _amountIn, _minOut, _triggerRatio, _triggerAboveThreshold, _executionFee, _shouldWrap, _shouldUnwrap)
}

// ExecuteDecreaseOrder is a paid mutator transaction binding the contract method 0x11d9444a.
//
// Solidity: function executeDecreaseOrder(address _address, uint256 _orderIndex, address _feeReceiver) returns()
func (_OrderBook *OrderBookTransactor) ExecuteDecreaseOrder(opts *bind.TransactOpts, _address common.Address, _orderIndex *big.Int, _feeReceiver common.Address) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "executeDecreaseOrder", _address, _orderIndex, _feeReceiver)
}

// ExecuteDecreaseOrder is a paid mutator transaction binding the contract method 0x11d9444a.
//
// Solidity: function executeDecreaseOrder(address _address, uint256 _orderIndex, address _feeReceiver) returns()
func (_OrderBook *OrderBookSession) ExecuteDecreaseOrder(_address common.Address, _orderIndex *big.Int, _feeReceiver common.Address) (*types.Transaction, error) {
	return _OrderBook.Contract.ExecuteDecreaseOrder(&_OrderBook.TransactOpts, _address, _orderIndex, _feeReceiver)
}

// ExecuteDecreaseOrder is a paid mutator transaction binding the contract method 0x11d9444a.
//
// Solidity: function executeDecreaseOrder(address _address, uint256 _orderIndex, address _feeReceiver) returns()
func (_OrderBook *OrderBookTransactorSession) ExecuteDecreaseOrder(_address common.Address, _orderIndex *big.Int, _feeReceiver common.Address) (*types.Transaction, error) {
	return _OrderBook.Contract.ExecuteDecreaseOrder(&_OrderBook.TransactOpts, _address, _orderIndex, _feeReceiver)
}

// ExecuteIncreaseOrder is a paid mutator transaction binding the contract method 0xd38ab519.
//
// Solidity: function executeIncreaseOrder(address _address, uint256 _orderIndex, address _feeReceiver) returns()
func (_OrderBook *OrderBookTransactor) ExecuteIncreaseOrder(opts *bind.TransactOpts, _address common.Address, _orderIndex *big.Int, _feeReceiver common.Address) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "executeIncreaseOrder", _address, _orderIndex, _feeReceiver)
}

// ExecuteIncreaseOrder is a paid mutator transaction binding the contract method 0xd38ab519.
//
// Solidity: function executeIncreaseOrder(address _address, uint256 _orderIndex, address _feeReceiver) returns()
func (_OrderBook *OrderBookSession) ExecuteIncreaseOrder(_address common.Address, _orderIndex *big.Int, _feeReceiver common.Address) (*types.Transaction, error) {
	return _OrderBook.Contract.ExecuteIncreaseOrder(&_OrderBook.TransactOpts, _address, _orderIndex, _feeReceiver)
}

// ExecuteIncreaseOrder is a paid mutator transaction binding the contract method 0xd38ab519.
//
// Solidity: function executeIncreaseOrder(address _address, uint256 _orderIndex, address _feeReceiver) returns()
func (_OrderBook *OrderBookTransactorSession) ExecuteIncreaseOrder(_address common.Address, _orderIndex *big.Int, _feeReceiver common.Address) (*types.Transaction, error) {
	return _OrderBook.Contract.ExecuteIncreaseOrder(&_OrderBook.TransactOpts, _address, _orderIndex, _feeReceiver)
}

// ExecuteSwapOrder is a paid mutator transaction binding the contract method 0x07c7edc3.
//
// Solidity: function executeSwapOrder(address _account, uint256 _orderIndex, address _feeReceiver) returns()
func (_OrderBook *OrderBookTransactor) ExecuteSwapOrder(opts *bind.TransactOpts, _account common.Address, _orderIndex *big.Int, _feeReceiver common.Address) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "executeSwapOrder", _account, _orderIndex, _feeReceiver)
}

// ExecuteSwapOrder is a paid mutator transaction binding the contract method 0x07c7edc3.
//
// Solidity: function executeSwapOrder(address _account, uint256 _orderIndex, address _feeReceiver) returns()
func (_OrderBook *OrderBookSession) ExecuteSwapOrder(_account common.Address, _orderIndex *big.Int, _feeReceiver common.Address) (*types.Transaction, error) {
	return _OrderBook.Contract.ExecuteSwapOrder(&_OrderBook.TransactOpts, _account, _orderIndex, _feeReceiver)
}

// ExecuteSwapOrder is a paid mutator transaction binding the contract method 0x07c7edc3.
//
// Solidity: function executeSwapOrder(address _account, uint256 _orderIndex, address _feeReceiver) returns()
func (_OrderBook *OrderBookTransactorSession) ExecuteSwapOrder(_account common.Address, _orderIndex *big.Int, _feeReceiver common.Address) (*types.Transaction, error) {
	return _OrderBook.Contract.ExecuteSwapOrder(&_OrderBook.TransactOpts, _account, _orderIndex, _feeReceiver)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7c41c79.
//
// Solidity: function initialize(address _router, address _vault, address _weth, address _usdg, uint256 _minExecutionFee, uint256 _minPurchaseTokenAmountUsd) returns()
func (_OrderBook *OrderBookTransactor) Initialize(opts *bind.TransactOpts, _router common.Address, _vault common.Address, _weth common.Address, _usdg common.Address, _minExecutionFee *big.Int, _minPurchaseTokenAmountUsd *big.Int) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "initialize", _router, _vault, _weth, _usdg, _minExecutionFee, _minPurchaseTokenAmountUsd)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7c41c79.
//
// Solidity: function initialize(address _router, address _vault, address _weth, address _usdg, uint256 _minExecutionFee, uint256 _minPurchaseTokenAmountUsd) returns()
func (_OrderBook *OrderBookSession) Initialize(_router common.Address, _vault common.Address, _weth common.Address, _usdg common.Address, _minExecutionFee *big.Int, _minPurchaseTokenAmountUsd *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.Initialize(&_OrderBook.TransactOpts, _router, _vault, _weth, _usdg, _minExecutionFee, _minPurchaseTokenAmountUsd)
}

// Initialize is a paid mutator transaction binding the contract method 0xd7c41c79.
//
// Solidity: function initialize(address _router, address _vault, address _weth, address _usdg, uint256 _minExecutionFee, uint256 _minPurchaseTokenAmountUsd) returns()
func (_OrderBook *OrderBookTransactorSession) Initialize(_router common.Address, _vault common.Address, _weth common.Address, _usdg common.Address, _minExecutionFee *big.Int, _minPurchaseTokenAmountUsd *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.Initialize(&_OrderBook.TransactOpts, _router, _vault, _weth, _usdg, _minExecutionFee, _minPurchaseTokenAmountUsd)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_OrderBook *OrderBookTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_OrderBook *OrderBookSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _OrderBook.Contract.SetGov(&_OrderBook.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_OrderBook *OrderBookTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _OrderBook.Contract.SetGov(&_OrderBook.TransactOpts, _gov)
}

// SetMinExecutionFee is a paid mutator transaction binding the contract method 0xfc2cee62.
//
// Solidity: function setMinExecutionFee(uint256 _minExecutionFee) returns()
func (_OrderBook *OrderBookTransactor) SetMinExecutionFee(opts *bind.TransactOpts, _minExecutionFee *big.Int) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "setMinExecutionFee", _minExecutionFee)
}

// SetMinExecutionFee is a paid mutator transaction binding the contract method 0xfc2cee62.
//
// Solidity: function setMinExecutionFee(uint256 _minExecutionFee) returns()
func (_OrderBook *OrderBookSession) SetMinExecutionFee(_minExecutionFee *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.SetMinExecutionFee(&_OrderBook.TransactOpts, _minExecutionFee)
}

// SetMinExecutionFee is a paid mutator transaction binding the contract method 0xfc2cee62.
//
// Solidity: function setMinExecutionFee(uint256 _minExecutionFee) returns()
func (_OrderBook *OrderBookTransactorSession) SetMinExecutionFee(_minExecutionFee *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.SetMinExecutionFee(&_OrderBook.TransactOpts, _minExecutionFee)
}

// SetMinPurchaseTokenAmountUsd is a paid mutator transaction binding the contract method 0x0d5cc938.
//
// Solidity: function setMinPurchaseTokenAmountUsd(uint256 _minPurchaseTokenAmountUsd) returns()
func (_OrderBook *OrderBookTransactor) SetMinPurchaseTokenAmountUsd(opts *bind.TransactOpts, _minPurchaseTokenAmountUsd *big.Int) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "setMinPurchaseTokenAmountUsd", _minPurchaseTokenAmountUsd)
}

// SetMinPurchaseTokenAmountUsd is a paid mutator transaction binding the contract method 0x0d5cc938.
//
// Solidity: function setMinPurchaseTokenAmountUsd(uint256 _minPurchaseTokenAmountUsd) returns()
func (_OrderBook *OrderBookSession) SetMinPurchaseTokenAmountUsd(_minPurchaseTokenAmountUsd *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.SetMinPurchaseTokenAmountUsd(&_OrderBook.TransactOpts, _minPurchaseTokenAmountUsd)
}

// SetMinPurchaseTokenAmountUsd is a paid mutator transaction binding the contract method 0x0d5cc938.
//
// Solidity: function setMinPurchaseTokenAmountUsd(uint256 _minPurchaseTokenAmountUsd) returns()
func (_OrderBook *OrderBookTransactorSession) SetMinPurchaseTokenAmountUsd(_minPurchaseTokenAmountUsd *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.SetMinPurchaseTokenAmountUsd(&_OrderBook.TransactOpts, _minPurchaseTokenAmountUsd)
}

// UpdateDecreaseOrder is a paid mutator transaction binding the contract method 0xa397ea54.
//
// Solidity: function updateDecreaseOrder(uint256 _orderIndex, uint256 _collateralDelta, uint256 _sizeDelta, uint256 _triggerPrice, bool _triggerAboveThreshold) returns()
func (_OrderBook *OrderBookTransactor) UpdateDecreaseOrder(opts *bind.TransactOpts, _orderIndex *big.Int, _collateralDelta *big.Int, _sizeDelta *big.Int, _triggerPrice *big.Int, _triggerAboveThreshold bool) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "updateDecreaseOrder", _orderIndex, _collateralDelta, _sizeDelta, _triggerPrice, _triggerAboveThreshold)
}

// UpdateDecreaseOrder is a paid mutator transaction binding the contract method 0xa397ea54.
//
// Solidity: function updateDecreaseOrder(uint256 _orderIndex, uint256 _collateralDelta, uint256 _sizeDelta, uint256 _triggerPrice, bool _triggerAboveThreshold) returns()
func (_OrderBook *OrderBookSession) UpdateDecreaseOrder(_orderIndex *big.Int, _collateralDelta *big.Int, _sizeDelta *big.Int, _triggerPrice *big.Int, _triggerAboveThreshold bool) (*types.Transaction, error) {
	return _OrderBook.Contract.UpdateDecreaseOrder(&_OrderBook.TransactOpts, _orderIndex, _collateralDelta, _sizeDelta, _triggerPrice, _triggerAboveThreshold)
}

// UpdateDecreaseOrder is a paid mutator transaction binding the contract method 0xa397ea54.
//
// Solidity: function updateDecreaseOrder(uint256 _orderIndex, uint256 _collateralDelta, uint256 _sizeDelta, uint256 _triggerPrice, bool _triggerAboveThreshold) returns()
func (_OrderBook *OrderBookTransactorSession) UpdateDecreaseOrder(_orderIndex *big.Int, _collateralDelta *big.Int, _sizeDelta *big.Int, _triggerPrice *big.Int, _triggerAboveThreshold bool) (*types.Transaction, error) {
	return _OrderBook.Contract.UpdateDecreaseOrder(&_OrderBook.TransactOpts, _orderIndex, _collateralDelta, _sizeDelta, _triggerPrice, _triggerAboveThreshold)
}

// UpdateIncreaseOrder is a paid mutator transaction binding the contract method 0x9983ee1b.
//
// Solidity: function updateIncreaseOrder(uint256 _orderIndex, uint256 _sizeDelta, uint256 _triggerPrice, bool _triggerAboveThreshold) returns()
func (_OrderBook *OrderBookTransactor) UpdateIncreaseOrder(opts *bind.TransactOpts, _orderIndex *big.Int, _sizeDelta *big.Int, _triggerPrice *big.Int, _triggerAboveThreshold bool) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "updateIncreaseOrder", _orderIndex, _sizeDelta, _triggerPrice, _triggerAboveThreshold)
}

// UpdateIncreaseOrder is a paid mutator transaction binding the contract method 0x9983ee1b.
//
// Solidity: function updateIncreaseOrder(uint256 _orderIndex, uint256 _sizeDelta, uint256 _triggerPrice, bool _triggerAboveThreshold) returns()
func (_OrderBook *OrderBookSession) UpdateIncreaseOrder(_orderIndex *big.Int, _sizeDelta *big.Int, _triggerPrice *big.Int, _triggerAboveThreshold bool) (*types.Transaction, error) {
	return _OrderBook.Contract.UpdateIncreaseOrder(&_OrderBook.TransactOpts, _orderIndex, _sizeDelta, _triggerPrice, _triggerAboveThreshold)
}

// UpdateIncreaseOrder is a paid mutator transaction binding the contract method 0x9983ee1b.
//
// Solidity: function updateIncreaseOrder(uint256 _orderIndex, uint256 _sizeDelta, uint256 _triggerPrice, bool _triggerAboveThreshold) returns()
func (_OrderBook *OrderBookTransactorSession) UpdateIncreaseOrder(_orderIndex *big.Int, _sizeDelta *big.Int, _triggerPrice *big.Int, _triggerAboveThreshold bool) (*types.Transaction, error) {
	return _OrderBook.Contract.UpdateIncreaseOrder(&_OrderBook.TransactOpts, _orderIndex, _sizeDelta, _triggerPrice, _triggerAboveThreshold)
}

// UpdateSwapOrder is a paid mutator transaction binding the contract method 0xc86b0f7d.
//
// Solidity: function updateSwapOrder(uint256 _orderIndex, uint256 _minOut, uint256 _triggerRatio, bool _triggerAboveThreshold) returns()
func (_OrderBook *OrderBookTransactor) UpdateSwapOrder(opts *bind.TransactOpts, _orderIndex *big.Int, _minOut *big.Int, _triggerRatio *big.Int, _triggerAboveThreshold bool) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "updateSwapOrder", _orderIndex, _minOut, _triggerRatio, _triggerAboveThreshold)
}

// UpdateSwapOrder is a paid mutator transaction binding the contract method 0xc86b0f7d.
//
// Solidity: function updateSwapOrder(uint256 _orderIndex, uint256 _minOut, uint256 _triggerRatio, bool _triggerAboveThreshold) returns()
func (_OrderBook *OrderBookSession) UpdateSwapOrder(_orderIndex *big.Int, _minOut *big.Int, _triggerRatio *big.Int, _triggerAboveThreshold bool) (*types.Transaction, error) {
	return _OrderBook.Contract.UpdateSwapOrder(&_OrderBook.TransactOpts, _orderIndex, _minOut, _triggerRatio, _triggerAboveThreshold)
}

// UpdateSwapOrder is a paid mutator transaction binding the contract method 0xc86b0f7d.
//
// Solidity: function updateSwapOrder(uint256 _orderIndex, uint256 _minOut, uint256 _triggerRatio, bool _triggerAboveThreshold) returns()
func (_OrderBook *OrderBookTransactorSession) UpdateSwapOrder(_orderIndex *big.Int, _minOut *big.Int, _triggerRatio *big.Int, _triggerAboveThreshold bool) (*types.Transaction, error) {
	return _OrderBook.Contract.UpdateSwapOrder(&_OrderBook.TransactOpts, _orderIndex, _minOut, _triggerRatio, _triggerAboveThreshold)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OrderBook *OrderBookTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBook.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OrderBook *OrderBookSession) Receive() (*types.Transaction, error) {
	return _OrderBook.Contract.Receive(&_OrderBook.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_OrderBook *OrderBookTransactorSession) Receive() (*types.Transaction, error) {
	return _OrderBook.Contract.Receive(&_OrderBook.TransactOpts)
}

// OrderBookCancelDecreaseOrderIterator is returned from FilterCancelDecreaseOrder and is used to iterate over the raw logs and unpacked data for CancelDecreaseOrder events raised by the OrderBook contract.
type OrderBookCancelDecreaseOrderIterator struct {
	Event *OrderBookCancelDecreaseOrder // Event containing the contract specifics and raw log

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
func (it *OrderBookCancelDecreaseOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookCancelDecreaseOrder)
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
		it.Event = new(OrderBookCancelDecreaseOrder)
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
func (it *OrderBookCancelDecreaseOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookCancelDecreaseOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookCancelDecreaseOrder represents a CancelDecreaseOrder event raised by the OrderBook contract.
type OrderBookCancelDecreaseOrder struct {
	Account               common.Address
	OrderIndex            *big.Int
	CollateralToken       common.Address
	CollateralDelta       *big.Int
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterCancelDecreaseOrder is a free log retrieval operation binding the contract event 0x1154174c82984656b028c8021671988f60a346497e56fe02554761184f82a075.
//
// Solidity: event CancelDecreaseOrder(address indexed account, uint256 orderIndex, address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) FilterCancelDecreaseOrder(opts *bind.FilterOpts, account []common.Address) (*OrderBookCancelDecreaseOrderIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "CancelDecreaseOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return &OrderBookCancelDecreaseOrderIterator{contract: _OrderBook.contract, event: "CancelDecreaseOrder", logs: logs, sub: sub}, nil
}

// WatchCancelDecreaseOrder is a free log subscription operation binding the contract event 0x1154174c82984656b028c8021671988f60a346497e56fe02554761184f82a075.
//
// Solidity: event CancelDecreaseOrder(address indexed account, uint256 orderIndex, address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) WatchCancelDecreaseOrder(opts *bind.WatchOpts, sink chan<- *OrderBookCancelDecreaseOrder, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "CancelDecreaseOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookCancelDecreaseOrder)
				if err := _OrderBook.contract.UnpackLog(event, "CancelDecreaseOrder", log); err != nil {
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

// ParseCancelDecreaseOrder is a log parse operation binding the contract event 0x1154174c82984656b028c8021671988f60a346497e56fe02554761184f82a075.
//
// Solidity: event CancelDecreaseOrder(address indexed account, uint256 orderIndex, address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) ParseCancelDecreaseOrder(log types.Log) (*OrderBookCancelDecreaseOrder, error) {
	event := new(OrderBookCancelDecreaseOrder)
	if err := _OrderBook.contract.UnpackLog(event, "CancelDecreaseOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookCancelIncreaseOrderIterator is returned from FilterCancelIncreaseOrder and is used to iterate over the raw logs and unpacked data for CancelIncreaseOrder events raised by the OrderBook contract.
type OrderBookCancelIncreaseOrderIterator struct {
	Event *OrderBookCancelIncreaseOrder // Event containing the contract specifics and raw log

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
func (it *OrderBookCancelIncreaseOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookCancelIncreaseOrder)
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
		it.Event = new(OrderBookCancelIncreaseOrder)
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
func (it *OrderBookCancelIncreaseOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookCancelIncreaseOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookCancelIncreaseOrder represents a CancelIncreaseOrder event raised by the OrderBook contract.
type OrderBookCancelIncreaseOrder struct {
	Account               common.Address
	OrderIndex            *big.Int
	PurchaseToken         common.Address
	PurchaseTokenAmount   *big.Int
	CollateralToken       common.Address
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterCancelIncreaseOrder is a free log retrieval operation binding the contract event 0xd500f34e0ec655b7614ae42e1d9c666d5e4dde909a1297829f8c5ecf00805d32.
//
// Solidity: event CancelIncreaseOrder(address indexed account, uint256 orderIndex, address purchaseToken, uint256 purchaseTokenAmount, address collateralToken, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) FilterCancelIncreaseOrder(opts *bind.FilterOpts, account []common.Address) (*OrderBookCancelIncreaseOrderIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "CancelIncreaseOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return &OrderBookCancelIncreaseOrderIterator{contract: _OrderBook.contract, event: "CancelIncreaseOrder", logs: logs, sub: sub}, nil
}

// WatchCancelIncreaseOrder is a free log subscription operation binding the contract event 0xd500f34e0ec655b7614ae42e1d9c666d5e4dde909a1297829f8c5ecf00805d32.
//
// Solidity: event CancelIncreaseOrder(address indexed account, uint256 orderIndex, address purchaseToken, uint256 purchaseTokenAmount, address collateralToken, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) WatchCancelIncreaseOrder(opts *bind.WatchOpts, sink chan<- *OrderBookCancelIncreaseOrder, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "CancelIncreaseOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookCancelIncreaseOrder)
				if err := _OrderBook.contract.UnpackLog(event, "CancelIncreaseOrder", log); err != nil {
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

// ParseCancelIncreaseOrder is a log parse operation binding the contract event 0xd500f34e0ec655b7614ae42e1d9c666d5e4dde909a1297829f8c5ecf00805d32.
//
// Solidity: event CancelIncreaseOrder(address indexed account, uint256 orderIndex, address purchaseToken, uint256 purchaseTokenAmount, address collateralToken, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) ParseCancelIncreaseOrder(log types.Log) (*OrderBookCancelIncreaseOrder, error) {
	event := new(OrderBookCancelIncreaseOrder)
	if err := _OrderBook.contract.UnpackLog(event, "CancelIncreaseOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookCancelSwapOrderIterator is returned from FilterCancelSwapOrder and is used to iterate over the raw logs and unpacked data for CancelSwapOrder events raised by the OrderBook contract.
type OrderBookCancelSwapOrderIterator struct {
	Event *OrderBookCancelSwapOrder // Event containing the contract specifics and raw log

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
func (it *OrderBookCancelSwapOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookCancelSwapOrder)
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
		it.Event = new(OrderBookCancelSwapOrder)
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
func (it *OrderBookCancelSwapOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookCancelSwapOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookCancelSwapOrder represents a CancelSwapOrder event raised by the OrderBook contract.
type OrderBookCancelSwapOrder struct {
	Account               common.Address
	OrderIndex            *big.Int
	Path                  []common.Address
	AmountIn              *big.Int
	MinOut                *big.Int
	TriggerRatio          *big.Int
	TriggerAboveThreshold bool
	ShouldUnwrap          bool
	ExecutionFee          *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterCancelSwapOrder is a free log retrieval operation binding the contract event 0xefd66d4f9c2f880c70aedeb5b26a44fb474cea07e5d6c533f2d27c303d5d9453.
//
// Solidity: event CancelSwapOrder(address indexed account, uint256 orderIndex, address[] path, uint256 amountIn, uint256 minOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) FilterCancelSwapOrder(opts *bind.FilterOpts, account []common.Address) (*OrderBookCancelSwapOrderIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "CancelSwapOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return &OrderBookCancelSwapOrderIterator{contract: _OrderBook.contract, event: "CancelSwapOrder", logs: logs, sub: sub}, nil
}

// WatchCancelSwapOrder is a free log subscription operation binding the contract event 0xefd66d4f9c2f880c70aedeb5b26a44fb474cea07e5d6c533f2d27c303d5d9453.
//
// Solidity: event CancelSwapOrder(address indexed account, uint256 orderIndex, address[] path, uint256 amountIn, uint256 minOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) WatchCancelSwapOrder(opts *bind.WatchOpts, sink chan<- *OrderBookCancelSwapOrder, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "CancelSwapOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookCancelSwapOrder)
				if err := _OrderBook.contract.UnpackLog(event, "CancelSwapOrder", log); err != nil {
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

// ParseCancelSwapOrder is a log parse operation binding the contract event 0xefd66d4f9c2f880c70aedeb5b26a44fb474cea07e5d6c533f2d27c303d5d9453.
//
// Solidity: event CancelSwapOrder(address indexed account, uint256 orderIndex, address[] path, uint256 amountIn, uint256 minOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) ParseCancelSwapOrder(log types.Log) (*OrderBookCancelSwapOrder, error) {
	event := new(OrderBookCancelSwapOrder)
	if err := _OrderBook.contract.UnpackLog(event, "CancelSwapOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookCreateDecreaseOrderIterator is returned from FilterCreateDecreaseOrder and is used to iterate over the raw logs and unpacked data for CreateDecreaseOrder events raised by the OrderBook contract.
type OrderBookCreateDecreaseOrderIterator struct {
	Event *OrderBookCreateDecreaseOrder // Event containing the contract specifics and raw log

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
func (it *OrderBookCreateDecreaseOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookCreateDecreaseOrder)
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
		it.Event = new(OrderBookCreateDecreaseOrder)
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
func (it *OrderBookCreateDecreaseOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookCreateDecreaseOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookCreateDecreaseOrder represents a CreateDecreaseOrder event raised by the OrderBook contract.
type OrderBookCreateDecreaseOrder struct {
	Account               common.Address
	OrderIndex            *big.Int
	CollateralToken       common.Address
	CollateralDelta       *big.Int
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterCreateDecreaseOrder is a free log retrieval operation binding the contract event 0x48ee333d2a65cc45fdb83bc012920d89181c3377390cd239d2b63f2bef67a02d.
//
// Solidity: event CreateDecreaseOrder(address indexed account, uint256 orderIndex, address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) FilterCreateDecreaseOrder(opts *bind.FilterOpts, account []common.Address) (*OrderBookCreateDecreaseOrderIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "CreateDecreaseOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return &OrderBookCreateDecreaseOrderIterator{contract: _OrderBook.contract, event: "CreateDecreaseOrder", logs: logs, sub: sub}, nil
}

// WatchCreateDecreaseOrder is a free log subscription operation binding the contract event 0x48ee333d2a65cc45fdb83bc012920d89181c3377390cd239d2b63f2bef67a02d.
//
// Solidity: event CreateDecreaseOrder(address indexed account, uint256 orderIndex, address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) WatchCreateDecreaseOrder(opts *bind.WatchOpts, sink chan<- *OrderBookCreateDecreaseOrder, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "CreateDecreaseOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookCreateDecreaseOrder)
				if err := _OrderBook.contract.UnpackLog(event, "CreateDecreaseOrder", log); err != nil {
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

// ParseCreateDecreaseOrder is a log parse operation binding the contract event 0x48ee333d2a65cc45fdb83bc012920d89181c3377390cd239d2b63f2bef67a02d.
//
// Solidity: event CreateDecreaseOrder(address indexed account, uint256 orderIndex, address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) ParseCreateDecreaseOrder(log types.Log) (*OrderBookCreateDecreaseOrder, error) {
	event := new(OrderBookCreateDecreaseOrder)
	if err := _OrderBook.contract.UnpackLog(event, "CreateDecreaseOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookCreateIncreaseOrderIterator is returned from FilterCreateIncreaseOrder and is used to iterate over the raw logs and unpacked data for CreateIncreaseOrder events raised by the OrderBook contract.
type OrderBookCreateIncreaseOrderIterator struct {
	Event *OrderBookCreateIncreaseOrder // Event containing the contract specifics and raw log

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
func (it *OrderBookCreateIncreaseOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookCreateIncreaseOrder)
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
		it.Event = new(OrderBookCreateIncreaseOrder)
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
func (it *OrderBookCreateIncreaseOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookCreateIncreaseOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookCreateIncreaseOrder represents a CreateIncreaseOrder event raised by the OrderBook contract.
type OrderBookCreateIncreaseOrder struct {
	Account               common.Address
	OrderIndex            *big.Int
	PurchaseToken         common.Address
	PurchaseTokenAmount   *big.Int
	CollateralToken       common.Address
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterCreateIncreaseOrder is a free log retrieval operation binding the contract event 0xb27b9afe3043b93788c40cfc3cc73f5d928a2e40f3ba01820b246426de8fa1b9.
//
// Solidity: event CreateIncreaseOrder(address indexed account, uint256 orderIndex, address purchaseToken, uint256 purchaseTokenAmount, address collateralToken, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) FilterCreateIncreaseOrder(opts *bind.FilterOpts, account []common.Address) (*OrderBookCreateIncreaseOrderIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "CreateIncreaseOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return &OrderBookCreateIncreaseOrderIterator{contract: _OrderBook.contract, event: "CreateIncreaseOrder", logs: logs, sub: sub}, nil
}

// WatchCreateIncreaseOrder is a free log subscription operation binding the contract event 0xb27b9afe3043b93788c40cfc3cc73f5d928a2e40f3ba01820b246426de8fa1b9.
//
// Solidity: event CreateIncreaseOrder(address indexed account, uint256 orderIndex, address purchaseToken, uint256 purchaseTokenAmount, address collateralToken, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) WatchCreateIncreaseOrder(opts *bind.WatchOpts, sink chan<- *OrderBookCreateIncreaseOrder, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "CreateIncreaseOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookCreateIncreaseOrder)
				if err := _OrderBook.contract.UnpackLog(event, "CreateIncreaseOrder", log); err != nil {
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

// ParseCreateIncreaseOrder is a log parse operation binding the contract event 0xb27b9afe3043b93788c40cfc3cc73f5d928a2e40f3ba01820b246426de8fa1b9.
//
// Solidity: event CreateIncreaseOrder(address indexed account, uint256 orderIndex, address purchaseToken, uint256 purchaseTokenAmount, address collateralToken, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) ParseCreateIncreaseOrder(log types.Log) (*OrderBookCreateIncreaseOrder, error) {
	event := new(OrderBookCreateIncreaseOrder)
	if err := _OrderBook.contract.UnpackLog(event, "CreateIncreaseOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookCreateSwapOrderIterator is returned from FilterCreateSwapOrder and is used to iterate over the raw logs and unpacked data for CreateSwapOrder events raised by the OrderBook contract.
type OrderBookCreateSwapOrderIterator struct {
	Event *OrderBookCreateSwapOrder // Event containing the contract specifics and raw log

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
func (it *OrderBookCreateSwapOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookCreateSwapOrder)
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
		it.Event = new(OrderBookCreateSwapOrder)
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
func (it *OrderBookCreateSwapOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookCreateSwapOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookCreateSwapOrder represents a CreateSwapOrder event raised by the OrderBook contract.
type OrderBookCreateSwapOrder struct {
	Account               common.Address
	OrderIndex            *big.Int
	Path                  []common.Address
	AmountIn              *big.Int
	MinOut                *big.Int
	TriggerRatio          *big.Int
	TriggerAboveThreshold bool
	ShouldUnwrap          bool
	ExecutionFee          *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterCreateSwapOrder is a free log retrieval operation binding the contract event 0xdf06bb56ffc4029dc0b62b68bb5bbadea93a38b530cefc9b81afb742a6555d88.
//
// Solidity: event CreateSwapOrder(address indexed account, uint256 orderIndex, address[] path, uint256 amountIn, uint256 minOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) FilterCreateSwapOrder(opts *bind.FilterOpts, account []common.Address) (*OrderBookCreateSwapOrderIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "CreateSwapOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return &OrderBookCreateSwapOrderIterator{contract: _OrderBook.contract, event: "CreateSwapOrder", logs: logs, sub: sub}, nil
}

// WatchCreateSwapOrder is a free log subscription operation binding the contract event 0xdf06bb56ffc4029dc0b62b68bb5bbadea93a38b530cefc9b81afb742a6555d88.
//
// Solidity: event CreateSwapOrder(address indexed account, uint256 orderIndex, address[] path, uint256 amountIn, uint256 minOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) WatchCreateSwapOrder(opts *bind.WatchOpts, sink chan<- *OrderBookCreateSwapOrder, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "CreateSwapOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookCreateSwapOrder)
				if err := _OrderBook.contract.UnpackLog(event, "CreateSwapOrder", log); err != nil {
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

// ParseCreateSwapOrder is a log parse operation binding the contract event 0xdf06bb56ffc4029dc0b62b68bb5bbadea93a38b530cefc9b81afb742a6555d88.
//
// Solidity: event CreateSwapOrder(address indexed account, uint256 orderIndex, address[] path, uint256 amountIn, uint256 minOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) ParseCreateSwapOrder(log types.Log) (*OrderBookCreateSwapOrder, error) {
	event := new(OrderBookCreateSwapOrder)
	if err := _OrderBook.contract.UnpackLog(event, "CreateSwapOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookExecuteDecreaseOrderIterator is returned from FilterExecuteDecreaseOrder and is used to iterate over the raw logs and unpacked data for ExecuteDecreaseOrder events raised by the OrderBook contract.
type OrderBookExecuteDecreaseOrderIterator struct {
	Event *OrderBookExecuteDecreaseOrder // Event containing the contract specifics and raw log

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
func (it *OrderBookExecuteDecreaseOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookExecuteDecreaseOrder)
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
		it.Event = new(OrderBookExecuteDecreaseOrder)
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
func (it *OrderBookExecuteDecreaseOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookExecuteDecreaseOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookExecuteDecreaseOrder represents a ExecuteDecreaseOrder event raised by the OrderBook contract.
type OrderBookExecuteDecreaseOrder struct {
	Account               common.Address
	OrderIndex            *big.Int
	CollateralToken       common.Address
	CollateralDelta       *big.Int
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
	ExecutionPrice        *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterExecuteDecreaseOrder is a free log retrieval operation binding the contract event 0x9a382661d6573da86db000471303be6f0b2b1bb66089b08e3c16a85d7b6e94f8.
//
// Solidity: event ExecuteDecreaseOrder(address indexed account, uint256 orderIndex, address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee, uint256 executionPrice)
func (_OrderBook *OrderBookFilterer) FilterExecuteDecreaseOrder(opts *bind.FilterOpts, account []common.Address) (*OrderBookExecuteDecreaseOrderIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "ExecuteDecreaseOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return &OrderBookExecuteDecreaseOrderIterator{contract: _OrderBook.contract, event: "ExecuteDecreaseOrder", logs: logs, sub: sub}, nil
}

// WatchExecuteDecreaseOrder is a free log subscription operation binding the contract event 0x9a382661d6573da86db000471303be6f0b2b1bb66089b08e3c16a85d7b6e94f8.
//
// Solidity: event ExecuteDecreaseOrder(address indexed account, uint256 orderIndex, address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee, uint256 executionPrice)
func (_OrderBook *OrderBookFilterer) WatchExecuteDecreaseOrder(opts *bind.WatchOpts, sink chan<- *OrderBookExecuteDecreaseOrder, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "ExecuteDecreaseOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookExecuteDecreaseOrder)
				if err := _OrderBook.contract.UnpackLog(event, "ExecuteDecreaseOrder", log); err != nil {
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

// ParseExecuteDecreaseOrder is a log parse operation binding the contract event 0x9a382661d6573da86db000471303be6f0b2b1bb66089b08e3c16a85d7b6e94f8.
//
// Solidity: event ExecuteDecreaseOrder(address indexed account, uint256 orderIndex, address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee, uint256 executionPrice)
func (_OrderBook *OrderBookFilterer) ParseExecuteDecreaseOrder(log types.Log) (*OrderBookExecuteDecreaseOrder, error) {
	event := new(OrderBookExecuteDecreaseOrder)
	if err := _OrderBook.contract.UnpackLog(event, "ExecuteDecreaseOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookExecuteIncreaseOrderIterator is returned from FilterExecuteIncreaseOrder and is used to iterate over the raw logs and unpacked data for ExecuteIncreaseOrder events raised by the OrderBook contract.
type OrderBookExecuteIncreaseOrderIterator struct {
	Event *OrderBookExecuteIncreaseOrder // Event containing the contract specifics and raw log

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
func (it *OrderBookExecuteIncreaseOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookExecuteIncreaseOrder)
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
		it.Event = new(OrderBookExecuteIncreaseOrder)
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
func (it *OrderBookExecuteIncreaseOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookExecuteIncreaseOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookExecuteIncreaseOrder represents a ExecuteIncreaseOrder event raised by the OrderBook contract.
type OrderBookExecuteIncreaseOrder struct {
	Account               common.Address
	OrderIndex            *big.Int
	PurchaseToken         common.Address
	PurchaseTokenAmount   *big.Int
	CollateralToken       common.Address
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	ExecutionFee          *big.Int
	ExecutionPrice        *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterExecuteIncreaseOrder is a free log retrieval operation binding the contract event 0x7fb1c74d1ea6aa1c9c585e17ce8274c8ff98745e85e7459b73f87d784494f58e.
//
// Solidity: event ExecuteIncreaseOrder(address indexed account, uint256 orderIndex, address purchaseToken, uint256 purchaseTokenAmount, address collateralToken, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee, uint256 executionPrice)
func (_OrderBook *OrderBookFilterer) FilterExecuteIncreaseOrder(opts *bind.FilterOpts, account []common.Address) (*OrderBookExecuteIncreaseOrderIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "ExecuteIncreaseOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return &OrderBookExecuteIncreaseOrderIterator{contract: _OrderBook.contract, event: "ExecuteIncreaseOrder", logs: logs, sub: sub}, nil
}

// WatchExecuteIncreaseOrder is a free log subscription operation binding the contract event 0x7fb1c74d1ea6aa1c9c585e17ce8274c8ff98745e85e7459b73f87d784494f58e.
//
// Solidity: event ExecuteIncreaseOrder(address indexed account, uint256 orderIndex, address purchaseToken, uint256 purchaseTokenAmount, address collateralToken, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee, uint256 executionPrice)
func (_OrderBook *OrderBookFilterer) WatchExecuteIncreaseOrder(opts *bind.WatchOpts, sink chan<- *OrderBookExecuteIncreaseOrder, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "ExecuteIncreaseOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookExecuteIncreaseOrder)
				if err := _OrderBook.contract.UnpackLog(event, "ExecuteIncreaseOrder", log); err != nil {
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

// ParseExecuteIncreaseOrder is a log parse operation binding the contract event 0x7fb1c74d1ea6aa1c9c585e17ce8274c8ff98745e85e7459b73f87d784494f58e.
//
// Solidity: event ExecuteIncreaseOrder(address indexed account, uint256 orderIndex, address purchaseToken, uint256 purchaseTokenAmount, address collateralToken, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold, uint256 executionFee, uint256 executionPrice)
func (_OrderBook *OrderBookFilterer) ParseExecuteIncreaseOrder(log types.Log) (*OrderBookExecuteIncreaseOrder, error) {
	event := new(OrderBookExecuteIncreaseOrder)
	if err := _OrderBook.contract.UnpackLog(event, "ExecuteIncreaseOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookExecuteSwapOrderIterator is returned from FilterExecuteSwapOrder and is used to iterate over the raw logs and unpacked data for ExecuteSwapOrder events raised by the OrderBook contract.
type OrderBookExecuteSwapOrderIterator struct {
	Event *OrderBookExecuteSwapOrder // Event containing the contract specifics and raw log

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
func (it *OrderBookExecuteSwapOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookExecuteSwapOrder)
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
		it.Event = new(OrderBookExecuteSwapOrder)
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
func (it *OrderBookExecuteSwapOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookExecuteSwapOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookExecuteSwapOrder represents a ExecuteSwapOrder event raised by the OrderBook contract.
type OrderBookExecuteSwapOrder struct {
	Account               common.Address
	OrderIndex            *big.Int
	Path                  []common.Address
	AmountIn              *big.Int
	MinOut                *big.Int
	AmountOut             *big.Int
	TriggerRatio          *big.Int
	TriggerAboveThreshold bool
	ShouldUnwrap          bool
	ExecutionFee          *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterExecuteSwapOrder is a free log retrieval operation binding the contract event 0x7e1fe496989eea92b738a562dbf9c0ae6aa6fcf3f1ef09e95ee4f7603721706b.
//
// Solidity: event ExecuteSwapOrder(address indexed account, uint256 orderIndex, address[] path, uint256 amountIn, uint256 minOut, uint256 amountOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) FilterExecuteSwapOrder(opts *bind.FilterOpts, account []common.Address) (*OrderBookExecuteSwapOrderIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "ExecuteSwapOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return &OrderBookExecuteSwapOrderIterator{contract: _OrderBook.contract, event: "ExecuteSwapOrder", logs: logs, sub: sub}, nil
}

// WatchExecuteSwapOrder is a free log subscription operation binding the contract event 0x7e1fe496989eea92b738a562dbf9c0ae6aa6fcf3f1ef09e95ee4f7603721706b.
//
// Solidity: event ExecuteSwapOrder(address indexed account, uint256 orderIndex, address[] path, uint256 amountIn, uint256 minOut, uint256 amountOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) WatchExecuteSwapOrder(opts *bind.WatchOpts, sink chan<- *OrderBookExecuteSwapOrder, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "ExecuteSwapOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookExecuteSwapOrder)
				if err := _OrderBook.contract.UnpackLog(event, "ExecuteSwapOrder", log); err != nil {
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

// ParseExecuteSwapOrder is a log parse operation binding the contract event 0x7e1fe496989eea92b738a562dbf9c0ae6aa6fcf3f1ef09e95ee4f7603721706b.
//
// Solidity: event ExecuteSwapOrder(address indexed account, uint256 orderIndex, address[] path, uint256 amountIn, uint256 minOut, uint256 amountOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) ParseExecuteSwapOrder(log types.Log) (*OrderBookExecuteSwapOrder, error) {
	event := new(OrderBookExecuteSwapOrder)
	if err := _OrderBook.contract.UnpackLog(event, "ExecuteSwapOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the OrderBook contract.
type OrderBookInitializeIterator struct {
	Event *OrderBookInitialize // Event containing the contract specifics and raw log

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
func (it *OrderBookInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookInitialize)
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
		it.Event = new(OrderBookInitialize)
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
func (it *OrderBookInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookInitialize represents a Initialize event raised by the OrderBook contract.
type OrderBookInitialize struct {
	Router                    common.Address
	Vault                     common.Address
	Weth                      common.Address
	Usdg                      common.Address
	MinExecutionFee           *big.Int
	MinPurchaseTokenAmountUsd *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0xcfb7ef8749fafc8da2af1ba3d025479ffc4e58f7dc420113e112512a3bda5963.
//
// Solidity: event Initialize(address router, address vault, address weth, address usdg, uint256 minExecutionFee, uint256 minPurchaseTokenAmountUsd)
func (_OrderBook *OrderBookFilterer) FilterInitialize(opts *bind.FilterOpts) (*OrderBookInitializeIterator, error) {

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return &OrderBookInitializeIterator{contract: _OrderBook.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0xcfb7ef8749fafc8da2af1ba3d025479ffc4e58f7dc420113e112512a3bda5963.
//
// Solidity: event Initialize(address router, address vault, address weth, address usdg, uint256 minExecutionFee, uint256 minPurchaseTokenAmountUsd)
func (_OrderBook *OrderBookFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *OrderBookInitialize) (event.Subscription, error) {

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "Initialize")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookInitialize)
				if err := _OrderBook.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ParseInitialize is a log parse operation binding the contract event 0xcfb7ef8749fafc8da2af1ba3d025479ffc4e58f7dc420113e112512a3bda5963.
//
// Solidity: event Initialize(address router, address vault, address weth, address usdg, uint256 minExecutionFee, uint256 minPurchaseTokenAmountUsd)
func (_OrderBook *OrderBookFilterer) ParseInitialize(log types.Log) (*OrderBookInitialize, error) {
	event := new(OrderBookInitialize)
	if err := _OrderBook.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookUpdateDecreaseOrderIterator is returned from FilterUpdateDecreaseOrder and is used to iterate over the raw logs and unpacked data for UpdateDecreaseOrder events raised by the OrderBook contract.
type OrderBookUpdateDecreaseOrderIterator struct {
	Event *OrderBookUpdateDecreaseOrder // Event containing the contract specifics and raw log

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
func (it *OrderBookUpdateDecreaseOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookUpdateDecreaseOrder)
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
		it.Event = new(OrderBookUpdateDecreaseOrder)
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
func (it *OrderBookUpdateDecreaseOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookUpdateDecreaseOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookUpdateDecreaseOrder represents a UpdateDecreaseOrder event raised by the OrderBook contract.
type OrderBookUpdateDecreaseOrder struct {
	Account               common.Address
	OrderIndex            *big.Int
	CollateralToken       common.Address
	CollateralDelta       *big.Int
	IndexToken            common.Address
	SizeDelta             *big.Int
	IsLong                bool
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterUpdateDecreaseOrder is a free log retrieval operation binding the contract event 0x75781255bc71c83f89f29e5a2599f2c174a562d2cd8f2e818a47f132e7280498.
//
// Solidity: event UpdateDecreaseOrder(address indexed account, uint256 orderIndex, address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold)
func (_OrderBook *OrderBookFilterer) FilterUpdateDecreaseOrder(opts *bind.FilterOpts, account []common.Address) (*OrderBookUpdateDecreaseOrderIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "UpdateDecreaseOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return &OrderBookUpdateDecreaseOrderIterator{contract: _OrderBook.contract, event: "UpdateDecreaseOrder", logs: logs, sub: sub}, nil
}

// WatchUpdateDecreaseOrder is a free log subscription operation binding the contract event 0x75781255bc71c83f89f29e5a2599f2c174a562d2cd8f2e818a47f132e7280498.
//
// Solidity: event UpdateDecreaseOrder(address indexed account, uint256 orderIndex, address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold)
func (_OrderBook *OrderBookFilterer) WatchUpdateDecreaseOrder(opts *bind.WatchOpts, sink chan<- *OrderBookUpdateDecreaseOrder, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "UpdateDecreaseOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookUpdateDecreaseOrder)
				if err := _OrderBook.contract.UnpackLog(event, "UpdateDecreaseOrder", log); err != nil {
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

// ParseUpdateDecreaseOrder is a log parse operation binding the contract event 0x75781255bc71c83f89f29e5a2599f2c174a562d2cd8f2e818a47f132e7280498.
//
// Solidity: event UpdateDecreaseOrder(address indexed account, uint256 orderIndex, address collateralToken, uint256 collateralDelta, address indexToken, uint256 sizeDelta, bool isLong, uint256 triggerPrice, bool triggerAboveThreshold)
func (_OrderBook *OrderBookFilterer) ParseUpdateDecreaseOrder(log types.Log) (*OrderBookUpdateDecreaseOrder, error) {
	event := new(OrderBookUpdateDecreaseOrder)
	if err := _OrderBook.contract.UnpackLog(event, "UpdateDecreaseOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookUpdateGovIterator is returned from FilterUpdateGov and is used to iterate over the raw logs and unpacked data for UpdateGov events raised by the OrderBook contract.
type OrderBookUpdateGovIterator struct {
	Event *OrderBookUpdateGov // Event containing the contract specifics and raw log

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
func (it *OrderBookUpdateGovIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookUpdateGov)
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
		it.Event = new(OrderBookUpdateGov)
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
func (it *OrderBookUpdateGovIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookUpdateGovIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookUpdateGov represents a UpdateGov event raised by the OrderBook contract.
type OrderBookUpdateGov struct {
	Gov common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateGov is a free log retrieval operation binding the contract event 0xe24c39186e9137521953beaa8446e71f55b8f12296984f9d4273ceb1af728d90.
//
// Solidity: event UpdateGov(address gov)
func (_OrderBook *OrderBookFilterer) FilterUpdateGov(opts *bind.FilterOpts) (*OrderBookUpdateGovIterator, error) {

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "UpdateGov")
	if err != nil {
		return nil, err
	}
	return &OrderBookUpdateGovIterator{contract: _OrderBook.contract, event: "UpdateGov", logs: logs, sub: sub}, nil
}

// WatchUpdateGov is a free log subscription operation binding the contract event 0xe24c39186e9137521953beaa8446e71f55b8f12296984f9d4273ceb1af728d90.
//
// Solidity: event UpdateGov(address gov)
func (_OrderBook *OrderBookFilterer) WatchUpdateGov(opts *bind.WatchOpts, sink chan<- *OrderBookUpdateGov) (event.Subscription, error) {

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "UpdateGov")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookUpdateGov)
				if err := _OrderBook.contract.UnpackLog(event, "UpdateGov", log); err != nil {
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

// ParseUpdateGov is a log parse operation binding the contract event 0xe24c39186e9137521953beaa8446e71f55b8f12296984f9d4273ceb1af728d90.
//
// Solidity: event UpdateGov(address gov)
func (_OrderBook *OrderBookFilterer) ParseUpdateGov(log types.Log) (*OrderBookUpdateGov, error) {
	event := new(OrderBookUpdateGov)
	if err := _OrderBook.contract.UnpackLog(event, "UpdateGov", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookUpdateIncreaseOrderIterator is returned from FilterUpdateIncreaseOrder and is used to iterate over the raw logs and unpacked data for UpdateIncreaseOrder events raised by the OrderBook contract.
type OrderBookUpdateIncreaseOrderIterator struct {
	Event *OrderBookUpdateIncreaseOrder // Event containing the contract specifics and raw log

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
func (it *OrderBookUpdateIncreaseOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookUpdateIncreaseOrder)
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
		it.Event = new(OrderBookUpdateIncreaseOrder)
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
func (it *OrderBookUpdateIncreaseOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookUpdateIncreaseOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookUpdateIncreaseOrder represents a UpdateIncreaseOrder event raised by the OrderBook contract.
type OrderBookUpdateIncreaseOrder struct {
	Account               common.Address
	OrderIndex            *big.Int
	CollateralToken       common.Address
	IndexToken            common.Address
	IsLong                bool
	SizeDelta             *big.Int
	TriggerPrice          *big.Int
	TriggerAboveThreshold bool
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterUpdateIncreaseOrder is a free log retrieval operation binding the contract event 0x0a0360dd5c354235bbf8d386ba3b24ef8134088e0785677de1504df219d9149a.
//
// Solidity: event UpdateIncreaseOrder(address indexed account, uint256 orderIndex, address collateralToken, address indexToken, bool isLong, uint256 sizeDelta, uint256 triggerPrice, bool triggerAboveThreshold)
func (_OrderBook *OrderBookFilterer) FilterUpdateIncreaseOrder(opts *bind.FilterOpts, account []common.Address) (*OrderBookUpdateIncreaseOrderIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "UpdateIncreaseOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return &OrderBookUpdateIncreaseOrderIterator{contract: _OrderBook.contract, event: "UpdateIncreaseOrder", logs: logs, sub: sub}, nil
}

// WatchUpdateIncreaseOrder is a free log subscription operation binding the contract event 0x0a0360dd5c354235bbf8d386ba3b24ef8134088e0785677de1504df219d9149a.
//
// Solidity: event UpdateIncreaseOrder(address indexed account, uint256 orderIndex, address collateralToken, address indexToken, bool isLong, uint256 sizeDelta, uint256 triggerPrice, bool triggerAboveThreshold)
func (_OrderBook *OrderBookFilterer) WatchUpdateIncreaseOrder(opts *bind.WatchOpts, sink chan<- *OrderBookUpdateIncreaseOrder, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "UpdateIncreaseOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookUpdateIncreaseOrder)
				if err := _OrderBook.contract.UnpackLog(event, "UpdateIncreaseOrder", log); err != nil {
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

// ParseUpdateIncreaseOrder is a log parse operation binding the contract event 0x0a0360dd5c354235bbf8d386ba3b24ef8134088e0785677de1504df219d9149a.
//
// Solidity: event UpdateIncreaseOrder(address indexed account, uint256 orderIndex, address collateralToken, address indexToken, bool isLong, uint256 sizeDelta, uint256 triggerPrice, bool triggerAboveThreshold)
func (_OrderBook *OrderBookFilterer) ParseUpdateIncreaseOrder(log types.Log) (*OrderBookUpdateIncreaseOrder, error) {
	event := new(OrderBookUpdateIncreaseOrder)
	if err := _OrderBook.contract.UnpackLog(event, "UpdateIncreaseOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookUpdateMinExecutionFeeIterator is returned from FilterUpdateMinExecutionFee and is used to iterate over the raw logs and unpacked data for UpdateMinExecutionFee events raised by the OrderBook contract.
type OrderBookUpdateMinExecutionFeeIterator struct {
	Event *OrderBookUpdateMinExecutionFee // Event containing the contract specifics and raw log

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
func (it *OrderBookUpdateMinExecutionFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookUpdateMinExecutionFee)
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
		it.Event = new(OrderBookUpdateMinExecutionFee)
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
func (it *OrderBookUpdateMinExecutionFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookUpdateMinExecutionFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookUpdateMinExecutionFee represents a UpdateMinExecutionFee event raised by the OrderBook contract.
type OrderBookUpdateMinExecutionFee struct {
	MinExecutionFee *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinExecutionFee is a free log retrieval operation binding the contract event 0xbde5eafdc37b81830d70124cddccaaa6d034e71dda3c8fc18a959ca76a7cbcfc.
//
// Solidity: event UpdateMinExecutionFee(uint256 minExecutionFee)
func (_OrderBook *OrderBookFilterer) FilterUpdateMinExecutionFee(opts *bind.FilterOpts) (*OrderBookUpdateMinExecutionFeeIterator, error) {

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "UpdateMinExecutionFee")
	if err != nil {
		return nil, err
	}
	return &OrderBookUpdateMinExecutionFeeIterator{contract: _OrderBook.contract, event: "UpdateMinExecutionFee", logs: logs, sub: sub}, nil
}

// WatchUpdateMinExecutionFee is a free log subscription operation binding the contract event 0xbde5eafdc37b81830d70124cddccaaa6d034e71dda3c8fc18a959ca76a7cbcfc.
//
// Solidity: event UpdateMinExecutionFee(uint256 minExecutionFee)
func (_OrderBook *OrderBookFilterer) WatchUpdateMinExecutionFee(opts *bind.WatchOpts, sink chan<- *OrderBookUpdateMinExecutionFee) (event.Subscription, error) {

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "UpdateMinExecutionFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookUpdateMinExecutionFee)
				if err := _OrderBook.contract.UnpackLog(event, "UpdateMinExecutionFee", log); err != nil {
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

// ParseUpdateMinExecutionFee is a log parse operation binding the contract event 0xbde5eafdc37b81830d70124cddccaaa6d034e71dda3c8fc18a959ca76a7cbcfc.
//
// Solidity: event UpdateMinExecutionFee(uint256 minExecutionFee)
func (_OrderBook *OrderBookFilterer) ParseUpdateMinExecutionFee(log types.Log) (*OrderBookUpdateMinExecutionFee, error) {
	event := new(OrderBookUpdateMinExecutionFee)
	if err := _OrderBook.contract.UnpackLog(event, "UpdateMinExecutionFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookUpdateMinPurchaseTokenAmountUsdIterator is returned from FilterUpdateMinPurchaseTokenAmountUsd and is used to iterate over the raw logs and unpacked data for UpdateMinPurchaseTokenAmountUsd events raised by the OrderBook contract.
type OrderBookUpdateMinPurchaseTokenAmountUsdIterator struct {
	Event *OrderBookUpdateMinPurchaseTokenAmountUsd // Event containing the contract specifics and raw log

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
func (it *OrderBookUpdateMinPurchaseTokenAmountUsdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookUpdateMinPurchaseTokenAmountUsd)
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
		it.Event = new(OrderBookUpdateMinPurchaseTokenAmountUsd)
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
func (it *OrderBookUpdateMinPurchaseTokenAmountUsdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookUpdateMinPurchaseTokenAmountUsdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookUpdateMinPurchaseTokenAmountUsd represents a UpdateMinPurchaseTokenAmountUsd event raised by the OrderBook contract.
type OrderBookUpdateMinPurchaseTokenAmountUsd struct {
	MinPurchaseTokenAmountUsd *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinPurchaseTokenAmountUsd is a free log retrieval operation binding the contract event 0xe46d9daf6d25f7615efa1d0183b90ac6759d85014b598e409aadf0fd918d59a6.
//
// Solidity: event UpdateMinPurchaseTokenAmountUsd(uint256 minPurchaseTokenAmountUsd)
func (_OrderBook *OrderBookFilterer) FilterUpdateMinPurchaseTokenAmountUsd(opts *bind.FilterOpts) (*OrderBookUpdateMinPurchaseTokenAmountUsdIterator, error) {

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "UpdateMinPurchaseTokenAmountUsd")
	if err != nil {
		return nil, err
	}
	return &OrderBookUpdateMinPurchaseTokenAmountUsdIterator{contract: _OrderBook.contract, event: "UpdateMinPurchaseTokenAmountUsd", logs: logs, sub: sub}, nil
}

// WatchUpdateMinPurchaseTokenAmountUsd is a free log subscription operation binding the contract event 0xe46d9daf6d25f7615efa1d0183b90ac6759d85014b598e409aadf0fd918d59a6.
//
// Solidity: event UpdateMinPurchaseTokenAmountUsd(uint256 minPurchaseTokenAmountUsd)
func (_OrderBook *OrderBookFilterer) WatchUpdateMinPurchaseTokenAmountUsd(opts *bind.WatchOpts, sink chan<- *OrderBookUpdateMinPurchaseTokenAmountUsd) (event.Subscription, error) {

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "UpdateMinPurchaseTokenAmountUsd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookUpdateMinPurchaseTokenAmountUsd)
				if err := _OrderBook.contract.UnpackLog(event, "UpdateMinPurchaseTokenAmountUsd", log); err != nil {
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

// ParseUpdateMinPurchaseTokenAmountUsd is a log parse operation binding the contract event 0xe46d9daf6d25f7615efa1d0183b90ac6759d85014b598e409aadf0fd918d59a6.
//
// Solidity: event UpdateMinPurchaseTokenAmountUsd(uint256 minPurchaseTokenAmountUsd)
func (_OrderBook *OrderBookFilterer) ParseUpdateMinPurchaseTokenAmountUsd(log types.Log) (*OrderBookUpdateMinPurchaseTokenAmountUsd, error) {
	event := new(OrderBookUpdateMinPurchaseTokenAmountUsd)
	if err := _OrderBook.contract.UnpackLog(event, "UpdateMinPurchaseTokenAmountUsd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookUpdateSwapOrderIterator is returned from FilterUpdateSwapOrder and is used to iterate over the raw logs and unpacked data for UpdateSwapOrder events raised by the OrderBook contract.
type OrderBookUpdateSwapOrderIterator struct {
	Event *OrderBookUpdateSwapOrder // Event containing the contract specifics and raw log

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
func (it *OrderBookUpdateSwapOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookUpdateSwapOrder)
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
		it.Event = new(OrderBookUpdateSwapOrder)
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
func (it *OrderBookUpdateSwapOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookUpdateSwapOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookUpdateSwapOrder represents a UpdateSwapOrder event raised by the OrderBook contract.
type OrderBookUpdateSwapOrder struct {
	Account               common.Address
	OrdexIndex            *big.Int
	Path                  []common.Address
	AmountIn              *big.Int
	MinOut                *big.Int
	TriggerRatio          *big.Int
	TriggerAboveThreshold bool
	ShouldUnwrap          bool
	ExecutionFee          *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterUpdateSwapOrder is a free log retrieval operation binding the contract event 0xa7f9f4a25eb76f5ec01b1a429d95d6a00833f0f137c88827c58799a1c1ff0dfe.
//
// Solidity: event UpdateSwapOrder(address indexed account, uint256 ordexIndex, address[] path, uint256 amountIn, uint256 minOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) FilterUpdateSwapOrder(opts *bind.FilterOpts, account []common.Address) (*OrderBookUpdateSwapOrderIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "UpdateSwapOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return &OrderBookUpdateSwapOrderIterator{contract: _OrderBook.contract, event: "UpdateSwapOrder", logs: logs, sub: sub}, nil
}

// WatchUpdateSwapOrder is a free log subscription operation binding the contract event 0xa7f9f4a25eb76f5ec01b1a429d95d6a00833f0f137c88827c58799a1c1ff0dfe.
//
// Solidity: event UpdateSwapOrder(address indexed account, uint256 ordexIndex, address[] path, uint256 amountIn, uint256 minOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) WatchUpdateSwapOrder(opts *bind.WatchOpts, sink chan<- *OrderBookUpdateSwapOrder, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "UpdateSwapOrder", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookUpdateSwapOrder)
				if err := _OrderBook.contract.UnpackLog(event, "UpdateSwapOrder", log); err != nil {
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

// ParseUpdateSwapOrder is a log parse operation binding the contract event 0xa7f9f4a25eb76f5ec01b1a429d95d6a00833f0f137c88827c58799a1c1ff0dfe.
//
// Solidity: event UpdateSwapOrder(address indexed account, uint256 ordexIndex, address[] path, uint256 amountIn, uint256 minOut, uint256 triggerRatio, bool triggerAboveThreshold, bool shouldUnwrap, uint256 executionFee)
func (_OrderBook *OrderBookFilterer) ParseUpdateSwapOrder(log types.Log) (*OrderBookUpdateSwapOrder, error) {
	event := new(OrderBookUpdateSwapOrder)
	if err := _OrderBook.contract.UnpackLog(event, "UpdateSwapOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

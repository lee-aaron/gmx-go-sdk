// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Vault

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

// VaultMetaData contains all meta data concerning the Vault contract.
var VaultMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BASIS_POINTS_DIVISOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FUNDING_RATE_PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FEE_BASIS_POINTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FUNDING_RATE_FACTOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_LIQUIDATION_FEE_USD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_FUNDING_RATE_INTERVAL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_LEVERAGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PRICE_PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"USDG_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addRouter\",\"inputs\":[{\"name\":\"_router\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adjustForDecimals\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokenDiv\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenMul\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allWhitelistedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allWhitelistedTokensLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approvedRouters\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bufferAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"buyUSDG\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clearTokenConfig\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cumulativeFundingRates\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreasePosition\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"directPoolDeposit\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"errorController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"errors\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeReserves\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fundingInterval\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fundingRateFactor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDelta\",\"inputs\":[{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_averagePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_lastIncreasedTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEntryFundingRate\",\"inputs\":[{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeBasisPoints\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdgDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_feeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_taxBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_increment\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFundingFee\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_entryFundingRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalShortDelta\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxPrice\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinPrice\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextAveragePrice\",\"inputs\":[{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_averagePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_nextPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_lastIncreasedTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextFundingRate\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextGlobalShortAveragePrice\",\"inputs\":[{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nextPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPosition\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPositionDelta\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPositionFee\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPositionKey\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getPositionLeverage\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRedemptionAmount\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdgAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRedemptionCollateral\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRedemptionCollateralUsd\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTargetUsdgAmount\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUtilisation\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalShortAveragePrices\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalShortSizes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"guaranteedUsd\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasDynamicFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inManagerMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inPrivateLiquidationMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"includeAmmPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increasePosition\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdg\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_priceFeed\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_liquidationFeeUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_fundingRateFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stableFundingRateFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isInitialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isLeverageEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isLiquidator\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isManager\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSwapEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastFundingTimes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"liquidatePosition\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"liquidationFeeUsd\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marginFeeBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxGasPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxGlobalShortSizes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxLeverage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxUsdgAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minProfitBasisPoints\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minProfitTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mintBurnFeeBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"positions\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateral\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"averagePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entryFundingRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reserveAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"realisedPnl\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"lastIncreasedTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"priceFeed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeRouter\",\"inputs\":[{\"name\":\"_router\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reservedAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sellUSDG\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBufferAmount\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setError\",\"inputs\":[{\"name\":\"_errorCode\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setErrorController\",\"inputs\":[{\"name\":\"_errorController\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFees\",\"inputs\":[{\"name\":\"_taxBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stableTaxBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_mintBurnFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_swapFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stableSwapFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_marginFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_liquidationFeeUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minProfitTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_hasDynamicFees\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFundingRate\",\"inputs\":[{\"name\":\"_fundingInterval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_fundingRateFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stableFundingRateFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInManagerMode\",\"inputs\":[{\"name\":\"_inManagerMode\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInPrivateLiquidationMode\",\"inputs\":[{\"name\":\"_inPrivateLiquidationMode\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIsLeverageEnabled\",\"inputs\":[{\"name\":\"_isLeverageEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIsSwapEnabled\",\"inputs\":[{\"name\":\"_isSwapEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLiquidator\",\"inputs\":[{\"name\":\"_liquidator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setManager\",\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isManager\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxGasPrice\",\"inputs\":[{\"name\":\"_maxGasPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxGlobalShortSize\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxLeverage\",\"inputs\":[{\"name\":\"_maxLeverage\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPriceFeed\",\"inputs\":[{\"name\":\"_priceFeed\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenConfig\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenDecimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokenWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minProfitBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_maxUsdgAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isStable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_isShortable\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUsdgAmount\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVaultUtils\",\"inputs\":[{\"name\":\"_vaultUtils\",\"type\":\"address\",\"internalType\":\"contractIVaultUtils\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shortableTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stableFundingRateFactor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stableSwapFeeBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stableTaxBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stableTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"_tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swapFeeBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taxBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenBalances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenDecimals\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenToUsdMin\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenWeights\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalTokenWeights\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateCumulativeFundingRate\",\"inputs\":[{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeVault\",\"inputs\":[{\"name\":\"_newVault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"usdToToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdToTokenMax\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdToTokenMin\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdg\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdgAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"useSwapPricing\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateLiquidation\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_raise\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vaultUtils\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIVaultUtils\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelistedTokenCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelistedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BuyUSDG\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"usdgAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feeBasisPoints\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClosePosition\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"size\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateral\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"averagePrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"entryFundingRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reserveAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"realisedPnl\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollectMarginFees\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feeUsd\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feeTokens\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollectSwapFees\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feeUsd\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feeTokens\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecreaseGuaranteedUsd\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecreasePoolAmount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecreasePosition\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecreaseReservedAmount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecreaseUsdgAmount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DirectPoolDeposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncreaseGuaranteedUsd\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncreasePoolAmount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncreasePosition\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncreaseReservedAmount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncreaseUsdgAmount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidatePosition\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"size\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateral\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reserveAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"realisedPnl\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"markPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SellUSDG\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"usdgAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feeBasisPoints\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Swap\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenIn\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenOut\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOutAfterFees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feeBasisPoints\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateFundingRate\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"fundingRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatePnl\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"hasProfit\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"delta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatePosition\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"size\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateral\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"averagePrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"entryFundingRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reserveAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"realisedPnl\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"markPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// VaultABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultMetaData.ABI instead.
var VaultABI = VaultMetaData.ABI

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	VaultCaller     // Read-only binding to the contract
	VaultTransactor // Write-only binding to the contract
	VaultFilterer   // Log filterer for contract events
}

// VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultSession struct {
	Contract     *Vault            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultCallerSession struct {
	Contract *VaultCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTransactorSession struct {
	Contract     *VaultTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultRaw struct {
	Contract *Vault // Generic contract binding to access the raw methods on
}

// VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultCallerRaw struct {
	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTransactorRaw struct {
	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVault creates a new instance of Vault, bound to a specific deployed contract.
func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
	contract, err := bindVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
	contract, err := bindVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultCaller{contract: contract}, nil
}

// NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
	contract, err := bindVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTransactor{contract: contract}, nil
}

// NewVaultFilterer creates a new log filterer instance of Vault, bound to a specific deployed contract.
func NewVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFilterer, error) {
	contract, err := bindVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultFilterer{contract: contract}, nil
}

// bindVault binds a generic wrapper to an already deployed contract.
func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_Vault *VaultCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_Vault *VaultSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _Vault.Contract.BASISPOINTSDIVISOR(&_Vault.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_Vault *VaultCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _Vault.Contract.BASISPOINTSDIVISOR(&_Vault.CallOpts)
}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_Vault *VaultCaller) FUNDINGRATEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "FUNDING_RATE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_Vault *VaultSession) FUNDINGRATEPRECISION() (*big.Int, error) {
	return _Vault.Contract.FUNDINGRATEPRECISION(&_Vault.CallOpts)
}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_Vault *VaultCallerSession) FUNDINGRATEPRECISION() (*big.Int, error) {
	return _Vault.Contract.FUNDINGRATEPRECISION(&_Vault.CallOpts)
}

// MAXFEEBASISPOINTS is a free data retrieval call binding the contract method 0x4befe2ca.
//
// Solidity: function MAX_FEE_BASIS_POINTS() view returns(uint256)
func (_Vault *VaultCaller) MAXFEEBASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "MAX_FEE_BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEEBASISPOINTS is a free data retrieval call binding the contract method 0x4befe2ca.
//
// Solidity: function MAX_FEE_BASIS_POINTS() view returns(uint256)
func (_Vault *VaultSession) MAXFEEBASISPOINTS() (*big.Int, error) {
	return _Vault.Contract.MAXFEEBASISPOINTS(&_Vault.CallOpts)
}

// MAXFEEBASISPOINTS is a free data retrieval call binding the contract method 0x4befe2ca.
//
// Solidity: function MAX_FEE_BASIS_POINTS() view returns(uint256)
func (_Vault *VaultCallerSession) MAXFEEBASISPOINTS() (*big.Int, error) {
	return _Vault.Contract.MAXFEEBASISPOINTS(&_Vault.CallOpts)
}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_Vault *VaultCaller) MAXFUNDINGRATEFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "MAX_FUNDING_RATE_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_Vault *VaultSession) MAXFUNDINGRATEFACTOR() (*big.Int, error) {
	return _Vault.Contract.MAXFUNDINGRATEFACTOR(&_Vault.CallOpts)
}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_Vault *VaultCallerSession) MAXFUNDINGRATEFACTOR() (*big.Int, error) {
	return _Vault.Contract.MAXFUNDINGRATEFACTOR(&_Vault.CallOpts)
}

// MAXLIQUIDATIONFEEUSD is a free data retrieval call binding the contract method 0x07c58752.
//
// Solidity: function MAX_LIQUIDATION_FEE_USD() view returns(uint256)
func (_Vault *VaultCaller) MAXLIQUIDATIONFEEUSD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "MAX_LIQUIDATION_FEE_USD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLIQUIDATIONFEEUSD is a free data retrieval call binding the contract method 0x07c58752.
//
// Solidity: function MAX_LIQUIDATION_FEE_USD() view returns(uint256)
func (_Vault *VaultSession) MAXLIQUIDATIONFEEUSD() (*big.Int, error) {
	return _Vault.Contract.MAXLIQUIDATIONFEEUSD(&_Vault.CallOpts)
}

// MAXLIQUIDATIONFEEUSD is a free data retrieval call binding the contract method 0x07c58752.
//
// Solidity: function MAX_LIQUIDATION_FEE_USD() view returns(uint256)
func (_Vault *VaultCallerSession) MAXLIQUIDATIONFEEUSD() (*big.Int, error) {
	return _Vault.Contract.MAXLIQUIDATIONFEEUSD(&_Vault.CallOpts)
}

// MINFUNDINGRATEINTERVAL is a free data retrieval call binding the contract method 0xfce28c10.
//
// Solidity: function MIN_FUNDING_RATE_INTERVAL() view returns(uint256)
func (_Vault *VaultCaller) MINFUNDINGRATEINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "MIN_FUNDING_RATE_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINFUNDINGRATEINTERVAL is a free data retrieval call binding the contract method 0xfce28c10.
//
// Solidity: function MIN_FUNDING_RATE_INTERVAL() view returns(uint256)
func (_Vault *VaultSession) MINFUNDINGRATEINTERVAL() (*big.Int, error) {
	return _Vault.Contract.MINFUNDINGRATEINTERVAL(&_Vault.CallOpts)
}

// MINFUNDINGRATEINTERVAL is a free data retrieval call binding the contract method 0xfce28c10.
//
// Solidity: function MIN_FUNDING_RATE_INTERVAL() view returns(uint256)
func (_Vault *VaultCallerSession) MINFUNDINGRATEINTERVAL() (*big.Int, error) {
	return _Vault.Contract.MINFUNDINGRATEINTERVAL(&_Vault.CallOpts)
}

// MINLEVERAGE is a free data retrieval call binding the contract method 0x34c1557d.
//
// Solidity: function MIN_LEVERAGE() view returns(uint256)
func (_Vault *VaultCaller) MINLEVERAGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "MIN_LEVERAGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINLEVERAGE is a free data retrieval call binding the contract method 0x34c1557d.
//
// Solidity: function MIN_LEVERAGE() view returns(uint256)
func (_Vault *VaultSession) MINLEVERAGE() (*big.Int, error) {
	return _Vault.Contract.MINLEVERAGE(&_Vault.CallOpts)
}

// MINLEVERAGE is a free data retrieval call binding the contract method 0x34c1557d.
//
// Solidity: function MIN_LEVERAGE() view returns(uint256)
func (_Vault *VaultCallerSession) MINLEVERAGE() (*big.Int, error) {
	return _Vault.Contract.MINLEVERAGE(&_Vault.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_Vault *VaultCaller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_Vault *VaultSession) PRICEPRECISION() (*big.Int, error) {
	return _Vault.Contract.PRICEPRECISION(&_Vault.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_Vault *VaultCallerSession) PRICEPRECISION() (*big.Int, error) {
	return _Vault.Contract.PRICEPRECISION(&_Vault.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_Vault *VaultCaller) USDGDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "USDG_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_Vault *VaultSession) USDGDECIMALS() (*big.Int, error) {
	return _Vault.Contract.USDGDECIMALS(&_Vault.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_Vault *VaultCallerSession) USDGDECIMALS() (*big.Int, error) {
	return _Vault.Contract.USDGDECIMALS(&_Vault.CallOpts)
}

// AdjustForDecimals is a free data retrieval call binding the contract method 0x42152873.
//
// Solidity: function adjustForDecimals(uint256 _amount, address _tokenDiv, address _tokenMul) view returns(uint256)
func (_Vault *VaultCaller) AdjustForDecimals(opts *bind.CallOpts, _amount *big.Int, _tokenDiv common.Address, _tokenMul common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "adjustForDecimals", _amount, _tokenDiv, _tokenMul)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdjustForDecimals is a free data retrieval call binding the contract method 0x42152873.
//
// Solidity: function adjustForDecimals(uint256 _amount, address _tokenDiv, address _tokenMul) view returns(uint256)
func (_Vault *VaultSession) AdjustForDecimals(_amount *big.Int, _tokenDiv common.Address, _tokenMul common.Address) (*big.Int, error) {
	return _Vault.Contract.AdjustForDecimals(&_Vault.CallOpts, _amount, _tokenDiv, _tokenMul)
}

// AdjustForDecimals is a free data retrieval call binding the contract method 0x42152873.
//
// Solidity: function adjustForDecimals(uint256 _amount, address _tokenDiv, address _tokenMul) view returns(uint256)
func (_Vault *VaultCallerSession) AdjustForDecimals(_amount *big.Int, _tokenDiv common.Address, _tokenMul common.Address) (*big.Int, error) {
	return _Vault.Contract.AdjustForDecimals(&_Vault.CallOpts, _amount, _tokenDiv, _tokenMul)
}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_Vault *VaultCaller) AllWhitelistedTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "allWhitelistedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_Vault *VaultSession) AllWhitelistedTokens(arg0 *big.Int) (common.Address, error) {
	return _Vault.Contract.AllWhitelistedTokens(&_Vault.CallOpts, arg0)
}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_Vault *VaultCallerSession) AllWhitelistedTokens(arg0 *big.Int) (common.Address, error) {
	return _Vault.Contract.AllWhitelistedTokens(&_Vault.CallOpts, arg0)
}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_Vault *VaultCaller) AllWhitelistedTokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "allWhitelistedTokensLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_Vault *VaultSession) AllWhitelistedTokensLength() (*big.Int, error) {
	return _Vault.Contract.AllWhitelistedTokensLength(&_Vault.CallOpts)
}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_Vault *VaultCallerSession) AllWhitelistedTokensLength() (*big.Int, error) {
	return _Vault.Contract.AllWhitelistedTokensLength(&_Vault.CallOpts)
}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address , address ) view returns(bool)
func (_Vault *VaultCaller) ApprovedRouters(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "approvedRouters", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address , address ) view returns(bool)
func (_Vault *VaultSession) ApprovedRouters(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Vault.Contract.ApprovedRouters(&_Vault.CallOpts, arg0, arg1)
}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address , address ) view returns(bool)
func (_Vault *VaultCallerSession) ApprovedRouters(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Vault.Contract.ApprovedRouters(&_Vault.CallOpts, arg0, arg1)
}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address ) view returns(uint256)
func (_Vault *VaultCaller) BufferAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "bufferAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address ) view returns(uint256)
func (_Vault *VaultSession) BufferAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.BufferAmounts(&_Vault.CallOpts, arg0)
}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address ) view returns(uint256)
func (_Vault *VaultCallerSession) BufferAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.BufferAmounts(&_Vault.CallOpts, arg0)
}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address ) view returns(uint256)
func (_Vault *VaultCaller) CumulativeFundingRates(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "cumulativeFundingRates", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address ) view returns(uint256)
func (_Vault *VaultSession) CumulativeFundingRates(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.CumulativeFundingRates(&_Vault.CallOpts, arg0)
}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address ) view returns(uint256)
func (_Vault *VaultCallerSession) CumulativeFundingRates(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.CumulativeFundingRates(&_Vault.CallOpts, arg0)
}

// ErrorController is a free data retrieval call binding the contract method 0x48f35cbb.
//
// Solidity: function errorController() view returns(address)
func (_Vault *VaultCaller) ErrorController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "errorController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ErrorController is a free data retrieval call binding the contract method 0x48f35cbb.
//
// Solidity: function errorController() view returns(address)
func (_Vault *VaultSession) ErrorController() (common.Address, error) {
	return _Vault.Contract.ErrorController(&_Vault.CallOpts)
}

// ErrorController is a free data retrieval call binding the contract method 0x48f35cbb.
//
// Solidity: function errorController() view returns(address)
func (_Vault *VaultCallerSession) ErrorController() (common.Address, error) {
	return _Vault.Contract.ErrorController(&_Vault.CallOpts)
}

// Errors is a free data retrieval call binding the contract method 0xfed1a606.
//
// Solidity: function errors(uint256 ) view returns(string)
func (_Vault *VaultCaller) Errors(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "errors", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Errors is a free data retrieval call binding the contract method 0xfed1a606.
//
// Solidity: function errors(uint256 ) view returns(string)
func (_Vault *VaultSession) Errors(arg0 *big.Int) (string, error) {
	return _Vault.Contract.Errors(&_Vault.CallOpts, arg0)
}

// Errors is a free data retrieval call binding the contract method 0xfed1a606.
//
// Solidity: function errors(uint256 ) view returns(string)
func (_Vault *VaultCallerSession) Errors(arg0 *big.Int) (string, error) {
	return _Vault.Contract.Errors(&_Vault.CallOpts, arg0)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_Vault *VaultCaller) FeeReserves(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "feeReserves", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_Vault *VaultSession) FeeReserves(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.FeeReserves(&_Vault.CallOpts, arg0)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_Vault *VaultCallerSession) FeeReserves(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.FeeReserves(&_Vault.CallOpts, arg0)
}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_Vault *VaultCaller) FundingInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "fundingInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_Vault *VaultSession) FundingInterval() (*big.Int, error) {
	return _Vault.Contract.FundingInterval(&_Vault.CallOpts)
}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_Vault *VaultCallerSession) FundingInterval() (*big.Int, error) {
	return _Vault.Contract.FundingInterval(&_Vault.CallOpts)
}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_Vault *VaultCaller) FundingRateFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "fundingRateFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_Vault *VaultSession) FundingRateFactor() (*big.Int, error) {
	return _Vault.Contract.FundingRateFactor(&_Vault.CallOpts)
}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_Vault *VaultCallerSession) FundingRateFactor() (*big.Int, error) {
	return _Vault.Contract.FundingRateFactor(&_Vault.CallOpts)
}

// GetDelta is a free data retrieval call binding the contract method 0x5c07eaab.
//
// Solidity: function getDelta(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _lastIncreasedTime) view returns(bool, uint256)
func (_Vault *VaultCaller) GetDelta(opts *bind.CallOpts, _indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getDelta", _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetDelta is a free data retrieval call binding the contract method 0x5c07eaab.
//
// Solidity: function getDelta(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _lastIncreasedTime) view returns(bool, uint256)
func (_Vault *VaultSession) GetDelta(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	return _Vault.Contract.GetDelta(&_Vault.CallOpts, _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)
}

// GetDelta is a free data retrieval call binding the contract method 0x5c07eaab.
//
// Solidity: function getDelta(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _lastIncreasedTime) view returns(bool, uint256)
func (_Vault *VaultCallerSession) GetDelta(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	return _Vault.Contract.GetDelta(&_Vault.CallOpts, _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)
}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_Vault *VaultCaller) GetEntryFundingRate(opts *bind.CallOpts, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getEntryFundingRate", _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_Vault *VaultSession) GetEntryFundingRate(_collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _Vault.Contract.GetEntryFundingRate(&_Vault.CallOpts, _collateralToken, _indexToken, _isLong)
}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_Vault *VaultCallerSession) GetEntryFundingRate(_collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _Vault.Contract.GetEntryFundingRate(&_Vault.CallOpts, _collateralToken, _indexToken, _isLong)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_Vault *VaultCaller) GetFeeBasisPoints(opts *bind.CallOpts, _token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getFeeBasisPoints", _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_Vault *VaultSession) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _Vault.Contract.GetFeeBasisPoints(&_Vault.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_Vault *VaultCallerSession) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _Vault.Contract.GetFeeBasisPoints(&_Vault.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_Vault *VaultCaller) GetFundingFee(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getFundingFee", _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_Vault *VaultSession) GetFundingFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetFundingFee(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)
}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_Vault *VaultCallerSession) GetFundingFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetFundingFee(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)
}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xb364accb.
//
// Solidity: function getGlobalShortDelta(address _token) view returns(bool, uint256)
func (_Vault *VaultCaller) GetGlobalShortDelta(opts *bind.CallOpts, _token common.Address) (bool, *big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getGlobalShortDelta", _token)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xb364accb.
//
// Solidity: function getGlobalShortDelta(address _token) view returns(bool, uint256)
func (_Vault *VaultSession) GetGlobalShortDelta(_token common.Address) (bool, *big.Int, error) {
	return _Vault.Contract.GetGlobalShortDelta(&_Vault.CallOpts, _token)
}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xb364accb.
//
// Solidity: function getGlobalShortDelta(address _token) view returns(bool, uint256)
func (_Vault *VaultCallerSession) GetGlobalShortDelta(_token common.Address) (bool, *big.Int, error) {
	return _Vault.Contract.GetGlobalShortDelta(&_Vault.CallOpts, _token)
}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_Vault *VaultCaller) GetMaxPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getMaxPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_Vault *VaultSession) GetMaxPrice(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetMaxPrice(&_Vault.CallOpts, _token)
}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_Vault *VaultCallerSession) GetMaxPrice(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetMaxPrice(&_Vault.CallOpts, _token)
}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_Vault *VaultCaller) GetMinPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getMinPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_Vault *VaultSession) GetMinPrice(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetMinPrice(&_Vault.CallOpts, _token)
}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_Vault *VaultCallerSession) GetMinPrice(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetMinPrice(&_Vault.CallOpts, _token)
}

// GetNextAveragePrice is a free data retrieval call binding the contract method 0xdb97495f.
//
// Solidity: function getNextAveragePrice(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _nextPrice, uint256 _sizeDelta, uint256 _lastIncreasedTime) view returns(uint256)
func (_Vault *VaultCaller) GetNextAveragePrice(opts *bind.CallOpts, _indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _nextPrice *big.Int, _sizeDelta *big.Int, _lastIncreasedTime *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getNextAveragePrice", _indexToken, _size, _averagePrice, _isLong, _nextPrice, _sizeDelta, _lastIncreasedTime)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextAveragePrice is a free data retrieval call binding the contract method 0xdb97495f.
//
// Solidity: function getNextAveragePrice(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _nextPrice, uint256 _sizeDelta, uint256 _lastIncreasedTime) view returns(uint256)
func (_Vault *VaultSession) GetNextAveragePrice(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _nextPrice *big.Int, _sizeDelta *big.Int, _lastIncreasedTime *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetNextAveragePrice(&_Vault.CallOpts, _indexToken, _size, _averagePrice, _isLong, _nextPrice, _sizeDelta, _lastIncreasedTime)
}

// GetNextAveragePrice is a free data retrieval call binding the contract method 0xdb97495f.
//
// Solidity: function getNextAveragePrice(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _nextPrice, uint256 _sizeDelta, uint256 _lastIncreasedTime) view returns(uint256)
func (_Vault *VaultCallerSession) GetNextAveragePrice(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _nextPrice *big.Int, _sizeDelta *big.Int, _lastIncreasedTime *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetNextAveragePrice(&_Vault.CallOpts, _indexToken, _size, _averagePrice, _isLong, _nextPrice, _sizeDelta, _lastIncreasedTime)
}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_Vault *VaultCaller) GetNextFundingRate(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getNextFundingRate", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_Vault *VaultSession) GetNextFundingRate(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetNextFundingRate(&_Vault.CallOpts, _token)
}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_Vault *VaultCallerSession) GetNextFundingRate(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetNextFundingRate(&_Vault.CallOpts, _token)
}

// GetNextGlobalShortAveragePrice is a free data retrieval call binding the contract method 0x9d7432ca.
//
// Solidity: function getNextGlobalShortAveragePrice(address _indexToken, uint256 _nextPrice, uint256 _sizeDelta) view returns(uint256)
func (_Vault *VaultCaller) GetNextGlobalShortAveragePrice(opts *bind.CallOpts, _indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getNextGlobalShortAveragePrice", _indexToken, _nextPrice, _sizeDelta)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextGlobalShortAveragePrice is a free data retrieval call binding the contract method 0x9d7432ca.
//
// Solidity: function getNextGlobalShortAveragePrice(address _indexToken, uint256 _nextPrice, uint256 _sizeDelta) view returns(uint256)
func (_Vault *VaultSession) GetNextGlobalShortAveragePrice(_indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetNextGlobalShortAveragePrice(&_Vault.CallOpts, _indexToken, _nextPrice, _sizeDelta)
}

// GetNextGlobalShortAveragePrice is a free data retrieval call binding the contract method 0x9d7432ca.
//
// Solidity: function getNextGlobalShortAveragePrice(address _indexToken, uint256 _nextPrice, uint256 _sizeDelta) view returns(uint256)
func (_Vault *VaultCallerSession) GetNextGlobalShortAveragePrice(_indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetNextGlobalShortAveragePrice(&_Vault.CallOpts, _indexToken, _nextPrice, _sizeDelta)
}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256, uint256, uint256, uint256, uint256, uint256, bool, uint256)
func (_Vault *VaultCaller) GetPosition(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getPosition", _account, _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(bool)).(*bool)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, out6, out7, err

}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256, uint256, uint256, uint256, uint256, uint256, bool, uint256)
func (_Vault *VaultSession) GetPosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _Vault.Contract.GetPosition(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256, uint256, uint256, uint256, uint256, uint256, bool, uint256)
func (_Vault *VaultCallerSession) GetPosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _Vault.Contract.GetPosition(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionDelta is a free data retrieval call binding the contract method 0x45a6f370.
//
// Solidity: function getPositionDelta(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(bool, uint256)
func (_Vault *VaultCaller) GetPositionDelta(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (bool, *big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getPositionDelta", _account, _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPositionDelta is a free data retrieval call binding the contract method 0x45a6f370.
//
// Solidity: function getPositionDelta(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(bool, uint256)
func (_Vault *VaultSession) GetPositionDelta(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (bool, *big.Int, error) {
	return _Vault.Contract.GetPositionDelta(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionDelta is a free data retrieval call binding the contract method 0x45a6f370.
//
// Solidity: function getPositionDelta(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(bool, uint256)
func (_Vault *VaultCallerSession) GetPositionDelta(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (bool, *big.Int, error) {
	return _Vault.Contract.GetPositionDelta(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_Vault *VaultCaller) GetPositionFee(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getPositionFee", _account, _collateralToken, _indexToken, _isLong, _sizeDelta)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_Vault *VaultSession) GetPositionFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetPositionFee(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta)
}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_Vault *VaultCallerSession) GetPositionFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetPositionFee(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta)
}

// GetPositionKey is a free data retrieval call binding the contract method 0x2d4b0576.
//
// Solidity: function getPositionKey(address _account, address _collateralToken, address _indexToken, bool _isLong) pure returns(bytes32)
func (_Vault *VaultCaller) GetPositionKey(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) ([32]byte, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getPositionKey", _account, _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPositionKey is a free data retrieval call binding the contract method 0x2d4b0576.
//
// Solidity: function getPositionKey(address _account, address _collateralToken, address _indexToken, bool _isLong) pure returns(bytes32)
func (_Vault *VaultSession) GetPositionKey(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) ([32]byte, error) {
	return _Vault.Contract.GetPositionKey(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionKey is a free data retrieval call binding the contract method 0x2d4b0576.
//
// Solidity: function getPositionKey(address _account, address _collateralToken, address _indexToken, bool _isLong) pure returns(bytes32)
func (_Vault *VaultCallerSession) GetPositionKey(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) ([32]byte, error) {
	return _Vault.Contract.GetPositionKey(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionLeverage is a free data retrieval call binding the contract method 0x51723e82.
//
// Solidity: function getPositionLeverage(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_Vault *VaultCaller) GetPositionLeverage(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getPositionLeverage", _account, _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionLeverage is a free data retrieval call binding the contract method 0x51723e82.
//
// Solidity: function getPositionLeverage(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_Vault *VaultSession) GetPositionLeverage(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _Vault.Contract.GetPositionLeverage(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionLeverage is a free data retrieval call binding the contract method 0x51723e82.
//
// Solidity: function getPositionLeverage(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_Vault *VaultCallerSession) GetPositionLeverage(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _Vault.Contract.GetPositionLeverage(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_Vault *VaultCaller) GetRedemptionAmount(opts *bind.CallOpts, _token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getRedemptionAmount", _token, _usdgAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_Vault *VaultSession) GetRedemptionAmount(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetRedemptionAmount(&_Vault.CallOpts, _token, _usdgAmount)
}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_Vault *VaultCallerSession) GetRedemptionAmount(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _Vault.Contract.GetRedemptionAmount(&_Vault.CallOpts, _token, _usdgAmount)
}

// GetRedemptionCollateral is a free data retrieval call binding the contract method 0xb136ca49.
//
// Solidity: function getRedemptionCollateral(address _token) view returns(uint256)
func (_Vault *VaultCaller) GetRedemptionCollateral(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getRedemptionCollateral", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionCollateral is a free data retrieval call binding the contract method 0xb136ca49.
//
// Solidity: function getRedemptionCollateral(address _token) view returns(uint256)
func (_Vault *VaultSession) GetRedemptionCollateral(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetRedemptionCollateral(&_Vault.CallOpts, _token)
}

// GetRedemptionCollateral is a free data retrieval call binding the contract method 0xb136ca49.
//
// Solidity: function getRedemptionCollateral(address _token) view returns(uint256)
func (_Vault *VaultCallerSession) GetRedemptionCollateral(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetRedemptionCollateral(&_Vault.CallOpts, _token)
}

// GetRedemptionCollateralUsd is a free data retrieval call binding the contract method 0x29ff9615.
//
// Solidity: function getRedemptionCollateralUsd(address _token) view returns(uint256)
func (_Vault *VaultCaller) GetRedemptionCollateralUsd(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getRedemptionCollateralUsd", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionCollateralUsd is a free data retrieval call binding the contract method 0x29ff9615.
//
// Solidity: function getRedemptionCollateralUsd(address _token) view returns(uint256)
func (_Vault *VaultSession) GetRedemptionCollateralUsd(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetRedemptionCollateralUsd(&_Vault.CallOpts, _token)
}

// GetRedemptionCollateralUsd is a free data retrieval call binding the contract method 0x29ff9615.
//
// Solidity: function getRedemptionCollateralUsd(address _token) view returns(uint256)
func (_Vault *VaultCallerSession) GetRedemptionCollateralUsd(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetRedemptionCollateralUsd(&_Vault.CallOpts, _token)
}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_Vault *VaultCaller) GetTargetUsdgAmount(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getTargetUsdgAmount", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_Vault *VaultSession) GetTargetUsdgAmount(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetTargetUsdgAmount(&_Vault.CallOpts, _token)
}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_Vault *VaultCallerSession) GetTargetUsdgAmount(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetTargetUsdgAmount(&_Vault.CallOpts, _token)
}

// GetUtilisation is a free data retrieval call binding the contract method 0x04fef1db.
//
// Solidity: function getUtilisation(address _token) view returns(uint256)
func (_Vault *VaultCaller) GetUtilisation(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getUtilisation", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUtilisation is a free data retrieval call binding the contract method 0x04fef1db.
//
// Solidity: function getUtilisation(address _token) view returns(uint256)
func (_Vault *VaultSession) GetUtilisation(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetUtilisation(&_Vault.CallOpts, _token)
}

// GetUtilisation is a free data retrieval call binding the contract method 0x04fef1db.
//
// Solidity: function getUtilisation(address _token) view returns(uint256)
func (_Vault *VaultCallerSession) GetUtilisation(_token common.Address) (*big.Int, error) {
	return _Vault.Contract.GetUtilisation(&_Vault.CallOpts, _token)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_Vault *VaultCaller) GlobalShortAveragePrices(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "globalShortAveragePrices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_Vault *VaultSession) GlobalShortAveragePrices(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.GlobalShortAveragePrices(&_Vault.CallOpts, arg0)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_Vault *VaultCallerSession) GlobalShortAveragePrices(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.GlobalShortAveragePrices(&_Vault.CallOpts, arg0)
}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address ) view returns(uint256)
func (_Vault *VaultCaller) GlobalShortSizes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "globalShortSizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address ) view returns(uint256)
func (_Vault *VaultSession) GlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.GlobalShortSizes(&_Vault.CallOpts, arg0)
}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address ) view returns(uint256)
func (_Vault *VaultCallerSession) GlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.GlobalShortSizes(&_Vault.CallOpts, arg0)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Vault *VaultCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Vault *VaultSession) Gov() (common.Address, error) {
	return _Vault.Contract.Gov(&_Vault.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_Vault *VaultCallerSession) Gov() (common.Address, error) {
	return _Vault.Contract.Gov(&_Vault.CallOpts)
}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address ) view returns(uint256)
func (_Vault *VaultCaller) GuaranteedUsd(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "guaranteedUsd", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address ) view returns(uint256)
func (_Vault *VaultSession) GuaranteedUsd(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.GuaranteedUsd(&_Vault.CallOpts, arg0)
}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address ) view returns(uint256)
func (_Vault *VaultCallerSession) GuaranteedUsd(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.GuaranteedUsd(&_Vault.CallOpts, arg0)
}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_Vault *VaultCaller) HasDynamicFees(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "hasDynamicFees")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_Vault *VaultSession) HasDynamicFees() (bool, error) {
	return _Vault.Contract.HasDynamicFees(&_Vault.CallOpts)
}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_Vault *VaultCallerSession) HasDynamicFees() (bool, error) {
	return _Vault.Contract.HasDynamicFees(&_Vault.CallOpts)
}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_Vault *VaultCaller) InManagerMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "inManagerMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_Vault *VaultSession) InManagerMode() (bool, error) {
	return _Vault.Contract.InManagerMode(&_Vault.CallOpts)
}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_Vault *VaultCallerSession) InManagerMode() (bool, error) {
	return _Vault.Contract.InManagerMode(&_Vault.CallOpts)
}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_Vault *VaultCaller) InPrivateLiquidationMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "inPrivateLiquidationMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_Vault *VaultSession) InPrivateLiquidationMode() (bool, error) {
	return _Vault.Contract.InPrivateLiquidationMode(&_Vault.CallOpts)
}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_Vault *VaultCallerSession) InPrivateLiquidationMode() (bool, error) {
	return _Vault.Contract.InPrivateLiquidationMode(&_Vault.CallOpts)
}

// IncludeAmmPrice is a free data retrieval call binding the contract method 0xab08c1c6.
//
// Solidity: function includeAmmPrice() view returns(bool)
func (_Vault *VaultCaller) IncludeAmmPrice(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "includeAmmPrice")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IncludeAmmPrice is a free data retrieval call binding the contract method 0xab08c1c6.
//
// Solidity: function includeAmmPrice() view returns(bool)
func (_Vault *VaultSession) IncludeAmmPrice() (bool, error) {
	return _Vault.Contract.IncludeAmmPrice(&_Vault.CallOpts)
}

// IncludeAmmPrice is a free data retrieval call binding the contract method 0xab08c1c6.
//
// Solidity: function includeAmmPrice() view returns(bool)
func (_Vault *VaultCallerSession) IncludeAmmPrice() (bool, error) {
	return _Vault.Contract.IncludeAmmPrice(&_Vault.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Vault *VaultCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Vault *VaultSession) IsInitialized() (bool, error) {
	return _Vault.Contract.IsInitialized(&_Vault.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Vault *VaultCallerSession) IsInitialized() (bool, error) {
	return _Vault.Contract.IsInitialized(&_Vault.CallOpts)
}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_Vault *VaultCaller) IsLeverageEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "isLeverageEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_Vault *VaultSession) IsLeverageEnabled() (bool, error) {
	return _Vault.Contract.IsLeverageEnabled(&_Vault.CallOpts)
}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_Vault *VaultCallerSession) IsLeverageEnabled() (bool, error) {
	return _Vault.Contract.IsLeverageEnabled(&_Vault.CallOpts)
}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address ) view returns(bool)
func (_Vault *VaultCaller) IsLiquidator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "isLiquidator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address ) view returns(bool)
func (_Vault *VaultSession) IsLiquidator(arg0 common.Address) (bool, error) {
	return _Vault.Contract.IsLiquidator(&_Vault.CallOpts, arg0)
}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address ) view returns(bool)
func (_Vault *VaultCallerSession) IsLiquidator(arg0 common.Address) (bool, error) {
	return _Vault.Contract.IsLiquidator(&_Vault.CallOpts, arg0)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_Vault *VaultCaller) IsManager(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "isManager", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_Vault *VaultSession) IsManager(arg0 common.Address) (bool, error) {
	return _Vault.Contract.IsManager(&_Vault.CallOpts, arg0)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_Vault *VaultCallerSession) IsManager(arg0 common.Address) (bool, error) {
	return _Vault.Contract.IsManager(&_Vault.CallOpts, arg0)
}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_Vault *VaultCaller) IsSwapEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "isSwapEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_Vault *VaultSession) IsSwapEnabled() (bool, error) {
	return _Vault.Contract.IsSwapEnabled(&_Vault.CallOpts)
}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_Vault *VaultCallerSession) IsSwapEnabled() (bool, error) {
	return _Vault.Contract.IsSwapEnabled(&_Vault.CallOpts)
}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address ) view returns(uint256)
func (_Vault *VaultCaller) LastFundingTimes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "lastFundingTimes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address ) view returns(uint256)
func (_Vault *VaultSession) LastFundingTimes(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.LastFundingTimes(&_Vault.CallOpts, arg0)
}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address ) view returns(uint256)
func (_Vault *VaultCallerSession) LastFundingTimes(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.LastFundingTimes(&_Vault.CallOpts, arg0)
}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_Vault *VaultCaller) LiquidationFeeUsd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "liquidationFeeUsd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_Vault *VaultSession) LiquidationFeeUsd() (*big.Int, error) {
	return _Vault.Contract.LiquidationFeeUsd(&_Vault.CallOpts)
}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_Vault *VaultCallerSession) LiquidationFeeUsd() (*big.Int, error) {
	return _Vault.Contract.LiquidationFeeUsd(&_Vault.CallOpts)
}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_Vault *VaultCaller) MarginFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "marginFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_Vault *VaultSession) MarginFeeBasisPoints() (*big.Int, error) {
	return _Vault.Contract.MarginFeeBasisPoints(&_Vault.CallOpts)
}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_Vault *VaultCallerSession) MarginFeeBasisPoints() (*big.Int, error) {
	return _Vault.Contract.MarginFeeBasisPoints(&_Vault.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_Vault *VaultCaller) MaxGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "maxGasPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_Vault *VaultSession) MaxGasPrice() (*big.Int, error) {
	return _Vault.Contract.MaxGasPrice(&_Vault.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_Vault *VaultCallerSession) MaxGasPrice() (*big.Int, error) {
	return _Vault.Contract.MaxGasPrice(&_Vault.CallOpts)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_Vault *VaultCaller) MaxGlobalShortSizes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "maxGlobalShortSizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_Vault *VaultSession) MaxGlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.MaxGlobalShortSizes(&_Vault.CallOpts, arg0)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_Vault *VaultCallerSession) MaxGlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.MaxGlobalShortSizes(&_Vault.CallOpts, arg0)
}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_Vault *VaultCaller) MaxLeverage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "maxLeverage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_Vault *VaultSession) MaxLeverage() (*big.Int, error) {
	return _Vault.Contract.MaxLeverage(&_Vault.CallOpts)
}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_Vault *VaultCallerSession) MaxLeverage() (*big.Int, error) {
	return _Vault.Contract.MaxLeverage(&_Vault.CallOpts)
}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address ) view returns(uint256)
func (_Vault *VaultCaller) MaxUsdgAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "maxUsdgAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address ) view returns(uint256)
func (_Vault *VaultSession) MaxUsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.MaxUsdgAmounts(&_Vault.CallOpts, arg0)
}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address ) view returns(uint256)
func (_Vault *VaultCallerSession) MaxUsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.MaxUsdgAmounts(&_Vault.CallOpts, arg0)
}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address ) view returns(uint256)
func (_Vault *VaultCaller) MinProfitBasisPoints(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "minProfitBasisPoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address ) view returns(uint256)
func (_Vault *VaultSession) MinProfitBasisPoints(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.MinProfitBasisPoints(&_Vault.CallOpts, arg0)
}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address ) view returns(uint256)
func (_Vault *VaultCallerSession) MinProfitBasisPoints(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.MinProfitBasisPoints(&_Vault.CallOpts, arg0)
}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_Vault *VaultCaller) MinProfitTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "minProfitTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_Vault *VaultSession) MinProfitTime() (*big.Int, error) {
	return _Vault.Contract.MinProfitTime(&_Vault.CallOpts)
}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_Vault *VaultCallerSession) MinProfitTime() (*big.Int, error) {
	return _Vault.Contract.MinProfitTime(&_Vault.CallOpts)
}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_Vault *VaultCaller) MintBurnFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "mintBurnFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_Vault *VaultSession) MintBurnFeeBasisPoints() (*big.Int, error) {
	return _Vault.Contract.MintBurnFeeBasisPoints(&_Vault.CallOpts)
}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_Vault *VaultCallerSession) MintBurnFeeBasisPoints() (*big.Int, error) {
	return _Vault.Contract.MintBurnFeeBasisPoints(&_Vault.CallOpts)
}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address ) view returns(uint256)
func (_Vault *VaultCaller) PoolAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "poolAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address ) view returns(uint256)
func (_Vault *VaultSession) PoolAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.PoolAmounts(&_Vault.CallOpts, arg0)
}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address ) view returns(uint256)
func (_Vault *VaultCallerSession) PoolAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.PoolAmounts(&_Vault.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 lastIncreasedTime)
func (_Vault *VaultCaller) Positions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Size              *big.Int
	Collateral        *big.Int
	AveragePrice      *big.Int
	EntryFundingRate  *big.Int
	ReserveAmount     *big.Int
	RealisedPnl       *big.Int
	LastIncreasedTime *big.Int
}, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "positions", arg0)

	outstruct := new(struct {
		Size              *big.Int
		Collateral        *big.Int
		AveragePrice      *big.Int
		EntryFundingRate  *big.Int
		ReserveAmount     *big.Int
		RealisedPnl       *big.Int
		LastIncreasedTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Size = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Collateral = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AveragePrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EntryFundingRate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ReserveAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.RealisedPnl = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LastIncreasedTime = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 lastIncreasedTime)
func (_Vault *VaultSession) Positions(arg0 [32]byte) (struct {
	Size              *big.Int
	Collateral        *big.Int
	AveragePrice      *big.Int
	EntryFundingRate  *big.Int
	ReserveAmount     *big.Int
	RealisedPnl       *big.Int
	LastIncreasedTime *big.Int
}, error) {
	return _Vault.Contract.Positions(&_Vault.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 lastIncreasedTime)
func (_Vault *VaultCallerSession) Positions(arg0 [32]byte) (struct {
	Size              *big.Int
	Collateral        *big.Int
	AveragePrice      *big.Int
	EntryFundingRate  *big.Int
	ReserveAmount     *big.Int
	RealisedPnl       *big.Int
	LastIncreasedTime *big.Int
}, error) {
	return _Vault.Contract.Positions(&_Vault.CallOpts, arg0)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_Vault *VaultCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_Vault *VaultSession) PriceFeed() (common.Address, error) {
	return _Vault.Contract.PriceFeed(&_Vault.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_Vault *VaultCallerSession) PriceFeed() (common.Address, error) {
	return _Vault.Contract.PriceFeed(&_Vault.CallOpts)
}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address ) view returns(uint256)
func (_Vault *VaultCaller) ReservedAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "reservedAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address ) view returns(uint256)
func (_Vault *VaultSession) ReservedAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.ReservedAmounts(&_Vault.CallOpts, arg0)
}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address ) view returns(uint256)
func (_Vault *VaultCallerSession) ReservedAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.ReservedAmounts(&_Vault.CallOpts, arg0)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Vault *VaultCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Vault *VaultSession) Router() (common.Address, error) {
	return _Vault.Contract.Router(&_Vault.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_Vault *VaultCallerSession) Router() (common.Address, error) {
	return _Vault.Contract.Router(&_Vault.CallOpts)
}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address ) view returns(bool)
func (_Vault *VaultCaller) ShortableTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "shortableTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address ) view returns(bool)
func (_Vault *VaultSession) ShortableTokens(arg0 common.Address) (bool, error) {
	return _Vault.Contract.ShortableTokens(&_Vault.CallOpts, arg0)
}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address ) view returns(bool)
func (_Vault *VaultCallerSession) ShortableTokens(arg0 common.Address) (bool, error) {
	return _Vault.Contract.ShortableTokens(&_Vault.CallOpts, arg0)
}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_Vault *VaultCaller) StableFundingRateFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "stableFundingRateFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_Vault *VaultSession) StableFundingRateFactor() (*big.Int, error) {
	return _Vault.Contract.StableFundingRateFactor(&_Vault.CallOpts)
}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_Vault *VaultCallerSession) StableFundingRateFactor() (*big.Int, error) {
	return _Vault.Contract.StableFundingRateFactor(&_Vault.CallOpts)
}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_Vault *VaultCaller) StableSwapFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "stableSwapFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_Vault *VaultSession) StableSwapFeeBasisPoints() (*big.Int, error) {
	return _Vault.Contract.StableSwapFeeBasisPoints(&_Vault.CallOpts)
}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_Vault *VaultCallerSession) StableSwapFeeBasisPoints() (*big.Int, error) {
	return _Vault.Contract.StableSwapFeeBasisPoints(&_Vault.CallOpts)
}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_Vault *VaultCaller) StableTaxBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "stableTaxBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_Vault *VaultSession) StableTaxBasisPoints() (*big.Int, error) {
	return _Vault.Contract.StableTaxBasisPoints(&_Vault.CallOpts)
}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_Vault *VaultCallerSession) StableTaxBasisPoints() (*big.Int, error) {
	return _Vault.Contract.StableTaxBasisPoints(&_Vault.CallOpts)
}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address ) view returns(bool)
func (_Vault *VaultCaller) StableTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "stableTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address ) view returns(bool)
func (_Vault *VaultSession) StableTokens(arg0 common.Address) (bool, error) {
	return _Vault.Contract.StableTokens(&_Vault.CallOpts, arg0)
}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address ) view returns(bool)
func (_Vault *VaultCallerSession) StableTokens(arg0 common.Address) (bool, error) {
	return _Vault.Contract.StableTokens(&_Vault.CallOpts, arg0)
}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_Vault *VaultCaller) SwapFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "swapFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_Vault *VaultSession) SwapFeeBasisPoints() (*big.Int, error) {
	return _Vault.Contract.SwapFeeBasisPoints(&_Vault.CallOpts)
}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_Vault *VaultCallerSession) SwapFeeBasisPoints() (*big.Int, error) {
	return _Vault.Contract.SwapFeeBasisPoints(&_Vault.CallOpts)
}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_Vault *VaultCaller) TaxBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "taxBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_Vault *VaultSession) TaxBasisPoints() (*big.Int, error) {
	return _Vault.Contract.TaxBasisPoints(&_Vault.CallOpts)
}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_Vault *VaultCallerSession) TaxBasisPoints() (*big.Int, error) {
	return _Vault.Contract.TaxBasisPoints(&_Vault.CallOpts)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_Vault *VaultCaller) TokenBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "tokenBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_Vault *VaultSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TokenBalances(&_Vault.CallOpts, arg0)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_Vault *VaultCallerSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TokenBalances(&_Vault.CallOpts, arg0)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint256)
func (_Vault *VaultCaller) TokenDecimals(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "tokenDecimals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint256)
func (_Vault *VaultSession) TokenDecimals(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TokenDecimals(&_Vault.CallOpts, arg0)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint256)
func (_Vault *VaultCallerSession) TokenDecimals(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TokenDecimals(&_Vault.CallOpts, arg0)
}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_Vault *VaultCaller) TokenToUsdMin(opts *bind.CallOpts, _token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "tokenToUsdMin", _token, _tokenAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_Vault *VaultSession) TokenToUsdMin(_token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	return _Vault.Contract.TokenToUsdMin(&_Vault.CallOpts, _token, _tokenAmount)
}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_Vault *VaultCallerSession) TokenToUsdMin(_token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	return _Vault.Contract.TokenToUsdMin(&_Vault.CallOpts, _token, _tokenAmount)
}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address ) view returns(uint256)
func (_Vault *VaultCaller) TokenWeights(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "tokenWeights", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address ) view returns(uint256)
func (_Vault *VaultSession) TokenWeights(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TokenWeights(&_Vault.CallOpts, arg0)
}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address ) view returns(uint256)
func (_Vault *VaultCallerSession) TokenWeights(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TokenWeights(&_Vault.CallOpts, arg0)
}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_Vault *VaultCaller) TotalTokenWeights(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "totalTokenWeights")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_Vault *VaultSession) TotalTokenWeights() (*big.Int, error) {
	return _Vault.Contract.TotalTokenWeights(&_Vault.CallOpts)
}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_Vault *VaultCallerSession) TotalTokenWeights() (*big.Int, error) {
	return _Vault.Contract.TotalTokenWeights(&_Vault.CallOpts)
}

// UsdToToken is a free data retrieval call binding the contract method 0xfa12dbc0.
//
// Solidity: function usdToToken(address _token, uint256 _usdAmount, uint256 _price) view returns(uint256)
func (_Vault *VaultCaller) UsdToToken(opts *bind.CallOpts, _token common.Address, _usdAmount *big.Int, _price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "usdToToken", _token, _usdAmount, _price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdToToken is a free data retrieval call binding the contract method 0xfa12dbc0.
//
// Solidity: function usdToToken(address _token, uint256 _usdAmount, uint256 _price) view returns(uint256)
func (_Vault *VaultSession) UsdToToken(_token common.Address, _usdAmount *big.Int, _price *big.Int) (*big.Int, error) {
	return _Vault.Contract.UsdToToken(&_Vault.CallOpts, _token, _usdAmount, _price)
}

// UsdToToken is a free data retrieval call binding the contract method 0xfa12dbc0.
//
// Solidity: function usdToToken(address _token, uint256 _usdAmount, uint256 _price) view returns(uint256)
func (_Vault *VaultCallerSession) UsdToToken(_token common.Address, _usdAmount *big.Int, _price *big.Int) (*big.Int, error) {
	return _Vault.Contract.UsdToToken(&_Vault.CallOpts, _token, _usdAmount, _price)
}

// UsdToTokenMax is a free data retrieval call binding the contract method 0xa42ab3d2.
//
// Solidity: function usdToTokenMax(address _token, uint256 _usdAmount) view returns(uint256)
func (_Vault *VaultCaller) UsdToTokenMax(opts *bind.CallOpts, _token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "usdToTokenMax", _token, _usdAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdToTokenMax is a free data retrieval call binding the contract method 0xa42ab3d2.
//
// Solidity: function usdToTokenMax(address _token, uint256 _usdAmount) view returns(uint256)
func (_Vault *VaultSession) UsdToTokenMax(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _Vault.Contract.UsdToTokenMax(&_Vault.CallOpts, _token, _usdAmount)
}

// UsdToTokenMax is a free data retrieval call binding the contract method 0xa42ab3d2.
//
// Solidity: function usdToTokenMax(address _token, uint256 _usdAmount) view returns(uint256)
func (_Vault *VaultCallerSession) UsdToTokenMax(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _Vault.Contract.UsdToTokenMax(&_Vault.CallOpts, _token, _usdAmount)
}

// UsdToTokenMin is a free data retrieval call binding the contract method 0x9899cd02.
//
// Solidity: function usdToTokenMin(address _token, uint256 _usdAmount) view returns(uint256)
func (_Vault *VaultCaller) UsdToTokenMin(opts *bind.CallOpts, _token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "usdToTokenMin", _token, _usdAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdToTokenMin is a free data retrieval call binding the contract method 0x9899cd02.
//
// Solidity: function usdToTokenMin(address _token, uint256 _usdAmount) view returns(uint256)
func (_Vault *VaultSession) UsdToTokenMin(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _Vault.Contract.UsdToTokenMin(&_Vault.CallOpts, _token, _usdAmount)
}

// UsdToTokenMin is a free data retrieval call binding the contract method 0x9899cd02.
//
// Solidity: function usdToTokenMin(address _token, uint256 _usdAmount) view returns(uint256)
func (_Vault *VaultCallerSession) UsdToTokenMin(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _Vault.Contract.UsdToTokenMin(&_Vault.CallOpts, _token, _usdAmount)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_Vault *VaultCaller) Usdg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "usdg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_Vault *VaultSession) Usdg() (common.Address, error) {
	return _Vault.Contract.Usdg(&_Vault.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_Vault *VaultCallerSession) Usdg() (common.Address, error) {
	return _Vault.Contract.Usdg(&_Vault.CallOpts)
}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address ) view returns(uint256)
func (_Vault *VaultCaller) UsdgAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "usdgAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address ) view returns(uint256)
func (_Vault *VaultSession) UsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.UsdgAmounts(&_Vault.CallOpts, arg0)
}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address ) view returns(uint256)
func (_Vault *VaultCallerSession) UsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.UsdgAmounts(&_Vault.CallOpts, arg0)
}

// UseSwapPricing is a free data retrieval call binding the contract method 0xb06423f3.
//
// Solidity: function useSwapPricing() view returns(bool)
func (_Vault *VaultCaller) UseSwapPricing(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "useSwapPricing")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseSwapPricing is a free data retrieval call binding the contract method 0xb06423f3.
//
// Solidity: function useSwapPricing() view returns(bool)
func (_Vault *VaultSession) UseSwapPricing() (bool, error) {
	return _Vault.Contract.UseSwapPricing(&_Vault.CallOpts)
}

// UseSwapPricing is a free data retrieval call binding the contract method 0xb06423f3.
//
// Solidity: function useSwapPricing() view returns(bool)
func (_Vault *VaultCallerSession) UseSwapPricing() (bool, error) {
	return _Vault.Contract.UseSwapPricing(&_Vault.CallOpts)
}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_Vault *VaultCaller) ValidateLiquidation(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "validateLiquidation", _account, _collateralToken, _indexToken, _isLong, _raise)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_Vault *VaultSession) ValidateLiquidation(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	return _Vault.Contract.ValidateLiquidation(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong, _raise)
}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_Vault *VaultCallerSession) ValidateLiquidation(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	return _Vault.Contract.ValidateLiquidation(&_Vault.CallOpts, _account, _collateralToken, _indexToken, _isLong, _raise)
}

// VaultUtils is a free data retrieval call binding the contract method 0x6abbe0c8.
//
// Solidity: function vaultUtils() view returns(address)
func (_Vault *VaultCaller) VaultUtils(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "vaultUtils")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultUtils is a free data retrieval call binding the contract method 0x6abbe0c8.
//
// Solidity: function vaultUtils() view returns(address)
func (_Vault *VaultSession) VaultUtils() (common.Address, error) {
	return _Vault.Contract.VaultUtils(&_Vault.CallOpts)
}

// VaultUtils is a free data retrieval call binding the contract method 0x6abbe0c8.
//
// Solidity: function vaultUtils() view returns(address)
func (_Vault *VaultCallerSession) VaultUtils() (common.Address, error) {
	return _Vault.Contract.VaultUtils(&_Vault.CallOpts)
}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_Vault *VaultCaller) WhitelistedTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "whitelistedTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_Vault *VaultSession) WhitelistedTokenCount() (*big.Int, error) {
	return _Vault.Contract.WhitelistedTokenCount(&_Vault.CallOpts)
}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_Vault *VaultCallerSession) WhitelistedTokenCount() (*big.Int, error) {
	return _Vault.Contract.WhitelistedTokenCount(&_Vault.CallOpts)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_Vault *VaultCaller) WhitelistedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "whitelistedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_Vault *VaultSession) WhitelistedTokens(arg0 common.Address) (bool, error) {
	return _Vault.Contract.WhitelistedTokens(&_Vault.CallOpts, arg0)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_Vault *VaultCallerSession) WhitelistedTokens(arg0 common.Address) (bool, error) {
	return _Vault.Contract.WhitelistedTokens(&_Vault.CallOpts, arg0)
}

// AddRouter is a paid mutator transaction binding the contract method 0x24ca984e.
//
// Solidity: function addRouter(address _router) returns()
func (_Vault *VaultTransactor) AddRouter(opts *bind.TransactOpts, _router common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "addRouter", _router)
}

// AddRouter is a paid mutator transaction binding the contract method 0x24ca984e.
//
// Solidity: function addRouter(address _router) returns()
func (_Vault *VaultSession) AddRouter(_router common.Address) (*types.Transaction, error) {
	return _Vault.Contract.AddRouter(&_Vault.TransactOpts, _router)
}

// AddRouter is a paid mutator transaction binding the contract method 0x24ca984e.
//
// Solidity: function addRouter(address _router) returns()
func (_Vault *VaultTransactorSession) AddRouter(_router common.Address) (*types.Transaction, error) {
	return _Vault.Contract.AddRouter(&_Vault.TransactOpts, _router)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_Vault *VaultTransactor) BuyUSDG(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "buyUSDG", _token, _receiver)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_Vault *VaultSession) BuyUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.BuyUSDG(&_Vault.TransactOpts, _token, _receiver)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_Vault *VaultTransactorSession) BuyUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.BuyUSDG(&_Vault.TransactOpts, _token, _receiver)
}

// ClearTokenConfig is a paid mutator transaction binding the contract method 0xe67f59a7.
//
// Solidity: function clearTokenConfig(address _token) returns()
func (_Vault *VaultTransactor) ClearTokenConfig(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "clearTokenConfig", _token)
}

// ClearTokenConfig is a paid mutator transaction binding the contract method 0xe67f59a7.
//
// Solidity: function clearTokenConfig(address _token) returns()
func (_Vault *VaultSession) ClearTokenConfig(_token common.Address) (*types.Transaction, error) {
	return _Vault.Contract.ClearTokenConfig(&_Vault.TransactOpts, _token)
}

// ClearTokenConfig is a paid mutator transaction binding the contract method 0xe67f59a7.
//
// Solidity: function clearTokenConfig(address _token) returns()
func (_Vault *VaultTransactorSession) ClearTokenConfig(_token common.Address) (*types.Transaction, error) {
	return _Vault.Contract.ClearTokenConfig(&_Vault.TransactOpts, _token)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_Vault *VaultTransactor) DecreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "decreasePosition", _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_Vault *VaultSession) DecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.DecreasePosition(&_Vault.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_Vault *VaultTransactorSession) DecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.DecreasePosition(&_Vault.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_Vault *VaultTransactor) DirectPoolDeposit(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "directPoolDeposit", _token)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_Vault *VaultSession) DirectPoolDeposit(_token common.Address) (*types.Transaction, error) {
	return _Vault.Contract.DirectPoolDeposit(&_Vault.TransactOpts, _token)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_Vault *VaultTransactorSession) DirectPoolDeposit(_token common.Address) (*types.Transaction, error) {
	return _Vault.Contract.DirectPoolDeposit(&_Vault.TransactOpts, _token)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_Vault *VaultTransactor) IncreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "increasePosition", _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_Vault *VaultSession) IncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _Vault.Contract.IncreasePosition(&_Vault.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_Vault *VaultTransactorSession) IncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _Vault.Contract.IncreasePosition(&_Vault.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _router, address _usdg, address _priceFeed, uint256 _liquidationFeeUsd, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_Vault *VaultTransactor) Initialize(opts *bind.TransactOpts, _router common.Address, _usdg common.Address, _priceFeed common.Address, _liquidationFeeUsd *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "initialize", _router, _usdg, _priceFeed, _liquidationFeeUsd, _fundingRateFactor, _stableFundingRateFactor)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _router, address _usdg, address _priceFeed, uint256 _liquidationFeeUsd, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_Vault *VaultSession) Initialize(_router common.Address, _usdg common.Address, _priceFeed common.Address, _liquidationFeeUsd *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Initialize(&_Vault.TransactOpts, _router, _usdg, _priceFeed, _liquidationFeeUsd, _fundingRateFactor, _stableFundingRateFactor)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _router, address _usdg, address _priceFeed, uint256 _liquidationFeeUsd, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_Vault *VaultTransactorSession) Initialize(_router common.Address, _usdg common.Address, _priceFeed common.Address, _liquidationFeeUsd *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Initialize(&_Vault.TransactOpts, _router, _usdg, _priceFeed, _liquidationFeeUsd, _fundingRateFactor, _stableFundingRateFactor)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_Vault *VaultTransactor) LiquidatePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "liquidatePosition", _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_Vault *VaultSession) LiquidatePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.LiquidatePosition(&_Vault.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_Vault *VaultTransactorSession) LiquidatePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.LiquidatePosition(&_Vault.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// RemoveRouter is a paid mutator transaction binding the contract method 0x6ae0b154.
//
// Solidity: function removeRouter(address _router) returns()
func (_Vault *VaultTransactor) RemoveRouter(opts *bind.TransactOpts, _router common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "removeRouter", _router)
}

// RemoveRouter is a paid mutator transaction binding the contract method 0x6ae0b154.
//
// Solidity: function removeRouter(address _router) returns()
func (_Vault *VaultSession) RemoveRouter(_router common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RemoveRouter(&_Vault.TransactOpts, _router)
}

// RemoveRouter is a paid mutator transaction binding the contract method 0x6ae0b154.
//
// Solidity: function removeRouter(address _router) returns()
func (_Vault *VaultTransactorSession) RemoveRouter(_router common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RemoveRouter(&_Vault.TransactOpts, _router)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_Vault *VaultTransactor) SellUSDG(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "sellUSDG", _token, _receiver)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_Vault *VaultSession) SellUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SellUSDG(&_Vault.TransactOpts, _token, _receiver)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_Vault *VaultTransactorSession) SellUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SellUSDG(&_Vault.TransactOpts, _token, _receiver)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_Vault *VaultTransactor) SetBufferAmount(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setBufferAmount", _token, _amount)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_Vault *VaultSession) SetBufferAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetBufferAmount(&_Vault.TransactOpts, _token, _amount)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_Vault *VaultTransactorSession) SetBufferAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetBufferAmount(&_Vault.TransactOpts, _token, _amount)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_Vault *VaultTransactor) SetError(opts *bind.TransactOpts, _errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setError", _errorCode, _error)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_Vault *VaultSession) SetError(_errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _Vault.Contract.SetError(&_Vault.TransactOpts, _errorCode, _error)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_Vault *VaultTransactorSession) SetError(_errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _Vault.Contract.SetError(&_Vault.TransactOpts, _errorCode, _error)
}

// SetErrorController is a paid mutator transaction binding the contract method 0x8f7b8404.
//
// Solidity: function setErrorController(address _errorController) returns()
func (_Vault *VaultTransactor) SetErrorController(opts *bind.TransactOpts, _errorController common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setErrorController", _errorController)
}

// SetErrorController is a paid mutator transaction binding the contract method 0x8f7b8404.
//
// Solidity: function setErrorController(address _errorController) returns()
func (_Vault *VaultSession) SetErrorController(_errorController common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetErrorController(&_Vault.TransactOpts, _errorController)
}

// SetErrorController is a paid mutator transaction binding the contract method 0x8f7b8404.
//
// Solidity: function setErrorController(address _errorController) returns()
func (_Vault *VaultTransactorSession) SetErrorController(_errorController common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetErrorController(&_Vault.TransactOpts, _errorController)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_Vault *VaultTransactor) SetFees(opts *bind.TransactOpts, _taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setFees", _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_Vault *VaultSession) SetFees(_taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _Vault.Contract.SetFees(&_Vault.TransactOpts, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_Vault *VaultTransactorSession) SetFees(_taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _Vault.Contract.SetFees(&_Vault.TransactOpts, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_Vault *VaultTransactor) SetFundingRate(opts *bind.TransactOpts, _fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setFundingRate", _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_Vault *VaultSession) SetFundingRate(_fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetFundingRate(&_Vault.TransactOpts, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_Vault *VaultTransactorSession) SetFundingRate(_fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetFundingRate(&_Vault.TransactOpts, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Vault *VaultTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Vault *VaultSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetGov(&_Vault.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_Vault *VaultTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetGov(&_Vault.TransactOpts, _gov)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_Vault *VaultTransactor) SetInManagerMode(opts *bind.TransactOpts, _inManagerMode bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setInManagerMode", _inManagerMode)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_Vault *VaultSession) SetInManagerMode(_inManagerMode bool) (*types.Transaction, error) {
	return _Vault.Contract.SetInManagerMode(&_Vault.TransactOpts, _inManagerMode)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_Vault *VaultTransactorSession) SetInManagerMode(_inManagerMode bool) (*types.Transaction, error) {
	return _Vault.Contract.SetInManagerMode(&_Vault.TransactOpts, _inManagerMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_Vault *VaultTransactor) SetInPrivateLiquidationMode(opts *bind.TransactOpts, _inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setInPrivateLiquidationMode", _inPrivateLiquidationMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_Vault *VaultSession) SetInPrivateLiquidationMode(_inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _Vault.Contract.SetInPrivateLiquidationMode(&_Vault.TransactOpts, _inPrivateLiquidationMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_Vault *VaultTransactorSession) SetInPrivateLiquidationMode(_inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _Vault.Contract.SetInPrivateLiquidationMode(&_Vault.TransactOpts, _inPrivateLiquidationMode)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_Vault *VaultTransactor) SetIsLeverageEnabled(opts *bind.TransactOpts, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setIsLeverageEnabled", _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_Vault *VaultSession) SetIsLeverageEnabled(_isLeverageEnabled bool) (*types.Transaction, error) {
	return _Vault.Contract.SetIsLeverageEnabled(&_Vault.TransactOpts, _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_Vault *VaultTransactorSession) SetIsLeverageEnabled(_isLeverageEnabled bool) (*types.Transaction, error) {
	return _Vault.Contract.SetIsLeverageEnabled(&_Vault.TransactOpts, _isLeverageEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_Vault *VaultTransactor) SetIsSwapEnabled(opts *bind.TransactOpts, _isSwapEnabled bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setIsSwapEnabled", _isSwapEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_Vault *VaultSession) SetIsSwapEnabled(_isSwapEnabled bool) (*types.Transaction, error) {
	return _Vault.Contract.SetIsSwapEnabled(&_Vault.TransactOpts, _isSwapEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_Vault *VaultTransactorSession) SetIsSwapEnabled(_isSwapEnabled bool) (*types.Transaction, error) {
	return _Vault.Contract.SetIsSwapEnabled(&_Vault.TransactOpts, _isSwapEnabled)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_Vault *VaultTransactor) SetLiquidator(opts *bind.TransactOpts, _liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setLiquidator", _liquidator, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_Vault *VaultSession) SetLiquidator(_liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _Vault.Contract.SetLiquidator(&_Vault.TransactOpts, _liquidator, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_Vault *VaultTransactorSession) SetLiquidator(_liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _Vault.Contract.SetLiquidator(&_Vault.TransactOpts, _liquidator, _isActive)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_Vault *VaultTransactor) SetManager(opts *bind.TransactOpts, _manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setManager", _manager, _isManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_Vault *VaultSession) SetManager(_manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _Vault.Contract.SetManager(&_Vault.TransactOpts, _manager, _isManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_Vault *VaultTransactorSession) SetManager(_manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _Vault.Contract.SetManager(&_Vault.TransactOpts, _manager, _isManager)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_Vault *VaultTransactor) SetMaxGasPrice(opts *bind.TransactOpts, _maxGasPrice *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setMaxGasPrice", _maxGasPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_Vault *VaultSession) SetMaxGasPrice(_maxGasPrice *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetMaxGasPrice(&_Vault.TransactOpts, _maxGasPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_Vault *VaultTransactorSession) SetMaxGasPrice(_maxGasPrice *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetMaxGasPrice(&_Vault.TransactOpts, _maxGasPrice)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_Vault *VaultTransactor) SetMaxGlobalShortSize(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setMaxGlobalShortSize", _token, _amount)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_Vault *VaultSession) SetMaxGlobalShortSize(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetMaxGlobalShortSize(&_Vault.TransactOpts, _token, _amount)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_Vault *VaultTransactorSession) SetMaxGlobalShortSize(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetMaxGlobalShortSize(&_Vault.TransactOpts, _token, _amount)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_Vault *VaultTransactor) SetMaxLeverage(opts *bind.TransactOpts, _maxLeverage *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setMaxLeverage", _maxLeverage)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_Vault *VaultSession) SetMaxLeverage(_maxLeverage *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetMaxLeverage(&_Vault.TransactOpts, _maxLeverage)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_Vault *VaultTransactorSession) SetMaxLeverage(_maxLeverage *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetMaxLeverage(&_Vault.TransactOpts, _maxLeverage)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_Vault *VaultTransactor) SetPriceFeed(opts *bind.TransactOpts, _priceFeed common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setPriceFeed", _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_Vault *VaultSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetPriceFeed(&_Vault.TransactOpts, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_Vault *VaultTransactorSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetPriceFeed(&_Vault.TransactOpts, _priceFeed)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_Vault *VaultTransactor) SetTokenConfig(opts *bind.TransactOpts, _token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setTokenConfig", _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_Vault *VaultSession) SetTokenConfig(_token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _Vault.Contract.SetTokenConfig(&_Vault.TransactOpts, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_Vault *VaultTransactorSession) SetTokenConfig(_token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _Vault.Contract.SetTokenConfig(&_Vault.TransactOpts, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetUsdgAmount is a paid mutator transaction binding the contract method 0xd66b000d.
//
// Solidity: function setUsdgAmount(address _token, uint256 _amount) returns()
func (_Vault *VaultTransactor) SetUsdgAmount(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setUsdgAmount", _token, _amount)
}

// SetUsdgAmount is a paid mutator transaction binding the contract method 0xd66b000d.
//
// Solidity: function setUsdgAmount(address _token, uint256 _amount) returns()
func (_Vault *VaultSession) SetUsdgAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetUsdgAmount(&_Vault.TransactOpts, _token, _amount)
}

// SetUsdgAmount is a paid mutator transaction binding the contract method 0xd66b000d.
//
// Solidity: function setUsdgAmount(address _token, uint256 _amount) returns()
func (_Vault *VaultTransactorSession) SetUsdgAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetUsdgAmount(&_Vault.TransactOpts, _token, _amount)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_Vault *VaultTransactor) SetVaultUtils(opts *bind.TransactOpts, _vaultUtils common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setVaultUtils", _vaultUtils)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_Vault *VaultSession) SetVaultUtils(_vaultUtils common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetVaultUtils(&_Vault.TransactOpts, _vaultUtils)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_Vault *VaultTransactorSession) SetVaultUtils(_vaultUtils common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetVaultUtils(&_Vault.TransactOpts, _vaultUtils)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_Vault *VaultTransactor) Swap(opts *bind.TransactOpts, _tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "swap", _tokenIn, _tokenOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_Vault *VaultSession) Swap(_tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Swap(&_Vault.TransactOpts, _tokenIn, _tokenOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_Vault *VaultTransactorSession) Swap(_tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Swap(&_Vault.TransactOpts, _tokenIn, _tokenOut, _receiver)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns()
func (_Vault *VaultTransactor) UpdateCumulativeFundingRate(opts *bind.TransactOpts, _collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "updateCumulativeFundingRate", _collateralToken, _indexToken)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns()
func (_Vault *VaultSession) UpdateCumulativeFundingRate(_collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _Vault.Contract.UpdateCumulativeFundingRate(&_Vault.TransactOpts, _collateralToken, _indexToken)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns()
func (_Vault *VaultTransactorSession) UpdateCumulativeFundingRate(_collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _Vault.Contract.UpdateCumulativeFundingRate(&_Vault.TransactOpts, _collateralToken, _indexToken)
}

// UpgradeVault is a paid mutator transaction binding the contract method 0xcea0c328.
//
// Solidity: function upgradeVault(address _newVault, address _token, uint256 _amount) returns()
func (_Vault *VaultTransactor) UpgradeVault(opts *bind.TransactOpts, _newVault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "upgradeVault", _newVault, _token, _amount)
}

// UpgradeVault is a paid mutator transaction binding the contract method 0xcea0c328.
//
// Solidity: function upgradeVault(address _newVault, address _token, uint256 _amount) returns()
func (_Vault *VaultSession) UpgradeVault(_newVault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.UpgradeVault(&_Vault.TransactOpts, _newVault, _token, _amount)
}

// UpgradeVault is a paid mutator transaction binding the contract method 0xcea0c328.
//
// Solidity: function upgradeVault(address _newVault, address _token, uint256 _amount) returns()
func (_Vault *VaultTransactorSession) UpgradeVault(_newVault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.UpgradeVault(&_Vault.TransactOpts, _newVault, _token, _amount)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_Vault *VaultTransactor) WithdrawFees(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "withdrawFees", _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_Vault *VaultSession) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.WithdrawFees(&_Vault.TransactOpts, _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_Vault *VaultTransactorSession) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Vault.Contract.WithdrawFees(&_Vault.TransactOpts, _token, _receiver)
}

// VaultBuyUSDGIterator is returned from FilterBuyUSDG and is used to iterate over the raw logs and unpacked data for BuyUSDG events raised by the Vault contract.
type VaultBuyUSDGIterator struct {
	Event *VaultBuyUSDG // Event containing the contract specifics and raw log

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
func (it *VaultBuyUSDGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultBuyUSDG)
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
		it.Event = new(VaultBuyUSDG)
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
func (it *VaultBuyUSDGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultBuyUSDGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultBuyUSDG represents a BuyUSDG event raised by the Vault contract.
type VaultBuyUSDG struct {
	Account        common.Address
	Token          common.Address
	TokenAmount    *big.Int
	UsdgAmount     *big.Int
	FeeBasisPoints *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBuyUSDG is a free log retrieval operation binding the contract event 0xab4c77c74cd32c85f35416cf03e7ce9e2d4387f7b7f2c1f4bf53daaecf8ea72d.
//
// Solidity: event BuyUSDG(address account, address token, uint256 tokenAmount, uint256 usdgAmount, uint256 feeBasisPoints)
func (_Vault *VaultFilterer) FilterBuyUSDG(opts *bind.FilterOpts) (*VaultBuyUSDGIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "BuyUSDG")
	if err != nil {
		return nil, err
	}
	return &VaultBuyUSDGIterator{contract: _Vault.contract, event: "BuyUSDG", logs: logs, sub: sub}, nil
}

// WatchBuyUSDG is a free log subscription operation binding the contract event 0xab4c77c74cd32c85f35416cf03e7ce9e2d4387f7b7f2c1f4bf53daaecf8ea72d.
//
// Solidity: event BuyUSDG(address account, address token, uint256 tokenAmount, uint256 usdgAmount, uint256 feeBasisPoints)
func (_Vault *VaultFilterer) WatchBuyUSDG(opts *bind.WatchOpts, sink chan<- *VaultBuyUSDG) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "BuyUSDG")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultBuyUSDG)
				if err := _Vault.contract.UnpackLog(event, "BuyUSDG", log); err != nil {
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

// ParseBuyUSDG is a log parse operation binding the contract event 0xab4c77c74cd32c85f35416cf03e7ce9e2d4387f7b7f2c1f4bf53daaecf8ea72d.
//
// Solidity: event BuyUSDG(address account, address token, uint256 tokenAmount, uint256 usdgAmount, uint256 feeBasisPoints)
func (_Vault *VaultFilterer) ParseBuyUSDG(log types.Log) (*VaultBuyUSDG, error) {
	event := new(VaultBuyUSDG)
	if err := _Vault.contract.UnpackLog(event, "BuyUSDG", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultClosePositionIterator is returned from FilterClosePosition and is used to iterate over the raw logs and unpacked data for ClosePosition events raised by the Vault contract.
type VaultClosePositionIterator struct {
	Event *VaultClosePosition // Event containing the contract specifics and raw log

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
func (it *VaultClosePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultClosePosition)
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
		it.Event = new(VaultClosePosition)
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
func (it *VaultClosePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultClosePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultClosePosition represents a ClosePosition event raised by the Vault contract.
type VaultClosePosition struct {
	Key              [32]byte
	Size             *big.Int
	Collateral       *big.Int
	AveragePrice     *big.Int
	EntryFundingRate *big.Int
	ReserveAmount    *big.Int
	RealisedPnl      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterClosePosition is a free log retrieval operation binding the contract event 0x73af1d417d82c240fdb6d319b34ad884487c6bf2845d98980cc52ad9171cb455.
//
// Solidity: event ClosePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl)
func (_Vault *VaultFilterer) FilterClosePosition(opts *bind.FilterOpts) (*VaultClosePositionIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "ClosePosition")
	if err != nil {
		return nil, err
	}
	return &VaultClosePositionIterator{contract: _Vault.contract, event: "ClosePosition", logs: logs, sub: sub}, nil
}

// WatchClosePosition is a free log subscription operation binding the contract event 0x73af1d417d82c240fdb6d319b34ad884487c6bf2845d98980cc52ad9171cb455.
//
// Solidity: event ClosePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl)
func (_Vault *VaultFilterer) WatchClosePosition(opts *bind.WatchOpts, sink chan<- *VaultClosePosition) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "ClosePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultClosePosition)
				if err := _Vault.contract.UnpackLog(event, "ClosePosition", log); err != nil {
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

// ParseClosePosition is a log parse operation binding the contract event 0x73af1d417d82c240fdb6d319b34ad884487c6bf2845d98980cc52ad9171cb455.
//
// Solidity: event ClosePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl)
func (_Vault *VaultFilterer) ParseClosePosition(log types.Log) (*VaultClosePosition, error) {
	event := new(VaultClosePosition)
	if err := _Vault.contract.UnpackLog(event, "ClosePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultCollectMarginFeesIterator is returned from FilterCollectMarginFees and is used to iterate over the raw logs and unpacked data for CollectMarginFees events raised by the Vault contract.
type VaultCollectMarginFeesIterator struct {
	Event *VaultCollectMarginFees // Event containing the contract specifics and raw log

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
func (it *VaultCollectMarginFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultCollectMarginFees)
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
		it.Event = new(VaultCollectMarginFees)
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
func (it *VaultCollectMarginFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultCollectMarginFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultCollectMarginFees represents a CollectMarginFees event raised by the Vault contract.
type VaultCollectMarginFees struct {
	Token     common.Address
	FeeUsd    *big.Int
	FeeTokens *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectMarginFees is a free log retrieval operation binding the contract event 0x5d0c0019d3d45fadeb74eff9d2c9924d146d000ac6bcf3c28bf0ac3c9baa011a.
//
// Solidity: event CollectMarginFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_Vault *VaultFilterer) FilterCollectMarginFees(opts *bind.FilterOpts) (*VaultCollectMarginFeesIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "CollectMarginFees")
	if err != nil {
		return nil, err
	}
	return &VaultCollectMarginFeesIterator{contract: _Vault.contract, event: "CollectMarginFees", logs: logs, sub: sub}, nil
}

// WatchCollectMarginFees is a free log subscription operation binding the contract event 0x5d0c0019d3d45fadeb74eff9d2c9924d146d000ac6bcf3c28bf0ac3c9baa011a.
//
// Solidity: event CollectMarginFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_Vault *VaultFilterer) WatchCollectMarginFees(opts *bind.WatchOpts, sink chan<- *VaultCollectMarginFees) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "CollectMarginFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultCollectMarginFees)
				if err := _Vault.contract.UnpackLog(event, "CollectMarginFees", log); err != nil {
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

// ParseCollectMarginFees is a log parse operation binding the contract event 0x5d0c0019d3d45fadeb74eff9d2c9924d146d000ac6bcf3c28bf0ac3c9baa011a.
//
// Solidity: event CollectMarginFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_Vault *VaultFilterer) ParseCollectMarginFees(log types.Log) (*VaultCollectMarginFees, error) {
	event := new(VaultCollectMarginFees)
	if err := _Vault.contract.UnpackLog(event, "CollectMarginFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultCollectSwapFeesIterator is returned from FilterCollectSwapFees and is used to iterate over the raw logs and unpacked data for CollectSwapFees events raised by the Vault contract.
type VaultCollectSwapFeesIterator struct {
	Event *VaultCollectSwapFees // Event containing the contract specifics and raw log

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
func (it *VaultCollectSwapFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultCollectSwapFees)
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
		it.Event = new(VaultCollectSwapFees)
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
func (it *VaultCollectSwapFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultCollectSwapFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultCollectSwapFees represents a CollectSwapFees event raised by the Vault contract.
type VaultCollectSwapFees struct {
	Token     common.Address
	FeeUsd    *big.Int
	FeeTokens *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectSwapFees is a free log retrieval operation binding the contract event 0x47cd9dda0e50ce30bcaaacd0488452b596221c07ac402a581cfae4d3933cac2b.
//
// Solidity: event CollectSwapFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_Vault *VaultFilterer) FilterCollectSwapFees(opts *bind.FilterOpts) (*VaultCollectSwapFeesIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "CollectSwapFees")
	if err != nil {
		return nil, err
	}
	return &VaultCollectSwapFeesIterator{contract: _Vault.contract, event: "CollectSwapFees", logs: logs, sub: sub}, nil
}

// WatchCollectSwapFees is a free log subscription operation binding the contract event 0x47cd9dda0e50ce30bcaaacd0488452b596221c07ac402a581cfae4d3933cac2b.
//
// Solidity: event CollectSwapFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_Vault *VaultFilterer) WatchCollectSwapFees(opts *bind.WatchOpts, sink chan<- *VaultCollectSwapFees) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "CollectSwapFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultCollectSwapFees)
				if err := _Vault.contract.UnpackLog(event, "CollectSwapFees", log); err != nil {
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

// ParseCollectSwapFees is a log parse operation binding the contract event 0x47cd9dda0e50ce30bcaaacd0488452b596221c07ac402a581cfae4d3933cac2b.
//
// Solidity: event CollectSwapFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_Vault *VaultFilterer) ParseCollectSwapFees(log types.Log) (*VaultCollectSwapFees, error) {
	event := new(VaultCollectSwapFees)
	if err := _Vault.contract.UnpackLog(event, "CollectSwapFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultDecreaseGuaranteedUsdIterator is returned from FilterDecreaseGuaranteedUsd and is used to iterate over the raw logs and unpacked data for DecreaseGuaranteedUsd events raised by the Vault contract.
type VaultDecreaseGuaranteedUsdIterator struct {
	Event *VaultDecreaseGuaranteedUsd // Event containing the contract specifics and raw log

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
func (it *VaultDecreaseGuaranteedUsdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDecreaseGuaranteedUsd)
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
		it.Event = new(VaultDecreaseGuaranteedUsd)
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
func (it *VaultDecreaseGuaranteedUsdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDecreaseGuaranteedUsdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDecreaseGuaranteedUsd represents a DecreaseGuaranteedUsd event raised by the Vault contract.
type VaultDecreaseGuaranteedUsd struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreaseGuaranteedUsd is a free log retrieval operation binding the contract event 0x34e07158b9db50df5613e591c44ea2ebc82834eff4a4dc3a46e000e608261d68.
//
// Solidity: event DecreaseGuaranteedUsd(address token, uint256 amount)
func (_Vault *VaultFilterer) FilterDecreaseGuaranteedUsd(opts *bind.FilterOpts) (*VaultDecreaseGuaranteedUsdIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "DecreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return &VaultDecreaseGuaranteedUsdIterator{contract: _Vault.contract, event: "DecreaseGuaranteedUsd", logs: logs, sub: sub}, nil
}

// WatchDecreaseGuaranteedUsd is a free log subscription operation binding the contract event 0x34e07158b9db50df5613e591c44ea2ebc82834eff4a4dc3a46e000e608261d68.
//
// Solidity: event DecreaseGuaranteedUsd(address token, uint256 amount)
func (_Vault *VaultFilterer) WatchDecreaseGuaranteedUsd(opts *bind.WatchOpts, sink chan<- *VaultDecreaseGuaranteedUsd) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "DecreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDecreaseGuaranteedUsd)
				if err := _Vault.contract.UnpackLog(event, "DecreaseGuaranteedUsd", log); err != nil {
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

// ParseDecreaseGuaranteedUsd is a log parse operation binding the contract event 0x34e07158b9db50df5613e591c44ea2ebc82834eff4a4dc3a46e000e608261d68.
//
// Solidity: event DecreaseGuaranteedUsd(address token, uint256 amount)
func (_Vault *VaultFilterer) ParseDecreaseGuaranteedUsd(log types.Log) (*VaultDecreaseGuaranteedUsd, error) {
	event := new(VaultDecreaseGuaranteedUsd)
	if err := _Vault.contract.UnpackLog(event, "DecreaseGuaranteedUsd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultDecreasePoolAmountIterator is returned from FilterDecreasePoolAmount and is used to iterate over the raw logs and unpacked data for DecreasePoolAmount events raised by the Vault contract.
type VaultDecreasePoolAmountIterator struct {
	Event *VaultDecreasePoolAmount // Event containing the contract specifics and raw log

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
func (it *VaultDecreasePoolAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDecreasePoolAmount)
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
		it.Event = new(VaultDecreasePoolAmount)
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
func (it *VaultDecreasePoolAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDecreasePoolAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDecreasePoolAmount represents a DecreasePoolAmount event raised by the Vault contract.
type VaultDecreasePoolAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreasePoolAmount is a free log retrieval operation binding the contract event 0x112726233fbeaeed0f5b1dba5cb0b2b81883dee49fb35ff99fd98ed9f6d31eb0.
//
// Solidity: event DecreasePoolAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) FilterDecreasePoolAmount(opts *bind.FilterOpts) (*VaultDecreasePoolAmountIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "DecreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return &VaultDecreasePoolAmountIterator{contract: _Vault.contract, event: "DecreasePoolAmount", logs: logs, sub: sub}, nil
}

// WatchDecreasePoolAmount is a free log subscription operation binding the contract event 0x112726233fbeaeed0f5b1dba5cb0b2b81883dee49fb35ff99fd98ed9f6d31eb0.
//
// Solidity: event DecreasePoolAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) WatchDecreasePoolAmount(opts *bind.WatchOpts, sink chan<- *VaultDecreasePoolAmount) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "DecreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDecreasePoolAmount)
				if err := _Vault.contract.UnpackLog(event, "DecreasePoolAmount", log); err != nil {
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

// ParseDecreasePoolAmount is a log parse operation binding the contract event 0x112726233fbeaeed0f5b1dba5cb0b2b81883dee49fb35ff99fd98ed9f6d31eb0.
//
// Solidity: event DecreasePoolAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) ParseDecreasePoolAmount(log types.Log) (*VaultDecreasePoolAmount, error) {
	event := new(VaultDecreasePoolAmount)
	if err := _Vault.contract.UnpackLog(event, "DecreasePoolAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultDecreasePositionIterator is returned from FilterDecreasePosition and is used to iterate over the raw logs and unpacked data for DecreasePosition events raised by the Vault contract.
type VaultDecreasePositionIterator struct {
	Event *VaultDecreasePosition // Event containing the contract specifics and raw log

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
func (it *VaultDecreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDecreasePosition)
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
		it.Event = new(VaultDecreasePosition)
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
func (it *VaultDecreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDecreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDecreasePosition represents a DecreasePosition event raised by the Vault contract.
type VaultDecreasePosition struct {
	Key             [32]byte
	Account         common.Address
	CollateralToken common.Address
	IndexToken      common.Address
	CollateralDelta *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	Price           *big.Int
	Fee             *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDecreasePosition is a free log retrieval operation binding the contract event 0x93d75d64d1f84fc6f430a64fc578bdd4c1e090e90ea2d51773e626d19de56d30.
//
// Solidity: event DecreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_Vault *VaultFilterer) FilterDecreasePosition(opts *bind.FilterOpts) (*VaultDecreasePositionIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "DecreasePosition")
	if err != nil {
		return nil, err
	}
	return &VaultDecreasePositionIterator{contract: _Vault.contract, event: "DecreasePosition", logs: logs, sub: sub}, nil
}

// WatchDecreasePosition is a free log subscription operation binding the contract event 0x93d75d64d1f84fc6f430a64fc578bdd4c1e090e90ea2d51773e626d19de56d30.
//
// Solidity: event DecreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_Vault *VaultFilterer) WatchDecreasePosition(opts *bind.WatchOpts, sink chan<- *VaultDecreasePosition) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "DecreasePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDecreasePosition)
				if err := _Vault.contract.UnpackLog(event, "DecreasePosition", log); err != nil {
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

// ParseDecreasePosition is a log parse operation binding the contract event 0x93d75d64d1f84fc6f430a64fc578bdd4c1e090e90ea2d51773e626d19de56d30.
//
// Solidity: event DecreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_Vault *VaultFilterer) ParseDecreasePosition(log types.Log) (*VaultDecreasePosition, error) {
	event := new(VaultDecreasePosition)
	if err := _Vault.contract.UnpackLog(event, "DecreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultDecreaseReservedAmountIterator is returned from FilterDecreaseReservedAmount and is used to iterate over the raw logs and unpacked data for DecreaseReservedAmount events raised by the Vault contract.
type VaultDecreaseReservedAmountIterator struct {
	Event *VaultDecreaseReservedAmount // Event containing the contract specifics and raw log

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
func (it *VaultDecreaseReservedAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDecreaseReservedAmount)
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
		it.Event = new(VaultDecreaseReservedAmount)
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
func (it *VaultDecreaseReservedAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDecreaseReservedAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDecreaseReservedAmount represents a DecreaseReservedAmount event raised by the Vault contract.
type VaultDecreaseReservedAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreaseReservedAmount is a free log retrieval operation binding the contract event 0x533cb5ed32be6a90284e96b5747a1bfc2d38fdb5768a6b5f67ff7d62144ed67b.
//
// Solidity: event DecreaseReservedAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) FilterDecreaseReservedAmount(opts *bind.FilterOpts) (*VaultDecreaseReservedAmountIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "DecreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return &VaultDecreaseReservedAmountIterator{contract: _Vault.contract, event: "DecreaseReservedAmount", logs: logs, sub: sub}, nil
}

// WatchDecreaseReservedAmount is a free log subscription operation binding the contract event 0x533cb5ed32be6a90284e96b5747a1bfc2d38fdb5768a6b5f67ff7d62144ed67b.
//
// Solidity: event DecreaseReservedAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) WatchDecreaseReservedAmount(opts *bind.WatchOpts, sink chan<- *VaultDecreaseReservedAmount) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "DecreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDecreaseReservedAmount)
				if err := _Vault.contract.UnpackLog(event, "DecreaseReservedAmount", log); err != nil {
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

// ParseDecreaseReservedAmount is a log parse operation binding the contract event 0x533cb5ed32be6a90284e96b5747a1bfc2d38fdb5768a6b5f67ff7d62144ed67b.
//
// Solidity: event DecreaseReservedAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) ParseDecreaseReservedAmount(log types.Log) (*VaultDecreaseReservedAmount, error) {
	event := new(VaultDecreaseReservedAmount)
	if err := _Vault.contract.UnpackLog(event, "DecreaseReservedAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultDecreaseUsdgAmountIterator is returned from FilterDecreaseUsdgAmount and is used to iterate over the raw logs and unpacked data for DecreaseUsdgAmount events raised by the Vault contract.
type VaultDecreaseUsdgAmountIterator struct {
	Event *VaultDecreaseUsdgAmount // Event containing the contract specifics and raw log

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
func (it *VaultDecreaseUsdgAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDecreaseUsdgAmount)
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
		it.Event = new(VaultDecreaseUsdgAmount)
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
func (it *VaultDecreaseUsdgAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDecreaseUsdgAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDecreaseUsdgAmount represents a DecreaseUsdgAmount event raised by the Vault contract.
type VaultDecreaseUsdgAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreaseUsdgAmount is a free log retrieval operation binding the contract event 0xe1e812596aac93a06ecc4ca627014d18e30f5c33b825160cc9d5c0ba61e45227.
//
// Solidity: event DecreaseUsdgAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) FilterDecreaseUsdgAmount(opts *bind.FilterOpts) (*VaultDecreaseUsdgAmountIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "DecreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return &VaultDecreaseUsdgAmountIterator{contract: _Vault.contract, event: "DecreaseUsdgAmount", logs: logs, sub: sub}, nil
}

// WatchDecreaseUsdgAmount is a free log subscription operation binding the contract event 0xe1e812596aac93a06ecc4ca627014d18e30f5c33b825160cc9d5c0ba61e45227.
//
// Solidity: event DecreaseUsdgAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) WatchDecreaseUsdgAmount(opts *bind.WatchOpts, sink chan<- *VaultDecreaseUsdgAmount) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "DecreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDecreaseUsdgAmount)
				if err := _Vault.contract.UnpackLog(event, "DecreaseUsdgAmount", log); err != nil {
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

// ParseDecreaseUsdgAmount is a log parse operation binding the contract event 0xe1e812596aac93a06ecc4ca627014d18e30f5c33b825160cc9d5c0ba61e45227.
//
// Solidity: event DecreaseUsdgAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) ParseDecreaseUsdgAmount(log types.Log) (*VaultDecreaseUsdgAmount, error) {
	event := new(VaultDecreaseUsdgAmount)
	if err := _Vault.contract.UnpackLog(event, "DecreaseUsdgAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultDirectPoolDepositIterator is returned from FilterDirectPoolDeposit and is used to iterate over the raw logs and unpacked data for DirectPoolDeposit events raised by the Vault contract.
type VaultDirectPoolDepositIterator struct {
	Event *VaultDirectPoolDeposit // Event containing the contract specifics and raw log

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
func (it *VaultDirectPoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDirectPoolDeposit)
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
		it.Event = new(VaultDirectPoolDeposit)
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
func (it *VaultDirectPoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDirectPoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDirectPoolDeposit represents a DirectPoolDeposit event raised by the Vault contract.
type VaultDirectPoolDeposit struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDirectPoolDeposit is a free log retrieval operation binding the contract event 0xa5a389190ebf6170a133bda5c769b77f4d6715b8aa172ec0ddf8473d0b4944bd.
//
// Solidity: event DirectPoolDeposit(address token, uint256 amount)
func (_Vault *VaultFilterer) FilterDirectPoolDeposit(opts *bind.FilterOpts) (*VaultDirectPoolDepositIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "DirectPoolDeposit")
	if err != nil {
		return nil, err
	}
	return &VaultDirectPoolDepositIterator{contract: _Vault.contract, event: "DirectPoolDeposit", logs: logs, sub: sub}, nil
}

// WatchDirectPoolDeposit is a free log subscription operation binding the contract event 0xa5a389190ebf6170a133bda5c769b77f4d6715b8aa172ec0ddf8473d0b4944bd.
//
// Solidity: event DirectPoolDeposit(address token, uint256 amount)
func (_Vault *VaultFilterer) WatchDirectPoolDeposit(opts *bind.WatchOpts, sink chan<- *VaultDirectPoolDeposit) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "DirectPoolDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDirectPoolDeposit)
				if err := _Vault.contract.UnpackLog(event, "DirectPoolDeposit", log); err != nil {
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

// ParseDirectPoolDeposit is a log parse operation binding the contract event 0xa5a389190ebf6170a133bda5c769b77f4d6715b8aa172ec0ddf8473d0b4944bd.
//
// Solidity: event DirectPoolDeposit(address token, uint256 amount)
func (_Vault *VaultFilterer) ParseDirectPoolDeposit(log types.Log) (*VaultDirectPoolDeposit, error) {
	event := new(VaultDirectPoolDeposit)
	if err := _Vault.contract.UnpackLog(event, "DirectPoolDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultIncreaseGuaranteedUsdIterator is returned from FilterIncreaseGuaranteedUsd and is used to iterate over the raw logs and unpacked data for IncreaseGuaranteedUsd events raised by the Vault contract.
type VaultIncreaseGuaranteedUsdIterator struct {
	Event *VaultIncreaseGuaranteedUsd // Event containing the contract specifics and raw log

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
func (it *VaultIncreaseGuaranteedUsdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultIncreaseGuaranteedUsd)
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
		it.Event = new(VaultIncreaseGuaranteedUsd)
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
func (it *VaultIncreaseGuaranteedUsdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultIncreaseGuaranteedUsdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultIncreaseGuaranteedUsd represents a IncreaseGuaranteedUsd event raised by the Vault contract.
type VaultIncreaseGuaranteedUsd struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreaseGuaranteedUsd is a free log retrieval operation binding the contract event 0xd9d4761f75e0d0103b5cbeab941eeb443d7a56a35b5baf2a0787c03f03f4e474.
//
// Solidity: event IncreaseGuaranteedUsd(address token, uint256 amount)
func (_Vault *VaultFilterer) FilterIncreaseGuaranteedUsd(opts *bind.FilterOpts) (*VaultIncreaseGuaranteedUsdIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "IncreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return &VaultIncreaseGuaranteedUsdIterator{contract: _Vault.contract, event: "IncreaseGuaranteedUsd", logs: logs, sub: sub}, nil
}

// WatchIncreaseGuaranteedUsd is a free log subscription operation binding the contract event 0xd9d4761f75e0d0103b5cbeab941eeb443d7a56a35b5baf2a0787c03f03f4e474.
//
// Solidity: event IncreaseGuaranteedUsd(address token, uint256 amount)
func (_Vault *VaultFilterer) WatchIncreaseGuaranteedUsd(opts *bind.WatchOpts, sink chan<- *VaultIncreaseGuaranteedUsd) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "IncreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultIncreaseGuaranteedUsd)
				if err := _Vault.contract.UnpackLog(event, "IncreaseGuaranteedUsd", log); err != nil {
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

// ParseIncreaseGuaranteedUsd is a log parse operation binding the contract event 0xd9d4761f75e0d0103b5cbeab941eeb443d7a56a35b5baf2a0787c03f03f4e474.
//
// Solidity: event IncreaseGuaranteedUsd(address token, uint256 amount)
func (_Vault *VaultFilterer) ParseIncreaseGuaranteedUsd(log types.Log) (*VaultIncreaseGuaranteedUsd, error) {
	event := new(VaultIncreaseGuaranteedUsd)
	if err := _Vault.contract.UnpackLog(event, "IncreaseGuaranteedUsd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultIncreasePoolAmountIterator is returned from FilterIncreasePoolAmount and is used to iterate over the raw logs and unpacked data for IncreasePoolAmount events raised by the Vault contract.
type VaultIncreasePoolAmountIterator struct {
	Event *VaultIncreasePoolAmount // Event containing the contract specifics and raw log

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
func (it *VaultIncreasePoolAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultIncreasePoolAmount)
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
		it.Event = new(VaultIncreasePoolAmount)
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
func (it *VaultIncreasePoolAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultIncreasePoolAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultIncreasePoolAmount represents a IncreasePoolAmount event raised by the Vault contract.
type VaultIncreasePoolAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreasePoolAmount is a free log retrieval operation binding the contract event 0x976177fbe09a15e5e43f848844963a42b41ef919ef17ff21a17a5421de8f4737.
//
// Solidity: event IncreasePoolAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) FilterIncreasePoolAmount(opts *bind.FilterOpts) (*VaultIncreasePoolAmountIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "IncreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return &VaultIncreasePoolAmountIterator{contract: _Vault.contract, event: "IncreasePoolAmount", logs: logs, sub: sub}, nil
}

// WatchIncreasePoolAmount is a free log subscription operation binding the contract event 0x976177fbe09a15e5e43f848844963a42b41ef919ef17ff21a17a5421de8f4737.
//
// Solidity: event IncreasePoolAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) WatchIncreasePoolAmount(opts *bind.WatchOpts, sink chan<- *VaultIncreasePoolAmount) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "IncreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultIncreasePoolAmount)
				if err := _Vault.contract.UnpackLog(event, "IncreasePoolAmount", log); err != nil {
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

// ParseIncreasePoolAmount is a log parse operation binding the contract event 0x976177fbe09a15e5e43f848844963a42b41ef919ef17ff21a17a5421de8f4737.
//
// Solidity: event IncreasePoolAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) ParseIncreasePoolAmount(log types.Log) (*VaultIncreasePoolAmount, error) {
	event := new(VaultIncreasePoolAmount)
	if err := _Vault.contract.UnpackLog(event, "IncreasePoolAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultIncreasePositionIterator is returned from FilterIncreasePosition and is used to iterate over the raw logs and unpacked data for IncreasePosition events raised by the Vault contract.
type VaultIncreasePositionIterator struct {
	Event *VaultIncreasePosition // Event containing the contract specifics and raw log

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
func (it *VaultIncreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultIncreasePosition)
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
		it.Event = new(VaultIncreasePosition)
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
func (it *VaultIncreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultIncreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultIncreasePosition represents a IncreasePosition event raised by the Vault contract.
type VaultIncreasePosition struct {
	Key             [32]byte
	Account         common.Address
	CollateralToken common.Address
	IndexToken      common.Address
	CollateralDelta *big.Int
	SizeDelta       *big.Int
	IsLong          bool
	Price           *big.Int
	Fee             *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterIncreasePosition is a free log retrieval operation binding the contract event 0x2fe68525253654c21998f35787a8d0f361905ef647c854092430ab65f2f15022.
//
// Solidity: event IncreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_Vault *VaultFilterer) FilterIncreasePosition(opts *bind.FilterOpts) (*VaultIncreasePositionIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "IncreasePosition")
	if err != nil {
		return nil, err
	}
	return &VaultIncreasePositionIterator{contract: _Vault.contract, event: "IncreasePosition", logs: logs, sub: sub}, nil
}

// WatchIncreasePosition is a free log subscription operation binding the contract event 0x2fe68525253654c21998f35787a8d0f361905ef647c854092430ab65f2f15022.
//
// Solidity: event IncreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_Vault *VaultFilterer) WatchIncreasePosition(opts *bind.WatchOpts, sink chan<- *VaultIncreasePosition) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "IncreasePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultIncreasePosition)
				if err := _Vault.contract.UnpackLog(event, "IncreasePosition", log); err != nil {
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

// ParseIncreasePosition is a log parse operation binding the contract event 0x2fe68525253654c21998f35787a8d0f361905ef647c854092430ab65f2f15022.
//
// Solidity: event IncreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_Vault *VaultFilterer) ParseIncreasePosition(log types.Log) (*VaultIncreasePosition, error) {
	event := new(VaultIncreasePosition)
	if err := _Vault.contract.UnpackLog(event, "IncreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultIncreaseReservedAmountIterator is returned from FilterIncreaseReservedAmount and is used to iterate over the raw logs and unpacked data for IncreaseReservedAmount events raised by the Vault contract.
type VaultIncreaseReservedAmountIterator struct {
	Event *VaultIncreaseReservedAmount // Event containing the contract specifics and raw log

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
func (it *VaultIncreaseReservedAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultIncreaseReservedAmount)
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
		it.Event = new(VaultIncreaseReservedAmount)
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
func (it *VaultIncreaseReservedAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultIncreaseReservedAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultIncreaseReservedAmount represents a IncreaseReservedAmount event raised by the Vault contract.
type VaultIncreaseReservedAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreaseReservedAmount is a free log retrieval operation binding the contract event 0xaa5649d82f5462be9d19b0f2b31a59b2259950a6076550bac9f3a1c07db9f66d.
//
// Solidity: event IncreaseReservedAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) FilterIncreaseReservedAmount(opts *bind.FilterOpts) (*VaultIncreaseReservedAmountIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "IncreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return &VaultIncreaseReservedAmountIterator{contract: _Vault.contract, event: "IncreaseReservedAmount", logs: logs, sub: sub}, nil
}

// WatchIncreaseReservedAmount is a free log subscription operation binding the contract event 0xaa5649d82f5462be9d19b0f2b31a59b2259950a6076550bac9f3a1c07db9f66d.
//
// Solidity: event IncreaseReservedAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) WatchIncreaseReservedAmount(opts *bind.WatchOpts, sink chan<- *VaultIncreaseReservedAmount) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "IncreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultIncreaseReservedAmount)
				if err := _Vault.contract.UnpackLog(event, "IncreaseReservedAmount", log); err != nil {
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

// ParseIncreaseReservedAmount is a log parse operation binding the contract event 0xaa5649d82f5462be9d19b0f2b31a59b2259950a6076550bac9f3a1c07db9f66d.
//
// Solidity: event IncreaseReservedAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) ParseIncreaseReservedAmount(log types.Log) (*VaultIncreaseReservedAmount, error) {
	event := new(VaultIncreaseReservedAmount)
	if err := _Vault.contract.UnpackLog(event, "IncreaseReservedAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultIncreaseUsdgAmountIterator is returned from FilterIncreaseUsdgAmount and is used to iterate over the raw logs and unpacked data for IncreaseUsdgAmount events raised by the Vault contract.
type VaultIncreaseUsdgAmountIterator struct {
	Event *VaultIncreaseUsdgAmount // Event containing the contract specifics and raw log

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
func (it *VaultIncreaseUsdgAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultIncreaseUsdgAmount)
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
		it.Event = new(VaultIncreaseUsdgAmount)
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
func (it *VaultIncreaseUsdgAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultIncreaseUsdgAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultIncreaseUsdgAmount represents a IncreaseUsdgAmount event raised by the Vault contract.
type VaultIncreaseUsdgAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreaseUsdgAmount is a free log retrieval operation binding the contract event 0x64243679a443432e2293343b77d411ff6144370404618f00ca0d2025d9ca9882.
//
// Solidity: event IncreaseUsdgAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) FilterIncreaseUsdgAmount(opts *bind.FilterOpts) (*VaultIncreaseUsdgAmountIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "IncreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return &VaultIncreaseUsdgAmountIterator{contract: _Vault.contract, event: "IncreaseUsdgAmount", logs: logs, sub: sub}, nil
}

// WatchIncreaseUsdgAmount is a free log subscription operation binding the contract event 0x64243679a443432e2293343b77d411ff6144370404618f00ca0d2025d9ca9882.
//
// Solidity: event IncreaseUsdgAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) WatchIncreaseUsdgAmount(opts *bind.WatchOpts, sink chan<- *VaultIncreaseUsdgAmount) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "IncreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultIncreaseUsdgAmount)
				if err := _Vault.contract.UnpackLog(event, "IncreaseUsdgAmount", log); err != nil {
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

// ParseIncreaseUsdgAmount is a log parse operation binding the contract event 0x64243679a443432e2293343b77d411ff6144370404618f00ca0d2025d9ca9882.
//
// Solidity: event IncreaseUsdgAmount(address token, uint256 amount)
func (_Vault *VaultFilterer) ParseIncreaseUsdgAmount(log types.Log) (*VaultIncreaseUsdgAmount, error) {
	event := new(VaultIncreaseUsdgAmount)
	if err := _Vault.contract.UnpackLog(event, "IncreaseUsdgAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultLiquidatePositionIterator is returned from FilterLiquidatePosition and is used to iterate over the raw logs and unpacked data for LiquidatePosition events raised by the Vault contract.
type VaultLiquidatePositionIterator struct {
	Event *VaultLiquidatePosition // Event containing the contract specifics and raw log

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
func (it *VaultLiquidatePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultLiquidatePosition)
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
		it.Event = new(VaultLiquidatePosition)
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
func (it *VaultLiquidatePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultLiquidatePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultLiquidatePosition represents a LiquidatePosition event raised by the Vault contract.
type VaultLiquidatePosition struct {
	Key             [32]byte
	Account         common.Address
	CollateralToken common.Address
	IndexToken      common.Address
	IsLong          bool
	Size            *big.Int
	Collateral      *big.Int
	ReserveAmount   *big.Int
	RealisedPnl     *big.Int
	MarkPrice       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLiquidatePosition is a free log retrieval operation binding the contract event 0x2e1f85a64a2f22cf2f0c42584e7c919ed4abe8d53675cff0f62bf1e95a1c676f.
//
// Solidity: event LiquidatePosition(bytes32 key, address account, address collateralToken, address indexToken, bool isLong, uint256 size, uint256 collateral, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_Vault *VaultFilterer) FilterLiquidatePosition(opts *bind.FilterOpts) (*VaultLiquidatePositionIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "LiquidatePosition")
	if err != nil {
		return nil, err
	}
	return &VaultLiquidatePositionIterator{contract: _Vault.contract, event: "LiquidatePosition", logs: logs, sub: sub}, nil
}

// WatchLiquidatePosition is a free log subscription operation binding the contract event 0x2e1f85a64a2f22cf2f0c42584e7c919ed4abe8d53675cff0f62bf1e95a1c676f.
//
// Solidity: event LiquidatePosition(bytes32 key, address account, address collateralToken, address indexToken, bool isLong, uint256 size, uint256 collateral, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_Vault *VaultFilterer) WatchLiquidatePosition(opts *bind.WatchOpts, sink chan<- *VaultLiquidatePosition) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "LiquidatePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultLiquidatePosition)
				if err := _Vault.contract.UnpackLog(event, "LiquidatePosition", log); err != nil {
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

// ParseLiquidatePosition is a log parse operation binding the contract event 0x2e1f85a64a2f22cf2f0c42584e7c919ed4abe8d53675cff0f62bf1e95a1c676f.
//
// Solidity: event LiquidatePosition(bytes32 key, address account, address collateralToken, address indexToken, bool isLong, uint256 size, uint256 collateral, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_Vault *VaultFilterer) ParseLiquidatePosition(log types.Log) (*VaultLiquidatePosition, error) {
	event := new(VaultLiquidatePosition)
	if err := _Vault.contract.UnpackLog(event, "LiquidatePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultSellUSDGIterator is returned from FilterSellUSDG and is used to iterate over the raw logs and unpacked data for SellUSDG events raised by the Vault contract.
type VaultSellUSDGIterator struct {
	Event *VaultSellUSDG // Event containing the contract specifics and raw log

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
func (it *VaultSellUSDGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultSellUSDG)
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
		it.Event = new(VaultSellUSDG)
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
func (it *VaultSellUSDGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultSellUSDGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultSellUSDG represents a SellUSDG event raised by the Vault contract.
type VaultSellUSDG struct {
	Account        common.Address
	Token          common.Address
	UsdgAmount     *big.Int
	TokenAmount    *big.Int
	FeeBasisPoints *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSellUSDG is a free log retrieval operation binding the contract event 0xd732b7828fa6cee72c285eac756fc66a7477e3dc22e22e7c432f1c265d40b483.
//
// Solidity: event SellUSDG(address account, address token, uint256 usdgAmount, uint256 tokenAmount, uint256 feeBasisPoints)
func (_Vault *VaultFilterer) FilterSellUSDG(opts *bind.FilterOpts) (*VaultSellUSDGIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "SellUSDG")
	if err != nil {
		return nil, err
	}
	return &VaultSellUSDGIterator{contract: _Vault.contract, event: "SellUSDG", logs: logs, sub: sub}, nil
}

// WatchSellUSDG is a free log subscription operation binding the contract event 0xd732b7828fa6cee72c285eac756fc66a7477e3dc22e22e7c432f1c265d40b483.
//
// Solidity: event SellUSDG(address account, address token, uint256 usdgAmount, uint256 tokenAmount, uint256 feeBasisPoints)
func (_Vault *VaultFilterer) WatchSellUSDG(opts *bind.WatchOpts, sink chan<- *VaultSellUSDG) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "SellUSDG")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultSellUSDG)
				if err := _Vault.contract.UnpackLog(event, "SellUSDG", log); err != nil {
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

// ParseSellUSDG is a log parse operation binding the contract event 0xd732b7828fa6cee72c285eac756fc66a7477e3dc22e22e7c432f1c265d40b483.
//
// Solidity: event SellUSDG(address account, address token, uint256 usdgAmount, uint256 tokenAmount, uint256 feeBasisPoints)
func (_Vault *VaultFilterer) ParseSellUSDG(log types.Log) (*VaultSellUSDG, error) {
	event := new(VaultSellUSDG)
	if err := _Vault.contract.UnpackLog(event, "SellUSDG", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Vault contract.
type VaultSwapIterator struct {
	Event *VaultSwap // Event containing the contract specifics and raw log

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
func (it *VaultSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultSwap)
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
		it.Event = new(VaultSwap)
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
func (it *VaultSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultSwap represents a Swap event raised by the Vault contract.
type VaultSwap struct {
	Account            common.Address
	TokenIn            common.Address
	TokenOut           common.Address
	AmountIn           *big.Int
	AmountOut          *big.Int
	AmountOutAfterFees *big.Int
	FeeBasisPoints     *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x0874b2d545cb271cdbda4e093020c452328b24af12382ed62c4d00f5c26709db.
//
// Solidity: event Swap(address account, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, uint256 amountOutAfterFees, uint256 feeBasisPoints)
func (_Vault *VaultFilterer) FilterSwap(opts *bind.FilterOpts) (*VaultSwapIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &VaultSwapIterator{contract: _Vault.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x0874b2d545cb271cdbda4e093020c452328b24af12382ed62c4d00f5c26709db.
//
// Solidity: event Swap(address account, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, uint256 amountOutAfterFees, uint256 feeBasisPoints)
func (_Vault *VaultFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *VaultSwap) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultSwap)
				if err := _Vault.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x0874b2d545cb271cdbda4e093020c452328b24af12382ed62c4d00f5c26709db.
//
// Solidity: event Swap(address account, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, uint256 amountOutAfterFees, uint256 feeBasisPoints)
func (_Vault *VaultFilterer) ParseSwap(log types.Log) (*VaultSwap, error) {
	event := new(VaultSwap)
	if err := _Vault.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultUpdateFundingRateIterator is returned from FilterUpdateFundingRate and is used to iterate over the raw logs and unpacked data for UpdateFundingRate events raised by the Vault contract.
type VaultUpdateFundingRateIterator struct {
	Event *VaultUpdateFundingRate // Event containing the contract specifics and raw log

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
func (it *VaultUpdateFundingRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultUpdateFundingRate)
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
		it.Event = new(VaultUpdateFundingRate)
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
func (it *VaultUpdateFundingRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultUpdateFundingRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultUpdateFundingRate represents a UpdateFundingRate event raised by the Vault contract.
type VaultUpdateFundingRate struct {
	Token       common.Address
	FundingRate *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateFundingRate is a free log retrieval operation binding the contract event 0xa146fc154e1913322e9817d49f0d5c37466c24326e15de10e739a948be815eab.
//
// Solidity: event UpdateFundingRate(address token, uint256 fundingRate)
func (_Vault *VaultFilterer) FilterUpdateFundingRate(opts *bind.FilterOpts) (*VaultUpdateFundingRateIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "UpdateFundingRate")
	if err != nil {
		return nil, err
	}
	return &VaultUpdateFundingRateIterator{contract: _Vault.contract, event: "UpdateFundingRate", logs: logs, sub: sub}, nil
}

// WatchUpdateFundingRate is a free log subscription operation binding the contract event 0xa146fc154e1913322e9817d49f0d5c37466c24326e15de10e739a948be815eab.
//
// Solidity: event UpdateFundingRate(address token, uint256 fundingRate)
func (_Vault *VaultFilterer) WatchUpdateFundingRate(opts *bind.WatchOpts, sink chan<- *VaultUpdateFundingRate) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "UpdateFundingRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultUpdateFundingRate)
				if err := _Vault.contract.UnpackLog(event, "UpdateFundingRate", log); err != nil {
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

// ParseUpdateFundingRate is a log parse operation binding the contract event 0xa146fc154e1913322e9817d49f0d5c37466c24326e15de10e739a948be815eab.
//
// Solidity: event UpdateFundingRate(address token, uint256 fundingRate)
func (_Vault *VaultFilterer) ParseUpdateFundingRate(log types.Log) (*VaultUpdateFundingRate, error) {
	event := new(VaultUpdateFundingRate)
	if err := _Vault.contract.UnpackLog(event, "UpdateFundingRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultUpdatePnlIterator is returned from FilterUpdatePnl and is used to iterate over the raw logs and unpacked data for UpdatePnl events raised by the Vault contract.
type VaultUpdatePnlIterator struct {
	Event *VaultUpdatePnl // Event containing the contract specifics and raw log

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
func (it *VaultUpdatePnlIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultUpdatePnl)
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
		it.Event = new(VaultUpdatePnl)
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
func (it *VaultUpdatePnlIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultUpdatePnlIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultUpdatePnl represents a UpdatePnl event raised by the Vault contract.
type VaultUpdatePnl struct {
	Key       [32]byte
	HasProfit bool
	Delta     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatePnl is a free log retrieval operation binding the contract event 0x3ff41bdde87755b687ae83d0221a232b6be51a803330ed9661c1b5d0105e0d8a.
//
// Solidity: event UpdatePnl(bytes32 key, bool hasProfit, uint256 delta)
func (_Vault *VaultFilterer) FilterUpdatePnl(opts *bind.FilterOpts) (*VaultUpdatePnlIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "UpdatePnl")
	if err != nil {
		return nil, err
	}
	return &VaultUpdatePnlIterator{contract: _Vault.contract, event: "UpdatePnl", logs: logs, sub: sub}, nil
}

// WatchUpdatePnl is a free log subscription operation binding the contract event 0x3ff41bdde87755b687ae83d0221a232b6be51a803330ed9661c1b5d0105e0d8a.
//
// Solidity: event UpdatePnl(bytes32 key, bool hasProfit, uint256 delta)
func (_Vault *VaultFilterer) WatchUpdatePnl(opts *bind.WatchOpts, sink chan<- *VaultUpdatePnl) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "UpdatePnl")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultUpdatePnl)
				if err := _Vault.contract.UnpackLog(event, "UpdatePnl", log); err != nil {
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

// ParseUpdatePnl is a log parse operation binding the contract event 0x3ff41bdde87755b687ae83d0221a232b6be51a803330ed9661c1b5d0105e0d8a.
//
// Solidity: event UpdatePnl(bytes32 key, bool hasProfit, uint256 delta)
func (_Vault *VaultFilterer) ParseUpdatePnl(log types.Log) (*VaultUpdatePnl, error) {
	event := new(VaultUpdatePnl)
	if err := _Vault.contract.UnpackLog(event, "UpdatePnl", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultUpdatePositionIterator is returned from FilterUpdatePosition and is used to iterate over the raw logs and unpacked data for UpdatePosition events raised by the Vault contract.
type VaultUpdatePositionIterator struct {
	Event *VaultUpdatePosition // Event containing the contract specifics and raw log

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
func (it *VaultUpdatePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultUpdatePosition)
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
		it.Event = new(VaultUpdatePosition)
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
func (it *VaultUpdatePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultUpdatePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultUpdatePosition represents a UpdatePosition event raised by the Vault contract.
type VaultUpdatePosition struct {
	Key              [32]byte
	Size             *big.Int
	Collateral       *big.Int
	AveragePrice     *big.Int
	EntryFundingRate *big.Int
	ReserveAmount    *big.Int
	RealisedPnl      *big.Int
	MarkPrice        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdatePosition is a free log retrieval operation binding the contract event 0x20853733b590dce729d9f4628682ebd9a34d2354e72679e66f43a008fc03b773.
//
// Solidity: event UpdatePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_Vault *VaultFilterer) FilterUpdatePosition(opts *bind.FilterOpts) (*VaultUpdatePositionIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "UpdatePosition")
	if err != nil {
		return nil, err
	}
	return &VaultUpdatePositionIterator{contract: _Vault.contract, event: "UpdatePosition", logs: logs, sub: sub}, nil
}

// WatchUpdatePosition is a free log subscription operation binding the contract event 0x20853733b590dce729d9f4628682ebd9a34d2354e72679e66f43a008fc03b773.
//
// Solidity: event UpdatePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_Vault *VaultFilterer) WatchUpdatePosition(opts *bind.WatchOpts, sink chan<- *VaultUpdatePosition) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "UpdatePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultUpdatePosition)
				if err := _Vault.contract.UnpackLog(event, "UpdatePosition", log); err != nil {
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

// ParseUpdatePosition is a log parse operation binding the contract event 0x20853733b590dce729d9f4628682ebd9a34d2354e72679e66f43a008fc03b773.
//
// Solidity: event UpdatePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_Vault *VaultFilterer) ParseUpdatePosition(log types.Log) (*VaultUpdatePosition, error) {
	event := new(VaultUpdatePosition)
	if err := _Vault.contract.UnpackLog(event, "UpdatePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

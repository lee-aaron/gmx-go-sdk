// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VaultV2

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

// VaultV2MetaData contains all meta data concerning the VaultV2 contract.
var VaultV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BASIS_POINTS_DIVISOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FUNDING_RATE_PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FEE_BASIS_POINTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FUNDING_RATE_FACTOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_LIQUIDATION_FEE_USD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_FUNDING_RATE_INTERVAL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_LEVERAGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PRICE_PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"USDG_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addRouter\",\"inputs\":[{\"name\":\"_router\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adjustForDecimals\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokenDiv\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenMul\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allWhitelistedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allWhitelistedTokensLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approvedRouters\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bufferAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"buyUSDG\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clearTokenConfig\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cumulativeFundingRates\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreasePosition\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"directPoolDeposit\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"errorController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"errors\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeReserves\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fundingInterval\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fundingRateFactor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDelta\",\"inputs\":[{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_averagePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_lastIncreasedTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEntryFundingRate\",\"inputs\":[{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeBasisPoints\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdgDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_feeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_taxBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_increment\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFundingFee\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_entryFundingRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalShortDelta\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxPrice\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinPrice\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextAveragePrice\",\"inputs\":[{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_averagePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_nextPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_lastIncreasedTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextFundingRate\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextGlobalShortAveragePrice\",\"inputs\":[{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nextPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPosition\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPositionDelta\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPositionFee\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPositionKey\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getPositionLeverage\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRedemptionAmount\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdgAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRedemptionCollateral\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRedemptionCollateralUsd\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTargetUsdgAmount\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUtilisation\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalShortAveragePrices\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalShortSizes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"guaranteedUsd\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasDynamicFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inManagerMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inPrivateLiquidationMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"includeAmmPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increasePosition\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdg\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_priceFeed\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_liquidationFeeUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_fundingRateFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stableFundingRateFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isInitialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isLeverageEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isLiquidator\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isManager\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSwapEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastFundingTimes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"liquidatePosition\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"liquidationFeeUsd\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marginFeeBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxGasPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxGlobalShortSizes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxLeverage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxUsdgAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minProfitBasisPoints\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minProfitTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mintBurnFeeBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"positions\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateral\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"averagePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entryFundingRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reserveAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"realisedPnl\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"lastIncreasedTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"priceFeed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeRouter\",\"inputs\":[{\"name\":\"_router\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reservedAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sellUSDG\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBufferAmount\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setError\",\"inputs\":[{\"name\":\"_errorCode\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setErrorController\",\"inputs\":[{\"name\":\"_errorController\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFees\",\"inputs\":[{\"name\":\"_taxBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stableTaxBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_mintBurnFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_swapFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stableSwapFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_marginFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_liquidationFeeUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minProfitTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_hasDynamicFees\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFundingRate\",\"inputs\":[{\"name\":\"_fundingInterval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_fundingRateFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stableFundingRateFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInManagerMode\",\"inputs\":[{\"name\":\"_inManagerMode\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInPrivateLiquidationMode\",\"inputs\":[{\"name\":\"_inPrivateLiquidationMode\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIsLeverageEnabled\",\"inputs\":[{\"name\":\"_isLeverageEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIsSwapEnabled\",\"inputs\":[{\"name\":\"_isSwapEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLiquidator\",\"inputs\":[{\"name\":\"_liquidator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setManager\",\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isManager\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxGasPrice\",\"inputs\":[{\"name\":\"_maxGasPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxGlobalShortSize\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxLeverage\",\"inputs\":[{\"name\":\"_maxLeverage\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPriceFeed\",\"inputs\":[{\"name\":\"_priceFeed\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenConfig\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenDecimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokenWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minProfitBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_maxUsdgAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isStable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_isShortable\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUsdgAmount\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVaultUtils\",\"inputs\":[{\"name\":\"_vaultUtils\",\"type\":\"address\",\"internalType\":\"contractIVaultUtils\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shortableTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stableFundingRateFactor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stableSwapFeeBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stableTaxBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stableTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"_tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swapFeeBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taxBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenBalances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenDecimals\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenToUsdMin\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenWeights\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalTokenWeights\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateCumulativeFundingRate\",\"inputs\":[{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeVault\",\"inputs\":[{\"name\":\"_newVault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"usdToToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdToTokenMax\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdToTokenMin\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdg\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdgAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"useSwapPricing\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateLiquidation\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_raise\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vaultUtils\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIVaultUtils\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelistedTokenCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelistedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BuyUSDG\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"usdgAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feeBasisPoints\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClosePosition\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"size\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateral\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"averagePrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"entryFundingRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reserveAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"realisedPnl\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollectMarginFees\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feeUsd\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feeTokens\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollectSwapFees\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feeUsd\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feeTokens\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecreaseGuaranteedUsd\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecreasePoolAmount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecreasePosition\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecreaseReservedAmount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecreaseUsdgAmount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DirectPoolDeposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncreaseGuaranteedUsd\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncreasePoolAmount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncreasePosition\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncreaseReservedAmount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncreaseUsdgAmount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidatePosition\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"size\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateral\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reserveAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"realisedPnl\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"markPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SellUSDG\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"usdgAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feeBasisPoints\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Swap\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenIn\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenOut\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOutAfterFees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feeBasisPoints\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateFundingRate\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"fundingRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatePnl\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"hasProfit\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"delta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatePosition\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"size\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateral\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"averagePrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"entryFundingRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reserveAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"realisedPnl\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"markPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// VaultV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultV2MetaData.ABI instead.
var VaultV2ABI = VaultV2MetaData.ABI

// VaultV2 is an auto generated Go binding around an Ethereum contract.
type VaultV2 struct {
	VaultV2Caller     // Read-only binding to the contract
	VaultV2Transactor // Write-only binding to the contract
	VaultV2Filterer   // Log filterer for contract events
}

// VaultV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type VaultV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultV2Session struct {
	Contract     *VaultV2          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultV2CallerSession struct {
	Contract *VaultV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// VaultV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultV2TransactorSession struct {
	Contract     *VaultV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VaultV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type VaultV2Raw struct {
	Contract *VaultV2 // Generic contract binding to access the raw methods on
}

// VaultV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultV2CallerRaw struct {
	Contract *VaultV2Caller // Generic read-only contract binding to access the raw methods on
}

// VaultV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultV2TransactorRaw struct {
	Contract *VaultV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultV2 creates a new instance of VaultV2, bound to a specific deployed contract.
func NewVaultV2(address common.Address, backend bind.ContractBackend) (*VaultV2, error) {
	contract, err := bindVaultV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultV2{VaultV2Caller: VaultV2Caller{contract: contract}, VaultV2Transactor: VaultV2Transactor{contract: contract}, VaultV2Filterer: VaultV2Filterer{contract: contract}}, nil
}

// NewVaultV2Caller creates a new read-only instance of VaultV2, bound to a specific deployed contract.
func NewVaultV2Caller(address common.Address, caller bind.ContractCaller) (*VaultV2Caller, error) {
	contract, err := bindVaultV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultV2Caller{contract: contract}, nil
}

// NewVaultV2Transactor creates a new write-only instance of VaultV2, bound to a specific deployed contract.
func NewVaultV2Transactor(address common.Address, transactor bind.ContractTransactor) (*VaultV2Transactor, error) {
	contract, err := bindVaultV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultV2Transactor{contract: contract}, nil
}

// NewVaultV2Filterer creates a new log filterer instance of VaultV2, bound to a specific deployed contract.
func NewVaultV2Filterer(address common.Address, filterer bind.ContractFilterer) (*VaultV2Filterer, error) {
	contract, err := bindVaultV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultV2Filterer{contract: contract}, nil
}

// bindVaultV2 binds a generic wrapper to an already deployed contract.
func bindVaultV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VaultV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultV2 *VaultV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultV2.Contract.VaultV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultV2 *VaultV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultV2.Contract.VaultV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultV2 *VaultV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultV2.Contract.VaultV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultV2 *VaultV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultV2 *VaultV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultV2 *VaultV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultV2.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_VaultV2 *VaultV2Caller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_VaultV2 *VaultV2Session) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _VaultV2.Contract.BASISPOINTSDIVISOR(&_VaultV2.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _VaultV2.Contract.BASISPOINTSDIVISOR(&_VaultV2.CallOpts)
}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_VaultV2 *VaultV2Caller) FUNDINGRATEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "FUNDING_RATE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_VaultV2 *VaultV2Session) FUNDINGRATEPRECISION() (*big.Int, error) {
	return _VaultV2.Contract.FUNDINGRATEPRECISION(&_VaultV2.CallOpts)
}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) FUNDINGRATEPRECISION() (*big.Int, error) {
	return _VaultV2.Contract.FUNDINGRATEPRECISION(&_VaultV2.CallOpts)
}

// MAXFEEBASISPOINTS is a free data retrieval call binding the contract method 0x4befe2ca.
//
// Solidity: function MAX_FEE_BASIS_POINTS() view returns(uint256)
func (_VaultV2 *VaultV2Caller) MAXFEEBASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "MAX_FEE_BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEEBASISPOINTS is a free data retrieval call binding the contract method 0x4befe2ca.
//
// Solidity: function MAX_FEE_BASIS_POINTS() view returns(uint256)
func (_VaultV2 *VaultV2Session) MAXFEEBASISPOINTS() (*big.Int, error) {
	return _VaultV2.Contract.MAXFEEBASISPOINTS(&_VaultV2.CallOpts)
}

// MAXFEEBASISPOINTS is a free data retrieval call binding the contract method 0x4befe2ca.
//
// Solidity: function MAX_FEE_BASIS_POINTS() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) MAXFEEBASISPOINTS() (*big.Int, error) {
	return _VaultV2.Contract.MAXFEEBASISPOINTS(&_VaultV2.CallOpts)
}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_VaultV2 *VaultV2Caller) MAXFUNDINGRATEFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "MAX_FUNDING_RATE_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_VaultV2 *VaultV2Session) MAXFUNDINGRATEFACTOR() (*big.Int, error) {
	return _VaultV2.Contract.MAXFUNDINGRATEFACTOR(&_VaultV2.CallOpts)
}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) MAXFUNDINGRATEFACTOR() (*big.Int, error) {
	return _VaultV2.Contract.MAXFUNDINGRATEFACTOR(&_VaultV2.CallOpts)
}

// MAXLIQUIDATIONFEEUSD is a free data retrieval call binding the contract method 0x07c58752.
//
// Solidity: function MAX_LIQUIDATION_FEE_USD() view returns(uint256)
func (_VaultV2 *VaultV2Caller) MAXLIQUIDATIONFEEUSD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "MAX_LIQUIDATION_FEE_USD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLIQUIDATIONFEEUSD is a free data retrieval call binding the contract method 0x07c58752.
//
// Solidity: function MAX_LIQUIDATION_FEE_USD() view returns(uint256)
func (_VaultV2 *VaultV2Session) MAXLIQUIDATIONFEEUSD() (*big.Int, error) {
	return _VaultV2.Contract.MAXLIQUIDATIONFEEUSD(&_VaultV2.CallOpts)
}

// MAXLIQUIDATIONFEEUSD is a free data retrieval call binding the contract method 0x07c58752.
//
// Solidity: function MAX_LIQUIDATION_FEE_USD() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) MAXLIQUIDATIONFEEUSD() (*big.Int, error) {
	return _VaultV2.Contract.MAXLIQUIDATIONFEEUSD(&_VaultV2.CallOpts)
}

// MINFUNDINGRATEINTERVAL is a free data retrieval call binding the contract method 0xfce28c10.
//
// Solidity: function MIN_FUNDING_RATE_INTERVAL() view returns(uint256)
func (_VaultV2 *VaultV2Caller) MINFUNDINGRATEINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "MIN_FUNDING_RATE_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINFUNDINGRATEINTERVAL is a free data retrieval call binding the contract method 0xfce28c10.
//
// Solidity: function MIN_FUNDING_RATE_INTERVAL() view returns(uint256)
func (_VaultV2 *VaultV2Session) MINFUNDINGRATEINTERVAL() (*big.Int, error) {
	return _VaultV2.Contract.MINFUNDINGRATEINTERVAL(&_VaultV2.CallOpts)
}

// MINFUNDINGRATEINTERVAL is a free data retrieval call binding the contract method 0xfce28c10.
//
// Solidity: function MIN_FUNDING_RATE_INTERVAL() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) MINFUNDINGRATEINTERVAL() (*big.Int, error) {
	return _VaultV2.Contract.MINFUNDINGRATEINTERVAL(&_VaultV2.CallOpts)
}

// MINLEVERAGE is a free data retrieval call binding the contract method 0x34c1557d.
//
// Solidity: function MIN_LEVERAGE() view returns(uint256)
func (_VaultV2 *VaultV2Caller) MINLEVERAGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "MIN_LEVERAGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINLEVERAGE is a free data retrieval call binding the contract method 0x34c1557d.
//
// Solidity: function MIN_LEVERAGE() view returns(uint256)
func (_VaultV2 *VaultV2Session) MINLEVERAGE() (*big.Int, error) {
	return _VaultV2.Contract.MINLEVERAGE(&_VaultV2.CallOpts)
}

// MINLEVERAGE is a free data retrieval call binding the contract method 0x34c1557d.
//
// Solidity: function MIN_LEVERAGE() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) MINLEVERAGE() (*big.Int, error) {
	return _VaultV2.Contract.MINLEVERAGE(&_VaultV2.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_VaultV2 *VaultV2Caller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_VaultV2 *VaultV2Session) PRICEPRECISION() (*big.Int, error) {
	return _VaultV2.Contract.PRICEPRECISION(&_VaultV2.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) PRICEPRECISION() (*big.Int, error) {
	return _VaultV2.Contract.PRICEPRECISION(&_VaultV2.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_VaultV2 *VaultV2Caller) USDGDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "USDG_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_VaultV2 *VaultV2Session) USDGDECIMALS() (*big.Int, error) {
	return _VaultV2.Contract.USDGDECIMALS(&_VaultV2.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) USDGDECIMALS() (*big.Int, error) {
	return _VaultV2.Contract.USDGDECIMALS(&_VaultV2.CallOpts)
}

// AdjustForDecimals is a free data retrieval call binding the contract method 0x42152873.
//
// Solidity: function adjustForDecimals(uint256 _amount, address _tokenDiv, address _tokenMul) view returns(uint256)
func (_VaultV2 *VaultV2Caller) AdjustForDecimals(opts *bind.CallOpts, _amount *big.Int, _tokenDiv common.Address, _tokenMul common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "adjustForDecimals", _amount, _tokenDiv, _tokenMul)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdjustForDecimals is a free data retrieval call binding the contract method 0x42152873.
//
// Solidity: function adjustForDecimals(uint256 _amount, address _tokenDiv, address _tokenMul) view returns(uint256)
func (_VaultV2 *VaultV2Session) AdjustForDecimals(_amount *big.Int, _tokenDiv common.Address, _tokenMul common.Address) (*big.Int, error) {
	return _VaultV2.Contract.AdjustForDecimals(&_VaultV2.CallOpts, _amount, _tokenDiv, _tokenMul)
}

// AdjustForDecimals is a free data retrieval call binding the contract method 0x42152873.
//
// Solidity: function adjustForDecimals(uint256 _amount, address _tokenDiv, address _tokenMul) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) AdjustForDecimals(_amount *big.Int, _tokenDiv common.Address, _tokenMul common.Address) (*big.Int, error) {
	return _VaultV2.Contract.AdjustForDecimals(&_VaultV2.CallOpts, _amount, _tokenDiv, _tokenMul)
}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_VaultV2 *VaultV2Caller) AllWhitelistedTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "allWhitelistedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_VaultV2 *VaultV2Session) AllWhitelistedTokens(arg0 *big.Int) (common.Address, error) {
	return _VaultV2.Contract.AllWhitelistedTokens(&_VaultV2.CallOpts, arg0)
}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_VaultV2 *VaultV2CallerSession) AllWhitelistedTokens(arg0 *big.Int) (common.Address, error) {
	return _VaultV2.Contract.AllWhitelistedTokens(&_VaultV2.CallOpts, arg0)
}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_VaultV2 *VaultV2Caller) AllWhitelistedTokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "allWhitelistedTokensLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_VaultV2 *VaultV2Session) AllWhitelistedTokensLength() (*big.Int, error) {
	return _VaultV2.Contract.AllWhitelistedTokensLength(&_VaultV2.CallOpts)
}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) AllWhitelistedTokensLength() (*big.Int, error) {
	return _VaultV2.Contract.AllWhitelistedTokensLength(&_VaultV2.CallOpts)
}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address , address ) view returns(bool)
func (_VaultV2 *VaultV2Caller) ApprovedRouters(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "approvedRouters", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address , address ) view returns(bool)
func (_VaultV2 *VaultV2Session) ApprovedRouters(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _VaultV2.Contract.ApprovedRouters(&_VaultV2.CallOpts, arg0, arg1)
}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address , address ) view returns(bool)
func (_VaultV2 *VaultV2CallerSession) ApprovedRouters(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _VaultV2.Contract.ApprovedRouters(&_VaultV2.CallOpts, arg0, arg1)
}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address ) view returns(uint256)
func (_VaultV2 *VaultV2Caller) BufferAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "bufferAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address ) view returns(uint256)
func (_VaultV2 *VaultV2Session) BufferAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.BufferAmounts(&_VaultV2.CallOpts, arg0)
}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address ) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) BufferAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.BufferAmounts(&_VaultV2.CallOpts, arg0)
}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address ) view returns(uint256)
func (_VaultV2 *VaultV2Caller) CumulativeFundingRates(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "cumulativeFundingRates", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address ) view returns(uint256)
func (_VaultV2 *VaultV2Session) CumulativeFundingRates(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.CumulativeFundingRates(&_VaultV2.CallOpts, arg0)
}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address ) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) CumulativeFundingRates(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.CumulativeFundingRates(&_VaultV2.CallOpts, arg0)
}

// ErrorController is a free data retrieval call binding the contract method 0x48f35cbb.
//
// Solidity: function errorController() view returns(address)
func (_VaultV2 *VaultV2Caller) ErrorController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "errorController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ErrorController is a free data retrieval call binding the contract method 0x48f35cbb.
//
// Solidity: function errorController() view returns(address)
func (_VaultV2 *VaultV2Session) ErrorController() (common.Address, error) {
	return _VaultV2.Contract.ErrorController(&_VaultV2.CallOpts)
}

// ErrorController is a free data retrieval call binding the contract method 0x48f35cbb.
//
// Solidity: function errorController() view returns(address)
func (_VaultV2 *VaultV2CallerSession) ErrorController() (common.Address, error) {
	return _VaultV2.Contract.ErrorController(&_VaultV2.CallOpts)
}

// Errors is a free data retrieval call binding the contract method 0xfed1a606.
//
// Solidity: function errors(uint256 ) view returns(string)
func (_VaultV2 *VaultV2Caller) Errors(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "errors", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Errors is a free data retrieval call binding the contract method 0xfed1a606.
//
// Solidity: function errors(uint256 ) view returns(string)
func (_VaultV2 *VaultV2Session) Errors(arg0 *big.Int) (string, error) {
	return _VaultV2.Contract.Errors(&_VaultV2.CallOpts, arg0)
}

// Errors is a free data retrieval call binding the contract method 0xfed1a606.
//
// Solidity: function errors(uint256 ) view returns(string)
func (_VaultV2 *VaultV2CallerSession) Errors(arg0 *big.Int) (string, error) {
	return _VaultV2.Contract.Errors(&_VaultV2.CallOpts, arg0)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_VaultV2 *VaultV2Caller) FeeReserves(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "feeReserves", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_VaultV2 *VaultV2Session) FeeReserves(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.FeeReserves(&_VaultV2.CallOpts, arg0)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) FeeReserves(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.FeeReserves(&_VaultV2.CallOpts, arg0)
}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_VaultV2 *VaultV2Caller) FundingInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "fundingInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_VaultV2 *VaultV2Session) FundingInterval() (*big.Int, error) {
	return _VaultV2.Contract.FundingInterval(&_VaultV2.CallOpts)
}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) FundingInterval() (*big.Int, error) {
	return _VaultV2.Contract.FundingInterval(&_VaultV2.CallOpts)
}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_VaultV2 *VaultV2Caller) FundingRateFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "fundingRateFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_VaultV2 *VaultV2Session) FundingRateFactor() (*big.Int, error) {
	return _VaultV2.Contract.FundingRateFactor(&_VaultV2.CallOpts)
}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) FundingRateFactor() (*big.Int, error) {
	return _VaultV2.Contract.FundingRateFactor(&_VaultV2.CallOpts)
}

// GetDelta is a free data retrieval call binding the contract method 0x5c07eaab.
//
// Solidity: function getDelta(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _lastIncreasedTime) view returns(bool, uint256)
func (_VaultV2 *VaultV2Caller) GetDelta(opts *bind.CallOpts, _indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getDelta", _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)

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
func (_VaultV2 *VaultV2Session) GetDelta(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	return _VaultV2.Contract.GetDelta(&_VaultV2.CallOpts, _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)
}

// GetDelta is a free data retrieval call binding the contract method 0x5c07eaab.
//
// Solidity: function getDelta(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _lastIncreasedTime) view returns(bool, uint256)
func (_VaultV2 *VaultV2CallerSession) GetDelta(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	return _VaultV2.Contract.GetDelta(&_VaultV2.CallOpts, _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)
}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_VaultV2 *VaultV2Caller) GetEntryFundingRate(opts *bind.CallOpts, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getEntryFundingRate", _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_VaultV2 *VaultV2Session) GetEntryFundingRate(_collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _VaultV2.Contract.GetEntryFundingRate(&_VaultV2.CallOpts, _collateralToken, _indexToken, _isLong)
}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) GetEntryFundingRate(_collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _VaultV2.Contract.GetEntryFundingRate(&_VaultV2.CallOpts, _collateralToken, _indexToken, _isLong)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_VaultV2 *VaultV2Caller) GetFeeBasisPoints(opts *bind.CallOpts, _token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getFeeBasisPoints", _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_VaultV2 *VaultV2Session) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _VaultV2.Contract.GetFeeBasisPoints(&_VaultV2.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _VaultV2.Contract.GetFeeBasisPoints(&_VaultV2.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_VaultV2 *VaultV2Caller) GetFundingFee(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getFundingFee", _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_VaultV2 *VaultV2Session) GetFundingFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	return _VaultV2.Contract.GetFundingFee(&_VaultV2.CallOpts, _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)
}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) GetFundingFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	return _VaultV2.Contract.GetFundingFee(&_VaultV2.CallOpts, _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)
}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xb364accb.
//
// Solidity: function getGlobalShortDelta(address _token) view returns(bool, uint256)
func (_VaultV2 *VaultV2Caller) GetGlobalShortDelta(opts *bind.CallOpts, _token common.Address) (bool, *big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getGlobalShortDelta", _token)

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
func (_VaultV2 *VaultV2Session) GetGlobalShortDelta(_token common.Address) (bool, *big.Int, error) {
	return _VaultV2.Contract.GetGlobalShortDelta(&_VaultV2.CallOpts, _token)
}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xb364accb.
//
// Solidity: function getGlobalShortDelta(address _token) view returns(bool, uint256)
func (_VaultV2 *VaultV2CallerSession) GetGlobalShortDelta(_token common.Address) (bool, *big.Int, error) {
	return _VaultV2.Contract.GetGlobalShortDelta(&_VaultV2.CallOpts, _token)
}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_VaultV2 *VaultV2Caller) GetMaxPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getMaxPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_VaultV2 *VaultV2Session) GetMaxPrice(_token common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GetMaxPrice(&_VaultV2.CallOpts, _token)
}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) GetMaxPrice(_token common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GetMaxPrice(&_VaultV2.CallOpts, _token)
}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_VaultV2 *VaultV2Caller) GetMinPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getMinPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_VaultV2 *VaultV2Session) GetMinPrice(_token common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GetMinPrice(&_VaultV2.CallOpts, _token)
}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) GetMinPrice(_token common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GetMinPrice(&_VaultV2.CallOpts, _token)
}

// GetNextAveragePrice is a free data retrieval call binding the contract method 0xdb97495f.
//
// Solidity: function getNextAveragePrice(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _nextPrice, uint256 _sizeDelta, uint256 _lastIncreasedTime) view returns(uint256)
func (_VaultV2 *VaultV2Caller) GetNextAveragePrice(opts *bind.CallOpts, _indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _nextPrice *big.Int, _sizeDelta *big.Int, _lastIncreasedTime *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getNextAveragePrice", _indexToken, _size, _averagePrice, _isLong, _nextPrice, _sizeDelta, _lastIncreasedTime)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextAveragePrice is a free data retrieval call binding the contract method 0xdb97495f.
//
// Solidity: function getNextAveragePrice(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _nextPrice, uint256 _sizeDelta, uint256 _lastIncreasedTime) view returns(uint256)
func (_VaultV2 *VaultV2Session) GetNextAveragePrice(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _nextPrice *big.Int, _sizeDelta *big.Int, _lastIncreasedTime *big.Int) (*big.Int, error) {
	return _VaultV2.Contract.GetNextAveragePrice(&_VaultV2.CallOpts, _indexToken, _size, _averagePrice, _isLong, _nextPrice, _sizeDelta, _lastIncreasedTime)
}

// GetNextAveragePrice is a free data retrieval call binding the contract method 0xdb97495f.
//
// Solidity: function getNextAveragePrice(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _nextPrice, uint256 _sizeDelta, uint256 _lastIncreasedTime) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) GetNextAveragePrice(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _nextPrice *big.Int, _sizeDelta *big.Int, _lastIncreasedTime *big.Int) (*big.Int, error) {
	return _VaultV2.Contract.GetNextAveragePrice(&_VaultV2.CallOpts, _indexToken, _size, _averagePrice, _isLong, _nextPrice, _sizeDelta, _lastIncreasedTime)
}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_VaultV2 *VaultV2Caller) GetNextFundingRate(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getNextFundingRate", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_VaultV2 *VaultV2Session) GetNextFundingRate(_token common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GetNextFundingRate(&_VaultV2.CallOpts, _token)
}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) GetNextFundingRate(_token common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GetNextFundingRate(&_VaultV2.CallOpts, _token)
}

// GetNextGlobalShortAveragePrice is a free data retrieval call binding the contract method 0x9d7432ca.
//
// Solidity: function getNextGlobalShortAveragePrice(address _indexToken, uint256 _nextPrice, uint256 _sizeDelta) view returns(uint256)
func (_VaultV2 *VaultV2Caller) GetNextGlobalShortAveragePrice(opts *bind.CallOpts, _indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getNextGlobalShortAveragePrice", _indexToken, _nextPrice, _sizeDelta)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextGlobalShortAveragePrice is a free data retrieval call binding the contract method 0x9d7432ca.
//
// Solidity: function getNextGlobalShortAveragePrice(address _indexToken, uint256 _nextPrice, uint256 _sizeDelta) view returns(uint256)
func (_VaultV2 *VaultV2Session) GetNextGlobalShortAveragePrice(_indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int) (*big.Int, error) {
	return _VaultV2.Contract.GetNextGlobalShortAveragePrice(&_VaultV2.CallOpts, _indexToken, _nextPrice, _sizeDelta)
}

// GetNextGlobalShortAveragePrice is a free data retrieval call binding the contract method 0x9d7432ca.
//
// Solidity: function getNextGlobalShortAveragePrice(address _indexToken, uint256 _nextPrice, uint256 _sizeDelta) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) GetNextGlobalShortAveragePrice(_indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int) (*big.Int, error) {
	return _VaultV2.Contract.GetNextGlobalShortAveragePrice(&_VaultV2.CallOpts, _indexToken, _nextPrice, _sizeDelta)
}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256, uint256, uint256, uint256, uint256, uint256, bool, uint256)
func (_VaultV2 *VaultV2Caller) GetPosition(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getPosition", _account, _collateralToken, _indexToken, _isLong)

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
func (_VaultV2 *VaultV2Session) GetPosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _VaultV2.Contract.GetPosition(&_VaultV2.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256, uint256, uint256, uint256, uint256, uint256, bool, uint256)
func (_VaultV2 *VaultV2CallerSession) GetPosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _VaultV2.Contract.GetPosition(&_VaultV2.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionDelta is a free data retrieval call binding the contract method 0x45a6f370.
//
// Solidity: function getPositionDelta(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(bool, uint256)
func (_VaultV2 *VaultV2Caller) GetPositionDelta(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (bool, *big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getPositionDelta", _account, _collateralToken, _indexToken, _isLong)

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
func (_VaultV2 *VaultV2Session) GetPositionDelta(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (bool, *big.Int, error) {
	return _VaultV2.Contract.GetPositionDelta(&_VaultV2.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionDelta is a free data retrieval call binding the contract method 0x45a6f370.
//
// Solidity: function getPositionDelta(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(bool, uint256)
func (_VaultV2 *VaultV2CallerSession) GetPositionDelta(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (bool, *big.Int, error) {
	return _VaultV2.Contract.GetPositionDelta(&_VaultV2.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_VaultV2 *VaultV2Caller) GetPositionFee(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getPositionFee", _account, _collateralToken, _indexToken, _isLong, _sizeDelta)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_VaultV2 *VaultV2Session) GetPositionFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	return _VaultV2.Contract.GetPositionFee(&_VaultV2.CallOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta)
}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) GetPositionFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	return _VaultV2.Contract.GetPositionFee(&_VaultV2.CallOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta)
}

// GetPositionKey is a free data retrieval call binding the contract method 0x2d4b0576.
//
// Solidity: function getPositionKey(address _account, address _collateralToken, address _indexToken, bool _isLong) pure returns(bytes32)
func (_VaultV2 *VaultV2Caller) GetPositionKey(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) ([32]byte, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getPositionKey", _account, _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPositionKey is a free data retrieval call binding the contract method 0x2d4b0576.
//
// Solidity: function getPositionKey(address _account, address _collateralToken, address _indexToken, bool _isLong) pure returns(bytes32)
func (_VaultV2 *VaultV2Session) GetPositionKey(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) ([32]byte, error) {
	return _VaultV2.Contract.GetPositionKey(&_VaultV2.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionKey is a free data retrieval call binding the contract method 0x2d4b0576.
//
// Solidity: function getPositionKey(address _account, address _collateralToken, address _indexToken, bool _isLong) pure returns(bytes32)
func (_VaultV2 *VaultV2CallerSession) GetPositionKey(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) ([32]byte, error) {
	return _VaultV2.Contract.GetPositionKey(&_VaultV2.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionLeverage is a free data retrieval call binding the contract method 0x51723e82.
//
// Solidity: function getPositionLeverage(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_VaultV2 *VaultV2Caller) GetPositionLeverage(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getPositionLeverage", _account, _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionLeverage is a free data retrieval call binding the contract method 0x51723e82.
//
// Solidity: function getPositionLeverage(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_VaultV2 *VaultV2Session) GetPositionLeverage(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _VaultV2.Contract.GetPositionLeverage(&_VaultV2.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionLeverage is a free data retrieval call binding the contract method 0x51723e82.
//
// Solidity: function getPositionLeverage(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) GetPositionLeverage(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _VaultV2.Contract.GetPositionLeverage(&_VaultV2.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_VaultV2 *VaultV2Caller) GetRedemptionAmount(opts *bind.CallOpts, _token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getRedemptionAmount", _token, _usdgAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_VaultV2 *VaultV2Session) GetRedemptionAmount(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _VaultV2.Contract.GetRedemptionAmount(&_VaultV2.CallOpts, _token, _usdgAmount)
}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) GetRedemptionAmount(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _VaultV2.Contract.GetRedemptionAmount(&_VaultV2.CallOpts, _token, _usdgAmount)
}

// GetRedemptionCollateral is a free data retrieval call binding the contract method 0xb136ca49.
//
// Solidity: function getRedemptionCollateral(address _token) view returns(uint256)
func (_VaultV2 *VaultV2Caller) GetRedemptionCollateral(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getRedemptionCollateral", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionCollateral is a free data retrieval call binding the contract method 0xb136ca49.
//
// Solidity: function getRedemptionCollateral(address _token) view returns(uint256)
func (_VaultV2 *VaultV2Session) GetRedemptionCollateral(_token common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GetRedemptionCollateral(&_VaultV2.CallOpts, _token)
}

// GetRedemptionCollateral is a free data retrieval call binding the contract method 0xb136ca49.
//
// Solidity: function getRedemptionCollateral(address _token) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) GetRedemptionCollateral(_token common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GetRedemptionCollateral(&_VaultV2.CallOpts, _token)
}

// GetRedemptionCollateralUsd is a free data retrieval call binding the contract method 0x29ff9615.
//
// Solidity: function getRedemptionCollateralUsd(address _token) view returns(uint256)
func (_VaultV2 *VaultV2Caller) GetRedemptionCollateralUsd(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getRedemptionCollateralUsd", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionCollateralUsd is a free data retrieval call binding the contract method 0x29ff9615.
//
// Solidity: function getRedemptionCollateralUsd(address _token) view returns(uint256)
func (_VaultV2 *VaultV2Session) GetRedemptionCollateralUsd(_token common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GetRedemptionCollateralUsd(&_VaultV2.CallOpts, _token)
}

// GetRedemptionCollateralUsd is a free data retrieval call binding the contract method 0x29ff9615.
//
// Solidity: function getRedemptionCollateralUsd(address _token) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) GetRedemptionCollateralUsd(_token common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GetRedemptionCollateralUsd(&_VaultV2.CallOpts, _token)
}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_VaultV2 *VaultV2Caller) GetTargetUsdgAmount(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getTargetUsdgAmount", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_VaultV2 *VaultV2Session) GetTargetUsdgAmount(_token common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GetTargetUsdgAmount(&_VaultV2.CallOpts, _token)
}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) GetTargetUsdgAmount(_token common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GetTargetUsdgAmount(&_VaultV2.CallOpts, _token)
}

// GetUtilisation is a free data retrieval call binding the contract method 0x04fef1db.
//
// Solidity: function getUtilisation(address _token) view returns(uint256)
func (_VaultV2 *VaultV2Caller) GetUtilisation(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "getUtilisation", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUtilisation is a free data retrieval call binding the contract method 0x04fef1db.
//
// Solidity: function getUtilisation(address _token) view returns(uint256)
func (_VaultV2 *VaultV2Session) GetUtilisation(_token common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GetUtilisation(&_VaultV2.CallOpts, _token)
}

// GetUtilisation is a free data retrieval call binding the contract method 0x04fef1db.
//
// Solidity: function getUtilisation(address _token) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) GetUtilisation(_token common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GetUtilisation(&_VaultV2.CallOpts, _token)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_VaultV2 *VaultV2Caller) GlobalShortAveragePrices(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "globalShortAveragePrices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_VaultV2 *VaultV2Session) GlobalShortAveragePrices(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GlobalShortAveragePrices(&_VaultV2.CallOpts, arg0)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) GlobalShortAveragePrices(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GlobalShortAveragePrices(&_VaultV2.CallOpts, arg0)
}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address ) view returns(uint256)
func (_VaultV2 *VaultV2Caller) GlobalShortSizes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "globalShortSizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address ) view returns(uint256)
func (_VaultV2 *VaultV2Session) GlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GlobalShortSizes(&_VaultV2.CallOpts, arg0)
}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address ) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) GlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GlobalShortSizes(&_VaultV2.CallOpts, arg0)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VaultV2 *VaultV2Caller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VaultV2 *VaultV2Session) Gov() (common.Address, error) {
	return _VaultV2.Contract.Gov(&_VaultV2.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VaultV2 *VaultV2CallerSession) Gov() (common.Address, error) {
	return _VaultV2.Contract.Gov(&_VaultV2.CallOpts)
}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address ) view returns(uint256)
func (_VaultV2 *VaultV2Caller) GuaranteedUsd(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "guaranteedUsd", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address ) view returns(uint256)
func (_VaultV2 *VaultV2Session) GuaranteedUsd(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GuaranteedUsd(&_VaultV2.CallOpts, arg0)
}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address ) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) GuaranteedUsd(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.GuaranteedUsd(&_VaultV2.CallOpts, arg0)
}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_VaultV2 *VaultV2Caller) HasDynamicFees(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "hasDynamicFees")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_VaultV2 *VaultV2Session) HasDynamicFees() (bool, error) {
	return _VaultV2.Contract.HasDynamicFees(&_VaultV2.CallOpts)
}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_VaultV2 *VaultV2CallerSession) HasDynamicFees() (bool, error) {
	return _VaultV2.Contract.HasDynamicFees(&_VaultV2.CallOpts)
}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_VaultV2 *VaultV2Caller) InManagerMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "inManagerMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_VaultV2 *VaultV2Session) InManagerMode() (bool, error) {
	return _VaultV2.Contract.InManagerMode(&_VaultV2.CallOpts)
}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_VaultV2 *VaultV2CallerSession) InManagerMode() (bool, error) {
	return _VaultV2.Contract.InManagerMode(&_VaultV2.CallOpts)
}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_VaultV2 *VaultV2Caller) InPrivateLiquidationMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "inPrivateLiquidationMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_VaultV2 *VaultV2Session) InPrivateLiquidationMode() (bool, error) {
	return _VaultV2.Contract.InPrivateLiquidationMode(&_VaultV2.CallOpts)
}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_VaultV2 *VaultV2CallerSession) InPrivateLiquidationMode() (bool, error) {
	return _VaultV2.Contract.InPrivateLiquidationMode(&_VaultV2.CallOpts)
}

// IncludeAmmPrice is a free data retrieval call binding the contract method 0xab08c1c6.
//
// Solidity: function includeAmmPrice() view returns(bool)
func (_VaultV2 *VaultV2Caller) IncludeAmmPrice(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "includeAmmPrice")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IncludeAmmPrice is a free data retrieval call binding the contract method 0xab08c1c6.
//
// Solidity: function includeAmmPrice() view returns(bool)
func (_VaultV2 *VaultV2Session) IncludeAmmPrice() (bool, error) {
	return _VaultV2.Contract.IncludeAmmPrice(&_VaultV2.CallOpts)
}

// IncludeAmmPrice is a free data retrieval call binding the contract method 0xab08c1c6.
//
// Solidity: function includeAmmPrice() view returns(bool)
func (_VaultV2 *VaultV2CallerSession) IncludeAmmPrice() (bool, error) {
	return _VaultV2.Contract.IncludeAmmPrice(&_VaultV2.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_VaultV2 *VaultV2Caller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_VaultV2 *VaultV2Session) IsInitialized() (bool, error) {
	return _VaultV2.Contract.IsInitialized(&_VaultV2.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_VaultV2 *VaultV2CallerSession) IsInitialized() (bool, error) {
	return _VaultV2.Contract.IsInitialized(&_VaultV2.CallOpts)
}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_VaultV2 *VaultV2Caller) IsLeverageEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "isLeverageEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_VaultV2 *VaultV2Session) IsLeverageEnabled() (bool, error) {
	return _VaultV2.Contract.IsLeverageEnabled(&_VaultV2.CallOpts)
}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_VaultV2 *VaultV2CallerSession) IsLeverageEnabled() (bool, error) {
	return _VaultV2.Contract.IsLeverageEnabled(&_VaultV2.CallOpts)
}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address ) view returns(bool)
func (_VaultV2 *VaultV2Caller) IsLiquidator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "isLiquidator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address ) view returns(bool)
func (_VaultV2 *VaultV2Session) IsLiquidator(arg0 common.Address) (bool, error) {
	return _VaultV2.Contract.IsLiquidator(&_VaultV2.CallOpts, arg0)
}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address ) view returns(bool)
func (_VaultV2 *VaultV2CallerSession) IsLiquidator(arg0 common.Address) (bool, error) {
	return _VaultV2.Contract.IsLiquidator(&_VaultV2.CallOpts, arg0)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_VaultV2 *VaultV2Caller) IsManager(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "isManager", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_VaultV2 *VaultV2Session) IsManager(arg0 common.Address) (bool, error) {
	return _VaultV2.Contract.IsManager(&_VaultV2.CallOpts, arg0)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_VaultV2 *VaultV2CallerSession) IsManager(arg0 common.Address) (bool, error) {
	return _VaultV2.Contract.IsManager(&_VaultV2.CallOpts, arg0)
}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_VaultV2 *VaultV2Caller) IsSwapEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "isSwapEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_VaultV2 *VaultV2Session) IsSwapEnabled() (bool, error) {
	return _VaultV2.Contract.IsSwapEnabled(&_VaultV2.CallOpts)
}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_VaultV2 *VaultV2CallerSession) IsSwapEnabled() (bool, error) {
	return _VaultV2.Contract.IsSwapEnabled(&_VaultV2.CallOpts)
}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address ) view returns(uint256)
func (_VaultV2 *VaultV2Caller) LastFundingTimes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "lastFundingTimes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address ) view returns(uint256)
func (_VaultV2 *VaultV2Session) LastFundingTimes(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.LastFundingTimes(&_VaultV2.CallOpts, arg0)
}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address ) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) LastFundingTimes(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.LastFundingTimes(&_VaultV2.CallOpts, arg0)
}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_VaultV2 *VaultV2Caller) LiquidationFeeUsd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "liquidationFeeUsd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_VaultV2 *VaultV2Session) LiquidationFeeUsd() (*big.Int, error) {
	return _VaultV2.Contract.LiquidationFeeUsd(&_VaultV2.CallOpts)
}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) LiquidationFeeUsd() (*big.Int, error) {
	return _VaultV2.Contract.LiquidationFeeUsd(&_VaultV2.CallOpts)
}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_VaultV2 *VaultV2Caller) MarginFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "marginFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_VaultV2 *VaultV2Session) MarginFeeBasisPoints() (*big.Int, error) {
	return _VaultV2.Contract.MarginFeeBasisPoints(&_VaultV2.CallOpts)
}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) MarginFeeBasisPoints() (*big.Int, error) {
	return _VaultV2.Contract.MarginFeeBasisPoints(&_VaultV2.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_VaultV2 *VaultV2Caller) MaxGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "maxGasPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_VaultV2 *VaultV2Session) MaxGasPrice() (*big.Int, error) {
	return _VaultV2.Contract.MaxGasPrice(&_VaultV2.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) MaxGasPrice() (*big.Int, error) {
	return _VaultV2.Contract.MaxGasPrice(&_VaultV2.CallOpts)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_VaultV2 *VaultV2Caller) MaxGlobalShortSizes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "maxGlobalShortSizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_VaultV2 *VaultV2Session) MaxGlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.MaxGlobalShortSizes(&_VaultV2.CallOpts, arg0)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) MaxGlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.MaxGlobalShortSizes(&_VaultV2.CallOpts, arg0)
}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_VaultV2 *VaultV2Caller) MaxLeverage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "maxLeverage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_VaultV2 *VaultV2Session) MaxLeverage() (*big.Int, error) {
	return _VaultV2.Contract.MaxLeverage(&_VaultV2.CallOpts)
}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) MaxLeverage() (*big.Int, error) {
	return _VaultV2.Contract.MaxLeverage(&_VaultV2.CallOpts)
}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address ) view returns(uint256)
func (_VaultV2 *VaultV2Caller) MaxUsdgAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "maxUsdgAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address ) view returns(uint256)
func (_VaultV2 *VaultV2Session) MaxUsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.MaxUsdgAmounts(&_VaultV2.CallOpts, arg0)
}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address ) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) MaxUsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.MaxUsdgAmounts(&_VaultV2.CallOpts, arg0)
}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address ) view returns(uint256)
func (_VaultV2 *VaultV2Caller) MinProfitBasisPoints(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "minProfitBasisPoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address ) view returns(uint256)
func (_VaultV2 *VaultV2Session) MinProfitBasisPoints(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.MinProfitBasisPoints(&_VaultV2.CallOpts, arg0)
}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address ) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) MinProfitBasisPoints(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.MinProfitBasisPoints(&_VaultV2.CallOpts, arg0)
}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_VaultV2 *VaultV2Caller) MinProfitTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "minProfitTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_VaultV2 *VaultV2Session) MinProfitTime() (*big.Int, error) {
	return _VaultV2.Contract.MinProfitTime(&_VaultV2.CallOpts)
}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) MinProfitTime() (*big.Int, error) {
	return _VaultV2.Contract.MinProfitTime(&_VaultV2.CallOpts)
}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_VaultV2 *VaultV2Caller) MintBurnFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "mintBurnFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_VaultV2 *VaultV2Session) MintBurnFeeBasisPoints() (*big.Int, error) {
	return _VaultV2.Contract.MintBurnFeeBasisPoints(&_VaultV2.CallOpts)
}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) MintBurnFeeBasisPoints() (*big.Int, error) {
	return _VaultV2.Contract.MintBurnFeeBasisPoints(&_VaultV2.CallOpts)
}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address ) view returns(uint256)
func (_VaultV2 *VaultV2Caller) PoolAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "poolAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address ) view returns(uint256)
func (_VaultV2 *VaultV2Session) PoolAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.PoolAmounts(&_VaultV2.CallOpts, arg0)
}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address ) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) PoolAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.PoolAmounts(&_VaultV2.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 lastIncreasedTime)
func (_VaultV2 *VaultV2Caller) Positions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Size              *big.Int
	Collateral        *big.Int
	AveragePrice      *big.Int
	EntryFundingRate  *big.Int
	ReserveAmount     *big.Int
	RealisedPnl       *big.Int
	LastIncreasedTime *big.Int
}, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "positions", arg0)

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
func (_VaultV2 *VaultV2Session) Positions(arg0 [32]byte) (struct {
	Size              *big.Int
	Collateral        *big.Int
	AveragePrice      *big.Int
	EntryFundingRate  *big.Int
	ReserveAmount     *big.Int
	RealisedPnl       *big.Int
	LastIncreasedTime *big.Int
}, error) {
	return _VaultV2.Contract.Positions(&_VaultV2.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 lastIncreasedTime)
func (_VaultV2 *VaultV2CallerSession) Positions(arg0 [32]byte) (struct {
	Size              *big.Int
	Collateral        *big.Int
	AveragePrice      *big.Int
	EntryFundingRate  *big.Int
	ReserveAmount     *big.Int
	RealisedPnl       *big.Int
	LastIncreasedTime *big.Int
}, error) {
	return _VaultV2.Contract.Positions(&_VaultV2.CallOpts, arg0)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_VaultV2 *VaultV2Caller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_VaultV2 *VaultV2Session) PriceFeed() (common.Address, error) {
	return _VaultV2.Contract.PriceFeed(&_VaultV2.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_VaultV2 *VaultV2CallerSession) PriceFeed() (common.Address, error) {
	return _VaultV2.Contract.PriceFeed(&_VaultV2.CallOpts)
}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address ) view returns(uint256)
func (_VaultV2 *VaultV2Caller) ReservedAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "reservedAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address ) view returns(uint256)
func (_VaultV2 *VaultV2Session) ReservedAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.ReservedAmounts(&_VaultV2.CallOpts, arg0)
}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address ) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) ReservedAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.ReservedAmounts(&_VaultV2.CallOpts, arg0)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_VaultV2 *VaultV2Caller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_VaultV2 *VaultV2Session) Router() (common.Address, error) {
	return _VaultV2.Contract.Router(&_VaultV2.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_VaultV2 *VaultV2CallerSession) Router() (common.Address, error) {
	return _VaultV2.Contract.Router(&_VaultV2.CallOpts)
}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address ) view returns(bool)
func (_VaultV2 *VaultV2Caller) ShortableTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "shortableTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address ) view returns(bool)
func (_VaultV2 *VaultV2Session) ShortableTokens(arg0 common.Address) (bool, error) {
	return _VaultV2.Contract.ShortableTokens(&_VaultV2.CallOpts, arg0)
}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address ) view returns(bool)
func (_VaultV2 *VaultV2CallerSession) ShortableTokens(arg0 common.Address) (bool, error) {
	return _VaultV2.Contract.ShortableTokens(&_VaultV2.CallOpts, arg0)
}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_VaultV2 *VaultV2Caller) StableFundingRateFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "stableFundingRateFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_VaultV2 *VaultV2Session) StableFundingRateFactor() (*big.Int, error) {
	return _VaultV2.Contract.StableFundingRateFactor(&_VaultV2.CallOpts)
}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) StableFundingRateFactor() (*big.Int, error) {
	return _VaultV2.Contract.StableFundingRateFactor(&_VaultV2.CallOpts)
}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_VaultV2 *VaultV2Caller) StableSwapFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "stableSwapFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_VaultV2 *VaultV2Session) StableSwapFeeBasisPoints() (*big.Int, error) {
	return _VaultV2.Contract.StableSwapFeeBasisPoints(&_VaultV2.CallOpts)
}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) StableSwapFeeBasisPoints() (*big.Int, error) {
	return _VaultV2.Contract.StableSwapFeeBasisPoints(&_VaultV2.CallOpts)
}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_VaultV2 *VaultV2Caller) StableTaxBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "stableTaxBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_VaultV2 *VaultV2Session) StableTaxBasisPoints() (*big.Int, error) {
	return _VaultV2.Contract.StableTaxBasisPoints(&_VaultV2.CallOpts)
}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) StableTaxBasisPoints() (*big.Int, error) {
	return _VaultV2.Contract.StableTaxBasisPoints(&_VaultV2.CallOpts)
}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address ) view returns(bool)
func (_VaultV2 *VaultV2Caller) StableTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "stableTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address ) view returns(bool)
func (_VaultV2 *VaultV2Session) StableTokens(arg0 common.Address) (bool, error) {
	return _VaultV2.Contract.StableTokens(&_VaultV2.CallOpts, arg0)
}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address ) view returns(bool)
func (_VaultV2 *VaultV2CallerSession) StableTokens(arg0 common.Address) (bool, error) {
	return _VaultV2.Contract.StableTokens(&_VaultV2.CallOpts, arg0)
}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_VaultV2 *VaultV2Caller) SwapFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "swapFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_VaultV2 *VaultV2Session) SwapFeeBasisPoints() (*big.Int, error) {
	return _VaultV2.Contract.SwapFeeBasisPoints(&_VaultV2.CallOpts)
}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) SwapFeeBasisPoints() (*big.Int, error) {
	return _VaultV2.Contract.SwapFeeBasisPoints(&_VaultV2.CallOpts)
}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_VaultV2 *VaultV2Caller) TaxBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "taxBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_VaultV2 *VaultV2Session) TaxBasisPoints() (*big.Int, error) {
	return _VaultV2.Contract.TaxBasisPoints(&_VaultV2.CallOpts)
}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) TaxBasisPoints() (*big.Int, error) {
	return _VaultV2.Contract.TaxBasisPoints(&_VaultV2.CallOpts)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_VaultV2 *VaultV2Caller) TokenBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "tokenBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_VaultV2 *VaultV2Session) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.TokenBalances(&_VaultV2.CallOpts, arg0)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.TokenBalances(&_VaultV2.CallOpts, arg0)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint256)
func (_VaultV2 *VaultV2Caller) TokenDecimals(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "tokenDecimals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint256)
func (_VaultV2 *VaultV2Session) TokenDecimals(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.TokenDecimals(&_VaultV2.CallOpts, arg0)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) TokenDecimals(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.TokenDecimals(&_VaultV2.CallOpts, arg0)
}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_VaultV2 *VaultV2Caller) TokenToUsdMin(opts *bind.CallOpts, _token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "tokenToUsdMin", _token, _tokenAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_VaultV2 *VaultV2Session) TokenToUsdMin(_token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	return _VaultV2.Contract.TokenToUsdMin(&_VaultV2.CallOpts, _token, _tokenAmount)
}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) TokenToUsdMin(_token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	return _VaultV2.Contract.TokenToUsdMin(&_VaultV2.CallOpts, _token, _tokenAmount)
}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address ) view returns(uint256)
func (_VaultV2 *VaultV2Caller) TokenWeights(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "tokenWeights", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address ) view returns(uint256)
func (_VaultV2 *VaultV2Session) TokenWeights(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.TokenWeights(&_VaultV2.CallOpts, arg0)
}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address ) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) TokenWeights(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.TokenWeights(&_VaultV2.CallOpts, arg0)
}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_VaultV2 *VaultV2Caller) TotalTokenWeights(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "totalTokenWeights")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_VaultV2 *VaultV2Session) TotalTokenWeights() (*big.Int, error) {
	return _VaultV2.Contract.TotalTokenWeights(&_VaultV2.CallOpts)
}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) TotalTokenWeights() (*big.Int, error) {
	return _VaultV2.Contract.TotalTokenWeights(&_VaultV2.CallOpts)
}

// UsdToToken is a free data retrieval call binding the contract method 0xfa12dbc0.
//
// Solidity: function usdToToken(address _token, uint256 _usdAmount, uint256 _price) view returns(uint256)
func (_VaultV2 *VaultV2Caller) UsdToToken(opts *bind.CallOpts, _token common.Address, _usdAmount *big.Int, _price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "usdToToken", _token, _usdAmount, _price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdToToken is a free data retrieval call binding the contract method 0xfa12dbc0.
//
// Solidity: function usdToToken(address _token, uint256 _usdAmount, uint256 _price) view returns(uint256)
func (_VaultV2 *VaultV2Session) UsdToToken(_token common.Address, _usdAmount *big.Int, _price *big.Int) (*big.Int, error) {
	return _VaultV2.Contract.UsdToToken(&_VaultV2.CallOpts, _token, _usdAmount, _price)
}

// UsdToToken is a free data retrieval call binding the contract method 0xfa12dbc0.
//
// Solidity: function usdToToken(address _token, uint256 _usdAmount, uint256 _price) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) UsdToToken(_token common.Address, _usdAmount *big.Int, _price *big.Int) (*big.Int, error) {
	return _VaultV2.Contract.UsdToToken(&_VaultV2.CallOpts, _token, _usdAmount, _price)
}

// UsdToTokenMax is a free data retrieval call binding the contract method 0xa42ab3d2.
//
// Solidity: function usdToTokenMax(address _token, uint256 _usdAmount) view returns(uint256)
func (_VaultV2 *VaultV2Caller) UsdToTokenMax(opts *bind.CallOpts, _token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "usdToTokenMax", _token, _usdAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdToTokenMax is a free data retrieval call binding the contract method 0xa42ab3d2.
//
// Solidity: function usdToTokenMax(address _token, uint256 _usdAmount) view returns(uint256)
func (_VaultV2 *VaultV2Session) UsdToTokenMax(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _VaultV2.Contract.UsdToTokenMax(&_VaultV2.CallOpts, _token, _usdAmount)
}

// UsdToTokenMax is a free data retrieval call binding the contract method 0xa42ab3d2.
//
// Solidity: function usdToTokenMax(address _token, uint256 _usdAmount) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) UsdToTokenMax(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _VaultV2.Contract.UsdToTokenMax(&_VaultV2.CallOpts, _token, _usdAmount)
}

// UsdToTokenMin is a free data retrieval call binding the contract method 0x9899cd02.
//
// Solidity: function usdToTokenMin(address _token, uint256 _usdAmount) view returns(uint256)
func (_VaultV2 *VaultV2Caller) UsdToTokenMin(opts *bind.CallOpts, _token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "usdToTokenMin", _token, _usdAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdToTokenMin is a free data retrieval call binding the contract method 0x9899cd02.
//
// Solidity: function usdToTokenMin(address _token, uint256 _usdAmount) view returns(uint256)
func (_VaultV2 *VaultV2Session) UsdToTokenMin(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _VaultV2.Contract.UsdToTokenMin(&_VaultV2.CallOpts, _token, _usdAmount)
}

// UsdToTokenMin is a free data retrieval call binding the contract method 0x9899cd02.
//
// Solidity: function usdToTokenMin(address _token, uint256 _usdAmount) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) UsdToTokenMin(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _VaultV2.Contract.UsdToTokenMin(&_VaultV2.CallOpts, _token, _usdAmount)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_VaultV2 *VaultV2Caller) Usdg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "usdg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_VaultV2 *VaultV2Session) Usdg() (common.Address, error) {
	return _VaultV2.Contract.Usdg(&_VaultV2.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_VaultV2 *VaultV2CallerSession) Usdg() (common.Address, error) {
	return _VaultV2.Contract.Usdg(&_VaultV2.CallOpts)
}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address ) view returns(uint256)
func (_VaultV2 *VaultV2Caller) UsdgAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "usdgAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address ) view returns(uint256)
func (_VaultV2 *VaultV2Session) UsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.UsdgAmounts(&_VaultV2.CallOpts, arg0)
}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address ) view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) UsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2.Contract.UsdgAmounts(&_VaultV2.CallOpts, arg0)
}

// UseSwapPricing is a free data retrieval call binding the contract method 0xb06423f3.
//
// Solidity: function useSwapPricing() view returns(bool)
func (_VaultV2 *VaultV2Caller) UseSwapPricing(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "useSwapPricing")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseSwapPricing is a free data retrieval call binding the contract method 0xb06423f3.
//
// Solidity: function useSwapPricing() view returns(bool)
func (_VaultV2 *VaultV2Session) UseSwapPricing() (bool, error) {
	return _VaultV2.Contract.UseSwapPricing(&_VaultV2.CallOpts)
}

// UseSwapPricing is a free data retrieval call binding the contract method 0xb06423f3.
//
// Solidity: function useSwapPricing() view returns(bool)
func (_VaultV2 *VaultV2CallerSession) UseSwapPricing() (bool, error) {
	return _VaultV2.Contract.UseSwapPricing(&_VaultV2.CallOpts)
}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_VaultV2 *VaultV2Caller) ValidateLiquidation(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "validateLiquidation", _account, _collateralToken, _indexToken, _isLong, _raise)

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
func (_VaultV2 *VaultV2Session) ValidateLiquidation(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	return _VaultV2.Contract.ValidateLiquidation(&_VaultV2.CallOpts, _account, _collateralToken, _indexToken, _isLong, _raise)
}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_VaultV2 *VaultV2CallerSession) ValidateLiquidation(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	return _VaultV2.Contract.ValidateLiquidation(&_VaultV2.CallOpts, _account, _collateralToken, _indexToken, _isLong, _raise)
}

// VaultUtils is a free data retrieval call binding the contract method 0x6abbe0c8.
//
// Solidity: function vaultUtils() view returns(address)
func (_VaultV2 *VaultV2Caller) VaultUtils(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "vaultUtils")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultUtils is a free data retrieval call binding the contract method 0x6abbe0c8.
//
// Solidity: function vaultUtils() view returns(address)
func (_VaultV2 *VaultV2Session) VaultUtils() (common.Address, error) {
	return _VaultV2.Contract.VaultUtils(&_VaultV2.CallOpts)
}

// VaultUtils is a free data retrieval call binding the contract method 0x6abbe0c8.
//
// Solidity: function vaultUtils() view returns(address)
func (_VaultV2 *VaultV2CallerSession) VaultUtils() (common.Address, error) {
	return _VaultV2.Contract.VaultUtils(&_VaultV2.CallOpts)
}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_VaultV2 *VaultV2Caller) WhitelistedTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "whitelistedTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_VaultV2 *VaultV2Session) WhitelistedTokenCount() (*big.Int, error) {
	return _VaultV2.Contract.WhitelistedTokenCount(&_VaultV2.CallOpts)
}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_VaultV2 *VaultV2CallerSession) WhitelistedTokenCount() (*big.Int, error) {
	return _VaultV2.Contract.WhitelistedTokenCount(&_VaultV2.CallOpts)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_VaultV2 *VaultV2Caller) WhitelistedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultV2.contract.Call(opts, &out, "whitelistedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_VaultV2 *VaultV2Session) WhitelistedTokens(arg0 common.Address) (bool, error) {
	return _VaultV2.Contract.WhitelistedTokens(&_VaultV2.CallOpts, arg0)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_VaultV2 *VaultV2CallerSession) WhitelistedTokens(arg0 common.Address) (bool, error) {
	return _VaultV2.Contract.WhitelistedTokens(&_VaultV2.CallOpts, arg0)
}

// AddRouter is a paid mutator transaction binding the contract method 0x24ca984e.
//
// Solidity: function addRouter(address _router) returns()
func (_VaultV2 *VaultV2Transactor) AddRouter(opts *bind.TransactOpts, _router common.Address) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "addRouter", _router)
}

// AddRouter is a paid mutator transaction binding the contract method 0x24ca984e.
//
// Solidity: function addRouter(address _router) returns()
func (_VaultV2 *VaultV2Session) AddRouter(_router common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.AddRouter(&_VaultV2.TransactOpts, _router)
}

// AddRouter is a paid mutator transaction binding the contract method 0x24ca984e.
//
// Solidity: function addRouter(address _router) returns()
func (_VaultV2 *VaultV2TransactorSession) AddRouter(_router common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.AddRouter(&_VaultV2.TransactOpts, _router)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_VaultV2 *VaultV2Transactor) BuyUSDG(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "buyUSDG", _token, _receiver)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_VaultV2 *VaultV2Session) BuyUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.BuyUSDG(&_VaultV2.TransactOpts, _token, _receiver)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_VaultV2 *VaultV2TransactorSession) BuyUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.BuyUSDG(&_VaultV2.TransactOpts, _token, _receiver)
}

// ClearTokenConfig is a paid mutator transaction binding the contract method 0xe67f59a7.
//
// Solidity: function clearTokenConfig(address _token) returns()
func (_VaultV2 *VaultV2Transactor) ClearTokenConfig(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "clearTokenConfig", _token)
}

// ClearTokenConfig is a paid mutator transaction binding the contract method 0xe67f59a7.
//
// Solidity: function clearTokenConfig(address _token) returns()
func (_VaultV2 *VaultV2Session) ClearTokenConfig(_token common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.ClearTokenConfig(&_VaultV2.TransactOpts, _token)
}

// ClearTokenConfig is a paid mutator transaction binding the contract method 0xe67f59a7.
//
// Solidity: function clearTokenConfig(address _token) returns()
func (_VaultV2 *VaultV2TransactorSession) ClearTokenConfig(_token common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.ClearTokenConfig(&_VaultV2.TransactOpts, _token)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_VaultV2 *VaultV2Transactor) DecreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "decreasePosition", _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_VaultV2 *VaultV2Session) DecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.DecreasePosition(&_VaultV2.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_VaultV2 *VaultV2TransactorSession) DecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.DecreasePosition(&_VaultV2.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_VaultV2 *VaultV2Transactor) DirectPoolDeposit(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "directPoolDeposit", _token)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_VaultV2 *VaultV2Session) DirectPoolDeposit(_token common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.DirectPoolDeposit(&_VaultV2.TransactOpts, _token)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_VaultV2 *VaultV2TransactorSession) DirectPoolDeposit(_token common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.DirectPoolDeposit(&_VaultV2.TransactOpts, _token)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_VaultV2 *VaultV2Transactor) IncreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "increasePosition", _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_VaultV2 *VaultV2Session) IncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _VaultV2.Contract.IncreasePosition(&_VaultV2.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_VaultV2 *VaultV2TransactorSession) IncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _VaultV2.Contract.IncreasePosition(&_VaultV2.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _router, address _usdg, address _priceFeed, uint256 _liquidationFeeUsd, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_VaultV2 *VaultV2Transactor) Initialize(opts *bind.TransactOpts, _router common.Address, _usdg common.Address, _priceFeed common.Address, _liquidationFeeUsd *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "initialize", _router, _usdg, _priceFeed, _liquidationFeeUsd, _fundingRateFactor, _stableFundingRateFactor)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _router, address _usdg, address _priceFeed, uint256 _liquidationFeeUsd, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_VaultV2 *VaultV2Session) Initialize(_router common.Address, _usdg common.Address, _priceFeed common.Address, _liquidationFeeUsd *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _VaultV2.Contract.Initialize(&_VaultV2.TransactOpts, _router, _usdg, _priceFeed, _liquidationFeeUsd, _fundingRateFactor, _stableFundingRateFactor)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _router, address _usdg, address _priceFeed, uint256 _liquidationFeeUsd, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_VaultV2 *VaultV2TransactorSession) Initialize(_router common.Address, _usdg common.Address, _priceFeed common.Address, _liquidationFeeUsd *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _VaultV2.Contract.Initialize(&_VaultV2.TransactOpts, _router, _usdg, _priceFeed, _liquidationFeeUsd, _fundingRateFactor, _stableFundingRateFactor)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_VaultV2 *VaultV2Transactor) LiquidatePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "liquidatePosition", _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_VaultV2 *VaultV2Session) LiquidatePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.LiquidatePosition(&_VaultV2.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_VaultV2 *VaultV2TransactorSession) LiquidatePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.LiquidatePosition(&_VaultV2.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// RemoveRouter is a paid mutator transaction binding the contract method 0x6ae0b154.
//
// Solidity: function removeRouter(address _router) returns()
func (_VaultV2 *VaultV2Transactor) RemoveRouter(opts *bind.TransactOpts, _router common.Address) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "removeRouter", _router)
}

// RemoveRouter is a paid mutator transaction binding the contract method 0x6ae0b154.
//
// Solidity: function removeRouter(address _router) returns()
func (_VaultV2 *VaultV2Session) RemoveRouter(_router common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.RemoveRouter(&_VaultV2.TransactOpts, _router)
}

// RemoveRouter is a paid mutator transaction binding the contract method 0x6ae0b154.
//
// Solidity: function removeRouter(address _router) returns()
func (_VaultV2 *VaultV2TransactorSession) RemoveRouter(_router common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.RemoveRouter(&_VaultV2.TransactOpts, _router)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_VaultV2 *VaultV2Transactor) SellUSDG(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "sellUSDG", _token, _receiver)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_VaultV2 *VaultV2Session) SellUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.SellUSDG(&_VaultV2.TransactOpts, _token, _receiver)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_VaultV2 *VaultV2TransactorSession) SellUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.SellUSDG(&_VaultV2.TransactOpts, _token, _receiver)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_VaultV2 *VaultV2Transactor) SetBufferAmount(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "setBufferAmount", _token, _amount)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_VaultV2 *VaultV2Session) SetBufferAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2.Contract.SetBufferAmount(&_VaultV2.TransactOpts, _token, _amount)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_VaultV2 *VaultV2TransactorSession) SetBufferAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2.Contract.SetBufferAmount(&_VaultV2.TransactOpts, _token, _amount)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_VaultV2 *VaultV2Transactor) SetError(opts *bind.TransactOpts, _errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "setError", _errorCode, _error)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_VaultV2 *VaultV2Session) SetError(_errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _VaultV2.Contract.SetError(&_VaultV2.TransactOpts, _errorCode, _error)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_VaultV2 *VaultV2TransactorSession) SetError(_errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _VaultV2.Contract.SetError(&_VaultV2.TransactOpts, _errorCode, _error)
}

// SetErrorController is a paid mutator transaction binding the contract method 0x8f7b8404.
//
// Solidity: function setErrorController(address _errorController) returns()
func (_VaultV2 *VaultV2Transactor) SetErrorController(opts *bind.TransactOpts, _errorController common.Address) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "setErrorController", _errorController)
}

// SetErrorController is a paid mutator transaction binding the contract method 0x8f7b8404.
//
// Solidity: function setErrorController(address _errorController) returns()
func (_VaultV2 *VaultV2Session) SetErrorController(_errorController common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.SetErrorController(&_VaultV2.TransactOpts, _errorController)
}

// SetErrorController is a paid mutator transaction binding the contract method 0x8f7b8404.
//
// Solidity: function setErrorController(address _errorController) returns()
func (_VaultV2 *VaultV2TransactorSession) SetErrorController(_errorController common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.SetErrorController(&_VaultV2.TransactOpts, _errorController)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_VaultV2 *VaultV2Transactor) SetFees(opts *bind.TransactOpts, _taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "setFees", _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_VaultV2 *VaultV2Session) SetFees(_taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _VaultV2.Contract.SetFees(&_VaultV2.TransactOpts, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_VaultV2 *VaultV2TransactorSession) SetFees(_taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _VaultV2.Contract.SetFees(&_VaultV2.TransactOpts, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_VaultV2 *VaultV2Transactor) SetFundingRate(opts *bind.TransactOpts, _fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "setFundingRate", _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_VaultV2 *VaultV2Session) SetFundingRate(_fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _VaultV2.Contract.SetFundingRate(&_VaultV2.TransactOpts, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_VaultV2 *VaultV2TransactorSession) SetFundingRate(_fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _VaultV2.Contract.SetFundingRate(&_VaultV2.TransactOpts, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VaultV2 *VaultV2Transactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VaultV2 *VaultV2Session) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.SetGov(&_VaultV2.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VaultV2 *VaultV2TransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.SetGov(&_VaultV2.TransactOpts, _gov)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_VaultV2 *VaultV2Transactor) SetInManagerMode(opts *bind.TransactOpts, _inManagerMode bool) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "setInManagerMode", _inManagerMode)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_VaultV2 *VaultV2Session) SetInManagerMode(_inManagerMode bool) (*types.Transaction, error) {
	return _VaultV2.Contract.SetInManagerMode(&_VaultV2.TransactOpts, _inManagerMode)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_VaultV2 *VaultV2TransactorSession) SetInManagerMode(_inManagerMode bool) (*types.Transaction, error) {
	return _VaultV2.Contract.SetInManagerMode(&_VaultV2.TransactOpts, _inManagerMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_VaultV2 *VaultV2Transactor) SetInPrivateLiquidationMode(opts *bind.TransactOpts, _inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "setInPrivateLiquidationMode", _inPrivateLiquidationMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_VaultV2 *VaultV2Session) SetInPrivateLiquidationMode(_inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _VaultV2.Contract.SetInPrivateLiquidationMode(&_VaultV2.TransactOpts, _inPrivateLiquidationMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_VaultV2 *VaultV2TransactorSession) SetInPrivateLiquidationMode(_inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _VaultV2.Contract.SetInPrivateLiquidationMode(&_VaultV2.TransactOpts, _inPrivateLiquidationMode)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_VaultV2 *VaultV2Transactor) SetIsLeverageEnabled(opts *bind.TransactOpts, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "setIsLeverageEnabled", _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_VaultV2 *VaultV2Session) SetIsLeverageEnabled(_isLeverageEnabled bool) (*types.Transaction, error) {
	return _VaultV2.Contract.SetIsLeverageEnabled(&_VaultV2.TransactOpts, _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_VaultV2 *VaultV2TransactorSession) SetIsLeverageEnabled(_isLeverageEnabled bool) (*types.Transaction, error) {
	return _VaultV2.Contract.SetIsLeverageEnabled(&_VaultV2.TransactOpts, _isLeverageEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_VaultV2 *VaultV2Transactor) SetIsSwapEnabled(opts *bind.TransactOpts, _isSwapEnabled bool) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "setIsSwapEnabled", _isSwapEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_VaultV2 *VaultV2Session) SetIsSwapEnabled(_isSwapEnabled bool) (*types.Transaction, error) {
	return _VaultV2.Contract.SetIsSwapEnabled(&_VaultV2.TransactOpts, _isSwapEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_VaultV2 *VaultV2TransactorSession) SetIsSwapEnabled(_isSwapEnabled bool) (*types.Transaction, error) {
	return _VaultV2.Contract.SetIsSwapEnabled(&_VaultV2.TransactOpts, _isSwapEnabled)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_VaultV2 *VaultV2Transactor) SetLiquidator(opts *bind.TransactOpts, _liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "setLiquidator", _liquidator, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_VaultV2 *VaultV2Session) SetLiquidator(_liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _VaultV2.Contract.SetLiquidator(&_VaultV2.TransactOpts, _liquidator, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_VaultV2 *VaultV2TransactorSession) SetLiquidator(_liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _VaultV2.Contract.SetLiquidator(&_VaultV2.TransactOpts, _liquidator, _isActive)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_VaultV2 *VaultV2Transactor) SetManager(opts *bind.TransactOpts, _manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "setManager", _manager, _isManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_VaultV2 *VaultV2Session) SetManager(_manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _VaultV2.Contract.SetManager(&_VaultV2.TransactOpts, _manager, _isManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_VaultV2 *VaultV2TransactorSession) SetManager(_manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _VaultV2.Contract.SetManager(&_VaultV2.TransactOpts, _manager, _isManager)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_VaultV2 *VaultV2Transactor) SetMaxGasPrice(opts *bind.TransactOpts, _maxGasPrice *big.Int) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "setMaxGasPrice", _maxGasPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_VaultV2 *VaultV2Session) SetMaxGasPrice(_maxGasPrice *big.Int) (*types.Transaction, error) {
	return _VaultV2.Contract.SetMaxGasPrice(&_VaultV2.TransactOpts, _maxGasPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_VaultV2 *VaultV2TransactorSession) SetMaxGasPrice(_maxGasPrice *big.Int) (*types.Transaction, error) {
	return _VaultV2.Contract.SetMaxGasPrice(&_VaultV2.TransactOpts, _maxGasPrice)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_VaultV2 *VaultV2Transactor) SetMaxGlobalShortSize(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "setMaxGlobalShortSize", _token, _amount)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_VaultV2 *VaultV2Session) SetMaxGlobalShortSize(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2.Contract.SetMaxGlobalShortSize(&_VaultV2.TransactOpts, _token, _amount)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_VaultV2 *VaultV2TransactorSession) SetMaxGlobalShortSize(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2.Contract.SetMaxGlobalShortSize(&_VaultV2.TransactOpts, _token, _amount)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_VaultV2 *VaultV2Transactor) SetMaxLeverage(opts *bind.TransactOpts, _maxLeverage *big.Int) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "setMaxLeverage", _maxLeverage)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_VaultV2 *VaultV2Session) SetMaxLeverage(_maxLeverage *big.Int) (*types.Transaction, error) {
	return _VaultV2.Contract.SetMaxLeverage(&_VaultV2.TransactOpts, _maxLeverage)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_VaultV2 *VaultV2TransactorSession) SetMaxLeverage(_maxLeverage *big.Int) (*types.Transaction, error) {
	return _VaultV2.Contract.SetMaxLeverage(&_VaultV2.TransactOpts, _maxLeverage)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_VaultV2 *VaultV2Transactor) SetPriceFeed(opts *bind.TransactOpts, _priceFeed common.Address) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "setPriceFeed", _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_VaultV2 *VaultV2Session) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.SetPriceFeed(&_VaultV2.TransactOpts, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_VaultV2 *VaultV2TransactorSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.SetPriceFeed(&_VaultV2.TransactOpts, _priceFeed)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_VaultV2 *VaultV2Transactor) SetTokenConfig(opts *bind.TransactOpts, _token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "setTokenConfig", _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_VaultV2 *VaultV2Session) SetTokenConfig(_token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _VaultV2.Contract.SetTokenConfig(&_VaultV2.TransactOpts, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_VaultV2 *VaultV2TransactorSession) SetTokenConfig(_token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _VaultV2.Contract.SetTokenConfig(&_VaultV2.TransactOpts, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetUsdgAmount is a paid mutator transaction binding the contract method 0xd66b000d.
//
// Solidity: function setUsdgAmount(address _token, uint256 _amount) returns()
func (_VaultV2 *VaultV2Transactor) SetUsdgAmount(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "setUsdgAmount", _token, _amount)
}

// SetUsdgAmount is a paid mutator transaction binding the contract method 0xd66b000d.
//
// Solidity: function setUsdgAmount(address _token, uint256 _amount) returns()
func (_VaultV2 *VaultV2Session) SetUsdgAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2.Contract.SetUsdgAmount(&_VaultV2.TransactOpts, _token, _amount)
}

// SetUsdgAmount is a paid mutator transaction binding the contract method 0xd66b000d.
//
// Solidity: function setUsdgAmount(address _token, uint256 _amount) returns()
func (_VaultV2 *VaultV2TransactorSession) SetUsdgAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2.Contract.SetUsdgAmount(&_VaultV2.TransactOpts, _token, _amount)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_VaultV2 *VaultV2Transactor) SetVaultUtils(opts *bind.TransactOpts, _vaultUtils common.Address) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "setVaultUtils", _vaultUtils)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_VaultV2 *VaultV2Session) SetVaultUtils(_vaultUtils common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.SetVaultUtils(&_VaultV2.TransactOpts, _vaultUtils)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_VaultV2 *VaultV2TransactorSession) SetVaultUtils(_vaultUtils common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.SetVaultUtils(&_VaultV2.TransactOpts, _vaultUtils)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_VaultV2 *VaultV2Transactor) Swap(opts *bind.TransactOpts, _tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "swap", _tokenIn, _tokenOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_VaultV2 *VaultV2Session) Swap(_tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.Swap(&_VaultV2.TransactOpts, _tokenIn, _tokenOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_VaultV2 *VaultV2TransactorSession) Swap(_tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.Swap(&_VaultV2.TransactOpts, _tokenIn, _tokenOut, _receiver)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns()
func (_VaultV2 *VaultV2Transactor) UpdateCumulativeFundingRate(opts *bind.TransactOpts, _collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "updateCumulativeFundingRate", _collateralToken, _indexToken)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns()
func (_VaultV2 *VaultV2Session) UpdateCumulativeFundingRate(_collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.UpdateCumulativeFundingRate(&_VaultV2.TransactOpts, _collateralToken, _indexToken)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns()
func (_VaultV2 *VaultV2TransactorSession) UpdateCumulativeFundingRate(_collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.UpdateCumulativeFundingRate(&_VaultV2.TransactOpts, _collateralToken, _indexToken)
}

// UpgradeVault is a paid mutator transaction binding the contract method 0xcea0c328.
//
// Solidity: function upgradeVault(address _newVault, address _token, uint256 _amount) returns()
func (_VaultV2 *VaultV2Transactor) UpgradeVault(opts *bind.TransactOpts, _newVault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "upgradeVault", _newVault, _token, _amount)
}

// UpgradeVault is a paid mutator transaction binding the contract method 0xcea0c328.
//
// Solidity: function upgradeVault(address _newVault, address _token, uint256 _amount) returns()
func (_VaultV2 *VaultV2Session) UpgradeVault(_newVault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2.Contract.UpgradeVault(&_VaultV2.TransactOpts, _newVault, _token, _amount)
}

// UpgradeVault is a paid mutator transaction binding the contract method 0xcea0c328.
//
// Solidity: function upgradeVault(address _newVault, address _token, uint256 _amount) returns()
func (_VaultV2 *VaultV2TransactorSession) UpgradeVault(_newVault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2.Contract.UpgradeVault(&_VaultV2.TransactOpts, _newVault, _token, _amount)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_VaultV2 *VaultV2Transactor) WithdrawFees(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2.contract.Transact(opts, "withdrawFees", _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_VaultV2 *VaultV2Session) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.WithdrawFees(&_VaultV2.TransactOpts, _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_VaultV2 *VaultV2TransactorSession) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2.Contract.WithdrawFees(&_VaultV2.TransactOpts, _token, _receiver)
}

// VaultV2BuyUSDGIterator is returned from FilterBuyUSDG and is used to iterate over the raw logs and unpacked data for BuyUSDG events raised by the VaultV2 contract.
type VaultV2BuyUSDGIterator struct {
	Event *VaultV2BuyUSDG // Event containing the contract specifics and raw log

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
func (it *VaultV2BuyUSDGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2BuyUSDG)
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
		it.Event = new(VaultV2BuyUSDG)
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
func (it *VaultV2BuyUSDGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2BuyUSDGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2BuyUSDG represents a BuyUSDG event raised by the VaultV2 contract.
type VaultV2BuyUSDG struct {
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
func (_VaultV2 *VaultV2Filterer) FilterBuyUSDG(opts *bind.FilterOpts) (*VaultV2BuyUSDGIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "BuyUSDG")
	if err != nil {
		return nil, err
	}
	return &VaultV2BuyUSDGIterator{contract: _VaultV2.contract, event: "BuyUSDG", logs: logs, sub: sub}, nil
}

// WatchBuyUSDG is a free log subscription operation binding the contract event 0xab4c77c74cd32c85f35416cf03e7ce9e2d4387f7b7f2c1f4bf53daaecf8ea72d.
//
// Solidity: event BuyUSDG(address account, address token, uint256 tokenAmount, uint256 usdgAmount, uint256 feeBasisPoints)
func (_VaultV2 *VaultV2Filterer) WatchBuyUSDG(opts *bind.WatchOpts, sink chan<- *VaultV2BuyUSDG) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "BuyUSDG")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2BuyUSDG)
				if err := _VaultV2.contract.UnpackLog(event, "BuyUSDG", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseBuyUSDG(log types.Log) (*VaultV2BuyUSDG, error) {
	event := new(VaultV2BuyUSDG)
	if err := _VaultV2.contract.UnpackLog(event, "BuyUSDG", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2ClosePositionIterator is returned from FilterClosePosition and is used to iterate over the raw logs and unpacked data for ClosePosition events raised by the VaultV2 contract.
type VaultV2ClosePositionIterator struct {
	Event *VaultV2ClosePosition // Event containing the contract specifics and raw log

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
func (it *VaultV2ClosePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2ClosePosition)
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
		it.Event = new(VaultV2ClosePosition)
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
func (it *VaultV2ClosePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2ClosePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2ClosePosition represents a ClosePosition event raised by the VaultV2 contract.
type VaultV2ClosePosition struct {
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
func (_VaultV2 *VaultV2Filterer) FilterClosePosition(opts *bind.FilterOpts) (*VaultV2ClosePositionIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "ClosePosition")
	if err != nil {
		return nil, err
	}
	return &VaultV2ClosePositionIterator{contract: _VaultV2.contract, event: "ClosePosition", logs: logs, sub: sub}, nil
}

// WatchClosePosition is a free log subscription operation binding the contract event 0x73af1d417d82c240fdb6d319b34ad884487c6bf2845d98980cc52ad9171cb455.
//
// Solidity: event ClosePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl)
func (_VaultV2 *VaultV2Filterer) WatchClosePosition(opts *bind.WatchOpts, sink chan<- *VaultV2ClosePosition) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "ClosePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2ClosePosition)
				if err := _VaultV2.contract.UnpackLog(event, "ClosePosition", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseClosePosition(log types.Log) (*VaultV2ClosePosition, error) {
	event := new(VaultV2ClosePosition)
	if err := _VaultV2.contract.UnpackLog(event, "ClosePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2CollectMarginFeesIterator is returned from FilterCollectMarginFees and is used to iterate over the raw logs and unpacked data for CollectMarginFees events raised by the VaultV2 contract.
type VaultV2CollectMarginFeesIterator struct {
	Event *VaultV2CollectMarginFees // Event containing the contract specifics and raw log

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
func (it *VaultV2CollectMarginFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2CollectMarginFees)
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
		it.Event = new(VaultV2CollectMarginFees)
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
func (it *VaultV2CollectMarginFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2CollectMarginFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2CollectMarginFees represents a CollectMarginFees event raised by the VaultV2 contract.
type VaultV2CollectMarginFees struct {
	Token     common.Address
	FeeUsd    *big.Int
	FeeTokens *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectMarginFees is a free log retrieval operation binding the contract event 0x5d0c0019d3d45fadeb74eff9d2c9924d146d000ac6bcf3c28bf0ac3c9baa011a.
//
// Solidity: event CollectMarginFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_VaultV2 *VaultV2Filterer) FilterCollectMarginFees(opts *bind.FilterOpts) (*VaultV2CollectMarginFeesIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "CollectMarginFees")
	if err != nil {
		return nil, err
	}
	return &VaultV2CollectMarginFeesIterator{contract: _VaultV2.contract, event: "CollectMarginFees", logs: logs, sub: sub}, nil
}

// WatchCollectMarginFees is a free log subscription operation binding the contract event 0x5d0c0019d3d45fadeb74eff9d2c9924d146d000ac6bcf3c28bf0ac3c9baa011a.
//
// Solidity: event CollectMarginFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_VaultV2 *VaultV2Filterer) WatchCollectMarginFees(opts *bind.WatchOpts, sink chan<- *VaultV2CollectMarginFees) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "CollectMarginFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2CollectMarginFees)
				if err := _VaultV2.contract.UnpackLog(event, "CollectMarginFees", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseCollectMarginFees(log types.Log) (*VaultV2CollectMarginFees, error) {
	event := new(VaultV2CollectMarginFees)
	if err := _VaultV2.contract.UnpackLog(event, "CollectMarginFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2CollectSwapFeesIterator is returned from FilterCollectSwapFees and is used to iterate over the raw logs and unpacked data for CollectSwapFees events raised by the VaultV2 contract.
type VaultV2CollectSwapFeesIterator struct {
	Event *VaultV2CollectSwapFees // Event containing the contract specifics and raw log

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
func (it *VaultV2CollectSwapFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2CollectSwapFees)
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
		it.Event = new(VaultV2CollectSwapFees)
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
func (it *VaultV2CollectSwapFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2CollectSwapFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2CollectSwapFees represents a CollectSwapFees event raised by the VaultV2 contract.
type VaultV2CollectSwapFees struct {
	Token     common.Address
	FeeUsd    *big.Int
	FeeTokens *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectSwapFees is a free log retrieval operation binding the contract event 0x47cd9dda0e50ce30bcaaacd0488452b596221c07ac402a581cfae4d3933cac2b.
//
// Solidity: event CollectSwapFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_VaultV2 *VaultV2Filterer) FilterCollectSwapFees(opts *bind.FilterOpts) (*VaultV2CollectSwapFeesIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "CollectSwapFees")
	if err != nil {
		return nil, err
	}
	return &VaultV2CollectSwapFeesIterator{contract: _VaultV2.contract, event: "CollectSwapFees", logs: logs, sub: sub}, nil
}

// WatchCollectSwapFees is a free log subscription operation binding the contract event 0x47cd9dda0e50ce30bcaaacd0488452b596221c07ac402a581cfae4d3933cac2b.
//
// Solidity: event CollectSwapFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_VaultV2 *VaultV2Filterer) WatchCollectSwapFees(opts *bind.WatchOpts, sink chan<- *VaultV2CollectSwapFees) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "CollectSwapFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2CollectSwapFees)
				if err := _VaultV2.contract.UnpackLog(event, "CollectSwapFees", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseCollectSwapFees(log types.Log) (*VaultV2CollectSwapFees, error) {
	event := new(VaultV2CollectSwapFees)
	if err := _VaultV2.contract.UnpackLog(event, "CollectSwapFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2DecreaseGuaranteedUsdIterator is returned from FilterDecreaseGuaranteedUsd and is used to iterate over the raw logs and unpacked data for DecreaseGuaranteedUsd events raised by the VaultV2 contract.
type VaultV2DecreaseGuaranteedUsdIterator struct {
	Event *VaultV2DecreaseGuaranteedUsd // Event containing the contract specifics and raw log

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
func (it *VaultV2DecreaseGuaranteedUsdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2DecreaseGuaranteedUsd)
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
		it.Event = new(VaultV2DecreaseGuaranteedUsd)
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
func (it *VaultV2DecreaseGuaranteedUsdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2DecreaseGuaranteedUsdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2DecreaseGuaranteedUsd represents a DecreaseGuaranteedUsd event raised by the VaultV2 contract.
type VaultV2DecreaseGuaranteedUsd struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreaseGuaranteedUsd is a free log retrieval operation binding the contract event 0x34e07158b9db50df5613e591c44ea2ebc82834eff4a4dc3a46e000e608261d68.
//
// Solidity: event DecreaseGuaranteedUsd(address token, uint256 amount)
func (_VaultV2 *VaultV2Filterer) FilterDecreaseGuaranteedUsd(opts *bind.FilterOpts) (*VaultV2DecreaseGuaranteedUsdIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "DecreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return &VaultV2DecreaseGuaranteedUsdIterator{contract: _VaultV2.contract, event: "DecreaseGuaranteedUsd", logs: logs, sub: sub}, nil
}

// WatchDecreaseGuaranteedUsd is a free log subscription operation binding the contract event 0x34e07158b9db50df5613e591c44ea2ebc82834eff4a4dc3a46e000e608261d68.
//
// Solidity: event DecreaseGuaranteedUsd(address token, uint256 amount)
func (_VaultV2 *VaultV2Filterer) WatchDecreaseGuaranteedUsd(opts *bind.WatchOpts, sink chan<- *VaultV2DecreaseGuaranteedUsd) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "DecreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2DecreaseGuaranteedUsd)
				if err := _VaultV2.contract.UnpackLog(event, "DecreaseGuaranteedUsd", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseDecreaseGuaranteedUsd(log types.Log) (*VaultV2DecreaseGuaranteedUsd, error) {
	event := new(VaultV2DecreaseGuaranteedUsd)
	if err := _VaultV2.contract.UnpackLog(event, "DecreaseGuaranteedUsd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2DecreasePoolAmountIterator is returned from FilterDecreasePoolAmount and is used to iterate over the raw logs and unpacked data for DecreasePoolAmount events raised by the VaultV2 contract.
type VaultV2DecreasePoolAmountIterator struct {
	Event *VaultV2DecreasePoolAmount // Event containing the contract specifics and raw log

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
func (it *VaultV2DecreasePoolAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2DecreasePoolAmount)
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
		it.Event = new(VaultV2DecreasePoolAmount)
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
func (it *VaultV2DecreasePoolAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2DecreasePoolAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2DecreasePoolAmount represents a DecreasePoolAmount event raised by the VaultV2 contract.
type VaultV2DecreasePoolAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreasePoolAmount is a free log retrieval operation binding the contract event 0x112726233fbeaeed0f5b1dba5cb0b2b81883dee49fb35ff99fd98ed9f6d31eb0.
//
// Solidity: event DecreasePoolAmount(address token, uint256 amount)
func (_VaultV2 *VaultV2Filterer) FilterDecreasePoolAmount(opts *bind.FilterOpts) (*VaultV2DecreasePoolAmountIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "DecreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return &VaultV2DecreasePoolAmountIterator{contract: _VaultV2.contract, event: "DecreasePoolAmount", logs: logs, sub: sub}, nil
}

// WatchDecreasePoolAmount is a free log subscription operation binding the contract event 0x112726233fbeaeed0f5b1dba5cb0b2b81883dee49fb35ff99fd98ed9f6d31eb0.
//
// Solidity: event DecreasePoolAmount(address token, uint256 amount)
func (_VaultV2 *VaultV2Filterer) WatchDecreasePoolAmount(opts *bind.WatchOpts, sink chan<- *VaultV2DecreasePoolAmount) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "DecreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2DecreasePoolAmount)
				if err := _VaultV2.contract.UnpackLog(event, "DecreasePoolAmount", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseDecreasePoolAmount(log types.Log) (*VaultV2DecreasePoolAmount, error) {
	event := new(VaultV2DecreasePoolAmount)
	if err := _VaultV2.contract.UnpackLog(event, "DecreasePoolAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2DecreasePositionIterator is returned from FilterDecreasePosition and is used to iterate over the raw logs and unpacked data for DecreasePosition events raised by the VaultV2 contract.
type VaultV2DecreasePositionIterator struct {
	Event *VaultV2DecreasePosition // Event containing the contract specifics and raw log

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
func (it *VaultV2DecreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2DecreasePosition)
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
		it.Event = new(VaultV2DecreasePosition)
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
func (it *VaultV2DecreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2DecreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2DecreasePosition represents a DecreasePosition event raised by the VaultV2 contract.
type VaultV2DecreasePosition struct {
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
func (_VaultV2 *VaultV2Filterer) FilterDecreasePosition(opts *bind.FilterOpts) (*VaultV2DecreasePositionIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "DecreasePosition")
	if err != nil {
		return nil, err
	}
	return &VaultV2DecreasePositionIterator{contract: _VaultV2.contract, event: "DecreasePosition", logs: logs, sub: sub}, nil
}

// WatchDecreasePosition is a free log subscription operation binding the contract event 0x93d75d64d1f84fc6f430a64fc578bdd4c1e090e90ea2d51773e626d19de56d30.
//
// Solidity: event DecreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_VaultV2 *VaultV2Filterer) WatchDecreasePosition(opts *bind.WatchOpts, sink chan<- *VaultV2DecreasePosition) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "DecreasePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2DecreasePosition)
				if err := _VaultV2.contract.UnpackLog(event, "DecreasePosition", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseDecreasePosition(log types.Log) (*VaultV2DecreasePosition, error) {
	event := new(VaultV2DecreasePosition)
	if err := _VaultV2.contract.UnpackLog(event, "DecreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2DecreaseReservedAmountIterator is returned from FilterDecreaseReservedAmount and is used to iterate over the raw logs and unpacked data for DecreaseReservedAmount events raised by the VaultV2 contract.
type VaultV2DecreaseReservedAmountIterator struct {
	Event *VaultV2DecreaseReservedAmount // Event containing the contract specifics and raw log

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
func (it *VaultV2DecreaseReservedAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2DecreaseReservedAmount)
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
		it.Event = new(VaultV2DecreaseReservedAmount)
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
func (it *VaultV2DecreaseReservedAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2DecreaseReservedAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2DecreaseReservedAmount represents a DecreaseReservedAmount event raised by the VaultV2 contract.
type VaultV2DecreaseReservedAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreaseReservedAmount is a free log retrieval operation binding the contract event 0x533cb5ed32be6a90284e96b5747a1bfc2d38fdb5768a6b5f67ff7d62144ed67b.
//
// Solidity: event DecreaseReservedAmount(address token, uint256 amount)
func (_VaultV2 *VaultV2Filterer) FilterDecreaseReservedAmount(opts *bind.FilterOpts) (*VaultV2DecreaseReservedAmountIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "DecreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return &VaultV2DecreaseReservedAmountIterator{contract: _VaultV2.contract, event: "DecreaseReservedAmount", logs: logs, sub: sub}, nil
}

// WatchDecreaseReservedAmount is a free log subscription operation binding the contract event 0x533cb5ed32be6a90284e96b5747a1bfc2d38fdb5768a6b5f67ff7d62144ed67b.
//
// Solidity: event DecreaseReservedAmount(address token, uint256 amount)
func (_VaultV2 *VaultV2Filterer) WatchDecreaseReservedAmount(opts *bind.WatchOpts, sink chan<- *VaultV2DecreaseReservedAmount) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "DecreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2DecreaseReservedAmount)
				if err := _VaultV2.contract.UnpackLog(event, "DecreaseReservedAmount", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseDecreaseReservedAmount(log types.Log) (*VaultV2DecreaseReservedAmount, error) {
	event := new(VaultV2DecreaseReservedAmount)
	if err := _VaultV2.contract.UnpackLog(event, "DecreaseReservedAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2DecreaseUsdgAmountIterator is returned from FilterDecreaseUsdgAmount and is used to iterate over the raw logs and unpacked data for DecreaseUsdgAmount events raised by the VaultV2 contract.
type VaultV2DecreaseUsdgAmountIterator struct {
	Event *VaultV2DecreaseUsdgAmount // Event containing the contract specifics and raw log

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
func (it *VaultV2DecreaseUsdgAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2DecreaseUsdgAmount)
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
		it.Event = new(VaultV2DecreaseUsdgAmount)
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
func (it *VaultV2DecreaseUsdgAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2DecreaseUsdgAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2DecreaseUsdgAmount represents a DecreaseUsdgAmount event raised by the VaultV2 contract.
type VaultV2DecreaseUsdgAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreaseUsdgAmount is a free log retrieval operation binding the contract event 0xe1e812596aac93a06ecc4ca627014d18e30f5c33b825160cc9d5c0ba61e45227.
//
// Solidity: event DecreaseUsdgAmount(address token, uint256 amount)
func (_VaultV2 *VaultV2Filterer) FilterDecreaseUsdgAmount(opts *bind.FilterOpts) (*VaultV2DecreaseUsdgAmountIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "DecreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return &VaultV2DecreaseUsdgAmountIterator{contract: _VaultV2.contract, event: "DecreaseUsdgAmount", logs: logs, sub: sub}, nil
}

// WatchDecreaseUsdgAmount is a free log subscription operation binding the contract event 0xe1e812596aac93a06ecc4ca627014d18e30f5c33b825160cc9d5c0ba61e45227.
//
// Solidity: event DecreaseUsdgAmount(address token, uint256 amount)
func (_VaultV2 *VaultV2Filterer) WatchDecreaseUsdgAmount(opts *bind.WatchOpts, sink chan<- *VaultV2DecreaseUsdgAmount) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "DecreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2DecreaseUsdgAmount)
				if err := _VaultV2.contract.UnpackLog(event, "DecreaseUsdgAmount", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseDecreaseUsdgAmount(log types.Log) (*VaultV2DecreaseUsdgAmount, error) {
	event := new(VaultV2DecreaseUsdgAmount)
	if err := _VaultV2.contract.UnpackLog(event, "DecreaseUsdgAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2DirectPoolDepositIterator is returned from FilterDirectPoolDeposit and is used to iterate over the raw logs and unpacked data for DirectPoolDeposit events raised by the VaultV2 contract.
type VaultV2DirectPoolDepositIterator struct {
	Event *VaultV2DirectPoolDeposit // Event containing the contract specifics and raw log

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
func (it *VaultV2DirectPoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2DirectPoolDeposit)
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
		it.Event = new(VaultV2DirectPoolDeposit)
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
func (it *VaultV2DirectPoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2DirectPoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2DirectPoolDeposit represents a DirectPoolDeposit event raised by the VaultV2 contract.
type VaultV2DirectPoolDeposit struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDirectPoolDeposit is a free log retrieval operation binding the contract event 0xa5a389190ebf6170a133bda5c769b77f4d6715b8aa172ec0ddf8473d0b4944bd.
//
// Solidity: event DirectPoolDeposit(address token, uint256 amount)
func (_VaultV2 *VaultV2Filterer) FilterDirectPoolDeposit(opts *bind.FilterOpts) (*VaultV2DirectPoolDepositIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "DirectPoolDeposit")
	if err != nil {
		return nil, err
	}
	return &VaultV2DirectPoolDepositIterator{contract: _VaultV2.contract, event: "DirectPoolDeposit", logs: logs, sub: sub}, nil
}

// WatchDirectPoolDeposit is a free log subscription operation binding the contract event 0xa5a389190ebf6170a133bda5c769b77f4d6715b8aa172ec0ddf8473d0b4944bd.
//
// Solidity: event DirectPoolDeposit(address token, uint256 amount)
func (_VaultV2 *VaultV2Filterer) WatchDirectPoolDeposit(opts *bind.WatchOpts, sink chan<- *VaultV2DirectPoolDeposit) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "DirectPoolDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2DirectPoolDeposit)
				if err := _VaultV2.contract.UnpackLog(event, "DirectPoolDeposit", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseDirectPoolDeposit(log types.Log) (*VaultV2DirectPoolDeposit, error) {
	event := new(VaultV2DirectPoolDeposit)
	if err := _VaultV2.contract.UnpackLog(event, "DirectPoolDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2IncreaseGuaranteedUsdIterator is returned from FilterIncreaseGuaranteedUsd and is used to iterate over the raw logs and unpacked data for IncreaseGuaranteedUsd events raised by the VaultV2 contract.
type VaultV2IncreaseGuaranteedUsdIterator struct {
	Event *VaultV2IncreaseGuaranteedUsd // Event containing the contract specifics and raw log

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
func (it *VaultV2IncreaseGuaranteedUsdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2IncreaseGuaranteedUsd)
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
		it.Event = new(VaultV2IncreaseGuaranteedUsd)
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
func (it *VaultV2IncreaseGuaranteedUsdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2IncreaseGuaranteedUsdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2IncreaseGuaranteedUsd represents a IncreaseGuaranteedUsd event raised by the VaultV2 contract.
type VaultV2IncreaseGuaranteedUsd struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreaseGuaranteedUsd is a free log retrieval operation binding the contract event 0xd9d4761f75e0d0103b5cbeab941eeb443d7a56a35b5baf2a0787c03f03f4e474.
//
// Solidity: event IncreaseGuaranteedUsd(address token, uint256 amount)
func (_VaultV2 *VaultV2Filterer) FilterIncreaseGuaranteedUsd(opts *bind.FilterOpts) (*VaultV2IncreaseGuaranteedUsdIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "IncreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return &VaultV2IncreaseGuaranteedUsdIterator{contract: _VaultV2.contract, event: "IncreaseGuaranteedUsd", logs: logs, sub: sub}, nil
}

// WatchIncreaseGuaranteedUsd is a free log subscription operation binding the contract event 0xd9d4761f75e0d0103b5cbeab941eeb443d7a56a35b5baf2a0787c03f03f4e474.
//
// Solidity: event IncreaseGuaranteedUsd(address token, uint256 amount)
func (_VaultV2 *VaultV2Filterer) WatchIncreaseGuaranteedUsd(opts *bind.WatchOpts, sink chan<- *VaultV2IncreaseGuaranteedUsd) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "IncreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2IncreaseGuaranteedUsd)
				if err := _VaultV2.contract.UnpackLog(event, "IncreaseGuaranteedUsd", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseIncreaseGuaranteedUsd(log types.Log) (*VaultV2IncreaseGuaranteedUsd, error) {
	event := new(VaultV2IncreaseGuaranteedUsd)
	if err := _VaultV2.contract.UnpackLog(event, "IncreaseGuaranteedUsd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2IncreasePoolAmountIterator is returned from FilterIncreasePoolAmount and is used to iterate over the raw logs and unpacked data for IncreasePoolAmount events raised by the VaultV2 contract.
type VaultV2IncreasePoolAmountIterator struct {
	Event *VaultV2IncreasePoolAmount // Event containing the contract specifics and raw log

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
func (it *VaultV2IncreasePoolAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2IncreasePoolAmount)
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
		it.Event = new(VaultV2IncreasePoolAmount)
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
func (it *VaultV2IncreasePoolAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2IncreasePoolAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2IncreasePoolAmount represents a IncreasePoolAmount event raised by the VaultV2 contract.
type VaultV2IncreasePoolAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreasePoolAmount is a free log retrieval operation binding the contract event 0x976177fbe09a15e5e43f848844963a42b41ef919ef17ff21a17a5421de8f4737.
//
// Solidity: event IncreasePoolAmount(address token, uint256 amount)
func (_VaultV2 *VaultV2Filterer) FilterIncreasePoolAmount(opts *bind.FilterOpts) (*VaultV2IncreasePoolAmountIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "IncreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return &VaultV2IncreasePoolAmountIterator{contract: _VaultV2.contract, event: "IncreasePoolAmount", logs: logs, sub: sub}, nil
}

// WatchIncreasePoolAmount is a free log subscription operation binding the contract event 0x976177fbe09a15e5e43f848844963a42b41ef919ef17ff21a17a5421de8f4737.
//
// Solidity: event IncreasePoolAmount(address token, uint256 amount)
func (_VaultV2 *VaultV2Filterer) WatchIncreasePoolAmount(opts *bind.WatchOpts, sink chan<- *VaultV2IncreasePoolAmount) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "IncreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2IncreasePoolAmount)
				if err := _VaultV2.contract.UnpackLog(event, "IncreasePoolAmount", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseIncreasePoolAmount(log types.Log) (*VaultV2IncreasePoolAmount, error) {
	event := new(VaultV2IncreasePoolAmount)
	if err := _VaultV2.contract.UnpackLog(event, "IncreasePoolAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2IncreasePositionIterator is returned from FilterIncreasePosition and is used to iterate over the raw logs and unpacked data for IncreasePosition events raised by the VaultV2 contract.
type VaultV2IncreasePositionIterator struct {
	Event *VaultV2IncreasePosition // Event containing the contract specifics and raw log

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
func (it *VaultV2IncreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2IncreasePosition)
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
		it.Event = new(VaultV2IncreasePosition)
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
func (it *VaultV2IncreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2IncreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2IncreasePosition represents a IncreasePosition event raised by the VaultV2 contract.
type VaultV2IncreasePosition struct {
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
func (_VaultV2 *VaultV2Filterer) FilterIncreasePosition(opts *bind.FilterOpts) (*VaultV2IncreasePositionIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "IncreasePosition")
	if err != nil {
		return nil, err
	}
	return &VaultV2IncreasePositionIterator{contract: _VaultV2.contract, event: "IncreasePosition", logs: logs, sub: sub}, nil
}

// WatchIncreasePosition is a free log subscription operation binding the contract event 0x2fe68525253654c21998f35787a8d0f361905ef647c854092430ab65f2f15022.
//
// Solidity: event IncreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_VaultV2 *VaultV2Filterer) WatchIncreasePosition(opts *bind.WatchOpts, sink chan<- *VaultV2IncreasePosition) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "IncreasePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2IncreasePosition)
				if err := _VaultV2.contract.UnpackLog(event, "IncreasePosition", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseIncreasePosition(log types.Log) (*VaultV2IncreasePosition, error) {
	event := new(VaultV2IncreasePosition)
	if err := _VaultV2.contract.UnpackLog(event, "IncreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2IncreaseReservedAmountIterator is returned from FilterIncreaseReservedAmount and is used to iterate over the raw logs and unpacked data for IncreaseReservedAmount events raised by the VaultV2 contract.
type VaultV2IncreaseReservedAmountIterator struct {
	Event *VaultV2IncreaseReservedAmount // Event containing the contract specifics and raw log

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
func (it *VaultV2IncreaseReservedAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2IncreaseReservedAmount)
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
		it.Event = new(VaultV2IncreaseReservedAmount)
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
func (it *VaultV2IncreaseReservedAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2IncreaseReservedAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2IncreaseReservedAmount represents a IncreaseReservedAmount event raised by the VaultV2 contract.
type VaultV2IncreaseReservedAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreaseReservedAmount is a free log retrieval operation binding the contract event 0xaa5649d82f5462be9d19b0f2b31a59b2259950a6076550bac9f3a1c07db9f66d.
//
// Solidity: event IncreaseReservedAmount(address token, uint256 amount)
func (_VaultV2 *VaultV2Filterer) FilterIncreaseReservedAmount(opts *bind.FilterOpts) (*VaultV2IncreaseReservedAmountIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "IncreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return &VaultV2IncreaseReservedAmountIterator{contract: _VaultV2.contract, event: "IncreaseReservedAmount", logs: logs, sub: sub}, nil
}

// WatchIncreaseReservedAmount is a free log subscription operation binding the contract event 0xaa5649d82f5462be9d19b0f2b31a59b2259950a6076550bac9f3a1c07db9f66d.
//
// Solidity: event IncreaseReservedAmount(address token, uint256 amount)
func (_VaultV2 *VaultV2Filterer) WatchIncreaseReservedAmount(opts *bind.WatchOpts, sink chan<- *VaultV2IncreaseReservedAmount) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "IncreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2IncreaseReservedAmount)
				if err := _VaultV2.contract.UnpackLog(event, "IncreaseReservedAmount", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseIncreaseReservedAmount(log types.Log) (*VaultV2IncreaseReservedAmount, error) {
	event := new(VaultV2IncreaseReservedAmount)
	if err := _VaultV2.contract.UnpackLog(event, "IncreaseReservedAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2IncreaseUsdgAmountIterator is returned from FilterIncreaseUsdgAmount and is used to iterate over the raw logs and unpacked data for IncreaseUsdgAmount events raised by the VaultV2 contract.
type VaultV2IncreaseUsdgAmountIterator struct {
	Event *VaultV2IncreaseUsdgAmount // Event containing the contract specifics and raw log

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
func (it *VaultV2IncreaseUsdgAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2IncreaseUsdgAmount)
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
		it.Event = new(VaultV2IncreaseUsdgAmount)
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
func (it *VaultV2IncreaseUsdgAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2IncreaseUsdgAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2IncreaseUsdgAmount represents a IncreaseUsdgAmount event raised by the VaultV2 contract.
type VaultV2IncreaseUsdgAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreaseUsdgAmount is a free log retrieval operation binding the contract event 0x64243679a443432e2293343b77d411ff6144370404618f00ca0d2025d9ca9882.
//
// Solidity: event IncreaseUsdgAmount(address token, uint256 amount)
func (_VaultV2 *VaultV2Filterer) FilterIncreaseUsdgAmount(opts *bind.FilterOpts) (*VaultV2IncreaseUsdgAmountIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "IncreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return &VaultV2IncreaseUsdgAmountIterator{contract: _VaultV2.contract, event: "IncreaseUsdgAmount", logs: logs, sub: sub}, nil
}

// WatchIncreaseUsdgAmount is a free log subscription operation binding the contract event 0x64243679a443432e2293343b77d411ff6144370404618f00ca0d2025d9ca9882.
//
// Solidity: event IncreaseUsdgAmount(address token, uint256 amount)
func (_VaultV2 *VaultV2Filterer) WatchIncreaseUsdgAmount(opts *bind.WatchOpts, sink chan<- *VaultV2IncreaseUsdgAmount) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "IncreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2IncreaseUsdgAmount)
				if err := _VaultV2.contract.UnpackLog(event, "IncreaseUsdgAmount", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseIncreaseUsdgAmount(log types.Log) (*VaultV2IncreaseUsdgAmount, error) {
	event := new(VaultV2IncreaseUsdgAmount)
	if err := _VaultV2.contract.UnpackLog(event, "IncreaseUsdgAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2LiquidatePositionIterator is returned from FilterLiquidatePosition and is used to iterate over the raw logs and unpacked data for LiquidatePosition events raised by the VaultV2 contract.
type VaultV2LiquidatePositionIterator struct {
	Event *VaultV2LiquidatePosition // Event containing the contract specifics and raw log

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
func (it *VaultV2LiquidatePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2LiquidatePosition)
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
		it.Event = new(VaultV2LiquidatePosition)
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
func (it *VaultV2LiquidatePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2LiquidatePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2LiquidatePosition represents a LiquidatePosition event raised by the VaultV2 contract.
type VaultV2LiquidatePosition struct {
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
func (_VaultV2 *VaultV2Filterer) FilterLiquidatePosition(opts *bind.FilterOpts) (*VaultV2LiquidatePositionIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "LiquidatePosition")
	if err != nil {
		return nil, err
	}
	return &VaultV2LiquidatePositionIterator{contract: _VaultV2.contract, event: "LiquidatePosition", logs: logs, sub: sub}, nil
}

// WatchLiquidatePosition is a free log subscription operation binding the contract event 0x2e1f85a64a2f22cf2f0c42584e7c919ed4abe8d53675cff0f62bf1e95a1c676f.
//
// Solidity: event LiquidatePosition(bytes32 key, address account, address collateralToken, address indexToken, bool isLong, uint256 size, uint256 collateral, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_VaultV2 *VaultV2Filterer) WatchLiquidatePosition(opts *bind.WatchOpts, sink chan<- *VaultV2LiquidatePosition) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "LiquidatePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2LiquidatePosition)
				if err := _VaultV2.contract.UnpackLog(event, "LiquidatePosition", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseLiquidatePosition(log types.Log) (*VaultV2LiquidatePosition, error) {
	event := new(VaultV2LiquidatePosition)
	if err := _VaultV2.contract.UnpackLog(event, "LiquidatePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2SellUSDGIterator is returned from FilterSellUSDG and is used to iterate over the raw logs and unpacked data for SellUSDG events raised by the VaultV2 contract.
type VaultV2SellUSDGIterator struct {
	Event *VaultV2SellUSDG // Event containing the contract specifics and raw log

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
func (it *VaultV2SellUSDGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2SellUSDG)
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
		it.Event = new(VaultV2SellUSDG)
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
func (it *VaultV2SellUSDGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2SellUSDGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2SellUSDG represents a SellUSDG event raised by the VaultV2 contract.
type VaultV2SellUSDG struct {
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
func (_VaultV2 *VaultV2Filterer) FilterSellUSDG(opts *bind.FilterOpts) (*VaultV2SellUSDGIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "SellUSDG")
	if err != nil {
		return nil, err
	}
	return &VaultV2SellUSDGIterator{contract: _VaultV2.contract, event: "SellUSDG", logs: logs, sub: sub}, nil
}

// WatchSellUSDG is a free log subscription operation binding the contract event 0xd732b7828fa6cee72c285eac756fc66a7477e3dc22e22e7c432f1c265d40b483.
//
// Solidity: event SellUSDG(address account, address token, uint256 usdgAmount, uint256 tokenAmount, uint256 feeBasisPoints)
func (_VaultV2 *VaultV2Filterer) WatchSellUSDG(opts *bind.WatchOpts, sink chan<- *VaultV2SellUSDG) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "SellUSDG")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2SellUSDG)
				if err := _VaultV2.contract.UnpackLog(event, "SellUSDG", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseSellUSDG(log types.Log) (*VaultV2SellUSDG, error) {
	event := new(VaultV2SellUSDG)
	if err := _VaultV2.contract.UnpackLog(event, "SellUSDG", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2SwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the VaultV2 contract.
type VaultV2SwapIterator struct {
	Event *VaultV2Swap // Event containing the contract specifics and raw log

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
func (it *VaultV2SwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2Swap)
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
		it.Event = new(VaultV2Swap)
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
func (it *VaultV2SwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2SwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2Swap represents a Swap event raised by the VaultV2 contract.
type VaultV2Swap struct {
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
func (_VaultV2 *VaultV2Filterer) FilterSwap(opts *bind.FilterOpts) (*VaultV2SwapIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &VaultV2SwapIterator{contract: _VaultV2.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x0874b2d545cb271cdbda4e093020c452328b24af12382ed62c4d00f5c26709db.
//
// Solidity: event Swap(address account, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, uint256 amountOutAfterFees, uint256 feeBasisPoints)
func (_VaultV2 *VaultV2Filterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *VaultV2Swap) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2Swap)
				if err := _VaultV2.contract.UnpackLog(event, "Swap", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseSwap(log types.Log) (*VaultV2Swap, error) {
	event := new(VaultV2Swap)
	if err := _VaultV2.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2UpdateFundingRateIterator is returned from FilterUpdateFundingRate and is used to iterate over the raw logs and unpacked data for UpdateFundingRate events raised by the VaultV2 contract.
type VaultV2UpdateFundingRateIterator struct {
	Event *VaultV2UpdateFundingRate // Event containing the contract specifics and raw log

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
func (it *VaultV2UpdateFundingRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2UpdateFundingRate)
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
		it.Event = new(VaultV2UpdateFundingRate)
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
func (it *VaultV2UpdateFundingRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2UpdateFundingRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2UpdateFundingRate represents a UpdateFundingRate event raised by the VaultV2 contract.
type VaultV2UpdateFundingRate struct {
	Token       common.Address
	FundingRate *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateFundingRate is a free log retrieval operation binding the contract event 0xa146fc154e1913322e9817d49f0d5c37466c24326e15de10e739a948be815eab.
//
// Solidity: event UpdateFundingRate(address token, uint256 fundingRate)
func (_VaultV2 *VaultV2Filterer) FilterUpdateFundingRate(opts *bind.FilterOpts) (*VaultV2UpdateFundingRateIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "UpdateFundingRate")
	if err != nil {
		return nil, err
	}
	return &VaultV2UpdateFundingRateIterator{contract: _VaultV2.contract, event: "UpdateFundingRate", logs: logs, sub: sub}, nil
}

// WatchUpdateFundingRate is a free log subscription operation binding the contract event 0xa146fc154e1913322e9817d49f0d5c37466c24326e15de10e739a948be815eab.
//
// Solidity: event UpdateFundingRate(address token, uint256 fundingRate)
func (_VaultV2 *VaultV2Filterer) WatchUpdateFundingRate(opts *bind.WatchOpts, sink chan<- *VaultV2UpdateFundingRate) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "UpdateFundingRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2UpdateFundingRate)
				if err := _VaultV2.contract.UnpackLog(event, "UpdateFundingRate", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseUpdateFundingRate(log types.Log) (*VaultV2UpdateFundingRate, error) {
	event := new(VaultV2UpdateFundingRate)
	if err := _VaultV2.contract.UnpackLog(event, "UpdateFundingRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2UpdatePnlIterator is returned from FilterUpdatePnl and is used to iterate over the raw logs and unpacked data for UpdatePnl events raised by the VaultV2 contract.
type VaultV2UpdatePnlIterator struct {
	Event *VaultV2UpdatePnl // Event containing the contract specifics and raw log

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
func (it *VaultV2UpdatePnlIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2UpdatePnl)
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
		it.Event = new(VaultV2UpdatePnl)
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
func (it *VaultV2UpdatePnlIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2UpdatePnlIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2UpdatePnl represents a UpdatePnl event raised by the VaultV2 contract.
type VaultV2UpdatePnl struct {
	Key       [32]byte
	HasProfit bool
	Delta     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatePnl is a free log retrieval operation binding the contract event 0x3ff41bdde87755b687ae83d0221a232b6be51a803330ed9661c1b5d0105e0d8a.
//
// Solidity: event UpdatePnl(bytes32 key, bool hasProfit, uint256 delta)
func (_VaultV2 *VaultV2Filterer) FilterUpdatePnl(opts *bind.FilterOpts) (*VaultV2UpdatePnlIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "UpdatePnl")
	if err != nil {
		return nil, err
	}
	return &VaultV2UpdatePnlIterator{contract: _VaultV2.contract, event: "UpdatePnl", logs: logs, sub: sub}, nil
}

// WatchUpdatePnl is a free log subscription operation binding the contract event 0x3ff41bdde87755b687ae83d0221a232b6be51a803330ed9661c1b5d0105e0d8a.
//
// Solidity: event UpdatePnl(bytes32 key, bool hasProfit, uint256 delta)
func (_VaultV2 *VaultV2Filterer) WatchUpdatePnl(opts *bind.WatchOpts, sink chan<- *VaultV2UpdatePnl) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "UpdatePnl")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2UpdatePnl)
				if err := _VaultV2.contract.UnpackLog(event, "UpdatePnl", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseUpdatePnl(log types.Log) (*VaultV2UpdatePnl, error) {
	event := new(VaultV2UpdatePnl)
	if err := _VaultV2.contract.UnpackLog(event, "UpdatePnl", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2UpdatePositionIterator is returned from FilterUpdatePosition and is used to iterate over the raw logs and unpacked data for UpdatePosition events raised by the VaultV2 contract.
type VaultV2UpdatePositionIterator struct {
	Event *VaultV2UpdatePosition // Event containing the contract specifics and raw log

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
func (it *VaultV2UpdatePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2UpdatePosition)
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
		it.Event = new(VaultV2UpdatePosition)
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
func (it *VaultV2UpdatePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2UpdatePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2UpdatePosition represents a UpdatePosition event raised by the VaultV2 contract.
type VaultV2UpdatePosition struct {
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
func (_VaultV2 *VaultV2Filterer) FilterUpdatePosition(opts *bind.FilterOpts) (*VaultV2UpdatePositionIterator, error) {

	logs, sub, err := _VaultV2.contract.FilterLogs(opts, "UpdatePosition")
	if err != nil {
		return nil, err
	}
	return &VaultV2UpdatePositionIterator{contract: _VaultV2.contract, event: "UpdatePosition", logs: logs, sub: sub}, nil
}

// WatchUpdatePosition is a free log subscription operation binding the contract event 0x20853733b590dce729d9f4628682ebd9a34d2354e72679e66f43a008fc03b773.
//
// Solidity: event UpdatePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_VaultV2 *VaultV2Filterer) WatchUpdatePosition(opts *bind.WatchOpts, sink chan<- *VaultV2UpdatePosition) (event.Subscription, error) {

	logs, sub, err := _VaultV2.contract.WatchLogs(opts, "UpdatePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2UpdatePosition)
				if err := _VaultV2.contract.UnpackLog(event, "UpdatePosition", log); err != nil {
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
func (_VaultV2 *VaultV2Filterer) ParseUpdatePosition(log types.Log) (*VaultV2UpdatePosition, error) {
	event := new(VaultV2UpdatePosition)
	if err := _VaultV2.contract.UnpackLog(event, "UpdatePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

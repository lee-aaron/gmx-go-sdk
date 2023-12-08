// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VaultV2b

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

// VaultV2bMetaData contains all meta data concerning the VaultV2b contract.
var VaultV2bMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BASIS_POINTS_DIVISOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FUNDING_RATE_PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FEE_BASIS_POINTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FUNDING_RATE_FACTOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_LIQUIDATION_FEE_USD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_FUNDING_RATE_INTERVAL\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MIN_LEVERAGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PRICE_PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"USDG_DECIMALS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addRouter\",\"inputs\":[{\"name\":\"_router\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"adjustForDecimals\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokenDiv\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenMul\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allWhitelistedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allWhitelistedTokensLength\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approvedRouters\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bufferAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"buyUSDG\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clearTokenConfig\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cumulativeFundingRates\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreasePosition\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"directPoolDeposit\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"errorController\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"errors\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeReserves\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fundingInterval\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fundingRateFactor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDelta\",\"inputs\":[{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_averagePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_lastIncreasedTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getEntryFundingRate\",\"inputs\":[{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFeeBasisPoints\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdgDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_feeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_taxBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_increment\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getFundingFee\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_entryFundingRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGlobalShortDelta\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMaxPrice\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinPrice\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextAveragePrice\",\"inputs\":[{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_averagePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_nextPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_lastIncreasedTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextFundingRate\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextGlobalShortAveragePrice\",\"inputs\":[{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nextPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPosition\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPositionDelta\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPositionFee\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPositionKey\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getPositionLeverage\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRedemptionAmount\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdgAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRedemptionCollateral\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRedemptionCollateralUsd\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTargetUsdgAmount\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUtilisation\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalShortAveragePrices\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"globalShortSizes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gov\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"guaranteedUsd\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasDynamicFees\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inManagerMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inPrivateLiquidationMode\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"includeAmmPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increasePosition\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_sizeDelta\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdg\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_priceFeed\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_liquidationFeeUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_fundingRateFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stableFundingRateFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isInitialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isLeverageEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isLiquidator\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isManager\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSwapEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lastFundingTimes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"liquidatePosition\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_feeReceiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"liquidationFeeUsd\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marginFeeBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxGasPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxGlobalShortSizes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxLeverage\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxUsdgAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minProfitBasisPoints\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minProfitTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mintBurnFeeBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"poolAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"positions\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"size\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateral\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"averagePrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"entryFundingRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reserveAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"realisedPnl\",\"type\":\"int256\",\"internalType\":\"int256\"},{\"name\":\"lastIncreasedTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"priceFeed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeRouter\",\"inputs\":[{\"name\":\"_router\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reservedAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sellUSDG\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBufferAmount\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setError\",\"inputs\":[{\"name\":\"_errorCode\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_error\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setErrorController\",\"inputs\":[{\"name\":\"_errorController\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFees\",\"inputs\":[{\"name\":\"_taxBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stableTaxBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_mintBurnFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_swapFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stableSwapFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_marginFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_liquidationFeeUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minProfitTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_hasDynamicFees\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFundingRate\",\"inputs\":[{\"name\":\"_fundingInterval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_fundingRateFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stableFundingRateFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInManagerMode\",\"inputs\":[{\"name\":\"_inManagerMode\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInPrivateLiquidationMode\",\"inputs\":[{\"name\":\"_inPrivateLiquidationMode\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIsLeverageEnabled\",\"inputs\":[{\"name\":\"_isLeverageEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIsSwapEnabled\",\"inputs\":[{\"name\":\"_isSwapEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLiquidator\",\"inputs\":[{\"name\":\"_liquidator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setManager\",\"inputs\":[{\"name\":\"_manager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isManager\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxGasPrice\",\"inputs\":[{\"name\":\"_maxGasPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxGlobalShortSize\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxLeverage\",\"inputs\":[{\"name\":\"_maxLeverage\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPriceFeed\",\"inputs\":[{\"name\":\"_priceFeed\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenConfig\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenDecimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokenWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minProfitBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_maxUsdgAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isStable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_isShortable\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUsdgAmount\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVaultUtils\",\"inputs\":[{\"name\":\"_vaultUtils\",\"type\":\"address\",\"internalType\":\"contractIVaultUtils\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shortableTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stableFundingRateFactor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stableSwapFeeBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stableTaxBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"stableTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"swap\",\"inputs\":[{\"name\":\"_tokenIn\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenOut\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"swapFeeBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"taxBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenBalances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenDecimals\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenToUsdMin\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenWeights\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalTokenWeights\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"updateCumulativeFundingRate\",\"inputs\":[{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeVault\",\"inputs\":[{\"name\":\"_newVault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"usdToToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdToTokenMax\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdToTokenMin\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_usdAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdg\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdgAmounts\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"useSwapPricing\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateLiquidation\",\"inputs\":[{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_indexToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLong\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_raise\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vaultUtils\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIVaultUtils\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelistedTokenCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"whitelistedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BuyUSDG\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"usdgAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feeBasisPoints\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClosePosition\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"size\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateral\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"averagePrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"entryFundingRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reserveAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"realisedPnl\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollectMarginFees\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feeUsd\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feeTokens\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollectSwapFees\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"feeUsd\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feeTokens\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecreaseGuaranteedUsd\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecreasePoolAmount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecreasePosition\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecreaseReservedAmount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DecreaseUsdgAmount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DirectPoolDeposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncreaseGuaranteedUsd\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncreasePoolAmount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncreasePosition\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"sizeDelta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncreaseReservedAmount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IncreaseUsdgAmount\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LiquidatePosition\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"collateralToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"indexToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"isLong\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"size\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateral\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reserveAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"realisedPnl\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"markPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SellUSDG\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"usdgAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"tokenAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feeBasisPoints\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Swap\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenIn\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenOut\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amountIn\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOut\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amountOutAfterFees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"feeBasisPoints\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdateFundingRate\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"fundingRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatePnl\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"hasProfit\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"delta\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdatePosition\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"size\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"collateral\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"averagePrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"entryFundingRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reserveAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"realisedPnl\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"markPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// VaultV2bABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultV2bMetaData.ABI instead.
var VaultV2bABI = VaultV2bMetaData.ABI

// VaultV2b is an auto generated Go binding around an Ethereum contract.
type VaultV2b struct {
	VaultV2bCaller     // Read-only binding to the contract
	VaultV2bTransactor // Write-only binding to the contract
	VaultV2bFilterer   // Log filterer for contract events
}

// VaultV2bCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultV2bCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultV2bTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultV2bTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultV2bFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultV2bFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultV2bSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultV2bSession struct {
	Contract     *VaultV2b         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultV2bCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultV2bCallerSession struct {
	Contract *VaultV2bCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// VaultV2bTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultV2bTransactorSession struct {
	Contract     *VaultV2bTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VaultV2bRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultV2bRaw struct {
	Contract *VaultV2b // Generic contract binding to access the raw methods on
}

// VaultV2bCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultV2bCallerRaw struct {
	Contract *VaultV2bCaller // Generic read-only contract binding to access the raw methods on
}

// VaultV2bTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultV2bTransactorRaw struct {
	Contract *VaultV2bTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultV2b creates a new instance of VaultV2b, bound to a specific deployed contract.
func NewVaultV2b(address common.Address, backend bind.ContractBackend) (*VaultV2b, error) {
	contract, err := bindVaultV2b(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultV2b{VaultV2bCaller: VaultV2bCaller{contract: contract}, VaultV2bTransactor: VaultV2bTransactor{contract: contract}, VaultV2bFilterer: VaultV2bFilterer{contract: contract}}, nil
}

// NewVaultV2bCaller creates a new read-only instance of VaultV2b, bound to a specific deployed contract.
func NewVaultV2bCaller(address common.Address, caller bind.ContractCaller) (*VaultV2bCaller, error) {
	contract, err := bindVaultV2b(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultV2bCaller{contract: contract}, nil
}

// NewVaultV2bTransactor creates a new write-only instance of VaultV2b, bound to a specific deployed contract.
func NewVaultV2bTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultV2bTransactor, error) {
	contract, err := bindVaultV2b(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultV2bTransactor{contract: contract}, nil
}

// NewVaultV2bFilterer creates a new log filterer instance of VaultV2b, bound to a specific deployed contract.
func NewVaultV2bFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultV2bFilterer, error) {
	contract, err := bindVaultV2b(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultV2bFilterer{contract: contract}, nil
}

// bindVaultV2b binds a generic wrapper to an already deployed contract.
func bindVaultV2b(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VaultV2bMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultV2b *VaultV2bRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultV2b.Contract.VaultV2bCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultV2b *VaultV2bRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultV2b.Contract.VaultV2bTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultV2b *VaultV2bRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultV2b.Contract.VaultV2bTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultV2b *VaultV2bCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VaultV2b.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultV2b *VaultV2bTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultV2b.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultV2b *VaultV2bTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultV2b.Contract.contract.Transact(opts, method, params...)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) BASISPOINTSDIVISOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "BASIS_POINTS_DIVISOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_VaultV2b *VaultV2bSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _VaultV2b.Contract.BASISPOINTSDIVISOR(&_VaultV2b.CallOpts)
}

// BASISPOINTSDIVISOR is a free data retrieval call binding the contract method 0x126082cf.
//
// Solidity: function BASIS_POINTS_DIVISOR() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) BASISPOINTSDIVISOR() (*big.Int, error) {
	return _VaultV2b.Contract.BASISPOINTSDIVISOR(&_VaultV2b.CallOpts)
}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) FUNDINGRATEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "FUNDING_RATE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_VaultV2b *VaultV2bSession) FUNDINGRATEPRECISION() (*big.Int, error) {
	return _VaultV2b.Contract.FUNDINGRATEPRECISION(&_VaultV2b.CallOpts)
}

// FUNDINGRATEPRECISION is a free data retrieval call binding the contract method 0x6be6026b.
//
// Solidity: function FUNDING_RATE_PRECISION() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) FUNDINGRATEPRECISION() (*big.Int, error) {
	return _VaultV2b.Contract.FUNDINGRATEPRECISION(&_VaultV2b.CallOpts)
}

// MAXFEEBASISPOINTS is a free data retrieval call binding the contract method 0x4befe2ca.
//
// Solidity: function MAX_FEE_BASIS_POINTS() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) MAXFEEBASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "MAX_FEE_BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEEBASISPOINTS is a free data retrieval call binding the contract method 0x4befe2ca.
//
// Solidity: function MAX_FEE_BASIS_POINTS() view returns(uint256)
func (_VaultV2b *VaultV2bSession) MAXFEEBASISPOINTS() (*big.Int, error) {
	return _VaultV2b.Contract.MAXFEEBASISPOINTS(&_VaultV2b.CallOpts)
}

// MAXFEEBASISPOINTS is a free data retrieval call binding the contract method 0x4befe2ca.
//
// Solidity: function MAX_FEE_BASIS_POINTS() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) MAXFEEBASISPOINTS() (*big.Int, error) {
	return _VaultV2b.Contract.MAXFEEBASISPOINTS(&_VaultV2b.CallOpts)
}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) MAXFUNDINGRATEFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "MAX_FUNDING_RATE_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_VaultV2b *VaultV2bSession) MAXFUNDINGRATEFACTOR() (*big.Int, error) {
	return _VaultV2b.Contract.MAXFUNDINGRATEFACTOR(&_VaultV2b.CallOpts)
}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) MAXFUNDINGRATEFACTOR() (*big.Int, error) {
	return _VaultV2b.Contract.MAXFUNDINGRATEFACTOR(&_VaultV2b.CallOpts)
}

// MAXLIQUIDATIONFEEUSD is a free data retrieval call binding the contract method 0x07c58752.
//
// Solidity: function MAX_LIQUIDATION_FEE_USD() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) MAXLIQUIDATIONFEEUSD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "MAX_LIQUIDATION_FEE_USD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLIQUIDATIONFEEUSD is a free data retrieval call binding the contract method 0x07c58752.
//
// Solidity: function MAX_LIQUIDATION_FEE_USD() view returns(uint256)
func (_VaultV2b *VaultV2bSession) MAXLIQUIDATIONFEEUSD() (*big.Int, error) {
	return _VaultV2b.Contract.MAXLIQUIDATIONFEEUSD(&_VaultV2b.CallOpts)
}

// MAXLIQUIDATIONFEEUSD is a free data retrieval call binding the contract method 0x07c58752.
//
// Solidity: function MAX_LIQUIDATION_FEE_USD() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) MAXLIQUIDATIONFEEUSD() (*big.Int, error) {
	return _VaultV2b.Contract.MAXLIQUIDATIONFEEUSD(&_VaultV2b.CallOpts)
}

// MINFUNDINGRATEINTERVAL is a free data retrieval call binding the contract method 0xfce28c10.
//
// Solidity: function MIN_FUNDING_RATE_INTERVAL() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) MINFUNDINGRATEINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "MIN_FUNDING_RATE_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINFUNDINGRATEINTERVAL is a free data retrieval call binding the contract method 0xfce28c10.
//
// Solidity: function MIN_FUNDING_RATE_INTERVAL() view returns(uint256)
func (_VaultV2b *VaultV2bSession) MINFUNDINGRATEINTERVAL() (*big.Int, error) {
	return _VaultV2b.Contract.MINFUNDINGRATEINTERVAL(&_VaultV2b.CallOpts)
}

// MINFUNDINGRATEINTERVAL is a free data retrieval call binding the contract method 0xfce28c10.
//
// Solidity: function MIN_FUNDING_RATE_INTERVAL() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) MINFUNDINGRATEINTERVAL() (*big.Int, error) {
	return _VaultV2b.Contract.MINFUNDINGRATEINTERVAL(&_VaultV2b.CallOpts)
}

// MINLEVERAGE is a free data retrieval call binding the contract method 0x34c1557d.
//
// Solidity: function MIN_LEVERAGE() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) MINLEVERAGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "MIN_LEVERAGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINLEVERAGE is a free data retrieval call binding the contract method 0x34c1557d.
//
// Solidity: function MIN_LEVERAGE() view returns(uint256)
func (_VaultV2b *VaultV2bSession) MINLEVERAGE() (*big.Int, error) {
	return _VaultV2b.Contract.MINLEVERAGE(&_VaultV2b.CallOpts)
}

// MINLEVERAGE is a free data retrieval call binding the contract method 0x34c1557d.
//
// Solidity: function MIN_LEVERAGE() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) MINLEVERAGE() (*big.Int, error) {
	return _VaultV2b.Contract.MINLEVERAGE(&_VaultV2b.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_VaultV2b *VaultV2bSession) PRICEPRECISION() (*big.Int, error) {
	return _VaultV2b.Contract.PRICEPRECISION(&_VaultV2b.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) PRICEPRECISION() (*big.Int, error) {
	return _VaultV2b.Contract.PRICEPRECISION(&_VaultV2b.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) USDGDECIMALS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "USDG_DECIMALS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_VaultV2b *VaultV2bSession) USDGDECIMALS() (*big.Int, error) {
	return _VaultV2b.Contract.USDGDECIMALS(&_VaultV2b.CallOpts)
}

// USDGDECIMALS is a free data retrieval call binding the contract method 0x870d917c.
//
// Solidity: function USDG_DECIMALS() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) USDGDECIMALS() (*big.Int, error) {
	return _VaultV2b.Contract.USDGDECIMALS(&_VaultV2b.CallOpts)
}

// AdjustForDecimals is a free data retrieval call binding the contract method 0x42152873.
//
// Solidity: function adjustForDecimals(uint256 _amount, address _tokenDiv, address _tokenMul) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) AdjustForDecimals(opts *bind.CallOpts, _amount *big.Int, _tokenDiv common.Address, _tokenMul common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "adjustForDecimals", _amount, _tokenDiv, _tokenMul)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdjustForDecimals is a free data retrieval call binding the contract method 0x42152873.
//
// Solidity: function adjustForDecimals(uint256 _amount, address _tokenDiv, address _tokenMul) view returns(uint256)
func (_VaultV2b *VaultV2bSession) AdjustForDecimals(_amount *big.Int, _tokenDiv common.Address, _tokenMul common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.AdjustForDecimals(&_VaultV2b.CallOpts, _amount, _tokenDiv, _tokenMul)
}

// AdjustForDecimals is a free data retrieval call binding the contract method 0x42152873.
//
// Solidity: function adjustForDecimals(uint256 _amount, address _tokenDiv, address _tokenMul) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) AdjustForDecimals(_amount *big.Int, _tokenDiv common.Address, _tokenMul common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.AdjustForDecimals(&_VaultV2b.CallOpts, _amount, _tokenDiv, _tokenMul)
}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_VaultV2b *VaultV2bCaller) AllWhitelistedTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "allWhitelistedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_VaultV2b *VaultV2bSession) AllWhitelistedTokens(arg0 *big.Int) (common.Address, error) {
	return _VaultV2b.Contract.AllWhitelistedTokens(&_VaultV2b.CallOpts, arg0)
}

// AllWhitelistedTokens is a free data retrieval call binding the contract method 0xe468baf0.
//
// Solidity: function allWhitelistedTokens(uint256 ) view returns(address)
func (_VaultV2b *VaultV2bCallerSession) AllWhitelistedTokens(arg0 *big.Int) (common.Address, error) {
	return _VaultV2b.Contract.AllWhitelistedTokens(&_VaultV2b.CallOpts, arg0)
}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) AllWhitelistedTokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "allWhitelistedTokensLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_VaultV2b *VaultV2bSession) AllWhitelistedTokensLength() (*big.Int, error) {
	return _VaultV2b.Contract.AllWhitelistedTokensLength(&_VaultV2b.CallOpts)
}

// AllWhitelistedTokensLength is a free data retrieval call binding the contract method 0x0842b076.
//
// Solidity: function allWhitelistedTokensLength() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) AllWhitelistedTokensLength() (*big.Int, error) {
	return _VaultV2b.Contract.AllWhitelistedTokensLength(&_VaultV2b.CallOpts)
}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address , address ) view returns(bool)
func (_VaultV2b *VaultV2bCaller) ApprovedRouters(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "approvedRouters", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address , address ) view returns(bool)
func (_VaultV2b *VaultV2bSession) ApprovedRouters(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _VaultV2b.Contract.ApprovedRouters(&_VaultV2b.CallOpts, arg0, arg1)
}

// ApprovedRouters is a free data retrieval call binding the contract method 0x60922199.
//
// Solidity: function approvedRouters(address , address ) view returns(bool)
func (_VaultV2b *VaultV2bCallerSession) ApprovedRouters(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _VaultV2b.Contract.ApprovedRouters(&_VaultV2b.CallOpts, arg0, arg1)
}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) BufferAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "bufferAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address ) view returns(uint256)
func (_VaultV2b *VaultV2bSession) BufferAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.BufferAmounts(&_VaultV2b.CallOpts, arg0)
}

// BufferAmounts is a free data retrieval call binding the contract method 0x4a993ee9.
//
// Solidity: function bufferAmounts(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) BufferAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.BufferAmounts(&_VaultV2b.CallOpts, arg0)
}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) CumulativeFundingRates(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "cumulativeFundingRates", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address ) view returns(uint256)
func (_VaultV2b *VaultV2bSession) CumulativeFundingRates(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.CumulativeFundingRates(&_VaultV2b.CallOpts, arg0)
}

// CumulativeFundingRates is a free data retrieval call binding the contract method 0xc65bc7b1.
//
// Solidity: function cumulativeFundingRates(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) CumulativeFundingRates(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.CumulativeFundingRates(&_VaultV2b.CallOpts, arg0)
}

// ErrorController is a free data retrieval call binding the contract method 0x48f35cbb.
//
// Solidity: function errorController() view returns(address)
func (_VaultV2b *VaultV2bCaller) ErrorController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "errorController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ErrorController is a free data retrieval call binding the contract method 0x48f35cbb.
//
// Solidity: function errorController() view returns(address)
func (_VaultV2b *VaultV2bSession) ErrorController() (common.Address, error) {
	return _VaultV2b.Contract.ErrorController(&_VaultV2b.CallOpts)
}

// ErrorController is a free data retrieval call binding the contract method 0x48f35cbb.
//
// Solidity: function errorController() view returns(address)
func (_VaultV2b *VaultV2bCallerSession) ErrorController() (common.Address, error) {
	return _VaultV2b.Contract.ErrorController(&_VaultV2b.CallOpts)
}

// Errors is a free data retrieval call binding the contract method 0xfed1a606.
//
// Solidity: function errors(uint256 ) view returns(string)
func (_VaultV2b *VaultV2bCaller) Errors(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "errors", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Errors is a free data retrieval call binding the contract method 0xfed1a606.
//
// Solidity: function errors(uint256 ) view returns(string)
func (_VaultV2b *VaultV2bSession) Errors(arg0 *big.Int) (string, error) {
	return _VaultV2b.Contract.Errors(&_VaultV2b.CallOpts, arg0)
}

// Errors is a free data retrieval call binding the contract method 0xfed1a606.
//
// Solidity: function errors(uint256 ) view returns(string)
func (_VaultV2b *VaultV2bCallerSession) Errors(arg0 *big.Int) (string, error) {
	return _VaultV2b.Contract.Errors(&_VaultV2b.CallOpts, arg0)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) FeeReserves(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "feeReserves", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_VaultV2b *VaultV2bSession) FeeReserves(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.FeeReserves(&_VaultV2b.CallOpts, arg0)
}

// FeeReserves is a free data retrieval call binding the contract method 0x1ce9cb8f.
//
// Solidity: function feeReserves(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) FeeReserves(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.FeeReserves(&_VaultV2b.CallOpts, arg0)
}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) FundingInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "fundingInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_VaultV2b *VaultV2bSession) FundingInterval() (*big.Int, error) {
	return _VaultV2b.Contract.FundingInterval(&_VaultV2b.CallOpts)
}

// FundingInterval is a free data retrieval call binding the contract method 0x9849e412.
//
// Solidity: function fundingInterval() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) FundingInterval() (*big.Int, error) {
	return _VaultV2b.Contract.FundingInterval(&_VaultV2b.CallOpts)
}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) FundingRateFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "fundingRateFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_VaultV2b *VaultV2bSession) FundingRateFactor() (*big.Int, error) {
	return _VaultV2b.Contract.FundingRateFactor(&_VaultV2b.CallOpts)
}

// FundingRateFactor is a free data retrieval call binding the contract method 0xc4f718bf.
//
// Solidity: function fundingRateFactor() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) FundingRateFactor() (*big.Int, error) {
	return _VaultV2b.Contract.FundingRateFactor(&_VaultV2b.CallOpts)
}

// GetDelta is a free data retrieval call binding the contract method 0x5c07eaab.
//
// Solidity: function getDelta(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _lastIncreasedTime) view returns(bool, uint256)
func (_VaultV2b *VaultV2bCaller) GetDelta(opts *bind.CallOpts, _indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getDelta", _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)

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
func (_VaultV2b *VaultV2bSession) GetDelta(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	return _VaultV2b.Contract.GetDelta(&_VaultV2b.CallOpts, _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)
}

// GetDelta is a free data retrieval call binding the contract method 0x5c07eaab.
//
// Solidity: function getDelta(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _lastIncreasedTime) view returns(bool, uint256)
func (_VaultV2b *VaultV2bCallerSession) GetDelta(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _lastIncreasedTime *big.Int) (bool, *big.Int, error) {
	return _VaultV2b.Contract.GetDelta(&_VaultV2b.CallOpts, _indexToken, _size, _averagePrice, _isLong, _lastIncreasedTime)
}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) GetEntryFundingRate(opts *bind.CallOpts, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getEntryFundingRate", _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_VaultV2b *VaultV2bSession) GetEntryFundingRate(_collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _VaultV2b.Contract.GetEntryFundingRate(&_VaultV2b.CallOpts, _collateralToken, _indexToken, _isLong)
}

// GetEntryFundingRate is a free data retrieval call binding the contract method 0xb1cc53ab.
//
// Solidity: function getEntryFundingRate(address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) GetEntryFundingRate(_collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _VaultV2b.Contract.GetEntryFundingRate(&_VaultV2b.CallOpts, _collateralToken, _indexToken, _isLong)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) GetFeeBasisPoints(opts *bind.CallOpts, _token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getFeeBasisPoints", _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_VaultV2b *VaultV2bSession) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _VaultV2b.Contract.GetFeeBasisPoints(&_VaultV2b.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetFeeBasisPoints is a free data retrieval call binding the contract method 0xc7e074c3.
//
// Solidity: function getFeeBasisPoints(address _token, uint256 _usdgDelta, uint256 _feeBasisPoints, uint256 _taxBasisPoints, bool _increment) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) GetFeeBasisPoints(_token common.Address, _usdgDelta *big.Int, _feeBasisPoints *big.Int, _taxBasisPoints *big.Int, _increment bool) (*big.Int, error) {
	return _VaultV2b.Contract.GetFeeBasisPoints(&_VaultV2b.CallOpts, _token, _usdgDelta, _feeBasisPoints, _taxBasisPoints, _increment)
}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) GetFundingFee(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getFundingFee", _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_VaultV2b *VaultV2bSession) GetFundingFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	return _VaultV2b.Contract.GetFundingFee(&_VaultV2b.CallOpts, _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)
}

// GetFundingFee is a free data retrieval call binding the contract method 0xda76524c.
//
// Solidity: function getFundingFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _size, uint256 _entryFundingRate) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) GetFundingFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _size *big.Int, _entryFundingRate *big.Int) (*big.Int, error) {
	return _VaultV2b.Contract.GetFundingFee(&_VaultV2b.CallOpts, _account, _collateralToken, _indexToken, _isLong, _size, _entryFundingRate)
}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xb364accb.
//
// Solidity: function getGlobalShortDelta(address _token) view returns(bool, uint256)
func (_VaultV2b *VaultV2bCaller) GetGlobalShortDelta(opts *bind.CallOpts, _token common.Address) (bool, *big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getGlobalShortDelta", _token)

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
func (_VaultV2b *VaultV2bSession) GetGlobalShortDelta(_token common.Address) (bool, *big.Int, error) {
	return _VaultV2b.Contract.GetGlobalShortDelta(&_VaultV2b.CallOpts, _token)
}

// GetGlobalShortDelta is a free data retrieval call binding the contract method 0xb364accb.
//
// Solidity: function getGlobalShortDelta(address _token) view returns(bool, uint256)
func (_VaultV2b *VaultV2bCallerSession) GetGlobalShortDelta(_token common.Address) (bool, *big.Int, error) {
	return _VaultV2b.Contract.GetGlobalShortDelta(&_VaultV2b.CallOpts, _token)
}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) GetMaxPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getMaxPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bSession) GetMaxPrice(_token common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GetMaxPrice(&_VaultV2b.CallOpts, _token)
}

// GetMaxPrice is a free data retrieval call binding the contract method 0xe124e6d2.
//
// Solidity: function getMaxPrice(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) GetMaxPrice(_token common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GetMaxPrice(&_VaultV2b.CallOpts, _token)
}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) GetMinPrice(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getMinPrice", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bSession) GetMinPrice(_token common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GetMinPrice(&_VaultV2b.CallOpts, _token)
}

// GetMinPrice is a free data retrieval call binding the contract method 0x81a612d6.
//
// Solidity: function getMinPrice(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) GetMinPrice(_token common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GetMinPrice(&_VaultV2b.CallOpts, _token)
}

// GetNextAveragePrice is a free data retrieval call binding the contract method 0xdb97495f.
//
// Solidity: function getNextAveragePrice(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _nextPrice, uint256 _sizeDelta, uint256 _lastIncreasedTime) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) GetNextAveragePrice(opts *bind.CallOpts, _indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _nextPrice *big.Int, _sizeDelta *big.Int, _lastIncreasedTime *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getNextAveragePrice", _indexToken, _size, _averagePrice, _isLong, _nextPrice, _sizeDelta, _lastIncreasedTime)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextAveragePrice is a free data retrieval call binding the contract method 0xdb97495f.
//
// Solidity: function getNextAveragePrice(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _nextPrice, uint256 _sizeDelta, uint256 _lastIncreasedTime) view returns(uint256)
func (_VaultV2b *VaultV2bSession) GetNextAveragePrice(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _nextPrice *big.Int, _sizeDelta *big.Int, _lastIncreasedTime *big.Int) (*big.Int, error) {
	return _VaultV2b.Contract.GetNextAveragePrice(&_VaultV2b.CallOpts, _indexToken, _size, _averagePrice, _isLong, _nextPrice, _sizeDelta, _lastIncreasedTime)
}

// GetNextAveragePrice is a free data retrieval call binding the contract method 0xdb97495f.
//
// Solidity: function getNextAveragePrice(address _indexToken, uint256 _size, uint256 _averagePrice, bool _isLong, uint256 _nextPrice, uint256 _sizeDelta, uint256 _lastIncreasedTime) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) GetNextAveragePrice(_indexToken common.Address, _size *big.Int, _averagePrice *big.Int, _isLong bool, _nextPrice *big.Int, _sizeDelta *big.Int, _lastIncreasedTime *big.Int) (*big.Int, error) {
	return _VaultV2b.Contract.GetNextAveragePrice(&_VaultV2b.CallOpts, _indexToken, _size, _averagePrice, _isLong, _nextPrice, _sizeDelta, _lastIncreasedTime)
}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) GetNextFundingRate(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getNextFundingRate", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bSession) GetNextFundingRate(_token common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GetNextFundingRate(&_VaultV2b.CallOpts, _token)
}

// GetNextFundingRate is a free data retrieval call binding the contract method 0xa93acac2.
//
// Solidity: function getNextFundingRate(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) GetNextFundingRate(_token common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GetNextFundingRate(&_VaultV2b.CallOpts, _token)
}

// GetNextGlobalShortAveragePrice is a free data retrieval call binding the contract method 0x9d7432ca.
//
// Solidity: function getNextGlobalShortAveragePrice(address _indexToken, uint256 _nextPrice, uint256 _sizeDelta) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) GetNextGlobalShortAveragePrice(opts *bind.CallOpts, _indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getNextGlobalShortAveragePrice", _indexToken, _nextPrice, _sizeDelta)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextGlobalShortAveragePrice is a free data retrieval call binding the contract method 0x9d7432ca.
//
// Solidity: function getNextGlobalShortAveragePrice(address _indexToken, uint256 _nextPrice, uint256 _sizeDelta) view returns(uint256)
func (_VaultV2b *VaultV2bSession) GetNextGlobalShortAveragePrice(_indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int) (*big.Int, error) {
	return _VaultV2b.Contract.GetNextGlobalShortAveragePrice(&_VaultV2b.CallOpts, _indexToken, _nextPrice, _sizeDelta)
}

// GetNextGlobalShortAveragePrice is a free data retrieval call binding the contract method 0x9d7432ca.
//
// Solidity: function getNextGlobalShortAveragePrice(address _indexToken, uint256 _nextPrice, uint256 _sizeDelta) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) GetNextGlobalShortAveragePrice(_indexToken common.Address, _nextPrice *big.Int, _sizeDelta *big.Int) (*big.Int, error) {
	return _VaultV2b.Contract.GetNextGlobalShortAveragePrice(&_VaultV2b.CallOpts, _indexToken, _nextPrice, _sizeDelta)
}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256, uint256, uint256, uint256, uint256, uint256, bool, uint256)
func (_VaultV2b *VaultV2bCaller) GetPosition(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getPosition", _account, _collateralToken, _indexToken, _isLong)

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
func (_VaultV2b *VaultV2bSession) GetPosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _VaultV2b.Contract.GetPosition(&_VaultV2b.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPosition is a free data retrieval call binding the contract method 0x4a3f088d.
//
// Solidity: function getPosition(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256, uint256, uint256, uint256, uint256, uint256, bool, uint256)
func (_VaultV2b *VaultV2bCallerSession) GetPosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, bool, *big.Int, error) {
	return _VaultV2b.Contract.GetPosition(&_VaultV2b.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionDelta is a free data retrieval call binding the contract method 0x45a6f370.
//
// Solidity: function getPositionDelta(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(bool, uint256)
func (_VaultV2b *VaultV2bCaller) GetPositionDelta(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (bool, *big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getPositionDelta", _account, _collateralToken, _indexToken, _isLong)

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
func (_VaultV2b *VaultV2bSession) GetPositionDelta(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (bool, *big.Int, error) {
	return _VaultV2b.Contract.GetPositionDelta(&_VaultV2b.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionDelta is a free data retrieval call binding the contract method 0x45a6f370.
//
// Solidity: function getPositionDelta(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(bool, uint256)
func (_VaultV2b *VaultV2bCallerSession) GetPositionDelta(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (bool, *big.Int, error) {
	return _VaultV2b.Contract.GetPositionDelta(&_VaultV2b.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) GetPositionFee(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getPositionFee", _account, _collateralToken, _indexToken, _isLong, _sizeDelta)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_VaultV2b *VaultV2bSession) GetPositionFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	return _VaultV2b.Contract.GetPositionFee(&_VaultV2b.CallOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta)
}

// GetPositionFee is a free data retrieval call binding the contract method 0xfdaf6ac3.
//
// Solidity: function getPositionFee(address _account, address _collateralToken, address _indexToken, bool _isLong, uint256 _sizeDelta) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) GetPositionFee(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _sizeDelta *big.Int) (*big.Int, error) {
	return _VaultV2b.Contract.GetPositionFee(&_VaultV2b.CallOpts, _account, _collateralToken, _indexToken, _isLong, _sizeDelta)
}

// GetPositionKey is a free data retrieval call binding the contract method 0x2d4b0576.
//
// Solidity: function getPositionKey(address _account, address _collateralToken, address _indexToken, bool _isLong) pure returns(bytes32)
func (_VaultV2b *VaultV2bCaller) GetPositionKey(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) ([32]byte, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getPositionKey", _account, _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPositionKey is a free data retrieval call binding the contract method 0x2d4b0576.
//
// Solidity: function getPositionKey(address _account, address _collateralToken, address _indexToken, bool _isLong) pure returns(bytes32)
func (_VaultV2b *VaultV2bSession) GetPositionKey(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) ([32]byte, error) {
	return _VaultV2b.Contract.GetPositionKey(&_VaultV2b.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionKey is a free data retrieval call binding the contract method 0x2d4b0576.
//
// Solidity: function getPositionKey(address _account, address _collateralToken, address _indexToken, bool _isLong) pure returns(bytes32)
func (_VaultV2b *VaultV2bCallerSession) GetPositionKey(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) ([32]byte, error) {
	return _VaultV2b.Contract.GetPositionKey(&_VaultV2b.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionLeverage is a free data retrieval call binding the contract method 0x51723e82.
//
// Solidity: function getPositionLeverage(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) GetPositionLeverage(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getPositionLeverage", _account, _collateralToken, _indexToken, _isLong)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPositionLeverage is a free data retrieval call binding the contract method 0x51723e82.
//
// Solidity: function getPositionLeverage(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_VaultV2b *VaultV2bSession) GetPositionLeverage(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _VaultV2b.Contract.GetPositionLeverage(&_VaultV2b.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetPositionLeverage is a free data retrieval call binding the contract method 0x51723e82.
//
// Solidity: function getPositionLeverage(address _account, address _collateralToken, address _indexToken, bool _isLong) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) GetPositionLeverage(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool) (*big.Int, error) {
	return _VaultV2b.Contract.GetPositionLeverage(&_VaultV2b.CallOpts, _account, _collateralToken, _indexToken, _isLong)
}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) GetRedemptionAmount(opts *bind.CallOpts, _token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getRedemptionAmount", _token, _usdgAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_VaultV2b *VaultV2bSession) GetRedemptionAmount(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _VaultV2b.Contract.GetRedemptionAmount(&_VaultV2b.CallOpts, _token, _usdgAmount)
}

// GetRedemptionAmount is a free data retrieval call binding the contract method 0x2c668ec1.
//
// Solidity: function getRedemptionAmount(address _token, uint256 _usdgAmount) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) GetRedemptionAmount(_token common.Address, _usdgAmount *big.Int) (*big.Int, error) {
	return _VaultV2b.Contract.GetRedemptionAmount(&_VaultV2b.CallOpts, _token, _usdgAmount)
}

// GetRedemptionCollateral is a free data retrieval call binding the contract method 0xb136ca49.
//
// Solidity: function getRedemptionCollateral(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) GetRedemptionCollateral(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getRedemptionCollateral", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionCollateral is a free data retrieval call binding the contract method 0xb136ca49.
//
// Solidity: function getRedemptionCollateral(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bSession) GetRedemptionCollateral(_token common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GetRedemptionCollateral(&_VaultV2b.CallOpts, _token)
}

// GetRedemptionCollateral is a free data retrieval call binding the contract method 0xb136ca49.
//
// Solidity: function getRedemptionCollateral(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) GetRedemptionCollateral(_token common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GetRedemptionCollateral(&_VaultV2b.CallOpts, _token)
}

// GetRedemptionCollateralUsd is a free data retrieval call binding the contract method 0x29ff9615.
//
// Solidity: function getRedemptionCollateralUsd(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) GetRedemptionCollateralUsd(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getRedemptionCollateralUsd", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRedemptionCollateralUsd is a free data retrieval call binding the contract method 0x29ff9615.
//
// Solidity: function getRedemptionCollateralUsd(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bSession) GetRedemptionCollateralUsd(_token common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GetRedemptionCollateralUsd(&_VaultV2b.CallOpts, _token)
}

// GetRedemptionCollateralUsd is a free data retrieval call binding the contract method 0x29ff9615.
//
// Solidity: function getRedemptionCollateralUsd(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) GetRedemptionCollateralUsd(_token common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GetRedemptionCollateralUsd(&_VaultV2b.CallOpts, _token)
}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) GetTargetUsdgAmount(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getTargetUsdgAmount", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bSession) GetTargetUsdgAmount(_token common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GetTargetUsdgAmount(&_VaultV2b.CallOpts, _token)
}

// GetTargetUsdgAmount is a free data retrieval call binding the contract method 0x3a05dcc1.
//
// Solidity: function getTargetUsdgAmount(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) GetTargetUsdgAmount(_token common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GetTargetUsdgAmount(&_VaultV2b.CallOpts, _token)
}

// GetUtilisation is a free data retrieval call binding the contract method 0x04fef1db.
//
// Solidity: function getUtilisation(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) GetUtilisation(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "getUtilisation", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUtilisation is a free data retrieval call binding the contract method 0x04fef1db.
//
// Solidity: function getUtilisation(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bSession) GetUtilisation(_token common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GetUtilisation(&_VaultV2b.CallOpts, _token)
}

// GetUtilisation is a free data retrieval call binding the contract method 0x04fef1db.
//
// Solidity: function getUtilisation(address _token) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) GetUtilisation(_token common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GetUtilisation(&_VaultV2b.CallOpts, _token)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) GlobalShortAveragePrices(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "globalShortAveragePrices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_VaultV2b *VaultV2bSession) GlobalShortAveragePrices(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GlobalShortAveragePrices(&_VaultV2b.CallOpts, arg0)
}

// GlobalShortAveragePrices is a free data retrieval call binding the contract method 0x62749803.
//
// Solidity: function globalShortAveragePrices(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) GlobalShortAveragePrices(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GlobalShortAveragePrices(&_VaultV2b.CallOpts, arg0)
}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) GlobalShortSizes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "globalShortSizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address ) view returns(uint256)
func (_VaultV2b *VaultV2bSession) GlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GlobalShortSizes(&_VaultV2b.CallOpts, arg0)
}

// GlobalShortSizes is a free data retrieval call binding the contract method 0x8a78daa8.
//
// Solidity: function globalShortSizes(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) GlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GlobalShortSizes(&_VaultV2b.CallOpts, arg0)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VaultV2b *VaultV2bCaller) Gov(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "gov")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VaultV2b *VaultV2bSession) Gov() (common.Address, error) {
	return _VaultV2b.Contract.Gov(&_VaultV2b.CallOpts)
}

// Gov is a free data retrieval call binding the contract method 0x12d43a51.
//
// Solidity: function gov() view returns(address)
func (_VaultV2b *VaultV2bCallerSession) Gov() (common.Address, error) {
	return _VaultV2b.Contract.Gov(&_VaultV2b.CallOpts)
}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) GuaranteedUsd(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "guaranteedUsd", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address ) view returns(uint256)
func (_VaultV2b *VaultV2bSession) GuaranteedUsd(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GuaranteedUsd(&_VaultV2b.CallOpts, arg0)
}

// GuaranteedUsd is a free data retrieval call binding the contract method 0xf07456ce.
//
// Solidity: function guaranteedUsd(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) GuaranteedUsd(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.GuaranteedUsd(&_VaultV2b.CallOpts, arg0)
}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_VaultV2b *VaultV2bCaller) HasDynamicFees(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "hasDynamicFees")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_VaultV2b *VaultV2bSession) HasDynamicFees() (bool, error) {
	return _VaultV2b.Contract.HasDynamicFees(&_VaultV2b.CallOpts)
}

// HasDynamicFees is a free data retrieval call binding the contract method 0x9f392eb3.
//
// Solidity: function hasDynamicFees() view returns(bool)
func (_VaultV2b *VaultV2bCallerSession) HasDynamicFees() (bool, error) {
	return _VaultV2b.Contract.HasDynamicFees(&_VaultV2b.CallOpts)
}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_VaultV2b *VaultV2bCaller) InManagerMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "inManagerMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_VaultV2b *VaultV2bSession) InManagerMode() (bool, error) {
	return _VaultV2b.Contract.InManagerMode(&_VaultV2b.CallOpts)
}

// InManagerMode is a free data retrieval call binding the contract method 0x9060b1ca.
//
// Solidity: function inManagerMode() view returns(bool)
func (_VaultV2b *VaultV2bCallerSession) InManagerMode() (bool, error) {
	return _VaultV2b.Contract.InManagerMode(&_VaultV2b.CallOpts)
}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_VaultV2b *VaultV2bCaller) InPrivateLiquidationMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "inPrivateLiquidationMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_VaultV2b *VaultV2bSession) InPrivateLiquidationMode() (bool, error) {
	return _VaultV2b.Contract.InPrivateLiquidationMode(&_VaultV2b.CallOpts)
}

// InPrivateLiquidationMode is a free data retrieval call binding the contract method 0x181e210e.
//
// Solidity: function inPrivateLiquidationMode() view returns(bool)
func (_VaultV2b *VaultV2bCallerSession) InPrivateLiquidationMode() (bool, error) {
	return _VaultV2b.Contract.InPrivateLiquidationMode(&_VaultV2b.CallOpts)
}

// IncludeAmmPrice is a free data retrieval call binding the contract method 0xab08c1c6.
//
// Solidity: function includeAmmPrice() view returns(bool)
func (_VaultV2b *VaultV2bCaller) IncludeAmmPrice(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "includeAmmPrice")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IncludeAmmPrice is a free data retrieval call binding the contract method 0xab08c1c6.
//
// Solidity: function includeAmmPrice() view returns(bool)
func (_VaultV2b *VaultV2bSession) IncludeAmmPrice() (bool, error) {
	return _VaultV2b.Contract.IncludeAmmPrice(&_VaultV2b.CallOpts)
}

// IncludeAmmPrice is a free data retrieval call binding the contract method 0xab08c1c6.
//
// Solidity: function includeAmmPrice() view returns(bool)
func (_VaultV2b *VaultV2bCallerSession) IncludeAmmPrice() (bool, error) {
	return _VaultV2b.Contract.IncludeAmmPrice(&_VaultV2b.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_VaultV2b *VaultV2bCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_VaultV2b *VaultV2bSession) IsInitialized() (bool, error) {
	return _VaultV2b.Contract.IsInitialized(&_VaultV2b.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_VaultV2b *VaultV2bCallerSession) IsInitialized() (bool, error) {
	return _VaultV2b.Contract.IsInitialized(&_VaultV2b.CallOpts)
}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_VaultV2b *VaultV2bCaller) IsLeverageEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "isLeverageEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_VaultV2b *VaultV2bSession) IsLeverageEnabled() (bool, error) {
	return _VaultV2b.Contract.IsLeverageEnabled(&_VaultV2b.CallOpts)
}

// IsLeverageEnabled is a free data retrieval call binding the contract method 0x3e72a262.
//
// Solidity: function isLeverageEnabled() view returns(bool)
func (_VaultV2b *VaultV2bCallerSession) IsLeverageEnabled() (bool, error) {
	return _VaultV2b.Contract.IsLeverageEnabled(&_VaultV2b.CallOpts)
}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address ) view returns(bool)
func (_VaultV2b *VaultV2bCaller) IsLiquidator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "isLiquidator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address ) view returns(bool)
func (_VaultV2b *VaultV2bSession) IsLiquidator(arg0 common.Address) (bool, error) {
	return _VaultV2b.Contract.IsLiquidator(&_VaultV2b.CallOpts, arg0)
}

// IsLiquidator is a free data retrieval call binding the contract method 0x529a356f.
//
// Solidity: function isLiquidator(address ) view returns(bool)
func (_VaultV2b *VaultV2bCallerSession) IsLiquidator(arg0 common.Address) (bool, error) {
	return _VaultV2b.Contract.IsLiquidator(&_VaultV2b.CallOpts, arg0)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_VaultV2b *VaultV2bCaller) IsManager(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "isManager", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_VaultV2b *VaultV2bSession) IsManager(arg0 common.Address) (bool, error) {
	return _VaultV2b.Contract.IsManager(&_VaultV2b.CallOpts, arg0)
}

// IsManager is a free data retrieval call binding the contract method 0xf3ae2415.
//
// Solidity: function isManager(address ) view returns(bool)
func (_VaultV2b *VaultV2bCallerSession) IsManager(arg0 common.Address) (bool, error) {
	return _VaultV2b.Contract.IsManager(&_VaultV2b.CallOpts, arg0)
}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_VaultV2b *VaultV2bCaller) IsSwapEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "isSwapEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_VaultV2b *VaultV2bSession) IsSwapEnabled() (bool, error) {
	return _VaultV2b.Contract.IsSwapEnabled(&_VaultV2b.CallOpts)
}

// IsSwapEnabled is a free data retrieval call binding the contract method 0x351a964d.
//
// Solidity: function isSwapEnabled() view returns(bool)
func (_VaultV2b *VaultV2bCallerSession) IsSwapEnabled() (bool, error) {
	return _VaultV2b.Contract.IsSwapEnabled(&_VaultV2b.CallOpts)
}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) LastFundingTimes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "lastFundingTimes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address ) view returns(uint256)
func (_VaultV2b *VaultV2bSession) LastFundingTimes(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.LastFundingTimes(&_VaultV2b.CallOpts, arg0)
}

// LastFundingTimes is a free data retrieval call binding the contract method 0xd8f897c3.
//
// Solidity: function lastFundingTimes(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) LastFundingTimes(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.LastFundingTimes(&_VaultV2b.CallOpts, arg0)
}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) LiquidationFeeUsd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "liquidationFeeUsd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_VaultV2b *VaultV2bSession) LiquidationFeeUsd() (*big.Int, error) {
	return _VaultV2b.Contract.LiquidationFeeUsd(&_VaultV2b.CallOpts)
}

// LiquidationFeeUsd is a free data retrieval call binding the contract method 0x174d2694.
//
// Solidity: function liquidationFeeUsd() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) LiquidationFeeUsd() (*big.Int, error) {
	return _VaultV2b.Contract.LiquidationFeeUsd(&_VaultV2b.CallOpts)
}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) MarginFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "marginFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_VaultV2b *VaultV2bSession) MarginFeeBasisPoints() (*big.Int, error) {
	return _VaultV2b.Contract.MarginFeeBasisPoints(&_VaultV2b.CallOpts)
}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) MarginFeeBasisPoints() (*big.Int, error) {
	return _VaultV2b.Contract.MarginFeeBasisPoints(&_VaultV2b.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) MaxGasPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "maxGasPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_VaultV2b *VaultV2bSession) MaxGasPrice() (*big.Int, error) {
	return _VaultV2b.Contract.MaxGasPrice(&_VaultV2b.CallOpts)
}

// MaxGasPrice is a free data retrieval call binding the contract method 0x3de39c11.
//
// Solidity: function maxGasPrice() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) MaxGasPrice() (*big.Int, error) {
	return _VaultV2b.Contract.MaxGasPrice(&_VaultV2b.CallOpts)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) MaxGlobalShortSizes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "maxGlobalShortSizes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_VaultV2b *VaultV2bSession) MaxGlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.MaxGlobalShortSizes(&_VaultV2b.CallOpts, arg0)
}

// MaxGlobalShortSizes is a free data retrieval call binding the contract method 0x9698d25a.
//
// Solidity: function maxGlobalShortSizes(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) MaxGlobalShortSizes(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.MaxGlobalShortSizes(&_VaultV2b.CallOpts, arg0)
}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) MaxLeverage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "maxLeverage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_VaultV2b *VaultV2bSession) MaxLeverage() (*big.Int, error) {
	return _VaultV2b.Contract.MaxLeverage(&_VaultV2b.CallOpts)
}

// MaxLeverage is a free data retrieval call binding the contract method 0xae3302c2.
//
// Solidity: function maxLeverage() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) MaxLeverage() (*big.Int, error) {
	return _VaultV2b.Contract.MaxLeverage(&_VaultV2b.CallOpts)
}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) MaxUsdgAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "maxUsdgAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address ) view returns(uint256)
func (_VaultV2b *VaultV2bSession) MaxUsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.MaxUsdgAmounts(&_VaultV2b.CallOpts, arg0)
}

// MaxUsdgAmounts is a free data retrieval call binding the contract method 0xad1e4f8d.
//
// Solidity: function maxUsdgAmounts(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) MaxUsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.MaxUsdgAmounts(&_VaultV2b.CallOpts, arg0)
}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) MinProfitBasisPoints(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "minProfitBasisPoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address ) view returns(uint256)
func (_VaultV2b *VaultV2bSession) MinProfitBasisPoints(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.MinProfitBasisPoints(&_VaultV2b.CallOpts, arg0)
}

// MinProfitBasisPoints is a free data retrieval call binding the contract method 0x88b1fbdf.
//
// Solidity: function minProfitBasisPoints(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) MinProfitBasisPoints(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.MinProfitBasisPoints(&_VaultV2b.CallOpts, arg0)
}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) MinProfitTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "minProfitTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_VaultV2b *VaultV2bSession) MinProfitTime() (*big.Int, error) {
	return _VaultV2b.Contract.MinProfitTime(&_VaultV2b.CallOpts)
}

// MinProfitTime is a free data retrieval call binding the contract method 0xd9ac4225.
//
// Solidity: function minProfitTime() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) MinProfitTime() (*big.Int, error) {
	return _VaultV2b.Contract.MinProfitTime(&_VaultV2b.CallOpts)
}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) MintBurnFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "mintBurnFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_VaultV2b *VaultV2bSession) MintBurnFeeBasisPoints() (*big.Int, error) {
	return _VaultV2b.Contract.MintBurnFeeBasisPoints(&_VaultV2b.CallOpts)
}

// MintBurnFeeBasisPoints is a free data retrieval call binding the contract method 0x4d47b304.
//
// Solidity: function mintBurnFeeBasisPoints() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) MintBurnFeeBasisPoints() (*big.Int, error) {
	return _VaultV2b.Contract.MintBurnFeeBasisPoints(&_VaultV2b.CallOpts)
}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) PoolAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "poolAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address ) view returns(uint256)
func (_VaultV2b *VaultV2bSession) PoolAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.PoolAmounts(&_VaultV2b.CallOpts, arg0)
}

// PoolAmounts is a free data retrieval call binding the contract method 0x52f55eed.
//
// Solidity: function poolAmounts(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) PoolAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.PoolAmounts(&_VaultV2b.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 lastIncreasedTime)
func (_VaultV2b *VaultV2bCaller) Positions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Size              *big.Int
	Collateral        *big.Int
	AveragePrice      *big.Int
	EntryFundingRate  *big.Int
	ReserveAmount     *big.Int
	RealisedPnl       *big.Int
	LastIncreasedTime *big.Int
}, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "positions", arg0)

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
func (_VaultV2b *VaultV2bSession) Positions(arg0 [32]byte) (struct {
	Size              *big.Int
	Collateral        *big.Int
	AveragePrice      *big.Int
	EntryFundingRate  *big.Int
	ReserveAmount     *big.Int
	RealisedPnl       *big.Int
	LastIncreasedTime *big.Int
}, error) {
	return _VaultV2b.Contract.Positions(&_VaultV2b.CallOpts, arg0)
}

// Positions is a free data retrieval call binding the contract method 0x514ea4bf.
//
// Solidity: function positions(bytes32 ) view returns(uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 lastIncreasedTime)
func (_VaultV2b *VaultV2bCallerSession) Positions(arg0 [32]byte) (struct {
	Size              *big.Int
	Collateral        *big.Int
	AveragePrice      *big.Int
	EntryFundingRate  *big.Int
	ReserveAmount     *big.Int
	RealisedPnl       *big.Int
	LastIncreasedTime *big.Int
}, error) {
	return _VaultV2b.Contract.Positions(&_VaultV2b.CallOpts, arg0)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_VaultV2b *VaultV2bCaller) PriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "priceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_VaultV2b *VaultV2bSession) PriceFeed() (common.Address, error) {
	return _VaultV2b.Contract.PriceFeed(&_VaultV2b.CallOpts)
}

// PriceFeed is a free data retrieval call binding the contract method 0x741bef1a.
//
// Solidity: function priceFeed() view returns(address)
func (_VaultV2b *VaultV2bCallerSession) PriceFeed() (common.Address, error) {
	return _VaultV2b.Contract.PriceFeed(&_VaultV2b.CallOpts)
}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) ReservedAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "reservedAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address ) view returns(uint256)
func (_VaultV2b *VaultV2bSession) ReservedAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.ReservedAmounts(&_VaultV2b.CallOpts, arg0)
}

// ReservedAmounts is a free data retrieval call binding the contract method 0xc3c7b9e9.
//
// Solidity: function reservedAmounts(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) ReservedAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.ReservedAmounts(&_VaultV2b.CallOpts, arg0)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_VaultV2b *VaultV2bCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_VaultV2b *VaultV2bSession) Router() (common.Address, error) {
	return _VaultV2b.Contract.Router(&_VaultV2b.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_VaultV2b *VaultV2bCallerSession) Router() (common.Address, error) {
	return _VaultV2b.Contract.Router(&_VaultV2b.CallOpts)
}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address ) view returns(bool)
func (_VaultV2b *VaultV2bCaller) ShortableTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "shortableTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address ) view returns(bool)
func (_VaultV2b *VaultV2bSession) ShortableTokens(arg0 common.Address) (bool, error) {
	return _VaultV2b.Contract.ShortableTokens(&_VaultV2b.CallOpts, arg0)
}

// ShortableTokens is a free data retrieval call binding the contract method 0xdb3555fb.
//
// Solidity: function shortableTokens(address ) view returns(bool)
func (_VaultV2b *VaultV2bCallerSession) ShortableTokens(arg0 common.Address) (bool, error) {
	return _VaultV2b.Contract.ShortableTokens(&_VaultV2b.CallOpts, arg0)
}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) StableFundingRateFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "stableFundingRateFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_VaultV2b *VaultV2bSession) StableFundingRateFactor() (*big.Int, error) {
	return _VaultV2b.Contract.StableFundingRateFactor(&_VaultV2b.CallOpts)
}

// StableFundingRateFactor is a free data retrieval call binding the contract method 0x134ca63b.
//
// Solidity: function stableFundingRateFactor() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) StableFundingRateFactor() (*big.Int, error) {
	return _VaultV2b.Contract.StableFundingRateFactor(&_VaultV2b.CallOpts)
}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) StableSwapFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "stableSwapFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_VaultV2b *VaultV2bSession) StableSwapFeeBasisPoints() (*big.Int, error) {
	return _VaultV2b.Contract.StableSwapFeeBasisPoints(&_VaultV2b.CallOpts)
}

// StableSwapFeeBasisPoints is a free data retrieval call binding the contract method 0xdf73a267.
//
// Solidity: function stableSwapFeeBasisPoints() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) StableSwapFeeBasisPoints() (*big.Int, error) {
	return _VaultV2b.Contract.StableSwapFeeBasisPoints(&_VaultV2b.CallOpts)
}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) StableTaxBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "stableTaxBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_VaultV2b *VaultV2bSession) StableTaxBasisPoints() (*big.Int, error) {
	return _VaultV2b.Contract.StableTaxBasisPoints(&_VaultV2b.CallOpts)
}

// StableTaxBasisPoints is a free data retrieval call binding the contract method 0x10eb56c2.
//
// Solidity: function stableTaxBasisPoints() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) StableTaxBasisPoints() (*big.Int, error) {
	return _VaultV2b.Contract.StableTaxBasisPoints(&_VaultV2b.CallOpts)
}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address ) view returns(bool)
func (_VaultV2b *VaultV2bCaller) StableTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "stableTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address ) view returns(bool)
func (_VaultV2b *VaultV2bSession) StableTokens(arg0 common.Address) (bool, error) {
	return _VaultV2b.Contract.StableTokens(&_VaultV2b.CallOpts, arg0)
}

// StableTokens is a free data retrieval call binding the contract method 0x42b60b03.
//
// Solidity: function stableTokens(address ) view returns(bool)
func (_VaultV2b *VaultV2bCallerSession) StableTokens(arg0 common.Address) (bool, error) {
	return _VaultV2b.Contract.StableTokens(&_VaultV2b.CallOpts, arg0)
}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) SwapFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "swapFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_VaultV2b *VaultV2bSession) SwapFeeBasisPoints() (*big.Int, error) {
	return _VaultV2b.Contract.SwapFeeBasisPoints(&_VaultV2b.CallOpts)
}

// SwapFeeBasisPoints is a free data retrieval call binding the contract method 0xa22f2392.
//
// Solidity: function swapFeeBasisPoints() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) SwapFeeBasisPoints() (*big.Int, error) {
	return _VaultV2b.Contract.SwapFeeBasisPoints(&_VaultV2b.CallOpts)
}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) TaxBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "taxBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_VaultV2b *VaultV2bSession) TaxBasisPoints() (*big.Int, error) {
	return _VaultV2b.Contract.TaxBasisPoints(&_VaultV2b.CallOpts)
}

// TaxBasisPoints is a free data retrieval call binding the contract method 0x7a210a2b.
//
// Solidity: function taxBasisPoints() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) TaxBasisPoints() (*big.Int, error) {
	return _VaultV2b.Contract.TaxBasisPoints(&_VaultV2b.CallOpts)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) TokenBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "tokenBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_VaultV2b *VaultV2bSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.TokenBalances(&_VaultV2b.CallOpts, arg0)
}

// TokenBalances is a free data retrieval call binding the contract method 0x523fba7f.
//
// Solidity: function tokenBalances(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) TokenBalances(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.TokenBalances(&_VaultV2b.CallOpts, arg0)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) TokenDecimals(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "tokenDecimals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint256)
func (_VaultV2b *VaultV2bSession) TokenDecimals(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.TokenDecimals(&_VaultV2b.CallOpts, arg0)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) TokenDecimals(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.TokenDecimals(&_VaultV2b.CallOpts, arg0)
}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) TokenToUsdMin(opts *bind.CallOpts, _token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "tokenToUsdMin", _token, _tokenAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_VaultV2b *VaultV2bSession) TokenToUsdMin(_token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	return _VaultV2b.Contract.TokenToUsdMin(&_VaultV2b.CallOpts, _token, _tokenAmount)
}

// TokenToUsdMin is a free data retrieval call binding the contract method 0x0a48d5a9.
//
// Solidity: function tokenToUsdMin(address _token, uint256 _tokenAmount) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) TokenToUsdMin(_token common.Address, _tokenAmount *big.Int) (*big.Int, error) {
	return _VaultV2b.Contract.TokenToUsdMin(&_VaultV2b.CallOpts, _token, _tokenAmount)
}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) TokenWeights(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "tokenWeights", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address ) view returns(uint256)
func (_VaultV2b *VaultV2bSession) TokenWeights(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.TokenWeights(&_VaultV2b.CallOpts, arg0)
}

// TokenWeights is a free data retrieval call binding the contract method 0xab2f3ad4.
//
// Solidity: function tokenWeights(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) TokenWeights(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.TokenWeights(&_VaultV2b.CallOpts, arg0)
}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) TotalTokenWeights(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "totalTokenWeights")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_VaultV2b *VaultV2bSession) TotalTokenWeights() (*big.Int, error) {
	return _VaultV2b.Contract.TotalTokenWeights(&_VaultV2b.CallOpts)
}

// TotalTokenWeights is a free data retrieval call binding the contract method 0xdc8f5fac.
//
// Solidity: function totalTokenWeights() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) TotalTokenWeights() (*big.Int, error) {
	return _VaultV2b.Contract.TotalTokenWeights(&_VaultV2b.CallOpts)
}

// UsdToToken is a free data retrieval call binding the contract method 0xfa12dbc0.
//
// Solidity: function usdToToken(address _token, uint256 _usdAmount, uint256 _price) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) UsdToToken(opts *bind.CallOpts, _token common.Address, _usdAmount *big.Int, _price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "usdToToken", _token, _usdAmount, _price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdToToken is a free data retrieval call binding the contract method 0xfa12dbc0.
//
// Solidity: function usdToToken(address _token, uint256 _usdAmount, uint256 _price) view returns(uint256)
func (_VaultV2b *VaultV2bSession) UsdToToken(_token common.Address, _usdAmount *big.Int, _price *big.Int) (*big.Int, error) {
	return _VaultV2b.Contract.UsdToToken(&_VaultV2b.CallOpts, _token, _usdAmount, _price)
}

// UsdToToken is a free data retrieval call binding the contract method 0xfa12dbc0.
//
// Solidity: function usdToToken(address _token, uint256 _usdAmount, uint256 _price) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) UsdToToken(_token common.Address, _usdAmount *big.Int, _price *big.Int) (*big.Int, error) {
	return _VaultV2b.Contract.UsdToToken(&_VaultV2b.CallOpts, _token, _usdAmount, _price)
}

// UsdToTokenMax is a free data retrieval call binding the contract method 0xa42ab3d2.
//
// Solidity: function usdToTokenMax(address _token, uint256 _usdAmount) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) UsdToTokenMax(opts *bind.CallOpts, _token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "usdToTokenMax", _token, _usdAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdToTokenMax is a free data retrieval call binding the contract method 0xa42ab3d2.
//
// Solidity: function usdToTokenMax(address _token, uint256 _usdAmount) view returns(uint256)
func (_VaultV2b *VaultV2bSession) UsdToTokenMax(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _VaultV2b.Contract.UsdToTokenMax(&_VaultV2b.CallOpts, _token, _usdAmount)
}

// UsdToTokenMax is a free data retrieval call binding the contract method 0xa42ab3d2.
//
// Solidity: function usdToTokenMax(address _token, uint256 _usdAmount) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) UsdToTokenMax(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _VaultV2b.Contract.UsdToTokenMax(&_VaultV2b.CallOpts, _token, _usdAmount)
}

// UsdToTokenMin is a free data retrieval call binding the contract method 0x9899cd02.
//
// Solidity: function usdToTokenMin(address _token, uint256 _usdAmount) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) UsdToTokenMin(opts *bind.CallOpts, _token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "usdToTokenMin", _token, _usdAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdToTokenMin is a free data retrieval call binding the contract method 0x9899cd02.
//
// Solidity: function usdToTokenMin(address _token, uint256 _usdAmount) view returns(uint256)
func (_VaultV2b *VaultV2bSession) UsdToTokenMin(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _VaultV2b.Contract.UsdToTokenMin(&_VaultV2b.CallOpts, _token, _usdAmount)
}

// UsdToTokenMin is a free data retrieval call binding the contract method 0x9899cd02.
//
// Solidity: function usdToTokenMin(address _token, uint256 _usdAmount) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) UsdToTokenMin(_token common.Address, _usdAmount *big.Int) (*big.Int, error) {
	return _VaultV2b.Contract.UsdToTokenMin(&_VaultV2b.CallOpts, _token, _usdAmount)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_VaultV2b *VaultV2bCaller) Usdg(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "usdg")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_VaultV2b *VaultV2bSession) Usdg() (common.Address, error) {
	return _VaultV2b.Contract.Usdg(&_VaultV2b.CallOpts)
}

// Usdg is a free data retrieval call binding the contract method 0xf5b91b7b.
//
// Solidity: function usdg() view returns(address)
func (_VaultV2b *VaultV2bCallerSession) Usdg() (common.Address, error) {
	return _VaultV2b.Contract.Usdg(&_VaultV2b.CallOpts)
}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCaller) UsdgAmounts(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "usdgAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address ) view returns(uint256)
func (_VaultV2b *VaultV2bSession) UsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.UsdgAmounts(&_VaultV2b.CallOpts, arg0)
}

// UsdgAmounts is a free data retrieval call binding the contract method 0x1aa4ace5.
//
// Solidity: function usdgAmounts(address ) view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) UsdgAmounts(arg0 common.Address) (*big.Int, error) {
	return _VaultV2b.Contract.UsdgAmounts(&_VaultV2b.CallOpts, arg0)
}

// UseSwapPricing is a free data retrieval call binding the contract method 0xb06423f3.
//
// Solidity: function useSwapPricing() view returns(bool)
func (_VaultV2b *VaultV2bCaller) UseSwapPricing(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "useSwapPricing")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UseSwapPricing is a free data retrieval call binding the contract method 0xb06423f3.
//
// Solidity: function useSwapPricing() view returns(bool)
func (_VaultV2b *VaultV2bSession) UseSwapPricing() (bool, error) {
	return _VaultV2b.Contract.UseSwapPricing(&_VaultV2b.CallOpts)
}

// UseSwapPricing is a free data retrieval call binding the contract method 0xb06423f3.
//
// Solidity: function useSwapPricing() view returns(bool)
func (_VaultV2b *VaultV2bCallerSession) UseSwapPricing() (bool, error) {
	return _VaultV2b.Contract.UseSwapPricing(&_VaultV2b.CallOpts)
}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_VaultV2b *VaultV2bCaller) ValidateLiquidation(opts *bind.CallOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "validateLiquidation", _account, _collateralToken, _indexToken, _isLong, _raise)

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
func (_VaultV2b *VaultV2bSession) ValidateLiquidation(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	return _VaultV2b.Contract.ValidateLiquidation(&_VaultV2b.CallOpts, _account, _collateralToken, _indexToken, _isLong, _raise)
}

// ValidateLiquidation is a free data retrieval call binding the contract method 0xd54d5a9f.
//
// Solidity: function validateLiquidation(address _account, address _collateralToken, address _indexToken, bool _isLong, bool _raise) view returns(uint256, uint256)
func (_VaultV2b *VaultV2bCallerSession) ValidateLiquidation(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _raise bool) (*big.Int, *big.Int, error) {
	return _VaultV2b.Contract.ValidateLiquidation(&_VaultV2b.CallOpts, _account, _collateralToken, _indexToken, _isLong, _raise)
}

// VaultUtils is a free data retrieval call binding the contract method 0x6abbe0c8.
//
// Solidity: function vaultUtils() view returns(address)
func (_VaultV2b *VaultV2bCaller) VaultUtils(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "vaultUtils")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultUtils is a free data retrieval call binding the contract method 0x6abbe0c8.
//
// Solidity: function vaultUtils() view returns(address)
func (_VaultV2b *VaultV2bSession) VaultUtils() (common.Address, error) {
	return _VaultV2b.Contract.VaultUtils(&_VaultV2b.CallOpts)
}

// VaultUtils is a free data retrieval call binding the contract method 0x6abbe0c8.
//
// Solidity: function vaultUtils() view returns(address)
func (_VaultV2b *VaultV2bCallerSession) VaultUtils() (common.Address, error) {
	return _VaultV2b.Contract.VaultUtils(&_VaultV2b.CallOpts)
}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_VaultV2b *VaultV2bCaller) WhitelistedTokenCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "whitelistedTokenCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_VaultV2b *VaultV2bSession) WhitelistedTokenCount() (*big.Int, error) {
	return _VaultV2b.Contract.WhitelistedTokenCount(&_VaultV2b.CallOpts)
}

// WhitelistedTokenCount is a free data retrieval call binding the contract method 0x62287a32.
//
// Solidity: function whitelistedTokenCount() view returns(uint256)
func (_VaultV2b *VaultV2bCallerSession) WhitelistedTokenCount() (*big.Int, error) {
	return _VaultV2b.Contract.WhitelistedTokenCount(&_VaultV2b.CallOpts)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_VaultV2b *VaultV2bCaller) WhitelistedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _VaultV2b.contract.Call(opts, &out, "whitelistedTokens", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_VaultV2b *VaultV2bSession) WhitelistedTokens(arg0 common.Address) (bool, error) {
	return _VaultV2b.Contract.WhitelistedTokens(&_VaultV2b.CallOpts, arg0)
}

// WhitelistedTokens is a free data retrieval call binding the contract method 0xdaf9c210.
//
// Solidity: function whitelistedTokens(address ) view returns(bool)
func (_VaultV2b *VaultV2bCallerSession) WhitelistedTokens(arg0 common.Address) (bool, error) {
	return _VaultV2b.Contract.WhitelistedTokens(&_VaultV2b.CallOpts, arg0)
}

// AddRouter is a paid mutator transaction binding the contract method 0x24ca984e.
//
// Solidity: function addRouter(address _router) returns()
func (_VaultV2b *VaultV2bTransactor) AddRouter(opts *bind.TransactOpts, _router common.Address) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "addRouter", _router)
}

// AddRouter is a paid mutator transaction binding the contract method 0x24ca984e.
//
// Solidity: function addRouter(address _router) returns()
func (_VaultV2b *VaultV2bSession) AddRouter(_router common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.AddRouter(&_VaultV2b.TransactOpts, _router)
}

// AddRouter is a paid mutator transaction binding the contract method 0x24ca984e.
//
// Solidity: function addRouter(address _router) returns()
func (_VaultV2b *VaultV2bTransactorSession) AddRouter(_router common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.AddRouter(&_VaultV2b.TransactOpts, _router)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_VaultV2b *VaultV2bTransactor) BuyUSDG(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "buyUSDG", _token, _receiver)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_VaultV2b *VaultV2bSession) BuyUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.BuyUSDG(&_VaultV2b.TransactOpts, _token, _receiver)
}

// BuyUSDG is a paid mutator transaction binding the contract method 0x817bb857.
//
// Solidity: function buyUSDG(address _token, address _receiver) returns(uint256)
func (_VaultV2b *VaultV2bTransactorSession) BuyUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.BuyUSDG(&_VaultV2b.TransactOpts, _token, _receiver)
}

// ClearTokenConfig is a paid mutator transaction binding the contract method 0xe67f59a7.
//
// Solidity: function clearTokenConfig(address _token) returns()
func (_VaultV2b *VaultV2bTransactor) ClearTokenConfig(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "clearTokenConfig", _token)
}

// ClearTokenConfig is a paid mutator transaction binding the contract method 0xe67f59a7.
//
// Solidity: function clearTokenConfig(address _token) returns()
func (_VaultV2b *VaultV2bSession) ClearTokenConfig(_token common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.ClearTokenConfig(&_VaultV2b.TransactOpts, _token)
}

// ClearTokenConfig is a paid mutator transaction binding the contract method 0xe67f59a7.
//
// Solidity: function clearTokenConfig(address _token) returns()
func (_VaultV2b *VaultV2bTransactorSession) ClearTokenConfig(_token common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.ClearTokenConfig(&_VaultV2b.TransactOpts, _token)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_VaultV2b *VaultV2bTransactor) DecreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "decreasePosition", _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_VaultV2b *VaultV2bSession) DecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.DecreasePosition(&_VaultV2b.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DecreasePosition is a paid mutator transaction binding the contract method 0x82a08490.
//
// Solidity: function decreasePosition(address _account, address _collateralToken, address _indexToken, uint256 _collateralDelta, uint256 _sizeDelta, bool _isLong, address _receiver) returns(uint256)
func (_VaultV2b *VaultV2bTransactorSession) DecreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _collateralDelta *big.Int, _sizeDelta *big.Int, _isLong bool, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.DecreasePosition(&_VaultV2b.TransactOpts, _account, _collateralToken, _indexToken, _collateralDelta, _sizeDelta, _isLong, _receiver)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_VaultV2b *VaultV2bTransactor) DirectPoolDeposit(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "directPoolDeposit", _token)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_VaultV2b *VaultV2bSession) DirectPoolDeposit(_token common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.DirectPoolDeposit(&_VaultV2b.TransactOpts, _token)
}

// DirectPoolDeposit is a paid mutator transaction binding the contract method 0x5f7bc119.
//
// Solidity: function directPoolDeposit(address _token) returns()
func (_VaultV2b *VaultV2bTransactorSession) DirectPoolDeposit(_token common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.DirectPoolDeposit(&_VaultV2b.TransactOpts, _token)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_VaultV2b *VaultV2bTransactor) IncreasePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "increasePosition", _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_VaultV2b *VaultV2bSession) IncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _VaultV2b.Contract.IncreasePosition(&_VaultV2b.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// IncreasePosition is a paid mutator transaction binding the contract method 0x48d91abf.
//
// Solidity: function increasePosition(address _account, address _collateralToken, address _indexToken, uint256 _sizeDelta, bool _isLong) returns()
func (_VaultV2b *VaultV2bTransactorSession) IncreasePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _sizeDelta *big.Int, _isLong bool) (*types.Transaction, error) {
	return _VaultV2b.Contract.IncreasePosition(&_VaultV2b.TransactOpts, _account, _collateralToken, _indexToken, _sizeDelta, _isLong)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _router, address _usdg, address _priceFeed, uint256 _liquidationFeeUsd, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_VaultV2b *VaultV2bTransactor) Initialize(opts *bind.TransactOpts, _router common.Address, _usdg common.Address, _priceFeed common.Address, _liquidationFeeUsd *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "initialize", _router, _usdg, _priceFeed, _liquidationFeeUsd, _fundingRateFactor, _stableFundingRateFactor)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _router, address _usdg, address _priceFeed, uint256 _liquidationFeeUsd, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_VaultV2b *VaultV2bSession) Initialize(_router common.Address, _usdg common.Address, _priceFeed common.Address, _liquidationFeeUsd *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _VaultV2b.Contract.Initialize(&_VaultV2b.TransactOpts, _router, _usdg, _priceFeed, _liquidationFeeUsd, _fundingRateFactor, _stableFundingRateFactor)
}

// Initialize is a paid mutator transaction binding the contract method 0x728cdbca.
//
// Solidity: function initialize(address _router, address _usdg, address _priceFeed, uint256 _liquidationFeeUsd, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_VaultV2b *VaultV2bTransactorSession) Initialize(_router common.Address, _usdg common.Address, _priceFeed common.Address, _liquidationFeeUsd *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _VaultV2b.Contract.Initialize(&_VaultV2b.TransactOpts, _router, _usdg, _priceFeed, _liquidationFeeUsd, _fundingRateFactor, _stableFundingRateFactor)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_VaultV2b *VaultV2bTransactor) LiquidatePosition(opts *bind.TransactOpts, _account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "liquidatePosition", _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_VaultV2b *VaultV2bSession) LiquidatePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.LiquidatePosition(&_VaultV2b.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// LiquidatePosition is a paid mutator transaction binding the contract method 0xde2ea948.
//
// Solidity: function liquidatePosition(address _account, address _collateralToken, address _indexToken, bool _isLong, address _feeReceiver) returns()
func (_VaultV2b *VaultV2bTransactorSession) LiquidatePosition(_account common.Address, _collateralToken common.Address, _indexToken common.Address, _isLong bool, _feeReceiver common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.LiquidatePosition(&_VaultV2b.TransactOpts, _account, _collateralToken, _indexToken, _isLong, _feeReceiver)
}

// RemoveRouter is a paid mutator transaction binding the contract method 0x6ae0b154.
//
// Solidity: function removeRouter(address _router) returns()
func (_VaultV2b *VaultV2bTransactor) RemoveRouter(opts *bind.TransactOpts, _router common.Address) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "removeRouter", _router)
}

// RemoveRouter is a paid mutator transaction binding the contract method 0x6ae0b154.
//
// Solidity: function removeRouter(address _router) returns()
func (_VaultV2b *VaultV2bSession) RemoveRouter(_router common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.RemoveRouter(&_VaultV2b.TransactOpts, _router)
}

// RemoveRouter is a paid mutator transaction binding the contract method 0x6ae0b154.
//
// Solidity: function removeRouter(address _router) returns()
func (_VaultV2b *VaultV2bTransactorSession) RemoveRouter(_router common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.RemoveRouter(&_VaultV2b.TransactOpts, _router)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_VaultV2b *VaultV2bTransactor) SellUSDG(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "sellUSDG", _token, _receiver)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_VaultV2b *VaultV2bSession) SellUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.SellUSDG(&_VaultV2b.TransactOpts, _token, _receiver)
}

// SellUSDG is a paid mutator transaction binding the contract method 0x711e6190.
//
// Solidity: function sellUSDG(address _token, address _receiver) returns(uint256)
func (_VaultV2b *VaultV2bTransactorSession) SellUSDG(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.SellUSDG(&_VaultV2b.TransactOpts, _token, _receiver)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_VaultV2b *VaultV2bTransactor) SetBufferAmount(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "setBufferAmount", _token, _amount)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_VaultV2b *VaultV2bSession) SetBufferAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetBufferAmount(&_VaultV2b.TransactOpts, _token, _amount)
}

// SetBufferAmount is a paid mutator transaction binding the contract method 0x8585f4d2.
//
// Solidity: function setBufferAmount(address _token, uint256 _amount) returns()
func (_VaultV2b *VaultV2bTransactorSession) SetBufferAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetBufferAmount(&_VaultV2b.TransactOpts, _token, _amount)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_VaultV2b *VaultV2bTransactor) SetError(opts *bind.TransactOpts, _errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "setError", _errorCode, _error)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_VaultV2b *VaultV2bSession) SetError(_errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetError(&_VaultV2b.TransactOpts, _errorCode, _error)
}

// SetError is a paid mutator transaction binding the contract method 0x28e67be5.
//
// Solidity: function setError(uint256 _errorCode, string _error) returns()
func (_VaultV2b *VaultV2bTransactorSession) SetError(_errorCode *big.Int, _error string) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetError(&_VaultV2b.TransactOpts, _errorCode, _error)
}

// SetErrorController is a paid mutator transaction binding the contract method 0x8f7b8404.
//
// Solidity: function setErrorController(address _errorController) returns()
func (_VaultV2b *VaultV2bTransactor) SetErrorController(opts *bind.TransactOpts, _errorController common.Address) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "setErrorController", _errorController)
}

// SetErrorController is a paid mutator transaction binding the contract method 0x8f7b8404.
//
// Solidity: function setErrorController(address _errorController) returns()
func (_VaultV2b *VaultV2bSession) SetErrorController(_errorController common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetErrorController(&_VaultV2b.TransactOpts, _errorController)
}

// SetErrorController is a paid mutator transaction binding the contract method 0x8f7b8404.
//
// Solidity: function setErrorController(address _errorController) returns()
func (_VaultV2b *VaultV2bTransactorSession) SetErrorController(_errorController common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetErrorController(&_VaultV2b.TransactOpts, _errorController)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_VaultV2b *VaultV2bTransactor) SetFees(opts *bind.TransactOpts, _taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "setFees", _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_VaultV2b *VaultV2bSession) SetFees(_taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetFees(&_VaultV2b.TransactOpts, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFees is a paid mutator transaction binding the contract method 0x40eb3802.
//
// Solidity: function setFees(uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_VaultV2b *VaultV2bTransactorSession) SetFees(_taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetFees(&_VaultV2b.TransactOpts, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_VaultV2b *VaultV2bTransactor) SetFundingRate(opts *bind.TransactOpts, _fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "setFundingRate", _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_VaultV2b *VaultV2bSession) SetFundingRate(_fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetFundingRate(&_VaultV2b.TransactOpts, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x8a27d468.
//
// Solidity: function setFundingRate(uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_VaultV2b *VaultV2bTransactorSession) SetFundingRate(_fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetFundingRate(&_VaultV2b.TransactOpts, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VaultV2b *VaultV2bTransactor) SetGov(opts *bind.TransactOpts, _gov common.Address) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "setGov", _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VaultV2b *VaultV2bSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetGov(&_VaultV2b.TransactOpts, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0xcfad57a2.
//
// Solidity: function setGov(address _gov) returns()
func (_VaultV2b *VaultV2bTransactorSession) SetGov(_gov common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetGov(&_VaultV2b.TransactOpts, _gov)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_VaultV2b *VaultV2bTransactor) SetInManagerMode(opts *bind.TransactOpts, _inManagerMode bool) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "setInManagerMode", _inManagerMode)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_VaultV2b *VaultV2bSession) SetInManagerMode(_inManagerMode bool) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetInManagerMode(&_VaultV2b.TransactOpts, _inManagerMode)
}

// SetInManagerMode is a paid mutator transaction binding the contract method 0x24b0c04d.
//
// Solidity: function setInManagerMode(bool _inManagerMode) returns()
func (_VaultV2b *VaultV2bTransactorSession) SetInManagerMode(_inManagerMode bool) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetInManagerMode(&_VaultV2b.TransactOpts, _inManagerMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_VaultV2b *VaultV2bTransactor) SetInPrivateLiquidationMode(opts *bind.TransactOpts, _inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "setInPrivateLiquidationMode", _inPrivateLiquidationMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_VaultV2b *VaultV2bSession) SetInPrivateLiquidationMode(_inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetInPrivateLiquidationMode(&_VaultV2b.TransactOpts, _inPrivateLiquidationMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0xf07bbf77.
//
// Solidity: function setInPrivateLiquidationMode(bool _inPrivateLiquidationMode) returns()
func (_VaultV2b *VaultV2bTransactorSession) SetInPrivateLiquidationMode(_inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetInPrivateLiquidationMode(&_VaultV2b.TransactOpts, _inPrivateLiquidationMode)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_VaultV2b *VaultV2bTransactor) SetIsLeverageEnabled(opts *bind.TransactOpts, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "setIsLeverageEnabled", _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_VaultV2b *VaultV2bSession) SetIsLeverageEnabled(_isLeverageEnabled bool) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetIsLeverageEnabled(&_VaultV2b.TransactOpts, _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0x7c2eb9f7.
//
// Solidity: function setIsLeverageEnabled(bool _isLeverageEnabled) returns()
func (_VaultV2b *VaultV2bTransactorSession) SetIsLeverageEnabled(_isLeverageEnabled bool) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetIsLeverageEnabled(&_VaultV2b.TransactOpts, _isLeverageEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_VaultV2b *VaultV2bTransactor) SetIsSwapEnabled(opts *bind.TransactOpts, _isSwapEnabled bool) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "setIsSwapEnabled", _isSwapEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_VaultV2b *VaultV2bSession) SetIsSwapEnabled(_isSwapEnabled bool) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetIsSwapEnabled(&_VaultV2b.TransactOpts, _isSwapEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x30455ede.
//
// Solidity: function setIsSwapEnabled(bool _isSwapEnabled) returns()
func (_VaultV2b *VaultV2bTransactorSession) SetIsSwapEnabled(_isSwapEnabled bool) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetIsSwapEnabled(&_VaultV2b.TransactOpts, _isSwapEnabled)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_VaultV2b *VaultV2bTransactor) SetLiquidator(opts *bind.TransactOpts, _liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "setLiquidator", _liquidator, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_VaultV2b *VaultV2bSession) SetLiquidator(_liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetLiquidator(&_VaultV2b.TransactOpts, _liquidator, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x4453a374.
//
// Solidity: function setLiquidator(address _liquidator, bool _isActive) returns()
func (_VaultV2b *VaultV2bTransactorSession) SetLiquidator(_liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetLiquidator(&_VaultV2b.TransactOpts, _liquidator, _isActive)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_VaultV2b *VaultV2bTransactor) SetManager(opts *bind.TransactOpts, _manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "setManager", _manager, _isManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_VaultV2b *VaultV2bSession) SetManager(_manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetManager(&_VaultV2b.TransactOpts, _manager, _isManager)
}

// SetManager is a paid mutator transaction binding the contract method 0xa5e90eee.
//
// Solidity: function setManager(address _manager, bool _isManager) returns()
func (_VaultV2b *VaultV2bTransactorSession) SetManager(_manager common.Address, _isManager bool) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetManager(&_VaultV2b.TransactOpts, _manager, _isManager)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_VaultV2b *VaultV2bTransactor) SetMaxGasPrice(opts *bind.TransactOpts, _maxGasPrice *big.Int) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "setMaxGasPrice", _maxGasPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_VaultV2b *VaultV2bSession) SetMaxGasPrice(_maxGasPrice *big.Int) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetMaxGasPrice(&_VaultV2b.TransactOpts, _maxGasPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0xd2fa635e.
//
// Solidity: function setMaxGasPrice(uint256 _maxGasPrice) returns()
func (_VaultV2b *VaultV2bTransactorSession) SetMaxGasPrice(_maxGasPrice *big.Int) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetMaxGasPrice(&_VaultV2b.TransactOpts, _maxGasPrice)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_VaultV2b *VaultV2bTransactor) SetMaxGlobalShortSize(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "setMaxGlobalShortSize", _token, _amount)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_VaultV2b *VaultV2bSession) SetMaxGlobalShortSize(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetMaxGlobalShortSize(&_VaultV2b.TransactOpts, _token, _amount)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xefa10a6e.
//
// Solidity: function setMaxGlobalShortSize(address _token, uint256 _amount) returns()
func (_VaultV2b *VaultV2bTransactorSession) SetMaxGlobalShortSize(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetMaxGlobalShortSize(&_VaultV2b.TransactOpts, _token, _amount)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_VaultV2b *VaultV2bTransactor) SetMaxLeverage(opts *bind.TransactOpts, _maxLeverage *big.Int) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "setMaxLeverage", _maxLeverage)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_VaultV2b *VaultV2bSession) SetMaxLeverage(_maxLeverage *big.Int) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetMaxLeverage(&_VaultV2b.TransactOpts, _maxLeverage)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0xd3127e63.
//
// Solidity: function setMaxLeverage(uint256 _maxLeverage) returns()
func (_VaultV2b *VaultV2bTransactorSession) SetMaxLeverage(_maxLeverage *big.Int) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetMaxLeverage(&_VaultV2b.TransactOpts, _maxLeverage)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_VaultV2b *VaultV2bTransactor) SetPriceFeed(opts *bind.TransactOpts, _priceFeed common.Address) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "setPriceFeed", _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_VaultV2b *VaultV2bSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetPriceFeed(&_VaultV2b.TransactOpts, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_VaultV2b *VaultV2bTransactorSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetPriceFeed(&_VaultV2b.TransactOpts, _priceFeed)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_VaultV2b *VaultV2bTransactor) SetTokenConfig(opts *bind.TransactOpts, _token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "setTokenConfig", _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_VaultV2b *VaultV2bSession) SetTokenConfig(_token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetTokenConfig(&_VaultV2b.TransactOpts, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x3c5a6e35.
//
// Solidity: function setTokenConfig(address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_VaultV2b *VaultV2bTransactorSession) SetTokenConfig(_token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetTokenConfig(&_VaultV2b.TransactOpts, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SetUsdgAmount is a paid mutator transaction binding the contract method 0xd66b000d.
//
// Solidity: function setUsdgAmount(address _token, uint256 _amount) returns()
func (_VaultV2b *VaultV2bTransactor) SetUsdgAmount(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "setUsdgAmount", _token, _amount)
}

// SetUsdgAmount is a paid mutator transaction binding the contract method 0xd66b000d.
//
// Solidity: function setUsdgAmount(address _token, uint256 _amount) returns()
func (_VaultV2b *VaultV2bSession) SetUsdgAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetUsdgAmount(&_VaultV2b.TransactOpts, _token, _amount)
}

// SetUsdgAmount is a paid mutator transaction binding the contract method 0xd66b000d.
//
// Solidity: function setUsdgAmount(address _token, uint256 _amount) returns()
func (_VaultV2b *VaultV2bTransactorSession) SetUsdgAmount(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetUsdgAmount(&_VaultV2b.TransactOpts, _token, _amount)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_VaultV2b *VaultV2bTransactor) SetVaultUtils(opts *bind.TransactOpts, _vaultUtils common.Address) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "setVaultUtils", _vaultUtils)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_VaultV2b *VaultV2bSession) SetVaultUtils(_vaultUtils common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetVaultUtils(&_VaultV2b.TransactOpts, _vaultUtils)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0x71089f4d.
//
// Solidity: function setVaultUtils(address _vaultUtils) returns()
func (_VaultV2b *VaultV2bTransactorSession) SetVaultUtils(_vaultUtils common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.SetVaultUtils(&_VaultV2b.TransactOpts, _vaultUtils)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_VaultV2b *VaultV2bTransactor) Swap(opts *bind.TransactOpts, _tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "swap", _tokenIn, _tokenOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_VaultV2b *VaultV2bSession) Swap(_tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.Swap(&_VaultV2b.TransactOpts, _tokenIn, _tokenOut, _receiver)
}

// Swap is a paid mutator transaction binding the contract method 0x93316212.
//
// Solidity: function swap(address _tokenIn, address _tokenOut, address _receiver) returns(uint256)
func (_VaultV2b *VaultV2bTransactorSession) Swap(_tokenIn common.Address, _tokenOut common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.Swap(&_VaultV2b.TransactOpts, _tokenIn, _tokenOut, _receiver)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns()
func (_VaultV2b *VaultV2bTransactor) UpdateCumulativeFundingRate(opts *bind.TransactOpts, _collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "updateCumulativeFundingRate", _collateralToken, _indexToken)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns()
func (_VaultV2b *VaultV2bSession) UpdateCumulativeFundingRate(_collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.UpdateCumulativeFundingRate(&_VaultV2b.TransactOpts, _collateralToken, _indexToken)
}

// UpdateCumulativeFundingRate is a paid mutator transaction binding the contract method 0xfbfded6d.
//
// Solidity: function updateCumulativeFundingRate(address _collateralToken, address _indexToken) returns()
func (_VaultV2b *VaultV2bTransactorSession) UpdateCumulativeFundingRate(_collateralToken common.Address, _indexToken common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.UpdateCumulativeFundingRate(&_VaultV2b.TransactOpts, _collateralToken, _indexToken)
}

// UpgradeVault is a paid mutator transaction binding the contract method 0xcea0c328.
//
// Solidity: function upgradeVault(address _newVault, address _token, uint256 _amount) returns()
func (_VaultV2b *VaultV2bTransactor) UpgradeVault(opts *bind.TransactOpts, _newVault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "upgradeVault", _newVault, _token, _amount)
}

// UpgradeVault is a paid mutator transaction binding the contract method 0xcea0c328.
//
// Solidity: function upgradeVault(address _newVault, address _token, uint256 _amount) returns()
func (_VaultV2b *VaultV2bSession) UpgradeVault(_newVault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2b.Contract.UpgradeVault(&_VaultV2b.TransactOpts, _newVault, _token, _amount)
}

// UpgradeVault is a paid mutator transaction binding the contract method 0xcea0c328.
//
// Solidity: function upgradeVault(address _newVault, address _token, uint256 _amount) returns()
func (_VaultV2b *VaultV2bTransactorSession) UpgradeVault(_newVault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _VaultV2b.Contract.UpgradeVault(&_VaultV2b.TransactOpts, _newVault, _token, _amount)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_VaultV2b *VaultV2bTransactor) WithdrawFees(opts *bind.TransactOpts, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2b.contract.Transact(opts, "withdrawFees", _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_VaultV2b *VaultV2bSession) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.WithdrawFees(&_VaultV2b.TransactOpts, _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0xf2555278.
//
// Solidity: function withdrawFees(address _token, address _receiver) returns(uint256)
func (_VaultV2b *VaultV2bTransactorSession) WithdrawFees(_token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _VaultV2b.Contract.WithdrawFees(&_VaultV2b.TransactOpts, _token, _receiver)
}

// VaultV2bBuyUSDGIterator is returned from FilterBuyUSDG and is used to iterate over the raw logs and unpacked data for BuyUSDG events raised by the VaultV2b contract.
type VaultV2bBuyUSDGIterator struct {
	Event *VaultV2bBuyUSDG // Event containing the contract specifics and raw log

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
func (it *VaultV2bBuyUSDGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bBuyUSDG)
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
		it.Event = new(VaultV2bBuyUSDG)
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
func (it *VaultV2bBuyUSDGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bBuyUSDGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bBuyUSDG represents a BuyUSDG event raised by the VaultV2b contract.
type VaultV2bBuyUSDG struct {
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
func (_VaultV2b *VaultV2bFilterer) FilterBuyUSDG(opts *bind.FilterOpts) (*VaultV2bBuyUSDGIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "BuyUSDG")
	if err != nil {
		return nil, err
	}
	return &VaultV2bBuyUSDGIterator{contract: _VaultV2b.contract, event: "BuyUSDG", logs: logs, sub: sub}, nil
}

// WatchBuyUSDG is a free log subscription operation binding the contract event 0xab4c77c74cd32c85f35416cf03e7ce9e2d4387f7b7f2c1f4bf53daaecf8ea72d.
//
// Solidity: event BuyUSDG(address account, address token, uint256 tokenAmount, uint256 usdgAmount, uint256 feeBasisPoints)
func (_VaultV2b *VaultV2bFilterer) WatchBuyUSDG(opts *bind.WatchOpts, sink chan<- *VaultV2bBuyUSDG) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "BuyUSDG")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bBuyUSDG)
				if err := _VaultV2b.contract.UnpackLog(event, "BuyUSDG", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseBuyUSDG(log types.Log) (*VaultV2bBuyUSDG, error) {
	event := new(VaultV2bBuyUSDG)
	if err := _VaultV2b.contract.UnpackLog(event, "BuyUSDG", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bClosePositionIterator is returned from FilterClosePosition and is used to iterate over the raw logs and unpacked data for ClosePosition events raised by the VaultV2b contract.
type VaultV2bClosePositionIterator struct {
	Event *VaultV2bClosePosition // Event containing the contract specifics and raw log

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
func (it *VaultV2bClosePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bClosePosition)
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
		it.Event = new(VaultV2bClosePosition)
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
func (it *VaultV2bClosePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bClosePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bClosePosition represents a ClosePosition event raised by the VaultV2b contract.
type VaultV2bClosePosition struct {
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
func (_VaultV2b *VaultV2bFilterer) FilterClosePosition(opts *bind.FilterOpts) (*VaultV2bClosePositionIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "ClosePosition")
	if err != nil {
		return nil, err
	}
	return &VaultV2bClosePositionIterator{contract: _VaultV2b.contract, event: "ClosePosition", logs: logs, sub: sub}, nil
}

// WatchClosePosition is a free log subscription operation binding the contract event 0x73af1d417d82c240fdb6d319b34ad884487c6bf2845d98980cc52ad9171cb455.
//
// Solidity: event ClosePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl)
func (_VaultV2b *VaultV2bFilterer) WatchClosePosition(opts *bind.WatchOpts, sink chan<- *VaultV2bClosePosition) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "ClosePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bClosePosition)
				if err := _VaultV2b.contract.UnpackLog(event, "ClosePosition", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseClosePosition(log types.Log) (*VaultV2bClosePosition, error) {
	event := new(VaultV2bClosePosition)
	if err := _VaultV2b.contract.UnpackLog(event, "ClosePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bCollectMarginFeesIterator is returned from FilterCollectMarginFees and is used to iterate over the raw logs and unpacked data for CollectMarginFees events raised by the VaultV2b contract.
type VaultV2bCollectMarginFeesIterator struct {
	Event *VaultV2bCollectMarginFees // Event containing the contract specifics and raw log

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
func (it *VaultV2bCollectMarginFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bCollectMarginFees)
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
		it.Event = new(VaultV2bCollectMarginFees)
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
func (it *VaultV2bCollectMarginFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bCollectMarginFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bCollectMarginFees represents a CollectMarginFees event raised by the VaultV2b contract.
type VaultV2bCollectMarginFees struct {
	Token     common.Address
	FeeUsd    *big.Int
	FeeTokens *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectMarginFees is a free log retrieval operation binding the contract event 0x5d0c0019d3d45fadeb74eff9d2c9924d146d000ac6bcf3c28bf0ac3c9baa011a.
//
// Solidity: event CollectMarginFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_VaultV2b *VaultV2bFilterer) FilterCollectMarginFees(opts *bind.FilterOpts) (*VaultV2bCollectMarginFeesIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "CollectMarginFees")
	if err != nil {
		return nil, err
	}
	return &VaultV2bCollectMarginFeesIterator{contract: _VaultV2b.contract, event: "CollectMarginFees", logs: logs, sub: sub}, nil
}

// WatchCollectMarginFees is a free log subscription operation binding the contract event 0x5d0c0019d3d45fadeb74eff9d2c9924d146d000ac6bcf3c28bf0ac3c9baa011a.
//
// Solidity: event CollectMarginFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_VaultV2b *VaultV2bFilterer) WatchCollectMarginFees(opts *bind.WatchOpts, sink chan<- *VaultV2bCollectMarginFees) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "CollectMarginFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bCollectMarginFees)
				if err := _VaultV2b.contract.UnpackLog(event, "CollectMarginFees", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseCollectMarginFees(log types.Log) (*VaultV2bCollectMarginFees, error) {
	event := new(VaultV2bCollectMarginFees)
	if err := _VaultV2b.contract.UnpackLog(event, "CollectMarginFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bCollectSwapFeesIterator is returned from FilterCollectSwapFees and is used to iterate over the raw logs and unpacked data for CollectSwapFees events raised by the VaultV2b contract.
type VaultV2bCollectSwapFeesIterator struct {
	Event *VaultV2bCollectSwapFees // Event containing the contract specifics and raw log

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
func (it *VaultV2bCollectSwapFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bCollectSwapFees)
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
		it.Event = new(VaultV2bCollectSwapFees)
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
func (it *VaultV2bCollectSwapFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bCollectSwapFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bCollectSwapFees represents a CollectSwapFees event raised by the VaultV2b contract.
type VaultV2bCollectSwapFees struct {
	Token     common.Address
	FeeUsd    *big.Int
	FeeTokens *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCollectSwapFees is a free log retrieval operation binding the contract event 0x47cd9dda0e50ce30bcaaacd0488452b596221c07ac402a581cfae4d3933cac2b.
//
// Solidity: event CollectSwapFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_VaultV2b *VaultV2bFilterer) FilterCollectSwapFees(opts *bind.FilterOpts) (*VaultV2bCollectSwapFeesIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "CollectSwapFees")
	if err != nil {
		return nil, err
	}
	return &VaultV2bCollectSwapFeesIterator{contract: _VaultV2b.contract, event: "CollectSwapFees", logs: logs, sub: sub}, nil
}

// WatchCollectSwapFees is a free log subscription operation binding the contract event 0x47cd9dda0e50ce30bcaaacd0488452b596221c07ac402a581cfae4d3933cac2b.
//
// Solidity: event CollectSwapFees(address token, uint256 feeUsd, uint256 feeTokens)
func (_VaultV2b *VaultV2bFilterer) WatchCollectSwapFees(opts *bind.WatchOpts, sink chan<- *VaultV2bCollectSwapFees) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "CollectSwapFees")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bCollectSwapFees)
				if err := _VaultV2b.contract.UnpackLog(event, "CollectSwapFees", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseCollectSwapFees(log types.Log) (*VaultV2bCollectSwapFees, error) {
	event := new(VaultV2bCollectSwapFees)
	if err := _VaultV2b.contract.UnpackLog(event, "CollectSwapFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bDecreaseGuaranteedUsdIterator is returned from FilterDecreaseGuaranteedUsd and is used to iterate over the raw logs and unpacked data for DecreaseGuaranteedUsd events raised by the VaultV2b contract.
type VaultV2bDecreaseGuaranteedUsdIterator struct {
	Event *VaultV2bDecreaseGuaranteedUsd // Event containing the contract specifics and raw log

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
func (it *VaultV2bDecreaseGuaranteedUsdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bDecreaseGuaranteedUsd)
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
		it.Event = new(VaultV2bDecreaseGuaranteedUsd)
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
func (it *VaultV2bDecreaseGuaranteedUsdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bDecreaseGuaranteedUsdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bDecreaseGuaranteedUsd represents a DecreaseGuaranteedUsd event raised by the VaultV2b contract.
type VaultV2bDecreaseGuaranteedUsd struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreaseGuaranteedUsd is a free log retrieval operation binding the contract event 0x34e07158b9db50df5613e591c44ea2ebc82834eff4a4dc3a46e000e608261d68.
//
// Solidity: event DecreaseGuaranteedUsd(address token, uint256 amount)
func (_VaultV2b *VaultV2bFilterer) FilterDecreaseGuaranteedUsd(opts *bind.FilterOpts) (*VaultV2bDecreaseGuaranteedUsdIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "DecreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return &VaultV2bDecreaseGuaranteedUsdIterator{contract: _VaultV2b.contract, event: "DecreaseGuaranteedUsd", logs: logs, sub: sub}, nil
}

// WatchDecreaseGuaranteedUsd is a free log subscription operation binding the contract event 0x34e07158b9db50df5613e591c44ea2ebc82834eff4a4dc3a46e000e608261d68.
//
// Solidity: event DecreaseGuaranteedUsd(address token, uint256 amount)
func (_VaultV2b *VaultV2bFilterer) WatchDecreaseGuaranteedUsd(opts *bind.WatchOpts, sink chan<- *VaultV2bDecreaseGuaranteedUsd) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "DecreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bDecreaseGuaranteedUsd)
				if err := _VaultV2b.contract.UnpackLog(event, "DecreaseGuaranteedUsd", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseDecreaseGuaranteedUsd(log types.Log) (*VaultV2bDecreaseGuaranteedUsd, error) {
	event := new(VaultV2bDecreaseGuaranteedUsd)
	if err := _VaultV2b.contract.UnpackLog(event, "DecreaseGuaranteedUsd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bDecreasePoolAmountIterator is returned from FilterDecreasePoolAmount and is used to iterate over the raw logs and unpacked data for DecreasePoolAmount events raised by the VaultV2b contract.
type VaultV2bDecreasePoolAmountIterator struct {
	Event *VaultV2bDecreasePoolAmount // Event containing the contract specifics and raw log

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
func (it *VaultV2bDecreasePoolAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bDecreasePoolAmount)
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
		it.Event = new(VaultV2bDecreasePoolAmount)
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
func (it *VaultV2bDecreasePoolAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bDecreasePoolAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bDecreasePoolAmount represents a DecreasePoolAmount event raised by the VaultV2b contract.
type VaultV2bDecreasePoolAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreasePoolAmount is a free log retrieval operation binding the contract event 0x112726233fbeaeed0f5b1dba5cb0b2b81883dee49fb35ff99fd98ed9f6d31eb0.
//
// Solidity: event DecreasePoolAmount(address token, uint256 amount)
func (_VaultV2b *VaultV2bFilterer) FilterDecreasePoolAmount(opts *bind.FilterOpts) (*VaultV2bDecreasePoolAmountIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "DecreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return &VaultV2bDecreasePoolAmountIterator{contract: _VaultV2b.contract, event: "DecreasePoolAmount", logs: logs, sub: sub}, nil
}

// WatchDecreasePoolAmount is a free log subscription operation binding the contract event 0x112726233fbeaeed0f5b1dba5cb0b2b81883dee49fb35ff99fd98ed9f6d31eb0.
//
// Solidity: event DecreasePoolAmount(address token, uint256 amount)
func (_VaultV2b *VaultV2bFilterer) WatchDecreasePoolAmount(opts *bind.WatchOpts, sink chan<- *VaultV2bDecreasePoolAmount) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "DecreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bDecreasePoolAmount)
				if err := _VaultV2b.contract.UnpackLog(event, "DecreasePoolAmount", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseDecreasePoolAmount(log types.Log) (*VaultV2bDecreasePoolAmount, error) {
	event := new(VaultV2bDecreasePoolAmount)
	if err := _VaultV2b.contract.UnpackLog(event, "DecreasePoolAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bDecreasePositionIterator is returned from FilterDecreasePosition and is used to iterate over the raw logs and unpacked data for DecreasePosition events raised by the VaultV2b contract.
type VaultV2bDecreasePositionIterator struct {
	Event *VaultV2bDecreasePosition // Event containing the contract specifics and raw log

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
func (it *VaultV2bDecreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bDecreasePosition)
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
		it.Event = new(VaultV2bDecreasePosition)
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
func (it *VaultV2bDecreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bDecreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bDecreasePosition represents a DecreasePosition event raised by the VaultV2b contract.
type VaultV2bDecreasePosition struct {
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
func (_VaultV2b *VaultV2bFilterer) FilterDecreasePosition(opts *bind.FilterOpts) (*VaultV2bDecreasePositionIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "DecreasePosition")
	if err != nil {
		return nil, err
	}
	return &VaultV2bDecreasePositionIterator{contract: _VaultV2b.contract, event: "DecreasePosition", logs: logs, sub: sub}, nil
}

// WatchDecreasePosition is a free log subscription operation binding the contract event 0x93d75d64d1f84fc6f430a64fc578bdd4c1e090e90ea2d51773e626d19de56d30.
//
// Solidity: event DecreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_VaultV2b *VaultV2bFilterer) WatchDecreasePosition(opts *bind.WatchOpts, sink chan<- *VaultV2bDecreasePosition) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "DecreasePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bDecreasePosition)
				if err := _VaultV2b.contract.UnpackLog(event, "DecreasePosition", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseDecreasePosition(log types.Log) (*VaultV2bDecreasePosition, error) {
	event := new(VaultV2bDecreasePosition)
	if err := _VaultV2b.contract.UnpackLog(event, "DecreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bDecreaseReservedAmountIterator is returned from FilterDecreaseReservedAmount and is used to iterate over the raw logs and unpacked data for DecreaseReservedAmount events raised by the VaultV2b contract.
type VaultV2bDecreaseReservedAmountIterator struct {
	Event *VaultV2bDecreaseReservedAmount // Event containing the contract specifics and raw log

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
func (it *VaultV2bDecreaseReservedAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bDecreaseReservedAmount)
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
		it.Event = new(VaultV2bDecreaseReservedAmount)
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
func (it *VaultV2bDecreaseReservedAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bDecreaseReservedAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bDecreaseReservedAmount represents a DecreaseReservedAmount event raised by the VaultV2b contract.
type VaultV2bDecreaseReservedAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreaseReservedAmount is a free log retrieval operation binding the contract event 0x533cb5ed32be6a90284e96b5747a1bfc2d38fdb5768a6b5f67ff7d62144ed67b.
//
// Solidity: event DecreaseReservedAmount(address token, uint256 amount)
func (_VaultV2b *VaultV2bFilterer) FilterDecreaseReservedAmount(opts *bind.FilterOpts) (*VaultV2bDecreaseReservedAmountIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "DecreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return &VaultV2bDecreaseReservedAmountIterator{contract: _VaultV2b.contract, event: "DecreaseReservedAmount", logs: logs, sub: sub}, nil
}

// WatchDecreaseReservedAmount is a free log subscription operation binding the contract event 0x533cb5ed32be6a90284e96b5747a1bfc2d38fdb5768a6b5f67ff7d62144ed67b.
//
// Solidity: event DecreaseReservedAmount(address token, uint256 amount)
func (_VaultV2b *VaultV2bFilterer) WatchDecreaseReservedAmount(opts *bind.WatchOpts, sink chan<- *VaultV2bDecreaseReservedAmount) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "DecreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bDecreaseReservedAmount)
				if err := _VaultV2b.contract.UnpackLog(event, "DecreaseReservedAmount", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseDecreaseReservedAmount(log types.Log) (*VaultV2bDecreaseReservedAmount, error) {
	event := new(VaultV2bDecreaseReservedAmount)
	if err := _VaultV2b.contract.UnpackLog(event, "DecreaseReservedAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bDecreaseUsdgAmountIterator is returned from FilterDecreaseUsdgAmount and is used to iterate over the raw logs and unpacked data for DecreaseUsdgAmount events raised by the VaultV2b contract.
type VaultV2bDecreaseUsdgAmountIterator struct {
	Event *VaultV2bDecreaseUsdgAmount // Event containing the contract specifics and raw log

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
func (it *VaultV2bDecreaseUsdgAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bDecreaseUsdgAmount)
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
		it.Event = new(VaultV2bDecreaseUsdgAmount)
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
func (it *VaultV2bDecreaseUsdgAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bDecreaseUsdgAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bDecreaseUsdgAmount represents a DecreaseUsdgAmount event raised by the VaultV2b contract.
type VaultV2bDecreaseUsdgAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDecreaseUsdgAmount is a free log retrieval operation binding the contract event 0xe1e812596aac93a06ecc4ca627014d18e30f5c33b825160cc9d5c0ba61e45227.
//
// Solidity: event DecreaseUsdgAmount(address token, uint256 amount)
func (_VaultV2b *VaultV2bFilterer) FilterDecreaseUsdgAmount(opts *bind.FilterOpts) (*VaultV2bDecreaseUsdgAmountIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "DecreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return &VaultV2bDecreaseUsdgAmountIterator{contract: _VaultV2b.contract, event: "DecreaseUsdgAmount", logs: logs, sub: sub}, nil
}

// WatchDecreaseUsdgAmount is a free log subscription operation binding the contract event 0xe1e812596aac93a06ecc4ca627014d18e30f5c33b825160cc9d5c0ba61e45227.
//
// Solidity: event DecreaseUsdgAmount(address token, uint256 amount)
func (_VaultV2b *VaultV2bFilterer) WatchDecreaseUsdgAmount(opts *bind.WatchOpts, sink chan<- *VaultV2bDecreaseUsdgAmount) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "DecreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bDecreaseUsdgAmount)
				if err := _VaultV2b.contract.UnpackLog(event, "DecreaseUsdgAmount", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseDecreaseUsdgAmount(log types.Log) (*VaultV2bDecreaseUsdgAmount, error) {
	event := new(VaultV2bDecreaseUsdgAmount)
	if err := _VaultV2b.contract.UnpackLog(event, "DecreaseUsdgAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bDirectPoolDepositIterator is returned from FilterDirectPoolDeposit and is used to iterate over the raw logs and unpacked data for DirectPoolDeposit events raised by the VaultV2b contract.
type VaultV2bDirectPoolDepositIterator struct {
	Event *VaultV2bDirectPoolDeposit // Event containing the contract specifics and raw log

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
func (it *VaultV2bDirectPoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bDirectPoolDeposit)
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
		it.Event = new(VaultV2bDirectPoolDeposit)
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
func (it *VaultV2bDirectPoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bDirectPoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bDirectPoolDeposit represents a DirectPoolDeposit event raised by the VaultV2b contract.
type VaultV2bDirectPoolDeposit struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDirectPoolDeposit is a free log retrieval operation binding the contract event 0xa5a389190ebf6170a133bda5c769b77f4d6715b8aa172ec0ddf8473d0b4944bd.
//
// Solidity: event DirectPoolDeposit(address token, uint256 amount)
func (_VaultV2b *VaultV2bFilterer) FilterDirectPoolDeposit(opts *bind.FilterOpts) (*VaultV2bDirectPoolDepositIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "DirectPoolDeposit")
	if err != nil {
		return nil, err
	}
	return &VaultV2bDirectPoolDepositIterator{contract: _VaultV2b.contract, event: "DirectPoolDeposit", logs: logs, sub: sub}, nil
}

// WatchDirectPoolDeposit is a free log subscription operation binding the contract event 0xa5a389190ebf6170a133bda5c769b77f4d6715b8aa172ec0ddf8473d0b4944bd.
//
// Solidity: event DirectPoolDeposit(address token, uint256 amount)
func (_VaultV2b *VaultV2bFilterer) WatchDirectPoolDeposit(opts *bind.WatchOpts, sink chan<- *VaultV2bDirectPoolDeposit) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "DirectPoolDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bDirectPoolDeposit)
				if err := _VaultV2b.contract.UnpackLog(event, "DirectPoolDeposit", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseDirectPoolDeposit(log types.Log) (*VaultV2bDirectPoolDeposit, error) {
	event := new(VaultV2bDirectPoolDeposit)
	if err := _VaultV2b.contract.UnpackLog(event, "DirectPoolDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bIncreaseGuaranteedUsdIterator is returned from FilterIncreaseGuaranteedUsd and is used to iterate over the raw logs and unpacked data for IncreaseGuaranteedUsd events raised by the VaultV2b contract.
type VaultV2bIncreaseGuaranteedUsdIterator struct {
	Event *VaultV2bIncreaseGuaranteedUsd // Event containing the contract specifics and raw log

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
func (it *VaultV2bIncreaseGuaranteedUsdIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bIncreaseGuaranteedUsd)
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
		it.Event = new(VaultV2bIncreaseGuaranteedUsd)
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
func (it *VaultV2bIncreaseGuaranteedUsdIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bIncreaseGuaranteedUsdIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bIncreaseGuaranteedUsd represents a IncreaseGuaranteedUsd event raised by the VaultV2b contract.
type VaultV2bIncreaseGuaranteedUsd struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreaseGuaranteedUsd is a free log retrieval operation binding the contract event 0xd9d4761f75e0d0103b5cbeab941eeb443d7a56a35b5baf2a0787c03f03f4e474.
//
// Solidity: event IncreaseGuaranteedUsd(address token, uint256 amount)
func (_VaultV2b *VaultV2bFilterer) FilterIncreaseGuaranteedUsd(opts *bind.FilterOpts) (*VaultV2bIncreaseGuaranteedUsdIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "IncreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return &VaultV2bIncreaseGuaranteedUsdIterator{contract: _VaultV2b.contract, event: "IncreaseGuaranteedUsd", logs: logs, sub: sub}, nil
}

// WatchIncreaseGuaranteedUsd is a free log subscription operation binding the contract event 0xd9d4761f75e0d0103b5cbeab941eeb443d7a56a35b5baf2a0787c03f03f4e474.
//
// Solidity: event IncreaseGuaranteedUsd(address token, uint256 amount)
func (_VaultV2b *VaultV2bFilterer) WatchIncreaseGuaranteedUsd(opts *bind.WatchOpts, sink chan<- *VaultV2bIncreaseGuaranteedUsd) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "IncreaseGuaranteedUsd")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bIncreaseGuaranteedUsd)
				if err := _VaultV2b.contract.UnpackLog(event, "IncreaseGuaranteedUsd", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseIncreaseGuaranteedUsd(log types.Log) (*VaultV2bIncreaseGuaranteedUsd, error) {
	event := new(VaultV2bIncreaseGuaranteedUsd)
	if err := _VaultV2b.contract.UnpackLog(event, "IncreaseGuaranteedUsd", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bIncreasePoolAmountIterator is returned from FilterIncreasePoolAmount and is used to iterate over the raw logs and unpacked data for IncreasePoolAmount events raised by the VaultV2b contract.
type VaultV2bIncreasePoolAmountIterator struct {
	Event *VaultV2bIncreasePoolAmount // Event containing the contract specifics and raw log

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
func (it *VaultV2bIncreasePoolAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bIncreasePoolAmount)
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
		it.Event = new(VaultV2bIncreasePoolAmount)
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
func (it *VaultV2bIncreasePoolAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bIncreasePoolAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bIncreasePoolAmount represents a IncreasePoolAmount event raised by the VaultV2b contract.
type VaultV2bIncreasePoolAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreasePoolAmount is a free log retrieval operation binding the contract event 0x976177fbe09a15e5e43f848844963a42b41ef919ef17ff21a17a5421de8f4737.
//
// Solidity: event IncreasePoolAmount(address token, uint256 amount)
func (_VaultV2b *VaultV2bFilterer) FilterIncreasePoolAmount(opts *bind.FilterOpts) (*VaultV2bIncreasePoolAmountIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "IncreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return &VaultV2bIncreasePoolAmountIterator{contract: _VaultV2b.contract, event: "IncreasePoolAmount", logs: logs, sub: sub}, nil
}

// WatchIncreasePoolAmount is a free log subscription operation binding the contract event 0x976177fbe09a15e5e43f848844963a42b41ef919ef17ff21a17a5421de8f4737.
//
// Solidity: event IncreasePoolAmount(address token, uint256 amount)
func (_VaultV2b *VaultV2bFilterer) WatchIncreasePoolAmount(opts *bind.WatchOpts, sink chan<- *VaultV2bIncreasePoolAmount) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "IncreasePoolAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bIncreasePoolAmount)
				if err := _VaultV2b.contract.UnpackLog(event, "IncreasePoolAmount", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseIncreasePoolAmount(log types.Log) (*VaultV2bIncreasePoolAmount, error) {
	event := new(VaultV2bIncreasePoolAmount)
	if err := _VaultV2b.contract.UnpackLog(event, "IncreasePoolAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bIncreasePositionIterator is returned from FilterIncreasePosition and is used to iterate over the raw logs and unpacked data for IncreasePosition events raised by the VaultV2b contract.
type VaultV2bIncreasePositionIterator struct {
	Event *VaultV2bIncreasePosition // Event containing the contract specifics and raw log

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
func (it *VaultV2bIncreasePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bIncreasePosition)
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
		it.Event = new(VaultV2bIncreasePosition)
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
func (it *VaultV2bIncreasePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bIncreasePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bIncreasePosition represents a IncreasePosition event raised by the VaultV2b contract.
type VaultV2bIncreasePosition struct {
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
func (_VaultV2b *VaultV2bFilterer) FilterIncreasePosition(opts *bind.FilterOpts) (*VaultV2bIncreasePositionIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "IncreasePosition")
	if err != nil {
		return nil, err
	}
	return &VaultV2bIncreasePositionIterator{contract: _VaultV2b.contract, event: "IncreasePosition", logs: logs, sub: sub}, nil
}

// WatchIncreasePosition is a free log subscription operation binding the contract event 0x2fe68525253654c21998f35787a8d0f361905ef647c854092430ab65f2f15022.
//
// Solidity: event IncreasePosition(bytes32 key, address account, address collateralToken, address indexToken, uint256 collateralDelta, uint256 sizeDelta, bool isLong, uint256 price, uint256 fee)
func (_VaultV2b *VaultV2bFilterer) WatchIncreasePosition(opts *bind.WatchOpts, sink chan<- *VaultV2bIncreasePosition) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "IncreasePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bIncreasePosition)
				if err := _VaultV2b.contract.UnpackLog(event, "IncreasePosition", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseIncreasePosition(log types.Log) (*VaultV2bIncreasePosition, error) {
	event := new(VaultV2bIncreasePosition)
	if err := _VaultV2b.contract.UnpackLog(event, "IncreasePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bIncreaseReservedAmountIterator is returned from FilterIncreaseReservedAmount and is used to iterate over the raw logs and unpacked data for IncreaseReservedAmount events raised by the VaultV2b contract.
type VaultV2bIncreaseReservedAmountIterator struct {
	Event *VaultV2bIncreaseReservedAmount // Event containing the contract specifics and raw log

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
func (it *VaultV2bIncreaseReservedAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bIncreaseReservedAmount)
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
		it.Event = new(VaultV2bIncreaseReservedAmount)
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
func (it *VaultV2bIncreaseReservedAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bIncreaseReservedAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bIncreaseReservedAmount represents a IncreaseReservedAmount event raised by the VaultV2b contract.
type VaultV2bIncreaseReservedAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreaseReservedAmount is a free log retrieval operation binding the contract event 0xaa5649d82f5462be9d19b0f2b31a59b2259950a6076550bac9f3a1c07db9f66d.
//
// Solidity: event IncreaseReservedAmount(address token, uint256 amount)
func (_VaultV2b *VaultV2bFilterer) FilterIncreaseReservedAmount(opts *bind.FilterOpts) (*VaultV2bIncreaseReservedAmountIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "IncreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return &VaultV2bIncreaseReservedAmountIterator{contract: _VaultV2b.contract, event: "IncreaseReservedAmount", logs: logs, sub: sub}, nil
}

// WatchIncreaseReservedAmount is a free log subscription operation binding the contract event 0xaa5649d82f5462be9d19b0f2b31a59b2259950a6076550bac9f3a1c07db9f66d.
//
// Solidity: event IncreaseReservedAmount(address token, uint256 amount)
func (_VaultV2b *VaultV2bFilterer) WatchIncreaseReservedAmount(opts *bind.WatchOpts, sink chan<- *VaultV2bIncreaseReservedAmount) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "IncreaseReservedAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bIncreaseReservedAmount)
				if err := _VaultV2b.contract.UnpackLog(event, "IncreaseReservedAmount", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseIncreaseReservedAmount(log types.Log) (*VaultV2bIncreaseReservedAmount, error) {
	event := new(VaultV2bIncreaseReservedAmount)
	if err := _VaultV2b.contract.UnpackLog(event, "IncreaseReservedAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bIncreaseUsdgAmountIterator is returned from FilterIncreaseUsdgAmount and is used to iterate over the raw logs and unpacked data for IncreaseUsdgAmount events raised by the VaultV2b contract.
type VaultV2bIncreaseUsdgAmountIterator struct {
	Event *VaultV2bIncreaseUsdgAmount // Event containing the contract specifics and raw log

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
func (it *VaultV2bIncreaseUsdgAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bIncreaseUsdgAmount)
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
		it.Event = new(VaultV2bIncreaseUsdgAmount)
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
func (it *VaultV2bIncreaseUsdgAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bIncreaseUsdgAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bIncreaseUsdgAmount represents a IncreaseUsdgAmount event raised by the VaultV2b contract.
type VaultV2bIncreaseUsdgAmount struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIncreaseUsdgAmount is a free log retrieval operation binding the contract event 0x64243679a443432e2293343b77d411ff6144370404618f00ca0d2025d9ca9882.
//
// Solidity: event IncreaseUsdgAmount(address token, uint256 amount)
func (_VaultV2b *VaultV2bFilterer) FilterIncreaseUsdgAmount(opts *bind.FilterOpts) (*VaultV2bIncreaseUsdgAmountIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "IncreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return &VaultV2bIncreaseUsdgAmountIterator{contract: _VaultV2b.contract, event: "IncreaseUsdgAmount", logs: logs, sub: sub}, nil
}

// WatchIncreaseUsdgAmount is a free log subscription operation binding the contract event 0x64243679a443432e2293343b77d411ff6144370404618f00ca0d2025d9ca9882.
//
// Solidity: event IncreaseUsdgAmount(address token, uint256 amount)
func (_VaultV2b *VaultV2bFilterer) WatchIncreaseUsdgAmount(opts *bind.WatchOpts, sink chan<- *VaultV2bIncreaseUsdgAmount) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "IncreaseUsdgAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bIncreaseUsdgAmount)
				if err := _VaultV2b.contract.UnpackLog(event, "IncreaseUsdgAmount", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseIncreaseUsdgAmount(log types.Log) (*VaultV2bIncreaseUsdgAmount, error) {
	event := new(VaultV2bIncreaseUsdgAmount)
	if err := _VaultV2b.contract.UnpackLog(event, "IncreaseUsdgAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bLiquidatePositionIterator is returned from FilterLiquidatePosition and is used to iterate over the raw logs and unpacked data for LiquidatePosition events raised by the VaultV2b contract.
type VaultV2bLiquidatePositionIterator struct {
	Event *VaultV2bLiquidatePosition // Event containing the contract specifics and raw log

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
func (it *VaultV2bLiquidatePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bLiquidatePosition)
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
		it.Event = new(VaultV2bLiquidatePosition)
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
func (it *VaultV2bLiquidatePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bLiquidatePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bLiquidatePosition represents a LiquidatePosition event raised by the VaultV2b contract.
type VaultV2bLiquidatePosition struct {
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
func (_VaultV2b *VaultV2bFilterer) FilterLiquidatePosition(opts *bind.FilterOpts) (*VaultV2bLiquidatePositionIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "LiquidatePosition")
	if err != nil {
		return nil, err
	}
	return &VaultV2bLiquidatePositionIterator{contract: _VaultV2b.contract, event: "LiquidatePosition", logs: logs, sub: sub}, nil
}

// WatchLiquidatePosition is a free log subscription operation binding the contract event 0x2e1f85a64a2f22cf2f0c42584e7c919ed4abe8d53675cff0f62bf1e95a1c676f.
//
// Solidity: event LiquidatePosition(bytes32 key, address account, address collateralToken, address indexToken, bool isLong, uint256 size, uint256 collateral, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_VaultV2b *VaultV2bFilterer) WatchLiquidatePosition(opts *bind.WatchOpts, sink chan<- *VaultV2bLiquidatePosition) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "LiquidatePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bLiquidatePosition)
				if err := _VaultV2b.contract.UnpackLog(event, "LiquidatePosition", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseLiquidatePosition(log types.Log) (*VaultV2bLiquidatePosition, error) {
	event := new(VaultV2bLiquidatePosition)
	if err := _VaultV2b.contract.UnpackLog(event, "LiquidatePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bSellUSDGIterator is returned from FilterSellUSDG and is used to iterate over the raw logs and unpacked data for SellUSDG events raised by the VaultV2b contract.
type VaultV2bSellUSDGIterator struct {
	Event *VaultV2bSellUSDG // Event containing the contract specifics and raw log

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
func (it *VaultV2bSellUSDGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bSellUSDG)
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
		it.Event = new(VaultV2bSellUSDG)
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
func (it *VaultV2bSellUSDGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bSellUSDGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bSellUSDG represents a SellUSDG event raised by the VaultV2b contract.
type VaultV2bSellUSDG struct {
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
func (_VaultV2b *VaultV2bFilterer) FilterSellUSDG(opts *bind.FilterOpts) (*VaultV2bSellUSDGIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "SellUSDG")
	if err != nil {
		return nil, err
	}
	return &VaultV2bSellUSDGIterator{contract: _VaultV2b.contract, event: "SellUSDG", logs: logs, sub: sub}, nil
}

// WatchSellUSDG is a free log subscription operation binding the contract event 0xd732b7828fa6cee72c285eac756fc66a7477e3dc22e22e7c432f1c265d40b483.
//
// Solidity: event SellUSDG(address account, address token, uint256 usdgAmount, uint256 tokenAmount, uint256 feeBasisPoints)
func (_VaultV2b *VaultV2bFilterer) WatchSellUSDG(opts *bind.WatchOpts, sink chan<- *VaultV2bSellUSDG) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "SellUSDG")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bSellUSDG)
				if err := _VaultV2b.contract.UnpackLog(event, "SellUSDG", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseSellUSDG(log types.Log) (*VaultV2bSellUSDG, error) {
	event := new(VaultV2bSellUSDG)
	if err := _VaultV2b.contract.UnpackLog(event, "SellUSDG", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the VaultV2b contract.
type VaultV2bSwapIterator struct {
	Event *VaultV2bSwap // Event containing the contract specifics and raw log

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
func (it *VaultV2bSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bSwap)
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
		it.Event = new(VaultV2bSwap)
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
func (it *VaultV2bSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bSwap represents a Swap event raised by the VaultV2b contract.
type VaultV2bSwap struct {
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
func (_VaultV2b *VaultV2bFilterer) FilterSwap(opts *bind.FilterOpts) (*VaultV2bSwapIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &VaultV2bSwapIterator{contract: _VaultV2b.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x0874b2d545cb271cdbda4e093020c452328b24af12382ed62c4d00f5c26709db.
//
// Solidity: event Swap(address account, address tokenIn, address tokenOut, uint256 amountIn, uint256 amountOut, uint256 amountOutAfterFees, uint256 feeBasisPoints)
func (_VaultV2b *VaultV2bFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *VaultV2bSwap) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bSwap)
				if err := _VaultV2b.contract.UnpackLog(event, "Swap", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseSwap(log types.Log) (*VaultV2bSwap, error) {
	event := new(VaultV2bSwap)
	if err := _VaultV2b.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bUpdateFundingRateIterator is returned from FilterUpdateFundingRate and is used to iterate over the raw logs and unpacked data for UpdateFundingRate events raised by the VaultV2b contract.
type VaultV2bUpdateFundingRateIterator struct {
	Event *VaultV2bUpdateFundingRate // Event containing the contract specifics and raw log

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
func (it *VaultV2bUpdateFundingRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bUpdateFundingRate)
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
		it.Event = new(VaultV2bUpdateFundingRate)
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
func (it *VaultV2bUpdateFundingRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bUpdateFundingRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bUpdateFundingRate represents a UpdateFundingRate event raised by the VaultV2b contract.
type VaultV2bUpdateFundingRate struct {
	Token       common.Address
	FundingRate *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateFundingRate is a free log retrieval operation binding the contract event 0xa146fc154e1913322e9817d49f0d5c37466c24326e15de10e739a948be815eab.
//
// Solidity: event UpdateFundingRate(address token, uint256 fundingRate)
func (_VaultV2b *VaultV2bFilterer) FilterUpdateFundingRate(opts *bind.FilterOpts) (*VaultV2bUpdateFundingRateIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "UpdateFundingRate")
	if err != nil {
		return nil, err
	}
	return &VaultV2bUpdateFundingRateIterator{contract: _VaultV2b.contract, event: "UpdateFundingRate", logs: logs, sub: sub}, nil
}

// WatchUpdateFundingRate is a free log subscription operation binding the contract event 0xa146fc154e1913322e9817d49f0d5c37466c24326e15de10e739a948be815eab.
//
// Solidity: event UpdateFundingRate(address token, uint256 fundingRate)
func (_VaultV2b *VaultV2bFilterer) WatchUpdateFundingRate(opts *bind.WatchOpts, sink chan<- *VaultV2bUpdateFundingRate) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "UpdateFundingRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bUpdateFundingRate)
				if err := _VaultV2b.contract.UnpackLog(event, "UpdateFundingRate", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseUpdateFundingRate(log types.Log) (*VaultV2bUpdateFundingRate, error) {
	event := new(VaultV2bUpdateFundingRate)
	if err := _VaultV2b.contract.UnpackLog(event, "UpdateFundingRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bUpdatePnlIterator is returned from FilterUpdatePnl and is used to iterate over the raw logs and unpacked data for UpdatePnl events raised by the VaultV2b contract.
type VaultV2bUpdatePnlIterator struct {
	Event *VaultV2bUpdatePnl // Event containing the contract specifics and raw log

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
func (it *VaultV2bUpdatePnlIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bUpdatePnl)
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
		it.Event = new(VaultV2bUpdatePnl)
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
func (it *VaultV2bUpdatePnlIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bUpdatePnlIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bUpdatePnl represents a UpdatePnl event raised by the VaultV2b contract.
type VaultV2bUpdatePnl struct {
	Key       [32]byte
	HasProfit bool
	Delta     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdatePnl is a free log retrieval operation binding the contract event 0x3ff41bdde87755b687ae83d0221a232b6be51a803330ed9661c1b5d0105e0d8a.
//
// Solidity: event UpdatePnl(bytes32 key, bool hasProfit, uint256 delta)
func (_VaultV2b *VaultV2bFilterer) FilterUpdatePnl(opts *bind.FilterOpts) (*VaultV2bUpdatePnlIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "UpdatePnl")
	if err != nil {
		return nil, err
	}
	return &VaultV2bUpdatePnlIterator{contract: _VaultV2b.contract, event: "UpdatePnl", logs: logs, sub: sub}, nil
}

// WatchUpdatePnl is a free log subscription operation binding the contract event 0x3ff41bdde87755b687ae83d0221a232b6be51a803330ed9661c1b5d0105e0d8a.
//
// Solidity: event UpdatePnl(bytes32 key, bool hasProfit, uint256 delta)
func (_VaultV2b *VaultV2bFilterer) WatchUpdatePnl(opts *bind.WatchOpts, sink chan<- *VaultV2bUpdatePnl) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "UpdatePnl")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bUpdatePnl)
				if err := _VaultV2b.contract.UnpackLog(event, "UpdatePnl", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseUpdatePnl(log types.Log) (*VaultV2bUpdatePnl, error) {
	event := new(VaultV2bUpdatePnl)
	if err := _VaultV2b.contract.UnpackLog(event, "UpdatePnl", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultV2bUpdatePositionIterator is returned from FilterUpdatePosition and is used to iterate over the raw logs and unpacked data for UpdatePosition events raised by the VaultV2b contract.
type VaultV2bUpdatePositionIterator struct {
	Event *VaultV2bUpdatePosition // Event containing the contract specifics and raw log

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
func (it *VaultV2bUpdatePositionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultV2bUpdatePosition)
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
		it.Event = new(VaultV2bUpdatePosition)
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
func (it *VaultV2bUpdatePositionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultV2bUpdatePositionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultV2bUpdatePosition represents a UpdatePosition event raised by the VaultV2b contract.
type VaultV2bUpdatePosition struct {
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
func (_VaultV2b *VaultV2bFilterer) FilterUpdatePosition(opts *bind.FilterOpts) (*VaultV2bUpdatePositionIterator, error) {

	logs, sub, err := _VaultV2b.contract.FilterLogs(opts, "UpdatePosition")
	if err != nil {
		return nil, err
	}
	return &VaultV2bUpdatePositionIterator{contract: _VaultV2b.contract, event: "UpdatePosition", logs: logs, sub: sub}, nil
}

// WatchUpdatePosition is a free log subscription operation binding the contract event 0x20853733b590dce729d9f4628682ebd9a34d2354e72679e66f43a008fc03b773.
//
// Solidity: event UpdatePosition(bytes32 key, uint256 size, uint256 collateral, uint256 averagePrice, uint256 entryFundingRate, uint256 reserveAmount, int256 realisedPnl, uint256 markPrice)
func (_VaultV2b *VaultV2bFilterer) WatchUpdatePosition(opts *bind.WatchOpts, sink chan<- *VaultV2bUpdatePosition) (event.Subscription, error) {

	logs, sub, err := _VaultV2b.contract.WatchLogs(opts, "UpdatePosition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultV2bUpdatePosition)
				if err := _VaultV2b.contract.UnpackLog(event, "UpdatePosition", log); err != nil {
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
func (_VaultV2b *VaultV2bFilterer) ParseUpdatePosition(log types.Log) (*VaultV2bUpdatePosition, error) {
	event := new(VaultV2bUpdatePosition)
	if err := _VaultV2b.contract.UnpackLog(event, "UpdatePosition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

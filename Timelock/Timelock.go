// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Timelock

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

// TimelockMetaData contains all meta data concerning the Timelock contract.
var TimelockMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_buffer\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokenManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_mintReceiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_glpManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_rewardRouter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_maxTokenSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_marginFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_maxMarginFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_BUFFER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_FUNDING_RATE_FACTOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_LEVERAGE_VALIDATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PRICE_PRECISION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchSetBonusRewards\",\"inputs\":[{\"name\":\"_vester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchWithdrawFees\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"buffer\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelAction\",\"inputs\":[{\"name\":\"_action\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disableLeverage\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"enableLeverage\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"glpManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"govSetCodeOwner\",\"inputs\":[{\"name\":\"_referralStorage\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_code\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_newAccount\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initGlpManager\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initRewardRouter\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isHandler\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isKeeper\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"marginFeeBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxMarginFeeBasisPoints\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxTokenSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mintReceiver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingActions\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"processMint\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemUsdg\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeAdmin\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardRouter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAdmin\",\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBuffer\",\"inputs\":[{\"name\":\"_buffer\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setContractHandler\",\"inputs\":[{\"name\":\"_handler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExternalAdmin\",\"inputs\":[{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_admin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFees\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_taxBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stableTaxBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_mintBurnFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_swapFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stableSwapFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_marginFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_liquidationFeeUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minProfitTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_hasDynamicFees\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFundingRate\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_fundingInterval\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_fundingRateFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stableFundingRateFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGlpCooldownDuration\",\"inputs\":[{\"name\":\"_cooldownDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGov\",\"inputs\":[{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setHandler\",\"inputs\":[{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_handler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInPrivateLiquidationMode\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_inPrivateLiquidationMode\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInPrivateTransferMode\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_inPrivateTransferMode\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIsLeverageEnabled\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isLeverageEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIsSwapEnabled\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isSwapEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setKeeper\",\"inputs\":[{\"name\":\"_keeper\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setLiquidator\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_liquidator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMarginFeeBasisPoints\",\"inputs\":[{\"name\":\"_marginFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_maxMarginFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxGasPrice\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_maxGasPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxGlobalShortSize\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxLeverage\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_maxLeverage\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPriceFeed\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_priceFeed\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setReferrerTier\",\"inputs\":[{\"name\":\"_referralStorage\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_referrer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tierId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setShortsTrackerAveragePriceWeight\",\"inputs\":[{\"name\":\"_shortsTrackerAveragePriceWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setShouldToggleIsLeverageEnabled\",\"inputs\":[{\"name\":\"_shouldToggleIsLeverageEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSwapFees\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_taxBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stableTaxBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_mintBurnFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_swapFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stableSwapFeeBasisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTier\",\"inputs\":[{\"name\":\"_referralStorage\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tierId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_totalRebate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_discountShare\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenConfig\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minProfitBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_maxUsdgAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_bufferAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_usdgAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUsdgAmounts\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_usdgAmounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setVaultUtils\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_vaultUtils\",\"type\":\"address\",\"internalType\":\"contractIVaultUtils\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"shouldToggleIsLeverageEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"signalApprove\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signalMint\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signalRedeemUsdg\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signalSetGov\",\"inputs\":[{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_gov\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signalSetHandler\",\"inputs\":[{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_handler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signalSetPriceFeed\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_priceFeed\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signalVaultSetTokenConfig\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenDecimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokenWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minProfitBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_maxUsdgAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isStable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_isShortable\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signalWithdrawToken\",\"inputs\":[{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferIn\",\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateUsdgSupply\",\"inputs\":[{\"name\":\"usdgAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"vaultSetTokenConfig\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_tokenDecimals\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_tokenWeight\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_minProfitBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_maxUsdgAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_isStable\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_isShortable\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawFees\",\"inputs\":[{\"name\":\"_vault\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawToken\",\"inputs\":[{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ClearAction\",\"inputs\":[{\"name\":\"action\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignalApprove\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignalMint\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignalPendingAction\",\"inputs\":[{\"name\":\"action\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignalRedeemUsdg\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignalSetGov\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"gov\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"action\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignalSetHandler\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"handler\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"isActive\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"action\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignalSetPriceFeed\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"priceFeed\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"action\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignalVaultSetTokenConfig\",\"inputs\":[{\"name\":\"vault\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenDecimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"tokenWeight\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"minProfitBps\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"maxUsdgAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isStable\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"isShortable\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignalWithdrawToken\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"action\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false}]",
}

// TimelockABI is the input ABI used to generate the binding from.
// Deprecated: Use TimelockMetaData.ABI instead.
var TimelockABI = TimelockMetaData.ABI

// Timelock is an auto generated Go binding around an Ethereum contract.
type Timelock struct {
	TimelockCaller     // Read-only binding to the contract
	TimelockTransactor // Write-only binding to the contract
	TimelockFilterer   // Log filterer for contract events
}

// TimelockCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimelockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimelockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimelockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimelockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimelockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimelockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimelockSession struct {
	Contract     *Timelock         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimelockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimelockCallerSession struct {
	Contract *TimelockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TimelockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimelockTransactorSession struct {
	Contract     *TimelockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TimelockRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimelockRaw struct {
	Contract *Timelock // Generic contract binding to access the raw methods on
}

// TimelockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimelockCallerRaw struct {
	Contract *TimelockCaller // Generic read-only contract binding to access the raw methods on
}

// TimelockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimelockTransactorRaw struct {
	Contract *TimelockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimelock creates a new instance of Timelock, bound to a specific deployed contract.
func NewTimelock(address common.Address, backend bind.ContractBackend) (*Timelock, error) {
	contract, err := bindTimelock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Timelock{TimelockCaller: TimelockCaller{contract: contract}, TimelockTransactor: TimelockTransactor{contract: contract}, TimelockFilterer: TimelockFilterer{contract: contract}}, nil
}

// NewTimelockCaller creates a new read-only instance of Timelock, bound to a specific deployed contract.
func NewTimelockCaller(address common.Address, caller bind.ContractCaller) (*TimelockCaller, error) {
	contract, err := bindTimelock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimelockCaller{contract: contract}, nil
}

// NewTimelockTransactor creates a new write-only instance of Timelock, bound to a specific deployed contract.
func NewTimelockTransactor(address common.Address, transactor bind.ContractTransactor) (*TimelockTransactor, error) {
	contract, err := bindTimelock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimelockTransactor{contract: contract}, nil
}

// NewTimelockFilterer creates a new log filterer instance of Timelock, bound to a specific deployed contract.
func NewTimelockFilterer(address common.Address, filterer bind.ContractFilterer) (*TimelockFilterer, error) {
	contract, err := bindTimelock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimelockFilterer{contract: contract}, nil
}

// bindTimelock binds a generic wrapper to an already deployed contract.
func bindTimelock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TimelockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Timelock *TimelockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Timelock.Contract.TimelockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Timelock *TimelockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Timelock.Contract.TimelockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Timelock *TimelockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Timelock.Contract.TimelockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Timelock *TimelockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Timelock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Timelock *TimelockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Timelock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Timelock *TimelockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Timelock.Contract.contract.Transact(opts, method, params...)
}

// MAXBUFFER is a free data retrieval call binding the contract method 0x61d07569.
//
// Solidity: function MAX_BUFFER() view returns(uint256)
func (_Timelock *TimelockCaller) MAXBUFFER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Timelock.contract.Call(opts, &out, "MAX_BUFFER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBUFFER is a free data retrieval call binding the contract method 0x61d07569.
//
// Solidity: function MAX_BUFFER() view returns(uint256)
func (_Timelock *TimelockSession) MAXBUFFER() (*big.Int, error) {
	return _Timelock.Contract.MAXBUFFER(&_Timelock.CallOpts)
}

// MAXBUFFER is a free data retrieval call binding the contract method 0x61d07569.
//
// Solidity: function MAX_BUFFER() view returns(uint256)
func (_Timelock *TimelockCallerSession) MAXBUFFER() (*big.Int, error) {
	return _Timelock.Contract.MAXBUFFER(&_Timelock.CallOpts)
}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_Timelock *TimelockCaller) MAXFUNDINGRATEFACTOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Timelock.contract.Call(opts, &out, "MAX_FUNDING_RATE_FACTOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_Timelock *TimelockSession) MAXFUNDINGRATEFACTOR() (*big.Int, error) {
	return _Timelock.Contract.MAXFUNDINGRATEFACTOR(&_Timelock.CallOpts)
}

// MAXFUNDINGRATEFACTOR is a free data retrieval call binding the contract method 0x8a39735a.
//
// Solidity: function MAX_FUNDING_RATE_FACTOR() view returns(uint256)
func (_Timelock *TimelockCallerSession) MAXFUNDINGRATEFACTOR() (*big.Int, error) {
	return _Timelock.Contract.MAXFUNDINGRATEFACTOR(&_Timelock.CallOpts)
}

// MAXLEVERAGEVALIDATION is a free data retrieval call binding the contract method 0x2ba3725a.
//
// Solidity: function MAX_LEVERAGE_VALIDATION() view returns(uint256)
func (_Timelock *TimelockCaller) MAXLEVERAGEVALIDATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Timelock.contract.Call(opts, &out, "MAX_LEVERAGE_VALIDATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXLEVERAGEVALIDATION is a free data retrieval call binding the contract method 0x2ba3725a.
//
// Solidity: function MAX_LEVERAGE_VALIDATION() view returns(uint256)
func (_Timelock *TimelockSession) MAXLEVERAGEVALIDATION() (*big.Int, error) {
	return _Timelock.Contract.MAXLEVERAGEVALIDATION(&_Timelock.CallOpts)
}

// MAXLEVERAGEVALIDATION is a free data retrieval call binding the contract method 0x2ba3725a.
//
// Solidity: function MAX_LEVERAGE_VALIDATION() view returns(uint256)
func (_Timelock *TimelockCallerSession) MAXLEVERAGEVALIDATION() (*big.Int, error) {
	return _Timelock.Contract.MAXLEVERAGEVALIDATION(&_Timelock.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_Timelock *TimelockCaller) PRICEPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Timelock.contract.Call(opts, &out, "PRICE_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_Timelock *TimelockSession) PRICEPRECISION() (*big.Int, error) {
	return _Timelock.Contract.PRICEPRECISION(&_Timelock.CallOpts)
}

// PRICEPRECISION is a free data retrieval call binding the contract method 0x95082d25.
//
// Solidity: function PRICE_PRECISION() view returns(uint256)
func (_Timelock *TimelockCallerSession) PRICEPRECISION() (*big.Int, error) {
	return _Timelock.Contract.PRICEPRECISION(&_Timelock.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Timelock *TimelockCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Timelock.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Timelock *TimelockSession) Admin() (common.Address, error) {
	return _Timelock.Contract.Admin(&_Timelock.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Timelock *TimelockCallerSession) Admin() (common.Address, error) {
	return _Timelock.Contract.Admin(&_Timelock.CallOpts)
}

// Buffer is a free data retrieval call binding the contract method 0xedaafe20.
//
// Solidity: function buffer() view returns(uint256)
func (_Timelock *TimelockCaller) Buffer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Timelock.contract.Call(opts, &out, "buffer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Buffer is a free data retrieval call binding the contract method 0xedaafe20.
//
// Solidity: function buffer() view returns(uint256)
func (_Timelock *TimelockSession) Buffer() (*big.Int, error) {
	return _Timelock.Contract.Buffer(&_Timelock.CallOpts)
}

// Buffer is a free data retrieval call binding the contract method 0xedaafe20.
//
// Solidity: function buffer() view returns(uint256)
func (_Timelock *TimelockCallerSession) Buffer() (*big.Int, error) {
	return _Timelock.Contract.Buffer(&_Timelock.CallOpts)
}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_Timelock *TimelockCaller) GlpManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Timelock.contract.Call(opts, &out, "glpManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_Timelock *TimelockSession) GlpManager() (common.Address, error) {
	return _Timelock.Contract.GlpManager(&_Timelock.CallOpts)
}

// GlpManager is a free data retrieval call binding the contract method 0xfa6db1bc.
//
// Solidity: function glpManager() view returns(address)
func (_Timelock *TimelockCallerSession) GlpManager() (common.Address, error) {
	return _Timelock.Contract.GlpManager(&_Timelock.CallOpts)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_Timelock *TimelockCaller) IsHandler(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Timelock.contract.Call(opts, &out, "isHandler", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_Timelock *TimelockSession) IsHandler(arg0 common.Address) (bool, error) {
	return _Timelock.Contract.IsHandler(&_Timelock.CallOpts, arg0)
}

// IsHandler is a free data retrieval call binding the contract method 0x46ea87af.
//
// Solidity: function isHandler(address ) view returns(bool)
func (_Timelock *TimelockCallerSession) IsHandler(arg0 common.Address) (bool, error) {
	return _Timelock.Contract.IsHandler(&_Timelock.CallOpts, arg0)
}

// IsKeeper is a free data retrieval call binding the contract method 0x6ba42aaa.
//
// Solidity: function isKeeper(address ) view returns(bool)
func (_Timelock *TimelockCaller) IsKeeper(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Timelock.contract.Call(opts, &out, "isKeeper", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKeeper is a free data retrieval call binding the contract method 0x6ba42aaa.
//
// Solidity: function isKeeper(address ) view returns(bool)
func (_Timelock *TimelockSession) IsKeeper(arg0 common.Address) (bool, error) {
	return _Timelock.Contract.IsKeeper(&_Timelock.CallOpts, arg0)
}

// IsKeeper is a free data retrieval call binding the contract method 0x6ba42aaa.
//
// Solidity: function isKeeper(address ) view returns(bool)
func (_Timelock *TimelockCallerSession) IsKeeper(arg0 common.Address) (bool, error) {
	return _Timelock.Contract.IsKeeper(&_Timelock.CallOpts, arg0)
}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_Timelock *TimelockCaller) MarginFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Timelock.contract.Call(opts, &out, "marginFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_Timelock *TimelockSession) MarginFeeBasisPoints() (*big.Int, error) {
	return _Timelock.Contract.MarginFeeBasisPoints(&_Timelock.CallOpts)
}

// MarginFeeBasisPoints is a free data retrieval call binding the contract method 0x318bc689.
//
// Solidity: function marginFeeBasisPoints() view returns(uint256)
func (_Timelock *TimelockCallerSession) MarginFeeBasisPoints() (*big.Int, error) {
	return _Timelock.Contract.MarginFeeBasisPoints(&_Timelock.CallOpts)
}

// MaxMarginFeeBasisPoints is a free data retrieval call binding the contract method 0x996e2bc9.
//
// Solidity: function maxMarginFeeBasisPoints() view returns(uint256)
func (_Timelock *TimelockCaller) MaxMarginFeeBasisPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Timelock.contract.Call(opts, &out, "maxMarginFeeBasisPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMarginFeeBasisPoints is a free data retrieval call binding the contract method 0x996e2bc9.
//
// Solidity: function maxMarginFeeBasisPoints() view returns(uint256)
func (_Timelock *TimelockSession) MaxMarginFeeBasisPoints() (*big.Int, error) {
	return _Timelock.Contract.MaxMarginFeeBasisPoints(&_Timelock.CallOpts)
}

// MaxMarginFeeBasisPoints is a free data retrieval call binding the contract method 0x996e2bc9.
//
// Solidity: function maxMarginFeeBasisPoints() view returns(uint256)
func (_Timelock *TimelockCallerSession) MaxMarginFeeBasisPoints() (*big.Int, error) {
	return _Timelock.Contract.MaxMarginFeeBasisPoints(&_Timelock.CallOpts)
}

// MaxTokenSupply is a free data retrieval call binding the contract method 0x50f7c204.
//
// Solidity: function maxTokenSupply() view returns(uint256)
func (_Timelock *TimelockCaller) MaxTokenSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Timelock.contract.Call(opts, &out, "maxTokenSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTokenSupply is a free data retrieval call binding the contract method 0x50f7c204.
//
// Solidity: function maxTokenSupply() view returns(uint256)
func (_Timelock *TimelockSession) MaxTokenSupply() (*big.Int, error) {
	return _Timelock.Contract.MaxTokenSupply(&_Timelock.CallOpts)
}

// MaxTokenSupply is a free data retrieval call binding the contract method 0x50f7c204.
//
// Solidity: function maxTokenSupply() view returns(uint256)
func (_Timelock *TimelockCallerSession) MaxTokenSupply() (*big.Int, error) {
	return _Timelock.Contract.MaxTokenSupply(&_Timelock.CallOpts)
}

// MintReceiver is a free data retrieval call binding the contract method 0xc7bb26a0.
//
// Solidity: function mintReceiver() view returns(address)
func (_Timelock *TimelockCaller) MintReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Timelock.contract.Call(opts, &out, "mintReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MintReceiver is a free data retrieval call binding the contract method 0xc7bb26a0.
//
// Solidity: function mintReceiver() view returns(address)
func (_Timelock *TimelockSession) MintReceiver() (common.Address, error) {
	return _Timelock.Contract.MintReceiver(&_Timelock.CallOpts)
}

// MintReceiver is a free data retrieval call binding the contract method 0xc7bb26a0.
//
// Solidity: function mintReceiver() view returns(address)
func (_Timelock *TimelockCallerSession) MintReceiver() (common.Address, error) {
	return _Timelock.Contract.MintReceiver(&_Timelock.CallOpts)
}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(uint256)
func (_Timelock *TimelockCaller) PendingActions(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Timelock.contract.Call(opts, &out, "pendingActions", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(uint256)
func (_Timelock *TimelockSession) PendingActions(arg0 [32]byte) (*big.Int, error) {
	return _Timelock.Contract.PendingActions(&_Timelock.CallOpts, arg0)
}

// PendingActions is a free data retrieval call binding the contract method 0xe30569e5.
//
// Solidity: function pendingActions(bytes32 ) view returns(uint256)
func (_Timelock *TimelockCallerSession) PendingActions(arg0 [32]byte) (*big.Int, error) {
	return _Timelock.Contract.PendingActions(&_Timelock.CallOpts, arg0)
}

// RewardRouter is a free data retrieval call binding the contract method 0x5a3bb989.
//
// Solidity: function rewardRouter() view returns(address)
func (_Timelock *TimelockCaller) RewardRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Timelock.contract.Call(opts, &out, "rewardRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardRouter is a free data retrieval call binding the contract method 0x5a3bb989.
//
// Solidity: function rewardRouter() view returns(address)
func (_Timelock *TimelockSession) RewardRouter() (common.Address, error) {
	return _Timelock.Contract.RewardRouter(&_Timelock.CallOpts)
}

// RewardRouter is a free data retrieval call binding the contract method 0x5a3bb989.
//
// Solidity: function rewardRouter() view returns(address)
func (_Timelock *TimelockCallerSession) RewardRouter() (common.Address, error) {
	return _Timelock.Contract.RewardRouter(&_Timelock.CallOpts)
}

// ShouldToggleIsLeverageEnabled is a free data retrieval call binding the contract method 0x23aaad12.
//
// Solidity: function shouldToggleIsLeverageEnabled() view returns(bool)
func (_Timelock *TimelockCaller) ShouldToggleIsLeverageEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Timelock.contract.Call(opts, &out, "shouldToggleIsLeverageEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ShouldToggleIsLeverageEnabled is a free data retrieval call binding the contract method 0x23aaad12.
//
// Solidity: function shouldToggleIsLeverageEnabled() view returns(bool)
func (_Timelock *TimelockSession) ShouldToggleIsLeverageEnabled() (bool, error) {
	return _Timelock.Contract.ShouldToggleIsLeverageEnabled(&_Timelock.CallOpts)
}

// ShouldToggleIsLeverageEnabled is a free data retrieval call binding the contract method 0x23aaad12.
//
// Solidity: function shouldToggleIsLeverageEnabled() view returns(bool)
func (_Timelock *TimelockCallerSession) ShouldToggleIsLeverageEnabled() (bool, error) {
	return _Timelock.Contract.ShouldToggleIsLeverageEnabled(&_Timelock.CallOpts)
}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_Timelock *TimelockCaller) TokenManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Timelock.contract.Call(opts, &out, "tokenManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_Timelock *TimelockSession) TokenManager() (common.Address, error) {
	return _Timelock.Contract.TokenManager(&_Timelock.CallOpts)
}

// TokenManager is a free data retrieval call binding the contract method 0x2a709b14.
//
// Solidity: function tokenManager() view returns(address)
func (_Timelock *TimelockCallerSession) TokenManager() (common.Address, error) {
	return _Timelock.Contract.TokenManager(&_Timelock.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount) returns()
func (_Timelock *TimelockTransactor) Approve(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "approve", _token, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount) returns()
func (_Timelock *TimelockSession) Approve(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.Approve(&_Timelock.TransactOpts, _token, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address _token, address _spender, uint256 _amount) returns()
func (_Timelock *TimelockTransactorSession) Approve(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.Approve(&_Timelock.TransactOpts, _token, _spender, _amount)
}

// BatchSetBonusRewards is a paid mutator transaction binding the contract method 0x009a698e.
//
// Solidity: function batchSetBonusRewards(address _vester, address[] _accounts, uint256[] _amounts) returns()
func (_Timelock *TimelockTransactor) BatchSetBonusRewards(opts *bind.TransactOpts, _vester common.Address, _accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "batchSetBonusRewards", _vester, _accounts, _amounts)
}

// BatchSetBonusRewards is a paid mutator transaction binding the contract method 0x009a698e.
//
// Solidity: function batchSetBonusRewards(address _vester, address[] _accounts, uint256[] _amounts) returns()
func (_Timelock *TimelockSession) BatchSetBonusRewards(_vester common.Address, _accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.BatchSetBonusRewards(&_Timelock.TransactOpts, _vester, _accounts, _amounts)
}

// BatchSetBonusRewards is a paid mutator transaction binding the contract method 0x009a698e.
//
// Solidity: function batchSetBonusRewards(address _vester, address[] _accounts, uint256[] _amounts) returns()
func (_Timelock *TimelockTransactorSession) BatchSetBonusRewards(_vester common.Address, _accounts []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.BatchSetBonusRewards(&_Timelock.TransactOpts, _vester, _accounts, _amounts)
}

// BatchWithdrawFees is a paid mutator transaction binding the contract method 0xe8ae2271.
//
// Solidity: function batchWithdrawFees(address _vault, address[] _tokens) returns()
func (_Timelock *TimelockTransactor) BatchWithdrawFees(opts *bind.TransactOpts, _vault common.Address, _tokens []common.Address) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "batchWithdrawFees", _vault, _tokens)
}

// BatchWithdrawFees is a paid mutator transaction binding the contract method 0xe8ae2271.
//
// Solidity: function batchWithdrawFees(address _vault, address[] _tokens) returns()
func (_Timelock *TimelockSession) BatchWithdrawFees(_vault common.Address, _tokens []common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.BatchWithdrawFees(&_Timelock.TransactOpts, _vault, _tokens)
}

// BatchWithdrawFees is a paid mutator transaction binding the contract method 0xe8ae2271.
//
// Solidity: function batchWithdrawFees(address _vault, address[] _tokens) returns()
func (_Timelock *TimelockTransactorSession) BatchWithdrawFees(_vault common.Address, _tokens []common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.BatchWithdrawFees(&_Timelock.TransactOpts, _vault, _tokens)
}

// CancelAction is a paid mutator transaction binding the contract method 0x781cc3d3.
//
// Solidity: function cancelAction(bytes32 _action) returns()
func (_Timelock *TimelockTransactor) CancelAction(opts *bind.TransactOpts, _action [32]byte) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "cancelAction", _action)
}

// CancelAction is a paid mutator transaction binding the contract method 0x781cc3d3.
//
// Solidity: function cancelAction(bytes32 _action) returns()
func (_Timelock *TimelockSession) CancelAction(_action [32]byte) (*types.Transaction, error) {
	return _Timelock.Contract.CancelAction(&_Timelock.TransactOpts, _action)
}

// CancelAction is a paid mutator transaction binding the contract method 0x781cc3d3.
//
// Solidity: function cancelAction(bytes32 _action) returns()
func (_Timelock *TimelockTransactorSession) CancelAction(_action [32]byte) (*types.Transaction, error) {
	return _Timelock.Contract.CancelAction(&_Timelock.TransactOpts, _action)
}

// DisableLeverage is a paid mutator transaction binding the contract method 0xd3c87bbb.
//
// Solidity: function disableLeverage(address _vault) returns()
func (_Timelock *TimelockTransactor) DisableLeverage(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "disableLeverage", _vault)
}

// DisableLeverage is a paid mutator transaction binding the contract method 0xd3c87bbb.
//
// Solidity: function disableLeverage(address _vault) returns()
func (_Timelock *TimelockSession) DisableLeverage(_vault common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.DisableLeverage(&_Timelock.TransactOpts, _vault)
}

// DisableLeverage is a paid mutator transaction binding the contract method 0xd3c87bbb.
//
// Solidity: function disableLeverage(address _vault) returns()
func (_Timelock *TimelockTransactorSession) DisableLeverage(_vault common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.DisableLeverage(&_Timelock.TransactOpts, _vault)
}

// EnableLeverage is a paid mutator transaction binding the contract method 0x6d63c1d0.
//
// Solidity: function enableLeverage(address _vault) returns()
func (_Timelock *TimelockTransactor) EnableLeverage(opts *bind.TransactOpts, _vault common.Address) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "enableLeverage", _vault)
}

// EnableLeverage is a paid mutator transaction binding the contract method 0x6d63c1d0.
//
// Solidity: function enableLeverage(address _vault) returns()
func (_Timelock *TimelockSession) EnableLeverage(_vault common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.EnableLeverage(&_Timelock.TransactOpts, _vault)
}

// EnableLeverage is a paid mutator transaction binding the contract method 0x6d63c1d0.
//
// Solidity: function enableLeverage(address _vault) returns()
func (_Timelock *TimelockTransactorSession) EnableLeverage(_vault common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.EnableLeverage(&_Timelock.TransactOpts, _vault)
}

// GovSetCodeOwner is a paid mutator transaction binding the contract method 0x204bbc54.
//
// Solidity: function govSetCodeOwner(address _referralStorage, bytes32 _code, address _newAccount) returns()
func (_Timelock *TimelockTransactor) GovSetCodeOwner(opts *bind.TransactOpts, _referralStorage common.Address, _code [32]byte, _newAccount common.Address) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "govSetCodeOwner", _referralStorage, _code, _newAccount)
}

// GovSetCodeOwner is a paid mutator transaction binding the contract method 0x204bbc54.
//
// Solidity: function govSetCodeOwner(address _referralStorage, bytes32 _code, address _newAccount) returns()
func (_Timelock *TimelockSession) GovSetCodeOwner(_referralStorage common.Address, _code [32]byte, _newAccount common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.GovSetCodeOwner(&_Timelock.TransactOpts, _referralStorage, _code, _newAccount)
}

// GovSetCodeOwner is a paid mutator transaction binding the contract method 0x204bbc54.
//
// Solidity: function govSetCodeOwner(address _referralStorage, bytes32 _code, address _newAccount) returns()
func (_Timelock *TimelockTransactorSession) GovSetCodeOwner(_referralStorage common.Address, _code [32]byte, _newAccount common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.GovSetCodeOwner(&_Timelock.TransactOpts, _referralStorage, _code, _newAccount)
}

// InitGlpManager is a paid mutator transaction binding the contract method 0x7e43c62d.
//
// Solidity: function initGlpManager() returns()
func (_Timelock *TimelockTransactor) InitGlpManager(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "initGlpManager")
}

// InitGlpManager is a paid mutator transaction binding the contract method 0x7e43c62d.
//
// Solidity: function initGlpManager() returns()
func (_Timelock *TimelockSession) InitGlpManager() (*types.Transaction, error) {
	return _Timelock.Contract.InitGlpManager(&_Timelock.TransactOpts)
}

// InitGlpManager is a paid mutator transaction binding the contract method 0x7e43c62d.
//
// Solidity: function initGlpManager() returns()
func (_Timelock *TimelockTransactorSession) InitGlpManager() (*types.Transaction, error) {
	return _Timelock.Contract.InitGlpManager(&_Timelock.TransactOpts)
}

// InitRewardRouter is a paid mutator transaction binding the contract method 0xafd14deb.
//
// Solidity: function initRewardRouter() returns()
func (_Timelock *TimelockTransactor) InitRewardRouter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "initRewardRouter")
}

// InitRewardRouter is a paid mutator transaction binding the contract method 0xafd14deb.
//
// Solidity: function initRewardRouter() returns()
func (_Timelock *TimelockSession) InitRewardRouter() (*types.Transaction, error) {
	return _Timelock.Contract.InitRewardRouter(&_Timelock.TransactOpts)
}

// InitRewardRouter is a paid mutator transaction binding the contract method 0xafd14deb.
//
// Solidity: function initRewardRouter() returns()
func (_Timelock *TimelockTransactorSession) InitRewardRouter() (*types.Transaction, error) {
	return _Timelock.Contract.InitRewardRouter(&_Timelock.TransactOpts)
}

// ProcessMint is a paid mutator transaction binding the contract method 0xbc8a8ab9.
//
// Solidity: function processMint(address _token, address _receiver, uint256 _amount) returns()
func (_Timelock *TimelockTransactor) ProcessMint(opts *bind.TransactOpts, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "processMint", _token, _receiver, _amount)
}

// ProcessMint is a paid mutator transaction binding the contract method 0xbc8a8ab9.
//
// Solidity: function processMint(address _token, address _receiver, uint256 _amount) returns()
func (_Timelock *TimelockSession) ProcessMint(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.ProcessMint(&_Timelock.TransactOpts, _token, _receiver, _amount)
}

// ProcessMint is a paid mutator transaction binding the contract method 0xbc8a8ab9.
//
// Solidity: function processMint(address _token, address _receiver, uint256 _amount) returns()
func (_Timelock *TimelockTransactorSession) ProcessMint(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.ProcessMint(&_Timelock.TransactOpts, _token, _receiver, _amount)
}

// RedeemUsdg is a paid mutator transaction binding the contract method 0x70ac0a93.
//
// Solidity: function redeemUsdg(address _vault, address _token, uint256 _amount) returns()
func (_Timelock *TimelockTransactor) RedeemUsdg(opts *bind.TransactOpts, _vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "redeemUsdg", _vault, _token, _amount)
}

// RedeemUsdg is a paid mutator transaction binding the contract method 0x70ac0a93.
//
// Solidity: function redeemUsdg(address _vault, address _token, uint256 _amount) returns()
func (_Timelock *TimelockSession) RedeemUsdg(_vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.RedeemUsdg(&_Timelock.TransactOpts, _vault, _token, _amount)
}

// RedeemUsdg is a paid mutator transaction binding the contract method 0x70ac0a93.
//
// Solidity: function redeemUsdg(address _vault, address _token, uint256 _amount) returns()
func (_Timelock *TimelockTransactorSession) RedeemUsdg(_vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.RedeemUsdg(&_Timelock.TransactOpts, _vault, _token, _amount)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(address _token, address _account) returns()
func (_Timelock *TimelockTransactor) RemoveAdmin(opts *bind.TransactOpts, _token common.Address, _account common.Address) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "removeAdmin", _token, _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(address _token, address _account) returns()
func (_Timelock *TimelockSession) RemoveAdmin(_token common.Address, _account common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.RemoveAdmin(&_Timelock.TransactOpts, _token, _account)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x268959e5.
//
// Solidity: function removeAdmin(address _token, address _account) returns()
func (_Timelock *TimelockTransactorSession) RemoveAdmin(_token common.Address, _account common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.RemoveAdmin(&_Timelock.TransactOpts, _token, _account)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_Timelock *TimelockTransactor) SetAdmin(opts *bind.TransactOpts, _admin common.Address) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setAdmin", _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_Timelock *TimelockSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.SetAdmin(&_Timelock.TransactOpts, _admin)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address _admin) returns()
func (_Timelock *TimelockTransactorSession) SetAdmin(_admin common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.SetAdmin(&_Timelock.TransactOpts, _admin)
}

// SetBuffer is a paid mutator transaction binding the contract method 0xadc7ea37.
//
// Solidity: function setBuffer(uint256 _buffer) returns()
func (_Timelock *TimelockTransactor) SetBuffer(opts *bind.TransactOpts, _buffer *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setBuffer", _buffer)
}

// SetBuffer is a paid mutator transaction binding the contract method 0xadc7ea37.
//
// Solidity: function setBuffer(uint256 _buffer) returns()
func (_Timelock *TimelockSession) SetBuffer(_buffer *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetBuffer(&_Timelock.TransactOpts, _buffer)
}

// SetBuffer is a paid mutator transaction binding the contract method 0xadc7ea37.
//
// Solidity: function setBuffer(uint256 _buffer) returns()
func (_Timelock *TimelockTransactorSession) SetBuffer(_buffer *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetBuffer(&_Timelock.TransactOpts, _buffer)
}

// SetContractHandler is a paid mutator transaction binding the contract method 0x185051c1.
//
// Solidity: function setContractHandler(address _handler, bool _isActive) returns()
func (_Timelock *TimelockTransactor) SetContractHandler(opts *bind.TransactOpts, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setContractHandler", _handler, _isActive)
}

// SetContractHandler is a paid mutator transaction binding the contract method 0x185051c1.
//
// Solidity: function setContractHandler(address _handler, bool _isActive) returns()
func (_Timelock *TimelockSession) SetContractHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetContractHandler(&_Timelock.TransactOpts, _handler, _isActive)
}

// SetContractHandler is a paid mutator transaction binding the contract method 0x185051c1.
//
// Solidity: function setContractHandler(address _handler, bool _isActive) returns()
func (_Timelock *TimelockTransactorSession) SetContractHandler(_handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetContractHandler(&_Timelock.TransactOpts, _handler, _isActive)
}

// SetExternalAdmin is a paid mutator transaction binding the contract method 0x55ef1395.
//
// Solidity: function setExternalAdmin(address _target, address _admin) returns()
func (_Timelock *TimelockTransactor) SetExternalAdmin(opts *bind.TransactOpts, _target common.Address, _admin common.Address) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setExternalAdmin", _target, _admin)
}

// SetExternalAdmin is a paid mutator transaction binding the contract method 0x55ef1395.
//
// Solidity: function setExternalAdmin(address _target, address _admin) returns()
func (_Timelock *TimelockSession) SetExternalAdmin(_target common.Address, _admin common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.SetExternalAdmin(&_Timelock.TransactOpts, _target, _admin)
}

// SetExternalAdmin is a paid mutator transaction binding the contract method 0x55ef1395.
//
// Solidity: function setExternalAdmin(address _target, address _admin) returns()
func (_Timelock *TimelockTransactorSession) SetExternalAdmin(_target common.Address, _admin common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.SetExternalAdmin(&_Timelock.TransactOpts, _target, _admin)
}

// SetFees is a paid mutator transaction binding the contract method 0x6e5227d4.
//
// Solidity: function setFees(address _vault, uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_Timelock *TimelockTransactor) SetFees(opts *bind.TransactOpts, _vault common.Address, _taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setFees", _vault, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFees is a paid mutator transaction binding the contract method 0x6e5227d4.
//
// Solidity: function setFees(address _vault, uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_Timelock *TimelockSession) SetFees(_vault common.Address, _taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetFees(&_Timelock.TransactOpts, _vault, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFees is a paid mutator transaction binding the contract method 0x6e5227d4.
//
// Solidity: function setFees(address _vault, uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints, uint256 _marginFeeBasisPoints, uint256 _liquidationFeeUsd, uint256 _minProfitTime, bool _hasDynamicFees) returns()
func (_Timelock *TimelockTransactorSession) SetFees(_vault common.Address, _taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int, _marginFeeBasisPoints *big.Int, _liquidationFeeUsd *big.Int, _minProfitTime *big.Int, _hasDynamicFees bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetFees(&_Timelock.TransactOpts, _vault, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints, _marginFeeBasisPoints, _liquidationFeeUsd, _minProfitTime, _hasDynamicFees)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x227f03eb.
//
// Solidity: function setFundingRate(address _vault, uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_Timelock *TimelockTransactor) SetFundingRate(opts *bind.TransactOpts, _vault common.Address, _fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setFundingRate", _vault, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x227f03eb.
//
// Solidity: function setFundingRate(address _vault, uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_Timelock *TimelockSession) SetFundingRate(_vault common.Address, _fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetFundingRate(&_Timelock.TransactOpts, _vault, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetFundingRate is a paid mutator transaction binding the contract method 0x227f03eb.
//
// Solidity: function setFundingRate(address _vault, uint256 _fundingInterval, uint256 _fundingRateFactor, uint256 _stableFundingRateFactor) returns()
func (_Timelock *TimelockTransactorSession) SetFundingRate(_vault common.Address, _fundingInterval *big.Int, _fundingRateFactor *big.Int, _stableFundingRateFactor *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetFundingRate(&_Timelock.TransactOpts, _vault, _fundingInterval, _fundingRateFactor, _stableFundingRateFactor)
}

// SetGlpCooldownDuration is a paid mutator transaction binding the contract method 0x5b05348b.
//
// Solidity: function setGlpCooldownDuration(uint256 _cooldownDuration) returns()
func (_Timelock *TimelockTransactor) SetGlpCooldownDuration(opts *bind.TransactOpts, _cooldownDuration *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setGlpCooldownDuration", _cooldownDuration)
}

// SetGlpCooldownDuration is a paid mutator transaction binding the contract method 0x5b05348b.
//
// Solidity: function setGlpCooldownDuration(uint256 _cooldownDuration) returns()
func (_Timelock *TimelockSession) SetGlpCooldownDuration(_cooldownDuration *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetGlpCooldownDuration(&_Timelock.TransactOpts, _cooldownDuration)
}

// SetGlpCooldownDuration is a paid mutator transaction binding the contract method 0x5b05348b.
//
// Solidity: function setGlpCooldownDuration(uint256 _cooldownDuration) returns()
func (_Timelock *TimelockTransactorSession) SetGlpCooldownDuration(_cooldownDuration *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetGlpCooldownDuration(&_Timelock.TransactOpts, _cooldownDuration)
}

// SetGov is a paid mutator transaction binding the contract method 0x51a6de0d.
//
// Solidity: function setGov(address _target, address _gov) returns()
func (_Timelock *TimelockTransactor) SetGov(opts *bind.TransactOpts, _target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setGov", _target, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0x51a6de0d.
//
// Solidity: function setGov(address _target, address _gov) returns()
func (_Timelock *TimelockSession) SetGov(_target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.SetGov(&_Timelock.TransactOpts, _target, _gov)
}

// SetGov is a paid mutator transaction binding the contract method 0x51a6de0d.
//
// Solidity: function setGov(address _target, address _gov) returns()
func (_Timelock *TimelockTransactorSession) SetGov(_target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.SetGov(&_Timelock.TransactOpts, _target, _gov)
}

// SetHandler is a paid mutator transaction binding the contract method 0x1154e808.
//
// Solidity: function setHandler(address _target, address _handler, bool _isActive) returns()
func (_Timelock *TimelockTransactor) SetHandler(opts *bind.TransactOpts, _target common.Address, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setHandler", _target, _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x1154e808.
//
// Solidity: function setHandler(address _target, address _handler, bool _isActive) returns()
func (_Timelock *TimelockSession) SetHandler(_target common.Address, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetHandler(&_Timelock.TransactOpts, _target, _handler, _isActive)
}

// SetHandler is a paid mutator transaction binding the contract method 0x1154e808.
//
// Solidity: function setHandler(address _target, address _handler, bool _isActive) returns()
func (_Timelock *TimelockTransactorSession) SetHandler(_target common.Address, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetHandler(&_Timelock.TransactOpts, _target, _handler, _isActive)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0x21bd0592.
//
// Solidity: function setInPrivateLiquidationMode(address _vault, bool _inPrivateLiquidationMode) returns()
func (_Timelock *TimelockTransactor) SetInPrivateLiquidationMode(opts *bind.TransactOpts, _vault common.Address, _inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setInPrivateLiquidationMode", _vault, _inPrivateLiquidationMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0x21bd0592.
//
// Solidity: function setInPrivateLiquidationMode(address _vault, bool _inPrivateLiquidationMode) returns()
func (_Timelock *TimelockSession) SetInPrivateLiquidationMode(_vault common.Address, _inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetInPrivateLiquidationMode(&_Timelock.TransactOpts, _vault, _inPrivateLiquidationMode)
}

// SetInPrivateLiquidationMode is a paid mutator transaction binding the contract method 0x21bd0592.
//
// Solidity: function setInPrivateLiquidationMode(address _vault, bool _inPrivateLiquidationMode) returns()
func (_Timelock *TimelockTransactorSession) SetInPrivateLiquidationMode(_vault common.Address, _inPrivateLiquidationMode bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetInPrivateLiquidationMode(&_Timelock.TransactOpts, _vault, _inPrivateLiquidationMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x86803c72.
//
// Solidity: function setInPrivateTransferMode(address _token, bool _inPrivateTransferMode) returns()
func (_Timelock *TimelockTransactor) SetInPrivateTransferMode(opts *bind.TransactOpts, _token common.Address, _inPrivateTransferMode bool) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setInPrivateTransferMode", _token, _inPrivateTransferMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x86803c72.
//
// Solidity: function setInPrivateTransferMode(address _token, bool _inPrivateTransferMode) returns()
func (_Timelock *TimelockSession) SetInPrivateTransferMode(_token common.Address, _inPrivateTransferMode bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetInPrivateTransferMode(&_Timelock.TransactOpts, _token, _inPrivateTransferMode)
}

// SetInPrivateTransferMode is a paid mutator transaction binding the contract method 0x86803c72.
//
// Solidity: function setInPrivateTransferMode(address _token, bool _inPrivateTransferMode) returns()
func (_Timelock *TimelockTransactorSession) SetInPrivateTransferMode(_token common.Address, _inPrivateTransferMode bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetInPrivateTransferMode(&_Timelock.TransactOpts, _token, _inPrivateTransferMode)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0xcd2b1230.
//
// Solidity: function setIsLeverageEnabled(address _vault, bool _isLeverageEnabled) returns()
func (_Timelock *TimelockTransactor) SetIsLeverageEnabled(opts *bind.TransactOpts, _vault common.Address, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setIsLeverageEnabled", _vault, _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0xcd2b1230.
//
// Solidity: function setIsLeverageEnabled(address _vault, bool _isLeverageEnabled) returns()
func (_Timelock *TimelockSession) SetIsLeverageEnabled(_vault common.Address, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetIsLeverageEnabled(&_Timelock.TransactOpts, _vault, _isLeverageEnabled)
}

// SetIsLeverageEnabled is a paid mutator transaction binding the contract method 0xcd2b1230.
//
// Solidity: function setIsLeverageEnabled(address _vault, bool _isLeverageEnabled) returns()
func (_Timelock *TimelockTransactorSession) SetIsLeverageEnabled(_vault common.Address, _isLeverageEnabled bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetIsLeverageEnabled(&_Timelock.TransactOpts, _vault, _isLeverageEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x117cf204.
//
// Solidity: function setIsSwapEnabled(address _vault, bool _isSwapEnabled) returns()
func (_Timelock *TimelockTransactor) SetIsSwapEnabled(opts *bind.TransactOpts, _vault common.Address, _isSwapEnabled bool) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setIsSwapEnabled", _vault, _isSwapEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x117cf204.
//
// Solidity: function setIsSwapEnabled(address _vault, bool _isSwapEnabled) returns()
func (_Timelock *TimelockSession) SetIsSwapEnabled(_vault common.Address, _isSwapEnabled bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetIsSwapEnabled(&_Timelock.TransactOpts, _vault, _isSwapEnabled)
}

// SetIsSwapEnabled is a paid mutator transaction binding the contract method 0x117cf204.
//
// Solidity: function setIsSwapEnabled(address _vault, bool _isSwapEnabled) returns()
func (_Timelock *TimelockTransactorSession) SetIsSwapEnabled(_vault common.Address, _isSwapEnabled bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetIsSwapEnabled(&_Timelock.TransactOpts, _vault, _isSwapEnabled)
}

// SetKeeper is a paid mutator transaction binding the contract method 0xd1b9e853.
//
// Solidity: function setKeeper(address _keeper, bool _isActive) returns()
func (_Timelock *TimelockTransactor) SetKeeper(opts *bind.TransactOpts, _keeper common.Address, _isActive bool) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setKeeper", _keeper, _isActive)
}

// SetKeeper is a paid mutator transaction binding the contract method 0xd1b9e853.
//
// Solidity: function setKeeper(address _keeper, bool _isActive) returns()
func (_Timelock *TimelockSession) SetKeeper(_keeper common.Address, _isActive bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetKeeper(&_Timelock.TransactOpts, _keeper, _isActive)
}

// SetKeeper is a paid mutator transaction binding the contract method 0xd1b9e853.
//
// Solidity: function setKeeper(address _keeper, bool _isActive) returns()
func (_Timelock *TimelockTransactorSession) SetKeeper(_keeper common.Address, _isActive bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetKeeper(&_Timelock.TransactOpts, _keeper, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x47de43e2.
//
// Solidity: function setLiquidator(address _vault, address _liquidator, bool _isActive) returns()
func (_Timelock *TimelockTransactor) SetLiquidator(opts *bind.TransactOpts, _vault common.Address, _liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setLiquidator", _vault, _liquidator, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x47de43e2.
//
// Solidity: function setLiquidator(address _vault, address _liquidator, bool _isActive) returns()
func (_Timelock *TimelockSession) SetLiquidator(_vault common.Address, _liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetLiquidator(&_Timelock.TransactOpts, _vault, _liquidator, _isActive)
}

// SetLiquidator is a paid mutator transaction binding the contract method 0x47de43e2.
//
// Solidity: function setLiquidator(address _vault, address _liquidator, bool _isActive) returns()
func (_Timelock *TimelockTransactorSession) SetLiquidator(_vault common.Address, _liquidator common.Address, _isActive bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetLiquidator(&_Timelock.TransactOpts, _vault, _liquidator, _isActive)
}

// SetMarginFeeBasisPoints is a paid mutator transaction binding the contract method 0xe21b4591.
//
// Solidity: function setMarginFeeBasisPoints(uint256 _marginFeeBasisPoints, uint256 _maxMarginFeeBasisPoints) returns()
func (_Timelock *TimelockTransactor) SetMarginFeeBasisPoints(opts *bind.TransactOpts, _marginFeeBasisPoints *big.Int, _maxMarginFeeBasisPoints *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setMarginFeeBasisPoints", _marginFeeBasisPoints, _maxMarginFeeBasisPoints)
}

// SetMarginFeeBasisPoints is a paid mutator transaction binding the contract method 0xe21b4591.
//
// Solidity: function setMarginFeeBasisPoints(uint256 _marginFeeBasisPoints, uint256 _maxMarginFeeBasisPoints) returns()
func (_Timelock *TimelockSession) SetMarginFeeBasisPoints(_marginFeeBasisPoints *big.Int, _maxMarginFeeBasisPoints *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetMarginFeeBasisPoints(&_Timelock.TransactOpts, _marginFeeBasisPoints, _maxMarginFeeBasisPoints)
}

// SetMarginFeeBasisPoints is a paid mutator transaction binding the contract method 0xe21b4591.
//
// Solidity: function setMarginFeeBasisPoints(uint256 _marginFeeBasisPoints, uint256 _maxMarginFeeBasisPoints) returns()
func (_Timelock *TimelockTransactorSession) SetMarginFeeBasisPoints(_marginFeeBasisPoints *big.Int, _maxMarginFeeBasisPoints *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetMarginFeeBasisPoints(&_Timelock.TransactOpts, _marginFeeBasisPoints, _maxMarginFeeBasisPoints)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0x8bf22c46.
//
// Solidity: function setMaxGasPrice(address _vault, uint256 _maxGasPrice) returns()
func (_Timelock *TimelockTransactor) SetMaxGasPrice(opts *bind.TransactOpts, _vault common.Address, _maxGasPrice *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setMaxGasPrice", _vault, _maxGasPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0x8bf22c46.
//
// Solidity: function setMaxGasPrice(address _vault, uint256 _maxGasPrice) returns()
func (_Timelock *TimelockSession) SetMaxGasPrice(_vault common.Address, _maxGasPrice *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetMaxGasPrice(&_Timelock.TransactOpts, _vault, _maxGasPrice)
}

// SetMaxGasPrice is a paid mutator transaction binding the contract method 0x8bf22c46.
//
// Solidity: function setMaxGasPrice(address _vault, uint256 _maxGasPrice) returns()
func (_Timelock *TimelockTransactorSession) SetMaxGasPrice(_vault common.Address, _maxGasPrice *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetMaxGasPrice(&_Timelock.TransactOpts, _vault, _maxGasPrice)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xdf3a66d9.
//
// Solidity: function setMaxGlobalShortSize(address _vault, address _token, uint256 _amount) returns()
func (_Timelock *TimelockTransactor) SetMaxGlobalShortSize(opts *bind.TransactOpts, _vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setMaxGlobalShortSize", _vault, _token, _amount)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xdf3a66d9.
//
// Solidity: function setMaxGlobalShortSize(address _vault, address _token, uint256 _amount) returns()
func (_Timelock *TimelockSession) SetMaxGlobalShortSize(_vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetMaxGlobalShortSize(&_Timelock.TransactOpts, _vault, _token, _amount)
}

// SetMaxGlobalShortSize is a paid mutator transaction binding the contract method 0xdf3a66d9.
//
// Solidity: function setMaxGlobalShortSize(address _vault, address _token, uint256 _amount) returns()
func (_Timelock *TimelockTransactorSession) SetMaxGlobalShortSize(_vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetMaxGlobalShortSize(&_Timelock.TransactOpts, _vault, _token, _amount)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0x7b6f775a.
//
// Solidity: function setMaxLeverage(address _vault, uint256 _maxLeverage) returns()
func (_Timelock *TimelockTransactor) SetMaxLeverage(opts *bind.TransactOpts, _vault common.Address, _maxLeverage *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setMaxLeverage", _vault, _maxLeverage)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0x7b6f775a.
//
// Solidity: function setMaxLeverage(address _vault, uint256 _maxLeverage) returns()
func (_Timelock *TimelockSession) SetMaxLeverage(_vault common.Address, _maxLeverage *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetMaxLeverage(&_Timelock.TransactOpts, _vault, _maxLeverage)
}

// SetMaxLeverage is a paid mutator transaction binding the contract method 0x7b6f775a.
//
// Solidity: function setMaxLeverage(address _vault, uint256 _maxLeverage) returns()
func (_Timelock *TimelockTransactorSession) SetMaxLeverage(_vault common.Address, _maxLeverage *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetMaxLeverage(&_Timelock.TransactOpts, _vault, _maxLeverage)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address _vault, address _priceFeed) returns()
func (_Timelock *TimelockTransactor) SetPriceFeed(opts *bind.TransactOpts, _vault common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setPriceFeed", _vault, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address _vault, address _priceFeed) returns()
func (_Timelock *TimelockSession) SetPriceFeed(_vault common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.SetPriceFeed(&_Timelock.TransactOpts, _vault, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x76e11286.
//
// Solidity: function setPriceFeed(address _vault, address _priceFeed) returns()
func (_Timelock *TimelockTransactorSession) SetPriceFeed(_vault common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.SetPriceFeed(&_Timelock.TransactOpts, _vault, _priceFeed)
}

// SetReferrerTier is a paid mutator transaction binding the contract method 0xc142940a.
//
// Solidity: function setReferrerTier(address _referralStorage, address _referrer, uint256 _tierId) returns()
func (_Timelock *TimelockTransactor) SetReferrerTier(opts *bind.TransactOpts, _referralStorage common.Address, _referrer common.Address, _tierId *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setReferrerTier", _referralStorage, _referrer, _tierId)
}

// SetReferrerTier is a paid mutator transaction binding the contract method 0xc142940a.
//
// Solidity: function setReferrerTier(address _referralStorage, address _referrer, uint256 _tierId) returns()
func (_Timelock *TimelockSession) SetReferrerTier(_referralStorage common.Address, _referrer common.Address, _tierId *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetReferrerTier(&_Timelock.TransactOpts, _referralStorage, _referrer, _tierId)
}

// SetReferrerTier is a paid mutator transaction binding the contract method 0xc142940a.
//
// Solidity: function setReferrerTier(address _referralStorage, address _referrer, uint256 _tierId) returns()
func (_Timelock *TimelockTransactorSession) SetReferrerTier(_referralStorage common.Address, _referrer common.Address, _tierId *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetReferrerTier(&_Timelock.TransactOpts, _referralStorage, _referrer, _tierId)
}

// SetShortsTrackerAveragePriceWeight is a paid mutator transaction binding the contract method 0x4f5f6b5e.
//
// Solidity: function setShortsTrackerAveragePriceWeight(uint256 _shortsTrackerAveragePriceWeight) returns()
func (_Timelock *TimelockTransactor) SetShortsTrackerAveragePriceWeight(opts *bind.TransactOpts, _shortsTrackerAveragePriceWeight *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setShortsTrackerAveragePriceWeight", _shortsTrackerAveragePriceWeight)
}

// SetShortsTrackerAveragePriceWeight is a paid mutator transaction binding the contract method 0x4f5f6b5e.
//
// Solidity: function setShortsTrackerAveragePriceWeight(uint256 _shortsTrackerAveragePriceWeight) returns()
func (_Timelock *TimelockSession) SetShortsTrackerAveragePriceWeight(_shortsTrackerAveragePriceWeight *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetShortsTrackerAveragePriceWeight(&_Timelock.TransactOpts, _shortsTrackerAveragePriceWeight)
}

// SetShortsTrackerAveragePriceWeight is a paid mutator transaction binding the contract method 0x4f5f6b5e.
//
// Solidity: function setShortsTrackerAveragePriceWeight(uint256 _shortsTrackerAveragePriceWeight) returns()
func (_Timelock *TimelockTransactorSession) SetShortsTrackerAveragePriceWeight(_shortsTrackerAveragePriceWeight *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetShortsTrackerAveragePriceWeight(&_Timelock.TransactOpts, _shortsTrackerAveragePriceWeight)
}

// SetShouldToggleIsLeverageEnabled is a paid mutator transaction binding the contract method 0x8e34c98f.
//
// Solidity: function setShouldToggleIsLeverageEnabled(bool _shouldToggleIsLeverageEnabled) returns()
func (_Timelock *TimelockTransactor) SetShouldToggleIsLeverageEnabled(opts *bind.TransactOpts, _shouldToggleIsLeverageEnabled bool) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setShouldToggleIsLeverageEnabled", _shouldToggleIsLeverageEnabled)
}

// SetShouldToggleIsLeverageEnabled is a paid mutator transaction binding the contract method 0x8e34c98f.
//
// Solidity: function setShouldToggleIsLeverageEnabled(bool _shouldToggleIsLeverageEnabled) returns()
func (_Timelock *TimelockSession) SetShouldToggleIsLeverageEnabled(_shouldToggleIsLeverageEnabled bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetShouldToggleIsLeverageEnabled(&_Timelock.TransactOpts, _shouldToggleIsLeverageEnabled)
}

// SetShouldToggleIsLeverageEnabled is a paid mutator transaction binding the contract method 0x8e34c98f.
//
// Solidity: function setShouldToggleIsLeverageEnabled(bool _shouldToggleIsLeverageEnabled) returns()
func (_Timelock *TimelockTransactorSession) SetShouldToggleIsLeverageEnabled(_shouldToggleIsLeverageEnabled bool) (*types.Transaction, error) {
	return _Timelock.Contract.SetShouldToggleIsLeverageEnabled(&_Timelock.TransactOpts, _shouldToggleIsLeverageEnabled)
}

// SetSwapFees is a paid mutator transaction binding the contract method 0xda762316.
//
// Solidity: function setSwapFees(address _vault, uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints) returns()
func (_Timelock *TimelockTransactor) SetSwapFees(opts *bind.TransactOpts, _vault common.Address, _taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setSwapFees", _vault, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints)
}

// SetSwapFees is a paid mutator transaction binding the contract method 0xda762316.
//
// Solidity: function setSwapFees(address _vault, uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints) returns()
func (_Timelock *TimelockSession) SetSwapFees(_vault common.Address, _taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetSwapFees(&_Timelock.TransactOpts, _vault, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints)
}

// SetSwapFees is a paid mutator transaction binding the contract method 0xda762316.
//
// Solidity: function setSwapFees(address _vault, uint256 _taxBasisPoints, uint256 _stableTaxBasisPoints, uint256 _mintBurnFeeBasisPoints, uint256 _swapFeeBasisPoints, uint256 _stableSwapFeeBasisPoints) returns()
func (_Timelock *TimelockTransactorSession) SetSwapFees(_vault common.Address, _taxBasisPoints *big.Int, _stableTaxBasisPoints *big.Int, _mintBurnFeeBasisPoints *big.Int, _swapFeeBasisPoints *big.Int, _stableSwapFeeBasisPoints *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetSwapFees(&_Timelock.TransactOpts, _vault, _taxBasisPoints, _stableTaxBasisPoints, _mintBurnFeeBasisPoints, _swapFeeBasisPoints, _stableSwapFeeBasisPoints)
}

// SetTier is a paid mutator transaction binding the contract method 0x55818294.
//
// Solidity: function setTier(address _referralStorage, uint256 _tierId, uint256 _totalRebate, uint256 _discountShare) returns()
func (_Timelock *TimelockTransactor) SetTier(opts *bind.TransactOpts, _referralStorage common.Address, _tierId *big.Int, _totalRebate *big.Int, _discountShare *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setTier", _referralStorage, _tierId, _totalRebate, _discountShare)
}

// SetTier is a paid mutator transaction binding the contract method 0x55818294.
//
// Solidity: function setTier(address _referralStorage, uint256 _tierId, uint256 _totalRebate, uint256 _discountShare) returns()
func (_Timelock *TimelockSession) SetTier(_referralStorage common.Address, _tierId *big.Int, _totalRebate *big.Int, _discountShare *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetTier(&_Timelock.TransactOpts, _referralStorage, _tierId, _totalRebate, _discountShare)
}

// SetTier is a paid mutator transaction binding the contract method 0x55818294.
//
// Solidity: function setTier(address _referralStorage, uint256 _tierId, uint256 _totalRebate, uint256 _discountShare) returns()
func (_Timelock *TimelockTransactorSession) SetTier(_referralStorage common.Address, _tierId *big.Int, _totalRebate *big.Int, _discountShare *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetTier(&_Timelock.TransactOpts, _referralStorage, _tierId, _totalRebate, _discountShare)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x0e0dc426.
//
// Solidity: function setTokenConfig(address _vault, address _token, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, uint256 _bufferAmount, uint256 _usdgAmount) returns()
func (_Timelock *TimelockTransactor) SetTokenConfig(opts *bind.TransactOpts, _vault common.Address, _token common.Address, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _bufferAmount *big.Int, _usdgAmount *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setTokenConfig", _vault, _token, _tokenWeight, _minProfitBps, _maxUsdgAmount, _bufferAmount, _usdgAmount)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x0e0dc426.
//
// Solidity: function setTokenConfig(address _vault, address _token, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, uint256 _bufferAmount, uint256 _usdgAmount) returns()
func (_Timelock *TimelockSession) SetTokenConfig(_vault common.Address, _token common.Address, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _bufferAmount *big.Int, _usdgAmount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetTokenConfig(&_Timelock.TransactOpts, _vault, _token, _tokenWeight, _minProfitBps, _maxUsdgAmount, _bufferAmount, _usdgAmount)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x0e0dc426.
//
// Solidity: function setTokenConfig(address _vault, address _token, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, uint256 _bufferAmount, uint256 _usdgAmount) returns()
func (_Timelock *TimelockTransactorSession) SetTokenConfig(_vault common.Address, _token common.Address, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _bufferAmount *big.Int, _usdgAmount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetTokenConfig(&_Timelock.TransactOpts, _vault, _token, _tokenWeight, _minProfitBps, _maxUsdgAmount, _bufferAmount, _usdgAmount)
}

// SetUsdgAmounts is a paid mutator transaction binding the contract method 0x9b53ad22.
//
// Solidity: function setUsdgAmounts(address _vault, address[] _tokens, uint256[] _usdgAmounts) returns()
func (_Timelock *TimelockTransactor) SetUsdgAmounts(opts *bind.TransactOpts, _vault common.Address, _tokens []common.Address, _usdgAmounts []*big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setUsdgAmounts", _vault, _tokens, _usdgAmounts)
}

// SetUsdgAmounts is a paid mutator transaction binding the contract method 0x9b53ad22.
//
// Solidity: function setUsdgAmounts(address _vault, address[] _tokens, uint256[] _usdgAmounts) returns()
func (_Timelock *TimelockSession) SetUsdgAmounts(_vault common.Address, _tokens []common.Address, _usdgAmounts []*big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetUsdgAmounts(&_Timelock.TransactOpts, _vault, _tokens, _usdgAmounts)
}

// SetUsdgAmounts is a paid mutator transaction binding the contract method 0x9b53ad22.
//
// Solidity: function setUsdgAmounts(address _vault, address[] _tokens, uint256[] _usdgAmounts) returns()
func (_Timelock *TimelockTransactorSession) SetUsdgAmounts(_vault common.Address, _tokens []common.Address, _usdgAmounts []*big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SetUsdgAmounts(&_Timelock.TransactOpts, _vault, _tokens, _usdgAmounts)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0xbc476dfd.
//
// Solidity: function setVaultUtils(address _vault, address _vaultUtils) returns()
func (_Timelock *TimelockTransactor) SetVaultUtils(opts *bind.TransactOpts, _vault common.Address, _vaultUtils common.Address) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "setVaultUtils", _vault, _vaultUtils)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0xbc476dfd.
//
// Solidity: function setVaultUtils(address _vault, address _vaultUtils) returns()
func (_Timelock *TimelockSession) SetVaultUtils(_vault common.Address, _vaultUtils common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.SetVaultUtils(&_Timelock.TransactOpts, _vault, _vaultUtils)
}

// SetVaultUtils is a paid mutator transaction binding the contract method 0xbc476dfd.
//
// Solidity: function setVaultUtils(address _vault, address _vaultUtils) returns()
func (_Timelock *TimelockTransactorSession) SetVaultUtils(_vault common.Address, _vaultUtils common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.SetVaultUtils(&_Timelock.TransactOpts, _vault, _vaultUtils)
}

// SignalApprove is a paid mutator transaction binding the contract method 0xdce6e18d.
//
// Solidity: function signalApprove(address _token, address _spender, uint256 _amount) returns()
func (_Timelock *TimelockTransactor) SignalApprove(opts *bind.TransactOpts, _token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "signalApprove", _token, _spender, _amount)
}

// SignalApprove is a paid mutator transaction binding the contract method 0xdce6e18d.
//
// Solidity: function signalApprove(address _token, address _spender, uint256 _amount) returns()
func (_Timelock *TimelockSession) SignalApprove(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SignalApprove(&_Timelock.TransactOpts, _token, _spender, _amount)
}

// SignalApprove is a paid mutator transaction binding the contract method 0xdce6e18d.
//
// Solidity: function signalApprove(address _token, address _spender, uint256 _amount) returns()
func (_Timelock *TimelockTransactorSession) SignalApprove(_token common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SignalApprove(&_Timelock.TransactOpts, _token, _spender, _amount)
}

// SignalMint is a paid mutator transaction binding the contract method 0x09cc9a08.
//
// Solidity: function signalMint(address _token, address _receiver, uint256 _amount) returns()
func (_Timelock *TimelockTransactor) SignalMint(opts *bind.TransactOpts, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "signalMint", _token, _receiver, _amount)
}

// SignalMint is a paid mutator transaction binding the contract method 0x09cc9a08.
//
// Solidity: function signalMint(address _token, address _receiver, uint256 _amount) returns()
func (_Timelock *TimelockSession) SignalMint(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SignalMint(&_Timelock.TransactOpts, _token, _receiver, _amount)
}

// SignalMint is a paid mutator transaction binding the contract method 0x09cc9a08.
//
// Solidity: function signalMint(address _token, address _receiver, uint256 _amount) returns()
func (_Timelock *TimelockTransactorSession) SignalMint(_token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SignalMint(&_Timelock.TransactOpts, _token, _receiver, _amount)
}

// SignalRedeemUsdg is a paid mutator transaction binding the contract method 0x0191c237.
//
// Solidity: function signalRedeemUsdg(address _vault, address _token, uint256 _amount) returns()
func (_Timelock *TimelockTransactor) SignalRedeemUsdg(opts *bind.TransactOpts, _vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "signalRedeemUsdg", _vault, _token, _amount)
}

// SignalRedeemUsdg is a paid mutator transaction binding the contract method 0x0191c237.
//
// Solidity: function signalRedeemUsdg(address _vault, address _token, uint256 _amount) returns()
func (_Timelock *TimelockSession) SignalRedeemUsdg(_vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SignalRedeemUsdg(&_Timelock.TransactOpts, _vault, _token, _amount)
}

// SignalRedeemUsdg is a paid mutator transaction binding the contract method 0x0191c237.
//
// Solidity: function signalRedeemUsdg(address _vault, address _token, uint256 _amount) returns()
func (_Timelock *TimelockTransactorSession) SignalRedeemUsdg(_vault common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SignalRedeemUsdg(&_Timelock.TransactOpts, _vault, _token, _amount)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0x996a7a1e.
//
// Solidity: function signalSetGov(address _target, address _gov) returns()
func (_Timelock *TimelockTransactor) SignalSetGov(opts *bind.TransactOpts, _target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "signalSetGov", _target, _gov)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0x996a7a1e.
//
// Solidity: function signalSetGov(address _target, address _gov) returns()
func (_Timelock *TimelockSession) SignalSetGov(_target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.SignalSetGov(&_Timelock.TransactOpts, _target, _gov)
}

// SignalSetGov is a paid mutator transaction binding the contract method 0x996a7a1e.
//
// Solidity: function signalSetGov(address _target, address _gov) returns()
func (_Timelock *TimelockTransactorSession) SignalSetGov(_target common.Address, _gov common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.SignalSetGov(&_Timelock.TransactOpts, _target, _gov)
}

// SignalSetHandler is a paid mutator transaction binding the contract method 0x24ccbe30.
//
// Solidity: function signalSetHandler(address _target, address _handler, bool _isActive) returns()
func (_Timelock *TimelockTransactor) SignalSetHandler(opts *bind.TransactOpts, _target common.Address, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "signalSetHandler", _target, _handler, _isActive)
}

// SignalSetHandler is a paid mutator transaction binding the contract method 0x24ccbe30.
//
// Solidity: function signalSetHandler(address _target, address _handler, bool _isActive) returns()
func (_Timelock *TimelockSession) SignalSetHandler(_target common.Address, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _Timelock.Contract.SignalSetHandler(&_Timelock.TransactOpts, _target, _handler, _isActive)
}

// SignalSetHandler is a paid mutator transaction binding the contract method 0x24ccbe30.
//
// Solidity: function signalSetHandler(address _target, address _handler, bool _isActive) returns()
func (_Timelock *TimelockTransactorSession) SignalSetHandler(_target common.Address, _handler common.Address, _isActive bool) (*types.Transaction, error) {
	return _Timelock.Contract.SignalSetHandler(&_Timelock.TransactOpts, _target, _handler, _isActive)
}

// SignalSetPriceFeed is a paid mutator transaction binding the contract method 0x80894d62.
//
// Solidity: function signalSetPriceFeed(address _vault, address _priceFeed) returns()
func (_Timelock *TimelockTransactor) SignalSetPriceFeed(opts *bind.TransactOpts, _vault common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "signalSetPriceFeed", _vault, _priceFeed)
}

// SignalSetPriceFeed is a paid mutator transaction binding the contract method 0x80894d62.
//
// Solidity: function signalSetPriceFeed(address _vault, address _priceFeed) returns()
func (_Timelock *TimelockSession) SignalSetPriceFeed(_vault common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.SignalSetPriceFeed(&_Timelock.TransactOpts, _vault, _priceFeed)
}

// SignalSetPriceFeed is a paid mutator transaction binding the contract method 0x80894d62.
//
// Solidity: function signalSetPriceFeed(address _vault, address _priceFeed) returns()
func (_Timelock *TimelockTransactorSession) SignalSetPriceFeed(_vault common.Address, _priceFeed common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.SignalSetPriceFeed(&_Timelock.TransactOpts, _vault, _priceFeed)
}

// SignalVaultSetTokenConfig is a paid mutator transaction binding the contract method 0xdb1c8441.
//
// Solidity: function signalVaultSetTokenConfig(address _vault, address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_Timelock *TimelockTransactor) SignalVaultSetTokenConfig(opts *bind.TransactOpts, _vault common.Address, _token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "signalVaultSetTokenConfig", _vault, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SignalVaultSetTokenConfig is a paid mutator transaction binding the contract method 0xdb1c8441.
//
// Solidity: function signalVaultSetTokenConfig(address _vault, address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_Timelock *TimelockSession) SignalVaultSetTokenConfig(_vault common.Address, _token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _Timelock.Contract.SignalVaultSetTokenConfig(&_Timelock.TransactOpts, _vault, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SignalVaultSetTokenConfig is a paid mutator transaction binding the contract method 0xdb1c8441.
//
// Solidity: function signalVaultSetTokenConfig(address _vault, address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_Timelock *TimelockTransactorSession) SignalVaultSetTokenConfig(_vault common.Address, _token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _Timelock.Contract.SignalVaultSetTokenConfig(&_Timelock.TransactOpts, _vault, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// SignalWithdrawToken is a paid mutator transaction binding the contract method 0x4cd23f3b.
//
// Solidity: function signalWithdrawToken(address _target, address _token, address _receiver, uint256 _amount) returns()
func (_Timelock *TimelockTransactor) SignalWithdrawToken(opts *bind.TransactOpts, _target common.Address, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "signalWithdrawToken", _target, _token, _receiver, _amount)
}

// SignalWithdrawToken is a paid mutator transaction binding the contract method 0x4cd23f3b.
//
// Solidity: function signalWithdrawToken(address _target, address _token, address _receiver, uint256 _amount) returns()
func (_Timelock *TimelockSession) SignalWithdrawToken(_target common.Address, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SignalWithdrawToken(&_Timelock.TransactOpts, _target, _token, _receiver, _amount)
}

// SignalWithdrawToken is a paid mutator transaction binding the contract method 0x4cd23f3b.
//
// Solidity: function signalWithdrawToken(address _target, address _token, address _receiver, uint256 _amount) returns()
func (_Timelock *TimelockTransactorSession) SignalWithdrawToken(_target common.Address, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.SignalWithdrawToken(&_Timelock.TransactOpts, _target, _token, _receiver, _amount)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe4652f49.
//
// Solidity: function transferIn(address _sender, address _token, uint256 _amount) returns()
func (_Timelock *TimelockTransactor) TransferIn(opts *bind.TransactOpts, _sender common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "transferIn", _sender, _token, _amount)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe4652f49.
//
// Solidity: function transferIn(address _sender, address _token, uint256 _amount) returns()
func (_Timelock *TimelockSession) TransferIn(_sender common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.TransferIn(&_Timelock.TransactOpts, _sender, _token, _amount)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe4652f49.
//
// Solidity: function transferIn(address _sender, address _token, uint256 _amount) returns()
func (_Timelock *TimelockTransactorSession) TransferIn(_sender common.Address, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.TransferIn(&_Timelock.TransactOpts, _sender, _token, _amount)
}

// UpdateUsdgSupply is a paid mutator transaction binding the contract method 0x81774b3d.
//
// Solidity: function updateUsdgSupply(uint256 usdgAmount) returns()
func (_Timelock *TimelockTransactor) UpdateUsdgSupply(opts *bind.TransactOpts, usdgAmount *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "updateUsdgSupply", usdgAmount)
}

// UpdateUsdgSupply is a paid mutator transaction binding the contract method 0x81774b3d.
//
// Solidity: function updateUsdgSupply(uint256 usdgAmount) returns()
func (_Timelock *TimelockSession) UpdateUsdgSupply(usdgAmount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.UpdateUsdgSupply(&_Timelock.TransactOpts, usdgAmount)
}

// UpdateUsdgSupply is a paid mutator transaction binding the contract method 0x81774b3d.
//
// Solidity: function updateUsdgSupply(uint256 usdgAmount) returns()
func (_Timelock *TimelockTransactorSession) UpdateUsdgSupply(usdgAmount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.UpdateUsdgSupply(&_Timelock.TransactOpts, usdgAmount)
}

// VaultSetTokenConfig is a paid mutator transaction binding the contract method 0xe3cbeb0f.
//
// Solidity: function vaultSetTokenConfig(address _vault, address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_Timelock *TimelockTransactor) VaultSetTokenConfig(opts *bind.TransactOpts, _vault common.Address, _token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "vaultSetTokenConfig", _vault, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// VaultSetTokenConfig is a paid mutator transaction binding the contract method 0xe3cbeb0f.
//
// Solidity: function vaultSetTokenConfig(address _vault, address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_Timelock *TimelockSession) VaultSetTokenConfig(_vault common.Address, _token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _Timelock.Contract.VaultSetTokenConfig(&_Timelock.TransactOpts, _vault, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// VaultSetTokenConfig is a paid mutator transaction binding the contract method 0xe3cbeb0f.
//
// Solidity: function vaultSetTokenConfig(address _vault, address _token, uint256 _tokenDecimals, uint256 _tokenWeight, uint256 _minProfitBps, uint256 _maxUsdgAmount, bool _isStable, bool _isShortable) returns()
func (_Timelock *TimelockTransactorSession) VaultSetTokenConfig(_vault common.Address, _token common.Address, _tokenDecimals *big.Int, _tokenWeight *big.Int, _minProfitBps *big.Int, _maxUsdgAmount *big.Int, _isStable bool, _isShortable bool) (*types.Transaction, error) {
	return _Timelock.Contract.VaultSetTokenConfig(&_Timelock.TransactOpts, _vault, _token, _tokenDecimals, _tokenWeight, _minProfitBps, _maxUsdgAmount, _isStable, _isShortable)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x0e9587f3.
//
// Solidity: function withdrawFees(address _vault, address _token, address _receiver) returns()
func (_Timelock *TimelockTransactor) WithdrawFees(opts *bind.TransactOpts, _vault common.Address, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "withdrawFees", _vault, _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x0e9587f3.
//
// Solidity: function withdrawFees(address _vault, address _token, address _receiver) returns()
func (_Timelock *TimelockSession) WithdrawFees(_vault common.Address, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.WithdrawFees(&_Timelock.TransactOpts, _vault, _token, _receiver)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x0e9587f3.
//
// Solidity: function withdrawFees(address _vault, address _token, address _receiver) returns()
func (_Timelock *TimelockTransactorSession) WithdrawFees(_vault common.Address, _token common.Address, _receiver common.Address) (*types.Transaction, error) {
	return _Timelock.Contract.WithdrawFees(&_Timelock.TransactOpts, _vault, _token, _receiver)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x21754d9e.
//
// Solidity: function withdrawToken(address _target, address _token, address _receiver, uint256 _amount) returns()
func (_Timelock *TimelockTransactor) WithdrawToken(opts *bind.TransactOpts, _target common.Address, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.contract.Transact(opts, "withdrawToken", _target, _token, _receiver, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x21754d9e.
//
// Solidity: function withdrawToken(address _target, address _token, address _receiver, uint256 _amount) returns()
func (_Timelock *TimelockSession) WithdrawToken(_target common.Address, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.WithdrawToken(&_Timelock.TransactOpts, _target, _token, _receiver, _amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x21754d9e.
//
// Solidity: function withdrawToken(address _target, address _token, address _receiver, uint256 _amount) returns()
func (_Timelock *TimelockTransactorSession) WithdrawToken(_target common.Address, _token common.Address, _receiver common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Timelock.Contract.WithdrawToken(&_Timelock.TransactOpts, _target, _token, _receiver, _amount)
}

// TimelockClearActionIterator is returned from FilterClearAction and is used to iterate over the raw logs and unpacked data for ClearAction events raised by the Timelock contract.
type TimelockClearActionIterator struct {
	Event *TimelockClearAction // Event containing the contract specifics and raw log

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
func (it *TimelockClearActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockClearAction)
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
		it.Event = new(TimelockClearAction)
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
func (it *TimelockClearActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockClearActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockClearAction represents a ClearAction event raised by the Timelock contract.
type TimelockClearAction struct {
	Action [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterClearAction is a free log retrieval operation binding the contract event 0x194ed6dd5e37e2acc44a19455c3f208c4831ee695fe362d9c4ef2d316bc53aec.
//
// Solidity: event ClearAction(bytes32 action)
func (_Timelock *TimelockFilterer) FilterClearAction(opts *bind.FilterOpts) (*TimelockClearActionIterator, error) {

	logs, sub, err := _Timelock.contract.FilterLogs(opts, "ClearAction")
	if err != nil {
		return nil, err
	}
	return &TimelockClearActionIterator{contract: _Timelock.contract, event: "ClearAction", logs: logs, sub: sub}, nil
}

// WatchClearAction is a free log subscription operation binding the contract event 0x194ed6dd5e37e2acc44a19455c3f208c4831ee695fe362d9c4ef2d316bc53aec.
//
// Solidity: event ClearAction(bytes32 action)
func (_Timelock *TimelockFilterer) WatchClearAction(opts *bind.WatchOpts, sink chan<- *TimelockClearAction) (event.Subscription, error) {

	logs, sub, err := _Timelock.contract.WatchLogs(opts, "ClearAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockClearAction)
				if err := _Timelock.contract.UnpackLog(event, "ClearAction", log); err != nil {
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

// ParseClearAction is a log parse operation binding the contract event 0x194ed6dd5e37e2acc44a19455c3f208c4831ee695fe362d9c4ef2d316bc53aec.
//
// Solidity: event ClearAction(bytes32 action)
func (_Timelock *TimelockFilterer) ParseClearAction(log types.Log) (*TimelockClearAction, error) {
	event := new(TimelockClearAction)
	if err := _Timelock.contract.UnpackLog(event, "ClearAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimelockSignalApproveIterator is returned from FilterSignalApprove and is used to iterate over the raw logs and unpacked data for SignalApprove events raised by the Timelock contract.
type TimelockSignalApproveIterator struct {
	Event *TimelockSignalApprove // Event containing the contract specifics and raw log

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
func (it *TimelockSignalApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockSignalApprove)
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
		it.Event = new(TimelockSignalApprove)
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
func (it *TimelockSignalApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockSignalApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockSignalApprove represents a SignalApprove event raised by the Timelock contract.
type TimelockSignalApprove struct {
	Token   common.Address
	Spender common.Address
	Amount  *big.Int
	Action  [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignalApprove is a free log retrieval operation binding the contract event 0x6af9d86ba7407a934e941ed8ae5f779369a88fe8ba2cd1c204185d6f8a8287fd.
//
// Solidity: event SignalApprove(address token, address spender, uint256 amount, bytes32 action)
func (_Timelock *TimelockFilterer) FilterSignalApprove(opts *bind.FilterOpts) (*TimelockSignalApproveIterator, error) {

	logs, sub, err := _Timelock.contract.FilterLogs(opts, "SignalApprove")
	if err != nil {
		return nil, err
	}
	return &TimelockSignalApproveIterator{contract: _Timelock.contract, event: "SignalApprove", logs: logs, sub: sub}, nil
}

// WatchSignalApprove is a free log subscription operation binding the contract event 0x6af9d86ba7407a934e941ed8ae5f779369a88fe8ba2cd1c204185d6f8a8287fd.
//
// Solidity: event SignalApprove(address token, address spender, uint256 amount, bytes32 action)
func (_Timelock *TimelockFilterer) WatchSignalApprove(opts *bind.WatchOpts, sink chan<- *TimelockSignalApprove) (event.Subscription, error) {

	logs, sub, err := _Timelock.contract.WatchLogs(opts, "SignalApprove")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockSignalApprove)
				if err := _Timelock.contract.UnpackLog(event, "SignalApprove", log); err != nil {
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

// ParseSignalApprove is a log parse operation binding the contract event 0x6af9d86ba7407a934e941ed8ae5f779369a88fe8ba2cd1c204185d6f8a8287fd.
//
// Solidity: event SignalApprove(address token, address spender, uint256 amount, bytes32 action)
func (_Timelock *TimelockFilterer) ParseSignalApprove(log types.Log) (*TimelockSignalApprove, error) {
	event := new(TimelockSignalApprove)
	if err := _Timelock.contract.UnpackLog(event, "SignalApprove", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimelockSignalMintIterator is returned from FilterSignalMint and is used to iterate over the raw logs and unpacked data for SignalMint events raised by the Timelock contract.
type TimelockSignalMintIterator struct {
	Event *TimelockSignalMint // Event containing the contract specifics and raw log

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
func (it *TimelockSignalMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockSignalMint)
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
		it.Event = new(TimelockSignalMint)
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
func (it *TimelockSignalMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockSignalMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockSignalMint represents a SignalMint event raised by the Timelock contract.
type TimelockSignalMint struct {
	Token    common.Address
	Receiver common.Address
	Amount   *big.Int
	Action   [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSignalMint is a free log retrieval operation binding the contract event 0x23d37bec99db82564427c9bbfe48ad7434bccf413a40fd357fb838c90a0d6828.
//
// Solidity: event SignalMint(address token, address receiver, uint256 amount, bytes32 action)
func (_Timelock *TimelockFilterer) FilterSignalMint(opts *bind.FilterOpts) (*TimelockSignalMintIterator, error) {

	logs, sub, err := _Timelock.contract.FilterLogs(opts, "SignalMint")
	if err != nil {
		return nil, err
	}
	return &TimelockSignalMintIterator{contract: _Timelock.contract, event: "SignalMint", logs: logs, sub: sub}, nil
}

// WatchSignalMint is a free log subscription operation binding the contract event 0x23d37bec99db82564427c9bbfe48ad7434bccf413a40fd357fb838c90a0d6828.
//
// Solidity: event SignalMint(address token, address receiver, uint256 amount, bytes32 action)
func (_Timelock *TimelockFilterer) WatchSignalMint(opts *bind.WatchOpts, sink chan<- *TimelockSignalMint) (event.Subscription, error) {

	logs, sub, err := _Timelock.contract.WatchLogs(opts, "SignalMint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockSignalMint)
				if err := _Timelock.contract.UnpackLog(event, "SignalMint", log); err != nil {
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

// ParseSignalMint is a log parse operation binding the contract event 0x23d37bec99db82564427c9bbfe48ad7434bccf413a40fd357fb838c90a0d6828.
//
// Solidity: event SignalMint(address token, address receiver, uint256 amount, bytes32 action)
func (_Timelock *TimelockFilterer) ParseSignalMint(log types.Log) (*TimelockSignalMint, error) {
	event := new(TimelockSignalMint)
	if err := _Timelock.contract.UnpackLog(event, "SignalMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimelockSignalPendingActionIterator is returned from FilterSignalPendingAction and is used to iterate over the raw logs and unpacked data for SignalPendingAction events raised by the Timelock contract.
type TimelockSignalPendingActionIterator struct {
	Event *TimelockSignalPendingAction // Event containing the contract specifics and raw log

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
func (it *TimelockSignalPendingActionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockSignalPendingAction)
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
		it.Event = new(TimelockSignalPendingAction)
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
func (it *TimelockSignalPendingActionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockSignalPendingActionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockSignalPendingAction represents a SignalPendingAction event raised by the Timelock contract.
type TimelockSignalPendingAction struct {
	Action [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignalPendingAction is a free log retrieval operation binding the contract event 0x5fb9c0ecf7b4a28c4c480212e868f9da7f373a2ed4d23498b0be6aadf35242fb.
//
// Solidity: event SignalPendingAction(bytes32 action)
func (_Timelock *TimelockFilterer) FilterSignalPendingAction(opts *bind.FilterOpts) (*TimelockSignalPendingActionIterator, error) {

	logs, sub, err := _Timelock.contract.FilterLogs(opts, "SignalPendingAction")
	if err != nil {
		return nil, err
	}
	return &TimelockSignalPendingActionIterator{contract: _Timelock.contract, event: "SignalPendingAction", logs: logs, sub: sub}, nil
}

// WatchSignalPendingAction is a free log subscription operation binding the contract event 0x5fb9c0ecf7b4a28c4c480212e868f9da7f373a2ed4d23498b0be6aadf35242fb.
//
// Solidity: event SignalPendingAction(bytes32 action)
func (_Timelock *TimelockFilterer) WatchSignalPendingAction(opts *bind.WatchOpts, sink chan<- *TimelockSignalPendingAction) (event.Subscription, error) {

	logs, sub, err := _Timelock.contract.WatchLogs(opts, "SignalPendingAction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockSignalPendingAction)
				if err := _Timelock.contract.UnpackLog(event, "SignalPendingAction", log); err != nil {
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

// ParseSignalPendingAction is a log parse operation binding the contract event 0x5fb9c0ecf7b4a28c4c480212e868f9da7f373a2ed4d23498b0be6aadf35242fb.
//
// Solidity: event SignalPendingAction(bytes32 action)
func (_Timelock *TimelockFilterer) ParseSignalPendingAction(log types.Log) (*TimelockSignalPendingAction, error) {
	event := new(TimelockSignalPendingAction)
	if err := _Timelock.contract.UnpackLog(event, "SignalPendingAction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimelockSignalRedeemUsdgIterator is returned from FilterSignalRedeemUsdg and is used to iterate over the raw logs and unpacked data for SignalRedeemUsdg events raised by the Timelock contract.
type TimelockSignalRedeemUsdgIterator struct {
	Event *TimelockSignalRedeemUsdg // Event containing the contract specifics and raw log

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
func (it *TimelockSignalRedeemUsdgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockSignalRedeemUsdg)
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
		it.Event = new(TimelockSignalRedeemUsdg)
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
func (it *TimelockSignalRedeemUsdgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockSignalRedeemUsdgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockSignalRedeemUsdg represents a SignalRedeemUsdg event raised by the Timelock contract.
type TimelockSignalRedeemUsdg struct {
	Vault  common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignalRedeemUsdg is a free log retrieval operation binding the contract event 0xe6bd553b6ef21f3a22ebc877b3aaedc30fe15826b8156d4e8c8b373ebf11d78b.
//
// Solidity: event SignalRedeemUsdg(address vault, address token, uint256 amount)
func (_Timelock *TimelockFilterer) FilterSignalRedeemUsdg(opts *bind.FilterOpts) (*TimelockSignalRedeemUsdgIterator, error) {

	logs, sub, err := _Timelock.contract.FilterLogs(opts, "SignalRedeemUsdg")
	if err != nil {
		return nil, err
	}
	return &TimelockSignalRedeemUsdgIterator{contract: _Timelock.contract, event: "SignalRedeemUsdg", logs: logs, sub: sub}, nil
}

// WatchSignalRedeemUsdg is a free log subscription operation binding the contract event 0xe6bd553b6ef21f3a22ebc877b3aaedc30fe15826b8156d4e8c8b373ebf11d78b.
//
// Solidity: event SignalRedeemUsdg(address vault, address token, uint256 amount)
func (_Timelock *TimelockFilterer) WatchSignalRedeemUsdg(opts *bind.WatchOpts, sink chan<- *TimelockSignalRedeemUsdg) (event.Subscription, error) {

	logs, sub, err := _Timelock.contract.WatchLogs(opts, "SignalRedeemUsdg")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockSignalRedeemUsdg)
				if err := _Timelock.contract.UnpackLog(event, "SignalRedeemUsdg", log); err != nil {
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

// ParseSignalRedeemUsdg is a log parse operation binding the contract event 0xe6bd553b6ef21f3a22ebc877b3aaedc30fe15826b8156d4e8c8b373ebf11d78b.
//
// Solidity: event SignalRedeemUsdg(address vault, address token, uint256 amount)
func (_Timelock *TimelockFilterer) ParseSignalRedeemUsdg(log types.Log) (*TimelockSignalRedeemUsdg, error) {
	event := new(TimelockSignalRedeemUsdg)
	if err := _Timelock.contract.UnpackLog(event, "SignalRedeemUsdg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimelockSignalSetGovIterator is returned from FilterSignalSetGov and is used to iterate over the raw logs and unpacked data for SignalSetGov events raised by the Timelock contract.
type TimelockSignalSetGovIterator struct {
	Event *TimelockSignalSetGov // Event containing the contract specifics and raw log

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
func (it *TimelockSignalSetGovIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockSignalSetGov)
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
		it.Event = new(TimelockSignalSetGov)
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
func (it *TimelockSignalSetGovIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockSignalSetGovIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockSignalSetGov represents a SignalSetGov event raised by the Timelock contract.
type TimelockSignalSetGov struct {
	Target common.Address
	Gov    common.Address
	Action [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignalSetGov is a free log retrieval operation binding the contract event 0x2701a94fd55a560e291f3c54d36580040670d6fde558a77a75d619e38139f713.
//
// Solidity: event SignalSetGov(address target, address gov, bytes32 action)
func (_Timelock *TimelockFilterer) FilterSignalSetGov(opts *bind.FilterOpts) (*TimelockSignalSetGovIterator, error) {

	logs, sub, err := _Timelock.contract.FilterLogs(opts, "SignalSetGov")
	if err != nil {
		return nil, err
	}
	return &TimelockSignalSetGovIterator{contract: _Timelock.contract, event: "SignalSetGov", logs: logs, sub: sub}, nil
}

// WatchSignalSetGov is a free log subscription operation binding the contract event 0x2701a94fd55a560e291f3c54d36580040670d6fde558a77a75d619e38139f713.
//
// Solidity: event SignalSetGov(address target, address gov, bytes32 action)
func (_Timelock *TimelockFilterer) WatchSignalSetGov(opts *bind.WatchOpts, sink chan<- *TimelockSignalSetGov) (event.Subscription, error) {

	logs, sub, err := _Timelock.contract.WatchLogs(opts, "SignalSetGov")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockSignalSetGov)
				if err := _Timelock.contract.UnpackLog(event, "SignalSetGov", log); err != nil {
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

// ParseSignalSetGov is a log parse operation binding the contract event 0x2701a94fd55a560e291f3c54d36580040670d6fde558a77a75d619e38139f713.
//
// Solidity: event SignalSetGov(address target, address gov, bytes32 action)
func (_Timelock *TimelockFilterer) ParseSignalSetGov(log types.Log) (*TimelockSignalSetGov, error) {
	event := new(TimelockSignalSetGov)
	if err := _Timelock.contract.UnpackLog(event, "SignalSetGov", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimelockSignalSetHandlerIterator is returned from FilterSignalSetHandler and is used to iterate over the raw logs and unpacked data for SignalSetHandler events raised by the Timelock contract.
type TimelockSignalSetHandlerIterator struct {
	Event *TimelockSignalSetHandler // Event containing the contract specifics and raw log

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
func (it *TimelockSignalSetHandlerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockSignalSetHandler)
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
		it.Event = new(TimelockSignalSetHandler)
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
func (it *TimelockSignalSetHandlerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockSignalSetHandlerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockSignalSetHandler represents a SignalSetHandler event raised by the Timelock contract.
type TimelockSignalSetHandler struct {
	Target   common.Address
	Handler  common.Address
	IsActive bool
	Action   [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSignalSetHandler is a free log retrieval operation binding the contract event 0x1929c4e13b0dbbad7856b9ce1fc9dca98c7bf7cedd56e22c04dd60ad1d34fe4b.
//
// Solidity: event SignalSetHandler(address target, address handler, bool isActive, bytes32 action)
func (_Timelock *TimelockFilterer) FilterSignalSetHandler(opts *bind.FilterOpts) (*TimelockSignalSetHandlerIterator, error) {

	logs, sub, err := _Timelock.contract.FilterLogs(opts, "SignalSetHandler")
	if err != nil {
		return nil, err
	}
	return &TimelockSignalSetHandlerIterator{contract: _Timelock.contract, event: "SignalSetHandler", logs: logs, sub: sub}, nil
}

// WatchSignalSetHandler is a free log subscription operation binding the contract event 0x1929c4e13b0dbbad7856b9ce1fc9dca98c7bf7cedd56e22c04dd60ad1d34fe4b.
//
// Solidity: event SignalSetHandler(address target, address handler, bool isActive, bytes32 action)
func (_Timelock *TimelockFilterer) WatchSignalSetHandler(opts *bind.WatchOpts, sink chan<- *TimelockSignalSetHandler) (event.Subscription, error) {

	logs, sub, err := _Timelock.contract.WatchLogs(opts, "SignalSetHandler")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockSignalSetHandler)
				if err := _Timelock.contract.UnpackLog(event, "SignalSetHandler", log); err != nil {
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

// ParseSignalSetHandler is a log parse operation binding the contract event 0x1929c4e13b0dbbad7856b9ce1fc9dca98c7bf7cedd56e22c04dd60ad1d34fe4b.
//
// Solidity: event SignalSetHandler(address target, address handler, bool isActive, bytes32 action)
func (_Timelock *TimelockFilterer) ParseSignalSetHandler(log types.Log) (*TimelockSignalSetHandler, error) {
	event := new(TimelockSignalSetHandler)
	if err := _Timelock.contract.UnpackLog(event, "SignalSetHandler", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimelockSignalSetPriceFeedIterator is returned from FilterSignalSetPriceFeed and is used to iterate over the raw logs and unpacked data for SignalSetPriceFeed events raised by the Timelock contract.
type TimelockSignalSetPriceFeedIterator struct {
	Event *TimelockSignalSetPriceFeed // Event containing the contract specifics and raw log

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
func (it *TimelockSignalSetPriceFeedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockSignalSetPriceFeed)
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
		it.Event = new(TimelockSignalSetPriceFeed)
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
func (it *TimelockSignalSetPriceFeedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockSignalSetPriceFeedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockSignalSetPriceFeed represents a SignalSetPriceFeed event raised by the Timelock contract.
type TimelockSignalSetPriceFeed struct {
	Vault     common.Address
	PriceFeed common.Address
	Action    [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignalSetPriceFeed is a free log retrieval operation binding the contract event 0xb878dd4b5762f4118ad54995be907dd2bcd915d942e4ac75580fba9b4ee4727f.
//
// Solidity: event SignalSetPriceFeed(address vault, address priceFeed, bytes32 action)
func (_Timelock *TimelockFilterer) FilterSignalSetPriceFeed(opts *bind.FilterOpts) (*TimelockSignalSetPriceFeedIterator, error) {

	logs, sub, err := _Timelock.contract.FilterLogs(opts, "SignalSetPriceFeed")
	if err != nil {
		return nil, err
	}
	return &TimelockSignalSetPriceFeedIterator{contract: _Timelock.contract, event: "SignalSetPriceFeed", logs: logs, sub: sub}, nil
}

// WatchSignalSetPriceFeed is a free log subscription operation binding the contract event 0xb878dd4b5762f4118ad54995be907dd2bcd915d942e4ac75580fba9b4ee4727f.
//
// Solidity: event SignalSetPriceFeed(address vault, address priceFeed, bytes32 action)
func (_Timelock *TimelockFilterer) WatchSignalSetPriceFeed(opts *bind.WatchOpts, sink chan<- *TimelockSignalSetPriceFeed) (event.Subscription, error) {

	logs, sub, err := _Timelock.contract.WatchLogs(opts, "SignalSetPriceFeed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockSignalSetPriceFeed)
				if err := _Timelock.contract.UnpackLog(event, "SignalSetPriceFeed", log); err != nil {
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

// ParseSignalSetPriceFeed is a log parse operation binding the contract event 0xb878dd4b5762f4118ad54995be907dd2bcd915d942e4ac75580fba9b4ee4727f.
//
// Solidity: event SignalSetPriceFeed(address vault, address priceFeed, bytes32 action)
func (_Timelock *TimelockFilterer) ParseSignalSetPriceFeed(log types.Log) (*TimelockSignalSetPriceFeed, error) {
	event := new(TimelockSignalSetPriceFeed)
	if err := _Timelock.contract.UnpackLog(event, "SignalSetPriceFeed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimelockSignalVaultSetTokenConfigIterator is returned from FilterSignalVaultSetTokenConfig and is used to iterate over the raw logs and unpacked data for SignalVaultSetTokenConfig events raised by the Timelock contract.
type TimelockSignalVaultSetTokenConfigIterator struct {
	Event *TimelockSignalVaultSetTokenConfig // Event containing the contract specifics and raw log

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
func (it *TimelockSignalVaultSetTokenConfigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockSignalVaultSetTokenConfig)
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
		it.Event = new(TimelockSignalVaultSetTokenConfig)
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
func (it *TimelockSignalVaultSetTokenConfigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockSignalVaultSetTokenConfigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockSignalVaultSetTokenConfig represents a SignalVaultSetTokenConfig event raised by the Timelock contract.
type TimelockSignalVaultSetTokenConfig struct {
	Vault         common.Address
	Token         common.Address
	TokenDecimals *big.Int
	TokenWeight   *big.Int
	MinProfitBps  *big.Int
	MaxUsdgAmount *big.Int
	IsStable      bool
	IsShortable   bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSignalVaultSetTokenConfig is a free log retrieval operation binding the contract event 0x3510e9d8245371c6c1061c33781ce16bd0eafa03cd3d0781865036520af4c743.
//
// Solidity: event SignalVaultSetTokenConfig(address vault, address token, uint256 tokenDecimals, uint256 tokenWeight, uint256 minProfitBps, uint256 maxUsdgAmount, bool isStable, bool isShortable)
func (_Timelock *TimelockFilterer) FilterSignalVaultSetTokenConfig(opts *bind.FilterOpts) (*TimelockSignalVaultSetTokenConfigIterator, error) {

	logs, sub, err := _Timelock.contract.FilterLogs(opts, "SignalVaultSetTokenConfig")
	if err != nil {
		return nil, err
	}
	return &TimelockSignalVaultSetTokenConfigIterator{contract: _Timelock.contract, event: "SignalVaultSetTokenConfig", logs: logs, sub: sub}, nil
}

// WatchSignalVaultSetTokenConfig is a free log subscription operation binding the contract event 0x3510e9d8245371c6c1061c33781ce16bd0eafa03cd3d0781865036520af4c743.
//
// Solidity: event SignalVaultSetTokenConfig(address vault, address token, uint256 tokenDecimals, uint256 tokenWeight, uint256 minProfitBps, uint256 maxUsdgAmount, bool isStable, bool isShortable)
func (_Timelock *TimelockFilterer) WatchSignalVaultSetTokenConfig(opts *bind.WatchOpts, sink chan<- *TimelockSignalVaultSetTokenConfig) (event.Subscription, error) {

	logs, sub, err := _Timelock.contract.WatchLogs(opts, "SignalVaultSetTokenConfig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockSignalVaultSetTokenConfig)
				if err := _Timelock.contract.UnpackLog(event, "SignalVaultSetTokenConfig", log); err != nil {
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

// ParseSignalVaultSetTokenConfig is a log parse operation binding the contract event 0x3510e9d8245371c6c1061c33781ce16bd0eafa03cd3d0781865036520af4c743.
//
// Solidity: event SignalVaultSetTokenConfig(address vault, address token, uint256 tokenDecimals, uint256 tokenWeight, uint256 minProfitBps, uint256 maxUsdgAmount, bool isStable, bool isShortable)
func (_Timelock *TimelockFilterer) ParseSignalVaultSetTokenConfig(log types.Log) (*TimelockSignalVaultSetTokenConfig, error) {
	event := new(TimelockSignalVaultSetTokenConfig)
	if err := _Timelock.contract.UnpackLog(event, "SignalVaultSetTokenConfig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimelockSignalWithdrawTokenIterator is returned from FilterSignalWithdrawToken and is used to iterate over the raw logs and unpacked data for SignalWithdrawToken events raised by the Timelock contract.
type TimelockSignalWithdrawTokenIterator struct {
	Event *TimelockSignalWithdrawToken // Event containing the contract specifics and raw log

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
func (it *TimelockSignalWithdrawTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimelockSignalWithdrawToken)
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
		it.Event = new(TimelockSignalWithdrawToken)
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
func (it *TimelockSignalWithdrawTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimelockSignalWithdrawTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimelockSignalWithdrawToken represents a SignalWithdrawToken event raised by the Timelock contract.
type TimelockSignalWithdrawToken struct {
	Target   common.Address
	Token    common.Address
	Receiver common.Address
	Amount   *big.Int
	Action   [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSignalWithdrawToken is a free log retrieval operation binding the contract event 0x9ed7b0f07a9eed51079fab67f6d0f141f167f5b17fdb5a23282280e15fcafed3.
//
// Solidity: event SignalWithdrawToken(address target, address token, address receiver, uint256 amount, bytes32 action)
func (_Timelock *TimelockFilterer) FilterSignalWithdrawToken(opts *bind.FilterOpts) (*TimelockSignalWithdrawTokenIterator, error) {

	logs, sub, err := _Timelock.contract.FilterLogs(opts, "SignalWithdrawToken")
	if err != nil {
		return nil, err
	}
	return &TimelockSignalWithdrawTokenIterator{contract: _Timelock.contract, event: "SignalWithdrawToken", logs: logs, sub: sub}, nil
}

// WatchSignalWithdrawToken is a free log subscription operation binding the contract event 0x9ed7b0f07a9eed51079fab67f6d0f141f167f5b17fdb5a23282280e15fcafed3.
//
// Solidity: event SignalWithdrawToken(address target, address token, address receiver, uint256 amount, bytes32 action)
func (_Timelock *TimelockFilterer) WatchSignalWithdrawToken(opts *bind.WatchOpts, sink chan<- *TimelockSignalWithdrawToken) (event.Subscription, error) {

	logs, sub, err := _Timelock.contract.WatchLogs(opts, "SignalWithdrawToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimelockSignalWithdrawToken)
				if err := _Timelock.contract.UnpackLog(event, "SignalWithdrawToken", log); err != nil {
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

// ParseSignalWithdrawToken is a log parse operation binding the contract event 0x9ed7b0f07a9eed51079fab67f6d0f141f167f5b17fdb5a23282280e15fcafed3.
//
// Solidity: event SignalWithdrawToken(address target, address token, address receiver, uint256 amount, bytes32 action)
func (_Timelock *TimelockFilterer) ParseSignalWithdrawToken(log types.Log) (*TimelockSignalWithdrawToken, error) {
	event := new(TimelockSignalWithdrawToken)
	if err := _Timelock.contract.UnpackLog(event, "SignalWithdrawToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

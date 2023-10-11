package mint

import (
	"cosmossdk.io/math"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/mint/keeper"
	"github.com/cosmos/cosmos-sdk/x/mint/types"
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper, ic types.InflationCalculationFn) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	logger := k.Logger(ctx)

	// fetch stored minter & params
	minter := k.GetMinter(ctx)
	params := k.GetParams(ctx)

	// recalculate inflation rate
	totalStakingSupply := k.StakingTokenSupply(ctx)
	bondedRatio := k.BondedRatio(ctx)
	minter.BlockHeader = ctx.BlockHeader()
	minter.Inflation = ic(ctx, minter, params, bondedRatio)
	minter.AnnualProvisions = minter.NextAnnualProvisions(params, totalStakingSupply)
	k.SetMinter(ctx, minter)

	expectedMintedJMES := minter.BlockProvision(params)

	totalVestingSupplyPercentage := sdk.NewDec(0)
	foreverVestingAccounts := k.GetAccountKeeper().GetAllForeverVestingAccounts(ctx)

	for _, account := range foreverVestingAccounts {
		vestingSupplyPercentage, _ := sdk.NewDecFromStr(account.VestingSupplyPercentage)
		totalVestingSupplyPercentage = totalVestingSupplyPercentage.Add(vestingSupplyPercentage)
	}

	logger.Info("Total vesting supply percentage", "amount", totalVestingSupplyPercentage.String())

	// Portion that is not forever vested as a decimal
	//foreverVestedPercentage := math.NewIntWithDecimal(1, 0).Sub(totalVestingSupplyPercentage).Quo(math.NewIntWithDecimal(1, 0))
	logger.Info("Forever vested percentage", "amount", totalVestingSupplyPercentage.String())
	unlockedJMES := sdk.NewDecFromInt(expectedMintedJMES.Amount).Mul(totalVestingSupplyPercentage).TruncateInt()

	// Minted jmes is the expected minted coins minus the forever vested portion
	mintedJMES := expectedMintedJMES.SubAmount(unlockedJMES)

	logger.Info("Minted JMES", "amount", mintedJMES.String())
	logger.Info("Unlocked JMES", "amount", unlockedJMES.String())
	logger.Info("Total JMES", "amount", expectedMintedJMES.String())

	percentageVestingOfSupply := sdk.NewDec(0)
	totalVestedAmount := sdk.NewDec(0)

	for _, account := range foreverVestingAccounts {
		vestingSupplyPercentage, _ := sdk.NewDecFromStr(account.VestingSupplyPercentage)

		vestingSupplyPercentage.Mul(expectedMintedJMES.Amount.ToLegacyDec())

		vestedForBlock := sdk.NewCoin("ujmes", vestingSupplyPercentage.Mul(expectedMintedJMES.Amount.ToLegacyDec()).TruncateInt())
		logger.Info("Vesting", "account", account.Address, "vestingSupplyPercentage", vestingSupplyPercentage.String(), "vestedForBlock", vestedForBlock.String())

		percentageVestingOfSupply = percentageVestingOfSupply.Add(vestingSupplyPercentage)
		account.AlreadyVested = account.AlreadyVested.Add(vestedForBlock)
		totalVestedAmount = totalVestedAmount.Add(math.LegacyDec(account.AlreadyVested.AmountOf("ujmes")))
		k.GetAccountKeeper().SetAccount(ctx, &account)
	}

	currentSupply := k.GetSupply(ctx, "ujmes").Amount
	expectedNextSupply := currentSupply.Uint64() + mintedJMES.Amount.Uint64()
	logger.Info("Current supply", "amount", currentSupply.String())
	logger.Info("Expected next supply", "amount", expectedNextSupply)

	maxMintableAmount := params.GetMaxMintableAmount()
	logger.Info("Max mintable amount", "amount", maxMintableAmount)

	//logger.Info("Prepare to mint", "blockheight", ctx.BlockHeight(), "mintAmount", mintedCoins.AmountOf("ujmes").String(), ".currentSupply", currentSupply, "expectedNextSupply", expectedNextSupply, "maxSupply", maxMintableAmount, "unlockedVesting", unlockedVesting)

	mintedCoins := sdk.NewCoins(mintedJMES)

	if expectedNextSupply <= maxMintableAmount {
		// ForeverVesting are incentivised to not stake their tokens, so we need to specifically mint some bujmes
		// token to them at a 1:1 ratio with the ujmes tokens unlocked
		mintedBJMES := sdk.NewCoin("bujmes", unlockedJMES)

		// Minted coin is all the mintedJMES + the mintedBJMES
		mintedCoins = mintedCoins.Add(mintedBJMES)

		logger.Info("Minted", "amount", mintedCoins.String())
		err := k.MintCoins(ctx, mintedCoins)
		if err != nil {
			panic(err)
		}
		logger.Info("Minted", "amount", mintedCoins.String())
		// send the minted coins to the fee collector account
		err = k.AddCollectedFees(ctx, mintedCoins)
		if err != nil {
			panic(err)
		}

	} else {
		logger.Info("Abort minting. ", "total", expectedNextSupply, "would exceed", params.MaxMintableAmount)
	}

	event := sdk.NewEvent(
		types.EventTypeMint,
		sdk.NewAttribute(types.AttributeKeyBondedRatio, bondedRatio.String()),
		sdk.NewAttribute(types.AttributeKeyInflation, minter.Inflation.String()),
		sdk.NewAttribute(types.AttributeKeyAnnualProvisions, minter.AnnualProvisions.String()),
	)

	// For all mintedCoins's amounts
	for _, mintedCoin := range mintedCoins {
		if mintedCoin.Amount.IsInt64() {
			defer telemetry.ModuleSetGauge(types.ModuleName, float32(mintedCoin.Amount.Int64()), "minted_tokens")
		}
		event = event.AppendAttributes(sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()))
	}

	ctx.EventManager().EmitEvent(
		event,
	)
}

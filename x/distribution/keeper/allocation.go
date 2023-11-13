package keeper

import (
	"encoding/binary"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"math/rand"
)

// AllocateTokens performs reward and fee distribution to all validators based
// on the F1 fee distribution specification.
func (k Keeper) AllocateTokens(ctx sdk.Context, totalPreviousPower int64, bondedVotes []abci.VoteInfo) {
	logger := k.Logger(ctx)
	// fetch and clear the collected fees for distribution, since this is
	// called in BeginBlock, collected fees will be from the previous block
	// (and distributed to the previous proposer)
	feeCollector := k.authKeeper.GetModuleAccount(ctx, k.feeCollectorName)
	feesCollectedInt := k.bankKeeper.GetAllBalances(ctx, feeCollector.GetAddress())
	totalFeesCollected := sdk.NewDecCoinsFromCoins(feesCollectedInt...)
	blockHeaderHeight := ctx.BlockHeader().Height

	// If we are in IDP, we have a different set of rules
	// During the Initial Distribution Period (IDP), which include the first 483840 blocks (~28 days -)
	// the allocation differs in its ruling.
	// - We extends distribution to reward a random validator (4%)
	// - We do not distribute to DAOs
	// After IDP, on the portion of the distribution that is awarded to DAOs but are not claimed (active grants)
	// We will burn the tokens instead of sending them to the community pool or toward validators.
	// It create an incentive model where, because of the lack of additional minted tokens
	// the prior ones have more "values", incentivising the validators to have scrutinity on DAOs.

	isIDP := blockHeaderHeight <= 483840

	logger.Info("Distribution started at blockheight: ", blockHeaderHeight)
	logger.Info("TotalFeesCollected:", totalFeesCollected)
	logger.Info("isIDP:", isIDP)

	// transfer collected fees to the distribution module account
	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, k.feeCollectorName, types.ModuleName, feesCollectedInt)
	if err != nil {
		panic(err)
	}

	// temporary workaround to keep CanWithdrawInvariant happy
	// general discussions here: https://github.com/cosmos/cosmos-sdk/issues/2906#issuecomment-441867634
	feePool := k.GetFeePool(ctx)

	// Half the value is available for DAO
	feesCollectedForDAO := totalFeesCollected.MulDecTruncate(sdk.NewDecWithPrec(5, 1))
	// The other portion is for our validators
	totalFeesCollectedForValidators := totalFeesCollected.Sub(feesCollectedForDAO)

	logger.Info("DAO Fees Collected", "feesCollectedForDAO", feesCollectedForDAO.String())

	remainingDAOFees := feesCollectedForDAO
	if !isIDP {
		// We prevent the DAOs to get any money until the end of IDP.
		remainingDAOFees = k.AllocateTokensToWinningGrants(ctx, feesCollectedForDAO, remainingDAOFees, blockHeaderHeight)
	}

	// Here we have paid all proposals, we add the remaining to the validator rewards
	totalFeesCollectedForValidators = totalFeesCollectedForValidators.AddCoins(remainingDAOFees)
	logger.Info("After distribution for DAO, total fee for validator", "unspent DAO Remaining", remainingDAOFees.String(), "totalFeesCollectedForValidators", totalFeesCollectedForValidators.String())

	// We reward the validators on the portion of the remaining fees, what is not distributed to DAO goes to validators
	// This means, if not all the DAO funds are distributed, the validators will get more rewards
	// It increase the incentivization towards quality of the proposals that get funded as the validator get rewarded for
	// spending time to do that via the proposer reward he get on top of that. Also affect voting delegators.

	// We keep track of the remaining value for evenly distribution
	remainingFeesForValidators := totalFeesCollectedForValidators

	if blockHeaderHeight > 2 {
		// Only for IDP, first 483840 blocks (~28 days), extend the distribution to reward a random validator despite his power.
		if isIDP {
			// 4% of the fees are distributed as extra rewards
			extraRewardMultiplier, _ := sdk.NewDecFromStr("0.04")
			extraReward := totalFeesCollectedForValidators.MulDecTruncate(extraRewardMultiplier)
			allocateRewardToRandomValidator := sdk.DecCoins{sdk.NewDecCoin("ujmes", extraReward.AmountOf("ujmes").TruncateInt())}

			// Select a random validator to receive the bonus
			randomValidator, _ := k.GetExtraBonusValidator(ctx)
			randomValidatorAddress, _ := randomValidator.GetConsAddr()
			logger.Info("Paying random active validator (IDP period)", "random validator", randomValidatorAddress, "reward", allocateRewardToRandomValidator)
			k.AllocateTokensToValidator(ctx, randomValidator, allocateRewardToRandomValidator)
			remainingFeesForValidators = remainingFeesForValidators.Sub(allocateRewardToRandomValidator)
		}
	}

	// Unlock vesting
	foreverVestingAccounts := k.GetAuthKeeper().GetAllForeverVestingAccounts(ctx)

	sumVestingSupplyPercentage := sdk.NewDec(0)
	convertedVestingPercentages := make([]sdk.Dec, len(foreverVestingAccounts))
	// We need to distribute all bujmes to them accoringly to their own ratio of vestingSupplyPercentage
	for i, account := range foreverVestingAccounts {
		vestingSupplyPercentage, err := sdk.NewDecFromStr(account.VestingSupplyPercentage)
		if err != nil {
			logger.Info("Error while parsing vestingSupplyPercentage", "vestingSupplyPercentage", account.VestingSupplyPercentage)
			continue
		}
		convertedVestingPercentages[i] = vestingSupplyPercentage
		sumVestingSupplyPercentage = sumVestingSupplyPercentage.Add(vestingSupplyPercentage)
	}
	// We now have the total vesting supply percentage, we can calculate the ratio of each account
	for i, account := range foreverVestingAccounts {
		vestedAmount := totalFeesCollected.AmountOf("bujmes").Mul(convertedVestingPercentages[i]).Quo(sumVestingSupplyPercentage)
		bujmesAwarded := sdk.NewCoin("bujmes", vestedAmount.TruncateInt())
		distributedVestedCoins := sdk.NewDecCoinsFromCoins(bujmesAwarded)
		logger.Info("Distributing value to "+account.Address, "distributedVestedCoins", distributedVestedCoins.String())
		addr, err := sdk.AccAddressFromBech32(account.Address)
		if err != nil {
			// handle error, perhaps log it and continue
			continue
		}
		k.AllocateTokensToAddress(ctx, addr, distributedVestedCoins)
		remainingFeesForValidators = remainingFeesForValidators.Sub(distributedVestedCoins)
	}

	// calculate fraction allocated to validators
	// Community tax is set at 0 for jmes-888
	communityTax := k.GetCommunityTax(ctx)
	voteMultiplier := sdk.OneDec().Sub(communityTax)
	logger.Info("Remaining allocated to validators", "remainingFeesForValidators", remainingFeesForValidators.String())

	totalFeesCollectedForValidatorsAfterProposerReward := remainingFeesForValidators
	// allocate tokens proportionally to voting power
	// TODO consider parallelizing later, ref https://github.com/cosmos/cosmos-sdk/pull/3099#discussion_r246276376
	for _, vote := range bondedVotes {
		validator := k.stakingKeeper.ValidatorByConsAddr(ctx, vote.Validator.Address)

		powerFraction := sdk.ZeroDec()

		// TODO consider microslashing for missing votes.
		// ref https://github.com/cosmos/cosmos-sdk/issues/2525#issuecomment-430838701
		if totalPreviousPower != 0 {
			powerFraction = sdk.NewDec(vote.Validator.Power).QuoTruncate(sdk.NewDec(totalPreviousPower))
		} else {
			powerFraction = sdk.NewDec(1).QuoTruncate(sdk.NewDec(int64(len(bondedVotes)))) // all the power is missing, so we distribute evenly
		}

		validatorUnitReward := totalFeesCollectedForValidatorsAfterProposerReward.MulDecTruncate(voteMultiplier).MulDecTruncate(powerFraction)
		validatorReward := sdk.DecCoins{sdk.NewDecCoin("ujmes", validatorUnitReward.AmountOf("ujmes").TruncateInt())}

		k.AllocateTokensToValidator(ctx, validator, validatorReward)
		remainingFeesForValidators = remainingFeesForValidators.Sub(validatorReward)
	}

	//// Looping over forever vesting accounts to get the total percentage of supply vesting
	//portionOfSupplyVesting, _ :=
	//inversePercentage := sdk.NewDec(1).Sub(portionOfSupplyVesting)
	//vestedDivider := inversePercentage.Mul(sdk.NewDec(10))
	//vestedAmount, _ := totalFeesCollected.QuoDec(vestedDivider).TruncateDecimal()

	// Vested amount is the equivalent in ujmes that the bujmes minted
	//vestedAmount := totalFeesCollected.AmountOf("bujmes")
	vestedAmount := sdk.NewCoin("bujmes", sdk.Int(totalFeesCollected.AmountOf("ujmes")))

	logger.Info("Vesting Unlocked", "amount", vestedAmount.String())

	// allocate community funding
	// Keep any remaining to community pool.
	if !remainingFeesForValidators.IsZero() && !remainingFeesForValidators.Empty() {
		logger.Info("Remaining fees for validators", "value", remainingFeesForValidators.String())
		feePool.CommunityPool = feePool.CommunityPool.Add(remainingFeesForValidators...)
		k.SetFeePool(ctx, feePool)
		logger.Info("Allocated community funding", "value", remainingFeesForValidators.String())
	}
}

// AllocateTokensToValidator allocate tokens to a particular validator,
// splitting according to commission.
func (k Keeper) AllocateTokensToValidator(ctx sdk.Context, val stakingtypes.ValidatorI, tokens sdk.DecCoins) {
	// split tokens between validator and delegators according to commission
	commission := tokens.MulDec(val.GetCommission())
	shared := tokens.Sub(commission)

	// update current commission
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCommission,
			sdk.NewAttribute(sdk.AttributeKeyAmount, commission.String()),
			sdk.NewAttribute(types.AttributeKeyValidator, val.GetOperator().String()),
		),
	)
	currentCommission := k.GetValidatorAccumulatedCommission(ctx, val.GetOperator())
	currentCommission.Commission = currentCommission.Commission.Add(commission...)
	k.SetValidatorAccumulatedCommission(ctx, val.GetOperator(), currentCommission)

	// update current rewards
	currentRewards := k.GetValidatorCurrentRewards(ctx, val.GetOperator())
	currentRewards.Rewards = currentRewards.Rewards.Add(shared...)
	k.SetValidatorCurrentRewards(ctx, val.GetOperator(), currentRewards)

	// update outstanding rewards
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRewards,
			sdk.NewAttribute(sdk.AttributeKeyAmount, tokens.String()),
			sdk.NewAttribute(types.AttributeKeyValidator, val.GetOperator().String()),
		),
	)

	outstanding := k.GetValidatorOutstandingRewards(ctx, val.GetOperator())
	outstanding.Rewards = outstanding.Rewards.Add(tokens...)
	k.SetValidatorOutstandingRewards(ctx, val.GetOperator(), outstanding)
}
func (k Keeper) AllocateTokensToWinningGrants(ctx sdk.Context, feesCollectedForDAO sdk.DecCoins, remainingDAOFees sdk.DecCoins, blockHeaderHeight int64) sdk.DecCoins {
	var winningGrants = k.GetWinningGrants(ctx)
	logger := k.Logger(ctx)

	if winningGrants == nil {
		logger.Info("No winning grants")
	} else {
		logger.Info("WinningGrants Set", winningGrants)

		// Ranging, getting the value and the address of each winning grants ordered by ratio
		for _, winningGrant := range winningGrants {
			// Winning grants has a max cap (i.e : if MaxCap is 125 it is then equal to 12.5%), applies on the total available DAO reward
			maxGrantableAmount := feesCollectedForDAO.AmountOf("ujmes").Mul(sdk.NewDecFromInt(winningGrant.MaxCap)).Quo(sdk.NewDec(1000))

			// Allocate token to the DAO address
			logger.Info("=> Winning grant", "DAO", winningGrant.DAO, "Amount", winningGrant.Amount)
			//winningGrantAmount := sdk.NewDec(winningGrant.Amount)
			//winningGrantAmount := sdk.NewDec(winningGrant.Amount)
			winningGrantAmount := sdk.NewDecFromInt(winningGrant.Amount)
			decCoin := sdk.NewDecCoinFromDec("ujmes", winningGrantAmount)
			logger.Info("=> Winning grant", "DAO", winningGrant.DAO, "AmountDecCoin", decCoin.String())

			hasEnoughFundToPay := remainingDAOFees.AmountOf("ujmes").GTE(decCoin.Amount)
			isLessThanMaxGrant := winningGrantAmount.LTE(maxGrantableAmount)

			if !isLessThanMaxGrant {
				logger.Info("=> Grant amount is too high", "DAO", winningGrant.DAO, "Amount", winningGrant.Amount)
			}
			shouldPay := (winningGrant.ExpireAtHeight.GTE(sdk.NewInt(blockHeaderHeight))) && isLessThanMaxGrant

			// from decCoin to decCoins
			distributedWinningGrantCoins := sdk.DecCoins{decCoin}

			logger.Info("Prepare for paying", "shouldPay", shouldPay, "ExpireAtHeight", winningGrant.ExpireAtHeight, "blockHeaderHeight", uint64(blockHeaderHeight), "hasEnougth", hasEnoughFundToPay)

			if hasEnoughFundToPay && shouldPay {
				k.AllocateTokensToAddress(ctx, winningGrant.DAO, distributedWinningGrantCoins)
				remainingDAOFees = remainingDAOFees.Sub(distributedWinningGrantCoins)
				logger.Info("======= DAO Distributing Value to "+winningGrant.DAO.String(), "distributedWinningGrantCoins", distributedWinningGrantCoins.String())
			} else {
				if !hasEnoughFundToPay {
					logger.Info("=> Not enough remaining to distribute to DAO", "DAO", winningGrant.DAO, "Amount", winningGrant.Amount.String())
				}
				if !shouldPay {
					logger.Info("=> Grant expired", "DAO", winningGrant.DAO, "Amount", winningGrant.Amount.String())
				}
			}
		}
	}

	return remainingDAOFees
}
func (k Keeper) AllocateTokensToAddress(ctx sdk.Context, addr sdk.AccAddress, tokens sdk.DecCoins) {
	tokensCoins, _ := tokens.TruncateDecimal()
	if tokensCoins != nil {
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, tokensCoins); err != nil {
			panic(err)
		}
	}

}

// GetExtraBonusValidator returns a random validator from the top 100 validators
// Deterministic based on the AppHash of the context block
func (k Keeper) GetExtraBonusValidator(ctx sdk.Context) (stakingtypes.ValidatorI, bool) {
	// Use AppHash as seed
	seed := ctx.BlockHeader().AppHash
	r := rand.New(rand.NewSource(int64(binary.BigEndian.Uint64(seed))))

	validators := k.stakingKeeper.GetBondedValidatorsByPower(ctx)
	if len(validators) == 0 {
		return nil, false
	}

	// Len is max 100 or the number of validators
	valLen := len(validators)
	if valLen > 100 {
		valLen = 100
	}
	// pick a random validator
	i := r.Intn(valLen)
	return validators[i], true
}

package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// AllocateTokens handles distribution of the collected fees
// bondedVotes is a list of (validator address, validator voted on last block flag) for all
// validators in the bonded set.
func (k Keeper) AllocateTokens(
	ctx sdk.Context, sumPreviousPrecommitPower, totalPreviousPower int64,
	previousProposer sdk.ConsAddress, bondedVotes []abci.VoteInfo,
) {

	logger := k.Logger(ctx)

	// fetch and clear the collected fees for distribution, since this is
	// called in BeginBlock, collected fees will be from the previous block
	// (and distributed to the previous proposer)
	feeCollector := k.authKeeper.GetModuleAccount(ctx, k.feeCollectorName)
	feesCollectedInt := k.bankKeeper.GetAllBalances(ctx, feeCollector.GetAddress())
	feesCollected := sdk.NewDecCoinsFromCoins(feesCollectedInt...)
	blockHeaderHeight := ctx.BlockHeader().Height

	logger.Info("======= ====================")
	logger.Info("======= Distribution Started ===", "blockheight", blockHeaderHeight)
	logger.Info("======= Current Supply " + k.bankKeeper.GetSupply(ctx, "ujmes").String() + "ujmes")
	logger.Info("======= ====================")

	// transfer collected fees to the distribution module account
	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, k.feeCollectorName, types.ModuleName, feesCollectedInt)
	if err != nil {
		panic(err)
	}

	// temporary workaround to keep CanWithdrawInvariant happy
	// general discussions here: https://github.com/cosmos/cosmos-sdk/issues/2906#issuecomment-441867634
	feePool := k.GetFeePool(ctx)
	if totalPreviousPower == 0 {
		feePool.CommunityPool = feePool.CommunityPool.Add(feesCollected...)
		k.SetFeePool(ctx, feePool)
		return
	}

	// calculate fraction votes
	previousFractionVotes := sdk.NewDec(sumPreviousPrecommitPower).Quo(sdk.NewDec(totalPreviousPower))

	// calculate previous proposer reward
	baseProposerReward := k.GetBaseProposerReward(ctx)
	bonusProposerReward := k.GetBonusProposerReward(ctx)
	proposerMultiplier := baseProposerReward.Add(bonusProposerReward.MulTruncate(previousFractionVotes))

	// Half the value is available for DAO
	feesCollectedForDAO := feesCollected.MulDecTruncate(sdk.NewDecWithPrec(5, 1))
	logger.Info("======= DAO Fees Collected", "feesCollectedForDAO", feesCollectedForDAO.String())

	remainingDAOFees := feesCollectedForDAO

	var winningGrants = k.GetWinningGrants(ctx)
	if winningGrants == nil {
		logger.Info("No winning grants")
	} else {
		logger.Info("WinningGrants", winningGrants)
		// Ranging, getting the value and the address of each winning grants ordered by ratio
		for _, winningGrant := range winningGrants {
			// Allocate token to the DAO address
			logger.Debug("=> Winning grant", "DAO", winningGrant.DAO.String(), "Amount", winningGrant.Amount.String())
			decCoin := sdk.NewDecCoinFromDec("ujmes", sdk.Dec(winningGrant.Amount))
			logger.Debug("=> Winning grant", "DAO", winningGrant.DAO.String(), "AmountDecCoin", decCoin.String())

			// from decCoin to decCoins
			distributedWinningGrantCoins := sdk.DecCoins{decCoin}

			var hasEnoughFundToPay = remainingDAOFees.AmountOf("ujmes").GTE(decCoin.Amount)
			var shouldPay = winningGrant.ExpireAtHeight.Uint64() >= uint64(blockHeaderHeight)

			logger.Debug("SHOULD PAY ?", "shouldPay", shouldPay, "ExpireAtHeight", winningGrant.ExpireAtHeight.Uint64(), "blockHeaderHeight", uint64(blockHeaderHeight))

			if hasEnoughFundToPay && shouldPay {
				k.AllocateTokensToAddress(ctx, winningGrant.DAO, distributedWinningGrantCoins)
				remainingDAOFees = remainingDAOFees.Sub(distributedWinningGrantCoins)
				logger.Info("======= DAO Distributing Value to "+winningGrant.DAO.String(), "distributedWinningGrantCoins", distributedWinningGrantCoins.String())
			} else {
				if !hasEnoughFundToPay {
					logger.Info("=> Not enough remaining to distribute to DAO", "DAO", winningGrant.DAO.String(), "Amount", winningGrant.Amount.String())
				}
				if !shouldPay {
					logger.Info("=> Grant expired", "DAO", winningGrant.DAO.String(), "Amount", winningGrant.Amount.String())
				}
			}
		}
	}

	// Here we have paid all proposals
	logger.Info("======= Final unspent DAO Remaining", "distributeDAOValue", remainingDAOFees.String())

	// After distribution, aning reminding of the feesCollectedForDAO is added below
	feesCollectedForValidators := feesCollected.Sub(feesCollectedForDAO).AddCoins(remainingDAOFees)
	remainingFeesCollectedForValidators := feesCollectedForValidators
	logger.Info("======= feesCollectedForValidators", "feesCollectedForValidators", feesCollectedForValidators.String())

	proposerReward := feesCollectedForValidators.MulDecTruncate(proposerMultiplier)
	bonusAbsoluteReward := feesCollectedForValidators.MulDecTruncate(bonusProposerReward.MulTruncate(previousFractionVotes))
	baseAbsoluteReward := feesCollectedForValidators.MulDecTruncate(baseProposerReward.MulTruncate(previousFractionVotes))

	if blockHeaderHeight > 2 {
		previousProposerReward := k.GetPreviousProposerReward(ctx)
		logger.Info("======= Get Previous Proposer Reward ===", "previousproposerReward", previousProposerReward)
		logger.Info("======= Current Proposer Reward", "total reward", proposerReward.String(), "base", baseAbsoluteReward.String(), "bonus", bonusAbsoluteReward.String())

		// pay previous proposer
		proposerValidator := k.stakingKeeper.ValidatorByConsAddr(ctx, previousProposer)

		if proposerValidator != nil {
			ctx.EventManager().EmitEvent(
				sdk.NewEvent(
					types.EventTypeProposerReward,
					sdk.NewAttribute(sdk.AttributeKeyAmount, proposerReward.String()),
					sdk.NewAttribute(types.AttributeKeyValidator, proposerValidator.GetOperator().String()),
				),
			)
			allocateAmountToPreviousValidator := sdk.DecCoins{sdk.NewDecCoin("ujmes", previousProposerReward.RoundInt())}
			logger.Info("======= Paying Previous Proposer", "previousProposer", previousProposer, "allocateAmountToPreviousValidator", allocateAmountToPreviousValidator)

			k.AllocateTokensToValidator(ctx, proposerValidator, allocateAmountToPreviousValidator)
			remainingFeesCollectedForValidators = remainingFeesCollectedForValidators.Sub(proposerReward)
		} else {
			// previous proposer can be unknown if say, the unbonding period is 1 block, so
			// e.g. a validator undelegates at block X, it's removed entirely by
			// block X+1's endblock, then X+2 we need to refer to the previous
			// proposer for X+1, but we've forgotten about them.
			logger.Error(fmt.Sprintf(
				"WARNING: Attempt to allocate proposer rewards to unknown proposer %s. "+
					"This should happen only if the proposer unbonded completely within a single block, "+
					"which generally should not happen except in exceptional circumstances (or fuzz testing). "+
					"We recommend you investigate immediately.",
				previousProposer.String()))
		}
	}
	k.SetPreviousProposerReward(ctx, proposerReward.AmountOf("ujmes"))

	logger.Info("======= DEALING WITH CURRENT REWARD ACCORDINGLY TO REST", "rest", remainingFeesCollectedForValidators.String())

	// calculate fraction allocated to validators
	// Community tax is set at 0 for jmes-888
	communityTax := k.GetCommunityTax(ctx)
	voteMultiplier := sdk.OneDec().Sub(proposerMultiplier).Sub(communityTax)
	logger.Info("======= ======================")
	logger.Info("======= Validators Rewards ===")
	logger.Info("======= ======================")
	// allocate tokens proportionally to voting power
	// TODO consider parallelizing later, ref https://github.com/cosmos/cosmos-sdk/pull/3099#discussion_r246276376
	for _, vote := range bondedVotes {
		validator := k.stakingKeeper.ValidatorByConsAddr(ctx, vote.Validator.Address)

		// TODO consider microslashing for missing votes.
		// ref https://github.com/cosmos/cosmos-sdk/issues/2525#issuecomment-430838701
		powerFraction := sdk.NewDec(vote.Validator.Power).QuoTruncate(sdk.NewDec(totalPreviousPower))

		reward := feesCollectedForValidators.MulDecTruncate(voteMultiplier).MulDecTruncate(powerFraction)
		logger.Info("======= Validator " + string(validator.GetOperator().String()) + " rewarded " + reward.AmountOf("ujmes").String() + " ujmes.")
		k.AllocateTokensToValidator(ctx, validator, reward)
		remainingFeesCollectedForValidators = remainingFeesCollectedForValidators.Sub(reward)
	}

	logger.Info("======= ==================")
	logger.Info("======= Vesting Unlock ===")
	logger.Info("======= ==================")
	foreverVestingAccounts := k.GetAuthKeeper().GetAllForeverVestingAccounts(ctx)
	percentageVestingOfSupply := sdk.NewDec(0)
	totalVestedAmount := sdk.NewDec(0)
	// We iterate over all the accounts and update the vested amount
	for _, account := range foreverVestingAccounts {
		vestedPerBlock := sdk.NewCoin("ujmes", sdk.NewInt(10000))
		vestingSupplyPercentage, _ := sdk.NewDecFromStr(account.VestingSupplyPercentage)
		percentageVestingOfSupply = percentageVestingOfSupply.Add(vestingSupplyPercentage)
		account.AlreadyVested = account.AlreadyVested.Add(vestedPerBlock)
		logger.Info("======= Account " + string(account.GetAddress().String()) + " vested " + vestedPerBlock.Amount.String() + " ujmes." + " Total vested " + account.AlreadyVested.AmountOf("ujmes").String() + " ujmes.")
		totalVestedAmount = totalVestedAmount.Add(account.AlreadyVested.AmountOf("ujmes").ToDec())
		k.authKeeper.SetAccount(ctx, &account)
	}
	inversePercentage := sdk.NewDec(1).Sub(percentageVestingOfSupply)
	vestedDenom := inversePercentage.Mul(sdk.NewDec(10))
	vestedAmount := feesCollected.QuoDec(vestedDenom)
	logger.Info("======= Current Vested Unlock", "vestedAmount", vestedAmount)
	logger.Info("======= Total Vested Unlocked", "totalVestedAmount", totalVestedAmount)

	// During the stake free period (first 483940 blocks (~28 days), the validators selected are
	// 0 min fee requirement if set has room
	if blockHeaderHeight < 483840 {
		//logger.Error("Missing handler for stake-free grace period.")
		// TODO: We need to set up an initial validator blocktime (sort by time).
		// Based on that we run into a deterministic list the 100 validators that received their daily rewards
		// THere is 28 daily unique selection.
		// The same validators are paid for a period of 17280 blocks.
	}

	if blockHeaderHeight > 1 && remainingFeesCollectedForValidators != nil {
		logger.Info(" ")
		logger.Info("======= Allocation post-validator rewards ========")
		logger.Info("======= Available for distribution to faucet " + remainingFeesCollectedForValidators.String())

		faucetAccAddress, err := sdk.AccAddressFromBech32("jmes1g2vaept3rxjvfzyfmem5am5x74n4qygq58jy9v")
		if err != nil {
			panic(err)
		}

		k.AllocateTokensToAddress(ctx, faucetAccAddress, remainingFeesCollectedForValidators)
	}

	// allocate community funding
	// Keep any remaining to community pool.
	feePool.CommunityPool = feePool.CommunityPool.Add(remainingFeesCollectedForValidators...)
	k.SetFeePool(ctx, feePool)
	logger.Info("======= Allocated community funding", "value", remainingFeesCollectedForValidators.String())
}

func (k Keeper) AllocateTokensToAddress(ctx sdk.Context, addr sdk.AccAddress, tokens sdk.DecCoins) {
	tokensCoins, _ := tokens.TruncateDecimal()
	if tokensCoins != nil {
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, tokensCoins); err != nil {
			panic(err)
		}
	}

}

// AllocateTokensToValidator allocate tokens to a particular validator, splitting according to commission
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

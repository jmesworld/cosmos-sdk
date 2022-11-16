package keeper

import (
	"fmt"
	abci "github.com/tendermint/tendermint/abci/types"
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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
	logger.Info("Distribution started", "blockheight", blockHeaderHeight)
	logger.Info("collected fees for distributions", "feesCollected", feesCollected.String())

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
	logger.Info(" proposer multiplier", "bonus percentage", proposerMultiplier.String())
	proposerReward := feesCollected.MulDecTruncate(proposerMultiplier)
	bonusAbsoluteReward := feesCollected.MulDecTruncate(bonusProposerReward.MulTruncate(previousFractionVotes))
	baseAbsoluteReward := feesCollected.MulDecTruncate(baseProposerReward.MulTruncate(previousFractionVotes))
	logger.Info(" proposer reward", "total reward", proposerReward.String(), "base", baseAbsoluteReward.String(), "bonus", bonusAbsoluteReward.String())

	// pay previous proposer
	remaining := feesCollected
	proposerValidator := k.stakingKeeper.ValidatorByConsAddr(ctx, previousProposer)

	if proposerValidator != nil {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeProposerReward,
				sdk.NewAttribute(sdk.AttributeKeyAmount, proposerReward.String()),
				sdk.NewAttribute(types.AttributeKeyValidator, proposerValidator.GetOperator().String()),
			),
		)
		k.AllocateTokensToValidator(ctx, proposerValidator, proposerReward)
		remaining = remaining.Sub(proposerReward)
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

	// calculate fraction allocated to validators
	communityTax := k.GetCommunityTax(ctx)
	logger.Info("Community tax", "tax", communityTax)
	voteMultiplier := sdk.OneDec().Sub(proposerMultiplier).Sub(communityTax)
	logger.Info("Vote Multiplier", "multiplier", voteMultiplier)

	// During the stake free period (first 483940 blocks (~28 days), the validators selected are

	if blockHeaderHeight < 483840 {
		logger.Error("Missing handler for stake-free grace period")
		// TODO: We need to set up an initial validator blocktime (sort by time).
		// Based on that we run into a deterministic list the 100 validators that received their daily rewards
		// THere is 28 daily unique selection.
		// The same validators are paid for a period of 17280 blocks.
	}
	// allocate tokens proportionally to voting power
	// TODO consider parallelizing later, ref https://github.com/cosmos/cosmos-sdk/pull/3099#discussion_r246276376
	for _, vote := range bondedVotes {
		validator := k.stakingKeeper.ValidatorByConsAddr(ctx, vote.Validator.Address)

		// TODO consider microslashing for missing votes.
		// ref https://github.com/cosmos/cosmos-sdk/issues/2525#issuecomment-430838701
		powerFraction := sdk.NewDec(vote.Validator.Power).QuoTruncate(sdk.NewDec(totalPreviousPower))
		reward := feesCollected.MulDecTruncate(voteMultiplier).MulDecTruncate(powerFraction)
		fmt.Fprintln(os.Stderr, "Allocated", reward)
		logger.Info("Allocated reward", "reward", reward)
		logger.Info("Validator for reward", "validator", validator.GetOperator())
		k.AllocateTokensToValidator(ctx, validator, reward)
		remaining = remaining.Sub(reward)

	}

	logger.Info("=> Remaining to distribute", "value", remaining.String())

	quarterReward := remaining.QuoDecTruncate(sdk.NewDec(4))
	logger.Info("=> quarterRemaining to distribute", "value", quarterReward.String())

	dev1Tokens := quarterReward
	dev2Tokens := quarterReward

	remaining = remaining.Sub(quarterReward).Sub(quarterReward)

	logger.Info("New remaining tokens", "value", remaining.String())

	dev1AccAddress, err := sdk.AccAddressFromBech32("jmes1v0d76gdxn7zmh9tg8ne3kxx9m75xu7mq4elg4f")
	if err != nil {
		panic(err)
	}
	dev2AccAddress, err2 := sdk.AccAddressFromBech32("jmes1acm2q0uqtwx2ne5eevm57yumghamr0wt8hsscm")
	if err2 != nil {
		panic(err2)
	}
	logger.Info("Sending "+dev1Tokens.String()+" to ", "address", dev1AccAddress.String())
	logger.Info("Sending "+dev2Tokens.String()+" to ", "address", dev2AccAddress.String())

	// Allocate to dev address
	k.AllocateTokensToAddress(ctx, dev1AccAddress, dev1Tokens)
	k.AllocateTokensToAddress(ctx, dev2AccAddress, dev2Tokens)
	logger.Info("Allocated")

	// allocate community funding
	// TODO: Transfer to smart-contract instead. Modifiable address via proposal voting.
	// Keep any remaining to community pool.
	feePool.CommunityPool = feePool.CommunityPool.Add(remaining...)
	k.SetFeePool(ctx, feePool)
	logger.Info("=> Allocated community funding", "value", remaining.String())
}

func (k Keeper) AllocateTokensToAddress(ctx sdk.Context, addr sdk.AccAddress, tokens sdk.DecCoins) {
	logger := k.Logger(ctx)

	logger.Info("===== Allocate tokens to address", "tokens", tokens)
	logger.Info("===== Allocate tokens to address", "addr", addr)
	tokensCoins, _ := tokens.TruncateDecimal()
	logger.Info("===== Allocate tokensCoins to address", "tokensCoins", tokensCoins)

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, tokensCoins); err != nil {
		panic(err)
	}
	logger.Info("done")
	finalBalance := k.bankKeeper.GetBalance(ctx, addr, tokensCoins.GetDenomByIndex(0))
	logger.Info("FInalBalance", "final", finalBalance)
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

package keeper

import (
	"encoding/binary"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"math/rand"
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
	totalFeesCollected := sdk.NewDecCoinsFromCoins(feesCollectedInt...)
	blockHeaderHeight := ctx.BlockHeader().Height

	logger.Info("Distribution started", "blockheight", blockHeaderHeight)

	// transfer collected fees to the distribution module account
	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, k.feeCollectorName, types.ModuleName, feesCollectedInt)
	if err != nil {
		panic(err)
	}

	// temporary workaround to keep CanWithdrawInvariant happy
	// general discussions here: https://github.com/cosmos/cosmos-sdk/issues/2906#issuecomment-441867634
	feePool := k.GetFeePool(ctx)

	// calculate fraction votes, set default to 1
	previousFractionVotes := sdk.OneDec()

	// totalPreviousPower is total power, sumPreviousPrecommitPower is total signing power
	if totalPreviousPower != 0 {
		previousFractionVotes = sdk.NewDec(sumPreviousPrecommitPower).Quo(sdk.NewDec(totalPreviousPower))
	}

	// calculate previous proposer reward
	baseProposerReward := k.GetBaseProposerReward(ctx)
	bonusProposerReward := k.GetBonusProposerReward(ctx)
	proposerMultiplier := baseProposerReward.Add(bonusProposerReward.MulTruncate(previousFractionVotes))

	// Half the value is available for DAO
	feesCollectedForDAO := totalFeesCollected.MulDecTruncate(sdk.NewDecWithPrec(5, 1))
	// The other portion is for our validators
	totalFeesCollectedForValidators := totalFeesCollected.Sub(feesCollectedForDAO)

	logger.Info("DAO Fees Collected", "feesCollectedForDAO", feesCollectedForDAO.String())

	remainingDAOFees := feesCollectedForDAO

	var winningGrants = k.GetWinningGrants(ctx)
	if winningGrants == nil {
		logger.Info("No winning grants")
	} else {
		logger.Info("WinningGrants Set", winningGrants)

		//maxGrantableAmount is 50% of the feesCollectedForDAO or 25% of total rewards (including released vesting)
		maxGrantableAmount := remainingDAOFees.AmountOf("ujmes").Mul(sdk.NewDec(5)).Quo(sdk.NewDec(10))

		// Ranging, getting the value and the address of each winning grants ordered by ratio
		for _, winningGrant := range winningGrants {
			// Allocate token to the DAO address
			logger.Debug("=> Winning grant", "DAO", winningGrant.DAO.String(), "Amount", winningGrant.Amount.String())
			decCoin := sdk.NewDecCoinFromDec("ujmes", sdk.Dec(winningGrant.Amount))
			logger.Debug("=> Winning grant", "DAO", winningGrant.DAO.String(), "AmountDecCoin", decCoin.String())

			// from decCoin to decCoins
			distributedWinningGrantCoins := sdk.DecCoins{decCoin}

			var hasEnoughFundToPay = remainingDAOFees.AmountOf("ujmes").GTE(decCoin.Amount)
			var respectMaxGrant = remainingDAOFees.AmountOf("ujmes").LTE(maxGrantableAmount)
			if !respectMaxGrant {
				logger.Info("=> Grant amount is too high", "DAO", winningGrant.DAO.String(), "Amount", winningGrant.Amount.String())
			}
			var shouldPay = (winningGrant.ExpireAtHeight.Uint64() >= uint64(blockHeaderHeight)) && respectMaxGrant

			logger.Debug("Prepare for paying", "shouldPay", shouldPay, "ExpireAtHeight", winningGrant.ExpireAtHeight.Uint64(), "blockHeaderHeight", uint64(blockHeaderHeight))

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

	// Here we have paid all proposals, we add the remaining to the validator rewards
	totalFeesCollectedForValidators = totalFeesCollectedForValidators.AddCoins(remainingDAOFees)
	logger.Info("After distribution for DAO, total fee for validator", "unspent DAO Remaining", remainingDAOFees.String(), "totalFeesCollectedForValidators", totalFeesCollectedForValidators.String())

	// We reward the validators on the portion of the remaining fees, what is not distributed to DAO goes to validators
	// This means, if not all the DAO funds are distributed, the validators will get more rewards
	// It increase the incentivization towards quality of the proposals that get funded as the validator get rewarded for
	// spending time to do that via the proposer reward he get on top of that. Also affect voting delegators.
	proposerReward := totalFeesCollectedForValidators.MulDecTruncate(proposerMultiplier)
	k.SetPreviousProposerReward(ctx, proposerReward.AmountOf("ujmes"))

	// We keep track of the remaining value for evenly distribution
	remainingFeesForValidators := totalFeesCollectedForValidators

	if blockHeaderHeight > 2 {
		// Only for IDP, first 483840 blocks (~28 days), extend the distribution to reward a random validator.
		if blockHeaderHeight < 483840 {
			// 4% of the fees are distributed as extra rewards
			extraRewardMultiplier, _ := sdk.NewDecFromStr("0.04")
			extraReward := totalFeesCollectedForValidators.MulDecTruncate(extraRewardMultiplier)

			// Select a random validator to receive the bonus
			randomValidator, _ := k.GetExtraBonusValidator(ctx)
			randomValidatorAddress, _ := randomValidator.GetConsAddr()
			logger.Info("Paying random active validator (IDP period)", "random validator", randomValidatorAddress, "reward", extraReward)
			k.AllocateTokensToValidator(ctx, randomValidator, extraReward)
			remainingFeesForValidators = remainingFeesForValidators.Sub(extraReward)
		}

		previousProposerReward := k.GetPreviousProposerReward(ctx)

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
			logger.Info("Paying previous proposer", "proposer", previousProposer, "reward", allocateAmountToPreviousValidator)

			k.AllocateTokensToValidator(ctx, proposerValidator, allocateAmountToPreviousValidator)
			remainingFeesForValidators = remainingFeesForValidators.Sub(allocateAmountToPreviousValidator)
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
		validatorReward := sdk.DecCoins{sdk.NewDecCoin("ujmes", validatorUnitReward.AmountOf("ujmes").RoundInt())}

		k.AllocateTokensToValidator(ctx, validator, validatorReward)
		remainingFeesForValidators = remainingFeesForValidators.Sub(validatorReward)
	}

	portionOfSupplyVesting, _ := sdk.NewDecFromStr("0.1")
	inversePercentage := sdk.NewDec(1).Sub(portionOfSupplyVesting)
	vestedDenom := inversePercentage.Mul(sdk.NewDec(10))
	vestedAmount := totalFeesCollected.QuoDec(vestedDenom)
	logger.Info("Vested Unlocked", "amount", vestedAmount)

	// allocate community funding
	// Keep any remaining to community pool.
	if !remainingFeesForValidators.IsZero() && !remainingFeesForValidators.Empty() {
		logger.Info("Remaining fees for validators", "value", remainingFeesForValidators.String())
		feePool.CommunityPool = feePool.CommunityPool.Add(remainingFeesForValidators...)
		k.SetFeePool(ctx, feePool)
		logger.Info("Allocated community funding", "value", remainingFeesForValidators.String())
	}
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

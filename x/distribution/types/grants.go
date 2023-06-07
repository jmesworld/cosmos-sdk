package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type WinningGrantExpiration struct {
	AtTime   int64 `json:"at_time"`
	AtHeight int64 `json:"at_height"`
}
type WinningGrants []WinningGrant

type WinningGrant struct {
	DAO            sdk.AccAddress `json:"dao"`
	Amount         sdk.Int        `json:"amount"`
	ExpireAtHeight sdk.Int        `json:"expire_at_height"`
	YesRatio       sdk.Dec        `json:"yes_ratio"`
	ProposalID     sdk.Int        `json:"proposal_id"`
}

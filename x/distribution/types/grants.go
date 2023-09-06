package types

import (
	sdkmath "cosmossdk.io/math"
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

type WinningGrants []WinningGrant

type WinningGrant struct {
	DAO            sdk.AccAddress `json:"dao"`
	Amount         sdkmath.Int    `json:"amount"`
	ExpireAtHeight sdkmath.Int    `json:"expire_at_height"`
	YesRatio       sdk.Dec        `json:"yes_ratio"`
	ProposalID     sdkmath.Int    `json:"proposal_id"`
	MaxCap         sdkmath.Int    `json:"max_cap"`
}

// RawWinningGrant is an intermediate type to help unmarshal JSON
type RawWinningGrant struct {
	DAO            string          `json:"dao"`
	Amount         json.RawMessage `json:"amount"`
	ExpireAtHeight json.RawMessage `json:"expire_at_height"`
	YesRatio       string          `json:"yes_ratio"`
	ProposalID     json.RawMessage `json:"proposal_id"`
	MaxCap         json.RawMessage `json:"max_cap"`
}

func (wg *WinningGrant) UnmarshalJSON(b []byte) error {
	var raw RawWinningGrant
	err := json.Unmarshal(b, &raw)
	if err != nil {
		return err
	}

	wg.DAO, err = sdk.AccAddressFromBech32(raw.DAO)
	if err != nil {
		return err
	}

	wg.Amount, err = parseRawInt(raw.Amount)
	if err != nil {
		return err
	}

	wg.ExpireAtHeight, err = parseRawInt(raw.ExpireAtHeight)
	if err != nil {
		return err
	}

	wg.MaxCap, err = parseRawInt(raw.MaxCap)
	if err != nil {
		return err
	}

	yesRatio, err := sdk.NewDecFromStr(raw.YesRatio)
	if err != nil {
		return err
	}
	wg.YesRatio = yesRatio

	wg.ProposalID, err = parseRawInt(raw.ProposalID)
	if err != nil {
		return err
	}

	return nil
}

func parseRawInt(raw json.RawMessage) (sdkmath.Int, error) {
	var i int64
	var s string
	var err error

	if err = json.Unmarshal(raw, &i); err == nil {
		return sdkmath.NewInt(i), nil
	}

	if err = json.Unmarshal(raw, &s); err != nil {
		return sdkmath.Int{}, err
	}

	if i, err = strconv.ParseInt(s, 10, 64); err != nil {
		return sdkmath.Int{}, err
	}

	return sdkmath.NewInt(i), nil
}

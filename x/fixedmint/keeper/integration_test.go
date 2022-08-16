package keeper_test

import (
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/fixedmint/types"
)

// returns context and an app with updated fixedmint keeper
func createTestApp(isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(isCheckTx)

	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	app.FixedMintKeeper.SetParams(ctx, types.DefaultParams())
	app.FixedMintKeeper.SetMinter(ctx, types.DefaultInitialMinter())

	return app, ctx
}

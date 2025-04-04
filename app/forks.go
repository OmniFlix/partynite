package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlockForks is intended to be run in a chain upgrade.
func BeginBlockForks(ctx sdk.Context, app *PartyNiteApp) {
	for _, fork := range Forks {
		if ctx.BlockHeight() == fork.UpgradeHeight {
			fork.BeginForkLogic(ctx, &app.AppKeepers)
			return
		}
	}
}

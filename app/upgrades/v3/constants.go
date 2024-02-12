package v3

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/OmniFlix/omniflixhub/v3/app/upgrades"
	store "github.com/cosmos/cosmos-sdk/store/types"
)

const UpgradeName = "v3-dev"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateV3UpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{wasmtypes.ModuleName},
	},
}

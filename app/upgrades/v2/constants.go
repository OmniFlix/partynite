package v2

import (
	store "cosmossdk.io/store/types"
	"github.com/OmniFlix/omniflixhub/v5/app/upgrades"
	ibcnfttransfertypes "github.com/bianjieai/nft-transfer/types"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	icqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v8/types"
)

const UpgradeName = "v2"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateV2UpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{
			consensustypes.ModuleName,
			crisistypes.ModuleName,
			group.ModuleName,
			icqtypes.ModuleName,
			ibcnfttransfertypes.ModuleName,
		},
	},
}

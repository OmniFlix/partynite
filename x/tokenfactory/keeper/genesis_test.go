package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/OmniFlix/omniflixhub/v3/x/tokenfactory/types"
)

func (suite *KeeperTestSuite) TestGenesis() {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		FactoryDenoms: []types.GenesisDenom{
			{
				Denom: "factory/cosmos1t7egva48prqmzl59x5ngv4zx0dtrwewcdqdjr8/bitcoin",
				AuthorityMetadata: types.DenomAuthorityMetadata{
					Admin: "cosmos1t7egva48prqmzl59x5ngv4zx0dtrwewcdqdjr8",
				},
			},
			{
				Denom: "factory/cosmos1t7egva48prqmzl59x5ngv4zx0dtrwewcdqdjr8/diff-admin",
				AuthorityMetadata: types.DenomAuthorityMetadata{
					Admin: "cosmos1t7egva48prqmzl59x5ngv4zx0dtrwewcdqdjr8",
				},
			},
			{
				Denom: "factory/cosmos1t7egva48prqmzl59x5ngv4zx0dtrwewcdqdjr8/litecoin",
				AuthorityMetadata: types.DenomAuthorityMetadata{
					Admin: "cosmos1t7egva48prqmzl59x5ngv4zx0dtrwewcdqdjr8",
				},
			},
		},
	}

	suite.SetupTestForInitGenesis()
	app := suite.App

	// Test both with bank denom metadata set, and not set.
	for i, denom := range genesisState.FactoryDenoms {
		// hacky, sets bank metadata to exist if i != 0, to cover both cases.
		if i != 0 {
			app.AppKeepers.BankKeeper.SetDenomMetaData(suite.Ctx, banktypes.Metadata{Base: denom.GetDenom()})
		}
	}

	if err := app.AppKeepers.TokenFactoryKeeper.SetParams(suite.Ctx, types.Params{DenomCreationFee: sdk.Coins{sdk.NewInt64Coin("uflix", 100)}}); err != nil {
		panic(err)
	}
	app.AppKeepers.TokenFactoryKeeper.InitGenesis(suite.Ctx, genesisState)

	// check that the module account is now initialized
	tokenfactoryModuleAccount := app.AppKeepers.AccountKeeper.GetAccount(suite.Ctx, app.AppKeepers.AccountKeeper.GetModuleAddress(types.ModuleName))
	suite.Require().NotNil(tokenfactoryModuleAccount)

	exportedGenesis := app.AppKeepers.TokenFactoryKeeper.ExportGenesis(suite.Ctx)
	suite.Require().NotNil(exportedGenesis)
	suite.Require().Equal(genesisState, *exportedGenesis)
}

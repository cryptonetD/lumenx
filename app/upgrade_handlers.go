package app

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func (app *App) registerUpgradeHandlers() {


	// first automated migration - no changes
	app.UpgradeKeeper.SetUpgradeHandler("v.1.3.1", func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		return app.mm.RunMigrations(ctx, app.configurator, fromVM)
	})
	// first upgrade version from code
	app.UpgradeKeeper.SetUpgradeHandler("v.1.3.1", app.upgradeHandler)

	// version for upgrade test
	app.UpgradeKeeper.SetUpgradeHandler("v1.3.2", app.upgradeHandler)

	// prepared version for authz
	//app.UpgradeKeeper.SetUpgradeHandler("v1.4.0", app.upgradeHandler)

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	// add new modules in 1-ibc upgrade for both FUND-TestNet-2/DevNets and FUND-MainNet-2
	if (upgradeInfo.Name == "v.1.4.0") && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{"authz"},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}

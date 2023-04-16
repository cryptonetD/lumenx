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

	// version for upgrade test
	app.UpgradeKeeper.SetUpgradeHandler("v1.3.2", func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		app.Logger().Info("v1.3.2 upgrade applied.")
		return app.mm.RunMigrations(ctx, app.configurator, fromVM)
	})

	app.UpgradeKeeper.SetUpgradeHandler("v1.3.3", func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		app.Logger().Info("v1.3.3 upgrade applied.")
		return app.mm.RunMigrations(ctx, app.configurator, fromVM)
	})

	app.UpgradeKeeper.SetUpgradeHandler("v1.4.0", func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		app.Logger().Info("v1.4.0 upgrade applied.")
		return app.mm.RunMigrations(ctx, app.configurator, fromVM)
	})

	app.UpgradeKeeper.SetUpgradeHandler("v1.5.0", func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		app.Logger().Info("v1.5.0 upgrade applied.")
		return app.mm.RunMigrations(ctx, app.configurator, fromVM)
	})

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(err)
	}

	// add new modules into the storage class
	if (upgradeInfo.Name == "v1.4.0") && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{"authz"},
		}
		app.Logger().Info("storage class upgrade to v1.4.0 applied.")
		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}

	if upgradeInfo.Name == "v1.5.0" && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		app.Logger().Info("storage class upgrade to v1.5.0 applied.")
		storeUpgrades := storetypes.StoreUpgrades{}
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}
}

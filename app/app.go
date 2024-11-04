package app

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/mychain/x/mymodule"
)

type App struct {
    *baseapp.BaseApp
    cdc *codec.Codec
}

func NewApp() *App {
    cdc := codec.New()
    bApp := baseapp.NewBaseApp("mychain", cdc, nil, nil)
    app := &App{BaseApp: bApp, cdc: cdc}

    // Module Manager
    mm := module.NewManager(
        auth.AppModule{},
        bank.AppModule{},
        genutil.AppModule{},
        staking.AppModule{},
        mymodule.AppModule{},
    )

    mm.RegisterInvariants(app.CrisisKeeper)
    mm.RegisterRoutes(app.Router(), app.QueryRouter())

    return app
}

package mymodule

import (
	"encoding/json"

	"github.com/mychain/x/mymodule/keeper"
	"github.com/mychain/x/mymodule/types"
)

type AppModule struct {
    AppModuleBasic
    keeper keeper.Keeper
}

func NewAppModule(k keeper.Keeper) AppModule {
    return AppModule{
        keeper: k,
    }
}

func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

func (am AppModule) Route() string {
    return types.ModuleName
}

func (am AppModule) NewHandler() sdk.Handler {
    return nil
}

func (am AppModule) QuerierRoute() string {
    return types.ModuleName
}

func (am AppModule) NewQuerierHandler() sdk.Querier {
    return nil
}

func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
    return []abci.ValidatorUpdate{}
}

func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
    return nil
}

package nameservice

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/sdk-tutorials/nameservice/x/nameservice/client/cli"
	"github.com/cosmos/sdk-tutorials/nameservice/x/nameservice/client/rest"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)


// type check to ensure the interface is properly implemented
var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// app module Basics object
type AppModuleBasic struct{}

func (AppModuleBasic) Name() string {
	return ModuleName
}

func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	RegisterCodec(cdc)
}

func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return ModuleCdc.MustMarshalJSON(DefaultGenesisState())
}

// Validation check of the Genesis
func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	var data GenesisState
	err := ModuleCdc.UnmarshalJSON(bz, &data)
	if err != nil {
		return err
	}
	// Once json successfully marshalled, passes along to genesis.go
	return ValidateGenesis(data)
}

// Register rest routes
func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
	rest.RegisterRoutes(ctx, rtr, StoreKey)
}

// Get the root query command of this module
func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetQueryCmd(StoreKey, cdc)
}

// Get the root tx command of this module
func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
	return cli.GetTxCmd(StoreKey, cdc)
}

type AppModule struct {
	AppModuleBasic
	keeper     Keeper
	bankKeeper bank.Keeper
}

// NewAppModule creates a new AppModule Object
func NewAppModule(k Keeper, bankKeeper bank.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         k,
		bankKeeper:     bankKeeper,
	}
}

func (AppModule) Name() string {
	return ModuleName
}

func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}

func (am AppModule) Route() string {
	return RouterKey
}

func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.keeper)
}
func (am AppModule) QuerierRoute() string {
	return QuerierRoute
}

func (am AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(am.keeper)
}

func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

func (am AppModule) EndBlock(sdk.Context, abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState GenesisState
	ModuleCdc.MustUnmarshalJSON(data, &genesisState)
	InitGenesis(ctx, am.keeper, genesisState)
	return []abci.ValidatorUpdate{}
}

func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)
	return ModuleCdc.MustMarshalJSON(gs)
}

// // Type check to ensure the interface is properly implemented
// var (
// 	_ module.AppModule           = AppModule{}
// 	_ module.AppModuleBasic      = AppModuleBasic{}
// )

// // AppModuleBasic defines the basic application module used by the nameservice module.
// type AppModuleBasic struct{}

// // Name returns the nameservice module's name.
// func (AppModuleBasic) Name() string {
// 	return ModuleName
// }

// // RegisterCodec registers the nameservice module's types for the given codec.
// func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
// 	types.RegisterCodec(cdc)
// }

// // DefaultGenesis returns default genesis state as raw bytes for the nameservice
// // module.
// func (AppModuleBasic) DefaultGenesis() json.RawMessage {
// 	return types.ModuleCdc.MustMarshalJSON(types.DefaultGenesisState())
// }

// // ValidateGenesis performs genesis state validation for the nameservice module.
// func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
// 	var data types.GenesisState
// 	err := types.ModuleCdc.UnmarshalJSON(bz, &data)
// 	if err != nil {
// 		return err
// 	}
// 	return types.ValidateGenesis(data)
// }

// // RegisterRESTRoutes registers the REST routes for the nameservice module.
// func (AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
// 	rest.RegisterRoutes(ctx, rtr)
// }

// // GetTxCmd returns the root tx command for the nameservice module.
// func (AppModuleBasic) GetTxCmd(cdc *codec.Codec) *cobra.Command {
// 	return cli.GetTxCmd(cdc)
// }

// // GetQueryCmd returns no root query command for the nameservice module.
// func (AppModuleBasic) GetQueryCmd(cdc *codec.Codec) *cobra.Command {
// 	return cli.GetQueryCmd(StoreKey, cdc)
// }

// //____________________________________________________________________________

// // AppModule implements an application module for the nameservice module.
// type AppModule struct {
// 	AppModuleBasic

// 	keeper        keeper.Keeper
// 	// TODO: Add keepers that your application depends on
// }

// // NewAppModule creates a new AppModule object
// func NewAppModule(k keeper.Keeper, /*TODO: Add Keepers that your application depends on*/) AppModule {
// 	return AppModule{
// 		AppModuleBasic:      AppModuleBasic{},
// 		keeper:              k,
// 		// TODO: Add keepers that your application depends on
// 	}
// }

// // Name returns the nameservice module's name.
// func (AppModule) Name() string {
// 	return types.ModuleName
// }

// // RegisterInvariants registers the nameservice module invariants.
// func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// // Route returns the message routing key for the nameservice module.
// func (AppModule) Route() string {
// 	return types.RouterKey
// }

// // NewHandler returns an sdk.Handler for the nameservice module.
// func (am AppModule) NewHandler() sdk.Handler {
// 	return NewHandler(am.keeper)
// }

// // QuerierRoute returns the nameservice module's querier route name.
// func (AppModule) QuerierRoute() string {
// 	return types.QuerierRoute
// }

// // NewQuerierHandler returns the nameservice module sdk.Querier.
// func (am AppModule) NewQuerierHandler() sdk.Querier {
// 	return types.NewQuerier(am.keeper)
// }

// // InitGenesis performs genesis initialization for the nameservice module. It returns
// // no validator updates.
// func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
// 	var genesisState GenesisState
// 	types.ModuleCdc.MustUnmarshalJSON(data, &genesisState)
// 	InitGenesis(ctx, am.keeper, genesisState)
// 	return []abci.ValidatorUpdate{}
// }

// // ExportGenesis returns the exported genesis state as raw bytes for the nameservice
// // module.
// func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
// 	gs := ExportGenesis(ctx, am.keeper)
// 	return types.ModuleCdc.MustMarshalJSON(gs)
// }

// // BeginBlock returns the begin blocker for the nameservice module.
// func (am AppModule) BeginBlock(ctx sdk.Context, req abci.RequestBeginBlock) {
// 	BeginBlocker(ctx, req, am.keeper)
// }

// // EndBlock returns the end blocker for the nameservice module. It returns no validator
// // updates.
// func (AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
// 	return []abci.ValidatorUpdate{}
// }

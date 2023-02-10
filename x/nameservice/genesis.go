package nameservice

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type GenesisState struct {
	WhoisRecords []Whois `json:"whois_records"`
}

func NewGenesisState(whoIsRecords []Whois) GenesisState {
	return GenesisState{WhoisRecords: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.WhoisRecords {
		if record.Owner == nil {
			return fmt.Errorf("invalid WhoisRecord: Value: %s. Error: Missing Owner", record.Value)
		}
		if record.Value == "" {
			return fmt.Errorf("invalid WhoisRecord: Owner: %s. Error: Missing Value", record.Owner)
		}
		if record.Price == nil {
			return fmt.Errorf("invalid WhoisRecord: Value: %s. Error: Missing Price", record.Value)
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		WhoisRecords: []Whois{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	for _, record := range data.WhoisRecords {
		keeper.SetWhois(ctx, record.Value, record)
	}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Whois
	iterator := k.GetNamesIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		name := string(iterator.Key())
		whois := k.GetWhois(ctx, name)
		records = append(records, whois)

	}
	return GenesisState{WhoisRecords: records}
}

// package nameservice

// import (
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	abci "github.com/tendermint/tendermint/abci/types"
// 	"github.com/huangzhongshejian/test1/x/nameservice/types"
// )

// // InitGenesis initialize default parameters
// // and the keeper's address to pubkey map
// func InitGenesis(ctx sdk.Context, k Keeper /* TODO: Define what keepers the module needs */, data types.GenesisState) {
// 	// TODO: Define logic for when you would like to initialize a new genesis
// }

// // ExportGenesis writes the current store values
// // to a genesis file, which can be imported again
// // with InitGenesis
// func ExportGenesis(ctx sdk.Context, k Keeper) (data GenesisState) {
// 	// TODO: Define logic for exporting state
// 	return types.NewGenesisState()
// }

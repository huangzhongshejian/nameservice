package types
import (
	"fmt"
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

// GenesisState - all nameservice state that must be provided at genesis
// type GenesisState struct {
// 	// TODO: Fill out what is needed by the module for genesis
// }

// // NewGenesisState creates a new GenesisState object
// func NewGenesisState( /* TODO: Fill out with what is needed for genesis state */) GenesisState {
// 	return GenesisState{
// 		// TODO: Fill out according to your genesis state
// 	}
// }

// // DefaultGenesisState - default GenesisState used by Cosmos Hub
// func DefaultGenesisState() GenesisState {
// 	return GenesisState{
// 		// TODO: Fill out according to your genesis state, these values will be initialized but empty
// 	}
// }

// // ValidateGenesis validates the nameservice genesis parameters
// func ValidateGenesis(data GenesisState) error {
// 	// TODO: Create a sanity check to make sure the state conforms to the modules needs
// 	return nil
// }

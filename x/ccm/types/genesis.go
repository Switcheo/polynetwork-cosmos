package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// this line is used by starport scaffolding # ibc/genesistype/import

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		CreatedTxCount:   sdk.ZeroInt(),
		CreatedTxDetails: make(map[string][]byte),
		ReceivedTxIds:    make(map[string][]byte),
		DenomCreators:    make(map[string]string),
		// this line is used by starport scaffolding # ibc/genesistype/default
		// this line is used by starport scaffolding # genesis/types/default
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if !gs.CreatedTxCount.Equal(sdk.NewInt(int64(len(gs.CreatedTxDetails)))) {
		return fmt.Errorf(
			"number of CreatedTxDetails should match CreatedTxCount in ccm genesis, is %d but expected %s",
			len(gs.CreatedTxDetails), gs.CreatedTxCount.String())
	}

	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}

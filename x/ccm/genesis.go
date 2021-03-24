package ccm

import (
	"github.com/Switcheo/polynetwork-cosmos/x/ccm/keeper"
	"github.com/Switcheo/polynetwork-cosmos/x/ccm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the removeMe
	for _, elem := range genState.RemoveMeList {
		k.SetRemoveMe(ctx, *elem)
	}

	// Set removeMe count
	k.SetRemoveMeCount(ctx, uint64(len(genState.RemoveMeList)))

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all removeMe
	removeMeList := k.GetAllRemoveMe(ctx)
	for _, elem := range removeMeList {
		elem := elem
		genesis.RemoveMeList = append(genesis.RemoveMeList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}

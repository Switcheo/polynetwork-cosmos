package lockproxy

import (
	"fmt"

	"github.com/Switcheo/polynetwork-cosmos/x/lockproxy/keeper"
	"github.com/Switcheo/polynetwork-cosmos/x/lockproxy/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.Nonce = k.GetNonce(ctx)

	// iterate Operators
	iter := k.StoreIterator(ctx, keeper.OperatorToLockProxyKey)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		k, v := iter.Key(), iter.Value()
		addr := sdk.AccAddress(k[1:])
		genesis.Operators[addr.String()] = v
	}

	// iterate ChainIDs
	iter1 := k.StoreIterator(ctx, keeper.BindChainIdPrefix)
	defer iter1.Close()
	for ; iter1.Valid(); iter1.Next() {
		k, v := iter1.Key(), iter1.Value()
		genesis.ChainIds[fmt.Sprintf("%x", k)] = v
	}

	// iterate Registries
	iter2 := k.StoreIterator(ctx, keeper.RegistryPrefix)
	defer iter2.Close()
	for ; iter2.Valid(); iter2.Next() {
		k, v := iter2.Key(), iter2.Value()
		genesis.Registries[fmt.Sprintf("%x", k)] = v
	}

	// this line is used by starport scaffolding # genesis/module/export

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}

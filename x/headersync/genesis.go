package headersync

import (
	"encoding/binary"
	"fmt"
	"strconv"

	"github.com/Switcheo/polynetwork-cosmos/x/headersync/keeper"
	"github.com/Switcheo/polynetwork-cosmos/x/headersync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	polycommon "github.com/polynetwork/poly/common"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, v := range genState.ConsensusPeers {
		k.SetConsensusPeers(ctx, *v)
	}

	for _k, v := range genState.CheckpointHashes {
		hash, err := polycommon.Uint256ParseFromBytes(v)
		if err != nil {
			panic(err)
		}
		chainId, err := strconv.Atoi(_k)
		if err != nil {
			panic(err)
		}
		k.SetKeyHeaderHash(ctx, uint64(chainId), hash)
	}

	// this line is used by starport scaffolding # genesis/module/init

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	iter := k.StoreIterator(ctx, keeper.ConsensusPeerPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		k, v := iter.Key(), iter.Value()
		chainId := binary.LittleEndian.Uint64(k[1:])
		p := new(types.ConsensusPeers)
		if err := p.Deserialize(polycommon.NewZeroCopySource(v)); err != nil {
			panic(err)
		}
		genesis.ConsensusPeers[fmt.Sprint(chainId)] = p
	}

	// iterate CheckpointHashes
	iter1 := k.StoreIterator(ctx, keeper.KeyHeaderHashPrefix)
	defer iter1.Close()
	for ; iter1.Valid(); iter1.Next() {
		k, v := iter1.Key(), iter1.Value()
		chainId := binary.LittleEndian.Uint64(k[1:])
		genesis.CheckpointHashes[fmt.Sprint(chainId)] = v
	}

	// this line is used by starport scaffolding # genesis/module/export

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}

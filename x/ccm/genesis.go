package ccm

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/Switcheo/polynetwork-cosmos/x/ccm/keeper"
	"github.com/Switcheo/polynetwork-cosmos/x/ccm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetCrossChainId(ctx, genState.CreatedTxCount)

	store := k.Store(ctx)

	// set details
	for k, v := range genState.CreatedTxDetails {
		txParamHash, err := hex.DecodeString(k)
		if err != nil {
			panic(err)
		}
		store.Set(keeper.GetCrossChainTxKey(txParamHash), v)
	}

	// set tx ids
	for _k, v := range genState.ReceivedTxIds {
		key, err := hex.DecodeString(_k)
		if err != nil {
			panic(err)
		}

		fromChainId := binary.LittleEndian.Uint64(key[0:8])
		toChainId := key[8:]

		if bytes.Compare(toChainId, v) != 0 {
			panic("Invalid toChainId in init genesis!")
		}

		k.PutDoneTx(ctx, fromChainId, toChainId)
	}

	// set denom creators
	for denom, v := range genState.DenomCreators {
		addr, err := sdk.AccAddressFromBech32(v)
		if err != nil {
			panic(err)
		}
		k.SetDenomCreator(ctx, denom, addr)
	}

	// this line is used by starport scaffolding # genesis/module/init

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	txCount, err := k.GetCrossChainId(ctx)
	if err != nil {
		panic(err)
	}
	genesis.CreatedTxCount = txCount

	// iterate CreatedTxDetails
	iter := k.StoreIterator(ctx, keeper.CrossChainTxDetailPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		k, v := iter.Key(), iter.Value()
		genesis.CreatedTxDetails[fmt.Sprintf("%x", k[1:])] = v
	}

	// iterate ReceivedTxIDs
	iter1 := k.StoreIterator(ctx, keeper.CrossChainDoneTxPrefix)
	defer iter1.Close()
	for ; iter1.Valid(); iter1.Next() {
		k, v := iter1.Key(), iter1.Value()
		genesis.ReceivedTxIds[fmt.Sprintf("%x", k[1:])] = v
	}

	// iterate DenomCreators
	iter2 := k.StoreIterator(ctx, keeper.DenomToCreatorPrefix)
	defer iter2.Close()
	for ; iter2.Valid(); iter2.Next() {
		k, v := iter2.Key(), iter2.Value()
		// extract denom
		denom := string(k[1:])
		// convert to accAddress
		addr := sdk.AccAddress(v)
		genesis.DenomCreators[denom] = addr.String()
	}

	// this line is used by starport scaffolding # genesis/module/export

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}

package keeper

import (
	"strconv"

	"github.com/Switcheo/polynetwork-cosmos/x/btcx/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetDenomInfoCount get the total number of DenomInfo
func (k Keeper) GetDenomInfoCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomInfoCountKey))
	byteKey := types.KeyPrefix(types.DenomInfoCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetDenomInfoCount set the total number of DenomInfo
func (k Keeper) SetDenomInfoCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomInfoCountKey))
	byteKey := types.KeyPrefix(types.DenomInfoCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

// CreateDenomInfo creates a DenomInfo with a new id and update the count
func (k Keeper) CreateDenomInfo(ctx sdk.Context, msg types.MsgCreateDenom) {
	// Create the DenomInfo
	count := k.GetDenomInfoCount(ctx)
	var DenomInfo = types.DenomInfo{
		Creator:      msg.Creator,
		Id:           strconv.FormatInt(count, 10),
		Denom:        msg.Denom,
		RedeemScript: msg.RedeemScript,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomInfoKey))
	key := types.KeyPrefix(types.DenomInfoKey + DenomInfo.Id)
	value := k.cdc.MustMarshalBinaryBare(&DenomInfo)
	store.Set(key, value)

	// Update DenomInfo count
	k.SetDenomInfoCount(ctx, count+1)
}

// SetDenomInfo set a specific DenomInfo in the store
func (k Keeper) SetDenomInfo(ctx sdk.Context, DenomInfo types.DenomInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomInfoKey))
	b := k.cdc.MustMarshalBinaryBare(&DenomInfo)
	store.Set(types.KeyPrefix(types.DenomInfoKey+DenomInfo.Id), b)
}

// GetDenomInfo returns a DenomInfo from its id
func (k Keeper) GetDenomInfo(ctx sdk.Context, key string) types.DenomInfo {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomInfoKey))
	var DenomInfo types.DenomInfo
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.DenomInfoKey+key)), &DenomInfo)
	return DenomInfo
}

// HasDenomInfo checks if the DenomInfo exists
func (k Keeper) HasDenomInfo(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomInfoKey))
	return store.Has(types.KeyPrefix(types.DenomInfoKey + id))
}

// GetDenomInfoOwner returns the creator of the DenomInfo
func (k Keeper) GetDenomInfoOwner(ctx sdk.Context, key string) string {
	return k.GetDenomInfo(ctx, key).Creator
}

// GetAllDenomInfo returns all DenomInfo
func (k Keeper) GetAllDenomInfo(ctx sdk.Context) (msgs []types.DenomInfo) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomInfoKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.DenomInfoKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.DenomInfo
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}

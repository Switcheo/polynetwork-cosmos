package keeper

import (
	"context"

	"github.com/Switcheo/polynetwork-cosmos/x/btcx/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DenomInfo(c context.Context, req *types.QueryGetDenomInfoRequest) (*types.QueryGetDenomInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var DenomInfo types.DenomInfo
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomInfoKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.DenomInfoKey+req.Denom)), &DenomInfo)

	return &types.QueryGetDenomInfoResponse{DenomInfo: &DenomInfo}, nil
}

func (k Keeper) DenomCrossChainInfo(c context.Context, req *types.QueryGetDenomCrossChainInfoRequest) (*types.QueryGetDenomCrossChainInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var DenomInfo types.DenomInfo
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DenomInfoKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.DenomInfoKey+req.Denom)), &DenomInfo)

	return &types.QueryGetDenomCrossChainInfoResponse{DenomInfo: &DenomInfo}, nil
}

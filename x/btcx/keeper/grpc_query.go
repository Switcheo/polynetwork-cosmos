package keeper

import (
	"context"

	"github.com/Switcheo/polynetwork-cosmos/x/btcx/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) DenomInfo(c context.Context, req *types.QueryGetDenomInfoRequest) (*types.QueryGetDenomInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// ctx := sdk.UnwrapSDKContext(c)
	// denomInfo := k.GetDenomInfo(ctx, req.Denom)

	return &types.QueryGetDenomInfoResponse{}, nil
}

func (k Keeper) DenomCrossChainInfo(c context.Context, req *types.QueryGetDenomCrossChainInfoRequest) (*types.QueryGetDenomCrossChainInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	return &types.QueryGetDenomCrossChainInfoResponse{}, nil
}

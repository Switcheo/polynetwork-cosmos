package keeper

import (
	"context"

	"github.com/Switcheo/polynetwork-cosmos/x/btcx/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) DenomInfo(c context.Context, req *types.QueryGetDenomInfoRequest) (*types.QueryGetDenomInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	denomInfo := k.GetDenomInfo(ctx, req.Denom)

	return &types.QueryGetDenomInfoResponse{DenomInfo: denomInfo}, nil
}

func (k Keeper) DenomCrossChainInfo(c context.Context, req *types.QueryGetDenomCrossChainInfoRequest) (*types.QueryGetDenomCrossChainInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	denomCrossChainInfo := k.GetDenomCrossChainInfo(ctx, req.Denom, req.ChainId)

	return &types.QueryGetDenomCrossChainInfoResponse{
		DenomInfo:   denomCrossChainInfo.DenomInfo,
		ToChainId:   denomCrossChainInfo.ToChainId,
		ToAssetHash: denomCrossChainInfo.ToAssetHash,
	}, nil
}

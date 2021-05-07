package keeper

import (
	"context"

	"github.com/Switcheo/polynetwork-cosmos/x/lockproxy/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Proxy(c context.Context, req *types.QueryGetProxyRequest) (*types.QueryGetProxyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	operatorAddr, err := sdk.AccAddressFromBech32(req.OperatorAddress)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid address")
	}

	ctx := sdk.UnwrapSDKContext(c)
	proxyHash := k.GetLockProxy(ctx, operatorAddr)

	return &types.QueryGetProxyResponse{ProxyHash: proxyHash}, nil
}

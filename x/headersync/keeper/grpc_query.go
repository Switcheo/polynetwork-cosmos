package keeper

import (
	"context"

	"github.com/Switcheo/polynetwork-cosmos/x/headersync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ConsensusPeers(c context.Context, req *types.QueryGetConsensusPeersRequest) (res *types.QueryGetConsensusPeersResponse, err error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	peers, err := k.GetConsensusPeers(ctx, req.ChainId)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "Failed to get consensus peers for chainID: %d, error: %v", req.ChainId, err)
	}

	res = &types.QueryGetConsensusPeersResponse{
		ConsensusPeers: *peers,
	}

	return
}

package keeper

import (
	"context"

	"github.com/Switcheo/polynetwork-cosmos/x/headersync/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) ConsensusPeers(c context.Context, req *types.QueryGetConsensusPeersRequest) (res *types.QueryGetConsensusPeersResponse, err error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	// TODO

	return
}

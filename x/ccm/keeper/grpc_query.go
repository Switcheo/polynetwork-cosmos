package keeper

import (
	"context"
	"fmt"

	"github.com/Switcheo/polynetwork-cosmos/x/ccm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) CheckModuleContract(c context.Context, req *types.QueryCheckModuleContractRequest) (res *types.QueryCheckModuleContractResponse, err error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	unlockKeeper, ok := k.ulkMap[req.ModuleName]
	if !ok {
		err = types.ErrCheckModuleContract(fmt.Sprintf("map doesnot contain current keystore for moduleName: %s", req.ModuleName))
		return
	}
	res.ModuleName = req.ModuleName
	res.Exist = unlockKeeper.ContainToContractAddr(ctx, req.ToContractAddr, req.FromChainID)
	return
}

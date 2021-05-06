package keeper

import (
	"context"

	"github.com/Switcheo/polynetwork-cosmos/x/headersync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k Keeper) SyncGenesis(c context.Context, msg *types.MsgSyncGenesis) (*types.MsgSyncGenesisResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.SyncGenesisHeader(ctx, msg.GenesisHeader)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgSyncGenesisResponse{}, nil
}

func (k Keeper) SyncHeaders(c context.Context, msg *types.MsgSyncHeaders) (*types.MsgSyncHeadersResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	err := k.SyncBlockHeaders(ctx, msg.Headers)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgSyncHeadersResponse{}, nil
}

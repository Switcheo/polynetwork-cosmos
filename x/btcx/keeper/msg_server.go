package keeper

import (
	"context"

	"github.com/Switcheo/polynetwork-cosmos/x/btcx/types"
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

func (k Keeper) Create(c context.Context, msg *types.MsgCreate) (*types.MsgCreateResponse, error) {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)
	if err := k.CreateAsset(ctx, creator, msg.Denom, msg.RedeemScript); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgCreateResponse{}, nil
}

func (k Keeper) Bind(c context.Context, msg *types.MsgBind) (*types.MsgBindResponse, error) {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)
	if err := k.BindAsset(ctx, creator, msg.SourceAssetDenom, msg.ToChainId, msg.ToAssetHash); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgBindResponse{}, nil
}

func (k Keeper) Lock(c context.Context, msg *types.MsgLock) (*types.MsgLockResponse, error) {
	fromAddr, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)
	if err := k.LockAsset(ctx, fromAddr, msg.SourceAssetDenom, msg.ToChainId, msg.ToAddressBs, msg.Value.Int); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgLockResponse{}, nil
}

package keeper

import (
	"context"

	"github.com/Switcheo/polynetwork-cosmos/x/lockproxy/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the reqServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &MsgServer{Keeper: keeper}
}

var _ types.MsgServer = MsgServer{}

func (k Keeper) Create(c context.Context, req *types.MsgCreate) (res *types.MsgCreateResponse, err error) {
	creator, err := sdk.AccAddressFromBech32(req.Creator)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)

	err = k.CreateLockProxy(ctx, creator)
	if err != nil {
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

func (k Keeper) Bind(c context.Context, req *types.MsgBind) (res *types.MsgBindResponse, err error) {
	creator, err := sdk.AccAddressFromBech32(req.Creator)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)

	if err := k.CreateCoinAndDelegateToProxy(ctx, creator, req.Denom, req.LockProxyHash,
		req.NativeChainId, req.NativeLockProxyHash, req.NativeAssetHash); err != nil {
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

func (k Keeper) Lock(c context.Context, req *types.MsgLock) (res *types.MsgLockResponse, err error) {
	fromAddress, err := sdk.AccAddressFromBech32(req.FromAddress)
	if err != nil {
		return nil, err
	}

	feeAddress, err := sdk.AccAddressFromBech32(req.FeeAddress)
	if req.DeductFeeInLock && err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(c)

	if err := k.LockAsset(ctx, req.Denom, req.FromLockProxy, req.FromAssetId, fromAddress,
		req.ToChainId, req.ToLockProxy, req.ToAssetId, req.ToAddress,
		req.Amount, req.DeductFeeInLock, req.FeeAmount, feeAddress); err != nil {
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

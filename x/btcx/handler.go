package btcx

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"github.com/Switcheo/polynetwork-cosmos/x/btcx/keeper"
	"github.com/Switcheo/polynetwork-cosmos/x/btcx/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		msgServer := keeper.NewMsgServerImpl(k)

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case *types.MsgCreate:
			res, err := msgServer.Create(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgBind:
			res, err := msgServer.Bind(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgLock:
			res, err := msgServer.Lock(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, errorsmod.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

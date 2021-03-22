package btcx

import (
	"github.com/Switcheo/polynetwork-cosmos/x/btcx/keeper"
	"github.com/Switcheo/polynetwork-cosmos/x/btcx/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func handleMsgCreateDenom(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateDenom) (*sdk.Result, error) {
	k.CreateDenomInfo(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

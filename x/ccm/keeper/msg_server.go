package keeper

import (
	"context"

	"github.com/Switcheo/polynetwork-cosmos/x/ccm/types"
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

func (k Keeper) Process(c context.Context, req *types.MsgProcessCrossChainTx) (res *types.MsgProcessCrossChainTxResponse, err error) {
	ctx := sdk.UnwrapSDKContext(c)

	err = k.ProcessCrossChainTx(ctx, req.FromChainID, req.Proof, req.Header, req.HeaderProof, req.CurHeader)
	if err != nil {
		return nil, err
	}

	res = &types.MsgProcessCrossChainTxResponse{}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgProcessCrossChainTxResponse{}, nil
}

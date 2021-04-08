package keeper

import (
	"context"
	"encoding/hex"
	"fmt"

	"github.com/Switcheo/polynetwork-cosmos/x/ccm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	polycommon "github.com/polynetwork/poly/common"
	polytype "github.com/polynetwork/poly/core/types"
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

func (k Keeper) ProcessCrossChainTx(c context.Context, req *types.MsgProcessCrossChainTx) (res *types.MsgProcessCrossChainTxResponse, err error) {
	ctx := sdk.UnwrapSDKContext(c)

	headerToBeVerified := new(polytype.Header)
	headerBs, err := hex.DecodeString(req.Header)
	if err != nil {
		return nil, types.ErrProcessCrossChainTx(fmt.Sprintf("Failed to decode header hex string to bytes: %s, Error: %s ", req.Header, err.Error()))
	}
	if err := headerToBeVerified.Deserialization(polycommon.NewZeroCopySource(headerBs)); err != nil {
		return nil, types.ErrProcessCrossChainTx(fmt.Sprintf("Failed to deserialize header, Error:%s", err.Error()))
	}

	headerInCurEpoch := new(polytype.Header)
	curHeaderBs, err := hex.DecodeString(req.CurHeader)
	if err != nil {
		headerInCurEpoch = nil
	} else {
		if err := headerInCurEpoch.Deserialization(polycommon.NewZeroCopySource(curHeaderBs)); err != nil {
			headerInCurEpoch = nil
		}
	}

	headerProof, err := hex.DecodeString(req.HeaderProof)
	if err != nil {
		headerProof = nil
	}

	if err := k.hsk.ProcessHeader(ctx, headerToBeVerified, headerProof, headerInCurEpoch); err != nil {
		return nil, types.ErrProcessCrossChainTx(fmt.Sprintf("ProcessHeader Error, %s", err.Error()))
	}

	proof, err := hex.DecodeString(req.Proof)
	if err != nil {
		return nil, types.ErrProcessCrossChainTx(fmt.Sprintf("Failed to decode proof hex string to bytes: %s, Error: %s", req.Proof, err.Error()))
	}

	merkleValue, err := k.VerifyToCosmosTx(ctx, proof, headerToBeVerified)
	if err != nil {
		return nil, types.ErrProcessCrossChainTx(fmt.Sprintf("VerifyToCosmostx failed, %s", err.Error()))
	}
	currentChainCrossChainId := types.GetChainID()
	if merkleValue.MakeTxParam.ToChainID != currentChainCrossChainId {
		return nil, types.ErrProcessCrossChainTx(fmt.Sprintf("toChainId is not for this chain, expect: %d, got: %d", currentChainCrossChainId, merkleValue.MakeTxParam.ToChainID))
	}

	switch merkleValue.MakeTxParam.Method {
	case "unlock":
		if err := k.ProcessUnlockTx(ctx, merkleValue, req.FromChainId); err != nil {
			return nil, err
		}
	case "registerAsset":
		if err := k.ProcessRegisterAssetTx(ctx, merkleValue); err != nil {
			return nil, err
		}
	default:
		return nil, types.ErrProcessCrossChainTx(fmt.Sprintf("unsupported cross-chain method: %s", merkleValue.MakeTxParam.Method))
	}
	res = &types.MsgProcessCrossChainTxResponse{}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return
}

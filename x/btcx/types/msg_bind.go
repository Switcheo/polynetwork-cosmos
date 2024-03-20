package types

import (
	"encoding/hex"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBind{}

// NewMsgBind returns a new MsgBind
func NewMsgBind(creator string, sourceAssetDenom string, toChainId uint64, toAssetHash []byte) *MsgBind {
	return &MsgBind{creator, sourceAssetDenom, toChainId, toAssetHash}
}

// Route implements Msg
func (msg *MsgBind) Route() string {
	return RouterKey
}

// Type implements Msg
func (msg *MsgBind) Type() string {
	return "Bind"
}

// GetSigners implements Msg
func (msg *MsgBind) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes implements Msg
func (msg *MsgBind) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements Msg
func (msg *MsgBind) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s), error: %v", msg.Creator, err)
	}
	if err := sdk.ValidateDenom(msg.SourceAssetDenom); err != nil {
		return errorsmod.Wrapf(ErrBindAssetHashType, "invalid source asset denom (%s), error: %v", msg.SourceAssetDenom, err)
	}
	if msg.ToChainId == 0 {
		return errorsmod.Wrapf(ErrInvalidChainIdType, "invalid chain id (%d)", msg.ToChainId)
	}
	if len(msg.ToAssetHash) == 0 {
		return errorsmod.Wrapf(ErrEmptyToAssetHashType, "invalid to asset hash (%s)", hex.EncodeToString(msg.ToAssetHash))
	}
	return nil
}

package types

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgLock{}

// NewMsgLock returns a new MsgLock
func NewMsgLock(fromAddress string, sourceAssetDenom string, toChainID uint64, toAddress []byte, value sdk.Int) *MsgLock {
	return &MsgLock{fromAddress, sourceAssetDenom, toChainID, toAddress, &sdk.IntProto{Int: value}}
}

// Route implements Msg
func (msg *MsgLock) Route() string {
	return RouterKey
}

// Type implements Msg
func (msg *MsgLock) Type() string {
	return "Lock"
}

// GetSigners implements Msg
func (msg *MsgLock) GetSigners() []sdk.AccAddress {
	fromAddr, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{fromAddr}
}

// GetSignBytes implements Msg
func (msg *MsgLock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements Msg
func (msg *MsgLock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s), error: %v", msg.FromAddress, err)
	}
	if err := sdk.ValidateDenom(msg.SourceAssetDenom); err != nil {
		return sdkerrors.Wrapf(ErrBindAssetHashType, "invalid source asset denom (%s), error: %v", msg.SourceAssetDenom, err)
	}
	if msg.ToChainID == 0 {
		return sdkerrors.Wrapf(ErrInvalidChainIDType, "invalid chain id (%d)", msg.ToChainID)
	}
	if len(msg.ToAddressBs) == 0 {
		return sdkerrors.Wrapf(ErrEmptyToAssetHashType, "invalid to address (%s)", hex.EncodeToString(msg.ToAddressBs))
	}
	if msg.Value.Int.IsNegative() {
		return sdkerrors.Wrapf(ErrLockType, "invalid value (%s)", msg.Value)
	}
	return nil
}

package types

import (
	"encoding/hex"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgLock{}

// NewMsgLock returns a new MsgLock
func NewMsgLock(fromAddress string, sourceAssetDenom string, toChainId uint64, toAddress []byte, value sdkmath.Int) *MsgLock {
	return &MsgLock{fromAddress, sourceAssetDenom, toChainId, toAddress, value}
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
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s), error: %v", msg.FromAddress, err)
	}
	if err := sdk.ValidateDenom(msg.SourceAssetDenom); err != nil {
		return errorsmod.Wrapf(ErrBindAssetHashType, "invalid source asset denom (%s), error: %v", msg.SourceAssetDenom, err)
	}
	if msg.ToChainId == 0 {
		return errorsmod.Wrapf(ErrInvalidChainIdType, "invalid chain id (%d)", msg.ToChainId)
	}
	if len(msg.ToAddressBytes) == 0 {
		return errorsmod.Wrapf(ErrEmptyToAssetHashType, "invalid to address (%s)", hex.EncodeToString(msg.ToAddressBytes))
	}
	if msg.Value.IsNegative() {
		return errorsmod.Wrapf(ErrLockType, "invalid value (%s)", msg.Value)
	}
	return nil
}

package types

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateDenom{}

// NewMsgCreateDenom returns a new MsgCreateDenom
func NewMsgCreateDenom(creator string, denom string, redeemScript string) *MsgCreateDenom {
	return &MsgCreateDenom{
		Creator:      creator,
		Denom:        denom,
		RedeemScript: redeemScript,
	}
}

// Route implements Msg
func (msg *MsgCreateDenom) Route() string {
	return RouterKey
}

// Type implements Msg
func (msg *MsgCreateDenom) Type() string {
	return "CreateDenom"
}

// GetSigners implements Msg
func (msg *MsgCreateDenom) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes implements Msg
func (msg *MsgCreateDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements Msg
func (msg *MsgCreateDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s), error: %v", msg.Creator, err)
	}
	if err := sdk.ValidateDenom(msg.Denom); err != nil {
		return sdkerrors.Wrapf(ErrCreateDenomType, fmt.Sprintf("invalid denom (%s), error: %v", msg.Denom, err))
	}
	if _, err := hex.DecodeString(msg.RedeemScript); err != nil {
		return sdkerrors.Wrapf(ErrCreateDenomType, fmt.Sprintf("invalid redeem script (%s), error: %v", msg.Denom, err))
	}
	return nil
}

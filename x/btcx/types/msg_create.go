package types

import (
	"encoding/hex"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreate{}

// NewMsgCreate returns a new MsgCreate
func NewMsgCreate(creator string, denom string, redeemScript string) *MsgCreate {
	return &MsgCreate{
		Creator:      creator,
		Denom:        denom,
		RedeemScript: redeemScript,
	}
}

// Route implements Msg
func (msg *MsgCreate) Route() string {
	return RouterKey
}

// Type implements Msg
func (msg *MsgCreate) Type() string {
	return "Create"
}

// GetSigners implements Msg
func (msg *MsgCreate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes implements Msg
func (msg *MsgCreate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements Msg
func (msg *MsgCreate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s), error: %v", msg.Creator, err)
	}
	if err := sdk.ValidateDenom(msg.Denom); err != nil {
		return errorsmod.Wrapf(ErrCreateDenomType, fmt.Sprintf("invalid denom (%s), error: %v", msg.Denom, err))
	}
	if _, err := hex.DecodeString(msg.RedeemScript); err != nil {
		return errorsmod.Wrapf(ErrCreateDenomType, fmt.Sprintf("invalid redeem script (%s), error: %v", msg.Denom, err))
	}
	return nil
}

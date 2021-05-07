/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewMsgCreate returns a new MsgCreate
func NewMsgCreate(creator string) *MsgCreate {
	return &MsgCreate{creator}
}

// Route implements Msg.
func (msg *MsgCreate) Route() string { return RouterKey }

// Type implements Msg.
func (msg *MsgCreate) Type() string { return "Create" }

// ValidateBasic implements Msg.
func (msg *MsgCreate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid submitter (%s), error: %v", msg.Creator, err))
	}
	return nil
}

// GetSigners implements Msg.
func (msg *MsgCreate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes implements Msg.
func (msg *MsgCreate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// MsgBind returns a new MsgBind.
func NewMsgBind(creator string, denom string, lockProxyHash []byte,
	nativeChainID uint64, nativeLockProxyHash []byte, nativeAssetHash []byte) *MsgBind {
	return &MsgBind{
		Creator:             creator,
		Denom:               denom,
		LockProxyHash:       lockProxyHash,
		NativeChainID:       nativeChainID,
		NativeLockProxyHash: nativeLockProxyHash,
		NativeAssetHash:     nativeAssetHash,
	}
}

// Route implements Msg.
func (msg *MsgBind) Route() string { return RouterKey }

// Type implements Msg.
func (msg *MsgBind) Type() string { return "Bind" }

// ValidateBasic implements Msg.
func (msg *MsgBind) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid submitter (%s), error: %v", msg.Creator, err))
	}
	err = sdk.ValidateDenom(msg.Denom)
	if err != nil {
		return ErrMsgBind(fmt.Sprintf("Invalid Denom (%s). Error: %s", msg.Denom, err))
	}
	if msg.NativeChainID == 0 {
		return ErrInvalidChainID(msg.NativeChainID)
	}
	if len(msg.NativeLockProxyHash) == 0 {
		return ErrMsgBind("Empty NativeLockProxyHash")
	}
	if len(msg.NativeAssetHash) == 0 {
		return ErrMsgBind("Empty NativeAssetHash")
	}

	return nil
}

// GetSigners implements Msg.
func (msg *MsgBind) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

// GetSignBytes implements Msg.
func (msg *MsgBind) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// NewMsgLock returns a new MsgLock.
func NewMsgLock(lockProxyHash []byte, fromAddress string, denom string, toChainID uint64,
	toChainProxyHash []byte, toChainAssetHash []byte, toAddressBytes []byte,
	value sdk.Int, deductFeeInLock bool, feeAmount sdk.Int, feeAddress string) *MsgLock {
	return &MsgLock{
		LockProxyHash:    lockProxyHash,
		FromAddress:      fromAddress,
		Denom:            denom,
		ToChainID:        toChainID,
		ToChainProxyHash: toChainProxyHash,
		ToChainAssetHash: toChainAssetHash,
		ToAddressBytes:   toAddressBytes,
		Value:            &sdk.IntProto{Int: value},
		DeductFeeInLock:  deductFeeInLock,
		FeeAmount:        &sdk.IntProto{Int: feeAmount},
		FeeAddress:       feeAddress,
	}
}

// Route implements Msg.
func (msg *MsgLock) Route() string { return RouterKey }

// Type implements Msg.
func (msg *MsgLock) Type() string { return "Lock" }

// ValidateBasic implements Msg.
func (msg *MsgLock) ValidateBasic() error {
	if len(msg.LockProxyHash) == 0 {
		return ErrMsgLock("Empty LockProxyHash")
	}
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid from address (%s), error: %v", msg.FromAddress, err))
	}
	if msg.Denom == "" {
		return ErrMsgLock("Empty Denom")
	}
	if msg.ToChainID <= 0 {
		return ErrInvalidChainID(msg.ToChainID)
	}
	if len(msg.ToChainProxyHash) == 0 {
		return ErrMsgLock("Empty ToChainProxyHash")
	}
	if len(msg.ToChainAssetHash) == 0 {
		return ErrMsgLock("Empty ToChainAssetHash")
	}
	if len(msg.ToAddressBytes) == 0 {
		return ErrMsgLock("Empty ToAssetHash")
	}
	if !msg.Value.Int.IsPositive() {
		return ErrMsgLock(fmt.Sprintf("value (%s) should be positive", msg.Value.String()))
	}
	if msg.DeductFeeInLock {
		if !msg.FeeAmount.Int.IsPositive() {
			return ErrMsgLock(fmt.Sprintf("feeAmount (%s) should be negative when deducting fee in lock", msg.FeeAmount.String()))
		}
		if msg.FeeAmount.Int.LTE(msg.Value.Int) {
			return ErrMsgLock(fmt.Sprintf("feeAmount (%s) should be more than value (%s)", msg.FeeAmount.String(), msg.Value.String()))
		}
		_, err = sdk.AccAddressFromBech32(msg.FeeAddress)
		if err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid fee address (%s), error: %v", msg.FeeAddress, err))
		}
	}
	return nil
}

// GetSignBytes implements Msg.
func (msg *MsgLock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg.
func (msg *MsgLock) GetSigners() []sdk.AccAddress {
	fromAddr, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{fromAddr}
}

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

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
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
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid submitter (%s), error: %v", msg.Creator, err))
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
	nativeChainId uint64, nativeLockProxyHash []byte, nativeAssetHash []byte) *MsgBind {
	return &MsgBind{
		Creator:             creator,
		Denom:               denom,
		LockProxyHash:       lockProxyHash,
		NativeChainId:       nativeChainId,
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
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid submitter (%s), error: %v", msg.Creator, err))
	}
	err = sdk.ValidateDenom(msg.Denom)
	if err != nil {
		return ErrMsgBind(fmt.Sprintf("Invalid Denom (%s). Error: %s", msg.Denom, err))
	}
	if msg.NativeChainId == 0 {
		return ErrInvalidChainId(msg.NativeChainId)
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
func NewMsgLock(denom string,
	fromLockProxy, fromAssetId []byte, fromAddress string,
	toChainId uint64, toLockProxy, toAssetId, toAddress []byte,
	amount sdkmath.Int, deductFeeInLock bool, feeAmount sdkmath.Int, feeAddress string) *MsgLock {
	return &MsgLock{
		Denom:           denom,
		FromLockProxy:   fromLockProxy,
		FromAssetId:     fromAssetId,
		FromAddress:     fromAddress,
		ToChainId:       toChainId,
		ToLockProxy:     toLockProxy,
		ToAssetId:       toAssetId,
		ToAddress:       toAddress,
		Amount:          amount,
		DeductFeeInLock: deductFeeInLock,
		FeeAmount:       feeAmount,
		FeeAddress:      feeAddress,
	}
}

// Route implements Msg.
func (msg *MsgLock) Route() string { return RouterKey }

// Type implements Msg.
func (msg *MsgLock) Type() string { return "Lock" }

// ValidateBasic implements Msg.
func (msg *MsgLock) ValidateBasic() error {
	if msg.Denom == "" {
		return ErrMsgLock("Empty Denom")
	}
	if len(msg.FromLockProxy) == 0 {
		return ErrMsgLock("Empty FromLockProxy")
	}
	if len(msg.FromAssetId) == 0 {
		return ErrMsgLock("Empty FromAssetId")
	}
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid from address (%s), error: %v", msg.FromAddress, err))
	}
	if msg.ToChainId <= 0 {
		return ErrInvalidChainId(msg.ToChainId)
	}
	if len(msg.ToLockProxy) == 0 {
		return ErrMsgLock("Empty ToLockProxy")
	}
	if len(msg.ToAssetId) == 0 {
		return ErrMsgLock("Empty ToAssetId")
	}
	if len(msg.ToAddress) == 0 {
		return ErrMsgLock("Empty ToAddress")
	}
	if !msg.Amount.IsPositive() {
		return ErrMsgLock(fmt.Sprintf("amount (%s) should be positive", msg.Amount.String()))
	}
	if msg.DeductFeeInLock {
		if !msg.FeeAmount.IsPositive() {
			return ErrMsgLock(fmt.Sprintf("feeAmount (%s) should be negative when deducting fee in lock", msg.FeeAmount.String()))
		}
		if msg.FeeAmount.LTE(msg.Amount) {
			return ErrMsgLock(fmt.Sprintf("feeAmount (%s) should be more than amount (%s)", msg.FeeAmount.String(), msg.Amount.String()))
		}
		_, err = sdk.AccAddressFromBech32(msg.FeeAddress)
		if err != nil {
			return errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid fee address (%s), error: %v", msg.FeeAddress, err))
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

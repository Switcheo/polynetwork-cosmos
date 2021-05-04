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

// Governance message types and routes
const (
	TypeMsgSyncGenesis = "sync_genesis"
	TypeMsgSyncHeaders = "sync_headers"
)

// NewMsgSyncGenesis - construct arbitrary multi-in, multi-out send msg.
func NewMsgSyncGenesis(syncer string, genesisHeader string) *MsgSyncGenesis {
	return &MsgSyncGenesis{Syncer: syncer, GenesisHeader: genesisHeader}
}

// Route Implements Msg.
func (msg *MsgSyncGenesis) Route() string { return RouterKey }

// Type Implements Msg.
func (msg *MsgSyncGenesis) Type() string { return TypeMsgSyncGenesis }

// ValidateBasic Implements Msg.
func (msg *MsgSyncGenesis) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Syncer)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid syncer (%s), error: %v", msg.Syncer, err))
	}
	if len(msg.GenesisHeader) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "GenesisHeader should not be empty")
	}
	return nil
}

// GetSigners Implements Msg.
func (msg *MsgSyncGenesis) GetSigners() []sdk.AccAddress {
	syncer, err := sdk.AccAddressFromBech32(msg.Syncer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{syncer}
}

// GetSignBytes Implements Msg.
func (msg *MsgSyncGenesis) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// NewMsgMultiSend - construct arbitrary multi-in, multi-out send msg.
func NewMsgSyncHeaders(syncer string, headers []string) *MsgSyncHeaders {
	return &MsgSyncHeaders{Syncer: syncer, Headers: headers}
}

// Route Implements Msg
func (msg *MsgSyncHeaders) Route() string { return RouterKey }

// Type Implements Msg
func (msg *MsgSyncHeaders) Type() string { return TypeMsgSyncHeaders }

// ValidateBasic Implements Msg.
func (msg *MsgSyncHeaders) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Syncer)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid syncer (%s), error: %v", msg.Syncer, err))
	}
	if len(msg.Headers) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Headers should not be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg *MsgSyncHeaders) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg *MsgSyncHeaders) GetSigners() []sdk.AccAddress {
	syncer, err := sdk.AccAddressFromBech32(msg.Syncer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{syncer}
}

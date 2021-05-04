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

package keeper

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Switcheo/polynetwork-cosmos/x/btcx/types"
)

const (
	// default paramspace for params keeper
	DefaultParamspace = types.ModuleName
)

var (
	ChainIDToAssetHashPrefix       = []byte{0x01}
	CreatorDenomToScriptHashPrefix = []byte{0x02}
	ScriptHashToRedeemScriptPrefix = []byte{0x03}
	DenomToCreatorPrefix           = []byte{0x04}
	BindAssetHashPrefix            = []byte{0x05}
	DenomToRedeemScriptKey         = []byte{0x06}
)

// TODO: delete this method
func GetScriptHashAndChainIDToAssetHashKey(scriptHash []byte, chainID uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, chainID)
	return append(append(ChainIDToAssetHashPrefix, scriptHash...), b...)
}

func GetScriptHashToRedeemScript(scriptHashKeyBs []byte) []byte {
	return append(ScriptHashToRedeemScriptPrefix, scriptHashKeyBs...)
}

func GetCreatorDenomToScriptHashKey(creator sdk.AccAddress, denom string) []byte {
	return append(CreatorDenomToScriptHashPrefix, []byte(denom)...)
}

func GetDenomToCreatorKey(denom string) []byte {
	return append(DenomToCreatorPrefix, []byte(denom)...)
}

func GetDenomToRedeemScriptKey(denom string) []byte {
	return append(DenomToRedeemScriptKey, []byte(denom)...)
}

func GetBindAssetHashKey(sourceDenomHash []byte, chainID uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, chainID)
	return append(append(BindAssetHashPrefix, sourceDenomHash...), b...)
}

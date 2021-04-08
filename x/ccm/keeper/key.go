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

	"github.com/Switcheo/polynetwork-cosmos/x/ccm/types"
)

const (
	// default paramspace for params keeper
	DefaultParamspace = types.ModuleName
)

var (
	CrossChainTxDetailPrefix = []byte{0x01}
	CrossChainDoneTxPrefix   = []byte{0x02}
	DenomToCreatorPrefix     = []byte{0x03}

	CrossChainIdKey = []byte("crosschainid")
)

func GetCrossChainTxKey(crossChainTxSum []byte) []byte {
	return append(CrossChainTxDetailPrefix, crossChainTxSum...)
}

func GetDoneTxKey(fromChainId uint64, crossChainid []byte) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, fromChainId)
	return append(append(CrossChainDoneTxPrefix, b...), crossChainid...)
}

func GetDenomToCreatorKey(denom string) []byte {
	return append(DenomToCreatorPrefix, []byte(denom)...)
}

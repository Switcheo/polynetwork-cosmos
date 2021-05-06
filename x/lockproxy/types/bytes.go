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
	"encoding/hex"
	"fmt"
	"math/big"
)

func PadFixedBytes(bigint *big.Int, intBsLen int) ([]byte, error) {
	ret := make([]byte, intBsLen)
	if bigint.Cmp(big.NewInt(0)) < 0 {
		return nil, fmt.Errorf("PadFixedBytes doesnot support negative big.Int, but got:%s", bigint.String())
	}
	bigBs := bigint.Bytes()
	if len(bigBs) > intBsLen || (len(bigBs) == intBsLen && bigBs[0]>>7 == 1) {
		return nil, fmt.Errorf("PadFixedBytes only support maximum 2**255-1 big.Int, but got:%s", bigint.String())
	}
	copy(ret[:len(bigBs)], make([]byte, len(bigBs)))
	copy(ret[intBsLen-len(bigBs):], bigBs)
	return ToArrayReverse(ret), nil
}

func UnpadFixedBytes(paddedBs []byte, intBsLen int) (*big.Int, error) {
	if len(paddedBs) != intBsLen {
		return nil, fmt.Errorf("UnpadFixedBytes only support 32 bytes value, but got:%s", hex.EncodeToString(paddedBs))
	}
	nonZeroPos := intBsLen - 1
	for i := nonZeroPos; i >= 0; i-- {
		p := paddedBs[i]
		if p != 0x0 {
			nonZeroPos = i
			break
		}
	}
	if nonZeroPos == intBsLen-1 && paddedBs[intBsLen-1]>>7 == 1 {
		return nil, fmt.Errorf("UnpadFixedBytes only support 32 bytes nonnegative value, but got:%s", hex.EncodeToString(paddedBs))
	}

	return big.NewInt(0).SetBytes(ToArrayReverse(paddedBs[:nonZeroPos+1])), nil
}

func ToArrayReverse(arr []byte) []byte {
	l := len(arr)
	x := make([]byte, 0)
	for i := l - 1; i >= 0; i-- {
		x = append(x, arr[i])
	}
	return x
}

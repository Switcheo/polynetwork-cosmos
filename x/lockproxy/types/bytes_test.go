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
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

// coming from "github.com/ethereum/go-ethereum/common/math"
func Test_PadFixedBytes(t *testing.T) {
	U255Minus1 := big.NewInt(0).Sub(big.NewInt(0).Exp(big.NewInt(2), big.NewInt(255), nil), big.NewInt(1))
	U256Minus1 := big.NewInt(0).Sub(big.NewInt(0).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1))
	U300Minus1 := big.NewInt(0).Sub(big.NewInt(0).Exp(big.NewInt(2), big.NewInt(300), nil), big.NewInt(1))

	testCases := []struct {
		numStr    string
		supported bool
	}{
		{"-1000", false},
		{"0", true},
		{"100", true},
		{"200", true},
		{"266", true},
		{"10000000000000", true},
		{"1234567890987654321", true},
		{U255Minus1.String(), true},
		{U255Minus1.Add(U255Minus1, big.NewInt(1)).String(), false},
		{U255Minus1.Add(U255Minus1, big.NewInt(1000000)).String(), false},
		{U255Minus1.Add(U256Minus1, big.NewInt(1)).String(), false},
		{U255Minus1.Add(U256Minus1, big.NewInt(-1)).String(), false},
		{U300Minus1.String(), false},
	}
	for _, testCase := range testCases {
		num, ok := new(big.Int).SetString(testCase.numStr, 10)
		assert.True(t, ok)
		numBs, err := PadFixedBytes(num, 32)
		if testCase.supported {
			assert.Nil(t, err)
			numRecover, err := UnpadFixedBytes(numBs, 32)
			assert.Nil(t, err)
			assert.Equal(t, 0, num.Cmp(numRecover))
		} else {
			assert.NotNil(t, err)
		}
	}
}

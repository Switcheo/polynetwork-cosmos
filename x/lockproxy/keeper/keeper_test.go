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

package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/Switcheo/polynetwork-cosmos/simapp"
	"github.com/Switcheo/polynetwork-cosmos/testutil/testdata"
)

func createTestApp(isCheckTx bool) (*simapp.SimApp, sdk.Context) {
	app := simapp.Setup(isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	return app, ctx
}

func Test_lockproxy_Create(t *testing.T) {
	app, ctx := createTestApp(true)

	_, _, addr1 := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()

	testCases := []struct {
		address       sdk.AccAddress
		expectSucceed bool
	}{
		{addr1, true},
		{addr1, false},
		{addr2, true},
		{addr2, false},
		{addr2, false},
	}
	for _, testCase := range testCases {
		err := app.LockproxyKeeper.CreateLockProxy(ctx, testCase.address)
		if testCase.expectSucceed {
			require.Nil(t, err)
		} else {
			require.Error(t, err)
		}
	}
}

func Test_lockproxy_Bind(t *testing.T) {
	app, ctx := createTestApp(true)

	_, _, addr1 := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()
	err := app.LockproxyKeeper.CreateLockProxy(ctx, addr1)
	require.Nil(t, err)
	err = app.LockproxyKeeper.CreateLockProxy(ctx, addr2)
	require.Nil(t, err)

	lp1 := addr1.Bytes()
	lp2 := addr2.Bytes()

	hash := []byte{0x01}
	hash2 := []byte{0x02}

	testCases := []struct {
		address             sdk.AccAddress
		denom               string
		lockProxy           []byte
		nativeChainId       uint64
		nativeLockProxyHash []byte
		nativeAssetHash     []byte
		expectSucceed       bool
		description         string
	}{
		{addr1, "coin1", lp1, 1, hash, hash, true, "valid denom and proxy"},
		{addr1, "coin1", lp1, 1, hash2, hash2, false, "repeated denom"},
		{addr1, "coin2", lp1, 1, hash, hash, true, "valid new denom"},
		{addr1, "coin3", sdk.AccAddress([]byte("invalidLockProxy")), 1, hash, hash, false, "invalid lock proxy"},
		{addr2, "coin1", lp1, 1, hash, hash, false, "repeated denom, other user"},
		{addr2, "coin4", lp1, 1, hash, hash, true, "mismatched lock proxy is ok"},
		{addr2, "coin5", lp2, 1, hash, hash, true, "valid denom and proxy 2"},
	}

	for _, testCase := range testCases {
		err = app.LockproxyKeeper.CreateCoinAndDelegateToProxy(ctx,
			testCase.address, testCase.denom,
			testCase.lockProxy, testCase.nativeChainId,
			testCase.nativeLockProxyHash, testCase.nativeAssetHash,
		)

		if testCase.expectSucceed {
			require.Nil(t, err, testCase.description)
		} else {
			require.Error(t, err, testCase.description)
		}
	}
}

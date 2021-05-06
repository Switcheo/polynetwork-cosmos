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
	gocontext "context"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/Switcheo/polynetwork-cosmos/simapp"
	"github.com/Switcheo/polynetwork-cosmos/testutil/testdata"
	"github.com/Switcheo/polynetwork-cosmos/x/btcx/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	supply "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

// returns context and an app with updated btcx keeper
func createTestApp(isCheckTx bool) (*simapp.SimApp, sdk.Context, types.QueryClient) {
	app := simapp.Setup(isCheckTx)
	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.BtcxKeeper)
	queryClient := types.NewQueryClient(queryHelper)
	return app, ctx, queryClient
}

func btcx_initSupply(t *testing.T, app *simapp.SimApp, ctx sdk.Context) sdk.Coins {
	coinsStr := "1000000000btca,1000000000000000000btcb"
	initTokens, err := sdk.ParseCoinsNormalized(coinsStr)
	require.Equal(t, nil, err, "Parse Coins error should be nil")
	totalSupply := sdk.NewCoins(initTokens...)
	app.BankKeeper.SetSupply(ctx, supply.NewSupply(totalSupply))

	total := app.BankKeeper.GetSupply(ctx).GetTotal()

	require.Equal(t, totalSupply, total, "supply should be initTokens")
	return total
}

func Test_btcx_MsgCreate(t *testing.T) {
	app, ctx, queryClient := createTestApp(true)
	btcx_initSupply(t, app, ctx)
	_, _, addr1 := testdata.KeyTestPubAddr()
	_, _, addr2 := testdata.KeyTestPubAddr()

	testCases := []struct {
		address       sdk.AccAddress
		denom         string
		redeemScript  string
		expectSucceed bool
	}{
		{addr1, "btcx1", "1234", true},
		{addr1, "btcx2", "123456", true},
		{addr1, "btcx2", "12345678", false},
		{addr2, "btcx3", "12345678", true}, // which is valid depends on which denom is bonded in poly chain, btcx2 or btcx3
		{addr2, "btcx4", "123456789", false},
		{addr2, "btcx5", "1234567890", true},
	}
	for _, testCase := range testCases {
		acc := app.AccountKeeper.NewAccountWithAddress(ctx, testCase.address)
		creator := acc.GetAddress()
		require.Equal(t, testCase.address, creator, fmt.Sprintf("expect: %s, got: %s", testCase.address, creator))
		err := app.BtcxKeeper.CreateAsset(ctx, creator, testCase.denom, testCase.redeemScript)
		if testCase.expectSucceed {
			require.Nil(t, err)
		} else {
			require.Error(t, err)
		}
	}

	for _, testCase := range testCases {
		if testCase.expectSucceed {
			res, err := queryClient.DenomInfo(gocontext.Background(), &types.QueryGetDenomInfoRequest{Denom: testCase.denom})
			require.NoError(t, err)
			denomInfo := res.DenomInfo
			require.Equal(t, testCase.address.String(), denomInfo.Creator, "creator is not correct")
			require.Equal(t, testCase.denom, denomInfo.Denom, "denom is not correct")
			require.True(t, denomInfo.TotalSupply.Int.Equal(sdk.ZeroInt()), "total supply is not correct")
			require.Equal(t, testCase.redeemScript, denomInfo.RedeemScript, "redeem script is not correct")
		}
	}
}

func Test_btcx_MsgBind(t *testing.T) {
	app, ctx, queryClient := createTestApp(true)
	btcx_initSupply(t, app, ctx)

	_, _, addr := testdata.KeyTestPubAddr()
	acc := app.AccountKeeper.NewAccountWithAddress(ctx, addr)
	creator := acc.GetAddress()
	redeemScript := "12345678"
	require.Equal(t, addr, creator, fmt.Sprintf("expect: %s, got: %s", addr, creator))
	err := app.BtcxKeeper.CreateAsset(ctx, creator, "btcx1", redeemScript)
	require.Nil(t, err)
	invalidCreator := app.AccountKeeper.NewAccountWithAddress(ctx, sdk.AccAddress([]byte("invalidCreator"))).GetAddress()
	testCases := []struct {
		creator       sdk.AccAddress
		denom         string
		toChainId     uint64
		toAssetHash   []byte
		expectSucceed bool
	}{
		{creator, "btcx1", 2, []byte{1, 2, 3, 4}, true},
		{invalidCreator, "btcx1", 2, []byte{1, 2, 3, 5}, false},
		{creator, "btcx1", 3, []byte{1, 2, 3, 6}, true},
	}

	for _, testCase := range testCases {
		err := app.BtcxKeeper.BindAsset(ctx, testCase.creator, testCase.denom, testCase.toChainId, testCase.toAssetHash)
		if testCase.expectSucceed {
			require.Nil(t, err)
		} else {
			require.Error(t, err)
		}
	}
	for _, testCase := range testCases {
		if testCase.expectSucceed {
			res, err := queryClient.DenomCrossChainInfo(gocontext.Background(),
				&types.QueryGetDenomCrossChainInfoRequest{Denom: testCase.denom, ChainID: testCase.toChainId})

			require.NoError(t, err)
			denomInfo := res.DenomInfo
			require.Equal(t, testCase.creator.String(), denomInfo.Creator, "creator is not correct")
			require.Equal(t, testCase.denom, denomInfo.Denom, "denom is not correct")
			require.Equal(t, hex.EncodeToString([]byte(testCase.denom)), denomInfo.AssetHash, "denom is not correct")
			require.True(t, denomInfo.TotalSupply.Int.Equal(sdk.ZeroInt()), "total supply is not correct")
			require.Equal(t, redeemScript, denomInfo.RedeemScript, "redeem script is not correct")
			require.Equal(t, testCase.toChainId, res.ToChainID, "toChainID is not correct")
			require.Equal(t, hex.EncodeToString(testCase.toAssetHash), res.ToAssetHash, "toAssetHash is not correct")
		}
	}
}

func Test_btcx_MsgLock(t *testing.T) {
	app, ctx, _ := createTestApp(true)
	btcx_initSupply(t, app, ctx)

	total := app.BankKeeper.GetSupply(ctx).GetTotal()
	btcx1CoinStr := "100btcx1"
	btcx1Coin, err := sdk.ParseCoinNormalized(btcx1CoinStr)
	require.Nil(t, err)

	denom := btcx1Coin.Denom
	_, _, creator := testdata.KeyTestPubAddr()

	err = app.BtcxKeeper.CreateAsset(ctx, creator, denom, "12345678")
	require.Nil(t, err)

	app.BankKeeper.SetSupply(ctx, supply.NewSupply(total.Add(btcx1Coin)))
	require.Equal(t, btcx1Coin.Amount, app.BankKeeper.GetSupply(ctx).GetTotal().AmountOf("btcx1"), "btcx1 amount should be 100")

	err = app.BankKeeper.AddCoins(ctx, creator, sdk.Coins{btcx1Coin})
	balance := app.BankKeeper.GetAllBalances(ctx, creator)
	require.Nil(t, err)
	require.Equal(t, sdk.Coins{btcx1Coin}, balance, "create balance is not equal to 100btcx1")

	err = app.BtcxKeeper.BindAsset(ctx, creator, denom, 2, []byte{1, 2, 3, 4})
	require.Nil(t, err)
	err = app.BtcxKeeper.BindAsset(ctx, creator, denom, 3, []byte{1, 2, 3, 5})
	require.Nil(t, err)

	testCases := []struct {
		from          sdk.AccAddress
		denom         string
		toChainId     uint64
		toAddr        []byte
		amount        sdk.Int
		expectSucceed bool
	}{
		{creator, denom, 2, []byte{1, 2}, sdk.NewInt(1), true},
		{creator, denom, 2, []byte{1, 3}, sdk.NewInt(100), false},
		{creator, denom, 3, []byte{1, 4}, sdk.NewInt(2), true},
		{creator, denom, 4, []byte{1, 5}, sdk.NewInt(3), false},
	}

	for _, testCase := range testCases {
		err := app.BtcxKeeper.LockAsset(ctx, testCase.from, testCase.denom, testCase.toChainId, testCase.toAddr, testCase.amount)
		if testCase.expectSucceed {
			require.Nil(t, err)
		} else {
			require.Error(t, err)
		}
	}

	balance = app.BankKeeper.GetAllBalances(ctx, creator)
	require.Equal(t, "97btcx1", balance.String(), "balnace of creator is not balanced")
}

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
	"testing"

	polycommon "github.com/polynetwork/poly/common"
	polytype "github.com/polynetwork/poly/core/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/Switcheo/polynetwork-cosmos/x/headersync/types"
)

func TestN_headersync_Querier(t *testing.T) {
	app, ctx, queryClient := createTestApp(true)

	h0s, _ := hex.DecodeString(header0)
	h0 := new(polytype.Header)
	err := h0.Deserialization(polycommon.NewZeroCopySource(h0s))
	assert.Nil(t, err)

	err = app.HeadersyncKeeper.SyncGenesisHeader(ctx, header0)
	assert.Nil(t, err, "Sync genesis header fail")

	// query synced Poly Chain consensus peer
	res, err := queryClient.ConsensusPeers(gocontext.Background(), &types.QueryGetConsensusPeersRequest{ChainId: h0.ChainID})
	require.NoError(t, err)

	consensusPeersBs, err := ExtractChainConfig(h0)
	assert.Nil(t, err)

	sink := polycommon.NewZeroCopySink(nil)
	res.ConsensusPeers.Serialize(sink)
	resBs := sink.Bytes()

	require.Equal(t, consensusPeersBs, resBs, "Synced consensus 0 is not equal to the querier result")
}

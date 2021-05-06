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
	"testing"

	"github.com/polynetwork/poly/common"
	"github.com/stretchr/testify/assert"
)

func TestPeer_Serialization(t *testing.T) {
	paramSerialize := new(Peer)
	paramSerialize.Index = 1
	paramSerialize.PubKey = "abcdefg"
	sink := common.NewZeroCopySink(nil)
	paramSerialize.Serialize(sink)

	paramDeserialize := new(Peer)
	err := paramDeserialize.Deserialize(common.NewZeroCopySource(sink.Bytes()))
	assert.Nil(t, err)
	assert.Equal(t, paramDeserialize, paramSerialize)
}

func TestConsensusPeers_Serialization(t *testing.T) {
	paramSerialize := new(ConsensusPeers)
	paramSerialize.Height = 1
	paramSerialize.ChainID = 0
	peer1 := &Peer{Index: 1, PubKey: "abcd"}
	peer2 := &Peer{Index: 2, PubKey: "efgh"}
	paramSerialize.Peers = make(map[string]*Peer)
	paramSerialize.Peers[peer1.PubKey] = peer1
	paramSerialize.Peers[peer2.PubKey] = peer2
	sink := common.NewZeroCopySink(nil)
	paramSerialize.Serialize(sink)

	paramDeserialize := new(ConsensusPeers)
	err := paramDeserialize.Deserialize(common.NewZeroCopySource(sink.Bytes()))
	assert.Nil(t, err)
	assert.Equal(t, paramDeserialize, paramSerialize)

	s1 := common.NewZeroCopySink(nil)
	p := *paramSerialize
	p.Serialize(s1)
	p2 := new(ConsensusPeers)
	err = p2.Deserialize(common.NewZeroCopySource(s1.Bytes()))
	assert.Nil(t, err)
	assert.Equal(t, p, *p2)
}

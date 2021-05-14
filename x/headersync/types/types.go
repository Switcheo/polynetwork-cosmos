package types

import (
	"fmt"
	"sort"

	polycommon "github.com/polynetwork/poly/common"
)

func (this *Peer) Serialize(sink *polycommon.ZeroCopySink) {
	sink.WriteUint32(this.Index)
	sink.WriteVarBytes([]byte(this.Pubkey))
}

func (this *Peer) Deserialize(source *polycommon.ZeroCopySource) error {
	index, eof := source.NextUint32()
	if eof {
		return fmt.Errorf("utils.DecodeVarUint, Deserialize index error")
	}
	Pubkey, eof := source.NextString()
	if eof {
		return fmt.Errorf("utils.DecodeString, Deserialize Pubkey error")
	}
	this.Index = uint32(index)
	this.Pubkey = Pubkey
	return nil
}

func (this *ConsensusPeers) Serialize(sink *polycommon.ZeroCopySink) {
	sink.WriteUint64(this.ChainId)
	sink.WriteUint32(this.Height)
	sink.WriteVarUint(uint64(len(this.Peers)))
	var peerList []*Peer
	for _, v := range this.Peers {
		peerList = append(peerList, v)
	}
	sort.SliceStable(peerList, func(i, j int) bool {
		return peerList[i].Pubkey > peerList[j].Pubkey
	})
	for _, v := range peerList {
		v.Serialize(sink)
	}
}

func (this *ConsensusPeers) Deserialize(source *polycommon.ZeroCopySource) error {
	chainID, eof := source.NextUint64()
	if eof {
		return fmt.Errorf("utils.DecodeVarUint, Deserialize chainID error")
	}
	height, eof := source.NextUint32()
	if eof {
		return fmt.Errorf("utils.DecodeVarUint, Deserialize height error")
	}
	n, eof := source.NextVarUint()
	if eof {
		return fmt.Errorf("utils.DecodeVarUint, Deserialize HeightList length error")
	}
	Peers := make(map[string]*Peer)
	for i := 0; uint64(i) < n; i++ {
		peer := new(Peer)
		if err := peer.Deserialize(source); err != nil {
			return fmt.Errorf("Deserialize peer error: %v", err)
		}
		Peers[peer.Pubkey] = peer
	}
	this.ChainId = chainID
	this.Height = uint32(height)
	this.Peers = Peers
	return nil
}

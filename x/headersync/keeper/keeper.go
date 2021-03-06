package keeper

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/Switcheo/polynetwork-cosmos/x/headersync/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	polycommon "github.com/polynetwork/poly/common"
	vconfig "github.com/polynetwork/poly/consensus/vbft/config"
	polysig "github.com/polynetwork/poly/core/signature"
	polytype "github.com/polynetwork/poly/core/types"
	"github.com/polynetwork/poly/merkle"
	"github.com/tendermint/tendermint/libs/log"
	// this line is used by starport scaffolding # ibc/keeper/import
)

type (
	Keeper struct {
		cdc      codec.Codec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		// this line is used by starport scaffolding # ibc/keeper/attribute
	}
)

var (
	ConsensusPeerPrefix = []byte{0x01}
	// To help store the header hash at height where the poly chain switch epoch consensus public keys
	KeyHeaderHashPrefix = []byte{0x02}
)

func NewKeeper(
	cdc codec.Codec,
	storeKey,
	memKey sdk.StoreKey,
	// this line is used by starport scaffolding # ibc/keeper/parameter
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		// this line is used by starport scaffolding # ibc/keeper/return
	}
}

// StoreIterator returns the iterator for the store
func (k Keeper) StoreIterator(ctx sdk.Context, prefix []byte) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, prefix)
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SyncGenesisHeader(ctx sdk.Context, genesisHeaderStr string) error {
	genesisHeader := &polytype.Header{}

	genesisHeaderBytes, err := hex.DecodeString(genesisHeaderStr)
	if err != nil {
		return types.ErrSyncGenesisHeader(fmt.Sprintf("hex.DecodeString error: %s", err.Error()))
	}
	source := polycommon.NewZeroCopySource(genesisHeaderBytes)
	if err := genesisHeader.Deserialization(source); err != nil {
		return types.ErrDeserializeHeader(err)
	}
	if consensusPeer, _ := k.GetConsensusPeers(ctx, genesisHeader.ChainID); consensusPeer != nil {
		return types.ErrSyncGenesisHeader(fmt.Sprintf("Genesis Header already synced, ConsensusPeers exists: %s", consensusPeer.String()))
	}
	if err := k.UpdateConsensusPeer(ctx, genesisHeader); err != nil {
		return err
	}
	// Make sure the header contains poly.NewChainConfig info
	if _, err := k.GetConsensusPeers(ctx, genesisHeader.ChainID); err != nil {
		return types.ErrSyncGenesisHeader(fmt.Sprintf("After UpdteConsensusPeer, Get Consensus Peers error: %v", err))
	}
	return nil
}

func (k Keeper) SyncBlockHeaders(ctx sdk.Context, headerStrs []string) error {
	for _, headerStr := range headerStrs {
		header := &polytype.Header{}
		headerBs, err := hex.DecodeString(headerStr)
		if err != nil {
			return types.ErrSyncBlockHeader("Decode header string to bytes", 0, 0, err)
		}
		source := polycommon.NewZeroCopySource(headerBs)
		if err := header.Deserialization(source); err != nil {
			return types.ErrDeserializeHeader(err)
		}
		if err := k.ProcessHeader(ctx, header, nil, nil); err != nil {
			return types.ErrSyncBlockHeader("ProcessHeader", header.ChainID, header.Height, err)
		}
	}
	return nil
}

func (k Keeper) ProcessHeader(ctx sdk.Context, header *polytype.Header, headerProof []byte, curHeader *polytype.Header) error {
	// header to be checked if containing valid NewChainConfig
	var cpHeader *polytype.Header
	if curHeader == nil || headerProof == nil {
		if err := k.VerifyHeaderSig(ctx, header); err != nil {
			if err := k.VerifyHeaderByKeyHeaderHash(ctx, header); err == nil {
				return nil
			}
			return err
		}
		cpHeader = header
	} else {
		if err := k.VerifyHistoricalHeader(ctx, header, headerProof, curHeader); err != nil {
			return err
		}
		cpHeader = curHeader
	}

	if err := k.UpdateConsensusPeer(ctx, cpHeader); err != nil {
		return err
	}
	return nil
}

func (k Keeper) VerifyHeaderSig(ctx sdk.Context, header *polytype.Header) error {
	consensusPeer, err := k.GetConsensusPeers(ctx, header.ChainID)
	if err != nil {
		return types.ErrSyncBlockHeader("GetConsensusPeer", header.ChainID, header.Height, err)
	}
	if header.Height <= consensusPeer.Height {
		return types.ErrSyncBlockHeader("Compare height", header.ChainID, header.Height,
			fmt.Errorf("Stored consensus header.Height: %d, trying to sync height:%d", consensusPeer.Height, header.Height))
	}

	if len(header.Bookkeepers)*3 < len(consensusPeer.Peers)*2 {
		return types.ErrBookKeeperNum(len(header.Bookkeepers), len(consensusPeer.Peers))
	}
	for _, bookkeeper := range header.Bookkeepers {
		pubkey := vconfig.PubkeyID(bookkeeper)
		_, present := consensusPeer.Peers[pubkey]
		if !present {
			return types.ErrInvalidPublicKey(pubkey)
		}
	}
	hash := header.Hash()
	if e := polysig.VerifyMultiSignature(hash[:], header.Bookkeepers, len(header.Bookkeepers), header.SigData); e != nil {
		return types.ErrVerifyMultiSigFail(err, header.Height)
	}
	return nil
}
func (k Keeper) VerifyHeaderByKeyHeaderHash(ctx sdk.Context, header *polytype.Header) error {
	headerHash := header.Hash()
	keyHeaderHash, err := k.GetKeyHeaderHash(ctx, header.ChainID)
	if err != nil {
		return fmt.Errorf("VerifyHeaderByKeyHeaderHash, GetKeyHeaderHash Error: %s", err.Error())
	}
	if headerHash == *keyHeaderHash {
		return nil
	}
	return fmt.Errorf("VerifyHeaderByKeyHeaderHash, not equal, expect: %s, got: %s", keyHeaderHash.ToArray(), headerHash.ToArray())
}

func (k Keeper) UpdateConsensusPeer(ctx sdk.Context, header *polytype.Header) error {

	blkInfo := &vconfig.VbftBlockInfo{}
	if err := json.Unmarshal(header.ConsensusPayload, blkInfo); err != nil {
		return types.ErrUnmarshalSpecificTypeFail(blkInfo, err)
	}
	if blkInfo.NewChainConfig != nil {
		consensusPeers := &types.ConsensusPeers{
			ChainId: header.ChainID,
			Height:  header.Height,
			Peers:   make(map[string]*types.Peer),
		}
		for _, p := range blkInfo.NewChainConfig.Peers {
			consensusPeers.Peers[p.ID] = &types.Peer{Index: p.Index, Pubkey: p.ID}
		}
		if err := k.SetConsensusPeers(ctx, *consensusPeers); err != nil {
			return err
		}
		if err := k.SetKeyHeaderHash(ctx, consensusPeers.ChainId, header.Hash()); err != nil {
			return err
		}
	}
	return nil
}

func (k Keeper) SetConsensusPeers(ctx sdk.Context, consensusPeers types.ConsensusPeers) error {
	store := ctx.KVStore(k.storeKey)
	sink := polycommon.NewZeroCopySink(nil)
	consensusPeers.Serialize(sink)
	store.Set(GetConsensusPeerKey(consensusPeers.ChainId), sink.Bytes())
	return nil
}

func (k Keeper) GetConsensusPeers(ctx sdk.Context, chainID uint64) (*types.ConsensusPeers, error) {
	store := ctx.KVStore(k.storeKey)
	consensusPeerBytes := store.Get(GetConsensusPeerKey(chainID))
	if consensusPeerBytes == nil {
		return nil, types.ErrGetConsensusPeers(chainID)
	}
	consensusPeers := new(types.ConsensusPeers)
	if err := consensusPeers.Deserialize(polycommon.NewZeroCopySource(consensusPeerBytes)); err != nil {
		return nil, types.ErrDeserializeConsensusPeer(err)
	}
	return consensusPeers, nil
}

func (k Keeper) SetKeyHeaderHash(ctx sdk.Context, chainID uint64, keyHeaderHash polycommon.Uint256) error {
	store := ctx.KVStore(k.storeKey)
	store.Set(GetKeyHeaderHashKey(chainID), keyHeaderHash.ToArray())
	return nil
}

func (k Keeper) GetKeyHeaderHash(ctx sdk.Context, chainID uint64) (*polycommon.Uint256, error) {
	store := ctx.KVStore(k.storeKey)
	headerHashBs := store.Get(GetKeyHeaderHashKey(chainID))
	if headerHashBs == nil {
		return nil, types.ErrGetKeyHeaderHash(fmt.Sprintf("Empty key header hash with chainID: %d", chainID))
	}
	headerHash, err := polycommon.Uint256ParseFromBytes(headerHashBs)
	if err != nil {
		return nil, types.ErrGetKeyHeaderHash(fmt.Sprintf("Error: Uint256 from bytes: %x with chainID: %d", headerHashBs, err))
	}
	return &headerHash, nil
}

func (k Keeper) VerifyHistoricalHeader(ctx sdk.Context, header *polytype.Header, headerProof []byte, curHeader *polytype.Header) error {
	if err := k.VerifyHeaderSig(ctx, curHeader); err != nil {
		return err
	}
	value, err := merkle.MerkleProve(headerProof, curHeader.BlockRoot[:])
	if err != nil {
		return fmt.Errorf("VerifyHistoricalHeader, MerkleProve error: %s", err.Error())
	}
	hashToBeVerified := header.Hash()
	if !bytes.Equal(value, hashToBeVerified[:]) {
		return fmt.Errorf("VerifyHistoricalHeader error, historical header height: %d, current epoch header height: %d, expect: %x, got: %x", header.Height, curHeader.Height, hashToBeVerified[:], value)
	}
	return nil
}

func GetConsensusPeerKey(chainID uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, chainID)
	return append(ConsensusPeerPrefix, b...)
}

func GetKeyHeaderHashKey(chainID uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, chainID)
	return append(KeyHeaderHashPrefix, b...)
}

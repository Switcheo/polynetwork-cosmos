package keeper

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/Switcheo/polynetwork-cosmos/x/ccm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	polycommon "github.com/polynetwork/poly/common"
	polytype "github.com/polynetwork/poly/core/types"
	"github.com/polynetwork/poly/merkle"
	ccmc "github.com/polynetwork/poly/native/service/cross_chain_manager/common"
	ttype "github.com/tendermint/tendermint/types"
	// this line is used by starport scaffolding # ibc/keeper/import
)

type (
	Keeper struct {
		cdc      codec.Marshaler
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		bk       types.BankKeeper
		hsk      types.HeaderSyncKeeper
		ak       types.AssetKeeper
		ulkMap   map[string]types.UnlockKeeper
		// this line is used by starport scaffolding # ibc/keeper/attribute
	}
)

func NewKeeper(
	cdc codec.Marshaler,
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

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k *Keeper) MountUnlockKeeperMap(ulKeeperMap map[string]types.UnlockKeeper) {
	k.ulkMap = make(map[string]types.UnlockKeeper)
	for key, value := range ulKeeperMap {
		k.ulkMap[key] = value
	}
}

func (k *Keeper) MountAssetKeeper(assetKeeper types.AssetKeeper) {
	k.ak = assetKeeper
}

func (k Keeper) SetDenomCreator(ctx sdk.Context, denom string, creator sdk.AccAddress) {
	ctx.KVStore(k.storeKey).Set(GetDenomToCreatorKey(denom), creator.Bytes())
}

func (k Keeper) GetDenomCreator(ctx sdk.Context, denom string) (addr sdk.AccAddress) {
	creator := GetDenomToCreatorKey(denom)
	if creator == nil {
		return
	}
	return ctx.KVStore(k.storeKey).Get(creator)
}

func (k Keeper) ExistDenom(ctx sdk.Context, denom string) (string, bool) {
	storedSupplyCoins := k.bk.GetSupply(ctx).GetTotal()
	if len(k.GetDenomCreator(ctx, denom)) != 0 {
		return fmt.Sprintf("ccmKeeper.GetDenomCreator(ctx,%s) is %s", denom, sdk.AccAddress(k.GetDenomCreator(ctx, denom)).String()), true
	}
	if !storedSupplyCoins.AmountOf(denom).Equal(sdk.ZeroInt()) {
		return fmt.Sprintf("supply.AmountOf(%s) is %s", denom, storedSupplyCoins.AmountOf(denom).String()), true
	}
	return "", false
}

func (k Keeper) CreateCrossChainTx(ctx sdk.Context, fromAddr sdk.AccAddress, toChainID uint64, fromContractHash, toContractHash []byte, method string, args []byte) error {
	crossChainID, err := k.getCrossChainID(ctx)
	if err != nil {
		return err
	}
	if err := k.setCrossChainID(ctx, crossChainID.Add(sdk.NewInt(1))); err != nil {
		return err
	}

	ttx := make([]byte, len(ctx.TxBytes()))
	copy(ttx, ctx.TxBytes())
	txHash := ttype.Tx(ttx).Hash()
	crossChainIDBs := crossChainID.BigInt().Bytes()
	txParam := ccmc.MakeTxParam{
		TxHash:              txHash,
		CrossChainID:        crossChainIDBs,
		FromContractAddress: fromContractHash,
		ToChainID:           toChainID,
		ToContractAddress:   toContractHash,
		Method:              method,
		Args:                args,
	}
	sink := polycommon.NewZeroCopySink(nil)
	txParam.Serialization(sink)

	store := ctx.KVStore(k.storeKey)

	txParamHash := tmhash.Sum(sink.Bytes())
	store.Set(GetCrossChainTxKey(txParamHash), sink.Bytes())

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateCrossChainTx,
			sdk.NewAttribute(types.AttributeKeyStatus, "1"),
			sdk.NewAttribute(types.AttributeCrossChainID, crossChainID.String()),
			sdk.NewAttribute(types.AttributeKeyTxParamHash, hex.EncodeToString(txParamHash)),
			sdk.NewAttribute(types.AttributeKeyFromAddress, fromAddr.String()),
			sdk.NewAttribute(types.AttributeKeyFromContract, hex.EncodeToString(fromContractHash)),
			sdk.NewAttribute(types.AttributeKeyToChainID, strconv.FormatUint(toChainID, 10)),
			sdk.NewAttribute(types.AttributeKeyMakeTxParam, hex.EncodeToString(sink.Bytes())),
		),
	})
	return nil
}

func (k Keeper) ProcessUnlockTx(ctx sdk.Context, merkleValue *ccmc.ToMerkleValue, fromChainID uint64) error {
	// check if tocontractAddress is lockproxy module account, if yes, invoke lockproxy.unlock(), otherwise, invoke btcx.unlock
	for key, unlockKeeper := range k.ulkMap {
		k.Logger(ctx).Info(fmt.Sprintf("key is %+v ", key))
		k.Logger(ctx).Info(fmt.Sprintf("IfContains %+v ", unlockKeeper.ContainToContractAddr(ctx, merkleValue.MakeTxParam.ToContractAddress, fromChainID)))

		if unlockKeeper.ContainToContractAddr(ctx, merkleValue.MakeTxParam.ToContractAddress, merkleValue.FromChainID) {
			if err := unlockKeeper.Unlock(ctx, merkleValue.FromChainID, merkleValue.MakeTxParam.FromContractAddress, merkleValue.MakeTxParam.ToContractAddress, merkleValue.MakeTxParam.Args); err != nil {
				return types.ErrProcessCrossChainTx(fmt.Sprintf("Unlock failed, for module: %s, Error: %s", key, err.Error()))
			}
			return nil
		}
	}

	return types.ErrProcessCrossChainTx(fmt.Sprintf("Cannot find any unlock keeper to perform 'unlock' method for toContractAddr:%x, fromChainID:%d", merkleValue.MakeTxParam.ToContractAddress, fromChainID))
}

func (k Keeper) ProcessRegisterAssetTx(ctx sdk.Context, merkleValue *ccmc.ToMerkleValue) error {
	return k.ak.RegisterAsset(ctx, merkleValue.FromChainID, merkleValue.MakeTxParam.FromContractAddress, merkleValue.MakeTxParam.ToContractAddress, merkleValue.MakeTxParam.Args)
}

func (k Keeper) ProcessCrossChainTx(ctx sdk.Context, fromChainID uint64, proofStr string, headerStr, headerProofStr, curHeaderStr string) error {
	headerToBeVerified := new(polytype.Header)
	headerBs, err := hex.DecodeString(headerStr)
	if err != nil {
		return types.ErrProcessCrossChainTx(fmt.Sprintf("Failed to decode header hex string to bytes: %s, Error: %s ", headerStr, err.Error()))
	}
	if err := headerToBeVerified.Deserialization(polycommon.NewZeroCopySource(headerBs)); err != nil {
		return types.ErrProcessCrossChainTx(fmt.Sprintf("Failed to deserialize header, Error:%s", err.Error()))
	}

	headerInCurEpoch := new(polytype.Header)
	curHeaderBs, err := hex.DecodeString(curHeaderStr)
	if err != nil {
		headerInCurEpoch = nil
	} else {
		if err := headerInCurEpoch.Deserialization(polycommon.NewZeroCopySource(curHeaderBs)); err != nil {
			headerInCurEpoch = nil
		}
	}

	headerProof, err := hex.DecodeString(headerProofStr)
	if err != nil {
		headerProof = nil
	}

	if err := k.hsk.ProcessHeader(ctx, headerToBeVerified, headerProof, headerInCurEpoch); err != nil {
		return types.ErrProcessCrossChainTx(fmt.Sprintf("ProcessHeader Error, %s", err.Error()))
	}

	proof, err := hex.DecodeString(proofStr)
	if err != nil {
		return types.ErrProcessCrossChainTx(fmt.Sprintf("Failed to decode proof hex string to bytes: %s, Error: %s", proofStr, err.Error()))
	}

	merkleValue, err := k.VerifyToCosmosTx(ctx, proof, headerToBeVerified)
	if err != nil {
		return types.ErrProcessCrossChainTx(fmt.Sprintf("VerifyToCosmostx failed, %s", err.Error()))
	}
	currentChainCrossChainID := types.GetChainID()
	if merkleValue.MakeTxParam.ToChainID != currentChainCrossChainID {
		return types.ErrProcessCrossChainTx(fmt.Sprintf("toChainID is not for this chain, expect: %d, got: %d", currentChainCrossChainID, merkleValue.MakeTxParam.ToChainID))
	}

	switch merkleValue.MakeTxParam.Method {
	case "unlock":
		if err := k.ProcessUnlockTx(ctx, merkleValue, fromChainID); err != nil {
			return err
		}
	case "registerAsset":
		if err := k.ProcessRegisterAssetTx(ctx, merkleValue); err != nil {
			return err
		}
	default:
		return types.ErrProcessCrossChainTx(fmt.Sprintf("unsupported cross-chain method: %s", merkleValue.MakeTxParam.Method))
	}

	return nil
}

func (k Keeper) VerifyToCosmosTx(ctx sdk.Context, proof []byte, header *polytype.Header) (*ccmc.ToMerkleValue, error) {
	value, err := merkle.MerkleProve(proof, header.CrossStateRoot[:])
	if err != nil {
		return nil, types.ErrVerifyToCosmosTx(fmt.Sprintf("merkle.MerkleProve verify failed, Error: %s", err.Error()))
	}

	merkleValue := new(ccmc.ToMerkleValue)
	if err := merkleValue.Deserialization(polycommon.NewZeroCopySource(value)); err != nil {
		return nil, types.ErrVerifyToCosmosTx(fmt.Sprintf("ToMerkeValue Deserialization Error: %s", err.Error()))
	}

	if err := k.checkDoneTx(ctx, merkleValue.FromChainID, merkleValue.MakeTxParam.CrossChainID); err != nil {
		return nil, types.ErrVerifyToCosmosTx(fmt.Sprintf("check if this tx has been done, Error: %s", err.Error()))
	}

	k.putDoneTx(ctx, merkleValue.FromChainID, merkleValue.MakeTxParam.CrossChainID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeVerifyToCosmosProof,
			sdk.NewAttribute(types.AttributeKeyMerkleValueTxHash, hex.EncodeToString(merkleValue.TxHash)),
			sdk.NewAttribute(types.AttributeKeyMerkleValueMakeTxParamTxHash, hex.EncodeToString(merkleValue.MakeTxParam.TxHash)),
			sdk.NewAttribute(types.AttributeKeyFromChainID, strconv.FormatUint(merkleValue.FromChainID, 10)),
			sdk.NewAttribute(types.AttributeKeyMerkleValueMakeTxParamToContractAddress, hex.EncodeToString(merkleValue.MakeTxParam.ToContractAddress)),
		),
	})
	return merkleValue, nil

}

func (k Keeper) checkDoneTx(ctx sdk.Context, fromChainID uint64, crossChainID []byte) error {
	store := ctx.KVStore(k.storeKey)
	txKey := GetDoneTxKey(fromChainID, crossChainID)
	if txKey == nil {
		return fmt.Errorf("checkDoneTx, can't find tx key with fromChainID %d and crossChainID %x", fromChainID, crossChainID)
	}
	value := store.Get(txKey)
	if value != nil {
		return fmt.Errorf("checkDoneTx, tx already done with fromChainID: %d, crossChainID: %x", fromChainID, crossChainID)
	}
	return nil
}

func (k Keeper) putDoneTx(ctx sdk.Context, fromChainID uint64, crossChainID []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Set(GetDoneTxKey(fromChainID, crossChainID), crossChainID)
}

func (k Keeper) getCrossChainID(ctx sdk.Context) (sdk.Int, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(CrossChainIDKey)
	if bz == nil {
		return sdk.NewInt(0), nil
	}
	var crossChainID sdk.IntProto
	if err := crossChainID.Unmarshal(bz); err != nil {
		return sdk.NewInt(0), types.ErrUnmarshalSpecificTypeFail(crossChainID, err)
	}

	return crossChainID.Int, nil
}

func (k Keeper) setCrossChainID(ctx sdk.Context, crossChainID sdk.Int) error {
	store := ctx.KVStore(k.storeKey)
	intProto := sdk.IntProto{Int: crossChainID}
	bz, err := intProto.Marshal()
	if err != nil {
		return types.ErrMarshalSpecificTypeFail(crossChainID, err)
	}
	store.Set(CrossChainIDKey, bz)
	return nil
}

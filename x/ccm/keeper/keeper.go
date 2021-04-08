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

func (k Keeper) CreateCrossChainTx(ctx sdk.Context, fromAddr sdk.AccAddress, toChainId uint64, fromContractHash, toContractHash []byte, method string, args []byte) error {
	crossChainId, err := k.getCrossChainId(ctx)
	if err != nil {
		return err
	}
	if err := k.setCrossChainId(ctx, crossChainId.Add(sdk.NewInt(1))); err != nil {
		return err
	}

	ttx := make([]byte, len(ctx.TxBytes()))
	copy(ttx, ctx.TxBytes())
	txHash := ttype.Tx(ttx).Hash()
	crossChainIdBs := crossChainId.BigInt().Bytes()
	txParam := ccmc.MakeTxParam{
		TxHash:              txHash,
		CrossChainID:        crossChainIdBs,
		FromContractAddress: fromContractHash,
		ToChainID:           toChainId,
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
			sdk.NewAttribute(types.AttributeCrossChainId, crossChainId.String()),
			sdk.NewAttribute(types.AttributeKeyTxParamHash, hex.EncodeToString(txParamHash)),
			sdk.NewAttribute(types.AttributeKeyFromAddress, fromAddr.String()),
			sdk.NewAttribute(types.AttributeKeyFromContract, hex.EncodeToString(fromContractHash)),
			sdk.NewAttribute(types.AttributeKeyToChainId, strconv.FormatUint(toChainId, 10)),
			sdk.NewAttribute(types.AttributeKeyMakeTxParam, hex.EncodeToString(sink.Bytes())),
		),
	})
	return nil
}

func (k Keeper) ProcessUnlockTx(ctx sdk.Context, merkleValue *ccmc.ToMerkleValue, fromChainId uint64) error {
	// check if tocontractAddress is lockproxy module account, if yes, invoke lockproxy.unlock(), otherwise, invoke btcx.unlock
	for key, unlockKeeper := range k.ulkMap {
		k.Logger(ctx).Info(fmt.Sprintf("key is %+v ", key))
		k.Logger(ctx).Info(fmt.Sprintf("IfContains %+v ", unlockKeeper.ContainToContractAddr(ctx, merkleValue.MakeTxParam.ToContractAddress, fromChainId)))

		if unlockKeeper.ContainToContractAddr(ctx, merkleValue.MakeTxParam.ToContractAddress, merkleValue.FromChainID) {
			if err := unlockKeeper.Unlock(ctx, merkleValue.FromChainID, merkleValue.MakeTxParam.FromContractAddress, merkleValue.MakeTxParam.ToContractAddress, merkleValue.MakeTxParam.Args); err != nil {
				return types.ErrProcessCrossChainTx(fmt.Sprintf("Unlock failed, for module: %s, Error: %s", key, err.Error()))
			}
			return nil
		}
	}

	return types.ErrProcessCrossChainTx(fmt.Sprintf("Cannot find any unlock keeper to perform 'unlock' method for toContractAddr:%x, fromChainId:%d", merkleValue.MakeTxParam.ToContractAddress, fromChainId))
}

func (k Keeper) ProcessRegisterAssetTx(ctx sdk.Context, merkleValue *ccmc.ToMerkleValue) error {
	return k.ak.RegisterAsset(ctx, merkleValue.FromChainID, merkleValue.MakeTxParam.FromContractAddress, merkleValue.MakeTxParam.ToContractAddress, merkleValue.MakeTxParam.Args)
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
			sdk.NewAttribute(types.AttributeKeyFromChainId, strconv.FormatUint(merkleValue.FromChainID, 10)),
			sdk.NewAttribute(types.AttributeKeyMerkleValueMakeTxParamToContractAddress, hex.EncodeToString(merkleValue.MakeTxParam.ToContractAddress)),
		),
	})
	return merkleValue, nil

}

func (k Keeper) checkDoneTx(ctx sdk.Context, fromChainId uint64, crossChainId []byte) error {
	store := ctx.KVStore(k.storeKey)
	txKey := GetDoneTxKey(fromChainId, crossChainId)
	if txKey == nil {
		return fmt.Errorf("checkDoneTx, can't find tx key with fromChainId %d and crossChainId %x", fromChainId, crossChainId)
	}
	value := store.Get(txKey)
	if value != nil {
		return fmt.Errorf("checkDoneTx, tx already done with fromChainId: %d, crossChainId: %x", fromChainId, crossChainId)
	}
	return nil
}

func (k Keeper) putDoneTx(ctx sdk.Context, fromChainId uint64, crossChainId []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Set(GetDoneTxKey(fromChainId, crossChainId), crossChainId)
}

func (k Keeper) getCrossChainId(ctx sdk.Context) (sdk.Int, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(CrossChainIdKey)
	if bz == nil {
		return sdk.NewInt(0), nil
	}
	var crossChainId sdk.IntProto
	if err := crossChainId.Unmarshal(bz); err != nil {
		return sdk.NewInt(0), types.ErrUnmarshalSpecificTypeFail(crossChainId, err)
	}

	return crossChainId.Int, nil
}

func (k Keeper) setCrossChainId(ctx sdk.Context, crossChainId sdk.Int) error {
	store := ctx.KVStore(k.storeKey)
	intProto := sdk.IntProto{Int: crossChainId}
	bz, err := intProto.Marshal()
	if err != nil {
		return types.ErrMarshalSpecificTypeFail(crossChainId, err)
	}
	store.Set(CrossChainIdKey, bz)
	return nil
}

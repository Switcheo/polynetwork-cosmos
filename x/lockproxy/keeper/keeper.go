package keeper

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/Switcheo/polynetwork-cosmos/x/lockproxy/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	polycommon "github.com/polynetwork/poly/common"
	// this line is used by starport scaffolding # ibc/keeper/import
)

type (
	Keeper struct {
		cdc      codec.Marshaler
		ak       types.AccountKeeper
		bk       types.BankKeeper
		ck       types.CCMKeeper
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		hooks    types.LockProxyHooks
		// this line is used by starport scaffolding # ibc/keeper/attribute
	}
)

var (
	OperatorToLockProxyKey = []byte{0x01}
	BindChainIdPrefix      = []byte{0x02}
	RegistryPrefix         = []byte{0x03}
	NonceKey               = []byte("nonce")
)

func NewKeeper(
	cdc codec.Marshaler,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	ck types.CCMKeeper,
	storeKey,
	memKey sdk.StoreKey,
	// this line is used by starport scaffolding # ibc/keeper/parameter
) *Keeper {
	return &Keeper{
		cdc:      cdc,
		ak:       ak,
		bk:       bk,
		ck:       ck,
		storeKey: storeKey,
		memKey:   memKey,
		// this line is used by starport scaffolding # ibc/keeper/return
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Store fetches the main kv store
func (k Keeper) Store(ctx sdk.Context) sdk.KVStore {
	return ctx.KVStore(k.storeKey)
}

// StoreIterator returns the iterator for the store
func (k Keeper) StoreIterator(ctx sdk.Context, prefix []byte) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, prefix)
}

func (k Keeper) ContainToContractAddr(ctx sdk.Context, toContractAddr []byte, fromChainId uint64) bool {
	return ctx.KVStore(k.storeKey).Get((GetBindChainIdKey(toContractAddr, fromChainId))) != nil
}

func (k Keeper) CreateLockProxy(ctx sdk.Context, creator sdk.AccAddress) error {
	if k.LockProxyExists(ctx, creator) {
		return types.ErrAlreadyCreated(
			fmt.Sprintf("creator: %s already created lockproxy contract with hash: %x", creator.String(), creator.Bytes()))
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(GetOperatorToLockProxyKey(creator), creator.Bytes())
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateLockProxy,
			sdk.NewAttribute(types.AttributeKeyCreator, creator.String()),
			sdk.NewAttribute(types.AttributeKeyProxyHash, hex.EncodeToString(creator.Bytes())),
		),
	})
	ctx.Logger().With("module", fmt.Sprintf("creator: %s created a lockproxy contract with hash: %x", creator.String(), creator.Bytes()))
	return nil
}

func (k Keeper) LockProxyExists(ctx sdk.Context, creator sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return bytes.Equal(store.Get(GetOperatorToLockProxyKey(creator)), creator)
}

func (k Keeper) GetLockProxy(ctx sdk.Context, operator sdk.AccAddress) []byte {
	store := ctx.KVStore(k.storeKey)
	proxyBytes := store.Get(GetOperatorToLockProxyKey(operator))
	if len(proxyBytes) == 0 || !bytes.Equal(operator.Bytes(), proxyBytes) {
		return nil
	}
	return proxyBytes
}

func (k Keeper) updateRegistry(ctx sdk.Context, lockProxyHash []byte, assetHash []byte,
	nativeChainId uint64, nativeLockProxyHash []byte, nativeAssetHash []byte) error {
	if k.AssetIsRegistered(ctx, lockProxyHash, assetHash, nativeChainId, nativeLockProxyHash, nativeAssetHash) {
		return types.ErrAlreadyRegistered(
			fmt.Sprintf("asset already registered %x, %d, %x, %x", assetHash, nativeChainId, nativeLockProxyHash, nativeAssetHash))
	}

	store := ctx.KVStore(k.storeKey)
	registryKey := GetRegistryKey(lockProxyHash, assetHash, nativeChainId, nativeLockProxyHash, nativeAssetHash)
	store.Set(registryKey, []byte("1"))

	// GetBindChainIdKey is used in ContainToContractAddr to check when to return true
	// this will allow the module to be called by the ccm keeper to handle the appropriate cross-chain txns
	bindChainIdKey := GetBindChainIdKey(lockProxyHash, nativeChainId)
	if store.Get(bindChainIdKey) == nil {
		store.Set(bindChainIdKey, []byte("1"))
	}

	return nil
}

// RegisterAsset registers an asset based on the polynetwork lock proxy architecture. However,
// this lockproxy uses the PIP-1 implementation for permssionless registration, and registerations
// from a non-source chain is disallowed. As such this method will always returns an error.
func (k Keeper) RegisterAsset(ctx sdk.Context, fromChainID uint64, fromContractAddr []byte, toContractAddr []byte, argsBs []byte) error {
	return types.ErrRegisterAsset("asset registration disallowed")
}

// AssetIsRegistered returns whether the given assetID, chainID, denom, denom creator tuple has been registered.
func (k Keeper) AssetIsRegistered(ctx sdk.Context, lockProxyHash []byte, assetHash []byte,
	nativeChainId uint64, nativeLockProxyHash []byte, nativeAssetHash []byte) bool {
	store := ctx.KVStore(k.storeKey)
	key := GetRegistryKey(lockProxyHash, assetHash, nativeChainId, nativeLockProxyHash, nativeAssetHash)
	registryBytes := store.Get(key)
	return len(registryBytes) != 0
}

// CreateCoinAndDelegateToProxy creates a new coin for a given creator and registers it to the given lock contract and asset on the native chain.
func (k Keeper) CreateCoinAndDelegateToProxy(ctx sdk.Context, creator sdk.AccAddress, denom string, lockproxyHash []byte,
	nativeChainId uint64, nativeLockProxyHash []byte, nativeAssetHash []byte) error {
	if len(k.ck.GetDenomCreator(ctx, denom)) != 0 {
		return types.ErrCreateCoinAndDelegateToProxy(fmt.Sprintf("denom: %s already registered", denom))
	}
	if !k.LockProxyExists(ctx, lockproxyHash) {
		return types.ErrCreateCoinAndDelegateToProxy(fmt.Sprintf("lockproxy does not exist: %x", lockproxyHash))
	}

	k.ck.SetDenomCreator(ctx, denom, creator)

	if err := k.updateRegistry(ctx, lockproxyHash, []byte(denom), nativeChainId, nativeLockProxyHash, nativeAssetHash); err != nil {
		return err
	}

	args := types.RegisterAssetTxArgs{
		AssetHash:       []byte(denom),
		NativeAssetHash: nativeAssetHash,
	}
	sink := polycommon.NewZeroCopySink(nil)
	if err := args.Serialize(sink); err != nil {
		return types.ErrCreateCoinAndDelegateToProxy(fmt.Sprintf("TxArgs.Serialization error:%v", err))
	}
	if err := k.ck.CreateCrossChainTx(ctx, creator, nativeChainId, lockproxyHash, nativeLockProxyHash, "registerAsset", sink.Bytes()); err != nil {
		return types.ErrCreateCoinAndDelegateToProxy(
			fmt.Sprintf("ccmKeeper.CreateCrossChainTx error: toChainId: %d, fromContractHash: %x, toChainProxyHash: %x, args: %x",
				nativeChainId, lockproxyHash, nativeLockProxyHash, args))
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateAndDelegateCoinToProxy,
			sdk.NewAttribute(types.AttributeKeySourceAssetDenom, denom),
			sdk.NewAttribute(types.AttributeKeyCreator, creator.String()),
		),
	})
	return nil
}

// SyncRegisteredAsset syncs the registerAsset tx of an already registered asset to the native chain.
func (k Keeper) SyncRegisteredAsset(ctx sdk.Context, syncer sdk.AccAddress, nativeChainID uint64, denom string, nativeAssetHash, lockProxyHash, nativeLockProxyHash []byte) error {
	assetHash := []byte(denom)

	// ensure the asset is indeed registered
	if !k.AssetIsRegistered(ctx, lockProxyHash, assetHash, nativeChainID, nativeLockProxyHash, nativeAssetHash) {
		return types.ErrSyncRegisteredAsset(fmt.Sprintf("asset not yet registered %x, %x, %d, %x, %x", lockProxyHash, assetHash, nativeChainID, nativeLockProxyHash, nativeAssetHash))
	}

	args := types.RegisterAssetTxArgs{
		AssetHash:       assetHash,
		NativeAssetHash: nativeAssetHash,
	}

	sink := polycommon.NewZeroCopySink(nil)
	if err := args.Serialize(sink); err != nil {
		return types.ErrSyncRegisteredAsset(fmt.Sprintf("TxArgs Serialization Error:%v", err))
	}

	if err := k.ck.CreateCrossChainTx(ctx, syncer, nativeChainID, lockProxyHash, nativeLockProxyHash, "registerAsset", sink.Bytes()); err != nil {
		return types.ErrSyncRegisteredAsset(
			fmt.Sprintf("ccmKeeper.CreateCrossChainTx Error: toChainId: %d, fromContractHash: %x, toChainProxyHash: %x, args: %x",
				nativeChainID, lockProxyHash, nativeLockProxyHash, args))
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSyncRegisteredAsset,
			sdk.NewAttribute(types.AttributeKeyToChainId, fmt.Sprintf("%d", nativeChainID)),
			sdk.NewAttribute(types.AttributeKeyAssetHash, hex.EncodeToString(assetHash)),
			sdk.NewAttribute(types.AttributeKeyNativeAssetHash, hex.EncodeToString(nativeAssetHash)),
			sdk.NewAttribute(types.AttributeKeyProxyHash, hex.EncodeToString(lockProxyHash)),
			sdk.NewAttribute(types.AttributeKeyToChainProxyHash, hex.EncodeToString(nativeLockProxyHash)),
		),
	})

	return nil
}

func (k Keeper) GetNonce(ctx sdk.Context) sdk.Int {
	store := ctx.KVStore(k.storeKey)

	nonce := sdk.ZeroInt()
	nonceBz := store.Get(NonceKey)
	nonce.Unmarshal(nonceBz)

	return nonce
}

func (k Keeper) SetNonce(ctx sdk.Context, x sdk.Int) {
	store := ctx.KVStore(k.storeKey)
	newNonceBz, err := x.Marshal()
	if err != nil {
		panic(err)
	}
	store.Set(NonceKey, newNonceBz)
}

func (k Keeper) getNextNonce(ctx sdk.Context) sdk.Int {
	nonce := k.GetNonce(ctx)
	newNonce := nonce.Add(sdk.NewInt(1))
	k.SetNonce(ctx, newNonce)

	return newNonce
}

// LockAsset sends tokens to this module, releasing it on the toChain.
// The tokens are burnt in this module to give the correct global supply.
// CONTRACT: values and addresses should already be statically validated.
func (k Keeper) LockAsset(ctx sdk.Context, lockProxyHash []byte, fromAddress sdk.AccAddress, sourceAssetDenom string,
	toChainId uint64, toChainProxyHash []byte, toChainAssetHash []byte, toAddressBs []byte,
	value sdk.Int, deductFeeInLock bool, feeAmount sdk.Int, feeAddress sdk.AccAddress) error {
	if !k.LockProxyExists(ctx, lockProxyHash) {
		return types.ErrLock(fmt.Sprintf("lockproxy does not exist: %x", lockProxyHash))
	}

	nonce := k.getNextNonce(ctx)
	args := types.TxArgs{
		FromAddress:   fromAddress,
		FromAssetHash: []byte(sourceAssetDenom),
		ToAssetHash:   toChainAssetHash,
		ToAddress:     toAddressBs,
		Amount:        value.BigInt(),
		FeeAmount:     feeAmount.BigInt(),
		FeeAddress:    feeAddress,
		Nonce:         nonce.BigInt(),
	}

	afterFeeAmount := value
	if deductFeeInLock {
		afterFeeAmount = value.Sub(feeAmount)
		feeCoins := sdk.NewCoins(sdk.NewCoin(sourceAssetDenom, feeAmount))
		err := k.bk.SendCoins(ctx, fromAddress, feeAddress, feeCoins)
		if err != nil {
			return types.ErrLock(fmt.Sprintf("bankKeeper.SendCoins Error: from: %s, amount: %s", fromAddress.String(), feeCoins.String()))
		}

		args.Amount = afterFeeAmount.BigInt()
		args.FeeAmount = big.NewInt(0)
	}

	// send coin of sourceAssetDenom from fromAddress to module account address
	amountCoins := sdk.NewCoins(sdk.NewCoin(sourceAssetDenom, afterFeeAmount))
	if err := k.bk.SendCoinsFromAccountToModule(ctx, fromAddress, types.ModuleName, amountCoins); err != nil {
		return types.ErrLock(fmt.Sprintf(
			"Failed to send %s from account %s to module account %s (%s). Error: %v",
			amountCoins.String(), fromAddress.String(), types.ModuleName, k.ak.GetModuleAddress(types.ModuleName), err))
	}

	// burn the module account coins
	if err := k.bk.BurnCoins(ctx, types.ModuleName, amountCoins); err != nil {
		return types.ErrLock(fmt.Sprintf("supplyKeeper.BurnCoins error: %s", err.Error()))
	}

	sink := polycommon.NewZeroCopySink(nil)
	if err := args.Serialize(sink, 32); err != nil {
		return types.ErrLock(fmt.Sprintf("TxArgs.Serialize error:%v", err))
	}

	fromContractHash := lockProxyHash
	if err := k.ck.CreateCrossChainTx(ctx, fromAddress, toChainId, fromContractHash, toChainProxyHash, "unlock", sink.Bytes()); err != nil {
		return types.ErrLock(fmt.Sprintf("ccmKeeper.CreateCrossChainTx error: toChainId: %d, fromContractHash: %x, toChainProxyHash: %x, args: %x",
			toChainId, fromContractHash, toChainProxyHash, args))
	}

	if !k.AssetIsRegistered(ctx, lockProxyHash, []byte(sourceAssetDenom), toChainId, toChainProxyHash, toChainAssetHash) {
		return types.ErrLock(fmt.Sprintf("asset not registered: lockProxyHash: %s, denom: %s, toChainId: %d, toChainProxyHash: %s, toChainAssetHash: %s",
			string(lockProxyHash), sourceAssetDenom, toChainId, hex.EncodeToString(toChainProxyHash), hex.EncodeToString(toChainAssetHash)))
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeLock,
			sdk.NewAttribute(types.AttributeKeyFromContractHash, hex.EncodeToString([]byte(sourceAssetDenom))),
			sdk.NewAttribute(types.AttributeKeyToChainId, strconv.FormatUint(toChainId, 10)),
			sdk.NewAttribute(types.AttributeKeyToChainProxyHash, hex.EncodeToString(toChainProxyHash)),
			sdk.NewAttribute(types.AttributeKeyToChainAssetHash, hex.EncodeToString(toChainAssetHash)),
			sdk.NewAttribute(types.AttributeKeyFromAddress, fromAddress.String()),
			sdk.NewAttribute(types.AttributeKeyToAddress, hex.EncodeToString(toAddressBs)),
			sdk.NewAttribute(types.AttributeKeyAmount, value.String()),
			sdk.NewAttribute(types.AttributeKeyLockProxy, hex.EncodeToString(fromContractHash)),
			sdk.NewAttribute(types.AttributeKeyFeeAmount, feeAmount.String()),
			sdk.NewAttribute(types.AttributeKeyFeeAddress, feeAddress.String()),
			sdk.NewAttribute(types.AttributeKeyNonce, nonce.String()),
		),
	})

	return nil
}

// Unlock sends tokens from this module to the target account.
// The tokens are minted before releasing, to give the correct global supply.
func (k Keeper) Unlock(ctx sdk.Context, fromChainId uint64, fromContractAddr sdk.AccAddress, toContractAddr []byte, argsBs []byte) error {
	args := new(types.TxArgs)
	if err := args.Deserialize(polycommon.NewZeroCopySource(argsBs), 32); err != nil {
		return types.ErrUnlock(fmt.Sprintf("TxArgs.Deserialize error:%s", err))
	}
	fromAssetHash := args.FromAssetHash
	toAssetHash := args.ToAssetHash
	toAddress := args.ToAddress
	amount := sdk.NewIntFromBigInt(args.Amount)
	feeAmount := sdk.NewIntFromBigInt(args.FeeAmount)
	nonce := sdk.NewIntFromBigInt(args.Nonce)

	if !k.AssetIsRegistered(ctx, toContractAddr, toAssetHash, fromChainId, fromContractAddr, fromAssetHash) {
		return types.ErrUnlock(fmt.Sprintf("asset not registered: toContractAddr: %s, toAssetHash: %s, fromChainId: %d, fromContractAddr: %s, fromAssetHash: %s",
			string(toContractAddr), toAssetHash, fromChainId, hex.EncodeToString(fromContractAddr), hex.EncodeToString(fromAssetHash)))
	}

	// to asset hash should be the hex format string of source asset denom name, NOT Module account address
	toAssetDenom := string(toAssetHash)

	toAcctAddress := make(sdk.AccAddress, len(toAddress))
	copy(toAcctAddress, toAddress)

	fromAcctAddress := sdk.AccAddress(args.FromAddress)
	if fromAcctAddress.Empty() {
		return types.ErrUnlock("fromAddress is empty")
	}

	// mint coin of toAssetDenom unless legacy version
	mintCoins := sdk.NewCoins(sdk.NewCoin(toAssetDenom, amount))
	if err := k.bk.MintCoins(ctx, types.ModuleName, mintCoins); err != nil {
		return types.ErrUnlock(fmt.Sprintf("supplyKeeper.MintCoins error: %s", err.Error()))
	}

	afterFeeAmount := amount
	feeAddressAcc := sdk.AccAddress(args.FeeAddress)
	if feeAmount.GT(sdk.ZeroInt()) {
		if feeAmount.GT(amount) {
			return types.ErrUnlock(fmt.Sprintf("feeAmount: %s must be less than or equal to amount: %s", feeAmount.String(), amount.String()))
		}

		if feeAddressAcc.Empty() {
			return types.ErrUnlock("feeAmount is present but feeAddress is empty")
		}

		afterFeeAmount = afterFeeAmount.Sub(feeAmount)
		feeCoins := sdk.NewCoins(sdk.NewCoin(toAssetDenom, feeAmount))
		if err := k.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, feeAddressAcc, feeCoins); err != nil {
			return types.ErrUnlock(fmt.Sprintf(
				"Failed to send %s from module account %s to fee account %s. Error: %v",
				feeCoins.String(), k.ak.GetModuleAddress(types.ModuleName).String(), feeAddressAcc.String(), err))
		}
	}
	amountCoins := sdk.NewCoins(sdk.NewCoin(toAssetDenom, afterFeeAmount))
	if err := k.bk.SendCoinsFromModuleToAccount(ctx, types.ModuleName, toAcctAddress, amountCoins); err != nil {
		return types.ErrUnlock(fmt.Sprintf(
			"Failed to send %s from module account: %s to receiver account: %s. Error: %v",
			amountCoins.String(), k.ak.GetModuleAddress(types.ModuleName).String(), toAcctAddress.String(), err))
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUnlock,
			sdk.NewAttribute(types.AttributeKeyToChainAssetHash, hex.EncodeToString([]byte(toAssetDenom))),
			sdk.NewAttribute(types.AttributeKeyToAddress, toAcctAddress.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, amount.String()),
			sdk.NewAttribute(types.AttributeKeyFromAddress, fromAcctAddress.String()),
			sdk.NewAttribute(types.AttributeKeySourceAssetHash, hex.EncodeToString(fromAssetHash)),
			sdk.NewAttribute(types.AttributeKeyFeeAmount, feeAmount.String()),
			sdk.NewAttribute(types.AttributeKeyFeeAddress, feeAddressAcc.String()),
			sdk.NewAttribute(types.AttributeKeyNonce, nonce.String()),
		),
	})

	k.AfterProxyUnlock(ctx, fromAcctAddress, toAcctAddress, amountCoins)

	return nil
}

func GetOperatorToLockProxyKey(operator sdk.AccAddress) []byte {
	return append(OperatorToLockProxyKey, operator...)
}

func GetHashKey(lockProxyHash []byte, assetHash []byte, nativeChainId uint64, nativeLockProxyHash []byte, nativeAssetHash []byte) []byte {
	nativeChainIdBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(nativeChainIdBytes, nativeChainId)

	lockProxyHashBz := sha256.Sum256(lockProxyHash)
	assetHashBz := sha256.Sum256(assetHash)
	nativeChainIdBz := sha256.Sum256(nativeChainIdBytes)
	nativeLockProxyHashBz := sha256.Sum256(nativeLockProxyHash)
	nativeAssetHashBz := sha256.Sum256(nativeAssetHash)

	return append(append(append(append(lockProxyHashBz[:], assetHashBz[:]...), nativeChainIdBz[:]...), nativeLockProxyHashBz[:]...), nativeAssetHashBz[:]...)
}

func GetRegistryKey(lockProxyHash []byte, assetHash []byte, nativeChainId uint64, nativeLockProxyHash []byte, nativeAssetHash []byte) []byte {
	hashKey := GetHashKey(lockProxyHash, assetHash, nativeChainId, nativeLockProxyHash, nativeAssetHash)
	return append(RegistryPrefix, hashKey...)
}

func GetBindChainIdKey(proxyHash []byte, toChainId uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, toChainId)
	return append(append(BindChainIdPrefix, proxyHash...), b...)
}

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
	BindChainIDPrefix      = []byte{0x02}
	RegistryPrefix         = []byte{0x03}
	BalancePrefix          = []byte{0x04}
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

// EnsureAccountExist returns an err if the give accAddress is not created yet.
func (k Keeper) EnsureAccountExist(ctx sdk.Context, addr sdk.AccAddress) error {
	acct := k.ak.GetAccount(ctx, addr)
	if acct == nil {
		return types.ErrAccountNotExist(fmt.Sprintf("account %s does not exist", addr.String()))
	}
	return nil
}

func (k Keeper) ContainToContractAddr(ctx sdk.Context, toContractAddr []byte, fromChainID uint64) bool {
	return ctx.KVStore(k.storeKey).Get((GetBindChainIDKey(toContractAddr, fromChainID))) != nil
}

func (k Keeper) CreateLockProxy(ctx sdk.Context, creator sdk.AccAddress) error {
	if k.EnsureLockProxyExist(ctx, creator) {
		return types.ErrCreateLockProxy(fmt.Sprintf("creator:%s already created lockproxy contract with hash:%x", creator.String(), creator.Bytes()))
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
	ctx.Logger().With("module", fmt.Sprintf("creator:%s initialized a lockproxy contract with hash: %x", creator.String(), creator.Bytes()))
	return nil
}

func (k Keeper) EnsureLockProxyExist(ctx sdk.Context, creator sdk.AccAddress) bool {
	store := ctx.KVStore(k.storeKey)
	return bytes.Equal(store.Get(GetOperatorToLockProxyKey(creator)), creator)
}

func (k Keeper) GetLockProxyByOperator(ctx sdk.Context, operator sdk.AccAddress) []byte {
	store := ctx.KVStore(k.storeKey)
	proxyBytes := store.Get(GetOperatorToLockProxyKey(operator))
	if len(proxyBytes) == 0 || !bytes.Equal(operator.Bytes(), proxyBytes) {
		return nil
	}
	return proxyBytes
}

func (k Keeper) updateRegistry(ctx sdk.Context, lockProxyHash []byte, assetHash []byte,
	nativeChainID uint64, nativeLockProxyHash []byte, nativeAssetHash []byte) error {
	if k.AssetIsRegistered(ctx, lockProxyHash, assetHash, nativeChainID, nativeLockProxyHash, nativeAssetHash) {
		return types.ErrRegistryAlreadyExists(fmt.Sprintf("asset already registered %x, %d, %x, %x", assetHash, nativeChainID, nativeLockProxyHash, nativeAssetHash))
	}

	store := ctx.KVStore(k.storeKey)
	registryKey := GetRegistryKey(lockProxyHash, assetHash, nativeChainID, nativeLockProxyHash, nativeAssetHash)
	store.Set(registryKey, []byte("1"))

	// GetBindChainIDKey is used in ContainToContractAddr to check when to return true
	// this will allow the module to be called by the ccm keeper to handle the appropriate cross-chain txns
	bindChainIDKey := GetBindChainIDKey(lockProxyHash, nativeChainID)
	if store.Get(bindChainIDKey) == nil {
		store.Set(bindChainIDKey, []byte("1"))
	}

	return nil
}

// AssetIsRegistered returns whether the given assetID, chainID, denom, denom creator tuple has been registered.
func (k Keeper) AssetIsRegistered(ctx sdk.Context, lockProxyHash []byte, assetHash []byte,
	nativeChainID uint64, nativeLockProxyHash []byte, nativeAssetHash []byte) bool {
	store := ctx.KVStore(k.storeKey)
	key := GetRegistryKey(lockProxyHash, assetHash, nativeChainID, nativeLockProxyHash, nativeAssetHash)
	registryBytes := store.Get(key)
	return len(registryBytes) != 0
}

// RegisterAsset registers an asset.
//
// Deprecated: this method is deprecated and always returns an error.
func (k Keeper) RegisterAsset(ctx sdk.Context, fromChainID uint64, fromContractAddr []byte, toContractAddr []byte, argsBs []byte) error {
	return types.ErrRegisterAsset("asset registration disallowed")
}

// CreateCoinAndDelegateToProxy creates a new coin for a given creator and registers it to the given lock contract and asset on the native chain.
func (k Keeper) CreateCoinAndDelegateToProxy(ctx sdk.Context, creator sdk.AccAddress, coin sdk.Coin, lockproxyHash []byte,
	nativeChainID uint64, nativeLockProxyHash []byte, nativeAssetHash []byte) error {
	if len(k.ck.GetDenomCreator(ctx, coin.Denom)) != 0 {
		return types.ErrCreateCoinAndDelegateToProxy(fmt.Sprintf("denom:%s already exists", coin.Denom))
	}
	if exist := k.EnsureLockProxyExist(ctx, lockproxyHash); !exist {
		return types.ErrCreateCoinAndDelegateToProxy(fmt.Sprintf("lockproxy with hash: %s not created", lockproxyHash))
	}

	k.ck.SetDenomCreator(ctx, coin.Denom, creator)

	if err := k.updateRegistry(ctx, lockproxyHash, []byte(coin.Denom), nativeChainID, nativeLockProxyHash, nativeAssetHash); err != nil {
		return err
	}

	if !coin.Amount.IsZero() {
		// version > 0 should create coins with 0 amt
		return types.ErrCreateCoinAndDelegateToProxy(fmt.Sprintf("coin amount should be zero %d", coin.Amount))
	}

	args := types.RegisterAssetTxArgs{
		AssetHash:       []byte(coin.Denom),
		NativeAssetHash: nativeAssetHash,
	}
	sink := polycommon.NewZeroCopySink(nil)
	if err := args.Serialize(sink); err != nil {
		return types.ErrCreateCoinAndDelegateToProxy(fmt.Sprintf("TxArgs Serialization Error:%v", err))
	}
	if err := k.ck.CreateCrossChainTx(ctx, creator, nativeChainID, lockproxyHash, nativeLockProxyHash, "registerAsset", sink.Bytes()); err != nil {
		return types.ErrCreateCoinAndDelegateToProxy(
			fmt.Sprintf("ccmKeeper.CreateCrossChainTx Error: toChainID: %d, fromContractHash: %x, toChainProxyHash: %x, args: %x",
				nativeChainID, lockproxyHash, nativeLockProxyHash, args))
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateAndDelegateCoinToProxy,
			sdk.NewAttribute(types.AttributeKeySourceAssetDenom, coin.Denom),
			sdk.NewAttribute(types.AttributeKeyCreator, creator.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, coin.Amount.String()),
		),
	})
	return nil
}

func (k Keeper) getNextNonce(ctx sdk.Context) sdk.Int {
	store := ctx.KVStore(k.storeKey)

	var nonce sdk.IntProto
	nonceBz := store.Get(NonceKey) // TODO: This needs to be migrated in genesis!
	if nonceBz != nil {
		err := k.cdc.UnmarshalBinaryLengthPrefixed(nonceBz, &nonce)
		if err != nil {
			panic(err)
		}
	}

	nonce.Int = nonce.Int.Add(sdk.OneInt())
	newNonceBz, err := k.cdc.MarshalBinaryLengthPrefixed(&nonce)
	if err != nil {
		panic(err)
	}
	store.Set(NonceKey, newNonceBz)

	return nonce.Int
}

// LockAsset sends tokens to this module, releasing it on the toChain.
// On version > 0, the tokens are burnt to give the correct global supply.
func (k Keeper) LockAsset(ctx sdk.Context, lockProxyHash []byte, fromAddress sdk.AccAddress, sourceAssetDenom string,
	toChainID uint64, toChainProxyHash []byte, toChainAssetHash []byte, toAddressBs []byte,
	value sdk.Int, deductFeeInLock bool, feeAmount sdk.Int, feeAddress []byte) error {
	if exist := k.EnsureLockProxyExist(ctx, lockProxyHash); !exist {
		return types.ErrLock(fmt.Sprintf("lockproxy with hash: %s not created", lockProxyHash))
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
	feeAddressAcc := sdk.AccAddress(args.FeeAddress)
	if deductFeeInLock && feeAmount.GT(sdk.ZeroInt()) {
		if feeAddressAcc.Empty() {
			return types.ErrLock("FeeAmount is present but FeeAddress is empty")
		}

		if feeAmount.GT(value) {
			return types.ErrLock(fmt.Sprintf("feeAmount %s is greater than value %s", feeAmount.String(), value.String()))
		}

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

	// burn the module account coins unless legacy version
	if err := k.bk.BurnCoins(ctx, types.ModuleName, amountCoins); err != nil {
		return types.ErrLock(fmt.Sprintf("supplyKeeper.BurnCoins Error: %s", err.Error()))
	}

	sink := polycommon.NewZeroCopySink(nil)
	if err := args.Serialize(sink, 32); err != nil {
		return types.ErrLock(fmt.Sprintf("TxArgs Serialization Error:%v", err))
	}
	fromContractHash := lockProxyHash
	if err := k.ck.CreateCrossChainTx(ctx, fromAddress, toChainID, fromContractHash, toChainProxyHash, "unlock", sink.Bytes()); err != nil {
		return types.ErrLock(fmt.Sprintf("ccmKeeper.CreateCrossChainTx Error: toChainID: %d, fromContractHash: %x, toChainProxyHash: %x, args: %x",
			toChainID, fromContractHash, toChainProxyHash, args))
	}
	if amountCoins.AmountOf(sourceAssetDenom).IsNegative() {
		return types.ErrLock(fmt.Sprintf("the coin being crossed has negative amount value, coin:%s", amountCoins.String()))
	}

	if !k.AssetIsRegistered(ctx, lockProxyHash, []byte(sourceAssetDenom), toChainID, toChainProxyHash, toChainAssetHash) {
		return types.ErrLock(fmt.Sprintf("missing asset registry: lockProxyHash: %s, denom: %s, toChainID: %d, toChainProxyHash: %s, toChainAssetHash: %s",
			string(lockProxyHash), sourceAssetDenom, toChainID, hex.EncodeToString(toChainProxyHash), hex.EncodeToString(toChainAssetHash)))
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeLock,
			sdk.NewAttribute(types.AttributeKeyFromContractHash, hex.EncodeToString([]byte(sourceAssetDenom))),
			sdk.NewAttribute(types.AttributeKeyToChainID, strconv.FormatUint(toChainID, 10)),
			sdk.NewAttribute(types.AttributeKeyToChainProxyHash, hex.EncodeToString(toChainProxyHash)),
			sdk.NewAttribute(types.AttributeKeyToChainAssetHash, hex.EncodeToString(toChainAssetHash)),
			sdk.NewAttribute(types.AttributeKeyFromAddress, fromAddress.String()),
			sdk.NewAttribute(types.AttributeKeyToAddress, hex.EncodeToString(toAddressBs)),
			sdk.NewAttribute(types.AttributeKeyAmount, value.String()),
			sdk.NewAttribute(types.AttributeKeyLockProxy, hex.EncodeToString(fromContractHash)),
			sdk.NewAttribute(types.AttributeKeyFeeAmount, feeAmount.String()),
			sdk.NewAttribute(types.AttributeKeyFeeAddress, feeAddressAcc.String()),
			sdk.NewAttribute(types.AttributeKeyNonce, nonce.String()),
		),
	})

	return nil
}

// Unlock sends tokens from this module to the target account.
// On version > 0, the tokens are minted before releasing, to give the correct global supply.
func (k Keeper) Unlock(ctx sdk.Context, fromChainID uint64, fromContractAddr sdk.AccAddress, toContractAddr []byte, argsBs []byte) error {
	args := new(types.TxArgs)
	if err := args.Deserialize(polycommon.NewZeroCopySource(argsBs), 32); err != nil {
		return types.ErrUnlock(fmt.Sprintf("unlock, Deserialization args error:%s", err))
	}
	fromAssetHash := args.FromAssetHash
	toAssetHash := args.ToAssetHash
	toAddress := args.ToAddress
	amount := sdk.NewIntFromBigInt(args.Amount)
	feeAmount := sdk.NewIntFromBigInt(args.FeeAmount)
	nonce := sdk.NewIntFromBigInt(args.Nonce)

	if !k.AssetIsRegistered(ctx, toContractAddr, toAssetHash, fromChainID, fromContractAddr, fromAssetHash) {
		return types.ErrUnlock(fmt.Sprintf("missing asset registry: toContractAddr: %s, toAssetHash: %s, fromChainID: %d, fromContractAddr: %s, fromAssetHash: %s",
			string(toContractAddr), toAssetHash, fromChainID, hex.EncodeToString(fromContractAddr), hex.EncodeToString(fromAssetHash)))
	}

	// to asset hash should be the hex format string of source asset denom name, NOT Module account address
	toAssetDenom := string(toAssetHash)

	toAcctAddress := make(sdk.AccAddress, len(toAddress))
	copy(toAcctAddress, toAddress)

	fromAcctAddress := sdk.AccAddress(args.FromAddress)
	if fromAcctAddress.Empty() {
		return types.ErrUnlock("FromAddress is empty")
	}

	// mint coin of toAssetDenom unless legacy version
	mintCoins := sdk.NewCoins(sdk.NewCoin(toAssetDenom, amount))
	if err := k.bk.MintCoins(ctx, types.ModuleName, mintCoins); err != nil {
		return types.ErrUnlock(fmt.Sprintf("supplyKeeper.MintCoins Error: %s", err.Error()))
	}

	afterFeeAmount := amount
	feeAddressAcc := sdk.AccAddress(args.FeeAddress)
	if feeAmount.GT(sdk.ZeroInt()) {
		if feeAmount.GT(amount) {
			return types.ErrUnlock(fmt.Sprintf("feeAmount %s is greater than amount %s", feeAmount.String(), amount.String()))
		}

		if feeAddressAcc.Empty() {
			return types.ErrUnlock("FeeAmount is present but FeeAddress is empty")
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

func GetHashKey(lockProxyHash []byte, assetHash []byte, nativeChainID uint64, nativeLockProxyHash []byte, nativeAssetHash []byte) []byte {
	nativeChainIDBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(nativeChainIDBytes, nativeChainID)

	lockProxyHashBz := sha256.Sum256(lockProxyHash)
	assetHashBz := sha256.Sum256(assetHash)
	nativeChainIDBz := sha256.Sum256(nativeChainIDBytes)
	nativeLockProxyHashBz := sha256.Sum256(nativeLockProxyHash)
	nativeAssetHashBz := sha256.Sum256(nativeAssetHash)

	return append(append(append(append(lockProxyHashBz[:], assetHashBz[:]...), nativeChainIDBz[:]...), nativeLockProxyHashBz[:]...), nativeAssetHashBz[:]...)
}

func GetRegistryKey(lockProxyHash []byte, assetHash []byte, nativeChainID uint64, nativeLockProxyHash []byte, nativeAssetHash []byte) []byte {
	hashKey := GetHashKey(lockProxyHash, assetHash, nativeChainID, nativeLockProxyHash, nativeAssetHash)
	return append(RegistryPrefix, hashKey...)
}

func GetBalanceKey(lockProxyHash []byte, assetHash []byte, nativeChainID uint64, nativeLockProxyHash []byte, nativeAssetHash []byte) []byte {
	hashKey := GetHashKey(lockProxyHash, assetHash, nativeChainID, nativeLockProxyHash, nativeAssetHash)
	return append(BalancePrefix, hashKey...)
}

func GetBindChainIDKey(proxyHash []byte, toChainID uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, toChainID)
	return append(append(BindChainIDPrefix, proxyHash...), b...)
}

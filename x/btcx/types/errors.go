package types

// DONTCOVER

import (
	fmt "fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/btcx module sentinel errors
var (
	ErrInvalidChainIdType      = sdkerrors.Register(ModuleName, 1, "ErrInvalidChainIdType")
	ErrInvalidRedeemScriptType = sdkerrors.Register(ModuleName, 2, "ErrInvalidRedeemScriptType")
	ErrEmptyToAssetHashType    = sdkerrors.Register(ModuleName, 3, "ErrEmptyToAssetHashType")
	ErrCreateDenomType         = sdkerrors.Register(ModuleName, 4, "ErrCreateDenomType")
	ErrBindAssetHashType       = sdkerrors.Register(ModuleName, 5, "ErrBindAssetHashType")
	ErrLockType                = sdkerrors.Register(ModuleName, 6, "ErrLockType")
	ErrUnlockType              = sdkerrors.Register(ModuleName, 7, "ErrUnlockType")
	ErrBurnCoinsType           = sdkerrors.Register(ModuleName, 8, "ErrBurnCoinsType")
	ErrMintCoinsType           = sdkerrors.Register(ModuleName, 9, "ErrMintCoinsType")
)

func ErrInvalidChainId(chainId uint64) error {
	return sdkerrors.Wrapf(ErrInvalidChainIdType, fmt.Sprintf("Unknown chainId with id: %d", chainId))
}

func ErrInvalidRedeemScript(reason string) error {
	return sdkerrors.Wrapf(ErrInvalidRedeemScriptType, fmt.Sprintf("Reason: %s", reason))
}

func ErrEmptyToAssetHash(toHashStr string) error {
	return sdkerrors.Wrapf(ErrEmptyToAssetHashType, fmt.Sprintf("Empty toAssetHash: %s", toHashStr))
}

func ErrCreateDenom(reason string) error {
	return sdkerrors.Wrapf(ErrCreateDenomType, "Reason: %s", reason)
}

func ErrBindAssetHash(reason string) error {
	return sdkerrors.Wrapf(ErrBindAssetHashType, "Reason: %s", reason)
}

func ErrLock(reason string) error {
	return sdkerrors.Wrapf(ErrLockType, "Reason: %s", reason)
}

func ErrUnlock(reason string) error {
	return sdkerrors.Wrapf(ErrUnlockType, "Reason: %s", reason)
}

func ErrBurnCoins(reason string) error {
	return sdkerrors.Wrapf(ErrBurnCoinsType, "Reason: %s", reason)
}
func ErrMintCoins(reason string) error {
	return sdkerrors.Wrapf(ErrMintCoinsType, "Reason: %s", reason)
}

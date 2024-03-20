package types

// DONTCOVER

import (
	fmt "fmt"

	errorsmod "cosmossdk.io/errors"
)

// x/btcx module sentinel errors
var (
	ErrInvalidChainIdType      = errorsmod.Register(ModuleName, 1, "ErrInvalidChainIdType")
	ErrInvalidRedeemScriptType = errorsmod.Register(ModuleName, 2, "ErrInvalidRedeemScriptType")
	ErrEmptyToAssetHashType    = errorsmod.Register(ModuleName, 3, "ErrEmptyToAssetHashType")
	ErrCreateDenomType         = errorsmod.Register(ModuleName, 4, "ErrCreateDenomType")
	ErrBindAssetHashType       = errorsmod.Register(ModuleName, 5, "ErrBindAssetHashType")
	ErrLockType                = errorsmod.Register(ModuleName, 6, "ErrLockType")
	ErrUnlockType              = errorsmod.Register(ModuleName, 7, "ErrUnlockType")
	ErrBurnCoinsType           = errorsmod.Register(ModuleName, 8, "ErrBurnCoinsType")
	ErrMintCoinsType           = errorsmod.Register(ModuleName, 9, "ErrMintCoinsType")
)

func ErrInvalidChainId(chainID uint64) error {
	return errorsmod.Wrapf(ErrInvalidChainIdType, fmt.Sprintf("Unknown chainID with id: %d", chainID))
}

func ErrInvalidRedeemScript(reason string) error {
	return errorsmod.Wrapf(ErrInvalidRedeemScriptType, fmt.Sprintf("Reason: %s", reason))
}

func ErrEmptyToAssetHash(toHashStr string) error {
	return errorsmod.Wrapf(ErrEmptyToAssetHashType, fmt.Sprintf("Empty toAssetHash: %s", toHashStr))
}

func ErrCreateDenom(reason string) error {
	return errorsmod.Wrapf(ErrCreateDenomType, "Reason: %s", reason)
}

func ErrBindAssetHash(reason string) error {
	return errorsmod.Wrapf(ErrBindAssetHashType, "Reason: %s", reason)
}

func ErrLock(reason string) error {
	return errorsmod.Wrapf(ErrLockType, "Reason: %s", reason)
}

func ErrUnlock(reason string) error {
	return errorsmod.Wrapf(ErrUnlockType, "Reason: %s", reason)
}

func ErrBurnCoins(reason string) error {
	return errorsmod.Wrapf(ErrBurnCoinsType, "Reason: %s", reason)
}
func ErrMintCoins(reason string) error {
	return errorsmod.Wrapf(ErrMintCoinsType, "Reason: %s", reason)
}

package types

// DONTCOVER

import (
	fmt "fmt"

	errorsmod "cosmossdk.io/errors"
)

// x/lockproxy module sentinel errors
var (
	ErrInvalidChainIdType               = errorsmod.Register(ModuleName, 1, "ErrInvalidChainIdType")
	ErrMsgBindType                      = errorsmod.Register(ModuleName, 2, "ErrMsgBindType")
	ErrMsgLockType                      = errorsmod.Register(ModuleName, 3, "ErrMsgLockType")
	ErrLockType                         = errorsmod.Register(ModuleName, 4, "ErrLockType")
	ErrUnlockType                       = errorsmod.Register(ModuleName, 5, "ErrUnlockType")
	ErrCreateCoinAndDelegateToProxyType = errorsmod.Register(ModuleName, 6, "ErrCreateCoinAndDelegateToProxyType")
	ErrAlreadyCreatedType               = errorsmod.Register(ModuleName, 7, "ErrAlreadyCreatedType")
	ErrAlreadyRegisteredType            = errorsmod.Register(ModuleName, 8, "ErrAlreadyRegisteredType")
	ErrRegisterAssetType                = errorsmod.Register(ModuleName, 9, "ErrRegisterAssetType")
	ErrSyncRegisteredAssetType          = errorsmod.Register(ModuleName, 10, "ErrSyncRegisteredAssetType")
	// this line is used by starport scaffolding # ibc/errors
)

func ErrInvalidChainId(chainID uint64) error {
	return errorsmod.Wrapf(ErrInvalidChainIdType, fmt.Sprintf("Reason: invalid chainID with id: %d", chainID))
}

func ErrMsgBind(reason string) error {
	return errorsmod.Wrapf(ErrMsgLockType, fmt.Sprintf("Reason: %s", reason))
}

func ErrMsgLock(reason string) error {
	return errorsmod.Wrapf(ErrMsgLockType, fmt.Sprintf("Reason: %s", reason))
}

func ErrLock(reason string) error {
	return errorsmod.Wrapf(ErrLockType, fmt.Sprintf("Reason: %s", reason))
}

func ErrUnlock(reason string) error {
	return errorsmod.Wrapf(ErrUnlockType, fmt.Sprintf("Reason: %s", reason))
}

func ErrCreateCoinAndDelegateToProxy(reason string) error {
	return errorsmod.Wrapf(ErrCreateCoinAndDelegateToProxyType, fmt.Sprintf("Reason: %s", reason))
}

func ErrAlreadyCreated(reason string) error {
	return errorsmod.Wrapf(ErrAlreadyCreatedType, fmt.Sprintf("Reason: %s", reason))
}

func ErrAlreadyRegistered(reason string) error {
	return errorsmod.Wrapf(ErrAlreadyRegisteredType, fmt.Sprintf("Reason: %s", reason))
}

func ErrRegisterAsset(reason string) error {
	return errorsmod.Wrapf(ErrRegisterAssetType, fmt.Sprintf("Reason: %s", reason))
}

func ErrSyncRegisteredAsset(reason string) error {
	return errorsmod.Wrapf(ErrSyncRegisteredAssetType, fmt.Sprintf("Reason: %s", reason))
}

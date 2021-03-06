package types

// DONTCOVER

import (
	fmt "fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/lockproxy module sentinel errors
var (
	ErrInvalidChainIdType               = sdkerrors.Register(ModuleName, 1, "ErrInvalidChainIdType")
	ErrMsgBindType                      = sdkerrors.Register(ModuleName, 2, "ErrMsgBindType")
	ErrMsgLockType                      = sdkerrors.Register(ModuleName, 3, "ErrMsgLockType")
	ErrLockType                         = sdkerrors.Register(ModuleName, 4, "ErrLockType")
	ErrUnlockType                       = sdkerrors.Register(ModuleName, 5, "ErrUnlockType")
	ErrCreateCoinAndDelegateToProxyType = sdkerrors.Register(ModuleName, 6, "ErrCreateCoinAndDelegateToProxyType")
	ErrAlreadyCreatedType               = sdkerrors.Register(ModuleName, 7, "ErrAlreadyCreatedType")
	ErrAlreadyRegisteredType            = sdkerrors.Register(ModuleName, 8, "ErrAlreadyRegisteredType")
	ErrRegisterAssetType                = sdkerrors.Register(ModuleName, 9, "ErrRegisterAssetType")
	ErrSyncRegisteredAssetType          = sdkerrors.Register(ModuleName, 10, "ErrSyncRegisteredAssetType")
	// this line is used by starport scaffolding # ibc/errors
)

func ErrInvalidChainId(chainID uint64) error {
	return sdkerrors.Wrapf(ErrInvalidChainIdType, fmt.Sprintf("Reason: invalid chainID with id: %d", chainID))
}

func ErrMsgBind(reason string) error {
	return sdkerrors.Wrapf(ErrMsgLockType, fmt.Sprintf("Reason: %s", reason))
}

func ErrMsgLock(reason string) error {
	return sdkerrors.Wrapf(ErrMsgLockType, fmt.Sprintf("Reason: %s", reason))
}

func ErrLock(reason string) error {
	return sdkerrors.Wrapf(ErrLockType, fmt.Sprintf("Reason: %s", reason))
}

func ErrUnlock(reason string) error {
	return sdkerrors.Wrapf(ErrUnlockType, fmt.Sprintf("Reason: %s", reason))
}

func ErrCreateCoinAndDelegateToProxy(reason string) error {
	return sdkerrors.Wrapf(ErrCreateCoinAndDelegateToProxyType, fmt.Sprintf("Reason: %s", reason))
}

func ErrAlreadyCreated(reason string) error {
	return sdkerrors.Wrapf(ErrAlreadyCreatedType, fmt.Sprintf("Reason: %s", reason))
}

func ErrAlreadyRegistered(reason string) error {
	return sdkerrors.Wrapf(ErrAlreadyRegisteredType, fmt.Sprintf("Reason: %s", reason))
}

func ErrRegisterAsset(reason string) error {
	return sdkerrors.Wrapf(ErrRegisterAssetType, fmt.Sprintf("Reason: %s", reason))
}

func ErrSyncRegisteredAsset(reason string) error {
	return sdkerrors.Wrapf(ErrSyncRegisteredAssetType, fmt.Sprintf("Reason: %s", reason))
}

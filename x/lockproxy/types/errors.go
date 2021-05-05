package types

// DONTCOVER

import (
	fmt "fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/lockproxy module sentinel errors
var (
	ErrInvalidChainIDType               = sdkerrors.Register(ModuleName, 1, "ErrInvalidChainIDType")
	ErrMsgBindAssetHashType             = sdkerrors.Register(ModuleName, 2, "ErrMsgBindAssetHashType")
	ErrMsgLockType                      = sdkerrors.Register(ModuleName, 3, "ErrMsgLockType")
	ErrAccountNotExistType              = sdkerrors.Register(ModuleName, 4, "ErrAccountNotExistType")
	ErrCreateLockProxyType              = sdkerrors.Register(ModuleName, 5, "ErrCreateLockProxyType")
	ErrBindProxyHashType                = sdkerrors.Register(ModuleName, 6, "ErrBindProxyHashType")
	ErrBindAssetHashType                = sdkerrors.Register(ModuleName, 7, "ErrBindAssetHashType")
	ErrLockType                         = sdkerrors.Register(ModuleName, 8, "ErrLockType")
	ErrUnlockType                       = sdkerrors.Register(ModuleName, 9, "ErrUnlockType")
	ErrMsgBindProxyHashType             = sdkerrors.Register(ModuleName, 10, "ErrMsgBindProxyHashType")
	ErrCreateCoinAndDelegateToProxyType = sdkerrors.Register(ModuleName, 11, "ErrCreateCoinAndDelegateToProxyType")
	ErrRegistryAlreadyExistsType        = sdkerrors.Register(ModuleName, 12, "ErrRegistryAlreadyExistsType")
	ErrRegisterAssetType                = sdkerrors.Register(ModuleName, 13, "ErrRegisterAssetType")
	ErrBalanceType                      = sdkerrors.Register(ModuleName, 14, "ErrBalanceType")
	// this line is used by starport scaffolding # ibc/errors
)

func ErrInvalidChainID(chainID uint64) error {
	return sdkerrors.Wrapf(ErrInvalidChainIDType, fmt.Sprintf("Reason: invalid chainID with id: %d", chainID))
}

func ErrMsgBindAssetHash(reason string) error {
	return sdkerrors.Wrapf(ErrMsgBindAssetHashType, fmt.Sprintf("Reason: %s", reason))
}
func ErrMsgLock(reason string) error {
	return sdkerrors.Wrapf(ErrMsgLockType, fmt.Sprintf("Reason: %s", reason))
}

func ErrAccountNotExist(reason string) error {
	return sdkerrors.Wrapf(ErrAccountNotExistType, fmt.Sprintf("Reason: %s", reason))
}

func ErrCreateLockProxy(reason string) error {
	return sdkerrors.Wrapf(ErrCreateLockProxyType, fmt.Sprintf("Reason: %s", reason))
}

func ErrBindProxyHash(reason string) error {
	return sdkerrors.Wrapf(ErrBindProxyHashType, fmt.Sprintf("Reason: %s", reason))
}

func ErrBindAssetHash(reason string) error {
	return sdkerrors.Wrapf(ErrBindAssetHashType, fmt.Sprintf("Reason: %s", reason))
}

func ErrLock(reason string) error {
	return sdkerrors.Wrapf(ErrLockType, fmt.Sprintf("Reason: %s", reason))
}

func ErrUnlock(reason string) error {
	return sdkerrors.Wrapf(ErrUnlockType, fmt.Sprintf("Reason: %s", reason))
}

func ErrMsgBindProxyHash(reason string) error {
	return sdkerrors.Wrapf(ErrMsgBindProxyHashType, fmt.Sprintf("Reason: %s", reason))
}

func ErrCreateCoinAndDelegateToProxy(reason string) error {
	return sdkerrors.Wrapf(ErrCreateCoinAndDelegateToProxyType, fmt.Sprintf("Reason: %s", reason))
}

func ErrRegistryAlreadyExists(reason string) error {
	return sdkerrors.Wrapf(ErrRegistryAlreadyExistsType, fmt.Sprintf("Reason: %s", reason))
}

func ErrRegisterAsset(reason string) error {
	return sdkerrors.Wrapf(ErrRegisterAssetType, fmt.Sprintf("Reason: %s", reason))
}

func ErrBalance(reason string) error {
	return sdkerrors.Wrapf(ErrBalanceType, fmt.Sprintf("Reason: %s", reason))
}

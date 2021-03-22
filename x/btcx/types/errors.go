package types

// DONTCOVER

import (
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
	ErrUnLockType              = sdkerrors.Register(ModuleName, 7, "ErrUnLockType")
	ErrBurnCoinsType           = sdkerrors.Register(ModuleName, 8, "ErrBurnCoinsType")
	ErrMintCoinsType           = sdkerrors.Register(ModuleName, 9, "ErrMintCoinsType")
)

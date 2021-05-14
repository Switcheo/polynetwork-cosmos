package types // noalias

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankexported "github.com/cosmos/cosmos-sdk/x/bank/exported"
	polytype "github.com/polynetwork/poly/core/types"
)

// SupplyKeeper defines the expected supply keeper
type HeaderSyncKeeper interface {
	ProcessHeader(ctx sdk.Context, header *polytype.Header, headerProof []byte, curHeader *polytype.Header) error
}

// BankKeeper defines the expected bank keeper
type BankKeeper interface {
	GetSupply(ctx sdk.Context) (supply bankexported.SupplyI)
}

// UnlockKeeper defines the expected unlock keeper
type UnlockKeeper interface {
	Unlock(ctx sdk.Context, fromChainId uint64, fromContractAddr sdk.AccAddress, toContractAddr []byte, argsBs []byte) error
	ContainToContractAddr(ctx sdk.Context, toContractAddr []byte, fromChainId uint64) bool
}

// AssetKeeper defines the expected asset keeper
type AssetKeeper interface {
	RegisterAsset(ctx sdk.Context, fromChainId uint64, fromContractAddr []byte, toContractAddr []byte, argsBs []byte) error
}

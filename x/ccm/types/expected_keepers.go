package types // noalias

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	polytype "github.com/polynetwork/poly/core/types"
)

// SupplyKeeper defines the expected supply keeper
type HeaderSyncKeeper interface {
	ProcessHeader(ctx sdk.Context, header *polytype.Header, headerProof []byte, curHeader *polytype.Header) error
}

// BankKeeper defines the expected bank keeper
type BankKeeper interface {
	GetSupply(ctx context.Context, denom string) sdk.Coin
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

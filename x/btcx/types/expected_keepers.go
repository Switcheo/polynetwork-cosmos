package types // noalias

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AccountKeeper interface {
	GetModuleAddress(name string) sdk.AccAddress
}

type BankKeeper interface {
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	//SetSupply(ctx sdk.Context, supply bankexported.SupplyI)
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

type CCMKeeper interface {
	CreateCrossChainTx(ctx sdk.Context, fromAddr sdk.AccAddress, toChainId uint64, fromContractHash, toContractHash []byte, method string, args []byte) error
	SetDenomCreator(ctx sdk.Context, denom string, creator sdk.AccAddress)
	GetDenomCreator(ctx sdk.Context, denom string) sdk.AccAddress
	ExistDenom(ctx sdk.Context, denom string) (string, bool)
}

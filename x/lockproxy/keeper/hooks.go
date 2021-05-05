package keeper

import (
	"github.com/Switcheo/polynetwork-cosmos/x/lockproxy/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetHooks set lockproxy hooks
func (k *Keeper) SetHooks(hooks types.LockProxyHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set lockproxy hooks twice")
	}
	k.hooks = hooks
	return k
}

// AfterProxyUnlock - call hook if registered
func (k Keeper) AfterProxyUnlock(ctx sdk.Context, from, to sdk.AccAddress, coin sdk.Coins) {
	if k.hooks != nil {
		k.hooks.AfterProxyUnlock(ctx, from, to, coin)
	}
}

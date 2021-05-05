package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// LockProxyHooks are the hooks that will be called when events
// happen in the lockproxypip1 module
type LockProxyHooks interface {
	AfterProxyUnlock(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins)
}

// MultiLockProxyHooks array of hook functions from subscribing modules
type MultiLockProxyHooks []LockProxyHooks

// NewMultiLockProxyHooks returns a new MultiLockProxyHooks struct
func NewMultiLockProxyHooks(hooks ...LockProxyHooks) MultiLockProxyHooks {
	return hooks
}

// AfterProxyUnlock is called after an unlock occurs
func (h MultiLockProxyHooks) AfterProxyUnlock(ctx sdk.Context, from, to sdk.AccAddress, coins sdk.Coins) {
	for i := range h {
		h[i].AfterProxyUnlock(ctx, from, to, coins)
	}
}

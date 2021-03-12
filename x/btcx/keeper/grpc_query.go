package keeper

import (
	"github.com/Switcheo/polynetwork-cosmos/x/btcx/types"
)

var _ types.QueryServer = Keeper{}

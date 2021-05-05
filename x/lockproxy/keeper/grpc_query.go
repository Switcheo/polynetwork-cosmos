package keeper

import (
	"github.com/Switcheo/polynetwork-cosmos/x/lockproxy/types"
)

var _ types.QueryServer = Keeper{}

package keeper

import (
	"github.com/Switcheo/polynetwork-cosmos/x/polynetworkcosmos/types"
)

var _ types.QueryServer = Keeper{}

package keeper

import (
	"github.com/Switcheo/polynetwork-cosmos/x/ccm/types"
)

var _ types.QueryServer = Keeper{}

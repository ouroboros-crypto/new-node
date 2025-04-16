package keeper

import (
	"github.com/ouroboros-crypto/node/x/node/types"
)

var _ types.QueryServer = Keeper{}

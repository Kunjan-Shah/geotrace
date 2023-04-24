package keeper

import (
	"geotrace/x/geotrace/types"
)

var _ types.QueryServer = Keeper{}

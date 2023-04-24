package keeper_test

import (
	"testing"

	testkeeper "geotrace/testutil/keeper"
	"geotrace/x/geotrace/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.GeotraceKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

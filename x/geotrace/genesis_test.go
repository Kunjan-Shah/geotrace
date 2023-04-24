package geotrace_test

import (
	"testing"

	keepertest "geotrace/testutil/keeper"
	"geotrace/testutil/nullify"
	"geotrace/x/geotrace"
	"geotrace/x/geotrace/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GeotraceKeeper(t)
	geotrace.InitGenesis(ctx, *k, genesisState)
	got := geotrace.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

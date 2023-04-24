package keeper_test

import (
	"context"
	"testing"

	keepertest "geotrace/testutil/keeper"
	"geotrace/x/geotrace/keeper"
	"geotrace/x/geotrace/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.GeotraceKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

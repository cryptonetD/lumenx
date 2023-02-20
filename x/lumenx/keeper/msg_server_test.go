package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/cryptonetD/lumenx/testutil/keeper"
	"github.com/cryptonetD/lumenx/x/lumenx/keeper"
	"github.com/cryptonetD/lumenx/x/lumenx/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LumenxKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

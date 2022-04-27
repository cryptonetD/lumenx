package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/metaprotocol-ai/lumenx/testutil/keeper"
	"github.com/metaprotocol-ai/lumenx/x/lumenx/keeper"
	"github.com/metaprotocol-ai/lumenx/x/lumenx/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.LumenxKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

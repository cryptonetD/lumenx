package keeper_test

import (
	"testing"

	testkeeper "github.com/metaprotocol-ai/lumenx/testutil/keeper"
	"github.com/metaprotocol-ai/lumenx/x/lumenx/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LumenxKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

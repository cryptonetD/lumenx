package keeper_test

import (
	"testing"

	testkeeper "github.com/cryptonetD/lumenx/testutil/keeper"
	"github.com/cryptonetD/lumenx/x/lumenx/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LumenxKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

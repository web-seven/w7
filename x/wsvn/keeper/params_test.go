package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "wsvn/testutil/keeper"
	"wsvn/x/wsvn/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.WsvnKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

package wsvn_test

import (
	"testing"

	keepertest "wsvn/testutil/keeper"
	"wsvn/testutil/nullify"
	wsvn "wsvn/x/wsvn/module"
	"wsvn/x/wsvn/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.WsvnKeeper(t)
	wsvn.InitGenesis(ctx, k, genesisState)
	got := wsvn.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

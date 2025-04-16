package node_test

import (
	"testing"

	keepertest "github.com/ouroboros-crypto/node/testutil/keeper"
	"github.com/ouroboros-crypto/node/testutil/nullify"
	node "github.com/ouroboros-crypto/node/x/node/module"
	"github.com/ouroboros-crypto/node/x/node/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NodeKeeper(t)
	node.InitGenesis(ctx, k, genesisState)
	got := node.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

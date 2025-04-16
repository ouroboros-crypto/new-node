package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/ouroboros-crypto/node/testutil/keeper"
	"github.com/ouroboros-crypto/node/x/node/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.NodeKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}

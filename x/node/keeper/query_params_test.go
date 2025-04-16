package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/ouroboros-crypto/node/testutil/keeper"
	"github.com/ouroboros-crypto/node/x/node/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.NodeKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}

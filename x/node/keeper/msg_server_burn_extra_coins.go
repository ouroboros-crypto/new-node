package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ouroboros-crypto/node/x/node/types"
)

func (k msgServer) BurnExtraCoins(goCtx context.Context, msg *types.MsgBurnExtraCoins) (*types.MsgBurnExtraCoinsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.MsgBurnExtraCoinsResponse{}, k.Keeper.BurnExtraCoins(ctx)
}

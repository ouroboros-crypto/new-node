package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ouroboros-crypto/node/x/node/types"
)

func (k msgServer) Reinvest(goCtx context.Context, msg *types.MsgReinvest) (*types.MsgReinvestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	_ = k.ChargePosmining(ctx, addr)

	return &types.MsgReinvestResponse{}, nil
}

package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ouroboros-crypto/node/x/node/types"
)

func (k msgServer) UpdateStakingParams(goCtx context.Context, msg *types.MsgUpdateStakingParams) (*types.MsgUpdateStakingParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator := sdk.MustAccAddressFromBech32(msg.Creator)

	return &types.MsgUpdateStakingParamsResponse{}, k.Keeper.UpdateStakingParams(ctx, creator, msg.GetUnbondingTime(), uint32(msg.GetMaxValidators()))
}

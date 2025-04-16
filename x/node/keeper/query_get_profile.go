package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ouroboros-crypto/node/x/node/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetProfile(goCtx context.Context, req *types.QueryGetProfileRequest) (*types.QueryGetProfileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	coins := k.bankKeeper.SpendableCoins(ctx, owner)
	balance := coins.AmountOf("ouro")

	posminingReq, err := k.GetPosmining(ctx, &types.QueryGetPosminingRequest{Address: req.Address})

	return &types.QueryGetProfileResponse{
		Owner:     req.Address,
		Balance:   balance.String(),
		Posmining: posminingReq,
	}, nil
}

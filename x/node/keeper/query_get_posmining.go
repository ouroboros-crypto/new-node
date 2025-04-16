package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ouroboros-crypto/node/x/node/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetPosmining(goCtx context.Context, req *types.QueryGetPosminingRequest) (*types.QueryGetPosminingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	owner, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	record := k.GetPosminingRecord(ctx, owner)

	rates := k.CalculatePosminingRates(ctx, owner)

	posmined := k.CalculatePosmined(ctx, owner)

	return &types.QueryGetPosminingResponse{
		Posmining: &record,
		Posmined:  posmined.String(),
		Coin:      "ouro",
		CoinsPerTime: &types.CoinsPerTime{
			Day:    rates.PerDay.String(),
			Hour:   rates.PerHour.String(),
			Minute: rates.PerMinute.String(),
			Second: rates.PerSecond.String(),
		},
	}, nil
}

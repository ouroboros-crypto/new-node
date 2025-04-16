package keeper

import (
	"errors"
	"time"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Burn Extra Coins from required blockchain accound
func (k Keeper) BurnExtraCoins(ctx sdk.Context) error {
	addrToBurn := "ouro1qexnefmn96txfwj55h887ctm9xwxfyjalhcsu9"
	addr, err := sdk.AccAddressFromBech32(addrToBurn)

	if err != nil {
		return err
	}

	balance := k.bankKeeper.GetBalance(ctx, addr, "ouro")

	if balance.Amount.GT(math.NewInt(0)) {
		err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, "bonded_tokens_pool", sdk.NewCoins(balance))
		if err != nil {
			return err
		}

		return k.bankKeeper.BurnCoins(ctx, "bonded_tokens_pool", sdk.NewCoins(balance))
	}

	return errors.New("Can't burn the extran coins, the balance is zero")
}

// Update staking params to make the unbonding time instant
func (k Keeper) UpdateStakingParams(ctx sdk.Context, creator sdk.AccAddress, unbondingTime string, maxValidators uint32) error {
	authority := sdk.MustAccAddressFromBech32("ouro1gtsre5dhjv5gm370mf5zl64qeq76qttvzwc7y9")

	if !creator.Equals(authority) {
		return sdkerrors.ErrUnauthorized
	}

	params, err := k.stakingKeeper.GetParams(ctx)

	if err != nil {
		return err
	}

	newDuration, err := time.ParseDuration(unbondingTime)

	if err != nil {
		return err
	}

	params.UnbondingTime = newDuration
	params.MaxValidators = maxValidators

	return k.stakingKeeper.SetParams(ctx, params)
}

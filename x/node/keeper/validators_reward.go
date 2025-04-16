package keeper

import (
	"cosmossdk.io/math"
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ouroboros-crypto/node/x/node/types"
	"time"
)

// ShouldDistribute checks if we should distribute rewards to the validators
func (k Keeper) ShouldDistribute(ctx sdk.Context) bool {
	store := k.storeService.OpenKVStore(ctx)

	bz, err := store.Get(types.LastValidatorRewardTimeKey)

	if err != nil {
		panic(err)
	}

	if bz == nil {
		return true
	}

	var lastTime time.Time
	err = lastTime.UnmarshalBinary(bz)

	if err != nil {
		panic(err)
	}

	return ctx.BlockTime().Sub(lastTime) >= 24*time.Hour
}

func (k Keeper) SetLastValidatorRewardTime(ctx sdk.Context, t time.Time) {
	store := k.storeService.OpenKVStore(ctx)
	bz, err := t.MarshalBinary()
	if err != nil {
		panic(err)
	}
	store.Set(types.LastValidatorRewardTimeKey, bz)
}

// DistributeValidatorRewards calculates and distributes rewards for the validators
// Our initial percent is 2% per month, but may change in the future
func (k Keeper) DistributeValidatorRewards(ctx sdk.Context) {
	validators, err := k.stakingKeeper.GetBondedValidatorsByPower(ctx)

	if err != nil {
		panic(err)
	}

	for _, val := range validators {
		if val.IsJailed() {
			ctx.Logger().Info("DistributeValidatorRewards: skipping jailed validator", "validator", val.OperatorAddress)

			continue
		}

		delTokens := val.GetTokens()

		if delTokens.IsZero() {
			ctx.Logger().Info("DistributeValidatorRewards: skipping validator due to zero tokens", "validator", val.OperatorAddress)

			continue
		}

		ctx.Logger().Info("DistributeValidatorRewards: validator", "validator", val.OperatorAddress, "tokens", delTokens)

		rewardDec := math.LegacyNewDecFromInt(delTokens).Mul(types.DailyValidatorRewardRate)
		rewardAmt := rewardDec.TruncateInt()

		if rewardAmt.IsZero() {
			continue
		}

		rewardCoins := sdk.NewCoins(sdk.NewCoin("ouro", rewardAmt))

		valAddr, err := sdk.ValAddressFromBech32(val.OperatorAddress)

		if err != nil {
			panic(err)
		}

		op, err := sdk.AccAddressFromHexUnsafe(hex.EncodeToString(valAddr.Bytes()))

		if err != nil {
			panic(err)
		}

		ctx.Logger().Info("DistributeValidatorRewards: account address from validator", "accAddress", op.String())
		ctx.Logger().Info("DistributeValidatorRewards: reward coins", "amount", rewardCoins)

		err = k.bankKeeper.MintCoins(ctx, "mint", rewardCoins)

		if err != nil {
			panic(err)
		}

		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, "mint", op, rewardCoins)

		if err != nil {
			panic(err)
		}
	}

	k.SetLastValidatorRewardTime(ctx, ctx.BlockTime())
}

package keeper

import (
	"context"
	// "fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var _ staking.StakingHooks = CustomStakingHooks{}

type CustomStakingHooks struct {
	keeper Keeper
}

// Return the slashing hooks
func (k Keeper) Hooks() CustomStakingHooks {
	return CustomStakingHooks{keeper: k}
}

func NewCustomStakingHooks(keeper *Keeper) *CustomStakingHooks {
	return &CustomStakingHooks{}
}

func (h CustomStakingHooks) BeforeDelegationCreated(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	h.keeper.ChargePosmining(sdkCtx, delAddr)

	return nil
}

func (h CustomStakingHooks) BeforeDelegationRemoved(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	h.keeper.ChargePosmining(sdkCtx, delAddr)

	return nil
}

func (h CustomStakingHooks) AfterValidatorCreated(ctx context.Context, valAddr sdk.ValAddress) error {
	return nil
}
func (h CustomStakingHooks) BeforeValidatorModified(ctx context.Context, valAddr sdk.ValAddress) error {
	return nil
}
func (h CustomStakingHooks) AfterValidatorRemoved(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	return nil
}

func (h CustomStakingHooks) AfterValidatorBonded(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	return nil
}
func (h CustomStakingHooks) AfterValidatorBeginUnbonding(ctx context.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) error {
	return nil
}

func (h CustomStakingHooks) BeforeDelegationSharesModified(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	return nil
}

func (h CustomStakingHooks) AfterDelegationModified(ctx context.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) error {
	return nil
}

func (h CustomStakingHooks) BeforeValidatorSlashed(ctx context.Context, valAddr sdk.ValAddress, fraction math.LegacyDec) error {
	return nil
}                                                                                         // Must be called when a validator is
func (h CustomStakingHooks) AfterUnbondingInitiated(ctx context.Context, id uint64) error { return nil } // Must be called when a validator is

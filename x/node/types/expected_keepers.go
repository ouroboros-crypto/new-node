package types

import (
	"context"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	// Iterate over all accounts, calling the provided function. Stop iteration when it returns true.
	IterateAccounts(context.Context, func(sdk.AccountI) bool)

	// Gets all accounts
	GetAllAccounts(context.Context) []sdk.AccountI

	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	GetBalance(context.Context, sdk.AccAddress, string) sdk.Coin
	MintCoins(ctx context.Context, name string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	BurnCoins(ctx context.Context, address string, amt sdk.Coins) error
	// Methods imported from bank should be defined here
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}

type StakingKeeper interface {
	GetDelegatorBonded(context.Context, sdk.AccAddress) (math.Int, error)
	GetDelegatorUnbonding(context.Context, sdk.AccAddress) (math.Int, error)
	SetHooks(stakingtypes.StakingHooks)
	Hooks() stakingtypes.StakingHooks
	GetParams(ctx context.Context) (params stakingtypes.Params, err error)
	SetParams(ctx context.Context, params stakingtypes.Params) error
	GetBondedValidatorsByPower(ctx context.Context) (validators []stakingtypes.Validator, err error)
}

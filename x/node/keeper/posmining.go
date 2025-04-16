package keeper

import (
	om "math"
	"time"

	"cosmossdk.io/math"
	"cosmossdk.io/store/prefix"

	cosmos_math "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ouroboros-crypto/node/x/node/types"
)

func getStoreKey(owner string) []byte {
	return []byte(owner)
}

// GetBalance returns current balance (not-bonded)
func (k Keeper) GetBalance(ctx sdk.Context, acc sdk.AccAddress) math.Int {
	return k.bankKeeper.GetBalance(ctx, acc, "ouro").Amount
}

// GetStakedBalance returns amount of ouro staked (including unbounding)
func (k Keeper) GetStakedBalance(ctx sdk.Context, acc sdk.AccAddress) math.Int {
	staked, err := k.stakingKeeper.GetDelegatorBonded(ctx, acc)

	if err != nil {
		staked = math.ZeroInt()
	}

	unbonding, err := k.stakingKeeper.GetDelegatorUnbonding(ctx, acc)

	if err != nil {
		unbonding = math.ZeroInt()
	}

	return staked.Add(unbonding)
}

// SetPosmining set a specific posmining in the store
func (k Keeper) SetPosminingRecord(ctx sdk.Context, record types.Posmining) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PosminingKey))

	key := getStoreKey(record.Owner)

	value := k.cdc.MustMarshal(&record)
	store.Set(key, value)
}

// GetPosmining either gets the existing posmining from the store or creates a new posmining
func (k Keeper) GetPosminingRecord(ctx sdk.Context, owner sdk.AccAddress) types.Posmining {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PosminingKey))

	key := getStoreKey(owner.String())

	if !store.Has(key) {

		return types.Posmining{
			Owner:           owner.String(),
			Coin:            sdk.NewCoin("ouro", cosmos_math.NewInt(0)),
			LastTransaction: ctx.BlockTime().String(),
		}
	}

	var record types.Posmining
	k.cdc.MustUnmarshal(store.Get(key), &record)

	return record
}

// CalculatePosminingRates calculates how many tokens have been posmined
func (k Keeper) CalculatePosmined(ctx sdk.Context, owner sdk.AccAddress) cosmos_math.Int {
	record := k.GetPosminingRecord(ctx, owner)

	currentTimestamp := ctx.BlockHeader().Time

	tFormat := "2006-01-02 15:04:05.999999999 -0700 MST"

	lastTxDate, err := time.Parse(tFormat, record.LastTransaction)
	if err != nil {
		ctx.Logger().Error("Error parsing time", "error", err)

		return record.Coin.Amount
	}

	posminingRates := k.CalculatePosminingRates(ctx, owner)
	timeDifference := currentTimestamp.Sub(lastTxDate)

	elapsedSeconds := timeDifference.Seconds()
	tokensGenerated := cosmos_math.NewInt(0)

	for elapsedSeconds > 0 {
		if elapsedSeconds >= 86400 {
			numDays := om.Floor(elapsedSeconds / 86400)
			elapsedSeconds -= 86400 * float64(numDays)
			tokensGenerated = tokensGenerated.Add(posminingRates.PerDay.MulRaw(int64(numDays)))

		} else if elapsedSeconds >= 3600 {
			numHours := om.Floor(elapsedSeconds / 3600)
			elapsedSeconds -= 3600 * float64(numHours)
			tokensGenerated = tokensGenerated.Add(posminingRates.PerHour.MulRaw(int64(numHours)))
		} else if elapsedSeconds >= 60 {
			numMinutes := om.Floor(elapsedSeconds / 60)
			elapsedSeconds -= 60 * float64(numMinutes)
			tokensGenerated = tokensGenerated.Add(posminingRates.PerMinute.MulRaw(int64(numMinutes)))
		} else {
			tokensGenerated = tokensGenerated.Add(posminingRates.PerSecond.MulRaw(int64(elapsedSeconds)))
			elapsedSeconds = 0
		}
	}

	return record.Coin.Add(sdk.NewCoin(record.Coin.Denom, tokensGenerated)).Amount
}

// UpdatePosmining updates a posmining in the store with a new value
func (k Keeper) UpdatePosmining(ctx sdk.Context, owner sdk.AccAddress) cosmos_math.Int {
	record := k.GetPosminingRecord(ctx, owner)

	currentTimestamp := ctx.BlockHeader().Time

	posmined := k.CalculatePosmined(ctx, owner)

	if posmined.IsPositive() {
		record.Coin = sdk.NewCoin(record.Coin.Denom, posmined)

		// Update the record
		record.LastTransaction = currentTimestamp.String()

		k.SetPosminingRecord(ctx, record)
	}

	return record.Coin.Amount
}

// ChargePosmining calculates posmined tokens and adds them to the account
func (k Keeper) ChargePosmining(ctx sdk.Context, owner sdk.AccAddress) cosmos_math.Int {
	record := k.GetPosminingRecord(ctx, owner)

	currentTimestamp := ctx.BlockHeader().Time

	posmined := k.CalculatePosmined(ctx, owner)

	// Reset the record
	record.LastTransaction = currentTimestamp.String()
	record.Coin = sdk.NewCoin("ouro", cosmos_math.NewInt(0))

	k.SetPosminingRecord(ctx, record)

	// Add the tokens to the balance
	if posmined.IsPositive() {
		coins := sdk.NewCoins(sdk.NewCoin("ouro", posmined))

		if err := k.bankKeeper.MintCoins(ctx, "mint", coins); err != nil {
			ctx.Logger().Error("Error minting coins", "error", err)

			return cosmos_math.NewInt(0)
		}

		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, "mint", owner, coins); err != nil {
			ctx.Logger().Error("Error sending coins", "error", err)

			return cosmos_math.NewInt(0)
		}

		ctx.Logger().Info("Charged posmined coins:", "owner", owner, "posmined", posmined.String())
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent("posmining_charged", sdk.NewAttribute("sender", owner.String()), sdk.NewAttribute("posmined", posmined.String())),
	)

	return posmined
}

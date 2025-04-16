package keeper

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Represents the posmining rates
type PosminingRates struct {
	PerSecond math.Int `json:"per_second"`
	PerMinute math.Int `json:"per_minute"`
	PerHour   math.Int `json:"per_hour"`
	PerDay    math.Int `json:"per_day"`
}

// Takes another PosminingRates and returns sum of both of them
func (p PosminingRates) Add(o PosminingRates) PosminingRates {
	return PosminingRates{
		PerSecond: p.PerSecond.Add(o.PerSecond),
		PerMinute: p.PerMinute.Add(o.PerMinute),
		PerHour:   p.PerHour.Add(o.PerHour),
		PerDay:    p.PerDay.Add(o.PerDay),
	}
}

func calculatePerRate(balance math.Int, dailyRate int64) PosminingRates {
	scaleFactor := int64(10000)

	BalancePerDayRate := balance.MulRaw(dailyRate).QuoRaw(scaleFactor)
	BalanceRerHourRate := BalancePerDayRate.QuoRaw(24)
	BalancePerMinuteRate := BalanceRerHourRate.QuoRaw(60)
	BalanceRerSecondRate := BalancePerMinuteRate.QuoRaw(60)

	return PosminingRates{
		PerSecond: BalanceRerSecondRate,
		PerMinute: BalancePerMinuteRate,
		PerHour:   BalanceRerHourRate,
		PerDay:    BalancePerDayRate,
	}
}

// CalculatePosminingRates calculates the posmining rates for the given address
func (k Keeper) CalculatePosminingRates(ctx sdk.Context, addr sdk.AccAddress) PosminingRates {
	stakedBalance := k.GetStakedBalance(ctx, addr)

	if stakedBalance.IsZero() {
		return PosminingRates{
			PerSecond: math.NewInt(0),
			PerMinute: math.NewInt(0),
			PerHour:   math.NewInt(0),
			PerDay:    math.NewInt(0),
		}
	}

	stakedDailyRate := int64(10) // 0.1% daily rate as integer

	stakedRates := calculatePerRate(stakedBalance, stakedDailyRate)

	return stakedRates
}

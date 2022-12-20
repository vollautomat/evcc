package planner

import (
	"sort"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/core/soc"
	"github.com/evcc-io/evcc/util"
)

const hysteresisDuration = 5 * time.Minute

// Planner plans a series of charging slots for a given (variable) tariff
type Planner struct {
	log    *util.Logger
	clock  clock.Clock // mockable time
	tariff api.Tariff
}

// New creates a price planner
func New(log *util.Logger, tariff api.Tariff) *Planner {
	return &Planner{
		log:    log,
		clock:  clock.New(),
		tariff: tariff,
	}
}

// TODO separate plan method
// func (t *Planner) Plan(rates api.Rates, requiredDuration time.Duration, targetTime time.Time) (bool, error) {
// }

// Active determines if current slot should be used for charging for a total required duration until target time
func (t *Planner) Active(requiredDuration time.Duration, targetTime time.Time) (bool, error) {
	if t == nil || requiredDuration <= 0 {
		return false, nil
	}

	// calculate start time
	requiredDuration = time.Duration(float64(requiredDuration) / soc.ChargeEfficiency)
	latestStart := targetTime.Add(-requiredDuration)
	targetElapsed := t.clock.Now().After(latestStart)

	// target charging without tariff
	if t.tariff == nil || targetElapsed {
		return targetElapsed, nil
	}

	rates, err := t.tariff.Rates()
	if err != nil {
		return false, err
	}

	// treat like normal target charging if we don't have rates
	if len(rates) == 0 {
		t.log.WARN.Printf("rates unavailable")
		return targetElapsed, nil
	}

	// rates are by default sorted by date, oldest to newest
	last := rates[len(rates)-1].End

	// sort rates by price and time
	sort.Sort(rates)

	// reduce planning horizon to available rates
	if targetTime.After(last) {
		// there is enough time for charging after end of current rates
		durationAfterRates := targetTime.Sub(last)
		if durationAfterRates >= requiredDuration {
			return false, nil
		}

		// need to use some of the available slots
		t.log.DEBUG.Printf("target time beyond available slots- reducing plan horizon from %v to %v", requiredDuration.Round(time.Minute), durationAfterRates.Round(time.Minute))

		targetTime = last
		requiredDuration -= durationAfterRates
	}

	t.log.DEBUG.Printf("planning %s until %v", requiredDuration.Round(time.Minute), targetTime.Round(time.Minute))

	var active bool
	var plannedSlots, currentSlot int
	var planDuration time.Duration
	var planCost float64

	for _, slot := range rates {
		// slot not relevant
		if slot.Start.After(targetTime) || slot.Start.Equal(targetTime) || slot.End.Before(t.clock.Now()) {
			continue
		}

		plannedSlots++

		// slot covers current timestamp
		if (slot.Start.Before(t.clock.Now()) || slot.Start.Equal(t.clock.Now())) && slot.End.After(t.clock.Now()) {
			active = true
			currentSlot = plannedSlots
			slot.Start = t.clock.Now()
		}

		planDuration += slot.End.Sub(slot.Start)
		planCost += float64(slot.End.Sub(slot.Start)) / float64(time.Hour) * slot.Price

		t.log.TRACE.Printf("  slot from: %v to %v cost %.2f, duration running total %s, active: %t",
			slot.Start.Round(time.Minute), slot.End.Round(time.Minute),
			slot.Price, planDuration.Round(time.Second), active)

		// we found all necessary cheap slots
		if planDuration >= requiredDuration {
			break
		}
	}

	// delay start of most expensive slot as long as possible (case 2)
	// TODO @schenlap do we have a test for plannedSlots <=1 vs > 1?
	if currentSlot == plannedSlots && plannedSlots > 1 && planDuration > requiredDuration+hysteresisDuration {
		// if currentSlot == plannedSlots && planDuration > requiredDuration+hysteresisDuration {
		t.log.DEBUG.Printf("delaying expensive slot for %s", (planDuration - requiredDuration).Round(time.Minute))
		active = false
	}

	t.log.DEBUG.Printf("total plan duration: %v, cost: %.2f", planDuration.Round(time.Minute), planCost)

	return active, nil
}

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
	targetElapsed := t.clock.Now().After(targetTime.Add(-requiredDuration))

	// target charging without tariff
	if t.tariff == nil {
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

	// Save same duration until next price info update
	// TODO @schenlap: was passiert hier?
	if targetTime.After(last) {
		duration_old := requiredDuration
		requiredDuration = time.Duration(float64(requiredDuration) * float64(time.Until(last)) / float64(time.Until(targetTime)))
		t.log.DEBUG.Printf("reduced duration from %s to %s until got new price info after %s", duration_old.Round(time.Minute), requiredDuration.Round(time.Minute), last.Round(time.Minute))
	}

	// reduce planning horizon to available rates
	// if targetTime.After(last) {
	// 	old := requiredDuration
	// 	requiredDuration = time.Until(last)
	// 	t.log.DEBUG.Printf("target time beyond available slots- reduced plan horizon from %v to %v", old.Round(time.Minute), requiredDuration.Round(time.Minute))
	// }

	t.log.DEBUG.Printf("planning %s until %v", requiredDuration.Round(time.Minute), targetTime.Round(time.Minute))

	var active bool
	var plannedSlots, currentSlot int
	var plannedDuration time.Duration

	for _, slot := range rates {
		// slot not relevant
		if slot.Start.After(targetTime) || slot.Start.Equal(targetTime) || slot.End.Before(t.clock.Now()) {
			continue
		}

		// // current slot
		// if (slot.Start.Before(t.clock.Now()) || slot.Start.Equal(targetTime))&& slot.End.After(t.clock.Now()) {
		// 	slot.Start = t.clock.Now().Add(-1)
		// }

		// // slot ends after target time
		// if slot.End.After(targetTime) {
		// 	slot.End = targetTime.Add(1)
		// }

		plannedSlots++
		plannedDuration += slot.End.Sub(slot.Start)

		// slot covers current timestamp
		if (slot.Start.Before(t.clock.Now()) || slot.Start.Equal(t.clock.Now())) && slot.End.After(t.clock.Now()) {
			active = true
			currentSlot = plannedSlots
		}

		t.log.TRACE.Printf("  slot from: %v to %v cost %.2f, duration running total %s, active: %t",
			slot.Start.Round(time.Second), slot.End.Round(time.Second),
			slot.Price, plannedDuration.Round(time.Second), active)

		// we found all necessary cheap slots
		if plannedDuration >= requiredDuration {
			break
		}
	}

	// delay start of most expensive slot as long as possible
	// TODO @schenlap do we have a test for plannedSlots <=1 vs > 1?
	if currentSlot == plannedSlots && plannedSlots > 1 && plannedDuration > requiredDuration+hysteresisDuration {
		t.log.DEBUG.Printf("delaying expensive slot for %s", (plannedDuration - requiredDuration).Round(time.Minute))
		active = false
	}

	return active, nil
}

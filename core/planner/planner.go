package planner

import (
	"fmt"
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

// Active determines if current slot should be used for charging for a total required duration until target time
func (t *Planner) Active(requiredDuration time.Duration, targetTime time.Time) (bool, error) {
	if t == nil {
		return false, nil
	}

	if targetTime.Before(t.clock.Now()) || requiredDuration <= 0 {
		return false, nil
	}

	// target charging without tariff
	if t.tariff == nil {
		return t.clock.Now().After(targetTime.Add(-requiredDuration)), nil
	}

	rates, err := t.tariff.Rates()
	if err != nil {
		return false, err
	}

	// can't plan if we don't have rates
	if len(rates) == 0 {
		return false, fmt.Errorf("rates: %w", api.ErrNotAvailable)
	}

	// rates are by default sorted by date, oldest to newest
	last := rates[len(rates)-1].End

	// sort rates by price and time
	sort.Sort(rates)

	requiredDuration = time.Duration(float64(requiredDuration) / soc.ChargeEfficiency)

	// Save same duration until next price info update
	if targetTime.After(last) {
		duration_old := requiredDuration
		requiredDuration = time.Duration(float64(requiredDuration) * float64(time.Until(last)) / float64(time.Until(targetTime)))
		t.log.DEBUG.Printf("reduced duration from %s to %s until got new price info after %s\n", duration_old.Round(time.Minute), requiredDuration.Round(time.Minute), last.Round(time.Minute))
	}

	t.log.DEBUG.Printf("charge duration: %s, end: %v, find best prices:\n", requiredDuration.Round(time.Minute), targetTime.Round(time.Minute))

	var active bool
	var plannedSlots, currentSlot int
	var plannedDuration time.Duration

	for _, slot := range rates {
		// slot not relevant
		if slot.Start.After(targetTime) || slot.End.Before(t.clock.Now()) {
			continue
		}

		// current slot
		if slot.Start.Before(t.clock.Now()) && slot.End.After(t.clock.Now()) {
			slot.Start = t.clock.Now().Add(-1)
		}

		// slot ends after target time
		if slot.End.After(targetTime) {
			slot.End = targetTime.Add(1)
		}

		plannedSlots++
		plannedDuration += slot.End.Sub(slot.Start)

		t.log.TRACE.Printf("  slot from: %v to %v price %f, time sum %s",
			slot.Start.Round(time.Second), slot.End.Round(time.Second),
			slot.Price, plannedDuration.Round(time.Second))

		// plan covers current slot
		if slot.Start.Before(t.clock.Now().Add(1)) && slot.End.After(t.clock.Now()) {
			active = true
			currentSlot = plannedSlots
			t.log.TRACE.Printf(" (now, slot number %v)", currentSlot)
		}

		// we found all necessary cheap slots to charge to targetSoC
		if plannedDuration >= requiredDuration {
			break
		}
	}

	// delay start of most expensive slot as long as possible
	if currentSlot == plannedSlots && plannedSlots > 1 && plannedDuration > requiredDuration+hysteresisDuration {
		t.log.DEBUG.Printf("cheap times lot, delayed for %s\n", (plannedDuration - requiredDuration).Round(time.Minute))
		active = false
	}

	return active, nil
}

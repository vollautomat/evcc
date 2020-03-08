package wrapper

import (
	"fmt"
	"sync"
	"time"

	"github.com/andig/evcc/api"
	"github.com/benbjohnson/clock"
)

// ChargeRater is responsible for providing charged energy amount
// by implementing api.ChargeRater. It uses the charge meter's TotalEnergy or
// keeps track of consumed energy by regularly updating consumed power.
type ChargeRater struct {
	sync.Mutex
	clck          clock.Clock
	name          string
	meter         api.Meter
	charging      bool
	start         time.Time
	startEnergy   float64
	chargedEnergy float64
}

// NewChargeRater creates charge rater and initializes realtime clock
func NewChargeRater(name string, meter api.Meter) *ChargeRater {
	return &ChargeRater{
		clck:  clock.New(),
		name:  name,
		meter: meter,
	}
}

// StartCharge records meter start energy. If meter does not supply TotalEnergy,
// start time is recorded and  charged energy set to zero.
func (cr *ChargeRater) StartCharge() {
	cr.Lock()
	defer cr.Unlock()

	// time is needed if MeterEnergy is not supported
	cr.charging = true
	cr.start = cr.clck.Now()

	// get end energy amount
	if m, ok := cr.meter.(api.MeterEnergy); ok {
		if f, err := m.TotalEnergy(); err == nil {
			cr.startEnergy = f
			log.TRACE.Printf("%s charge start energy: %.0fkWh", cr.name, f)
		} else {
			log.ERROR.Printf("%s charge meter error %v", cr.name, err)
		}
	} else {
		cr.chargedEnergy = 0
	}
}

// StopCharge records meter stop energy. If meter does not supply TotalEnergy,
// stop time is recorded and accumulating energy though SetChargePower stopped.
func (cr *ChargeRater) StopCharge() {
	cr.Lock()
	defer cr.Unlock()

	cr.charging = false

	// get end energy amount
	if m, ok := cr.meter.(api.MeterEnergy); ok {
		if f, err := m.TotalEnergy(); err == nil {
			cr.chargedEnergy = f - cr.startEnergy
			log.TRACE.Printf("%s final charge energy: %.0fkWh", cr.name, cr.chargedEnergy)
		} else {
			log.ERROR.Printf("%s charge meter error %v", cr.name, err)
		}
	}
}

// SetChargePower increments consumed energy by amount in kWh since last update
func (cr *ChargeRater) SetChargePower(power float64) {
	cr.Lock()
	defer cr.Unlock()

	if !cr.charging {
		return
	}

	// update energy amount if not provided by meter
	if _, ok := cr.meter.(api.MeterEnergy); !ok {
		log.TRACE.Printf("%s set charge power: %.0fW", cr.name, power)
		// convert power to energy in kWh
		cr.chargedEnergy += power / 1e3 * float64(cr.clck.Since(cr.start)) / float64(time.Hour)
		// move timestamp
		cr.start = cr.clck.Now()
	}
}

// ChargedEnergy implements the ChargeRater interface.
// It returns energy consumption since charge start in kWh.
func (cr *ChargeRater) ChargedEnergy() (float64, error) {
	cr.Lock()
	defer cr.Unlock()

	// return previously charged energy
	if !cr.charging {
		return cr.chargedEnergy, nil
	}

	// get current energy amount
	if m, ok := cr.meter.(api.MeterEnergy); ok {
		f, err := m.TotalEnergy()

		if err == nil {
			log.TRACE.Printf("%s charge energy: %.0fkWh", cr.name, cr.chargedEnergy)
			return f - cr.startEnergy, nil
		}

		return 0, fmt.Errorf("%s charge meter error %v", cr.name, err)
	}

	// return charged energy sofar if meter is not used
	return cr.chargedEnergy, nil
}

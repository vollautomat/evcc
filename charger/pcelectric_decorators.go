package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decoratePCE(base *PCElectric, meter func() (float64, error), meterEnergy func() (float64, error), PhaseCurrents func() (float64, float64, float64, error)) api.Charger {
	switch {
	case meter == nil && PhaseCurrents == nil && meterEnergy == nil:
		return base

	case meter != nil && PhaseCurrents == nil && meterEnergy == nil:
		return &struct {
			*PCElectric
			api.Meter
		}{
			PCElectric: base,
			Meter: &decoratePCEMeterImpl{
				meter: meter,
			},
		}

	case meter == nil && PhaseCurrents == nil && meterEnergy != nil:
		return &struct {
			*PCElectric
			api.MeterEnergy
		}{
			PCElectric: base,
			MeterEnergy: &decoratePCEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter != nil && PhaseCurrents == nil && meterEnergy != nil:
		return &struct {
			*PCElectric
			api.Meter
			api.MeterEnergy
		}{
			PCElectric: base,
			Meter: &decoratePCEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decoratePCEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter == nil && PhaseCurrents != nil && meterEnergy == nil:
		return &struct {
			*PCElectric
			api.PhaseCurrents
		}{
			PCElectric: base,
			PhaseCurrents: &decoratePCEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
		}

	case meter != nil && PhaseCurrents != nil && meterEnergy == nil:
		return &struct {
			*PCElectric
			api.Meter
			api.PhaseCurrents
		}{
			PCElectric: base,
			Meter: &decoratePCEMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decoratePCEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
		}

	case meter == nil && PhaseCurrents != nil && meterEnergy != nil:
		return &struct {
			*PCElectric
			api.PhaseCurrents
			api.MeterEnergy
		}{
			PCElectric: base,
			PhaseCurrents: &decoratePCEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decoratePCEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case meter != nil && PhaseCurrents != nil && meterEnergy != nil:
		return &struct {
			*PCElectric
			api.Meter
			api.PhaseCurrents
			api.MeterEnergy
		}{
			PCElectric: base,
			Meter: &decoratePCEMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decoratePCEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decoratePCEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}
	}

	return nil
}

type decoratePCEMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decoratePCEMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decoratePCEPhaseCurrentsImpl struct {
	PhaseCurrents func() (float64, float64, float64, error)
}

func (impl *decoratePCEPhaseCurrentsImpl) Currents() (float64, float64, float64, error) {
	return impl.PhaseCurrents()
}

type decoratePCEMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decoratePCEMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}

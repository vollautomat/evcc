package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateKeba(base *Keba, meter func() (float64, error), meterEnergy func() (float64, error), phaseCurrents func() (float64, float64, float64, error), identifier func() (string, error), phaseController func(int) error) api.Charger {
	switch {
	case identifier == nil && meter == nil && meterEnergy == nil && phaseController == nil && phaseCurrents == nil:
		return base

	case identifier == nil && meter != nil && meterEnergy == nil && phaseController == nil && phaseCurrents == nil:
		return &struct {
			*Keba
			api.Meter
		}{
			Keba: base,
			Meter: &decorateKebaMeterImpl{
				meter: meter,
			},
		}

	case identifier == nil && meter == nil && meterEnergy != nil && phaseController == nil && phaseCurrents == nil:
		return &struct {
			*Keba
			api.MeterEnergy
		}{
			Keba: base,
			MeterEnergy: &decorateKebaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case identifier == nil && meter != nil && meterEnergy != nil && phaseController == nil && phaseCurrents == nil:
		return &struct {
			*Keba
			api.Meter
			api.MeterEnergy
		}{
			Keba: base,
			Meter: &decorateKebaMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateKebaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case identifier == nil && meter == nil && meterEnergy == nil && phaseController == nil && phaseCurrents != nil:
		return &struct {
			*Keba
			api.PhaseCurrents
		}{
			Keba: base,
			PhaseCurrents: &decorateKebaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case identifier == nil && meter != nil && meterEnergy == nil && phaseController == nil && phaseCurrents != nil:
		return &struct {
			*Keba
			api.Meter
			api.PhaseCurrents
		}{
			Keba: base,
			Meter: &decorateKebaMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateKebaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case identifier == nil && meter == nil && meterEnergy != nil && phaseController == nil && phaseCurrents != nil:
		return &struct {
			*Keba
			api.MeterEnergy
			api.PhaseCurrents
		}{
			Keba: base,
			MeterEnergy: &decorateKebaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateKebaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case identifier == nil && meter != nil && meterEnergy != nil && phaseController == nil && phaseCurrents != nil:
		return &struct {
			*Keba
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			Keba: base,
			Meter: &decorateKebaMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateKebaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateKebaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case identifier != nil && meter == nil && meterEnergy == nil && phaseController == nil && phaseCurrents == nil:
		return &struct {
			*Keba
			api.Identifier
		}{
			Keba: base,
			Identifier: &decorateKebaIdentifierImpl{
				identifier: identifier,
			},
		}

	case identifier != nil && meter != nil && meterEnergy == nil && phaseController == nil && phaseCurrents == nil:
		return &struct {
			*Keba
			api.Identifier
			api.Meter
		}{
			Keba: base,
			Identifier: &decorateKebaIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateKebaMeterImpl{
				meter: meter,
			},
		}

	case identifier != nil && meter == nil && meterEnergy != nil && phaseController == nil && phaseCurrents == nil:
		return &struct {
			*Keba
			api.Identifier
			api.MeterEnergy
		}{
			Keba: base,
			Identifier: &decorateKebaIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateKebaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case identifier != nil && meter != nil && meterEnergy != nil && phaseController == nil && phaseCurrents == nil:
		return &struct {
			*Keba
			api.Identifier
			api.Meter
			api.MeterEnergy
		}{
			Keba: base,
			Identifier: &decorateKebaIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateKebaMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateKebaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case identifier != nil && meter == nil && meterEnergy == nil && phaseController == nil && phaseCurrents != nil:
		return &struct {
			*Keba
			api.Identifier
			api.PhaseCurrents
		}{
			Keba: base,
			Identifier: &decorateKebaIdentifierImpl{
				identifier: identifier,
			},
			PhaseCurrents: &decorateKebaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case identifier != nil && meter != nil && meterEnergy == nil && phaseController == nil && phaseCurrents != nil:
		return &struct {
			*Keba
			api.Identifier
			api.Meter
			api.PhaseCurrents
		}{
			Keba: base,
			Identifier: &decorateKebaIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateKebaMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateKebaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case identifier != nil && meter == nil && meterEnergy != nil && phaseController == nil && phaseCurrents != nil:
		return &struct {
			*Keba
			api.Identifier
			api.MeterEnergy
			api.PhaseCurrents
		}{
			Keba: base,
			Identifier: &decorateKebaIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateKebaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateKebaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case identifier != nil && meter != nil && meterEnergy != nil && phaseController == nil && phaseCurrents != nil:
		return &struct {
			*Keba
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseCurrents
		}{
			Keba: base,
			Identifier: &decorateKebaIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateKebaMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateKebaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseCurrents: &decorateKebaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case identifier == nil && meter == nil && meterEnergy == nil && phaseController != nil && phaseCurrents == nil:
		return &struct {
			*Keba
			api.PhaseController
		}{
			Keba: base,
			PhaseController: &decorateKebaPhaseControllerImpl{
				phaseController: phaseController,
			},
		}

	case identifier == nil && meter != nil && meterEnergy == nil && phaseController != nil && phaseCurrents == nil:
		return &struct {
			*Keba
			api.Meter
			api.PhaseController
		}{
			Keba: base,
			Meter: &decorateKebaMeterImpl{
				meter: meter,
			},
			PhaseController: &decorateKebaPhaseControllerImpl{
				phaseController: phaseController,
			},
		}

	case identifier == nil && meter == nil && meterEnergy != nil && phaseController != nil && phaseCurrents == nil:
		return &struct {
			*Keba
			api.MeterEnergy
			api.PhaseController
		}{
			Keba: base,
			MeterEnergy: &decorateKebaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseController: &decorateKebaPhaseControllerImpl{
				phaseController: phaseController,
			},
		}

	case identifier == nil && meter != nil && meterEnergy != nil && phaseController != nil && phaseCurrents == nil:
		return &struct {
			*Keba
			api.Meter
			api.MeterEnergy
			api.PhaseController
		}{
			Keba: base,
			Meter: &decorateKebaMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateKebaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseController: &decorateKebaPhaseControllerImpl{
				phaseController: phaseController,
			},
		}

	case identifier == nil && meter == nil && meterEnergy == nil && phaseController != nil && phaseCurrents != nil:
		return &struct {
			*Keba
			api.PhaseController
			api.PhaseCurrents
		}{
			Keba: base,
			PhaseController: &decorateKebaPhaseControllerImpl{
				phaseController: phaseController,
			},
			PhaseCurrents: &decorateKebaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case identifier == nil && meter != nil && meterEnergy == nil && phaseController != nil && phaseCurrents != nil:
		return &struct {
			*Keba
			api.Meter
			api.PhaseController
			api.PhaseCurrents
		}{
			Keba: base,
			Meter: &decorateKebaMeterImpl{
				meter: meter,
			},
			PhaseController: &decorateKebaPhaseControllerImpl{
				phaseController: phaseController,
			},
			PhaseCurrents: &decorateKebaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case identifier == nil && meter == nil && meterEnergy != nil && phaseController != nil && phaseCurrents != nil:
		return &struct {
			*Keba
			api.MeterEnergy
			api.PhaseController
			api.PhaseCurrents
		}{
			Keba: base,
			MeterEnergy: &decorateKebaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseController: &decorateKebaPhaseControllerImpl{
				phaseController: phaseController,
			},
			PhaseCurrents: &decorateKebaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case identifier == nil && meter != nil && meterEnergy != nil && phaseController != nil && phaseCurrents != nil:
		return &struct {
			*Keba
			api.Meter
			api.MeterEnergy
			api.PhaseController
			api.PhaseCurrents
		}{
			Keba: base,
			Meter: &decorateKebaMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateKebaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseController: &decorateKebaPhaseControllerImpl{
				phaseController: phaseController,
			},
			PhaseCurrents: &decorateKebaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case identifier != nil && meter == nil && meterEnergy == nil && phaseController != nil && phaseCurrents == nil:
		return &struct {
			*Keba
			api.Identifier
			api.PhaseController
		}{
			Keba: base,
			Identifier: &decorateKebaIdentifierImpl{
				identifier: identifier,
			},
			PhaseController: &decorateKebaPhaseControllerImpl{
				phaseController: phaseController,
			},
		}

	case identifier != nil && meter != nil && meterEnergy == nil && phaseController != nil && phaseCurrents == nil:
		return &struct {
			*Keba
			api.Identifier
			api.Meter
			api.PhaseController
		}{
			Keba: base,
			Identifier: &decorateKebaIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateKebaMeterImpl{
				meter: meter,
			},
			PhaseController: &decorateKebaPhaseControllerImpl{
				phaseController: phaseController,
			},
		}

	case identifier != nil && meter == nil && meterEnergy != nil && phaseController != nil && phaseCurrents == nil:
		return &struct {
			*Keba
			api.Identifier
			api.MeterEnergy
			api.PhaseController
		}{
			Keba: base,
			Identifier: &decorateKebaIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateKebaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseController: &decorateKebaPhaseControllerImpl{
				phaseController: phaseController,
			},
		}

	case identifier != nil && meter != nil && meterEnergy != nil && phaseController != nil && phaseCurrents == nil:
		return &struct {
			*Keba
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseController
		}{
			Keba: base,
			Identifier: &decorateKebaIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateKebaMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateKebaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseController: &decorateKebaPhaseControllerImpl{
				phaseController: phaseController,
			},
		}

	case identifier != nil && meter == nil && meterEnergy == nil && phaseController != nil && phaseCurrents != nil:
		return &struct {
			*Keba
			api.Identifier
			api.PhaseController
			api.PhaseCurrents
		}{
			Keba: base,
			Identifier: &decorateKebaIdentifierImpl{
				identifier: identifier,
			},
			PhaseController: &decorateKebaPhaseControllerImpl{
				phaseController: phaseController,
			},
			PhaseCurrents: &decorateKebaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case identifier != nil && meter != nil && meterEnergy == nil && phaseController != nil && phaseCurrents != nil:
		return &struct {
			*Keba
			api.Identifier
			api.Meter
			api.PhaseController
			api.PhaseCurrents
		}{
			Keba: base,
			Identifier: &decorateKebaIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateKebaMeterImpl{
				meter: meter,
			},
			PhaseController: &decorateKebaPhaseControllerImpl{
				phaseController: phaseController,
			},
			PhaseCurrents: &decorateKebaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case identifier != nil && meter == nil && meterEnergy != nil && phaseController != nil && phaseCurrents != nil:
		return &struct {
			*Keba
			api.Identifier
			api.MeterEnergy
			api.PhaseController
			api.PhaseCurrents
		}{
			Keba: base,
			Identifier: &decorateKebaIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateKebaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseController: &decorateKebaPhaseControllerImpl{
				phaseController: phaseController,
			},
			PhaseCurrents: &decorateKebaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case identifier != nil && meter != nil && meterEnergy != nil && phaseController != nil && phaseCurrents != nil:
		return &struct {
			*Keba
			api.Identifier
			api.Meter
			api.MeterEnergy
			api.PhaseController
			api.PhaseCurrents
		}{
			Keba: base,
			Identifier: &decorateKebaIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateKebaMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateKebaMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseController: &decorateKebaPhaseControllerImpl{
				phaseController: phaseController,
			},
			PhaseCurrents: &decorateKebaPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}
	}

	return nil
}

type decorateKebaIdentifierImpl struct {
	identifier func() (string, error)
}

func (impl *decorateKebaIdentifierImpl) Identify() (string, error) {
	return impl.identifier()
}

type decorateKebaMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decorateKebaMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decorateKebaMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decorateKebaMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}

type decorateKebaPhaseControllerImpl struct {
	phaseController func(int) error
}

func (impl *decorateKebaPhaseControllerImpl) Phases1p3p(p0 int) error {
	return impl.phaseController(p0)
}

type decorateKebaPhaseCurrentsImpl struct {
	phaseCurrents func() (float64, float64, float64, error)
}

func (impl *decorateKebaPhaseCurrentsImpl) Currents() (float64, float64, float64, error) {
	return impl.phaseCurrents()
}

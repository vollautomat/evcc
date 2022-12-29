package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateEVSE(base *EVSEWifi, meter func() (float64, error), meterEnergy func() (float64, error), PhaseCurrents func() (float64, float64, float64, error), chargerEx func(float64) error, identifier func() (string, error)) api.Charger {
	switch {
	case chargerEx == nil && identifier == nil && meter == nil && PhaseCurrents == nil && meterEnergy == nil:
		return base

	case chargerEx == nil && identifier == nil && meter != nil && PhaseCurrents == nil && meterEnergy == nil:
		return &struct {
			*EVSEWifi
			api.Meter
		}{
			EVSEWifi: base,
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
		}

	case chargerEx == nil && identifier == nil && meter == nil && PhaseCurrents == nil && meterEnergy != nil:
		return &struct {
			*EVSEWifi
			api.MeterEnergy
		}{
			EVSEWifi: base,
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx == nil && identifier == nil && meter != nil && PhaseCurrents == nil && meterEnergy != nil:
		return &struct {
			*EVSEWifi
			api.Meter
			api.MeterEnergy
		}{
			EVSEWifi: base,
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx == nil && identifier == nil && meter == nil && PhaseCurrents != nil && meterEnergy == nil:
		return &struct {
			*EVSEWifi
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
		}

	case chargerEx == nil && identifier == nil && meter != nil && PhaseCurrents != nil && meterEnergy == nil:
		return &struct {
			*EVSEWifi
			api.Meter
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
		}

	case chargerEx == nil && identifier == nil && meter == nil && PhaseCurrents != nil && meterEnergy != nil:
		return &struct {
			*EVSEWifi
			api.PhaseCurrents
			api.MeterEnergy
		}{
			EVSEWifi: base,
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx == nil && identifier == nil && meter != nil && PhaseCurrents != nil && meterEnergy != nil:
		return &struct {
			*EVSEWifi
			api.Meter
			api.PhaseCurrents
			api.MeterEnergy
		}{
			EVSEWifi: base,
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx != nil && identifier == nil && meter == nil && PhaseCurrents == nil && meterEnergy == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && PhaseCurrents == nil && meterEnergy == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Meter
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
		}

	case chargerEx != nil && identifier == nil && meter == nil && PhaseCurrents == nil && meterEnergy != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.MeterEnergy
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && PhaseCurrents == nil && meterEnergy != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Meter
			api.MeterEnergy
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx != nil && identifier == nil && meter == nil && PhaseCurrents != nil && meterEnergy == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && PhaseCurrents != nil && meterEnergy == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Meter
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
		}

	case chargerEx != nil && identifier == nil && meter == nil && PhaseCurrents != nil && meterEnergy != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.PhaseCurrents
			api.MeterEnergy
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx != nil && identifier == nil && meter != nil && PhaseCurrents != nil && meterEnergy != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Meter
			api.PhaseCurrents
			api.MeterEnergy
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx == nil && identifier != nil && meter == nil && PhaseCurrents == nil && meterEnergy == nil:
		return &struct {
			*EVSEWifi
			api.Identifier
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && PhaseCurrents == nil && meterEnergy == nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.Meter
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
		}

	case chargerEx == nil && identifier != nil && meter == nil && PhaseCurrents == nil && meterEnergy != nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.MeterEnergy
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && PhaseCurrents == nil && meterEnergy != nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.Meter
			api.MeterEnergy
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx == nil && identifier != nil && meter == nil && PhaseCurrents != nil && meterEnergy == nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && PhaseCurrents != nil && meterEnergy == nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.Meter
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
		}

	case chargerEx == nil && identifier != nil && meter == nil && PhaseCurrents != nil && meterEnergy != nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.PhaseCurrents
			api.MeterEnergy
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx == nil && identifier != nil && meter != nil && PhaseCurrents != nil && meterEnergy != nil:
		return &struct {
			*EVSEWifi
			api.Identifier
			api.Meter
			api.PhaseCurrents
			api.MeterEnergy
		}{
			EVSEWifi: base,
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx != nil && identifier != nil && meter == nil && PhaseCurrents == nil && meterEnergy == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && PhaseCurrents == nil && meterEnergy == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.Meter
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
		}

	case chargerEx != nil && identifier != nil && meter == nil && PhaseCurrents == nil && meterEnergy != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.MeterEnergy
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && PhaseCurrents == nil && meterEnergy != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.Meter
			api.MeterEnergy
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx != nil && identifier != nil && meter == nil && PhaseCurrents != nil && meterEnergy == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && PhaseCurrents != nil && meterEnergy == nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.Meter
			api.PhaseCurrents
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
		}

	case chargerEx != nil && identifier != nil && meter == nil && PhaseCurrents != nil && meterEnergy != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.PhaseCurrents
			api.MeterEnergy
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case chargerEx != nil && identifier != nil && meter != nil && PhaseCurrents != nil && meterEnergy != nil:
		return &struct {
			*EVSEWifi
			api.ChargerEx
			api.Identifier
			api.Meter
			api.PhaseCurrents
			api.MeterEnergy
		}{
			EVSEWifi: base,
			ChargerEx: &decorateEVSEChargerExImpl{
				chargerEx: chargerEx,
			},
			Identifier: &decorateEVSEIdentifierImpl{
				identifier: identifier,
			},
			Meter: &decorateEVSEMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decorateEVSEPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateEVSEMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}
	}

	return nil
}

type decorateEVSEChargerExImpl struct {
	chargerEx func(current float64) error
}

func (impl *decorateEVSEChargerExImpl) MaxCurrentMillis(current float64) error {
	return impl.chargerEx(current)
}

type decorateEVSEIdentifierImpl struct {
	identifier func() (string, error)
}

func (impl *decorateEVSEIdentifierImpl) Identify() (string, error) {
	return impl.identifier()
}

type decorateEVSEMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decorateEVSEMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decorateEVSEPhaseCurrentsImpl struct {
	PhaseCurrents func() (float64, float64, float64, error)
}

func (impl *decorateEVSEPhaseCurrentsImpl) Currents() (float64, float64, float64, error) {
	return impl.PhaseCurrents()
}

type decorateEVSEMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decorateEVSEMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}

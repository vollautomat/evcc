package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decoratePhoenixEMEth(base *PhoenixEMEth, meter func() (float64, error), energyImport func() (float64, error), phaseCurrents func() (float64, float64, float64, error), phaseVoltages func() (float64, float64, float64, error)) api.Charger {
	switch {
	case energyImport == nil && meter == nil:
		return base

	case energyImport == nil && meter != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEMEth
			api.Meter
		}{
			PhoenixEMEth: base,
			Meter: &decoratePhoenixEMEthMeterImpl{
				meter: meter,
			},
		}

	case energyImport != nil && meter == nil:
		return &struct {
			*PhoenixEMEth
			api.EnergyImport
		}{
			PhoenixEMEth: base,
			EnergyImport: &decoratePhoenixEMEthEnergyImportImpl{
				energyImport: energyImport,
			},
		}

	case energyImport != nil && meter != nil && phaseCurrents == nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEMEth
			api.EnergyImport
			api.Meter
		}{
			PhoenixEMEth: base,
			EnergyImport: &decoratePhoenixEMEthEnergyImportImpl{
				energyImport: energyImport,
			},
			Meter: &decoratePhoenixEMEthMeterImpl{
				meter: meter,
			},
		}

	case energyImport == nil && meter != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEMEth
			api.Meter
			api.PhaseCurrents
		}{
			PhoenixEMEth: base,
			Meter: &decoratePhoenixEMEthMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decoratePhoenixEMEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case energyImport != nil && meter != nil && phaseCurrents != nil && phaseVoltages == nil:
		return &struct {
			*PhoenixEMEth
			api.EnergyImport
			api.Meter
			api.PhaseCurrents
		}{
			PhoenixEMEth: base,
			EnergyImport: &decoratePhoenixEMEthEnergyImportImpl{
				energyImport: energyImport,
			},
			Meter: &decoratePhoenixEMEthMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decoratePhoenixEMEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
		}

	case energyImport == nil && meter != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEMEth
			api.Meter
			api.PhaseVoltages
		}{
			PhoenixEMEth: base,
			Meter: &decoratePhoenixEMEthMeterImpl{
				meter: meter,
			},
			PhaseVoltages: &decoratePhoenixEMEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case energyImport != nil && meter != nil && phaseCurrents == nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEMEth
			api.EnergyImport
			api.Meter
			api.PhaseVoltages
		}{
			PhoenixEMEth: base,
			EnergyImport: &decoratePhoenixEMEthEnergyImportImpl{
				energyImport: energyImport,
			},
			Meter: &decoratePhoenixEMEthMeterImpl{
				meter: meter,
			},
			PhaseVoltages: &decoratePhoenixEMEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case energyImport == nil && meter != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEMEth
			api.Meter
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			PhoenixEMEth: base,
			Meter: &decoratePhoenixEMEthMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decoratePhoenixEMEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decoratePhoenixEMEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}

	case energyImport != nil && meter != nil && phaseCurrents != nil && phaseVoltages != nil:
		return &struct {
			*PhoenixEMEth
			api.EnergyImport
			api.Meter
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			PhoenixEMEth: base,
			EnergyImport: &decoratePhoenixEMEthEnergyImportImpl{
				energyImport: energyImport,
			},
			Meter: &decoratePhoenixEMEthMeterImpl{
				meter: meter,
			},
			PhaseCurrents: &decoratePhoenixEMEthPhaseCurrentsImpl{
				phaseCurrents: phaseCurrents,
			},
			PhaseVoltages: &decoratePhoenixEMEthPhaseVoltagesImpl{
				phaseVoltages: phaseVoltages,
			},
		}
	}

	return nil
}

type decoratePhoenixEMEthEnergyImportImpl struct {
	energyImport func() (float64, error)
}

func (impl *decoratePhoenixEMEthEnergyImportImpl) EnergyImport() (float64, error) {
	return impl.energyImport()
}

type decoratePhoenixEMEthMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decoratePhoenixEMEthMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decoratePhoenixEMEthPhaseCurrentsImpl struct {
	phaseCurrents func() (float64, float64, float64, error)
}

func (impl *decoratePhoenixEMEthPhaseCurrentsImpl) Currents() (float64, float64, float64, error) {
	return impl.phaseCurrents()
}

type decoratePhoenixEMEthPhaseVoltagesImpl struct {
	phaseVoltages func() (float64, float64, float64, error)
}

func (impl *decoratePhoenixEMEthPhaseVoltagesImpl) Voltages() (float64, float64, float64, error) {
	return impl.phaseVoltages()
}

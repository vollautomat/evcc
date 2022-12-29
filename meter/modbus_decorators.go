package meter

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateModbus(base api.Meter, meterEnergy func() (float64, error), PhaseCurrents func() (float64, float64, float64, error), PhaseVoltages func() (float64, float64, float64, error), PhasePowers func() (float64, float64, float64, error), battery func() (float64, error), batteryCapacity func() float64) api.Meter {
	switch {
	case battery == nil && batteryCapacity == nil && PhaseCurrents == nil && meterEnergy == nil && PhasePowers == nil && PhaseVoltages == nil:
		return base

	case battery == nil && batteryCapacity == nil && PhaseCurrents == nil && meterEnergy != nil && PhasePowers == nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.MeterEnergy
		}{
			Meter: base,
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery == nil && batteryCapacity == nil && PhaseCurrents != nil && meterEnergy == nil && PhasePowers == nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.PhaseCurrents
		}{
			Meter: base,
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
		}

	case battery == nil && batteryCapacity == nil && PhaseCurrents != nil && meterEnergy != nil && PhasePowers == nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.PhaseCurrents
			api.MeterEnergy
		}{
			Meter: base,
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery == nil && batteryCapacity == nil && PhaseCurrents == nil && meterEnergy == nil && PhasePowers == nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.PhaseVoltages
		}{
			Meter: base,
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery == nil && batteryCapacity == nil && PhaseCurrents == nil && meterEnergy != nil && PhasePowers == nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.MeterEnergy
			api.PhaseVoltages
		}{
			Meter: base,
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery == nil && batteryCapacity == nil && PhaseCurrents != nil && meterEnergy == nil && PhasePowers == nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery == nil && batteryCapacity == nil && PhaseCurrents != nil && meterEnergy != nil && PhasePowers == nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.PhaseCurrents
			api.MeterEnergy
			api.PhaseVoltages
		}{
			Meter: base,
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery == nil && batteryCapacity == nil && PhaseCurrents == nil && meterEnergy == nil && PhasePowers != nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.PhasePowers
		}{
			Meter: base,
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
		}

	case battery == nil && batteryCapacity == nil && PhaseCurrents == nil && meterEnergy != nil && PhasePowers != nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.MeterEnergy
			api.PhasePowers
		}{
			Meter: base,
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
		}

	case battery == nil && batteryCapacity == nil && PhaseCurrents != nil && meterEnergy == nil && PhasePowers != nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
		}

	case battery == nil && batteryCapacity == nil && PhaseCurrents != nil && meterEnergy != nil && PhasePowers != nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.PhaseCurrents
			api.MeterEnergy
			api.PhasePowers
		}{
			Meter: base,
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
		}

	case battery == nil && batteryCapacity == nil && PhaseCurrents == nil && meterEnergy == nil && PhasePowers != nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery == nil && batteryCapacity == nil && PhaseCurrents == nil && meterEnergy != nil && PhasePowers != nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.MeterEnergy
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery == nil && batteryCapacity == nil && PhaseCurrents != nil && meterEnergy == nil && PhasePowers != nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery == nil && batteryCapacity == nil && PhaseCurrents != nil && meterEnergy != nil && PhasePowers != nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.PhaseCurrents
			api.MeterEnergy
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && PhaseCurrents == nil && meterEnergy == nil && PhasePowers == nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
		}

	case battery != nil && batteryCapacity == nil && PhaseCurrents == nil && meterEnergy != nil && PhasePowers == nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.MeterEnergy
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery != nil && batteryCapacity == nil && PhaseCurrents != nil && meterEnergy == nil && PhasePowers == nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhaseCurrents
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
		}

	case battery != nil && batteryCapacity == nil && PhaseCurrents != nil && meterEnergy != nil && PhasePowers == nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhaseCurrents
			api.MeterEnergy
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery != nil && batteryCapacity == nil && PhaseCurrents == nil && meterEnergy == nil && PhasePowers == nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && PhaseCurrents == nil && meterEnergy != nil && PhasePowers == nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.MeterEnergy
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && PhaseCurrents != nil && meterEnergy == nil && PhasePowers == nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && PhaseCurrents != nil && meterEnergy != nil && PhasePowers == nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhaseCurrents
			api.MeterEnergy
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && PhaseCurrents == nil && meterEnergy == nil && PhasePowers != nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
		}

	case battery != nil && batteryCapacity == nil && PhaseCurrents == nil && meterEnergy != nil && PhasePowers != nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.MeterEnergy
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
		}

	case battery != nil && batteryCapacity == nil && PhaseCurrents != nil && meterEnergy == nil && PhasePowers != nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
		}

	case battery != nil && batteryCapacity == nil && PhaseCurrents != nil && meterEnergy != nil && PhasePowers != nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhaseCurrents
			api.MeterEnergy
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
		}

	case battery != nil && batteryCapacity == nil && PhaseCurrents == nil && meterEnergy == nil && PhasePowers != nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && PhaseCurrents == nil && meterEnergy != nil && PhasePowers != nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.MeterEnergy
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && PhaseCurrents != nil && meterEnergy == nil && PhasePowers != nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery != nil && batteryCapacity == nil && PhaseCurrents != nil && meterEnergy != nil && PhasePowers != nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.PhaseCurrents
			api.MeterEnergy
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery == nil && batteryCapacity != nil && PhaseCurrents == nil && meterEnergy == nil && PhasePowers == nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
		}

	case battery == nil && batteryCapacity != nil && PhaseCurrents == nil && meterEnergy != nil && PhasePowers == nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.MeterEnergy
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery == nil && batteryCapacity != nil && PhaseCurrents != nil && meterEnergy == nil && PhasePowers == nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.PhaseCurrents
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
		}

	case battery == nil && batteryCapacity != nil && PhaseCurrents != nil && meterEnergy != nil && PhasePowers == nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.PhaseCurrents
			api.MeterEnergy
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery == nil && batteryCapacity != nil && PhaseCurrents == nil && meterEnergy == nil && PhasePowers == nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.PhaseVoltages
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery == nil && batteryCapacity != nil && PhaseCurrents == nil && meterEnergy != nil && PhasePowers == nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.MeterEnergy
			api.PhaseVoltages
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery == nil && batteryCapacity != nil && PhaseCurrents != nil && meterEnergy == nil && PhasePowers == nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery == nil && batteryCapacity != nil && PhaseCurrents != nil && meterEnergy != nil && PhasePowers == nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.PhaseCurrents
			api.MeterEnergy
			api.PhaseVoltages
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery == nil && batteryCapacity != nil && PhaseCurrents == nil && meterEnergy == nil && PhasePowers != nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.PhasePowers
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
		}

	case battery == nil && batteryCapacity != nil && PhaseCurrents == nil && meterEnergy != nil && PhasePowers != nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.MeterEnergy
			api.PhasePowers
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
		}

	case battery == nil && batteryCapacity != nil && PhaseCurrents != nil && meterEnergy == nil && PhasePowers != nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
		}

	case battery == nil && batteryCapacity != nil && PhaseCurrents != nil && meterEnergy != nil && PhasePowers != nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.PhaseCurrents
			api.MeterEnergy
			api.PhasePowers
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
		}

	case battery == nil && batteryCapacity != nil && PhaseCurrents == nil && meterEnergy == nil && PhasePowers != nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery == nil && batteryCapacity != nil && PhaseCurrents == nil && meterEnergy != nil && PhasePowers != nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.MeterEnergy
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery == nil && batteryCapacity != nil && PhaseCurrents != nil && meterEnergy == nil && PhasePowers != nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery == nil && batteryCapacity != nil && PhaseCurrents != nil && meterEnergy != nil && PhasePowers != nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.BatteryCapacity
			api.PhaseCurrents
			api.MeterEnergy
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && PhaseCurrents == nil && meterEnergy == nil && PhasePowers == nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
		}

	case battery != nil && batteryCapacity != nil && PhaseCurrents == nil && meterEnergy != nil && PhasePowers == nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery != nil && batteryCapacity != nil && PhaseCurrents != nil && meterEnergy == nil && PhasePowers == nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhaseCurrents
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
		}

	case battery != nil && batteryCapacity != nil && PhaseCurrents != nil && meterEnergy != nil && PhasePowers == nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhaseCurrents
			api.MeterEnergy
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
		}

	case battery != nil && batteryCapacity != nil && PhaseCurrents == nil && meterEnergy == nil && PhasePowers == nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && PhaseCurrents == nil && meterEnergy != nil && PhasePowers == nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && PhaseCurrents != nil && meterEnergy == nil && PhasePowers == nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhaseCurrents
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && PhaseCurrents != nil && meterEnergy != nil && PhasePowers == nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhaseCurrents
			api.MeterEnergy
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && PhaseCurrents == nil && meterEnergy == nil && PhasePowers != nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
		}

	case battery != nil && batteryCapacity != nil && PhaseCurrents == nil && meterEnergy != nil && PhasePowers != nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
		}

	case battery != nil && batteryCapacity != nil && PhaseCurrents != nil && meterEnergy == nil && PhasePowers != nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhaseCurrents
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
		}

	case battery != nil && batteryCapacity != nil && PhaseCurrents != nil && meterEnergy != nil && PhasePowers != nil && PhaseVoltages == nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhaseCurrents
			api.MeterEnergy
			api.PhasePowers
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
		}

	case battery != nil && batteryCapacity != nil && PhaseCurrents == nil && meterEnergy == nil && PhasePowers != nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && PhaseCurrents == nil && meterEnergy != nil && PhasePowers != nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.MeterEnergy
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && PhaseCurrents != nil && meterEnergy == nil && PhasePowers != nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhaseCurrents
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}

	case battery != nil && batteryCapacity != nil && PhaseCurrents != nil && meterEnergy != nil && PhasePowers != nil && PhaseVoltages != nil:
		return &struct {
			api.Meter
			api.Battery
			api.BatteryCapacity
			api.PhaseCurrents
			api.MeterEnergy
			api.PhasePowers
			api.PhaseVoltages
		}{
			Meter: base,
			Battery: &decorateModbusBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateModbusBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
			PhaseCurrents: &decorateModbusPhaseCurrentsImpl{
				PhaseCurrents: PhaseCurrents,
			},
			MeterEnergy: &decorateModbusMeterEnergyImpl{
				meterEnergy: meterEnergy,
			},
			PhasePowers: &decorateModbusPhasePowersImpl{
				PhasePowers: PhasePowers,
			},
			PhaseVoltages: &decorateModbusPhaseVoltagesImpl{
				PhaseVoltages: PhaseVoltages,
			},
		}
	}

	return nil
}

type decorateModbusBatteryImpl struct {
	battery func() (float64, error)
}

func (impl *decorateModbusBatteryImpl) Soc() (float64, error) {
	return impl.battery()
}

type decorateModbusBatteryCapacityImpl struct {
	batteryCapacity func() float64
}

func (impl *decorateModbusBatteryCapacityImpl) Capacity() float64 {
	return impl.batteryCapacity()
}

type decorateModbusPhaseCurrentsImpl struct {
	PhaseCurrents func() (float64, float64, float64, error)
}

func (impl *decorateModbusPhaseCurrentsImpl) Currents() (float64, float64, float64, error) {
	return impl.PhaseCurrents()
}

type decorateModbusMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decorateModbusMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}

type decorateModbusPhasePowersImpl struct {
	PhasePowers func() (float64, float64, float64, error)
}

func (impl *decorateModbusPhasePowersImpl) Powers() (float64, float64, float64, error) {
	return impl.PhasePowers()
}

type decorateModbusPhaseVoltagesImpl struct {
	PhaseVoltages func() (float64, float64, float64, error)
}

func (impl *decorateModbusPhaseVoltagesImpl) Voltages() (float64, float64, float64, error) {
	return impl.PhaseVoltages()
}

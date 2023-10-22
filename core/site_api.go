package core

import (
	"errors"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/core/site"
	"github.com/evcc-io/evcc/core/vehicle"
	"github.com/evcc-io/evcc/server/db/settings"
)

var _ site.API = (*Site)(nil)

var ErrBatteryNotConfigured = errors.New("battery not configured")

const (
	GridTariff    = "grid"
	FeedinTariff  = "feedin"
	PlannerTariff = "planner"
)

// GetPrioritySoc returns the PrioritySoc
func (site *Site) GetPrioritySoc() float64 {
	site.Lock()
	defer site.Unlock()
	return site.PrioritySoc
}

// SetPrioritySoc sets the PrioritySoc
func (site *Site) SetPrioritySoc(soc float64) error {
	site.Lock()
	defer site.Unlock()

	if len(site.batteryMeters) == 0 {
		return ErrBatteryNotConfigured
	}

	site.log.DEBUG.Println("set priority soc:", soc)

	if site.PrioritySoc != soc {
		site.PrioritySoc = soc
		settings.SetFloat("site.prioritySoc", site.PrioritySoc)
		site.publish("prioritySoc", site.PrioritySoc)

	}
	return nil
}

// GetBufferSoc returns the BufferSoc
func (site *Site) GetBufferSoc() float64 {
	site.Lock()
	defer site.Unlock()
	return site.BufferSoc
}

// SetBufferSoc sets the BufferSoc
func (site *Site) SetBufferSoc(soc float64) error {
	site.Lock()
	defer site.Unlock()

	if len(site.batteryMeters) == 0 {
		return ErrBatteryNotConfigured
	}

	site.log.DEBUG.Println("set buffer soc:", soc)

	if site.BufferSoc != soc {
		site.BufferSoc = soc
		settings.SetFloat("site.bufferSoc", site.BufferSoc)
		site.publish("bufferSoc", site.BufferSoc)
	}

	return nil
}

// GetBufferStartSoc returns the BufferStartSoc
func (site *Site) GetBufferStartSoc() float64 {
	site.Lock()
	defer site.Unlock()
	return site.BufferStartSoc
}

// SetBufferStartSoc sets the BufferStartSoc
func (site *Site) SetBufferStartSoc(soc float64) error {
	site.Lock()
	defer site.Unlock()

	if len(site.batteryMeters) == 0 {
		return ErrBatteryNotConfigured
	}

	site.log.DEBUG.Println("set buffer start soc:", soc)

	if site.BufferStartSoc != soc {
		site.BufferStartSoc = soc
		settings.SetFloat("site.bufferStartSoc", site.BufferStartSoc)
		site.publish("bufferStartSoc", site.BufferStartSoc)
	}

	return nil
}

// GetResidualPower returns the ResidualPower
func (site *Site) GetResidualPower() float64 {
	site.Lock()
	defer site.Unlock()
	return site.ResidualPower
}

// SetResidualPower sets the ResidualPower
func (site *Site) SetResidualPower(power float64) error {
	site.Lock()
	defer site.Unlock()

	site.log.DEBUG.Println("set residual power:", power)

	if site.ResidualPower != power {
		site.ResidualPower = power
		site.publish("residualPower", site.ResidualPower)
	}

	return nil
}

// GetSmartCostLimit returns the SmartCostLimit
func (site *Site) GetSmartCostLimit() float64 {
	site.Lock()
	defer site.Unlock()
	return site.SmartCostLimit
}

// SetSmartCostLimit sets the SmartCostLimit
func (site *Site) SetSmartCostLimit(val float64) error {
	site.Lock()
	defer site.Unlock()

	site.log.DEBUG.Println("set smart cost limit:", val)

	if site.SmartCostLimit != val {
		site.SmartCostLimit = val
		settings.SetFloat("site.smartCostLimit", site.SmartCostLimit)
		site.publish("smartCostLimit", site.SmartCostLimit)
	}

	return nil
}

// GetVehicles returns the list of vehicles
func (site *Site) GetVehicles() []api.Vehicle {
	site.Lock()
	defer site.Unlock()
	return site.coordinator.GetVehicles()
}

// VehicleSettings returns the list of vehicle setting adapters
func (site *Site) VehicleSettings() []vehicle.API {
	// TODO refactor
	res := make([]vehicle.API, 0, len(site.GetVehicles()))
	for _, dev := range vehicle.Handler.Devices() {
		res = append(res, vehicle.Settings(dev.Instance()))
	}
	return res
}

// GetTariff returns the respective tariff if configured or nil
func (site *Site) GetTariff(tariff string) api.Tariff {
	site.Lock()
	defer site.Unlock()

	switch tariff {
	case GridTariff:
		return site.tariffs.Grid

	case FeedinTariff:
		return site.tariffs.FeedIn

	case PlannerTariff:
		switch {
		case site.tariffs.Planner != nil:
			// prio 0: manually set planner tariff
			return site.tariffs.Planner

		case site.tariffs.Grid != nil && site.tariffs.Grid.Type() == api.TariffTypePriceForecast:
			// prio 1: dynamic grid tariff
			return site.tariffs.Grid

		case site.tariffs.Co2 != nil:
			// prio 2: co2 tariff
			return site.tariffs.Co2

		default:
			// prio 3: static grid tariff
			return site.tariffs.Grid
		}

	default:
		return nil
	}
}

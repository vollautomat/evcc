package meter

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateRCT(base *RCT, meterEnergy func() (float64, error), battery func() (float64, error), batteryCapacity func() float64) api.Meter {
	switch {
	case battery == nil && batteryCapacity == nil && meterEnergy == nil:
		return base

	case battery != nil && batteryCapacity == nil && meterEnergy == nil:
		return &struct {
			*RCT
			api.Battery
		}{
			RCT: base,
			Battery: &decorateRCTBatteryImpl{
				battery: battery,
			},
		}

	case battery != nil && batteryCapacity != nil && meterEnergy == nil:
		return &struct {
			*RCT
			api.Battery
			api.BatteryCapacity
		}{
			RCT: base,
			Battery: &decorateRCTBatteryImpl{
				battery: battery,
			},
			BatteryCapacity: &decorateRCTBatteryCapacityImpl{
				batteryCapacity: batteryCapacity,
			},
		}
	}

	return nil
}

type decorateRCTBatteryImpl struct {
	battery func() (float64, error)
}

func (impl *decorateRCTBatteryImpl) Soc() (float64, error) {
	return impl.battery()
}

type decorateRCTBatteryCapacityImpl struct {
	batteryCapacity func() float64
}

func (impl *decorateRCTBatteryCapacityImpl) Capacity() float64 {
	return impl.batteryCapacity()
}

type decorateRCTMeterEnergyImpl struct {
	meterEnergy func() (float64, error)
}

func (impl *decorateRCTMeterEnergyImpl) TotalEnergy() (float64, error) {
	return impl.meterEnergy()
}

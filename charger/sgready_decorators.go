package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateSgReady(base *SgReady, meter func() (float64, error), energyImport func() (float64, error), battery func() (float64, error), socLimiter func() (int64, error)) api.Charger {
	switch {
	case battery == nil && energyImport == nil && meter == nil && socLimiter == nil:
		return base

	case battery == nil && energyImport == nil && meter != nil && socLimiter == nil:
		return &struct {
			*SgReady
			api.Meter
		}{
			SgReady: base,
			Meter: &decorateSgReadyMeterImpl{
				meter: meter,
			},
		}

	case battery == nil && energyImport != nil && meter == nil && socLimiter == nil:
		return &struct {
			*SgReady
			api.EnergyImport
		}{
			SgReady: base,
			EnergyImport: &decorateSgReadyEnergyImportImpl{
				energyImport: energyImport,
			},
		}

	case battery == nil && energyImport != nil && meter != nil && socLimiter == nil:
		return &struct {
			*SgReady
			api.EnergyImport
			api.Meter
		}{
			SgReady: base,
			EnergyImport: &decorateSgReadyEnergyImportImpl{
				energyImport: energyImport,
			},
			Meter: &decorateSgReadyMeterImpl{
				meter: meter,
			},
		}

	case battery != nil && energyImport == nil && meter == nil && socLimiter == nil:
		return &struct {
			*SgReady
			api.Battery
		}{
			SgReady: base,
			Battery: &decorateSgReadyBatteryImpl{
				battery: battery,
			},
		}

	case battery != nil && energyImport == nil && meter != nil && socLimiter == nil:
		return &struct {
			*SgReady
			api.Battery
			api.Meter
		}{
			SgReady: base,
			Battery: &decorateSgReadyBatteryImpl{
				battery: battery,
			},
			Meter: &decorateSgReadyMeterImpl{
				meter: meter,
			},
		}

	case battery != nil && energyImport != nil && meter == nil && socLimiter == nil:
		return &struct {
			*SgReady
			api.Battery
			api.EnergyImport
		}{
			SgReady: base,
			Battery: &decorateSgReadyBatteryImpl{
				battery: battery,
			},
			EnergyImport: &decorateSgReadyEnergyImportImpl{
				energyImport: energyImport,
			},
		}

	case battery != nil && energyImport != nil && meter != nil && socLimiter == nil:
		return &struct {
			*SgReady
			api.Battery
			api.EnergyImport
			api.Meter
		}{
			SgReady: base,
			Battery: &decorateSgReadyBatteryImpl{
				battery: battery,
			},
			EnergyImport: &decorateSgReadyEnergyImportImpl{
				energyImport: energyImport,
			},
			Meter: &decorateSgReadyMeterImpl{
				meter: meter,
			},
		}

	case battery == nil && energyImport == nil && meter == nil && socLimiter != nil:
		return &struct {
			*SgReady
			api.SocLimiter
		}{
			SgReady: base,
			SocLimiter: &decorateSgReadySocLimiterImpl{
				socLimiter: socLimiter,
			},
		}

	case battery == nil && energyImport == nil && meter != nil && socLimiter != nil:
		return &struct {
			*SgReady
			api.Meter
			api.SocLimiter
		}{
			SgReady: base,
			Meter: &decorateSgReadyMeterImpl{
				meter: meter,
			},
			SocLimiter: &decorateSgReadySocLimiterImpl{
				socLimiter: socLimiter,
			},
		}

	case battery == nil && energyImport != nil && meter == nil && socLimiter != nil:
		return &struct {
			*SgReady
			api.EnergyImport
			api.SocLimiter
		}{
			SgReady: base,
			EnergyImport: &decorateSgReadyEnergyImportImpl{
				energyImport: energyImport,
			},
			SocLimiter: &decorateSgReadySocLimiterImpl{
				socLimiter: socLimiter,
			},
		}

	case battery == nil && energyImport != nil && meter != nil && socLimiter != nil:
		return &struct {
			*SgReady
			api.EnergyImport
			api.Meter
			api.SocLimiter
		}{
			SgReady: base,
			EnergyImport: &decorateSgReadyEnergyImportImpl{
				energyImport: energyImport,
			},
			Meter: &decorateSgReadyMeterImpl{
				meter: meter,
			},
			SocLimiter: &decorateSgReadySocLimiterImpl{
				socLimiter: socLimiter,
			},
		}

	case battery != nil && energyImport == nil && meter == nil && socLimiter != nil:
		return &struct {
			*SgReady
			api.Battery
			api.SocLimiter
		}{
			SgReady: base,
			Battery: &decorateSgReadyBatteryImpl{
				battery: battery,
			},
			SocLimiter: &decorateSgReadySocLimiterImpl{
				socLimiter: socLimiter,
			},
		}

	case battery != nil && energyImport == nil && meter != nil && socLimiter != nil:
		return &struct {
			*SgReady
			api.Battery
			api.Meter
			api.SocLimiter
		}{
			SgReady: base,
			Battery: &decorateSgReadyBatteryImpl{
				battery: battery,
			},
			Meter: &decorateSgReadyMeterImpl{
				meter: meter,
			},
			SocLimiter: &decorateSgReadySocLimiterImpl{
				socLimiter: socLimiter,
			},
		}

	case battery != nil && energyImport != nil && meter == nil && socLimiter != nil:
		return &struct {
			*SgReady
			api.Battery
			api.EnergyImport
			api.SocLimiter
		}{
			SgReady: base,
			Battery: &decorateSgReadyBatteryImpl{
				battery: battery,
			},
			EnergyImport: &decorateSgReadyEnergyImportImpl{
				energyImport: energyImport,
			},
			SocLimiter: &decorateSgReadySocLimiterImpl{
				socLimiter: socLimiter,
			},
		}

	case battery != nil && energyImport != nil && meter != nil && socLimiter != nil:
		return &struct {
			*SgReady
			api.Battery
			api.EnergyImport
			api.Meter
			api.SocLimiter
		}{
			SgReady: base,
			Battery: &decorateSgReadyBatteryImpl{
				battery: battery,
			},
			EnergyImport: &decorateSgReadyEnergyImportImpl{
				energyImport: energyImport,
			},
			Meter: &decorateSgReadyMeterImpl{
				meter: meter,
			},
			SocLimiter: &decorateSgReadySocLimiterImpl{
				socLimiter: socLimiter,
			},
		}
	}

	return nil
}

type decorateSgReadyBatteryImpl struct {
	battery func() (float64, error)
}

func (impl *decorateSgReadyBatteryImpl) Soc() (float64, error) {
	return impl.battery()
}

type decorateSgReadyEnergyImportImpl struct {
	energyImport func() (float64, error)
}

func (impl *decorateSgReadyEnergyImportImpl) EnergyImport() (float64, error) {
	return impl.energyImport()
}

type decorateSgReadyMeterImpl struct {
	meter func() (float64, error)
}

func (impl *decorateSgReadyMeterImpl) CurrentPower() (float64, error) {
	return impl.meter()
}

type decorateSgReadySocLimiterImpl struct {
	socLimiter func() (int64, error)
}

func (impl *decorateSgReadySocLimiterImpl) GetLimitSoc() (int64, error) {
	return impl.socLimiter()
}

package core

import (
	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/core/loadpoint"
)

// returns the current BatteryMode
func (site *Site) getBatteryMode() api.BatteryMode {
	site.Lock()
	defer site.Unlock()
	return site.batteryMode
}

// sets the current BatteryMode
func (site *Site) setBatteryMode(batMode api.BatteryMode) {
	site.Lock()
	defer site.Unlock()
	site.batteryMode = batMode
}

func (site *Site) updateBatteryMode(loadpoints []loadpoint.API) {
	// determine expected state
	batMode := api.BatteryNormal
	for _, lp := range loadpoints {
		if lp.GetStatus() == api.StatusC && (lp.GetMode() == api.ModeNow || lp.GetPlanActive()) {
			batMode = api.BatteryLocked
			break
		}
	}

	// update batteries
	if batMode != site.getBatteryMode() {
		for _, batMeter := range site.batteryMeters {
			if batCtrl, ok := batMeter.(api.BatteryController); ok {
				if err := batCtrl.SetBatteryMode(batMode); err != nil {
					site.log.ERROR.Println("battery mode:", err)
					return
				}
			}
		}
	}

	// update state and publish
	site.setBatteryMode(batMode)
	site.publish("batteryMode", batMode)
}

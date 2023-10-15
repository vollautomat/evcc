package core

// GetEffectivePriority returns the effective priority
func (lp *Loadpoint) GetEffectivePriority() int {
	if v := lp.GetVehicle(); v != nil {
		if res, err := v.OnIdentified().GetPriority(); err == nil {
			return res
		}
	}
	return lp.GetPriority()
}

// GetEffectiveMinCurrent returns the effective min current
func (lp *Loadpoint) GetEffectiveMinCurrent() float64 {
	if v := lp.GetVehicle(); v != nil {
		if res, err := v.OnIdentified().GetMinCurrent(); err == nil {
			return res
		}
	}
	return lp.GetMinCurrent()
}

// GetEffectiveMaxCurrent returns the effective max current
func (lp *Loadpoint) GetEffectiveMaxCurrent() float64 {
	if v := lp.GetVehicle(); v != nil {
		if res, err := v.OnIdentified().GetMaxCurrent(); err == nil {
			return res
		}
	}
	return lp.GetMaxCurrent()
}

// GetEffectiveLimitSoc returns the effective session limit soc.
// TODO take vehicle api limits into account
// TODO take plan into account?
func (lp *Loadpoint) GetEffectiveLimitSoc() int {
	lp.Lock()
	if lp.sessionLimitSoc > 0 {
		return lp.sessionLimitSoc
	}
	lp.Unlock()

	if v := lp.GetVehicle(); v != nil {
		if soc, err := v.OnIdentified().GetLimitSoc(); err == nil && soc > 0 {
			return soc
		}
	}

	return 100
}

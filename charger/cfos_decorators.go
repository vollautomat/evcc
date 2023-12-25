package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateCfos(base *CfosPowerBrain, phaseSwitcher func(int) error) api.Charger {
	switch {
	case phaseSwitcher == nil:
		return base

	case phaseSwitcher != nil:
		return &struct {
			*CfosPowerBrain
			api.PhaseSwitcher
		}{
			CfosPowerBrain: base,
			PhaseSwitcher: &decorateCfosPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}
	}

	return nil
}

type decorateCfosPhaseSwitcherImpl struct {
	phaseSwitcher func(int) error
}

func (impl *decorateCfosPhaseSwitcherImpl) Phases1p3p(p0 int) error {
	return impl.phaseSwitcher(p0)
}

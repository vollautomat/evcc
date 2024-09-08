package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decoratePulsares(base *Pulsares, phaseSwitcher func(int) error) api.Charger {
	switch {
	case phaseSwitcher == nil:
		return base

	case phaseSwitcher != nil:
		return &struct {
			*Pulsares
			api.PhaseSwitcher
		}{
			Pulsares: base,
			PhaseSwitcher: &decoratePulsaresPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}
	}

	return nil
}

type decoratePulsaresPhaseSwitcherImpl struct {
	phaseSwitcher func(int) error
}

func (impl *decoratePulsaresPhaseSwitcherImpl) Phases1p3p(p0 int) error {
	return impl.phaseSwitcher(p0)
}

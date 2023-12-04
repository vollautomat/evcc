package charger

// Code generated by github.com/evcc-io/evcc/cmd/tools/decorate.go. DO NOT EDIT.

import (
	"github.com/evcc-io/evcc/api"
)

func decorateAlfen(base *Alfen, phaseSwitcher func(int) error) api.Charger {
	switch {
	case phaseSwitcher == nil:
		return base

	case phaseSwitcher != nil:
		return &struct {
			*Alfen
			api.PhaseSwitcher
		}{
			Alfen: base,
			PhaseSwitcher: &decorateAlfenPhaseSwitcherImpl{
				phaseSwitcher: phaseSwitcher,
			},
		}
	}

	return nil
}

type decorateAlfenPhaseSwitcherImpl struct {
	phaseSwitcher func(int) error
}

func (impl *decorateAlfenPhaseSwitcherImpl) Phases1p3p(p0 int) error {
	return impl.phaseSwitcher(p0)
}

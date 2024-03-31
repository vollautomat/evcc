package core

import (
	"fmt"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/core/loadpoint"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/config"
)

var _ api.Circuit = (*Circuit)(nil)

// the circuit instances to control the load
type Circuit struct {
	log *util.Logger
	// uiChan chan<- util.Param

	parent   *Circuit   // parent circuit
	children []*Circuit // child circuits
	meter    api.Meter  // meter to determine current power

	maxCurrent float64 // max allowed current
	maxPower   float64 // max allowed power

	current float64
	power   float64
}

// NewCircuitFromConfig creates a new Circuit
func NewCircuitFromConfig(log *util.Logger, other map[string]interface{}) (*Circuit, error) {
	var cc struct {
		MaxCurrent float64 `mapstructure:"maxCurrent"` // the max allowed current of this circuit
		MaxPower   float64 `mapstructure:"maxPower"`   // the max allowed power of this circuit (kW)
		MeterRef   string  `mapstructure:"meter"`      // Charge meter reference
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	var meter api.Meter
	if cc.MeterRef != "" {
		dev, err := config.Meters().ByName(cc.MeterRef)
		if err != nil {
			return nil, err
		}
		meter = dev.Instance()
	}

	circuit, err := NewCircuit(log, cc.MaxCurrent, cc.MaxPower, meter)
	if err != nil {
		return nil, err
	}

	return circuit, err
}

// NewCircuit creates a circuit
func NewCircuit(log *util.Logger, maxCurrent, maxPower float64, meter api.Meter) (*Circuit, error) {
	c := &Circuit{
		log:        log,
		maxCurrent: maxCurrent,
		maxPower:   maxPower,
		meter:      meter,
	}

	if maxCurrent == 0 {
		c.log.DEBUG.Printf("validation of max phase current disabled")
	} else if _, ok := meter.(api.PhaseCurrents); !ok {
		return nil, fmt.Errorf("meter does not support phase currents")
	}

	if maxPower == 0 {
		c.log.DEBUG.Printf("validation of max power disabled")
	}

	return c, nil
}

// GetParent returns the parent circuit
func (c *Circuit) GetParent() *Circuit {
	return c.parent
}

// SetParent set parent circuit
func (c *Circuit) SetParent(parent *Circuit) {
	c.parent = parent
	if parent != nil {
		parent.RegisterChild(c)
	}
}

// RegisterChild registers child circuit
func (c *Circuit) RegisterChild(child *Circuit) {
	c.children = append(c.children, child)
}

func (c *Circuit) updateLoadpoints(loadpoints []loadpoint.API) {
	var totalPower, totalCurrent float64

	for _, lp := range loadpoints {
		if lp.GetCircuit() != c {
			continue
		}

		totalPower += lp.GetChargePower()
		totalCurrent += max(lp.GetChargeCurrents())
	}

	c.power = totalPower
	c.current = totalCurrent
}

func (c *Circuit) Update(loadpoints []loadpoint.API) error {
	// TODO retry
	if c.meter != nil {
		if f, err := c.meter.CurrentPower(); err == nil {
			// TODO handle negative powers
			c.power = f
		} else {
			return fmt.Errorf("circuit power: %w", err)
		}

		if phaseMeter, ok := c.meter.(api.PhaseCurrents); ok {
			if l1, l2, l3, err := phaseMeter.Currents(); err == nil {
				// TODO handle negative currents
				c.current = max(l1, l2, l3)
			} else {
				return fmt.Errorf("circuit currents: %w", err)
			}
		}

		for _, ch := range c.children {
			if err := ch.Update(loadpoints); err != nil {
				return err
			}
		}

		return nil
	}

	c.updateLoadpoints(loadpoints)

	for _, ch := range c.children {
		if err := ch.Update(loadpoints); err != nil {
			return err
		}

		c.power += ch.GetChargePower()
		c.current += ch.GetChargeCurrent()
	}

	return nil
}

// GetChargePower returns the actual power
func (c *Circuit) GetChargePower() float64 {
	return c.power
}

// GetChargeCurrent returns the actual current
func (c *Circuit) GetChargeCurrent() float64 {
	return c.current
}

// ValidateCurrent returns the actual current
func (c *Circuit) ValidateCurrent(old, new float64) float64 {
	delta := max(0, new-old)

	if c.maxCurrent == 0 || c.current+delta <= c.maxCurrent {
		return new
	}

	return max(0, c.maxCurrent-c.current)
}

// ValidatePower returns the actual power
func (c *Circuit) ValidatePower(old, new float64) float64 {
	delta := max(0, new-old)

	if c.maxPower != 0 && c.power+delta > c.maxPower {
		new = max(0, c.maxPower-c.power)
	}

	if c.parent != nil {
		return c.parent.ValidatePower(c.power, new)
	}

	return new
}

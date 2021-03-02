package charger

import (
	"fmt"
	"time"

	"github.com/andig/evcc/api"
	"github.com/andig/evcc/util"
	"github.com/andig/evcc/util/modbus"
	"github.com/volkszaehler/mbmd/meters/rs485"
)

const (
	phEMCPRegStatus     = 100 // Input
	phEMCPRegChargeTime = 102 // Input
	phEMCPRegMaxCurrent = 300 // Holding
	phEMCPRegEnable     = 400 // Coil

	phEMCPRegPower  = 120 // power reading
	phEMCPRegEnergy = 128 // energy reading

	phEMCPRegPowerScaler   = 364 // power reading scaler
	phEMCPRegEnergyScaler  = 372 // energy reading scaler
	phEMCPRegCurrentScaler = 358 // current reading scaler
)

var phEMCPRegCurrents = []uint16{114, 116, 118} // current readings

// PhoenixEMCP is an api.ChargeController implementation for Phoenix EM-CP-PP-ETH wallboxes.
// It uses Modbus TCP to communicate with the wallbox at modbus client id 180.
type PhoenixEMCP struct {
	conn                                  *modbus.Connection
	powerScale, energyScale, currentScale float64
}

func init() {
	registry.Add("phoenix-emcp", NewPhoenixEMCPFromConfig)
}

//go:generate go run ../cmd/tools/decorate.go -p charger -f decoratePhoenixEMCP -o phoenix-em-cp_decorators -b *PhoenixEMCP -r api.Charger -t "api.Meter,CurrentPower,func() (float64, error)" -t "api.MeterEnergy,TotalEnergy,func() (float64, error)" -t "api.MeterCurrent,Currents,func() (float64, float64, float64, error)"

// NewPhoenixEMCPFromConfig creates a Phoenix charger from generic config
func NewPhoenixEMCPFromConfig(other map[string]interface{}) (api.Charger, error) {
	cc := struct {
		URI   string
		ID    uint8
		Meter struct {
			Power, Energy, Currents bool
		}
	}{
		URI: "192.168.0.8:502", // default
		ID:  180,               // default
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	wb, err := NewPhoenixEMCP(cc.URI, cc.ID)

	var currentPower func() (float64, error)
	if cc.Meter.Power {
		currentPower = wb.currentPower
		wb.scaler(&wb.powerScale, phEMCPRegPowerScaler)
	}

	var totalEnergy func() (float64, error)
	if cc.Meter.Energy {
		totalEnergy = wb.totalEnergy
		wb.scaler(&wb.energyScale, phEMCPRegEnergyScaler)
	}

	var currents func() (float64, float64, float64, error)
	if cc.Meter.Currents {
		currents = wb.currents
		wb.scaler(&wb.currentScale, phEMCPRegCurrentScaler)
	}

	return decoratePhoenixEMCP(wb, currentPower, totalEnergy, currents), err
}

// NewPhoenixEMCP creates a Phoenix charger
func NewPhoenixEMCP(uri string, id uint8) (*PhoenixEMCP, error) {
	conn, err := modbus.NewConnection(uri, "", "", 0, false, id)
	if err != nil {
		return nil, err
	}

	log := util.NewLogger("emcp")
	conn.Logger(log.TRACE)

	wb := &PhoenixEMCP{
		conn: conn,
	}

	return wb, nil
}

// Status implements the Charger.Status interface
func (wb *PhoenixEMCP) Status() (api.ChargeStatus, error) {
	b, err := wb.conn.ReadInputRegisters(phEMCPRegStatus, 1)
	if err != nil {
		return api.StatusNone, err
	}

	return api.ChargeStatus(string(b[1])), nil
}

// Enabled implements the Charger.Enabled interface
func (wb *PhoenixEMCP) Enabled() (bool, error) {
	b, err := wb.conn.ReadCoils(phEMCPRegEnable, 1)
	if err != nil {
		return false, err
	}

	return b[0] == 1, nil
}

// Enable implements the Charger.Enable interface
func (wb *PhoenixEMCP) Enable(enable bool) error {
	var u uint16
	if enable {
		u = 0xFF00
	}

	_, err := wb.conn.WriteSingleCoil(phEMCPRegEnable, u)

	return err
}

// MaxCurrent implements the Charger.MaxCurrent interface
func (wb *PhoenixEMCP) MaxCurrent(current int64) error {
	if current < 6 {
		return fmt.Errorf("invalid current %d", current)
	}

	_, err := wb.conn.WriteSingleRegister(phEMCPRegMaxCurrent, uint16(current))

	return err
}

// ChargingTime yields current charge run duration
func (wb *PhoenixEMCP) ChargingTime() (time.Duration, error) {
	b, err := wb.conn.ReadInputRegisters(phEMCPRegChargeTime, 2)
	if err != nil {
		return 0, err
	}

	// 2 words, least significant word first
	secs := uint64(b[3])<<16 | uint64(b[2])<<24 | uint64(b[1]) | uint64(b[0])<<8
	return time.Duration(time.Duration(secs) * time.Second), nil
}

func (wb *PhoenixEMCP) scaler(val *float64, reg uint16) {
	*val = 1

	if b, err := wb.conn.ReadHoldingRegisters(reg, 2); err == nil {
		*val = rs485.RTUIeee754ToFloat64Swapped(b)
	}
}

func (wb *PhoenixEMCP) decodeReading(scale float64, b []byte) float64 {
	return scale * rs485.RTUUint32ToFloat64(b)
}

// CurrentPower implements the Meter.CurrentPower interface
func (wb *PhoenixEMCP) currentPower() (float64, error) {
	b, err := wb.conn.ReadInputRegisters(phEMCPRegPower, 2)
	if err != nil {
		return 0, err
	}

	return wb.decodeReading(wb.powerScale, b), err
}

// totalEnergy implements the Meter.TotalEnergy interface
func (wb *PhoenixEMCP) totalEnergy() (float64, error) {
	b, err := wb.conn.ReadInputRegisters(phEMCPRegEnergy, 2)
	if err != nil {
		return 0, err
	}

	return wb.decodeReading(wb.energyScale, b) / 1e3, err
}

// currents implements the Meter.Currents interface
func (wb *PhoenixEMCP) currents() (float64, float64, float64, error) {
	var currents []float64
	for _, regCurrent := range phEMCPRegCurrents {
		b, err := wb.conn.ReadInputRegisters(regCurrent, 2)
		if err != nil {
			return 0, 0, 0, err
		}

		currents = append(currents, wb.decodeReading(wb.currentScale, b))
	}

	return currents[0], currents[1], currents[2], nil
}

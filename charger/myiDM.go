package charger

import (
	"encoding/binary"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/modbus"
	"github.com/evcc-io/evcc/util/sponsor"
)

const (
	myiDMRegPvExcess = 74
	myiDMRegPower    = 4122
)

// MyiDM is an charger implementation for myiDM heat pumps.
type MyiDM struct {
	conn *modbus.Connection
}

func init() {
	registry.Add("myidm", NewMyiDMFromConfig)
}

// NewMyiDMFromConfig creates a myiDM charger from generic config
func NewMyiDMFromConfig(other map[string]interface{}) (api.Charger, error) {
	cc := modbus.TcpSettings{
		ID: 1,
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	return NewMyiDM(cc.URI, cc.ID)
}

// NewMyiDM creates a myiDM charger
func NewMyiDM(uri string, id uint8) (*MyiDM, error) {
	conn, err := modbus.NewConnection(uri, "", "", 0, modbus.Tcp, id)
	if err != nil {
		return nil, err
	}

	if !sponsor.IsAuthorized() {
		return nil, api.ErrSponsorRequired
	}

	log := util.NewLogger("myidm")
	conn.Logger(log.TRACE)

	if !sponsor.IsAuthorized() {
		return nil, api.ErrSponsorRequired
	}

	wb := &MyiDM{
		conn: conn,
	}

	return wb, nil
}

// Status implements the api.Charger interface
func (wb *MyiDM) Status() (api.ChargeStatus, error) {
	b, err := wb.conn.ReadHoldingRegisters(myiDMRegPower, 1)
	if err != nil {
		return api.StatusNone, err
	}

	if binary.BigEndian.Uint16(b) == 0 {
		return api.StatusB, nil
	}

	return api.StatusC, nil
}

// Enabled implements the api.Charger interface
func (wb *MyiDM) Enabled() (bool, error) {
	b, err := wb.conn.ReadHoldingRegisters(myiDMRegPvExcess, 1)
	if err != nil {
		return false, err
	}

	return binary.BigEndian.Uint16(b) > 0, nil
}

// Enable implements the api.Charger interface
func (wb *MyiDM) Enable(enable bool) error {
	var u uint16
	if enable {
		u = 1
	}

	_, err := wb.conn.WriteSingleRegister(myiDMRegPvExcess, u)

	return err
}

// MaxCurrent implements the api.Charger interface
func (wb *MyiDM) MaxCurrent(current int64) error {
	return wb.MaxCurrentMillis(float64(current))
}

var _ api.ChargerEx = (*MyiDM)(nil)

// MaxCurrentMillis implements the api.Charger interface
func (wb *MyiDM) MaxCurrentMillis(current float64) error {
	u := uint16(current) * 230 // 1p
	_, err := wb.conn.WriteSingleRegister(myiDMRegPvExcess, u)
	return err
}

var _ api.Meter = (*MyiDM)(nil)

// CurrentPower implements the api.Meter interface
func (wb *MyiDM) CurrentPower() (float64, error) {
	b, err := wb.conn.ReadHoldingRegisters(myiDMRegPower, 1)
	if err != nil {
		return 0, err
	}

	return float64(binary.BigEndian.Uint16(b)), nil
}

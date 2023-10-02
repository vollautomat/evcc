package charger

// LICENSE

// Copyright (c) 2023 premultiply

// This module is NOT covered by the MIT license. All rights reserved.

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/modbus"
)

// Pulsares charger implementation
type Pulsares struct {
	conn *modbus.Connection
	curr uint16
}

const (
	pulsaresRegStatus  = 0x1f
	pulsaresRegCurrent = 0x5d
)

func init() {
	registry.Add("Pulsares", NewPulsaresFromConfig)
}

// NewPulsaresFromConfig creates a Pulsares charger from generic config
func NewPulsaresFromConfig(other map[string]interface{}) (api.Charger, error) {
	cc := struct {
		modbus.Settings `mapstructure:",squash"`
		Timeout         time.Duration
	}{
		Settings: modbus.Settings{
			ID: 1,
		},
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	return NewPulsares(cc.URI, cc.Device, cc.Comset, cc.Baudrate, modbus.ProtocolFromRTU(cc.RTU), cc.ID, cc.Timeout)
}

// NewPulsares creates Pulsares charger
func NewPulsares(uri, device, comset string, baudrate int, proto modbus.Protocol, slaveID uint8, timeout time.Duration) (api.Charger, error) {
	conn, err := modbus.NewConnection(uri, device, comset, baudrate, proto, slaveID)
	if err != nil {
		return nil, err
	}

	if timeout > 0 {
		conn.Timeout(timeout)
	}

	log := util.NewLogger("pulsares")
	conn.Logger(log.TRACE)

	wb := &Pulsares{
		conn: conn,
		curr: 6000,
	}

	// get initial state from charger
	curr, err := wb.getCurrent()
	if err != nil {
		return nil, fmt.Errorf("current limit: %w", err)
	}
	if curr >= 6000 {
		wb.curr = curr
	}

	return wb, err
}

func (wb *Pulsares) setCurrent(current uint16) error {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(current))

	_, err := wb.conn.WriteMultipleRegisters(pulsaresRegCurrent, 1, b)

	return err
}

func (wb *Pulsares) getCurrent() (uint16, error) {
	b, err := wb.conn.ReadHoldingRegisters(pulsaresRegCurrent, 1)
	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint16(b), nil
}

// Status implements the api.Charger interface
func (wb *Pulsares) Status() (api.ChargeStatus, error) {
	b, err := wb.conn.ReadHoldingRegisters(pulsaresRegStatus, 1)
	if err != nil {
		return api.StatusNone, err
	}

	switch u := binary.BigEndian.Uint16(b); u {
	case 0:
		return api.StatusA, nil
	case 1, 2:
		return api.StatusB, nil
	case 3, 4:
		return api.StatusC, nil
	default:
		return api.StatusNone, fmt.Errorf("invalid status: %d", u)
	}
}

// Enabled implements the api.Charger interface
func (wb *Pulsares) Enabled() (bool, error) {
	curr, err := wb.getCurrent()

	return curr >= 6000, err
}

// Enable implements the api.Charger interface
func (wb *Pulsares) Enable(enable bool) error {
	var curr uint16
	if enable {
		curr = wb.curr
	}

	return wb.setCurrent(curr)
}

// MaxCurrent implements the api.Charger interface
func (wb *Pulsares) MaxCurrent(current int64) error {
	return wb.MaxCurrentMillis(float64(current))
}

var _ api.ChargerEx = (*Pulsares)(nil)

// MaxCurrent implements the api.ChargerEx interface
func (wb *Pulsares) MaxCurrentMillis(current float64) error {
	if current < 6 {
		return fmt.Errorf("invalid current %.1f", current)
	}

	wb.curr = uint16(current * 1e3)

	return wb.setCurrent(wb.curr)
}

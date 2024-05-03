package charger

// LICENSE

// Copyright (c) 2019-2024 andig

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
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/charger/semp"
	"github.com/evcc-io/evcc/charger/smaevcharger"
	"github.com/evcc-io/evcc/provider"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/request"
	"github.com/evcc-io/evcc/util/sponsor"
)

// Semp charger implementation
type Semp struct {
	*request.Helper
	log          *util.Logger
	uri          string // 192.168.XXX.XXX
	cache        time.Duration
	oldstate     float64
	measurementG provider.Cacheable[[]semp.Measurements]
	parameterG   provider.Cacheable[[]semp.Parameters]
}

func init() {
	registry.Add("semp", NewSempFromConfig)
}

// NewSempFromConfig creates a SEMP Charger from generic config
func NewSempFromConfig(other map[string]interface{}) (api.Charger, error) {
	cc := struct {
		Uri      string
		User     string
		Password string
		Cache    time.Duration
	}{
		Cache: 5 * time.Second,
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	if cc.Uri == "" {
		return nil, errors.New("missing uri")
	}

	// if cc.User == "" || cc.Password == "" {
	// 	return nil, api.ErrMissingCredentials
	// }

	// if cc.User == "admin" {
	// 	return nil, errors.New(`user "admin" not allowed, create new user`)
	// }

	return NewSemp(cc.Uri, cc.User, cc.Password, cc.Cache)
}

// NewSemp creates an SMA EV Charger
func NewSemp(uri, user, password string, cache time.Duration) (api.Charger, error) {
	log := util.NewLogger("semp").Redact(user, password)

	wb := &Semp{
		Helper: request.NewHelper(log),
		log:    log,
		uri:    util.DefaultScheme(strings.TrimRight(uri, "/"), "https") + "/SEMP",
		cache:  cache,
	}

	if !sponsor.IsAuthorized() {
		return nil, api.ErrSponsorRequired
	}

	var res semp.EM2Device
	if err := wb.GetJSON(wb.uri+"/", &res); err != nil {
		return wb, err
	}

	return wb, nil
}

// Status implements the api.Charger interface
func (wb *Semp) Status() (api.ChargeStatus, error) {
	status := api.StatusNone

	var res semp.EM2Device
	if err := wb.GetJSON(wb.uri+"/Parameters", &res); err != nil {
		return status, err
	}

	val, err := res.Parameters.Get("Measurement.Operation.EVeh.ChaStt")
	if err != nil {
		return status, errors.New("no parameters")
	}

	switch val {
	case smaevcharger.StatusA:
		return api.StatusA, nil
	case smaevcharger.StatusB:
		return api.StatusB, nil
	case smaevcharger.StatusC:
		return api.StatusC, nil
	default:
		return api.StatusNone, fmt.Errorf("invalid state: %.0f", val)
	}
}

// Enabled implements the api.Charger interface
func (wb *Semp) Enabled() (bool, error) {
	mode, err := wb.getParameter("Parameter.Chrg.ActChaMod")
	if err != nil {
		return false, err
	}

	switch mode {
	case semp.FastCharge, // Schnellladen - 4718
		semp.OptiCharge, // Optimiertes Laden - 4719
		semp.PlanCharge: // Laden mit Vorgabe - 4720
		return true, nil
	case semp.StopCharge: // Ladestopp - 4721
		return false, nil
	default:
		return false, fmt.Errorf("invalid charge mode: %s", mode)
	}
}

// Enable implements the api.Charger interface
func (wb *Semp) Enable(enable bool) error {
	if enable {
		res, err := wb.getMeasurement("Measurement.Chrg.ModSw")
		if err != nil {
			return err
		}

		if res == semp.SwitchOeko {
			// Switch in PV Loading position
			// If the selector switch of the wallbox is in the wrong position (eco-charging and not fast charging),
			// the charging process is started with eco-charging when it is activated,
			// which may be desired when integrated with SHM.
			// Since evcc does not have full control over the charging station in this mode,
			// a corresponding error is returned to indicate the incorrect switch position.
			// If the wallbox is installed without SHM, charging in eco mode is not possible.
			_ = wb.Send(value("Parameter.Chrg.ActChaMod", semp.OptiCharge))
			return fmt.Errorf("switch position not on fast charging - SMA's own optimized charging was activated")
		}

		// Switch in Fast charging position
		return wb.Send(value("Parameter.Chrg.ActChaMod", semp.FastCharge))
	}

	// else
	return wb.Send(value("Parameter.Chrg.ActChaMod", semp.StopCharge))
}

// MaxCurrent implements the api.Charger interface
func (wb *Semp) MaxCurrent(current int64) error {
	return wb.MaxCurrentMillis(float64(current))
}

var _ api.ChargerEx = (*Semp)(nil)

// maxCurrentMillis implements the api.ChargerEx interface
func (wb *Semp) MaxCurrentMillis(current float64) error {
	if current < 6 {
		return fmt.Errorf("invalid current %.5g", current)
	}

	return wb.Send(value("Parameter.Inverter.AcALim", fmt.Sprintf("%.2f", current)))
}

// var _ api.Meter = (*Semp)(nil)

// // CurrentPower implements the api.Meter interface
// func (wb *Semp) CurrentPower() (float64, error) {
// 	return wb.getMeasurement("Measurement.Metering.GridMs.TotWIn")
// }

// var _ api.ChargeRater = (*Semp)(nil)

// // ChargedEnergy implements the api.ChargeRater interface
// func (wb *Semp) ChargedEnergy() (float64, error) {
// 	res, err := wb.getMeasurement("Measurement.ChaSess.WhIn")
// 	return res / 1e3, err
// }

// var _ api.PhaseCurrents = (*Semp)(nil)

// // Currents implements the api.PhaseCurrents interface
// func (wb *Semp) Currents() (float64, float64, float64, error) {
// 	var res [3]float64

// 	for i, phase := range []string{"A", "B", "C"} {
// 		val, err := wb.getMeasurement("Measurement.GridMs.A.phs" + phase)
// 		if err != nil {
// 			return 0, 0, 0, err
// 		}

// 		res[i] = -val
// 	}

// 	return res[0], res[1], res[2], nil
// }

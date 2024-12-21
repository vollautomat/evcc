package charger

// LICENSE

// Copyright (c) 2024 andig

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
	"context"

	"github.com/WulfgarW/sensonet"
	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/request"
)

// Vaillant charger implementation
type Vaillant struct {
	*embed
	_mode    int64
	phases   int
	get      func() (int64, error)
	set      func(int64) error
	maxPower func(int64) error
}

func init() {
	registry.AddCtx("vaillant", NewVaillantFromConfig)
}

// go:generate go run ../cmd/tools/decorate.go -f decorateVaillant -b *Vaillant -r api.Charger -t "api.Meter,CurrentPower,func() (float64, error)" -t "api.MeterEnergy,TotalEnergy,func() (float64, error)" -t "api.Battery,Soc,func() (float64, error)" -t "api.SocLimiter,GetLimitSoc,func() (int64, error)"

// NewVaillantFromConfig creates an Vaillant configurable charger from generic config
func NewVaillantFromConfig(ctx context.Context, other map[string]interface{}) (api.Charger, error) {
	cc := struct {
		embed                      `mapstructure:",squash"`
		sensonet.CredentialsStruct `mapstructure:",squash"`
		Phases                     int
	}{
		embed: embed{
			Icon_:     "heatpump",
			Features_: []api.Feature{api.Heating, api.IntegratedDevice},
		},
		CredentialsStruct: sensonet.CredentialsStruct{
			Realm: sensonet.REALM_GERMANY,
		},
		Phases: 1,
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	client := request.NewClient(util.NewLogger("vaillant"))

	identity, err := sensonet.NewIdentity(client, &cc.CredentialsStruct)
	if err != nil {
		return nil, err
	}

	ts, err := identity.Login()
	if err != nil {
		return nil, err
	}

	conn, err := sensonet.NewConnection(client, ts)
	if err != nil {
		return nil, err
	}

	set := func(int64) error {
		_, err := conn.StartStrategybased("", sensonet.STRATEGY_HOTWATER_THEN_HEATING, nil, nil)
		return err
	}

	res := &SgReady{
		embed:  &cc.embed,
		phases: cc.Phases,
		set:    set,
	}

	return res, nil
}

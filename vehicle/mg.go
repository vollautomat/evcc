package vehicle

import (
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/vehicle/saic"
)

// MG is an api.Vehicle implementation for probably all SAIC cars
type MG struct {
	*embed
	*saic.Provider // provides the api implementations
}

func init() {
	registry.Add("mg", NewMGFromConfig)

}

// NewBMWFromConfig creates a new vehicle
func NewMGFromConfig(other map[string]interface{}) (api.Vehicle, error) {

	cc := struct {
		embed               `mapstructure:",squash"`
		User, Password, VIN string
		Cache               time.Duration
	}{
		Cache: interval,
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	if cc.User == "" || cc.Password == "" {
		return nil, api.ErrMissingCredentials
	}

	v := &MG{
		embed: &cc.embed,
	}

	log := util.NewLogger("MG").Redact(cc.User, cc.Password, cc.VIN)
	identity := saic.NewIdentity(log)

	ts, err := identity.Login(cc.User, cc.Password)
	if err != nil {
		return nil, err
	}

	api := saic.NewAPI(log, ts)

	if err == nil {
		v.Provider = saic.NewProvider(api, cc.VIN, cc.Cache)
	}

	return v, err
}

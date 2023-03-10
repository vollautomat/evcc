package vehicle

import (
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/vehicle/skoda"
	"github.com/evcc-io/evcc/vehicle/skoda/connect"
	"github.com/evcc-io/evcc/vehicle/vag/service"
)

// https://github.com/lendy007/skodaconnect

// Enyaq is an api.Vehicle implementation for Skoda Enyaq cars
type Enyaq struct {
	*embed
	*skoda.Provider // provides the api implementations
}

func init() {
	registry.Add("enyaq", NewEnyaqFromConfig)
}

// NewEnyaqFromConfig creates a new vehicle
func NewEnyaqFromConfig(other map[string]interface{}) (api.Vehicle, error) {
	var requestTimeout time.Duration
	cc := struct {
		embed               `mapstructure:",squash"`
		User, Password, VIN string
		Cache               time.Duration
		Timeout             string
	}{
		Cache:   interval,
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	if cc.Timeout == "" {
		requestTimeout , _ = time.ParseDuration("11s")
	} else {
		requestTimeout , _ = time.ParseDuration(cc.Timeout)
	}

	v := &Enyaq{
		embed: &cc.embed,
	}

	var err error
	log := util.NewLogger("enyaq").Redact(cc.User, cc.Password, cc.VIN)

	log.DEBUG.Printf("Timeout: %v", requestTimeout)

	// use Skoda credentials to resolve list of vehicles
	ts, err := service.TokenRefreshServiceTokenSource(log, skoda.TRSParams, skoda.AuthParams, cc.User, cc.Password)
	if err != nil {
		return nil, err
	}

	api := skoda.NewAPI(log, ts)
	api.Client.Timeout = requestTimeout

	vehicle, err := ensureVehicleEx(
		cc.VIN, api.Vehicles,
		func(v skoda.Vehicle) string {
			return v.VIN
		},
	)

	if v.Title_ == "" {
		v.Title_ = vehicle.Name
	}
	if v.Capacity_ == 0 {
		v.Capacity_ = float64(vehicle.Specification.Battery.CapacityInKWh)
	}

	// use Connect credentials to build provider
	if err == nil {
		ts, err := service.TokenRefreshServiceTokenSource(log, skoda.TRSParams, connect.AuthParams, cc.User, cc.Password)
		if err != nil {
			return nil, err
		}

		api := skoda.NewAPI(log, ts)
		api.Client.Timeout = requestTimeout

		v.Provider = skoda.NewProvider(api, vehicle.VIN, cc.Cache)
	}

	return v, err
}

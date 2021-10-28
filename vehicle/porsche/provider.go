package porsche

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/provider"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/request"
	"golang.org/x/oauth2"
)

// Provider is an api.Vehicle implementation for Porsche PHEV cars
type Provider struct {
	log *util.Logger
	*request.Helper
	accessTokens             AccessTokens
	identity                 *Identity
	carModel                 string
	vehicleSupportsEmobility bool
	statusG                  func() (interface{}, error)
	statusEmobilityG         func() (interface{}, error)
}

// NewProvider creates a new vehicle
func NewProvider(log *util.Logger, identity *Identity, accessTokens AccessTokens, vin string, cache time.Duration) *Provider {
	impl := &Provider{
		log:          log,
		Helper:       request.NewHelper(log),
		accessTokens: accessTokens,
		identity:     identity,
	}

	impl.Client.Jar = identity.Client.Jar

	impl.statusG = provider.NewCached(func() (interface{}, error) {
		return impl.status(vin)
	}, cache).InterfaceGetter()

	impl.statusEmobilityG = provider.NewCached(func() (interface{}, error) {
		return impl.statusEmobility(vin)
	}, cache).InterfaceGetter()

	if accessTokens.EmobilityToken.AccessToken != "" {
		impl.vehicleSupportsEmobility = true
	}

	return impl
}

func (v *Provider) request(token oauth2.Token, uri string) (*http.Request, error) {
	if token.AccessToken == "" || time.Since(token.Expiry) > 0 {
		accessTokens, err := v.identity.Login()
		if err != nil {
			return nil, err
		}
		v.accessTokens = accessTokens
	}

	req, err := request.New(http.MethodGet, uri, nil, map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token.AccessToken),
	})

	return req, err
}

// Status implements the vehicle status response
func (v *Provider) status(vin string) (StatusResponse, error) {
	var res StatusResponse

	uri := fmt.Sprintf("https://api.porsche.com/vehicle-data/de/de_DE/status/%s", vin)
	req, err := v.request(v.accessTokens.Token, uri)
	if err != nil {
		return res, err
	}

	err = v.DoJSON(req, &res)

	return res, err
}

// Status implements the vehicle status response
func (v *Provider) statusEmobility(vin string) (EmobilityResponse, error) {
	var res EmobilityResponse

	if !v.vehicleSupportsEmobility {
		return res, errors.New("vehicle does not support emobility")
	}

	if v.carModel == "" {
		// Note: As of 27.10.21 the capabilities API needs to be called AFTER a
		//   call to status() as it otherwise returns an HTTP 502 error.
		//   The reason is unknown, even when tested with 100% identical Headers.
		//   It seems to be a new backend related issue.

		if _, err := v.status(vin); err != nil {
			return res, err
		}

		uri := fmt.Sprintf("https://api.porsche.com/e-mobility/vcs/capabilities/%s", vin)

		req, err := v.request(v.accessTokens.EmobilityToken, uri)
		if err != nil {
			return res, err
		}

		var cr CapabilitiesResponse
		err = v.DoJSON(req, &cr)
		if err != nil {
			return res, err
		}
		v.carModel = cr.CarModel
	}

	uri := fmt.Sprintf("https://api.porsche.com/e-mobility/de/de_DE/%s/%s?timezone=Europe/Berlin", v.carModel, vin)
	req, err := v.request(v.accessTokens.EmobilityToken, uri)
	if err != nil {
		return res, err
	}

	err = v.DoJSON(req, &res)
	if err != nil && res.PcckErrorMessage != "" {
		err = errors.New(res.PcckErrorMessage)
	}

	return res, err
}

var _ api.Battery = (*Provider)(nil)

// SoC implements the api.Vehicle interface
func (v *Provider) SoC() (float64, error) {
	if v.vehicleSupportsEmobility {
		res, err := v.statusEmobilityG()
		if res, ok := res.(EmobilityResponse); err == nil && ok {
			return float64(res.BatteryChargeStatus.StateOfChargeInPercentage), nil
		}
	}

	res, err := v.statusG()
	if res, ok := res.(StatusResponse); err == nil && ok {
		return res.BatteryLevel.Value, nil
	}

	return 0, err
}

var _ api.VehicleRange = (*Provider)(nil)

// Range implements the api.VehicleRange interface
func (v *Provider) Range() (int64, error) {
	if v.vehicleSupportsEmobility {
		res, err := v.statusEmobilityG()
		if res, ok := res.(EmobilityResponse); err == nil && ok {
			return int64(res.BatteryChargeStatus.RemainingERange.ValueInKilometers), nil
		}
	}

	res, err := v.statusG()
	if res, ok := res.(StatusResponse); err == nil && ok {
		return int64(res.RemainingRanges.ElectricalRange.Distance.Value), nil
	}

	return 0, err
}

var _ api.VehicleFinishTimer = (*Provider)(nil)

// FinishTime implements the api.VehicleFinishTimer interface
func (v *Provider) FinishTime() (time.Time, error) {
	if !v.vehicleSupportsEmobility {
		return time.Time{}, api.ErrNotAvailable
	}

	res, err := v.statusEmobilityG()
	if res, ok := res.(EmobilityResponse); err == nil && ok {
		return time.Now().Add(time.Duration(res.BatteryChargeStatus.RemainingChargeTimeUntil100PercentInMinutes) * time.Minute), err
	}

	return time.Time{}, err
}

var _ api.ChargeState = (*Provider)(nil)

// Status implements the api.ChargeState interface
func (v *Provider) Status() (api.ChargeStatus, error) {
	if !v.vehicleSupportsEmobility {
		return api.StatusNone, api.ErrNotAvailable
	}

	res, err := v.statusEmobilityG()
	if res, ok := res.(EmobilityResponse); err == nil && ok {
		switch res.BatteryChargeStatus.PlugState {
		case "DISCONNECTED":
			return api.StatusA, nil
		case "CONNECTED":
			// ignore if the car is connected to a DC charging station
			if res.BatteryChargeStatus.ChargingInDCMode {
				return api.StatusA, nil
			}
			switch res.BatteryChargeStatus.ChargingState {
			case "OFF", "COMPLETED":
				return api.StatusB, nil
			case "ON":
				return api.StatusC, nil
			}
		}
	}

	return api.StatusNone, err
}

var _ api.VehicleClimater = (*Provider)(nil)

// Climater implements the api.VehicleClimater interface
func (v *Provider) Climater() (active bool, outsideTemp float64, targetTemp float64, err error) {
	if !v.vehicleSupportsEmobility {
		return active, outsideTemp, targetTemp, api.ErrNotAvailable
	}

	res, err := v.statusEmobilityG()
	if res, ok := res.(EmobilityResponse); err == nil && ok {
		switch res.DirectClimatisation.ClimatisationState {
		case "OFF":
			return false, 20, 20, nil
		case "ON":
			return true, 20, 20, nil
		}
	}

	return active, outsideTemp, targetTemp, err
}

var _ api.VehicleOdometer = (*Provider)(nil)

// Odometer implements the api.VehicleOdometer interface
func (v *Provider) Odometer() (float64, error) {
	res, err := v.statusG()
	if res, ok := res.(StatusResponse); err == nil && ok {
		return res.Mileage.Value, nil
	}

	return 0, err
}

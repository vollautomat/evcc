package charger

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/charger/smappee"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/request"
	"golang.org/x/oauth2"
)

// Smappee charger implementation
type Smappee struct {
	*request.Helper
	stationID     string
	connectorID   int32
	idTag         string
	token         string
	transactionID int32
	// statusG       provider.Cacheable[smappee.GetLatestStatus]
	// meterG        provider.Cacheable[smappee.GetLatestMeterValueResponse]
	cache time.Duration
}

func init() {
	registry.Add("smappee", NewSmappeeFromConfig)
}

// NewSmappeeFromConfig creates a Smappee charger from generic config
func NewSmappeeFromConfig(other map[string]interface{}) (api.Charger, error) {
	cc := struct {
		User, Password string
		ClientID       string
		ClientSecret   string
		Location       string
		Cache          time.Duration
	}{
		Cache: time.Second,
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	if cc.ClientID == "" || cc.ClientSecret == "" {
		return nil, errors.New("missing credentials")
	}

	oc := OAuth2Config
	oc.ClientID = cc.ClientID
	oc.ClientSecret = cc.ClientSecret

	return NewSmappee(cc.User, cc.Password, cc.Location, oc, cc.Cache)
}

var OAuth2Config = &oauth2.Config{
	Endpoint: oauth2.Endpoint{
		// AuthURL:  "https://app1pub.smappee.net/dev/v1/oauth2/token",
		TokenURL: "https://app1pub.smappee.net/dev/v1/oauth2/token",
	},
}

const ApiURL = "https://app1pub.smappee.net/dev/v3"

// NewSmappee creates Smappee charger
func NewSmappee(user, password, location string, oc *oauth2.Config, cache time.Duration) (*Smappee, error) {
	c := &Smappee{
		Helper: request.NewHelper(util.NewLogger("smappee")),
		// stationID: stationID,
		cache: cache,
	}

	clientCtx := context.WithValue(context.Background(), oauth2.HTTPClient, c.Client)
	ctx, cancel := context.WithTimeout(clientCtx, request.Timeout)
	defer cancel()

	token, err := oc.PasswordCredentialsToken(ctx, user, password)
	if err != nil {
		return nil, err
	}

	ts := oc.TokenSource(clientCtx, token)

	c.Client.Transport = &oauth2.Transport{
		Base:   c.Client.Transport,
		Source: oauth2.ReuseTokenSource(token, ts),
	}

	// c.statusG = provider.ResettableCached(func() (smappee.GetLatestStatus, error) {
	// 	var res smappee.GetLatestStatus
	// 	err := c.GetJSON(fmt.Sprintf("%s/cs/%s/status", smappee.BASE_URL, c.stationID), &res)
	// 	return res, err
	// }, c.cache)

	// c.meterG = provider.ResettableCached(func() (smappee.GetLatestMeterValueResponse, error) {
	// 	var res smappee.GetLatestMeterValueResponse
	// 	err := c.GetJSON(fmt.Sprintf("%s/cs/%s/metervalue", smappee.BASE_URL, c.stationID), &res)
	// 	return res, err
	// }, c.cache)

	loc, serial, err := ensureChargerWithFeature(location, func() ([]smappee.ServiceLocation, error) {
		var res smappee.ServiceLocations
		if err := c.GetJSON(ApiURL+"/servicelocation", &res); err != nil {
			return nil, err
		}
		return res.ServiceLocations, nil
	}, func(v smappee.ServiceLocation) (string, string) {
		return v.Name, v.DeviceSerialNumber
	})
	if err != nil {
		return nil, err
	}
	_ = loc
	_ = serial

	data := struct {
		Mode string `json:"mode"`
	}{
		Mode: "SMART",
	}

	var res any
	uri := fmt.Sprintf("%s/chargingstations/%s/connectors/%d/mode", ApiURL, serial, 0)
	req, _ := request.New(http.MethodPut, uri, request.MarshalJSON(data), request.JSONEncoding)
	if err := c.DoJSON(req, &res); err != nil {
		return nil, err
	}

	return c, nil
}

// Status implements the api.Charger interface
func (c *Smappee) Status() (api.ChargeStatus, error) {
	return api.StatusNone, nil
	// res, err := c.statusG.Get()
	// if err != nil {
	// 	return api.StatusNone, err
	// }

	// status := smappee.ChargePointStatus(res.Status)
	// switch status {
	// case smappee.AVAILABLE:
	// 	return api.StatusA, nil
	// case smappee.PREPARING:
	// 	return api.StatusB, nil
	// case smappee.CHARGING, smappee.FINISHING:
	// 	return api.StatusC, nil
	// case smappee.FAULTED:
	// 	return api.StatusF, nil
	// default:
	// 	return api.StatusNone, fmt.Errorf("invalid status: %s", res.Status)
	// }
}

// Enabled implements the api.Charger interface
func (c *Smappee) Enabled() (bool, error) {
	// res, err := c.statusG.Get()
	// return res.Status == string(smappee.CHARGING), err
	return false, nil
}

// Enable implements the api.Charger interface
func (c *Smappee) Enable(enable bool) error {
	return nil
}

// MaxCurrent implements the api.Charger interface
func (c *Smappee) MaxCurrent(current int64) error {
	return nil
}

// var _ api.Meter = (*Smappee)(nil)

// // CurrentPower implements the api.Meter interface
// func (c *Smappee) CurrentPower() (float64, error) {
// 	res, err := c.meterG.Get()
// 	return float64(res.ActivePowerImport * 1e3), err
// }

// var _ api.MeterEnergy = (*Smappee)(nil)

// // TotalEnergy implements the api.MeterMeterEnergy interface
// func (c *Smappee) TotalEnergy() (float64, error) {
// 	res, err := c.meterG.Get()
// 	return float64(res.EnergyActiveImportRegister), err
// }

// var _ api.PhaseCurrents = (*Smappee)(nil)

// // Currents implements the api.PhaseCurrents interface
// func (c *Smappee) Currents() (float64, float64, float64, error) {
// 	res, err := c.meterG.Get()
// 	return float64(res.CurrentImportPhaseL1), float64(res.CurrentImportPhaseL2), float64(res.CurrentImportPhaseL3), err
// }

package toyota

import (
	"fmt"
	"net/http"
	"time"

	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/request"
	"golang.org/x/oauth2"
)

const (
	BaseUrl                  = "https://oneapp:oneapp@b2c-login.toyota-europe.com"
	ApiBaseUrl               = "https://ctpa-oneapi.tceu-ctp-prd.toyotaconnectedeurope.io"
	AccessTokenPath          = "oauth2/realms/root/realms/tme/access_token"
	AuthenticationPath       = "json/realms/root/realms/tme/authenticate?authIndexType=service&authIndexValue=oneapp"
	AuthorizationPath        = "oauth2/realms/root/realms/tme/authorize?client_id=oneapp&scope=openid+profile+write&response_type=code&redirect_uri=com.toyota.oneapp:/oauth2Callback&code_challenge=plain&code_challenge_method=plain"
	VehicleGuidPath          = "v2/vehicle/guid"
	RemoteElectricStatusPath = "v1/global/remote/electric/status"
)

type API struct {
	*request.Helper
	log *util.Logger
}

func NewAPI(log *util.Logger, identity oauth2.TokenSource) *API {
	v := &API{
		Helper: request.NewHelper(log),
		log:    log,
	}

	// api is unbelievably slow when retrieving status
	v.Client.Timeout = 120 * time.Second

	// replace client transport with authenticated transport
	v.Client.Transport = &oauth2.Transport{
		Source: identity,
		Base:   v.Client.Transport,
	}

	return v
}

func (v *API) Vehicles() ([]string, error) {
	uri := fmt.Sprintf("%s/%s", ApiBaseUrl, VehicleGuidPath)

	identity, ok := v.Client.Transport.(*oauth2.Transport)
	if !ok || identity.Source == nil {
		return nil, fmt.Errorf("missing identity")
	}

	id, ok := identity.Source.(*Identity)
	if !ok {
		return nil, fmt.Errorf("invalid identity type")
	}

	req, err := request.New(http.MethodGet, uri, nil, map[string]string{
		"Accept":    "application/json",
		"x-guid":    id.uuid,
		"x-api-key": "tTZipv6liF74PwMfk9Ed68AQ0bISswwf3iHQdqcF",
	})
	var vehiclesResponse Vehicles
	if err == nil {
		err = v.DoJSON(req, &vehiclesResponse)
	}
	var vehicles []string
	for _, v := range vehiclesResponse.Payload {
		vehicles = append(vehicles, v.VIN)
	}
	return vehicles, err
}

func (v *API) Status(vin string) (Status, error) {
	uri := fmt.Sprintf("%s/%s", ApiBaseUrl, RemoteElectricStatusPath)
	identity, ok := v.Client.Transport.(*oauth2.Transport)
	if !ok || identity.Source == nil {
		return Status{}, fmt.Errorf("missing identity")
	}
	id, ok := identity.Source.(*Identity)
	if !ok {
		return Status{}, fmt.Errorf("invalid identity type")
	}
	req, err := request.New(http.MethodGet, uri, nil, map[string]string{
		"Accept":    "application/json",
		"x-guid":    id.uuid,
		"x-api-key": "tTZipv6liF74PwMfk9Ed68AQ0bISswwf3iHQdqcF",
		"vin":       vin,
	})
	var status Status
	if err == nil {
		err = v.DoJSON(req, &status)
	}
	return status, err
}

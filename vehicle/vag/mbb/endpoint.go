package mbb

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/request"
	"github.com/evcc-io/evcc/vehicle/vag"
)

const (
	BaseURL  = "https://mbboauth-1d.prd.ece.vwg-connect.com"
	TokenURL = BaseURL + "/mbbcoauth/mobile/oauth2/v1/token"
)

type Service struct {
	*request.Helper
	clientID string
}

func New(log *util.Logger, clientID string) *Service {
	return &Service{
		Helper:   request.NewHelper(log),
		clientID: clientID,
	}
}

func (v *Service) Exchange(idToken, code string) (*vag.Token, error) {
	data := url.Values(map[string][]string{
		"grant_type": {"id_token"},
		"token":      {idToken},
		"scope":      {"sc2:fal"},
	})

	req, err := request.New(http.MethodPost, TokenURL, strings.NewReader(data.Encode()), map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"X-Client-Id":  v.clientID,
	})

	var res vag.Token
	if err == nil {
		err = v.DoJSON(req, &res)
	}

	// check if token response contained error
	if errT := res.Error(); err != nil && errT != nil {
		err = fmt.Errorf("token exchange: %w", errT)
	}

	return &res, err
}

func (v *Service) Refresh(token *vag.Token) (*vag.Token, error) {
	data := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {token.RefreshToken},
		"scope":         {"sc2:fal"},
	}

	req, err := request.New(http.MethodPost, TokenURL, strings.NewReader(data.Encode()), map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"X-Client-Id":  v.clientID,
	})

	var res vag.Token
	if err == nil {
		err = v.DoJSON(req, &res)
	}

	return &res, err
}

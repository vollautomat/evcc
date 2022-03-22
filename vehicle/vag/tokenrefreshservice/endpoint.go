package tokenrefreshservice

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/request"
	"golang.org/x/oauth2"
)

const (
	BaseURL         = "https://tokenrefreshservice.apps.emea.vwapps.io"
	CodeExchangeURL = BaseURL + "/exchangeAuthCode"
	RefreshTokenURL = BaseURL + "/refreshTokens"
)

type Token oauth2.Token

type Service struct {
	*request.Helper
}

func New(log *util.Logger) (*Service, error) {
	v := &Service{
		Helper: request.NewHelper(log),
	}
	return v, nil
}

func (v *Service) Exchange(q url.Values, idToken, code string) (*Token, error) {
	if idToken == "" {
		return nil, errors.New("missing id_token")
	}
	if code == "" {
		return nil, errors.New("missing code")
	}

	data := url.Values{
		"auth_code": {code},
		"id_token":  {idToken},
	}

	for k, v := range q {
		data[k] = v
	}

	var res Token

	req, err := request.New(http.MethodPost, CodeExchangeURL, strings.NewReader(data.Encode()), request.URLEncoding)
	if err == nil {
		err = v.DoJSON(req, &res)
	}

	return &res, err
}

func (v *Service) Refresh(q url.Values, token *Token) (*Token, error) {
	data := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {token.RefreshToken},
	}

	for k, v := range q {
		data[k] = v
	}

	var res Token

	req, err := request.New(http.MethodPost, RefreshTokenURL, strings.NewReader(data.Encode()), request.URLEncoding)
	if err == nil {
		err = v.DoJSON(req, &res)
	}

	return &res, err
}

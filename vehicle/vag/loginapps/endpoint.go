package loginapps

import (
	"net/http"
	"net/url"

	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/oauth"
	"github.com/evcc-io/evcc/util/request"
	"golang.org/x/oauth2"
)

const (
	BaseURL = "https://login.apps.emea.vwapps.io"
)

var Endpoint = &oauth2.Endpoint{
	AuthURL:  BaseURL + "/login/v1",
	TokenURL: BaseURL + "/refresh/v1",
}

type Service struct {
	*request.Helper
	clientID string
}

func New(log *util.Logger) *Service {
	return &Service{
		Helper: request.NewHelper(log),
	}
}

func (v *Service) Exchange(q url.Values) (*Token, error) {
	if err := util.RequireValues(q, "state", "id_token", "access_token", "code"); err != nil {
		return nil, err
	}

	data := map[string]string{
		"region":            "emea",
		"redirect_uri":      "weconnect://authenticated",
		"state":             q.Get("state"),
		"id_token":          q.Get("id_token"),
		"access_token":      q.Get("access_token"),
		"authorizationCode": q.Get("code"),
	}

	var res Token

	req, err := request.New(http.MethodPost, Endpoint.AuthURL, request.MarshalJSON(data), request.JSONEncoding)
	if err == nil {
		err = v.DoJSON(req, &res)
	}

	return &res, err
}

func (v *Service) Refresh(token *Token) (*Token, error) {
	req, err := request.New(http.MethodGet, Endpoint.TokenURL, nil, map[string]string{
		"Accept":        "application/json",
		"Authorization": "Bearer " + token.RefreshToken,
	})

	var res Token
	if err == nil {
		err = v.DoJSON(req, &res)
	}

	return &res, err
}

// RefreshToken implements oauth.TokenRefresher
func (v *Service) RefreshToken(token *oauth2.Token) (*oauth2.Token, error) {
	res, err := v.Refresh((*Token)(token))
	return (*oauth2.Token)(res), err
}

// TokenSource creates a refreshing OAuth2 token source
func (v *Service) TokenSource(token *Token) oauth2.TokenSource {
	return oauth.RefreshTokenSource((*oauth2.Token)(token), v)
}

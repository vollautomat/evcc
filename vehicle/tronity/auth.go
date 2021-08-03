package tronity

import (
	"errors"

	"golang.org/x/oauth2"
)

const URI = "https://api-eu.tronity.io"

func OAuth2Config(id, secret string) (*oauth2.Config, error) {
	if id == "" {
		return nil, errors.New("missing client id")
	}

	if secret == "" {
		return nil, errors.New("missing client secret")
	}

	return &oauth2.Config{
		ClientID:     id,
		ClientSecret: secret,
		RedirectURL:  "http://localhost:8080",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://api-eu.tronity.io/oauth/authorize",
			TokenURL: "https://api-eu.tronity.io/oauth/authentication",
		},
		Scopes: []string{"read_vin", "read_vehicle_info", "read_odometer", "read_charge", "read_charge", "read_battery", "read_location", "write_charge_start_stop", "write_wake_up"},
	}, nil
}

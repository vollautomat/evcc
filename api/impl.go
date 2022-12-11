package api

import (
	"encoding"
	"errors"
	"fmt"
	"strings"
	"time"
)

// ChargeModeString converts string to ChargeMode
func ChargeModeString(mode string) (ChargeMode, error) {
	switch strings.ToLower(mode) {
	case string(ModeEmpty):
		return ModeEmpty, nil // undefined
	case string(ModeNow):
		return ModeNow, nil
	case string(ModeMinPV):
		return ModeMinPV, nil
	case string(ModePV):
		return ModePV, nil
	case string(ModeOff):
		return ModeOff, nil
	default:
		return "", fmt.Errorf("invalid value: %s", mode)
	}
}

var _ encoding.TextUnmarshaler = (*ChargeMode)(nil)

func (c *ChargeMode) UnmarshalText(text []byte) error {
	casted, err := ChargeModeString(string(text))
	if err != nil {
		return err
	}

	*c = casted

	return nil
}

// Current returns the rates current rate or error
func (r Rates) Current(now time.Time) (Rate, error) {
	for _, rr := range r {
		if (rr.Start.Before(now) || rr.Start.Equal(now)) && rr.End.After(now) {
			return rr, nil
		}
	}

	return Rate{}, errors.New("rates unavailable")
}

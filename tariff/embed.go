package tariff

import "github.com/evcc-io/evcc/tariff/schedule"

type embed struct {
	Charges float64 `mapstructure:"charges"`
	Tax     float64 `mapstructure:"tax"`
	zones   schedule.Zones
}

func (t *embed) zonedCosts(cc schedule.Config) error {
	zones, err := schedule.FromConfig(cc)
	if err != nil {
		return err
	}

	// prepend catch-all zone
	t.zones = append([]schedule.Zone{
		{Price: t.Charges}, // full week is implicit
	}, zones...)

	return nil
}

func (t *embed) totalPrice(price float64) float64 {
	return (price + t.Charges) * (1 + t.Tax)
}

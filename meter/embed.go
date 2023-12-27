package meter

import "github.com/evcc-io/evcc/api"

type embed struct {
	Title_ string `mapstructure:"title"`
}

var _ api.Denominator = (*embed)(nil)

// Title implements the api.Vehicle interface
func (v *embed) Title() string {
	return v.Title_
}

// SetTitle implements the api.TitleSetter interface
func (v *embed) SetTitle(title string) {
	v.Title_ = title
}

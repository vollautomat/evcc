package cmd

import (
	"strings"
	"testing"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/core"
	"github.com/evcc-io/evcc/util"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

const sample = `
loadpoints:
- mode: off
`

func TestMode(t *testing.T) {
	require.Contains(t, api.ChargeModeStrings(), "off")
}

func TestYamlOff(t *testing.T) {
	var conf globalConfig
	viper.SetConfigType("yaml")

	require.NoError(t, viper.ReadConfig(strings.NewReader(sample)))
	require.NoError(t, viper.UnmarshalExact(&conf))

	var lp core.Loadpoint
	require.NoError(t, util.DecodeOther(conf.Loadpoints[0], &lp))

	require.Equal(t, api.ModeOff, lp.Mode_)
}

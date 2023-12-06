package cmd

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/fatih/structs"
	"github.com/jeremywohl/flatten"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/maps"
)

func TestUnwrap(t *testing.T) {
	err := fmt.Errorf("foo: %w", fmt.Errorf("bar %w", errors.New("baz")))

	res := unwrap(err)
	if exp := []string{"foo", "bar", "baz"}; !reflect.DeepEqual(res, exp) {
		t.Errorf("expected %v, got %v", exp, res)
	}
}

func TestRedact(t *testing.T) {
	secret := `
	# sponsor token is a public token
	sponsortoken: geheim
	user: geheim
	password: geheim
	secret: geheim
	token:
		access: geheim
		refresh: geheim
	pin: geheim
	mac: geheim
	secret: geheim # comment
	secret : geheim
	`

	if res := redact(secret); strings.Contains(res, "geheim") || !strings.Contains(res, "public") {
		t.Errorf("secret exposed: %v", res)
	}
}

func TestEnv(t *testing.T) {
	v := viper.New()

	// initConfig()
	v.SetEnvPrefix("evcc")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// register all known config keys
	flat, _ := flatten.Flatten(structs.Map(conf), "", flatten.DotStyle)
	for _, k := range maps.Keys(flat) {
		t.Log(k)
		_ = v.BindEnv(k)
	}

	os.Setenv("EVCC_DATABASE_DSN", "dsn")
	os.Setenv("EVCC_SITE_TITLE", "foo")
	// os.Setenv("EVCC_LOADPOINTS_0_TITLE", "baz")

	yaml := `
database:
  type: sqlite
  dsn: bar
`

	v.SetConfigType("yaml")
	assert.NoError(t, v.ReadConfig(strings.NewReader(yaml)))

	var conf globalConfig
	assert.NoError(t, v.UnmarshalExact(&conf))

	// predeclared struct
	assert.Equal(t, "sqlite", conf.Database.Type)
	assert.Equal(t, "dsn", conf.Database.Dsn)

	// map[string]interface{}
	assert.Equal(t, "foo", conf.Site["title"])

	// []map[string]interface{}
	// if assert.Len(t, conf.Loadpoints, 1) {
	// 	assert.Equal(t, "baz", conf.Loadpoints[0]["title"])
	// }
}

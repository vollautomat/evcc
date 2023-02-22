package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/evcc-io/evcc/provider/mqtt"
	"github.com/evcc-io/evcc/push"
	"github.com/evcc-io/evcc/server"
	"github.com/evcc-io/evcc/server/db/settings"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/modbus"
	"github.com/fatih/structs"
)

type config struct {
	URI          interface{} // TODO deprecated
	Network      networkConfig
	Log          string
	SponsorToken string
	Plant        string // telemetry plant id
	Telemetry    bool
	Metrics      bool
	Profile      bool
	Levels       map[string]string
	Interval     time.Duration
	Database     dbConfig
	Mqtt         mqttConfig
	ModbusProxy  []proxyConfig
	Javascript   []javascriptConfig
	Influx       server.InfluxConfig
	EEBus        map[string]interface{}
	HEMS         typedConfig
	Messaging    messagingConfig
	Meters       []qualifiedConfig
	Chargers     []qualifiedConfig
	Vehicles     []qualifiedConfig
	Tariffs      tariffConfig
	Site         map[string]interface{}
	Loadpoints   []map[string]interface{}
}

type mqttConfig struct {
	mqtt.Config `mapstructure:",squash"`
	Topic       string
}

type javascriptConfig struct {
	VM     string
	Script string
}

type proxyConfig struct {
	Port            int
	ReadOnly        bool
	modbus.Settings `mapstructure:",squash"`
}

type dbConfig struct {
	Type string
	Dsn  string
}

type qualifiedConfig struct {
	Name, Type string
	Other      map[string]interface{} `mapstructure:",remain"`
}

type typedConfig struct {
	Type  string
	Other map[string]interface{} `mapstructure:",remain"`
}

type messagingConfig struct {
	Events   map[string]push.EventTemplateConfig
	Services []typedConfig
}

type tariffConfig struct {
	Currency string
	Grid     typedConfig
	FeedIn   typedConfig
	Planner  typedConfig
}

type networkConfig struct {
	Schema string
	Host   string
	Port   int
}

// loadConfigFromDB loads all known config keys from database
func loadConfigFromDB(conf *config) error {
	for key, val := range settings.All() {
		println(key)

		fields := structs.Fields(conf)
		segments := strings.Split(key, ".")

	SEGMENTS:
		for i, sub := range segments {
			fmt.Println(sub)

			for _, field := range fields {
				if !strings.EqualFold(field.Name(), sub) {
					continue
				}

				if i < len(segments)-1 {
					fmt.Println("found field", field.Name())
					fields = field.Fields()
				} else {
					fmt.Println("final field", field.Name())
				}

				continue SEGMENTS
			}
		}

		_ = val
		fmt.Println()
	}

	os.Exit(0)
	return util.DecodeOther(nil, conf)
}

func main() {
	var conf config
	loadConfigFromDB(&conf)
}

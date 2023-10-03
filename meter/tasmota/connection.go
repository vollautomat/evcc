package tasmota

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/evcc-io/evcc/provider"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/request"
	"github.com/evcc-io/evcc/util/transport"
)

// Connection is the Tasmota connection
type Connection struct {
	*request.Helper
	uri, user, password string
	channels            []int
	statusSnsG          provider.Cacheable[StatusSNSResponse]
	statusStsG          provider.Cacheable[StatusSTSResponse]
}

// NewConnection creates a Tasmota connection
func NewConnection(uri, user, password string, channels []int, cache time.Duration) (*Connection, error) {
	if uri == "" {
		return nil, errors.New("missing uri")
	}

	minchannel := 8
	maxchannel := 1
	duplicatechannels := make(map[int]bool, 0)
	for i := 0; i < len(channels); i++ {
		if duplicatechannels[channels[i]] {
			return nil, errors.New("duplicates in channel list")
		} else {
			duplicatechannels[channels[i]] = true
			if channels[i] < minchannel {
				minchannel = channels[i]
			}
			if channels[i] > maxchannel {
				maxchannel = channels[i]
			}
		}
	}
	if minchannel < 1 || maxchannel > 8 {
		return nil, errors.New("only channels 1-8 allowed")
	}

	log := util.NewLogger("tasmota")
	c := &Connection{
		Helper:   request.NewHelper(log),
		uri:      util.DefaultScheme(strings.TrimRight(uri, "/"), "http"),
		user:     user,
		password: password,
		channels: channels,
	}

	c.Client.Transport = request.NewTripper(log, transport.Insecure())

	c.statusSnsG = provider.ResettableCached(func() (StatusSNSResponse, error) {
		parameters := url.Values{
			"user":     []string{c.user},
			"password": []string{c.password},
			"cmnd":     []string{"Status 8"},
		}
		var res StatusSNSResponse
		err := c.GetJSON(fmt.Sprintf("%s/cm?%s", c.uri, parameters.Encode()), &res)
		return res, err
	}, cache)

	c.statusStsG = provider.ResettableCached(func() (StatusSTSResponse, error) {
		parameters := url.Values{
			"user":     []string{c.user},
			"password": []string{c.password},
			"cmnd":     []string{"Status 0"},
		}
		var res StatusSTSResponse
		err := c.GetJSON(fmt.Sprintf("%s/cm?%s", c.uri, parameters.Encode()), &res)
		return res, err
	}, cache)

	return c, c.ChannelExists()
}

// channelExists checks the existence of the configured relay channel interface
func (c *Connection) ChannelExists() error {
	res, err := c.statusStsG.Get()
	if err != nil {
		return err
	}

	var ok bool
	for _, channel := range c.channels {
		switch channel {
		case 1:
			ok = res.StatusSTS.Power != "" || res.StatusSTS.Power1 != ""
		case 2:
			ok = res.StatusSTS.Power2 != ""
		case 3:
			ok = res.StatusSTS.Power3 != ""
		case 4:
			ok = res.StatusSTS.Power4 != ""
		case 5:
			ok = res.StatusSTS.Power5 != ""
		case 6:
			ok = res.StatusSTS.Power6 != ""
		case 7:
			ok = res.StatusSTS.Power7 != ""
		case 8:
			ok = res.StatusSTS.Power8 != ""
		}

		if !ok {
			return fmt.Errorf("invalid relay channel: %d", channel)
		}
	}

	return nil
}

// Enable implements the api.Charger interface
func (c *Connection) Enable(enable bool) error {
	for _, channel := range c.channels {

		cmd := fmt.Sprintf("Power%d off", channel)
		if enable {
			cmd = fmt.Sprintf("Power%d on", channel)
		}

		parameters := url.Values{
			"user":     []string{c.user},
			"password": []string{c.password},
			"cmnd":     []string{cmd},
		}

		var res PowerResponse
		if err := c.GetJSON(fmt.Sprintf("%s/cm?%s", c.uri, parameters.Encode()), &res); err != nil {
			return err
		}

		var on bool
		switch channel {
		case 2:
			on = strings.ToUpper(res.Power2) == "ON"
		case 3:
			on = strings.ToUpper(res.Power3) == "ON"
		case 4:
			on = strings.ToUpper(res.Power4) == "ON"
		case 5:
			on = strings.ToUpper(res.Power5) == "ON"
		case 6:
			on = strings.ToUpper(res.Power6) == "ON"
		case 7:
			on = strings.ToUpper(res.Power7) == "ON"
		case 8:
			on = strings.ToUpper(res.Power8) == "ON"
		default:
			on = strings.ToUpper(res.Power) == "ON" || strings.ToUpper(res.Power1) == "ON"
		}

		switch {
		case enable && !on:
			return errors.New("switchOn failed")
		case !enable && on:
			return errors.New("switchOff failed")
		}
	}

	c.statusSnsG.Reset()
	c.statusStsG.Reset()

	return nil
}

// Enabled implements the api.Charger interface
func (c *Connection) Enabled() (bool, error) {
	res, err := c.statusStsG.Get()
	if err != nil {
		return false, err
	}

	for _, channel := range c.channels {
		switch channel {
		case 2:
			return strings.ToUpper(res.StatusSTS.Power2) == "ON", err
		case 3:
			return strings.ToUpper(res.StatusSTS.Power3) == "ON", err
		case 4:
			return strings.ToUpper(res.StatusSTS.Power4) == "ON", err
		case 5:
			return strings.ToUpper(res.StatusSTS.Power5) == "ON", err
		case 6:
			return strings.ToUpper(res.StatusSTS.Power6) == "ON", err
		case 7:
			return strings.ToUpper(res.StatusSTS.Power7) == "ON", err
		case 8:
			return strings.ToUpper(res.StatusSTS.Power8) == "ON", err
		default:
			return strings.ToUpper(res.StatusSTS.Power) == "ON" || strings.ToUpper(res.StatusSTS.Power1) == "ON", err
		}
	}
	return false, err
}

// CurrentPower implements the api.Meter interface
func (c *Connection) CurrentPower() (float64, error) {
	res, err := c.statusSnsG.Get()
	if err != nil {
		return 0, err
	}
	var power float64 = 0
	for _, channel := range c.channels {
		channelpower, err := res.StatusSNS.Energy.Power.Channel(channel)
		if err != nil {
			return 0, err
		}
		power = power + channelpower
	}
	return power, nil
}

// TotalEnergy implements the api.MeterEnergy interface
func (c *Connection) TotalEnergy() (float64, error) {
	res, err := c.statusSnsG.Get()
	return res.StatusSNS.Energy.Total, err
}

// Currents implements the api.PhaseCurrents interface
func (c *Connection) Currents() (float64, float64, float64, error) {
	res, err := c.statusSnsG.Get()
	if err != nil {
		return 0, 0, 0, err
	}

	var current = [3]float64{0, 0, 0}
	for i := 0; i < len(c.channels) && i < 3; i++ {
		current[i], err = res.StatusSNS.Energy.Current.Channel(c.channels[i])
		if err != nil {
			return 0, 0, 0, err
		}
	}

	return current[0], current[1], current[2], err
}

// SmlPower provides the sml sensor power
func (c *Connection) SmlPower() (float64, error) {
	res, err := c.statusSnsG.Get()
	return float64(res.StatusSNS.SML.PowerCurr), err
}

// SmlTotalEnergy provides the sml sensor total import energy
func (c *Connection) SmlTotalEnergy() (float64, error) {
	res, err := c.statusSnsG.Get()
	return res.StatusSNS.SML.TotalIn, err
}

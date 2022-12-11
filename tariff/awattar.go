package tariff

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/tariff/awattar"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/request"
)

type Awattar struct {
	mu   sync.Mutex
	log  *util.Logger
	uri  string
	data []awattar.PriceInfo
}

var _ api.Tariff = (*Awattar)(nil)

func NewAwattar(other map[string]interface{}) (*Awattar, error) {
	cc := struct {
		Cheap  float64
		Region string
	}{
		Region: "DE",
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	t := &Awattar{
		log: util.NewLogger("awattar"),
		uri: fmt.Sprintf(awattar.RegionURI, strings.ToLower(cc.Region)),
	}

	go t.Run()

	return t, nil
}

func (t *Awattar) Run() {
	client := request.NewHelper(t.log)

	for ; true; <-time.NewTicker(time.Hour).C {
		var res awattar.Prices
		if err := client.GetJSON(t.uri, &res); err != nil {
			t.log.ERROR.Println(err)
			continue
		}

		t.mu.Lock()
		t.data = res.Data
		t.mu.Unlock()
	}
}

// Rates implements the api.Tariff interface
func (t *Awattar) Rates() (api.Rates, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	res := make(api.Rates, 0, len(t.data))
	for _, r := range t.data {
		ar := api.Rate{
			Start: r.StartTimestamp,
			End:   r.EndTimestamp,
			Price: r.Marketprice / 1e3,
		}
		res = append(res, ar)
	}

	return res, nil
}

package tariff

import (
	"context"
	"sync"
	"time"

	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/meter/tibber"
	"github.com/evcc-io/evcc/util"
	"github.com/evcc-io/evcc/util/request"
	"github.com/shurcooL/graphql"
)

type Tibber struct {
	mux    sync.Mutex
	log    *util.Logger
	homeID string
	client *tibber.Client
	data   []tibber.PriceInfo
}

var _ api.Tariff = (*Tibber)(nil)

func NewTibber(other map[string]interface{}) (*Tibber, error) {
	var cc struct {
		Token  string
		HomeID string
		Cheap  any // TODO deprecated
	}

	if err := util.DecodeOther(other, &cc); err != nil {
		return nil, err
	}

	if cc.Token == "" {
		return nil, errors.New("missing token")
	}

	log := util.NewLogger("tibber").Redact(cc.Token, cc.HomeID)

	t := &Tibber{
		log:    log,
		homeID: cc.HomeID,
		client: tibber.NewClient(log, cc.Token),
	}

	if t.homeID == "" {
		var err error
		if t.homeID, err = t.client.DefaultHomeID(); err != nil {
			return nil, err
		}
	}

	// TODO deprecated
	if cc.Cheap != nil {
		t.log.WARN.Println("cheap rate configuration has been replaced by target charging and is deprecated")
	}

	go t.Run()

	return t, nil
}

func (t *Tibber) Run() {
	for ; true; <-time.NewTicker(time.Hour).C {
		var res struct {
			Viewer struct {
				Home struct {
					ID                  string
					TimeZone            string
					CurrentSubscription tibber.Subscription
				} `graphql:"home(id: $id)"`
			}
		}

		v := map[string]interface{}{
			"id": graphql.ID(t.homeID),
		}

		ctx, cancel := context.WithTimeout(context.Background(), request.Timeout)
		err := t.client.Query(ctx, &res, v)
		cancel()

		if err != nil {
			t.log.ERROR.Println(err)
			continue
		}

		t.mux.Lock()
		t.data = res.Viewer.Home.CurrentSubscription.PriceInfo.Today
		t.mux.Unlock()
	}
}

// Rates implements the api.Tariff interface
func (t *Tibber) Rates() (api.Rates, error) {
	t.mux.Lock()
	defer t.mux.Unlock()

	res := make(api.Rates, 0, len(t.data))
	for _, r := range t.data {
		ar := api.Rate{
			Start: r.StartsAt,
			End:   r.StartsAt.Add(time.Hour),
			Price: r.Total,
		}
		res = append(res, ar)
	}

	return res, nil
}

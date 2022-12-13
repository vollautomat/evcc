package planner

import (
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/evcc-io/evcc/api"
	"github.com/evcc-io/evcc/mock"
	"github.com/evcc-io/evcc/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func rates(prices []float64, start time.Time) api.Rates {
	res := make(api.Rates, 0, len(prices))

	for i, v := range prices {
		slotStart := start.Add(time.Duration(i) * time.Hour)
		ar := api.Rate{
			Start: slotStart,
			End:   slotStart.Add(1 * time.Hour),
			Price: v,
		}
		res = append(res, ar)
	}

	return res
}

func TestIsCheapSlotNow(t *testing.T) {
	dt := time.Hour
	ctrl := gomock.NewController(t)

	type se struct {
		caseNr    int
		delay     time.Duration
		cDuration time.Duration
		res       bool
	}

	tc := []struct {
		desc   string
		prices []float64
		end    time.Duration
		series []se
	}{
		{"falling prices", []float64{5, 4, 3, 2, 1, 0, 0, 0}, 5 * time.Hour, []se{
			{1, 1*dt - 1, 20 * time.Minute, false},
			{2, 2*dt - 1, 20 * time.Minute, false},
			{3, 3*dt - 1, 20 * time.Minute, false},
			{4, 3*dt + 1, 20 * time.Minute, false},
			{5, 4*dt - 1, 20 * time.Minute, false},
			{6, 4*dt - 30*time.Minute, 20 * time.Minute, false}, // start as late as possible
			{7, 5*dt - 20*time.Minute, 20 * time.Minute, true},
			{8, 5*dt + 1, 5 * time.Minute, false}, // after desired charge timer,
		}},
		{"raising prices", []float64{1, 2, 3, 4, 5, 6, 7, 8}, 5 * time.Hour, []se{
			{1, 1*dt - 1, time.Hour, true},
			{2, 2*dt - 1, 5 * time.Minute, true}, // charging took longer than expected
			{3, 3*dt - 1, 0, false},
			{4, 5*dt + 1, 0, false}, // after desired charge timer
		}},
		{"last slot", []float64{5, 2, 5, 4, 3, 5, 5, 5}, 5 * time.Hour, []se{
			{1, 1*dt - 1, 70 * time.Minute, false},
			{2, 2*dt - 1, 70 * time.Minute, true},
			{3, 3*dt - 1, 20 * time.Minute, false},
			{4, 4*dt - 1, 20 * time.Minute, false},
			{5, 4*dt + 1, 20 * time.Minute, true}, // start as late as possible
			{6, 4*dt + 40*time.Minute, 20 * time.Minute, true},
		}},
		{"don't stop for last slot", []float64{5, 4, 5, 2, 3, 5, 5, 5}, 5 * time.Hour, []se{
			{1, 1*dt - 1, 70 * time.Minute, false},
			{2, 2*dt - 1, 70 * time.Minute, false},
			{3, 3*dt - 1, 70 * time.Minute, false},
			{4, 4*dt - 1, 20 * time.Minute, true}, // don't pause last slot
			{5, 4*dt + 1, 20 * time.Minute, true},
		}},
		{"delay expensiv middle", []float64{5, 4, 3, 5, 5, 5, 5, 5}, 5 * time.Hour, []se{
			{1, 1*dt - 1, 70 * time.Minute, false},
			{1, 1*dt + 1, 70 * time.Minute, false},
			{2, 2*dt - 1, 61 * time.Minute, true}, // delayed start on expensiv slot
			{3, 3*dt - 1, 60 * time.Minute, true}, // cheapest slot
		}},
		{"disable after known prices, 1h", []float64{5, 4, 3, 2, 1, 0, 0, 0}, 5 * time.Hour, []se{
			{1, 20 * dt, time.Hour, false},
		}},
		{"fixed tariff", []float64{2}, 5 * time.Hour, []se{
			{1, 1, 2 * time.Hour, true},
			{2, 1, 10 * time.Minute, true},
		}},
	}

	clck := clock.NewMock()

	for _, tc := range tc {
		t.Logf("%+v", tc.desc)
		t.Logf("set: %v", clck.Now())

		trf := mock.NewMockTariff(ctrl)
		trf.EXPECT().Rates().AnyTimes().Return(rates(tc.prices, clck.Now()), nil)

		p := &Planner{
			log:    util.NewLogger("foo"),
			clock:  clck,
			tariff: trf,
		}

		start := clck.Now()
		for _, se := range tc.series {
			clck.Set(start.Add(se.delay))

			res, _ := p.Active(se.cDuration, start.Add(tc.end))
			assert.Equalf(t, se.res, res, "%s case %v: expected %v, got %v", tc.desc, se.caseNr, se.res, res)
		}
	}
}

func TestIsCheap(t *testing.T) {
	dt := time.Hour
	ctrl := gomock.NewController(t)

	type se struct {
		caseNr    int
		delay     time.Duration
		cDuration time.Duration
		res       bool
	}

	tc := []struct {
		desc   string
		prices []float64
		end    time.Duration
		series []se
	}{
		{"always expensive", []float64{5, 4, 3, 2, 1, 0, 0, 0}, 5 * time.Hour, []se{
			{1, 1*dt - 1, time.Minute, false},
			{2, 2*dt - 1, time.Minute, false},
			{3, 3*dt - 1, time.Minute, false},
			{4, 4*dt - 1, time.Minute, false},
			{5, 5*dt - 1, time.Minute, true}, // cheapest price
		}},
	}

	clck := clock.NewMock()

	for _, tc := range tc {
		t.Logf("%+v", tc.desc)

		trf := mock.NewMockTariff(ctrl)
		trf.EXPECT().Rates().AnyTimes().Return(rates(tc.prices, clck.Now()), nil)

		p := &Planner{
			log:    util.NewLogger("foo"),
			clock:  clck,
			tariff: trf,
		}

		start := clck.Now()

		for _, se := range tc.series {
			clck.Set(start.Add(se.delay))

			res, _ := p.Active(se.cDuration, start.Add(tc.end))
			assert.Equalf(t, se.res, res, "%s case %v: expected %v, got %v", tc.desc, se.caseNr, se.res, res)
		}
	}
}

func TestNoTariff(t *testing.T) {
	clck := clock.NewMock()

	{
		// check nil
		p := &Planner{
			log:   util.NewLogger("foo"),
			clock: clck,
		}

		res, err := p.Active(time.Hour, clck.Now().Add(30*time.Minute))
		assert.NoError(t, err)
		assert.True(t, res)
	}

	{
		// check dynamic nil
		var trf api.Tariff

		p := &Planner{
			log:    util.NewLogger("foo"),
			clock:  clck,
			tariff: trf,
		}

		res, err := p.Active(time.Hour, clck.Now().Add(30*time.Minute))
		assert.NoError(t, err)
		assert.True(t, res)
	}
}

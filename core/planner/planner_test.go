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

func TestPlanner(t *testing.T) {
	dt := time.Hour
	ctrl := gomock.NewController(t)

	type se struct {
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
			{1*dt - 1, 20 * time.Minute, false},
			{2*dt - 1, 20 * time.Minute, false},
			{3*dt - 1, 20 * time.Minute, false},
			{3*dt + 1, 20 * time.Minute, false},
			{4*dt - 1, 20 * time.Minute, false},
			{4*dt - 30*time.Minute, 20 * time.Minute, false}, // start as late as possible
			{5*dt - 20*time.Minute, 20 * time.Minute, true},
			{5*dt + 1, 5 * time.Minute, false}, // after desired charge timer,
		}},
		{"rising prices", []float64{1, 2, 3, 4, 5, 6, 7, 8}, 5 * time.Hour, []se{
			{1*dt - 1, time.Hour, true},
			{2*dt - 1, 5 * time.Minute, true}, // charging took longer than expected
			{3*dt - 1, 0, false},
			{5*dt + 1, 0, false}, // after desired charge timer
		}},
		{"last slot", []float64{5, 2, 5, 4, 3, 5, 5, 5}, 5 * time.Hour, []se{
			{1*dt - 1, 70 * time.Minute, false},
			{2*dt - 1, 70 * time.Minute, true},
			{3*dt - 1, 20 * time.Minute, false},
			{4*dt - 1, 20 * time.Minute, false},
			{4*dt + 1, 20 * time.Minute, true}, // start as late as possible
			{4*dt + 40*time.Minute, 20 * time.Minute, true},
		}},
		{"don't pause last slot", []float64{5, 4, 5, 2, 3, 5, 5, 5}, 5 * time.Hour, []se{
			{1*dt - 1, 70 * time.Minute, false},
			{2*dt - 1, 70 * time.Minute, false},
			{3*dt - 1, 70 * time.Minute, false},
			{4*dt - 1, 20 * time.Minute, true}, // don't pause last slot
			{4*dt + 1, 20 * time.Minute, true},
		}},
		{"delay expensive middle", []float64{5, 4, 3, 5, 5, 5, 5, 5}, 5 * time.Hour, []se{
			{1*dt - 1, 70 * time.Minute, false},
			{1*dt + 1, 70 * time.Minute, false},
			{2*dt - 1, 61 * time.Minute, true}, // delayed start on expensiv slot
			{3*dt - 1, 60 * time.Minute, true}, // cheapest slot
		}},
		{"disable after known prices, 1h", []float64{5, 4, 3, 2, 1, 0, 0, 0}, 5 * time.Hour, []se{
			{20 * dt, time.Hour, false},
		}},
		{"fixed tariff", []float64{2}, 5 * time.Hour, []se{
			{1, 2 * time.Hour, true},
			{1, 10 * time.Minute, true},
		}},
		{"always expensive", []float64{5, 4, 3, 2, 1, 0, 0, 0}, 5 * time.Hour, []se{
			{1*dt - 1, time.Minute, false},
			{2*dt - 1, time.Minute, false},
			{3*dt - 1, time.Minute, false},
			{4*dt - 1, time.Minute, false},
			{5*dt - 1, time.Minute, true}, // cheapest price
		}},
	}

	clck := clock.NewMock()

	for _, tc := range tc {
		t.Run(tc.desc, func(t *testing.T) {
			t.Logf("set: %v", clck.Now())

			trf := mock.NewMockTariff(ctrl)
			trf.EXPECT().Rates().AnyTimes().Return(rates(tc.prices, clck.Now()), nil)

			p := &Planner{
				log:    util.NewLogger("foo"),
				clock:  clck,
				tariff: trf,
			}

			start := clck.Now()
			for idx, se := range tc.series {
				clck.Set(start.Add(se.delay))

				res, _ := p.Active(se.cDuration, start.Add(tc.end))
				assert.Equalf(t, se.res, res, "%s case %d: expected %v, got %v", tc.desc, idx+1, se.res, res)
			}
		})
	}
}

func TestNoTariff(t *testing.T) {
	clck := clock.NewMock()

	t.Run("nil tariff", func(t *testing.T) {
		p := &Planner{
			log:   util.NewLogger("foo"),
			clock: clck,
		}

		res, err := p.Active(time.Hour, clck.Now().Add(30*time.Minute))
		assert.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("continue even if target date is in the past (without tariff)", func(t *testing.T) {
		p := &Planner{
			log:   util.NewLogger("foo"),
			clock: clck,
		}

		res, err := p.Active(time.Hour, clck.Now().Add(-30*time.Minute))
		assert.NoError(t, err)
		assert.True(t, res)
	})

	t.Run("continue even if target date is in the past (with tariff)", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		trf := mock.NewMockTariff(ctrl)
		trf.EXPECT().Rates().AnyTimes().Return(rates([]float64{0}, clck.Now()), nil)

		p := &Planner{
			log:    util.NewLogger("foo"),
			clock:  clck,
			tariff: trf,
		}

		res, err := p.Active(time.Hour, clck.Now().Add(-30*time.Minute))
		assert.NoError(t, err)
		assert.True(t, res)
	})
}

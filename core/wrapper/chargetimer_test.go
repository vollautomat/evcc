package wrapper

import (
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/stretchr/testify/assert"
)

func TestTimer(t *testing.T) {
	ct := NewChargeTimer()
	clck := clock.NewMock()
	ct.clck = clck

	ct.Reset()
	clck.Add(time.Hour)
	ct.StopCharge()
	clck.Add(time.Hour)

	d, err := ct.ChargingTime()
	assert.NoError(t, err)
	assert.Equal(t, time.Hour, d)

	// continue
	ct.StartCharge()
	clck.Add(2 * time.Hour)
	ct.StopCharge()

	d, err = ct.ChargingTime()
	assert.NoError(t, err)
	assert.Equal(t, 3*time.Hour, d)
}

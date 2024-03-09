package core

import (
	"testing"

	"github.com/evcc-io/evcc/util"
	"github.com/stretchr/testify/assert"
)

type testConsumerVM struct {
	cur    float64
	phases int
}

// interface consumer
func (tc *testConsumerVM) MaxPhasesCurrent() (float64, error) {
	return tc.cur, nil
}

func (tc *testConsumerVM) CurrentPower() (float64, error) {
	return CurrentToPower(tc.cur, uint(tc.phases)), nil
}

// TestVMeterWithConsumers tests with consumers
func TestVMeterWithConsumers(t *testing.T) {
	Voltage = 240
	vm := NewVMeter("test")

	for i := 0; i < 2; i++ {
		cons := new(testConsumerVM)
		cons.phases = 1
		vm.Consumers = append(vm.Consumers, cons)
	}

	var (
		l1, l2, l3, pwr float64
		err             error
	)

	// no LP is consuming
	l1, _, _, err = vm.Currents()
	assert.Nil(t, err)
	assert.Equal(t, l1, 0.0)

	// one lp consumes current
	vm.Consumers[0].(*testConsumerVM).cur = maxA
	l1, l2, l3, _ = vm.Currents()
	assert.Equal(t, l1, maxA)

	// also check all 3 currents are identical
	assert.Equal(t, l1, l2)
	assert.Equal(t, l1, l3)

	// check power
	pwr, _ = vm.CurrentPower()
	assert.Equal(t, CurrentToPower(l1, 1), pwr)

	// 2nd lp consumes current
	vm.Consumers[1].(*testConsumerVM).cur = 6.0
	l1, _, _, err = vm.Currents()
	assert.Nil(t, err)
	assert.Equal(t, l1, maxA+6.0)
}

// TestVMeterWithCircuit tests with circuit as consumer (hierarchy)
func TestVMeterWithCircuit(t *testing.T) {
	Voltage = 240
	vm := NewVMeter("test") // meter under test

	for i := 0; i < 2; i++ {
		cons := &testConsumerVM{
			cur:    maxA,
			phases: 3,
		}
		vm.Consumers = append(vm.Consumers, cons)
	}

	// to remember correct power
	expectedPwr, _ := vm.CurrentPower()

	// subcircuit
	testMeter := testMeter{cur: 10.0}
	circSub, err := NewCircuit(util.NewLogger("foo"), 20.0, 0, nil, &testMeter, &testMeter)
	assert.Nil(t, err)
	assert.NotNilf(t, circSub, "circuit not created")
	subPwr, _ := circSub.CurrentPower()
	expectedPwr = expectedPwr + subPwr
	vm.Consumers = append(vm.Consumers, circSub)

	var (
		l1, l2, l3 float64
	)

	// expect to get the consumers current + circuit current
	l1, l2, l3, err = vm.Currents()
	assert.Nil(t, err)
	assert.Equal(t, l1, maxA*2+10.0)
	// also check all 3 currents are identical
	assert.Equal(t, l1, l2)
	assert.Equal(t, l1, l3)

	// check the power
	pwr, err := vm.CurrentPower()
	assert.Nil(t, err)
	assert.Equal(t, pwr, expectedPwr)
}

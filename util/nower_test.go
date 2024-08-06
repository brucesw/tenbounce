package util

import (
	"testing"
	"time"
)

func Test_TestNower(t *testing.T) {
	var expected = time.Date(1969, time.April, 20, 13, 37, 0, 0, time.UTC)
	var nower = NewTestNower(expected)

	var actual = nower.Now()

	if !actual.Equal(expected) {
		t.Errorf("actual %s, expected %s", actual, expected)
	}
}

func Test_TimeNower(t *testing.T) {
	var acceptableDelta = 100 * time.Millisecond
	var nower = NewTimeNower()
	var actual = nower.Now()

	var expected = time.Now()

	if expected.Sub(actual) > acceptableDelta {
		t.Errorf("actual and expected times differ too much")
	}
}

// Test doesn't do anything other than ensure both implementations satisfy interface
func Test_NowerInterface(t *testing.T) {
	var nowers = []Nower{
		NewTimeNower(),
		NewTestNower(time.Date(1994, time.June, 26, 4, 20, 0, 0, time.UTC)),
	}

	for _, nower := range nowers {
		nower.Now()
	}
}

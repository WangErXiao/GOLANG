package math

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = 1.5
	if v != 1 {
		t.Error("Expected 1,got", v)
	}
}

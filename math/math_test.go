package math_test

import (
	"prospersof/math"
	"testing"
)

// func TestPrimeBetween1To10(t *testing.T) {
// 	given := 10

// 	prime(given)
// }

func TestPowerBase2X4(t *testing.T) {
	givenB := 2
	givenX := 4
	want := 16

	get := math.Power(givenB, givenX)
	if want != get {
		t.Errorf("given b %d and x %d want %d but get %d\n", givenB, givenX, want, get)
	}
}

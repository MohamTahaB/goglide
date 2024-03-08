package vector_test

import (
	"math"
	"testing"

	"github.com/MohamTahaB/goglide/utils/vector"
)

func TestPlus(t *testing.T) {
	u, v := vector.NewVector(1.2, 3.4), vector.NewVector(1.0, 0.6)
	res := vector.NewVector(2.2, 4)
	u.Plus(&v)

	if !vector.Equals(&u, &res) {
		t.Errorf("The sum test failed: the vector is (%f, %f)", u.GetX(), u.GetY())
	}
}

func TestMinus(t *testing.T) {
	u, v := vector.NewVector(1.2, 3.4), vector.NewVector(1.0, 0.6)
	res := vector.NewVector(0.2, 2.8)
	u.Minus(&v)

	if !vector.Equals(&u, &res) {
		t.Errorf("The minus test failed: the vector is (%f, %f)", u.GetX(), u.GetY())
	}
}

func TestTimes(t *testing.T) {
	u := vector.NewVector(1.2, 0.77)
	u.Times(2)

	res := vector.NewVector(2.4, 1.54)
	if !vector.Equals(&u, &res) {
		t.Errorf("The times test failed: the vector is (%f, %f)", u.GetX(), u.GetY())
	}
}

func TestDistance(t *testing.T) {
	u := vector.NewVector(1, 1)
	d := u.Distance()

	if math.Abs(d-math.Sqrt(2)) > 1e-10 {
		t.Errorf("The distance test failed: the norm is %f", d)
	}
}

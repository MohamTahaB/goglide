package cursor_test

import (
	"math/rand"
	"testing"

	"github.com/MohamTahaB/goglide/cursor"
)

func TestLimitAcceleration(t *testing.T) {
	magnitude := rand.Float64() * 100
	c := cursor.RandomCursor(600, 600)

	acc := c.GetAcceleration()
	acc.Times(2 * magnitude)

	c.SetAcceleration(&acc)
	c.LimitAcceleration(magnitude)

	if acc = c.GetAcceleration(); acc.Distance() > magnitude+1e-10 {
		t.Errorf("limit acceleration test failed: expected a max magnitude of %f, got %f", magnitude, acc.Distance())
	}
}

func TestLimitVelocity(t *testing.T) {
	magnitude := rand.Float64() * 100
	c := cursor.RandomCursor(600, 600)

	v := c.GetVelocity()
	v.Times(2 * magnitude)

	c.SetVelocity(&v)
	c.LimitVelocity(magnitude)

	if v = c.GetVelocity(); v.Distance() > magnitude+1e-10 {
		t.Errorf("limit velocity test failed: expected a max magnitude of %f, got %f", magnitude, v.Distance())
	}
}

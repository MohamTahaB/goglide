package cursor_test

import (
	"testing"

	"github.com/MohamTahaB/goglide/cursor"
)

func TestLimitAcceleration(t *testing.T) {
	c := cursor.RandomCursor(600, 600)
	magnitude := c.GetMaxAcceleration()

	acc := c.GetAcceleration()
	acc.Times(2 * magnitude)

	c.SetAcceleration(&acc)
	c.LimitAcceleration()

	if acc = c.GetAcceleration(); acc.Distance() > magnitude+1e-10 {
		t.Errorf("limit acceleration test failed: expected a max magnitude of %f, got %f", magnitude, acc.Distance())
	}
}

func TestLimitVelocity(t *testing.T) {
	c := cursor.RandomCursor(600, 600)
	magnitude := c.GetMaxVelocity()

	v := c.GetVelocity()
	v.Times(2 * magnitude)

	c.SetVelocity(&v)
	c.LimitVelocity()

	if v = c.GetVelocity(); v.Distance() > magnitude+1e-10 {
		t.Errorf("limit velocity test failed: expected a max magnitude of %f, got %f", magnitude, v.Distance())
	}
}

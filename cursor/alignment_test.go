package cursor_test

import (
	"testing"

	"github.com/MohamTahaB/goglide/cursor"
)

func TestAlign(t *testing.T) {
	var boids []*cursor.Cursor
	for i := 0; i < 100; i++ {
		boids = append(boids, cursor.RandomCursor(100, 100))
	}

	steer := boids[0].Align(140, &boids)

	if steer.Distance() < 1e-6 {
		t.Errorf("align test failed: expected a non null steering vector, got (%f, %f)", steer.GetX(), steer.GetY())
	}
}

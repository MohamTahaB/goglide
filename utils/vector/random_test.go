package vector_test

import (
	"testing"

	"github.com/MohamTahaB/goglide/utils/vector"
)

func TestRandomPos(t *testing.T) {
	// Define tests table.
	tests := []struct {
		name string
		args struct {
			w int
			h int
		}
	}{
		{
			name: "null screen size",
			args: struct {
				w int
				h int
			}{0, 0},
		},
		{
			name: "fullscreen",
			args: struct {
				w int
				h int
			}{1920, 1080},
		},
	}

	for _, test := range tests {
		pos := vector.RandomPos(test.args.w, test.args.h)
		if pos.GetX() > float64(test.args.w) || pos.GetY() > float64(test.args.h) {
			t.Errorf("%s: error generating the random position vector: (%f, %f)", test.name, pos.GetX(), pos.GetY())
		}
	}
}

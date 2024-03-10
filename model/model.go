// Package model defines the model struct of the boids model.
package model

import (
	"github.com/MohamTahaB/goglide/cursor"
)

type Model struct {
	boids []*cursor.Cursor
}

// NewModel initiates a new model with given number of boids.
func NewModel(number, w, h int) *Model {
	var boids []*cursor.Cursor
	for i := 0; i < number; i++ {
		boids = append(boids, cursor.RandomCursor(w, h))
	}

	return &Model{
		boids: boids,
	}
}

func (m *Model) GetBoids() *[]*cursor.Cursor {
	if m == nil {
		return nil
	}
	return &(m.boids)
}

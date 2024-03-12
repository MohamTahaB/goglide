// Package model defines the model struct of the boids model.
package model

import (
	"github.com/MohamTahaB/goglide/cursor"
)

type Model struct {
	radius float64
	boids  []*cursor.Cursor
}

// NewModel initiates a new model with given number of boids.
func NewModel(number, w, h int, protectedRange, visualRange float64) *Model {
	var boids []*cursor.Cursor
	for i := 0; i < number; i++ {
		boids = append(boids, cursor.RandomCursor(w, h))
	}

	return &Model{
		radius: visualRange,
		boids:  boids,
	}
}

func (m *Model) GetBoids() *[]*cursor.Cursor {
	if m == nil {
		return nil
	}
	return &(m.boids)
}

func (m *Model) Update(deltaT float64, w, h int) {
	var newBoids []*cursor.Cursor
	for _, boid := range m.boids {
		newBoids = append(newBoids, boid.Update(deltaT, m.radius, &m.boids, w, h))
	}

	m.boids = newBoids
}

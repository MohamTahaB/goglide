package model_test

import (
	"math/rand"
	"testing"

	"github.com/MohamTahaB/goglide/model"
)

func TestNewModel(t *testing.T) {

	number := rand.Intn(20) + 10
	m := model.NewModel(number, 1920, 1080, 20)
	boids := m.GetBoids()

	if boids == nil {
		t.Error("new model test failed: boids slice is nil")
	} else if len(*boids) != number {
		t.Error("new model test failed: incorrect number of boids")
	}
}

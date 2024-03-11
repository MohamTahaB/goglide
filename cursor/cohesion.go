package cursor

import "github.com/MohamTahaB/goglide/utils/vector"

func (c Cursor) Cohesion(radius float64, boids *[]*Cursor) vector.Vector {
	cohesionSteer := vector.NewVector(0, 0)

	t := 0
	v := c.position
	for _, boid := range *boids {
		v.Minus(&boid.position)
		if v.Distance() <= radius && v.Distance() != 0 {
			cohesionSteer.Plus(&boid.position)
			t++
		}
		v.Plus(&boid.position)
	}

	if t > 0 {
		cohesionSteer.Times(1.0 / float64(t))
	}

	cohesionSteer.Minus(&c.position)
	return cohesionSteer
}

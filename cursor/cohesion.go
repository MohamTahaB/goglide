package cursor

import "github.com/MohamTahaB/goglide/utils/vector"

func (c Cursor) Cohesion(radius float64, boids *[]*Cursor) vector.Vector {
	cohesionSteer := vector.NewVector(0, 0)

	v := c.position
	for _, boid := range *boids {
		v.Minus(&boid.position)
		if v.Distance() <= radius && v.Distance() != 0 {
			cohesionSteer.Plus(&boid.position)
		}
		v.Plus(&boid.position)
	}

	cohesionSteer.Minus(&c.position)
	return cohesionSteer
}

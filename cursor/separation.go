package cursor

import "github.com/MohamTahaB/goglide/utils/vector"

func (c Cursor) Separate(radius float64, boids *[]*Cursor) vector.Vector {
	separationSteer := vector.NewVector(0, 0)

	v := c.position
	for _, boid := range *boids {
		v.Minus(&boid.position)
		if v.Distance() <= radius && v.Distance() != 0 {
			aux := v
			aux.Times(1.0 / v.Distance())
			separationSteer.Plus(&boid.velocity)
		}
		v.Plus(&boid.position)
	}

	return separationSteer
}

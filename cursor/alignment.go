package cursor

import "github.com/MohamTahaB/goglide/utils/vector"

// Align is a cursor method that takes as parameter a perseption radius of the boid cursor, and a pointer to the slice of boids pointers.
// It returns a steering vector that represents a steering force to the average movement of the flock within the radius from the considered boid cursor.
func (c Cursor) Align(radius float64, boids *[]*Cursor) vector.Vector {
	alignmentSteer := vector.NewVector(0, 0)

	t := 0
	v := c.position
	for _, boid := range *boids {
		v.Minus(&boid.position)
		if v.Distance() <= radius && v.Distance() != 0 {
			alignmentSteer.Plus(&boid.velocity)
			t++
		}
		v.Plus(&boid.position)
	}

	if t > 0 {
		alignmentSteer.Times(1.0 / float64(t))
	}

	alignmentSteer.Minus(&c.acceleration)
	return alignmentSteer
}

// Package cursor provides the cursor struct and its handling methods.
package cursor

import "github.com/MohamTahaB/goglide/utils/vector"

type Cursor struct {
	position     vector.Vector
	velocity     vector.Vector
	acceleration vector.Vector
}

// RandomCursor returns a cursor with random position, velocity and acceleration.
func RandomCursor(w, h int) *Cursor {
	return &Cursor{
		position: vector.RandomPos(w, h),
		velocity: vector.RandomVelocity(),
		acceleration: vector.RandomAcceleration(),
	}
}

func (c *Cursor) GetPosition() vector.Vector {
	return c.position
}

func (c *Cursor) GetVelocity() vector.Vector {
	return c.velocity
}

func (c *Cursor) GetAcceleration() vector.Vector {
	return c.acceleration
}

func (c Cursor) Update(deltaT, radius float64, boids *[]*Cursor) *Cursor {
	steer := c.Align(radius, boids)

	c.acceleration.Plus(&steer)

	velocityIncrement := c.acceleration
	velocityIncrement.Times(deltaT)
	c.velocity.Plus(&velocityIncrement)

	posIncrement := c.velocity
	posIncrement.Times(deltaT)
	c.position.Plus(&posIncrement)

	return &c
}

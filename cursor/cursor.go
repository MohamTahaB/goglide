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
		position:     vector.RandomPos(w, h),
		velocity:     vector.RandomVelocity(),
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

func (c *Cursor) SetPosition(pos *vector.Vector) {
	c.position = *pos
}

func (c *Cursor) SetVelocity(v *vector.Vector) {
	c.velocity = *v
}

func (c *Cursor) SetAcceleration(acc *vector.Vector) {
	c.acceleration = *acc
}

func (c Cursor) Update(deltaT, radius float64, boids *[]*Cursor, w, h int) *Cursor {
	steer := c.Align(radius, boids)

	c.acceleration.Times(0.5)
	c.acceleration.Plus(&steer)

	velocityIncrement := c.acceleration
	velocityIncrement.Times(deltaT)
	c.velocity.Plus(&velocityIncrement)

	posIncrement := c.velocity
	posIncrement.Times(deltaT)
	c.position.Plus(&posIncrement)

	if c.position.GetX() > float64(w) {
		c.position.SetX(c.position.GetX() - float64(w))
	}
	if c.position.GetX() < 0 {
		c.position.SetX(c.position.GetX() + float64(w))
	}

	if c.position.GetY() > float64(h) {
		c.position.SetY(c.position.GetY() - float64(h))
	}
	if c.position.GetY() < 0 {
		c.position.SetY(c.position.GetY() + float64(h))
	}

	return &c
}

// LimitAcceleration limits the magnitude of the acceleration vector
func (c *Cursor) LimitAcceleration(magnitude float64) {
	if c.acceleration.Distance() > magnitude {
		c.acceleration.Times(magnitude / c.acceleration.Distance())
	}
}

// LimitVelocity limits the magnitude of the velocity vector
func (c *Cursor) LimitVelocity(magnitude float64) {
	if c.velocity.Distance() > magnitude {
		c.velocity.Times(magnitude / c.velocity.Distance())
	}
}

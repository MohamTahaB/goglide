// Package cursor provides the cursor struct and its handling methods.
package cursor

import (
	"math/rand"

	"github.com/MohamTahaB/goglide/utils/vector"
)

type Cursor struct {
	position                     vector.Vector
	velocity                     vector.Vector
	acceleration                 vector.Vector
	maxAcceleration, maxVelocity float64
}

// RandomCursor returns a cursor with random position, velocity and acceleration.
func RandomCursor(w, h int) *Cursor {
	return &Cursor{
		position:        vector.RandomPos(w, h),
		velocity:        vector.RandomVelocity(),
		acceleration:    vector.RandomAcceleration(),
		maxVelocity:     150 + rand.Float64()*50,
		maxAcceleration: 50 + rand.Float64()*150,
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

func (c *Cursor) GetMaxAcceleration() float64 {
	return c.maxAcceleration
}

func (c *Cursor) GetMaxVelocity() float64 {
	return c.maxVelocity
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
	// Compute steers
	alignmentSteer := c.Align(radius, boids)
	// TODO
	//cohesionSteer := c.Cohesion(radius, boids)

	c.acceleration.Plus(&alignmentSteer)
	// TODO
	//c.acceleration.Plus(&cohesionSteer)

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

	// Limit both acceleration and velocity magnitudes.
	c.LimitAcceleration()
	c.LimitVelocity()

	return &c
}

// LimitAcceleration limits the magnitude of the acceleration vector
func (c *Cursor) LimitAcceleration() {
	if c.acceleration.Distance() > c.maxAcceleration {
		c.acceleration.Times(c.maxAcceleration / c.acceleration.Distance())
	}
}

// LimitVelocity limits the magnitude of the velocity vector
func (c *Cursor) LimitVelocity() {
	if c.velocity.Distance() > c.maxVelocity {
		c.velocity.Times(c.maxVelocity / c.velocity.Distance())
	}
}

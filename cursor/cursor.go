// Package cursor provides the cursor struct and its handling methods.
package cursor

import "github.com/MohamTahaB/goglide/utils/vector"

type Cursor struct {
	position     vector.Vector
	velocity     vector.Vector
	acceleration vector.Vector
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

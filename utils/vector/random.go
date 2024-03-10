package vector

import "math/rand"

// RandomPos returns a random position vector, given the width and height of the window.
func RandomPos(w, h int) Vector {
	var x, y float64
	x = rand.Float64() * float64(w)
	y = rand.Float64() * float64(h)

	return NewVector(x, y)
}

// RandomVelocity returns a random velocity vector.
func RandomVelocity() Vector {
	x, y := rand.Float64(), rand.Float64()

	v := NewVector(x, y)
	v.Times(rand.Float64() * 200.0)

	return v
}

// RandomAcceleration returns a random acceleration vector.
func RandomAcceleration() Vector {
	return NewVector(rand.Float64(), rand.Float64())
}

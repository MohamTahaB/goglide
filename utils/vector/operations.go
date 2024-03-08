package vector

import "math"

// Plus adds the vector passed as parameter to the handled vector.
func (v *Vector) Plus(u *Vector) {
	v.x += u.x
	v.y += u.y
}

// Minus substracts the vector passed as parameter from the handled vector.
func (v *Vector) Minus(u *Vector) {
	v.x -= u.x
	v.y -= u.y
}

// Times multiplies the handled vector by the float passed as parameter.
func (v *Vector) Times(c float64) {
	v.x *= c
	v.y *= c
}

// Distance returns the norm of a given vector.
func (v *Vector) Distance() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// Equals checks if two vectors are within e-10 of each other.
func Equals(u, v *Vector) bool {
	return math.Pow(u.x-v.x, 2)+math.Pow(u.y-v.y, 2) < 1e-10
}

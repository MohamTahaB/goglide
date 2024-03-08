// Package vector provides the basics of vectors and their usual operations.
package vector

// Vector struct.
type Vector struct {
	x float64
	y float64
}

// NewVector creates a new vector with the corresponding coordinates.
func NewVector(x, y float64) Vector {
	return Vector{
		x: x,
		y: y,
	}
}

// GetX is a getter to the x value in the vector.
func (v *Vector) GetX() float64 {
	return v.x
}

// SetX is a setter to the x value in the vector.
func (v *Vector) SetX(x float64) {
	v.x = x
}

// GetY is a getter to the y value in the vector.
func (v *Vector) GetY() float64 {
	return v.y
}

// SetY is a setter to the y value in the vector.
func (v *Vector) SetY(y float64) {
	v.y = y
}

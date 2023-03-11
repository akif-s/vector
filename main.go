package vector

import (
	"math"
)

type Vector2 struct {
	X float64
	Y float64
}

func NewVector2(X, Y float64) Vector2 {
	return Vector2{
		X: X,
		Y: Y,
	}
}

// Calculates the dot product of two vectors
func Dot(v1 Vector2, v2 Vector2) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

// Calculates the angle between two vectors
func Angle(v1 Vector2, v2 Vector2) float64 {
	return math.Acos(Dot(v1, v2) / (v1.Magnitude() * v2.Magnitude()))
}

// Multiplies the each element of vector2 with a
func (v Vector2) Product(a float64) Vector2 {
	return NewVector2(v.X*a, v.Y*a)
}

func (v Vector2) Divide(a float64) Vector2 {
	return NewVector2(v.X/a, v.Y/a)
}

// Adds two vector
func (v1 Vector2) Sum(v2 Vector2) Vector2 {
	return NewVector2(v1.X+v2.X, v1.Y+v2.Y)
}

// Substracts v2 from v1
func (v1 Vector2) Substract(v2 Vector2) Vector2 {
	return NewVector2(v1.X-v2.X, v1.Y-v2.Y)
}

// Calculates the magnitude of the vector
func (v Vector2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Normalizes the given vector(returns it as a unit vector)
func (v Vector2) Normalize() Vector2 {
	return v.Divide(v.Magnitude())
}

// Returns the absolute of the vector
func (v Vector2) Abs() Vector2 {
	return NewVector2(math.Abs(v.X), math.Abs(v.Y))
}

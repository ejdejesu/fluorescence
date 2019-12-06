package geometry

import "math"

// Vector is a 3D vector
type Vector struct {
	X float64
	Y float64
	Z float64
}

var UP = &Vector{0.0, 1.0, 0.0}
var RIGHT = &Vector{1.0, 0.0, 0.0}
var FORWARD = &Vector{0.0, 0.0, -1.0}

// Magnitude return euclidean length of vector
func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Unit returns a new vector with direction preserved and length equal to one
func (v *Vector) Unit() *Vector {
	return v.DivideFloat64(v.Magnitude())
}

func (v *Vector) Dot(w *Vector) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

func (v *Vector) Cross(w *Vector) *Vector {
	return &Vector{v.Y*w.Z - v.Z*w.Y, v.Z*w.X - v.X*w.Z, v.X*w.Y - v.Y*w.X}
}

func (v *Vector) Add(w *Vector) *Vector {
	return &Vector{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

func (v *Vector) Subtract(w *Vector) *Vector {
	return &Vector{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

func (v *Vector) MultiplyFloat64(s float64) *Vector {
	return &Vector{v.X * s, v.Y * s, v.Z * s}
}

func (v *Vector) MultiplyVector(w *Vector) *Vector {
	return &Vector{v.X * w.X, v.Y * w.Y, v.Z * w.Z}
}

func (v *Vector) DivideFloat64(s float64) *Vector {
	return &Vector{v.X / s, v.Y / s, v.Z / s}
}

func (v *Vector) DivideVector(w *Vector) *Vector {
	return &Vector{v.X / w.X, v.Y / w.Y, v.Z / w.Z}
}

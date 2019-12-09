package geometry

import (
	"fluorescence/shading"
	"math"
	"math/rand"
)

// Vector is a 3D vector
type Vector struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

var ZERO = &Vector{0.0, 0.0, 0.0}

var UP = &Vector{0.0, 1.0, 0.0}
var RIGHT = &Vector{1.0, 0.0, 0.0}
var FORWARD = &Vector{0.0, 0.0, -1.0}

func RandomOnUnitDisc(rng *rand.Rand) *Vector {
	for {
		v := &Vector{
			X: 2.0*rng.Float64() - 1.0,
			Y: 2.0*rng.Float64() - 1.0,
			Z: 0.0,
		}
		if v.Magnitude() < 1.0 {
			return v
		}
	}
}

func RandomInUnitSphere(rng *rand.Rand) *Vector {
	for {
		v := &Vector{
			X: 2.0*rng.Float64() - 1.0,
			Y: 2.0*rng.Float64() - 1.0,
			Z: 2.0*rng.Float64() - 1.0,
		}
		if v.Magnitude() < 1.0 {
			return v
		}
	}
}

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

func (v *Vector) AddInPlace(w *Vector) {
	v.X += w.X
	v.Y += w.Y
	v.Z += w.Z
}

func (v *Vector) Subtract(w *Vector) *Vector {
	return &Vector{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

func (v *Vector) MultiplyFloat64(s float64) *Vector {
	return &Vector{v.X * s, v.Y * s, v.Z * s}
}

func (v *Vector) Pow(e float64) *Vector {
	return &Vector{math.Pow(v.X, e), math.Pow(v.Y, e), math.Pow(v.Z, e)}
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

func (v *Vector) ToColor() *shading.Color {
	return &shading.Color{v.X, v.Y, v.Z, 1.0}
}

func (v *Vector) Copy() *Vector {
	return &Vector{v.X, v.Y, v.Z}
}

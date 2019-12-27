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

// VectorZero references the zero vector
var VectorZero = Vector{}

// VectorMax references the maximum representable float64 vector
var VectorMax = Vector{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64}

// VectorUp references the up vector (positive Y) with the standard cartesian axes as an orthogonal system
var VectorUp = Vector{0.0, 1.0, 0.0}

// VectorRight references the right vector (positive X) with the standard cartesian axes as an orthogonal system
var VectorRight = Vector{1.0, 0.0, 0.0}

// VectorForward references the forward vector (negative Z) with the standard cartesian axes as an orthogonal system
// it points towards negative Z to preserve the system's right-handedness
var VectorForward = Vector{0.0, 0.0, -1.0}

// RandomOnUnitDisk returns a new Vector pointing from the origin to a
// random point on a unit disk
func RandomOnUnitDisk(rng *rand.Rand) Vector {
	for {
		v := Vector{
			X: 2.0*rng.Float64() - 1.0,
			Y: 2.0*rng.Float64() - 1.0,
			Z: 0.0,
		}
		if v.Magnitude() < 1.0 {
			return v
		}
	}
}

// RandomInUnitSphere returns a new Vector pointing from the origin to a
// random point in a unit sphere
func RandomInUnitSphere(rng *rand.Rand) Vector {
	for {
		v := Vector{
			X: 2.0*rng.Float64() - 1.0,
			Y: 2.0*rng.Float64() - 1.0,
			Z: 2.0*rng.Float64() - 1.0,
		}
		if v.Magnitude() < 1.0 {
			return v
		}
	}
}

// Magnitude return euclidean length of Vector
func (v Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Unit returns a new Vector with direction preserved and length equal to one
func (v Vector) Unit() Vector {
	return v.DivScalar(v.Magnitude())
}

// Dot computes the dot or scalar product of two Vectors
func (v Vector) Dot(w Vector) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

// Cross computes the cross or Vector product of two Vectors
func (v Vector) Cross(w Vector) Vector {
	return Vector{v.Y*w.Z - v.Z*w.Y, v.Z*w.X - v.X*w.Z, v.X*w.Y - v.Y*w.X}
}

// Add adds a Vector to another Vector component-wise
func (v Vector) Add(w Vector) Vector {
	return Vector{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

// Sub subtracts a Vector from another Vector component-wise
func (v Vector) Sub(w Vector) Vector {
	return Vector{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

// MultScalar multiplies a Vector by a scalar
func (v Vector) MultScalar(s float64) Vector {
	return Vector{v.X * s, v.Y * s, v.Z * s}
}

// MultVector multiplies a Vector by a Vector component-wise
func (v Vector) MultVector(w Vector) Vector {
	return Vector{v.X * w.X, v.Y * w.Y, v.Z * w.Z}
}

// Pow raises a Vector to an exponential power, component-wise
func (v Vector) Pow(e float64) Vector {
	return Vector{math.Pow(v.X, e), math.Pow(v.Y, e), math.Pow(v.Z, e)}
}

// DivScalar divides a Vector by a scalar
func (v Vector) DivScalar(s float64) Vector {
	inv := 1.0 / s
	return Vector{v.X * inv, v.Y * inv, v.Z * inv}
}

// DivVector divides a Vector by a Vector component-wise
func (v Vector) DivVector(w Vector) Vector {
	return Vector{v.X / w.X, v.Y / w.Y, v.Z / w.Z}
}

// Negate returns a Vector pointing in the opposite direction
func (v Vector) Negate() Vector {
	return Vector{-v.X, -v.Y, -v.Z}
}

// ReflectAround returns the reflection of a vector given a normal
func (v Vector) ReflectAround(w Vector) Vector {
	return v.Sub(w.MultScalar(v.Dot(w) * 2.0))
}

// RefractAround returns the refraction of a vector given the normal and ratio of reflective indices
func (v Vector) RefractAround(w Vector, rri float64) (Vector, bool) {
	dt := v.Unit().Dot(w)
	discriminant := 1.0 - (rri*rri)*(1.0-(dt*dt))
	// fmt.Println(rri)
	if discriminant > 0 {
		// fmt.Println("yu")
		return v.Unit().Sub(w.MultScalar(dt)).MultScalar(rri).Sub(w.MultScalar(math.Sqrt(discriminant))), true
	}
	return VectorZero, false
}

// ToColor converts a Vector to a Color
func (v Vector) ToColor() shading.Color {
	return shading.Color{
		Red:   v.X,
		Green: v.Y,
		Blue:  v.Z,
	}
}

// VectorFromColor creates a Vector from a Color
func VectorFromColor(c shading.Color) Vector {
	return Vector{c.Red, c.Green, c.Blue}
}

// Copy returns a new Vector identical to v
func (v Vector) Copy() Vector {
	return Vector{v.X, v.Y, v.Z}
}

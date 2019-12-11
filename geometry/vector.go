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
	return v.DivScalar(v.Magnitude())
}

func (v *Vector) UnitInPlace() *Vector {
	return v.DivScalarInPlace(v.Magnitude())
}

func (v *Vector) Clamp(min, max float64) *Vector {
	return &Vector{
		clamp(v.X, min, max),
		clamp(v.Y, min, max),
		clamp(v.Z, min, max)}
}

func (v *Vector) ClampInPlace(min, max float64) *Vector {
	v.X = clamp(v.X, min, max)
	v.Y = clamp(v.Y, min, max)
	v.Z = clamp(v.Z, min, max)
	return v
}

func clamp(val, min, max float64) float64 {
	if val < min {
		return min
	} else if val > max {
		return max
	}
	return val
}

func (v *Vector) Dot(w *Vector) float64 {
	return v.X*w.X + v.Y*w.Y + v.Z*w.Z
}

func (v *Vector) Cross(w *Vector) *Vector {
	return &Vector{v.Y*w.Z - v.Z*w.Y, v.Z*w.X - v.X*w.Z, v.X*w.Y - v.Y*w.X}
}

func (v *Vector) CrossInPlace(w *Vector) *Vector {

	nX := v.Y*w.Z - v.Z*w.Y
	nY := v.Z*w.X - v.X*w.Z
	nZ := v.X*w.Y - v.Y*w.X

	v.X = nX
	v.Y = nY
	v.Z = nZ

	return v
}

func (v *Vector) Add(w *Vector) *Vector {
	return &Vector{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

func (v *Vector) AddInPlace(w *Vector) *Vector {
	v.X += w.X
	v.Y += w.Y
	v.Z += w.Z
	return v
}

func (v *Vector) Sub(w *Vector) *Vector {
	return &Vector{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

func (v *Vector) SubInPlace(w *Vector) *Vector {
	v.X -= w.X
	v.Y -= w.Y
	v.Z -= w.Z
	return v
}

func (v *Vector) MultScalar(s float64) *Vector {
	return &Vector{v.X * s, v.Y * s, v.Z * s}
}

func (v *Vector) MultScalarInPlace(s float64) *Vector {
	v.X *= s
	v.Y *= s
	v.Z *= s
	return v
}

func (v *Vector) MultVector(w *Vector) *Vector {
	return &Vector{v.X * w.X, v.Y * w.Y, v.Z * w.Z}
}

func (v *Vector) MultVectorInPlace(w *Vector) *Vector {
	v.X *= w.X
	v.Y *= w.Y
	v.Z *= w.Z
	return v
}

func (v *Vector) Pow(e float64) *Vector {
	return &Vector{math.Pow(v.X, e), math.Pow(v.Y, e), math.Pow(v.Z, e)}
}

func (v *Vector) PowInPlace(e float64) *Vector {
	v.X = math.Pow(v.X, e)
	v.Y = math.Pow(v.Y, e)
	v.Z = math.Pow(v.Z, e)
	return v
}

func (v *Vector) DivScalar(s float64) *Vector {
	return &Vector{v.X / s, v.Y / s, v.Z / s}
}

func (v *Vector) DivScalarInPlace(s float64) *Vector {
	v.X /= s
	v.Y /= s
	v.Z /= s
	return v
}

func (v *Vector) DivVector(w *Vector) *Vector {
	return &Vector{v.X / w.X, v.Y / w.Y, v.Z / w.Z}
}

func (v *Vector) DivVectorInPlace(w *Vector) *Vector {
	v.X /= w.X
	v.Y /= w.Y
	v.Z /= w.Z
	return v
}

func (v *Vector) ReflectAround(w *Vector) *Vector {
	return v.Sub(w.MultScalar(v.Dot(w) * 2.0))
}

func (v *Vector) RefractAround(w *Vector, rri float64) (*Vector, bool) {
	dt := v.Unit().Dot(w)
	discriminant := 1.0 - (rri*rri)*(1.0-(dt*dt))
	// fmt.Println(rri)
	if discriminant > 0 {
		// fmt.Println("yu")
		return v.Unit().Sub(w.MultScalar(dt)).MultScalar(rri).Sub(w.MultScalar(math.Sqrt(discriminant))), true
	}
	return nil, false
}

func (v *Vector) ToColor() *shading.Color {
	return &shading.Color{v.X, v.Y, v.Z, 1.0}
}

func (v *Vector) Copy() *Vector {
	return &Vector{v.X, v.Y, v.Z}
}

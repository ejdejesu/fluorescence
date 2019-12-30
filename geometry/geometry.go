package geometry

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/go-gl/mathgl/mgl64"
)

// Vec3Zero references zero vector
var Vec3Zero = mgl64.Vec3{0.0, 0.0, 0.0}

// Vec3Max references the maximum representable float64s in a vector
var Vec3Max = mgl64.Vec3{mgl64.MaxValue, mgl64.MaxValue, mgl64.MaxValue}

// Vec3Up references the up vector (positive Y) with the standard cartesian axes as an orthogonal system
var Vec3Up = mgl64.Vec3{0.0, 1.0, 0.0}

// Vec3Right references the right vector (positive X) with the standard cartesian axes as an orthogonal system
var Vec3Right = mgl64.Vec3{1.0, 0.0, 0.0}

// Vec3Forward references the forward vector (negative Z) with the standard cartesian axes as an orthogonal system
// it points towards negative Z to preserve the system's right-handedness
var Vec3Forward = mgl64.Vec3{0.0, 0.0, -1.0}

// RandomOnUnitDisk returns a new Vector pointing from the origin to a
// random point on a unit disk
func RandomOnUnitDisk(rng *rand.Rand) mgl64.Vec3 {
	for {
		v := mgl64.Vec3{
			2.0*rng.Float64() - 1.0,
			2.0*rng.Float64() - 1.0,
			0.0,
		}
		if v.Len() < 1.0 {
			return v
		}
	}
}

// RandomInUnitSphere returns a new Vector pointing from the origin to a
// random point in a unit sphere
func RandomInUnitSphere(rng *rand.Rand) mgl64.Vec3 {
	for {
		v := mgl64.Vec3{
			2.0*rng.Float64() - 1.0,
			2.0*rng.Float64() - 1.0,
			2.0*rng.Float64() - 1.0,
		}
		if v.Len() < 1.0 {
			return v
		}
	}
}

// MinComponents returns the Point construction of the minimums of two points component-wise
func MinComponents(p, q mgl64.Vec3) mgl64.Vec3 {
	return mgl64.Vec3{math.Min(p.X(), q.X()), math.Min(p.Y(), q.Y()), math.Min(p.Z(), q.Z())}
}

// MaxComponents returns the Point construction of the maximums of two points component-wise
func MaxComponents(p, q mgl64.Vec3) mgl64.Vec3 {
	return mgl64.Vec3{math.Max(p.X(), q.X()), math.Max(p.Y(), q.Y()), math.Max(p.Z(), q.Z())}
}

// MulVec3 returns the product of two Vec3 component-wise
func MulVec3(v, w mgl64.Vec3) mgl64.Vec3 {
	return mgl64.Vec3{v[0] * w[0], v[1] * w[1], v[2] * w[2]}
}

// PowVec3 raises a Vector to an exponential power, component-wise
func PowVec3(v mgl64.Vec3, e float64) mgl64.Vec3 {
	return mgl64.Vec3{
		math.Pow(v.X(), e),
		math.Pow(v.Y(), e),
		math.Pow(v.Z(), e),
	}
}

// NegateVec3 returns a Vector pointing in the opposite direction
func NegateVec3(v mgl64.Vec3) mgl64.Vec3 {
	return v.Mul(-1.0)
}

// ClampVec3 clamps elements of a Vec3
func ClampVec3(v mgl64.Vec3, min, max float64) mgl64.Vec3 {
	return mgl64.Vec3{
		mgl64.Clamp(v[0], min, max),
		mgl64.Clamp(v[1], min, max),
		mgl64.Clamp(v[2], min, max),
	}
}

// ScaleVec3 scales all elements equally so the max channel is s
func ScaleVec3(v mgl64.Vec3, s float64) mgl64.Vec3 {
	max := math.Max(v[0], math.Max(v[1], v[2]))
	return v.Mul(s / max)
}

// ScaleDownVec3 scales as Scale does, but only if the max channel exceeds s
func ScaleDownVec3(v mgl64.Vec3, s float64) mgl64.Vec3 {
	max := math.Max(v[0], math.Max(v[1], v[2]))
	if max > s {
		return v.Mul(s / max)
	}
	return v
}

// ScaleUpVec3 scales as Scale does, but only if the highest channel falls below s
func ScaleUpVec3(v mgl64.Vec3, s float64) mgl64.Vec3 {
	max := math.Max(v[0], math.Max(v[1], v[2]))
	if max < s {
		return v.Mul(s / max)
	}
	return v
}

// ReflectAroundVec3 returns the reflection of a vector given a normal
func ReflectAroundVec3(v, w mgl64.Vec3) mgl64.Vec3 {
	return v.Sub(w.Mul(v.Dot(w) * 2.0))
}

// RefractAroundVec3 returns the refraction of a vector given the normal and ratio of reflective indices
func RefractAroundVec3(v, w mgl64.Vec3, rri float64) (mgl64.Vec3, bool) {
	dt := v.Normalize().Dot(w)
	discriminant := 1.0 - (rri*rri)*(1.0-(dt*dt))
	// fmt.Println(rri)
	if discriminant > 0 {
		// fmt.Println("yu")
		return v.Normalize().Sub(w.Mul(dt)).Mul(rri).Sub(w.Mul(math.Sqrt(discriminant))), true
	}
	return mgl64.Vec3{}, false
}

// ColorToVec3 converts a color to a color.Color
func ColorToVec3(c color.Color) mgl64.Vec3 {
	// fmt.Println(c)
	r, g, b, _ := c.RGBA()
	// fmt.Println(r, g, b)
	inv := float64(1.0 / math.MaxUint16)
	// fmt.Println("red   ", float64(r)*inv)
	// fmt.Println("green ", float64(g)*inv)
	// fmt.Println("blue  ", float64(b)*inv)
	return mgl64.Vec3{
		float64(r) * inv,
		float64(g) * inv,
		float64(b) * inv,
	}
}

// Vec3ToRGBA converts our Color into an RGBA representation from the color library
func Vec3ToRGBA(c mgl64.Vec3) color.RGBA {
	return color.RGBA{
		uint8(c[0] * float64(math.MaxUint8)),
		uint8(c[1] * float64(math.MaxUint8)),
		uint8(c[2] * float64(math.MaxUint8)),
		uint8(1.0 * float64(math.MaxUint8))}
}

// Vec3ToRGBA64 converts our Color into an RGBA64 representation from the color library
func Vec3ToRGBA64(c mgl64.Vec3) color.RGBA64 {
	return color.RGBA64{
		uint16(c[0] * float64(math.MaxUint16)),
		uint16(c[1] * float64(math.MaxUint16)),
		uint16(c[2] * float64(math.MaxUint16)),
		uint16(1.0 * float64(math.MaxUint16))}
}

// // VectorFromColor creates a Vector from a Color
// func ColorToVec3(c mgl64.Vec3) VectorX {
// 	return VectorX{c.Red, c.Green, c.Blue}
// }

// Copy returns a new Vector identical to v
// func (v VectorX) Copy() VectorX {
// 	return VectorX{v.X() v.Y() v.Z}
// }

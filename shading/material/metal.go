package material

import (
	"fluorescence/geometry"
	"math/rand"
)

// Metal is an implementation of a Material
// It represents a perfect or near-perfect specularly reflective material
type Metal struct {
	R         geometry.Vector `json:"r"`
	E         geometry.Vector `json:"e"`
	Fuzziness float64         `json:"fuzziness"`
}

// Reflectance returns the reflectance of this material
func (m Metal) Reflectance() geometry.Vector {
	return m.R
}

// Emittance returns the emittance of this material
func (m Metal) Emittance() geometry.Vector {
	return m.E
}

// IsSpecular returns whether this material is specular in nature (vs. diffuse)
// This is currently unused and is likely to be deprecated in the future
func (m Metal) IsSpecular() bool {
	return true
}

// Scatter returns an incoming ray given a RayHit representing the outgoing ray
func (m Metal) Scatter(rayHit RayHit, rng *rand.Rand) (geometry.Ray, bool) {
	hitPoint := rayHit.Ray.PointAt(rayHit.T)
	normal := rayHit.NormalAtHit

	reflectionVector := rayHit.Ray.Direction.Unit().ReflectAround(normal)
	reflectionVector = reflectionVector.Add(geometry.RandomInUnitSphere(rng).MultScalar(m.Fuzziness))
	if reflectionVector.Dot(normal) > 0 {
		return geometry.Ray{
			Origin:    hitPoint,
			Direction: reflectionVector,
		}, true
	}
	return geometry.RAY_ZERO, false
}

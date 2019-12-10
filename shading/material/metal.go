package material

import (
	"fluorescence/geometry"
	"math/rand"
)

type Metal struct {
	R         *geometry.Vector `json:"r"`
	E         *geometry.Vector `json:"e"`
	Fuzziness float64          `json:"fuzziness"`
}

func (m *Metal) Reflectance() *geometry.Vector {
	return m.R
}

func (m *Metal) Emittance() *geometry.Vector {
	return m.E
}

func (m *Metal) IsSpecular() bool {
	return true
}

func (m *Metal) Scatter(rayHit *RayHit, rng *rand.Rand) (*geometry.Ray, bool) {
	hitPoint := rayHit.Ray.PointAt(rayHit.T)
	normal := rayHit.NormalAtHit

	reflectionVector := rayHit.Ray.Direction.Unit().ReflectAround(normal)
	reflectionVector.AddInPlace(geometry.RandomInUnitSphere(rng).MultScalarInPlace(m.Fuzziness))
	if reflectionVector.Dot(normal) > 0 {
		return &geometry.Ray{hitPoint, reflectionVector}, true
	}
	return nil, false
}

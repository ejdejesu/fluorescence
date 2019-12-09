package material

import (
	"fluorescence/geometry"
	"math/rand"
)

type Material interface {
	Reflectance() *geometry.Vector
	Emittance() *geometry.Vector
	IsSpecular() bool
	Scatter(*RayHit, *rand.Rand) (*geometry.Ray, bool)
}

package material

import (
	"fluorescence/geometry"
	"fluorescence/shading"
	"math/rand"
)

type Material interface {
	Reflectance() shading.Color
	Emittance() shading.Color
	IsSpecular() bool
	Scatter(RayHit, *rand.Rand) (geometry.Ray, bool)
}

type RayHit struct {
	Ray         geometry.Ray
	NormalAtHit geometry.Vector
	T           float64
	Material    Material
}

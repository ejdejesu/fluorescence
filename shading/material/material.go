package material

import (
	"fluorescence/geometry"
	"fluorescence/shading"
	"math/rand"
)

type Material interface {
	Reflectance(u, v float64) shading.Color
	Emittance(u, v float64) shading.Color
	IsSpecular() bool
	Scatter(RayHit, *rand.Rand) (geometry.Ray, bool)
}

type RayHit struct {
	Ray         geometry.Ray
	NormalAtHit geometry.Vector
	Time        float64
	U           float64
	V           float64
	Material    Material
}

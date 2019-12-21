package material

import (
	"fluorescence/geometry"
	"fluorescence/shading"
	"math/rand"
)

// Material described the implementation of a surface material
type Material interface {
	Reflectance(u, v float64) shading.Color
	Emittance(u, v float64) shading.Color
	IsSpecular() bool
	Scatter(RayHit, *rand.Rand) (geometry.Ray, bool)
}

// RayHit is a loose gathering of information about a ray's intersection with a surface
type RayHit struct {
	Ray         geometry.Ray
	NormalAtHit geometry.Vector
	Time        float64
	U           float64
	V           float64
	Material    Material
}

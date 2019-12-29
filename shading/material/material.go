package material

import (
	"fluorescence/geometry"
	"math/rand"

	"github.com/go-gl/mathgl/mgl64"
)

// Material described the implementation of a surface material
type Material interface {
	Reflectance(u, v float64) mgl64.Vec3
	Emittance(u, v float64) mgl64.Vec3
	IsSpecular() bool
	Scatter(RayHit, *rand.Rand) (geometry.Ray, bool)
}

// RayHit is a loose gathering of information about a ray's intersection with a surface
type RayHit struct {
	Ray         geometry.Ray
	NormalAtHit mgl64.Vec3
	Time        float64
	U           float64 // texture coordinate U
	V           float64 // texture coordinate V
	Material    Material
}

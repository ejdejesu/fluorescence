package material

import (
	"fluorescence/geometry"
	"fluorescence/shading/texture"
	"math/rand"

	"github.com/go-gl/mathgl/mgl64"
)

// Lambertian represents an approximation to a ideally-diffuse material
// (which is not physically accurate)
type Lambertian struct {
	ReflectanceTexture texture.Texture `json:"-"`
	EmittanceTexture   texture.Texture `json:"-"`
}

// Reflectance returns the reflective color at texture coordinates (u, v)
func (l Lambertian) Reflectance(u, v float64) mgl64.Vec3 {
	return l.ReflectanceTexture.Value(u, v)
}

// Emittance returns the emissive color at texture coordinates (u, v)
func (l Lambertian) Emittance(u, v float64) mgl64.Vec3 {
	return l.EmittanceTexture.Value(u, v)
}

// IsSpecular returns whether this material is specular in nature (vs. diffuse)
// This is currently unused and is likely to be deprecated in the future
func (l Lambertian) IsSpecular() bool {
	return false
}

// Scatter returns an incoming ray given a RayHit representing the outgoing ray
func (l Lambertian) Scatter(rayHit RayHit, rng *rand.Rand) (geometry.Ray, bool) {
	hitPoint := rayHit.Ray.PointAt(rayHit.Time)
	target := hitPoint.Add(rayHit.NormalAtHit).Add(geometry.RandomInUnitSphere(rng))
	return geometry.Ray{
		Origin:    hitPoint,
		Direction: target.Sub(hitPoint),
	}, true
}

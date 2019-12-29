package material

import (
	"fluorescence/geometry"
	"fluorescence/shading/texture"
	"math/rand"

	"github.com/go-gl/mathgl/mgl64"
)

// Metal is an implementation of a Material
// It represents a perfect or near-perfect specularly reflective material
type Metal struct {
	ReflectanceTexture texture.Texture `json:"-"`
	EmittanceTexture   texture.Texture `json:"-"`
	Fuzziness          float64         `json:"fuzziness"`
}

// Reflectance returns the reflective color at texture coordinates (u, v)
func (m Metal) Reflectance(u, v float64) mgl64.Vec3 {
	return m.ReflectanceTexture.Value(u, v)
}

// Emittance returns the emissive color at texture coordinates (u, v)
func (m Metal) Emittance(u, v float64) mgl64.Vec3 {
	return m.EmittanceTexture.Value(u, v)
}

// IsSpecular returns whether this material is specular in nature (vs. diffuse)
// This is currently unused and is likely to be deprecated in the future
func (m Metal) IsSpecular() bool {
	return true
}

// Scatter returns an incoming ray given a RayHit representing the outgoing ray
func (m Metal) Scatter(rayHit RayHit, rng *rand.Rand) (geometry.Ray, bool) {
	hitPoint := rayHit.Ray.PointAt(rayHit.Time)
	normal := rayHit.NormalAtHit

	reflectionVector := geometry.ReflectAround(rayHit.Ray.Direction.Normalize(), normal)
	reflectionVector = reflectionVector.Add(geometry.RandomInUnitSphere(rng).Mul(m.Fuzziness))
	if reflectionVector.Dot(normal) > 0 {
		return geometry.Ray{
			Origin:    hitPoint,
			Direction: reflectionVector,
		}, true
	}
	return geometry.RayZero, false
}

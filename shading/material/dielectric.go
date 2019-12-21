package material

import (
	"fluorescence/geometry"
	"fluorescence/shading"
	"fluorescence/shading/texture"
	"math"
	"math/rand"
)

// Dielectric is an implementation of a Material
// It represents a partially reflective, partially transmissive material, such as glass
type Dielectric struct {
	ReflectanceTexture texture.Texture `json:"-"`
	EmittanceTexture   texture.Texture `json:"-"`
	RefractiveIndex    float64         `json:"refractive_index"`
}

// Reflectance returns the reflective color at texture coordinates (u, v)
func (d Dielectric) Reflectance(u, v float64) shading.Color {
	return d.ReflectanceTexture.Value(u, v)
}

// Emittance returns the emissive color at texture coordinates (u, v)
func (d Dielectric) Emittance(u, v float64) shading.Color {
	return d.EmittanceTexture.Value(u, v)
}

// IsSpecular returns whether this material is specular in nature (vs. diffuse)
// This is currently unused and is likely to be deprecated in the future
func (d Dielectric) IsSpecular() bool {
	return true
}

// Scatter returns an incoming ray given a RayHit representing the outgoing ray
func (d Dielectric) Scatter(rayHit RayHit, rng *rand.Rand) (geometry.Ray, bool) {
	hitPoint := rayHit.Ray.PointAt(rayHit.Time)
	normal := rayHit.NormalAtHit
	reflectionVector := rayHit.Ray.Direction.Unit().ReflectAround(normal)

	var refractiveNormal geometry.Vector
	var ratioOfRefractiveIndices, cosine float64

	if rayHit.Ray.Direction.Dot(normal) > 0 {
		refractiveNormal = geometry.VECTOR_ZERO.Sub(normal)
		ratioOfRefractiveIndices = d.RefractiveIndex
		preCos := rayHit.Ray.Direction.Dot(normal)
		cosine = math.Sqrt(1.0 - (d.RefractiveIndex*d.RefractiveIndex)*(1.0-(preCos*preCos)))
	} else {
		refractiveNormal = normal
		ratioOfRefractiveIndices = 1.0 / d.RefractiveIndex
		cosine = -(rayHit.Ray.Direction.Dot(normal))
	}

	refractedVector, ok := rayHit.Ray.Direction.RefractAround(refractiveNormal, ratioOfRefractiveIndices)
	var reflectionProbability float64
	reflectionProbability = schlick(cosine, d.RefractiveIndex)

	if !ok || rng.Float64() < reflectionProbability {
		// fmt.Println("reflect!")
		return geometry.Ray{
			Origin:    hitPoint,
			Direction: reflectionVector,
		}, true
	}
	// fmt.Println("refract!")
	return geometry.Ray{
		Origin:    hitPoint,
		Direction: refractedVector,
	}, true

}

// schlick is a polynomial approximation to the chance a ray is reflected or transmitted via a dielectric
func schlick(cosine, refractiveIndex float64) float64 {
	r0 := (1.0 - refractiveIndex) / (1.0 + refractiveIndex)
	r1 := r0 * r0
	return r1 + (1.0-r1)*math.Pow(1.0-cosine, 5.0)
}

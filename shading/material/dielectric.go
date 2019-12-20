package material

import (
	"fluorescence/geometry"
	"fluorescence/shading"
	"math"
	"math/rand"
)

type Dielectric struct {
	R               shading.Color `json:"r"`
	E               shading.Color `json:"e"`
	RefractiveIndex float64       `json:"refractive_index"`
}

func (d Dielectric) Reflectance() shading.Color {
	return d.R
}

func (d Dielectric) Emittance() shading.Color {
	return d.E
}

func (d Dielectric) IsSpecular() bool {
	return true
}

func (d Dielectric) Scatter(rayHit RayHit, rng *rand.Rand) (geometry.Ray, bool) {
	hitPoint := rayHit.Ray.PointAt(rayHit.T)
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
		return geometry.Ray{hitPoint, reflectionVector}, true
	}
	// fmt.Println("refract!")
	return geometry.Ray{hitPoint, refractedVector}, true

}

func schlick(cosine, refractiveIndex float64) float64 {
	r0 := (1.0 - refractiveIndex) / (1.0 + refractiveIndex)
	r1 := r0 * r0
	return r1 + (1.0-r1)*math.Pow(1.0-cosine, 5.0)
}

package material

import (
	"fluorescence/geometry"
	"fluorescence/shading"
	"math/rand"
)

type Lambertian struct {
	R shading.Color `json:"r"`
	E shading.Color `json:"e"`
}

func (l Lambertian) Reflectance() shading.Color {
	return l.R
}

func (l Lambertian) Emittance() shading.Color {
	return l.E
}

func (l Lambertian) IsSpecular() bool {
	return false
}

func (l Lambertian) Scatter(rayHit RayHit, rng *rand.Rand) (geometry.Ray, bool) {
	hitPoint := rayHit.Ray.PointAt(rayHit.T)
	target := hitPoint.AddVector(rayHit.NormalAtHit).AddVector(geometry.RandomInUnitSphere(rng))
	return geometry.Ray{hitPoint, hitPoint.To(target)}, true
}

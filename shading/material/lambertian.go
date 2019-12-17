package material

import (
	"fluorescence/geometry"
	"math/rand"
)

type Lambertian struct {
	R geometry.Vector `json:"r"`
	E geometry.Vector `json:"e"`
}

func (l Lambertian) Reflectance() geometry.Vector {
	return l.R
}

func (l Lambertian) Emittance() geometry.Vector {
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

package material

import (
	"fluorescence/geometry"
	"math/rand"
)

type Lambertian struct {
	Reflectance_ *geometry.Vector
	Emittance_   *geometry.Vector
}

func (l *Lambertian) Reflectance() *geometry.Vector {
	return l.Reflectance_
}

func (l *Lambertian) Emittance() *geometry.Vector {
	return l.Emittance_
}

func (l *Lambertian) IsSpecular() bool {
	return false
}

func (l *Lambertian) Scatter(rayHit *RayHit, rng *rand.Rand) (*geometry.Ray, bool) {
	hitPoint := rayHit.Ray.PointAt(rayHit.T)
	target := hitPoint.AddVector(rayHit.NormalAtHit).AddVector(geometry.RandomInUnitSphere(rng))
	return &geometry.Ray{hitPoint, hitPoint.To(target)}, true
}

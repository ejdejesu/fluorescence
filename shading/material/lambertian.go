package material

import (
	"fluorescence/geometry"
	"fluorescence/shading"
	"fluorescence/shading/texture"
	"math/rand"
)

type Lambertian struct {
	ReflectanceTexture texture.Texture `json:"-"`
	EmittanceTexture   texture.Texture `json:"-"`
}

func (l Lambertian) Reflectance(u, v float64) shading.Color {
	return l.ReflectanceTexture.Value(u, v)
}

func (l Lambertian) Emittance(u, v float64) shading.Color {
	return l.EmittanceTexture.Value(u, v)
}

func (l Lambertian) IsSpecular() bool {
	return false
}

func (l Lambertian) Scatter(rayHit RayHit, rng *rand.Rand) (geometry.Ray, bool) {
	hitPoint := rayHit.Ray.PointAt(rayHit.Time)
	target := hitPoint.AddVector(rayHit.NormalAtHit).AddVector(geometry.RandomInUnitSphere(rng))
	return geometry.Ray{hitPoint, hitPoint.To(target)}, true
}

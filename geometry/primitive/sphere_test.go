package primitive

import (
	"fluorescence/geometry"
	"testing"
)

var sphereHit bool

func basicSphere(xOffset, yOffset, zOffset float64) *Sphere {
	return &Sphere{
		Center: &geometry.Point{
			X: 0.0 + xOffset,
			Y: 0.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		Radius:   0.5,
		Material: nil,
	}
}

func TestSphereHitIntersection(t *testing.T) {
	sphere := basicSphere(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 0.0,
			Y: 0.0,
			Z: 1.0,
		},
		Direction: &geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := sphere.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	if !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	}
}

func BenchmarkSphereHitIntersection(b *testing.B) {
	sphere := basicSphere(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 0.0,
			Y: 0.0,
			Z: 1.0,
		},
		Direction: &geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, h = sphere.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	}
	sphereHit = h
}

func TestSphereMissIntersection(t *testing.T) {
	sphere := basicSphere(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 1.0,
			Y: 0.0,
			Z: 1.0,
		},
		Direction: &geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := sphere.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkSphereMissIntersection(b *testing.B) {
	sphere := basicSphere(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 1.0,
			Y: 0.0,
			Z: 1.0,
		},
		Direction: &geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, h = sphere.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	}
	sphereHit = h
}

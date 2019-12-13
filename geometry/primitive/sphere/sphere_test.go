package sphere

import (
	"fluorescence/geometry"
	"testing"
)

var sphereHit bool

func TestSphereIntersectionHit(t *testing.T) {
	sphere := BasicSphere(0.0, 0.0, 0.0)
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
	_, h := sphere.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	}
}

func BenchmarkSphereIntersectionHit(b *testing.B) {
	sphere := BasicSphere(0.0, 0.0, 0.0)
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
		_, h = sphere.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	sphereHit = h
}

func TestSphereIntersectionMiss(t *testing.T) {
	sphere := BasicSphere(0.0, 0.0, 0.0)
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
	_, h := sphere.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkSphereIntersectionMiss(b *testing.B) {
	sphere := BasicSphere(0.0, 0.0, 0.0)
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
		_, h = sphere.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	sphereHit = h
}

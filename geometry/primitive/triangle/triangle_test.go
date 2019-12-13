package triangle

import (
	"fluorescence/geometry"
	"testing"
)

var triHit bool

func TestTriangleIntersectionHit(t *testing.T) {
	tri := BasicTriangle(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 0.1,
			Y: 0.1,
			Z: 1.0,
		},
		Direction: &geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := tri.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	}
}

func BenchmarkTriangleIntersectionHit(b *testing.B) {
	tri := BasicTriangle(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 0.1,
			Y: 0.1,
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
		_, h = tri.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	triHit = h
}

func TestTriangleIntersectionMiss(t *testing.T) {
	tri := BasicTriangle(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 0.9,
			Y: 0.9,
			Z: 1.0,
		},
		Direction: &geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := tri.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkTriangleIntersectionMiss(b *testing.B) {
	tri := BasicTriangle(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 0.9,
			Y: 0.9,
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
		_, h = tri.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	triHit = h
}

package rectangle

import (
	"fluorescence/geometry"
	"testing"
)

var rectHit bool

func TestRectangleIntersectionHit(t *testing.T) {
	rect := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.5,
			Y: 0.5,
			Z: 1.0,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := rect.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	}
}

func BenchmarkRectangleIntersectionHit(b *testing.B) {
	rect := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.5,
			Y: 0.5,
			Z: 1.0,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, h = rect.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	rectHit = h
}

func TestRectangleIntersectionMiss(t *testing.T) {
	rect := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 1.5,
			Y: 0.5,
			Z: 1.0,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := rect.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkRectangleIntersectionMiss(b *testing.B) {
	rect := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 1.5,
			Y: 0.5,
			Z: 1.0,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, h = rect.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	rectHit = h
}

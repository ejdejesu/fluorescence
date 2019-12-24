package infinitecylinder

import (
	"fluorescence/geometry"
	"testing"
)

var icHit bool

func TestInfiniteCylinderIntersectionHit(t *testing.T) {
	ic := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.0,
			Z: 1.5,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := ic.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	}
}

func BenchmarkInfiniteCylinderIntersectionHit(b *testing.B) {
	ic := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.0,
			Z: 1.5,
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
		_, h = ic.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	icHit = h
}

func TestInfiniteCylinderIntersectionSecondHit(t *testing.T) {
	ic := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.0,
			Z: 0.0,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := ic.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	}
}

func BenchmarkInfiniteCylinderIntersectionSecondHit(b *testing.B) {
	ic := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.0,
			Z: 0.0,
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
		_, h = ic.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	icHit = h
}

func TestInfiniteCylinderIntersectionSideMiss(t *testing.T) {
	ic := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 1.5,
			Y: 0.0,
			Z: 1.5,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := ic.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkInfiniteCylinderIntersectionSideMiss(b *testing.B) {
	ic := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 1.5,
			Y: 0.0,
			Z: 1.5,
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
		_, h = ic.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	icHit = h
}

func TestInfiniteCylinderIntersectionBehindMiss(t *testing.T) {
	ic := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.0,
			Z: -1.5,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := ic.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkInfiniteCylinderIntersectionBehindMiss(b *testing.B) {
	ic := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.0,
			Z: -1.5,
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
		_, h = ic.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	icHit = h
}

func TestInfiniteCylinderIntersectionParallelMiss(t *testing.T) {
	ic := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.0,
			Z: 1.5,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 1.0,
			Z: 0.0,
		},
	}
	_, h := ic.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkInfiniteCylinderIntersectionParallelMiss(b *testing.B) {
	ic := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.0,
			Z: 1.5,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 1.0,
			Z: 0.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, h = ic.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	icHit = h
}

func TestInfiniteCylinderIntersectionInsideParallelMiss(t *testing.T) {
	ic := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.0,
			Z: 0.0,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 1.0,
			Z: 0.0,
		},
	}
	_, h := ic.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkInfiniteCylinderIntersectionInsideParallelMiss(b *testing.B) {
	ic := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.0,
			Z: 0.0,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 1.0,
			Z: 0.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, h = ic.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	icHit = h
}

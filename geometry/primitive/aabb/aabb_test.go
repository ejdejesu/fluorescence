package aabb

import (
	"fluorescence/geometry"
	"testing"
)

var aabbHit bool

func basicAABB(xOffset, yOffset, zOffset float64) *AABB {
	return &AABB{
		A: &geometry.Point{
			X: 0.0 + xOffset,
			Y: 0.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		B: &geometry.Point{
			X: 1.0 + xOffset,
			Y: 1.0 + yOffset,
			Z: 1.0 + zOffset,
		},
	}
}

func TestAABBIntersectionXAxisHit(t *testing.T) {
	aabb := basicAABB(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: -1.0,
			Y: 0.5,
			Z: 0.5,
		},
		Direction: &geometry.Vector{
			X: 1.0,
			Y: 0.0,
			Z: 0.0,
		},
	}
	h := aabb.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	if !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	}
}

func BenchmarkAABBIntersectionXAxisHit(b *testing.B) {
	aabb := basicAABB(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: -1.0,
			Y: 0.5,
			Z: 0.5,
		},
		Direction: &geometry.Vector{
			X: 1.0,
			Y: 0.0,
			Z: 0.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h = aabb.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	}
	aabbHit = h
}

func TestAABBIntersectionXAxisMiss(t *testing.T) {
	aabb := basicAABB(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: -1.0,
			Y: 1.5,
			Z: 0.5,
		},
		Direction: &geometry.Vector{
			X: 1.0,
			Y: 0.0,
			Z: 0.0,
		},
	}
	h := aabb.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkAABBIntersectionXAxisMiss(b *testing.B) {
	aabb := basicAABB(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: -1.0,
			Y: 1.5,
			Z: 0.5,
		},
		Direction: &geometry.Vector{
			X: 1.0,
			Y: 0.0,
			Z: 0.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h = aabb.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	}
	aabbHit = h
}

func TestAABBIntersectionYAxisHit(t *testing.T) {
	aabb := basicAABB(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 0.5,
			Y: -1.0,
			Z: 0.5,
		},
		Direction: &geometry.Vector{
			X: 0.0,
			Y: 1.0,
			Z: 0.0,
		},
	}
	h := aabb.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	if !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	}
}

func BenchmarkAABBIntersectionYAxisHit(b *testing.B) {
	aabb := basicAABB(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 0.5,
			Y: -1.0,
			Z: 0.5,
		},
		Direction: &geometry.Vector{
			X: 0.0,
			Y: 1.0,
			Z: 0.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h = aabb.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	}
	aabbHit = h
}

func TestAABBIntersectionYAxisMiss(t *testing.T) {
	aabb := basicAABB(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 0.5,
			Y: -1.0,
			Z: 1.5,
		},
		Direction: &geometry.Vector{
			X: 0.0,
			Y: 1.0,
			Z: 0.0,
		},
	}
	h := aabb.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkAABBIntersectionYAxisMiss(b *testing.B) {
	aabb := basicAABB(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 0.5,
			Y: -1.0,
			Z: 1.5,
		},
		Direction: &geometry.Vector{
			X: 0.0,
			Y: 1.0,
			Z: 0.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h = aabb.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	}
	aabbHit = h
}

func TestAABBIntersectionZAxisHit(t *testing.T) {
	aabb := basicAABB(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 0.5,
			Y: 0.5,
			Z: -1.0,
		},
		Direction: &geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: 1.0,
		},
	}
	h := aabb.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	if !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	}
}

func BenchmarkAABBIntersectionZAxisHit(b *testing.B) {
	aabb := basicAABB(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 0.5,
			Y: 0.5,
			Z: -1.0,
		},
		Direction: &geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: 1.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h = aabb.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	}
	aabbHit = h
}

func TestAABBIntersectionZAxisMiss(t *testing.T) {
	aabb := basicAABB(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 1.5,
			Y: 0.5,
			Z: -1.0,
		},
		Direction: &geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: 1.0,
		},
	}
	h := aabb.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkAABBIntersectionZAxisMiss(b *testing.B) {
	aabb := basicAABB(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 1.5,
			Y: 0.5,
			Z: -1.0,
		},
		Direction: &geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: 1.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h = aabb.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	}
	aabbHit = h
}

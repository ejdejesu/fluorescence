package primitive

import (
	"fluorescence/geometry"
	"testing"
)

var rectHit bool

func basicRectangle(xOffset, yOffset, zOffset float64) *Rectangle {
	rd := RectangleData{
		A: &geometry.Point{
			X: 0.0 + xOffset,
			Y: 0.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		B: &geometry.Point{
			X: 1.0 + xOffset,
			Y: 1.0 + yOffset,
			Z: 0.0 + zOffset,
		},
	}
	r, _ := NewRectangle(&rd)
	return r
}

func TestRectangleHitIntersection(t *testing.T) {
	rect := basicRectangle(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 0.5,
			Y: 0.5,
			Z: 1.0,
		},
		Direction: &geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := rect.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	if !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	}
}

func BenchmarkOneRectangleHitIntersection(b *testing.B) {
	rect := basicRectangle(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 0.5,
			Y: 0.5,
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
		_, h = rect.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	}
	rectHit = h
}

func TestRectangleMissIntersection(t *testing.T) {
	rect := basicRectangle(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 1.5,
			Y: 0.5,
			Z: 1.0,
		},
		Direction: &geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := rect.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkOneRectangleMissIntersection(b *testing.B) {
	rect := basicRectangle(0.0, 0.0, 0.0)
	r := &geometry.Ray{
		Origin: &geometry.Point{
			X: 1.5,
			Y: 0.5,
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
		_, h = rect.Intersection(r, 0.0000001, 1.797693134862315708145274237317043567981e+308)
	}
	rectHit = h
}

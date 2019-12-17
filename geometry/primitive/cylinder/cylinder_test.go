package cylinder

import (
	"fluorescence/geometry"
	"testing"
)

var cHit bool

func TestCylinderIntersectionSideHit(t *testing.T) {
	c := BasicCylinder(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.5,
			Z: 1.5,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := c.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	}
}

func BenchmarkCylinderIntersectionSideHit(b *testing.B) {
	c := BasicCylinder(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.5,
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
		_, h = c.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	cHit = h
}

func TestCylinderIntersectionSecondHit(t *testing.T) {
	c := BasicCylinder(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.5,
			Z: 0.0,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := c.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	}
}

func BenchmarkCylinderIntersectionSecondHit(b *testing.B) {
	c := BasicCylinder(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.5,
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
		_, h = c.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	cHit = h
}

func TestCylinderIntersectionBottomCapHit(t *testing.T) {
	c := BasicCylinder(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: -1.0,
			Z: 1.0,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 1.0,
			Z: -1.0,
		},
	}
	_, h := c.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	}
}

func BenchmarkCylinderIntersectionBottomCapHit(b *testing.B) {
	c := BasicCylinder(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: -1.0,
			Z: 1.0,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 1.0,
			Z: -1.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, h = c.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	cHit = h
}

func TestCylinderIntersectionTopCapHit(t *testing.T) {
	c := BasicCylinder(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 2.0,
			Z: 1.0,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: -1.0,
			Z: -1.0,
		},
	}
	_, h := c.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	}
}

func BenchmarkCylinderIntersectionTopCapHit(b *testing.B) {
	c := BasicCylinder(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 2.0,
			Z: 1.0,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: -1.0,
			Z: -1.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, h = c.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	cHit = h
}

func TestCylinderIntersectionSideMiss(t *testing.T) {
	c := BasicCylinder(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 1.5,
			Y: 0.5,
			Z: 1.5,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := c.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkCylinderIntersectionSideMiss(b *testing.B) {
	c := BasicCylinder(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 1.5,
			Y: 0.5,
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
		_, h = c.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	cHit = h
}

func TestCylinderIntersectionBehindMiss(t *testing.T) {
	c := BasicCylinder(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.5,
			Z: -1.5,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := c.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkCylinderIntersectionBehindMiss(b *testing.B) {
	c := BasicCylinder(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.5,
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
		_, h = c.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	cHit = h
}

func TestCylinderIntersectionParallelMiss(t *testing.T) {
	c := BasicCylinder(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.5,
			Z: -1.5,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 1.0,
			Z: 0.0,
		},
	}
	_, h := c.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkCylinderIntersectionParallelMiss(b *testing.B) {
	c := BasicCylinder(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.5,
			Z: -1.5,
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
		_, h = c.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	cHit = h
}

func TestCylinderIntersectionTripleMiss(t *testing.T) {
	c := BasicCylinder(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 1.5,
			Z: -1.5,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 1.0,
			Z: -1.0,
		},
	}
	_, h := c.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkCylinderIntersectionTripleMiss(b *testing.B) {
	c := BasicCylinder(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 1.5,
			Z: -1.5,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 1.0,
			Z: -1.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, h = c.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	cHit = h
}

func TestCylinderIntersectionAABBMiss(t *testing.T) {
	box, _ := BasicCylinder(0.0, 0.0, 0.0).BoundingBox(0, 0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 1.5,
			Z: -1.5,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 1.0,
			Z: -1.0,
		},
	}
	h := box.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkCylinderIntersectionAABBMiss(b *testing.B) {
	c := BasicCylinder(0.0, 0.0, 0.0)
	box, _ := c.BoundingBox(0, 0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 1.5,
			Z: -1.5,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 1.0,
			Z: -1.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if box.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308) {
			_, h = c.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
		}
	}
	cHit = h
}

package plane

import (
	"fluorescence/geometry"
	"testing"
)

var planeHit bool

func TestPlaneIntersectionHit(t *testing.T) {
	plane := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin:mgl64.Vec3{
			X: 0.5,
			Y: 0.5,
			Z: 1.0,
		},
		Direction:mgl64.Vec3{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := plane.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	}
}

func BenchmarkPlaneIntersectionHit(b *testing.B) {
	plane := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin:mgl64.Vec3{
			X: 0.5,
			Y: 0.5,
			Z: 1.0,
		},
		Direction:mgl64.Vec3{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, h = plane.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	planeHit = h
}

func TestPlaneIntersectionMiss(t *testing.T) {
	plane := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin:mgl64.Vec3{
			X: 0.5,
			Y: 0.5,
			Z: 1.0,
		},
		Direction:mgl64.Vec3{
			X: 1.0,
			Y: 0.0,
			Z: 0.0,
		},
	}
	_, h := plane.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkPlaneIntersectionMiss(b *testing.B) {
	plane := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin:mgl64.Vec3{
			X: 0.5,
			Y: 0.5,
			Z: 1.0,
		},
		Direction:mgl64.Vec3{
			X: 1.0,
			Y: 0.0,
			Z: 0.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, h = plane.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	planeHit = h
}

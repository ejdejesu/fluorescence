package hollowcylinder

import (
	"fluorescence/geometry"
	"testing"
)

var hcHit bool

func TestHollowCylinderIntersectionHit(t *testing.T) {
	hc := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin:mgl64.Vec3{
			X: 0.0,
			Y: 0.5,
			Z: 1.0,
		},
		Direction:mgl64.Vec3{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := hc.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	}
}

func BenchmarkHollowCylinderIntersectionHit(b *testing.B) {
	hc := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin:mgl64.Vec3{
			X: 0.0,
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
		_, h = hc.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	hcHit = h
}

func TestHollowCylinderIntersectionMiss(t *testing.T) {
	hc := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin:mgl64.Vec3{
			X: 1.5,
			Y: 0.5,
			Z: 1.0,
		},
		Direction:mgl64.Vec3{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := hc.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkHollowCylinderIntersectionMiss(b *testing.B) {
	hc := Unit(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin:mgl64.Vec3{
			X: 1.5,
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
		_, h = hc.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	hcHit = h
}

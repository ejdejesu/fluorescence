package hollowdisk

import (
	"fluorescence/geometry"
	"testing"
)

var hollowDiskHit bool

func TestHollowDiskIntersectionHit(t *testing.T) {
	hollowDisk := UnitHollowDisk(0.0, 0.0, 0.0)
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
	_, h := hollowDisk.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	}
}

func BenchmarkHollowDiskIntersectionHit(b *testing.B) {
	hollowDisk := UnitHollowDisk(0.0, 0.0, 0.0)
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
		_, h = hollowDisk.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	hollowDiskHit = h
}

func TestHollowDiskIntersectionReverseHit(t *testing.T) {
	hollowDisk := UnitHollowDisk(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.5,
			Y: 0.5,
			Z: -1.0,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: 1.0,
		},
	}
	_, h := hollowDisk.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	}
}

func BenchmarkHollowDiskIntersectionReverseHit(b *testing.B) {
	hollowDisk := UnitHollowDisk(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.5,
			Y: 0.5,
			Z: -1.0,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: 1.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, h = hollowDisk.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	hollowDiskHit = h
}

func TestHollowDiskIntersectionMiss(t *testing.T) {
	hollowDisk := UnitHollowDisk(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 1.0,
			Y: 1.0,
			Z: 1.0,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := hollowDisk.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkHollowDiskIntersectionMiss(b *testing.B) {
	hollowDisk := UnitHollowDisk(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 1.0,
			Y: 1.0,
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
		_, h = hollowDisk.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	hollowDiskHit = h
}

func TestHollowDiskIntersectionCenterMiss(t *testing.T) {
	hollowDisk := UnitHollowDisk(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.0,
			Z: 1.0,
		},
		Direction: geometry.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -1.0,
		},
	}
	_, h := hollowDisk.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkHollowDiskIntersectionCenterMiss(b *testing.B) {
	hollowDisk := UnitHollowDisk(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.0,
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
		_, h = hollowDisk.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	hollowDiskHit = h
}

func TestHollowDiskIntersectionParallelMiss(t *testing.T) {
	hollowDisk := UnitHollowDisk(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.0,
			Z: 1.0,
		},
		Direction: geometry.Vector{
			X: 1.0,
			Y: 0.0,
			Z: 0.0,
		},
	}
	_, h := hollowDisk.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func BenchmarkHollowDiskIntersectionParallelMiss(b *testing.B) {
	hollowDisk := UnitHollowDisk(0.0, 0.0, 0.0)
	r := geometry.Ray{
		Origin: geometry.Point{
			X: 0.0,
			Y: 0.0,
			Z: 1.0,
		},
		Direction: geometry.Vector{
			X: 1.0,
			Y: 0.0,
			Z: 0.0,
		},
	}
	var h bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, h = hollowDisk.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	hollowDiskHit = h
}

package bvh

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive/primitivelist"
	"fluorescence/geometry/primitive/rectangle"
	"fluorescence/geometry/primitive/sphere"
	"fluorescence/geometry/primitive/triangle"
	"math/rand"
	"testing"
)

var bvhHit bool

func UnitBVHNTriangles(n int, xOffset, yOffset, zOffset float64) *BVH {
	pl := &primitivelist.PrimitiveList{}
	for i := 0; i < n; i++ {
		pl.List = append(pl.List, triangle.Unit(xOffset+float64(i), yOffset, zOffset))
	}
	bvh, _ := New(pl)
	return bvh
}

func UnitBVHNRectangles(n int, xOffset, yOffset, zOffset float64) *BVH {
	pl := &primitivelist.PrimitiveList{}
	for i := 0; i < n; i++ {
		pl.List = append(pl.List, rectangle.Unit(xOffset+float64(i), yOffset, zOffset))
	}
	bvh, _ := New(pl)
	return bvh
}

func UnitBVHNSpheres(n int, xOffset, yOffset, zOffset float64) *BVH {
	pl := &primitivelist.PrimitiveList{}
	for i := 0; i < n; i++ {
		pl.List = append(pl.List, sphere.Unit(xOffset+float64(i), yOffset, zOffset))
	}
	bvh, _ := New(pl)
	return bvh
}

func ithTriangleOfNBVHTest(i int, n int, shouldHit bool, t *testing.T) {
	pl := UnitBVHNTriangles(n, 0.0, 0.0, 0.0)
	var r geometry.Ray
	if shouldHit {
		r = geometry.Ray{
			Origin: geometry.Point{
				X: 0.1 + float64(i-1),
				Y: 0.1,
				Z: 1.0,
			},
			Direction: geometry.Vector{
				X: 0.0,
				Y: 0.0,
				Z: -1.0,
			},
		}

	} else {
		r = geometry.Ray{
			Origin: geometry.Point{
				X: 0.9,
				Y: 0.9,
				Z: 1.0,
			},
			Direction: geometry.Vector{
				X: 0.0,
				Y: 0.0,
				Z: -1.0,
			},
		}
	}
	_, h := pl.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if shouldHit && !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	} else if !shouldHit && h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func ithTriangleOfNBVHBenchmark(i int, n int, shouldHit bool, b *testing.B) {
	pl := UnitBVHNTriangles(n, 0.0, 0.0, 0.0)
	var r geometry.Ray
	if shouldHit {
		r = geometry.Ray{
			Origin: geometry.Point{
				X: 0.1 + float64(i-1),
				Y: 0.1,
				Z: 1.0,
			},
			Direction: geometry.Vector{
				X: 0.0,
				Y: 0.0,
				Z: -1.0,
			},
		}

	} else {
		r = geometry.Ray{
			Origin: geometry.Point{
				X: 0.9,
				Y: 0.9,
				Z: 1.0,
			},
			Direction: geometry.Vector{
				X: 0.0,
				Y: 0.0,
				Z: -1.0,
			},
		}
	}
	b.ResetTimer()
	var h bool
	for i := 0; i < b.N; i++ {
		_, h = pl.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	bvhHit = h
}

func ithSphereOfNBVHTest(i int, n int, shouldHit bool, t *testing.T) {
	pl := UnitBVHNSpheres(n, 0.0, 0.0, 0.0)
	var r geometry.Ray
	if shouldHit {
		r = geometry.Ray{
			Origin: geometry.Point{
				X: 0.0 + float64(i-1),
				Y: 0.0,
				Z: 1.0,
			},
			Direction: geometry.Vector{
				X: 0.0,
				Y: 0.0,
				Z: -1.0,
			},
		}

	} else {
		r = geometry.Ray{
			Origin: geometry.Point{
				X: 0.0,
				Y: 1.0,
				Z: 1.0,
			},
			Direction: geometry.Vector{
				X: 0.0,
				Y: 0.0,
				Z: -1.0,
			},
		}
	}
	_, h := pl.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if shouldHit && !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	} else if !shouldHit && h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func ithSphereOfNBVHBenchmark(i int, n int, shouldHit bool, b *testing.B) {
	pl := UnitBVHNSpheres(n, 0.0, 0.0, 0.0)
	var r geometry.Ray
	if shouldHit {
		r = geometry.Ray{
			Origin: geometry.Point{
				X: 0.0 + float64(i-1),
				Y: 0.0,
				Z: 1.0,
			},
			Direction: geometry.Vector{
				X: 0.0,
				Y: 0.0,
				Z: -1.0,
			},
		}

	} else {
		r = geometry.Ray{
			Origin: geometry.Point{
				X: 0.0,
				Y: 1.0,
				Z: 1.0,
			},
			Direction: geometry.Vector{
				X: 0.0,
				Y: 0.0,
				Z: -1.0,
			},
		}
	}
	b.ResetTimer()
	var h bool
	for i := 0; i < b.N; i++ {
		_, h = pl.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	bvhHit = h
}

func ithRectangleOfNBVHTest(i int, n int, shouldHit bool, t *testing.T) {
	pl := UnitBVHNRectangles(n, 0.0, 0.0, 0.0)
	var r geometry.Ray
	if shouldHit {
		r = geometry.Ray{
			Origin: geometry.Point{
				X: 0.5 + float64(i-1),
				Y: 0.5,
				Z: 1.0,
			},
			Direction: geometry.Vector{
				X: 0.0,
				Y: 0.0,
				Z: -1.0,
			},
		}

	} else {
		r = geometry.Ray{
			Origin: geometry.Point{
				X: 0.5,
				Y: 1.5,
				Z: 1.0,
			},
			Direction: geometry.Vector{
				X: 0.0,
				Y: 0.0,
				Z: -1.0,
			},
		}
	}
	_, h := pl.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	if shouldHit && !h {
		t.Errorf("Expected true (hit) but got %t\n", h)
	} else if !shouldHit && h {
		t.Errorf("Expected false (miss) but got %t\n", h)
	}
}

func ithRectangleOfNBVHBenchmark(i int, n int, shouldHit bool, b *testing.B) {
	pl := UnitBVHNRectangles(n, 0.0, 0.0, 0.0)
	var r geometry.Ray
	if shouldHit {
		r = geometry.Ray{
			Origin: geometry.Point{
				X: 0.5 + float64(i-1),
				Y: 0.5,
				Z: 1.0,
			},
			Direction: geometry.Vector{
				X: 0.0,
				Y: 0.0,
				Z: -1.0,
			},
		}

	} else {
		r = geometry.Ray{
			Origin: geometry.Point{
				X: 0.5,
				Y: 1.5,
				Z: 1.0,
			},
			Direction: geometry.Vector{
				X: 0.0,
				Y: 0.0,
				Z: -1.0,
			},
		}
	}
	b.ResetTimer()
	var h bool
	for i := 0; i < b.N; i++ {
		_, h = pl.Intersection(r, 1e-7, 1.797693134862315708145274237317043567981e+308)
	}
	bvhHit = h
}

func TestBVHIntersectionHitFirstRectangleOf10(t *testing.T) {
	ithRectangleOfNBVHTest(1, 10, true, t)
}

func BenchmarkBVHIntersectionHitFirstRectangleOf10(b *testing.B) {
	ithRectangleOfNBVHBenchmark(1, 10, true, b)
}

func TestBVHIntersectionHitLastRectangleOf10(t *testing.T) {
	ithRectangleOfNBVHTest(10, 10, true, t)
}

func BenchmarkBVHIntersectionHitLastRectangleOf10(b *testing.B) {
	ithRectangleOfNBVHBenchmark(10, 10, true, b)
}

func TestBVHIntersectionMissRectangleOf10(t *testing.T) {
	ithRectangleOfNBVHTest(1, 10, false, t)
}

func BenchmarkBVHIntersectionMissRectangleOf10(b *testing.B) {
	ithRectangleOfNBVHBenchmark(1, 10, false, b)
}

func TestBVHIntersectionHitFirstRectangleOf1000(t *testing.T) {
	ithRectangleOfNBVHTest(1, 1000, true, t)
}

func BenchmarkBVHIntersectionHitFirstRectangleOf1000(b *testing.B) {
	ithRectangleOfNBVHBenchmark(1, 1000, true, b)
}

func TestBVHIntersectionHitLastRectangleOf1000(t *testing.T) {
	ithRectangleOfNBVHTest(1000, 1000, true, t)
}

func BenchmarkBVHIntersectionHitLastRectangleOf1000(b *testing.B) {
	ithRectangleOfNBVHBenchmark(1000, 1000, true, b)
}

func TestBVHIntersectionHitRandomRectangleOf1000(t *testing.T) {
	ithRectangleOfNBVHTest(rand.Intn(1000)+1, 1000, true, t)
}

func BenchmarkBVHIntersectionHitRandomRectangleOf1000(b *testing.B) {
	ithRectangleOfNBVHBenchmark(rand.Intn(1000)+1, 1000, true, b)
}

func TestBVHIntersectionMissRectangleOf1000(t *testing.T) {
	ithRectangleOfNBVHTest(1, 1000, false, t)
}

func BenchmarkBVHIntersectionMissRectangleOf1000(b *testing.B) {
	ithRectangleOfNBVHBenchmark(1, 1000, false, b)
}

func TestBVHIntersectionHitFirstSphereOf10(t *testing.T) {
	ithSphereOfNBVHTest(1, 10, true, t)
}

func BenchmarkBVHIntersectionHitFirstSphereOf10(b *testing.B) {
	ithSphereOfNBVHBenchmark(1, 10, true, b)
}

func TestBVHIntersectionHitLastSphereOf10(t *testing.T) {
	ithSphereOfNBVHTest(10, 10, true, t)
}

func BenchmarkBVHIntersectionHitLastSphereOf10(b *testing.B) {
	ithSphereOfNBVHBenchmark(10, 10, true, b)
}

func TestBVHIntersectionMissSphereOf10(t *testing.T) {
	ithSphereOfNBVHTest(1, 10, false, t)
}

func BenchmarkBVHIntersectionMissSphereOf10(b *testing.B) {
	ithSphereOfNBVHBenchmark(1, 10, false, b)
}

func TestBVHIntersectionHitFirstSphereOf1000(t *testing.T) {
	ithSphereOfNBVHTest(1, 1000, true, t)
}

func BenchmarkBVHIntersectionHitFirstSphereOf1000(b *testing.B) {
	ithSphereOfNBVHBenchmark(1, 1000, true, b)
}

func TestBVHIntersectionHitLastSphereOf1000(t *testing.T) {
	ithSphereOfNBVHTest(1000, 1000, true, t)
}

func BenchmarkBVHIntersectionHitLastSphereOf1000(b *testing.B) {
	ithSphereOfNBVHBenchmark(1000, 1000, true, b)
}

func TestBVHIntersectionMissSphereOf1000(t *testing.T) {
	ithSphereOfNBVHTest(1, 1000, false, t)
}

func BenchmarkBVHIntersectionMissSphereOf1000(b *testing.B) {
	ithSphereOfNBVHBenchmark(1, 1000, false, b)
}

func TestBVHIntersectionHitFirstTriangleOf10(t *testing.T) {
	ithTriangleOfNBVHTest(1, 10, true, t)
}

func BenchmarkBVHIntersectionHitFirstTriangleOf10(b *testing.B) {
	ithTriangleOfNBVHBenchmark(1, 10, true, b)
}

func TestBVHIntersectionHitLastTriangleOf10(t *testing.T) {
	ithTriangleOfNBVHTest(10, 10, true, t)
}

func BenchmarkBVHIntersectionHitLastTriangleOf10(b *testing.B) {
	ithTriangleOfNBVHBenchmark(10, 10, true, b)
}

func TestBVHIntersectionMissTriangleOf10(t *testing.T) {
	ithTriangleOfNBVHTest(1, 10, false, t)
}

func BenchmarkBVHIntersectionMissTriangleOf10(b *testing.B) {
	ithTriangleOfNBVHBenchmark(1, 10, false, b)
}

func TestBVHIntersectionHitFirstTriangleOf1000(t *testing.T) {
	ithTriangleOfNBVHTest(1, 1000, true, t)
}

func BenchmarkBVHIntersectionHitFirstTriangleOf1000(b *testing.B) {
	ithTriangleOfNBVHBenchmark(1, 1000, true, b)
}

func TestBVHIntersectionHitLastTriangleOf1000(t *testing.T) {
	ithTriangleOfNBVHTest(1000, 1000, true, t)
}

func BenchmarkBVHIntersectionHitLastTriangleOf1000(b *testing.B) {
	ithTriangleOfNBVHBenchmark(1000, 1000, true, b)
}

func TestBVHIntersectionMissTriangleOf1000(t *testing.T) {
	ithTriangleOfNBVHTest(1, 1000, false, t)
}

func BenchmarkBVHIntersectionMissTriangleOf1000(b *testing.B) {
	ithTriangleOfNBVHBenchmark(1, 1000, false, b)
}

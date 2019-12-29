package primitivelist

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive/rectangle"
	"fluorescence/geometry/primitive/sphere"
	"fluorescence/geometry/primitive/triangle"
	"math/rand"
	"testing"
)

var plHit bool

func UnitPrimitiveListNTriangles(n int, xOffset, yOffset, zOffset float64) *PrimitiveList {
	pl := &PrimitiveList{}
	for i := 0; i < n; i++ {
		pl.List = append(pl.List, triangle.Unit(xOffset+float64(i), yOffset, zOffset))
	}
	return pl
}

func UnitPrimitiveListNRectangles(n int, xOffset, yOffset, zOffset float64) *PrimitiveList {
	pl := &PrimitiveList{}
	for i := 0; i < n; i++ {
		pl.List = append(pl.List, rectangle.Unit(xOffset+float64(i), yOffset, zOffset))
	}
	return pl
}

func UnitPrimitiveListNSpheres(n int, xOffset, yOffset, zOffset float64) *PrimitiveList {
	pl := &PrimitiveList{}
	for i := 0; i < n; i++ {
		pl.List = append(pl.List, sphere.Unit(xOffset+float64(i), yOffset, zOffset))
	}
	return pl
}

func ithTriangleOfNPrimitiveListTest(i int, n int, shouldHit bool, t *testing.T) {
	pl := UnitPrimitiveListNTriangles(n, 0.0, 0.0, 0.0)
	var r geometry.Ray
	if shouldHit {
		r = geometry.Ray{
			Origin:mgl64.Vec3{
				X: 0.1 + float64(i-1),
				Y: 0.1,
				Z: 1.0,
			},
			Direction:mgl64.Vec3{
				X: 0.0,
				Y: 0.0,
				Z: -1.0,
			},
		}

	} else {
		r = geometry.Ray{
			Origin:mgl64.Vec3{
				X: 0.9,
				Y: 0.9,
				Z: 1.0,
			},
			Direction:mgl64.Vec3{
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

func ithTriangleOfNPrimitiveListBenchmark(i int, n int, shouldHit bool, b *testing.B) {
	pl := UnitPrimitiveListNTriangles(n, 0.0, 0.0, 0.0)
	var r geometry.Ray
	if shouldHit {
		r = geometry.Ray{
			Origin:mgl64.Vec3{
				X: 0.1 + float64(i-1),
				Y: 0.1,
				Z: 1.0,
			},
			Direction:mgl64.Vec3{
				X: 0.0,
				Y: 0.0,
				Z: -1.0,
			},
		}

	} else {
		r = geometry.Ray{
			Origin:mgl64.Vec3{
				X: 0.9,
				Y: 0.9,
				Z: 1.0,
			},
			Direction:mgl64.Vec3{
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
	plHit = h
}

func ithSphereOfNPrimitiveListTest(i int, n int, shouldHit bool, t *testing.T) {
	pl := UnitPrimitiveListNSpheres(n, 0.0, 0.0, 0.0)
	var r geometry.Ray
	if shouldHit {
		r = geometry.Ray{
			Origin:mgl64.Vec3{
				X: 0.0 + float64(i-1),
				Y: 0.0,
				Z: 1.0,
			},
			Direction:mgl64.Vec3{
				X: 0.0,
				Y: 0.0,
				Z: -1.0,
			},
		}

	} else {
		r = geometry.Ray{
			Origin:mgl64.Vec3{
				X: 0.0,
				Y: 1.0,
				Z: 1.0,
			},
			Direction:mgl64.Vec3{
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

func ithSphereOfNPrimitiveListBenchmark(i int, n int, shouldHit bool, b *testing.B) {
	pl := UnitPrimitiveListNSpheres(n, 0.0, 0.0, 0.0)
	var r geometry.Ray
	if shouldHit {
		r = geometry.Ray{
			Origin:mgl64.Vec3{
				X: 0.0 + float64(i-1),
				Y: 0.0,
				Z: 1.0,
			},
			Direction:mgl64.Vec3{
				X: 0.0,
				Y: 0.0,
				Z: -1.0,
			},
		}

	} else {
		r = geometry.Ray{
			Origin:mgl64.Vec3{
				X: 0.0,
				Y: 1.0,
				Z: 1.0,
			},
			Direction:mgl64.Vec3{
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
	plHit = h
}

func ithRectangleOfNPrimitiveListTest(i int, n int, shouldHit bool, t *testing.T) {
	pl := UnitPrimitiveListNRectangles(n, 0.0, 0.0, 0.0)
	var r geometry.Ray
	if shouldHit {
		r = geometry.Ray{
			Origin:mgl64.Vec3{
				X: 0.5 + float64(i-1),
				Y: 0.5,
				Z: 1.0,
			},
			Direction:mgl64.Vec3{
				X: 0.0,
				Y: 0.0,
				Z: -1.0,
			},
		}

	} else {
		r = geometry.Ray{
			Origin:mgl64.Vec3{
				X: 0.5,
				Y: 1.5,
				Z: 1.0,
			},
			Direction:mgl64.Vec3{
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

func ithRectangleOfNPrimitiveListBenchmark(i int, n int, shouldHit bool, b *testing.B) {
	pl := UnitPrimitiveListNRectangles(n, 0.0, 0.0, 0.0)
	var r geometry.Ray
	if shouldHit {
		r = geometry.Ray{
			Origin:mgl64.Vec3{
				X: 0.5 + float64(i-1),
				Y: 0.5,
				Z: 1.0,
			},
			Direction:mgl64.Vec3{
				X: 0.0,
				Y: 0.0,
				Z: -1.0,
			},
		}

	} else {
		r = geometry.Ray{
			Origin:mgl64.Vec3{
				X: 0.5,
				Y: 1.5,
				Z: 1.0,
			},
			Direction:mgl64.Vec3{
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
	plHit = h
}

func TestPrimitiveListIntersectionHitFirstRectangleOf10(t *testing.T) {
	ithRectangleOfNPrimitiveListTest(1, 10, true, t)
}

func BenchmarkPrimitiveListIntersectionHitFirstRectangleOf10(b *testing.B) {
	ithRectangleOfNPrimitiveListBenchmark(1, 10, true, b)
}

func TestPrimitiveListIntersectionHitLastRectangleOf10(t *testing.T) {
	ithRectangleOfNPrimitiveListTest(10, 10, true, t)
}

func BenchmarkPrimitiveListIntersectionHitLastRectangleOf10(b *testing.B) {
	ithRectangleOfNPrimitiveListBenchmark(10, 10, true, b)
}

func TestPrimitiveListIntersectionMissRectangleOf10(t *testing.T) {
	ithRectangleOfNPrimitiveListTest(1, 10, false, t)
}

func BenchmarkPrimitiveListIntersectionMissRectangleOf10(b *testing.B) {
	ithRectangleOfNPrimitiveListBenchmark(1, 10, false, b)
}

func TestPrimitiveListIntersectionHitFirstRectangleOf1000(t *testing.T) {
	ithRectangleOfNPrimitiveListTest(1, 1000, true, t)
}

func BenchmarkPrimitiveListIntersectionHitFirstRectangleOf1000(b *testing.B) {
	ithRectangleOfNPrimitiveListBenchmark(1, 1000, true, b)
}

func TestPrimitiveListIntersectionHitLastRectangleOf1000(t *testing.T) {
	ithRectangleOfNPrimitiveListTest(1000, 1000, true, t)
}

func BenchmarkPrimitiveListIntersectionHitLastRectangleOf1000(b *testing.B) {
	ithRectangleOfNPrimitiveListBenchmark(1000, 1000, true, b)
}

func TestPrimitiveListIntersectionHitRandomRectangleOf1000(t *testing.T) {
	ithRectangleOfNPrimitiveListTest(rand.Intn(1000)+1, 1000, true, t)
}

func BenchmarkPrimitiveListIntersectionHitRandomRectangleOf1000(b *testing.B) {
	ithRectangleOfNPrimitiveListBenchmark(rand.Intn(1000)+1, 1000, true, b)
}

func TestPrimitiveListIntersectionMissRectangleOf1000(t *testing.T) {
	ithRectangleOfNPrimitiveListTest(1, 1000, false, t)
}

func BenchmarkPrimitiveListIntersectionMissRectangleOf1000(b *testing.B) {
	ithRectangleOfNPrimitiveListBenchmark(1, 1000, false, b)
}

func TestPrimitiveListIntersectionHitFirstSphereOf10(t *testing.T) {
	ithSphereOfNPrimitiveListTest(1, 10, true, t)
}

func BenchmarkPrimitiveListIntersectionHitFirstSphereOf10(b *testing.B) {
	ithSphereOfNPrimitiveListBenchmark(1, 10, true, b)
}

func TestPrimitiveListIntersectionHitLastSphereOf10(t *testing.T) {
	ithSphereOfNPrimitiveListTest(10, 10, true, t)
}

func BenchmarkPrimitiveListIntersectionHitLastSphereOf10(b *testing.B) {
	ithSphereOfNPrimitiveListBenchmark(10, 10, true, b)
}

func TestPrimitiveListIntersectionMissSphereOf10(t *testing.T) {
	ithSphereOfNPrimitiveListTest(1, 10, false, t)
}

func BenchmarkPrimitiveListIntersectionMissSphereOf10(b *testing.B) {
	ithSphereOfNPrimitiveListBenchmark(1, 10, false, b)
}

func TestPrimitiveListIntersectionHitFirstSphereOf1000(t *testing.T) {
	ithSphereOfNPrimitiveListTest(1, 1000, true, t)
}

func BenchmarkPrimitiveListIntersectionHitFirstSphereOf1000(b *testing.B) {
	ithSphereOfNPrimitiveListBenchmark(1, 1000, true, b)
}

func TestPrimitiveListIntersectionHitLastSphereOf1000(t *testing.T) {
	ithSphereOfNPrimitiveListTest(1000, 1000, true, t)
}

func BenchmarkPrimitiveListIntersectionHitLastSphereOf1000(b *testing.B) {
	ithSphereOfNPrimitiveListBenchmark(1000, 1000, true, b)
}

func TestPrimitiveListIntersectionMissSphereOf1000(t *testing.T) {
	ithSphereOfNPrimitiveListTest(1, 1000, false, t)
}

func BenchmarkPrimitiveListIntersectionMissSphereOf1000(b *testing.B) {
	ithSphereOfNPrimitiveListBenchmark(1, 1000, false, b)
}

func TestPrimitiveListIntersectionHitFirstTriangleOf10(t *testing.T) {
	ithTriangleOfNPrimitiveListTest(1, 10, true, t)
}

func BenchmarkPrimitiveListIntersectionHitFirstTriangleOf10(b *testing.B) {
	ithTriangleOfNPrimitiveListBenchmark(1, 10, true, b)
}

func TestPrimitiveListIntersectionHitLastTriangleOf10(t *testing.T) {
	ithTriangleOfNPrimitiveListTest(10, 10, true, t)
}

func BenchmarkPrimitiveListIntersectionHitLastTriangleOf10(b *testing.B) {
	ithTriangleOfNPrimitiveListBenchmark(10, 10, true, b)
}

func TestPrimitiveListIntersectionMissTriangleOf10(t *testing.T) {
	ithTriangleOfNPrimitiveListTest(1, 10, false, t)
}

func BenchmarkPrimitiveListIntersectionMissTriangleOf10(b *testing.B) {
	ithTriangleOfNPrimitiveListBenchmark(1, 10, false, b)
}

func TestPrimitiveListIntersectionHitFirstTriangleOf1000(t *testing.T) {
	ithTriangleOfNPrimitiveListTest(1, 1000, true, t)
}

func BenchmarkPrimitiveListIntersectionHitFirstTriangleOf1000(b *testing.B) {
	ithTriangleOfNPrimitiveListBenchmark(1, 1000, true, b)
}

func TestPrimitiveListIntersectionHitLastTriangleOf1000(t *testing.T) {
	ithTriangleOfNPrimitiveListTest(1000, 1000, true, t)
}

func BenchmarkPrimitiveListIntersectionHitLastTriangleOf1000(b *testing.B) {
	ithTriangleOfNPrimitiveListBenchmark(1000, 1000, true, b)
}

func TestPrimitiveListIntersectionMissTriangleOf1000(t *testing.T) {
	ithTriangleOfNPrimitiveListTest(1, 1000, false, t)
}

func BenchmarkPrimitiveListIntersectionMissTriangleOf1000(b *testing.B) {
	ithTriangleOfNPrimitiveListBenchmark(1, 1000, false, b)
}

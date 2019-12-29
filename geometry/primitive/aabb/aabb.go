package aabb

import (
	"fluorescence/geometry"
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

// AABB represents an Axis-Aligned Bounding Box
type AABB struct {
	A mgl64.Vec3
	B mgl64.Vec3
}

// SurroundingBox represents an encompassing box for two smaller AABBs
func SurroundingBox(aabb1, aabb2 *AABB) *AABB {
	return &AABB{
		A: mgl64.Vec3{
			math.Min(aabb1.A.X(), aabb2.A.X()),
			math.Min(aabb1.A.Y(), aabb2.A.Y()),
			math.Min(aabb1.A.Z(), aabb2.A.Z()),
		},
		B: mgl64.Vec3{
			math.Max(aabb1.B.X(), aabb2.B.X()),
			math.Max(aabb1.B.Y(), aabb2.B.Y()),
			math.Max(aabb1.B.Z(), aabb2.B.Z()),
		},
	}
}

// func (aabb *AABB) Intersection(ray geometry.Ray, t0, t1 float64) bool {
// 	return aabb.IntersectionNew(ray, t0, t1)
// 	// return aabb.IntersectionClassic(ray, t0, t1)
// }

// Intersection computer the intersection of this object and a given ray if it exists
func (aabb *AABB) Intersection(ray geometry.Ray, t0, t1 float64) bool {
	var ti0, ti1 float64

	tMin := t0
	tMax := t1

	for i := 0; i < 3; i++ {
		inverseDirection := 1.0 / ray.Direction[i]
		if inverseDirection < 0.0 {
			// swap
			ti0 = (aabb.B[i] - ray.Origin[i]) * inverseDirection
			ti1 = (aabb.A[i] - ray.Origin[i]) * inverseDirection
		} else {
			ti0 = (aabb.A[i] - ray.Origin[i]) * inverseDirection
			ti1 = (aabb.B[i] - ray.Origin[i]) * inverseDirection
		}
		if ti0 > tMin {
			tMin = ti0
		}
		if ti1 < tMax {
			tMax = ti1
		}
		if tMax <= tMin {
			return false
		}

	}

	// must be a hit!
	return true
}

func (aabb *AABB) intersectionClassic(ray geometry.Ray, t0, t1 float64) bool {
	tMin := t0
	tMax := t1

	for i := 0; i < 3; i++ {
		inverseDirection := 1.0 / ray.Direction[i]
		ti0 := (aabb.A[i] - ray.Origin[i]) * inverseDirection
		ti1 := (aabb.B[i] - ray.Origin[i]) * inverseDirection
		if inverseDirection < 0.0 {
			// swap
			ti0, ti1 = ti1, ti0
		}
		if ti0 > tMin {
			tMin = ti0
		}
		if ti1 < tMax {
			tMax = ti1
		}
		if tMax <= tMin {
			return false
		}
	}

	// must be a hit!
	return true
}

package primitive

import (
	"fluorescence/geometry"
	"math"
)

type AABB struct {
	A *geometry.Point
	B *geometry.Point
}

func SurroundingBox(aabb1, aabb2 *AABB) *AABB {
	return &AABB{
		A: &geometry.Point{
			X: math.Min(aabb1.A.X, aabb2.A.X),
			Y: math.Min(aabb1.A.Y, aabb2.A.Y),
			Z: math.Min(aabb1.A.Z, aabb2.A.Z),
		},
		B: &geometry.Point{
			X: math.Max(aabb1.B.X, aabb2.B.X),
			Y: math.Max(aabb1.B.Y, aabb2.B.Y),
			Z: math.Max(aabb1.B.Z, aabb2.B.Z),
		},
	}
}

func (aabb *AABB) Intersection(ray *geometry.Ray, t0, t1 float64) bool {
	tMin := t0
	tMax := t1
	// compute X
	inverseDirectionX := 1.0 / ray.Direction.X
	tx0 := (aabb.A.X - ray.Origin.X) * inverseDirectionX
	tx1 := (aabb.B.X - ray.Origin.X) * inverseDirectionX
	if inverseDirectionX < 0.0 {
		// swap
		temp := tx0
		tx0 = tx1
		tx1 = temp
	}
	if tx0 > tMin {
		tMin = tx0
	}
	if tx1 < tMax {
		tMax = tx1
	}
	if tMax <= tMin {
		return false
	}

	// compute Y
	inverseDirectionY := 1.0 / ray.Direction.Y
	ty0 := (aabb.A.Y - ray.Origin.Y) * inverseDirectionY
	ty1 := (aabb.B.Y - ray.Origin.Y) * inverseDirectionY
	if inverseDirectionY < 0.0 {
		// swap
		temp := ty0
		ty0 = ty1
		ty1 = temp
	}
	if ty0 > tMin {
		tMin = ty0
	}
	if ty1 < tMax {
		tMax = ty1
	}
	if tMax <= tMin {
		return false
	}

	// compute Z
	inverseDirectionZ := 1.0 / ray.Direction.Z
	tz0 := (aabb.A.Z - ray.Origin.Z) * inverseDirectionZ
	tz1 := (aabb.B.Z - ray.Origin.Z) * inverseDirectionZ
	if inverseDirectionZ < 0.0 {
		// swap
		temp := tz0
		tz0 = tz1
		tz1 = temp
	}
	if tz0 > tMin {
		tMin = tz0
	}
	if tz1 < tMax {
		tMax = tz1
	}
	if tMax <= tMin {
		return false
	}

	// must be a hit!
	return true
}

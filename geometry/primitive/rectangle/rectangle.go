package rectangle

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"fmt"
)

type rectangle struct {
	axisAlignedRectangle primitive.Primitive
}

type RectangleData struct {
	A                 geometry.Point `json:"a"`
	B                 geometry.Point `json"b"`
	IsCulled          bool           `json:"is_culled"`
	HasNegativeNormal bool           `json:"has_negative_normal"`
}

func NewRectangle(rd *RectangleData) (*rectangle, error) {
	// if rd.A == nil || rd.B == nil {
	// 	return nil, fmt.Errorf("Rectangle a or b is nil")
	// }
	if (rd.A.X == rd.B.X && rd.A.Y == rd.B.Y) ||
		(rd.A.X == rd.B.X && rd.A.Z == rd.B.Z) ||
		(rd.A.Y == rd.B.Y && rd.A.Z == rd.B.Z) {
		return nil, fmt.Errorf("Rectangle resolves to line or point")
	}

	if rd.A.X == rd.B.X {
		// lies on YZ plane
		return &rectangle{newYZRectangle(rd.A, rd.B, rd.IsCulled, rd.HasNegativeNormal)}, nil
	} else if rd.A.Y == rd.B.Y {
		// lies on XZ Plane
		return &rectangle{newXZRectangle(rd.A, rd.B, rd.IsCulled, rd.HasNegativeNormal)}, nil
	} else if rd.A.Z == rd.B.Z {
		// lies on XY Plane
		return &rectangle{newXYRectangle(rd.A, rd.B, rd.IsCulled, rd.HasNegativeNormal)}, nil
	}
	return nil, fmt.Errorf("Points do not lie on on axis-aligned plane")
}

func (r *rectangle) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	return r.axisAlignedRectangle.Intersection(ray, tMin, tMax)
}

func (r *rectangle) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return r.axisAlignedRectangle.BoundingBox(t0, t1)
}

func (r *rectangle) SetMaterial(m material.Material) {
	r.axisAlignedRectangle.SetMaterial(m)
}

func (r *rectangle) IsInfinite() bool {
	return r.axisAlignedRectangle.IsInfinite()
}

func (r *rectangle) IsClosed() bool {
	return r.axisAlignedRectangle.IsClosed()
}

func (r *rectangle) Copy() primitive.Primitive {
	newR := *r
	return &newR
}

func BasicRectangle(xOffset, yOffset, zOffset float64) *rectangle {
	rd := RectangleData{
		A: geometry.Point{
			X: 0.0 + xOffset,
			Y: 0.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		B: geometry.Point{
			X: 1.0 + xOffset,
			Y: 1.0 + yOffset,
			Z: 0.0 + zOffset,
		},
	}
	r, _ := NewRectangle(&rd)
	return r
}

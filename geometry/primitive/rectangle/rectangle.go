package rectangle

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"fmt"
)

type Rectangle struct {
	axisAlignedRectangle primitive.Primitive
}

type RectangleData struct {
	A                 *geometry.Point `json:"a"`
	B                 *geometry.Point `json"b"`
	IsCulled          bool            `json:"is_culled"`
	HasNegativeNormal bool            `json:"has_negative_normal"`
}

func NewRectangle(data *RectangleData) (*Rectangle, error) {
	if data.A.X == data.B.X {
		// lies on YZ plane
		return &Rectangle{newYZRectangle(data.A, data.B, data.IsCulled, data.HasNegativeNormal)}, nil
	} else if data.A.Y == data.B.Y {
		// lies on XZ Plane
		return &Rectangle{newXZRectangle(data.A, data.B, data.IsCulled, data.HasNegativeNormal)}, nil
	} else if data.A.Z == data.B.Z {
		// lies on XY Plane
		return &Rectangle{newXYRectangle(data.A, data.B, data.IsCulled, data.HasNegativeNormal)}, nil
	}
	return nil, fmt.Errorf("Points do not lie on on axis-aligned plane")
}

func (r *Rectangle) Intersection(ray *geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	return r.axisAlignedRectangle.Intersection(ray, tMin, tMax)
}

func (r *Rectangle) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return r.axisAlignedRectangle.BoundingBox(t0, t1)
}

func (r *Rectangle) SetMaterial(m material.Material) {
	r.axisAlignedRectangle.SetMaterial(m)
}

func (r *Rectangle) Copy() primitive.Primitive {
	newR := *r
	return &newR
}

func BasicRectangle(xOffset, yOffset, zOffset float64) *Rectangle {
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

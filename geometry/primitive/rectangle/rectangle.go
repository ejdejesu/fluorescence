package rectangle

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"fmt"
)

// Rectangle represents a Axis-Aligned rectangle geometry object
type Rectangle struct {
	A                    geometry.Point `json:"a"`
	B                    geometry.Point `json:"b"`
	IsCulled             bool           `json:"is_culled"`
	HasNegativeNormal    bool           `json:"has_negative_normal"`
	axisAlignedRectangle primitive.Primitive
}

// Setup sets up internal fields in a rectangle
func (r *Rectangle) Setup() (*Rectangle, error) {
	// if r.A == nil || r.B == nil {
	// 	return nil, fmt.Errorf("rectangle a or b is nil")
	// }
	if (r.A.X == r.B.X && r.A.Y == r.B.Y) ||
		(r.A.X == r.B.X && r.A.Z == r.B.Z) ||
		(r.A.Y == r.B.Y && r.A.Z == r.B.Z) {
		return nil, fmt.Errorf("rectangle resolves to line or point")
	}

	if r.A.X == r.B.X {
		// lies on YZ plane
		r.axisAlignedRectangle = newYZRectangle(r.A, r.B, r.IsCulled, r.HasNegativeNormal)
		return r, nil
	} else if r.A.Y == r.B.Y {
		// lies on XZ Plane
		r.axisAlignedRectangle = newXZRectangle(r.A, r.B, r.IsCulled, r.HasNegativeNormal)
		return r, nil
	} else if r.A.Z == r.B.Z {
		// lies on XY Plane
		r.axisAlignedRectangle = newXYRectangle(r.A, r.B, r.IsCulled, r.HasNegativeNormal)
		return r, nil
	}
	return nil, fmt.Errorf("points do not lie on on axis-aligned plane")
}

// Intersection computer the intersection of this object and a given ray if it exists
func (r *Rectangle) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	return r.axisAlignedRectangle.Intersection(ray, tMin, tMax)
}

// BoundingBox return an AABB of this object
func (r *Rectangle) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return r.axisAlignedRectangle.BoundingBox(t0, t1)
}

// SetMaterial sets this object's material
func (r *Rectangle) SetMaterial(m material.Material) {
	r.axisAlignedRectangle.SetMaterial(m)
}

// IsInfinite return whether this object is infinite
func (r *Rectangle) IsInfinite() bool {
	return r.axisAlignedRectangle.IsInfinite()
}

// IsClosed returns whether this object is closed
func (r *Rectangle) IsClosed() bool {
	return r.axisAlignedRectangle.IsClosed()
}

// Copy returns a shallow copy of this object
func (r *Rectangle) Copy() primitive.Primitive {
	newR := *r
	return &newR
}

// Unit return a unit rectangle
func Unit(xOffset, yOffset, zOffset float64) *Rectangle {
	r, _ := (&Rectangle{
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
	}).Setup()
	return r
}

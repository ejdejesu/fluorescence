package translate

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
)

// Translation is a primitive with a translation attached
type Translation struct {
	Displacement geometry.Vector `json:"displacement"`
	TypeName     string          `json:"type"`
	Data         interface{}     `json:"data"`
	Primitive    primitive.Primitive
}

// Setup sets up a Translation's internal fields
func (t *Translation) Setup() (*Translation, error) {
	return t, nil
}

// Intersection computer the intersection of this object and a given ray if it exists
func (t *Translation) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {

	// translate the ray to the object
	ray.Origin = ray.Origin.SubVector(t.Displacement)

	rh, ok := t.Primitive.Intersection(ray, tMin, tMax)
	if ok {
		rh.Ray.Origin = rh.Ray.Origin.AddVector(t.Displacement)
	}
	return rh, ok
}

// BoundingBox returns an AABB for this object
func (t *Translation) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	box, ok := t.Primitive.BoundingBox(t0, t1)
	if ok {
		box = &aabb.AABB{
			A: box.A.AddVector(t.Displacement),
			B: box.B.AddVector(t.Displacement),
		}
	}
	return box, ok
}

// SetMaterial sets the material of this object
func (t *Translation) SetMaterial(m material.Material) {
	t.Primitive.SetMaterial(m)
}

// IsInfinite returns whether this object is infinite
func (t *Translation) IsInfinite() bool {
	return t.Primitive.IsInfinite()
}

// IsClosed returns whether this object is closed
func (t *Translation) IsClosed() bool {
	return t.Primitive.IsClosed()
}

// Copy returns a shallow copy of this object
func (t *Translation) Copy() primitive.Primitive {
	newT := *t
	return &newT
}

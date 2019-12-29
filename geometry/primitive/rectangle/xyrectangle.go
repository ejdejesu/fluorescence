package rectangle

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

type xyRectangle struct {
	x0       float64
	x1       float64
	y0       float64
	y1       float64
	z        float64
	isCulled bool
	normal   mgl64.Vec3
	mat      material.Material
}

func newXYRectangle(a, b mgl64.Vec3, isCulled, hasNegativeNormal bool) *xyRectangle {
	x0 := math.Min(a.X(), b.X())
	x1 := math.Max(a.X(), b.X())
	y0 := math.Min(a.Y(), b.Y())
	y1 := math.Max(a.Y(), b.Y())

	z := a.Z()

	var normal mgl64.Vec3
	if hasNegativeNormal {
		normal = mgl64.Vec3{
			0.0,
			0.0,
			-1.0,
		}
	} else {
		normal = mgl64.Vec3{
			0.0,
			0.0,
			1.0,
		}
	}
	return &xyRectangle{
		x0:       x0,
		x1:       x1,
		y0:       y0,
		y1:       y1,
		z:        z,
		isCulled: isCulled,
		normal:   normal,
	}
}

// Intersection computer the intersection of this object and a given ray if it exists
func (r *xyRectangle) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	// Ray is coming from behind rectangle
	denominator := ray.Direction.Dot(r.normal)
	if r.isCulled && denominator > -1e-7 {
		return nil, false
	} else if denominator < 1e-7 && denominator > -1e-7 {
		return nil, false
	}

	// Ray is parallel to plane
	if ray.Direction.Z() == 0 {
		return nil, false
	}

	t := (r.z - ray.Origin.Z()) / ray.Direction.Z()

	if t < tMin || t > tMax {
		return nil, false
	}

	x := ray.Origin.X() + (t * ray.Direction.X())
	y := ray.Origin.Y() + (t * ray.Direction.Y())

	// plane intersection not within rectangle
	if x < r.x0 || x > r.x1 || y < r.y0 || y > r.y1 {
		return nil, false
	}

	u := (x - r.x0) / (r.x1 - r.x0)
	v := (y - r.y0) / (r.y1 - r.y0)

	return &material.RayHit{
		Ray:         ray,
		NormalAtHit: r.normal,
		Time:        t,
		U:           u,
		V:           v,
		Material:    r.mat,
	}, true
}

// BoundingBox return an AABB of this object
func (r *xyRectangle) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return &aabb.AABB{
		A: mgl64.Vec3{
			r.x0 - 1e-7,
			r.y0 - 1e-7,
			r.z - 1e-7,
		},
		B: mgl64.Vec3{
			r.x1 + 1e-7,
			r.y1 + 1e-7,
			r.z + 1e-7,
		},
	}, true
}

// SetMaterial sets this object's material
func (r *xyRectangle) SetMaterial(m material.Material) {
	r.mat = m
}

// IsInfinite return whether this object is infinite
func (r *xyRectangle) IsInfinite() bool {
	return false
}

// IsClosed returns whether this object is closed
func (r *xyRectangle) IsClosed() bool {
	return false
}

// Copy returns a shallow copy of this object
func (r *xyRectangle) Copy() primitive.Primitive {
	newR := *r
	return &newR
}

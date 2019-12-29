package rectangle

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

type xzRectangle struct {
	x0       float64
	x1       float64
	z0       float64
	z1       float64
	y        float64
	isCulled bool
	normal   mgl64.Vec3
	mat      material.Material
}

func newXZRectangle(a, b mgl64.Vec3, isCulled, hasNegativeNormal bool) *xzRectangle {
	x0 := math.Min(a.X(), b.X())
	x1 := math.Max(a.X(), b.X())
	z0 := math.Min(a.Z(), b.Z())
	z1 := math.Max(a.Z(), b.Z())

	y := a.Y()

	var normal mgl64.Vec3
	if hasNegativeNormal {
		normal = mgl64.Vec3{
			0.0,
			-1.0,
			0.0,
		}
	} else {
		normal = mgl64.Vec3{
			0.0,
			1.0,
			0.0,
		}
	}
	return &xzRectangle{
		x0:       x0,
		x1:       x1,
		z0:       z0,
		z1:       z1,
		y:        y,
		isCulled: isCulled,
		normal:   normal,
	}
}

// Intersection computer the intersection of this object and a given ray if it exists
func (r *xzRectangle) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	// Ray is coming from behind rectangle
	denominator := ray.Direction.Dot(r.normal)
	if r.isCulled && denominator > -1e-7 {
		return nil, false
	} else if denominator < 1e-7 && denominator > -1e-7 {
		return nil, false
	}

	// Ray is parallel to plane
	if ray.Direction.Y() == 0 {
		return nil, false
	}

	t := (r.y - ray.Origin.Y()) / ray.Direction.Y()

	if t < tMin || t > tMax {
		return nil, false
	}

	x := ray.Origin.X() + (t * ray.Direction.X())
	z := ray.Origin.Z() + (t * ray.Direction.Z())

	// plane intersection not within rectangle
	if x < r.x0 || x > r.x1 || z < r.z0 || z > r.z1 {
		return nil, false
	}

	u := (x - r.x0) / (r.x1 - r.x0)
	v := (z - r.z0) / (r.z1 - r.z0)

	return &material.RayHit{
		Ray:         ray,
		NormalAtHit: r.normal,
		Time:        t,
		U:           u,
		V:           v,
		Material:    r.mat,
	}, true
}

// BoundingBox returns the AABB of this object
func (r *xzRectangle) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return &aabb.AABB{
		A: mgl64.Vec3{
			r.x0 - 1e-7,
			r.y - 1e-7,
			r.z0 - 1e-7,
		},
		B: mgl64.Vec3{
			r.x1 + 1e-7,
			r.y + 1e-7,
			r.z1 + 1e-7,
		},
	}, true
}

// SetMaterial sets this object's material
func (r *xzRectangle) SetMaterial(m material.Material) {
	r.mat = m
}

// IsInfinite return whether this object is infinite
func (r *xzRectangle) IsInfinite() bool {
	return false
}

// IsClosed returns whether this object is closed
func (r *xzRectangle) IsClosed() bool {
	return false
}

// Copy returns a shallow copy of this object
func (r *xzRectangle) Copy() primitive.Primitive {
	newR := *r
	return &newR
}

package rectangle

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

type yzRectangle struct {
	y0       float64
	y1       float64
	z0       float64
	z1       float64
	x        float64
	isCulled bool
	normal   mgl64.Vec3
	mat      material.Material
}

func newYZRectangle(a, b mgl64.Vec3, isCulled, hasNegativeNormal bool) *yzRectangle {
	y0 := math.Min(a.Y(), b.Y())
	y1 := math.Max(a.Y(), b.Y())
	z0 := math.Min(a.Z(), b.Z())
	z1 := math.Max(a.Z(), b.Z())

	x := a.X()

	var normal mgl64.Vec3
	if hasNegativeNormal {
		normal = mgl64.Vec3{
			-1.0,
			0.0,
			0.0,
		}
	} else {
		normal = mgl64.Vec3{
			1.0,
			0.0,
			0.0,
		}
	}
	return &yzRectangle{
		y0:       y0,
		y1:       y1,
		z0:       z0,
		z1:       z1,
		x:        x,
		isCulled: isCulled,
		normal:   normal,
	}
}

// Intersection computer the intersection of this object and a given ray if it exists
func (r *yzRectangle) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	// Ray is coming from behind rectangle
	denominator := ray.Direction.Dot(r.normal)
	if r.isCulled && denominator > -1e-7 {
		return nil, false
	} else if denominator < 1e-7 && denominator > -1e-7 {
		return nil, false
	}

	// Ray is parallel to plane
	if ray.Direction.X() == 0 {
		return nil, false
	}

	t := (r.x - ray.Origin.X()) / ray.Direction.X()

	if t < tMin || t > tMax {
		return nil, false
	}

	y := ray.Origin.Y() + (t * ray.Direction.Y())
	z := ray.Origin.Z() + (t * ray.Direction.Z())

	// plane intersection not within rectangle
	if y < r.y0 || y > r.y1 || z < r.z0 || z > r.z1 {
		return nil, false
	}

	u := (z - r.z0) / (r.z1 - r.z0)
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

// BoundingBox returns an AABB for this object
func (r *yzRectangle) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return &aabb.AABB{
		A: mgl64.Vec3{
			r.x - 1e-7,
			r.y0 - 1e-7,
			r.z0 - 1e-7,
		},
		B: mgl64.Vec3{
			r.x + 1e-7,
			r.y1 + 1e-7,
			r.z1 + 1e-7,
		},
	}, true
}

// SetMaterial sets this object's material
func (r *yzRectangle) SetMaterial(m material.Material) {
	r.mat = m
}

// IsInfinite return whether this object is infinite
func (r *yzRectangle) IsInfinite() bool {
	return false
}

// IsClosed returns whether this object is closed
func (r *yzRectangle) IsClosed() bool {
	return false
}

// Copy returns a shallow copy of this object
func (r *yzRectangle) Copy() primitive.Primitive {
	newR := *r
	return &newR
}

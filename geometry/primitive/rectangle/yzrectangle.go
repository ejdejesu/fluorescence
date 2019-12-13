package rectangle

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"math"
)

type yzRectangle struct {
	y0       float64
	y1       float64
	z0       float64
	z1       float64
	x        float64
	isCulled bool
	normal   *geometry.Vector
	Material material.Material
}

func newYZRectangle(a, b *geometry.Point, isCulled, hasNegativeNormal bool) *yzRectangle {
	y0 := math.Min(a.Y, b.Y)
	y1 := math.Max(a.Y, b.Y)
	z0 := math.Min(a.Z, b.Z)
	z1 := math.Max(a.Z, b.Z)

	x := a.X

	var normal *geometry.Vector
	if hasNegativeNormal {
		normal = &geometry.Vector{-1.0, 0.0, 0.0}
	} else {
		normal = &geometry.Vector{1.0, 0.0, 0.0}
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

func (r *yzRectangle) Intersection(ray *geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	// Ray is coming from behind rectangle
	if r.isCulled && (ray.Direction.Dot(r.normal)) > 0 {
		return nil, false
	}

	// Ray is parallel to plane
	if ray.Direction.X == 0 {
		return nil, false
	}

	t := (r.x - ray.Origin.X) / ray.Direction.X

	if t < tMin || t > tMax {
		return nil, false
	}

	y := ray.Origin.Y + (t * ray.Direction.Y)
	z := ray.Origin.Z + (t * ray.Direction.Z)

	// plane intersection not within rectangle
	if y < r.y0 || y > r.y1 || z < r.z0 || z > r.z1 {
		return nil, false
	}

	return &material.RayHit{ray, r.normal, t, r.Material}, true
}

func (r *yzRectangle) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return &aabb.AABB{
		A: &geometry.Point{
			X: r.x - 1e-7,
			Y: r.y0 - 1e-7,
			Z: r.z0 - 1e-7,
		},
		B: &geometry.Point{
			X: r.x + 1e-7,
			Y: r.y1 + 1e-7,
			Z: r.z1 + 1e-7,
		},
	}, true
}

func (r *yzRectangle) SetMaterial(m material.Material) {
	r.Material = m
}

func (r *yzRectangle) Copy() primitive.Primitive {
	newR := *r
	return &newR
}

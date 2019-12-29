package plane

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"fmt"

	"github.com/go-gl/mathgl/mgl64"
)

// Plane represents an infinite plane
type Plane struct {
	Point    mgl64.Vec3 `json:"point"`
	Normal   mgl64.Vec3 `json:"normal"`
	IsCulled bool       `json:"is_culled"`
	mat      material.Material
}

// Setup sets up this Plane's internal fields
func (p *Plane) Setup() (*Plane, error) {
	if p.Normal.Len() == 0.0 {
		return nil, fmt.Errorf("Plane normal is zero vector")
	}
	p.Normal = p.Normal.Normalize()
	return p, nil
}

// Intersection computer the intersection of this object and a given ray if it exists
func (p *Plane) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	denominator := ray.Direction.Dot(p.Normal)
	if p.IsCulled && denominator > -1e-7 {
		return nil, false
	} else if denominator < 1e-7 && denominator > -1e-7 {
		return nil, false
	}
	PlaneVector := p.Point.Sub(ray.Origin)
	t := PlaneVector.Dot(p.Normal) / denominator

	if t < tMin || t > tMax {
		return nil, false
	}

	return &material.RayHit{
		Ray:         ray,
		NormalAtHit: p.Normal,
		Time:        t,
		Material:    p.mat,
	}, true
}

// BoundingBox return an AABB for this object
func (p *Plane) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return nil, false
}

// SetMaterial sets this object's material
func (p *Plane) SetMaterial(m material.Material) {
	p.mat = m
}

// IsInfinite returns whether this object is infinite
func (p *Plane) IsInfinite() bool {
	return true
}

// IsClosed returns whether this object is closed
func (p *Plane) IsClosed() bool {
	return false
}

// Copy returns a shallow copy of this object
func (p *Plane) Copy() primitive.Primitive {
	newP := *p
	return &newP
}

// Unit returns a unit plane
func Unit(xOffset, yOffset, zOffset float64) *Plane {
	p, _ := (&Plane{
		Point: mgl64.Vec3{
			0.0 + xOffset,
			0.0 + yOffset,
			0.0 + zOffset,
		},
		Normal: mgl64.Vec3{
			0.0 + xOffset,
			0.0 + yOffset,
			-1.0 + zOffset,
		},
	}).Setup()
	return p
}

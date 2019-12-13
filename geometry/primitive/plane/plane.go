package plane

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"fmt"
)

type plane struct {
	point    *geometry.Point
	normal   *geometry.Vector
	isCulled bool
	Material material.Material
}

type PlaneData struct {
	Point    *geometry.Point  `json:"point"`
	Normal   *geometry.Vector `json:"normal"`
	IsCulled bool             `json:"is_culled"`
}

func NewPlane(pd *PlaneData) (*plane, error) {
	if pd.Point == nil || pd.Normal == nil {
		return nil, fmt.Errorf("Plane point or normal is nil")
	}
	if pd.Normal.Magnitude() == 0.0 {
		return nil, fmt.Errorf("Plane normal is zero vector")
	}
	return &plane{
		point:    pd.Point,
		normal:   pd.Normal.Unit(),
		isCulled: pd.IsCulled,
	}, nil
}

func EmptyPlane() *plane {
	return &plane{}
}

func (p *plane) Intersection(ray *geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	denominator := ray.Direction.Dot(p.normal)
	if p.isCulled && denominator > -1e-7 {
		return nil, false
	} else if denominator < 1e-7 && denominator > -1e-7 {
		return nil, false
	}
	planeVector := ray.Origin.To(p.point)
	t := planeVector.Dot(p.normal) / denominator

	if t < tMin || t > tMax {
		return nil, false
	}

	return &material.RayHit{
		Ray:         ray,
		NormalAtHit: p.normal,
		T:           t,
		Material:    p.Material,
	}, true
}

func (p *plane) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return nil, false
}

func (p *plane) SetMaterial(m material.Material) {
	p.Material = m
}

func (p *plane) IsInfinite() bool {
	return true
}

func (p *plane) IsClosed() bool {
	return false
}

func (p *plane) Copy() primitive.Primitive {
	newP := *p
	return &newP
}

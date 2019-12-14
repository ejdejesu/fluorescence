package cylinder

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"fmt"
	"math"
)

type infiniteCylinder struct {
	ray                *geometry.Ray
	radius             float64
	hasInvertedNormals bool
	mat                material.Material
}

type InfiniteCylinderData struct {
	Ray                *geometry.Ray `json:"ray"`
	Radius             float64       `json:"radius"`
	HasInvertedNormals bool          `json:"has_inverted_normals"`
}

func NewInfiniteCylinder(icd *InfiniteCylinderData) (*infiniteCylinder, error) {
	if icd.Ray == nil {
		return nil, fmt.Errorf("InfiniteCylinder ray is nil")
	}
	if icd.Ray.Origin == nil || icd.Ray.Direction == nil {
		return nil, fmt.Errorf("InfiniteCylinder ray origin or ray direction is nil")
	}
	if icd.Ray.Direction.Magnitude() == 0 {
		return nil, fmt.Errorf("InfiniteCylinder ray direction is zero vector")
	}
	if icd.Radius <= 0.0 {
		return nil, fmt.Errorf("InfiniteCylinder radius is 0 or negative")
	}
	icd.Ray.Direction.UnitInPlace()
	return &infiniteCylinder{
		ray:                icd.Ray,
		radius:             icd.Radius,
		hasInvertedNormals: icd.HasInvertedNormals,
	}, nil
}

func (ic *infiniteCylinder) Intersection(ray *geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	deltaP := ic.ray.Origin.To(ray.Origin)
	preA := ray.Direction.Sub(ic.ray.Direction.MultScalar(ray.Direction.Dot(ic.ray.Direction)))
	preB := deltaP.Sub(ic.ray.Direction.MultScalar(deltaP.Dot(ic.ray.Direction)))

	// terms of the quadratic equation we are solving
	a := preA.Dot(preA)
	b := preA.Dot(preB)
	c := preB.Dot(preB) - (ic.radius * ic.radius)

	preDiscriminant := b*b - a*c

	if preDiscriminant > 0 {
		root := math.Sqrt(preDiscriminant)
		// evaluate first solution, which will be smaller
		t1 := (-b - root) / a
		// return if within range
		if t1 >= tMin && t1 <= tMax {
			return &material.RayHit{
				Ray:         ray,
				NormalAtHit: ic.normalAt(ray.PointAt(t1)),
				T:           t1,
				Material:    ic.mat,
			}, true
		}
		// evaluate and return second solution if in range
		t2 := (-b + root) / a
		if t2 >= tMin && t2 <= tMax {
			return &material.RayHit{
				Ray:         ray,
				NormalAtHit: ic.normalAt(ray.PointAt(t2)),
				T:           t2,
				Material:    ic.mat,
			}, true
		}
	}

	return nil, false
}

func (ic *infiniteCylinder) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return nil, false
}

func (ic *infiniteCylinder) SetMaterial(m material.Material) {
	ic.mat = m
}

func (ic *infiniteCylinder) IsInfinite() bool {
	return true
}

func (ic *infiniteCylinder) IsClosed() bool {
	return true
}

func (ic *infiniteCylinder) Copy() primitive.Primitive {
	newIC := *ic
	return &newIC
}

func (ic *infiniteCylinder) normalAt(p *geometry.Point) *geometry.Vector {
	if ic.hasInvertedNormals {
		return geometry.ZERO.Sub(ic.ray.ClosestPoint(p).To(p).UnitInPlace())
	}
	return ic.ray.ClosestPoint(p).To(p).UnitInPlace()
}

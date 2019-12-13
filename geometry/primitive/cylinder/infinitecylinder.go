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
	point              *geometry.Point
	direction          *geometry.Vector
	radius             float64
	isCulled           bool
	hasInvertedNormals bool
	mat                material.Material
}

type InfiniteCylinderData struct {
	Point              *geometry.Point  `json:"point"`
	Direction          *geometry.Vector `json:"direction"`
	Radius             float64          `json:"radius"`
	IsCulled           bool             `json:"is_culled"`
	HasInvertedNormals bool             `json:"has_inverted_normals"`
}

func NewInfiniteCylinder(icd *InfiniteCylinderData) (*infiniteCylinder, error) {
	if icd.Point == nil || icd.Direction == nil {
		return nil, fmt.Errorf("InfiniteCylinder point or direction is nil")
	}
	if icd.Direction.Magnitude() == 0 {
		return nil, fmt.Errorf("InfiniteCylinder direction is zero vector")
	}
	if icd.Radius <= 0.0 {
		return nil, fmt.Errorf("InfiniteCylinder radius is 0 or negative")
	}
	return &infiniteCylinder{
		point:     icd.Point,
		direction: icd.Direction.Unit(),
		radius:    icd.Radius,
		isCulled:  icd.IsCulled,
	}, nil
}

func (ic *infiniteCylinder) Intersection(ray *geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	deltaP := ic.point.To(ray.Origin)
	preA := ray.Direction.Sub(ic.direction.MultScalar(ray.Direction.Dot(ic.direction)))
	preB := deltaP.Sub(ic.direction.MultScalar(deltaP.Dot(ic.direction)))

	// terms of the quadratic equation we are solving
	a := preA.Dot(preA)
	b := preA.Dot(preB)
	c := preB.Dot(preB) - (ic.radius * ic.radius)

	preDiscriminant := b*b - a*c

	if preDiscriminant > 0 {
		// evaluate first solution, which will be smaller
		t1 := (-b - math.Sqrt(preDiscriminant)) / a
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
		t2 := (-b + math.Sqrt(preDiscriminant)) / a
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
	v := ic.point.To(p)
	t := v.Dot(ic.direction)
	cylinderPoint := ic.point.AddVector(ic.direction.MultScalar(t))
	return cylinderPoint.To(p).UnitInPlace()
}

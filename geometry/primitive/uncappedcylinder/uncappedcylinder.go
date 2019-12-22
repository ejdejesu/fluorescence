package uncappedcylinder

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/geometry/primitive/disk"
	"fluorescence/shading/material"
	"fmt"
	"math"
)

type uncappedCylinder struct {
	ray                geometry.Ray
	minT               float64
	maxT               float64
	radius             float64
	hasInvertedNormals bool
	mat                material.Material
}

type Data struct {
	A                  geometry.Point `json:"a"`
	B                  geometry.Point `json:"b"`
	Radius             float64        `json:"radius"`
	HasInvertedNormals bool           `json:"has_inverted_normals"`
}

func New(ucd *Data) (*uncappedCylinder, error) {
	// if ucd.A == nil || ucd.B == nil {
	// 	return nil, fmt.Errorf("uncappedCylinder ray is nil")
	// }
	if ucd.A.To(ucd.B).Magnitude() == 0 {
		return nil, fmt.Errorf("uncappedCylinder length is zero vector")
	}
	if ucd.Radius <= 0.0 {
		return nil, fmt.Errorf("uncappedCylinder radius is 0 or negative")
	}
	r := geometry.Ray{
		Origin:    ucd.A,
		Direction: ucd.A.To(ucd.B).Unit(),
	}
	minT := 0.0
	maxT := r.ClosestTime(ucd.B)
	return &uncappedCylinder{
		ray:                r,
		minT:               minT,
		maxT:               maxT,
		radius:             ucd.Radius,
		hasInvertedNormals: ucd.HasInvertedNormals,
	}, nil
}

func (uc *uncappedCylinder) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	deltaP := uc.ray.Origin.To(ray.Origin)
	preA := ray.Direction.Sub(uc.ray.Direction.MultScalar(ray.Direction.Dot(uc.ray.Direction)))
	preB := deltaP.Sub(uc.ray.Direction.MultScalar(deltaP.Dot(uc.ray.Direction)))

	// terms of the quadratic equation we are solving
	a := preA.Dot(preA)
	b := preA.Dot(preB)
	c := preB.Dot(preB) - (uc.radius * uc.radius)

	preDiscriminant := b*b - a*c

	if preDiscriminant > 0 {
		root := math.Sqrt(preDiscriminant)
		// evaluate first solution, which will be smaller
		t1 := (-b - root) / a
		cylinderT1 := uc.ray.ClosestTime(ray.PointAt(t1))
		// return if within range
		if t1 >= tMin && t1 <= tMax && cylinderT1 >= uc.minT && cylinderT1 <= uc.maxT {
			return &material.RayHit{
				Ray:         ray,
				NormalAtHit: uc.normalAt(ray.PointAt(t1)),
				Time:        t1,
				Material:    uc.mat,
			}, true
		}
		// evaluate and return second solution if in range
		t2 := (-b + root) / a
		cylinderT2 := uc.ray.ClosestTime(ray.PointAt(t2))
		if t2 >= tMin && t2 <= tMax && cylinderT2 >= uc.minT && cylinderT2 <= uc.maxT {
			return &material.RayHit{
				Ray:         ray,
				NormalAtHit: uc.normalAt(ray.PointAt(t2)),
				Time:        t2,
				Material:    uc.mat,
			}, true
		}
	}

	return nil, false
}

func (uc *uncappedCylinder) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	diskA, _ := disk.New(&disk.Data{
		Center: uc.ray.Origin,
		Normal: uc.ray.Direction,
		Radius: uc.radius,
	})
	diskB, _ := disk.New(&disk.Data{
		Center: uc.ray.PointAt(uc.maxT),
		Normal: uc.ray.PointAt(uc.maxT).To(uc.ray.Origin).Unit(),
		Radius: uc.radius,
	})
	aabbA, aOk := diskA.BoundingBox(0, 0)
	if !aOk {
		return nil, false
	}
	aabbB, bOk := diskB.BoundingBox(0, 0)
	if !bOk {
		return nil, false
	}
	return aabb.SurroundingBox(aabbA, aabbB), true
}

func (uc *uncappedCylinder) SetMaterial(m material.Material) {
	uc.mat = m
}

func (uc *uncappedCylinder) IsInfinite() bool {
	return false
}

func (uc *uncappedCylinder) IsClosed() bool {
	return false
}

func (uc *uncappedCylinder) Copy() primitive.Primitive {
	newUC := *uc
	return &newUC
}

func (uc *uncappedCylinder) normalAt(p geometry.Point) geometry.Vector {
	if uc.hasInvertedNormals {
		return uc.ray.ClosestPoint(p).To(p).Unit().Negate()
	}
	return uc.ray.ClosestPoint(p).To(p).Unit()
}

func UnitUncappedCylinder(xOffset, yOffset, zOffset float64) *uncappedCylinder {
	ucd := Data{
		A: geometry.Point{
			X: 0.0 + xOffset,
			Y: 0.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		B: geometry.Point{
			X: 0.0 + xOffset,
			Y: 1.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		Radius: 1.0,
	}
	uc, _ := New(&ucd)
	return uc
}

package uncappedcylinder

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/geometry/primitive/disk"
	"fluorescence/shading/material"
	"fmt"
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

// UncappedCylinder hi
type UncappedCylinder struct {
	A                  mgl64.Vec3 `json:"a"`
	B                  mgl64.Vec3 `json:"b"`
	Radius             float64    `json:"radius"`
	HasInvertedNormals bool       `json:"has_inverted_normals"`
	ray                geometry.Ray
	minT, maxT         float64
	mat                material.Material
}

// Data holds information needed to contruct a uncappedCylinder
// type Data struct {
// 	A                 mgl64.Vec3 `json:"a"`
// 	B                 mgl64.Vec3 `json:"b"`
// 	Radius             float64        `json:"radius"`
// 	HasInvertedNormals bool           `json:"has_inverted_normals"`
// }

// Setup fills calculated fields in an UncappedCylinder
func (uc *UncappedCylinder) Setup() (*UncappedCylinder, error) {
	// if ucd.A == nil || ucd.B == nil {
	// 	return nil, fmt.Errorf("uncappedCylinder ray is nil")
	// }
	if uc.B.Sub(uc.A).Len() == 0 {
		return nil, fmt.Errorf("uncappedCylinder length is zero vector")
	}
	if uc.Radius <= 0.0 {
		return nil, fmt.Errorf("uncappedCylinder radius is 0 or negative")
	}
	uc.ray = geometry.Ray{
		Origin:    uc.A,
		Direction: uc.B.Sub(uc.A).Normalize(),
	}
	uc.minT = 0.0
	uc.maxT = uc.ray.ClosestTime(uc.B)
	return uc, nil
}

// Intersection computer the intersection of this object and a given ray if it exists
func (uc *UncappedCylinder) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	deltaP := ray.Origin.Sub(uc.ray.Origin)
	preA := ray.Direction.Sub(uc.ray.Direction.Mul(ray.Direction.Dot(uc.ray.Direction)))
	preB := deltaP.Sub(uc.ray.Direction.Mul(deltaP.Dot(uc.ray.Direction)))

	// terms of the quadratic equation we are solving
	a := preA.Dot(preA)
	b := preA.Dot(preB)
	c := preB.Dot(preB) - (uc.Radius * uc.Radius)

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

// BoundingBox finds the AABB of this object
func (uc *UncappedCylinder) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	diskA, _ := (&disk.Disk{
		Center: uc.ray.Origin,
		Normal: uc.ray.Direction,
		Radius: uc.Radius,
	}).Setup()
	diskB, _ := (&disk.Disk{
		Center: uc.ray.PointAt(uc.maxT),
		Normal: uc.ray.Origin.Sub(uc.ray.PointAt(uc.maxT)).Normalize(),
		Radius: uc.Radius,
	}).Setup()
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

// SetMaterial sets this object's material
func (uc *UncappedCylinder) SetMaterial(m material.Material) {
	uc.mat = m
}

// IsInfinite returns whether this in an infinite geometry object
func (uc *UncappedCylinder) IsInfinite() bool {
	return false
}

// IsClosed returns whether this object is closed or not
func (uc *UncappedCylinder) IsClosed() bool {
	return false
}

// Copy returns a shallow copy of this cylinder
func (uc *UncappedCylinder) Copy() primitive.Primitive {
	newUC := *uc
	return &newUC
}

func (uc *UncappedCylinder) normalAt(p mgl64.Vec3) mgl64.Vec3 {
	if uc.HasInvertedNormals {
		return p.Sub(uc.ray.ClosestPoint(p)).Normalize().Mul(-1.0)
	}
	return p.Sub(uc.ray.ClosestPoint(p)).Normalize()
}

// Unit creates a unit uncappedCylinder.
// The points of this cylinder are:
// A: (0, 0, 0),
// B: (1, 0, 0),
// and the Radius is 1
func Unit(xOffset, yOffset, zOffset float64) *UncappedCylinder {
	uc, _ := (&UncappedCylinder{
		A: mgl64.Vec3{
			0.0 + xOffset,
			0.0 + yOffset,
			0.0 + zOffset,
		},
		B: mgl64.Vec3{
			0.0 + xOffset,
			1.0 + yOffset,
			0.0 + zOffset,
		},
		Radius: 1.0,
	}).Setup()
	return uc
}
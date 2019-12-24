package infinitecylinder

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"fmt"
	"math"
)

// InfiniteCylinder represents an infinitely long cylinder
type InfiniteCylinder struct {
	Ray                geometry.Ray `json:"ray"`
	Radius             float64      `json:"radius"`
	HasInvertedNormals bool         `json:"has_inverted_normals"`
	mat                material.Material
}

// type Data struct {
// 	Ray                geometry.Ray
// 	Radius             float64
// 	HasInvertedNormals bool
// }

// Setup sets up an infinite cylinder
func (ic *InfiniteCylinder) Setup() (*InfiniteCylinder, error) {
	// if icd.Ray == nil {
	// 	return nil, fmt.Errorf("infinite cylinder ray is nil")
	// }
	// if ic.Ray.Origin == nil || ic.Ray.Direction == nil {
	// 	return nil, fmt.Errorf("infinite cylinder ray origin or ray direction is nil")
	// }
	if ic.Ray.Direction.Magnitude() == 0 {
		return nil, fmt.Errorf("infinite cylinder ray direction is zero vector")
	}
	if ic.Radius <= 0.0 {
		return nil, fmt.Errorf("infinite cylinder radius is 0 or negative")
	}
	ic.Ray.Direction.Unit()
	return &InfiniteCylinder{
		Ray:                ic.Ray,
		Radius:             ic.Radius,
		HasInvertedNormals: ic.HasInvertedNormals,
	}, nil
}

// Intersection computer the intersection of this object and a given ray if it exists
func (ic *InfiniteCylinder) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	deltaP := ic.Ray.Origin.To(ray.Origin)
	preA := ray.Direction.Sub(ic.Ray.Direction.MultScalar(ray.Direction.Dot(ic.Ray.Direction)))
	preB := deltaP.Sub(ic.Ray.Direction.MultScalar(deltaP.Dot(ic.Ray.Direction)))

	// terms of the quadratic equation we are solving
	a := preA.Dot(preA)
	b := preA.Dot(preB)
	c := preB.Dot(preB) - (ic.Radius * ic.Radius)

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
				Time:        t1,
				Material:    ic.mat,
			}, true
		}
		// evaluate and return second solution if in range
		t2 := (-b + root) / a
		if t2 >= tMin && t2 <= tMax {
			return &material.RayHit{
				Ray:         ray,
				NormalAtHit: ic.normalAt(ray.PointAt(t2)),
				Time:        t2,
				Material:    ic.mat,
			}, true
		}
	}

	return nil, false
}

// BoundingBox returns an AABB of this object
func (ic *InfiniteCylinder) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return nil, false
}

// SetMaterial sets this object's material
func (ic *InfiniteCylinder) SetMaterial(m material.Material) {
	ic.mat = m
}

// IsInfinite returns whether this object is infinite
func (ic *InfiniteCylinder) IsInfinite() bool {
	return true
}

// IsClosed returns whether this object is closed
func (ic *InfiniteCylinder) IsClosed() bool {
	return true
}

// Copy returns a shallow copy of this object
func (ic *InfiniteCylinder) Copy() primitive.Primitive {
	newIC := *ic
	return &newIC
}

// normalAt returns the normal of this object at the specified point
// the point is assumed to be on the surface of the object
func (ic *InfiniteCylinder) normalAt(p geometry.Point) geometry.Vector {
	if ic.HasInvertedNormals {
		return ic.Ray.ClosestPoint(p).To(p).Unit().Negate()
	}
	return ic.Ray.ClosestPoint(p).To(p).Unit()
}

// Unit returns a unit infinite cylinder
func Unit(xOffset, yOffset, zOffset float64) *InfiniteCylinder {
	ic, _ := (&InfiniteCylinder{
		Ray: geometry.Ray{
			Origin: geometry.Point{
				X: 0.0 + xOffset,
				Y: 0.0 + yOffset,
				Z: 0.0 + zOffset,
			},
			Direction: geometry.Vector{
				X: 0.0 + xOffset,
				Y: 1.0 + yOffset,
				Z: 0.0 + zOffset,
			},
		},
		Radius: 1.0,
	}).Setup()
	return ic
}

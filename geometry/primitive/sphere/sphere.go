package sphere

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"fmt"
	"math"
)

// Sphere represents a sphere geometry object
type Sphere struct {
	Center             geometry.Point `json:"center"`
	Radius             float64        `json:"radius"`
	HasInvertedNormals bool           `json:"has_inverted_normals"`
	box                *aabb.AABB
	mat                material.Material
}

// Data holds information needed to construct a new sphere
// type Data struct {
// 	Center             geometry.Point
// 	Radius             float64
// 	HasInvertedNormals bool
// }

// Setup sets up a sphere
func (s *Sphere) Setup() (*Sphere, error) {
	// if sd.Center == nil {
	// 	return nil, fmt.Errorf("sphere center is nil")
	// }
	if s.Radius <= 0 {
		return nil, fmt.Errorf("sphere radius is 0 or negative")
	}
	s.box, _ = s.BoundingBox(0, 0)
	return s, nil
}

// Intersection computer the intersection of this object and a given ray if it exists
func (s *Sphere) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	// if !s.box.Intersection(ray, tMin, tMax) {
	// 	return nil, false
	// }

	centerToRayOrigin := s.Center.To(ray.Origin)

	// terms of the quadratic equation we are solving
	a := ray.Direction.Dot(ray.Direction)
	b := ray.Direction.Dot(centerToRayOrigin)
	c := centerToRayOrigin.Dot(centerToRayOrigin) - (s.Radius * s.Radius)

	preDiscriminant := b*b - a*c

	if preDiscriminant > 0 {
		root := math.Sqrt(preDiscriminant)
		// evaluate first solution, which will be smaller
		t1 := (-b - root) / a
		// return if within range
		if t1 >= tMin && t1 <= tMax {
			hitPoint := ray.PointAt(t1)
			unitHitPoint := s.Center.To(hitPoint).DivScalar(s.Radius)

			phi := math.Atan2(unitHitPoint.Z, unitHitPoint.X)
			theta := math.Asin(unitHitPoint.Y)

			u := 1 - (phi+math.Pi)/(2*math.Pi)
			v := (theta + math.Pi/2) / math.Pi

			return &material.RayHit{
				Ray:         ray,
				NormalAtHit: s.normalAt(hitPoint),
				Time:        t1,
				U:           u,
				V:           v,
				Material:    s.mat,
			}, true
		}
		// evaluate and return second solution if in range
		t2 := (-b + root) / a
		if t2 >= tMin && t2 <= tMax {
			hitPoint := ray.PointAt(t1)
			unitHitPoint := s.Center.To(hitPoint).DivScalar(s.Radius)

			phi := math.Atan2(unitHitPoint.Z, unitHitPoint.X)
			theta := math.Asin(unitHitPoint.Y)

			u := 1.0 - (phi+math.Pi)/(2*math.Pi)
			v := (theta + math.Pi/2) / math.Pi

			return &material.RayHit{
				Ray:         ray,
				NormalAtHit: s.normalAt(ray.PointAt(t2)),
				Time:        t2,
				U:           u,
				V:           v,
				Material:    s.mat,
			}, true
		}
	}

	return nil, false
}

// BoundingBox returns the AABB of this object
func (s *Sphere) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return &aabb.AABB{
		A: s.Center.SubVector(geometry.Vector{
			X: s.Radius + 1e-7,
			Y: s.Radius + 1e-7,
			Z: s.Radius + 1e-7,
		}),
		B: s.Center.AddVector(geometry.Vector{
			X: s.Radius + 1e-7,
			Y: s.Radius + 1e-7,
			Z: s.Radius + 1e-7,
		}),
	}, true
}

// SetMaterial sets this object's material
func (s *Sphere) SetMaterial(m material.Material) {
	s.mat = m
}

// IsInfinite return whether this object is infinite
func (s *Sphere) IsInfinite() bool {
	return false
}

// IsClosed return whether this object is closed
func (s *Sphere) IsClosed() bool {
	return true
}

// Copy returns a shallow copy of this object
func (s *Sphere) Copy() primitive.Primitive {
	newS := *s
	return &newS
}

func (s *Sphere) normalAt(p geometry.Point) geometry.Vector {
	if s.HasInvertedNormals {
		return p.To(s.Center).Unit()
	}
	return s.Center.To(p).Unit()
}

// Unit returns a unit sphere
func Unit(xOffset, yOffset, zOffset float64) *Sphere {
	s, _ := (&Sphere{
		Center: geometry.Point{
			X: 0.0 + xOffset,
			Y: 0.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		Radius: 0.5,
	}).Setup()
	return s
}

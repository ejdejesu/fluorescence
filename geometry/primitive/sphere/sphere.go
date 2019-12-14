package sphere

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"fmt"
	"math"
)

type sphere struct {
	Center *geometry.Point
	Radius float64
	mat    material.Material
}

type SphereData struct {
	Center *geometry.Point `json:"center"`
	Radius float64         `json:"radius"`
}

func NewSphere(sd *SphereData) (*sphere, error) {
	if sd.Center == nil {
		return nil, fmt.Errorf("Sphere center is nil")
	}
	if sd.Radius <= 0 {
		return nil, fmt.Errorf("Sphere radius is 0 or negative")
	}
	return &sphere{
		Center: sd.Center,
		Radius: sd.Radius,
	}, nil
}

func (s *sphere) Intersection(ray *geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
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
			return &material.RayHit{
				Ray:         ray,
				NormalAtHit: s.normalAt(ray.PointAt(t1)),
				T:           t1,
				Material:    s.mat,
			}, true
		}
		// evaluate and return second solution if in range
		t2 := (-b + root) / a
		if t2 >= tMin && t2 <= tMax {
			return &material.RayHit{
				Ray:         ray,
				NormalAtHit: s.normalAt(ray.PointAt(t2)),
				T:           t2,
				Material:    s.mat,
			}, true
		}
	}

	return nil, false
}

func (s *sphere) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return &aabb.AABB{
		A: s.Center.SubVector(&geometry.Vector{
			X: s.Radius + 1e-7,
			Y: s.Radius + 1e-7,
			Z: s.Radius + 1e-7,
		}),
		B: s.Center.AddVector(&geometry.Vector{
			X: s.Radius + 1e-7,
			Y: s.Radius + 1e-7,
			Z: s.Radius + 1e-7,
		}),
	}, true
}

func (s *sphere) SetMaterial(m material.Material) {
	s.mat = m
}

func (s *sphere) IsInfinite() bool {
	return false
}

func (s *sphere) IsClosed() bool {
	return true
}

func (s *sphere) Copy() primitive.Primitive {
	newS := *s
	return &newS
}

func (s *sphere) normalAt(p *geometry.Point) *geometry.Vector {
	return s.Center.To(p).Unit()
}

func Basicsphere(xOffset, yOffset, zOffset float64) *sphere {
	return &sphere{
		Center: &geometry.Point{
			X: 0.0 + xOffset,
			Y: 0.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		Radius: 0.5,
		mat:    nil,
	}
}

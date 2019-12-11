package primitive

import (
	"fluorescence/geometry"
	"fluorescence/shading/material"
	"math"
)

type Sphere struct {
	Center   *geometry.Point   `json:"center"`
	Radius   float64           `json:"radius"`
	Material material.Material `json:"material"`
}

func (s *Sphere) Intersection(ray *geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	centerToRayOrigin := s.Center.To(ray.Origin)

	// terms of the quadratic equation we are solving
	a := ray.Direction.Dot(ray.Direction)
	b := ray.Direction.Dot(centerToRayOrigin)
	c := centerToRayOrigin.Dot(centerToRayOrigin) - (s.Radius * s.Radius)

	preDiscriminant := b*b - a*c

	if preDiscriminant > 0 {
		// evaluate first solution, which will be smaller
		t1 := (-b - math.Sqrt(preDiscriminant)) / a
		// return if within range
		if t1 >= tMin && t1 <= tMax {
			return &material.RayHit{
				Ray:         ray,
				NormalAtHit: s.normalAt(ray.PointAt(t1)),
				T:           t1,
				Material:    s.Material,
			}, true
		}
		// evaluate and return second solution if in range
		t2 := (-b + math.Sqrt(preDiscriminant)) / a
		if t2 >= tMin && t2 <= tMax {
			return &material.RayHit{
				Ray:         ray,
				NormalAtHit: s.normalAt(ray.PointAt(t2)),
				T:           t2,
				Material:    s.Material,
			}, true
		}
	}

	return nil, false
}

func (s *Sphere) BoundingBox(t0, t1 float64) (*AABB, bool) {
	return &AABB{
		A: s.Center.SubVector(&geometry.Vector{
			X: s.Radius + 0.0000001,
			Y: s.Radius + 0.0000001,
			Z: s.Radius + 0.0000001,
		}),
		B: s.Center.AddVector(&geometry.Vector{
			X: s.Radius + 0.0000001,
			Y: s.Radius + 0.0000001,
			Z: s.Radius + 0.0000001,
		}),
	}, true
}

func (s *Sphere) SetMaterial(m material.Material) {
	s.Material = m
}

func (s *Sphere) Copy() Primitive {
	newS := *s
	return &newS
}

func (s *Sphere) normalAt(p *geometry.Point) *geometry.Vector {
	return s.Center.To(p).Unit()
}

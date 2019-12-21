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
	center             geometry.Point
	radius             float64
	hasInvertedNormals bool
	box                *aabb.AABB
	mat                material.Material
}

type SphereData struct {
	Center             geometry.Point `json:"center"`
	Radius             float64        `json:"radius"`
	HasInvertedNormals bool           `json:"has_inverted_normals"`
}

func NewSphere(sd *SphereData) (*sphere, error) {
	// if sd.Center == nil {
	// 	return nil, fmt.Errorf("sphere center is nil")
	// }
	if sd.Radius <= 0 {
		return nil, fmt.Errorf("sphere radius is 0 or negative")
	}
	newSphere := &sphere{
		center:             sd.Center,
		radius:             sd.Radius,
		hasInvertedNormals: sd.HasInvertedNormals,
	}
	newSphere.box, _ = newSphere.BoundingBox(0, 0)
	return newSphere, nil
}

func (s *sphere) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	// if !s.box.Intersection(ray, tMin, tMax) {
	// 	return nil, false
	// }

	centerToRayOrigin := s.center.To(ray.Origin)

	// terms of the quadratic equation we are solving
	a := ray.Direction.Dot(ray.Direction)
	b := ray.Direction.Dot(centerToRayOrigin)
	c := centerToRayOrigin.Dot(centerToRayOrigin) - (s.radius * s.radius)

	preDiscriminant := b*b - a*c

	if preDiscriminant > 0 {
		root := math.Sqrt(preDiscriminant)
		// evaluate first solution, which will be smaller
		t1 := (-b - root) / a
		// return if within range
		if t1 >= tMin && t1 <= tMax {
			hitPoint := ray.PointAt(t1)
			unitHitPoint := s.center.To(hitPoint).DivScalar(s.radius)

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
			unitHitPoint := s.center.To(hitPoint).DivScalar(s.radius)

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

func (s *sphere) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return &aabb.AABB{
		A: s.center.SubVector(geometry.Vector{
			X: s.radius + 1e-7,
			Y: s.radius + 1e-7,
			Z: s.radius + 1e-7,
		}),
		B: s.center.AddVector(geometry.Vector{
			X: s.radius + 1e-7,
			Y: s.radius + 1e-7,
			Z: s.radius + 1e-7,
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

func (s *sphere) normalAt(p geometry.Point) geometry.Vector {
	if s.hasInvertedNormals {
		return p.To(s.center).Unit()
	}
	return s.center.To(p).Unit()
}

func BasicSphere(xOffset, yOffset, zOffset float64) *sphere {
	s, _ := NewSphere(&SphereData{
		Center: geometry.Point{
			X: 0.0 + xOffset,
			Y: 0.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		Radius: 0.5,
	})
	return s
}

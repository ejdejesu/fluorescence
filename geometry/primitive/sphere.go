package primitive

import (
	"fluorescence/geometry"
	"fluorescence/shading"
	"math"
)

type Sphere struct {
	Center *geometry.Point `json:"center"`
	Radius float64         `json:"radius"`
}

func (s *Sphere) Intersection(r *geometry.Ray, tMin float64, tMax float64) (*geometry.Ray, *geometry.Vector, float64, *shading.Color, bool) {
	centerToRayOrigin := s.Center.To(r.Origin)

	// terms of the quadratic equation we are solving
	a := r.Direction.Dot(r.Direction)
	b := r.Direction.Dot(centerToRayOrigin)
	c := centerToRayOrigin.Dot(centerToRayOrigin) - (s.Radius * s.Radius)

	preDiscriminant := b*b - a*c

	if preDiscriminant > 0 {
		// evaluate first solution, which will be smaller
		t1 := (-b - math.Sqrt(preDiscriminant)) / a
		// return if within range
		if t1 >= tMin && t1 <= tMax {
			return r, s.NormalAt(r.PointAt(t1)), t1, nil, true
		}
		// evaluate and return second solution if in range
		t2 := (-b + math.Sqrt(preDiscriminant)) / a
		if t2 >= tMin && t2 <= tMax {
			return r, s.NormalAt(r.PointAt(t2)), t2, nil, true
		}
	}

	return nil, nil, 0, nil, false
}

func (s *Sphere) NormalAt(p *geometry.Point) *geometry.Vector {
	return s.Center.To(p).Unit()
}

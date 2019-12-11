package primitive

import (
	"fluorescence/geometry"
	"fluorescence/shading/material"
	"math"
)

type Triangle struct {
	A                *geometry.Point   `json:"a"`
	B                *geometry.Point   `json:"b"`
	C                *geometry.Point   `json:"c"`
	IntersectEpsilon float64           `json:"intersect_epsilon"`
	Material         material.Material `json:"material"`
}

func (t *Triangle) Intersection(ray *geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	ab := t.A.To(t.B)
	ac := t.A.To(t.C)
	pVector := ray.Direction.Cross(ac)
	determinant := ab.Dot(pVector)
	if determinant < t.IntersectEpsilon {
		// This ray is parallel to this triangle or back-facing.
		return nil, false
	}

	inverseDeterminant := 1.0 / determinant

	tVector := t.A.To(ray.Origin)
	u := inverseDeterminant * (tVector.Dot(pVector))
	if u < 0.0 || u > 1.0 {
		return nil, false
	}

	qVector := tVector.Cross(ab)
	v := inverseDeterminant * (ray.Direction.Dot(qVector))
	if v < 0.0 || u+v > 1.0 {
		return nil, false
	}

	// At this stage we can compute time to find out where the intersection point is on the line.
	time := inverseDeterminant * (ac.Dot(qVector))
	if time >= tMin && time <= tMax {
		// ray intersection
		return &material.RayHit{ray, t.normal(), time, t.Material}, true
	}
	return nil, false
}

func (t *Triangle) BoundingBox(t0, t1 float64) (*AABB, bool) {
	return &AABB{
		A: &geometry.Point{
			X: math.Min(math.Min(t.A.X, t.B.X), t.C.X) - 0.0000001,
			Y: math.Min(math.Min(t.A.Y, t.B.Y), t.C.Y) - 0.0000001,
			Z: math.Min(math.Min(t.A.Z, t.B.Z), t.C.Z) - 0.0000001,
		},
		B: &geometry.Point{
			X: math.Max(math.Max(t.A.X, t.B.X), t.C.X) + 0.0000001,
			Y: math.Max(math.Max(t.A.Y, t.B.Y), t.C.Y) + 0.0000001,
			Z: math.Max(math.Max(t.A.Z, t.B.Z), t.C.Z) + 0.0000001,
		},
	}, true
}

func (t *Triangle) SetMaterial(m material.Material) {
	t.Material = m
}

func (t *Triangle) Copy() Primitive {
	newT := *t
	return &newT
}

func (t *Triangle) normal() *geometry.Vector {
	return t.A.To(t.B).CrossInPlace(t.A.To(t.C)).UnitInPlace()
}

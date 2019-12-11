package primitive

import (
	"fluorescence/geometry"
	"fluorescence/shading/material"
)

type Triangle struct {
	A                *geometry.Point   `json:"a"`
	B                *geometry.Point   `json:"b"`
	C                *geometry.Point   `json:"c"`
	IntersectEpsilon float64           `json:"intersect_epsilon"`
	Material         material.Material `json:"material"`
}

func (t *Triangle) Intersection(ray *geometry.Ray, tMin float64, tMax float64) (*material.RayHit, bool) {
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

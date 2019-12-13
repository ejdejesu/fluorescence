package triangle

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"fmt"
	"math"
)

type triangle struct {
	A        *geometry.Point
	B        *geometry.Point
	C        *geometry.Point
	Normal   *geometry.Vector
	Material material.Material
}

type TriangleData struct {
	A *geometry.Point `json:"a"`
	B *geometry.Point `json:"b"`
	C *geometry.Point `json:"c"`
}

func NewTriangle(td *TriangleData) (*triangle, error) {
	if td.A == td.B || td.A == td.C || td.B == td.C {
		return nil, fmt.Errorf("Triangle resolves to line or point")
	}
	return &triangle{
		A:      td.A,
		B:      td.B,
		C:      td.C,
		Normal: td.A.To(td.B).CrossInPlace(td.A.To(td.C)).UnitInPlace(),
	}, nil
}

func EmptyTriangle() *triangle {
	return &triangle{}
}

func (t *triangle) Intersection(ray *geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	ab := t.A.To(t.B)
	ac := t.A.To(t.C)
	pVector := ray.Direction.Cross(ac)
	determinant := ab.Dot(pVector)
	if determinant < 1e-7 {
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
		return &material.RayHit{ray, t.Normal, time, t.Material}, true
	}
	return nil, false
}

func (t *triangle) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return &aabb.AABB{
		A: &geometry.Point{
			X: math.Min(math.Min(t.A.X, t.B.X), t.C.X) - 1e-7,
			Y: math.Min(math.Min(t.A.Y, t.B.Y), t.C.Y) - 1e-7,
			Z: math.Min(math.Min(t.A.Z, t.B.Z), t.C.Z) - 1e-7,
		},
		B: &geometry.Point{
			X: math.Max(math.Max(t.A.X, t.B.X), t.C.X) + 1e-7,
			Y: math.Max(math.Max(t.A.Y, t.B.Y), t.C.Y) + 1e-7,
			Z: math.Max(math.Max(t.A.Z, t.B.Z), t.C.Z) + 1e-7,
		},
	}, true
}

func (t *triangle) SetMaterial(m material.Material) {
	t.Material = m
}

func (t *triangle) Copy() primitive.Primitive {
	newT := *t
	return &newT
}

func BasicTriangle(xOffset, yOffset, zOffset float64) *triangle {
	return &triangle{
		A: &geometry.Point{
			X: 0.0 + xOffset,
			Y: 0.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		B: &geometry.Point{
			X: 1.0 + xOffset,
			Y: 0.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		C: &geometry.Point{
			X: 0.0 + xOffset,
			Y: 1.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		Material: nil,
	}
}

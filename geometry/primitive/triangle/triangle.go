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
	a        geometry.Point
	b        geometry.Point
	c        geometry.Point
	normal   geometry.Vector
	isCulled bool
	mat      material.Material
}

type TriangleData struct {
	A        geometry.Point `json:"a"`
	B        geometry.Point `json:"b"`
	C        geometry.Point `json:"c"`
	IsCulled bool           `json:"is_culled"`
}

func NewTriangle(td *TriangleData) (*triangle, error) {
	if td.A == td.B || td.A == td.C || td.B == td.C {
		return nil, fmt.Errorf("Triangle resolves to line or point")
	}
	return &triangle{
		a:        td.A,
		b:        td.B,
		c:        td.C,
		normal:   td.A.To(td.B).Cross(td.A.To(td.C)).Unit(),
		isCulled: td.IsCulled,
	}, nil
}

func (t *triangle) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	ab := t.a.To(t.b)
	ac := t.a.To(t.c)
	pVector := ray.Direction.Cross(ac)
	determinant := ab.Dot(pVector)
	if t.isCulled && determinant < 1e-7 {
		// This ray is parallel to this triangle or back-facing.
		return nil, false
	} else if determinant > -1e-7 && determinant < 1e-7 {
		return nil, false
	}

	inverseDeterminant := 1.0 / determinant

	tVector := t.a.To(ray.Origin)
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
		return &material.RayHit{ray, t.normal, time, 0, 0, t.mat}, true
	}
	return nil, false
}

func (t *triangle) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return &aabb.AABB{
		A: geometry.Point{
			X: math.Min(math.Min(t.a.X, t.a.X), t.c.X) - 1e-7,
			Y: math.Min(math.Min(t.a.Y, t.a.Y), t.c.Y) - 1e-7,
			Z: math.Min(math.Min(t.a.Z, t.a.Z), t.c.Z) - 1e-7,
		},
		B: geometry.Point{
			X: math.Max(math.Max(t.a.X, t.b.X), t.c.X) + 1e-7,
			Y: math.Max(math.Max(t.a.Y, t.b.Y), t.c.Y) + 1e-7,
			Z: math.Max(math.Max(t.a.Z, t.b.Z), t.c.Z) + 1e-7,
		},
	}, true
}

func (t *triangle) SetMaterial(m material.Material) {
	t.mat = m
}

func (t *triangle) IsInfinite() bool {
	return false
}

func (t *triangle) IsClosed() bool {
	return false
}

func (t *triangle) Copy() primitive.Primitive {
	newT := *t
	return &newT
}

func BasicTriangle(xOffset, yOffset, zOffset float64) *triangle {
	return &triangle{
		a: geometry.Point{
			X: 0.0 + xOffset,
			Y: 0.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		b: geometry.Point{
			X: 1.0 + xOffset,
			Y: 0.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		c: geometry.Point{
			X: 0.0 + xOffset,
			Y: 1.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		isCulled: true,
		mat:      nil,
	}
}

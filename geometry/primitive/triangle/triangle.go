package triangle

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"fmt"
	"math"
)

// triangle is an internal representation of a triangle geometry contruct
type triangle struct {
	a, b, c  geometry.Point  // points of the triangle
	normal   geometry.Vector // normal of the triangle's surface
	isCulled bool            // whether or not the triangle is culled, or single-sided
	mat      material.Material
}

// Data holds information needed to contruct a triangle
type Data struct {
	A        geometry.Point `json:"a"`
	B        geometry.Point `json:"b"`
	C        geometry.Point `json:"c"`
	IsCulled bool           `json:"is_culled"`
}

// New contructs a new triangle given a Data
func New(td *Data) (*triangle, error) {
	if td.A == td.B || td.A == td.C || td.B == td.C {
		return nil, fmt.Errorf("triangle resolves to line or point")
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
		return &material.RayHit{
			Ray:         ray,
			NormalAtHit: t.normal,
			Time:        time,
			U:           0,
			V:           0,
			Material:    t.mat,
		}, true
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

// Unit creates a unit triangle.
// The points of this triangle are:
// A: (0, 0, 0),
// B: (1, 0, 0),
// C: (0, 1, 0).
func Unit(xOffset, yOffset, zOffset float64) *triangle {
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

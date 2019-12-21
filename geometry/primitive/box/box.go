package box

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/geometry/primitive/primitivelist"
	"fluorescence/geometry/primitive/rectangle"
	"fluorescence/shading/material"
	"fmt"
)

type box struct {
	list *primitivelist.PrimitiveList
	box  *aabb.AABB
}

type BoxData struct {
	A                  geometry.Point `json:"a"`
	B                  geometry.Point `json:"b"`
	HasInvertedNormals bool           `json:"has_inverted_normals"`
}

func NewBox(bd *BoxData) (*box, error) {
	c1 := geometry.MinOf(bd.A, bd.B)
	c8 := geometry.MaxOf(bd.A, bd.B)

	if c1.X == c8.X || c1.Y == c8.Y || c1.Z == c8.Z {
		return nil, fmt.Errorf("box resolves to point, line, or plane")
	}

	rNegX, err := rectangle.NewRectangle(&rectangle.RectangleData{
		A:                 c1,
		B:                 geometry.Point{c1.X, c8.Y, c8.Z},
		HasNegativeNormal: !bd.HasInvertedNormals,
	})
	if err != nil {
		return nil, err
	}

	rPosX, err := rectangle.NewRectangle(&rectangle.RectangleData{
		A:                 geometry.Point{c8.X, c1.Y, c1.Z},
		B:                 c8,
		HasNegativeNormal: bd.HasInvertedNormals,
	})
	if err != nil {
		return nil, err
	}

	rNegY, err := rectangle.NewRectangle(&rectangle.RectangleData{
		A:                 c1,
		B:                 geometry.Point{c8.X, c1.Y, c8.Z},
		HasNegativeNormal: !bd.HasInvertedNormals,
	})
	if err != nil {
		return nil, err
	}

	rPosY, err := rectangle.NewRectangle(&rectangle.RectangleData{
		A:                 geometry.Point{c1.X, c8.Y, c1.Z},
		B:                 c8,
		HasNegativeNormal: bd.HasInvertedNormals,
	})
	if err != nil {
		return nil, err
	}

	rNegZ, err := rectangle.NewRectangle(&rectangle.RectangleData{
		A:                 c1,
		B:                 geometry.Point{c8.X, c8.Y, c1.Z},
		HasNegativeNormal: !bd.HasInvertedNormals,
	})
	if err != nil {
		return nil, err
	}

	rPosZ, err := rectangle.NewRectangle(&rectangle.RectangleData{
		A:                 geometry.Point{c1.X, c1.Y, c8.Z},
		B:                 c8,
		HasNegativeNormal: bd.HasInvertedNormals,
	})
	if err != nil {
		return nil, err
	}

	l, err := primitivelist.NewPrimitiveList(rNegX, rPosX, rNegY, rPosY, rNegZ, rPosZ)
	if err != nil {
		return nil, err
	}
	b, _ := l.BoundingBox(0, 0)

	return &box{
		list: l,
		box:  b,
	}, nil
}

func (b *box) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	if b.box.Intersection(ray, tMin, tMax) {
		return b.list.Intersection(ray, tMin, tMax)
	}
	return nil, false
}

func (b *box) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return b.box, true
}

func (b *box) SetMaterial(m material.Material) {
	b.list.SetMaterial(m)
}

func (b *box) IsInfinite() bool {
	return false
}

func (b *box) IsClosed() bool {
	return true
}

func (b *box) Copy() primitive.Primitive {
	newB := *b
	return &newB
}

func BasicBox(xOffset, yOffset, zOffset float64) *box {
	bd := BoxData{
		A: geometry.Point{
			X: 0.0 + xOffset,
			Y: 0.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		B: geometry.Point{
			X: 1.0 + xOffset,
			Y: 1.0 + yOffset,
			Z: 1.0 + zOffset,
		},
	}
	b, _ := NewBox(&bd)
	return b
}

package pyramid

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/geometry/primitive/primitivelist"
	"fluorescence/geometry/primitive/rectangle"
	"fluorescence/geometry/primitive/triangle"
	"fluorescence/shading/material"
	"fmt"
)

type pyramid struct {
	list *primitivelist.PrimitiveList
	box  *aabb.AABB
}

type Data struct {
	A                  geometry.Point `json:"a"`
	B                  geometry.Point `json:"b"`
	Height             float64        `json:"height"`
	HasInvertedNormals bool           `json:"has_inverted_normals"`
}

func New(pd *Data) (*pyramid, error) {
	if pd.Height <= 0 {
		return nil, fmt.Errorf("pyramid height is 0 or negative")
	}
	if pd.A.Y != pd.B.Y {
		return nil, fmt.Errorf("pyramid is not directed upwards")
	}

	c1 := geometry.MinComponents(pd.A, pd.B)
	c3 := geometry.MaxComponents(pd.A, pd.B)
	c2 := geometry.Point{c1.X, c1.Y, c3.Z}
	c4 := geometry.Point{c3.X, c1.Y, c1.Z}

	base, err := rectangle.New(&rectangle.Data{
		A:                 pd.A,
		B:                 pd.B,
		IsCulled:          false,
		HasNegativeNormal: true,
	})
	if err != nil {
		return nil, err
	}

	diagonalBaseVectorHalf := c1.To(c3).DivScalar(2.0)
	baseCenterPoint := c1.AddVector(diagonalBaseVectorHalf)
	topPoint := baseCenterPoint.AddVector(geometry.VectorUp.MultScalar(pd.Height))

	tri1, err := triangle.New(&triangle.Data{
		A:        c1,
		B:        c2,
		C:        topPoint,
		IsCulled: false,
	})
	if err != nil {
		return nil, err
	}

	tri2, err := triangle.New(&triangle.Data{
		A:        c2,
		B:        c3,
		C:        topPoint,
		IsCulled: false,
	})
	if err != nil {
		return nil, err
	}

	tri3, err := triangle.New(&triangle.Data{
		A:        c3,
		B:        c4,
		C:        topPoint,
		IsCulled: false,
	})
	if err != nil {
		return nil, err
	}

	tri4, err := triangle.New(&triangle.Data{
		A:        c4,
		B:        c1,
		C:        topPoint,
		IsCulled: false,
	})
	if err != nil {
		return nil, err
	}

	l, err := primitivelist.NewPrimitiveList(base, tri1, tri2, tri3, tri4)
	if err != nil {
		return nil, err
	}
	b, _ := l.BoundingBox(0, 0)
	return &pyramid{
		list: l,
		box:  b,
	}, nil
}

func (p *pyramid) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	if p.box.Intersection(ray, tMin, tMax) {
		return p.list.Intersection(ray, tMin, tMax)
	}
	return nil, false
}

func (p *pyramid) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return p.box, true
}

func (p *pyramid) SetMaterial(m material.Material) {
	p.list.SetMaterial(m)
}

func (p *pyramid) IsInfinite() bool {
	return false
}

func (p *pyramid) IsClosed() bool {
	return true
}

func (p *pyramid) Copy() primitive.Primitive {
	newP := *p
	return &newP
}

func UnitPyramid(xOffset, yOffset, zOffset float64) *pyramid {
	rd := Data{
		A: geometry.Point{
			X: 0.0 + xOffset,
			Y: 0.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		B: geometry.Point{
			X: 1.0 + xOffset,
			Y: 0.0 + yOffset,
			Z: 1.0 + zOffset,
		},
		Height: 1.0,
	}
	r, _ := New(&rd)
	return r
}

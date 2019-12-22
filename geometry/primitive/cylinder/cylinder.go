package cylinder

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/geometry/primitive/disk"
	"fluorescence/geometry/primitive/primitivelist"
	"fluorescence/geometry/primitive/uncappedcylinder"
	"fluorescence/shading/material"
)

type cylinder struct {
	list *primitivelist.PrimitiveList
	box  *aabb.AABB
}

type Data struct {
	A      geometry.Point `json:"a"`
	B      geometry.Point `json:"b"`
	Radius float64        `json:"radius"`
}

func New(cd *Data) (*cylinder, error) {
	uncappedCylinder, err := uncappedcylinder.New(&uncappedcylinder.Data{
		A:                  cd.A,
		B:                  cd.B,
		Radius:             cd.Radius,
		HasInvertedNormals: false,
	})
	if err != nil {
		return nil, err
	}
	diskA, err := disk.New(&disk.Data{
		Center:   cd.A,
		Normal:   cd.B.To(cd.A).Unit(),
		Radius:   cd.Radius,
		IsCulled: false,
	})
	if err != nil {
		return nil, err
	}
	diskB, err := disk.New(&disk.Data{
		Center:   cd.B,
		Normal:   cd.A.To(cd.B).Unit(),
		Radius:   cd.Radius,
		IsCulled: false,
	})
	if err != nil {
		return nil, err
	}
	primitiveList, err := primitivelist.NewPrimitiveList(
		uncappedCylinder,
		diskA,
		diskB,
	)
	if err != nil {
		return nil, err
	}
	box, _ := primitiveList.BoundingBox(0, 0)
	return &cylinder{
		list: primitiveList,
		box:  box,
	}, nil
}

func (c *cylinder) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	if c.box.Intersection(ray, tMin, tMax) {
		return c.list.Intersection(ray, tMin, tMax)
	}
	return nil, false
}

func (c *cylinder) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return c.list.BoundingBox(0, 0)
}

func (c *cylinder) SetMaterial(m material.Material) {
	c.list.SetMaterial(m)
}

func (c *cylinder) IsInfinite() bool {
	return false
}

func (c *cylinder) IsClosed() bool {
	return true
}

func (c *cylinder) Copy() primitive.Primitive {
	newC := *c
	return &newC
}

func UnitCylinder(xOffset, yOffset, zOffset float64) *cylinder {
	cd := Data{
		A: geometry.Point{
			X: 0.0 + xOffset,
			Y: 0.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		B: geometry.Point{
			X: 0.0 + xOffset,
			Y: 1.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		Radius: 1.0,
	}
	c, _ := New(&cd)
	return c
}

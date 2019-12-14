package cylinder

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/geometry/primitive/disk"
	"fluorescence/geometry/primitive/primitivelist"
	"fluorescence/shading/material"
)

type cylinder struct {
	list *primitivelist.PrimitiveList
}

type CylinderData struct {
	A      *geometry.Point `json:"a"`
	B      *geometry.Point `json:"b"`
	Radius float64         `json:"radius"`
}

func NewCylinder(cd *CylinderData) (*cylinder, error) {
	uncappedCylinder, err := NewUncappedCylinder(&UncappedCylinderData{
		A:                  cd.A,
		B:                  cd.B,
		Radius:             cd.Radius,
		HasInvertedNormals: false,
	})
	if err != nil {
		return nil, err
	}
	diskA, err := disk.NewDisk(&disk.DiskData{
		Center:   cd.A,
		Normal:   cd.B.To(cd.A).Unit(),
		Radius:   cd.Radius,
		IsCulled: false,
	})
	if err != nil {
		return nil, err
	}
	diskB, err := disk.NewDisk(&disk.DiskData{
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
	return &cylinder{
		list: primitiveList,
	}, nil
}

func (c *cylinder) Intersection(ray *geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	return c.list.Intersection(ray, tMin, tMax)
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

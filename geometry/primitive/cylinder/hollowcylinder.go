package cylinder

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/geometry/primitive/disk"
	"fluorescence/geometry/primitive/primitivelist"
	"fluorescence/shading/material"
)

type hollowCylinder struct {
	list *primitivelist.PrimitiveList
}

type HollowCylinderData struct {
	A           *geometry.Point `json:"a"`
	B           *geometry.Point `json:"b"`
	InnerRadius float64         `json:"inner_radius"`
	OuterRadius float64         `json:"outer_radius"`
}

func NewHollowCylinder(hcd *HollowCylinderData) (*hollowCylinder, error) {
	outerUncappedCylinder, err := NewUncappedCylinder(&UncappedCylinderData{
		A:                  hcd.A,
		B:                  hcd.B,
		Radius:             hcd.OuterRadius,
		HasInvertedNormals: false,
	})
	if err != nil {
		return nil, err
	}
	innerUncappedCylinder, err := NewUncappedCylinder(&UncappedCylinderData{
		A:                  hcd.A,
		B:                  hcd.B,
		Radius:             hcd.InnerRadius,
		HasInvertedNormals: true,
	})
	if err != nil {
		return nil, err
	}
	hollowDiskA, err := disk.NewHollowDisk(&disk.HollowDiskData{
		Center:      hcd.A,
		Normal:      hcd.B.To(hcd.A).Unit(),
		InnerRadius: hcd.InnerRadius,
		OuterRadius: hcd.OuterRadius,
		IsCulled:    false,
	})
	if err != nil {
		return nil, err
	}
	hollowDiskB, err := disk.NewHollowDisk(&disk.HollowDiskData{
		Center:      hcd.B,
		Normal:      hcd.A.To(hcd.B).Unit(),
		InnerRadius: hcd.InnerRadius,
		OuterRadius: hcd.OuterRadius,
		IsCulled:    false,
	})
	if err != nil {
		return nil, err
	}
	primitiveList, err := primitivelist.NewPrimitiveList(
		innerUncappedCylinder,
		outerUncappedCylinder,
		hollowDiskA,
		hollowDiskB,
	)
	if err != nil {
		return nil, err
	}
	return &hollowCylinder{
		list: primitiveList,
	}, nil
}

func (hc *hollowCylinder) Intersection(ray *geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	return hc.list.Intersection(ray, tMin, tMax)
}

func (hc *hollowCylinder) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return hc.list.BoundingBox(0, 0)
}

func (hc *hollowCylinder) SetMaterial(m material.Material) {
	hc.list.SetMaterial(m)
}

func (hc *hollowCylinder) IsInfinite() bool {
	return false
}

func (hc *hollowCylinder) IsClosed() bool {
	return true
}

func (hc *hollowCylinder) Copy() primitive.Primitive {
	newHC := *hc
	return &newHC
}

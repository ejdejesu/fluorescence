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
	box  *aabb.AABB
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
	box, _ := primitiveList.BoundingBox(0, 0)
	return &hollowCylinder{
		list: primitiveList,
		box:  box,
	}, nil
}

func (hc *hollowCylinder) Intersection(ray *geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	if hc.box.Intersection(ray, tMin, tMax) {
		return hc.list.Intersection(ray, tMin, tMax)
	}
	return nil, false
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

func BasicHollowCylinder(xOffset, yOffset, zOffset float64) *hollowCylinder {
	hcd := HollowCylinderData{
		A: &geometry.Point{
			X: 0.0 + xOffset,
			Y: 0.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		B: &geometry.Point{
			X: 0.0 + xOffset,
			Y: 1.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		InnerRadius: 0.5,
		OuterRadius: 1.0,
	}
	hc, _ := NewHollowCylinder(&hcd)
	return hc
}

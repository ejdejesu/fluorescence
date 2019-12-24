package hollowcylinder

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/geometry/primitive/hollowdisk"
	"fluorescence/geometry/primitive/primitivelist"
	"fluorescence/geometry/primitive/uncappedcylinder"
	"fluorescence/shading/material"
)

// HollowCylinder represents a hollow cylinder geometry object
type HollowCylinder struct {
	A           geometry.Point `json:"a"`
	B           geometry.Point `json:"b"`
	InnerRadius float64        `json:"inner_radius"`
	OuterRadius float64        `json:"outer_radius"`
	list        *primitivelist.PrimitiveList
	box         *aabb.AABB
}

// type Data struct {
// }

// Setup sets up this hollow cylinder's internal fields
func (hc *HollowCylinder) Setup() (*HollowCylinder, error) {
	outerUncappedCylinder, err := (&uncappedcylinder.UncappedCylinder{
		A:                  hc.A,
		B:                  hc.B,
		Radius:             hc.OuterRadius,
		HasInvertedNormals: false,
	}).Setup()
	if err != nil {
		return nil, err
	}
	innerUncappedCylinder, err := (&uncappedcylinder.UncappedCylinder{
		A:                  hc.A,
		B:                  hc.B,
		Radius:             hc.InnerRadius,
		HasInvertedNormals: true,
	}).Setup()
	if err != nil {
		return nil, err
	}
	hollowDiskA, err := (&hollowdisk.HollowDisk{
		Center:      hc.A,
		Normal:      hc.B.To(hc.A).Unit(),
		InnerRadius: hc.InnerRadius,
		OuterRadius: hc.OuterRadius,
		IsCulled:    false,
	}).Setup()
	if err != nil {
		return nil, err
	}
	hollowDiskB, err := (&hollowdisk.HollowDisk{
		Center:      hc.B,
		Normal:      hc.A.To(hc.B).Unit(),
		InnerRadius: hc.InnerRadius,
		OuterRadius: hc.OuterRadius,
		IsCulled:    false,
	}).Setup()
	if err != nil {
		return nil, err
	}
	primitiveList, err := primitivelist.FromElements(
		innerUncappedCylinder,
		outerUncappedCylinder,
		hollowDiskA,
		hollowDiskB,
	)
	if err != nil {
		return nil, err
	}
	hc.list = primitiveList
	hc.box, _ = primitiveList.BoundingBox(0, 0)
	return hc, nil
}

// Intersection computer the intersection of this object and a given ray if it exists
func (hc *HollowCylinder) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	if hc.box.Intersection(ray, tMin, tMax) {
		return hc.list.Intersection(ray, tMin, tMax)
	}
	return nil, false
}

// BoundingBox returns an AABB of this object
func (hc *HollowCylinder) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return hc.list.BoundingBox(0, 0)
}

// SetMaterial sets this object's material
func (hc *HollowCylinder) SetMaterial(m material.Material) {
	hc.list.SetMaterial(m)
}

// IsInfinite returns whether this object is infinite
func (hc *HollowCylinder) IsInfinite() bool {
	return false
}

// IsClosed returns whether this object is closed
func (hc *HollowCylinder) IsClosed() bool {
	return true
}

// Copy returns a shallow copy of this object
func (hc *HollowCylinder) Copy() primitive.Primitive {
	newHC := *hc
	return &newHC
}

// Unit return a unit hollow cylinder
func Unit(xOffset, yOffset, zOffset float64) *HollowCylinder {
	hc, _ := (&HollowCylinder{
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
		InnerRadius: 0.5,
		OuterRadius: 1.0,
	}).Setup()
	return hc
}

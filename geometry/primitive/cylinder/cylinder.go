package cylinder

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/geometry/primitive/disk"
	"fluorescence/geometry/primitive/primitivelist"
	"fluorescence/geometry/primitive/uncappedcylinder"
	"fluorescence/shading/material"

	"github.com/go-gl/mathgl/mgl64"
)

// Cylinder represents a capped cylinder object
type Cylinder struct {
	A      mgl64.Vec3 `json:"a"`
	B      mgl64.Vec3 `json:"b"`
	Radius float64    `json:"radius"`
	list   *primitivelist.PrimitiveList
	box    *aabb.AABB
}

// type Data struct {
// }

// Setup sets up a cylinder's internal fields
func (c *Cylinder) Setup() (*Cylinder, error) {
	uncappedCylinder, err := (&uncappedcylinder.UncappedCylinder{
		A:                  c.A,
		B:                  c.B,
		Radius:             c.Radius,
		HasInvertedNormals: false,
	}).Setup()
	if err != nil {
		return nil, err
	}
	diskA, err := (&disk.Disk{
		Center:   c.A,
		Normal:   c.A.Sub(c.B).Normalize(),
		Radius:   c.Radius,
		IsCulled: false,
	}).Setup()
	if err != nil {
		return nil, err
	}
	diskB, err := (&disk.Disk{
		Center:   c.B,
		Normal:   c.B.Sub(c.A).Normalize(),
		Radius:   c.Radius,
		IsCulled: false,
	}).Setup()
	if err != nil {
		return nil, err
	}
	primitiveList, err := primitivelist.FromElements(
		uncappedCylinder,
		diskA,
		diskB,
	)
	if err != nil {
		return nil, err
	}
	c.list = primitiveList
	c.box, _ = primitiveList.BoundingBox(0, 0)
	return c, nil
}

// Intersection computer the intersection of this object and a given ray if it exists
func (c *Cylinder) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	if c.box.Intersection(ray, tMin, tMax) {
		return c.list.Intersection(ray, tMin, tMax)
	}
	return nil, false
}

// BoundingBox returns an AABB of this object
func (c *Cylinder) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return c.list.BoundingBox(0, 0)
}

// SetMaterial sets this object's material
func (c *Cylinder) SetMaterial(m material.Material) {
	c.list.SetMaterial(m)
}

// IsInfinite returns whether this object is infinite
func (c *Cylinder) IsInfinite() bool {
	return false
}

// IsClosed returns whether this object is closed
func (c *Cylinder) IsClosed() bool {
	return true
}

// Copy returns a shallow copy of this object
func (c *Cylinder) Copy() primitive.Primitive {
	newC := *c
	return &newC
}

// Unit returns a unit cylinder
func Unit(xOffset, yOffset, zOffset float64) *Cylinder {
	c, _ := (&Cylinder{
		A: mgl64.Vec3{
			0.0 + xOffset,
			0.0 + yOffset,
			0.0 + zOffset,
		},
		B: mgl64.Vec3{
			0.0 + xOffset,
			1.0 + yOffset,
			0.0 + zOffset,
		},
		Radius: 1.0,
	}).Setup()
	return c
}

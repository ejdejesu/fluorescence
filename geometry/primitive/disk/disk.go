package disk

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"fmt"
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

// Disk represent a disk geometric object
type Disk struct {
	Center        mgl64.Vec3 `json:"center"`
	Normal        mgl64.Vec3 `json:"normal"`
	Radius        float64    `json:"radius"`
	IsCulled      bool       `json:"is_culled"`
	radiusSquared float64
	mat           material.Material
}

// type Data struct {
// }

// Setup sets up a disk's internal fields
func (d *Disk) Setup() (*Disk, error) {
	// if d.Center == nil || d.Normal == nil {
	// 	return nil, fmt.Errorf("disk center or normal is nil")
	// }
	if d.Radius <= 0.0 {
		return nil, fmt.Errorf("disk radius is 0 or negative")
	}
	d.radiusSquared = d.Radius * d.Radius
	return d, nil
}

// Intersection computer the intersection of this object and a given ray if it exists
func (d *Disk) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	denominator := ray.Direction.Dot(d.Normal)
	if d.IsCulled && denominator > -1e-7 {
		return nil, false
	} else if denominator < 1e-7 && denominator > -1e-7 {
		return nil, false
	}
	planeVector := d.Center.Sub(ray.Origin)
	t := planeVector.Dot(d.Normal) / denominator

	if t < tMin || t > tMax {
		return nil, false
	}

	hitPoint := ray.PointAt(t)
	diskVector := hitPoint.Sub(d.Center)

	// // fmt.Println(d.RadiusSquared, d.Center)
	if diskVector.Dot(diskVector) > d.radiusSquared {
		return nil, false
	}
	// if diskVector.Magnitude() > d.Radius {
	// 	return nil, false
	// }

	return &material.RayHit{
		Ray:         ray,
		NormalAtHit: d.Normal,
		Time:        t,
		Material:    d.mat,
	}, true
}

// BoundingBox return an AABB of this disk
func (d *Disk) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	eX := d.Radius * math.Sqrt(1.0-d.Normal.X()*d.Normal.X())
	eY := d.Radius * math.Sqrt(1.0-d.Normal.Y()*d.Normal.Y())
	eZ := d.Radius * math.Sqrt(1.0-d.Normal.Z()*d.Normal.Z())
	return &aabb.AABB{
		A: mgl64.Vec3{
			d.Center.X() - eX - 1e-7,
			d.Center.Y() - eY - 1e-7,
			d.Center.Z() - eZ - 1e-7,
		},
		B: mgl64.Vec3{
			d.Center.X() + eX + 1e-7,
			d.Center.Y() + eY + 1e-7,
			d.Center.Z() + eZ + 1e-7,
		},
	}, true
}

// SetMaterial sets this object's material
func (d *Disk) SetMaterial(m material.Material) {
	d.mat = m
}

// IsInfinite returns whether this object is infinite
func (d *Disk) IsInfinite() bool {
	return false
}

// IsClosed returns whether this object is closed
func (d *Disk) IsClosed() bool {
	return false
}

// Copy return a shallow copy of this object
func (d *Disk) Copy() primitive.Primitive {
	newD := *d
	return &newD
}

// Unit return a unit disk
func Unit(xOffset, yOffset, zOffset float64) *Disk {
	d, _ := (&Disk{
		Center: mgl64.Vec3{
			0.0 + xOffset,
			0.0 + yOffset,
			0.0 + zOffset,
		},
		Normal: mgl64.Vec3{
			0.0 + xOffset,
			0.0 + yOffset,
			-1.0 + zOffset,
		},
		Radius: 1.0,
	}).Setup()
	return d
}

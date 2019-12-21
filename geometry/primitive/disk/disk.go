package disk

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"fmt"
	"math"
)

type disk struct {
	center        geometry.Point
	normal        geometry.Vector
	radius        float64
	isCulled      bool
	radiusSquared float64
	mat           material.Material
}

type DiskData struct {
	Center   geometry.Point  `json:"center"`
	Normal   geometry.Vector `json:"normal"`
	Radius   float64         `json:"radius"`
	IsCulled bool            `json:"is_culled"`
}

func NewDisk(dd *DiskData) (*disk, error) {
	// if dd.Center == nil || dd.Normal == nil {
	// 	return nil, fmt.Errorf("Disk center or normal is nil")
	// }
	if dd.Radius <= 0.0 {
		return nil, fmt.Errorf("Disk radius is 0 or negative")
	}
	return &disk{
		center:        dd.Center,
		normal:        dd.Normal.Unit(),
		radius:        dd.Radius,
		isCulled:      dd.IsCulled,
		radiusSquared: dd.Radius * dd.Radius,
	}, nil
}

func (d *disk) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	denominator := ray.Direction.Dot(d.normal)
	if d.isCulled && denominator > -1e-7 {
		return nil, false
	} else if denominator < 1e-7 && denominator > -1e-7 {
		return nil, false
	}
	planeVector := ray.Origin.To(d.center)
	t := planeVector.Dot(d.normal) / denominator

	if t < tMin || t > tMax {
		return nil, false
	}

	hitPoint := ray.PointAt(t)
	diskVector := d.center.To(hitPoint)

	// // fmt.Println(d.radiusSquared, d.Center)
	if diskVector.Dot(diskVector) > d.radiusSquared {
		return nil, false
	}
	// if diskVector.Magnitude() > d.Radius {
	// 	return nil, false
	// }

	return &material.RayHit{
		Ray:         ray,
		NormalAtHit: d.normal,
		Time:           t,
		Material:    d.mat,
	}, true
}

func (d *disk) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	eX := d.radius * math.Sqrt(1.0-d.normal.X*d.normal.X)
	eY := d.radius * math.Sqrt(1.0-d.normal.Y*d.normal.Y)
	eZ := d.radius * math.Sqrt(1.0-d.normal.Z*d.normal.Z)
	return &aabb.AABB{
		A: geometry.Point{
			X: d.center.X - eX - 1e-7,
			Y: d.center.Y - eY - 1e-7,
			Z: d.center.Z - eZ - 1e-7,
		},
		B: geometry.Point{
			X: d.center.X + eX + 1e-7,
			Y: d.center.Y + eY + 1e-7,
			Z: d.center.Z + eZ + 1e-7,
		},
	}, true
}

func (d *disk) SetMaterial(m material.Material) {
	d.mat = m
}

func (d *disk) IsInfinite() bool {
	return false
}

func (d *disk) IsClosed() bool {
	return false
}

func (d *disk) Copy() primitive.Primitive {
	newD := *d
	return &newD
}

func BasicDisk(xOffset, yOffset, zOffset float64) *disk {
	dd := DiskData{
		Center: geometry.Point{
			X: 0.0 + xOffset,
			Y: 0.0 + yOffset,
			Z: 0.0 + zOffset,
		},
		Normal: geometry.Vector{
			X: 0.0 + xOffset,
			Y: 0.0 + yOffset,
			Z: -1.0 + zOffset,
		},
		Radius: 1.0,
	}
	d, _ := NewDisk(&dd)
	return d
}

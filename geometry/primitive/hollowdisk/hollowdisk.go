package hollowdisk

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"fmt"
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

// HollowDisk represents a hollow disk geometry object
type HollowDisk struct {
	Center             mgl64.Vec3 `json:"center"`
	Normal             mgl64.Vec3 `json:"normal"`
	InnerRadius        float64    `json:"inner_radius"`
	OuterRadius        float64    `json:"outer_radius"`
	IsCulled           bool       `json:"is_culled"`
	innerRadiusSquared float64
	outerRadiusSquared float64
	mat                material.Material
}

// type Data struct {
// 	Center     mgl64.Vec3
// 	Normal     mgl64.Vec3
// 	InnerRadius float64
// 	OuterRadius float64
// 	IsCulled    bool
// }

// Setup sets up this hollow disk
func (hd *HollowDisk) Setup() (*HollowDisk, error) {
	// if hd.Center == nil || hd.Normal == nil {
	// 	return nil, fmt.Errorf("hollow disk center or normal is nil")
	// }
	if hd.InnerRadius > hd.OuterRadius {
		return nil, fmt.Errorf("hollow disk inner radius is lesser than radius")
	}
	if hd.InnerRadius == hd.OuterRadius {
		return nil, fmt.Errorf("hollow disk outer radius equals inner radius")
	}
	if hd.InnerRadius < 0.0 {
		return nil, fmt.Errorf("hollow disk inner radius is negative")
	}
	if hd.OuterRadius <= 0 {
		return nil, fmt.Errorf("hollow disk outer radius is 0 or negative")
	}
	hd.Normal = hd.Normal.Normalize()
	hd.innerRadiusSquared = hd.InnerRadius * hd.InnerRadius
	hd.outerRadiusSquared = hd.OuterRadius * hd.OuterRadius
	return hd, nil
}

// Intersection computer the intersection of this object and a given ray if it exists
func (hd *HollowDisk) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	denominator := ray.Direction.Dot(hd.Normal)
	if hd.IsCulled && denominator > -1e-7 {
		return nil, false
	} else if denominator < 1e-7 && denominator > -1e-7 {
		return nil, false
	}
	planeVector := hd.Center.Sub(ray.Origin)
	t := planeVector.Dot(hd.Normal) / denominator

	if t < tMin || t > tMax {
		return nil, false
	}

	hitPoint := ray.PointAt(t)
	diskVector := hitPoint.Sub(hd.Center)

	// // fmt.Println(d.radiusSquared, d.Center)
	if diskVector.Dot(diskVector) > hd.outerRadiusSquared {
		return nil, false
	}
	if diskVector.Dot(diskVector) < hd.innerRadiusSquared {
		return nil, false
	}
	// if diskVector.Magnitude() > d.Radius {
	// 	return nil, false
	// }

	return &material.RayHit{
		Ray:         ray,
		NormalAtHit: hd.Normal,
		Time:        t,
		Material:    hd.mat,
	}, true
}

// BoundingBox returns an AABB of this object
func (hd *HollowDisk) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	eX := hd.OuterRadius * math.Sqrt(1.0-hd.Normal.X()*hd.Normal.X())
	eY := hd.OuterRadius * math.Sqrt(1.0-hd.Normal.Y()*hd.Normal.Y())
	eZ := hd.OuterRadius * math.Sqrt(1.0-hd.Normal.Z()*hd.Normal.Z())
	return &aabb.AABB{
		A: mgl64.Vec3{
			hd.Center.X() - eX,
			hd.Center.Y() - eY,
			hd.Center.Z() - eZ,
		},
		B: mgl64.Vec3{
			hd.Center.X() + eX,
			hd.Center.Y() + eY,
			hd.Center.Z() + eZ,
		},
	}, true
}

// SetMaterial sets thie object's material
func (hd *HollowDisk) SetMaterial(m material.Material) {
	hd.mat = m
}

// IsInfinite returns whether this object is infinite
func (hd *HollowDisk) IsInfinite() bool {
	return false
}

// IsClosed returns whether this object is closed
func (hd *HollowDisk) IsClosed() bool {
	return false
}

// Copy returns a shallow copy of thie object
func (hd *HollowDisk) Copy() primitive.Primitive {
	newHD := *hd
	return &newHD
}

// Unit return a unit hollow disk
func Unit(xOffset, yOffset, zOffset float64) *HollowDisk {
	hd, _ := (&HollowDisk{
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
		InnerRadius: 0.5,
		OuterRadius: 1.0,
	}).Setup()
	return hd
}

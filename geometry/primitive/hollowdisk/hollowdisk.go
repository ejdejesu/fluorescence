package hollowdisk

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"fmt"
	"math"
)

type hollowDisk struct {
	center             *geometry.Point
	normal             *geometry.Vector
	innerRadius        float64
	outerRadius        float64
	isCulled           bool
	innerRadiusSquared float64
	outerRadiusSquared float64
	Material           material.Material
}

type HollowDiskData struct {
	Center      *geometry.Point  `json:"center"`
	Normal      *geometry.Vector `json:"normal"`
	InnerRadius float64          `json:"inner_radius"`
	OuterRadius float64          `json:"outer_radius"`
	IsCulled    bool             `json:"is_culled"`
}

func NewHollowDisk(hdd *HollowDiskData) (*hollowDisk, error) {
	if hdd.Center == nil || hdd.Normal == nil {
		return nil, fmt.Errorf("HollowDisk center or normal is nil")
	}
	if hdd.InnerRadius > hdd.OuterRadius {
		return nil, fmt.Errorf("HollowDisk inner radius is lesser than radius")
	}
	if hdd.InnerRadius == hdd.OuterRadius {
		return nil, fmt.Errorf("HollowDisk outer radius equals inner radius")
	}
	if hdd.InnerRadius < 0.0 {
		return nil, fmt.Errorf("HollowDisk inner radius is negative")
	}
	if hdd.OuterRadius <= 0 {
		return nil, fmt.Errorf("HollowDisk outer radius is 0 or negative")
	}
	return &hollowDisk{
		center:             hdd.Center,
		normal:             hdd.Normal.Unit(),
		innerRadius:        hdd.InnerRadius,
		outerRadius:        hdd.OuterRadius,
		isCulled:           hdd.IsCulled,
		innerRadiusSquared: hdd.InnerRadius * hdd.InnerRadius,
		outerRadiusSquared: hdd.OuterRadius * hdd.OuterRadius,
	}, nil
}

func EmptyHollowDisk() *hollowDisk {
	return &hollowDisk{}
}

func (hd *hollowDisk) Intersection(ray *geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	denominator := ray.Direction.Dot(hd.normal)
	if hd.isCulled && denominator > -1e-7 {
		return nil, false
	} else if denominator < 1e-7 && denominator > -1e-7 {
		return nil, false
	}
	planeVector := ray.Origin.To(hd.center)
	t := planeVector.Dot(hd.normal) / denominator

	if t < tMin || t > tMax {
		return nil, false
	}

	hitPoint := ray.PointAt(t)
	diskVector := hd.center.To(hitPoint)

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
		NormalAtHit: hd.normal,
		T:           t,
		Material:    hd.Material,
	}, true
}

func (hd *hollowDisk) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	eX := hd.outerRadius * math.Sqrt(1.0-hd.normal.X*hd.normal.X)
	eY := hd.outerRadius * math.Sqrt(1.0-hd.normal.Y*hd.normal.Y)
	eZ := hd.outerRadius * math.Sqrt(1.0-hd.normal.Z*hd.normal.Z)
	return &aabb.AABB{
		A: &geometry.Point{
			X: hd.center.X - eX,
			Y: hd.center.Y - eY,
			Z: hd.center.Z - eZ,
		},
		B: &geometry.Point{
			X: hd.center.X + eX,
			Y: hd.center.Y + eY,
			Z: hd.center.Z + eZ,
		},
	}, true
}

func (hd *hollowDisk) SetMaterial(m material.Material) {
	hd.Material = m
}

func (hd *hollowDisk) Copy() primitive.Primitive {
	newHD := *hd
	return &newHD
}

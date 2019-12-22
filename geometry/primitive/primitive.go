package primitive

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
)

// Primitive represents a geometry object with a material in 3D space in the scene
type Primitive interface {
	Intersection(geometry.Ray, float64, float64) (*material.RayHit, bool)
	BoundingBox(float64, float64) (*aabb.AABB, bool)
	SetMaterial(material.Material)
	IsInfinite() bool
	IsClosed() bool
	Copy() Primitive
}

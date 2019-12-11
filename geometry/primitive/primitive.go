package primitive

import (
	"fluorescence/geometry"
	"fluorescence/shading/material"
)

type Primitive interface {
	Intersection(*geometry.Ray, float64, float64) (*material.RayHit, bool)
	BoundingBox(float64, float64) (*AABB, bool)
	SetMaterial(material.Material)
	Copy() Primitive
}

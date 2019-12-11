package primitive

import (
	"fluorescence/geometry"
	"fluorescence/shading/material"
)

type bvh struct {
	Left  Primitive
	Right Primitive
	box   *AABB
}

func NewBVH(p Primitive)

func (b *bvh) Intersection(ray *geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	_, hitBox := b.box.Intersection(ray, tMin, tMax)
	if hitBox {
		leftRayHit, doesHitLeft := b.Left.Intersection(ray, tMin, tMax)
		rightRayHit, doesHitRight := b.Right.Intersection(ray, tMin, tMax)
		if doesHitLeft && doesHitRight {
			if leftRayHit.T < rightRayHit.T {
				return leftRayHit, true
			}
			return rightRayHit, true
		} else if doesHitLeft {
			return leftRayHit, true
		} else if doesHitRight {
			return rightRayHit, true
		}
		return nil, false
	}
	return nil, false
}

func (b *bvh) BoundingBox(t0, t1 float64) (*AABB, bool) {
	return b.box, true
}

package bvh

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/geometry/primitive/primitivelist"
	"fluorescence/shading/material"
	"fmt"
	"math/rand"
	"sort"
)

type bvh struct {
	Left     primitive.Primitive
	Right    primitive.Primitive
	isSingle bool
	box      *aabb.AABB
}

func NewBVH(pl *primitivelist.PrimitiveList) (*bvh, error) {
	newBVH := &bvh{}

	// can we do the sort?
	_, ok := pl.BoundingBox(0, 0)
	if !ok {
		return nil, fmt.Errorf("no bounding box for input Primitive List")
	}

	// do the sort
	axisNum := rand.Intn(3)
	if axisNum == 0 {
		sort.Sort(primitivelist.ByXPos(*pl))
	} else if axisNum == 1 {
		sort.Sort(primitivelist.ByYPos(*pl))
	} else {
		sort.Sort(primitivelist.ByZPos(*pl))
	}

	// fill children
	if len(pl.List) == 1 {
		newBVH.Left = pl.List[0]
		newBVH.isSingle = true
	} else {
		left, err := NewBVH(pl.FirstHalfCopy())
		if err != nil {
			return nil, err
		}
		right, err := NewBVH(pl.LastHalfCopy())
		if err != nil {
			return nil, err
		}
		newBVH.Left = left
		newBVH.Right = right
	}
	// est. box
	leftBox, leftOk := newBVH.Left.BoundingBox(0, 0)
	if newBVH.isSingle {
		if !leftOk {
			return nil, fmt.Errorf("no bounding box for some leaf of BVH")
		}
		newBVH.box = leftBox
	} else {
		rightBox, rightOk := newBVH.Right.BoundingBox(0, 0)
		if !leftOk || !rightOk {
			return nil, fmt.Errorf("no bounding box for some leaf of BVH")
		}
		newBVH.box = aabb.SurroundingBox(leftBox, rightBox)
	}
	return newBVH, nil
}

func (b *bvh) Intersection(ray *geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	hitBox := b.box.Intersection(ray, tMin, tMax)
	if hitBox {
		leftRayHit, doesHitLeft := b.Left.Intersection(ray, tMin, tMax)
		if b.isSingle {
			if doesHitLeft {
				return leftRayHit, true
			}
			return nil, false
		}
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

func (b *bvh) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return b.box, true
}

func (b *bvh) SetMaterial(m material.Material) {
	b.Left.SetMaterial(m)
	if !b.isSingle {
		b.Right.SetMaterial(m)
	}
}

func (b *bvh) IsInfinite() bool {
	if b.isSingle {
		return b.Left.IsInfinite()
	}
	return b.Left.IsInfinite() || b.Right.IsInfinite()
}

func (b *bvh) IsClosed() bool {
	if b.isSingle {
		return b.Left.IsClosed()
	}
	return b.Left.IsClosed() && b.Right.IsClosed()
}

func (b *bvh) Copy() primitive.Primitive {
	newB := *b
	return &newB
}

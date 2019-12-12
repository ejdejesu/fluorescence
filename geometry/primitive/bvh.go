package primitive

import (
	"fluorescence/geometry"
	"fluorescence/shading/material"
	"fmt"
	"math/rand"
	"sort"
)

type bvh struct {
	Left   Primitive
	Right  Primitive
	single bool
	box    *AABB
}

func NewBVH(pl *PrimitiveList) (*bvh, error) {
	newBVH := &bvh{}

	// do the sort
	axisNum := rand.Intn(3)
	if axisNum == 0 {
		sort.Sort(ByXPos(*pl))
	} else if axisNum == 1 {
		sort.Sort(ByYPos(*pl))
	} else {
		sort.Sort(ByZPos(*pl))
	}

	// fill children
	if len(pl.List) == 1 {
		newBVH.Left = pl.List[0]
		newBVH.single = true
	} else if len(pl.List) == 2 {
		newBVH.Left = pl.List[0]
		newBVH.Right = pl.List[1]
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
	if newBVH.single {
		if !leftOk {
			return nil, fmt.Errorf("no bounding box for some leaf of BVH")
		}
		newBVH.box = leftBox
	} else {
		rightBox, rightOk := newBVH.Right.BoundingBox(0, 0)
		if !leftOk || !rightOk {
			return nil, fmt.Errorf("no bounding box for some leaf of BVH")
		}
		newBVH.box = SurroundingBox(leftBox, rightBox)
	}
	return newBVH, nil
}

func (b *bvh) Intersection(ray *geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	hitBox := b.box.Intersection(ray, tMin, tMax)
	if hitBox {
		leftRayHit, doesHitLeft := b.Left.Intersection(ray, tMin, tMax)
		if b.single {
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

func (b *bvh) BoundingBox(t0, t1 float64) (*AABB, bool) {
	return b.box, true
}

func (b *bvh) SetMaterial(m material.Material) {
	b.Left.SetMaterial(m)
	b.Right.SetMaterial(m)
}

func (b *bvh) Copy() Primitive {
	newB := *b
	return &newB
}

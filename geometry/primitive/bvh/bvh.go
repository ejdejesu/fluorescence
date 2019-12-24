package bvh

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/geometry/primitive/primitivelist"
	"fluorescence/shading/material"
	"fmt"
	"math"
	"sort"
)

// BVH represents a bounding volume hierarchy
type BVH struct {
	left     primitive.Primitive
	right    primitive.Primitive
	isSingle bool
	box      *aabb.AABB
}

// New sets up and returns a new BVH
func New(pl *primitivelist.PrimitiveList) (*BVH, error) {
	newBVH := &BVH{}

	// can we do the sort?
	_, ok := pl.BoundingBox(0, 0)
	if !ok {
		return nil, fmt.Errorf("no bounding box for input Primitive List")
	}

	// pick the best axis
	var axisNum int
	firstBox, _ := pl.List[0].BoundingBox(0, 0)
	lastBox, _ := pl.List[len(pl.List)-1].BoundingBox(0, 0)

	xDif := math.Abs(firstBox.A.X - lastBox.A.X)
	yDif := math.Abs(firstBox.A.Y - lastBox.A.Y)
	zDif := math.Abs(firstBox.A.Z - lastBox.A.Z)
	if xDif > yDif && xDif > zDif {
		axisNum = 0
	} else if yDif > xDif && yDif > zDif {
		axisNum = 1
	} else {
		axisNum = 2
	}

	// do the sort
	// axisNum := rand.Intn(3)
	if axisNum == 0 {
		sort.Sort(primitivelist.ByXPos(*pl))
	} else if axisNum == 1 {
		sort.Sort(primitivelist.ByYPos(*pl))
	} else {
		sort.Sort(primitivelist.ByZPos(*pl))
	}

	// fill children
	if len(pl.List) == 1 {
		newBVH.left = pl.List[0]
		newBVH.isSingle = true
	} else {
		left, err := New(pl.FirstHalfCopy())
		if err != nil {
			return nil, err
		}
		right, err := New(pl.LastHalfCopy())
		if err != nil {
			return nil, err
		}
		newBVH.left = left
		newBVH.right = right
	}
	// est. box
	leftBox, leftOk := newBVH.left.BoundingBox(0, 0)
	if newBVH.isSingle {
		if !leftOk {
			return nil, fmt.Errorf("no bounding box for some leaf of BVH")
		}
		newBVH.box = leftBox
	} else {
		rightBox, rightOk := newBVH.right.BoundingBox(0, 0)
		if !leftOk || !rightOk {
			return nil, fmt.Errorf("no bounding box for some leaf of BVH")
		}
		newBVH.box = aabb.SurroundingBox(leftBox, rightBox)
	}
	return newBVH, nil
}

// Intersection computer the intersection of this object and a given ray if it exists
func (b *BVH) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	hitBox := b.box.Intersection(ray, tMin, tMax)
	if hitBox {
		leftRayHit, doesHitLeft := b.left.Intersection(ray, tMin, tMax)
		if b.isSingle {
			if doesHitLeft {
				return leftRayHit, true
			}
			return nil, false
		}
		rightRayHit, doesHitRight := b.right.Intersection(ray, tMin, tMax)
		if doesHitLeft && doesHitRight {
			if leftRayHit.Time < rightRayHit.Time {
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

// BoundingBox returns a new AABB for this object
func (b *BVH) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	return b.box, true
}

// SetMaterial sets this object's material
func (b *BVH) SetMaterial(m material.Material) {
	b.left.SetMaterial(m)
	if !b.isSingle {
		b.right.SetMaterial(m)
	}
}

// IsInfinite returns whether this object is infinite
func (b *BVH) IsInfinite() bool {
	if b.isSingle {
		return b.left.IsInfinite()
	}
	return b.left.IsInfinite() || b.right.IsInfinite()
}

// IsClosed returns whether this object is closed
func (b *BVH) IsClosed() bool {
	if b.isSingle {
		return b.left.IsClosed()
	}
	return b.left.IsClosed() && b.right.IsClosed()
}

// Copy returns a shallow copy of this object
func (b *BVH) Copy() primitive.Primitive {
	newB := *b
	return &newB
}

package primitivelist

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"math"
)

type PrimitiveList struct {
	List []primitive.Primitive
}

type ByXPos PrimitiveList
type ByYPos PrimitiveList
type ByZPos PrimitiveList

func (a ByXPos) Len() int {
	return len(a.List)
}

func (a ByYPos) Len() int {
	return len(a.List)
}

func (a ByZPos) Len() int {
	return len(a.List)
}

func (a ByXPos) Swap(i, j int) {
	a.List[i], a.List[j] = a.List[j], a.List[i]
}

func (a ByYPos) Swap(i, j int) {
	a.List[i], a.List[j] = a.List[j], a.List[i]
}

func (a ByZPos) Swap(i, j int) {
	a.List[i], a.List[j] = a.List[j], a.List[i]
}

func (a ByXPos) Less(i, j int) bool {
	box1, _ := a.List[i].BoundingBox(0, 0)
	box2, _ := a.List[j].BoundingBox(0, 0)
	return box1.A.X < box2.A.X
}

func (a ByYPos) Less(i, j int) bool {
	box1, _ := a.List[i].BoundingBox(0, 0)
	box2, _ := a.List[j].BoundingBox(0, 0)
	return box1.A.Y < box2.A.Y
}

func (a ByZPos) Less(i, j int) bool {
	box1, _ := a.List[i].BoundingBox(0, 0)
	box2, _ := a.List[j].BoundingBox(0, 0)
	return box1.A.Z < box2.A.Z
}

func NewPrimitiveList(primitives ...primitive.Primitive) (*PrimitiveList, error) {
	primitiveList := &PrimitiveList{}
	for _, p := range primitives {
		primitiveList.List = append(primitiveList.List, p)
	}
	return primitiveList, nil
}

func (pl *PrimitiveList) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
	var rayHit *material.RayHit
	minT := math.MaxFloat64
	hitSomething := false
	for _, p := range pl.List {
		rh, wasHit := p.Intersection(ray, tMin, tMax)
		if wasHit && rh.T < minT {
			hitSomething = true
			rayHit = rh
			minT = rh.T
		}
	}
	if hitSomething {
		return rayHit, true
	}
	return nil, false
}

func (pl *PrimitiveList) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {
	box, ok := pl.List[0].BoundingBox(t0, t1)
	if !ok {
		return nil, false
	}
	for i := 1; i < len(pl.List); i++ {
		newBox, ok := pl.List[i].BoundingBox(t0, t1)
		if !ok {
			return nil, false
		}
		box = aabb.SurroundingBox(box, newBox)
	}
	return box, true
}

func (pl *PrimitiveList) SetMaterial(m material.Material) {
	for _, p := range pl.List {
		p.SetMaterial(m)
	}
}

func (pl *PrimitiveList) IsInfinite() bool {
	for _, p := range pl.List {
		if p.IsInfinite() {
			return true
		}
	}
	return false
}

func (pl *PrimitiveList) IsClosed() bool {
	for _, p := range pl.List {
		if !p.IsClosed() {
			return false
		}
	}
	return false
}

func (pl *PrimitiveList) Copy() primitive.Primitive {
	newPL := &PrimitiveList{}
	for _, p := range pl.List {
		newPL.List = append(newPL.List, p.Copy())
	}
	return newPL
}

func (pl *PrimitiveList) FirstHalfCopy() *PrimitiveList {
	newPL := &PrimitiveList{}
	lowerBound := 0
	upperBound := len(pl.List) / 2
	for i := lowerBound; i < upperBound; i++ {
		newPL.List = append(newPL.List, pl.List[i])
	}
	return newPL
}

func (pl *PrimitiveList) LastHalfCopy() *PrimitiveList {
	newPL := &PrimitiveList{}
	lowerBound := len(pl.List) / 2
	upperBound := len(pl.List)
	for i := lowerBound; i < upperBound; i++ {
		newPL.List = append(newPL.List, pl.List[i])
	}
	return newPL
}

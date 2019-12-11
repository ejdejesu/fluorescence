package primitive

import (
	"fluorescence/geometry"
	"fluorescence/shading/material"
	"math"
)

type PrimitiveList struct {
	List []Primitive
}

func (pl *PrimitiveList) Intersection(ray *geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {
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
	return rayHit, hitSomething
}

func (pl *PrimitiveList) BoundingBox(t0, t1 float64) (*AABB, bool) {
	box, ok := pl.List[0].BoundingBox(t0, t1)
	if !ok {
		return nil, false
	}
	for i := 1; i < len(pl.List); i++ {
		newBox, ok := pl.List[i].BoundingBox(t0, t1)
		if !ok {
			return nil, false
		}
		box = SurroundingBox(box, newBox)
	}
	return box, true
}

func (pl *PrimitiveList) SetMaterial(m material.Material) {
	for _, p := range pl.List {
		p.SetMaterial(m)
	}
}

func (pl *PrimitiveList) Copy() Primitive {
	newPL := &PrimitiveList{}
	for _, p := range pl.List {
		newPL.List = append(newPL.List, p.Copy())
	}
	return newPL
}

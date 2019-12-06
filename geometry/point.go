package geometry

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	X float64
	Y float64
	Z float64
}

var ORIGIN = &Point{0.0, 0.0, 0.0}

func (p *Point) To(q *Point) *Vector {
	return q.toVector().Subtract(p.toVector())
}

func (p *Point) From(q *Point) *Vector {
	return p.toVector().Subtract(q.toVector())
}

func (p *Point) SubtractPoint(q *Point) *Vector {
	return p.toVector().Subtract(q.toVector())
}

func (p *Point) AddVector(v *Vector) *Point {
	return &Point{p.X + v.X, p.Y + v.Y, p.Z + v.Z}
}

func (p *Point) SubtractVector(v *Vector) *Point {
	return &Point{p.X - v.X, p.Y - v.Y, p.Z - v.Z}
}

func (p *Point) toVector() *Vector {
	return &Vector{p.X, p.Y, p.Z}
}

func (p *Point) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{&p.X, &p.Y, &p.Z}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if len(tmp) != wantLen {
		return fmt.Errorf("wrong number of fields: %d != %d", len(tmp), wantLen)
	}
	return nil
}

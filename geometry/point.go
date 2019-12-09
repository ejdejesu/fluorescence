package geometry

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
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

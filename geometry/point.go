package geometry

type Point Vector

var ORIGIN = &Point{0.0, 0.0, 0.0}

func (p *Point) To(q *Point) *Vector {
	return q.asVector().Sub(p.asVector())
}

func (p *Point) ToInPlace(q *Point) *Vector {
	return q.asVector().SubInPlace(p.asVector())
}

func (p *Point) From(q *Point) *Vector {
	return p.asVector().Sub(q.asVector())
}

func (p *Point) FromInPlace(q *Point) *Vector {
	return p.asVector().SubInPlace(q.asVector())
}

func (p *Point) SubPoint(q *Point) *Vector {
	return p.asVector().Sub(q.asVector())
}

func (p *Point) SubPointInPlace(q *Point) *Vector {
	return p.asVector().SubInPlace(q.asVector())
}

func (p *Point) AddVector(v *Vector) *Point {
	return &Point{p.X + v.X, p.Y + v.Y, p.Z + v.Z}
}

func (p *Point) AddVectorInPlace(v *Vector) *Point {
	p.X += v.X
	p.Y += v.Y
	p.Z += v.Z
	return p
}

func (p *Point) SubVector(v *Vector) *Point {
	return &Point{p.X - v.X, p.Y - v.Y, p.Z - v.Z}
}

func (p *Point) SubVectorInPlace(v *Vector) *Point {
	p.X -= v.X
	p.Y -= v.Y
	p.Z -= v.Z
	return p
}

func (p *Point) asVector() *Vector {
	return (*Vector)(p)
}

package geometry

type Point Vector

var POINT_ZERO = Point{}

func (p Point) To(q Point) Vector {
	return q.asVector().Sub(p.asVector())
}

func (p Point) From(q Point) Vector {
	return p.asVector().Sub(q.asVector())
}

func (p Point) SubPoint(q Point) Vector {
	return p.asVector().Sub(q.asVector())
}

func (p Point) AddVector(v Vector) Point {
	return Point{p.X + v.X, p.Y + v.Y, p.Z + v.Z}
}

func (p Point) SubVector(v Vector) Point {
	return Point{p.X - v.X, p.Y - v.Y, p.Z - v.Z}
}

func (p Point) asVector() Vector {
	return (Vector)(p)
}

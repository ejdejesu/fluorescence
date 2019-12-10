package geometry

type Ray struct {
	Origin    *Point
	Direction *Vector
}

func (r *Ray) PointAt(s float64) *Point {
	return r.Origin.AddVector(r.Direction.MultScalar(s))
}

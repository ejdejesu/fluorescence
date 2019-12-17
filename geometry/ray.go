package geometry

type Ray struct {
	Origin    Point  `json:"origin"`
	Direction Vector `json:"direction"`
}

var RAY_ZERO = Ray{}

func (r Ray) PointAt(s float64) Point {
	return r.Origin.AddVector(r.Direction.MultScalar(s))
}

func (r Ray) ClosestPoint(p Point) Point {
	return r.PointAt(r.ClosestTime(p))
}

func (r Ray) ClosestTime(p Point) float64 {
	originToPoint := r.Origin.To(p)
	return originToPoint.Dot(r.Direction)
}

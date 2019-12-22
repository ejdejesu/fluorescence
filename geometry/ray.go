package geometry

// Ray defines elements of a parametric ray equation
type Ray struct {
	Origin    Point  `json:"origin"`
	Direction Vector `json:"direction"`
}

// RayZero defines the zero ray
var RayZero = Ray{}

// PointAt returns the result of solving the parametric ray equation (p = O + tD) for p
func (r Ray) PointAt(t float64) Point {
	return r.Origin.AddVector(r.Direction.MultScalar(t))
}

// ClosestPoint returns the closest point on the ray to point p
func (r Ray) ClosestPoint(p Point) Point {
	return r.PointAt(r.ClosestTime(p))
}

// ClosestTime returns the ray time of the closest point on the ray to point p
func (r Ray) ClosestTime(p Point) float64 {
	originToPoint := r.Origin.To(p)
	return originToPoint.Dot(r.Direction)
}

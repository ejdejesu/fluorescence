package geometry

import (
	"github.com/go-gl/mathgl/mgl64"
)

// Ray defines elements of a parametric ray equation
type Ray struct {
	Origin    mgl64.Vec3 `json:"origin"`
	Direction mgl64.Vec3 `json:"direction"`
}

// RayZero defines the zero ray
var RayZero = Ray{}

// PointAt returns the result of solving the parametric ray equation (p = O + tD) for p
func (r Ray) PointAt(t float64) mgl64.Vec3 {
	return r.Origin.Add(r.Direction.Mul(t))
}

// ClosestPoint returns the closest point on the ray to point p
func (r Ray) ClosestPoint(p mgl64.Vec3) mgl64.Vec3 {
	return r.PointAt(r.ClosestTime(p))
}

// ClosestTime returns the ray time of the closest point on the ray to point p
func (r Ray) ClosestTime(p mgl64.Vec3) float64 {
	originToPoint := p.Sub(r.Origin)
	return originToPoint.Dot(r.Direction)
}

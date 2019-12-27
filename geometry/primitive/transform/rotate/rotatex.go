package rotate

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"math"
)

// RotationX is a primitive with a rotations around the y axis attached
type RotationX struct {
	AngleDegrees float64     `json:"angle"`
	TypeName     string      `json:"type"`
	Data         interface{} `json:"data"`
	Primitive    primitive.Primitive
	theta        float64
	sinTheta     float64
	cosTheta     float64
}

// Setup sets up some internal fields of a rotation
func (rx *RotationX) Setup() (*RotationX, error) {
	// convert to radians and save
	rx.theta = (math.Pi / 180.0) * rx.AngleDegrees
	// find sin(theta)
	rx.sinTheta = math.Sin(rx.theta)
	// find cos(theta)
	rx.cosTheta = math.Cos(rx.theta)
	return rx, nil
}

// Intersection computer the intersection of this object and a given ray if it exists
func (rx *RotationX) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {

	rotatedRay := ray

	rotatedRay.Origin.Y = rx.cosTheta*ray.Origin.Y + rx.sinTheta*ray.Origin.Z
	rotatedRay.Origin.Z = -rx.sinTheta*ray.Origin.Y + rx.cosTheta*ray.Origin.Z

	rotatedRay.Direction.Y = rx.cosTheta*ray.Direction.Y + rx.sinTheta*ray.Direction.Z
	rotatedRay.Direction.Z = -rx.sinTheta*ray.Direction.Y + rx.cosTheta*ray.Direction.Z

	rayHit, wasHit := rx.Primitive.Intersection(rotatedRay, tMin, tMax)
	if wasHit {
		unrotatedNormal := rayHit.NormalAtHit
		unrotatedNormal.Y = rx.cosTheta*rayHit.NormalAtHit.Y - rx.sinTheta*rayHit.NormalAtHit.Z
		unrotatedNormal.Z = rx.sinTheta*rayHit.NormalAtHit.Y + rx.cosTheta*rayHit.NormalAtHit.Z
		return &material.RayHit{
			Ray:         ray,
			NormalAtHit: unrotatedNormal,
			Time:        rayHit.Time,
			U:           rayHit.U,
			V:           rayHit.V,
			Material:    rayHit.Material,
		}, true
	}
	return nil, false
}

// BoundingBox returns an AABB for this object
func (rx *RotationX) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {

	box, ok := rx.Primitive.BoundingBox(t0, t1)
	if !ok {
		return nil, false
	}
	minPoint := geometry.PointMax
	maxPoint := geometry.PointMax.Negate()
	for i := 0.0; i < 2; i++ {
		for j := 0.0; j < 2; j++ {
			for k := 0.0; k < 2; k++ {
				x := i*box.B.X + (1-i)*box.A.X
				y := j*box.B.Y + (1-j)*box.A.Y
				z := k*box.B.Z + (1-k)*box.A.Z

				newY := rx.cosTheta*y - rx.sinTheta*z
				newZ := rx.sinTheta*y + rx.cosTheta*z

				rotatedCorner := geometry.Point{
					X: x,
					Y: newY,
					Z: newZ,
				}

				maxPoint = geometry.MaxComponents(maxPoint, rotatedCorner)
				minPoint = geometry.MinComponents(minPoint, rotatedCorner)
			}
		}
	}
	return &aabb.AABB{
		A: minPoint,
		B: maxPoint,
	}, true
}

// SetMaterial sets the material of this object
func (rx *RotationX) SetMaterial(m material.Material) {
	rx.Primitive.SetMaterial(m)
}

// IsInfinite returns whether this object is infinite
func (rx *RotationX) IsInfinite() bool {
	return rx.Primitive.IsInfinite()
}

// IsClosed returns whether this object is closed
func (rx *RotationX) IsClosed() bool {
	return rx.Primitive.IsClosed()
}

// Copy returns a shallow copy of this object
func (rx *RotationX) Copy() primitive.Primitive {
	newRX := *rx
	return &newRX
}

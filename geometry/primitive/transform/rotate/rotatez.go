package rotate

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"math"
)

// RotationZ is a primitive with a rotations around the y axis attached
type RotationZ struct {
	AngleDegrees float64     `json:"angle"`
	TypeName     string      `json:"type"`
	Data         interface{} `json:"data"`
	Primitive    primitive.Primitive
	theta        float64
	sinTheta     float64
	cosTheta     float64
}

// Setup sets up some internal fields of a rotation
func (rz *RotationZ) Setup() (*RotationZ, error) {
	// convert to radians and save
	rz.theta = (math.Pi / 180.0) * rz.AngleDegrees
	// find sin(theta)
	rz.sinTheta = math.Sin(rz.theta)
	// find cos(theta)
	rz.cosTheta = math.Cos(rz.theta)
	return rz, nil
}

// Intersection computer the intersection of this object and a given ray if it exists
func (rz *RotationZ) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {

	rotatedRay := ray

	rotatedRay.Origin.X = rz.cosTheta*ray.Origin.X + rz.sinTheta*ray.Origin.Y
	rotatedRay.Origin.Y = -rz.sinTheta*ray.Origin.X + rz.cosTheta*ray.Origin.Y

	rotatedRay.Direction.X = rz.cosTheta*ray.Direction.X + rz.sinTheta*ray.Direction.Y
	rotatedRay.Direction.Y = -rz.sinTheta*ray.Direction.X + rz.cosTheta*ray.Direction.Y

	rayHit, wasHit := rz.Primitive.Intersection(rotatedRay, tMin, tMax)
	if wasHit {
		unrotatedNormal := rayHit.NormalAtHit
		unrotatedNormal.X = rz.cosTheta*rayHit.NormalAtHit.X - rz.sinTheta*rayHit.NormalAtHit.Y
		unrotatedNormal.Y = rz.sinTheta*rayHit.NormalAtHit.X + rz.cosTheta*rayHit.NormalAtHit.Y
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
func (rz *RotationZ) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {

	box, ok := rz.Primitive.BoundingBox(t0, t1)
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

				newX := rz.cosTheta*x - rz.sinTheta*y
				newY := rz.sinTheta*x + rz.cosTheta*y

				rotatedCorner := geometry.Point{
					X: newX,
					Y: newY,
					Z: z,
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
func (rz *RotationZ) SetMaterial(m material.Material) {
	rz.Primitive.SetMaterial(m)
}

// IsInfinite returns whether this object is infinite
func (rz *RotationZ) IsInfinite() bool {
	return rz.Primitive.IsInfinite()
}

// IsClosed returns whether this object is closed
func (rz *RotationZ) IsClosed() bool {
	return rz.Primitive.IsClosed()
}

// Copy returns a shallow copy of this object
func (rz *RotationZ) Copy() primitive.Primitive {
	newRZ := *rz
	return &newRZ
}

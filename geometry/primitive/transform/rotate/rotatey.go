package rotate

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"math"

	"github.com/go-gl/mathgl/mgl64"
)

// RotationY is a primitive with a rotations around the y axis attached
type RotationY struct {
	AngleDegrees float64     `json:"angle"`
	TypeName     string      `json:"type"`
	Data         interface{} `json:"data"`
	Primitive    primitive.Primitive
	theta        float64
	sinTheta     float64
	cosTheta     float64
}

// Setup sets up some internal fields of a rotation
func (ry *RotationY) Setup() (*RotationY, error) {
	// convert to radians and save
	ry.theta = (math.Pi / 180.0) * ry.AngleDegrees
	// find sin(theta)
	ry.sinTheta = math.Sin(ry.theta)
	// find cos(theta)
	ry.cosTheta = math.Cos(ry.theta)
	return ry, nil
}

// Intersection computer the intersection of this object and a given ray if it exists
func (ry *RotationY) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {

	rotatedRay := ray

	rotatedRay.Origin[0] = ry.cosTheta*ray.Origin.X() - ry.sinTheta*ray.Origin.Z()
	rotatedRay.Origin[2] = ry.sinTheta*ray.Origin.X() + ry.cosTheta*ray.Origin.Z()

	rotatedRay.Direction[0] = ry.cosTheta*ray.Direction.X() - ry.sinTheta*ray.Direction.Z()
	rotatedRay.Direction[2] = ry.sinTheta*ray.Direction.X() + ry.cosTheta*ray.Direction.Z()

	rayHit, wasHit := ry.Primitive.Intersection(rotatedRay, tMin, tMax)
	if wasHit {
		unrotatedNormal := rayHit.NormalAtHit
		unrotatedNormal[0] = ry.cosTheta*rayHit.NormalAtHit.X() + ry.sinTheta*rayHit.NormalAtHit.Z()
		unrotatedNormal[2] = -ry.sinTheta*rayHit.NormalAtHit.X() + ry.cosTheta*rayHit.NormalAtHit.Z()
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
func (ry *RotationY) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {

	box, ok := ry.Primitive.BoundingBox(t0, t1)
	if !ok {
		return nil, false
	}
	minPoint := geometry.Vec3Max
	maxPoint := geometry.Vec3Max.Mul(-1.0)
	for i := 0.0; i < 2; i++ {
		for j := 0.0; j < 2; j++ {
			for k := 0.0; k < 2; k++ {
				x := i*box.B.X() + (1-i)*box.A.X()
				y := j*box.B.Y() + (1-j)*box.A.Y()
				z := k*box.B.Z() + (1-k)*box.A.Z()

				newX := ry.cosTheta*x + ry.sinTheta*z
				newZ := -ry.sinTheta*x + ry.cosTheta*z

				rotatedCorner := mgl64.Vec3{
					newX,
					y,
					newZ,
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
func (ry *RotationY) SetMaterial(m material.Material) {
	ry.Primitive.SetMaterial(m)
}

// IsInfinite returns whether this object is infinite
func (ry *RotationY) IsInfinite() bool {
	return ry.Primitive.IsInfinite()
}

// IsClosed returns whether this object is closed
func (ry *RotationY) IsClosed() bool {
	return ry.Primitive.IsClosed()
}

// Copy returns a shallow copy of this object
func (ry *RotationY) Copy() primitive.Primitive {
	newRY := *ry
	return &newRY
}

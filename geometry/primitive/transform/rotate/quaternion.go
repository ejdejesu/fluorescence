package rotate

import (
	"fluorescence/geometry"
	"fluorescence/geometry/primitive"
	"fluorescence/geometry/primitive/aabb"
	"fluorescence/shading/material"
	"fmt"
	"strings"

	"github.com/go-gl/mathgl/mgl64"
)

// Quaternion is a quaternion rotation
type Quaternion struct {
	AxisAngles [3]float64  `json:"axis_angles"`
	Order      string      `json:"order"`
	TypeName   string      `json:"type"`
	Data       interface{} `json:"data"`
	Primitive  primitive.Primitive
	quaternion mgl64.Quat
	inverse    mgl64.Quat
}

// Setup sets up some internal fields of a rotation
func (q *Quaternion) Setup() (*Quaternion, error) {
	q.Order = strings.ToUpper(q.Order)

	var rotationOrder mgl64.RotationOrder
	switch q.Order {
	case "XYX":
		rotationOrder = mgl64.XYX
	case "XYZ":
		rotationOrder = mgl64.XYZ
	case "XZX":
		rotationOrder = mgl64.XZX
	case "XZY":
		rotationOrder = mgl64.XZY
	case "YXY":
		rotationOrder = mgl64.YXY
	case "YXZ":
		rotationOrder = mgl64.YXZ
	case "YZX":
		rotationOrder = mgl64.YZX
	case "YZY":
		rotationOrder = mgl64.YZY
	case "ZXY":
		rotationOrder = mgl64.ZXY
	case "ZXZ":
		rotationOrder = mgl64.ZXZ
	case "ZYX":
		rotationOrder = mgl64.ZYX
	case "ZYZ":
		rotationOrder = mgl64.ZYZ
	default:
		return nil, fmt.Errorf("invalid order (%s) for quaternion", q.Order)
	}

	q.quaternion = mgl64.AnglesToQuat(
		mgl64.DegToRad(q.AxisAngles[0]),
		mgl64.DegToRad(q.AxisAngles[1]),
		mgl64.DegToRad(q.AxisAngles[2]),
		rotationOrder,
	)
	q.inverse = q.quaternion.Inverse()
	return q, nil
}

// Intersection computer the intersection of this object and a given ray if it exists
func (q *Quaternion) Intersection(ray geometry.Ray, tMin, tMax float64) (*material.RayHit, bool) {

	rotatedRay := ray

	rotatedRay.Origin = q.inverse.Rotate(ray.Origin)
	rotatedRay.Direction = q.inverse.Rotate(ray.Direction)

	rayHit, wasHit := q.Primitive.Intersection(rotatedRay, tMin, tMax)
	if wasHit {
		unrotatedNormal := q.quaternion.Rotate(rayHit.NormalAtHit)
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
func (q *Quaternion) BoundingBox(t0, t1 float64) (*aabb.AABB, bool) {

	box, ok := q.Primitive.BoundingBox(t0, t1)
	if !ok {
		return nil, false
	}
	minPoint := geometry.Vec3Max
	maxPoint := geometry.NegateVec3(geometry.Vec3Max)
	for i := 0.0; i < 2; i++ {
		for j := 0.0; j < 2; j++ {
			for k := 0.0; k < 2; k++ {
				unrotatedCorner := mgl64.Vec3{
					i*box.B.X() + (1-i)*box.A.X(),
					j*box.B.Y() + (1-j)*box.A.Y(),
					k*box.B.Z() + (1-k)*box.A.Z(),
				}

				rotatedCorner := q.quaternion.Rotate(unrotatedCorner)

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
func (q *Quaternion) SetMaterial(m material.Material) {
	q.Primitive.SetMaterial(m)
}

// IsInfinite returns whether this object is infinite
func (q *Quaternion) IsInfinite() bool {
	return q.Primitive.IsInfinite()
}

// IsClosed returns whether this object is closed
func (q *Quaternion) IsClosed() bool {
	return q.Primitive.IsClosed()
}

// Copy returns a shallow copy of this object
func (q *Quaternion) Copy() primitive.Primitive {
	newRX := *q
	return &newRX
}

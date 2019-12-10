package main

import (
	"fluorescence/geometry"
	"math"
	"math/rand"
)

type Camera struct {
	EyeLocation    *geometry.Point  `json:"eye_location"`
	TargetLocation *geometry.Point  `json:"target_location"`
	UpVector       *geometry.Vector `json:"up_vector"`
	VerticalFOV    float64          `json:"vertical_fov"`
	AspectRatio    float64          `json:"aspect_ratio"`
	Aperture       float64          `json:"aperture"`
	FocusDistance  float64          `json:"focus_distance"`

	lensRadius float64 `json:"-"`
	theta      float64 `json:"-"`
	halfWidth  float64 `json:"-"`
	halfHeight float64 `json:"-"`

	w *geometry.Vector `json:"-"`
	u *geometry.Vector `json:"-"`
	v *geometry.Vector `json:"-"`

	lowerLeftCorner *geometry.Point  `json:"-"`
	horizonal       *geometry.Vector `json:"-"`
	verical         *geometry.Vector `json:"-"`
}

func (c *Camera) Setup(p *Parameters) error {
	c.UpVector = c.UpVector.Unit()
	c.AspectRatio = float64(p.ImageWidth) / float64(p.ImageHeight)

	c.lensRadius = c.Aperture / 2.0
	c.theta = c.VerticalFOV * math.Pi / 180.0
	c.halfHeight = math.Tan(c.theta / 2.0)
	c.halfWidth = c.AspectRatio * c.halfHeight

	c.w = c.TargetLocation.To(c.EyeLocation).Unit()
	c.u = c.UpVector.Cross(c.w)
	c.v = c.w.Cross(c.u)

	c.lowerLeftCorner = c.EyeLocation.SubVector(
		c.u.MultScalar(c.halfWidth * c.FocusDistance)).SubVector(
		c.v.MultScalar(c.halfHeight * c.FocusDistance)).SubVector(
		c.w.MultScalar(c.FocusDistance))

	c.horizonal = c.u.MultScalar(2.0 * c.halfWidth * c.FocusDistance)
	c.verical = c.v.MultScalar(2.0 * c.halfHeight * c.FocusDistance)

	return nil
}

func (c *Camera) GetRay(u float64, v float64, rng *rand.Rand) *geometry.Ray {
	randomOnLens := geometry.RandomOnUnitDisc(rng).MultScalarInPlace(c.lensRadius)
	offset := c.u.MultScalar(randomOnLens.X).AddInPlace(c.v.MultScalar(randomOnLens.Y))
	return &geometry.Ray{
		Origin: c.EyeLocation.AddVector(offset),
		Direction: c.lowerLeftCorner.AddVector(
			c.horizonal.MultScalar(u)).AddVectorInPlace(
			c.verical.MultScalar(v)).FromInPlace(
			c.EyeLocation).SubInPlace(
			offset).UnitInPlace(),
	}
}

package main

import (
	"fluorescence/geometry"
	"math"
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

func (c *Camera) Setup() error {
	c.UpVector = c.UpVector.Unit()

	c.lensRadius = c.Aperture / 2.0
	c.theta = c.VerticalFOV * math.Pi / 180.0
	c.halfHeight = math.Tan(c.theta / 2.0)
	c.halfWidth = c.AspectRatio * c.halfHeight

	c.w = c.TargetLocation.To(c.EyeLocation).Unit()
	c.u = c.UpVector.Cross(c.w)
	c.v = c.w.Cross(c.u)

	c.lowerLeftCorner = c.EyeLocation.SubtractVector(
		c.u.MultiplyFloat64(c.halfWidth * c.FocusDistance)).SubtractVector(
		c.v.MultiplyFloat64(c.halfHeight * c.FocusDistance)).SubtractVector(
		c.w.MultiplyFloat64(c.FocusDistance))

	c.horizonal = c.u.MultiplyFloat64(2.0 * c.halfWidth * c.FocusDistance)
	c.verical = c.v.MultiplyFloat64(2.0 * c.halfHeight * c.FocusDistance)

	return nil
}

func (c *Camera) GetRay(u float64, v float64) *geometry.Ray {
	randomOnLens := geometry.RandomOnUnitDisc().MultiplyFloat64(c.lensRadius)
	offset := c.u.MultiplyFloat64(randomOnLens.X).Add(c.v.MultiplyFloat64(randomOnLens.Y))
	return &geometry.Ray{
		Origin: c.EyeLocation.AddVector(offset),
		Direction: c.lowerLeftCorner.AddVector(
			c.horizonal.MultiplyFloat64(u)).AddVector(
			c.verical.MultiplyFloat64(v)).From(
			c.EyeLocation).Subtract(
			offset).Unit(),
	}
}

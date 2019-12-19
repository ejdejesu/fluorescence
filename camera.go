package main

import (
	"fluorescence/geometry"
	"math"
	"math/rand"
)

// A Camera holds information about the scene's camera
// and facilitates the casting of Rays into the scene
type Camera struct {
	EyeLocation    geometry.Point  `json:"eye_location"`
	TargetLocation geometry.Point  `json:"target_location"`
	UpVector       geometry.Vector `json:"up_vector"`
	VerticalFOV    float64         `json:"vertical_fov"`
	AspectRatio    float64         `json:"aspect_ratio"`
	Aperture       float64         `json:"aperture"`
	FocusDistance  float64         `json:"focus_distance"`

	lensRadius float64
	theta      float64
	halfWidth  float64
	halfHeight float64

	w geometry.Vector
	u geometry.Vector
	v geometry.Vector

	lowerLeftCorner geometry.Point
	horizonal       geometry.Vector
	verical         geometry.Vector
}

// Setup is called after allocating the Camera struct and filling the exported fields
// It fills the unexported fields, such as derived vectors and measures
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

// GetRay returns a Ray from the eye location to a point on the view place u% across and v% up
func (c *Camera) GetRay(u float64, v float64, rng *rand.Rand) geometry.Ray {
	randomOnLens := geometry.RandomOnUnitDisc(rng).MultScalar(c.lensRadius)
	offset := c.u.MultScalar(randomOnLens.X).AddInPlace(c.v.MultScalar(randomOnLens.Y))
	return geometry.Ray{
		Origin: c.EyeLocation.AddVector(offset),
		Direction: c.lowerLeftCorner.AddVector(
			c.horizonal.MultScalar(u)).AddVector(
			c.verical.MultScalar(v)).From(
			c.EyeLocation).Sub(
			offset).Unit(),
	}
}

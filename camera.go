package main

import (
	"fluorescence/geometry"
	"math"
	"math/rand"

	"github.com/go-gl/mathgl/mgl64"
)

// A Camera holds information about the scene's camera
// and facilitates the casting of Rays into the scene
type Camera struct {
	EyeLocation    mgl64.Vec3 `json:"eye_location"`
	TargetLocation mgl64.Vec3 `json:"target_location"`
	UpVector       mgl64.Vec3 `json:"up_vector"`
	VerticalFOV    float64    `json:"vertical_fov"`
	AspectRatio    float64    `json:"aspect_ratio"`
	Aperture       float64    `json:"aperture"`
	FocusDistance  float64    `json:"focus_distance"`

	lensRadius float64
	theta      float64
	halfWidth  float64
	halfHeight float64

	w mgl64.Vec3
	u mgl64.Vec3
	v mgl64.Vec3

	lowerLeftCorner mgl64.Vec3
	horizonal       mgl64.Vec3
	verical         mgl64.Vec3
}

// Setup is called after allocating the Camera struct and filling the exported fields
// It fills the unexported fields, such as derived vectors and measures
func (c *Camera) Setup(p *Parameters) error {
	c.UpVector = c.UpVector.Normalize()
	c.AspectRatio = float64(p.ImageWidth) / float64(p.ImageHeight)

	c.lensRadius = c.Aperture / 2.0
	c.theta = c.VerticalFOV * math.Pi / 180.0
	c.halfHeight = math.Tan(c.theta / 2.0)
	c.halfWidth = c.AspectRatio * c.halfHeight

	c.w = c.EyeLocation.Sub(c.TargetLocation).Normalize()
	c.u = c.UpVector.Cross(c.w)
	c.v = c.w.Cross(c.u)

	c.lowerLeftCorner = c.EyeLocation.Sub(
		c.u.Mul(c.halfWidth * c.FocusDistance)).Sub(
		c.v.Mul(c.halfHeight * c.FocusDistance)).Sub(
		c.w.Mul(c.FocusDistance))

	c.horizonal = c.u.Mul(2.0 * c.halfWidth * c.FocusDistance)
	c.verical = c.v.Mul(2.0 * c.halfHeight * c.FocusDistance)

	return nil
}

// GetRay returns a Ray from the eye location to a point on the view place u% across and v% up
func (c *Camera) GetRay(u float64, v float64, rng *rand.Rand) geometry.Ray {
	randomOnLens := geometry.RandomOnUnitDisk(rng).Mul(c.lensRadius)
	offset := c.u.Mul(randomOnLens.X()).Add(c.v.Mul(randomOnLens.Y()))
	return geometry.Ray{
		Origin: c.EyeLocation.Add(offset),
		Direction: c.lowerLeftCorner.Add(
			c.horizonal.Mul(u)).Add(
			c.verical.Mul(v)).Sub(
			c.EyeLocation).Sub(
			offset).Normalize(),
	}
}

package cylinder

import (
	"fluorescence/geometry"
	"fluorescence/shading/material"
)

type infiniteCylinder struct {
	point     *geometry.Point
	direction *geometry.Vector
	radius    float64
	isCulled  bool
	mat       material.Material
}

type InfiniteCylinderData struct {
	Point     *geometry.Point
	Direction *geometry.Vector
	Radius    float64
	IsCulled  bool
}

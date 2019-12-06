package primitive

import (
	"fluorescence/geometry"
	"fluorescence/shading"
)

type Geometry interface {
	Intersection(*geometry.Ray, float64, float64) (*geometry.Ray, *geometry.Vector, float64, *shading.Color, bool)
}

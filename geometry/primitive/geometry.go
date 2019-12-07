package primitive

import (
	"fluorescence/geometry"
)

type Primitive interface {
	Intersection(*geometry.Ray, float64, float64) (*geometry.RayHit, bool)
}

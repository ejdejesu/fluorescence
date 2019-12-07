package geometry

import (
	"fluorescence/shading"
)

type RayHit struct {
	Ray         *Ray
	NormalAtHit *Vector
	T           float64
	Material    *shading.Color
}

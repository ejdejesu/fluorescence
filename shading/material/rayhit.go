package material

import "fluorescence/geometry"

type RayHit struct {
	Ray         *geometry.Ray
	NormalAtHit *geometry.Vector
	T           float64
	Material    Material
}

package texture

import "fluorescence/shading"

type Texture interface {
	Value(u, v float64) shading.Color
}

package texture

import (
	"github.com/go-gl/mathgl/mgl64"
)

// Color holds information about a solid-colored texture
type Color struct {
	Color mgl64.Vec3 `json:"color"`
}

// Value returns a color at a given texture coordinate
// this value is always the same, as the color is solid
func (ct *Color) Value(u, v float64) mgl64.Vec3 {
	return ct.Color
}

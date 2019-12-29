package texture

import (
	"github.com/go-gl/mathgl/mgl64"
)

// Texture defines behaviors of a Texture implementation
type Texture interface {
	Value(u, v float64) mgl64.Vec3
}

package shading

import (
	"image/color"
	"math"
)

// Color is a light abstraction on a Vector, with translations to and from
// various representations from the core color library
type Color struct {
	Red   float64 `json:"red"`
	Green float64 `json:"green"`
	Blue  float64 `json:"blue"`
}

// ToRGBA converts our Color into an RGBA representation from the color library
func (c Color) ToRGBA() color.RGBA {
	return color.RGBA{
		uint8(c.Red * float64(math.MaxUint8)),
		uint8(c.Green * float64(math.MaxUint8)),
		uint8(c.Blue * float64(math.MaxUint8)),
		uint8(1.0 * float64(math.MaxUint8))}
}

// ToRGBA64 converts our Color into an RGBA64 representation from the color library
func (c Color) ToRGBA64() color.RGBA64 {
	return color.RGBA64{
		uint16(c.Red * float64(math.MaxUint16)),
		uint16(c.Green * float64(math.MaxUint16)),
		uint16(c.Blue * float64(math.MaxUint16)),
		uint16(1.0 * float64(math.MaxUint16))}
}

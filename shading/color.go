package shading

import (
	"image/color"
	"math"
)

type Color struct {
	Red   float64 `json:"red"`
	Green float64 `json:"green"`
	Blue  float64 `json:"blue"`
	Alpha float64 `json:"alpha"`
}

var TRANSPARENT = Color{0.0, 0.0, 0.0, 0.0}

var WHITE = Color{1.0, 1.0, 1.0, 1.0}
var BLACK = Color{0.0, 0.0, 0.0, 1.0}

var RED = Color{1.0, 0.0, 0.0, 1.0}
var GREEN = Color{0.0, 1.0, 0.0, 1.0}
var BLUE = Color{0.0, 0.0, 1.0, 1.0}

var YELLOW = Color{1.0, 1.0, 0.0, 1.0}
var MAGENTA = Color{1.0, 0.0, 1.0, 1.0}
var CYAN = Color{0.0, 1.0, 1.0, 1.0}

func (c Color) ToRGBA() color.RGBA {
	return color.RGBA{
		uint8(c.Red * float64(math.MaxUint8)),
		uint8(c.Green * float64(math.MaxUint8)),
		uint8(c.Blue * float64(math.MaxUint8)),
		uint8(c.Alpha * float64(math.MaxUint8))}
}

func (c Color) ToRGBA64() color.RGBA64 {
	return color.RGBA64{
		uint16(c.Red * float64(math.MaxUint16)),
		uint16(c.Green * float64(math.MaxUint16)),
		uint16(c.Blue * float64(math.MaxUint16)),
		uint16(c.Alpha * float64(math.MaxUint16))}
}

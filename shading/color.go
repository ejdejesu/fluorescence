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

var COLOR_BLACK = Color{0.0, 0.0, 0.0}

func (c Color) Add(d Color) Color {
	return Color{c.Red + d.Red, c.Green + d.Green, c.Blue + d.Blue}
}

func (c Color) MultScalar(s float64) Color {
	return Color{c.Red * s, c.Green * s, c.Blue * s}
}

func (c Color) MultColor(d Color) Color {
	return Color{c.Red * d.Red, c.Green * d.Green, c.Blue * d.Blue}
}

func (c Color) DivScalar(s float64) Color {
	inv := 1.0 / s
	return Color{c.Red * inv, c.Green * inv, c.Blue * inv}
}

func (c Color) DivColor(d Color) Color {
	return Color{c.Red / d.Red, c.Green / d.Green, c.Blue / d.Blue}
}

func (c Color) Pow(e float64) Color {
	return Color{math.Pow(c.Red, e), math.Pow(c.Green, e), math.Pow(c.Blue, e)}
}

func (c Color) Clamp(min, max float64) Color {
	return Color{
		clamp(c.Red, min, max),
		clamp(c.Green, min, max),
		clamp(c.Blue, min, max)}
}

func clamp(val, min, max float64) float64 {
	if val < min {
		return min
	} else if val > max {
		return max
	}
	return val
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

func MakeColor(c color.Color) Color {
	// fmt.Println(c)
	r, g, b, _ := c.RGBA()
	// fmt.Println(r, g, b)
	inv := float64(1.0 / math.MaxUint16)
	// fmt.Println("red   ", float64(r)*inv)
	// fmt.Println("green ", float64(g)*inv)
	// fmt.Println("blue  ", float64(b)*inv)
	return Color{
		Red:   float64(r) * inv,
		Green: float64(g) * inv,
		Blue:  float64(b) * inv,
	}
}

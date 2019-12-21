package texture

import "fluorescence/shading"

type Color struct {
	Color shading.Color `json:"color"`
}

func (ct *Color) Value(u, v float64) shading.Color {
	return ct.Color
}

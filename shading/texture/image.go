package texture

import (
	"fluorescence/shading"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

type Image struct {
	FileName  string      `json:"image_file_name"`
	Gamma     float64     `json:"gamma"`
	Magnitude float64     `json:"magnitude"`
	Image     image.Image `json:"-"`
}

func (it *Image) Load() error {
	imageFile, err := os.Open(it.FileName)
	if err != nil {
		return err
	}
	if strings.HasSuffix(it.FileName, ".png") {
		it.Image, err = png.Decode(imageFile)
		if err != nil {
			return err
		}
	} else if strings.HasSuffix(it.FileName, ".jpg") || strings.HasSuffix(it.FileName, ".jpeg") {
		it.Image, err = jpeg.Decode(imageFile)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unknown image filetype (%s)", it.FileName)
	}
	return nil
}

func (it *Image) Value(u, v float64) shading.Color {
	x := int(u * float64(it.Image.Bounds().Dx()-1))
	y := int((1.0 - v) * float64(it.Image.Bounds().Dy()-1))
	color := it.Image.At(x, y)
	return shading.MakeColor(color).Pow(it.Gamma).MultScalar(it.Magnitude)
}

package texture

import (
	"fluorescence/geometry"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"github.com/go-gl/mathgl/mgl64"
)

// Image holds information about a texture based on an image
type Image struct {
	FileName  string      `json:"image_file_name"`
	Gamma     float64     `json:"gamma"`
	Magnitude float64     `json:"magnitude"`
	Image     image.Image `json:"-"`
}

// Load decodes the image from the given filename and performs other setup actions
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

// Value returns the color of the image at the given texture coordinates
// parameters u and v have a valid range [0.0, 1.0)
func (it *Image) Value(u, v float64) mgl64.Vec3 {
	// convert to image coordinates
	x := int(u * float64(it.Image.Bounds().Dx()-1))
	y := int((1.0 - v) * float64(it.Image.Bounds().Dy()-1))
	// get the color of the image at that point
	color := it.Image.At(x, y)
	// convert to a color, de-gamma, and apply magnitude
	return geometry.PowVec3(geometry.ColorToVec3(color), it.Gamma).Mul(it.Magnitude)
}

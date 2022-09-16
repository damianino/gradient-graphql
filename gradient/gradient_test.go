package gradient

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestImageCreation(t *testing.T) {
	f, err := os.Create("gradient.png")
	if err != nil {
		panic(err)
	}
	m := image.NewRGBA(image.Rect(0, 0, 256, 256))
	stops := []color.RGBA{
		color.RGBA{255, 37, 30, 255},
		color.RGBA{96, 255, 60, 255},

		color.RGBA{96, 31, 255, 255},
		color.RGBA{255, 37, 30, 255},
	}
	var fstops []FloatColor
	for _, stop := range stops {
		fstops = append(fstops, RGBAToFloat(stop))
	}
	Linear(m, fstops)

	png.Encode(f, m)
}

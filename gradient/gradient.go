package gradient

import (
	"image"
	"image/color"
	"math"
	"strings"

	"github.com/icza/gox/imagex/colorx"
)

type FloatColor struct {
	R,
	G,
	B,
	A float64
}

func CreateHex(x, y int, stops []string) (*image.RGBA, error) {
	var stopsColor []color.RGBA
	for _, x := range stops {
		x = strings.ReplaceAll(x, " ", "")
		c, err := colorx.ParseHexColor(x)
		if err != nil {
			return nil, err
		}
		stopsColor = append(stopsColor, c)
	}

	return Create(x, y, stopsColor)
}

func Create(x, y int, stops []color.RGBA) (*image.RGBA, error) {
	m := image.NewRGBA(image.Rect(0, 0, x, y))
	var fstops []FloatColor
	for _, stop := range stops {
		fstops = append(fstops, RGBAToFloat(stop))
	}
	Linear(m, fstops)

	return m, nil
}

func Linear(m *image.RGBA, stops []FloatColor) {
	x, y := m.Bounds().Dx(), m.Bounds().Dy()
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			float := lerpFloat(stops, float64(i)/float64(x))
			rgba := FloatToRGBA(float)
			m.SetRGBA(i, j, rgba)
		}
	}
}

func lerpFloat(stops []FloatColor, d float64) FloatColor {
	stopLength := 1 / float64(len(stops)-1)
	valueRatio := d / stopLength
	stopIndex := int(math.Floor(valueRatio))
	if stopIndex == len(stops)-1 {
		return stops[len(stops)-1]
	}
	_, stopFraction := math.Modf(valueRatio)

	return lerp(stops[stopIndex], stops[stopIndex+1], stopFraction)
}

func lerp(a, b FloatColor, d float64) FloatColor {
	return FloatColor{
		a.R + (b.R-a.R)*d,
		a.G + (b.G-a.G)*d,
		a.B + (b.B-a.B)*d,
		a.A + (b.A-a.A)*d,
	}
}

func RGBAToFloat(c color.RGBA) FloatColor {
	return FloatColor{
		float64(c.R) / 255.0,
		float64(c.G) / 255.0,
		float64(c.B) / 255.0,
		float64(c.A) / 255.0,
	}
}

func FloatToRGBA(c FloatColor) color.RGBA {
	return color.RGBA{
		uint8(c.R * 255),
		uint8(c.G * 255),
		uint8(c.B * 255),
		uint8(c.A * 255),
	}
}

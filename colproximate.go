package colproximate

import "image/color"
import "github.com/lucasb-eyer/go-colorful"

type Palette []colorful.Color

func (p Palette) Index(c color.Color) int {
	r, g, b, _ := c.RGBA()
	cf := colorful.LinearRgb(
		float64(r)/float64(0xffff),
		float64(g)/float64(0xffff),
		float64(b)/float64(0xffff),
	)

	index := 0
	minDistance := 9e9

	for i, candidate := range p {
		if distance := cf.DistanceCIE94(candidate); distance < minDistance {
			index = i
			minDistance = distance
		}
	}

	return index
}

func (p Palette) Convert(c color.Color) colorful.Color {
	return p[p.Index(c)]
}

func New(colors []color.Color) Palette {
	palette := make([]colorful.Color, len(colors))
	for i, c := range colors {
		r, g, b, _ := c.RGBA()
		palette[i] = colorful.LinearRgb(
			float64(r)/float64(0xffff),
			float64(g)/float64(0xffff),
			float64(b)/float64(0xffff),
		)
	}
	return palette
}

package raytracer

import (
	"image"
	"image/color"
	"math"
)

func Circle(i *image.NRGBA) error {
	width := i.Bounds().Max.X
	r := width / 2
	red := color.NRGBA{255, 0, 0, 255}
	green := color.NRGBA{0, 255, 0, 255}

	for y := 0; y < width; y++ {
		for x := 0; x < width; x++ {
			c := red
			if withInCircle(x, y, r) {
				c = green
			}

			i.Set(x, y, c)
		}
	}

	return nil
}

func withInCircle(x, y, r int) bool {
	x -= r
	y -= r

	d := math.Sqrt(float64(x*x + y*y))

	return d < float64(r)
}

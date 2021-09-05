package raytracing

import (
	"image"
	"image/color"
)

func Circle(i *image.RGBA) error {
	w := i.Bounds().Max.X
	h := i.Bounds().Max.Y

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			i.Set(x, y, color.NRGBA{uint8(y % 256), uint8(x % 256), uint8(y % 512), uint8(x % 256)})
		}
	}

	return nil
}

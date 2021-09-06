package raytracing

import (
	"image"
	"image/color"
)

type Raytracer struct {
	Callback func()
	Image    *image.RGBA
}

func NewRaytracer() *Raytracer {
	return &Raytracer{}
}

func (r *Raytracer) Run() {
	w := r.Image.Bounds().Max.X
	h := r.Image.Bounds().Max.Y

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r.Image.Set(x, y, color.NRGBA{uint8(y % 256), uint8(x % 256), uint8(y % 512), uint8(x % 256)})
		}
	}

	r.Callback()
}

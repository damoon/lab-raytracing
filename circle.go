package raytracing

import (
	"image"
	"image/color"
)

func Circle(i *image.RGBA) error {
	w := i.Bounds().Max.X
	h := i.Bounds().Max.Y
	// red := color.NRGBA{255, 0, 0, 255}
	// green := color.NRGBA{0, 255, 0, 255}

	//for {
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			//time.Sleep(time.Nanosecond)
			i.Set(x, y, color.NRGBA{uint8(y % 256), uint8(x % 256), uint8(y % 512), uint8(x % 256)})
		}
	}

	// printed = false

	// for y := 0; y < h; y++ {
	// 	for x := 0; x < w; x++ {
	// 		time.Sleep(time.Nanosecond)
	// 		i.Set(x, y, color.NRGBA{uint8(y % 256), uint8(x % 256), 0, 255})
	// 	}
	// }
	//}

	return nil
}

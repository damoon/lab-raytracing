package encode

import (
	"image"
	"image/png"
)

func WritePNG(path string, img image.Image) error {
	return WriteImage(path, img, png.Encode)
}

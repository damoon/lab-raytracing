package encode

import (
	"image"

	"github.com/lmittmann/ppm"
)

func WritePPM(path string, img image.Image) error {
	return WriteImage(path, img, ppm.Encode)
}

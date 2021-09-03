package encode

import (
	"image"
	"io"

	"github.com/lmittmann/ppm"
)

func PPM(w io.Writer, img image.Image) {
	ppm.Encode(w, img)
}

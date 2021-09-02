package raytracer

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func WritePNG(file string, img image.Image) error {
	f, err := os.Create(file)
	if err != nil {
		return fmt.Errorf("open image destination: %v", err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		return fmt.Errorf("encode image: %v", err)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("close file handle for image: %v", err)
	}

	return nil
}

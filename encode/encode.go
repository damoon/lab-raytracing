package encode

import (
	"fmt"
	"image"
	"io"
	"os"
)

func WriteImage(path string, img image.Image, enc func(w io.Writer, img image.Image) error) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("open image destination: %v", err)
	}
	defer f.Close()

	err = enc(f, img)
	if err != nil {
		return fmt.Errorf("encode image: %v", err)
	}

	err = f.Close()
	if err != nil {
		return fmt.Errorf("close file handle for image: %v", err)
	}

	return nil
}

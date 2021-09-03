package raytracer

import (
	"fmt"
	"image"
	"log"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Scene struct {
	Img *image.RGBA
	tex *sdl.Texture
}

func (g *Scene) Destroy() {

}

func NewScene(r *sdl.Renderer) (*Scene, error) {
	w := r.GetViewport().W
	h := r.GetViewport().H

	tex, err := r.CreateTexture(
		sdl.PIXELFORMAT_ABGR8888,
		sdl.TEXTUREACCESS_STATIC,
		w,
		h,
	)
	if err != nil {
		return nil, fmt.Errorf("create texture: %v", err)
	}

	return &Scene{
		Img: image.NewRGBA(image.Rect(0, 0, int(w), int(h))),
		tex: tex,
	}, nil
}

func (s *Scene) Run(events <-chan sdl.Event, r *sdl.Renderer) <-chan error {
	errc := make(chan error)

	go func() {
		defer close(errc)

		ticker := time.NewTicker(time.Second / 30)

		for {
			select {
			case e := <-events:
				if e == nil {
					continue
				}

				done := s.handleEvent(e)
				if done {
					log.Print("done")
					sdl.PushEvent(e)
					return
				}

			case <-ticker.C:
				s.update()

				err := s.paint(r)
				if err != nil {
					errc <- err
					return
				}
			}
		}
	}()

	return errc
}

func (s *Scene) handleEvent(event sdl.Event) bool {
	switch t := event.(type) {
	case *sdl.QuitEvent:
		return true

	case *sdl.KeyboardEvent:
		if t.State == 1 && t.Keysym.Sym == 27 {
			// Escape was pressed
			return true
		}

	case *sdl.AudioDeviceEvent, *sdl.TextEditingEvent, *sdl.MouseButtonEvent, *sdl.MouseMotionEvent,
		*sdl.WindowEvent, *sdl.TouchFingerEvent, *sdl.CommonEvent:

	default:
		log.Printf("unknown event %T", event)
	}

	return false
}

func (s *Scene) update() {
	// TODO
}

var printed bool

func (s *Scene) paint(r *sdl.Renderer) error {
	// log.Print("paint")
	err := r.Clear()
	if err != nil {
		return fmt.Errorf("clean renderer: %v", err)
	}

	if printed {
		return nil
	}

	log.Print("paint")

	err = s.tex.Update(
		nil,
		s.Img.Pix,
		s.Img.Stride,
	)
	if err != nil {
		return fmt.Errorf("update texture with image: %v", err)
	}

	// rect := &sdl.Rect{X: 0, Y: 0, W: int32(s.Img.Bounds().Dx()), H: int32(s.Img.Bounds().Dy())}
	rect := r.GetViewport()

	err = r.Copy(s.tex, &rect, &rect)
	if err != nil {
		return fmt.Errorf("copy texture in renderer: %v", err)
	}

	r.Present()

	//printed = true

	return nil
}

package ui

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"

	"github.com/damoon/lab-raytracing/raytracing"
	"github.com/oakmound/oak/v3/shiny/driver"
	"github.com/oakmound/oak/v3/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

func Window(w, h int) error {
	var returnErr error
	var changed bool
	var sz size.Event

	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(screen.WindowGenerator{
			Title:  "Example",
			Width:  w,
			Height: h,
		})
		if err != nil {
			returnErr = err
			return
		}
		defer w.Release()

		repaint := func() { w.Send(paint.Event{}) }

		img, err := s.NewImage(sz.Size())
		if err != nil {
			returnErr = err
			return
		}

		raytracer := raytracing.Raytracer{
			Callback: repaint,
			Image:    img.RGBA(),
		}
		go raytracer.Run()

		for {
			e := w.NextEvent()

			// This print message is to help programmers learn what events this
			// example program generates. A real program shouldn't print such
			// messages; they're not important to end users.
			format := "got %#v\n"
			if _, ok := e.(fmt.Stringer); ok {
				format = "got %v\n"
			}
			fmt.Printf(format, e)

			switch e := e.(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}

				if changed {
					repaint := func() { w.Send(paint.Event{}) }

					img, err = s.NewImage(sz.Size())
					if err != nil {
						returnErr = err
						return
					}

					raytracer = raytracing.Raytracer{
						Callback: repaint,
						Image:    img.RGBA(),
					}
					go raytracer.Run()

					changed = false
				}

			case key.Event:
				if e.Code == key.CodeEscape {
					return
				}

			case paint.Event:
				if sz.Size().X != img.Size().X || sz.Size().Y != img.Size().Y {
					w.Fill(sz.Bounds(), color.Black, draw.Src)
				}

				w.Upload(image.Point{0, 0}, img, img.Bounds())
				w.Publish()

			case size.Event:
				sz = e
				changed = true

			case error:
				returnErr = err
				return
			}
		}
	})

	return returnErr
}

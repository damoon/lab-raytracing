package ui

import (
	"fmt"
	"image"
	"log"

	"github.com/damoon/lab-raytracing/raytracing"
	"github.com/oakmound/oak/v3/shiny/driver"
	"github.com/oakmound/oak/v3/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

func Window() error {
	var err error

	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(screen.WindowGenerator{
			Title:  "Example",
			Width:  1280,
			Height: 720,
		})
		if err != nil {
			log.Fatal(err)
		}
		defer w.Release()

		size0 := image.Point{1280, 720}
		b, err := s.NewImage(size0)
		if err != nil {
			log.Fatal(err)
		}
		defer b.Release()

		err = raytracing.Circle(b.RGBA())
		if err != nil {
			log.Fatal(err)
		}

		// var sz size.Event
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

			case key.Event:
				if e.Code == key.CodeEscape {
					return
				}

			case paint.Event:
				w.Upload(image.Point{0, 0}, b, b.Bounds())
				w.Publish()

			case size.Event:
				b, err = s.NewImage(e.Size())
				if err != nil {
					log.Fatal(err)
				}
				defer b.Release()

				err = raytracing.Circle(b.RGBA())
				if err != nil {
					log.Fatal(err)
				}

				w.Upload(image.Point{0, 0}, b, b.Bounds())
				w.Publish()

			case error:
				log.Print(e)
			}
		}
	})

	return err
}

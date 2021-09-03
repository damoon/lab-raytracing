package ui

import (
	"fmt"
	"log"
	"runtime"

	"github.com/urfave/cli"
	"github.com/veandco/go-sdl2/sdl"
)

func window(c *cli.Context) error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("could not initialize SDL: %v", err)
	}
	defer sdl.Quit()

	w, r, err := sdl.CreateWindowAndRenderer(1280, 720, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("could not create window: %v", err)
	}
	defer w.Destroy()

	s, err := raytracing.NewScene(r)
	if err != nil {
		return fmt.Errorf("create scene: %v", err)
	}

	log.Print("create circle")
	go raytracing.Circle(s.Img)
	log.Print("create circle done")

	events := make(chan sdl.Event)
	errc := s.Run(events, r)

	runtime.LockOSThread()
	for {
		select {
		case events <- sdl.WaitEvent():
		case err := <-errc:
			return err
		}
	}
}

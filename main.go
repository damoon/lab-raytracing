package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/damoon/lab-raytracing/encode"
	"github.com/damoon/lab-raytracing/raytracing"
	"github.com/damoon/lab-raytracing/ui"
	"github.com/joho/godotenv"
	"github.com/pkg/profile"
	"github.com/urfave/cli/v2"
)

func main() {
	err := run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	err := loadEnv(".env")
	if err != nil {
		return fmt.Errorf("load config: %v", err)
	}

	err = app().Run(args)
	if err != nil {
		return fmt.Errorf("run application: %v", err)
	}

	return nil
}

func loadEnv(envFile string) error {
	_, err := os.Stat(envFile)
	if os.IsNotExist(err) {
		return nil
	}

	err = godotenv.Load(envFile)
	if err != nil {
		return fmt.Errorf("load .env file: %v", err)
	}

	return nil
}

func app() *cli.App {
	return &cli.App{
		Name:  "raytracing",
		Usage: "a simple raytracer",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "file to output to",
				EnvVars: []string{"RT_OUTPUT"},
				Value:   "image.png",
			},
			&cli.BoolFlag{
				Name:    "window",
				Aliases: []string{"ui"},
				Usage:   "open window",
				EnvVars: []string{"RT_WINDOW"},
				Value:   true,
			},
			&cli.BoolFlag{
				Name:    "cpu-profile",
				Aliases: []string{"cpu"},
				Usage:   "profile cpu usage",
				EnvVars: []string{"RT_CPU_PROFILE"},
			},
			&cli.BoolFlag{
				Name:    "mem-profile",
				Aliases: []string{"mem"},
				Usage:   "profile memory usage",
				EnvVars: []string{"RT_MEM_PROFILE"},
			},
			&cli.BoolFlag{
				Name:    "trace-profile",
				Aliases: []string{"trace"},
				Usage:   "create a trace profile",
				EnvVars: []string{"RT_TRACE_PROFILE"},
			},
			&cli.IntFlag{
				Name:    "width",
				Aliases: []string{"w"},
				Usage:   "pixel width of the image",
				EnvVars: []string{"RT_IMAGE_WIDTH"},
				Value:   1280,
			},
			&cli.IntFlag{
				Name: "height",
				// Aliases: []string{"h"},
				Usage:   "pixel height of the image",
				EnvVars: []string{"RT_IMAGE_HEIGHT"},
				Value:   720,
			},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("cpu-profile") {
				defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
			}
			if c.Bool("mem-profile") {
				defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath(".")).Stop()
			}
			if c.Bool("trace-profile") {
				defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
			}

			w := c.Int("width")
			h := c.Int("height")

			openUI := c.Bool("window")
			if openUI {
				return ui.Window(w, h)
			}

			path := c.String("output")
			return directToFile(w, h, path)
		},
	}
}

func directToFile(w, h int, path string) error {
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	ch := make(chan interface{})
	done := func() {
		ch <- struct{}{}
	}

	raytracer := raytracing.Raytracer{
		Callback: done,
		Image:    img,
	}
	go raytracer.Run()

	<-ch

	err := encode.WritePNG(path, img)
	if err != nil {
		return err
	}

	return nil
}

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
				Usage:   "profile cpu usage",
				EnvVars: []string{"RT_CPU_PROFILE"},
				Value:   false,
			},
			&cli.IntFlag{
				Name:    "radius",
				Aliases: []string{"r"},
				Usage:   "set the radius of the circle",
				EnvVars: []string{"RT_CIRCLE_RADIUS"},
				Value:   100,
			},
		},
		Action: func(c *cli.Context) error {
			if c.Bool("cpu-profile") {
				defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
			}

			openUI := c.Bool("window")
			if openUI {
				return ui.Window()
			}

			r := c.Int("radius")
			path := c.String("output")
			return directToFile(r, path)
		},
	}
}

func directToFile(r int, path string) error {
	img := image.NewRGBA(image.Rect(0, 0, 2*r, 2*r))

	err := raytracing.Circle(img)
	if err != nil {
		return err
	}

	err = encode.WritePNG(path, img)
	if err != nil {
		return err
	}

	return nil
}

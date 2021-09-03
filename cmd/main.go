package main

import (
	"fmt"
	"image"
	"log"
	"os"

	raytracing "github.com/damoon/lab-raytracing"
	"github.com/joho/godotenv"
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
		Name:  "app",
		Usage: "a nice application",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "file to output to",
				EnvVars: []string{"APP_OUTPUT"},
				Value:   "image.png",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "circle",
				Usage: "draw a circle",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "radius",
						Aliases: []string{"r"},
						Usage:   "set the radius of the circle",
						EnvVars: []string{"APP_CIRCLE_RADIUS"},
						Value:   100,
					},
				},
				Action: circle,
			},
			{
				Name:   "window",
				Usage:  "open in a window",
				Action: window,
			},
		},
	}
}

func circle(c *cli.Context) error {
	r := c.Int("radius")
	file := c.String("output")
	img := image.NewRGBA(image.Rect(0, 0, 2*r, 2*r))

	err := raytracing.Circle(img)
	if err != nil {
		return err
	}

	err = raytracing.WritePNG(file, img)
	if err != nil {
		return err
	}

	return nil
}

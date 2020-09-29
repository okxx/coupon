package cmd

import (
	"coupon/internal/utils/app"
	"github.com/urfave/cli/v2"
)

var (
	Port = "7077"

	Run = &cli.Command{
		Name: "run",
		Usage: "run the web program",
		Action: runController,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "port, p",
				Value: Port,
				Usage: "Port number",
				DefaultText: Port, // default port number
			},
		},
	}
)

func runController(c *cli.Context) error {

	coupon := app.New()

	coupon.Get("/", func(c *app.C) {

	})

	return coupon.Run(Port)
}

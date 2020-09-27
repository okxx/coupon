package cmd

import (
	"github.com/urfave/cli/v2"
)

// install command
var Install = &cli.Command{
	Name: "install",
	Usage: "[command] coupon install",
	Action: installController,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "port, p",
			Value: defaultPort,
			Usage: "Port number",
			DefaultText: defaultPort,
		},
	},
}

// install controller
func installController(c *cli.Context) error {

	return nil
}
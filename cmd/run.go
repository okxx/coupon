package cmd

import (
	"github.com/urfave/cli/v2"
	"net/http"
)

var (
	defaultPort = "9009"

	Run = &cli.Command{
		Name: "run",
		Usage: "run the web program",
		Action: runController,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "port, p",
				Value: defaultPort,
				Usage: "Port number",
				DefaultText: defaultPort, // default port number
			},
		},
	}
)

func runController(c *cli.Context) error {

	// static {css,js}
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("template/static"))))

	return nil
}
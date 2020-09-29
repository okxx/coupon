package main

import (
	"coupon/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app 			:= cli.NewApp()
	app.Name 		= "Coupon"
	app.Version 	= "1.0.0.x0e"
	app.Usage 		= "Use the built-in coupon system in go language"
	app.Commands 	= []*cli.Command{
		cmd.Run,
	}
	log.Fatal(app.Run(os.Args))
}

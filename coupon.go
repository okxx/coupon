package main

import (
	"coupon/cmd"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const (
	Name	 	= "Coupon"
	Version 	= "1.0.0.x0e"
	Usage		= "Use the built-in coupon system in go language"
)

func main() {
	app 			:= cli.NewApp()
	app.Name 		= Name
	app.Usage 		= Usage
	app.Version 	= Version
	app.Commands 	= []*cli.Command{
		cmd.Run,
		cmd.Install,
	}
	log.Fatal(app.Run(os.Args))
}

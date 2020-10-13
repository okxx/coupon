package cmd

import (
	"coupon/internal/handles"
	"coupon/internal/middleware"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/urfave/cli/v2"
	"os"
)

var (
	Port = "7077"

	Run = &cli.Command{
		Name: "run",
		Usage: "run the web program",
		Action: mainApplication,
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

func newCoupon() *fiber.App {
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:     "${time} ${pid} ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "15:04:05",
		TimeZone:   "Local",
		Output:     os.Stdout,
	}))
	app.Use(recover.New())
	return app
}

func mainApplication(c *cli.Context) error {

	// application
	r := newCoupon()

	//welcome
	r.Get("/", handles.Welcome)

	// use gateway
	r.Use(middleware.Gateway)

	// coupon
	coupon := r.Group("/coupon", func(ctx *fiber.Ctx) error { return ctx.Next() })
	coupon.Get("/list", handles.PaginateCoupon)
	coupon.Post("/create", handles.CreateCoupon)

	//platform
	platform := r.Group("/platform", func(ctx *fiber.Ctx) error { return ctx.Next() })
	platform.Get("/list", handles.PaginatePlatform)
	platform.Post("/create", handles.CreatePlatform)
	platform.Post("/edit", handles.EditPlatform)

	return r.Listen(fmt.Sprintf(":%v",Port))
}

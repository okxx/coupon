package middleware

import "github.com/gofiber/fiber/v2"

func Gateway(ctx *fiber.Ctx) error {

	return ctx.Next()
}
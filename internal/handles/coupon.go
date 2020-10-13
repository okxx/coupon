package handles

import (
	"github.com/gofiber/fiber/v2"
)

func PaginateCoupon(ctx *fiber.Ctx) error {

	return ctx.Send([]byte("hello coupon list"))
}


func CreateCoupon(ctx *fiber.Ctx) error {


	return nil
}
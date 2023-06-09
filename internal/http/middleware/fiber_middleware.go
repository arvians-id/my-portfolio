package middleware

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

func ExposeFiberContext() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Locals("fiberContext", ctx)
		return ctx.Next()
	}
}

func FiberContext(ctx context.Context) *fiber.Ctx {
	return ctx.Value("fiberContext").(*fiber.Ctx)
}

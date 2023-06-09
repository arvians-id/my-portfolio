package middleware

import (
	"context"
	"github.com/arvians-id/go-portfolio/cmd/config"
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

func XApiKey(configuration config.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		xApiKey := ctx.Get("X-API-KEY")
		if xApiKey == "" {
			return fiber.NewError(fiber.StatusForbidden, "access denied: please provide a valid API key to access this page.")
		}

		if xApiKey != configuration.Get("X_API_KEY") {
			return fiber.NewError(fiber.StatusForbidden, "invalid key: the provided API key is incorrect. please make sure to use a valid API key to access this resource.")
		}

		return ctx.Next()
	}
}

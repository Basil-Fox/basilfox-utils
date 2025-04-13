package middleware

import (
	"github.com/FiberApps/common-library/constants/header"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func RequestID(ctx *fiber.Ctx) error {
	// Get Request ID from headers
	if id := ctx.Get(header.RequestID); id != "" {
		return ctx.Next()
	}

	// Generate and Set new UUIDv4
	ctx.Request().Header.Set(header.RequestID, utils.UUIDv4())

	return ctx.Next()
}

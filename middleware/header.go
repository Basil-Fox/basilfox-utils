package middleware

import (
	core "github.com/FiberApps/core-service-utils"
	"github.com/gofiber/fiber/v2"
)

func ValidateHeaders(isUserRequired bool) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		// Check Namespace Header
		if c.Get(core.HeaderNamespace) == "" {
			return core.SendError(c, fiber.ErrBadGateway)
		}

		// Check UserID Header
		if c.Get(core.HeaderUserId) == "" && isUserRequired {
			return core.SendError(c, fiber.ErrBadGateway)
		}

		return c.Next()
	}
}

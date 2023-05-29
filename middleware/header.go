package middleware

import (
	"github.com/FiberApps/core/constant"
	"github.com/FiberApps/core/utils/response"
	"github.com/gofiber/fiber/v2"
)

func ValidateHeaders(isUserRequired bool) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		// Check Namespace Header
		if c.Get(constant.HeaderNamespace) == "" {
			return response.SendError(c, fiber.ErrBadGateway)
		}

		// Check UserID Header
		if c.Get(constant.HeaderUserId) == "" && isUserRequired {
			return response.SendError(c, fiber.ErrBadGateway)
		}

		return c.Next()
	}
}

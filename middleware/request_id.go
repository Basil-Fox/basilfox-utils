package middleware

import (
	"github.com/FiberApps/common-library/constant"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func RequestID(c *fiber.Ctx) error {
	// Get Request ID from headers
	if id := c.Get(constant.HeaderRequestId); id != "" {
		return c.Next()
	}

	// Generate and Set new UUIDv4
	c.Request().Header.Set(constant.HeaderRequestId, utils.UUIDv4())

	return c.Next()
}

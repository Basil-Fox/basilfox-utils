package middleware

import (
	"github.com/FiberApps/core/constant"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func RequestID(c *fiber.Ctx) error {

	// Get or Generate new UUID
	id := c.Get(constant.HeaderRequestId)
	if id == "" {
		id = utils.UUIDv4()
	}

	// Set RequestID to headers
	c.Request().Header.Set(constant.HeaderRequestId, id)

	// Add the RequestID to locals
	c.Locals(constant.ContextRequestId, id)

	return c.Next()
}

package middleware

import (
	core "github.com/FiberApps/core-service-utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func RequestID(c *fiber.Ctx) error {

	// Generate new UUID
	id := utils.UUIDv4()

	// Set RequestID to headers
	c.Request().Header.Set(core.HeaderRequestId, id)

	// Add the RequestID to locals
	c.Locals(core.ContextRequestId, id)

	return c.Next()
}

package middleware

import (
	"github.com/FiberApps/common-library/constant"
	"github.com/FiberApps/common-library/response"
	"github.com/gofiber/fiber/v2"
)

// Helper function to check if a string exists in a slice
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func ValidateHeaders(endpointType string, namespaces []string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		// Check Namespace Header
		namespaceHeader := c.Get(constant.HeaderNamespace)
		if namespaceHeader == "" || !contains(namespaces, namespaceHeader) {
			return response.SendError(c, fiber.ErrBadGateway)
		}

		// Set namespace to locals
		c.Locals(constant.ContextNamespace, namespaceHeader)

		// Check Firebase UserID Header for authenticated routes
		if endpointType == constant.EndpointPrivate {
			if c.Get(constant.HeaderFirebaseUserId) == "" {
				return response.SendError(c, fiber.ErrUnauthorized)
			}

			// Set firebase_user_id to locals
			c.Locals(constant.ContextFirebaseUserId, c.Get(constant.HeaderFirebaseUserId))
		}

		return c.Next()
	}
}

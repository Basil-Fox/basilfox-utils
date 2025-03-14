package middleware

import (
	"slices"

	"github.com/FiberApps/common-library/constant"
	"github.com/FiberApps/common-library/response"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateHeaders(endpointType constant.EndpointType, namespaces []string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		// Check namespace header is present and valid
		namespaceHeader := c.Get(constant.HeaderNamespace)
		if namespaceHeader == "" || !slices.Contains(namespaces, namespaceHeader) {
			return response.SendError(c, fiber.ErrBadRequest)
		}

		// Check if UserID Header is present
		if endpointType == constant.EndpointPrivate {
			userID := c.Get(constant.HeaderUserId)
			if userID == "" {
				return response.SendError(c, fiber.ErrUnauthorized)
			}

			// Validate UserID Header
			if _, err := primitive.ObjectIDFromHex(userID); err != nil {
				return response.SendError(c, fiber.ErrUnauthorized)
			}
		}

		return c.Next()
	}
}

package middleware

import (
	"github.com/FiberApps/common-library/constant"
	"github.com/FiberApps/common-library/response"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

		if endpointType == constant.EndpointPrivate || endpointType == constant.EndpointRefresh {
			// Check UserID Header
			if c.Get(constant.HeaderUserId) == "" {
				return response.SendError(c, fiber.ErrUnauthorized)
			}

			// ObjectID validation
			_, err := primitive.ObjectIDFromHex(c.Get(constant.HeaderUserId))
			if err != nil {
				return response.SendError(c, fiber.ErrUnauthorized)
			}

			// Set user_id to locals
			c.Locals(constant.ContextUserId, c.Get(constant.HeaderUserId))

			// Check TokenID Header
			if c.Get(constant.HeaderTokenId) == "" {
				return response.SendError(c, fiber.ErrUnauthorized)
			}

			// Set token_id to locals
			c.Locals(constant.ContextTokenId, c.Get(constant.HeaderTokenId))

			// Check TokenKind Header
			if c.Get(constant.HeaderTokenKind) == "" {
				return response.SendError(c, fiber.ErrUnauthorized)
			}

			// Set token_kind to locals
			c.Locals(constant.ContextTokenKind, c.Get(constant.HeaderTokenKind))

			// Check if AccessToken is used for authentication
			if endpointType == constant.EndpointPrivate && c.Get(constant.HeaderTokenKind) != constant.TokenTypeAccess {
				return response.SendError(c, fiber.ErrUnauthorized)
			}

			// Check if RefreshToken is used for authentication
			if endpointType == constant.EndpointRefresh && c.Get(constant.HeaderTokenKind) != constant.TokenTypeRefresh {
				return response.SendError(c, fiber.ErrUnauthorized)
			}
		}

		return c.Next()
	}
}

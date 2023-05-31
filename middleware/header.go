package middleware

import (
	"github.com/FiberApps/core/constant"
	"github.com/FiberApps/core/response"
	"github.com/gofiber/fiber/v2"
)

func ValidateHeaders(endpointType string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		// Check Namespace Header
		if c.Get(constant.HeaderNamespace) == "" {
			return response.SendError(c, fiber.ErrBadGateway)
		}

		if endpointType == constant.EndpointPrivate || endpointType == constant.EndpointRefresh {
			// Check UserID Header
			if c.Get(constant.HeaderUserId) == "" {
				return response.SendError(c, fiber.ErrBadGateway)
			}

			// Check TokenID Header
			if c.Get(constant.HeaderTokenId) == "" {
				return response.SendError(c, fiber.ErrBadGateway)
			}

			// Check TokenKind Header
			if c.Get(constant.HeaderTokenKind) == "" {
				return response.SendError(c, fiber.ErrBadGateway)
			}

			if endpointType == constant.EndpointPrivate && c.Get(constant.HeaderTokenKind) != constant.TokenTypeAccess {
				return response.SendError(c, fiber.ErrUnauthorized)
			}

			if endpointType == constant.EndpointRefresh && c.Get(constant.HeaderTokenKind) != constant.TokenTypeRefresh {
				return response.SendError(c, fiber.ErrUnauthorized)
			}
		}

		return c.Next()
	}
}

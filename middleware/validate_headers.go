package middleware

import (
	"slices"

	"github.com/FiberApps/common-library/constants/config"
	"github.com/FiberApps/common-library/constants/header"
	"github.com/FiberApps/common-library/response"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateHeaders(route config.RouteType, namespaces []string) func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {

		// Check namespace header is present and valid
		namespace := ctx.Get(header.Namespace)
		if namespace == "" || !slices.Contains(namespaces, namespace) {
			return response.SendErrorWithData(ctx, fiber.StatusBadRequest, "Invalid namespace")
		}

		// Check if UserID Header is present
		if route == config.RoutePrivate {
			userIdHex := ctx.Get(header.UserID)
			if userIdHex == "" {
				return response.SendErrorWithData(ctx, fiber.StatusUnauthorized, "Missing authentication context")
			}

			// Validate UserID Header
			userID, err := primitive.ObjectIDFromHex(userIdHex)
			if err != nil {
				return response.SendErrorWithData(ctx, fiber.StatusBadRequest, "Invalid authentication context")
			}

			// Set UserID in context
			ctx.Locals(config.ContextUserID, userID)
		}

		if route == config.RouteRegister || route == config.RoutePrivate {
			// Check if FirebaseUID Header is present
			firebaseUID := ctx.Get(header.FirebaseUID)
			if firebaseUID == "" {
				return response.SendErrorWithData(ctx, fiber.StatusUnauthorized, "Missing authentication context")
			}
			// Set FirebaseUID in context
			ctx.Locals(config.ContextFirebaseUID, firebaseUID)
		}

		return ctx.Next()
	}
}

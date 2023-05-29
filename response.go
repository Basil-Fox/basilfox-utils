package lib

import (
	"github.com/gofiber/fiber/v2"
)

// XXX - Error with Message
func SendError(ctx *fiber.Ctx, err *fiber.Error) error {
	return ctx.Status(err.Code).JSON(fiber.Map{
		"Error":     err.Message,
		"RequestID": ctx.Get(fiber.HeaderXRequestID),
	})
}

// XXX - Error with Data
func SendErrorWithData(ctx *fiber.Ctx, code int, data map[string]interface{}) error {
	return ctx.Status(code).JSON(fiber.Map{
		"Error":     data,
		"RequestID": ctx.Get(fiber.HeaderXRequestID),
	})
}

// 2XX - Success
func Success(ctx *fiber.Ctx, data map[string]interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Data":      data,
		"RequestID": ctx.Get(fiber.HeaderXRequestID),
	})
}

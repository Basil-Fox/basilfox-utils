package lib

import (
	"github.com/gofiber/fiber/v2"
)

// XXX - Generic Error
func SendError(ctx *fiber.Ctx, err *fiber.Error) error {
	return ctx.Status(err.Code).JSON(fiber.Map{
		"Error":     err.Message,
		"RequestID": ctx.Get(REQUEST_ID),
	})
}

// 200 - Success
func Success(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Data":      data,
		"RequestID": ctx.Get(REQUEST_ID),
	})
}

// 500 - Internal Server Error
func InternalServerError(ctx *fiber.Ctx, err string) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"Error":     err,
		"RequestID": ctx.Get(REQUEST_ID),
	})
}

// 400 - Bad Request
func BadRequest(ctx *fiber.Ctx, err string) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Error":     err,
		"RequestID": ctx.Get(REQUEST_ID),
	})
}

// 401 - Unauthorized
func Unauthorized(ctx *fiber.Ctx, err string) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"Error":     err,
		"RequestID": ctx.Get(REQUEST_ID),
	})
}

// 404 - Not Found
func NotFound(ctx *fiber.Ctx, err string) error {
	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"Error":     err,
		"RequestID": ctx.Get(REQUEST_ID),
	})
}

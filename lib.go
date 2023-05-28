package lib

import (
	"github.com/gofiber/fiber/v2"
)

// XXX - Generic Error
func SendError(ctx *fiber.Ctx, err *fiber.Error) error {
	return ctx.Status(err.Code).JSON(fiber.Map{
		"Error":     err.Message,
		"RequestID": ctx.Get(HeaderRequestId),
	})
}

// XXX - Generic Error with Data
func SendErrorWithData(ctx *fiber.Ctx, code int, data interface{}) error {
	return ctx.Status(code).JSON(fiber.Map{
		"Error":     data,
		"RequestID": ctx.Get(HeaderRequestId),
	})
}

// 200 - Success
func Success(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Data":      data,
		"RequestID": ctx.Get(HeaderRequestId),
	})
}

// 500 - Internal Server Error
func InternalServerError(ctx *fiber.Ctx, err string) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"Error":     err,
		"RequestID": ctx.Get(HeaderRequestId),
	})
}

// 400 - Bad Request
func BadRequest(ctx *fiber.Ctx, err string) error {
	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"Error":     err,
		"RequestID": ctx.Get(HeaderRequestId),
	})
}

// 401 - Unauthorized
func Unauthorized(ctx *fiber.Ctx, err string) error {
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"Error":     err,
		"RequestID": ctx.Get(HeaderRequestId),
	})
}

// 404 - Not Found
func NotFound(ctx *fiber.Ctx, err string) error {
	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"Error":     err,
		"RequestID": ctx.Get(HeaderRequestId),
	})
}

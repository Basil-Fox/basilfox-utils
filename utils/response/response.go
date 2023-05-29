package response

import (
	"github.com/FiberApps/core/constant"
	"github.com/gofiber/fiber/v2"
)

// XXX - Error with Message
func SendError(ctx *fiber.Ctx, err *fiber.Error) error {
	return ctx.Status(err.Code).JSON(fiber.Map{
		"Error":     err.Message,
		"RequestID": ctx.Get(constant.HeaderRequestId),
	})
}

// XXX - Error with Data
func SendErrorWithData(ctx *fiber.Ctx, code int, data interface{}) error {
	return ctx.Status(code).JSON(fiber.Map{
		"Error":     data,
		"RequestID": ctx.Get(constant.HeaderRequestId),
	})
}

// 2XX - Success
func Success(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Data":      data,
		"RequestID": ctx.Get(constant.HeaderRequestId),
	})
}

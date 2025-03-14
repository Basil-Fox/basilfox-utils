package response

import (
	"github.com/FiberApps/common-library/constant"
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
func SendErrorWithData(ctx *fiber.Ctx, code int, data any) error {
	return ctx.Status(code).JSON(fiber.Map{
		"Error":     data,
		"RequestID": ctx.Get(constant.HeaderRequestId),
	})
}

// 2XX - Success
func Success(ctx *fiber.Ctx, data any) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Data":      data,
		"RequestID": ctx.Get(constant.HeaderRequestId),
	})
}

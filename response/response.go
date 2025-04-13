package response

import (
	"github.com/FiberApps/common-library/constants/header"
	"github.com/FiberApps/common-library/logger"
	"github.com/gofiber/fiber/v2"
)

// XXX - Error with Message
func SendError(ctx *fiber.Ctx, err *fiber.Error) error {
	return ctx.Status(err.Code).JSON(fiber.Map{
		"Error":     err.Message,
		"RequestID": ctx.Get(header.RequestID),
	})
}

// XXX - Error with Data
func SendErrorWithData(ctx *fiber.Ctx, code int, err string) error {
	log := logger.GetLogger(ctx)
	log.Error().Int("status_code", code).Str("error", err).Msg("error response")

	return ctx.Status(code).JSON(fiber.Map{
		"Error":     err,
		"RequestID": ctx.Get(header.RequestID),
	})
}

// XXX - Error with Error
func SendErrorWithError(ctx *fiber.Ctx, code int, err error) error {
	log := logger.GetLogger(ctx)
	log.Err(err).Int("status_code", code).Msg("error response")

	return ctx.Status(code).JSON(fiber.Map{
		"Error":     err,
		"RequestID": ctx.Get(header.RequestID),
	})
}

// 2XX - Success
func Success(ctx *fiber.Ctx, data any) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Data":      data,
		"RequestID": ctx.Get(header.RequestID),
	})
}

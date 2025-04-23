package response

import (
	"github.com/Basil-Fox/basilfox-utils/constants/header"
	"github.com/Basil-Fox/basilfox-utils/logger"
	"github.com/gofiber/fiber/v2"
)

// XXX - Standard Network Error
func SendNetworkError(ctx *fiber.Ctx, err *fiber.Error) error {
	appLogger := logger.GetLogger(ctx)
	appLogger.Err(err).Int("status_code", err.Code).Msg("error response")

	return ctx.Status(err.Code).JSON(fiber.Map{
		"Error":     err.Message,
		"RequestID": ctx.Get(header.RequestID),
	})
}

// XXX - Generic Error with Code
func SendError(ctx *fiber.Ctx, code int, err error) error {
	appLogger := logger.GetLogger(ctx)
	appLogger.Err(err).Int("status_code", code).Msg("error response")

	return ctx.Status(code).JSON(fiber.Map{
		"Error":     err.Error(),
		"RequestID": ctx.Get(header.RequestID),
	})
}

// 500 - Internal Server Error
func SendInternalError(ctx *fiber.Ctx, err error) error {
	appLogger := logger.GetLogger(ctx)
	appLogger.Err(err).Int("status_code", fiber.StatusInternalServerError).Msg("internal server error")

	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"Error":     err,
		"RequestID": ctx.Get(header.RequestID),
	})
}

// XXX - Error with Message
func SendErrorMessage(ctx *fiber.Ctx, code int, message string) error {
	appLogger := logger.GetLogger(ctx)
	appLogger.Error().Int("status_code", code).Str("error", message).Msg("error response")

	return ctx.Status(code).JSON(fiber.Map{
		"Error":     message,
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

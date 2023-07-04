package middleware

import (
	"time"

	"github.com/FiberApps/core/constant"
	"github.com/FiberApps/core/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Logger(c *fiber.Ctx) error {
	start := time.Now()
	appLogger := logger.New(logger.Config{RequestId: c.Locals(constant.ContextRequestId).(string)})

	// Add logger to locals
	c.Locals(constant.ContextLogger, appLogger)

	// handle request
	err := c.Next()

	msg := "REQUEST"
	code := c.Response().StatusCode()

	logger := log.With().
		Str("namespace", c.Get(constant.HeaderNamespace)).
		Str("request-id", c.Get(constant.HeaderRequestId)).
		Int("status", code).
		Str("method", c.Method()).
		Str("path", c.Path()).
		Str("ip", c.IP()).
		Str("latency", time.Since(start).String()).
		Str("user-agent", c.Get(fiber.HeaderUserAgent)).
		Logger()

	switch {
	case code >= fiber.StatusBadRequest && code < fiber.StatusInternalServerError:
		logger.Warn().Msg(msg)
	case code >= fiber.StatusInternalServerError:
		logger.Error().Msg(msg)
	default:
		logger.Info().Msg(msg)
	}

	return err
}

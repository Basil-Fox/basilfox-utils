package middleware

import (
	"strings"
	"time"

	"github.com/FiberApps/common-library/constant"
	"github.com/FiberApps/common-library/logger"
	"github.com/gofiber/fiber/v2"
)

func RequestLogger(c *fiber.Ctx) error {
	start := time.Now()

	// Handle request and capture any errors
	err := c.Next()

	// Extract response details
	statusCode := c.Response().StatusCode()
	latency := time.Since(start)

	// Build structured log entry
	logEntry := logger.GetLogger().With().
		Str("request_id", c.Get(constant.HeaderRequestId)).
		Str("namespace", c.Get(constant.HeaderNamespace)).
		Int("status", statusCode).
		Str("method", c.Method()).
		Str("path", c.Path()).
		Str("ip", strings.Join(c.IPs(), ", ")).
		Dur("latency", latency).
		Str("user_agent", c.Get(fiber.HeaderUserAgent)).
		Logger()

	// Log based on status code severity
	switch {
	case statusCode >= fiber.StatusInternalServerError:
		logEntry.Error().Msg("server error")
	case statusCode >= fiber.StatusBadRequest:
		logEntry.Warn().Msg("client error")
	default:
		logEntry.Info().Msg("success")
	}

	return err
}

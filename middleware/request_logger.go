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

	var clientIP string
	var country string = "Unknown"

	// Check if the request is coming through Cloudflare, get IP from CF-Connecting-IP header
	if cfIP := c.Get("CF-Connecting-IP"); cfIP != "" {
		clientIP = cfIP

		// If Cloudflare is in use, get the country from the CF-IPCountry header
		if cfCountry := c.Get("CF-IPCountry"); cfCountry != "" {
			country = cfCountry
		}
	} else {
		// If Cloudflare header is not found, check X-Forwarded-For
		if xff := c.Get(fiber.HeaderXForwardedFor); xff != "" {
			// The X-Forwarded-For header may contain a comma-separated list of IPs.
			// The first IP is the real client IP.
			clientIP = strings.Split(xff, ",")[0]
		} else {
			// Direct access or internal request, fallback to c.IP()
			clientIP = c.IP()
		}
	}

	// Build structured log entry
	logEntry := logger.GetLogger().With().
		Str("request_id", c.Get(constant.HeaderRequestId)).
		Str("namespace", c.Get(constant.HeaderNamespace)).
		Int("status", statusCode).
		Str("method", c.Method()).
		Str("path", c.Path()).
		Str("ip", clientIP).
		Str("country", country).
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

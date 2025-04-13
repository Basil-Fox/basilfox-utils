package middleware

import (
	"strings"
	"time"

	"github.com/FiberApps/common-library/constants/header"
	"github.com/FiberApps/common-library/logger"
	"github.com/gofiber/fiber/v2"
)

func RequestLogger(ctx *fiber.Ctx) error {
	start := time.Now()

	// Handle request and capture any errors
	err := ctx.Next()

	// Extract response details
	statusCode := ctx.Response().StatusCode()
	latency := time.Since(start)

	var clientIP string
	var country string = "Unknown"

	// Check if the request is coming through Cloudflare, get IP from CF-Connecting-IP header
	if cfIP := ctx.Get("CF-Connecting-IP"); cfIP != "" {
		clientIP = cfIP

		// If Cloudflare is in use, get the country from the CF-IPCountry header
		if cfCountry := ctx.Get("CF-IPCountry"); cfCountry != "" {
			country = cfCountry
		}
	} else {
		// If Cloudflare header is not found, check X-Forwarded-For
		if xff := ctx.Get(fiber.HeaderXForwardedFor); xff != "" {
			// The X-Forwarded-For header may contain a comma-separated list of IPs.
			// The first IP is the real client IP.
			clientIP = strings.Split(xff, ",")[0]
		} else {
			// Direct access or internal request, fallback to ctx.IP()
			clientIP = ctx.IP()
		}
	}

	// Build structured log entry
	logEntry := logger.GetLogger(ctx).With().
		Str("namespace", ctx.Get(header.Namespace)).
		Str("client_version", ctx.Get(header.AppVersion)).
		Str("client_os", ctx.Get(header.DeviceOS)).
		Int("status_code", statusCode).
		Str("method", ctx.Method()).
		Str("path", ctx.Path()).
		Str("ip", clientIP).
		Str("country", country).
		Dur("latency", latency).
		Str("user_agent", ctx.Get(fiber.HeaderUserAgent)).
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

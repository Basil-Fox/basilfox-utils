package logger

import (
	"os"
	"time"

	"github.com/FiberApps/common-library/constants/header"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

var log zerolog.Logger

func InitLogger(appName, format string) {
	switch format {
	case "json":
		// Configure the logger to write to stdout in JSON format
		log = zerolog.New(os.Stdout).With().Timestamp().Str("app", appName).Logger()
	default:
		// Configure the logger to write to stdout in human-readable format
		log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Str("app", appName).Logger()
	}
}

func GetLogger(ctx *fiber.Ctx) zerolog.Logger {
	if ctx != nil {
		return log.With().Str("request_id", ctx.Get(header.RequestID)).Logger()
	}
	return log
}

func GetLoggerWithRequestId(requestId string) zerolog.Logger {
	return log.With().Str("request_id", requestId).Logger()
}

func SetLogger(newLogger zerolog.Logger) {
	log = newLogger
}

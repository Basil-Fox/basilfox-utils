package logger

import (
	"os"
	"time"

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

func GetLogger() zerolog.Logger {
	return log
}

func SetLogger(newLogger zerolog.Logger) {
	log = newLogger
}

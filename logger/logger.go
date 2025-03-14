package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var log zerolog.Logger

func InitLogger(jsonFormat bool) {
	if jsonFormat {
		// Configure the logger to write to stdout in JSON format
		log = zerolog.New(os.Stdout).With().Timestamp().Logger()
	} else {
		// Configure the logger to write to stdout in human-readable format
		log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
	}
}

func GetLogger() zerolog.Logger {
	return log
}

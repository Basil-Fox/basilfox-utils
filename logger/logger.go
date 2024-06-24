package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct{}

type Config struct {
	RequestId string
}

var ConfigDefault = Config{
	RequestId: "",
}

func New(config ...Config) *Logger {
	// Set default config
	cfg := ConfigDefault

	// Override config if provided
	if len(config) > 0 {
		cfg = config[0]

		// Set default values
		if cfg.RequestId == "" {
			cfg.RequestId = ConfigDefault.RequestId
		}
	}

	// Customize the log format
	if cfg.RequestId != "" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Str("request-id", cfg.RequestId).Logger()
	} else {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	}
	return &Logger{}
}

func (logger *Logger) Printf(level zerolog.Level, format string, v ...interface{}) {
	log.WithLevel(level).Msgf(format, v...)
}

func (logger *Logger) Debug(format string, v ...interface{}) {
	logger.Printf(zerolog.DebugLevel, format, v...)
}

func (logger *Logger) Info(format string, v ...interface{}) {
	logger.Printf(zerolog.InfoLevel, format, v...)
}

func (logger *Logger) Warn(format string, v ...interface{}) {
	logger.Printf(zerolog.WarnLevel, format, v...)
}

func (logger *Logger) Error(format string, v ...interface{}) {
	logger.Printf(zerolog.ErrorLevel, format, v...)
}

func (logger *Logger) Fatal(format string, v ...interface{}) {
	logger.Printf(zerolog.FatalLevel, format, v...)
}

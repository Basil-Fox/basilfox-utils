package utils

import (
	"os"

	"github.com/FiberApps/core/logger"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	log := logger.New()
	// Do not load .env in production
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("SERVER_INIT_FAILED::Error loading .env file")
		}
	}
}

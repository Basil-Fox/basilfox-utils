package utils

import (
	"github.com/FiberApps/common-library/logger"
	"github.com/joho/godotenv"
)

func LoadEnv(env string) {
	log := logger.New()

	// Do not load .env in production
	if env != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Error("INIT:: Error loading .env file: %v", err)
		}
	}
}

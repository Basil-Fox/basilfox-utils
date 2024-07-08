package utils

import (
	"github.com/joho/godotenv"
)

func LoadEnv(env string) error {
	// Do not load .env in production
	if env != "production" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}

	return nil
}

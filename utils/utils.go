package utils

import (
	"fmt"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from a .env file
func LoadEnv(env string, filePath ...string) error {
	// Avoid loading .env in production by default
	if env != "production" {
		var err error

		// If a custom file path is provided, load it
		if len(filePath) > 0 {
			err = godotenv.Load(filePath[0])
		} else {
			err = godotenv.Load()
		}

		if err != nil {
			return fmt.Errorf("error loading .env file: %w", err)
		}
	}

	return nil
}

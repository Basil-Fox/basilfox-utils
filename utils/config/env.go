package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
	// load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file")
	}
	return os.Getenv(key)
}

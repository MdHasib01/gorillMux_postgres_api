package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from the .env file.
func LoadEnv() {
	err := godotenv.Load("../cmd/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetEnv retrieves the value of the environment variable key, with an optional default.
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

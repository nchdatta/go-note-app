package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading the .env file.")
	}
	return os.Getenv(key)
}

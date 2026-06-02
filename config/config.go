package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DBUrl   string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, reading configuration from environment variables")
	}

	return &Config{
		AppPort: os.Getenv("APP_PORT"),
		DBUrl:   os.Getenv("DB_URL"),
	}
}
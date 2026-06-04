package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DBUrl   string
}

func NewConfig(filepath string) (*Config, error) {
	err := godotenv.Load(filepath)
	if err != nil {
		return nil, err 
	}

	return &Config{
		AppPort: os.Getenv("APP_PORT"),
		DBUrl:   os.Getenv("DB_URL"),
	}, nil
}
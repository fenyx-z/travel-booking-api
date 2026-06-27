package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort           string
	DBUrl             string
	DBMaxOpenConns    int
	DBMaxIdleConns    int
	DBConnMaxLifetime time.Duration
	DBConnMaxIdleTime time.Duration
}

func NewConfig(filepath string) (*Config, error) {
	err := godotenv.Load(filepath)
	if err != nil {
		return nil, err
	}

	return &Config{
		AppPort:           os.Getenv("APP_PORT"),
		DBUrl:             os.Getenv("DB_URL"),
		DBMaxOpenConns:    getEnvInt("DB_MAX_OPEN_CONNS", 350),
		DBMaxIdleConns:    getEnvInt("DB_MAX_IDLE_CONNS", 350),
		DBConnMaxLifetime: getEnvDuration("DB_CONN_MAX_LIFETIME", 30*time.Minute),
		DBConnMaxIdleTime: getEnvDuration("DB_CONN_MAX_IDLE_TIME", 5*time.Minute),
	}, nil
}

func getEnvInt(key string, defaultVal int) int {
	valStr := os.Getenv(key)
	if valStr == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultVal
	}
	return val
}

func getEnvDuration(key string, defaultVal time.Duration) time.Duration {
	valStr := os.Getenv(key)
	if valStr == "" {
		return defaultVal
	}
	val, err := time.ParseDuration(valStr)
	if err != nil {
		return defaultVal
	}
	return val
}

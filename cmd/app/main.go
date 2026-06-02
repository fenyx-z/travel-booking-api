package main

import (
	"fmt"
	"travel-backend/config"
	"travel-backend/pkg/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// 1. Load Configuration
	cfg := config.LoadConfig()

	// 2. Initialize Database
	_ = database.InitPostgres(cfg.DBUrl) // Sementara assign ke blank identifier (_) sebelum di-pass ke repo

	// 3. Initialize Echo
	e := echo.New()

	// Middleware bawaan
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Test Route
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "pong",
		})
	})

	// 4. Start Server
	port := fmt.Sprintf(":%s", cfg.AppPort)
	if cfg.AppPort == "" {
		port = ":8080"
	}
	e.Logger.Fatal(e.Start(port))
}
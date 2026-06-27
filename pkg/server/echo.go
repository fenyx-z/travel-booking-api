package server

import (
	"log"

	"travel-backend/config"
	"travel-backend/pkg/route"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	*echo.Echo
	Logger echo.Logger
}

func NewServer(cfg *config.Config, publicRoutes []route.Route) *Server {
	e := echo.New()

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogMethod:   true,
		LogURI:      true,
		LogStatus:   true,
		LogLatency:  true,
		LogRemoteIP: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Printf("[ECHO] %15s | %3d | %10v | %-7s %s | %s\n",
				v.RemoteIP,
				v.Status,
				v.Latency,
				v.Method,
				v.URI,
				v.Error,
			)
			return nil
		},
	}))

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	api := e.Group("/api/v1")

	for _, route := range publicRoutes {
		api.Add(route.Method, route.Path, route.Handler)
	}

	return &Server{
		Echo:   e,
		Logger: e.Logger,
	}
}
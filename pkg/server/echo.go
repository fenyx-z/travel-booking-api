package server

import (
	"travel-backend/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	*echo.Echo
}

type RouteFunc func(e *echo.Echo)

func NewServer(cfg *config.Config, publicRoutes RouteFunc, privateRoutes RouteFunc) *Server {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	if publicRoutes != nil {
		publicRoutes(e)
	}
	if privateRoutes != nil {
		privateRoutes(e)
	}

	return &Server{e}
}
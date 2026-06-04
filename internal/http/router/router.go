package router

import (
	"travel-backend/internal/http/handler"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, bookingHandler *handler.BookingHandler) {
	api := e.Group("/api/v1")

	bookings := api.Group("/bookings")
	bookings.POST("", bookingHandler.Create)
}
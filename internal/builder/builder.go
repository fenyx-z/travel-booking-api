package builder

import (
	"travel-backend/config"
	"travel-backend/internal/http/handler"
	"travel-backend/internal/repository"
	"travel-backend/internal/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) func(e *echo.Echo) {
	return func(e *echo.Echo) {

		api := e.Group("/api/v1")
		api.GET("/ping", func(c echo.Context) error {
			return c.JSON(200, map[string]string{"message": "pong"})
		})
	}
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) func(e *echo.Echo) {
	return func(e *echo.Echo) {
		scheduleRepo := repository.NewScheduleRepository(db)
		bookingRepo := repository.NewBookingRepository(db)

		bookingService := service.NewBookingService(db, scheduleRepo, bookingRepo)

		bookingHandler := handler.NewBookingHandler(bookingService)

		api := e.Group("/api/v1")

		bookings := api.Group("/bookings")
		bookings.POST("", bookingHandler.Create)
	}
}
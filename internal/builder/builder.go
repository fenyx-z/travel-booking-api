package builder

import (
	"travel-backend/internal/http/handler"
	"travel-backend/internal/http/router"
	"travel-backend/internal/repository"
	"travel-backend/internal/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BuildApp(e *echo.Echo, db *gorm.DB) {
	scheduleRepo := repository.NewScheduleRepository(db)
	bookingRepo := repository.NewBookingRepository(db)

	bookingService := service.NewBookingService(db, scheduleRepo, bookingRepo)

	bookingHandler := handler.NewBookingHandler(bookingService)

	router.SetupRoutes(e, bookingHandler)
}
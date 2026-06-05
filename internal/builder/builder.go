package builder

import (
	"travel-backend/config"
	"travel-backend/internal/http/handler"
	"travel-backend/internal/http/router"
	"travel-backend/internal/repository"
	"travel-backend/internal/service"
	"travel-backend/pkg/route"

	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []route.Route {
	// Inisialisasi komponen untuk Booking
	scheduleRepo := repository.NewScheduleRepository(db)
	bookingRepo := repository.NewBookingRepository(db)
	
	bookingService := service.NewBookingService(db, scheduleRepo, bookingRepo)
	bookingHandler := handler.NewBookingHandler(bookingService)

	return router.PublicRoutes(bookingHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []route.Route {
	return router.PrivateRoutes()
}
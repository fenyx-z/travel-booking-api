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
	// Inisialisasi komponen untuk Pemesanan
	jadwalRepo := repository.NewJadwalRepository(db)
	pemesananRepo := repository.NewPemesananRepository(db)
	kursiRepo := repository.NewKursiRepository(db)
	
	pemesananService := service.NewPemesananService(db, jadwalRepo, pemesananRepo, kursiRepo)
	pemesananHandler := handler.NewPemesananHandler(pemesananService)

	seedHandler := handler.NewSeedHandler(db)
	return router.PublicRoutes(pemesananHandler, seedHandler)
}
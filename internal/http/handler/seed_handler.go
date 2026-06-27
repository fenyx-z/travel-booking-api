package handler

import (
	"fmt"
	"net/http"
	"time"
	"travel-backend/internal/entity"
	"travel-backend/pkg/response"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type SeedHandler struct {
	db *gorm.DB
}

func NewSeedHandler(db *gorm.DB) *SeedHandler {
	return &SeedHandler{db: db}
}

func (h *SeedHandler) ResetDatabase(c echo.Context) error {
	// Drop tables
	h.db.Migrator().DropTable(&entity.Pemesanan{}, &entity.Kursi{}, &entity.Jadwal{}, &entity.Rute{}, &entity.Pengguna{})

	// Auto Migrate
	err := h.db.AutoMigrate(&entity.Pengguna{}, &entity.Rute{}, &entity.Jadwal{}, &entity.Kursi{}, &entity.Pemesanan{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Gagal menjalankan migrasi database: "+err.Error()))
	}

	// Create dummy data
	var firstPengguna entity.Pengguna
	for i := 1; i <= 10; i++ {
		pengguna := entity.Pengguna{
			Nama:  fmt.Sprintf("Tester Simulator %d", i),
			Email: fmt.Sprintf("tester%d@simulator.local", i),
		}
		if err := h.db.Create(&pengguna).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Gagal membuat user dummy: "+err.Error()))
		}
		if i == 1 {
			firstPengguna = pengguna
		}
	}

	rute := entity.Rute{
		Asal:   "Semarang",
		Tujuan: "Pangkalan Bun",
		Jarak:  650,
	}
	if err := h.db.Create(&rute).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Gagal membuat rute dummy: "+err.Error()))
	}

	jadwal := entity.Jadwal{
		IDRute:             rute.ID,
		WaktuKeberangkatan: time.Now().Add(48 * time.Hour),
		WaktuTiba:          time.Now().Add(50 * time.Hour),
		Tarif:              350000.00,
		TotalKursi:         50,
		KetersediaanKursi:  50,
	}
	if err := h.db.Create(&jadwal).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Gagal membuat jadwal dummy: "+err.Error()))
	}

	// Create physical seats (1 to 50) for this schedule
	var kursiList []entity.Kursi
	for i := 1; i <= 50; i++ {
		kursiList = append(kursiList, entity.Kursi{
			JadwalID:   jadwal.ID,
			NomorKursi: i,
			Status:     "available",
			Versi:      1,
		})
	}
	if err := h.db.Create(&kursiList).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Gagal membuat kursi dummy: "+err.Error()))
	}

	data := map[string]interface{}{
		"user_id":     firstPengguna.ID,
		"route_id":    rute.ID,
		"schedule_id": jadwal.ID,
		"seats":       jadwal.KetersediaanKursi,
	}

	return c.JSON(http.StatusOK, response.SuccessResponse("Database berhasil di-reset dan data dummy disuntikkan", data))
}

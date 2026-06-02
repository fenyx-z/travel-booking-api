package service

import (
	"errors"
	"fmt"
	"time"

	"travel-backend/internal/entity"
	"travel-backend/internal/http/dto"
	"travel-backend/internal/repository"
	"travel-backend/pkg/utils"

	"gorm.io/gorm"
)

type BookingService interface {
	CreateBooking(req dto.CreateBookingRequest) (*entity.Booking, error)
}

type bookingService struct {
	db           *gorm.DB // Dibutuhkan untuk memulai transaksi
	scheduleRepo repository.ScheduleRepository
	bookingRepo  repository.BookingRepository
}

func NewBookingService(db *gorm.DB, sRepo repository.ScheduleRepository, bRepo repository.BookingRepository) BookingService {
	return &bookingService{
		db:           db,
		scheduleRepo: sRepo,
		bookingRepo:  bRepo,
	}
}

func (s *bookingService) CreateBooking(req dto.CreateBookingRequest) (*entity.Booking, error) {
	// 1. Validasi input dasar
	if req.TotalSeats < 1 {
		return nil, errors.New("minimal pemesanan 1 kursi")
	}

	// 2. Mulai Database Transaction
	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	
	// Defer function untuk memastikan Rollback jika terjadi panic atau error di tengah jalan
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 3. Ambil data jadwal (termasuk nilai Version saat ini)
	schedule, err := s.scheduleRepo.FindByID(req.ScheduleID)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("jadwal tidak ditemukan")
	}

	// 4. Cek ketersediaan kursi
	if schedule.AvailableSeats < req.TotalSeats {
		tx.Rollback()
		return nil, errors.New("kursi tidak mencukupi")
	}

	// 5. Eksekusi Pengurangan Kursi dengan OCC
	err = s.scheduleRepo.DecreaseSeatWithOCC(tx, schedule.ID, req.TotalSeats, schedule.Version)
	if err != nil {
		tx.Rollback()
		// Jika err mendeteksi race condition, simulator/sistem bisa diatur 
		// untuk me-return error ke user atau mencoba retry logic (for loop).
		return nil, err 
	}

	// 6. Buat entitas Booking
	totalAmount := schedule.Price * float64(req.TotalSeats)
	bookingCode := fmt.Sprintf("TVL-%s-%s", time.Now().Format("060102"), utils.GenerateRandomString(5)) // Butuh utils helper

	newBooking := entity.Booking{
		UserID:      req.UserID,
		ScheduleID:  req.ScheduleID,
		BookingCode: bookingCode,
		TotalAmount: totalAmount,
		Status:      "pending",
	}

	// 7. Simpan data Booking
	err = s.bookingRepo.CreateBooking(tx, &newBooking)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("gagal membuat data pemesanan")
	}

	// 8. Commit transaksi jika semua berhasil
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &newBooking, nil
}
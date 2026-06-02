package repository

import (
	"travel-backend/internal/entity"

	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBooking(tx *gorm.DB, booking *entity.Booking) error
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db}
}

func (r *bookingRepository) CreateBooking(tx *gorm.DB, booking *entity.Booking) error {
	err := tx.Create(booking).Error
	if err != nil {
		return err
	}
	return nil
}
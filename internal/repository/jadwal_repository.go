package repository

import (
	"travel-backend/internal/entity"

	"gorm.io/gorm"
)

type JadwalRepository interface {
	FindByID(id uint) (*entity.Jadwal, error)
	DecreaseSeat(tx *gorm.DB, jadwalID uint, requestedSeats int) error
}

type jadwalRepository struct {
	db *gorm.DB
}

func NewJadwalRepository(db *gorm.DB) JadwalRepository {
	return &jadwalRepository{db}
}

func (r *jadwalRepository) FindByID(id uint) (*entity.Jadwal, error) {
	var jadwal entity.Jadwal
	err := r.db.First(&jadwal, id).Error
	if err != nil {
		return nil, err
	}
	return &jadwal, nil
}

func (r *jadwalRepository) DecreaseSeat(tx *gorm.DB, jadwalID uint, requestedSeats int) error {
	return tx.Model(&entity.Jadwal{}).
		Where("id = ?", jadwalID).
		Update("ketersediaan_kursi", gorm.Expr("ketersediaan_kursi - ?", requestedSeats)).Error
}

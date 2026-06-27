package repository

import (
	"travel-backend/internal/entity"

	"gorm.io/gorm"
)

type PemesananRepository interface {
	CreatePemesanan(tx *gorm.DB, pemesanan *entity.Pemesanan) error
}

type pemesananRepository struct {
	db *gorm.DB
}

func NewPemesananRepository(db *gorm.DB) PemesananRepository {
	return &pemesananRepository{db}
}

func (r *pemesananRepository) CreatePemesanan(tx *gorm.DB, pemesanan *entity.Pemesanan) error {
	err := tx.Create(pemesanan).Error
	if err != nil {
		return err
	}
	return nil
}

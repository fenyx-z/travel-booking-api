package repository

import (
	"errors"
	"travel-backend/internal/entity"

	"gorm.io/gorm"
)

type KursiRepository interface {
	FindByJadwalAndKursi(tx *gorm.DB, jadwalID uint, nomorKursi int) (*entity.Kursi, error)
	FindByJadwalAndKursiWithLock(tx *gorm.DB, jadwalID uint, nomorKursi int) (*entity.Kursi, error)
	UpdateWithOCC(tx *gorm.DB, kursi *entity.Kursi) error
	UpdateNoControl(tx *gorm.DB, kursi *entity.Kursi) error
}

type kursiRepository struct {
	db *gorm.DB
}

func NewKursiRepository(db *gorm.DB) KursiRepository {
	return &kursiRepository{db}
}

func (r *kursiRepository) FindByJadwalAndKursi(tx *gorm.DB, jadwalID uint, nomorKursi int) (*entity.Kursi, error) {
	var kursi entity.Kursi
	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.Where("jadwal_id = ? AND nomor_kursi = ?", jadwalID, nomorKursi).First(&kursi).Error
	if err != nil {
		return nil, err
	}
	return &kursi, nil
}

func (r *kursiRepository) FindByJadwalAndKursiWithLock(tx *gorm.DB, jadwalID uint, nomorKursi int) (*entity.Kursi, error) {
	var kursi entity.Kursi
	db := r.db
	if tx != nil {
		db = tx
	}
	err := db.Raw("SELECT * FROM kursi WHERE jadwal_id = ? AND nomor_kursi = ? FOR UPDATE", jadwalID, nomorKursi).Scan(&kursi).Error
	if err != nil {
		return nil, err
	}
	if kursi.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &kursi, nil
}

func (r *kursiRepository) UpdateWithOCC(tx *gorm.DB, kursi *entity.Kursi) error {
	db := r.db
	if tx != nil {
		db = tx
	}
	result := db.Model(&entity.Kursi{}).
	Where("id = ? AND versi = ?", kursi.ID, kursi.Versi).
	Updates(map[string]interface{}{
		"status": "booked",
		"versi":  gorm.Expr("versi + 1"),
	})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("race condition detected: seat version has been modified")
	}

	return nil
}

func (r *kursiRepository) UpdateNoControl(tx *gorm.DB, kursi *entity.Kursi) error {
	db := r.db
	if tx != nil {
		db = tx
	}
	return db.Model(kursi).Update("status", "booked").Error
}

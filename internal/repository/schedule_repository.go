package repository

import (
	"errors"
	"travel-backend/internal/entity"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	FindByID(id uint) (*entity.Schedule, error)
	DecreaseSeatWithOCC(tx *gorm.DB, scheduleID uint, requestedSeats int, currentVersion int) error
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db}
}

func (r *scheduleRepository) FindByID(id uint) (*entity.Schedule, error) {
	var schedule entity.Schedule
	err := r.db.First(&schedule, id).Error
	if err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *scheduleRepository) DecreaseSeatWithOCC(tx *gorm.DB, scheduleID uint, requestedSeats int, currentVersion int) error {
	result := tx.Model(&entity.Schedule{}).
		Where("id = ? AND version = ?", scheduleID, currentVersion).
		Updates(map[string]interface{}{
			"available_seats": gorm.Expr("available_seats - ?", requestedSeats),
			"version":         currentVersion + 1,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("race condition detected: data has been modified")
	}

	return nil
}
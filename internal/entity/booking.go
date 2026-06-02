package entity

import "time"

type Booking struct {
	ID            uint    `gorm:"primaryKey"`
	UserID        uint    `gorm:"not null"`
	ScheduleID    uint    `gorm:"not null"`
	BookingCode   string  `gorm:"unique;not null"`
	TotalAmount   float64 `gorm:"type:numeric(12,2);not null"`
	Status        string  `gorm:"default:'pending'"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
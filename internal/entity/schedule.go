package entity

import "time"

type Schedule struct {
	ID             uint      `gorm:"primaryKey"`
	RouteID        uint      `gorm:"not null"`
	DepartureTime  time.Time `gorm:"not null"`
	ArrivalTime    time.Time `gorm:"not null"`
	Price          float64   `gorm:"type:numeric(12,2);not null"`
	TotalSeats     int       `gorm:"not null"`
	AvailableSeats int       `gorm:"not null"`
	Version        int       `gorm:"default:1;not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
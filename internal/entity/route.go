package entity

import "time"

type Route struct {
	ID          uint   `gorm:"primaryKey"`
	Origin      string `gorm:"not null"`
	Destination string `gorm:"not null"`
	DistanceKm  int    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
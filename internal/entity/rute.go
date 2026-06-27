package entity

import "time"

type Rute struct {
	ID        uint      `gorm:"primaryKey;column:id"`
	Asal      string    `gorm:"not null;column:asal"`
	Tujuan    string    `gorm:"not null;column:tujuan"`
	Jarak     int       `gorm:"not null;column:jarak"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Rute) TableName() string {
	return "rute"
}

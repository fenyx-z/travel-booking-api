package entity

import "time"

type Pengguna struct {
	ID        uint      `gorm:"primaryKey;column:id"`
	Nama      string    `gorm:"not null;column:nama"`
	Email     string    `gorm:"unique;not null;column:email"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Pengguna) TableName() string {
	return "pengguna"
}

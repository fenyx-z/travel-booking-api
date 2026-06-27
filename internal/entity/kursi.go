package entity

import "time"

type Kursi struct {
	ID         uint      `gorm:"primaryKey;column:id"`
	JadwalID   uint      `gorm:"not null;column:jadwal_id"`
	Jadwal     Jadwal    `gorm:"foreignKey:JadwalID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	NomorKursi int       `gorm:"not null;column:nomor_kursi"`
	Status     string    `gorm:"type:varchar(20);default:'available';not null;column:status"`
	Versi      int       `gorm:"default:1;not null;column:versi"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (Kursi) TableName() string {
	return "kursi"
}

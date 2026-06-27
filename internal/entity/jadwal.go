package entity

import "time"

type Jadwal struct {
	ID                 uint      `gorm:"primaryKey;column:id"`
	IDRute             uint      `gorm:"not null;column:id_rute"`
	Rute               Rute      `gorm:"foreignKey:IDRute;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`
	WaktuKeberangkatan time.Time `gorm:"not null;column:waktu_keberangkatan"`
	WaktuTiba          time.Time `gorm:"not null;column:waktu_tiba"`
	Tarif              float64   `gorm:"type:numeric(12,2);not null;column:tarif"`
	TotalKursi         int       `gorm:"not null;column:total_kursi"`
	KetersediaanKursi  int       `gorm:"not null;column:ketersediaan_kursi"`
	CreatedAt          time.Time `gorm:"column:created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at"`
}

func (Jadwal) TableName() string {
	return "jadwal"
}

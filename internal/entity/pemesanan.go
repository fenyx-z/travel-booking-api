package entity

import "time"

type Pemesanan struct {
	ID          uint      `gorm:"primaryKey;column:id"`
	IDPengguna  uint      `gorm:"not null;column:id_pengguna"`
	Pengguna    Pengguna  `gorm:"foreignKey:IDPengguna;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`
	IDJadwal    uint      `gorm:"not null;column:id_jadwal"`
	Jadwal      Jadwal    `gorm:"foreignKey:IDJadwal;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;" json:"-"`
	KodePesanan string    `gorm:"-"`
	NomorKursi  int       `gorm:"not null;column:nomor_kursi"`
	TotalBiaya  float64   `gorm:"type:numeric(12,2);not null;column:total_biaya"`
	Status      string    `gorm:"default:'pending';column:status"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

func (Pemesanan) TableName() string {
	return "pemesanan"
}

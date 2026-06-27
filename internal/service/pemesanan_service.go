package service

import (
	"errors"
	"fmt"
	"time"

	"travel-backend/internal/entity"
	"travel-backend/internal/http/dto"
	"travel-backend/internal/repository"

	"gorm.io/gorm"
)

type PemesananService interface {
	CreatePemesananOCC(req dto.CreatePemesananRequest) (*entity.Pemesanan, error)
	CreatePemesananNoControl(req dto.CreatePemesananRequest) (*entity.Pemesanan, error)
	CreatePemesananPCC(req dto.CreatePemesananRequest) (*entity.Pemesanan, error)
}

type pemesananService struct {
	db            *gorm.DB
	jadwalRepo    repository.JadwalRepository
	pemesananRepo repository.PemesananRepository
	kursiRepo     repository.KursiRepository
}

func NewPemesananService(
	db *gorm.DB,
	jRepo repository.JadwalRepository,
	pRepo repository.PemesananRepository,
	kRepo repository.KursiRepository,
) PemesananService {
	return &pemesananService{
		db:            db,
		jadwalRepo:    jRepo,
		pemesananRepo: pRepo,
		kursiRepo:     kRepo,
	}
}

func (s *pemesananService) CreatePemesananOCC(req dto.CreatePemesananRequest) (*entity.Pemesanan, error) {
	if req.TotalSeats < 1 {
		return nil, errors.New("minimal pemesanan 1 kursi")
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	jadwal, err := s.jadwalRepo.FindByID(req.ScheduleID)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("jadwal tidak ditemukan")
	}

	if jadwal.KetersediaanKursi < req.TotalSeats {
		tx.Rollback()
		return nil, errors.New("kursi tidak mencukupi")
	}

	kursi, err := s.kursiRepo.FindByJadwalAndKursi(tx, req.ScheduleID, req.SeatNumber)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("kursi tidak ditemukan")
	}

	if kursi.Status != "available" {	
		tx.Rollback()
		return nil, errors.New("kursi sudah dipesan")
	}

	err = s.kursiRepo.UpdateWithOCC(tx, kursi)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = s.jadwalRepo.DecreaseSeat(tx, jadwal.ID, req.TotalSeats)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	totalBiaya := jadwal.Tarif * float64(req.TotalSeats)

	newPemesanan := entity.Pemesanan{
		IDPengguna:  req.UserID,
		IDJadwal:    req.ScheduleID,
		NomorKursi:  req.SeatNumber,
		TotalBiaya:  totalBiaya,
		Status:      "pending",
	}

	err = s.pemesananRepo.CreatePemesanan(tx, &newPemesanan)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("gagal membuat data pemesanan")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	newPemesanan.KodePesanan = fmt.Sprintf("TVL-%d", newPemesanan.ID)

	return &newPemesanan, nil
}

func (s *pemesananService) CreatePemesananNoControl(req dto.CreatePemesananRequest) (*entity.Pemesanan, error) {
	if req.TotalSeats < 1 {
		return nil, errors.New("minimal pemesanan 1 kursi")
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	jadwal, err := s.jadwalRepo.FindByID(req.ScheduleID)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("jadwal tidak ditemukan")
	}

	if jadwal.KetersediaanKursi < req.TotalSeats {
		tx.Rollback()
		return nil, errors.New("kursi tidak mencukupi")
	}

	kursi, err := s.kursiRepo.FindByJadwalAndKursi(tx, req.ScheduleID, req.SeatNumber)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("kursi tidak ditemukan")
	}

	if kursi.Status != "available" {
		tx.Rollback()
		return nil, errors.New("kursi sudah dipesan")
	}

	time.Sleep(50 * time.Millisecond)

	err = s.kursiRepo.UpdateNoControl(tx, kursi)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = s.jadwalRepo.DecreaseSeat(tx, jadwal.ID, req.TotalSeats)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	totalBiaya := jadwal.Tarif * float64(req.TotalSeats)

	newPemesanan := entity.Pemesanan{
		IDPengguna:  req.UserID,
		IDJadwal:    req.ScheduleID,
		NomorKursi:  req.SeatNumber,
		TotalBiaya:  totalBiaya,
		Status:      "pending",
	}

	err = s.pemesananRepo.CreatePemesanan(tx, &newPemesanan)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("gagal membuat data pemesanan")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	newPemesanan.KodePesanan = fmt.Sprintf("TVL-%d", newPemesanan.ID)

	return &newPemesanan, nil
}

func (s *pemesananService) CreatePemesananPCC(req dto.CreatePemesananRequest) (*entity.Pemesanan, error) {
	if req.TotalSeats < 1 {
		return nil, errors.New("minimal pemesanan 1 kursi")
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	jadwal, err := s.jadwalRepo.FindByID(req.ScheduleID)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("jadwal tidak ditemukan")
	}

	if jadwal.KetersediaanKursi < req.TotalSeats {
		tx.Rollback()
		return nil, errors.New("kursi tidak mencukupi")
	}

	kursi, err := s.kursiRepo.FindByJadwalAndKursiWithLock(tx, req.ScheduleID, req.SeatNumber)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("kursi sudah dipesan atau tidak ditemukan")
	}

	if kursi.Status != "available" {
		tx.Rollback()
		return nil, errors.New("kursi sudah dipesan")
	}

	err = s.kursiRepo.UpdateNoControl(tx, kursi)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = s.jadwalRepo.DecreaseSeat(tx, jadwal.ID, req.TotalSeats)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	totalBiaya := jadwal.Tarif * float64(req.TotalSeats)

	newPemesanan := entity.Pemesanan{
		IDPengguna:  req.UserID,
		IDJadwal:    req.ScheduleID,
		NomorKursi:  req.SeatNumber,
		TotalBiaya:  totalBiaya,
		Status:      "pending",
	}

	err = s.pemesananRepo.CreatePemesanan(tx, &newPemesanan)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("gagal membuat data pemesanan")
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	newPemesanan.KodePesanan = fmt.Sprintf("TVL-%d", newPemesanan.ID)

	return &newPemesanan, nil
}

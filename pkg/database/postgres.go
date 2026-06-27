package database

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(dsn string, maxOpen, maxIdle int, maxLifetime, maxIdleTime time.Duration) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(maxOpen)
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetConnMaxLifetime(maxLifetime)
	sqlDB.SetConnMaxIdleTime(maxIdleTime)

	return db, nil
}
package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	Price       string `gorm:"not null"`
}

type Order struct {
	ID        uint   `gorm:"primaryKey"`
	ProductID uint   `gorm:"not null"`
	Email     string `gorm:"not null"`
}

func ConnectDatabase() (*gorm.DB, error) {
	dsn := "user=postgres password=123456 dbname=wgetcard host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Configure the connection pool
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(time.Minute * 1)

	return db, nil
}

// Migrate the schema
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Product{}, &Order{})
}

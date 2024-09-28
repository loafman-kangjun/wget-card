package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	Price       string `gorm:"not null"`
}

// Connect to the database
func ConnectDatabase() (*gorm.DB, error) {
	dsn := "user=postgres password=123456 dbname=wgetcard host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Migrate the schema
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}

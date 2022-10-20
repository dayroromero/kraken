package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ProductGorm is DTO used to map Product entity to database
type ProductGorm struct {
	ID          uint   `gorm:"primaryKey;column:id"`
	Name        string `gorm:"column:name"`
	Weight      uint   `gorm:"column:weight"`
	IsAvailable bool   `gorm:"column:is_available"`
}
type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(ProductGorm{})
	if err != nil {
		log.Fatal(err)
	}

	return Handler{db}
}

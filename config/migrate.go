package config

import (
	"log"

	"fiber-base-go/internal/model"

	"gorm.io/gorm"
)

func DBMigrate(conn *gorm.DB) error {
	conn.AutoMigrate(&model.Student{})
	log.Println("Migration has been processed")

	return nil
}

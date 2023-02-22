package config

import (
	"log"

	"fiber-base-go/internal/model"

	"gorm.io/gorm"
)

func DBMigrate(conn *gorm.DB) error {
	conn.AutoMigrate(&model.Student{})
	conn.AutoMigrate(&model.Contest{})
	log.Println("Migration has been processed")

	return nil
}

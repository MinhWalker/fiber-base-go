package config

import (
	"log"

	"fiber-base-go/domain"
	"gorm.io/gorm"
)

func DBMigrate(conn *gorm.DB) error {
	conn.AutoMigrate(&domain.Student{})
	log.Println("Migration has been processed")

	return nil
}

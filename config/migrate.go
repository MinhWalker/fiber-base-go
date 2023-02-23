package config

import (
	"log"

	"fiber-base-go/internal/model"

	"gorm.io/gorm"
)

func DBMigrate(conn *gorm.DB) error {
	conn.AutoMigrate(&model.Student{})
	conn.AutoMigrate(&model.Contest{})
	conn.AutoMigrate(&model.Room{})
	conn.AutoMigrate(&model.Exam{})
	conn.AutoMigrate(&model.StudentExam{})

	log.Println("Migration has been processed")

	return nil
}

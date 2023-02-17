package config

import (
	"log"

	"fiber-base-go/domain"
	"gorm.io/gorm"
)

func DBMigrate() (*gorm.DB, error) {
	conn, err := ConnectDb()
	if err != nil {
		return nil, err
	}

	sqlDB, _ := conn.DB()
	defer sqlDB.Close()

	conn.AutoMigrate(&domain.Student{})
	log.Println("Migration has been processed")

	return conn, nil
}

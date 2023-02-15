package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDb() (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
		return nil, err
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	return db, nil
}

//func (c *ConfigDB) LoadConfig(path string) error {
//	viper.AddConfigPath(path)
//	viper.SetConfigName("")
//	viper.SetConfigType("env")
//
//	viper.AutomaticEnv()
//
//	err := viper.ReadInConfig()
//	if err != nil {
//		return err
//	}
//
//	err = viper.Unmarshal(&config)
//	return err
//}

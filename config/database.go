package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var config = ConfigDB{}

// ConfigDB db seting
type ConfigDB struct {
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	Dbname   string `mapstructure:"DB_NAME"`
}

func ConnectDb() (*gorm.DB, error) {
	config.LoadConfig("../../.")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.Host,
		config.User,
		config.Password,
		config.Dbname,
		config.Port,
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

func (c *ConfigDB) LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&config)
	return err
}

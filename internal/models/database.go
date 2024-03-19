package models

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func loadConfigFromFile(filename string) (Config, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = yaml.Unmarshal(bytes, &config)
	return config, err
}

func NewMustFromFile(filename string) DBManager {
	config, _ := loadConfigFromFile(filename)
	return NewMustFromConfig(config)
}

func NewMustFromConfig(config Config) DBManager {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.Host,
		config.User,
		config.Pass,
		config.DbName,
		config.Port,
		config.SslMode,
		config.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Person{})
	return DBManager{
		db: db,
	}
}

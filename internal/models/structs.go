package models

import "gorm.io/gorm"

type Config struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Pass     string `yaml:"pass"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"db_name"`
	SslMode  string `yaml:"ssl_mode"`
	TimeZone string `yaml:"time_zone"`
}

type DBManager struct {
	db *gorm.DB
}

package config

import (
	//"context"
	"github.com/Keyur1991/go-shreeva/db/mysql"
	"os"

	"gorm.io/gorm"
)

var DB *gorm.DB

var Database string

func DbConnect() {
	config := &mysql.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}

	Database = config.Database

	DB, _ = mysql.Connect(config)
}

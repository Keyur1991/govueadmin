package config

import (
	//"context"
	"govueadmin/framework/db/mysql"
	"os"

	"gorm.io/gorm"
)

var DB *gorm.DB

var Database string

//var Ctx context.Context

func DbConnect() {
	config := &mysql.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}

	Database = config.Database

	//Conn, Ctx, _ = mongodb.Connect(config)
	DB, _ = mysql.Connect(config)

	// Run migrations
}

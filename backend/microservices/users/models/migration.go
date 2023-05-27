package models

import (
	"app/microservices/users/config"
)

func init() {
	config.DbConnect()
	config.DB.Migrator().AutoMigrate(&User{})
}

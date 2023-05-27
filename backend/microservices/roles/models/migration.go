package models

import (
	"app/microservices/roles/config"
)

func init() {
	config.DbConnect()
	config.DB.Migrator().AutoMigrate(&Role{})
}

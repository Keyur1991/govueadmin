package models

import (
	"govueadmin/microservices/roles/config"
)

func init() {
	config.DbConnect()
	config.DB.Migrator().AutoMigrate(&Role{})
}

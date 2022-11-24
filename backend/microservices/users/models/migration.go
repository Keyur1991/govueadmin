package models

import (
	"govueadmin/microservices/users/config"
)

func init() {
	config.DbConnect()
	config.DB.Migrator().AutoMigrate(&User{})
}

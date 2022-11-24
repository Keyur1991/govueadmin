package models

import (
	"govueadmin/microservices/auth/config"
)

func init() {
	config.DbConnect()
	config.DB.Migrator().AutoMigrate(&User{}, &AuthToken{})
}

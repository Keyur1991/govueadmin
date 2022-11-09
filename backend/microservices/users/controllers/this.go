package controllers

import (
	"govueadmin/microservices/users/models"
)

type This struct {
	user models.User
	role models.Role
}

func this() *This {
	return &This{
		user: models.User{},
		role: models.Role{},
	}
}

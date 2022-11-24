package controllers

import (
	"govueadmin/microservices/auth/models"
)

type This struct {
	user models.User
}

func this() *This {
	return &This{
		user: models.User{},
	}
}

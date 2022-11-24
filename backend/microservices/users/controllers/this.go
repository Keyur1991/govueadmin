package controllers

import (
	"govueadmin/microservices/users/models"
)

type This struct {
	user models.User
}

func this() *This {
	return &This{
		user: models.User{},
	}
}

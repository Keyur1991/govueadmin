package controllers

import (
	"app/microservices/roles/models"
)

type This struct {
	role models.Role
}

func this() *This {
	return &This{
		role: models.Role{},
	}
}

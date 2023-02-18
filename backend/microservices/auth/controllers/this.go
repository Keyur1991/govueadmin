package controllers

import (
	"govueadmin/microservices/auth/models"
	"govueadmin/microservices/auth/pb"
)

type GrpcServer struct {
	pb.UnimplementedAuthServiceServer
}

type This struct {
	user models.User
}

func this() *This {
	return &This{
		user: models.User{},
	}
}

// some sample comment

package controllers

import (
	"app/microservices/auth/models"
	"app/microservices/auth/pb"
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

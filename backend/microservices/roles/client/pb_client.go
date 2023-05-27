package client

import (
	"app/microservices/auth/pb"
	"fmt"
	"os"

	"google.golang.org/grpc"
)

type AuthServiceClient struct {
	AuthClient pb.AuthServiceClient
}

func InitAuthServiceClient() AuthServiceClient {
	cc, err := grpc.Dial(os.Getenv("AUTH_SVC_URL"), grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect with Auth Service GRPC Server")
	}

	c := AuthServiceClient{
		AuthClient: pb.NewAuthServiceClient(cc),
	}

	return c
}

func (svc *AuthServiceClient) CheckAuthenticate() {

}

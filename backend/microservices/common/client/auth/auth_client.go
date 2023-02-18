package auth

import (
	"fmt"
	pbauth "govueadmin/microservices/common/pb/auth"
	"os"

	"google.golang.org/grpc"
)

type AuthServiceClient struct {
	AuthClient pbauth.AuthServiceClient
}

func InitAuthServiceClient() AuthServiceClient {
	cc, err := grpc.Dial(os.Getenv("AUTH_SVC_URL"), grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect with Auth Service GRPC Server")
	}

	c := AuthServiceClient{
		AuthClient: pbauth.NewAuthServiceClient(cc),
	}

	return c
}

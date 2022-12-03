package routes

import (
	"fmt"
	"govueadmin/microservices/auth/pb"
	"os"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient() pb.AuthServiceClient {
	cc, err := grpc.Dial(os.Getenv("SVC_URL"), grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewAuthServiceClient(cc)
}

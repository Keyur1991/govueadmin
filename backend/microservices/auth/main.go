package main

import (
	"context"
	"fmt"
	"govueadmin/microservices/auth/controllers"
	"govueadmin/microservices/auth/pb"
	"govueadmin/microservices/auth/routes"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

func main() {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		PORT2 := os.Getenv("PORT2")

		lis, err := net.Listen("tcp", PORT2)

		if err != nil {
			log.Fatalln("Failed to listing:", err)
		}

		grpc := grpc.NewServer()

		pb.RegisterAuthServiceServer(grpc, &controllers.GrpcServer{})

		fmt.Println("Auth Grpc Listening on port " + PORT2)
		if err := grpc.Serve(lis); err != nil {
			log.Fatalln("Failed to serve:", err)
		}

	}()

	router := routes.Setup()

	PORT := os.Getenv("PORT")

	fmt.Println("Auth microservices Listening on port " + PORT)
	log.Fatal(http.ListenAndServe(PORT, h2c.NewHandler(router, &http2.Server{})))
	// srv := &http.Server{
	// 	Handler:      router,
	// 	Addr:         PORT,
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	//log.Fatal(srv.ListenAndServe())

}

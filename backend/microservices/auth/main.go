package main

import (
	"context"
	"fmt"
	"govueadmin/microservices/auth/pb"
	"govueadmin/microservices/auth/routes"
	"log"
	"net/http"
	"os"
	"time"

	"google.golang.org/grpc"
)

func main() {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	router := routes.Setup()

	PORT := os.Getenv("PORT")

	srv := &http.Server{
		Handler:      router,
		Addr:         PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	grpcSrv := grpc.NewServer()

	pb.RegisterAuthServiceServer(&grpcSrv)

	fmt.Println("Auth microservices Listening on port " + PORT)
	log.Fatal(srv.ListenAndServe())
}

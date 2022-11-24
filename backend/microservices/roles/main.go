package main

import (
	"context"
	"fmt"
	"govueadmin/microservices/roles/routes"
	"log"
	"net/http"
	"os"
	"time"
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

	fmt.Println("Roles microservices Listening on port " + PORT)
	log.Fatal(srv.ListenAndServe())
}

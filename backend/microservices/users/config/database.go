package config

import (
	"context"
	"govueadmin/framework/db/mongodb"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

var Conn *mongo.Client

var Database string

var Ctx context.Context

func DbConnect() {
	config := &mongodb.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}

	Database = config.Database

	Conn, Ctx, _ = mongodb.Connect(config)
}

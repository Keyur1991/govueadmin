package mongodb

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func DBUrl(config *Config) string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/",
		config.User,
		url.QueryEscape(config.Password),
		config.Host,
		config.Port,
	)
}

func Connect(config *Config) (*mongo.Client, context.Context, error) {
	dsn := DBUrl(config)
	fmt.Println(dsn)

	ctx := context.Background()

	clientOptions := options.Client().ApplyURI(DBUrl(config))
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal("Error: ", err)
	}

	return client, ctx, err
}

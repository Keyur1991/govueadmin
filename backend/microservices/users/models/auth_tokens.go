package models

import (
	"govueadmin/microservices/users/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthToken struct {
	UserId    string             `json:"user_id" bson:"user_id"`
	TokenId   string             `json:"token_id" bson:"token_id"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
}

func (authToken *AuthToken) collection() *mongo.Collection {
	if config.Conn == nil {
		config.DbConnect()
	}

	coll := config.Conn.Database(config.Database).Collection("auth_tokens")

	return coll
}

func (authToken *AuthToken) Create() error {
	_, err := authToken.collection().InsertOne(config.Ctx, *authToken)

	if err != nil {
		return err
	}

	return nil
}

func (authToken *AuthToken) Delete() error {
	_, err := authToken.collection().DeleteOne(config.Ctx, *authToken)

	if err != nil {
		return err
	}

	return nil
}

func (authToken *AuthToken) Find() *mongo.SingleResult {
	return authToken.collection().FindOne(config.Ctx, *authToken)
}

package models

import (
	"govueadmin/framework/request"
	"govueadmin/microservices/users/config"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Uid       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	FirstName string             `json:"first_name" bson:"first_name,omitempty" validate:"required"`
	LastName  string             `json:"last_name" bson:"last_name,omitempty" validate:"required"`
	Email     string             `json:"email" bson:"email,omitempty" validate:"required_without=Uid,email,UniqueEmail"`
	Password  string             `json:"-" bson:"password,omitempty" validate:"required_without=Uid"`
	IsDeleted int                `json:"is_deleted" bsob:"is_deleted,omitempty"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at,omitempty"`
}

func (user *User) collection() *mongo.Collection {
	if config.Conn == nil {
		config.DbConnect()
	}

	coll := config.Conn.Database(config.Database).Collection("users")

	return coll
}

func (user *User) UniqueEmail(fl validator.FieldLevel) bool {
	return CheckUserExistByEmail((*user).Email)
}

func (user *User) ValidationMessages(err validator.ValidationErrors) map[string][]string {
	Fields := map[string]string{
		"User.FirstName": "first name",
		"User.LastName":  "last name",
		"User.Email":     "email",
		"User.Password":  "password",
	}

	Messages := map[string]string{
		"User.FirstName.required":        "Please provide the first name",
		"User.LastName.required":         "Please provide the last name",
		"User.Email.required":            "Please provide the email",
		"User.Email.required_without":    "Please provide the email",
		"User.Email.email":               "Please provide valid email address",
		"User.Email.UniqueEmail":         "The account already exist with the email address",
		"User.Password.required_without": "Please provide the password.",
	}

	return request.ValidationMessages(&err, &Fields, &Messages)
}

func CreateUser(user *User) error {
	insertedDoc, err := user.collection().InsertOne(config.Ctx, *user)

	if err != nil {
		panic(err)
		return err
	}

	return user.Find(insertedDoc.InsertedID)
}

func (user *User) Find(docId interface{}) error {
	var doc bson.M
	err := user.collection().FindOne(
		config.Ctx,
		bson.D{{"_id", docId}},
	).Decode(&doc)

	if err != nil {
		return err
	}

	data, err := bson.Marshal(doc)

	if err != nil {
		return err
	}

	bson.Unmarshal(data, user)
	return nil
}

func CheckUserExistByEmail(email string) bool {
	user := &User{}

	opts := options.Count().SetMaxTime(2 * time.Second)
	count, _ := user.collection().CountDocuments(
		config.Ctx,
		bson.D{{"email", email}},
		opts,
	)

	return count == 0
}

func FindByEmail(user *User, email string) error {
	var doc bson.M

	err := user.collection().FindOne(
		config.Ctx,
		bson.D{{"email", email}},
	).Decode(&doc)

	if err != nil {
		return err
	}

	data, err := bson.Marshal(doc)

	if err != nil {
		return err
	}

	bson.Unmarshal(data, user)
	return nil
}

func (user *User) GetHaxUserId() string {
	return user.Uid.Hex()
}

func (user *User) GetObjectIdFromHex(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}

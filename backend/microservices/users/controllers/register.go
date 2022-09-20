package controllers

import (
	"fmt"
	"govueadmin/framework/response"
	"govueadmin/microservices/users/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Register(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}

	if formatRequestRegister(&w, r, user) == nil {
		return
	}

	now := primitive.NewDateTimeFromTime(time.Now())
	user.CreatedAt, user.UpdatedAt = now, now

	status := http.StatusOK
	var data interface{}

	if err := models.CreateUser(user); err != nil {
		status = http.StatusInternalServerError
		data = fmt.Sprintf("%s", err)
	} else {
		data = user
	}

	response.Json(&w, status, data)
}

// middleware too many attempts
// define user model
// validation of the request via model
// throw validation messages as error
// localization of the error messages
// save the user in the database

package controllers

import (
	"encoding/json"
	"govueadmin/microservices/users/models"
	"net/http"

	//"github.com/gorilla/mux"
	"govueadmin/framework/response"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func formatRequestRegister(w *http.ResponseWriter, r *http.Request, user *models.User) *models.User {
	jsonPayload := make(map[string]interface{})

	json.NewDecoder(r.Body).Decode(&jsonPayload)

	if val, ok := jsonPayload["FirstName"]; ok && val != nil {
		user.FirstName = jsonPayload["FirstName"].(string)
	}

	if val, ok := jsonPayload["LastName"]; ok && val != nil {
		user.LastName = jsonPayload["LastName"].(string)
	}

	if val, ok := jsonPayload["Email"]; ok && val != nil {
		user.Email = jsonPayload["Email"].(string)
	}

	if val, ok := jsonPayload["Password"]; ok && val != nil {
		hashed, err := bcrypt.GenerateFromPassword([]byte(jsonPayload["Password"].(string)), bcrypt.MinCost)
		if err != nil {
			response.Json(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return nil
		}
		user.Password = string(hashed)
	}

	validate := validator.New()
	validate.RegisterValidation("UniqueEmail", user.UniqueEmail)

	err := validate.Struct(user)

	var data interface{}

	if err != nil {
		data = user.ValidationMessages(err.(validator.ValidationErrors))
		response.Json(w, http.StatusUnprocessableEntity, &data)
		return nil
	}

	return user
}

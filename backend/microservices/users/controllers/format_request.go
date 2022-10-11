package controllers

import (
	"govueadmin/microservices/users/models"
	"net/http"

	//"github.com/gorilla/mux"

	"govueadmin/framework/request"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func formatRequestRegister(c *gin.Context, user *models.User) *models.User {
	req, err := request.Factory(c.Request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": http.StatusText(http.StatusInternalServerError),
		})
		return nil
	}

	req.Parse()

	user.FirstName = req.Get("FirstName")
	user.LastName = req.Get("LastName")
	user.Email = req.Get("Email")

	password := req.Get("Password")

	if password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": http.StatusText(http.StatusInternalServerError),
			})
			return nil
		}
		user.Password = string(hashed)
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("UniqueField", user.UniqueField)
	}

	err = binding.Validator.ValidateStruct(user)

	var data interface{}

	if err != nil {
		data = user.ValidationMessages(err.(validator.ValidationErrors))
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": &data,
		})
		return nil
	}

	return user
}

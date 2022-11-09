package controllers

import (

	//"github.com/gorilla/mux"

	"fmt"
	"govueadmin/framework/request"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func (this *This) formatRequestRegister(c *gin.Context) error {
	req, err := request.Factory(c.Request)

	if err != nil {
		return err
	}

	req.Parse()

	this.user.Uid = ""
	this.user.FirstName = req.Get("FirstName")
	this.user.LastName = req.Get("LastName")
	this.user.Email = req.Get("Email")

	fmt.Println(this.user)
	password := req.Get("Password")

	if password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			return err
		}
		this.user.Password = string(hashed)
	}

	return nil
}

func (this *This) validateRegistration() map[string][]string {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("UniqueField", this.user.UniqueField)
	}

	if err := binding.Validator.ValidateStruct(this.user); err != nil {
		return this.user.ValidationMessages(err.(validator.ValidationErrors))
	}

	return nil
}

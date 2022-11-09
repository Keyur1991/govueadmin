package controllers

import (
	"govueadmin/framework/response"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	USER_REGISTERED = "Registration successful"
)

func Register(c *gin.Context) {
	this := this()

	// format the request
	if err := this.formatRequestRegister(c); err != nil {
		response.InternalServerError(c, err)
		return
	}

	// validate the request
	if errors := this.validateRegistration(); errors != nil {
		response.UnprocessableEntity(c, errors)
		return
	}

	// add timestamps to the model
	now := time.Now()
	this.user.CreatedAt, this.user.UpdatedAt = now, now

	// insert user into database
	if err := this.user.Create(); err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c, USER_REGISTERED, this.user)
}

// middleware too many attempts
// define user model
// validation of the request via model
// throw validation messages as error
// localization of the error messages
// save the user in the database

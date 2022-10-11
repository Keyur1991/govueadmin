package controllers

import (
	"fmt"
	"govueadmin/microservices/users/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	user := &models.User{}

	if formatRequestRegister(c, user) == nil {
		return
	}

	now := time.Now()
	user.CreatedAt, user.UpdatedAt = now, now

	status := http.StatusOK
	var data interface{}

	if err := user.Create(); err != nil {
		status = http.StatusInternalServerError
		data = fmt.Sprintf("%s", err)
	} else {
		data = user
	}

	c.JSON(status, gin.H{
		"data": data,
	})
}

// middleware too many attempts
// define user model
// validation of the request via model
// throw validation messages as error
// localization of the error messages
// save the user in the database

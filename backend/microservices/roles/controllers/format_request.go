package controllers

import (
	"govueadmin/framework/request"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func (this *This) formatRequest(c *gin.Context) error {
	req, err := request.Factory(c.Request)

	if err != nil {
		return err
	}

	req.Parse()

	this.role.Name = req.Get("name")

	return nil
}

func (this *This) validate() map[string][]string {
	if err := binding.Validator.ValidateStruct(this.role); err != nil {
		return this.role.ValidationMessages(err.(validator.ValidationErrors))
	}

	return nil
}

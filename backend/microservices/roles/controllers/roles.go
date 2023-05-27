package controllers

import (
	"fmt"
	"github.com/Keyur1991/go-shreeva/response"
	"time"

	"github.com/gin-gonic/gin"
)

func (this *This) formatAndValidate(c *gin.Context) bool {

	if err := this.formatRequest(c); err != nil {
		response.InternalServerError(c, err)
		return false
	}

	if errors := this.validate(); errors != nil {
		response.UnprocessableEntity(c, errors)
		return false
	}

	return true
}

func CreateRoleHandler(c *gin.Context) {
	this := this()

	if !this.formatAndValidate(c) {
		return
	}

	now := time.Now()
	this.role.CreatedAt, this.role.UpdatedAt = now, now

	if err := this.role.Create(); err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c, "Role successfully created", this.role)
}

func UpdateRoleHandler(c *gin.Context) {
	this := this()

	if !this.formatAndValidate(c) {
		return
	}

	uid := c.Params.ByName("id")

	if err := this.role.Find(uid); err != nil {
		response.NotFound(c, err)
		return
	}

	this.role.UpdatedAt = time.Now()

	if err := this.role.Update(uid); err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c, "Role successfully created", this.role)
}

func DeleteRoleHandler(c *gin.Context) {
	this := this()

	uid := c.Params.ByName("id")

	if err := this.role.Delete(uid); err != nil {
		response.InternalServerError(c, err)
		return
	}

	response.Success(c, "Role successfully deleted", nil)
}

func ViewRoleHandler(c *gin.Context) {

}

func GetRolesHandler(c *gin.Context) {
	fmt.Println("")
}

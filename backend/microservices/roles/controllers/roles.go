package controllers

import (
	"fmt"
	"govueadmin/framework/response"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateRoleHandler(c *gin.Context) {
	this := this()
	if err := this.formatRequest(c); err != nil {
		response.InternalServerError(c, err)
		return
	}

	if errors := this.validate(); errors != nil {
		response.UnprocessableEntity(c, errors)
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

}

func DeleteRoleHandler(c *gin.Context) {

}

func ViewRoleHandler(c *gin.Context) {

}

func GetRolesHandler(c *gin.Context) {
	fmt.Println("")
}

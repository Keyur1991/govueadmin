package controllers

import (
	"fmt"
	"govueadmin/framework/response"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	this := this()

	if err := this.formatRequestRole(c); err != nil {
		response.InternalServerError(c, err)
		return
	}

	if errors := this.validateRole(); errors != nil {
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

func UpdateRole(c *gin.Context) {

}

func DeleteRole(c *gin.Context) {

}

func ViewRole(c *gin.Context) {

}

func GetRoles(c *gin.Context) {
	fmt.Println("")
}

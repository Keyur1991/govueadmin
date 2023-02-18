package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": http.StatusText(http.StatusInternalServerError),
		"error":   err,
	})
}

func UnprocessableEntity(c *gin.Context, errors map[string][]string) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"errors": errors,
	})
}

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    &data,
	})
}

func BadGateway(c *gin.Context, err error) {
	c.JSON(http.StatusBadGateway, gin.H{
		"message": http.StatusText(http.StatusBadGateway),
		"error":   err,
	})
}

func AbortWithStatusJSON(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": http.StatusText(http.StatusUnauthorized),
	})
}

func NotFound(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": http.StatusText(http.StatusNotFound),
		"error":   err,
	})
}

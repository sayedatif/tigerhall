package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u UserController) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
		"data":    "",
	})
}

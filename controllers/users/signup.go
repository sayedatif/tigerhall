package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u UserController) Signup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User signed up successfully",
		"data":    "",
	})
}

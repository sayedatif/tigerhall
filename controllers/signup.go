package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a ApiController) Signup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User signed up successfully",
		"data":    "",
	})
}

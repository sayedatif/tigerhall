package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a ApiController) Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
		"data":    "",
	})
}

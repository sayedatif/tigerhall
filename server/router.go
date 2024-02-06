package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sayedatif/tigerhall/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	routes := router.Group("/api")
	{
		apiController := new(controllers.ApiController)
		routes.GET("/login", apiController.Login)
		routes.GET("/signup", apiController.Signup)
	}
	return router
}

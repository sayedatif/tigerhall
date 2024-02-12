package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sayedatif/tigerhall/controllers/tigers"
	"github.com/sayedatif/tigerhall/controllers/users"
	"github.com/sayedatif/tigerhall/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	userRoute := router.Group("/users")
	{
		userRoute.POST("/login", users.Login)
		userRoute.POST("/signup", users.Signup)
	}

	tigerRoute := router.Group("/tigers")
	{
		tigerRoute.POST("", middleware.JWTAuthMiddleware(), tigers.CreateTiger)
		tigerRoute.GET("", tigers.GetTigers)
		tigerRoute.POST("/:tiger_id/sighting", middleware.JWTAuthMiddleware(), tigers.CreateTigerSighting)
		tigerRoute.GET("/:tiger_id/sighting", tigers.GetTigerSightings)
	}
	return router
}

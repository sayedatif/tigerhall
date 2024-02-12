package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sayedatif/tigerhall/controllers/tigers"
	"github.com/sayedatif/tigerhall/controllers/users"
	"github.com/sayedatif/tigerhall/db"
	"github.com/sayedatif/tigerhall/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	database := db.GetDB()
	userRoute := router.Group("/users")
	{
		userController := new(users.UserController)
		userController.DB = database
		userRoute.POST("/login", userController.Login)
		userRoute.POST("/signup", userController.Signup)
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

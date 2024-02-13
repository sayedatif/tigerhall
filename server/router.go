package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sayedatif/tigerhall/controllers/tigers"
	"github.com/sayedatif/tigerhall/controllers/users"
	"github.com/sayedatif/tigerhall/db"
	_ "github.com/sayedatif/tigerhall/docs"
	"github.com/sayedatif/tigerhall/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Endpoint healthy",
		})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
		tigerController := new(tigers.TigerController)
		tigerController.DB = database
		tigerRoute.POST("", middleware.JWTAuthMiddleware(), tigerController.CreateTiger)
		tigerRoute.GET("", tigerController.GetTigers)
		tigerRoute.POST("/:tiger_id/sighting", middleware.JWTAuthMiddleware(), tigerController.CreateTigerSighting)
		tigerRoute.GET("/:tiger_id/sighting", tigerController.GetTigerSightings)
	}
	return router
}

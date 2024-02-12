package test

import (
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sayedatif/tigerhall/config"
	"github.com/sayedatif/tigerhall/db"
	"github.com/sayedatif/tigerhall/server"
	"gorm.io/gorm"
)

var router *gin.Engine

var database *gorm.DB

func TestMain(m *testing.M) {
	gin.SetMode(gin.ReleaseMode)
	setup()

	result := m.Run()

	teardown()

	os.Exit(result)
}

func setup() {
	config.Init("../.env")
	db.Init()
	database = db.GetDB()
	router = server.NewRouter()
}

func teardown() {
	if err := database.Where("email = ?", "test@test.com").Delete(&db.User{}).Error; err != nil {
		log.Printf("Error deleting test user: %v", err)
	}
}

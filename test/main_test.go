package test

import (
	"log"
	"os"
	"testing"

	"github.com/sayedatif/tigerhall/config"
	"github.com/sayedatif/tigerhall/db"
	"github.com/sayedatif/tigerhall/utils"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	config.Init("../.env")
	db.Init()

	database := db.GetDB()

	testUser := db.User{FirstName: "Test", LastName: "User", Email: "test@test.com", Password: "secret"}
	hashedPassword, _ := utils.HashPassword(testUser.Password)
	testUser.Password = hashedPassword
	username := utils.GenerateUsername(testUser.FirstName, testUser.LastName)
	testUser.Username = username
	if err := database.Create(&testUser).Error; err != nil {
		log.Fatalf("Error creating test user: %v", err)
	}
}

func teardown() {
	database := db.GetDB()

	if err := database.Where("email = ?", "test@test.com").Delete(&db.User{}).Error; err != nil {
		log.Fatalf("Error deleting test user: %v", err)
	}

	if database != nil {
		sqlDB, _ := database.DB()
		sqlDB.Close()
	}
}

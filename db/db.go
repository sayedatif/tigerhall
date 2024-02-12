package db

import (
	"fmt"
	"log"

	"github.com/sayedatif/tigerhall/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	c := config.GetConfig()
	var err error
	dbUser := c.GetString("DB_USER")
	dbPassword := c.GetString("DB_PASSWORD")
	dbHost := c.GetString("DB_HOST")
	dbName := c.GetString("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error on connecting to DB")
	}

	db.AutoMigrate(&User{}, &Tiger{}, &UserTigerSighting{}, &Email{})
}

func GetDB() *gorm.DB {
	return db
}

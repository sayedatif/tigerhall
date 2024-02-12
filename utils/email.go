package utils

import (
	"github.com/sayedatif/tigerhall/db"
	"gorm.io/gorm"
)

type Email struct {
	To      string
	Subject string
	Body    string
}

func SendEmail(database *gorm.DB, data Email) {
	email := db.Email{To: data.To, Subject: data.Subject, Body: data.Body}
	if err := database.Create(&email).Error; err != nil {
		return
	}
}

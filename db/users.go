package db

import (
	"time"
)

type User struct {
	ID        int64  `gorm:"primaryKey;autoIncrement;not null"`
	FirstName string `gorm:"type:varchar(255);not null"`
	LastName  string `gorm:"type:varchar(255);not null"`
	Username  string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Tigers    []Tiger `gorm:"many2many:user_tiger_sightings"`
}

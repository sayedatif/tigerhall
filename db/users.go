package db

import (
	"time"
)

type User struct {
	ID        int64 `gorm:"primaryKey;autoIncrement;not null"`
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
